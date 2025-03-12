package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	level := flag.String("level", "CRITICAL", "log level to filter for")
	flag.Parse()
	f, err := os.Open("./log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	bufferReader := bufio.NewReader(f)
	for line, err := bufferReader.ReadString('\n'); err == nil; line, err = bufferReader.ReadString('\n') {
		if strings.Contains(line, *level) {
			fmt.Println(line)
		}
	}
}
