package advent2020

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Tileset []Tile

func (t Tileset) FindCorners() []int {
	allEdges := map[Edge][]int{}

	for _, tile := range t {
		edges := tile.Edges()

		for _, edge := range edges {
			flip := edge.Flip()
			allEdges[edge] = append(allEdges[edge], tile.ID)
			allEdges[flip] = append(allEdges[flip], tile.ID)
		}
	}

	tilesWithOutlierEdges := map[int]int{}

	for _, tileIDs := range allEdges {
		if len(tileIDs) < 2 {
			tilesWithOutlierEdges[tileIDs[0]]++
		}

	}

	var corners []int

	for tileID, unmatchedEdges := range tilesWithOutlierEdges {
		if unmatchedEdges > 2 && unmatchedEdges < 5 {
			corners = append(corners, tileID)
		}
	}

	return corners
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
