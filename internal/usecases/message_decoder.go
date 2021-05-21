package usecases

import (
	"errors"
	"fmt"
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

	size := len(messages[0])
	var wordsNotRepeat []string

	for i := 0; i < size; i++ {
		for j := 0; j < len(messages); j++ {
			fmt.Println(messages[j][i])
			if messages[j][i] != "" && !decoder.constains(wordsNotRepeat, messages[j][i]) {
				wordsNotRepeat = append(wordsNotRepeat, messages[j][i])
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
