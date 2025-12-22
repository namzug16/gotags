package gotags

import "strings"

type attrs struct{}

// Defines attributes
var X = attrs{}

func (attrs) Attr(name string, values ...string) HTML {
	return &AttributeComponent{
		name:   name,
		values: values,
	}
}

func (attrs) If(condition bool, values ...string) string {
	if condition {
		return strings.Join(values, " ")
	}
	return ""
}

func (attrs) Accept(values ...string) HTML {
	return X.Attr("accept", values...)
}
func (attrs) Accept_charset(values ...string) HTML {
	return X.Attr("accept-charset", values...)
}
func (attrs) Accesskey(values ...string) HTML {
	return X.Attr("accesskey", values...)
}
func (attrs) Action(values ...string) HTML {
	return X.Attr("action", values...)
}
func (attrs) Align(values ...string) HTML {
	return X.Attr("align", values...)
}
func (attrs) Allow(values ...string) HTML {
	return X.Attr("allow", values...)
}
func (attrs) Alt(values ...string) HTML {
	return X.Attr("alt", values...)
}
func (attrs) As(values ...string) HTML {
	return X.Attr("as", values...)
}
func (attrs) Async(values ...string) HTML {
	return X.Attr("async", values...)
}
func (attrs) Autocapitalize(values ...string) HTML {
	return X.Attr("autocapitalize", values...)
}
func (attrs) Autocomplete(values ...string) HTML {
	return X.Attr("autocomplete", values...)
}
func (attrs) Autoplay(values ...string) HTML {
	return X.Attr("autoplay", values...)
}
func (attrs) Background(values ...string) HTML {
	return X.Attr("background", values...)
}
func (attrs) Border(values ...string) HTML {
	return X.Attr("border", values...)
}
func (attrs) Capture(values ...string) HTML {
	return X.Attr("capture", values...)
}
func (attrs) Charset(values ...string) HTML {
	return X.Attr("charset", values...)
}
func (attrs) Checked(values ...string) HTML {
	return X.Attr("checked", values...)
}
func (attrs) Cite(values ...string) HTML {
	return X.Attr("cite", values...)
}
func (attrs) Class(values ...string) HTML {
	return X.Attr("class", values...)
}
func (attrs) Color(values ...string) HTML {
	return X.Attr("color", values...)
}
func (attrs) Cols(values ...string) HTML {
	return X.Attr("cols", values...)
}
func (attrs) Colspan(values ...string) HTML {
	return X.Attr("colspan", values...)
}
func (attrs) Content(values ...string) HTML {
	return X.Attr("content", values...)
}
func (attrs) Contenteditable(values ...string) HTML {
	return X.Attr("contenteditable", values...)
}
func (attrs) Controls(values ...string) HTML {
	return X.Attr("controls", values...)
}
func (attrs) Coords(values ...string) HTML {
	return X.Attr("coords", values...)
}
func (attrs) Crossorigin(values ...string) HTML {
	return X.Attr("crossorigin", values...)
}
func (attrs) Csp(values ...string) HTML {
	return X.Attr("csp", values...)
}
func (attrs) Data(values ...string) HTML {
	return X.Attr("data", values...)
}
func (attrs) Datetime(values ...string) HTML {
	return X.Attr("datetime", values...)
}
func (attrs) Decoding(values ...string) HTML {
	return X.Attr("decoding", values...)
}
func (attrs) Default(values ...string) HTML {
	return X.Attr("default", values...)
}
func (attrs) Defer(values ...string) HTML {
	return X.Attr("defer", values...)
}
func (attrs) Dir(values ...string) HTML {
	return X.Attr("dir", values...)
}
func (attrs) Dirname(values ...string) HTML {
	return X.Attr("dirname", values...)
}
func (attrs) Disabled(values ...string) HTML {
	return X.Attr("disabled", values...)
}
func (attrs) Download(values ...string) HTML {
	return X.Attr("download", values...)
}
func (attrs) Draggable(values ...string) HTML {
	return X.Attr("draggable", values...)
}
func (attrs) Enctype(values ...string) HTML {
	return X.Attr("enctype", values...)
}
func (attrs) Enterkeyhint(values ...string) HTML {
	return X.Attr("enterkeyhint", values...)
}
func (attrs) For(values ...string) HTML {
	return X.Attr("for", values...)
}
func (attrs) Form(values ...string) HTML {
	return X.Attr("form", values...)
}
func (attrs) Formaction(values ...string) HTML {
	return X.Attr("formaction", values...)
}
func (attrs) Formenctype(values ...string) HTML {
	return X.Attr("formenctype", values...)
}
func (attrs) Formmethod(values ...string) HTML {
	return X.Attr("formmethod", values...)
}
func (attrs) Formnovalidate(values ...string) HTML {
	return X.Attr("formnovalidate", values...)
}
func (attrs) Formtarget(values ...string) HTML {
	return X.Attr("formtarget", values...)
}
func (attrs) Headers(values ...string) HTML {
	return X.Attr("headers", values...)
}
func (attrs) Height(values ...string) HTML {
	return X.Attr("height", values...)
}
func (attrs) Hidden(values ...string) HTML {
	return X.Attr("hidden", values...)
}
func (attrs) High(values ...string) HTML {
	return X.Attr("high", values...)
}
func (attrs) Href(values ...string) HTML {
	return X.Attr("href", values...)
}
func (attrs) Hreflang(values ...string) HTML {
	return X.Attr("hreflang", values...)
}
func (attrs) Http_equiv(values ...string) HTML {
	return X.Attr("http-equiv", values...)
}
func (attrs) Id(values ...string) HTML {
	return X.Attr("id", values...)
}
func (attrs) Integrity(values ...string) HTML {
	return X.Attr("integrity", values...)
}
func (attrs) Intrinsicsize(values ...string) HTML {
	return X.Attr("intrinsicsize", values...)
}
func (attrs) Inputmode(values ...string) HTML {
	return X.Attr("inputmode", values...)
}
func (attrs) Ismap(values ...string) HTML {
	return X.Attr("ismap", values...)
}
func (attrs) Itemprop(values ...string) HTML {
	return X.Attr("itemprop", values...)
}
func (attrs) Kind(values ...string) HTML {
	return X.Attr("kind", values...)
}
func (attrs) Label(values ...string) HTML {
	return X.Attr("label", values...)
}
func (attrs) Lang(values ...string) HTML {
	return X.Attr("lang", values...)
}
func (attrs) Language(values ...string) HTML {
	return X.Attr("language", values...)
}
func (attrs) Loading(values ...string) HTML {
	return X.Attr("loading", values...)
}
func (attrs) List(values ...string) HTML {
	return X.Attr("list", values...)
}
func (attrs) Loop(values ...string) HTML {
	return X.Attr("loop", values...)
}
func (attrs) Low(values ...string) HTML {
	return X.Attr("low", values...)
}
func (attrs) Max(values ...string) HTML {
	return X.Attr("max", values...)
}
func (attrs) Maxlength(values ...string) HTML {
	return X.Attr("maxlength", values...)
}
func (attrs) Minlength(values ...string) HTML {
	return X.Attr("minlength", values...)
}
func (attrs) Media(values ...string) HTML {
	return X.Attr("media", values...)
}
func (attrs) Method(values ...string) HTML {
	return X.Attr("method", values...)
}
func (attrs) Min(values ...string) HTML {
	return X.Attr("min", values...)
}
func (attrs) Multiple(values ...string) HTML {
	return X.Attr("multiple", values...)
}
func (attrs) Muted(values ...string) HTML {
	return X.Attr("muted", values...)
}
func (attrs) Name(values ...string) HTML {
	return X.Attr("name", values...)
}
func (attrs) Novalidate(values ...string) HTML {
	return X.Attr("novalidate", values...)
}
func (attrs) Open(values ...string) HTML {
	return X.Attr("open", values...)
}
func (attrs) Optimum(values ...string) HTML {
	return X.Attr("optimum", values...)
}
func (attrs) Pattern(values ...string) HTML {
	return X.Attr("pattern", values...)
}
func (attrs) Ping(values ...string) HTML {
	return X.Attr("ping", values...)
}
func (attrs) Placeholder(values ...string) HTML {
	return X.Attr("placeholder", values...)
}
func (attrs) Playsinline(values ...string) HTML {
	return X.Attr("playsinline", values...)
}
func (attrs) Poster(values ...string) HTML {
	return X.Attr("poster", values...)
}
func (attrs) Preload(values ...string) HTML {
	return X.Attr("preload", values...)
}
func (attrs) Readonly(values ...string) HTML {
	return X.Attr("readonly", values...)
}
func (attrs) Referrerpolicy(values ...string) HTML {
	return X.Attr("referrerpolicy", values...)
}
func (attrs) Rel(values ...string) HTML {
	return X.Attr("rel", values...)
}
func (attrs) Required(values ...string) HTML {
	return X.Attr("required", values...)
}
func (attrs) Reversed(values ...string) HTML {
	return X.Attr("reversed", values...)
}
func (attrs) Role(values ...string) HTML {
	return X.Attr("role", values...)
}
func (attrs) Rows(values ...string) HTML {
	return X.Attr("rows", values...)
}
func (attrs) Rowspan(values ...string) HTML {
	return X.Attr("rowspan", values...)
}
func (attrs) Sandbox(values ...string) HTML {
	return X.Attr("sandbox", values...)
}
func (attrs) Scope(values ...string) HTML {
	return X.Attr("scope", values...)
}
func (attrs) Scoped(values ...string) HTML {
	return X.Attr("scoped", values...)
}
func (attrs) Selected(values ...string) HTML {
	return X.Attr("selected", values...)
}
func (attrs) Shape(values ...string) HTML {
	return X.Attr("shape", values...)
}
func (attrs) Size(values ...string) HTML {
	return X.Attr("size", values...)
}
func (attrs) Sizes(values ...string) HTML {
	return X.Attr("sizes", values...)
}
func (attrs) Slot(values ...string) HTML {
	return X.Attr("slot", values...)
}
func (attrs) Span(values ...string) HTML {
	return X.Attr("span", values...)
}
func (attrs) Spellcheck(values ...string) HTML {
	return X.Attr("spellcheck", values...)
}
func (attrs) Src(values ...string) HTML {
	return X.Attr("src", values...)
}
func (attrs) Srcdoc(values ...string) HTML {
	return X.Attr("srcdoc", values...)
}
func (attrs) Srclang(values ...string) HTML {
	return X.Attr("srclang", values...)
}
func (attrs) Srcset(values ...string) HTML {
	return X.Attr("srcset", values...)
}
func (attrs) Start(values ...string) HTML {
	return X.Attr("start", values...)
}
func (attrs) Step(values ...string) HTML {
	return X.Attr("step", values...)
}
func (attrs) Style(values ...string) HTML {
	return X.Attr("style", values...)
}
func (attrs) Summary(values ...string) HTML {
	return X.Attr("summary", values...)
}
func (attrs) Tabindex(values ...string) HTML {
	return X.Attr("tabindex", values...)
}
func (attrs) Target(values ...string) HTML {
	return X.Attr("target", values...)
}
func (attrs) Title(values ...string) HTML {
	return X.Attr("title", values...)
}
func (attrs) Translate(values ...string) HTML {
	return X.Attr("translate", values...)
}
func (attrs) Type(values ...string) HTML {
	return X.Attr("type", values...)
}
func (attrs) Usemap(values ...string) HTML {
	return X.Attr("usemap", values...)
}
func (attrs) Value(values ...string) HTML {
	return X.Attr("value", values...)
}
func (attrs) Width(values ...string) HTML {
	return X.Attr("width", values...)
}
func (attrs) Wrap(values ...string) HTML {
	return X.Attr("wrap", values...)
}
func (attrs) Anchor(values ...string) HTML {
	return X.Attr("anchor", values...)
}
func (attrs) Autocorrect(values ...string) HTML {
	return X.Attr("autocorrect", values...)
}
func (attrs) Autofocus(values ...string) HTML {
	return X.Attr("autofocus", values...)
}
func (attrs) Exportparts(values ...string) HTML {
	return X.Attr("exportparts", values...)
}
func (attrs) Inert(values ...string) HTML {
	return X.Attr("inert", values...)
}
func (attrs) Is(values ...string) HTML {
	return X.Attr("is", values...)
}
func (attrs) Itemid(values ...string) HTML {
	return X.Attr("itemid", values...)
}
func (attrs) Itemref(values ...string) HTML {
	return X.Attr("itemref", values...)
}
func (attrs) Itemscope(values ...string) HTML {
	return X.Attr("itemscope", values...)
}
func (attrs) Itemtype(values ...string) HTML {
	return X.Attr("itemtype", values...)
}
func (attrs) Nonce(values ...string) HTML {
	return X.Attr("nonce", values...)
}
func (attrs) Part(values ...string) HTML {
	return X.Attr("part", values...)
}
func (attrs) Popover(values ...string) HTML {
	return X.Attr("popover", values...)
}
func (attrs) Virtualkeyboardpolicy(values ...string) HTML {
	return X.Attr("virtualkeyboardpolicy", values...)
}
func (attrs) Writingsuggestions(values ...string) HTML {
	return X.Attr("writingsuggestions", values...)
}
func (attrs) Property(values ...string) HTML {
	return X.Attr("property", values...)
}
