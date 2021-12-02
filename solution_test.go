package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func TestT(t *testing.T) {
	file, err := os.Open("tcases.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tcase := strings.Fields(scanner.Text())
		runIt(tcase[0], tcase[1], t)
	}

}
