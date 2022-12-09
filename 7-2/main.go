package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ipinfil/aoc2022/loading"
)

type directory struct {
	Path string
	Size int
	Children map[string]*directory
	Parent *directory
}

var pwd string = ""
var baseDir directory
var minSufficient int

func Pwd(cdArg string, currentDir *directory) *directory {
	if cdArg == ".." {
		pwdSplit := strings.Split(strings.TrimLeft(pwd, "/"), "/")
		pwd = "/" + strings.Join(pwdSplit[:len(pwdSplit) - 1], "/")

		return currentDir.Parent
	}

	split := strings.Split(cdArg, "")
	if split[0] == "/" {
		pwd = cdArg
		return &baseDir
	}

	if pwd != "/" {
		pwd += "/"
	}

	pwd += cdArg

	newDirectory := &directory{
		Path: pwd,
		Size: 0,
		Children: make(map[string]*directory),
		Parent: currentDir,
	}
	currentDir.Children[cdArg] = newDirectory

	return newDirectory
}

func GetDirectorySize(dir *directory) int {
	if len(dir.Children) == 0 {
		if dir.Size >= 1111105 && (dir.Size < minSufficient || minSufficient == 0) {
			minSufficient = dir.Size
		}

		return dir.Size
	}

	currentSize := dir.Size
	for _, el := range dir.Children {
		currentSize += GetDirectorySize(el)
	}

	if currentSize >= 1111105 && (currentSize < minSufficient || minSufficient == 0) {
		minSufficient = currentSize
	}
	return currentSize
}

func main() {
	file := loading.LoadFile("7-1/input.txt")
	defer file.Close()
	scanner := loading.NewLineScanner(file)
	var currentDir *directory

	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")

		if split[0] == "$" { // command
			if split[1] == "cd" {
				currentDir = Pwd(split[2], currentDir)

				if currentDir.Path == "" {
					currentDir.Path = pwd
					currentDir.Children = make(map[string]*directory)
				}
				continue
			} else {
				continue
			}
		}

		if split[0] == "dir" {
			continue
		}

		fileSize, err := strconv.Atoi(split[0])
		if err != nil {
			panic(err)
		}

		currentDir.Size += fileSize
	}

	fmt.Println(GetDirectorySize(&baseDir))
	fmt.Println(minSufficient)
}

