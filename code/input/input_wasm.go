// +build: wasm

package input

import "github.com/hajimehoshi/ebiten/v2"
import "math"
import "image/color"
import "fmt"

var Inside bool
var jx int = 270
var jy int = 130
var X float64
var Y float64
var IsMoving bool = false
var Shooting bool

func IsTargeting() bool {
	x, y := ebiten.TouchPosition(0)
	return !(x == 0 && y == 0) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}
func TargetPos() (int, int) {
	x, y := ebiten.TouchPosition(0)
	x2, y2 := ebiten.CursorPosition()
	return int(math.Max(float64(x), float64(x2))), int(math.Max(float64(y), float64(y2)))
}
func ShootPos() (int, int) {
	x2, y2 := ebiten.CursorPosition()
	if Inside {
		x, y := ebiten.TouchPosition(1)
		return int(math.Max(float64(x), float64(x2))), int(math.Max(float64(y), float64(y2)))
	} else {
		x, y := ebiten.TouchPosition(0)
		return int(math.Max(float64(x), float64(x2))), int(math.Max(float64(y), float64(y2)))
	}
}

func UpdateInput() {
	x, y := TargetPos()
	X, Y = 0, 0
	Inside = false
	fmt.Println(Shooting)
	if math.Sqrt(float64((x-jx)*(x-jx)+(y-jy)*(y-jy))) < 100 {
		Inside = true
		IsMoving = true
		X = math.Cos(math.Atan2(float64(y-jy), float64(x-jx)))
		Y = math.Sin(math.Atan2(float64(y-jy), float64(x-jx)))
		fmt.Println(math.Atan2(float64(y-jy), float64(x-jx)))
		fmt.Println("\n")
		fmt.Println(Y)
		fmt.Println("\n")
		fmt.Println(X)
		fmt.Println("\n")
	}
	Shooting = (len(ebiten.AppendTouchIDs([]ebiten.TouchID{})) == 1 && !Inside) || (len(ebiten.AppendTouchIDs([]ebiten.TouchID{})) == 2)
}

func Draw(screen *ebiten.Image) {
	purpleCol := color.RGBA{255, 0, 255, 255}

	for x := jx - 20; x < jx+20; x++ {
		for y := jy - 20; y < jy+20; y++ {
			screen.Set(x, y, purpleCol)
		}
	}
}
