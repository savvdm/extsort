package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {

	num := flag.Int("n", 10, "Number of strings to generate")
	max := flag.Int("m", 120, "Max string length")
	seed := flag.Int64("s", 1, "Random generator's seed")

	flag.Parse()

	rand.Seed(*seed)

	for i := 0; i < *num; i++ {
		strlen := rand.Intn(*max)
		s := make([]byte, strlen)
		for i := 0; i < strlen; i++ {
			s[i] = byte(rand.Intn('z'-'a'+1) + 'a')
		}
		fmt.Println(string(s))
	}
}
