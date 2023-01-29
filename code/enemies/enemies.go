package enemies

import (
	"DownToTheCenter/fs"
	"DownToTheCenter/items"
	"DownToTheCenter/mapRenderer"
	"bytes"
	"image"
	"math/rand"

	paths "github.com/MarkChase3/original-paths-but-importable"
	"github.com/hajimehoshi/ebiten/v2"
)

type EnemyType struct {
	spr  []*ebiten.Image
	updt func(enemy *Enemy)
}
type Enemy struct {
	classification EnemyType
	X              uint16
	Y              uint16
	flip           bool
	itens          [2]items.Item
	currentSprite  uint16
}

var NEnemies uint16 = uint16(rand.Intn(100))

var Enemies []Enemy
var zombie EnemyType

func NewEnemie(path string, updt func(item *Enemy)) EnemyType {
	c, _, _ := image.DecodeConfig(bytes.NewReader(fs.LoadFile(path)))
	tilesetImage, _, _ := image.Decode(bytes.NewReader(fs.LoadFile(path)))
	tileset := ebiten.NewImageFromImage(tilesetImage)

	spr := []*ebiten.Image{}
	spr = make([]*ebiten.Image, c.Width/16*c.Width/8+c.Height)
	for i := 0; i < c.Width/16; i++ {
		for j := 0; j < c.Height/16; j++ {
			clip := image.Rectangle{image.Point{i * 16, j * 16}, image.Point{i*16 + 16, j*16 + 16}}
			spr[j*c.Width/8+i] = ebiten.NewImageFromImage(tileset.SubImage(clip))
		}
	}
	return EnemyType{spr, updt}
}
func init() {
	zombie = NewEnemie("images/zombie.png", func(enemy *Enemy) {

	})
}
func Start() {
	NEnemies = 1
	Enemies = append(Enemies, Enemy{})
	Enemies = append(Enemies, Enemy{})
	Enemies[0] = Enemy{
		zombie,
		300, 300,
		false,
		[2]items.Item{
			items.Item{
				Classification: items.FireBall,
				X:              98, Y: 103,
				Active:        false,
				Onthefloor:    false,
				State:         []byte{},
				CurrentSprite: 0,
				Hand:          ebiten.MouseButtonLeft,
				Projectiles:   []items.Projectil{}},
			items.Item{
				Classification: items.FireBall,
				X:              98, Y: 103,
				Active:        false,
				Onthefloor:    false,
				State:         []byte{},
				CurrentSprite: 0,
				Hand:          ebiten.MouseButtonLeft,
				Projectiles:   []items.Projectil{}}},
		0}

	/*for i := 0; i < NEnemies; i++ {


	}*/
}
func Update() {
	var str [][]int8
	var reallines []string
	for i := 0; i < 10; i++ {
		str = append(str, mapRenderer.Filteredlayers[1][(i+int(Enemies[0].Y/16))*int(mapRenderer.Width)-5:(i+int(Enemies[0].Y/16))*int(mapRenderer.Width)+5])
	}
	for i := 0; i < 10; i++ {
		var bits8 []byte
		for j := 0; j < 10; j++ {
			bits8 = append(bits8, byte(str[i][j]))
		}
		reallines = append(reallines, string(bits8))
	}
	grid := paths.NewGridFromStringArrays(reallines, 16, 16)
	grid.CellsByRune(1)
}

func Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	item1op := &ebiten.DrawImageOptions{}
	item2op := &ebiten.DrawImageOptions{}
	if !Enemies[0].flip {
		op.GeoM.Scale(float64(1), float64(1))
		op.GeoM.Translate(float64(Enemies[0].X)-float64(mapRenderer.CamX), float64(Enemies[0].Y)-float64(mapRenderer.CamY))
		item1op.GeoM.Translate(float64(Enemies[0].itens[0].X)-float64(mapRenderer.CamX), float64(Enemies[0].itens[0].Y)-float64(mapRenderer.CamY))
		item2op.GeoM.Translate(float64(Enemies[0].itens[1].X)-float64(mapRenderer.CamX), float64(Enemies[0].itens[1].Y)-float64(mapRenderer.CamY))
	} else {
		op.GeoM.Scale(-1, 1)
		item1op.GeoM.Scale(-1, 1)
		item2op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(Enemies[0].X)+float64(16.0)-float64(mapRenderer.CamX), float64(Enemies[0].itens[0].Y)-float64(mapRenderer.CamY))
		item1op.GeoM.Translate(float64(Enemies[0].itens[0].X)-float64(mapRenderer.CamX)+float64(20), float64(Enemies[0].itens[0].Y)-float64(mapRenderer.CamY))
		item2op.GeoM.Translate(float64(Enemies[0].itens[1].X)-float64(mapRenderer.CamX)+float64(5), float64(Enemies[0].itens[1].Y)-float64(mapRenderer.CamY))
	}
	screen.DrawImage(Enemies[0].classification.spr[Enemies[0].currentSprite], op)
	screen.DrawImage(Enemies[0].itens[0].Classification.Spr[Enemies[0].itens[0].CurrentSprite], item1op)
	screen.DrawImage(Enemies[0].itens[1].Classification.Spr[Enemies[0].itens[1].CurrentSprite], item2op)

}
