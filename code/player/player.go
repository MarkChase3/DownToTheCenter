package player

import (
	"DownToTheCenter/fs"
	"DownToTheCenter/input"
	"DownToTheCenter/items"
	"DownToTheCenter/mapRenderer"
	"bytes"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// import "fmt"
var spr []*ebiten.Image
var currentSprite uint8 = 0
var X, Y float64
var flip bool
var itens [2]items.Item

func maX(Y int16, X int16) int16 {
	if X > Y {
		return X
	}
	return Y
}

func min(Y int16, X int16) int16 {
	if X < Y {
		return X
	}
	return Y
}
func loadSprite(tileset *ebiten.Image, clip image.Rectangle) {
	spr = append(spr, nil)
	spr[len(spr)-1] = tileset
	spr[len(spr)-1] = ebiten.NewImageFromImage(spr[len(spr)-1].SubImage(clip))
}

func init() {
	items.LoadItems()
	itens[0] = items.Item{items.FireBall, 98, 103, false, false, []byte{}, 0, ebiten.MouseButtonLeft, []items.Projectil{}}
	itens[1] = items.Item{items.Bow, 105, 103, false, false, []byte{}, 0, ebiten.MouseButtonRight, []items.Projectil{}}
	flip = false
	X, Y = 100, 100
	img, _, _ := image.Decode(bytes.NewReader(fs.LoadFile("images/player.png")))
	tileset := ebiten.NewImageFromImage(img)
	loadSprite(tileset, image.Rectangle{image.Point{0, 0}, image.Point{16, 16}})
	loadSprite(tileset, image.Rectangle{image.Point{16, 0}, image.Point{32, 16}})
	loadSprite(tileset, image.Rectangle{image.Point{0, 16}, image.Point{16, 32}})
	loadSprite(tileset, image.Rectangle{image.Point{16, 16}, image.Point{32, 32}})
}

func Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	item1op := &ebiten.DrawImageOptions{}
	item2op := &ebiten.DrawImageOptions{}
	if !flip {
		op.GeoM.Scale(float64(1), float64(1))
		op.GeoM.Translate(float64(X)-float64(mapRenderer.CamX), float64(Y)-float64(mapRenderer.CamY))
		item1op.GeoM.Translate(float64(itens[0].X)-float64(mapRenderer.CamX), float64(itens[0].Y)-float64(mapRenderer.CamY))
		item2op.GeoM.Translate(float64(itens[1].X)-float64(mapRenderer.CamX), float64(itens[1].Y)-float64(mapRenderer.CamY))
	} else {
		op.GeoM.Scale(-1, 1)
		item1op.GeoM.Scale(-1, 1)
		item2op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(X)+float64(16.0)-float64(mapRenderer.CamX), float64(Y)-float64(mapRenderer.CamY))
		item1op.GeoM.Translate(float64(itens[0].X)-float64(mapRenderer.CamX)+float64(20), float64(itens[0].Y)-float64(mapRenderer.CamY))
		item2op.GeoM.Translate(float64(itens[1].X)-float64(mapRenderer.CamX)+float64(5), float64(itens[1].Y)-float64(mapRenderer.CamY))
	}
	screen.DrawImage(spr[currentSprite], op)
	screen.DrawImage(itens[0].Classification.Spr[itens[0].CurrentSprite], item1op)
	screen.DrawImage(itens[1].Classification.Spr[itens[1].CurrentSprite], item2op)
	for j := 0; j < 2; j++ {
		for i := 0; i < len(itens[j].Projectiles); i++ {
			projecop := &ebiten.DrawImageOptions{}
			projecop.GeoM.Translate(float64(itens[j].Projectiles[i].X)-float64(mapRenderer.CamX), float64(itens[j].Projectiles[i].Y)-float64(mapRenderer.CamY))
			screen.DrawImage(itens[j].Classification.Spr[itens[j].CurrentSprite], projecop)
		}
	}
}

func Update() {
	itens[0].Classification.Update(&itens[0], flip)
	itens[1].Classification.Update(&itens[1], flip)
	for j := 0; j < 2; j++ {
		for i := 0; i < len(itens[j].Projectiles); i++ {
			itens[j].Projectiles[i].X += itens[j].Projectiles[i].SpdX
			itens[j].Projectiles[i].Y += itens[j].Projectiles[i].SpdY
			if mapRenderer.Overlaps(int16(itens[j].Projectiles[i].X), int16(itens[j].Projectiles[i].Y)) {
				itens[j].Classification.WhenCollide(&itens[j].Projectiles, uint16(i))
			}
		}
	}
	var colX float64
	colX = X
	if input.IsMoving {
		Y += input.Y * 1.5
		itens[0].Y += float32(input.Y * 1.5)
		itens[1].Y += float32(input.Y * 1.5)
	}
	if mapRenderer.Overlaps(int16(colX), int16(Y)) {
		Y -= input.Y * 1.5
		itens[0].Y -= float32(input.Y * 1.5)
		itens[1].Y -= float32(input.Y * 1.5)
	}
	if input.IsMoving {
		X += input.X * 1.5
		colX += input.X * 1.5
		///flip = false
		itens[0].X += float32(input.X * 1.5)
		itens[1].X += float32(input.X * 1.5)
	}
	if mapRenderer.Overlaps(int16(colX), int16(Y)) {
		X -= input.X * 3
		colX -= input.X * 3
		itens[0].X -= float32(input.X * 1.5)
		itens[1].X -= float32(input.X * 1.5)
	}
	mapRenderer.CamX = min(maX(0, int16(X-160)), mapRenderer.Width*16-20*16)
	mapRenderer.CamY = min(maX(0, int16(Y-90)), mapRenderer.Height*16-11*16)
}
