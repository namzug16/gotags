package htmx

const (
	// EventAbort send this event to an element to abort a request.
	EventAbort = "htmx:abort"

	// EventAfterOnLoad triggered after an AJAX request has completed processing a successful response.
	EventAfterOnLoad = "htmx:afterOnLoad"

	// EventAfterProcessNode triggered after htmx has initialized a node.
	EventAfterProcessNode = "htmx:afterProcessNode"

	// EventAfterRequest triggered after an AJAX request has completed.
	EventAfterRequest = "htmx:afterRequest"

	// EventAfterSettle triggered after the DOM has settled.
	EventAfterSettle = "htmx:afterSettle"

	// EventAfterSwap triggered after new content has been swapped in.
	EventAfterSwap = "htmx:afterSwap"

	// EventBeforeCleanupElement triggered before htmx disables an element or removes it from the DOM.
	EventBeforeCleanupElement = "htmx:beforeCleanupElement"

	// EventBeforeOnLoad triggered before any response processing occurs.
	EventBeforeOnLoad = "htmx:beforeOnLoad"

	// EventBeforeProcessNode triggered before htmx initializes a node.
	EventBeforeProcessNode = "htmx:beforeProcessNode"

	// EventBeforeRequest triggered before an AJAX request is made.
	EventBeforeRequest = "htmx:beforeRequest"

	// EventBeforeSwap triggered before a swap is done, allows you to configure the swap.
	EventBeforeSwap = "htmx:beforeSwap"

	// EventBeforeSend triggered just before an ajax request is sent.
	EventBeforeSend = "htmx:beforeSend"

	// EventBeforeTransition triggered before the View Transition wrapped swap occurs.
	EventBeforeTransition = "htmx:beforeTransition"

	// EventConfigRequest triggered before the request, allows you to customize parameters, headers.
	EventConfigRequest = "htmx:configRequest"

	// EventConfirm triggered after a trigger occurs on an element, allows you to cancel (or delay) issuing the AJAX request.
	EventConfirm = "htmx:confirm"

	// EventHistoryCacheError triggered on an error during cache writing.
	EventHistoryCacheError = "htmx:historyCacheError"

	// EventHistoryCacheHit triggered on a cache hit in the history subsystem.
	EventHistoryCacheHit = "htmx:historyCacheHit"

	// EventHistoryCacheMiss triggered on a cache miss in the history subsystem.
	EventHistoryCacheMiss = "htmx:historyCacheMiss"

	// EventHistoryCacheMissLoadError triggered on a unsuccessful remote retrieval.
	EventHistoryCacheMissLoadError = "htmx:historyCacheMissLoadError"

	// EventHistoryCacheMissLoad triggered on a successful remote retrieval.
	EventHistoryCacheMissLoad = "htmx:historyCacheMissLoad"

	// EventHistoryRestore triggered when htmx handles a history restoration action.
	EventHistoryRestore = "htmx:historyRestore"

	// EventBeforeHistorySave triggered before content is saved to the history cache.
	EventBeforeHistorySave = "htmx:beforeHistorySave"

	// EventLoad triggered when new content is added to the DOM.
	EventLoad = "htmx:load"

	// EventNoSSESourceError triggered when an element refers to a SSE event in its trigger,
	// but no parent SSE source has been defined.
	EventNoSSESourceError = "htmx:noSSESourceError"

	// EventOnLoadError triggered when an exception occurs during the onLoad handling in htmx.
	EventOnLoadError = "htmx:onLoadError"

	// EventOOBAfterSwap triggered after an out of band element as been swapped in.
	EventOOBAfterSwap = "htmx:oobAfterSwap"

	// EventOOBBeforeSwap triggered before an out of band element swap is done,
	// allows you to configure the swap.
	EventOOBBeforeSwap = "htmx:oobBeforeSwap"

	// EventOOBErrorNoTarget triggered when an out of band element does not have
	// a matching ID in the current DOM.
	EventOOBErrorNoTarget = "htmx:oobErrorNoTarget"

	// EventPrompt triggered after a prompt is shown.
	EventPrompt = "htmx:prompt"

	// EventPushedIntoHistory triggered after a url is pushed into history.
	EventPushedIntoHistory = "htmx:pushedIntoHistory"

	// EventReplacedInHistory triggered after a url is replaced in history.
	EventReplacedInHistory = "htmx:replacedInHistory"

	// EventResponseError triggered when an HTTP response error
	// (non-200 or 300 response code) occurs.
	EventResponseError = "htmx:responseError"

	// EventSendAbort triggered when a request is aborted.
	EventSendAbort = "htmx:sendAbort"

	// EventSendError triggered when a network error prevents an HTTP request from happening.
	EventSendError = "htmx:sendError"

	// EventSSEError triggered when an error occurs with a SSE source.
	EventSSEError = "htmx:sseError"

	// EventSSEOpen triggered when a SSE source is opened.
	EventSSEOpen = "htmx:sseOpen"

	// EventSwapError triggered when an error occurs during the swap phase.
	EventSwapError = "htmx:swapError"

	// EventTargetError triggered when an invalid target is specified.
	EventTargetError = "htmx:targetError"

	// EventTimeout triggered when a request timeout occurs.
	EventTimeout = "htmx:timeout"

	// EventValidationValidate triggered before an element is validated.
	EventValidationValidate = "htmx:validation:validate"

	// EventValidationFailed triggered when an element fails validation.
	EventValidationFailed = "htmx:validation:failed"

	// EventValidationHalted triggered when a request is halted due to validation errors.
	EventValidationHalted = "htmx:validation:halted"

	// EventXHRAbort triggered when an ajax request aborts.
	EventXHRAbort = "htmx:xhr:abort"

	// EventXHRLoadEnd triggered when an ajax request ends.
	EventXHRLoadEnd = "htmx:xhr:loadend"

	// EventXHRLoadStart triggered when an ajax request starts.
	EventXHRLoadStart = "htmx:xhr:loadstart"

	// EventXHRProgress triggered periodically during an ajax request
	// that supports progress events.
	EventXHRProgress = "htmx:xhr:progress"
)
