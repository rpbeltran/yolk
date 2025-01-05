package utils

import (
	"errors"
	"fmt"
)

var ErrStripAngleQuotedGotEmptyString = errors.New("StripAngleQuotes() got the empty string which is invalid")
var ErrStripAngleQuotedGotInvalidString = errors.New("StripAngleQuotes() got a string which was not surounded by '<' and '>'")

func StripAngleQuotes(quoted string) (string, error) {
	if l := len(quoted); l == 0 {
		return "", ErrStripAngleQuotedGotEmptyString
	} else if quoted[0] == '<' && quoted[l-1] == '>' {
		return quoted[1 : l-1], nil
	} else {
		return "", fmt.Errorf("%w: %q", ErrStripAngleQuotedGotInvalidString, quoted)
	}
}
