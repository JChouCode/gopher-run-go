package gopher_anim

import (
	"encoding/csv"
	// "fmt"
	"github.com/JChouCode/gopher-run-go/gopher"
	"github.com/faiface/pixel"
	"image"
	_ "image/png"
	"io"
	"math"
	"os"
	"strconv"
)

func loadAnimationSheet(imgPath, csvPath string, fWidth float64) (sheet pixel.Picture, anims map[string][]pixel.Rect, err error) {
	// open and load the spritesheet
	sheetFile, err := os.Open(imgPath)
	if err != nil {
		return nil, nil, err
	}
	defer sheetFile.Close()
	sheetImg, _, err := image.Decode(sheetFile)
	if err != nil {
		return nil, nil, err
	}
	sheet = pixel.PictureDataFromImage(sheetImg)

	// create a slice of frames inside the spritesheet
	var frames []pixel.Rect
	for x := 0.0; x+fWidth <= sheet.Bounds().Max.X; x += fWidth {
		frames = append(frames, pixel.R(
			x,
			0,
			x+fWidth,
			sheet.Bounds().H(),
		))
	}

	descFile, err := os.Open(csvPath)
	if err != nil {
		return nil, nil, err
	}
	defer descFile.Close()

	anims = make(map[string][]pixel.Rect)

	// load the animation information, name and interval inside the spritesheet
	desc := csv.NewReader(descFile)
	for {
		anim, err := desc.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, nil, err
		}
		name := anim[0]
		start, _ := strconv.Atoi(anim[1])
		end, _ := strconv.Atoi(anim[2])
		anims[name] = frames[start : end+1]
	}
	return sheet, anims, nil
}

type move_state int

const (
	idle move_state = iota
	run
	jump
)

var rate = 1.0 / 10
var eyeRate = 3.0
var blink = false
var counter = 0.0

type GopherAnim struct {
	sprite *pixel.Sprite
	frame  pixel.Rect
	state  move_state

	sheet pixel.Picture
	anims map[string][]pixel.Rect

	dir float64
}

// Initialize GopherAnim
func New(imgPath string, csvPath string) GopherAnim {
	sheet, anims, err := loadAnimationSheet(imgPath, csvPath, 12)
	if err != nil {
		panic(err)
	}
	return GopherAnim{pixel.NewSprite(nil, pixel.Rect{}), pixel.Rect{}, idle, sheet, anims, 0}
}

func (ga *GopherAnim) Update(g gopher.Gopher, dt float64) {
	counter += dt

	var tempState move_state
	//Update state
	switch {
	case g.IsJump():
		tempState = jump
	case g.GetVel().Len() > 0:
		tempState = run
	case g.GetVel().Len() == 0:
		tempState = idle
	}

	if tempState != ga.state {
		ga.state = tempState
		counter = 0
	}

	switch ga.state {
	case idle:
		// if blink = true {
		// 	ga.frame = ga.animas["Front"][1]
		// }
		if int(math.Floor(counter))%4 == 0 {
			ga.frame = ga.anims["Front"][1]
			counter += 0.3
		} else {
			ga.frame = ga.anims["Front"][0]
		}
		// ga.frame = ga.anims["Front"][0]
	case run:
		// fmt.Print("run")
		i := int(math.Floor(counter / rate))
		ga.frame = ga.anims["Run"][i%len(ga.anims["Run"])]
	case jump:
		// fmt.Print("jump")
		i := 0
		switch {
		case g.GetVel().Y < gopher.GetJumpY()*1/3:
			i += 2
		case g.GetVel().Y < gopher.GetJumpY()*2/3:
			i++
		}
		ga.frame = ga.anims["Jump"][i]
	}

	ga.dir = g.GetDir()
}

func (ga *GopherAnim) Draw(t pixel.Target, g gopher.Gopher) {
	// fmt.Print(ga.frame)
	ga.sprite.Set(ga.sheet, ga.frame)
	ga.sprite.Draw(t, pixel.IM.
		ScaledXY(pixel.ZV, pixel.V(
			g.GetBody().W()/ga.sprite.Frame().W(),
			g.GetBody().H()/ga.sprite.Frame().H(),
		)).
		ScaledXY(pixel.ZV, pixel.V(-ga.dir, 1)).
		Moved(g.GetBody().Center()),
	)
}

func (ga *GopherAnim) GetSheet() pixel.Picture {
	return ga.sheet
}
