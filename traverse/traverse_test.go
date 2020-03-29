package traverse

import (
	"fmt"
	"reflect"
	"testing"

	mathInt "github.com/ikngtty/benchmark-go-dfs/math/int"
	"github.com/ikngtty/benchmark-go-dfs/tree"
)

const bigTreeDepth = 20

func TestFindPath(t *testing.T) {
	functions := []struct {
		name string
		body func(*tree.Node, tree.NodeValue) []tree.NodeValue
	}{
		{"recursive", FindPathRec},
		{"loop", FindPathLoop},
	}

	depth := 5
	cases := []struct {
		name      string
		destValue tree.NodeValue
		want      []tree.NodeValue
	}{
		{"top", 1, []tree.NodeValue{1}},
		{"halfway", 10, []tree.NodeValue{1, 2, 5, 10}},
		{"most left leaf", 16, []tree.NodeValue{1, 2, 4, 8, 16}},
		{"most right leaf", 31, []tree.NodeValue{1, 3, 7, 15, 31}},
		{"not found", 42, nil},
	}

	for _, f := range functions {
		for _, c := range cases {
			caseName := fmt.Sprintf("%s:%s", f.name, c.name)
			t.Run(caseName, func(t *testing.T) {
				node := tree.GenerateATree(depth)
				got := f.body(node, c.destValue)
				if !reflect.DeepEqual(got, c.want) {
					t.Errorf("want: %#v, got: %#v", c.want, got)
				}
			})
		}
	}
}

func BenchmarkFindPathRec(b *testing.B) {
	node := tree.GenerateATree(bigTreeDepth)
	destValue := tree.NodeValue(mathInt.Pow(2, bigTreeDepth) - 1) // most right leaf
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindPathRec(node, destValue)
	}
}

func BenchmarkFindPathLoop(b *testing.B) {
	node := tree.GenerateATree(bigTreeDepth)
	destValue := tree.NodeValue(mathInt.Pow(2, bigTreeDepth) - 1) // most right leaf
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindPathLoop(node, destValue)
	}
}
