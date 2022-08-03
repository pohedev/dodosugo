package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"unicode/utf8"
)

const (
	dodo = "ドド"
	suko = "スコ"
)

func main() {
	run()
}

func run() {
	var (
		dodosuko     = []string{dodo, suko}
		matchPattern = "011101110111"
	)

	rand.Seed(time.Now().UnixNano())

	for {
		var generatedPhrase, generatedPattern bytes.Buffer
		for i := 0; i < utf8.RuneCountInString(matchPattern); i++ {
			randomNum := rand.Intn(len(dodosuko))

			generatedPhrase.WriteString(dodosuko[randomNum])
			generatedPattern.WriteString(strconv.Itoa(randomNum))
		}

		fmt.Println(generatedPhrase.String())
		if generatedPattern.String() == matchPattern {
			fmt.Println("ラブ注入♡")
			break
		}
	}
}
