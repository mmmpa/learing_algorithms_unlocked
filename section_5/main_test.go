package main

import (
	"testing"
	"github.com/k0kubun/pp"
	"math/rand"
)


func shuffle(data []*Node) []*Node {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}

	return data
}


func TestCompute(t *testing.T) {
	names := []string{
		"",
		"undershorts",
		"socks",
		"compression shorts",
		"hose",
		"cup",
		"pants",
		"skates",
		"leg pads",
		"T-shirt",
		"chest pad",
		"sweater",
		"mask",
		"catch glove",
		"blocker",
	}
	nodes := make([]*Node, len(names))

	for i, n := range names {
		nodes[i] = NewNode(n)
	}

	nodes[1].AddNear(nodes[3])
	nodes[2].AddNear(nodes[4])
	nodes[3].AddNear(nodes[4])
	nodes[3].AddNear(nodes[5])
	nodes[4].AddNear(nodes[6])
	nodes[5].AddNear(nodes[6])
	nodes[6].AddNear(nodes[7])
	nodes[7].AddNear(nodes[8])
	nodes[8].AddNear(nodes[13])
	nodes[9].AddNear(nodes[10])
	nodes[10].AddNear(nodes[11])
	nodes[11].AddNear(nodes[12])
	nodes[12].AddNear(nodes[13])
	nodes[13].AddNear(nodes[14])

	shuffle(nodes[1:])
	pp.Println(topologicalSort(nodes[1:]))
}
