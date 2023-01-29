package items

import (
	"DownToTheCenter/fs"
	"DownToTheCenter/input"
	"DownToTheCenter/mapRenderer"
	"bytes"
	"image"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type ItemType struct {
	Spr         []*ebiten.Image
	Update      func(item *Item, flip bool)
	WhenCollide func(projectiles *[]Projectil, i uint16)
}
type Projectil struct {
	X         float32
	Y         float32
	SpdX      float32
	SpdY      float32
	IsBouncer bool
	Life      uint16
	//SprIndex uint8
}
type Item struct {
	Classification ItemType
	X              float32
	Y              float32
	Active         bool
	Onthefloor     bool
	State          []byte
	CurrentSprite  uint8
	Hand           ebiten.MouseButton
	Projectiles    []Projectil
}

var Sword ItemType
var Bow ItemType
var FireBall ItemType

func NewItem(path string, updt func(item *Item, flip bool), cold func(projectiles *[]Projectil, i uint16)) ItemType {
	c, _, _ := image.DecodeConfig(bytes.NewReader(fs.LoadFile(path)))
	tilesetImage, _, _ := image.Decode(bytes.NewReader(fs.LoadFile(path)))
	tileset := ebiten.NewImageFromImage(tilesetImage)

	spr := []*ebiten.Image{}
	for i := 0; i < c.Width/8; i++ {
		for j := 0; j < c.Height/8; j++ {
			clip := image.Rectangle{image.Point{i * 8, j * 8}, image.Point{i*8 + 8, j*8 + 8}}
			spr = append(spr, ebiten.NewImageFromImage(tileset.SubImage(clip)))
		}
	}
	return ItemType{spr, updt, cold}
}
func LoadItems() {
	Sword = NewItem("images/sword.png", func(item *Item, flip bool) {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			item.CurrentSprite = 2
		} else {
			item.CurrentSprite = 0
		}
	}, func(projectiles *[]Projectil, i uint16) {})
	Bow = NewItem("images/bow.png", func(item *Item, flip bool) {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			item.CurrentSprite = 2
			posx, posy := ebiten.CursorPosition()
			item.Projectiles = append(item.Projectiles, Projectil{item.X, item.Y, float32(math.Cos(math.Atan2(float64(posy+int(mapRenderer.CamY)-int(item.Y)), float64(posx+int(mapRenderer.CamX)-int(item.X))))) * 3, float32(math.Sin(math.Atan2(float64(posy+int(mapRenderer.CamY)-int(item.Y)), float64(posx+int(mapRenderer.CamX)-int(item.X))))) * 3, false, 4})
		} else {
			item.CurrentSprite = 0
		}
	}, func(projectiles *[]Projectil, i uint16) {
		*projectiles = append((*projectiles)[:i], (*projectiles)[i+1:]...)
	})
	FireBall = NewItem("images/fireball.png", func(item *Item, flip bool) {
		if input.IsTargeting() {
			item.CurrentSprite = 2
			print(input.TargetPos())
			posx, posy := input.TargetPos()
			item.Projectiles = append(item.Projectiles, Projectil{item.X, item.Y, float32(math.Cos(math.Atan2(float64(posy+int(mapRenderer.CamY)-int(item.Y)), float64(posx+int(mapRenderer.CamX)-int(item.X))))) * 3, float32(math.Sin(math.Atan2(float64(posy+int(mapRenderer.CamY)-int(item.Y)), float64(posx+int(mapRenderer.CamX)-int(item.X))))) * 3, true, 4})
		} else {
			item.CurrentSprite = 0
		}
	}, func(projectiles *[]Projectil, i uint16) {
		if rand.Intn(2) < 1 {
			(*projectiles)[i].SpdX = -(*projectiles)[i].SpdX
		} else {
			(*projectiles)[i].SpdY = -(*projectiles)[i].SpdY
		}
		(*projectiles)[i].Life--
		if (*projectiles)[i].Life <= 0 {
			*projectiles = append((*projectiles)[:i], (*projectiles)[i+1:]...)
		}
	})
}
