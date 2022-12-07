package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const marker = 14

func readFile(filename string) []string {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	return lines
}

func findStartOfPacket(line string) int {
	var i int
	for i = marker; i < len(line); i++ {
		var bad = false
		var four = line[i-marker : i]
		keys := make(map[rune]bool)
		for _, v := range four {
			if keys[v] {
				bad = true
				break
			}
			keys[v] = true
		}
		if !bad {
			return i
		}
	}
	return len(line)
}

type File struct {
	Name string
	Size int
}

type Folder struct {
	Name         string
	Files        []File
	Folders      map[string]*Folder
	ParentFolder *Folder
}

func newFolder(name string, parent *Folder) *Folder {
	return &Folder{name, []File{}, make(map[string]*Folder), parent}
}

func findCandidates(pointer *Folder, dirs []string) (int, []string) {
	log.Println(pointer.Name, len(pointer.Files))
	size := 0
	for _, folder := range pointer.Folders {
		log.Println(folder.Name)
		s, d := findCandidates(folder, dirs)
		size += s
		dirs = append(dirs, d...)
	}

	for _, file := range pointer.Files {
		size += file.Size
	}

	if size <= 100000 {
		dirs = append(dirs, pointer.Name)
	}

	return size, dirs

}

func main() {

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)

	lines := readFile("input.txt")
	// skip the first line, it goes to root
	lines = lines[1:]
	root := Folder{"/", []File{}, make(map[string]*Folder), nil}
	pointer := &root

	// build directory tree
	for _, line := range lines {
		// log.Println(root)
		// log.Println(num, line, pointer.Name)
		if line == "cd .." {
			pointer = pointer.ParentFolder
		} else if strings.HasPrefix(line, "$ cd") {
			// log.Println(pointer.Folders, line[4:])
			for k, f := range pointer.Folders {
				if k == line[4:] {
					pointer = f
					break
				}
			}
			// log.Println(pointer.Name)
		} else if strings.HasPrefix(line, "dir ") {
			name := line[4:]
			pointer.Folders[name] = newFolder(name, pointer)
			// log.Println(pointer.Folders[line[4:]], pointer.Folders[line[4:]].Name, pointer.Folders[line[4:]].ParentFolder.Name)
		} else if !strings.HasPrefix(line, "$ ls") {
			s := strings.Split(line, " ")
			size, _ := strconv.Atoi(s[0])
			name := s[1]
			pointer.Files = append(pointer.Files, File{name, size})
		}
	}
	for _, f := range root.Folders {
		log.Println(f.Name, len(f.Files), len(f.Folders))
	}
	// walk directory tree
	// var dirs []string
	// log.Println(findCandidates(&root, dirs))
}
