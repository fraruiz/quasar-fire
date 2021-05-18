package usecases

import "strings"

type MessageDecoder struct {
}

func NewMessageDecoder() MessageDecoder {
	return MessageDecoder{}
}

func (decoder MessageDecoder) Decode(messages [][]string) string {
	wordsNotRepeat := make([]string, 0)

	for item := 0; item < len(messages); item++ {
		for subitem := 0; subitem < len(messages[item]); subitem++ {
			if !decoder.constains(wordsNotRepeat, messages[item][subitem]) {
				wordsNotRepeat = append(wordsNotRepeat, messages[item][subitem])
			}
		}
	}

	content := strings.Join(wordsNotRepeat[:], " ")

	return content
}

func (decoder MessageDecoder) constains(values []string, value string) bool {
	for i := 0; i < len(values); i++ {
		if values[i] == value {
			return true
		}
	}

	return false
}
