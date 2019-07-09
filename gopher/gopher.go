package gopher

import (
	// "fmt"
	"github.com/faiface/pixel"
)

// Gravity constant
const gravity = 700

// Velocity of gopher
var (
	runX  float64 = 230
	jumpY float64 = 350
)

type Gopher struct {
	body pixel.Rect
	vel  pixel.Vec
	jump bool
}

// Initialize Gopher
func New() Gopher {
	// return Gopher{pixel.R(-6, -7, 6, 7), pixel.V(0, 0), false}
	return Gopher{pixel.R(1, 1, 61, 71), pixel.V(0, 0), false}
}

// Gopher Jump
func (g *Gopher) Jump() {
	g.vel.Y += jumpY
	g.body = g.body.Moved(pixel.V(0, g.vel.Y))
	g.jump = true
}

// func (g *Gopher) Update(dir int, dt float64) {
// 	// fmt.Print(g.body)
// 	switch dir {
// 	// Running forward
// 	case 1:
// 		g.vel.X += runX
// 		// Running backward
// 	case -1:
// 		g.vel.X -= runX
// 		// Not moving
// 	case 0:
// 		g.vel.X = 0
// 	}
// 	// Apply gravity
// 	if g.jump {
// 		g.vel.Y -= gravity * dt
// 	}
// 	// Check if hit ground
// 	if g.body.Min.Y <= 1 {
// 		g.vel.Y = 0
// 		g.jump = false
// 	}

// 	g.body = g.body.Moved(g.vel.Scaled(dt))
// }

func (g *Gopher) Update(ctrl pixel.Vec, dt float64) {
	// fmt.Print(g.body)
	switch {
	// Running forward
	case ctrl.X > 0:
		g.vel.X = runX
		// Running backward
	case ctrl.X < 0:
		g.vel.X = -runX
		// Not moving
	case ctrl.X == 0:
		g.vel.X = 0
	}
	// Apply gravity
	if g.jump {
		g.vel.Y -= gravity * dt
	}

	g.body = g.body.Moved(g.vel.Scaled(dt))

	// Check if hit ground
	if g.body.Min.Y <= 1 {
		g.vel.Y = 0
		g.body = g.body.Moved(pixel.V(0, 1-g.body.Min.Y))
		g.jump = false
	}

	if !g.jump && ctrl.Y > 0 {
		g.vel.Y = jumpY
		g.jump = true
	}
}

func (g *Gopher) IsJump() bool {
	return g.jump
}

func (g *Gopher) GetBody() pixel.Rect {
	return g.body
}

func (g *Gopher) GetVel() pixel.Vec {
	return g.vel
}

func (g *Gopher) GetDir() float64 {
	switch {
	case g.vel.X > 0:
		return 1
	case g.vel.X < 0:
		return -1
	case g.vel.X == 0:
		return 1
	}
	return 1
}

func GetJumpY() float64 {
	return jumpY
}
