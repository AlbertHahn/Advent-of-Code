package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	//"math"
)

func catch(e error){
	if e != nil {
		panic(e)
	}
}

func main() {

	readFile, err := os.Open("input.txt")
	catch(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var ListOne []int
	var ListTwo []int

    for fileScanner.Scan() {
		lineLeft, err := strconv.Atoi(fileScanner.Text()[0:5])
		lineRight, err := strconv.Atoi(fileScanner.Text()[8:13])

		if err != nil {
			panic(err)
		}

		ListOne = append(ListOne, lineLeft)
		ListTwo = append(ListTwo, lineRight)
    }

	fmt.Println("Column Left: ", ListOne)
	fmt.Println("Column Right: ", ListTwo)

	readFile.Close()

	var totalDistance = getTotalDistance(ListOne, ListTwo)
	fmt.Println(totalDistance)

}


func getTotalDistance(ListOne []int, ListTwo []int) int {

	var totalDistance int = 0

	for len(ListOne) > 0 && len(ListTwo) > 0 {

		minIndexOne := 0
		minIndexTwo := 0
	
		for i, v := range ListOne {
			if (v < ListOne[minIndexOne]) {
				minIndexOne  = i
			}
		}
	
		for i, v := range ListTwo {
			if (v < ListTwo[minIndexTwo]) {
				minIndexTwo  = i
			}
		}
	
		distance := getDistance(ListOne[minIndexOne], ListTwo[minIndexTwo])
	
		fmt.Println("The lowest number is:", ListOne[minIndexOne])
		fmt.Println("The lowest number is:", ListTwo[minIndexTwo])
		fmt.Println("Distance between the lowest", distance)
	
		ListOne = append(ListOne[:minIndexOne], ListOne[minIndexOne+1:]...)
		ListTwo = append(ListTwo[:minIndexTwo], ListTwo[minIndexTwo+1:]...)

		totalDistance += distance

	}

	return totalDistance
}

func getDistance(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}