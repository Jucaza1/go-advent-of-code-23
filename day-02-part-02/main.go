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
	turns := strings.Split(splits[1], ";")
    colorMin:=map[string]int{
        "red":0,
        "green":0,
        "blue":0,
    }

	for _, turn := range turns {
		for _, color := range strings.Split(turn, ",") {
			colorSplit := strings.Split(strings.TrimSpace(color), " ")
			val, err := strconv.Atoi(colorSplit[0])
			if err != nil {
				return 0
			}
            if colorMin[colorSplit[1]]< val{
                colorMin[colorSplit[1]]= val
            }
		}
	}
    power:= colorMin["red"]*colorMin["green"]*colorMin["blue"]
	return power
}
