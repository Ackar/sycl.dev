package main

import (
	"fmt"
	"io"
)

type color int

const (
	black color = iota
	blue
	brown
	white
	green
)

func (c color) GetCode() int {
	switch c {
	case blue:
		return 39
	case brown:
		return 130
	case white:
		return 15
	case green:
		return 70
	default:
		return 0
	}
}

type Canvas struct {
	bg     [][]byte
	fg     [][]byte
	chars  [][]byte
	width  int
	height int
	writer io.Writer
}

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

func (c *Canvas) DisplayCanvas() {
	for y, line := range c.chars {
		lastBg := -1
		lastFg := -1
		for x, char := range line {
			bg := color(int(c.bg[y][x]) - '0').GetCode()
			if bg != lastBg {
				fmt.Fprintf(c.writer, "\033[48;5;%dm", bg)
				lastBg = bg
			}
			fg := color(int(c.fg[y][x]) - '0').GetCode()
			if fg != lastFg {
				fmt.Fprintf(c.writer, "\033[38;5;%dm", fg)
				lastFg = fg
			}
			fmt.Fprintf(c.writer, "%c", char)
		}
		fmt.Fprintf(c.writer, "\n")
	}
}

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
