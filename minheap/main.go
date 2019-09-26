package main

import "fmt"

type Node struct {
	Value int
}

// Init min heap
func Init(nodes []Node) {
	for i:= len(nodes) / 2 - 1; i >= 0; i-- {
		down(nodes, i, len(nodes))
	}
}

// Downheap
func down(nodes []Node, i, n int) {
	// sub-tree root: r
	// left child: 2 * r + 1
	// right child: 2 * r + 2
	r := i

	// last node index
	last := n - 1

	for 2 * r + 1 <= last {
		j := 2 * r + 1
		if j + 1 <= last && nodes[j + 1].Value < nodes[j].Value {
			j += 1
		}
		if nodes[r].Value > nodes[j].Value {
			nodes[r].Value, nodes[j].Value = nodes[j].Value, nodes[r].Value
			r = j
		} else {
			break
		}
	}
}

// Upheap
func up(nodes []Node, j int) {
	// node: r
	// parent: (r - 1) / 2
	r := j

	for (r - 1) / 2 >= 0 {
		p := (r - 1) / 2
		if nodes[p].Value > nodes[r].Value {
			nodes[p].Value, nodes[r].Value = nodes[r].Value, nodes[p].Value
			r = p
		} else {
			break
		}
	}
}

// Pop the top element
func Pop(nodes []Node) (Node, []Node) {
	nodes[0].Value, nodes[len(nodes) - 1].Value = nodes[len(nodes) - 1].Value, nodes[0].Value
	down(nodes, 0, len(nodes) - 1)
	return nodes[len(nodes) - 1], nodes[:len(nodes) - 1]
}

// Push a new element
func Push(node Node, nodes []Node) []Node {
	nodes = append(nodes, node)
	up(nodes, len(nodes) - 1)
	return nodes
}

// Remove node
func Remove(node Node, nodes []Node) []Node {
	result := make([]Node, 0)
	for _, n := range nodes {
		if n.Value != node.Value {
			result = append(result, Node{n.Value})
		}
	}
	Init(result)
	return result
}

// Check if nodes[] is a min heap
func Check(nodes []Node) {
	if isMinHeap(nodes, 0) {
		fmt.Printf("It's min heap\n\n")
	} else {
		fmt.Printf("It's not min heap\n\n")
	}
}

// Check recursively (start from nodes[r])
func isMinHeap(nodes []Node, r int) bool {
	if r > len(nodes) / 2 - 1 {
		return true
	}

	// left child
	left := 2 * r + 1
	if left <= len(nodes) - 1 && nodes[left].Value < nodes[r].Value {
		return false
	}
	// right chile
	right := 2 * r + 2
	if right <= len(nodes) - 1 && nodes[right].Value < nodes[r].Value {
		return false
	}

	return isMinHeap(nodes, left) && isMinHeap(nodes, right)
}

func main() {
	nodes := []Node{{4}, {5}, {8}, {7}, {9}, {3}, {10}, {2}, {6}, {1}}

	fmt.Println("=== Create Array ===")
	fmt.Printf("nodes: %v\n", nodes)
	Check(nodes)

	fmt.Println("=== Init ===")
	Init(nodes)
	fmt.Printf("nodes: %v\n", nodes)
	Check(nodes)

	fmt.Println("=== Pop ===")
	min, nodes := Pop(nodes)
	fmt.Printf("min: %v\n", min)
	fmt.Printf("nodes: %v\n", nodes)
	Check(nodes)

	fmt.Println("=== Push ===")
	valueToPush := 0
	nodes = Push(Node{valueToPush}, nodes)
	fmt.Printf("value to push: %v\n", valueToPush)
	fmt.Printf("nodes: %v\n", nodes)
	Check(nodes)

	fmt.Println("=== Remove ===")
	valueToRemove := 4
	nodes = Remove(Node{valueToRemove}, nodes)
	fmt.Printf("value to remove: %v\n", valueToRemove)
	fmt.Printf("nodes: %v\n", nodes)
	Check(nodes)
}