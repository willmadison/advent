package advent2022

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Directory struct {
	name      string
	parent    *Directory
	subdirs   map[string]*Directory
	files     map[string]int
	totalSize int
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		name:    name,
		parent:  parent,
		subdirs: make(map[string]*Directory),
		files:   make(map[string]int),
	}
}

func (d *Directory) calculateSize() int {
	if d.totalSize > 0 {
		return d.totalSize
	}

	size := 0

	for _, fileSize := range d.files {
		size += fileSize
	}

	for _, subdir := range d.subdirs {
		size += subdir.calculateSize()
	}

	d.totalSize = size

	return size
}

func parseDirectoryStructure(input io.Reader) (*Directory, error) {
	scanner := bufio.NewScanner(input)

	root := NewDirectory("/", nil)
	current := root

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "$") {
			parts := strings.Fields(line)
			if len(parts) < 2 {
				continue
			}

			command := parts[1]

			if command == "cd" {
				if len(parts) < 3 {
					continue
				}
				target := parts[2]

				switch target {
				case "/":
					current = root
				case "..":
					if current.parent != nil {
						current = current.parent
					}
				default:
					if _, exists := current.subdirs[target]; !exists {
						current.subdirs[target] = NewDirectory(target, current)
					}
					current = current.subdirs[target]
				}
			}
		} else {
			parts := strings.Fields(line)
			if len(parts) < 2 {
				continue
			}

			if parts[0] == "dir" {
				dirName := parts[1]
				if _, exists := current.subdirs[dirName]; !exists {
					current.subdirs[dirName] = NewDirectory(dirName, current)
				}
			} else {
				size, err := strconv.Atoi(parts[0])
				if err == nil {
					fileName := parts[1]
					current.files[fileName] = size
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	root.calculateSize()

	return root, nil
}

func FindTotalDirectorySizeWithConstraint(input io.Reader, maxSize int) (int, error) {
	root, err := parseDirectoryStructure(input)
	if err != nil {
		return 0, err
	}

	totalSize := 0

	var collectSizes func(*Directory)

	collectSizes = func(dir *Directory) {
		if dir.totalSize <= maxSize {
			totalSize += dir.totalSize
		}

		for _, subdir := range dir.subdirs {
			collectSizes(subdir)
		}
	}

	collectSizes(root)

	return totalSize, nil
}

func FindSmallestDirectoryToDelete(input io.Reader, neededSpace int) (int, error) {
	root, err := parseDirectoryStructure(input)
	if err != nil {
		return 0, err
	}

	totalDiskSpace := 70_000_000
	currentFreeSpace := totalDiskSpace - root.totalSize
	spaceToFree := neededSpace - currentFreeSpace

	if spaceToFree <= 0 {
		return 0, nil
	}

	smallestSize := root.totalSize

	var findSmallest func(*Directory)

	findSmallest = func(dir *Directory) {
		if dir.totalSize >= spaceToFree && dir.totalSize < smallestSize {
			smallestSize = dir.totalSize
		}

		for _, subdir := range dir.subdirs {
			findSmallest(subdir)
		}
	}

	findSmallest(root)

	return smallestSize, nil
}
