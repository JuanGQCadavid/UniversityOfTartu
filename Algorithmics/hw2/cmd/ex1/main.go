package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const (
	PYTHON_FILE_NAME = "repositories/base_python_code.py"
)

var (
	QUEUE_NUMBER = 5
	STACK_NUMBER = 3
)

func main() {
	cmd := exec.Command("python3", PYTHON_FILE_NAME, strconv.Itoa(QUEUE_NUMBER), strconv.Itoa(STACK_NUMBER))
	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalln(err.Error())
	}

	response := strings.Split(string(out), "\n")
	for i := 0; i < 10; i++ {
		log.Println(response[i])
	}
}
