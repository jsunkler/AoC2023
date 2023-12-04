package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jsunkler/AoC2023/solvers"
)

func main() {
	var dayFlag = flag.Int("day", 1, "Enter the day of december to process.")
	flag.Parse()

	p1, p2, err := solvers.Solve(*dayFlag)
	if err != nil {
		fmt.Printf("Execution failed. Error: %w\n", err)
		os.Exit(1)
	}
	fmt.Printf("Solved!\nPart 1\n%s\n\nPart 2\n%s\n\n", p1, p2)
}
