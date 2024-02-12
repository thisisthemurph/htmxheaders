package htmxheaders

import (
	"fmt"
)

// Swap represents the type of content swap method used in HTMX.
// It enumerates different ways in which content can be swapped on the client-side
// without a full page reload. Each swap method has its own meaning and effect
// on how content is updated or manipulated.
//
// For more information see: https://htmx.org/attributes/hx-swap/
type Swap int64

const (
	SwapInnerHTML Swap = iota
	SwapOuterHTML
	SwapBeforeBegin
	SwapAfterBegin
	SwapBeforeEnd
	SwapAfterEnd
	SwapDelete
	SwapNone
)

// String returns a string representation of the Swap value.
// If the Swap value is not recognized, it returns "innerHTML" by default.
func (s Swap) String() string {
	switch s {
	case SwapOuterHTML:
		return "outerHTML"
	case SwapBeforeBegin:
		return "beforebegin"
	case SwapAfterBegin:
		return "afterbegin"
	case SwapBeforeEnd:
		return "beforeend"
	case SwapAfterEnd:
		return "afterend"
	case SwapDelete:
		return "delete"
	case SwapNone:
		return "none"
	default:
		return "innerHTML"
	}
}

// SwapFromString converts a string representation to a Swap value.
// If the provided string does not match any known Swap values, it returns SwapInnerHTML by default
// along with an error indicating the invalid string value.
func SwapFromString(s string) (Swap, error) {
	switch s {
	case "innerHTML":
		return SwapInnerHTML, nil
	case "outerHTML":
		return SwapOuterHTML, nil
	case "beforebegin":
		return SwapBeforeBegin, nil
	case "afterbegin":
		return SwapAfterBegin, nil
	case "beforeend":
		return SwapBeforeEnd, nil
	case "afterend":
		return SwapAfterEnd, nil
	case "delete":
		return SwapDelete, nil
	case "none":
		return SwapNone, nil
	default:
		return SwapInnerHTML, fmt.Errorf("invalid Swap value: %q", s)
	}
}

// StringToSwap is identical to SwapFromString and converts a string representation to a Swap value.
// If the provided string does not match any known Swap values, it returns SwapInnerHTML by default
// along with an error indicating the invalid string value.
func StringToSwap(s string) (Swap, error) {
	return SwapFromString(s)
}
