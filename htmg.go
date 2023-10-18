package htmg

import (
	"fmt"
	"io"
)

type Context struct {
	io.Writer
}

func New(w io.Writer) Context {
	return Context{w}
}

func (c *Context) write(str string) {
	c.Write([]byte(str))
}

func (c *Context) writeWithAttributes(str string, attributes []string) {
	c.write("<")
	c.write(str)
	for _, attribute := range attributes {
		c.write(" ")
		c.write(attribute)
	}
	c.write(">")
}

func (c *Context) write_(str string) {
	c.write("</")
	c.write(str)
	c.write(">")
}

func (c *Context) Text(str string) {
	c.write(str)
}

func (c *Context) Textf(format string, a ...any) {
	c.write(fmt.Sprintf(format, a...))
}

func (c *Context) Html(attributes ...string) {
	c.write("<!DOCTYPE html>")
	c.writeWithAttributes("html", attributes)
}

func (c *Context) Html_() {
	c.write_("html")
}
func (c *Context) A(attributes ...string) {
	c.writeWithAttributes("a", attributes)
}

func (c *Context) A_() {
	c.write_("a")
}

func (c *Context) Abbr(attributes ...string) {
	c.writeWithAttributes("abbr", attributes)
}

func (c *Context) Abbr_() {
	c.write_("abbr")
}

func (c *Context) Address(attributes ...string) {
	c.writeWithAttributes("address", attributes)
}

func (c *Context) Address_() {
	c.write_("address")
}

func (c *Context) Area(attributes ...string) {
	c.writeWithAttributes("area", attributes)
}

func (c *Context) Article(attributes ...string) {
	c.writeWithAttributes("article", attributes)
}

func (c *Context) Article_() {
	c.write_("article")
}

func (c *Context) Aside(attributes ...string) {
	c.writeWithAttributes("aside", attributes)
}

func (c *Context) Aside_() {
	c.write_("aside")
}

func (c *Context) Audio(attributes ...string) {
	c.writeWithAttributes("audio", attributes)
}

func (c *Context) Audio_() {
	c.write_("audio")
}

func (c *Context) B(attributes ...string) {
	c.writeWithAttributes("b", attributes)
}

func (c *Context) B_() {
	c.write_("b")
}

func (c *Context) Base(attributes ...string) {
	c.writeWithAttributes("base", attributes)
}

func (c *Context) Bdi(attributes ...string) {
	c.writeWithAttributes("bdi", attributes)
}

func (c *Context) Bdi_() {
	c.write_("bdi")
}

func (c *Context) Bdo(attributes ...string) {
	c.writeWithAttributes("bdo", attributes)
}

func (c *Context) Bdo_() {
	c.write_("bdo")
}

func (c *Context) Blockquote(attributes ...string) {
	c.writeWithAttributes("blockquote", attributes)
}

func (c *Context) Blockquote_() {
	c.write_("blockquote")
}

func (c *Context) Body(attributes ...string) {
	c.writeWithAttributes("body", attributes)
}

func (c *Context) Body_() {
	c.write_("body")
}

func (c *Context) Br(attributes ...string) {
	c.writeWithAttributes("br", attributes)
}

func (c *Context) Button(attributes ...string) {
	c.writeWithAttributes("button", attributes)
}

func (c *Context) Button_() {
	c.write_("button")
}

func (c *Context) Canvas(attributes ...string) {
	c.writeWithAttributes("canvas", attributes)
}

func (c *Context) Canvas_() {
	c.write_("canvas")
}

func (c *Context) Caption(attributes ...string) {
	c.writeWithAttributes("caption", attributes)
}

func (c *Context) Caption_() {
	c.write_("caption")
}

func (c *Context) Cite(attributes ...string) {
	c.writeWithAttributes("cite", attributes)
}

func (c *Context) Cite_() {
	c.write_("cite")
}

func (c *Context) Code(attributes ...string) {
	c.writeWithAttributes("code", attributes)
}

func (c *Context) Code_() {
	c.write_("code")
}

func (c *Context) Col(attributes ...string) {
	c.writeWithAttributes("col", attributes)
}

func (c *Context) Colgroup(attributes ...string) {
	c.writeWithAttributes("colgroup", attributes)
}

func (c *Context) Colgroup_() {
	c.write_("colgroup")
}

func (c *Context) Data(attributes ...string) {
	c.writeWithAttributes("data", attributes)
}

func (c *Context) Data_() {
	c.write_("data")
}

func (c *Context) Datalist(attributes ...string) {
	c.writeWithAttributes("datalist", attributes)
}

func (c *Context) Datalist_() {
	c.write_("datalist")
}

func (c *Context) Dd(attributes ...string) {
	c.writeWithAttributes("dd", attributes)
}

func (c *Context) Dd_() {
	c.write_("dd")
}

func (c *Context) Del(attributes ...string) {
	c.writeWithAttributes("del", attributes)
}

func (c *Context) Del_() {
	c.write_("del")
}

func (c *Context) Details(attributes ...string) {
	c.writeWithAttributes("details", attributes)
}

func (c *Context) Details_() {
	c.write_("details")
}

func (c *Context) Dfn(attributes ...string) {
	c.writeWithAttributes("dfn", attributes)
}

func (c *Context) Dfn_() {
	c.write_("dfn")
}

func (c *Context) Dialog(attributes ...string) {
	c.writeWithAttributes("dialog", attributes)
}

func (c *Context) Dialog_() {
	c.write_("dialog")
}

func (c *Context) Div(attributes ...string) {
	c.writeWithAttributes("div", attributes)
}

func (c *Context) Div_() {
	c.write_("div")
}

func (c *Context) Dl(attributes ...string) {
	c.writeWithAttributes("dl", attributes)
}

func (c *Context) Dl_() {
	c.write_("dl")
}

func (c *Context) Dt(attributes ...string) {
	c.writeWithAttributes("dt", attributes)
}

func (c *Context) Dt_() {
	c.write_("dt")
}

func (c *Context) Em(attributes ...string) {
	c.writeWithAttributes("em", attributes)
}

func (c *Context) Em_() {
	c.write_("em")
}

func (c *Context) Embed(attributes ...string) {
	c.writeWithAttributes("embed", attributes)
}

func (c *Context) Fieldset(attributes ...string) {
	c.writeWithAttributes("fieldset", attributes)
}

func (c *Context) Fieldset_() {
	c.write_("fieldset")
}

func (c *Context) Figcaption(attributes ...string) {
	c.writeWithAttributes("figcaption", attributes)
}

func (c *Context) Figcaption_() {
	c.write_("figcaption")
}

func (c *Context) Figure(attributes ...string) {
	c.writeWithAttributes("figure", attributes)
}

func (c *Context) Figure_() {
	c.write_("figure")
}

func (c *Context) Footer(attributes ...string) {
	c.writeWithAttributes("footer", attributes)
}

func (c *Context) Footer_() {
	c.write_("footer")
}

func (c *Context) Form(attributes ...string) {
	c.writeWithAttributes("form", attributes)
}

func (c *Context) Form_() {
	c.write_("form")
}

func (c *Context) H1(attributes ...string) {
	c.writeWithAttributes("h1", attributes)
}

func (c *Context) H1_() {
	c.write_("h1")
}

func (c *Context) H2(attributes ...string) {
	c.writeWithAttributes("h2", attributes)
}

func (c *Context) H2_() {
	c.write_("h2")
}

func (c *Context) H3(attributes ...string) {
	c.writeWithAttributes("h3", attributes)
}

func (c *Context) H3_() {
	c.write_("h3")
}

func (c *Context) H4(attributes ...string) {
	c.writeWithAttributes("h4", attributes)
}

func (c *Context) H4_() {
	c.write_("h4")
}

func (c *Context) H5(attributes ...string) {
	c.writeWithAttributes("h5", attributes)
}

func (c *Context) H5_() {
	c.write_("h5")
}

func (c *Context) H6(attributes ...string) {
	c.writeWithAttributes("h6", attributes)
}

func (c *Context) H6_() {
	c.write_("h6")
}

func (c *Context) Head(attributes ...string) {
	c.writeWithAttributes("head", attributes)
}

func (c *Context) Head_() {
	c.write_("head")
}

func (c *Context) Header(attributes ...string) {
	c.writeWithAttributes("header", attributes)
}

func (c *Context) Header_() {
	c.write_("header")
}

func (c *Context) Hgroup(attributes ...string) {
	c.writeWithAttributes("hgroup", attributes)
}

func (c *Context) Hgroup_() {
	c.write_("hgroup")
}

func (c *Context) Hr(attributes ...string) {
	c.writeWithAttributes("hr", attributes)
}

func (c *Context) I(attributes ...string) {
	c.writeWithAttributes("i", attributes)
}

func (c *Context) I_() {
	c.write_("i")
}

func (c *Context) Iframe(attributes ...string) {
	c.writeWithAttributes("iframe", attributes)
}

func (c *Context) Iframe_() {
	c.write_("iframe")
}

func (c *Context) Img(attributes ...string) {
	c.writeWithAttributes("img", attributes)
}

func (c *Context) Input(attributes ...string) {
	c.writeWithAttributes("input", attributes)
}

func (c *Context) Ins(attributes ...string) {
	c.writeWithAttributes("ins", attributes)
}

func (c *Context) Ins_() {
	c.write_("ins")
}

func (c *Context) Kbd(attributes ...string) {
	c.writeWithAttributes("kbd", attributes)
}

func (c *Context) Kbd_() {
	c.write_("kbd")
}

func (c *Context) Label(attributes ...string) {
	c.writeWithAttributes("label", attributes)
}

func (c *Context) Label_() {
	c.write_("label")
}

func (c *Context) Legend(attributes ...string) {
	c.writeWithAttributes("legend", attributes)
}

func (c *Context) Legend_() {
	c.write_("legend")
}

func (c *Context) Li(attributes ...string) {
	c.writeWithAttributes("li", attributes)
}

func (c *Context) Li_() {
	c.write_("li")
}

func (c *Context) Link(attributes ...string) {
	c.writeWithAttributes("link", attributes)
}

func (c *Context) Main(attributes ...string) {
	c.writeWithAttributes("main", attributes)
}

func (c *Context) Main_() {
	c.write_("main")
}

func (c *Context) Map(attributes ...string) {
	c.writeWithAttributes("map", attributes)
}

func (c *Context) Map_() {
	c.write_("map")
}

func (c *Context) Mark(attributes ...string) {
	c.writeWithAttributes("mark", attributes)
}

func (c *Context) Mark_() {
	c.write_("mark")
}

func (c *Context) Meta(attributes ...string) {
	c.writeWithAttributes("meta", attributes)
}

func (c *Context) Meter(attributes ...string) {
	c.writeWithAttributes("meter", attributes)
}

func (c *Context) Meter_() {
	c.write_("meter")
}

func (c *Context) Nav(attributes ...string) {
	c.writeWithAttributes("nav", attributes)
}

func (c *Context) Nav_() {
	c.write_("nav")
}

func (c *Context) Noscript(attributes ...string) {
	c.writeWithAttributes("noscript", attributes)
}

func (c *Context) Noscript_() {
	c.write_("noscript")
}

func (c *Context) Object(attributes ...string) {
	c.writeWithAttributes("object", attributes)
}

func (c *Context) Object_() {
	c.write_("object")
}

func (c *Context) Ol(attributes ...string) {
	c.writeWithAttributes("ol", attributes)
}

func (c *Context) Ol_() {
	c.write_("ol")
}

func (c *Context) Optgroup(attributes ...string) {
	c.writeWithAttributes("optgroup", attributes)
}

func (c *Context) Optgroup_() {
	c.write_("optgroup")
}

func (c *Context) Option(attributes ...string) {
	c.writeWithAttributes("option", attributes)
}

func (c *Context) Option_() {
	c.write_("option")
}

func (c *Context) Output(attributes ...string) {
	c.writeWithAttributes("output", attributes)
}

func (c *Context) Output_() {
	c.write_("output")
}

func (c *Context) P(attributes ...string) {
	c.writeWithAttributes("p", attributes)
}

func (c *Context) P_() {
	c.write_("p")
}

func (c *Context) Param(attributes ...string) {
	c.writeWithAttributes("param", attributes)
}

func (c *Context) Picture(attributes ...string) {
	c.writeWithAttributes("picture", attributes)
}

func (c *Context) Picture_() {
	c.write_("picture")
}

func (c *Context) Pre(attributes ...string) {
	c.writeWithAttributes("pre", attributes)
}

func (c *Context) Pre_() {
	c.write_("pre")
}

func (c *Context) Progress(attributes ...string) {
	c.writeWithAttributes("progress", attributes)
}

func (c *Context) Progress_() {
	c.write_("progress")
}

func (c *Context) Q(attributes ...string) {
	c.writeWithAttributes("q", attributes)
}

func (c *Context) Q_() {
	c.write_("q")
}

func (c *Context) Rp(attributes ...string) {
	c.writeWithAttributes("rp", attributes)
}

func (c *Context) Rp_() {
	c.write_("rp")
}

func (c *Context) Rt(attributes ...string) {
	c.writeWithAttributes("rt", attributes)
}

func (c *Context) Rt_() {
	c.write_("rt")
}

func (c *Context) Ruby(attributes ...string) {
	c.writeWithAttributes("ruby", attributes)
}

func (c *Context) Ruby_() {
	c.write_("ruby")
}

func (c *Context) S(attributes ...string) {
	c.writeWithAttributes("s", attributes)
}

func (c *Context) S_() {
	c.write_("s")
}

func (c *Context) Samp(attributes ...string) {
	c.writeWithAttributes("samp", attributes)
}

func (c *Context) Samp_() {
	c.write_("samp")
}

func (c *Context) Script(attributes ...string) {
	c.writeWithAttributes("script", attributes)
}

func (c *Context) Script_() {
	c.write_("script")
}

func (c *Context) Section(attributes ...string) {
	c.writeWithAttributes("section", attributes)
}

func (c *Context) Section_() {
	c.write_("section")
}

func (c *Context) Select(attributes ...string) {
	c.writeWithAttributes("select", attributes)
}

func (c *Context) Select_() {
	c.write_("select")
}

func (c *Context) Small(attributes ...string) {
	c.writeWithAttributes("small", attributes)
}

func (c *Context) Small_() {
	c.write_("small")
}

func (c *Context) Source(attributes ...string) {
	c.writeWithAttributes("source", attributes)
}

func (c *Context) Span(attributes ...string) {
	c.writeWithAttributes("span", attributes)
}

func (c *Context) Span_() {
	c.write_("span")
}

func (c *Context) Strong(attributes ...string) {
	c.writeWithAttributes("strong", attributes)
}

func (c *Context) Strong_() {
	c.write_("strong")
}

func (c *Context) Style(attributes ...string) {
	c.writeWithAttributes("style", attributes)
}

func (c *Context) Style_() {
	c.write_("style")
}

func (c *Context) Sub(attributes ...string) {
	c.writeWithAttributes("sub", attributes)
}

func (c *Context) Sub_() {
	c.write_("sub")
}

func (c *Context) Summary(attributes ...string) {
	c.writeWithAttributes("summary", attributes)
}

func (c *Context) Summary_() {
	c.write_("summary")
}

func (c *Context) Sup(attributes ...string) {
	c.writeWithAttributes("sup", attributes)
}

func (c *Context) Sup_() {
	c.write_("sup")
}

func (c *Context) Svg(attributes ...string) {
	c.writeWithAttributes("svg", attributes)
}

func (c *Context) Svg_() {
	c.write_("svg")
}

func (c *Context) Table(attributes ...string) {
	c.writeWithAttributes("table", attributes)
}

func (c *Context) Table_() {
	c.write_("table")
}

func (c *Context) Tbody(attributes ...string) {
	c.writeWithAttributes("tbody", attributes)
}

func (c *Context) Tbody_() {
	c.write_("tbody")
}

func (c *Context) Td(attributes ...string) {
	c.writeWithAttributes("td", attributes)
}

func (c *Context) Td_() {
	c.write_("td")
}

func (c *Context) Template(attributes ...string) {
	c.writeWithAttributes("template", attributes)
}

func (c *Context) Template_() {
	c.write_("template")
}

func (c *Context) Textarea(attributes ...string) {
	c.writeWithAttributes("textarea", attributes)
}

func (c *Context) Textarea_() {
	c.write_("textarea")
}

func (c *Context) Tfoot(attributes ...string) {
	c.writeWithAttributes("tfoot", attributes)
}

func (c *Context) Tfoot_() {
	c.write_("tfoot")
}

func (c *Context) Th(attributes ...string) {
	c.writeWithAttributes("th", attributes)
}

func (c *Context) Th_() {
	c.write_("th")
}

func (c *Context) Thead(attributes ...string) {
	c.writeWithAttributes("thead", attributes)
}

func (c *Context) Thead_() {
	c.write_("thead")
}

func (c *Context) Time(attributes ...string) {
	c.writeWithAttributes("time", attributes)
}

func (c *Context) Time_() {
	c.write_("time")
}

func (c *Context) Title(attributes ...string) {
	c.writeWithAttributes("title", attributes)
}

func (c *Context) Title_() {
	c.write_("title")
}

func (c *Context) Tr(attributes ...string) {
	c.writeWithAttributes("tr", attributes)
}

func (c *Context) Tr_() {
	c.write_("tr")
}

func (c *Context) Track(attributes ...string) {
	c.writeWithAttributes("track", attributes)
}

func (c *Context) U(attributes ...string) {
	c.writeWithAttributes("u", attributes)
}

func (c *Context) U_() {
	c.write_("u")
}

func (c *Context) Ul(attributes ...string) {
	c.writeWithAttributes("ul", attributes)
}

func (c *Context) Ul_() {
	c.write_("ul")
}

func (c *Context) Var(attributes ...string) {
	c.writeWithAttributes("var", attributes)
}

func (c *Context) Var_() {
	c.write_("var")
}

func (c *Context) Video(attributes ...string) {
	c.writeWithAttributes("video", attributes)
}

func (c *Context) Video_() {
	c.write_("video")
}

func (c *Context) Wbr(attributes ...string) {
	c.writeWithAttributes("wbr", attributes)
}
