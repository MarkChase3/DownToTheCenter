package main

import "github.com/hajimehoshi/ebiten/v2"
//import "github.com/hajimehoshi/ebiten/v2/ebitenutil"
import "image/color"
import "DownToTheCenter/player"
import "DownToTheCenter/mapRenderer"

type Game struct{}

func (g *Game) Update() error {
    player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})
	mapRenderer.Draw(screen)
	player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 180
}

func main() {
	ebiten.SetWindowSize(640, 360)
	ebiten.SetWindowTitle("Down To The Center")
	ebiten.RunGame(&Game{})
}