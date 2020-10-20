package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func intArg(i int) int {
	input := os.Args[i]
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid number: %s\n", input)
		os.Exit(1)
	}
	return num
}

func main() {
	if len(os.Args) < 2 {
		println("Specify the number of strings to generate")
		os.Exit(1)
	}
	num := intArg(1)

	max := 160
	if len(os.Args) > 2 {
		max = intArg(2)
	}

	var seed int64 = 1
	if len(os.Args) > 3 {
		seed = int64(intArg(3))
	}
	rand.Seed(seed)

	for i := 0; i < num; i++ {
		strlen := rand.Intn(max)
		s := make([]byte, strlen)
		for i := 0; i < strlen; i++ {
			s[i] = byte(rand.Intn('z'-'a'+1) + 'a')
		}
		fmt.Println(string(s))
	}
}
