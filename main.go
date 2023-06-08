package main

import (
	"fmt"
	"goForestRanger/pkg/entities"
	"goForestRanger/pkg/game"
	"log"
	"os"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	ranger      entities.Ranger
	background  entities.Background
	initCalled  int
	last_move_y int
)

func init() {
	ranger.Init(0, 0)
	background.Init()
	initCalled++
}

type Game struct{}

func (g *Game) Update() error {
	var move_y int

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(0)
	}

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
	background.Draw(screen)
	ranger.Draw(screen)

	msg := fmt.Sprintf("ranger[x:%d/y:%d]", ranger.X, ranger.Y)
	ebitenutil.DebugPrintAt(screen, msg, 500, 0)

	msg = fmt.Sprintf("initCalled:%d", initCalled)
	ebitenutil.DebugPrintAt(screen, msg, 500, 25)

	msg = fmt.Sprintf("move_y:%d", last_move_y)
	ebitenutil.DebugPrintAt(screen, msg, 500, 50)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return game.ScreenWidth, game.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(game.ScreenWidth, game.ScreenHeight)
	ebiten.SetWindowTitle(game.Caption)
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
