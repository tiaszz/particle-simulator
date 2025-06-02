package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct{}

func drawCircle(screen *ebiten.Image, cx, cy, r float64, clr color.Color) {
	vector.DrawFilledCircle(screen, float32(cx), float32(cy), float32(r), clr, false) // Draw a filled circle
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawCircle(screen, 50, 50, 25, color.RGBA{R: 255, A: 255})
}

func (g *Game) Layout(screenWidth, screenHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
