package usecases

import (
	"errors"

	core "github.com/franciscoruizar/quasar-fire/internal"
)

type LocationFinder struct {
	repository             core.SateliteRepository
	intersectionCalculator IntersectionCalculator
}

func NewLocationFinder(repository core.SateliteRepository) LocationFinder {
	return LocationFinder{
		repository:             repository,
		intersectionCalculator: NewIntersectionCalculator(),
	}
}

func (finder LocationFinder) Find(distances []float64) (core.Position, error) {
	satelites := finder.repository.FindAll()

	circles := finder.getRadioBetweenDistanceAndPosition(satelites, distances)

	positions := finder.findIntersectionsBetweenCircles(circles)

	response, err := finder.findMoreOcurrences(positions)

	if err != nil {
		return core.Position{}, errors.New("cannot find position")
	}

	return response, nil
}

func (finder LocationFinder) getRadioBetweenDistanceAndPosition(satelites []core.Satelite, distances []float64) []core.Circle {
	var circles []core.Circle
	for i := 0; i < len(satelites); i++ {
		circles = append(circles, core.NewCircle(satelites[i].Position.X, satelites[i].Position.Y, distances[i]))
	}

	return circles
}

func (finder LocationFinder) findIntersectionsBetweenCircles(circles []core.Circle) [][]core.Position {
	var positions [][]core.Position
	for i := 0; i < len(circles); i++ {
		for j := 0; j < len(circles); j++ {
			if j != i {
				intersect, err := finder.intersectionCalculator.Calculate(circles[i], circles[j])

				if err == nil {
					positions = append(positions, intersect)
				}
			}
		}
	}

	return positions
}

func (finder LocationFinder) findMoreOcurrences(arrays [][]core.Position) (core.Position, error) {
	var array []core.Position

	for i := 0; i < len(arrays); i++ {
		for j := 0; j < len(arrays[i]); j++ {
			array = append(array, arrays[i][j])
		}
	}

	if array == nil {
		return core.Position{}, errors.New("error into array")
	}

	position := finder.findMoreOcurrencesIntoArray(array)

	return position, nil
}

func (finder LocationFinder) findMoreOcurrencesIntoArray(array []core.Position) core.Position {
	position := array[0]

	for i := 0; i < len(array); i++ {
		if finder.occurrences(position, array) < finder.occurrences(array[i], array) {
			position = array[i]
		}
	}

	return position
}

func (finder LocationFinder) occurrences(position core.Position, array []core.Position) int {
	counter := 0

	for i := 0; i < len(array); i++ {
		if position.X == array[i].X && position.Y == array[i].Y {
			counter = counter + 1
		}
	}

	return counter
}
