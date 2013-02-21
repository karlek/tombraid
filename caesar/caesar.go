// Package caesar implements attacks against the caesar cipher
package caesar

import "strings"

import "github.com/karlek/tombraid/broken"
import "github.com/karlek/tombraid/frequency"
import "github.com/mewkiz/pkg/errorsutil"

//
func FrequencyAttack(chars []string, charset, cipher string) (guesses []string, err error) {
	//
	for _, commonChar := range frequency.Top3 {
		for _, char := range chars {
			pos := strings.Index(charset, char)
			if pos == -1 {
				return nil, errorsutil.Errorf("Character `%c` is not present in charset: `%s`", char, charset)
			}

			commonCharPos := strings.Index(charset, commonChar)
			if commonCharPos == -1 {
				return nil, errorsutil.Errorf("Character `%c` is not present in charset: `%s`", commonChar, charset)
			}

			unwindBy := commonCharPos - pos
			unciphered, err := broken.Caesar(charset, cipher, unwindBy)
			if err != nil {
				return nil, err
			}

			guesses = append(guesses, unciphered)
		}

	}

	return guesses, nil
}

// Preforms a brute force attack by shifting the cipher foreach
// character in the charset
func BruteAttack(charset, cipher string) (guesses []string, err error) {
	csLen := len(charset)
	for i := 0; i < csLen; i++ {
		newCipher, err := broken.Caesar(charset, cipher, i)
		if err != nil {
			return nil, err
		}

		guesses = append(guesses, newCipher)
	}

	return guesses, nil
}
