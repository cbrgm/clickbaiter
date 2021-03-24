package main

import (
	"flag"
	"fmt"
	"github.com/cbrgm/clickbaiter/clickbaiter"
	"math/rand"
	"time"
)

var hashtagsFlag = flag.Bool("hashtags", false, "use hashtags")

func main() {
	flag.Parse()
	rand.Seed(time.Now().Unix())

	b := *hashtagsFlag

	fmt.Println(b)

	cbg := clickbaiter.NewGenerator()
	cbg.UseHashtags(b)

	fmt.Println(cbg.RandomSentence())
}
