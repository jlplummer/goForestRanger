package main

import (
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	ScreenWidth  = 600
	ScreenHeight = 300
)

var (
	RangerImage *ebiten.Image
)

func init() {
	var err error
	RangerImage, _, err = ebitenutil.NewImageFromFile("media/ranger.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(5, 5)
	screen.DrawImage(RangerImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("goForestRanger")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
