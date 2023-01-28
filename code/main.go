package main

import (
	"DownToTheCenter/enemies"
	"DownToTheCenter/mapRenderer"
	"DownToTheCenter/player"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

//import "github.com/hajimehoshi/ebiten/v2/ebitenutil"

type Game struct{}

func (g *Game) Update() error {
	player.Update()
	enemies.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	mapRenderer.Draw(screen)
	enemies.Draw(screen)
	player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 180
}

func main() {
	enemies.Start()
	ebiten.SetWindowSize(640, 360)
	ebiten.SetWindowTitle("Down To The Center")
	ebiten.RunGame(&Game{})
}
