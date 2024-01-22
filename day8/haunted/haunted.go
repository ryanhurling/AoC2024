package haunted

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

type Choice rune

const (
	LEFT  Choice = 'L'
	RIGHT Choice = 'R'
)

type Instructions []Choice

type Graph map[string]Node

func (g Graph) getFirstEntry() Node {
	return g["AAA"]
}

func (g Graph) GetAllStartingNodes() []Node {
	var nodes []Node
	for s, n := range g {
		if strings.LastIndexByte(s, 'A') == 2 {
			nodes = append(nodes, n)
		}
	}
	return nodes
}

type Node struct {
	id, leftKey, rightKey string
}

func CreateInstructions(str string) Instructions {
	return Instructions(str)
}

func CreateGraph(entries []string) Graph {
	graph := make(Graph)
	for i, entry := range entries {
		if node, err := createNode(entry); err == nil {
			if i == 0 {
				graph["_ROOT_"] = node
			}
			if _, ok := graph[node.id]; !ok {
				graph[node.id] = node
			}
		}
	}
	return graph
}

func createNode(entry string) (Node, error) {
	regex := regexp.MustCompile("([A-Z]{3})")
	matches := regex.FindAllString(entry, 3)
	if matches[0] == matches[1] && matches[0] == matches[2] {
		return Node{}, fmt.Errorf("cycle Node")
	}
	return Node{id: matches[0], leftKey: matches[1], rightKey: matches[2]}, nil
}

func CountSteps(graph Graph, instructions Instructions) int {
	node := graph.getFirstEntry()
	fmt.Printf("starting at Node: %s\n", node.id)
	var stepCount int
	for i := 0; i < len(instructions); i++ {
		currentInstr := instructions[i]
		switch currentInstr {
		case LEFT:
			node = graph[node.leftKey]
			break
		case RIGHT:
			node = graph[node.rightKey]
			break
		}
		stepCount++
		fmt.Printf("steps: %v\n", stepCount)
		if node.id == "ZZZ" {
			return stepCount
		}
		if i+1 == len(instructions) {
			fmt.Println("resetting to start of instructions")
			i = -1
		}
	}
	return stepCount
}

func walk(startingNode Node, graph Graph, instructions Instructions) int {
	var stepCount int
	for i := 0; i < len(instructions); i++ {
		currentInstr := instructions[i]
		switch currentInstr {
		case LEFT:
			startingNode = graph[startingNode.leftKey]
			break
		case RIGHT:
			startingNode = graph[startingNode.rightKey]
			break
		}
		stepCount++
		fmt.Printf("steps: %v\n", stepCount)
		if strings.LastIndex(startingNode.id, "Z") == 2 {
			return stepCount
		}
		if i+1 == len(instructions) {
			fmt.Println("resetting to start of instructions")
			i = -1
		}
	}
	return stepCount
}

func CountAllSteps(graph Graph, instructions Instructions) int {
	startingNodes := graph.GetAllStartingNodes()
	wG := &sync.WaitGroup{}
	wG.Add(len(startingNodes))
	steps := make([]int, len(startingNodes))
	for i, node := range startingNodes {
		func(n Node, index int, group *sync.WaitGroup) {
			steps[i] = walk(node, graph, instructions)
			wG.Done()
		}(node, i, wG)
	}

	wG.Wait()
	return LCM(steps[0], steps[1], steps[2:]...)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
