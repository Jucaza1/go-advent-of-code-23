package main

import (
	"bufio"
	"fmt"
	"log"
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
	var lineAcc int
	for scanner.Scan() {
		lineAcc += processLine(scanner.Bytes())
	}
	fmt.Println("input processed -> ", lineAcc)
}
func processLine(l []byte) int {
	splits := strings.Split(string(l), ":")
	game := splits[0]
    gameSplit:= strings.Split(game, " ")
	gameID, err := strconv.Atoi(gameSplit[1])
	if err != nil {
		return 0
	}
	replacer := strings.NewReplacer("red", "12", "green", "13", "blue", "14")
	turns := strings.Split(replacer.Replace(splits[1]), ";")

	for _, turn := range turns {
		for _, color := range strings.Split(turn, ",") {
			colorSplit := strings.Split(strings.TrimSpace(color), " ")
			val, err := strconv.Atoi(colorSplit[0])
			if err != nil {
				return 0
			}
			limit, err := strconv.Atoi(colorSplit[1])
			if err != nil {
				return 0
			}
			if val > limit {
				return 0
			}
		}
	}
	return gameID
}
