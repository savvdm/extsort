package main

import (
	"bufio"
	"container/heap"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
)

// heap element: a single line read from an input bunch
type Input struct {
	line    string
	scanner *bufio.Scanner
}

// read next line of input
// return true if the line is read
func (input *Input) next() bool {
	scan := input.scanner
	if scan.Scan() {
		input.line = scan.Text()
	} else {
		if err := scan.Err(); err != nil {
			log.Fatal(err)
		}
		return false // end of file
	}
	return true
}

// heap storage
type InputHeap []Input

// heap implementation
func (h InputHeap) Len() int           { return len(h) }
func (h InputHeap) Less(i, j int) bool { return h[i].line < h[j].line }
func (h InputHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *InputHeap) Push(x interface{}) {
	*h = append(*h, x.(Input))
}
func (h *InputHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// write buf to temp file
// rewind the file afterwards, to read from the beginning
func write(lines []string) (file *os.File) {
	file, err := ioutil.TempFile("", "sort")
	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(file)

	// write data
	for _, line := range lines {
		if _, err = fmt.Fprintln(w, line); err != nil {
			log.Fatal(err)
		}
	}

	// flush data to disk
	if err = w.Flush(); err != nil {
		log.Fatal(err)
	}

	// rewind to the beginning of the file
	if _, err := file.Seek(0, io.SeekStart); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	num := flag.Int("n", 5000, "Number of strings to read in one bunch")
	flag.Parse()

	buf := make([]string, 0, *num)

	files := make([]*os.File, 10)

	// read input in bunches, sort, and write each bunch to a temporary file
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		buf = append(buf, scanner.Text())
		if len(buf) == cap(buf) {
			sort.Strings(buf)
			files = append(files, write(buf))
			buf = buf[:0]
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// write remaining lines
	if len(buf) > 0 {
		sort.Strings(buf)
		files = append(files, write(buf))
	}

	// read first lines of each bunch into a heap, and sort it
	h := make(InputHeap, 0, len(files))
	for _, file := range files {
		scanner := bufio.NewScanner(file)
		input := Input{"", scanner}
		if input.next() {
			h = append(h, input)
		}
	}
	heap.Init(&h)

	// read line from the heap, replacing with the next one from the same bunch
	// repeat until all inputs are exausted
	for h.Len() > 0 {
		fmt.Println(h[0].line)
		if h[0].next() {
			// adjust heap
			heap.Fix(&h, 0)
		} else {
			// no more lines in this input bunch
			heap.Pop(&h)
		}
	}

	// clean up temp files
	for _, file := range files {
		if err := file.Close(); err != nil {
			log.Println(err)
		}
		if err := os.Remove(file.Name()); err != nil {
			log.Println(err)
		}
	}
}
