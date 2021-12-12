package aoc012

import (
	util "adventofcode/util"
	"fmt"
	"strings"
)

type node struct {
	connections []string
}

const startNode string = "start"
const endNode string = "end"

func challengeA(input []string) int {

	nodes := make(map[string]node)

	for _, line := range input {
		connection := strings.Split(line, "-")
		if val, prs := nodes[connection[0]]; prs {
			nodes[connection[0]] = node{connections: append(val.connections, connection[1])}
		} else {
			connections := make([]string, 1)
			connections[0] = connection[1]
			nodes[connection[0]] = node{connections: connections}
		}

		if val, prs := nodes[connection[1]]; prs {
			nodes[connection[1]] = node{connections: append(val.connections, connection[0])}
		} else {
			connections := make([]string, 1)
			connections[0] = connection[0]
			nodes[connection[1]] = node{connections: connections}
		}
	}

	fmt.Printf("challengeA(): nodes = %v\n", nodes)

	paths := make([]string, 0)

	var walkTheNode func(path string, currentNode node)
	walkTheNode = func(path string, currentNode node) {
		for _, connection := range currentNode.connections {
			currentPath := path + "->" + connection
			if connection == endNode {
				paths = append(paths, currentPath)
			} else if connection != startNode {
				if util.IsUpper(connection) || strings.Count(path, "->"+connection) < 1 {
					walkTheNode(currentPath, nodes[connection])
				}
			}
		}
	}

	walkTheNode(startNode, nodes[startNode])

	fmt.Printf("challengeA(): paths = %v\n", paths)

	return len(paths)
}

func challengeB(input []string) int {

	nodes := make(map[string]node)

	for _, line := range input {
		connection := strings.Split(line, "-")
		if val, prs := nodes[connection[0]]; prs {
			nodes[connection[0]] = node{connections: append(val.connections, connection[1])}
		} else {
			connections := make([]string, 1)
			connections[0] = connection[1]
			nodes[connection[0]] = node{connections: connections}
		}

		if val, prs := nodes[connection[1]]; prs {
			nodes[connection[1]] = node{connections: append(val.connections, connection[0])}
		} else {
			connections := make([]string, 1)
			connections[0] = connection[0]
			nodes[connection[1]] = node{connections: connections}
		}
	}

	fmt.Printf("challengeA(): nodes = %v\n", nodes)

	smallCaves := make([]string, 0)

	for nodeName, _ := range nodes {
		if util.IsLower(nodeName) {
			smallCaves = append(smallCaves, nodeName)
		}
	}

	paths := make(map[string]int)

	var walkTheNode func(path string, currentNode node)
	walkTheNode = func(path string, currentNode node) {
		containsNoSmallCaveTwice := true
		for _, smallCave := range smallCaves {
			if strings.Count(path, "->"+smallCave) >= 2 {
				containsNoSmallCaveTwice = false
				break
			}
		}
		for _, connection := range currentNode.connections {
			currentPath := path + "->" + connection
			if connection == endNode {
				if val, ok := paths[currentPath]; ok {
					val++
				} else {
					paths[currentPath] = 1
				}
			} else if connection != startNode {
				if util.IsUpper(connection) || containsNoSmallCaveTwice || strings.Count(path, "->"+connection) == 0 {
					walkTheNode(currentPath, nodes[connection])
				}
			}
		}
	}

	walkTheNode(startNode, nodes[startNode])

	fmt.Printf("challengeA(): paths = %v\n", paths)

	return len(paths)
}
