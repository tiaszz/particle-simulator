package main

type Particle struct {
	x, y   float64
	vx, vy float64
	ax, ay float64
}

const Gravity = 500.0
const (
	ScreenWidth  = 1280
	ScreenHeight = 960
	Dt           = 1.0 / 60.0
)

type Game struct {
	particle Particle
}
