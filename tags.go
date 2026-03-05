package gotags

import "fmt"

func normalizeTagContent(tag string, content []any) []HTML {
	c, err := normalizeContent(content)
	if err != nil {
		panic(fmt.Sprintf(
			err.Error()+"; in Tag <%s>",
			tag,
		))
	}
	return c
}

func Base(content ...any) HTML {
	return NewTagComponent("base", true, normalizeTagContent("base", content)...)
}
func Head(content ...any) HTML {
	return NewTagComponent("head", false, normalizeTagContent("head", content)...)
}
func Link(content ...any) HTML {
	return NewTagComponent("link", true, normalizeTagContent("link", content)...)
}
func Meta(content ...any) HTML {
	return NewTagComponent("meta", true, normalizeTagContent("meta", content)...)
}
func Style(content ...any) HTML {
	return NewTagComponent("style", false, normalizeTagContent("style", content)...)
}
func Title(content ...any) HTML {
	return NewTagComponent("title", false, normalizeTagContent("title", content)...)
}
func Body(content ...any) HTML {
	return NewTagComponent("body", false, normalizeTagContent("body", content)...)
}
func Address(content ...any) HTML {
	return NewTagComponent("address", false, normalizeTagContent("address", content)...)
}
func Article(content ...any) HTML {
	return NewTagComponent("article", false, normalizeTagContent("article", content)...)
}
func Aside(content ...any) HTML {
	return NewTagComponent("aside", false, normalizeTagContent("aside", content)...)
}
func Footer(content ...any) HTML {
	return NewTagComponent("footer", false, normalizeTagContent("footer", content)...)
}
func Header(content ...any) HTML {
	return NewTagComponent("header", false, normalizeTagContent("header", content)...)
}
func H1(content ...any) HTML {
	return NewTagComponent("h1", false, normalizeTagContent("h1", content)...)
}
func H2(content ...any) HTML {
	return NewTagComponent("h2", false, normalizeTagContent("h2", content)...)
}
func H3(content ...any) HTML {
	return NewTagComponent("h3", false, normalizeTagContent("h3", content)...)
}
func H4(content ...any) HTML {
	return NewTagComponent("h4", false, normalizeTagContent("h4", content)...)
}
func H5(content ...any) HTML {
	return NewTagComponent("h5", false, normalizeTagContent("h5", content)...)
}
func H6(content ...any) HTML {
	return NewTagComponent("h6", false, normalizeTagContent("h6", content)...)
}
func Hgroup(content ...any) HTML {
	return NewTagComponent("hgroup", false, normalizeTagContent("hgroup", content)...)
}
func Main(content ...any) HTML {
	return NewTagComponent("main", false, normalizeTagContent("main", content)...)
}
func Nav(content ...any) HTML {
	return NewTagComponent("nav", false, normalizeTagContent("nav", content)...)
}
func Section(content ...any) HTML {
	return NewTagComponent("section", false, normalizeTagContent("section", content)...)
}
func Search(content ...any) HTML {
	return NewTagComponent("search", false, normalizeTagContent("search", content)...)
}
func Blockquote(content ...any) HTML {
	return NewTagComponent("blockquote", false, normalizeTagContent("blockquote", content)...)
}
func Dd(content ...any) HTML {
	return NewTagComponent("dd", false, normalizeTagContent("dd", content)...)
}
func Div(content ...any) HTML {
	return NewTagComponent("div", false, normalizeTagContent("div", content)...)
}
func Dl(content ...any) HTML {
	return NewTagComponent("dl", false, normalizeTagContent("dl", content)...)
}
func Dt(content ...any) HTML {
	return NewTagComponent("dt", false, normalizeTagContent("dt", content)...)
}
func Figcaption(content ...any) HTML {
	return NewTagComponent("figcaption", false, normalizeTagContent("figcaption", content)...)
}
func Figure(content ...any) HTML {
	return NewTagComponent("figure", false, normalizeTagContent("figure", content)...)
}
func Hr(content ...any) HTML {
	return NewTagComponent("hr", true, normalizeTagContent("hr", content)...)
}
func Li(content ...any) HTML {
	return NewTagComponent("li", false, normalizeTagContent("li", content)...)
}
func Menu(content ...any) HTML {
	return NewTagComponent("menu", false, normalizeTagContent("menu", content)...)
}
func Ol(content ...any) HTML {
	return NewTagComponent("ol", false, normalizeTagContent("ol", content)...)
}
func P(content ...any) HTML {
	return NewTagComponent("p", false, normalizeTagContent("p", content)...)
}
func Pre(content ...any) HTML {
	return NewTagComponent("pre", false, normalizeTagContent("pre", content)...)
}
func Ul(content ...any) HTML {
	return NewTagComponent("ul", false, normalizeTagContent("ul", content)...)
}
func A(content ...any) HTML {
	return NewTagComponent("a", false, normalizeTagContent("a", content)...)
}
func Abbr(content ...any) HTML {
	return NewTagComponent("abbr", false, normalizeTagContent("abbr", content)...)
}
func B(content ...any) HTML {
	return NewTagComponent("b", false, normalizeTagContent("b", content)...)
}
func Bdi(content ...any) HTML {
	return NewTagComponent("bdi", false, normalizeTagContent("bdi", content)...)
}
func Bdo(content ...any) HTML {
	return NewTagComponent("bdo", false, normalizeTagContent("bdo", content)...)
}
func Br(content ...any) HTML {
	return NewTagComponent("br", true, normalizeTagContent("br", content)...)
}
func Cite(content ...any) HTML {
	return NewTagComponent("cite", false, normalizeTagContent("cite", content)...)
}
func Code(content ...any) HTML {
	return NewTagComponent("code", false, normalizeTagContent("code", content)...)
}
func Data(content ...any) HTML {
	return NewTagComponent("data", false, normalizeTagContent("data", content)...)
}
func Dfn(content ...any) HTML {
	return NewTagComponent("dfn", false, normalizeTagContent("dfn", content)...)
}
func Em(content ...any) HTML {
	return NewTagComponent("em", false, normalizeTagContent("em", content)...)
}
func I(content ...any) HTML {
	return NewTagComponent("i", false, normalizeTagContent("i", content)...)
}
func Kbd(content ...any) HTML {
	return NewTagComponent("kbd", false, normalizeTagContent("kbd", content)...)
}
func Mark(content ...any) HTML {
	return NewTagComponent("mark", false, normalizeTagContent("mark", content)...)
}
func Q(content ...any) HTML {
	return NewTagComponent("q", false, normalizeTagContent("q", content)...)
}
func Rp(content ...any) HTML {
	return NewTagComponent("rp", false, normalizeTagContent("rp", content)...)
}
func Rt(content ...any) HTML {
	return NewTagComponent("rt", false, normalizeTagContent("rt", content)...)
}
func Ruby(content ...any) HTML {
	return NewTagComponent("ruby", false, normalizeTagContent("ruby", content)...)
}
func S(content ...any) HTML {
	return NewTagComponent("s", false, normalizeTagContent("s", content)...)
}
func Samp(content ...any) HTML {
	return NewTagComponent("samp", false, normalizeTagContent("samp", content)...)
}
func Small(content ...any) HTML {
	return NewTagComponent("small", false, normalizeTagContent("small", content)...)
}
func Span(content ...any) HTML {
	return NewTagComponent("span", false, normalizeTagContent("span", content)...)
}
func Strong(content ...any) HTML {
	return NewTagComponent("strong", false, normalizeTagContent("strong", content)...)
}
func Sub(content ...any) HTML {
	return NewTagComponent("sub", false, normalizeTagContent("sub", content)...)
}
func Sup(content ...any) HTML {
	return NewTagComponent("sup", false, normalizeTagContent("sup", content)...)
}
func Time(content ...any) HTML {
	return NewTagComponent("time", false, normalizeTagContent("time", content)...)
}
func U(content ...any) HTML {
	return NewTagComponent("u", false, normalizeTagContent("u", content)...)
}
func Var(content ...any) HTML {
	return NewTagComponent("var", false, normalizeTagContent("var", content)...)
}
func Wbr(content ...any) HTML {
	return NewTagComponent("wbr", true, normalizeTagContent("wbr", content)...)
}
func Area(content ...any) HTML {
	return NewTagComponent("area", true, normalizeTagContent("area", content)...)
}
func Audio(content ...any) HTML {
	return NewTagComponent("audio", false, normalizeTagContent("audio", content)...)
}
func Img(content ...any) HTML {
	return NewTagComponent("img", true, normalizeTagContent("img", content)...)
}
func Map(content ...any) HTML {
	return NewTagComponent("map", false, normalizeTagContent("map", content)...)
}
func Track(content ...any) HTML {
	return NewTagComponent("track", true, normalizeTagContent("track", content)...)
}
func Video(content ...any) HTML {
	return NewTagComponent("video", false, normalizeTagContent("video", content)...)
}
func Embed(content ...any) HTML {
	return NewTagComponent("embed", true, normalizeTagContent("embed", content)...)
}
func Fencedframe(content ...any) HTML {
	return NewTagComponent("fencedframe", false, normalizeTagContent("fencedframe", content)...)
}
func Iframe(content ...any) HTML {
	return NewTagComponent("iframe", false, normalizeTagContent("iframe", content)...)
}
func Object(content ...any) HTML {
	return NewTagComponent("object", false, normalizeTagContent("object", content)...)
}
func Picture(content ...any) HTML {
	return NewTagComponent("picture", false, normalizeTagContent("picture", content)...)
}
func Source(content ...any) HTML {
	return NewTagComponent("source", true, normalizeTagContent("source", content)...)
}
func Svg(content ...any) HTML {
	return NewTagComponent("svg", false, normalizeTagContent("svg", content)...)
}
func Math(content ...any) HTML {
	return NewTagComponent("math", false, normalizeTagContent("math", content)...)
}
func Canvas(content ...any) HTML {
	return NewTagComponent("canvas", false, normalizeTagContent("canvas", content)...)
}
func Noscript(content ...any) HTML {
	return NewTagComponent("noscript", false, normalizeTagContent("noscript", content)...)
}
func Script(content ...any) HTML {
	return NewTagComponent("script", false, normalizeTagContent("script", content)...)
}
func Del(content ...any) HTML {
	return NewTagComponent("del", false, normalizeTagContent("del", content)...)
}
func Ins(content ...any) HTML {
	return NewTagComponent("ins", false, normalizeTagContent("ins", content)...)
}
func Caption(content ...any) HTML {
	return NewTagComponent("caption", false, normalizeTagContent("caption", content)...)
}
func Col(content ...any) HTML {
	return NewTagComponent("col", true, normalizeTagContent("col", content)...)
}
func Colgroup(content ...any) HTML {
	return NewTagComponent("colgroup", false, normalizeTagContent("colgroup", content)...)
}
func Table(content ...any) HTML {
	return NewTagComponent("table", false, normalizeTagContent("table", content)...)
}
func Tbody(content ...any) HTML {
	return NewTagComponent("tbody", false, normalizeTagContent("tbody", content)...)
}
func Td(content ...any) HTML {
	return NewTagComponent("td", false, normalizeTagContent("td", content)...)
}
func Tfoot(content ...any) HTML {
	return NewTagComponent("tfoot", false, normalizeTagContent("tfoot", content)...)
}
func Th(content ...any) HTML {
	return NewTagComponent("th", false, normalizeTagContent("th", content)...)
}
func Thead(content ...any) HTML {
	return NewTagComponent("thead", false, normalizeTagContent("thead", content)...)
}
func Tr(content ...any) HTML {
	return NewTagComponent("tr", false, normalizeTagContent("tr", content)...)
}
func Button(content ...any) HTML {
	return NewTagComponent("button", false, normalizeTagContent("button", content)...)
}
func Datalist(content ...any) HTML {
	return NewTagComponent("datalist", false, normalizeTagContent("datalist", content)...)
}
func Fieldset(content ...any) HTML {
	return NewTagComponent("fieldset", false, normalizeTagContent("fieldset", content)...)
}
func Form(content ...any) HTML {
	return NewTagComponent("form", false, normalizeTagContent("form", content)...)
}
func Input(content ...any) HTML {
	return NewTagComponent("input", true, normalizeTagContent("input", content)...)
}
func Label(content ...any) HTML {
	return NewTagComponent("label", false, normalizeTagContent("label", content)...)
}
func Legend(content ...any) HTML {
	return NewTagComponent("legend", false, normalizeTagContent("legend", content)...)
}
func Meter(content ...any) HTML {
	return NewTagComponent("meter", false, normalizeTagContent("meter", content)...)
}
func Optgroup(content ...any) HTML {
	return NewTagComponent("optgroup", false, normalizeTagContent("optgroup", content)...)
}
func Option(content ...any) HTML {
	return NewTagComponent("option", false, normalizeTagContent("option", content)...)
}
func Output(content ...any) HTML {
	return NewTagComponent("output", false, normalizeTagContent("output", content)...)
}
func Progress(content ...any) HTML {
	return NewTagComponent("progress", false, normalizeTagContent("progress", content)...)
}
func Select(content ...any) HTML {
	return NewTagComponent("select", false, normalizeTagContent("select", content)...)
}
func Textarea(content ...any) HTML {
	return NewTagComponent("textarea", false, normalizeTagContent("textarea", content)...)
}
func Details(content ...any) HTML {
	return NewTagComponent("details", false, normalizeTagContent("details", content)...)
}
func Dialog(content ...any) HTML {
	return NewTagComponent("dialog", false, normalizeTagContent("dialog", content)...)
}
func Summary(content ...any) HTML {
	return NewTagComponent("summary", false, normalizeTagContent("summary", content)...)
}
func Slot(content ...any) HTML {
	return NewTagComponent("slot", false, normalizeTagContent("slot", content)...)
}
func Template(content ...any) HTML {
	return NewTagComponent("template", false, normalizeTagContent("template", content)...)
}
