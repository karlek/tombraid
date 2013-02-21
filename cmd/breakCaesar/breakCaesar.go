package main

import "fmt"
import "io/ioutil"
import "log"
import "os"

import "github.com/karlek/tombraid/broken"
import "github.com/karlek/tombraid/frequency"
import "github.com/karlek/tombraid/caesar"

const (
	charset = "abcdefghijklmnopqrstuvwxyz"
)

func main() {
	err := breakCaesar()
	if err != nil {
		log.Println(err)
	}
}

func breakCaesar() (err error) {
	buf, err := ioutil.ReadFile(os.Getenv("GOPATH") + "/src/github.com/karlek/tombraid/message.txt")
	if err != nil {
		return err
	}

	message := string(buf)

	csLen := len(charset)
	isFreqBroken := false
	isBruteBroken := false
	for i := 0; i < csLen; i++ {
		cipher, err := broken.Caesar(charset, message, i)
		if err != nil {
			return err
		}

		commonChars := frequency.EnglishAnalysis(cipher)
		freqGuesses, err := caesar.FrequencyAttack(commonChars, charset, cipher)
		if err != nil {
			return err
		}

		bruteGuesses, err := caesar.BruteAttack(charset, cipher)
		if err != nil {
			return err
		}

		for _, guess := range freqGuesses {
			if guess == message {
				isFreqBroken = true
			}
		}

		for _, guess := range bruteGuesses {
			if guess == message {
				isBruteBroken = true
			}
		}
	}

	switch {
	case isFreqBroken && isBruteBroken:
		fmt.Println("All attacks successful!")
	case isFreqBroken:
		fmt.Println("Only frequency attack worked!")
	case isBruteBroken:
		fmt.Println("Only brute force attack worked!")
	default:
		fmt.Println("No attack broke the cipher!")
	}

	return nil
}
