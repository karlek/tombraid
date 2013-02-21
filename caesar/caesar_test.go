// Test cases for word list attack with different type of hashes
package caesar

import "fmt"
import "io/ioutil"
import "log"
import "strings"
import "testing"

import "github.com/karlek/tombraid/broken"
import "github.com/karlek/tombraid/frequency"
import "github.com/mewkiz/pkg/errorsutil"

func TestFrequencyAttack(t *testing.T) {
	const (
		charset = "abcdefghijklmnopqrstuvwxyz"
	)

	message, err := ioutil.ReadFile(os.Getenv("GOPATH") + "src/github.com/karlek/tombraid/message.txt")

	csLen := len(charset)
	for i := 0; i < csLen; i++ {
		cipher, err := broken.Caesar(charset, message, i)
		if err != nil {
			return err
		}

		chars := frequency.EnglishAnalysis(cipher)
		if err != nil {
			return err
		}

		guesses, err := FrequencyAttack(chars, charset, cipher)
		if err != nil {
			return t.Errorf("FrequencyAttack(%v, %v, %v, %v) = %v, want %v\n\n", chars, charset, cipher, unciphered, message)
		}

		for _, unciphered := range guesses {
			if unciphered == message {
				return nil
			}
		}
	}
}

func TestBruteAttack(t *testing.T) {
	const (
		charset = "abcdefghijklmnopqrstuvwxyz"
	)

	message, err := ioutil.ReadFile(os.Getenv("GOPATH") + "src/github.com/karlek/tombraid/message.txt")

	csLen := len(charset)
	for i := 0; i < csLen; i++ {

	message, err := BruteAttack(charset, cipher)
	if err != nil {
		t.Errorf("BruteAttack(%v, %v) - failed with error: %s\n", charset, cipher, err.Error())
	}
}

