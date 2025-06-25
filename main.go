package main

import (
	"image/color"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	mx, my := ebiten.CursorPosition()

	if inpututil.IsKeyJustPressed(ebiten.KeyE) {
		g.isAttractorMode = !g.isAttractorMode
	}
	var gravityStrength float64
	if g.isAttractorMode {
		gravityStrength = 800.0
	} else {
		gravityStrength = -800.0
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		newParticle := &Particle{
			x:  float64(mx),
			y:  float64(my),
			vx: (rand.Float64() - 0.5) * 4,
			vy: (rand.Float64() - 0.5) * 4,
		}

		g.particle = append(g.particle, newParticle)
	}

	for _, p := range g.particle {
		dx := float64(mx) - p.x
		dy := float64(my) - p.y

		dist := math.Sqrt(dx*dx + dy*dy)

		if dist > 1 {
			normDx := dx / dist
			normDy := dy / dist

			force := gravityStrength / (dist*dist + 10000) * 10000

			accelX := normDx * force
			accelY := normDy * force

			p.ax += accelX
			p.ay += accelY

			if !g.isAttractorMode {
				p.ay -= 2.1
			}
		}

		p.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, p := range g.particle {
		vector.DrawFilledCircle(screen, float32(p.x), float32(p.y), 10, color.RGBA{255, 0, 0, 255}, false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 1280, 960
}

func main() {
	game := &Game{
		particle:        make([]*Particle, 0),
		isAttractorMode: true,
	}

	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Particle simulator")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
