package frequency

import "sort"

type charFreq struct {
	char string
	freq int
}

type charFreqs []charFreq

func (cfs charFreqs) Len() int {
	return len(cfs)
}

func (cfs charFreqs) Swap(i, j int) {
	cfs[i], cfs[j] = cfs[j], cfs[i]
}

func (cfs charFreqs) Less(i, j int) bool {
	return cfs[i].freq > cfs[j].freq
}

var Top3 = []string{
	"e",
	"t",
	"a",
}

func EnglishAnalysis(cipher string) (chars []string) {
	m := map[string]int{}
	for _, char := range cipher {
		m[string(char)]++
	}

	var cfs charFreqs
	for char, freq := range m {
		cfs = append(cfs, charFreq{char, freq})
	}

	sort.Sort(cfs)

	return []string{cfs[0].char, cfs[1].char, cfs[2].char}
}

/// Problem with caesar cipher occured
// Creates a charset string of all characters used in cipher
// func findCharset(cipher string) (charset string) {
// 	m := map[string]bool{}
// 	for _, char := range cipher {
// 		m[string(char)] = true
// 	}

// 	var tmpCharset []string
// 	for char, _ := range m {
// 		tmpCharset = append(tmpCharset, char)
// 	}

// 	sort.Strings(tmpCharset)
// 	for _, str := range tmpCharset {
// 		charset += str
// 	}

// 	fmt.Println(charset)

// 	return charset
// }
