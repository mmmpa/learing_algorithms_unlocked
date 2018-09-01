package main

import "github.com/k0kubun/pp"

type Node struct {
	Name  string
	Nears *Linker
	In    int
}

func NewNode(name string) *Node {
	return &Node{
		Name:  name,
		Nears: &Linker{},
	}
}

func (o *Node) AddNear(node *Node) {
	node.Increment()
	pp.Println(node.Name,node.In)
	o.Nears.Push(&Item{Node: node})
}

func (o *Node) NearHead() *Item {
	return o.Nears.Head
}

func (o *Node) Increment() {
	o.In++
}

type Linker struct {
	Head *Item
	Tail *Item
	Size int
}

type Item struct {
	Node *Node
	Next *Item
}

func (o *Linker) Push(item *Item) {
	if o.Head == nil {
		o.Head = item
		o.Tail = item
	} else {
		o.Tail.Next = item
		o.Tail = item
	}
	o.Size++
}

func (o *Linker) Shift() *Item {
	head := o.Head
	o.Head = head.Next
	o.Size--
	return head
}

func (o *Linker) Has() bool {
	return o.Head != nil
}

func topologicalSort(nodes []*Node) []string {
	queue := Linker{}

	restNode := len(nodes)
	rests := make([]int, restNode)
	linker := make([]string, restNode)
	indexes := map[string]int{}

	pos := 0

	for i, n := range nodes {
		rests[i] = n.In
		indexes[n.Name] = i

		if n.In == 0 {
			pp.Println(n.Name)
			item := &Item{Node: n}
			queue.Push(item)

			linker[pos] = n.Name
			pos++
		}
	}

	for queue.Size != 0 {
		node := queue.Shift().Node

		now := node.NearHead()
		for now != nil {
			i := indexes[now.Node.Name]
			rests[i]--

			if rests[i] == 0 {
				item := &Item{Node: now.Node}
				queue.Push(item)

				linker[pos] = now.Node.Name
				pos++
			}

			now = now.Next
		}
	}

	return linker
}
