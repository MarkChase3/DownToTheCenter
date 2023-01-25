package player

import "github.com/hajimehoshi/ebiten/v2"
// import "github.com/hajimehoshi/ebiten/v2/ebitenutil"
import "image"
import "bytes"
import "DownToTheCenter/mapRenderer"
import "DownToTheCenter/fs"
import "DownToTheCenter/items"
var spr []*ebiten.Image
var currentSprite uint8 = 0
var x, y float64
var flip bool
var itens[2] items.Item
func max(y int16, x int16) int16 {
    if x > y {
		return x
	}
	return y
}

func min(y int16, x int16) int16 {
    if x < y {
		return x
	}
	return y
}
func loadSprite(tileset *ebiten.Image, clip image.Rectangle){
	spr = append(spr, nil)
	spr[len(spr)-1] = tileset
	spr[len(spr)-1] = ebiten.NewImageFromImage(spr[len(spr)-1].SubImage(clip))
}

func init() {
	items.LoadItems()
	itens[0] = items.Item{items.Sword, 100, 100, false, false, "inital", 0}
	flip = false
	x, y = 100, 100
	img, _, _ := image.Decode(bytes.NewReader(fs.LoadFile("images/player.png")))
	tileset := ebiten.NewImageFromImage(img)
	loadSprite(tileset, image.Rectangle{image.Point{0, 0}, image.Point{16, 16}})
	loadSprite(tileset, image.Rectangle{image.Point{16, 0}, image.Point{32, 16}})
	loadSprite(tileset, image.Rectangle{image.Point{0, 16}, image.Point{16, 32}})
	loadSprite(tileset, image.Rectangle{image.Point{16, 16}, image.Point{32, 32}})
}

func Draw(screen *ebiten.Image){
    itemop := &ebiten.DrawImageOptions{}
	itemop.GeoM.Translate(float64(x) - float64(mapRenderer.CamX), float64(y) - float64(mapRenderer.CamY))
	screen.DrawImage(itens[0].Classification.Spr[itens[0].CurrentSprite], itemop)
	op := &ebiten.DrawImageOptions{}
    if !flip {
	    op.GeoM.Scale(float64(1), float64(1))
        op.GeoM.Translate(float64(x) - float64(mapRenderer.CamX), float64(y) - float64(mapRenderer.CamY))
	} else {
	    op.GeoM.Scale(-1, 1)
        op.GeoM.Translate(float64(x) + float64(16.0) - float64(mapRenderer.CamX), float64(y) - float64(mapRenderer.CamY))
	}
	screen.DrawImage(spr[currentSprite], op)
}

func Update(){
    itens[0].Classification.Update(itens[0])
	var colx float64
	if flip {
		colx = x
	}  else {
		colx = x	
	}
	if(ebiten.IsKeyPressed(ebiten.KeyW)){
		y-=1.5
	}
	if mapRenderer.Overlaps(int16(colx), int16(y)) {
		y += 1.5
	}
    if(ebiten.IsKeyPressed(ebiten.KeyS)){
		y+=1.5
	}
	if mapRenderer.Overlaps(int16(colx), int16(y)) {
		y -= 1.5
	}
    if(ebiten.IsKeyPressed(ebiten.KeyA)){
		x-=1.5
		colx-=1.5
		flip = false
	}
	if mapRenderer.Overlaps(int16(colx), int16(y)) {
		x += 1.5
		colx+=1.5
	}
	if(ebiten.IsKeyPressed(ebiten.KeyD)){
		x+=1.5
		colx+=1.5
		flip = true
	}
	if mapRenderer.Overlaps(int16(colx), int16(y)) {
		x -= 1.5
		colx -= 1.5
	}
    mapRenderer.CamX = min(max(0, int16(x - 160)), mapRenderer.Width*16-20*16)
    mapRenderer.CamY = min(max(0, int16(y - 90)), mapRenderer.Height*16-11*16)
}