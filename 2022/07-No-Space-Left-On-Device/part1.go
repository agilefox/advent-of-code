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

func getFullPath(dir *Folder) string {
	path := dir.Name
	if dir.ParentFolder == nil {
		return path
	}
	for p := dir.ParentFolder; p.ParentFolder != nil; p = p.ParentFolder {
		path = fmt.Sprintf("%s/%s", p.Name, path)
	}
	return path
}

var sum = 0

func findCandidates(cur_dir *Folder) int {
	size := 0

	// count all the files in the directory
	for _, file := range cur_dir.Files {
		size += file.Size
	}
	sum += size
	// log.Println(fmt.Sprintf("%s->%d", cur_dir.Name, size))

	// go down into the sub directory
	for _, folder := range cur_dir.Folders {
		size += findCandidates(folder)
	}

	// log.Println(getFullPath(cur_dir))
	// log.Println("\t", size, size > 27322861)
	if size >= 2677139 {
		log.Println(size, cur_dir.Name, size, "*******************")
	}
	// if size >= 100000 {
	// 	log.Println(cur_dir.Name, size)
	// 	sum += size
	// }
	return size

}

func main() {

	log.SetOutput(os.Stdout)
	log.SetFlags(log.Lshortfile)

	lines := readFile("input.txt")
	// skip the first line, it goes to root
	lines = lines[1:]
	root := Folder{"/", []File{}, make(map[string]*Folder), nil}
	cur_dir := &root

	// build directory tree
	for _, line := range lines {
		// log.Println(root)
		// log.Println(num, line, cur_dir.Name)
		if line == "$ cd .." {
			cur_dir = cur_dir.ParentFolder
			// log.Println("Now in ", cur_dir.Name)
		} else if strings.HasPrefix(line, "$ cd") {
			name := line[5:]
			cur_dir = cur_dir.Folders[name]
		} else if strings.HasPrefix(line, "dir ") {
			name := line[4:]
			// log.Println(fmt.Sprintf("Add %s to %s", name, getFullPath(cur_dir)))
			cur_dir.Folders[name] = newFolder(name, cur_dir)

			// log.Println(n+2, "Created dir", getFullPath(cur_dir.Folders[name]))
			// log.Println(cur_dir.Folders[line[4:]], cur_dir.Folders[line[4:]].Name, cur_dir.Folders[line[4:]].ParentFolder.Name)
		} else if !strings.HasPrefix(line, "$ ls") {
			s := strings.Split(line, " ")
			size, _ := strconv.Atoi(s[0])
			name := s[1]
			cur_dir.Files = append(cur_dir.Files, File{name, size})
		}
	}
	// for _, f := range root.Folders {
	// 	log.Println(f.Name, len(f.Files), len(f.Folders))
	// }
	// walk directory tree

	// log.Println(findCandidates(&root), sum)
	size := findCandidates(&root)
	log.Println(size, sum)
}

// 70,000,000 - 42,677,139 = 27,322,861
// 30000000 - 27322861 = 2677139
