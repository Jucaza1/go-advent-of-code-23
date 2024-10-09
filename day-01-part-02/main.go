package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

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
	var lineAcc uint32
	for scanner.Scan() {
		lineAcc += uint32(processLine(scanner.Bytes()))
	}
	fmt.Println("input processed -> ", lineAcc)
	//    res:= processLine([]byte("twone"))
	// fmt.Println("input processed -> ", res) //should be 21

}
func processLine(b []byte) int {
	pool := [20]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
		"0",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"7",
		"8",
		"9",
	}
	var indexFirst int = math.MaxInt
	var indexLast int = -1
	var valueFirst int
	var valueLast int
	str := string(b)
	for i, p := range pool {
		indexf := strings.Index(str, p)
		indexl := strings.LastIndex(str, p)
		if indexf < indexFirst && indexf != -1 {
			indexFirst = indexf
			valueFirst = i % 10
		}
		if indexl > indexLast && indexl != -1 {
			indexLast = indexl
			valueLast = i % 10
		}
	}
	return valueFirst*10 + valueLast
}
