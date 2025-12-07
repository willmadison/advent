package advent2025

import (
	"bufio"
	"io"
	"sync"

	"github.com/google/uuid"

	"github.com/willmadison/advent/internal/location"
)

const StartIndicator rune = 'S'
const Splitter rune = '^'

type TachyonBeam struct {
	ID       string
	Location location.Coordinate
}

type laboratoryMap struct {
	SplittersByLocation map[location.Coordinate]struct{}
	StartingLocation    location.Coordinate
	MaxRows             int
	MaxCols             int
}

func parseLaboratoryMap(r io.Reader) (*laboratoryMap, error) {
	scanner := bufio.NewScanner(r)

	var maxRows int
	var maxCols int
	var row int

	var startingLocation location.Coordinate
	splittersByLocation := map[location.Coordinate]struct{}{}

	for scanner.Scan() {
		line := scanner.Text()

		if maxCols == 0 {
			maxCols = len(line)
		}

		for col, value := range line {
			switch rune(value) {
			case StartIndicator:
				startingLocation = location.Coordinate{Row: row, Col: col}
			case Splitter:
				splittersByLocation[location.Coordinate{Row: row, Col: col}] = struct{}{}
			}
		}

		row++
	}

	maxRows = row

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &laboratoryMap{
		SplittersByLocation: splittersByLocation,
		StartingLocation:    startingLocation,
		MaxRows:             maxRows,
		MaxCols:             maxCols,
	}, nil
}

func CountTachyonBeamSplits(r io.Reader) (int, error) {
	labMap, err := parseLaboratoryMap(r)
	if err != nil {
		return 0, err
	}

	var totalSplits int
	var beamsByLocation = map[location.Coordinate]TachyonBeam{}
	var splitCollisions = map[location.Coordinate]struct{}{}

	initialBeam := TachyonBeam{ID: uuid.New().String(), Location: location.Coordinate{Row: labMap.StartingLocation.Row + 1, Col: labMap.StartingLocation.Col}}
	beamsByLocation[initialBeam.Location] = initialBeam

	var wg sync.WaitGroup
	var beamsMut sync.Mutex
	splits := make(chan int)

	wg.Add(1)
	go placeBeam(initialBeam, labMap.MaxRows, labMap.MaxCols, labMap.SplittersByLocation, splitCollisions, beamsByLocation, &beamsMut, splits, &wg)

	go func() {
		wg.Wait()
		close(splits)
	}()

	for split := range splits {
		totalSplits += split
	}

	return totalSplits, nil
}

func CountQuantumTimelines(r io.Reader) (int, error) {
	labMap, err := parseLaboratoryMap(r)
	if err != nil {
		return 0, err
	}

	memo := make(map[location.Coordinate]int)

	initialLocation := location.Coordinate{Row: labMap.StartingLocation.Row + 1, Col: labMap.StartingLocation.Col}
	totalPaths := countPaths(initialLocation, labMap.MaxRows, labMap.MaxCols, labMap.SplittersByLocation, memo)

	return totalPaths, nil
}

func countPaths(coord location.Coordinate, maxRows, maxCols int, splitters map[location.Coordinate]struct{}, memo map[location.Coordinate]int) int {
	if count, exists := memo[coord]; exists {
		return count
	}

	if coord.Row >= maxRows || coord.Col < 0 || coord.Col >= maxCols {
		return 1
	}

	var totalPaths int

	if _, isSplitter := splitters[coord]; isSplitter {
		leftCoord := location.Coordinate{Row: coord.Row + 1, Col: coord.Col - 1}
		rightCoord := location.Coordinate{Row: coord.Row + 1, Col: coord.Col + 1}

		leftPaths := countPaths(leftCoord, maxRows, maxCols, splitters, memo)
		rightPaths := countPaths(rightCoord, maxRows, maxCols, splitters, memo)

		totalPaths = leftPaths + rightPaths
	} else {
		nextCoord := location.Coordinate{Row: coord.Row + 1, Col: coord.Col}
		totalPaths = countPaths(nextCoord, maxRows, maxCols, splitters, memo)
	}

	memo[coord] = totalPaths
	return totalPaths
}

func placeBeam(beam TachyonBeam, maxRows, maxCols int, splittersByLocation, splitCollisions map[location.Coordinate]struct{}, beamsByLocation map[location.Coordinate]TachyonBeam, beamsMut *sync.Mutex, splits chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	col := beam.Location.Col
	splitCount := 0

	for row := beam.Location.Row; row < maxRows; row++ {
		coord := location.Coordinate{Row: row, Col: col}

		if _, isSplitter := splittersByLocation[coord]; isSplitter {
			beamsMut.Lock()
			_, alreadySplit := splitCollisions[coord]
			if !alreadySplit {
				splitCollisions[coord] = struct{}{}
			}
			beamsMut.Unlock()

			if alreadySplit {
				return
			}

			splitCount++

			leftCol := col - 1

			if leftCol >= 0 {
				leftCoord := location.Coordinate{Row: row + 1, Col: leftCol}
				leftBeam := TachyonBeam{ID: uuid.New().String(), Location: leftCoord}
				if registerBeam(leftBeam, leftCoord, beamsByLocation, beamsMut) {
					wg.Add(1)
					go placeBeam(
						leftBeam,
						maxRows, maxCols,
						splittersByLocation,
						splitCollisions,
						beamsByLocation,
						beamsMut,
						splits,
						wg,
					)
				}
			}

			rightCol := col + 1
			if rightCol < maxCols {
				rightCoord := location.Coordinate{Row: row + 1, Col: rightCol}
				rightBeam := TachyonBeam{ID: uuid.New().String(), Location: rightCoord}
				if registerBeam(rightBeam, rightCoord, beamsByLocation, beamsMut) {
					wg.Add(1)
					go placeBeam(
						rightBeam,
						maxRows, maxCols,
						splittersByLocation,
						splitCollisions,
						beamsByLocation,
						beamsMut,
						splits,
						wg,
					)
				}
			}

			splits <- splitCount

			return
		} else {
			beam.Location.Row = row
		}
	}
}

func registerBeam(
	beam TachyonBeam,
	coord location.Coordinate,
	beamsByLocation map[location.Coordinate]TachyonBeam,
	beamsMu *sync.Mutex,
) bool {
	beamsMu.Lock()
	defer beamsMu.Unlock()

	if _, exists := beamsByLocation[coord]; exists {
		return false
	}

	beamsByLocation[coord] = beam
	return true
}
