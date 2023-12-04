package main

import (
	"flag"
	"log/slog"

	"github.com/jsunkler/AoC2023/solvers"
)

func main() {
	var dayFlag = flag.Int("day", 1, "Enter the day of december to process.")
	flag.Parse()

	slog.Info("Solving...", "day", *dayFlag)
	p1, p2 := solvers.Solve(*dayFlag)
	slog.Info("Solved.", "Part 1", p1, "Part 2", p2)
}
