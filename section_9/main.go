package main

import (
	"fmt"
	"github.com/k0kubun/pp"
)

func huffman(a string) string {
	buildHuffman(a)

	return ""
}

func buildHuffman(a string) *HuffmanNode {
	bytes := []byte(a)
	counts := map[uint8]int{}

	for _, b := range bytes {
		counts[b]++
	}

	heap := NewHeap(len(counts))

	for k, v := range counts {
		heap.Insert(&HuffmanNode{
			Byte:  k,
			Count: v,
			Per:   float32(v) / float32(len(bytes)),
		})
	}

	var root *HuffmanNode
	for {
		n1, _ := heap.Pick()
		n2, err2 := heap.Pick()

		if err2 != nil {
			root = n1.(*HuffmanNode)
			break
		}

		nn1 := n1.(*HuffmanNode)
		nn2 := n2.(*HuffmanNode)

		if nn2.Per < nn1.Per {
			nn2, nn1 = nn1, nn2
		}

		heap.Insert(&HuffmanNode{
			Zero:  nn1,
			One:   nn2,
			Count: nn1.Count + nn2.Count,
			Per:   nn1.Per + nn2.Per,
		})
	}

	pp.Println(root)
	hmap := buildHuffmanMap(root, 0, 0, map[byte]MapBody{})

	for k, v := range hmap {
		fmt.Printf(fmt.Sprintf("%%v %%0%db\n", v.Length), string(k), v.Code)
	}

	return root
}

func buildHuffmanMap(node *HuffmanNode, pre int, length int, m map[byte]MapBody) map[byte]MapBody {
	if node == nil {
		return m
	}

	if node.Leaf() {
		m[node.Byte] = MapBody{
			Code:   pre,
			Length: length,
		}
		return m
	}

	buildHuffmanMap(node.Zero, pre<<1, length+1, m)
	buildHuffmanMap(node.One, (pre<<1)+1, length+1, m)

	return m
}

type MapBody struct {
	Code   int
	Length int
}

type HuffmanNode struct {
	Byte  uint8
	Count int
	Per   float32
	Zero  *HuffmanNode
	One   *HuffmanNode
}

func (o *HuffmanNode) Leaf() bool {
	return o.Zero == nil && o.One == nil
}

func (o *HuffmanNode) Key() int {
	return int(o.Per * 100)
}

func (o *HuffmanNode) SetKey(n int) {
	o.Count = n
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
