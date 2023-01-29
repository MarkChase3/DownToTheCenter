//go:build !js
// +build !js

package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var Angle float64
var IsMoving bool
var X float64
var Y float64

func IsTargeting() bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func TargetPos() (int, int) {
	return ebiten.CursorPosition()
}

func UpdateInput() {
	X, Y = 0, 0
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		Y -= 1.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		X += 1.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		X -= 1.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		Y += 1.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyD) {
		IsMoving = true
	} else {
		IsMoving = false
	}
}

func Draw(screen *ebiten.Image) {

}
