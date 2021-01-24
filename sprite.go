package main

var treeChars = []string{
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
}
var treeFg = []string{
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
	"               ",
}
var treeBg = []string{
	"    44444444   ",
	"   4444444444  ",
	"  444444444444 ",
	"  44444444444  ",
	"    4444444    ",
	"      222      ",
	"      222      ",
	"      222      ",
	"      222      ",
}

var cloudChars = []string{
	"               ",
	"               ",
	"               ",
	"               ",
}
var cloudFg = []string{
	"               ",
	"               ",
	"               ",
	"               ",
}
var cloudBg = []string{
	"    3333333    ",
	"  333333333333 ",
	"333333333333   ",
	"   33333333    ",
}

var gopherChars = []string{
	"          ",
	"          ",
	"   .  .   ",
	"    --    ",
	"    ll    ",
	"          ",
	"          ",
}

var gopherBg = []string{
	"  1    1  ",
	" 11111111 ",
	" 11311311 ",
	"2111111112",
	" 11111111 ",
	" 11111111 ",
	"  2    2  ",
}
var gopherFg = []string{
	"  4    4  ",
	" 44444444 ",
	" 44444444 ",
	"4444224444",
	" 44433444 ",
	" 44444444 ",
	"  2    2  ",
}

var gopherChars2 = []string{
	"          ",
	"          ",
	"   .  .   ",
	"    --    ",
	"    ll    ",
	"          ",
	"  /       ",
}

var gopherBg2 = []string{
	"  1    1  ",
	" 11111111 ",
	" 11311311 ",
	"2111111112",
	" 11111111 ",
	" 11111111 ",
	"       2  ",
}
var gopherFg2 = []string{
	"  4    4  ",
	" 44444444 ",
	" 44444444 ",
	"4444224444",
	" 44433444 ",
	" 44444444 ",
	"  2    2  ",
}

var gopherChars3 = []string{
	"          ",
	"          ",
	"   .  .   ",
	"    --    ",
	"    ll    ",
	"          ",
	"       /  ",
}

var gopherBg3 = []string{
	"  1    1  ",
	" 11111111 ",
	" 11311311 ",
	"2111111112",
	" 11111111 ",
	" 11111111 ",
	"  2       ",
}
var gopherFg3 = []string{
	"  4    4  ",
	" 44444444 ",
	" 44444444 ",
	"4444224444",
	" 44433444 ",
	" 44444444 ",
	"  2    2  ",
}

var kiwiChars = []string{
	"   .  ",
	"     \\",
	"| |   ",
}

var kiwiFg = []string{
	"      ",
	"     4",
	"2 2   ",
}

var kiwiBg = []string{
	"2222  ",
	"222   ",
	"      ",
}

type Sprite struct {
	fgColors []string
	bgColors []string
	chars    []string
	posX     int
	posY     int
}
