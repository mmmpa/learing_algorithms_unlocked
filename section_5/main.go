package main

import (
	"github.com/k0kubun/pp"
	"math"
	"fmt"
)

type Node struct {
	Name   string
	Nears  *Linker
	In     int
	Weight int
}

func NewNode(name string) *Node {
	return &Node{
		Name:  name,
		Nears: &Linker{},
	}
}

func NewNodeWithWeight(name string, weight int) *Node {
	return &Node{
		Name:   name,
		Nears:  &Linker{},
		Weight: weight,
	}
}

func (o *Node) AddNear(node *Node) {
	o.Nears.Push(&Item{Node: node, Weight: o.Weight})
	node.Increment()
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
	Node   *Node
	Weight int // to this
	Next   *Item
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

func topologicalSort(nodes []*Node) []*Node {
	queue := Linker{}

	restNode := len(nodes)
	rests := make([]int, restNode)
	linker := make([]*Node, restNode)
	indexes := map[string]int{}

	pos := 0

	for i, n := range nodes {
		rests[i] = n.In
		indexes[n.Name] = i

		if n.In == 0 {
			item := &Item{Node: n}
			queue.Push(item)

			linker[pos] = n
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

				linker[pos] = now.Node
				pos++
			}

			now = now.Next
		}
	}

	return linker
}

func namePrinter(nodes []*Node) {
	for _, n := range nodes {
		pp.Println(n.Name)
	}
}

type CriticalPathResult struct {
	Path   *PathItem
	Weight int
	Step   int
}

func (o CriticalPathResult) PathList() []string {
	path := make([]string, o.Step)

	pos := 0
	for p := o.Path; p != nil; p = p.Next {
		path[pos] = p.Name
		pos++
	}

	return path
}

func (o *CriticalPathResult) Add(name string) {
	prev := o.Path
	o.Path = &PathItem{Name: name}

	if prev != nil {
		o.Path.Next = prev
	}
	o.Step++
}

type PathItem struct {
	Name string
	Next *PathItem
}

func criticalPath(nodes []*Node) CriticalPathResult {
	start := NewNodeWithWeight("start", 0)
	end := NewNodeWithWeight("end", 0)

	nameMap := map[string]*Node{}
	shortest := map[string]int{}
	nearest := map[string]*Node{}

	for _, node := range nodes {
		nameMap[node.Name] = node
		shortest[node.Name] = math.MaxInt32
		nearest[node.Name] = nil

		if node.In == 0 {
			start.AddNear(node)
		}

		if !node.Nears.Has() {
			node.AddNear(end)
		}
	}

	shortest["start"] = 0

	newNodes := append([]*Node{start}, nodes...)
	newNodes = append(newNodes, end)

	list := topologicalSort(newNodes)

	for _, visited := range list {
		next := visited.NearHead()

		for next != nil {
			nextName := next.Node.Name

			if nearest[nextName] == nil {
				nearest[nextName] = visited
				shortest[nextName] = -visited.Weight
				continue
			}

			preShortest := shortest[nextName]
			nextShortest := shortest[visited.Name] - next.Weight

			if preShortest > nextShortest {
				nearest[nextName] = visited
				shortest[nextName] = nextShortest
			}
			next = next.Next
		}
	}

	result := CriticalPathResult{Weight: -shortest["end"]}

	tail := end
	pos := 0
	for {
		result.Add(tail.Name)

		if tail.Name == "start" {
			break
		}

		tail = nearest[tail.Name]
		pos++
	}

	return result
}

type N2 struct {
	Index    int
	Name     string
	Edges    *[]Edge
	Shortest int
}

func (o *N2) Key() int {
	return o.Shortest
}

func (o *N2) SetKey(n int) {
	o.Shortest = n
}

type Edge struct {
	To     *N2
	Weight int
}

type H struct {
	KeyValue int
	Value    int
}

func (o *H) Key() int {
	return o.KeyValue
}

func (o *H) SetKey(n int) {
	o.KeyValue = n
}

type HeapItem interface {
	Key() int
	SetKey(int)
}

type Heap struct {
	Body []HeapItem
	Pos  int
}

func NewHeap(l int) *Heap {
	return &Heap{
		Body: make([]HeapItem, l),
		Pos:  0,
	}
}

func (o *Heap) Insert(h HeapItem) {
	o.Body[o.Pos] = h

	now := o.Pos

	for now != 0 {
		parent := (now - 1) / 2

		if o.Body[parent].Key() > o.Body[now].Key() {
			o.Body[now], o.Body[parent] = o.Body[parent], o.Body[now]
		}
		now = now / 2
	}

	o.Pos++
}

func (o *Heap) Decrease(h HeapItem, n int) {

	now := 0
	for i, item := range o.Body {
		if item == h {
			now = i
			item.SetKey(n)
			break
		}
	}

	for now != 0 {
		parent := (now - 1) / 2

		if o.Body[parent].Key() > o.Body[now].Key() {
			o.Body[now], o.Body[parent] = o.Body[parent], o.Body[now]
		}
		now = now / 2
	}
}

func (o *Heap) Pick() (HeapItem, error) {
	if o.Pos == 0 {
		return nil, fmt.Errorf("%v", "")
	}
	o.Pos--
	re := o.Body[0]
	o.Body[0] = o.Body[o.Pos]

	now := 0
	next := 1

	for next < o.Pos {
		nowValue := o.Body[now].Key()
		leftValue := o.Body[next].Key()
		rightValue := o.Body[next+1].Key()

		if leftValue > rightValue {
			next++
		}

		nextValue := o.Body[next].Key()

		if nowValue > nextValue {
			o.Body[now], o.Body[next] = o.Body[next], o.Body[now]
		} else {
			break
		}

		now = next
		next = now*2 + 1
	}

	return re, nil
}

func (o *Heap) SortedKeys() []int {
	re := make([]int, o.Pos)

	for i, n := range o.SortedBody() {
		re[i] = n.Key()
	}

	return re
}

func (o *Heap) SortedBody() []HeapItem {
	re := make([]HeapItem, o.Pos)

	for i, _ := range re {
		h, _ := o.Pick()
		re[i] = h
	}

	return re
}

func heapSort(data []HeapItem) *Heap {
	heap := NewHeap(len(data))

	for _, h := range data {
		heap.Insert(h)
	}

	return heap
}

func heapSort3(data []*H) *Heap {
	re := make([]HeapItem, len(data))

	for i, _ := range re {
		re[i] = data[i]
	}

	return heapSort(re)
}

func heapSort2(data []*N2) *Heap {
	re := make([]HeapItem, len(data))

	for i, _ := range re {
		re[i] = data[i]
	}

	return heapSort(re)
}

type Set struct {
	Body []*N2
	Map  map[string]bool
	Size int
}

func (o *Set) Set(nodes []*N2) *Set {
	o.Size = len(nodes)
	o.Body = make([]*N2, o.Size)
	copy(o.Body, nodes)
	o.Map = map[string]bool{}
	for _, n := range nodes {
		o.Map[n.Name] = true
	}
	return o
}

func (o *Set) HasAny() bool {
	return o.Size != 0
}

func (o *Set) Has(name string) bool {
	return o.Map[name]
}

func (o *Set) Del(node *N2) *N2 {
	newer := make([]*N2, o.Size-1)

	i := 0
	for _, n := range o.Body {
		if n.Name != node.Name {
			newer[i] = n
			i++
		}
	}
	o.Body = newer
	o.Size--
	o.Map[node.Name] = false
	return node
}

func dijkstra(nodes []*N2, start int) map[string]int {
	nearest := map[string]*N2{}
	rest := (&Set{}).Set(nodes)
	startNode := nodes[start]

	for _, n := range nodes {
		n.Shortest = math.MaxInt32
	}
	startNode.Shortest = 0
	heap := heapSort2(nodes)

	h, _ := heap.Pick()
	visited := h.(*N2)

	for {
		rest.Del(visited)
		heap = NewHeap(len(nodes))

		for _, edge := range *visited.Edges {
			nextNode := edge.To
			nextName := nextNode.Name

			if !rest.Has(nextNode.Name) {
				continue
			}

			pre := nearest[nextNode.Name]
			nextShortest := visited.Shortest + edge.Weight

			if pre == nil {
				nearest[nextName] = visited
				nextNode.Shortest = nextShortest
			} else {
				if nextShortest < nextNode.Shortest {
					nearest[nextName] = visited
					nextNode.Shortest = nextShortest
				}
			}

			heap.Insert(nextNode)
		}

		h, err := heap.Pick()
		if err != nil {
			break
		}
		visited = h.(*N2)
	}
	shortest := map[string]int{}

	for _, n := range nodes {
		shortest[n.Name] = n.Shortest
	}

	return shortest
}
