package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	part1(file)
	file.Seek(0, io.SeekStart)
	part2(file)
}

func part1(input io.Reader) {
	fileSystem := createFileSystem(input)

	sizes := make(map[string]int)
	calcualteSize("/", fileSystem, sizes)

	var result int
	for _, v := range sizes {
		if v <= 100000 {
			result += v
		}
	}
	fmt.Println(result)
}

func createFileSystem(input io.Reader) map[string][]string {
	scanner := bufio.NewScanner(input)
	fileStructure := make(map[string][]string)

	var dir string
	var path []string
	for scanner.Scan() {
		row := scanner.Text()
		if row[0] == '$' { // command
			var command, argument string
			fmt.Sscanf(row, "$ %s %s", &command, &argument)
			if command == "cd" {
				if argument == ".." {
					path = path[:len(path)-1]
					dir = strings.Join(path, "/")
				} else {
					// root is an empty string in the path to prevent starting pathname with // instead of /
					if argument == "/" {
						path = append(path, "")
						dir = "/"
					} else {
						path = append(path, argument)
						dir = strings.Join(path, "/")
					}
				}
			}
		} else { // ls output
			fileStructure[dir] = append(fileStructure[dir], row)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return fileStructure
}

func calcualteSize(dir string, filestructure map[string][]string, sizes map[string]int) int {
	for _, file := range filestructure[dir] {
		// a dir has a dir prefix, files have a size as prefix
		if file[0] == 'd' {
			var dirName, fullDir string
			fmt.Sscanf(file, "dir %s", &dirName)
			if dir == "/" {
				fullDir = "/" + dirName
			} else {
				fullDir = dir + "/" + dirName
			}
			sizes[fullDir] = calcualteSize(fullDir, filestructure, sizes)
			sizes[dir] += sizes[fullDir]
		} else {
			var size int
			var filename string
			fmt.Sscanf(file, "%d %s", &size, &filename)
			sizes[dir] += size
		}
	}
	return sizes[dir]
}

func part2(input io.Reader) {
	fileSystem := createFileSystem(input)

	sizes := make(map[string]int)
	totalSize := calcualteSize("/", fileSystem, sizes)
	freeSpace := 70000000 - totalSize
	toDelete := 30000000 - freeSpace

	sizeSlice := make([]int, 0, len(sizes))
	for _, v := range sizes {
		sizeSlice = append(sizeSlice, v)
	}

	sort.Slice(sizeSlice, func(i, j int) bool {
		return sizeSlice[j] > sizeSlice[i]
	})

	var result int
	for _, v := range sizeSlice {
		if v >= toDelete {
			result = v
			break
		}
	}

	fmt.Println(result)
}
