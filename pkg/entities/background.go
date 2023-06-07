package entities

import (
	"goForestRanger/pkg/game"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Background struct {
	Rows            int
	Columns         int
	Image           *ebiten.Image
	BackgroundTiles [][]*ebiten.Image
	Inited          bool
}

func (b *Background) Init() {
	var err error
	b.Image, _, err = ebitenutil.NewImageFromFile("media/grass_8x8.png")
	if err != nil {
		log.Fatal("Cannot open grass_8x8.png")
	}

	b.Rows = game.ScreenHeight / game.TileHeight  // Y
	b.Columns = game.ScreenWidth / game.TileWidth // X
	// fmt.Printf("b.Rows: %d\nb.Columns: %d\n", b.Rows, b.Columns)

	// // generate a random background
	b.BackgroundTiles = make([][]*ebiten.Image, b.Rows)
	// fmt.Printf("b.BackgroundTiles = %v\n", b.BackgroundTiles)
	for j, _ := range b.BackgroundTiles {

		b.BackgroundTiles[j] = make([]*ebiten.Image, b.Columns)
		// fmt.Printf("b.BackgroundTiles[%d] = %v\n", j, b.BackgroundTiles[j])
		for g, _ := range b.BackgroundTiles[j] {
			startX := 0
			startY := 0
			endX := 7
			endY := 7

			b.BackgroundTiles[j][g] = ebiten.NewImageFromImage(b.Image.SubImage(image.Rect(startX, startY, endX, endY)))
			// fmt.Printf("in init(): b.BackgroundTiles[%d][%d] = %v\n", j, g, b.BackgroundTiles[j][g])
		}
	}

	// for j := 0; j < b.Columns; j++ {
	// 	innerLen := b.Rows + 1
	// 	b.BackgroundTiles[j] = make([]*ebiten.Image, innerLen)
	// 	for g := 0; g < innerLen; g++ {

	// 		/*
	// 		// 			r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 		// 			tilePtr := r1.Intn(game.MaxGrassTiles + 1)

	// 		// 			startX := tilePtr * game.TileHeight
	// 		// 			startY := 0
	// 		// 			endX := startX + game.TileHeight
	// 		// 			endY := (startY + game.TileWidth) - 1
	// 		// 			fmt.Printf(`tilePtr: %d
	// 		// startX: %d
	// 		// startY: %d
	// 		// endX: %d
	// 		// endY: %d\n\n`, tilePtr, startX, startY, endX, endY)
	// 		*/
	// 		startX := 0
	// 		startY := 0
	// 		endX := 7
	// 		endY := 7
	// 		b.BackgroundTiles[j][g] = ebiten.NewImageFromImage(b.Image.SubImage(image.Rect(startX, startY, endX, endY)))
	// 	}

	b.Inited = true
}

// 	fmt.Println(b.BackgroundTiles)
// }

func (b *Background) Update() {}

func (b *Background) Draw(screen *ebiten.Image) {
	// Remember: X = columns, Y = rows
	var x, y int
	x = -game.TileWidth
	y = -game.TileHeight

	// rows (Y)
	for j, l := range b.BackgroundTiles {
		// columns (X)
		for g, _ := range l {
			// fmt.Printf("b.BackgroundTiles[%d][%d] = %v\n", j, g, b.BackgroundTiles[j][g])
			op := &ebiten.DrawImageOptions{}

			// figure out how to make first loop NOT add game.Tile* to the translate
			// without an IF?
			// op.GeoM.Translate(float64(x)+game.TileWidth, float64(y)+game.TileHeight)
			op.GeoM.Translate(float64(x)+game.TileWidth, float64(y)+game.TileHeight)
			op.GeoM.Scale(game.SpriteFactor, game.SpriteFactor)
			// op.GeoM.Translate((float64(x)+game.TileWidth)*game.SpriteFactor, (float64(y)+game.TileHeight)*game.SpriteFactor)

			// fmt.Printf("%d:%d>\n%v\n", j, g, op.GeoM.String())
			screen.DrawImage(b.BackgroundTiles[j][g], op)

			x += game.TileWidth
		}
		y += game.TileHeight
		x = -game.TileWidth
	}
	// log.Fatal("Stopping after first draw")

}
