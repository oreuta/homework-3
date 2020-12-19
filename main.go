package main

import (
	"errors"
	"fmt"
)

var (
	dirX           = [4]int{1, 0, -1, 0}
	dirY           = [4]int{0, 1, 0, -1}
	errZeroLength  = errors.New("zero length")
	errInvalidSize = errors.New("not a n*n matrix")
)

func transform(m [][]int) ([]int, error) {
	if len(m) == 0 {
		return nil, errZeroLength
	}

	xLen, yLen := len(m[0]), len(m)
	if xLen != yLen {
		return nil, errInvalidSize
	}

	for _, row := range m {
		if len(row) != xLen {
			return nil, errInvalidSize
		}
	}

	N := xLen
	seen := make([][]bool, N)
	for i := 0; i < N; i++ {
		seen[i] = make([]bool, N)
	}

	iterations, dirIdx, x, y := N*N, 0, 0, 0
	output := []int{}
	for iterations > 0 {
		output = append(output, m[y][x])
		seen[y][x] = true

		newX, newY := x+dirX[dirIdx], y+dirY[dirIdx]

		if newX >= N || newX < 0 || newY >= N || newY < 0 || seen[newY][newX] {
			dirIdx = (dirIdx + 1) % 4
			x, y = x+dirX[dirIdx], y+dirY[dirIdx]
		} else {
			x, y = newX, newY
		}
		iterations--
	}

	return output, nil
}

func main() {
	fmt.Print("Matrix size: ")
	var N int
	fmt.Scan(&N)

	m := make([][]int, N)
	for i := range m {
		fmt.Printf("Row %d\n", i+1)
		m[i] = make([]int, N)

		for j := range m[i] {
			_, err := fmt.Scan(&m[i][j])
			if err != nil {
				fmt.Println("Invalid input")
				return
			}
		}
	}

	transformed, err := transform(m)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Transformed: %v\n", transformed)
	}
}
