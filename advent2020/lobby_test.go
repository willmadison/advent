package advent2020

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDirections(t *testing.T) {
	given := `sesenwnenenewseeswwswswwnenewsewsw`

	expected := []HexDirection{
		Southeast,
		Southeast,
		Northwest,
		Northeast,
		Northeast,
		Northeast,
		Westward,
		Southeast,
		Eastward,
		Southwest,
		Westward,
		Southwest,
		Southwest,
		Westward,
		Northeast,
		Northeast,
		Westward,
		Southeast,
		Westward,
		Southwest,
	}

	actual := ParseDirections(given)
	assert.Equal(t, expected, actual)
}

func TestParseAllDirections(t *testing.T) {
	given := strings.NewReader(`sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`)

	floor := ParseAllDirections(given)

	for _, direction := range floor.AllDirections {
		floor.Follow(direction)
	}

	expectedBlackCount := 10

	assert.Equal(t, expectedBlackCount, floor.GetBlackCount())

	for i := 0; i < 100; i++ {
		floor.Rotate()
	}

	expectedBlackCount = 2208

	assert.Equal(t, expectedBlackCount, floor.GetBlackCount())
}
