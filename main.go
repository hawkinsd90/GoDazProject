package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

func main() {
	var word, word2, samePosMatch, sameLetterMatch, noMatchString string
	var samePosIndex, sameLetterIndex, randomIndex, wordLenth int
	var samePos, sameLetter bool
	noMatches := make([]string, 0)

	//get liat of 5 letter words
	wordLenth = 5
	words := getWords(wordLenth)

	//get Word to guess
	randomIndex = rand.Intn(len(words))
	word = words[randomIndex]

	//get example guess case
	for !samePos || !sameLetter || word == word2 || samePosMatch == sameLetterMatch {
		randomIndex = rand.Intn(len(words))
		word2 = words[randomIndex]
		setHints(word, word2, &noMatches, &samePosMatch, &samePosIndex, &samePos, &sameLetterMatch, &sameLetterIndex, &sameLetter)
	}

	noMatchString = getNoMatchString(noMatches)
	samePos = false
	sameLetter = false

	//launch game
	fmt.Println("Welcome to Devon Hawkins' Wordle.")
	fmt.Printf("\nYou are given 5 tries to guess a %d letter word.\n", wordLenth)
	fmt.Printf("We'll start by example using the word '%s'\n%s\n", word2, word2)
	fmt.Printf("Answer has letter '%s' in the same position", samePosMatch)
	fmt.Printf("\nAnswer has letter '%s'but in a different position\n", sameLetterMatch)
	fmt.Println(noMatchString)
	fmt.Printf("Please enter a %d letter word.\n", wordLenth)

	scanner := bufio.NewScanner(os.Stdin)
	tries := 0
	maxTries := 5
	length := 0
	guess := ""
	for scanner.Scan() && tries < maxTries {
		guess = strings.ToLower(scanner.Text())
		length = len(guess)
		if length != wordLenth {
			fmt.Println("Word must be 5 characters long")
			continue
		}
		if guess == word {
			fmt.Println("You Win!")
			break
		}

		setHints(word, guess, &noMatches, &samePosMatch, &samePosIndex, &samePos, &sameLetterMatch, &sameLetterIndex, &sameLetter)
		fmt.Println(word2)
		if samePos {
			fmt.Printf("\nAnswer has the letter '%s' in the same position", samePosMatch)
		}
		if sameLetter {
			fmt.Printf("\nAnswer has the letter '%s' but in a different position\n", sameLetterMatch)
		}
		if len(noMatches) > 0 {
			fmt.Println(getNoMatchString(noMatches))
		}

		tries++
		fmt.Printf("\n%d tries left.\n", int(maxTries-tries))
	}
	if scanner.Err() != nil {
		/*handle error*/
		fmt.Println("ERROR! YOU BROKE IT")
	}

	if tries == 5 && guess != word {
		fmt.Printf("\n\nGAME OVER!\n\nThe correct answer is '%s'", word)

	}
}

func getWords(wordLength int) (words []string) {

	words = make([]string, 0, 10000)
	word := ""
	body, err := os.Open("words_alpha.txt") //io.ReadFile("words_alpha.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	// fmt.Println(string(body))

	//get 5 letter words
	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		word = scanner.Text()
		if len(word) != 5 {
			continue
		}
		words = append(words, word)
	}
	defer body.Close()
	return words
}

func setHints(baseWord string, compWord string, noMatches *[]string, matchLetter *string, matchIndex *int, matchBool *bool, sameLetterDiffPos *string, sameLetterDiffPosIndex *int, sameLetterDiffPosBool *bool) {
	match := false
	*noMatches = make([]string, 0)
	*matchBool = false
	*sameLetterDiffPosBool = false

	//fmt.Printf("\n comparing '%s' to '%s'", compWord, baseWord)
	for i, x := range compWord {
		//fmt.Println(i, " => ", string(c))
		match = false
		for j, y := range baseWord {
			if x == y {
				match = true
			}
			if !*matchBool && string(x) == string(y) && int(i) == int(j) {
				*matchLetter = string(x)
				*matchIndex = int(i)
				*matchBool = true
				break
			}

			if !*sameLetterDiffPosBool && string(x) == string(y) && int(i) != int(j) {
				*sameLetterDiffPos = string(y)
				*sameLetterDiffPosIndex = int(j)
				*sameLetterDiffPosBool = true
				//fmt.Printf("\n sameLetterDiffPosBool match on '%s' at '%d'", string(y), int(j))
				break
			}
		}

		if !match {
			*noMatches = append(*noMatches, string(x))
		}
	}
}

func getNoMatchString(noMatches []string) (noMatch string) {

	var buffer bytes.Buffer

	for i, noMatch := range noMatches {
		if i == 0 {
			buffer.WriteString(fmt.Sprintf("'%s'", noMatch))
		} else {
			buffer.WriteString(fmt.Sprintf(", '%s'", noMatch))
		}
	}

	if len(buffer.String()) > 0 {
		buffer.WriteString(" are not in the string.")
	}

	return buffer.String()

}
