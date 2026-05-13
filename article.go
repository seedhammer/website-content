package content

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"

	"seedhammer.com/website/content/article"
	"seedhammer.com/website/content/page"
)

type Article struct {
	Slug        string
	Title       string
	Author      string
	AuthorLink  string
	Description string
	Image       string
	Type        string
	Published   time.Time
	Content     template.HTML
}

var allArticles struct {
	bySlug map[string]Article
	byDate []Article
	tmpl   *template.Template
}

type Section struct {
	ID   string
	Name string
}

type TOC []struct {
	Slug     string
	Name     string
	Index    int
	Sections []Section
}

func init() {
	allArticles.tmpl = template.Must(template.ParseFS(page.FS, "icons.html", "base.html", "article.html")).Lookup("base")
	arts, err := loadArticles(article.FS)
	if err != nil {
		panic(err)
	}
	allArticles.byDate = arts
	allArticles.bySlug = make(map[string]Article)
	for _, a := range arts {
		allArticles.bySlug[a.Slug] = a
	}
}

func docsHandler(urlPath string, fs fs.FS, toc TOC) http.Handler {
	arts, err := loadArticles(fs)
	if err != nil {
		panic(err)
	}
	bySlug := make(map[string]Article)
	for _, a := range arts {
		bySlug[a.Slug] = a
	}
	for i := range toc {
		toc[i].Index = i + 1
	}
	tmpl, ok := Templates[path.Base(urlPath)]
	if !ok {
		panic("template not found")
	}
	return CachingHandler(LoggingFunc(func(w http.ResponseWriter, r *http.Request) error {
		slug := strings.TrimPrefix(r.URL.Path, "/"+urlPath+"/")
		md := slug
		switch md {
		case "":
			md = "index"
		case "index":
			return ErrNotFound
		}
		doc, ok := bySlug[md]
		if !ok {
			return ErrNotFound
		}
		args := struct {
			Base       string
			TOC        TOC
			ActiveSlug string
			Doc        Article
		}{
			Base:       "/" + urlPath,
			ActiveSlug: slug,
			TOC:        toc,
			Doc:        doc,
		}
		if err := RenderPageTemplate(w, r, tmpl, args); err != nil {
			return fmt.Errorf("docs: template failed: %v", err)
		}
		return nil
	}))
}

func renderArticleTOC(w http.ResponseWriter, r *http.Request) error {
	if err := RenderPageTemplate(w, r, Templates["articles"], allArticles.byDate); err != nil {
		return fmt.Errorf("article: template failed: %v", err)
	}
	return nil
}

func renderArticles(w http.ResponseWriter, r *http.Request) error {
	name := strings.TrimPrefix(r.URL.Path, "/article/")
	if name == "" {
		return renderArticleTOC(w, r)
	}
	art, ok := allArticles.bySlug[name]
	if !ok {
		return ErrNotFound
	}
	if err := RenderPageTemplate(w, r, allArticles.tmpl, art); err != nil {
		return fmt.Errorf("article: template failed: %v", err)
	}
	return nil
}

func loadArticles(f fs.FS) ([]Article, error) {
	var articles []Article
	err := fs.WalkDir(f, ".", func(path string, d fs.DirEntry, err error) error {
		ext := filepath.Ext(path)
		if ext != ".md" {
			return err
		}
		base := filepath.Base(path)
		name := strings.TrimSuffix(base, ext)
		md := goldmark.New(
			goldmark.WithExtensions(
				meta.Meta,
				extension.Table,
			),
			goldmark.WithParserOptions(parser.WithHeadingAttribute()),
			goldmark.WithRendererOptions(html.WithUnsafe()),
		)
		src, err := fs.ReadFile(f, path)
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		context := parser.NewContext()
		if err := md.Convert(src, &buf, parser.WithContext(context)); err != nil {
			return fmt.Errorf("article: %s: %w", path, err)
		}
		m := meta.Get(context)
		art := Article{
			Slug:    name,
			Content: template.HTML(buf.String()),
		}
		if d, ok := m["published"]; ok {
			published, err := time.ParseInLocation("2006-01-02", d.(string), time.UTC)
			if err != nil {
				return fmt.Errorf("article: %s: failed to parse `published` field: %w", path, err)
			}
			art.Published = published
		}
		if title, ok := m["title"]; ok {
			art.Title = title.(string)
		}
		if desc, ok := m["description"]; ok {
			art.Description = desc.(string)
		}
		if desc, ok := m["author"]; ok {
			art.Author = desc.(string)
		}
		if desc, ok := m["authorlink"]; ok {
			art.AuthorLink = desc.(string)
		}
		if img, ok := m["image"]; ok {
			art.Image = img.(string)
		}
		articles = append(articles, art)
		return err
	})
	sort.SliceStable(articles, func(i, j int) bool {
		return articles[i].Published.After(articles[j].Published)
	})
	return articles, err
}
