package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/ipinfil/aoc2022/loading"
)

type coordinate struct {
	x int
	y int
}

type hashset map[coordinate]struct{}

var visitedCoors hashset = make(hashset)

func (c *coordinate) MoveUp() {
	c.y++
}

func (c *coordinate) MoveDown() {
	c.y--
}

func (c *coordinate) MoveLeft() {
	c.x--
}

func (c *coordinate) MoveRight() {
	c.x++
}

func (c *coordinate) MoveRightUp() {
	c.x++
	c.y++
}

func (c *coordinate) MoveRightDown() {
	c.x++
	c.y--
}

func (c *coordinate) MoveLeftUp() {
	c.x--
	c.y++
}

func (c *coordinate) MoveLeftDown() {
	c.x--
	c.y--
}

func Distance(c1 *coordinate, c2 *coordinate) int {
	return int(math.Max(math.Abs(float64(c1.x - c2.x)), math.Abs(float64(c1.y - c2.y))))
}

func (c *coordinate) MoveAccordingly(h *coordinate, isTail bool) {
	if isTail {
		visitedCoors[coordinate{c.x, c.y}] = struct{}{}
	}

	if Distance(c, h) < 2 {
		return
	}

	if c.y == h.y {
		if h.x > c.x {
			c.MoveRight()
		} else {
			c.MoveLeft()
		}
	} else if c.x == h.x {
		if h.y > c.y {
			c.MoveUp()
		} else {
			c.MoveDown()
		}
	} else {
		if h.y > c.y && h.x > c.x {
			c.MoveRightUp()
		} else if h.y > c.y && h.x < c.x {
			c.MoveLeftUp()
		} else if h.y < c.y && h.x > c.x {
			c.MoveRightDown()
		} else {
			c.MoveLeftDown()
		}
	}

	if isTail {
		visitedCoors[coordinate{c.x, c.y}] = struct{}{}
	}
}

func DoOperation(h *coordinate, knots []*coordinate, direction string) {
	if direction == "R" {
		h.MoveRight()
	} else if direction == "L" {
		h.MoveLeft()
	} else if direction == "D" {
		h.MoveDown()
	} else {
		h.MoveUp()
	}

	tmpHead := h
	for i, knot := range knots {
		knot.MoveAccordingly(tmpHead, i == 8)
		tmpHead = knot
	}
}

func main() {
	file := loading.LoadFile("9-1/input.txt")
	scanner := loading.NewLineScanner(file)
	defer file.Close()

	head := &coordinate{}
	knots := []*coordinate{
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
		{},
	}

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")

		number, err := strconv.Atoi(text[1])
		if err != nil {
			panic(err)
		}

		for i := 0; i < number; i++ {
			DoOperation(head, knots, text[0])
		}
	}
	fmt.Println(len(visitedCoors))
}