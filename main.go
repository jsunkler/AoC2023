package main

import (
	"flag"
	"fmt"

	"github.com/jsunkler/AoC2023/solvers"
)

func main() {
	var dayFlag = flag.Int("day", 1, "Enter the day of december to process.")
	flag.Parse()

	fmt.Printf("Solution for day %d is %s.\n", *dayFlag, solvers.Solve(*dayFlag))
}
