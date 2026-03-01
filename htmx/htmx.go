package htmx

import (
	gt "github.com/namzug16/gotags"
)

// Issues a GET to the specified URL
func Get(value string) gt.HTML {
	return gt.X.Attr("hx-get", value)
}

// Issues a POST to the specified URL
func Post(value string) gt.HTML {
	return gt.X.Attr("hx-post", value)
}

// Handle events with inline scripts on elements
func On(event, value string) gt.HTML {
	return gt.X.Attr("hx-on:"+event, value)
}

// Push a URL into the browser location bar to create history
func PushUrl(value string) gt.HTML {
	return gt.X.Attr("hx-push-url", value)
}

// Select content to swap in from a response
func Select(value string) gt.HTML {
	return gt.X.Attr("hx-select", value)
}

// Select content to swap in from a response, somewhere other than the target (out of band)
func SelectOob(value string) gt.HTML {
	return gt.X.Attr("hx-select-oob", value)
}

// Controls how content will swap in (outerHTML, beforeend, afterend, …)
func Swap(value string) gt.HTML {
	return gt.X.Attr("hx-swap", value)
}

// Mark element to swap in from a response (out of band)
func SwapOob(value string) gt.HTML {
	return gt.X.Attr("hx-swap-oob", value)
}

// Specifies the target element to be swapped
func Target(value string) gt.HTML {
	return gt.X.Attr("hx-target", value)
}

// Specifies the event that triggers the request
func Trigger(value string) gt.HTML {
	return gt.X.Attr("hx-trigger", value)
}

// Add values to submit with the request (JSON format)
func Vals(value string) gt.HTML {
	return gt.X.Attr("hx-vals", value)
}

// Add progressive enhancement for links and forms
func Boost(value string) gt.HTML {
	return gt.X.Attr("hx-boost", value)
}

// Shows a confirm() dialog before issuing a request
func Confirm(value string) gt.HTML {
	return gt.X.Attr("hx-confirm", value)
}

// Issues a DELETE to the specified URL
func Delete(value string) gt.HTML {
	return gt.X.Attr("hx-delete", value)
}

// Disables htmx processing for the given node and any children nodes
func Disable() gt.HTML {
	return gt.X.Attr("hx-disable")
}

// Adds the disabled attribute to the specified elements while a request is in flight
func DisabledElt(value string) gt.HTML {
	return gt.X.Attr("hx-disabled-elt", value)
}

// Control and disable automatic attribute inheritance for child nodes
func Disinherit(value string) gt.HTML {
	return gt.X.Attr("hx-disinherit", value)
}

// Changes the request encoding type
func Encoding(value string) gt.HTML {
	return gt.X.Attr("hx-encoding", value)
}

// Extensions to use for this element
func Ext(value string) gt.HTML {
	return gt.X.Attr("hx-ext", value)
}

// Adds to the headers that will be submitted with the request
func Headers(value string) gt.HTML {
	return gt.X.Attr("hx-headers", value)
}

// Prevent sensitive data being saved to the history cache
func History(value string) gt.HTML {
	return gt.X.Attr("hx-history", value)
}

// The element to snapshot and restore during history navigation
func HistoryElt(value string) gt.HTML {
	return gt.X.Attr("hx-history-elt", value)
}

// Include additional data in requests
func Include(value string) gt.HTML {
	return gt.X.Attr("hx-include", value)
}

// The element to put the htmx-request class on during the request
func Indicator(value string) gt.HTML {
	return gt.X.Attr("hx-indicator", value)
}

// Control and enable automatic attribute inheritance for child nodes if it has been disabled by default
func Inherit(value string) gt.HTML {
	return gt.X.Attr("hx-inherit", value)
}

// Filters the parameters that will be submitted with a request
func Params(value string) gt.HTML {
	return gt.X.Attr("hx-params", value)
}

// Issues a PATCH to the specified URL
func Patch(value string) gt.HTML {
	return gt.X.Attr("hx-patch", value)
}

// Specifies elements to keep unchanged between requests
func Preserve(value string) gt.HTML {
	return gt.X.Attr("hx-preserve", value)
}

// Shows a prompt() before submitting a request
func Prompt(value string) gt.HTML {
	return gt.X.Attr("hx-prompt", value)
}

// Issues a PUT to the specified URL
func Put(value string) gt.HTML {
	return gt.X.Attr("hx-put", value)
}

// Replace the URL in the browser location bar
func ReplaceUrl(value string) gt.HTML {
	return gt.X.Attr("hx-replace-url", value)
}

// Configures various aspects of the request
func Request(value string) gt.HTML {
	return gt.X.Attr("hx-request", value)
}

// Control how requests made by different elements are synchronized
func Sync(value string) gt.HTML {
	return gt.X.Attr("hx-sync", value)
}

// Force elements to validate themselves before a request
func Validate(value string) gt.HTML {
	return gt.X.Attr("hx-validate", value)
}
