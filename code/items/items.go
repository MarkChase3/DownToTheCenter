package items

import "image"
import "bytes"
import "DownToTheCenter/fs"
import "github.com/hajimehoshi/ebiten/v2"
type ItemType struct {
	Spr []*ebiten.Image
    Update func (item *Item, flip bool)
}
type Projectil struct {
    X float32
    Y float32
	SpdX float32
	SpdY float32
}
type Item struct {
    Classification ItemType
	X float32
	Y float32
	Active bool
	Onthefloor bool
	State []byte
    CurrentSprite uint8
	Hand ebiten.MouseButton
	Projectiles []Projectil
}
var Sword ItemType
var Bow ItemType

func NewItem(path string, updt func (item *Item, flip bool)) ItemType {
	c, _, _ := image.DecodeConfig(bytes.NewReader(fs.LoadFile(path)))
	tilesetImage, _,  _ := image.Decode(bytes.NewReader(fs.LoadFile(path)))
	tileset := ebiten.NewImageFromImage(tilesetImage)

	spr := []*ebiten.Image{}
	for i := 0; i < c.Width/8; i++ {
		for j := 0; j < c.Height/8; j++ {
            clip := image.Rectangle{image.Point{i*8, j*8}, image.Point{i*8+8, j*8+8}}
			spr = append(spr, ebiten.NewImageFromImage(tileset.SubImage(clip)))
		}	
	}
    return ItemType{spr, updt}
}
func LoadItems(){
    Sword = NewItem("images/sword.png", func (item *Item, flip bool) {
		if ebiten.IsMouseButtonPressed(item.Hand) {
			item.CurrentSprite = 2
		} else {
			item.CurrentSprite = 0
		}
	})
	Bow = NewItem("images/bow.png", func (item *Item, flip bool) {
		if ebiten.IsMouseButtonPressed(item.Hand) {
			item.CurrentSprite = 2
			item.Projectiles = append(item.Projectiles, Projectil{item.X, item.Y, 0.001, 0.0001})
		} else {
			item.CurrentSprite = 0
		}
	})
}