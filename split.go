package caseconv

import "unicode"

func isWordCharacter(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsNumber(r)
}

func characterFlags(r rune) (number, upper, lower bool) {
	return unicode.IsNumber(r), unicode.IsUpper(r), unicode.IsLower(r)
}

func SplitWords(s string) []string {
	for s != "" && !isWordCharacter([]rune(s)[0]) {
		s = s[1:]
	}

	if s == "" {
		return nil
	}

	word := []rune(s)[:1]

	var words [][]rune
	for _, r := range s[1:] {
		rn, ru, rl := characterFlags(r)
		if !rn && !ru && !rl {
			if len(word) > 0 {
				words = append(words, word)
				word = nil
			}

			continue
		} else if len(word) > 0 {
			if pn, pu, pl := characterFlags(word[len(word)-1]); pn {
				if !rn {
					words = append(words, word)
					word = nil
				}
			} else if pu {
				if !rn && rl && len(word) > 1 {
					words = append(words, word[:len(word)-1])
					word = word[len(word)-1:]
				}
			} else if pl {
				if ru || rn {
					words = append(words, word)
					word = nil
				}
			} else {
				panic("unreachable")
			}
		}

		word = append(word, r)
	}

	if len(word) > 0 {
		words = append(words, word)
	}

	out := make([]string, len(words))
	for i, word := range words {
		out[i] = string(word)
	}

	return out
}
