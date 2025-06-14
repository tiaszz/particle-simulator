package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (p *Particle) Update() {
	// gravity
	p.ay = Gravity

	// update velocity
	p.vx += p.ax * Dt
	p.vy += p.ay * Dt

	// update position
	p.x += p.vx * Dt
	p.y += p.vy * Dt

	if p.x-10 < 0 {
		p.x = 10
		p.vx = -p.vx * 0.8
	}

	if p.x+10 > ScreenWidth {
		p.x = ScreenWidth - 10
		p.vx = -p.vx * 0.8
	}

	if p.y-10 < 0 {
		p.y = 10
		p.vy = -p.vy * 0.8
	}

	if p.y+10 > ScreenHeight {
		p.y = ScreenHeight - 10
		p.vy = -p.vy * 0.8
	}

	// reset acceleration
	p.ax = 0
	p.ay = 0
}

func (g *Game) Update() error {
	g.particle.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	vector.DrawFilledCircle(screen, float32(g.particle.x), float32(g.particle.y), 10, color.RGBA{255, 0, 0, 255}, false)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 960
}

func main() {
	game := &Game{}
	game.particle.x = ScreenWidth / 2
	game.particle.y = ScreenHeight / 2

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Particle simulator")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
