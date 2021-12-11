package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func solvePart1(octopuses [][]int) {
	steps := 100
	flashes := 0

	for i := 0; i < steps; i++ {
		for y := 0; y < len(octopuses); y++ {
			for x := 0; x < len(octopuses[y]); x++ {
				increaseEnergy(octopuses, x, y)
			}
		}

		zeroriseLitOctopus(octopuses)
		flashes += countFlashes(octopuses)
	}

	// printOctopuses(octopuses)
	fmt.Println("PART 1 ANSWER")
	fmt.Println(flashes)
}

func solvePart2(octopuses [][]int) {
	for i := 0; ; i++ {
		for y := 0; y < len(octopuses); y++ {
			for x := 0; x < len(octopuses[y]); x++ {
				increaseEnergy(octopuses, x, y)
			}
		}

		zeroriseLitOctopus(octopuses)

		if isAllLit(octopuses) {
			fmt.Println("PART 2 ANSWER")
			fmt.Println(i + 1)

			return
		}
	}
}

func increaseEnergy(octopuses [][]int, x int, y int) {
	octopuses[y][x]++

	if octopuses[y][x] == 10 {
		// Check top left
		if (x-1 >= 0) && (y-1 >= 0) && octopuses[y-1][x-1] <= 9 {
			increaseEnergy(octopuses, x-1, y-1)
		}

		// Check top
		if y-1 >= 0 && octopuses[y-1][x] <= 9 {
			increaseEnergy(octopuses, x, y-1)
		}

		// Check top right
		if y-1 >= 0 && x+1 < len(octopuses[y-1]) && octopuses[y-1][x+1] <= 9 {
			increaseEnergy(octopuses, x+1, y-1)
		}

		// Check right
		if x+1 < len(octopuses[y]) && octopuses[y][x+1] <= 9 {
			increaseEnergy(octopuses, x+1, y)
		}

		// Check bottom right
		if y+1 < len(octopuses) && x+1 < len(octopuses[y+1]) && octopuses[y+1][x+1] <= 9 {
			increaseEnergy(octopuses, x+1, y+1)
		}

		// Check bottom
		if y+1 < len(octopuses) && octopuses[y+1][x] <= 9 {
			increaseEnergy(octopuses, x, y+1)
		}

		// Check bottom left
		if y+1 < len(octopuses) && x-1 >= 0 && octopuses[y+1][x-1] <= 9 {
			increaseEnergy(octopuses, x-1, y+1)
		}

		// Check left
		if x-1 >= 0 && octopuses[y][x-1] <= 9 {
			increaseEnergy(octopuses, x-1, y)
		}
	}
}

func countFlashes(octopuses [][]int) int {
	count := 0

	for i := 0; i < len(octopuses); i++ {
		for j := 0; j < len(octopuses[i]); j++ {
			if octopuses[i][j] == 0 {
				count++
			}
		}
	}

	return count
}

func zeroriseLitOctopus(octopuses [][]int) {
	for y := 0; y < len(octopuses); y++ {
		for x := 0; x < len(octopuses[y]); x++ {
			if octopuses[y][x] > 9 {
				octopuses[y][x] = 0
			}
		}
	}
}

func isAllLit(octopuses [][]int) bool {
	for y := 0; y < len(octopuses); y++ {
		for x := 0; x < len(octopuses[y]); x++ {
			if octopuses[y][x] != 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}
	defer f.Close()

	octopuses1 := make([][]int, 0)
	octopuses2 := make([][]int, 0)

	r := bufio.NewReader(f)
	for {
		l, err := r.ReadString('\n')
		if err != nil && len(l) == 0 {
			if err == io.EOF {
				break
			}

			log.Fatalf("failed to read input file: %s", err)
		}

		l = strings.TrimSpace(l)
		energies1 := make([]int, 0, len(l))
		energies2 := make([]int, 0, len(l))
		for _, r := range l {
			energies1 = append(energies1, int(r-'0'))
			energies2 = append(energies2, int(r-'0'))
		}

		octopuses1 = append(octopuses1, energies1)
		octopuses2 = append(octopuses2, energies2)
	}

	solvePart1(octopuses1)
	solvePart2(octopuses2)
}
