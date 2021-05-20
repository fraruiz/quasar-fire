package usecases_test

import (
	"testing"

	"github.com/franciscoruizar/quasar-fire/internal/usecases"
	"github.com/stretchr/testify/assert"
)

func Test_MessageDecoder_Succeed(t *testing.T) {
	message1 := []string{"este", "es", "un", "mensaje"}
	message2 := []string{"este", "", "un", "mensaje"}
	message3 := []string{"", "es", "", "mensaje"}

	messages := [][]string{message1, message2, message3}

	decoder := usecases.NewMessageDecoder()
	actual, _ := decoder.Decode(messages...)

	expected := "este es un mensaje"

	assert.Equal(t, expected, actual)
}
