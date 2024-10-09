package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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
	var mat = make([][]rune, 0, 140)
	var result int
	for i := 0; scanner.Scan(); i++ {
		mat = append(mat, []rune(scanner.Text()))
	}
	result = addParts(mat)
	fmt.Println("input processed -> ", result)
}

func addParts(mat [][]rune) int {
	var acc int
	var currentNum []rune
	var skip int
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if !isADigit(mat[i][j]) {
				continue
			}
			currentNum = []rune{mat[i][j]}
			for k := j+1; k < len(mat[0]); k++ {
				skip = k - j //save j skip for later add to j
				if !isADigit(mat[i][k]) {
					break
				}
				currentNum = append(currentNum, mat[i][k])
			}
			if hasSymbol(mat, i, j, len(currentNum)) {
				currentInt, err := strconv.Atoi(string(currentNum))
				if err != nil {
					log.Fatalf("expected to be a number in the []rune, got %v\n", currentNum)
				}
				acc += currentInt
			}
			j += skip
			skip = 0
			// then j++
		}
	}
	//find number
	//collect it
	//look for a symbol in the surroundings
	return acc
}
func isADigit(r rune) bool {
	return r >= '0' && r <= '9'
}
func hasSymbol(mat [][]rune, i int, j int, nlen int) bool {
	imax := len(mat)
	jmax := len(mat[0])
	for t := i - 1; t < i+2; t++ {
		for k := j - 1; k < j+nlen+1; k++ {
			if k >= jmax || t >= imax || k < 0 || t < 0 || (t == i && k-j < nlen && k-j > -1) {
				continue
			}
			if mat[t][k] != '.' && !isADigit(mat[t][k]) {
				return true
			}
		}
	}
	return false
}
