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

	// dir, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Current working directory:", dir)
    _,filename,_,ok := runtime.Caller(0)
    if !ok{
        log.Fatal("error: could not get caller filename")
    }
    basedir:= filepath.Dir(filename)
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

}

func processLine(l []byte) uint8 {
	str := string(l)
	var first uint8
	var last uint8
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			first = uint8(str[i] - '0')
			break
		}
	}
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] >= '0' && str[i] <= '9' {
			last = uint8(str[i] - '0')
			break
		}
	}
	return first*10 + last
}
