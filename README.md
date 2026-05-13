## Public content of seedhammer.com


### Run local viewer

Command `viewer` runs a webserver at http://localhost:8080 with
the content.

```sh
$ go run ./cmd/viewer
```

The [Nix](https://nixos.org) helper script runs the viewer
and reloads it on file changes.

```sh
$ nix run .#

```
