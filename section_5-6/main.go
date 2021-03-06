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
	From   *N2
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
		} else {
			break
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
	startNode := nodes[start]

	for _, n := range nodes {
		n.Shortest = math.MaxInt32
	}
	startNode.Shortest = 0
	heap := heapSort2(nodes)

	for {
		h, err := heap.Pick()
		if err != nil {
			break
		}
		visited := h.(*N2)

		for _, edge := range *visited.Edges {
			nextNode := edge.To
			nextName := nextNode.Name

			nextShortest := visited.Shortest + edge.Weight

			if nextShortest < nextNode.Shortest {
				nearest[nextName] = visited
				heap.Decrease(nextNode, nextShortest)
			}
		}
	}
	shortest := map[string]int{}

	for _, n := range nodes {
		shortest[n.Name] = n.Shortest
	}

	return shortest
}

func bellmanFord(nodes []*N2, start int) (map[string]*N2, map[string]int) {
	nearest := map[string]*N2{}
	shortest := map[string]int{}
	for _, n := range nodes {
		shortest[n.Name] = math.MaxInt32
	}
	startNode := nodes[start]
	shortest[startNode.Name] = 0

	for range nodes {
		relaxes(nodes, nearest, shortest)
	}

	return nearest, shortest
}

func relaxes(nodes []*N2, nearest map[string]*N2, shortest map[string]int) {
	for _, visited := range nodes {
		for _, edge := range *visited.Edges {
			nextNode := edge.To
			nextName := nextNode.Name

			nextShortest := shortest[visited.Name] + edge.Weight

			if nextShortest < shortest[nextName] {
				nearest[nextName] = visited
				shortest[nextName] = nextShortest
			}
		}
	}
}

type Cycle struct {
	Name string
	Next *Cycle
	Step int
}

func findNegativeCycle(nodes []*N2, nearest map[string]*N2, shortest map[string]int) []string {
	base := map[string]int{}
	for k, n := range shortest {
		base[k] = n
	}
	relaxes(nodes, nearest, shortest)

	paths := map[string]bool{}

	now := nodes[0]
	for {
		if paths[now.Name] {
			break
		}

		paths[now.Name] = true
		now = nearest[now.Name]
	}

	cycleCheck := map[string]bool{}
	var cycle *Cycle
	var tail *Cycle
	for {
		if cycleCheck[now.Name] {
			break
		}

		cycleCheck[now.Name] = true

		if cycle == nil {
			cycle = &Cycle{Name: now.Name, Step: 1}
			tail = cycle
		} else {
			cycle.Next = &Cycle{Name: now.Name, Step: cycle.Step + 1}
			cycle = cycle.Next
		}

		now = nearest[now.Name]
	}

	steps := make([]string, cycle.Step)
	i := 0
	for i < cycle.Step {
		steps[i] = tail.Name

		tail = tail.Next
		i++
	}

	return steps
}

func floyd(nodes [][]int) [][]int {
	count := len(nodes) + 1

	shortest := make([][][]int, count)
	pred := make([][][]*int, count)

	for i, _ := range shortest {
		shortest[i] = make([][]int, count)
		pred[i] = make([][]*int, count)

		for j, _ := range shortest {
			shortest[i][j] = make([]int, count)
			pred[i][j] = make([]*int, count)
		}
	}

	for i := 1; i < count; i++ {
		for j := 1; j < count; j++ {
			w := nodes[i-1][j-1]
			shortest[i][j][0] = w

			if w != math.MaxInt32 && w != 0 {
				pred[i][j][0] = ip(i)
			}
		}
	}

	for x := 1; x < count; x++ {
		for i := 1; i < count; i++ {
			for j := 1; j < count; j++ {
				next := shortest[i][x][x-1] + shortest[x][j][x-1]

				if shortest[i][j][x-1] > next {
					shortest[i][j][x] = next
					pred[i][j][x] = ip(x)
				} else {
					shortest[i][j][x] = shortest[i][j][x-1]
					pred[i][j][x] = pred[i][j][x-1]
				}
			}
		}
	}

	l := count - 1
	shortestResult := make([][]int, l)
	for i, _ := range shortestResult {
		shortestResult[i] = make([]int, l)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			shortestResult[i][j] = shortest[i+1][j+1][l]
		}
	}

	return shortestResult
}

func ip(i int) *int {
	return &i
}

func printShortest2(shortest [][][]int) {
	fmt.Println("shortest")
	for i, l := range shortest {
		if i == 0 {
			continue
		}
		fmt.Println(l)
	}
	fmt.Print("\n")
}

func printShortest(shortest [][][]int, x int) {
	fmt.Println("shortest")
	for i, l := range shortest {
		if i == 0 {
			continue
		}
		for ii, ll := range l {
			if ii == 0 {
				continue
			}
			v := ll[x]
			if v >= math.MaxInt32 {
				fmt.Print("- ")
			} else {
				fmt.Print(v, " ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func printPred(shortest [][][]*int, x int) {
	fmt.Println("pred")
	for i, l := range shortest {
		if i == 0 {
			continue
		}
		for ii, ll := range l {
			if ii == 0 {
				continue
			}
			v := ll[x]
			if v == nil {
				fmt.Print("- ")
			} else {
				fmt.Print(*v, " ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
