package main

import (
	"fmt"
	"io"
)

// Color is an ASCII color
type Color int

// Color constants
const (
	Black Color = iota
	Blue
	Brown
	White
	Green
)

// GetCode returns the ANSI code associated with the color
// See https://en.wikipedia.org/wiki/ANSI_escape_code
func (c Color) GetCode() int {
	switch c {
	case Blue:
		return 39
	case Brown:
		return 130
	case White:
		return 15
	case Green:
		return 70
	default:
		return 0
	}
}

// Canvas is an ASCII canvas
type Canvas struct {
	bg     [][]byte
	fg     [][]byte
	chars  [][]byte
	width  int
	height int
	writer io.Writer
}

// NewCanvas creates a new Canvas with the given dimensions
func NewCanvas(width, height int, writer io.Writer) *Canvas {
	res := &Canvas{
		width:  width,
		height: height,
		bg:     make([][]byte, height),
		fg:     make([][]byte, height),
		chars:  make([][]byte, height),
		writer: writer,
	}
	for i := 0; i < height; i++ {
		res.bg[i] = make([]byte, width)
		res.fg[i] = make([]byte, width)
		res.chars[i] = make([]byte, width)
		for j, _ := range res.chars[i] {
			res.chars[i][j] = ' '
		}
	}

	return res
}

// Clear reset the canvas to its initial blank state
func (c *Canvas) Clear() {
	for i := 0; i < c.height; i++ {
		c.bg[i] = make([]byte, c.width)
		c.fg[i] = make([]byte, c.width)
		c.chars[i] = make([]byte, c.width)
		for j, _ := range c.chars[i] {
			c.chars[i][j] = ' '
		}
	}
}

// DisplayCanvas writes the canvas to its writer
func (c *Canvas) DisplayCanvas() {
	for y, line := range c.chars {
		lastBg := -1
		lastFg := -1
		for x, char := range line {
			bg := Color(int(c.bg[y][x]) - '0').GetCode()
			if bg != lastBg {
				fmt.Fprintf(c.writer, "\033[48;5;%dm", bg)
				lastBg = bg
			}
			fg := Color(int(c.fg[y][x]) - '0').GetCode()
			if fg != lastFg {
				fmt.Fprintf(c.writer, "\033[38;5;%dm", fg)
				lastFg = fg
			}
			fmt.Fprintf(c.writer, "%c", char)
		}
		fmt.Fprintf(c.writer, "\033[0m\n")
	}
}

// DrawSprite draws a sprite into the canvas
func (c *Canvas) DrawSprite(s *Sprite) {
	for y, line := range s.chars {
		posY := y + s.posY
		if posY < 0 || posY >= c.height {
			continue
		}
		for x, char := range line {
			posX := x + s.posX
			if posX < 0 || posX >= c.width {
				continue
			}
			if s.bgColors[y][x] == ' ' &&
				s.fgColors[y][x] == ' ' &&
				char == ' ' {
				continue
			}
			c.bg[y+s.posY][x+s.posX] = s.bgColors[y][x]
			c.fg[y+s.posY][x+s.posX] = s.fgColors[y][x]
			c.chars[y+s.posY][x+s.posX] = byte(char)
		}
	}
}
