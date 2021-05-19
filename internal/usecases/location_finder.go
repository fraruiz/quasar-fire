package usecases

import (
	"errors"

	domain "github.com/franciscoruizar/quasar-fire/internal/domain"
)

type LocationFinder struct {
	intersectionCalculator IntersectionCalculator
	satellitesSearcher     SatellitesByNamesSearcher
}

func NewLocationFinder(repository domain.SateliteRepository) LocationFinder {
	return LocationFinder{
		intersectionCalculator: NewIntersectionCalculator(),
		satellitesSearcher:     NewSatellitesByNamesSearcher(repository),
	}
}

func (finder LocationFinder) Find(distances []float64, sattelites []string) (domain.Position, error) {
	satelites, err := finder.satellitesSearcher.Search(sattelites)

	if err != nil {
		return domain.Position{}, err
	}

	circles, err := finder.getRadioBetweenDistanceAndPosition(satelites, distances)

	if err != nil {
		return domain.Position{}, err
	}

	positions := finder.findIntersectionsBetweenCircles(circles)

	response, err := finder.findMoreOcurrences(positions)

	if err != nil {
		return domain.Position{}, errors.New("position could not be found")
	}

	return response, nil
}

func (finder LocationFinder) getRadioBetweenDistanceAndPosition(satelites []domain.Satelite, distances []float64) ([]domain.Circle, error) {
	var circles []domain.Circle
	for i := 0; i < len(satelites); i++ {
		circle, err := domain.NewCircle(satelites[i].Position().X().Value(), satelites[i].Position().Y().Value(), distances[i])

		if err != nil {
			return nil, err
		}

		circles = append(circles, circle)
	}

	return circles, nil
}

func (finder LocationFinder) findIntersectionsBetweenCircles(circles []domain.Circle) [][]domain.Position {
	var positions [][]domain.Position
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

func (finder LocationFinder) findMoreOcurrences(arrays [][]domain.Position) (domain.Position, error) {
	var array []domain.Position

	for i := 0; i < len(arrays); i++ {
		for j := 0; j < len(arrays[i]); j++ {
			array = append(array, arrays[i][j])
		}
	}

	if array == nil {
		return domain.Position{}, errors.New("error into array")
	}

	position := finder.findMoreOcurrencesIntoArray(array)

	return position, nil
}

func (finder LocationFinder) findMoreOcurrencesIntoArray(array []domain.Position) domain.Position {
	position := array[0]

	for i := 0; i < len(array); i++ {
		if finder.occurrences(position, array) < finder.occurrences(array[i], array) {
			position = array[i]
		}
	}

	return position
}

func (finder LocationFinder) occurrences(position domain.Position, array []domain.Position) int {
	counter := 0

	for i := 0; i < len(array); i++ {
		if position.X().Value() == array[i].X().Value() && position.Y().Value() == array[i].Y().Value() {
			counter = counter + 1
		}
	}

	return counter
}
