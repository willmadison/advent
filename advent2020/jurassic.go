package advent2020

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/willmadison/advent/internal/location"
)

type Tileset []Tile

type Side int

const (
	TOP Side = iota
	RIGHT
	BOTTOM
	LEFT
)

func (t Tileset) FindCorners() ([]int, map[int]struct{}) {
	tilesWithOutlierEdges := map[int]int{}

	for _, tileIDs := range t.AllEdges() {
		if len(tileIDs) < 2 {
			tilesWithOutlierEdges[tileIDs[0]]++
		}

	}

	var corners []int
	sides := map[int]struct{}{}

	for tileID, unmatchedEdges := range tilesWithOutlierEdges {
		if unmatchedEdges <= 2 {
			sides[tileID] = struct{}{}
		}

		if unmatchedEdges > 2 && unmatchedEdges < 5 {
			corners = append(corners, tileID)
		}
	}

	return corners, sides
}

func (t Tileset) AllEdges() map[Edge][]int {
	allEdges := map[Edge][]int{}

	for _, tile := range t {
		edges := tile.Edges()

		for _, edge := range edges {
			flip := edge.Flip()
			allEdges[edge] = append(allEdges[edge], tile.ID)
			allEdges[flip] = append(allEdges[flip], tile.ID)
		}
	}

	return allEdges
}

type CoordinateQueue interface {
	Enqueue(location.Coordinate)
	Dequeue() (location.Coordinate, error)
	Peek() (location.Coordinate, error)
	Size() int
}

type coordinateQueue struct {
	data []location.Coordinate
	size int
}

func NewCoordinateQueue(coordinates ...location.Coordinate) CoordinateQueue {
	return &coordinateQueue{data: coordinates, size: len(coordinates)}
}

func (c *coordinateQueue) Enqueue(value location.Coordinate) {
	c.data = append(c.data, value)
	c.size++
}

func (c *coordinateQueue) Dequeue() (location.Coordinate, error) {
	if c.size > 0 {
		value := c.data[0]
		c.size--
		c.data = c.data[1:]
		return value, nil
	}

	return location.Coordinate{}, errors.New("No Such Element")
}

func (c *coordinateQueue) Peek() (location.Coordinate, error) {
	if c.size > 0 {
		value := c.data[0]
		return value, nil
	}

	return location.Coordinate{}, errors.New("No Such Element")
}

func (c coordinateQueue) Size() int {
	return c.size
}

func (t *Tileset) ProperlyArrange() {
	allLocations := []location.Coordinate{}
	indiciesByLocation := map[location.Coordinate]int{}

	var locationsAdded int

	var maxRow, maxCol int

	for row := 0; row*row < len(*t); row++ {
		for col := 0; col*col < len(*t); col++ {
			allLocations = append(allLocations, location.Coordinate{row, col})
			indiciesByLocation[location.Coordinate{row, col}] = locationsAdded
			locationsAdded++

			if col > maxCol {
				maxCol = col
			}
		}

		if row > maxRow {
			maxRow = row
		}
	}

	tilesByID := map[int]Tile{}

	for _, tile := range *t {
		tilesByID[tile.ID] = tile
	}

	corners, _ := t.FindCorners()

	allEdges := t.AllEdges()

	topLeftID := math.MinInt64

	for _, corner := range corners {
		tile := tilesByID[corner]
		rightEdge := tile.Edges()[RIGHT]
		flippedRightEdge := rightEdge.Flip()

		bottomEdge := tile.Edges()[BOTTOM]
		flippedBottomEdge := bottomEdge.Flip()

		matchesRight := allEdges[rightEdge]
		matchesRight = append(matchesRight, allEdges[flippedRightEdge]...)

		matchesBottom := allEdges[bottomEdge]
		matchesBottom = append(matchesBottom, allEdges[flippedBottomEdge]...)

		var rightMatches, bottomMatches int

		for _, match := range matchesRight {
			if match != corner {
				rightMatches++
			}
		}

		for _, match := range matchesBottom {
			if match != corner {
				bottomMatches++
			}
		}

		if rightMatches < 2 || bottomMatches < 2 {
			continue
		}

		topLeftID = corner
		break
	}

	tilesByLocation := map[location.Coordinate]Tile{}
	locationsByTileID := map[int]location.Coordinate{}

	topLeft := tilesByID[topLeftID]
	tilesByLocation[location.Coordinate{}] = topLeft

	placed := map[int]struct{}{
		topLeft.ID: {},
	}

	queue := NewCoordinateQueue(allLocations...)

	for queue.Size() > 0 {
		currentLocation, _ := queue.Dequeue()

		if _, tileInLocation := tilesByLocation[currentLocation]; tileInLocation {
			continue
		}

		for id, tile := range tilesByID {
			var tileInPlace bool

			if _, tileInPlace = placed[id]; tileInPlace {
				continue
			}

			right := location.Coordinate{currentLocation.Row, currentLocation.Col + 1}
			bottom := location.Coordinate{currentLocation.Row + 1, currentLocation.Col}
			left := location.Coordinate{currentLocation.Row, currentLocation.Col - 1}
			up := location.Coordinate{currentLocation.Row - 1, currentLocation.Col}

			rightTile, hasRight := tilesByLocation[right]
			bottomTile, hasBottom := tilesByLocation[bottom]
			leftTile, hasLeft := tilesByLocation[left]
			aboveTile, hasUp := tilesByLocation[up]

			if hasRight {
				leftEdge := rightTile.Edges()[LEFT]

				for _, e := range tile.Edges() {
					var needsTransformation bool

					rightEdge := tile.Edges()[RIGHT]

					if e == leftEdge {
						placed[tile.ID] = struct{}{}
						locationsByTileID[tile.ID] = currentLocation
						tilesByLocation[currentLocation] = tile
						needsTransformation = e != rightEdge
					}

					if e.Flip() == leftEdge {
						placed[tile.ID] = struct{}{}
						locationsByTileID[tile.ID] = currentLocation
						tilesByLocation[currentLocation] = tile
						needsTransformation = true
					}

					if needsTransformation {
						var rotations int
						var flipped bool

						for rightEdge != leftEdge {
							if rotations < 4 {
								tile.Rotate()
								rotations++
							} else if !flipped {
								tile.Flip()
								flipped = true
							} else {
								panic("unable to find a proper orientation")
							}

							rightEdge = tile.Edges()[RIGHT]
						}
					}
				}
			}

			_, tileInPlace = placed[tile.ID]

			if hasBottom && !tileInPlace {
				topEdge := bottomTile.Edges()[TOP]

				for _, e := range tile.Edges() {
					var needsTransformation bool

					bottomEdge := tile.Edges()[BOTTOM]

					if e == topEdge {
						placed[tile.ID] = struct{}{}
						locationsByTileID[tile.ID] = currentLocation
						tilesByLocation[currentLocation] = tile
						needsTransformation = e != bottomEdge
					}

					if e.Flip() == topEdge {
						placed[tile.ID] = struct{}{}
						locationsByTileID[tile.ID] = currentLocation
						tilesByLocation[currentLocation] = tile
						needsTransformation = true
					}

					if needsTransformation {
						var rotations int
						var flipped bool

						for bottomEdge != topEdge {
							if rotations < 4 {
								tile.Rotate()
								rotations++
							} else if !flipped {
								tile.Flip()
								flipped = true
							} else {
								panic("unable to find a proper orientation")
							}

							bottomEdge = tile.Edges()[BOTTOM]
						}
					}
				}
			}

			_, tileInPlace = placed[tile.ID]

			if hasLeft && !tileInPlace {
				rightEdge := leftTile.Edges()[RIGHT]

				for _, e := range tile.Edges() {
					var needsTransformation bool

					leftEdge := tile.Edges()[LEFT]

					if e == rightEdge {
						placed[tile.ID] = struct{}{}
						locationsByTileID[tile.ID] = currentLocation
						tilesByLocation[currentLocation] = tile
						needsTransformation = e != leftEdge
					}

					if e.Flip() == rightEdge {
						placed[tile.ID] = struct{}{}
						locationsByTileID[tile.ID] = currentLocation
						tilesByLocation[currentLocation] = tile
						needsTransformation = true
					}

					if needsTransformation {
						var rotations int
						var flipped bool

						for leftEdge != rightEdge {
							if rotations < 4 {
								tile.Rotate()
								rotations++
							} else if !flipped {
								tile.Flip()
								flipped = true
							} else {
								panic("unable to find a proper orientation")
							}

							leftEdge = tile.Edges()[LEFT]
						}
					}
				}
			}

			_, tileInPlace = placed[tile.ID]

			if hasUp && !tileInPlace {
				bottomEdge := aboveTile.Edges()[BOTTOM]

				for _, e := range tile.Edges() {
					var needsTransformation bool

					topEdge := tile.Edges()[TOP]

					if e == bottomEdge {
						placed[tile.ID] = struct{}{}
						locationsByTileID[tile.ID] = currentLocation
						tilesByLocation[currentLocation] = tile
						needsTransformation = e != topEdge
					}

					if e.Flip() == bottomEdge {
						placed[tile.ID] = struct{}{}
						locationsByTileID[tile.ID] = currentLocation
						tilesByLocation[currentLocation] = tile
						needsTransformation = true
					}

					if needsTransformation {
						var rotations int
						var flipped bool

						for topEdge != bottomEdge {
							if rotations < 4 {
								tile.Rotate()
								rotations++
							} else if !flipped {
								tile.Flip()
								flipped = true
							} else {
								panic("unable to find a proper orientation")
							}

							topEdge = tile.Edges()[TOP]
						}
					}
				}
			}
		}
	}

	for _, location := range allLocations {
		index := indiciesByLocation[location]
		tile := tilesByLocation[location]

		(*t)[index] = tile
	}
}

func (t Tileset) Print() string {
	var buf bytes.Buffer

	var batches [][]Tile

	tilesPerBatch := int(math.Sqrt(float64(len(t))))

	var batch []Tile
	for _, tile := range t {
		if len(batch) == tilesPerBatch {
			batches = append(batches, batch)
			batch = []Tile{}
		}

		batch = append(batch, tile)
	}

	batches = append(batches, batch)

	for _, batch := range batches {
		for row := 0; row < len(t[0].Pixels); row++ {
			for _, tile := range batch {
				buf.WriteString(tile.Print(row))
			}

			buf.WriteRune('\n')
		}
	}

	return buf.String()
}

type Tile struct {
	ID     int
	Pixels [][]Pixel
}

type Edge uint16

func (e Edge) Flip(bits ...int) Edge {
	var flipped Edge
	numBits := 10

	if len(bits) > 0 {
		numBits = bits[0]
	}

	for i := 0; i < numBits; i++ {
		flipped |= ((e >> i) & 1) << (numBits - 1 - i)
	}

	return flipped
}

func (t Tile) Edges() [4]Edge {
	var top, right, bottom, left Edge

	for i, pixel := range t.Pixels[0] {
		if pixel == Pixel('#') {
			top |= 1 << i
		}
	}

	for i, pixel := range t.Pixels[len(t.Pixels)-1] {
		if pixel == Pixel('#') {
			bottom |= 1 << i
		}
	}

	for i, pixels := range t.Pixels {
		if pixels[0] == Pixel('#') {
			left |= 1 << i
		}

		if pixels[len(pixels)-1] == Pixel('#') {
			right |= 1 << i
		}
	}

	return [4]Edge{top, right, bottom, left}
}

func (t Tile) PrintAll() string {
	var buf bytes.Buffer

	for i := range t.Pixels {
		buf.WriteString(t.Print(i))
		buf.WriteRune('\n')
	}

	return buf.String()
}

func (t Tile) Print(row int) string {
	var buf bytes.Buffer

	for _, pixel := range t.Pixels[row] {
		buf.WriteRune(rune(pixel))
	}

	return buf.String()
}

func (t *Tile) Rotate() {
	N := len(t.Pixels[0])
	for i := 0; i < N/2; i++ {
		for j := i; j < N-i-1; j++ {
			t.Pixels[i][j], t.Pixels[N-1-j][i], t.Pixels[N-1-i][N-1-j], t.Pixels[j][N-1-i] = t.Pixels[N-1-j][i], t.Pixels[N-1-i][N-1-j], t.Pixels[j][N-1-i], t.Pixels[i][j]
		}
	}
}

func (t *Tile) Flip() {
	for i, j := 0, len(t.Pixels[0])-1; i < j; i, j = i+1, j-1 {
		t.Pixels[i], t.Pixels[j] = t.Pixels[j], t.Pixels[i]
	}
}

type Pixel rune

func ParseImageTiles(r io.Reader) Tileset {
	var tiles []Tile

	scanner := bufio.NewScanner(r)

	newTile := true

	var tile Tile

	for scanner.Scan() {
		if newTile {
			tile = Tile{}
			header := scanner.Text()
			tile.ID = extractTileIDFrom(header)
			newTile = false
			continue
		}

		value := scanner.Text()

		if value == "" {
			newTile = true
			tiles = append(tiles, tile)
			continue
		}

		var row int

		tile.Pixels = append(tile.Pixels, parsePixels(value))

		row++
	}

	tiles = append(tiles, tile)

	return Tileset(tiles)
}

func extractTileIDFrom(value string) int {
	tileIDParts := strings.Split(value, " ")
	id, _ := strconv.Atoi(strings.TrimRight(tileIDParts[1], ":"))
	return id
}

func parsePixels(row string) []Pixel {
	var pixels []Pixel

	for _, v := range row {
		pixels = append(pixels, Pixel(v))
	}

	return pixels
}
