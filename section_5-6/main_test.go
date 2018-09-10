package main

import (
	"testing"
	"math/rand"
	"github.com/k0kubun/pp"
	"math"
)

func shuffle(data []*Node) []*Node {
	n := len(data)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}

	return data
}

func eq(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, an := range a {
		if an != b[i] {
			return false
		}
	}
	return true
}

func eq2(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, an := range a {
		if an != b[i] {
			return false
		}
	}
	return true
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
	namePrinter(topologicalSort(nodes[1:]))
}

type N struct {
	Name   string
	Weight int
}

func TestCompute2(t *testing.T) {
	names := []struct {
		Name   string
		Weight int
		Nears  []string
	}{
		{"マリネードを合わせる", 2, []string{"チキンをマリネードにつける"}},
		{"にんにくを刻む", 4, []string{"にんにくと生姜を加える"}},
		{"生姜を刻む", 3, []string{"にんにくと生姜を加える"}},
		{"人参を刻む", 4, []string{"人参セロリピーナッツに火を入れる"}},
		{"セロリを刻む", 3, []string{"人参セロリピーナッツに火を入れる"}},
		{"ピーナッツを洗う", 2, []string{"人参セロリピーナッツに火を入れる"}},
		{"クッキングソースを合わせる", 3, []string{"クッキングソースをかける"}},
		{"チキンを焼く", 6, []string{"チキンをマリネードにつける"}},
		{"チキンをマリネードにつける", 15, []string{"チキンに少し火を入れる"}},
		{"チキンに少し火を入れる", 4, []string{"にんにくと生姜を加える"}},
		{"にんにくと生姜を加える", 1, []string{"チキンを仕上げる"}},
		{"チキンを仕上げる", 2, []string{"チキンを外す"}},
		{"チキンを外す", 1, []string{"人参セロリピーナッツに火を入れる"}},
		{"人参セロリピーナッツに火を入れる", 4, []string{"チキンを戻す"}},
		{"チキンを戻す", 1, []string{"クッキングソースをかける"}},
		{"クッキングソースをかける", 1, []string{"ソースが濃くなるまで火を入れる"}},
		{"ソースが濃くなるまで火を入れる", 3, []string{"料理を火から外す"}},
		{"料理を火から外す", 1, []string{}},
	}
	nodes := make([]*Node, len(names))

	for i, n := range names {
		nodes[i] = NewNodeWithWeight(n.Name, n.Weight)
	}

	for i, n := range names {
		for _, nn := range n.Nears {
			for _, node := range nodes {
				if node.Name == nn {
					nodes[i].AddNear(node)
				}
			}
		}
	}

	shuffle(nodes)

	re := criticalPath(nodes)
	pp.Println(re.PathList())

	if re.Weight != 39 {
		t.Fail()
	}
}

func TestCompute3(t *testing.T) {
	hs := []*H{
		{0, 0},
		{2, 0},
		{8, 0},
		{1, 0},
		{3, 0},
		{0, 0},
	}
	re := heapSort3(hs)

	if !eq(re.SortedKeys(), []int{0, 0, 1, 2, 3, 8}) {
		t.Fail()
	}
	re2 := heapSort3(hs)
	re2.Decrease(hs[2], 1)

	if !eq(re2.SortedKeys(), []int{0, 0, 1, 1, 2, 3}) {
		t.Fail()
	}
}

func TestCompute4(t *testing.T) {
	nodes := map[string]*N2{
		"s": {0, "s", nil, 0},
		"t": {1, "t", nil, 0},
		"x": {2, "x", nil, 0},
		"y": {3, "y", nil, 0},
		"z": {4, "z", nil, 0},
	}

	nodes["s"].Edges = &[]Edge{{nodes["t"], 6, nil}, {nodes["y"], 4, nil}}
	nodes["t"].Edges = &[]Edge{{nodes["x"], 3, nil}, {nodes["y"], 2, nil}}
	nodes["x"].Edges = &[]Edge{{nodes["z"], 4, nil}}
	nodes["y"].Edges = &[]Edge{{nodes["t"], 1, nil}, {nodes["x"], 9, nil}, {nodes["z"], 3, nil}}
	nodes["z"].Edges = &[]Edge{{nodes["s"], 2, nil}, {nodes["x"], 5, nil}}

	ns := make([]*N2, 5)
	for _, n := range nodes {
		ns[n.Index] = n
	}

	ex := map[string]int{
		"s": 0,
		"t": 5,
		"x": 8,
		"y": 4,
		"z": 7,
	}

	for k, n := range dijkstra(ns, 0) {
		if ex[k] != n {
			t.Fail()
		}
	}
}

func TestCompute5(t *testing.T) {
	nodes := map[string]*N2{
		"s": {0, "s", nil, 0},
		"t": {1, "t", nil, 0},
		"x": {2, "x", nil, 0},
		"y": {3, "y", nil, 0},
		"z": {4, "z", nil, 0},
	}

	nodes["s"].Edges = &[]Edge{{nodes["t"], 6, nodes["s"]}, {nodes["y"], 7, nodes["s"]}}
	nodes["t"].Edges = &[]Edge{{nodes["x"], 5, nodes["t"]}, {nodes["y"], 8, nodes["t"]}, {nodes["z"], -4, nodes["t"]}}
	nodes["x"].Edges = &[]Edge{{nodes["t"], -2, nodes["x"]}}
	nodes["y"].Edges = &[]Edge{{nodes["x"], -3, nodes["y"]}, {nodes["z"], 9, nodes["y"]}}
	nodes["z"].Edges = &[]Edge{{nodes["s"], 2, nodes["z"]}, {nodes["x"], 7, nodes["z"]}}

	ns := make([]*N2, 5)
	for _, n := range nodes {
		ns[n.Index] = n
	}

	ex := map[string]int{
		"s": 0,
		"t": 2,
		"x": 4,
		"y": 7,
		"z": -2,
	}

	_, shortest := bellmanFord(ns, 0)

	for k, n := range shortest {
		if ex[k] != n {
			t.Fail()
		}
	}
}
func TestCompute6(t *testing.T) {
	nodes := map[string]*N2{
		"s": {0, "s", nil, 0},
		"t": {1, "t", nil, 0},
		"x": {2, "x", nil, 0},
		"y": {3, "y", nil, 0},
		"z": {4, "z", nil, 0},
	}

	nodes["s"].Edges = &[]Edge{{nodes["t"], 6, nodes["s"]}, {nodes["y"], 7, nodes["s"]}}
	nodes["t"].Edges = &[]Edge{{nodes["x"], 5, nodes["t"]}, {nodes["y"], 8, nodes["t"]}, {nodes["z"], -4, nodes["t"]}}
	nodes["x"].Edges = &[]Edge{{nodes["t"], -2, nodes["x"]}}
	nodes["y"].Edges = &[]Edge{{nodes["x"], -3, nodes["y"]}, {nodes["z"], 9, nodes["y"]}}
	nodes["z"].Edges = &[]Edge{{nodes["s"], 2, nodes["z"]}, {nodes["x"], 5, nodes["z"]}}

	ns := make([]*N2, 5)
	for _, n := range nodes {
		ns[n.Index] = n
	}

	nearest, shortest := bellmanFord(ns, 0)

	if !eq2(findNegativeCycle(ns, nearest, shortest), []string{"z", "t", "x"}) {
		t.Fail()
	}
}

func TestCompute7(t *testing.T) {
	nodes := [][]int{
		{0, 3, 8, math.MaxInt32},
		{math.MaxInt32, 0, math.MaxInt32, 1},
		{math.MaxInt32, 4, 0, math.MaxInt32},
		{2, math.MaxInt32, -5, 0},
	}

	ex := [][]int{
		{0, 3, -1, 4},
		{3, 0, -4, 1},
		{7, 4, 0, 5},
		{2, -1, -5, 0},
	}

	for i, n := range floyd(nodes) {
		if !eq(ex[i], n) {
			t.Fail()
		}
	}
}
