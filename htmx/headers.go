package htmx

const (
	// HeaderHXBoosted indicates that the request is via an element using hx-boost.
	HeaderHXBoosted = "HX-Boosted"

	// HeaderHXCurrentURL is the current URL of the browser.
	HeaderHXCurrentURL = "HX-Current-URL"

	// HeaderHXHistoryRestoreRequest is “true” if the request is for history
	// restoration after a miss in the local history cache.
	HeaderHXHistoryRestoreRequest = "HX-History-Restore-Request"

	// HeaderHXPrompt is the user response to an hx-prompt.
	HeaderHXPrompt = "HX-Prompt"

	// HeaderHXRequest is always “true”.
	HeaderHXRequest = "HX-Request"

	// HeaderHXTarget is the id of the target element if it exists.
	HeaderHXTarget = "HX-Target"

	// HeaderHXTriggerName is the name of the triggered element if it exists.
	HeaderHXTriggerName = "HX-Trigger-Name"

	// [REQUEST] HeaderHXTrigger is the id of the triggered element if it exists.
	// [RESPONSE] HeaderHXTrigger allows you to trigger client-side events.
	HeaderHXTrigger = "HX-Trigger"

	// HeaderHXLocation allows you to do a client-side redirect that does not do
	// a full page reload.
	HeaderHXLocation = "HX-Location"

	// HeaderHXPushURL pushes a new url into the history stack.
	HeaderHXPushURL = "HX-Push-Url"

	// HeaderHXRedirect can be used to do a client-side redirect to a new location.
	HeaderHXRedirect = "HX-Redirect"

	// HeaderHXRefresh if set to “true” the client-side will do a full refresh
	// of the page.
	HeaderHXRefresh = "HX-Refresh"

	// HeaderHXReplaceURL replaces the current URL in the location bar.
	HeaderHXReplaceURL = "HX-Replace-Url"

	// HeaderHXReswap allows you to specify how the response will be swapped.
	// See hx-swap for possible values.
	HeaderHXReswap = "HX-Reswap"

	// HeaderHXRetarget a CSS selector that updates the target of the content
	// update to a different element on the page.
	HeaderHXRetarget = "HX-Retarget"

	// HeaderHXReselect a CSS selector that allows you to choose which part of
	// the response is used to be swapped in. Overrides an existing hx-select
	// on the triggering element.
	HeaderHXReselect = "HX-Reselect"

	// HeaderHXTriggerAfterSettle allows you to trigger client-side events after
	// the settle step.
	HeaderHXTriggerAfterSettle = "HX-Trigger-After-Settle"

	// HeaderHXTriggerAfterSwap allows you to trigger client-side events after
	// the swap step.
	HeaderHXTriggerAfterSwap = "HX-Trigger-After-Swap"
)
