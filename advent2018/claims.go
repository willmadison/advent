package advent2018

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Claim struct {
	ID            int
	Location      Point
	Width, Height int
	points        []Point
}

type Point struct {
	Row, Col int
}

type Claimset struct {
	claims []Claim
}

func NewClaimsetFrom(r io.Reader) Claimset {
	var claims []Claim

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		claim := ParseClaim(scanner.Text())
		claims = append(claims, claim)
	}

	return NewClaimset(claims...)
}

func NewClaimset(claims ...Claim) Claimset {
	return Claimset{claims: claims}
}

func (c Claimset) OverlappingRegion() int {
	return len(c.findCommonPoints())
}

func (c Claimset) findCommonPoints() map[Point]struct{} {
	commonPoints := map[Point]struct{}{}
	seen := map[Point]struct{}{}

	for _, claim := range c.claims {
		for _, p := range claim.points {
			if _, present := seen[p]; !present {
				seen[p] = struct{}{}
			} else {
				commonPoints[p] = struct{}{}
			}
		}
	}

	return commonPoints
}

func (c Claimset) FindNonOverlappingClaim() int {
	commonPoints := c.findCommonPoints()

	for _, claim := range c.claims {
		var overlapping bool

		for _, p := range claim.points {
			if _, present := commonPoints[p]; present {
				overlapping = true
				break
			}
		}

		if !overlapping {
			return claim.ID
		}
	}

	return -1
}

func ParseClaim(rawClaim string) Claim {
	fields := strings.Fields(rawClaim)

	rawID := fields[0]
	id, _ := strconv.Atoi(strings.TrimLeft(rawID, "#"))

	rawLocation := fields[2]
	locationParts := strings.Split(rawLocation, ",")

	col, _ := strconv.Atoi(locationParts[0])
	row, _ := strconv.Atoi(strings.TrimSuffix(locationParts[1], ":"))

	location := Point{row, col}

	rawDimensions := fields[3]
	dimensionParts := strings.Split(rawDimensions, "x")

	width, _ := strconv.Atoi(dimensionParts[0])
	height, _ := strconv.Atoi(dimensionParts[1])

	points := determinePointsInRegion(location, width, height)

	return Claim{ID: id,
		Location: location,
		Width:    width, Height: height,
		points: points,
	}
}

func determinePointsInRegion(location Point, width, height int) []Point {
	var points []Point

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			points = append(points, Point{location.Row + row, location.Col + col})
		}
	}

	return points
}
