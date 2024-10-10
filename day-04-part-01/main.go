package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func main() {

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("error: could not get caller filename")
	}
	basedir := filepath.Dir(filename)
	// file, err := os.Open(basedir + "/input-test.txt") //should result in 281
	file, err := os.Open(basedir + "/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var acc int
	for i := 0; scanner.Scan(); i++ {
		acc += processLine(scanner.Text())
	}
	fmt.Println("input processed -> ", acc)
}

func processLine(l string) int {
	splitedCard := strings.Split(l, ":")
	numbersSplited := strings.Split(splitedCard[1], "|")
    //replace all double spaces with spaces and split by space
	winnerText := strings.Split(strings.ReplaceAll(strings.TrimSpace(numbersSplited[0]), "  ", " "), " ")
	myNumbersText := strings.Split(strings.ReplaceAll(strings.TrimSpace(numbersSplited[1]), "  ", " "), " ")
	winnerSlice := make([]int, 0, 10)
	myNumbersSlice := make([]int, 0, 25)
	for _, n := range winnerText {
		val, err := strconv.Atoi(n)
		if err != nil {
			panic("expected to be a number")
		}
		winnerSlice = append(winnerSlice, val)
	}
	for _, n := range myNumbersText {
		val, err := strconv.Atoi(n)
		if err != nil {
			panic("expected to be a number")
		}
		myNumbersSlice = append(myNumbersSlice, val)
	}
	count := 0
	for _, n := range myNumbersSlice {
		for _, w := range winnerSlice {
			if n == w {
				count++
			}
		}
	}
	if count == 0 {
		return 0
	}
	return int(math.Pow(2, float64(count-1)))
}
