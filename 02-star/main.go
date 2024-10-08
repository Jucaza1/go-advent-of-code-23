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
    res:= processLine([]byte("twone"))
	fmt.Println("input processed -> ", res)

}
