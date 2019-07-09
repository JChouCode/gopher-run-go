package main

import (
	"github.com/JChouCode/gopher-run-go/gopher"
	"github.com/JChouCode/gopher-run-go/gopher_anim"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	// "image"
	_ "image/png"
	// "math/rand"
	// "os"
	"time"
)

//Initialize window
func initWindow(t string, w float64, h float64) *pixelgl.Window {
	cfg := pixelgl.WindowConfig{
		Title:  t,
		Bounds: pixel.R(0, 0, w, h),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	// win.SetSmooth(true)
	return win
}

func run() {

	win := initWindow("Gopher-Run-GO", 1024, 768)

	gopher := gopher.New()
	anim := gopher_anim.New("sheet.png", "sheet.csv")

	imd := imdraw.New(anim.GetSheet())

	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(colornames.Blanchedalmond)

		// dir := 0

		// if win.Pressed(pixelgl.KeyLeft) {
		// 	dir--
		// }
		// if win.Pressed(pixelgl.KeyRight) {
		// 	dir++
		// }
		// if win.JustPressed(pixelgl.KeyUp) {
		// 	if !gopher.IsJump() {
		// 		gopher.Jump()
		// 	}
		// }

		ctrl := pixel.ZV
		if win.Pressed(pixelgl.KeyLeft) {
			ctrl.X--
		}
		if win.Pressed(pixelgl.KeyRight) {
			ctrl.X++
		}
		if win.JustPressed(pixelgl.KeyUp) {
			ctrl.Y++
		}

		// gopher.Update(dir, dt)
		gopher.Update(ctrl, dt)
		anim.Update(gopher, dt)

		imd.Clear()

		anim.Draw(imd, gopher)

		imd.Draw(win)

		// imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
