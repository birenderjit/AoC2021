package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func loadDataFromFile(name string) []string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var data []string

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

func part1(list []string) int {
	horizontal, depth := 0, 0
	for _, line := range list {
		var directions = strings.Split(line, " ")
		val, _ := strconv.Atoi(directions[1])
		//fmt.Println(directions[0], directions[1])
		switch directions[0] {
		case "forward":
			horizontal += val
		case "down":
			depth += val
		case "up":
			depth -= val
		}
	}
	return horizontal * depth
}

func part2(list []string) int {
	horizontal, depth, aim := 0, 0, 0
	for _, line := range list {
		var directions = strings.Split(line, " ")
		val, _ := strconv.Atoi(directions[1])
		//fmt.Println(directions[0], directions[1])
		switch directions[0] {
		case "forward":
			horizontal += val
			if aim > 0 {
				depth = depth + (aim * val)
			}
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}
	return horizontal * depth
}
func main() {
	//data := loadDataFromFile("sample.txt")
	data := loadDataFromFile("input.txt")

	fmt.Println("Part 1 -- ", part1(data))
	fmt.Println("Part 2 -- ", part2(data))
}
