package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var amountRepit int
	errFor4 := 0
	errFor5 := 0
	goodFourNumberCar := []string{}
	goodFiveNumberCar := []string{}
	fmt.Scanln(&amountRepit)
	scaner := bufio.NewScanner(os.Stdin)
	for i := 0; i < amountRepit; i++ {
		scaner.Scan()
		inputString := scaner.Text()
		goodFiveNumberCar, errFor5 = ChekingLineFiveSimvol(inputString)
		goodFourNumberCar, errFor4 = ChekingLineFourSimvol(inputString)

		if errFor4 == 0 {
			for _, v := range goodFourNumberCar {
				fmt.Print(v + " ")
			}
		} else if errFor5 == 0 {
			for _, v := range goodFiveNumberCar {
				fmt.Print(v + " ")
			}
		} else {
			fmt.Println("-")
		}
	}
}

func ChekingLineFourSimvol(inputString string) ([]string, int) {
	var (
		line            string
		lengInputString int
		detail          int
	)
	goodCarNumbers := []string{}
	err := 0
	lengInputString = len(inputString)
	if lengInputString%4 == 0 {
		for i := 0; i <= lengInputString; i = i + 3 {
			line = inputString[:4]
			inputString = inputString[4:]
			detail = DetailFourLine(line)
			if detail == 1 {
				goodCarNumbers = append(goodCarNumbers, line)
			} else {
				err = 1
			}
		}
	} else {
		err = 1
	}

	return goodCarNumbers, err
}

func ChekingLineFiveSimvol(inputString string) ([]string, int) {
	var (
		line            string
		lengInputString int
		detail          int
	)
	goodCarNumbers := []string{}
	err := 0
	lengInputString = len(inputString)
	if lengInputString%5 == 0 {
		for i := 0; i <= lengInputString; i = i + 4 {
			line = inputString[:5]
			inputString = inputString[5:]
			detail = DetailFiveLine(line)
			if detail == 1 {
				goodCarNumbers = append(goodCarNumbers, line)
			} else {
				err = 1
			}
		}
	} else {
		err = 1
	}

	return goodCarNumbers, err
}

func DetailFourLine(line string) int {
	var result int
	firstIndexNumber := strings.IndexAny(line, "0123456789")
	lastIndexNumber := strings.LastIndexAny(line, "0123456789")
	if firstIndexNumber == lastIndexNumber && firstIndexNumber == 1 {
		result = 1
	} else {
		result = 0
	}

	return result
}

func DetailFiveLine(line string) int {
	var result int
	firstIndexNumber := strings.IndexAny(line, "0123456789")
	lastIndexNumber := strings.LastIndexAny(line, "0123456789")
	if firstIndexNumber == 1 && lastIndexNumber == 2 {
		result = 1
	} else {
		result = 0
	}

	return result
}
