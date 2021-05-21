package usecases_test

import (
	"math"
	"testing"

	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
	usecases "github.com/franciscoruizar/quasar-fire/internal/usecases"
	"github.com/stretchr/testify/assert"
)

func Test_DistanceCalculator_Succeed(t *testing.T) {
	positionA, _ := domain.NewPosition(-500, -200)
	positionB, _ := domain.NewPosition(100, -100)
	expected := 608.0

	calculator := usecases.NewDistanceCalculator()

	actual := math.Round(calculator.Calculate(positionA, positionB))

	assert.Equal(t, expected, actual)
}
