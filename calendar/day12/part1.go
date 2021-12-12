package adventofcode

import (
	util "adventofcode/util/common"
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cave struct {
	name   string
	isBig  bool
	egress []*cave
}

func printCave_r(node *cave, visited map[string]bool) {
	if visited[node.name] {
		return
	}
	visited[node.name] = true

	util.Dprintf("%s\n", node.name)
	for _, e := range node.egress {
		util.Dprintf("%s ", e.name)
	}
	util.Dprintf("\n~%s\n\n", node.name)
	util.Dprintf("\n")
	for _, e := range node.egress {
		printCave_r(e, visited)
	}
}

func printCaves(node *cave) {
	var visited map[string]bool = make(map[string]bool, 100)
	printCave_r(node, visited)
}

func findPaths_r(node, end *cave, visited map[string]bool, debug string) ([]string, bool) {
	debug += "," + node.name
	if visited[node.name] && !node.isBig {
		util.Dprintf("%v %s V\n", visited, debug)
		return nil, false
	}
	if node.name == end.name {
		npaths := []string{node.name}
		util.Dprintf("%s STOP\n", debug)
		return npaths, true
	}
	visited[node.name] = true
	paths := make([]string, 0, 100)
	found := false
	for _, e := range node.egress {
		var nVisited map[string]bool = make(map[string]bool, 100)
		for k, v := range visited {
			nVisited[k] = v
		}
		npaths, valid := findPaths_r(e, end, nVisited, debug)
		if valid {
			for i := range npaths {
				npaths[i] = node.name + "," + npaths[i]
			}
			found = true
			paths = append(paths, npaths...)
		}
	}
	return paths, found
}

func findPaths(node, end *cave) []string {
	var visited map[string]bool = make(map[string]bool, 100)
	paths, valid := findPaths_r(node, end, visited, "")
	if valid {
		return paths
	} else {
		panic(1)
	}
}

func Part1(filename string) string {
	// STDOUT MUST BE FLUSHED MANUALLY!!!
	defer util.SdoutFlush()

	f, err := os.Open(filename)
	util.Check_error(err)
	defer f.Close()

	file_scanner := bufio.NewScanner(f)

	var caveGraph map[string]*cave = make(map[string]*cave)
	for file_scanner.Scan() {
		line := file_scanner.Text()
		lineSp := strings.Split(line, "-")
		c1 := lineSp[0]
		c2 := lineSp[1]
		if _, exists := caveGraph[c1]; !exists {
			iB := strings.ToUpper(c1) == c1
			caveGraph[c1] = &cave{name: c1, isBig: iB, egress: make([]*cave, 0, 10)}
		}
		if _, exists := caveGraph[c2]; !exists {
			iB := strings.ToUpper(c2) == c2
			caveGraph[c2] = &cave{name: c2, isBig: iB, egress: make([]*cave, 0, 10)}
		}
		caveGraph[c1].egress = append(caveGraph[c1].egress, caveGraph[c2])
		caveGraph[c2].egress = append(caveGraph[c2].egress, caveGraph[c1])
	}
	paths := findPaths(caveGraph["start"], caveGraph["end"])
	util.Dprintf("\n\n")
	for _, p := range paths {
		util.Dprintf("%s\n", p)
	}
	return fmt.Sprintf("%d", len(paths))
}
