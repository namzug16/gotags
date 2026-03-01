package gotags

import "fmt"

func normalizeContent(tag string, content []any) []HTML {
	res := []HTML{}
	for _, c := range content {
		switch v := c.(type) {
		case HTML:
			res = append(res, v)
		case []HTML:
			res = append(res, v...)
		case string:
			res = append(res, Text(v))
		default:
			panic(fmt.Sprintf(
				"gotags: invalid child type %T in <%s>; expected string, HTML or []HTML",
				c,
				tag,
			))
		}
	}
	return res
}

func Base(content ...any) HTML {
	return NewTagComponent("base", true, normalizeContent("base", content)...)
}
func Head(content ...any) HTML {
	return NewTagComponent("head", false, normalizeContent("head", content)...)
}
func Link(content ...any) HTML {
	return NewTagComponent("link", true, normalizeContent("link", content)...)
}
func Meta(content ...any) HTML {
	return NewTagComponent("meta", true, normalizeContent("meta", content)...)
}
func Style(content ...any) HTML {
	return NewTagComponent("style", false, normalizeContent("style", content)...)
}
func Title(content ...any) HTML {
	return NewTagComponent("title", false, normalizeContent("title", content)...)
}
func Body(content ...any) HTML {
	return NewTagComponent("body", false, normalizeContent("body", content)...)
}
func Address(content ...any) HTML {
	return NewTagComponent("address", false, normalizeContent("address", content)...)
}
func Article(content ...any) HTML {
	return NewTagComponent("article", false, normalizeContent("article", content)...)
}
func Aside(content ...any) HTML {
	return NewTagComponent("aside", false, normalizeContent("aside", content)...)
}
func Footer(content ...any) HTML {
	return NewTagComponent("footer", false, normalizeContent("footer", content)...)
}
func Header(content ...any) HTML {
	return NewTagComponent("header", false, normalizeContent("header", content)...)
}
func H1(content ...any) HTML {
	return NewTagComponent("h1", false, normalizeContent("h1", content)...)
}
func H2(content ...any) HTML {
	return NewTagComponent("h2", false, normalizeContent("h2", content)...)
}
func H3(content ...any) HTML {
	return NewTagComponent("h3", false, normalizeContent("h3", content)...)
}
func H4(content ...any) HTML {
	return NewTagComponent("h4", false, normalizeContent("h4", content)...)
}
func H5(content ...any) HTML {
	return NewTagComponent("h5", false, normalizeContent("h5", content)...)
}
func H6(content ...any) HTML {
	return NewTagComponent("h6", false, normalizeContent("h6", content)...)
}
func Hgroup(content ...any) HTML {
	return NewTagComponent("hgroup", false, normalizeContent("hgroup", content)...)
}
func Main(content ...any) HTML {
	return NewTagComponent("main", false, normalizeContent("main", content)...)
}
func Nav(content ...any) HTML {
	return NewTagComponent("nav", false, normalizeContent("nav", content)...)
}
func Section(content ...any) HTML {
	return NewTagComponent("section", false, normalizeContent("section", content)...)
}
func Search(content ...any) HTML {
	return NewTagComponent("search", false, normalizeContent("search", content)...)
}
func Blockquote(content ...any) HTML {
	return NewTagComponent("blockquote", false, normalizeContent("blockquote", content)...)
}
func Dd(content ...any) HTML {
	return NewTagComponent("dd", false, normalizeContent("dd", content)...)
}
func Div(content ...any) HTML {
	return NewTagComponent("div", false, normalizeContent("div", content)...)
}
func Dl(content ...any) HTML {
	return NewTagComponent("dl", false, normalizeContent("dl", content)...)
}
func Dt(content ...any) HTML {
	return NewTagComponent("dt", false, normalizeContent("dt", content)...)
}
func Figcaption(content ...any) HTML {
	return NewTagComponent("figcaption", false, normalizeContent("figcaption", content)...)
}
func Figure(content ...any) HTML {
	return NewTagComponent("figure", false, normalizeContent("figure", content)...)
}
func Hr(content ...any) HTML {
	return NewTagComponent("hr", true, normalizeContent("hr", content)...)
}
func Li(content ...any) HTML {
	return NewTagComponent("li", false, normalizeContent("li", content)...)
}
func Menu(content ...any) HTML {
	return NewTagComponent("menu", false, normalizeContent("menu", content)...)
}
func Ol(content ...any) HTML {
	return NewTagComponent("ol", false, normalizeContent("ol", content)...)
}
func P(content ...any) HTML {
	return NewTagComponent("p", false, normalizeContent("p", content)...)
}
func Pre(content ...any) HTML {
	return NewTagComponent("pre", false, normalizeContent("pre", content)...)
}
func Ul(content ...any) HTML {
	return NewTagComponent("ul", false, normalizeContent("ul", content)...)
}
func A(content ...any) HTML {
	return NewTagComponent("a", false, normalizeContent("a", content)...)
}
func Abbr(content ...any) HTML {
	return NewTagComponent("abbr", false, normalizeContent("abbr", content)...)
}
func B(content ...any) HTML {
	return NewTagComponent("b", false, normalizeContent("b", content)...)
}
func Bdi(content ...any) HTML {
	return NewTagComponent("bdi", false, normalizeContent("bdi", content)...)
}
func Bdo(content ...any) HTML {
	return NewTagComponent("bdo", false, normalizeContent("bdo", content)...)
}
func Br(content ...any) HTML {
	return NewTagComponent("br", true, normalizeContent("br", content)...)
}
func Cite(content ...any) HTML {
	return NewTagComponent("cite", false, normalizeContent("cite", content)...)
}
func Code(content ...any) HTML {
	return NewTagComponent("code", false, normalizeContent("code", content)...)
}
func Data(content ...any) HTML {
	return NewTagComponent("data", false, normalizeContent("data", content)...)
}
func Dfn(content ...any) HTML {
	return NewTagComponent("dfn", false, normalizeContent("dfn", content)...)
}
func Em(content ...any) HTML {
	return NewTagComponent("em", false, normalizeContent("em", content)...)
}
func I(content ...any) HTML {
	return NewTagComponent("i", false, normalizeContent("i", content)...)
}
func Kbd(content ...any) HTML {
	return NewTagComponent("kbd", false, normalizeContent("kbd", content)...)
}
func Mark(content ...any) HTML {
	return NewTagComponent("mark", false, normalizeContent("mark", content)...)
}
func Q(content ...any) HTML {
	return NewTagComponent("q", false, normalizeContent("q", content)...)
}
func Rp(content ...any) HTML {
	return NewTagComponent("rp", false, normalizeContent("rp", content)...)
}
func Rt(content ...any) HTML {
	return NewTagComponent("rt", false, normalizeContent("rt", content)...)
}
func Ruby(content ...any) HTML {
	return NewTagComponent("ruby", false, normalizeContent("ruby", content)...)
}
func S(content ...any) HTML {
	return NewTagComponent("s", false, normalizeContent("s", content)...)
}
func Samp(content ...any) HTML {
	return NewTagComponent("samp", false, normalizeContent("samp", content)...)
}
func Small(content ...any) HTML {
	return NewTagComponent("small", false, normalizeContent("small", content)...)
}
func Span(content ...any) HTML {
	return NewTagComponent("span", false, normalizeContent("span", content)...)
}
func Strong(content ...any) HTML {
	return NewTagComponent("strong", false, normalizeContent("strong", content)...)
}
func Sub(content ...any) HTML {
	return NewTagComponent("sub", false, normalizeContent("sub", content)...)
}
func Sup(content ...any) HTML {
	return NewTagComponent("sup", false, normalizeContent("sup", content)...)
}
func Time(content ...any) HTML {
	return NewTagComponent("time", false, normalizeContent("time", content)...)
}
func U(content ...any) HTML {
	return NewTagComponent("u", false, normalizeContent("u", content)...)
}
func Var(content ...any) HTML {
	return NewTagComponent("var", false, normalizeContent("var", content)...)
}
func Wbr(content ...any) HTML {
	return NewTagComponent("wbr", true, normalizeContent("wbr", content)...)
}
func Area(content ...any) HTML {
	return NewTagComponent("area", true, normalizeContent("area", content)...)
}
func Audio(content ...any) HTML {
	return NewTagComponent("audio", false, normalizeContent("audio", content)...)
}
func Img(content ...any) HTML {
	return NewTagComponent("img", true, normalizeContent("img", content)...)
}
func Map(content ...any) HTML {
	return NewTagComponent("map", false, normalizeContent("map", content)...)
}
func Track(content ...any) HTML {
	return NewTagComponent("track", true, normalizeContent("track", content)...)
}
func Video(content ...any) HTML {
	return NewTagComponent("video", false, normalizeContent("video", content)...)
}
func Embed(content ...any) HTML {
	return NewTagComponent("embed", true, normalizeContent("embed", content)...)
}
func Fencedframe(content ...any) HTML {
	return NewTagComponent("fencedframe", false, normalizeContent("fencedframe", content)...)
}
func Iframe(content ...any) HTML {
	return NewTagComponent("iframe", false, normalizeContent("iframe", content)...)
}
func Object(content ...any) HTML {
	return NewTagComponent("object", false, normalizeContent("object", content)...)
}
func Picture(content ...any) HTML {
	return NewTagComponent("picture", false, normalizeContent("picture", content)...)
}
func Source(content ...any) HTML {
	return NewTagComponent("source", true, normalizeContent("source", content)...)
}
func Svg(content ...any) HTML {
	return NewTagComponent("svg", false, normalizeContent("svg", content)...)
}
func Math(content ...any) HTML {
	return NewTagComponent("math", false, normalizeContent("math", content)...)
}
func Canvas(content ...any) HTML {
	return NewTagComponent("canvas", false, normalizeContent("canvas", content)...)
}
func Noscript(content ...any) HTML {
	return NewTagComponent("noscript", false, normalizeContent("noscript", content)...)
}
func Script(content ...any) HTML {
	return NewTagComponent("script", false, normalizeContent("script", content)...)
}
func Del(content ...any) HTML {
	return NewTagComponent("del", false, normalizeContent("del", content)...)
}
func Ins(content ...any) HTML {
	return NewTagComponent("ins", false, normalizeContent("ins", content)...)
}
func Caption(content ...any) HTML {
	return NewTagComponent("caption", false, normalizeContent("caption", content)...)
}
func Col(content ...any) HTML {
	return NewTagComponent("col", true, normalizeContent("col", content)...)
}
func Colgroup(content ...any) HTML {
	return NewTagComponent("colgroup", false, normalizeContent("colgroup", content)...)
}
func Table(content ...any) HTML {
	return NewTagComponent("table", false, normalizeContent("table", content)...)
}
func Tbody(content ...any) HTML {
	return NewTagComponent("tbody", false, normalizeContent("tbody", content)...)
}
func Td(content ...any) HTML {
	return NewTagComponent("td", false, normalizeContent("td", content)...)
}
func Tfoot(content ...any) HTML {
	return NewTagComponent("tfoot", false, normalizeContent("tfoot", content)...)
}
func Th(content ...any) HTML {
	return NewTagComponent("th", false, normalizeContent("th", content)...)
}
func Thead(content ...any) HTML {
	return NewTagComponent("thead", false, normalizeContent("thead", content)...)
}
func Tr(content ...any) HTML {
	return NewTagComponent("tr", false, normalizeContent("tr", content)...)
}
func Button(content ...any) HTML {
	return NewTagComponent("button", false, normalizeContent("button", content)...)
}
func Datalist(content ...any) HTML {
	return NewTagComponent("datalist", false, normalizeContent("datalist", content)...)
}
func Fieldset(content ...any) HTML {
	return NewTagComponent("fieldset", false, normalizeContent("fieldset", content)...)
}
func Form(content ...any) HTML {
	return NewTagComponent("form", false, normalizeContent("form", content)...)
}
func Input(content ...any) HTML {
	return NewTagComponent("input", true, normalizeContent("input", content)...)
}
func Label(content ...any) HTML {
	return NewTagComponent("label", false, normalizeContent("label", content)...)
}
func Legend(content ...any) HTML {
	return NewTagComponent("legend", false, normalizeContent("legend", content)...)
}
func Meter(content ...any) HTML {
	return NewTagComponent("meter", false, normalizeContent("meter", content)...)
}
func Optgroup(content ...any) HTML {
	return NewTagComponent("optgroup", false, normalizeContent("optgroup", content)...)
}
func Option(content ...any) HTML {
	return NewTagComponent("option", false, normalizeContent("option", content)...)
}
func Output(content ...any) HTML {
	return NewTagComponent("output", false, normalizeContent("output", content)...)
}
func Progress(content ...any) HTML {
	return NewTagComponent("progress", false, normalizeContent("progress", content)...)
}
func Select(content ...any) HTML {
	return NewTagComponent("select", false, normalizeContent("select", content)...)
}
func Textarea(content ...any) HTML {
	return NewTagComponent("textarea", false, normalizeContent("textarea", content)...)
}
func Details(content ...any) HTML {
	return NewTagComponent("details", false, normalizeContent("details", content)...)
}
func Dialog(content ...any) HTML {
	return NewTagComponent("dialog", false, normalizeContent("dialog", content)...)
}
func Summary(content ...any) HTML {
	return NewTagComponent("summary", false, normalizeContent("summary", content)...)
}
func Slot(content ...any) HTML {
	return NewTagComponent("slot", false, normalizeContent("slot", content)...)
}
func Template(content ...any) HTML {
	return NewTagComponent("template", false, normalizeContent("template", content)...)
}
