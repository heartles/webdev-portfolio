## Libraries used:

- Font Awesome
- Golang standard library (for html templating)

## Tutorials used:

- https://www.w3schools.com/howto/howto_js_topnav_responsive.asp

## Building:

The build process depends on Go. If you do not have it installed, grab it from your
package manager or the [official site](https://golang.org).

To build the static site files, execute:

`go run main.go`

This will build a static copy of the site to a subdirectory called `dist/`. To host
this, use whatever static site hosting you wish.
