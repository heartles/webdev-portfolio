## Libraries used:

- Font Awesome
- Golang standard library (for html templating)
- Bootstrap (in contact form page)

## Tutorials used:

- https://www.w3schools.com/howto/howto_js_topnav_responsive.asp

## Building:

To build on x86_64 Linux, run `bin/build` from the root directory of the project.

This will build a static copy of the site to a subdirectory called `dist/`. To host
this, use whatever static site hosting you wish.

### On other systems

If you need a build for a different system, you will need to compile the build agent
yourself. To do so, make sure [Go](https://golang.org) is installed, then run

`go run build.go`
