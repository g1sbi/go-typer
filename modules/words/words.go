package words

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Generate() string {
	words, err := readWordsFromFile("wordsList")
	if err != nil {
		log.Fatal("Error reading words from file:", err)
	}

	rand.Seed(time.Now().UnixNano())
	randomWords := getRandomWords(words, 35)

	output := strings.Join(randomWords, " ")
	return output
}

func readWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return words, nil
}

func getRandomWords(words []string, count int) []string {
	var randomWords []string
	for i := 0; i < count; i++ {
		randomIndex := rand.Intn(len(words))
		randomWords = append(randomWords, words[randomIndex])
	}

	return randomWords
}
