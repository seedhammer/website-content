package content

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/btcsuite/btcd/address/v2/bech32"
	"github.com/gorilla/feeds"
	"github.com/lpar/gzipped/v2"
	"seedhammer.com/website/content/about"
	"seedhammer.com/website/content/doc/manual"
	"seedhammer.com/website/content/getstarted"
	"seedhammer.com/website/content/page"
)

//go:embed static
var staticFiles embed.FS

var Templates map[string]*template.Template

type LoggingFunc func(http.ResponseWriter, *http.Request) error

var ErrNotFound = errors.New("resource not found")

var nostrID []byte

const (
	domain    = "seedhammer.com"
	nostrNPUB = "npub1z072s0nvldeva7a6qek3kj358f0h5gm640kkllkk52h0qrjtnw8q38wfm4"
)

func init() {
	_, decoded, err := bech32.Decode(nostrNPUB)
	if err != nil {
		panic(err)
	}
	dec2, err := bech32.ConvertBits(decoded, 5, 8, false)
	if err != nil {
		panic(err)
	}
	nostrID = dec2
}

func init() {
	p, err := LoadTemplates(page.FS)
	if err != nil {
		panic(err)
	}
	Templates = p
}

func Register(mux *http.ServeMux) {
	mux.Handle("GET /static/", CachingHandler(gzipped.FileServer(gzipped.FS(staticFiles))))
	mux.Handle("GET /favicon.ico", http.NotFoundHandler())
	mux.Handle("GET /article/", CachingHandler(LoggingFunc(renderArticles)))
	mux.Handle("GET /article/rss", CachingHandler(LoggingFunc(renderRssArticles)))
	mux.Handle("GET /tip", http.RedirectHandler("/get-started/engrave-plate", http.StatusFound))
	mux.Handle("GET /get-started/", docsHandler("get-started", getstarted.FS, getstartedTOC))
	mux.Handle("GET /doc/manual/", docsHandler("doc/manual", manual.FS, manualTOC))
	mux.Handle("GET /doc/", http.RedirectHandler("/doc/manual/", http.StatusFound))
	mux.Handle("GET /about/", docsHandler("about", about.FS, aboutTOC))

	mux.Handle("GET /website/content", goImportHandler(http.NotFoundHandler()))

	mux.Handle("GET /.well-known/nostr.json", http.HandlerFunc(nip05Handler))

	mux.Handle("GET /{page}", goImportHandler(pageHandler()))
	mux.Handle("GET /{$}", goImportHandler(pageHandler()))
}

func goImportHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("go-get") != "1" {
			h.ServeHTTP(w, r)
			return
		}
		module := r.URL.Path
		var repo string
		switch r.URL.Path {
		case "/":
			repo = "https://github.com/seedhammer/seedhammer"
		case "/website/content":
			repo = "https://github.com/seedhammer/website-content"
		default:
			http.NotFound(w, r)
			return
		}
		root := "seedhammer.com" + module
		fmt.Fprintf(w, `<html><head>
<meta name="go-import" content="%[1]s git %[2]s">
<meta name="go-source" content="%[1]s _ %[2]s/tree/main{/dir} %[2]s/tree/main{/dir}/{file}#L{line}">
</head></html>`, root, repo)
	})
}

func RenderPageTemplate(w io.Writer, r *http.Request, tmpl *template.Template, args any) error {
	if tmpl == nil {
		return ErrNotFound
	}
	u := r.URL
	u.Host = domain
	u.Scheme = "https"
	pargs := struct {
		PageURL *url.URL
		Nav     []Nav
		NavURL  string
		Content any
	}{
		PageURL: u,
		Nav:     topNav,
		Content: args,
	}
	for _, nav := range pargs.Nav {
		if strings.HasPrefix(r.URL.Path, nav.URL) {
			pargs.NavURL = nav.URL
			break
		}
	}
	return tmpl.Lookup("base").Execute(w, pargs)
}

func (f LoggingFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := f(w, r); err != nil {
		if errors.Is(err, ErrNotFound) {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		log.Printf("ERROR: %v: %v", r.URL, err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
}

func nip05Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, `{
  "names": {
    "_": "%x"
  }
}`, nostrID)
}

func renderRssArticles(w http.ResponseWriter, r *http.Request) error {
	w.Header().Add("Content-Type", "application/rss+xml")
	feed := &feeds.Feed{
		Title:       "Articles from SeedHammer",
		Description: "Latest news, updates and insights from SeedHammer Cold Storage.",
		Author:      &feeds.Author{Name: "", Email: ""},
		Link:        &feeds.Link{Href: "https://" + domain},
	}
	arts := allArticles.byDate
	const limit = 100
	if len(arts) > limit {
		arts = arts[:limit]
	}
	for _, a := range arts {
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       a.Title,
			Link:        &feeds.Link{Href: fmt.Sprintf("https://%s/article/%s", domain, a.Slug)},
			Author:      &feeds.Author{Name: a.Author, Email: a.AuthorLink},
			Description: a.Description,
			Created:     a.Published,
		})
	}
	return feed.WriteRss(w)
}

func LoadTemplates(dir fs.FS) (map[string]*template.Template, error) {
	pages := make(map[string]*template.Template)
	err := fs.WalkDir(dir, ".", func(path string, d fs.DirEntry, err error) error {
		ext := filepath.Ext(path)
		if ext != ".html" {
			return err
		}
		base := filepath.Base(path)
		if base == "base.html" {
			return err
		}
		tmpl, err := template.ParseFS(dir, "icons.html", "base.html", path)
		if err != nil {
			return err
		}
		name := strings.TrimSuffix(base, ext)
		pages[name] = tmpl
		return err
	})
	return pages, err
}

func pageHandler() http.Handler {
	return CachingHandler(LoggingFunc(func(w http.ResponseWriter, r *http.Request) error {
		tname := r.PathValue("page")
		if tname == "" {
			tname = "sh2"
		}
		tmpl, ok := Templates[tname]
		if !ok {
			return ErrNotFound
		}
		if err := RenderPageTemplate(w, r, tmpl, nil); err != nil {
			return fmt.Errorf("page: template failed: %v", err)
		}
		return nil
	}))
}

func CachingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "public, max-age=3600")
		h.ServeHTTP(w, r)
	})
}
