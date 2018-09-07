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
