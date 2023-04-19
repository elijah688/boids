package main

import (
	"boids/config"
	"boids/domain/boid"
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	green   = color.RGBA{10, 255, 50, 255}
	boids   [config.BOID_COUNT]*boid.Boid
	boidMap [config.SCREEN_WIDTH + 1][config.SCREEN_HEIGHT + 1]int
	rWLock  = new(sync.RWMutex)
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, b := range boids {
		screen.Set(int(b.GetPosition().GetX()+1), int(b.GetPosition().GetY()), green)
		screen.Set(int(b.GetPosition().GetX()-1), int(b.GetPosition().GetY()), green)
		screen.Set(int(b.GetPosition().GetX()), int(b.GetPosition().GetY()-1), green)
		screen.Set(int(b.GetPosition().GetX()), int(b.GetPosition().GetY()+1), green)
	}
}

func (g *Game) Layout(_, _ int) (w, h int) {
	return config.SCREEN_WIDTH, config.SCREEN_HEIGHT
}

func main() {
	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}

	for i := 0; i < config.BOID_COUNT; i++ {
		boid.NewBoid(i, &boidMap, &boids, rWLock)
	}
	ebiten.SetWindowSize(config.SCREEN_WIDTH*2, config.SCREEN_HEIGHT*2)
	ebiten.SetWindowTitle("Boids in a box")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
