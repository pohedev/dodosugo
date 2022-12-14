package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"unicode/utf8"
)

func main() {
	var (
		dodosuko     = []string{"ドド", "スコ"}
		matchPattern = "011101110111"
	)

	rand.Seed(time.Now().UnixNano())

	for {
		var generatedPhrase, generatedPattern bytes.Buffer
		for i := 0; i < utf8.RuneCountInString(matchPattern); i++ {
			randomNum := rand.Intn(2)

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
