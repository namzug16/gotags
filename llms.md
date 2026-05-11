# gotags — LLM Agent Guide

## Overview

`gotags` is a Go package for building HTML with pure functions. No templates, no runtime reflection — just composable functions that return HTML.

**Core Interface:**

```go
type HTML interface {
    String() string
}
```

Everything in this package implements `HTML`. You call `.String()` on the root to get the final HTML string.

---

## Tags

All HTML tags are functions that accept `...any`. They normalize content automatically:

- **string** → wrapped as `Text` (HTML-escaped automatically)
- **HTML** → rendered as-is
- **[]HTML** → flattened into children
- **anything else** → panics with a clear error message

```go
Div(
    "hello",           // → <span>hello</span> (escaped)
    Span("world"),     // → <span>world</span>
    []HTML{P(), Q()},  // → <p></p><q></q>
)
```

### Void Tags (self-closing)

These never render children even if you pass content (content is silently ignored):

`Base`, `Br`, `Col`, `Embed`, `Hr`, `Img`, `Input`, `Area`, `Track`, `Source`, `Wbr`

```go
Br()           // → <br />
Img(X.Src("cat.png"))  // → <img src="cat.png" />
```

### Full Tag List

**Document:** `Doc`, `Html`, `Head`, `Body`, `Base`, `Link`, `Meta`, `Style`, `Title`, `Script`, `Noscript`

**Sections:** `Address`, `Article`, `Aside`, `Footer`, `Header`, `H1`–`H6`, `Hgroup`, `Main`, `Nav`, `Section`, `Search`

**Text:** `P`, `Span`, `Div`, `Blockquote`, `Pre`, `Code`, `Br`, `Hr`, `Abb`, `B`, `Bdi`, `Bdo`, `Cite`, `Data`, `Dfn`, `Em`, `I`, `Kbd`, `Mark`, `Q`, `Rp`, `Rt`, `Ruby`, `S`, `Samp`, `Small`, `Strong`, `Sub`, `Sup`, `Time`, `U`, `Var`, `Wbr`

**Lists:** `Ul`, `Ol`, `Li`, `Dl`, `Dt`, `Dd`, `Menu`

**Tables:** `Table`, `Caption`, `Col`, `Colgroup`, `Thead`, `Tbody`, `Tfoot`, `Tr`, `Th`, `Td`

**Forms:** `Form`, `Input`, `Label`, `Button`, `Select`, `Option`, `Optgroup`, `Textarea`, `Fieldset`, `Legend`, `Datalist`, `Output`, `Progress`, `Meter`

**Media:** `Img`, `Audio`, `Video`, `Track`, `Source`, `Embed`, `Iframe`, `Picture`, `Map`, `Fencedframe`, `Object`, `Canvas`, `Svg`, `Math`

**Interactive:** `Details`, `Dialog`, `Summary`, `Slot`, `Template`

**Content:** `Figure`, `Figcaption`, `Del`, `Ins`, `A`

**Custom/XML:** `NewTagComponent(name string, isVoid bool, children ...HTML)` — create any tag

```go
// Custom XML-like tag
func RSSFeed(content ...HTML) HTML {
    return NewTagComponent("feed", false, content...)
}
```

---

## Attributes

Access via the `X` object (aliased as `attrs`):

```go
X.Class("btn", "btn-primary")
X.Id("main")
X.Data("user-id", "123")
```

### Available Attributes

`Accept`, `Accept_charset`, `Accesskey`, `Action`, `Align`, `Allow`, `Alt`, `As`, `Async`, `Autocapitalize`, `Autocomplete`, `Autofocus`, `Autoplay`, `Background`, `Border`, `Capture`, `Charset`, `Checked`, `Cite`, `Class`, `Color`, `Cols`, `Colspan`, `Content`, `Contenteditable`, `Controls`, `Coords`, `Crossorigin`, `Csp`, `Data`, `Datetime`, `Decoding`, `Default`, `Defer`, `Dir`, `Dirname`, `Disabled`, `Download`, `Draggable`, `Enctype`, `Enterkeyhint`, `For`, `Form`, `Formaction`, `Formenctype`, `Formmethod`, `Formnovalidate`, `Formtarget`, `Headers`, `Height`, `Hidden`, `High`, `Href`, `Hreflang`, `Http_equiv`, `Id`, `Integrity`, `Intrinsicsize`, `Inputmode`, `Ismap`, `Itemprop`, `Itemid`, `Itemref`, `Itemscope`, `Itemtype`, `Kind`, `Label`, `Lang`, `Language`, `List`, `Loading`, `Loop`, `Low`, `Max`, `Maxlength`, `Min`, `Minlength`, `Media`, `Method`, `Multiple`, `Muted`, `Name`, `Novalidate`, `Open`, `Optimum`, `Part`, `Pattern`, `Ping`, `Placeholder`, `Playsinline`, `Poster`, `Preload`, `Property`, `Readonly`, `Referrerpolicy`, `Rel`, `Required`, `Reversed`, `Role`, `Rows`, `Rowspan`, `Sandbox`, `Scope`, `Scoped`, `Selected`, `Shape`, `Size`, `Sizes`, `Slot`, `Span`, `Spellcheck`, `Src`, `Srcdoc`, `Srclang`, `Srcset`, `Start`, `Step`, `Style`, `Summary`, `Tabindex`, `Target`, `Title`, `Translate`, `Type`, `Usemap`, `Value`, `Virtualkeyboardpolicy`, `Width`, `Wrap`, `Writingsuggestions`, `Anchor`, `Autocorrect`, `Exportparts`, `Inert`, `Is`, `Nonce`, `Popover`

### Custom Attributes

```go
X.Attr("data-custom", "value")
X.Attr("aria-label", "Close dialog")
```

### Conditional Attributes with `X.If`

```go
X.Class(
    "btn",
    X.If(isPrimary, "btn-primary"),
    X.If(!isPrimary, "btn-secondary"),
)
```

`X.If(condition bool, values ...string)` returns the joined values if true, empty string if false. This works inside any attribute call.

---

## Component Helpers

### `Doc(components ...HTML) *DocComponent`

Renders a complete HTML document with DOCTYPE:

```go
Doc(
    Head(Title("Page")),
    Body(Div("content")),
)
// → <!DOCTYPE html><html><head><title>Page</title></head><body><div>content</div></body></html>
```

### `Fragment(components ...any) *FragmentComponent`

Compose multiple root elements without a wrapper tag. Useful for htmx responses:

```go
return Fragment(
    H1("Title"),
    Form(Input(X.Name("email")), Button("Submit")),
)
```

Note: `AttributeComponent` cannot be placed in `Fragment` — it panics if you try.

### `If(condition bool, components ...any) []HTML`

Returns children if true, empty slice if false. For conditional **components**:

```go
Div(
    "text",
    If(isLoggedIn, Span("Welcome!"), Span("Please log in")),
)
```

### `IfLazy(condition bool, component func() HTML) HTML`

Lazy version — the function only runs when condition is true. Use for expensive computations:

```go
Div(
    IfLazy(isExpanded, func() HTML {
        return HeavyComponent()
    }),
)
```

### `Range[T any](items []T, mapper func(int, T) HTML) []HTML`

Transform a slice into a slice of HTML:

```go
users := []string{"Alice", "Bob", "Carol"}

Ul(
    Range(users, func(i int, name string) HTML {
        return Li(X.Class("user"), name)
    }),
)
```

### `AddToTag(target HTML, components ...HTML) HTML`

Mutate a tag in place by appending attributes or children. Returns the same target. Panics if target is not a `*TagComponent`:

```go
card := Div(X.Class("card"))
AddToTag(card,
    H2("Title"),
    P("Content"),
)
// Useful for htmx OOB swap scenarios
```

### `Text(s string) Text`

Explicit text node (auto-escaped). Strings are implicitly wrapped, but use `Text` for clarity or when you need to pass text as a value:

```go
Span(Text("<script>"))  // → <span>&lt;script&gt;</span>
```

### `Raw(s string) Raw`

Unescaped raw HTML. Use with caution:

```go
Span(Raw("<b>bold</b>"))  // → <span><b>bold</b></span>
```

### `NewTagComponent(name string, isVoid bool, children ...HTML) *TagComponent`

Create any tag directly. `isVoid = true` for self-closing tags:

```go
custom := NewTagComponent("my-element", false, "content")
```

---

## htmx Integration

Import the package:

```go
import hx "github.com/namzug16/gotags/htmx"
```

### Attributes

`hx-get`, `hx-post`, `hx-put`, `hx-delete`, `hx-patch`

```go
hx.Get("/users")
hx.Post("/submit")
hx.Patch("/update/1")
hx.Delete("/delete/1")
```

`hx-on:{event}` — inline event handler

```go
hx.On("click", "console.log('clicked')")
hx.On("mouseover", "this.classList.add('hover')")
```

`hx-trigger` — what triggers the request

```go
hx.Trigger("click")
hx.Trigger("every 5s")
hx.Trigger("revealed")  // for lazy loading
```

`hx-target` — CSS selector for target element

```go
hx.Target("#results")
hx.Target("closest .item")
```

`hx-swap` — how content is swapped

```go
hx.Swap("innerHTML")
hx.Swap("outerHTML")
hx.Swap("beforeend")   // append
hx.Swap("afterend")    // prepend
hx.Swap("delete")
hx.Swap("none")
```

`hx-swap-oob` — swap element out of band (to different target)

```go
hx.SwapOob("#sidebar")
```

`hx-select` / `hx-select-oob` — select subset from response

```go
hx.Select("#content")
hx.SelectOob("#toast")
```

`hx-push-url` — push URL to history

```go
hx.PushUrl("true")
```

`hx-replace-url` — replace current URL

```go
hx.ReplaceUrl("/new-path")
```

`hx-boost` — progressive enhancement for links/forms

```go
hx.Boost("true")
```

`hx-confirm` — show confirm dialog before request

```go
hx.Confirm("Are you sure?")
```

`hx-prompt` — show prompt, submit value with request

```go
hx.Prompt("Enter your name:")
```

`hx-vals` — add JSON values to request

```go
hx.Vals(`{"extra": true}`)
```

`hx-headers` — add custom headers

```go
hx.Headers(`{"X-Custom": "value"}`)
```

`hx-params` — filter which params to submit

```go
hx.Params("only", "email", "name")   // only these
hx.Params("none")                    // exclude all
hx.Params("not", "csrf")             // exclude specific
```

`hx-include` — include additional elements in request

```go
hx.Include("#search-input")
```

`hx-ext` — use htmx extensions

```go
hx.Ext("alpine-morph")
```

`hx-disinherit` / `hx-inherit` — control attribute inheritance

```go
hx.Disinherit("hx-get")
hx.Inherit("hx-swap")
```

`hx-history` — control history caching

```go
hx.History("false")
```

`hx-history-elt` — element to snapshot for history

```go
hx.HistoryElt("#main")
```

`hx-indicator` — element to show during request

```go
hx.Indicator(".spinner")
```

`hx-disabled-elt` — disable element during request

```go
hx.DisabledElt("closest button")
```

`hx-sync` — synchronize requests between elements

```go
hx.Sync("closest form")
```

`hx-request` — configure request behavior

```go
hx.Request(`{"timeout": 5000}`)
```

`hx-encoding` — change encoding type

```go
hx.Encoding("multipart/form-data")
```

`hx-validate` — force validation before request

```go
hx.Validate("true")
```

`hx-preserve` — keep element state across swaps

```go
hx.Preserve("#editor")
```

`hx-disable` — disable htmx processing on element

```go
hx.Disable()
```

---

### htmx Events

Available in `htmx` package:

```go
hx.EventAbort
hx.EventAfterOnLoad
hx.EventAfterProcessNode
hx.EventAfterRequest
hx.EventAfterSettle
hx.EventAfterSwap
hx.EventBeforeCleanupElement
hx.EventBeforeOnLoad
hx.EventBeforeProcessNode
hx.EventBeforeRequest
hx.EventBeforeSwap
hx.EventBeforeSend
hx.EventBeforeTransition
hx.EventConfigRequest
hx.EventConfirm
hx.EventHistoryCacheError
hx.EventHistoryCacheHit
hx.EventHistoryCacheMiss
hx.EventHistoryCacheMissLoadError
hx.EventHistoryCacheMissLoad
hx.EventHistoryRestore
hx.EventBeforeHistorySave
hx.EventLoad
hx.EventNoSSESourceError
hx.EventOnLoadError
hx.EventOOBAfterSwap
hx.EventOOBBeforeSwap
hx.EventOOBErrorNoTarget
hx.EventPrompt
hx.EventPushedIntoHistory
hx.EventReplacedInHistory
hx.EventResponseError
hx.EventSendAbort
hx.EventSendError
hx.EventSSEError
hx.EventSSEOpen
hx.EventSwapError
hx.EventTargetError
hx.EventTimeout
hx.EventValidationValidate
hx.EventValidationFailed
hx.EventValidationHalted
hx.EventXHRAbort
hx.EventXHRLoadEnd
hx.EventXHRLoadStart
hx.EventXHRProgress
```

---

### htmx Headers

Use in your Go HTTP handlers to control htmx behavior from the server:

**Request headers** (read from incoming request):
```go
hx.HeaderHXRequest        // "true" if htmx request
hx.HeaderHXBoosted        // "true" if via hx-boost
hx.HeaderHXCurrentURL     // current browser URL
hx.HeaderHXHistoryRestoreRequest  // "true" if history restoration
hx.HeaderHXPrompt         // user response to hx-prompt
hx.HeaderHXTarget         // id of target element
hx.HeaderHXTriggerName    // name of triggered element
hx.HeaderHXTrigger        // id of triggered element
```

**Response headers** (set in your handler):
```go
hx.HeaderHXLocation       // client-side redirect
hx.HeaderHXPushURL        // push new URL to history
hx.HeaderHXRedirect       // client-side redirect
hx.HeaderHXRefresh        // "true" to full page refresh
hx.HeaderHXReplaceURL     // replace current URL
hx.HeaderHXReswap         // override swap method
hx.HeaderHXRetarget       // retarget to different element
hx.HeaderHXReselect       // reselect content from response
hx.HeaderHXTrigger        // trigger client-side events
hx.HeaderHXTriggerAfterSettle   // trigger after settle
hx.HeaderHXTriggerAfterSwap     // trigger after swap
```

---

## Common Patterns

### Page Layout

```go
Doc(
    Head(
        Title("My Page"),
        Link(X.Rel("stylesheet"), X.Href("/styles.css")),
    ),
    Body(
        Header(Nav(UL(
            LI(A(X.Href("/"), "Home")),
            LI(A(X.Href("/about"), "About")),
        ))),
        Main(Section(H1("Welcome"))),
        Footer(P("© 2026")),
    ),
)
```

### Conditional Classes

```go
Button(
    X.Class(
        "btn",
        X.If(isActive, "btn-active"),
        X.If(!isActive, "btn-inactive"),
    ),
    "Click me",
)
```

### List Rendering

```go
items := []string{"a", "b", "c"}
UL(
    Range(items, func(i int, s string) HTML {
        return LI(X.Data("index", fmt.Sprintf("%d", i)), s)
    }),
)
```

### htmx Form

```go
Form(
    X.Action("/submit"),
    X.Method("POST"),
    Input(X.Name("email"), X.Placeholder("email@example.com")),
    Button(
        hx.Post("/submit"),
        hx.Trigger("click"),
        hx.Swap("outerHTML"),
        "Submit",
    ),
)
```

### htmx Response with Multiple Elements

```go
return Fragment(
    H2("Results"),
    UL(
        Range(items, func(_ int, item string) HTML {
            return LI(item)
        }),
    ),
)
```

### Lazy Conditional Content

```go
Div(
    IfLazy(showDetails, func() HTML {
        return HeavyListComponent()
    }),
)
```

### Mutate Tag In-Place (htmx OOB)

```go
card := Div(X.Id("card-1"), X.Class("card"))
AddToTag(card,
    H2("Updated Title"),  // children go to tag
    HxSwapOob("#sidebar"), // attributes also work
)
```