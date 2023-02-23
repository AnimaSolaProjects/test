package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type point struct {
	x int
	y int
}

func checkIfIn(tempPoint point, arr []point) bool {
	isIn := false
	for _, point := range arr {
		if point.x == tempPoint.x && point.y == tempPoint.y {
			isIn = true
		}
	}
	return isIn
}

func filling() []point {
	var length int
	fmt.Fscan(os.Stdin, &length)
	rand.Seed(time.Now().Unix())
	arr := make([]point, 0, length)

	for {
		tempPoint := point{rand.Intn(10), rand.Intn(10)}

		if checkIfIn(tempPoint, arr) {
			fmt.Println("Point is already present and will be ignored")
			continue
		}
		arr = append(arr, tempPoint)

		if !(len(arr) < length) {
			break
		}
	}

	return arr
}

func main() {
	fmt.Print("Enter the number of elements: ")
	arr := filling()
	counter := 0

	for a := 0; a < len(arr); a++ {
		for b := a + 1; b < len(arr); b++ {
			if arr[a].x == arr[b].x {
				for a1 := a + 1; a1 < len(arr); a1++ {
					for b1 := b + 1; b1 < len(arr); b1++ {
						if arr[a1].x == arr[b1].x {
							if (arr[a].x == arr[a1].x && arr[a].x == arr[a1].x) || (arr[b].x == arr[b1].x && arr[b].x == arr[b1].x) {
								if (arr[a].y-arr[a1].y == arr[b].y-arr[b1].y) || (arr[a].y-arr[b1].y == arr[b].y-arr[a1].y) {
									counter++
								}
							}
						}
					}
				}
				arr = append(arr[:b], arr[b+1:]...)
				arr = append(arr[:a], arr[a+1:]...)
			}
		}
	}
	fmt.Println("Counter:", counter)
}
