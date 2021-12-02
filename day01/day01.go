package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func loadDataFromFile(name string) []int {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var data []int

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln(err)
		}
		data = append(data, num)
	}

	return data
}

func part1(data []int) int {
	counter := 0
	for i := 0; i < len(data)-1; i++ {
		if data[i] < data[i+1] {
			counter++
		}
	}
	return counter
}

func part2(data []int) int {
	counter := 0
	for i := 0; i < len(data)-3; i++ {
		if data[i]+data[i+1]+data[i+2] < data[i+1]+data[i+2]+data[i+3] {
			counter++
		}
	}
	return counter
}

func main() {
	data := loadDataFromFile("input.txt")
	//data := loadDataFromFile("input_sample.txt")

	fmt.Println(part1(data))
	fmt.Println(part2(data))

}
