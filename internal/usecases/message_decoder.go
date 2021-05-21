package usecases

import (
	"errors"
	"strings"
)

type MessageDecoder struct {
}

func NewMessageDecoder() MessageDecoder {
	return MessageDecoder{}
}

func (decoder MessageDecoder) Decode(messages ...[]string) (string, error) {
	err := decoder.ensureMessages(messages...)

	if err != nil {
		return "", err
	}

	wordsNotRepeat := make([]string, 0)

	for item := 0; item < len(messages); item++ {
		for subitem := 0; subitem < len(messages[item]); subitem++ {
			if !decoder.constains(wordsNotRepeat, messages[item][subitem]) {
				wordsNotRepeat = append(wordsNotRepeat, messages[item][subitem])
			}
		}
	}

	content := strings.Join(wordsNotRepeat[:], " ")
	content = strings.Trim(content, " ")
	return content, nil
}

func (decoder MessageDecoder) ensureMessages(messages ...[]string) error {
	if messages == nil || messages[0] == nil {
		return errors.New("message not determined")
	}

	for i := 1; i < len(messages); i++ {
		if len(messages[i-1]) != len(messages[i]) {
			return errors.New("message not determined")
		}
	}

	return nil
}

func (decoder MessageDecoder) constains(values []string, value string) bool {
	for i := 0; i < len(values); i++ {
		if values[i] == value {
			return true
		}
	}

	return false
}
