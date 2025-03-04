
package main

import (
	"fmt"
	"bufio"
	"github.com/raflyhangga/go-car-benchmark/car"
	"os"
)

var uniqueList []string
var indexList []int
var uniqueMap = make(map[string]int)

func loadValue(key interface{}, slotNo int) interface{} {
	return uniqueList[key.(int)]
}

func main() {
	file, err := os.Open("cloudPhysicsIO.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if idx, exists := uniqueMap[line]; exists {
			indexList = append(indexList, idx)
		} else {
			uniqueMap[line] = len(uniqueList)
			uniqueList = append(uniqueList, line)
			indexList = append(indexList, uniqueMap[line])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var numSlots [9]int;
	numSlots[0] = 1000
	numSlots[1] = 5000
	numSlots[2] = 8000
	numSlots[3] = 10000
	numSlots[4] = 20000
	numSlots[5] = 30000
	numSlots[6] = 40000
	numSlots[7] = 50000
	numSlots[8] = 900


	for _, value := range numSlots {
		car := car.NewCAR(value)
		car.SetLoadValue(loadValue)
		
		for _, value := range indexList {
			car.Load(value)
		}
	
		reqCount, missCount := car.GetCount()
		missRate := float64(missCount) / float64(reqCount)
		fmt.Printf("[i] Cache Size: %d Miss Rate: %.4f\n", value, missRate)
	}
}
