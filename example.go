package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func HighestYear(yearStrings []string) (maxYear int, numberAlive int) {
	m := make(map[int]int)

	for _, v := range yearStrings {
		yearRange := strings.Split(v, ":")
		start, _ := strconv.Atoi(yearRange[0])
		end, _ := strconv.Atoi(yearRange[1])
		m[start]++
		m[end]--
	}

	si := make([]int, 0, len(m))
	for i := range m {
		si = append(si, i)
	}
	sort.Ints(si)

	counter := 0
	for _, key := range si {
		value := m[key]
		if value < 0 {
			if counter > numberAlive {
				numberAlive = counter
				maxYear = key
			}
		}
		counter += value
	}

	return maxYear, numberAlive
}

func YearList(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("\ndataset: '%s' doesn't exist \n", filepath)
		return nil
	}
	defer file.Close()

	fmt.Printf("\ndataset: '%s'\n", filepath)

	var yearList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		year := scanner.Text()
		yearList = append(yearList, year)
		fmt.Println(year)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return yearList
}

func HighestYearTest(filepath string) {
	if yearList := YearList(filepath); yearList != nil {
		highestYear, count := HighestYear(yearList)
		fmt.Printf("The year with the most number of people alive for dataset %s is %v with %v people alive\n", filepath, highestYear, count)
	}
}

func main() {
	HighestYearTest("/Users/kamau/Documents/GoPath/src/github.com/hopeforsenegal/exampleYears/1.txt")
	HighestYearTest("/Users/kamau/Documents/GoPath/src/github.com/hopeforsenegal/exampleYears/2.txt")
	HighestYearTest("/Users/kamau/Documents/GoPath/src/github.com/hopeforsenegal/exampleYears/3.txt")
	HighestYearTest("/Users/kamau/Documents/GoPath/src/github.com/hopeforsenegal/exampleYears/4.txt")
	HighestYearTest("/Users/kamau/Documents/GoPath/src/github.com/hopeforsenegal/exampleYears/5.txt")
	HighestYearTest("/Users/kamau/Documents/GoPath/src/github.com/hopeforsenegal/exampleYears/6.txt")
}
