package main

import (
	"bufio"
	_ "container/heap"
	"flag"
	"fmt"
	"os"
	"sort"
)

type StrHeap []string

func (h StrHeap) Len() int           { return len(h) }
func (h StrHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h StrHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *StrHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}

func (h *StrHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// exit with the specified code in case of error
func check(e error, code int) {
	if e != nil {
		fmt.Println(e)
		os.Exit(code)
	}
}

func main() {
	num := flag.Int("n", 10, "Number of strings to read in one bunch")
	flag.Parse()

	buf := make([]string, 0, *num)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buf = append(buf, scanner.Text())
		if len(buf) == cap(buf) {
			break
		}
	}
	check(scanner.Err(), 4)

	sort.Strings(buf)

	for _, v := range buf {
		fmt.Println(v)
	}
}
