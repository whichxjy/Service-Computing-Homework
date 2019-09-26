package main

import (
    "testing"
    "reflect"
)

func TestSort(t *testing.T) {
	// Create Array
	nodes := []Node{{4}, {5}, {8}, {7}, {9}, {3}, {10}, {2}, {6}, {1}}
	// Init
	Init(nodes)
	if !reflect.DeepEqual(nodes, []Node{{1}, {2}, {3}, {4}, {5}, {8}, {10}, {7}, {6}, {9}}) {
		t.Errorf("Fail to init\n")
	}
	// Pop
	min, nodes := Pop(nodes)
	if (min != Node{1} || !reflect.DeepEqual(nodes, []Node{{2}, {4}, {3}, {6}, {5}, {8}, {10}, {7}, {9}})) {
		t.Errorf("Fail to pop\n")
	}
	// Push
	nodes = Push(Node{0}, nodes)
	if !reflect.DeepEqual(nodes, []Node{{0}, {2}, {3}, {6}, {4}, {8}, {10}, {7}, {9}, {5}}) {
		t.Errorf("Fail to push\n")
	}
	// Remove
	nodes = Remove(Node{4}, nodes)
	if !reflect.DeepEqual(nodes, []Node{{0}, {2}, {3}, {5}, {8}, {10}, {7}, {9}, {6}}) {
		t.Errorf("Fail to push\n")
	}
}