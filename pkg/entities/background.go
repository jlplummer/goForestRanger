package entities

import "goForestRanger/pkg/game"

type Background struct {
	Rows    int
	Columns int
}

func (b *Background) Init() {
	b.Rows = game.ScreenHeight / game.TileHeight
	b.Columns = game.ScreenWidth / game.TileWidth
}

func (b *Background) Update() {}

func (b *Background) Draw() {}
