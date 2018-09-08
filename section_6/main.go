package main

import (
	"fmt"
)

func lcs(a string, b string) string {
	x := []rune(" " + a)
	y := []rune(" " + b)

	xl := len(x)
	yl := len(y)

	lcss := make([][]int, yl)
	lcss[0] = make([]int, xl)

	for yi := 1; yi < yl; yi++ {
		lcss[yi] = make([]int, xl)

		for xi := 1; xi < xl; xi++ {
			switch {
			case x[xi] == y[yi]:
				lcss[yi][xi] = lcss[yi-1][xi-1] + 1
			case lcss[yi-1][xi] > lcss[yi][xi-1]:
				lcss[yi][xi] = lcss[yi-1][xi]
			default:
				lcss[yi][xi] = lcss[yi][xi-1]
			}
		}
	}

	return pick(x, y, xl-1, yl-1, lcss, "")
}

const (
	DEL      = "del"
	INS      = "ins"
	REP      = "rep"
	COP      = "copy"
	END      = "end"
	DEL_COST = 2
	INS_COST = 2
	REP_COST = 1
	COP_COST = -1
)

func pick(x []rune, y []rune, xi, yi int, lcss [][]int, result string) string {
	switch {
	case lcss[yi][xi] == 0:
		return result
	case x[xi] == y[yi]:
		return pick(x, y, xi-1, yi-1, lcss, string(x[xi])+result)
	case lcss[yi-1][xi] > lcss[yi][xi-1]:
		return pick(x, y, xi, yi-1, lcss, result)
	default:
		return pick(x, y, xi-1, yi, lcss, result)
	}
}

func replace(a string, b string) [][]string {
	x := []rune(" " + b)
	y := []rune(" " + a)

	xl := len(x)
	yl := len(y)

	costs := make([][]int, yl)
	actions := make([][][]string, yl)

	for i := 0; i < yl; i++ {
		costs[i] = make([]int, xl)
		actions[i] = make([][]string, xl)
	}

	costs[0][0] = 0
	actions[0][0] = []string{END}

	for xi := 1; xi < xl; xi++ {
		costs[0][xi] = costs[0][xi-1] + INS_COST
		actions[0][xi] = []string{INS, string(x[xi])}
	}

	for yi := 1; yi < yl; yi++ {
		costs[yi][0] = costs[yi-1][0] + DEL_COST
		actions[yi][0] = []string{DEL, string(y[yi])}
	}

	for xi := 1; xi < xl; xi++ {
		for yi := 1; yi < yl; yi++ {
			cost := costs[yi-1][xi-1]

			if x[xi] == y[yi] {
				costs[yi][xi] = cost + COP_COST
				actions[yi][xi] = []string{COP, string(x[xi])}
			} else {
				costs[yi][xi] = cost + REP_COST
				actions[yi][xi] = []string{REP, fmt.Sprintf("%v %v", string(y[yi]), string(x[xi]))}
			}

			if costs[yi][xi] > costs[yi-1][xi]+2 {
				costs[yi][xi] = costs[yi-1][xi] + DEL_COST
				actions[yi][xi] = []string{DEL, string(y[yi])}
			}

			if costs[yi][xi] > costs[yi][xi-1]+2 {
				costs[yi][xi] = costs[yi][xi-1] + INS_COST
				actions[yi][xi] = []string{INS, string(x[xi])}
			}
		}
	}

	return trace(x, y, xl-1, yl-1, actions, [][]string{})
}

func trace(x []rune, y []rune, xi, yi int, actions [][][]string, result [][]string) [][]string {
	action := actions[yi][xi]
	switch {
	case action[0] == END:
		return result
	case action[0] == COP, action[0] == REP:
		return trace(x, y, xi-1, yi-1, actions, append([][]string{action}, result...))
	case action[0] == DEL:
		return trace(x, y, xi, yi-1, actions, append([][]string{action}, result...))
	default:
		return trace(x, y, xi, yi-1, actions, append([][]string{action}, result...))
	}
}

func pattern(a string, pat string) []int {
	text := []rune(a)
	table, chars := nextStateTable(pat)
	pl := len(pat)
	results := []int{}

	state := 0
	for i, char := range text {
		if !chars[string(char)] {
			state = 0
			continue
		}

		nextState := table[state][string(char)]
		if nextState == pl {
			results = append(results, i-state)
			state = 0
		} else {
			state = nextState
		}
	}

	return results
}

func nextState(expected string, actual string) int {
	ex := []rune(expected)
	ac := []rune(actual)
	acl := len(ac)

	l := acl

	for l > 0 {
		if string(ex[0:l]) == string(ac[acl-l:]) {
			return l
		}
		l--
	}

	return l
}

func nextStateTable(a string) ([]map[string]int, map[string]bool) {
	xs := []rune(a)
	chars := map[string]bool{}
	table := make([]map[string]int, len(xs))

	for _, x := range xs {
		chars[string(x)] = true
	}
	for i, _ := range table {
		table[i] = map[string]int{}
	}

	for i, _ := range table {
		table[i] = map[string]int{}
		ac := string(xs[0:i])

		for k, _ := range chars {
			next := nextState(a, fmt.Sprintf("%s%v", ac, k))
			table[i][string(k)] = next
		}
	}

	return table, chars
}
