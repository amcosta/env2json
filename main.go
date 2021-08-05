package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	envs := make(map[string]string)
	filepath := os.Args[1]

	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		key, value := SplitValue(line)
		envs[key] = value
	}

	outputRaw, err := json.Marshal(envs)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(outputRaw))
}

func SplitValue(line string) (string, string) {
	values := strings.Split(line, "=")
	return values[0], strings.Trim(values[1], "\"")
}
