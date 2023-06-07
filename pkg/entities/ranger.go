package entities

import (
	"goForestRanger/pkg/game"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Ranger struct {
	X, Y, Health int
	Image        *ebiten.Image
}

func (r *Ranger) Init(x, y int) {
	var err error

	r.Health = 5
	r.X = x
	r.Y = y

	r.Image, _, err = ebitenutil.NewImageFromFile("media/ranger_8x8.png")
	if err != nil {
		log.Fatal("Cannot open ranger_8x8.png")
	}
}

func (r *Ranger) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Translate(float64(r.X), float64(r.Y))
	op.GeoM.Scale(game.SpriteFactor, game.SpriteFactor)
	screen.DrawImage(r.Image, op)
}

func (r *Ranger) Update(y int) {
	// don't let sprite move off screen
	if r.Y+y < 0 {
		y = 0
	}

	// dividing screen size by sprite factor and subtracting
	// sprite height seems to work?
	if r.Y+y > (game.ScreenHeight/game.SpriteFactor)-game.PlayerMove {
		y = 0
	}
	r.Y += y
}

func (r *Ranger) Height() int {
	return r.Image.Bounds().Dy()
}

func (r *Ranger) Width() int {
	return r.Image.Bounds().Dx()
}
