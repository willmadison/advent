package advent2022_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willmadison/advent/advent2022"
)

func TestFindTotalDirectorySizeWithConstraint(t *testing.T) {
	given := strings.NewReader(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`)

	expectedTotalSize := 95437

	totalSize, err := advent2022.FindTotalDirectorySizeWithConstraint(given, 100_000)

	assert.Nil(t, err)

	assert.Equal(t, expectedTotalSize, totalSize)
}

func TestFindSmallestDirectoryToDelete(t *testing.T) {
	given := strings.NewReader(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`)

	expectedSize := 24933642

	size, err := advent2022.FindSmallestDirectoryToDelete(given, 30_000_000)

	assert.Nil(t, err)

	assert.Equal(t, expectedSize, size)
}
