// +build: wasm

package input

import "github.com/hajimehoshi/ebiten/v2"
import "math"
import "image/color"
import "fmt"

var jx int = 270
var jy int = 130
var X float64
var Y float64
var IsMoving bool

func IsTargeting() bool {
	x, y := ebiten.TouchPosition(0)
	return !(x == 0 && y == 0) || ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
}
func TargetPos() (int, int) {
	x, y := ebiten.TouchPosition(0)
	x2, y2 := ebiten.CursorPosition()
	return int(math.Max(float64(x), float64(x2))), int(math.Max(float64(y), float64(y2)))
}

func UpdateInput() {
	x, y := TargetPos()
	X, Y = 0, 0
	if math.Sqrt(float64((x-jx)*(x-jx)+(y-jy)*(y-jy))) < 100 {
		IsMoving = IsTargeting()
		X = math.Cos(math.Atan2(float64(y-jy), float64(x-jx)))
		Y = math.Sin(math.Atan2(float64(y-jy), float64(x-jx)))
		fmt.Println(math.Atan2(float64(y-jy), float64(x-jx)))
		fmt.Println("\n")
		fmt.Println(Y)
		fmt.Println("\n")
		fmt.Println(X)
		fmt.Println("\n")
	}
}

func Draw(screen *ebiten.Image) {
	purpleCol := color.RGBA{255, 0, 255, 255}

	for x := jx - 20; x < jx+20; x++ {
		for y := jy - 20; y < jy+20; y++ {
			screen.Set(x, y, purpleCol)
		}
	}
	for x := +X*10 - 20; x < X*10+20; x++ {
		for y := Y*10 - 20; y < Y*10+20; y++ {
			screen.Set(int(x), int(y), purpleCol)
		}
	}
}
