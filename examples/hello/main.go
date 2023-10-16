package main

import (
	"net/http"

	"github.com/htmgolang/htmg"
)

// local component
func myHtmlHead(g htmg.HtmgCtx, title string) {
	g.Head()
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
	g.Head_()
}

// http handler
func helloHandler(w http.ResponseWriter, req *http.Request) {
	g := htmg.New(w) // pick any letter you like
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
