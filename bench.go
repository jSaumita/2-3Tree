package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
	"github.com/jsaumap/2-3Tree"

	st "github.com/jsaumap/2-3Tree/2-3Tree.go"
)

var n int

func init() {
	flag.IntVar(&n, "n", 10000, "Number of elements to insert")
}

func main() {
	flag.Parse()
	elements := make([]int, n)
	for i := range elements {
		elements[i] = rand.Int()
	}

	ini := time.Now()
	var t treebench.ST
	t = st.New()
	d := time.Now().Sub(ini)
	fmt.Printf("Tree created [%v]\n", d)

	fmt.Printf("Inserting data")
	ini = time.Now()
	for _, v := range elements {
		t.Insert(v)
	}
	d = time.Now().Sub(ini)
	fmt.Printf(" [%v]\n", d)

	fmt.Printf("Searching data")
	ini = time.Now()
	for _, v := range elements {
		if !t.Search(v) {
			fmt.Printf(" ... FAILED!\n")
			fmt.Printf("%d wasn't found!!\n", v)
		}
	}
	d = time.Now().Sub(ini)
	fmt.Printf(" [%v]\n", d)
}
