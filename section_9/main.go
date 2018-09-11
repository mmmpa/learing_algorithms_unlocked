package main

import (
	"fmt"
	"math"
)

func huffman(a string) ([]byte, *HuffmanNode, uint8) {
	tree, byteMap := buildHuffman(a)

	bytes := []byte(a)

	c := make([]byte, 0, 100)
	s := uint8(0)
	p := uint8(8)
	for _, b := range bytes {
		v := byteMap[b]

		for i := 0; i < int(v.Length); i++ {
			p--
			b := (v.Code >> uint8(i)) & 1
			s += uint8(b << p)

			if p == 0 {
				c = append(c, s)
				p = uint8(8)
				s = uint8(0)
			}
		}
	}

	if p != 0 {
		c = append(c, s)
	}

	return c, tree, p
}

func deHuffman(bytes []byte, tree *HuffmanNode, rest uint8) string {
	result := make([]byte, 0, 100)

	now := tree
	l := len(bytes) - 1
	for ii, by := range bytes {
		top := 0
		if ii == l {
			top = int(rest)
		}

		for i := 7; i >= top; i-- {
			b := (by >> uint8(i)) & 1

			if b == 0 {
				now = now.Zero
			} else {
				now = now.One
			}

			if now.Leaf() {
				result = append(result, now.Byte)
				now = tree
			}
		}
	}

	return string(result)
}

func buildHuffman(a string) (*HuffmanNode, map[byte]MapBody) {
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

	hmap := buildHuffmanMap(root, 0, 0, map[byte]MapBody{})

	return root, hmap
}

func buildHuffmanMap(node *HuffmanNode, pre int, length uint, m map[byte]MapBody) map[byte]MapBody {
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

	buildHuffmanMap(node.Zero, pre, length+1, m)
	buildHuffmanMap(node.One, pre+(1<<length), length+1, m)

	return m
}

type MapBody struct {
	Code   int
	Length uint
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

func uintPointer(n uint32) *uint32 {
	return &n
}

func initialMap() (map[string]*uint32, []string) {
	m := map[string]*uint32{}
	m2 := make([]string, math.MaxUint16)
	for i := uint32(0); i < math.MaxUint8; i++ {
		m[string(byte(i))] = uintPointer(i)
		m2[i] = string(byte(i))
	}

	return m, m2
}

func lzw(a string) []uint32 {
	m, _ := initialMap()
	bytes := []byte(a)
	result := make([]uint32, 0, 100)

	initial := uint32(math.MaxUint8)

	pre := string(bytes[0])
	for i := 1; i < len(bytes); i++ {
		now := string(bytes[i])

		if m[pre+now] == nil {
			result = append(result, *m[pre])
			initial++
			m[pre+now] = uintPointer(initial)
			pre = now
		} else {
			pre = pre + now
		}
	}
	result = append(result, *m[pre])

	return result
}

func delzw(bytes []uint32) string {
	s, m := initialMap()

	initial := uint32(math.MaxUint8)

	result := m[bytes[0]]
	pre := m[bytes[0]]
	nowHead := ""
	for i := 1; i < len(bytes); i++ {
		now := m[bytes[i]]

		if now == "" {
			built := m[bytes[i-1]] + head(m[bytes[i-1]])
			result += built
			now = built
		} else {
			result += now
		}

		nowHead = head(now)
		if s[pre+nowHead] == nil {
			initial++
			s[pre+nowHead] = uintPointer(initial)
			m[initial] = pre + nowHead
			pre = now
		} else {
			pre = pre + now
		}
	}

	return result
}

func head(s string) string {
	if s == "" {
		return s
	}

	return string([]byte(s)[0])
}
