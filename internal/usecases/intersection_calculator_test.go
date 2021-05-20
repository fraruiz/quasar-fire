package usecases_test

import (
	"testing"

	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
	usecases "github.com/franciscoruizar/quasar-fire/internal/usecases"
	"github.com/stretchr/testify/assert"
)

func Test_IntersectionCalculador_Succeed(t *testing.T) {
	circleA, _ := domain.NewCircle(-500, -200, 538.5164807134505)
	circleB, _ := domain.NewCircle(100, -100, 141.4213562373095)

	calculator := usecases.NewIntersectionCalculator()

	actual, _ := calculator.Calculate(circleA, circleB)

	positionA, _ := domain.NewPosition(38, -227)
	positionB, _ := domain.NewPosition(0, 0)
	expected := []domain.Position{positionA, positionB}

	assert.Equal(t, expected, actual)
}
