package main

import (
	"bufio"
	_ "container/heap"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

func writeBunch(lines []string) (file *os.File) {
	file, err := ioutil.TempFile("", "sort")
	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(file)

	// write data
	for _, str := range lines {
		if _, err = fmt.Fprintln(file, str); err != nil {
			log.Fatal(err)
		}
	}

	if err = w.Flush(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Written %s\n", file.Name())
	return
}

func main() {
	num := flag.Int("n", 10, "Number of strings to read in one bunch")
	flag.Parse()

	buf := make([]string, 0, *num)

	files := make([]*os.File, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buf = append(buf, scanner.Text())
		if len(buf) == cap(buf) {
			sort.Strings(buf)
			files = append(files, writeBunch(buf))
			buf = buf[:0]
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// clean up temp files
	for _, file := range files {
		if err := file.Close(); err != nil {
			log.Println(err)
		}
		if err := os.Remove(file.Name()); err != nil {
			log.Println(err)
		}
		log.Printf("Removed %s\n", file.Name())
	}
}
