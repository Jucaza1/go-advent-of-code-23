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
	var currentNum []rune
	var skip int
	var gearStore = newAsterikStore()
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if !isADigit(mat[i][j]) {
				continue
			}
			currentNum = []rune{mat[i][j]}
			for k := j + 1; k < len(mat[0]); k++ {
				skip = k - j //save j skip for later add to j
				if !isADigit(mat[i][k]) {
					break
				}
				currentNum = append(currentNum, mat[i][k])
			}
			seekSymbolAsteriskAndStore(mat, i, j, currentNum, &gearStore)
			j += skip
			skip = 0
			// then j++
		}
	}
	return gearStore.sumAll2products()
}
func isADigit(r rune) bool {
	return r >= '0' && r <= '9'
}

type asteriskStore struct {
	store map[asteriskCoord][]int
}
type asteriskCoord struct {
	x int
	y int
}

func newAsterikStore()asteriskStore{
    return asteriskStore{
        store: make(map[asteriskCoord][]int),
    }
}
func (a *asteriskStore) storeVal(p asteriskCoord, val int) {
	a.store[p] = append(a.store[p], val)
}
func (a *asteriskStore) sumAll2products() int {
	var acc int
	for _, v := range a.store {
		if len(v) == 2 {
			acc += v[0] * v[1]
		}
	}
	return acc
}
func seekSymbolAsteriskAndStore(mat [][]rune, i int, j int, currentNum []rune, store *asteriskStore) {
	imax := len(mat)
	jmax := len(mat[0])
	nlen := len(currentNum)
	for t := i - 1; t < i+2; t++ {
		for k := j - 1; k < j+nlen+1; k++ {
			if k >= jmax || t >= imax || k < 0 || t < 0 || (t == i && k-j < nlen && k-j > -1) {
				continue
			}
			if mat[t][k] == '*' {
				currentInt, err := strconv.Atoi(string(currentNum))
				if err != nil {
					log.Fatalf("expected to be a number in the []rune, got %v\n", currentNum)
				}
                store.storeVal(asteriskCoord{x: t, y: k},currentInt)
			}
		}
	}
}
