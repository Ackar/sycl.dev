package main

import (
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type terminalDisplayer struct {
	w      io.Writer
	canvas *Canvas
}

func (d *terminalDisplayer) clear() {
	d.writeAndFlush("\033c")
}

func (d *terminalDisplayer) writeAndFlush(s string) {
	d.w.Write([]byte(s))
	d.flush()
}

func (d *terminalDisplayer) write(s string) {
	d.w.Write([]byte(s))
}

func (d *terminalDisplayer) flush() {
	if f, ok := d.w.(http.Flusher); ok {
		f.Flush()
	}
}

func (d *terminalDisplayer) writeHeader() {
	d.write("                              \033[32mSylvain Cleymans\033[0m\n")
	d.write("                              ----------------\n")
	d.write("                             Software Engineer\n")
	d.write("                              contact@sycl.dev\n")
	d.flush()
}

func (d *terminalDisplayer) render() {
	gopherSprites := []*Sprite{
		{
			bgColors: gopherBg,
			fgColors: gopherFg,
			chars:    gopherChars,
			posX:     30,
			posY:     18,
		},
		{
			bgColors: gopherBg2,
			fgColors: gopherFg2,
			chars:    gopherChars2,
			posX:     30,
			posY:     18,
		},
		{
			bgColors: gopherBg3,
			fgColors: gopherFg3,
			chars:    gopherChars3,
			posX:     30,
			posY:     18,
		},
	}

	treeSprite := &Sprite{
		bgColors: treeBg,
		fgColors: treeFg,
		chars:    treeChars,
		posX:     75,
		posY:     15,
	}
	cloudSprites := []*Sprite{
		{
			bgColors: cloudBg,
			fgColors: cloudFg,
			chars:    cloudChars,
			posX:     75,
			posY:     0,
		},
		{
			bgColors: cloudBg,
			fgColors: cloudFg,
			chars:    cloudChars,
			posX:     55,
			posY:     3,
		},
	}
	kiwiSprite := &Sprite{
		bgColors: kiwiBg,
		fgColors: kiwiFg,
		chars:    kiwiChars,
		posX:     100,
		posY:     22,
	}
	step := 0
	gopher := 0
	begin := time.Now()
	for {
		d.clear()
		d.writeHeader()
		d.canvas.Clear()
		d.canvas.DrawSprite(kiwiSprite)
		if step%3 == 0 {
			gopher = (gopher + 1) % len(gopherSprites)
		}
		for _, cloud := range cloudSprites {
			d.canvas.DrawSprite(cloud)
		}
		d.canvas.DrawSprite(treeSprite)
		d.canvas.DrawSprite(gopherSprites[gopher])
		d.canvas.DisplayCanvas()
		if step%2 == 0 {
			treeSprite.posX--
			kiwiSprite.posX--
		}

		if treeSprite.posX < -15 {
			treeSprite.posX = rand.Intn(80) + 80
		}

		if kiwiSprite.posX < -15 {
			kiwiSprite.posX = rand.Intn(100) + 80
		}

		for _, cloud := range cloudSprites {
			cloud.posX--
			if cloud.posX < -20 {
				cloud.posX = rand.Intn(80) + 80
				cloud.posY = rand.Intn(10)
			}
		}
		step++
		d.flush()
		time.Sleep(50 * time.Millisecond)
		if time.Now().Sub(begin).Seconds() > 30 {
			break
		}
	}
	d.writeAndFlush("\033[0mSee you later!\n")
}

func consoleHandler(w http.ResponseWriter, r *http.Request) {
	d := terminalDisplayer{w, NewCanvas(80, 25, w)}
	d.clear()
	d.writeHeader()
	d.render()
}

type polyHandler struct {
	console http.HandlerFunc
	web     http.HandlerFunc
}

func (h *polyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	userAgent := r.Header.Get("user-agent")
	log.WithField("user-agent", userAgent).Info("Received request")
	if !strings.Contains(userAgent, "curl") {
		h.web(w, r)
	} else {
		h.console(w, r)
	}
}

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Starting...")
	fs := http.FileServer(http.Dir("static"))
	multi := &polyHandler{consoleHandler, fs.ServeHTTP}
	if err := http.ListenAndServe(":8080", multi); err != nil {
		panic(err)
	}
}
