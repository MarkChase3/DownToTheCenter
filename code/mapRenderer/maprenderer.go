package mapRenderer

import (
	"DownToTheCenter/fs"
	"encoding/json"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var spr []*ebiten.Image

func loadSprite(tileset *ebiten.Image, clip image.Rectangle) {
	spr = append(spr, nil)
	spr[len(spr)-1] = tileset
	spr[len(spr)-1] = ebiten.NewImageFromImage(spr[len(spr)-1].SubImage(clip))
}

var Height int16
var Width int16
var Filteredlayers [][]int8
var CamX int16 = 0
var CamY int16 = 0

type EntityLayer struct {
	Eid            string `json:"_eid"`
	Entities       []int8 `json:"entities"`
	GridCellWidth  int8   `json:"gridCellWidth"`
	GridCellHeight int8   `json:"gridCellHeight"`
	GridCellsX     int8   `json:"gridCellsX"`
	GridCellsY     int8   `json:"gridCellsY"`
	Name           string `json:"name"`
	OffsetX        int8   `json:"offsetX"`
	OffsetY        int8   `json:"offsetY"`
}
type TileLayer struct {
	Eid            string `json:"_eid"`
	Tiles          []int8 `json:"data"`
	GridCellWidth  int8   `json:"gridCellWidth"`
	GridCellHeight int8   `json:"gridCellHeight"`
	GridCellsX     int8   `json:"gridCellsX"`
	GridCellsY     int8   `json:"gridCellsY"`
	Name           string `json:"name"`
	OffsetX        int8   `json:"offsetX"`
	OffsetY        int8   `json:"offsetY"`
	Tileset        string `json:"tileset"`
	ArrayMode      int8   `json:"arrayMode"`
	ExportMode     int8   `json:"exportMode"`
}

func init() {
	var jsonfile []byte
	jsonfile = fs.LoadFile("maps/jsons/map1.json")
	var gamemap map[string]interface{}
	json.Unmarshal([]byte(jsonfile), &gamemap)
	layers, _ := gamemap["layers"].([]interface{})
	entitylayer := layers[0].(map[string]interface{})
	entitylayerjson, _ := json.Marshal(entitylayer)
	var realentitylayer EntityLayer
	json.Unmarshal([]byte(entitylayerjson), &realentitylayer)
	floorlayer := layers[1].(map[string]interface{})
	floorlayerjson, _ := json.Marshal(floorlayer)
	var realfloorlayer TileLayer
	json.Unmarshal([]byte(floorlayerjson), &realfloorlayer)
	walllayer := layers[2].(map[string]interface{})
	walllayerjson, _ := json.Marshal(walllayer)
	var realwalllayer TileLayer
	json.Unmarshal([]byte(walllayerjson), &realwalllayer)
	Width = int16(realwalllayer.GridCellsX)
	Height = int16(realwalllayer.GridCellsY)
	Filteredlayers = append(Filteredlayers, realentitylayer.Entities)
	Filteredlayers = append(Filteredlayers, realfloorlayer.Tiles)
	Filteredlayers = append(Filteredlayers, realwalllayer.Tiles)
	last := make([]int8, (Width+1)*2)
	for i := 0; i < int((Width+1)*2); i++ {
		last[i] = -1
	}
	Filteredlayers[0] = append(Filteredlayers[0], last...)
	Filteredlayers[1] = append(Filteredlayers[1], last...)
	Filteredlayers[2] = append(Filteredlayers[2], last...)
	tileset, _, _ := ebitenutil.NewImageFromFile("images/tileset1.png")
	loadSprite(tileset, image.Rectangle{image.Point{48, 48}, image.Point{64, 64}})
	loadSprite(tileset, image.Rectangle{image.Point{0, 0}, image.Point{16, 16}})
	loadSprite(tileset, image.Rectangle{image.Point{16, 0}, image.Point{32, 16}})
	loadSprite(tileset, image.Rectangle{image.Point{32, 0}, image.Point{48, 16}})
	loadSprite(tileset, image.Rectangle{image.Point{48, 0}, image.Point{64, 16}})
	loadSprite(tileset, image.Rectangle{image.Point{0, 16}, image.Point{16, 32}})
	loadSprite(tileset, image.Rectangle{image.Point{16, 16}, image.Point{32, 32}})
	loadSprite(tileset, image.Rectangle{image.Point{32, 16}, image.Point{48, 32}})
	loadSprite(tileset, image.Rectangle{image.Point{48, 16}, image.Point{64, 32}})
	loadSprite(tileset, image.Rectangle{image.Point{0, 32}, image.Point{16, 48}})
	loadSprite(tileset, image.Rectangle{image.Point{16, 32}, image.Point{32, 48}})
	loadSprite(tileset, image.Rectangle{image.Point{32, 32}, image.Point{48, 48}})
	loadSprite(tileset, image.Rectangle{image.Point{48, 32}, image.Point{64, 48}})
	loadSprite(tileset, image.Rectangle{image.Point{0, 48}, image.Point{16, 64}})
	loadSprite(tileset, image.Rectangle{image.Point{16, 48}, image.Point{32, 64}})
	loadSprite(tileset, image.Rectangle{image.Point{32, 48}, image.Point{48, 64}})
	loadSprite(tileset, image.Rectangle{image.Point{48, 48}, image.Point{64, 64}})
}

func Draw(screen *ebiten.Image) {
	for i := CamX / 16; i < CamX/16+21; i++ {
		for j := CamY / 16; j < CamY/16+13; j++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Scale(float64(1), float64(1))
			op.GeoM.Translate(float64(int16(i)*16-CamX), float64(int16(j)*16-(CamY)))
			screen.DrawImage(spr[Filteredlayers[2][(int16(j))*Width+(int16(i))]+1], op)
			screen.DrawImage(spr[Filteredlayers[1][(int16(j))*Width+(int16(i))]+1], op)
		}
	}
}
func max(y int16, x int16) int16 {
	if x > y {
		return x
	}
	return y
}
func Overlaps(x int16, y int16) bool {
	for i := max(0, x/16-3); i < x/16+3; i++ {
		for j := max(0, y/16-3); j < y/16+3; j++ {
			if Filteredlayers[1][int(j*Width)+int(i)] != -1 &&
				((int16(i))*16)+16 > x && x+16 > ((int16(i))*16) &&
				((int16(j))*16)+16 > y && y+16 > ((int16(j))*16) {
				return true
			}
		}
	}
	return false
}
