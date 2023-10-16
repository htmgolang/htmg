# htmg (HyperText Markup Golang)

Warning: this text may not be suitable for all audiences (no templating engine involved).

Using this library, aside from string concatenation, is the most straightforward way to output HTML content to any io.Writer.

## Features

- Pure Go
- Fast
- Dead simple
- Fully imperative (but can look declarative)
- Compatible with any control structure (for, if, switch)
- Works with any io.Writer
- Compatible with any editor
- Integrates well with htmx

## Usage

Acquire the library using `go get`:

```shell
go get github.com/htmgolang/htmg
```

Everything you need to know about htmg is demonstrated in this example:

```go
package main

import (
	"net/http"

	"github.com/htmgolang/htmg"
)

// local component (optional)
func myHtmlHead(g htmg.HtmgCtx, title string) {
	g.Head() // opening tag
	{
		g.Meta("charset='UTF-8'")
		g.Meta("name='viewport'", "content='width=device-width, initial-scale=1.0'")
		g.Title()
		{
			g.Text(title)
		}
		g.Title_()
		g.Style()
		{
			g.Text("html {padding:2rem; margin:auto; font-family:system-ui;}")
		}
		g.Style_()
	}
	g.Head_() // closing tag
}

// HTTP handler
func helloHandler(w http.ResponseWriter, req *http.Request) {
	g := htmg.New(w) // choose any letter you prefer
	g.Html("lang='en'")
	{
		myHtmlHead(g, "htmg demo")
		g.Body()
		{
			g.H1()
			{
				g.Textf("Request URL path: %s", req.URL.Path)
			}
			g.H1_()
			g.Ul()
			for _, rune := range req.URL.Path {
				g.Li()
				{
					g.Text(string(rune))
				}
				g.Li_()
			}
			g.Ul_()
		}
		g.Body_()
	}
	g.Html_()
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
}
```

## License

MIT