package main

import (
	"fmt"
	"github.com/cbrgm/clickbaiter/clickbaiter"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	gen := clickbaiter.NewGenerator()
	fmt.Println(gen.RandomSentence())
}
