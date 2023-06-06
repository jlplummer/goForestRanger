package main

import (
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	Caption      = "goForestRanger"
	ScreenWidth  = 600
	ScreenHeight = 300

	PlayerHeight = 24
	PlayerWidth  = 21

	TileHeight = 4
	TileWidth  = 8

	ItemHeight = 8
	ItemWidth  = 8

	SpriteFactor = 3.0

	FootmanHeight = 8
	FootmanWidth  = 8

	KnightHeight = 16
	KnightWidth  = 19

	EnemyCooldown     = 200
	CollisionDistance = 20
	ArrowCooldown     = 10
	DaggerCooldown    = 10

	PlayerMove = 8 * SpriteFactor

	ButtonCooldown = 500 // milliseconds
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
	ebiten.SetWindowTitle(Caption)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
