package main

import (
	"testing"
	"fmt"
)

func TestCompute(t *testing.T) {
	rows := []struct {
		a  string
		b  string
		ex string
	}{
		{
			"GTACCGTCA",
			"CATCGA",
			"TCGA",
		},
	}

	for _, row := range rows {
		ac := lcs(row.a, row.b)

		if ac != row.ex {
			fmt.Printf("%+v \n", ac)
			t.Fail()
		}
	}
}

func TestCompute2(t *testing.T) {
	rows := []struct {
		a  string
		b  string
		ex [][]string
	}{
		{
			"ACAAGC",
			"CCGT",
			[][]string{{DEL, "A"}, {COP, "C"}, {DEL, "A"}, {REP, "A C"}, {COP, "G"}, {REP, "C T"}},
		},
	}

	for _, row := range rows {
		ac := replace(row.a, row.b)

		for i, _ := range ac {
			if row.ex[i][0] != ac[i][0] || row.ex[i][1] != ac[i][1] {
				fmt.Printf("%+v %+v \n", row.ex[i], ac[i])
				t.Fail()
			}
		}
	}
}

func TestCompute3(t *testing.T) {
	rows := []struct {
		a  string
		b  string
		ex []int
	}{
		{
			"GTAACAGTAAACG",
			"AAC",
			[]int{2, 9},
		},
	}

	for _, row := range rows {
		ac := pattern(row.a, row.b)

		if len(ac) == 0 {
			t.Fail()
		}

		for i, _ := range ac {
			if row.ex[i] != ac[i] {
				fmt.Printf("%+v %+v \n", row.ex[i], ac[i])
				t.Fail()
			}
		}
	}
}

func TestCompute4(t *testing.T) {
	rows := []struct {
		a  string
		ex []map[string]int
	}{
		{
			"ACACAGA",
			[]map[string]int{
				{"A": 1, "C": 0, "G": 0},
				{"A": 1, "C": 2, "G": 0},
				{"A": 3, "C": 0, "G": 0},
				{"A": 1, "C": 4, "G": 0},
				{"A": 5, "C": 0, "G": 0},
				{"A": 1, "C": 4, "G": 6},
				{"A": 7, "C": 0, "G": 0},
				{"A": 1, "C": 2, "G": 0},
			},
		},
	}

	for _, row := range rows {
		ac, m := nextStateTable(row.a)

		for i, _ := range ac {
			for k, _ := range m {
				if ac[i][k] != row.ex[i][k] {
					fmt.Printf("%+v %+v \n", ac[i][k], row.ex[i][k])
					t.Fail()
				}
			}
		}
	}
}

func TestCompute5(t *testing.T) {
	rows := []struct {
		a  string
		b  string
		ex int
	}{
		{"ACA", "ACG", 0},
		{"ACA", "A", 1},
		{"ACACGAA", "ACACA", 3},
		{"ACACGAA", "ACACG", 5},
		{"ACACGAA", "ACACGT", 0},
	}

	for _, row := range rows {
		ac := nextState(row.a, row.b)
		if ac != row.ex {
			fmt.Printf("%+v %+v \n", ac, row.ex)
			t.Fail()
		}
	}
}
