package utils

import (
	"errors"
	"fmt"
)

var ErrDeserializeNameGotEmptyString = errors.New("DeserializeName() got the empty string which is invalid")
var ErrDeserializeNameGotInvalidString = errors.New("DeserializeName() got a string which was not surounded by '<' and '>'")
var ErrDeserializeNameGotEmptyName = errors.New("DeserializeName() got an empty name which is not allowed")

func DeserializeName(quoted string) (string, error) {
	if l := len(quoted); l == 0 {
		return "", ErrDeserializeNameGotEmptyString
	} else if quoted[0] != '<' || quoted[l-1] != '>' {
		return "", fmt.Errorf("%w: %q", ErrDeserializeNameGotInvalidString, quoted)
	} else if l == 2 {
		return "", fmt.Errorf("%w: %q", ErrDeserializeNameGotEmptyName, quoted)
	} else {
		return quoted[1 : l-1], nil
	}
}

func SerializeName(name string) string {
	return fmt.Sprintf("<%s>", name)
}
