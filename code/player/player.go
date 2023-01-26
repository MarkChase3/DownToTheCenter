package player

import "github.com/hajimehoshi/ebiten/v2"
import "image"
import "bytes"
import "fmt"
import "DownToTheCenter/mapRenderer"
import "DownToTheCenter/fs"
import "DownToTheCenter/items"
var spr []*ebiten.Image
var currentSprite uint8 = 0
var x, y float64
var flip bool
var itens [2]items.Item
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
	itens[0] = items.Item{items.Sword, 98, 103, false, false, []byte{}, 0, ebiten.MouseButtonLeft, []items.Projectil{}}
	itens[1] = items.Item{items.Bow, 106, 103, false, false, []byte{}, 0, ebiten.MouseButtonRight, []items.Projectil{}}
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
	op := &ebiten.DrawImageOptions{}
	item1op := &ebiten.DrawImageOptions{}
	item2op := &ebiten.DrawImageOptions{}
	if !flip {
	    op.GeoM.Scale(float64(1), float64(1))
        op.GeoM.Translate(float64(x) - float64(mapRenderer.CamX), float64(y) - float64(mapRenderer.CamY))
		item1op.GeoM.Translate(float64(itens[0].X) - float64(mapRenderer.CamX), float64(itens[0].Y) - float64(mapRenderer.CamY))
		item2op.GeoM.Translate(float64(itens[1].X) - float64(mapRenderer.CamX), float64(itens[1].Y) - float64(mapRenderer.CamY))
	} else {
	    op.GeoM.Scale(-1, 1)
	    item1op.GeoM.Scale(-1, 1)
	    item2op.GeoM.Scale(-1, 1)
        op.GeoM.Translate(float64(x) + float64(16.0) - float64(mapRenderer.CamX), float64(y) - float64(mapRenderer.CamY))
		item1op.GeoM.Translate(float64(itens[0].X) - float64(mapRenderer.CamX) + float64(20), float64(itens[0].Y) - float64(mapRenderer.CamY))
		item2op.GeoM.Translate(float64(itens[1].X) - float64(mapRenderer.CamX) + float64(5), float64(itens[1].Y) - float64(mapRenderer.CamY))
	}
	screen.DrawImage(spr[currentSprite], op)
	screen.DrawImage(itens[0].Classification.Spr[itens[0].CurrentSprite], item1op)
	screen.DrawImage(itens[1].Classification.Spr[itens[1].CurrentSprite], item2op)
    for i := 0; i < len(itens[0].Projectiles); i++ {
		println("!KJhdhnjvgbbbbnngj")
		projecop := &ebiten.DrawImageOptions{}
        projecop.GeoM.Translate(float64(itens[0].Projectiles[i].X) - float64(mapRenderer.CamX), float64(itens[0].Projectiles[i].X) - float64(mapRenderer.CamY))
		screen.DrawImage(itens[0].Classification.Spr[itens[0].CurrentSprite], projecop)
	}

    for i := 0; i < len(itens[1].Projectiles); i++ {
		fmt.Println(float64(itens[1].Projectiles[i].X) - float64(mapRenderer.CamX))
		projecop := &ebiten.DrawImageOptions{}
        projecop.GeoM.Translate(float64(itens[1].Projectiles[i].X) - float64(mapRenderer.CamX), float64(itens[1].Projectiles[i].X) - float64(mapRenderer.CamY))
		screen.DrawImage(itens[1].Classification.Spr[itens[1].CurrentSprite], projecop)
	}

}

func Update(){
    itens[0].Classification.Update(&itens[0], flip)
    itens[1].Classification.Update(&itens[1], flip)
	for i := 0; i < len(itens[0].Projectiles); i++ {
		itens[0].Projectiles[i].X += itens[0].Projectiles[i].SpdX
		itens[0].Projectiles[i].Y += itens[0].Projectiles[i].SpdY
	}
	for i := 0; i < len(itens[1].Projectiles); i++ {
	    itens[1].Projectiles[i].X += itens[1].Projectiles[i].SpdX
	    itens[1].Projectiles[i].Y += itens[1].Projectiles[i].SpdY
	}
	var colx float64
	if flip {
		colx = x
	}  else {
		colx = x	
	}
	if(ebiten.IsKeyPressed(ebiten.KeyW)){
		y-=1.5
		itens[0].Y -= 1.5
		itens[1].Y -= 1.5
	}
	if mapRenderer.Overlaps(int16(colx), int16(y)) {
		y += 1.5
		itens[0].Y += 1.5
		itens[1].Y += 1.5
	}
    if(ebiten.IsKeyPressed(ebiten.KeyS)){
		y+=1.5
		itens[0].Y += 1.5
		itens[1].Y += 1.5
	}
	if mapRenderer.Overlaps(int16(colx), int16(y)) {
		y -= 1.5
		itens[0].Y -= 1.5
		itens[1].Y -= 1.5
	}
    if(ebiten.IsKeyPressed(ebiten.KeyA)){
		x-=1.5
		colx-=1.5
		flip = false
		itens[0].X -= 1.5
		itens[1].X -= 1.5
	}
	if mapRenderer.Overlaps(int16(colx), int16(y)) {
		x += 1.5
		colx+=1.5
		itens[0].X += 1.5
		itens[1].X += 1.5
	}
	if(ebiten.IsKeyPressed(ebiten.KeyD)){
		x+=1.5
		colx+=1.5
		flip = true
		itens[0].X += 1.5
		itens[1].X += 1.5
	}
	if mapRenderer.Overlaps(int16(colx), int16(y)) {
		x -= 1.5
		colx -= 1.5
		itens[0].X -= 1.5
		itens[1].X -= 1.5
	}
    mapRenderer.CamX = min(max(0, int16(x - 160)), mapRenderer.Width*16-20*16)
    mapRenderer.CamY = min(max(0, int16(y - 90)), mapRenderer.Height*16-11*16)
}