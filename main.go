package main

import (
	"fmt"
	"goForestRanger/pkg/entities"
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
	ranger      entities.Ranger
	initCalled  int
	last_move_y int
)

func init() {
	ranger.Init(0, 0)
	initCalled++
}

type Game struct{}

func (g *Game) Update() error {
	var move_y int

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		move_y -= ranger.Height()
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		move_y += ranger.Height()
	}

	ranger.Update(move_y)

	last_move_y = move_y
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ranger.Draw(screen)

	msg := fmt.Sprintf("ranger[x:%d/y:%d]", ranger.X, ranger.Y)
	ebitenutil.DebugPrintAt(screen, msg, 500, 0)

	msg = fmt.Sprintf("initCalled:%d", initCalled)
	ebitenutil.DebugPrintAt(screen, msg, 500, 25)

	msg = fmt.Sprintf("move_y:%d", last_move_y)
	ebitenutil.DebugPrintAt(screen, msg, 500, 50)

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
