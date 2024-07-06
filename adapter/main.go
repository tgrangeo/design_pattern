package main

import (
	"fmt"
	"math"
)

// This class expects a round peg
type RoundHole struct {
	radius int
}

func (r *RoundHole) fits(peg RoundPeg) bool {
	return peg.radius <= r.radius
}

// Different classes
type RoundPeg struct {
	radius int
}

type SquarePeg struct {
	width int
}

// Adapter from square to round peg
type SquarePegAdapter struct {
	squarePeg SquarePeg
}

func (s *SquarePegAdapter) convertToRound() RoundPeg {
	return RoundPeg{
		radius: int(float64(s.squarePeg.width) * math.Sqrt(2) / 2),
	}
}

func main() {
	hole := RoundHole{radius: 12}
	littlePeg := RoundPeg{radius: 10}
	bigPeg := RoundPeg{radius: 20}

	fmt.Println("little peg fits:", hole.fits(littlePeg))
	fmt.Println("big peg fits:", hole.fits(bigPeg))

	littleSquare := SquarePeg{width: 2}
	bigSquare := SquarePeg{width: 20}

	littleAdapter := SquarePegAdapter{squarePeg: littleSquare}
	bigAdapter := SquarePegAdapter{squarePeg: bigSquare}

	littleSquareConverted := littleAdapter.convertToRound()
	bigSquareConverted := bigAdapter.convertToRound()

	fmt.Println("little square peg fits:", hole.fits(littleSquareConverted))
	fmt.Println("big square peg fits:", hole.fits(bigSquareConverted))
}
