package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	gen := clickbaiter.NewGenerator()
	fmt.Println(gen.RandomSentence())
}
