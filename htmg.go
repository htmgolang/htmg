package htmg

import (
	"fmt"
	"io"
)

type HtmgCtx struct {
	writer io.Writer
}

func New(w io.Writer) HtmgCtx {
	return HtmgCtx{writer: w}
}

func (c *HtmgCtx) write(str string) {
	c.writer.Write([]byte(str))
}

func (c *HtmgCtx) writeWithAttributes(str string, attributes []string) {
	c.write("<")
	c.write(str)
	for _, attribute := range attributes {
		c.write(" ")
		c.write(attribute)
	}
	c.write(">")
}

func (c *HtmgCtx) write_(str string) {
	c.write("</")
	c.write(str)
	c.write(">")
}

func (c *HtmgCtx) Text(str string) {
	c.write(str)
}

func (c *HtmgCtx) Textf(format string, a ...any) {
	c.write(fmt.Sprintf(format, a...))
}

func (c *HtmgCtx) Html(attributes ...string) {
	c.write("<!DOCTYPE html>")
	c.writeWithAttributes("html", attributes)
}

func (c *HtmgCtx) Html_() {
	c.write_("html")
}
func (c *HtmgCtx) A(attributes ...string) {
	c.writeWithAttributes("a", attributes)
}

func (c *HtmgCtx) A_() {
	c.write_("a")
}

func (c *HtmgCtx) Abbr(attributes ...string) {
	c.writeWithAttributes("abbr", attributes)
}

func (c *HtmgCtx) Abbr_() {
	c.write_("abbr")
}

func (c *HtmgCtx) Address(attributes ...string) {
	c.writeWithAttributes("address", attributes)
}

func (c *HtmgCtx) Address_() {
	c.write_("address")
}

func (c *HtmgCtx) Area(attributes ...string) {
	c.writeWithAttributes("area", attributes)
}

func (c *HtmgCtx) Article(attributes ...string) {
	c.writeWithAttributes("article", attributes)
}

func (c *HtmgCtx) Article_() {
	c.write_("article")
}

func (c *HtmgCtx) Aside(attributes ...string) {
	c.writeWithAttributes("aside", attributes)
}

func (c *HtmgCtx) Aside_() {
	c.write_("aside")
}

func (c *HtmgCtx) Audio(attributes ...string) {
	c.writeWithAttributes("audio", attributes)
}

func (c *HtmgCtx) Audio_() {
	c.write_("audio")
}

func (c *HtmgCtx) B(attributes ...string) {
	c.writeWithAttributes("b", attributes)
}

func (c *HtmgCtx) B_() {
	c.write_("b")
}

func (c *HtmgCtx) Base(attributes ...string) {
	c.writeWithAttributes("base", attributes)
}

func (c *HtmgCtx) Bdi(attributes ...string) {
	c.writeWithAttributes("bdi", attributes)
}

func (c *HtmgCtx) Bdi_() {
	c.write_("bdi")
}

func (c *HtmgCtx) Bdo(attributes ...string) {
	c.writeWithAttributes("bdo", attributes)
}

func (c *HtmgCtx) Bdo_() {
	c.write_("bdo")
}

func (c *HtmgCtx) Blockquote(attributes ...string) {
	c.writeWithAttributes("blockquote", attributes)
}

func (c *HtmgCtx) Blockquote_() {
	c.write_("blockquote")
}

func (c *HtmgCtx) Body(attributes ...string) {
	c.writeWithAttributes("body", attributes)
}

func (c *HtmgCtx) Body_() {
	c.write_("body")
}

func (c *HtmgCtx) Br(attributes ...string) {
	c.writeWithAttributes("br", attributes)
}

func (c *HtmgCtx) Button(attributes ...string) {
	c.writeWithAttributes("button", attributes)
}

func (c *HtmgCtx) Button_() {
	c.write_("button")
}

func (c *HtmgCtx) Canvas(attributes ...string) {
	c.writeWithAttributes("canvas", attributes)
}

func (c *HtmgCtx) Canvas_() {
	c.write_("canvas")
}

func (c *HtmgCtx) Caption(attributes ...string) {
	c.writeWithAttributes("caption", attributes)
}

func (c *HtmgCtx) Caption_() {
	c.write_("caption")
}

func (c *HtmgCtx) Cite(attributes ...string) {
	c.writeWithAttributes("cite", attributes)
}

func (c *HtmgCtx) Cite_() {
	c.write_("cite")
}

func (c *HtmgCtx) Code(attributes ...string) {
	c.writeWithAttributes("code", attributes)
}

func (c *HtmgCtx) Code_() {
	c.write_("code")
}

func (c *HtmgCtx) Col(attributes ...string) {
	c.writeWithAttributes("col", attributes)
}

func (c *HtmgCtx) Colgroup(attributes ...string) {
	c.writeWithAttributes("colgroup", attributes)
}

func (c *HtmgCtx) Colgroup_() {
	c.write_("colgroup")
}

func (c *HtmgCtx) Data(attributes ...string) {
	c.writeWithAttributes("data", attributes)
}

func (c *HtmgCtx) Data_() {
	c.write_("data")
}

func (c *HtmgCtx) Datalist(attributes ...string) {
	c.writeWithAttributes("datalist", attributes)
}

func (c *HtmgCtx) Datalist_() {
	c.write_("datalist")
}

func (c *HtmgCtx) Dd(attributes ...string) {
	c.writeWithAttributes("dd", attributes)
}

func (c *HtmgCtx) Dd_() {
	c.write_("dd")
}

func (c *HtmgCtx) Del(attributes ...string) {
	c.writeWithAttributes("del", attributes)
}

func (c *HtmgCtx) Del_() {
	c.write_("del")
}

func (c *HtmgCtx) Details(attributes ...string) {
	c.writeWithAttributes("details", attributes)
}

func (c *HtmgCtx) Details_() {
	c.write_("details")
}

func (c *HtmgCtx) Dfn(attributes ...string) {
	c.writeWithAttributes("dfn", attributes)
}

func (c *HtmgCtx) Dfn_() {
	c.write_("dfn")
}

func (c *HtmgCtx) Dialog(attributes ...string) {
	c.writeWithAttributes("dialog", attributes)
}

func (c *HtmgCtx) Dialog_() {
	c.write_("dialog")
}

func (c *HtmgCtx) Div(attributes ...string) {
	c.writeWithAttributes("div", attributes)
}

func (c *HtmgCtx) Div_() {
	c.write_("div")
}

func (c *HtmgCtx) Dl(attributes ...string) {
	c.writeWithAttributes("dl", attributes)
}

func (c *HtmgCtx) Dl_() {
	c.write_("dl")
}

func (c *HtmgCtx) Dt(attributes ...string) {
	c.writeWithAttributes("dt", attributes)
}

func (c *HtmgCtx) Dt_() {
	c.write_("dt")
}

func (c *HtmgCtx) Em(attributes ...string) {
	c.writeWithAttributes("em", attributes)
}

func (c *HtmgCtx) Em_() {
	c.write_("em")
}

func (c *HtmgCtx) Embed(attributes ...string) {
	c.writeWithAttributes("embed", attributes)
}

func (c *HtmgCtx) Fieldset(attributes ...string) {
	c.writeWithAttributes("fieldset", attributes)
}

func (c *HtmgCtx) Fieldset_() {
	c.write_("fieldset")
}

func (c *HtmgCtx) Figcaption(attributes ...string) {
	c.writeWithAttributes("figcaption", attributes)
}

func (c *HtmgCtx) Figcaption_() {
	c.write_("figcaption")
}

func (c *HtmgCtx) Figure(attributes ...string) {
	c.writeWithAttributes("figure", attributes)
}

func (c *HtmgCtx) Figure_() {
	c.write_("figure")
}

func (c *HtmgCtx) Footer(attributes ...string) {
	c.writeWithAttributes("footer", attributes)
}

func (c *HtmgCtx) Footer_() {
	c.write_("footer")
}

func (c *HtmgCtx) Form(attributes ...string) {
	c.writeWithAttributes("form", attributes)
}

func (c *HtmgCtx) Form_() {
	c.write_("form")
}

func (c *HtmgCtx) H1(attributes ...string) {
	c.writeWithAttributes("h1", attributes)
}

func (c *HtmgCtx) H1_() {
	c.write_("h1")
}

func (c *HtmgCtx) H2(attributes ...string) {
	c.writeWithAttributes("h2", attributes)
}

func (c *HtmgCtx) H2_() {
	c.write_("h2")
}

func (c *HtmgCtx) H3(attributes ...string) {
	c.writeWithAttributes("h3", attributes)
}

func (c *HtmgCtx) H3_() {
	c.write_("h3")
}

func (c *HtmgCtx) H4(attributes ...string) {
	c.writeWithAttributes("h4", attributes)
}

func (c *HtmgCtx) H4_() {
	c.write_("h4")
}

func (c *HtmgCtx) H5(attributes ...string) {
	c.writeWithAttributes("h5", attributes)
}

func (c *HtmgCtx) H5_() {
	c.write_("h5")
}

func (c *HtmgCtx) H6(attributes ...string) {
	c.writeWithAttributes("h6", attributes)
}

func (c *HtmgCtx) H6_() {
	c.write_("h6")
}

func (c *HtmgCtx) Head(attributes ...string) {
	c.writeWithAttributes("head", attributes)
}

func (c *HtmgCtx) Head_() {
	c.write_("head")
}

func (c *HtmgCtx) Header(attributes ...string) {
	c.writeWithAttributes("header", attributes)
}

func (c *HtmgCtx) Header_() {
	c.write_("header")
}

func (c *HtmgCtx) Hgroup(attributes ...string) {
	c.writeWithAttributes("hgroup", attributes)
}

func (c *HtmgCtx) Hgroup_() {
	c.write_("hgroup")
}

func (c *HtmgCtx) Hr(attributes ...string) {
	c.writeWithAttributes("hr", attributes)
}

func (c *HtmgCtx) I(attributes ...string) {
	c.writeWithAttributes("i", attributes)
}

func (c *HtmgCtx) I_() {
	c.write_("i")
}

func (c *HtmgCtx) Iframe(attributes ...string) {
	c.writeWithAttributes("iframe", attributes)
}

func (c *HtmgCtx) Iframe_() {
	c.write_("iframe")
}

func (c *HtmgCtx) Img(attributes ...string) {
	c.writeWithAttributes("img", attributes)
}

func (c *HtmgCtx) Input(attributes ...string) {
	c.writeWithAttributes("input", attributes)
}

func (c *HtmgCtx) Ins(attributes ...string) {
	c.writeWithAttributes("ins", attributes)
}

func (c *HtmgCtx) Ins_() {
	c.write_("ins")
}

func (c *HtmgCtx) Kbd(attributes ...string) {
	c.writeWithAttributes("kbd", attributes)
}

func (c *HtmgCtx) Kbd_() {
	c.write_("kbd")
}

func (c *HtmgCtx) Label(attributes ...string) {
	c.writeWithAttributes("label", attributes)
}

func (c *HtmgCtx) Label_() {
	c.write_("label")
}

func (c *HtmgCtx) Legend(attributes ...string) {
	c.writeWithAttributes("legend", attributes)
}

func (c *HtmgCtx) Legend_() {
	c.write_("legend")
}

func (c *HtmgCtx) Li(attributes ...string) {
	c.writeWithAttributes("li", attributes)
}

func (c *HtmgCtx) Li_() {
	c.write_("li")
}

func (c *HtmgCtx) Link(attributes ...string) {
	c.writeWithAttributes("link", attributes)
}

func (c *HtmgCtx) Main(attributes ...string) {
	c.writeWithAttributes("main", attributes)
}

func (c *HtmgCtx) Main_() {
	c.write_("main")
}

func (c *HtmgCtx) Map(attributes ...string) {
	c.writeWithAttributes("map", attributes)
}

func (c *HtmgCtx) Map_() {
	c.write_("map")
}

func (c *HtmgCtx) Mark(attributes ...string) {
	c.writeWithAttributes("mark", attributes)
}

func (c *HtmgCtx) Mark_() {
	c.write_("mark")
}

func (c *HtmgCtx) Meta(attributes ...string) {
	c.writeWithAttributes("meta", attributes)
}

func (c *HtmgCtx) Meter(attributes ...string) {
	c.writeWithAttributes("meter", attributes)
}

func (c *HtmgCtx) Meter_() {
	c.write_("meter")
}

func (c *HtmgCtx) Nav(attributes ...string) {
	c.writeWithAttributes("nav", attributes)
}

func (c *HtmgCtx) Nav_() {
	c.write_("nav")
}

func (c *HtmgCtx) Noscript(attributes ...string) {
	c.writeWithAttributes("noscript", attributes)
}

func (c *HtmgCtx) Noscript_() {
	c.write_("noscript")
}

func (c *HtmgCtx) Object(attributes ...string) {
	c.writeWithAttributes("object", attributes)
}

func (c *HtmgCtx) Object_() {
	c.write_("object")
}

func (c *HtmgCtx) Ol(attributes ...string) {
	c.writeWithAttributes("ol", attributes)
}

func (c *HtmgCtx) Ol_() {
	c.write_("ol")
}

func (c *HtmgCtx) Optgroup(attributes ...string) {
	c.writeWithAttributes("optgroup", attributes)
}

func (c *HtmgCtx) Optgroup_() {
	c.write_("optgroup")
}

func (c *HtmgCtx) Option(attributes ...string) {
	c.writeWithAttributes("option", attributes)
}

func (c *HtmgCtx) Option_() {
	c.write_("option")
}

func (c *HtmgCtx) Output(attributes ...string) {
	c.writeWithAttributes("output", attributes)
}

func (c *HtmgCtx) Output_() {
	c.write_("output")
}

func (c *HtmgCtx) P(attributes ...string) {
	c.writeWithAttributes("p", attributes)
}

func (c *HtmgCtx) P_() {
	c.write_("p")
}

func (c *HtmgCtx) Param(attributes ...string) {
	c.writeWithAttributes("param", attributes)
}

func (c *HtmgCtx) Picture(attributes ...string) {
	c.writeWithAttributes("picture", attributes)
}

func (c *HtmgCtx) Picture_() {
	c.write_("picture")
}

func (c *HtmgCtx) Pre(attributes ...string) {
	c.writeWithAttributes("pre", attributes)
}

func (c *HtmgCtx) Pre_() {
	c.write_("pre")
}

func (c *HtmgCtx) Progress(attributes ...string) {
	c.writeWithAttributes("progress", attributes)
}

func (c *HtmgCtx) Progress_() {
	c.write_("progress")
}

func (c *HtmgCtx) Q(attributes ...string) {
	c.writeWithAttributes("q", attributes)
}

func (c *HtmgCtx) Q_() {
	c.write_("q")
}

func (c *HtmgCtx) Rp(attributes ...string) {
	c.writeWithAttributes("rp", attributes)
}

func (c *HtmgCtx) Rp_() {
	c.write_("rp")
}

func (c *HtmgCtx) Rt(attributes ...string) {
	c.writeWithAttributes("rt", attributes)
}

func (c *HtmgCtx) Rt_() {
	c.write_("rt")
}

func (c *HtmgCtx) Ruby(attributes ...string) {
	c.writeWithAttributes("ruby", attributes)
}

func (c *HtmgCtx) Ruby_() {
	c.write_("ruby")
}

func (c *HtmgCtx) S(attributes ...string) {
	c.writeWithAttributes("s", attributes)
}

func (c *HtmgCtx) S_() {
	c.write_("s")
}

func (c *HtmgCtx) Samp(attributes ...string) {
	c.writeWithAttributes("samp", attributes)
}

func (c *HtmgCtx) Samp_() {
	c.write_("samp")
}

func (c *HtmgCtx) Script(attributes ...string) {
	c.writeWithAttributes("script", attributes)
}

func (c *HtmgCtx) Script_() {
	c.write_("script")
}

func (c *HtmgCtx) Section(attributes ...string) {
	c.writeWithAttributes("section", attributes)
}

func (c *HtmgCtx) Section_() {
	c.write_("section")
}

func (c *HtmgCtx) Select(attributes ...string) {
	c.writeWithAttributes("select", attributes)
}

func (c *HtmgCtx) Select_() {
	c.write_("select")
}

func (c *HtmgCtx) Small(attributes ...string) {
	c.writeWithAttributes("small", attributes)
}

func (c *HtmgCtx) Small_() {
	c.write_("small")
}

func (c *HtmgCtx) Source(attributes ...string) {
	c.writeWithAttributes("source", attributes)
}

func (c *HtmgCtx) Span(attributes ...string) {
	c.writeWithAttributes("span", attributes)
}

func (c *HtmgCtx) Span_() {
	c.write_("span")
}

func (c *HtmgCtx) Strong(attributes ...string) {
	c.writeWithAttributes("strong", attributes)
}

func (c *HtmgCtx) Strong_() {
	c.write_("strong")
}

func (c *HtmgCtx) Style(attributes ...string) {
	c.writeWithAttributes("style", attributes)
}

func (c *HtmgCtx) Style_() {
	c.write_("style")
}

func (c *HtmgCtx) Sub(attributes ...string) {
	c.writeWithAttributes("sub", attributes)
}

func (c *HtmgCtx) Sub_() {
	c.write_("sub")
}

func (c *HtmgCtx) Summary(attributes ...string) {
	c.writeWithAttributes("summary", attributes)
}

func (c *HtmgCtx) Summary_() {
	c.write_("summary")
}

func (c *HtmgCtx) Sup(attributes ...string) {
	c.writeWithAttributes("sup", attributes)
}

func (c *HtmgCtx) Sup_() {
	c.write_("sup")
}

func (c *HtmgCtx) Svg(attributes ...string) {
	c.writeWithAttributes("svg", attributes)
}

func (c *HtmgCtx) Svg_() {
	c.write_("svg")
}

func (c *HtmgCtx) Table(attributes ...string) {
	c.writeWithAttributes("table", attributes)
}

func (c *HtmgCtx) Table_() {
	c.write_("table")
}

func (c *HtmgCtx) Tbody(attributes ...string) {
	c.writeWithAttributes("tbody", attributes)
}

func (c *HtmgCtx) Tbody_() {
	c.write_("tbody")
}

func (c *HtmgCtx) Td(attributes ...string) {
	c.writeWithAttributes("td", attributes)
}

func (c *HtmgCtx) Td_() {
	c.write_("td")
}

func (c *HtmgCtx) Template(attributes ...string) {
	c.writeWithAttributes("template", attributes)
}

func (c *HtmgCtx) Template_() {
	c.write_("template")
}

func (c *HtmgCtx) Textarea(attributes ...string) {
	c.writeWithAttributes("textarea", attributes)
}

func (c *HtmgCtx) Textarea_() {
	c.write_("textarea")
}

func (c *HtmgCtx) Tfoot(attributes ...string) {
	c.writeWithAttributes("tfoot", attributes)
}

func (c *HtmgCtx) Tfoot_() {
	c.write_("tfoot")
}

func (c *HtmgCtx) Th(attributes ...string) {
	c.writeWithAttributes("th", attributes)
}

func (c *HtmgCtx) Th_() {
	c.write_("th")
}

func (c *HtmgCtx) Thead(attributes ...string) {
	c.writeWithAttributes("thead", attributes)
}

func (c *HtmgCtx) Thead_() {
	c.write_("thead")
}

func (c *HtmgCtx) Time(attributes ...string) {
	c.writeWithAttributes("time", attributes)
}

func (c *HtmgCtx) Time_() {
	c.write_("time")
}

func (c *HtmgCtx) Title(attributes ...string) {
	c.writeWithAttributes("title", attributes)
}

func (c *HtmgCtx) Title_() {
	c.write_("title")
}

func (c *HtmgCtx) Tr(attributes ...string) {
	c.writeWithAttributes("tr", attributes)
}

func (c *HtmgCtx) Tr_() {
	c.write_("tr")
}

func (c *HtmgCtx) Track(attributes ...string) {
	c.writeWithAttributes("track", attributes)
}

func (c *HtmgCtx) U(attributes ...string) {
	c.writeWithAttributes("u", attributes)
}

func (c *HtmgCtx) U_() {
	c.write_("u")
}

func (c *HtmgCtx) Ul(attributes ...string) {
	c.writeWithAttributes("ul", attributes)
}

func (c *HtmgCtx) Ul_() {
	c.write_("ul")
}

func (c *HtmgCtx) Var(attributes ...string) {
	c.writeWithAttributes("var", attributes)
}

func (c *HtmgCtx) Var_() {
	c.write_("var")
}

func (c *HtmgCtx) Video(attributes ...string) {
	c.writeWithAttributes("video", attributes)
}

func (c *HtmgCtx) Video_() {
	c.write_("video")
}

func (c *HtmgCtx) Wbr(attributes ...string) {
	c.writeWithAttributes("wbr", attributes)
}
