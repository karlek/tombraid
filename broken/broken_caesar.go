// Package broken implements unsafe encryption algorithms.
// It's only purpose is to write attacks against each algorithm and prove
// that they are broken.
//
// Warning: package broken should not be used except in educational purpose!
package broken

import "strings"

import "github.com/mewkiz/pkg/errorsutil"

// Caesar shifts each character in the message some fixed number of
// positions down the charset.
func Caesar(charset, message string, shiftBy int) (cipher string, err error) {
	// When shift overflows the slice index we use the modulo of the
	// charset length to simulate rotation.
	// Example:
	// When Y + 2 becomes A; (24 + 2) % 26 = 0		-> charset[0]	= 'a'
	// Or when A - 1 becomes Z; (26 - 1) % 26 = 25	-> charset[25]	= 'z'
	csLen := len(charset)

	// For each character in message
	// Get it's position in the charset
	// The new character is it's position in the charset + shiftBy
	// Simulate rotation with modulo and subtraction.
	for _, mChar := range message {
		pos := strings.IndexRune(charset, mChar)
		if pos == -1 {
			return "", errorsutil.Errorf("Character `%c` is not present in charset: `%s`", mChar, charset)
		}

		posShift := pos + shiftBy
		if posShift < 0 {
			// Example:
			// 26 + (-1) = 25; charset[25] = 'z'
			posShift = csLen + posShift
		}

		// Concatenate the final cipher string
		cipher += string(charset[posShift%csLen])
	}

	return cipher, nil
}
