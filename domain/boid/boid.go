package boid

import (
	"boids/config"
	"boids/domain/vector2d"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Boid struct {
	position *vector2d.Vector2D
	velocity *vector2d.Vector2D
	id       int
}

func (b *Boid) GetPosition() *vector2d.Vector2D {
	return b.position
}

func (b *Boid) calcAcceleration(boidsMap *[config.SCREEN_WIDTH + 1][config.SCREEN_HEIGHT + 1]int, boids *[config.BOID_COUNT]*Boid, rWlock *sync.RWMutex) *vector2d.Vector2D {
	upper, lower := b.position.AddScalar(config.VIEW_RADIUS), b.position.AddScalar(-config.VIEW_RADIUS)
	avgPosition, avgVelocity, separation := vector2d.NewVec(0, 0), vector2d.NewVec(0, 0), vector2d.NewVec(0, 0)
	count := 0.0
	rWlock.RLock()
	for i := math.Max(lower.GetX(), 0); i <= math.Min(upper.GetX(), config.SCREEN_WIDTH); i++ {
		for j := math.Max(lower.GetY(), 0); j <= math.Min(upper.GetY(), config.SCREEN_HEIGHT); j++ {
			if otherBoidId := boidsMap[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != b.id {
				if dist := boids[otherBoidId].position.Distance(b.position); dist < config.VIEW_RADIUS {
					count++
					avgVelocity = avgVelocity.Add(boids[otherBoidId].velocity)
					avgPosition = avgPosition.Add(boids[otherBoidId].position)
					separation = separation.Add(b.position.Subtract(boids[otherBoidId].position).DivisionScalar(dist))
				}
			}
		}
	}
	rWlock.RUnlock()

	accel := vector2d.NewVec(b.borderBounce(b.position.GetX(), config.SCREEN_WIDTH), b.borderBounce(b.position.GetY(), config.SCREEN_HEIGHT))
	if count > 0 {
		avgPosition, avgVelocity = avgPosition.DivisionScalar(count), avgVelocity.DivisionScalar(count)
		accelAlignment := avgVelocity.Subtract(b.velocity).MultiplyScalar(config.ADJ_RATE)
		accelCohesion := avgPosition.Subtract(b.position).MultiplyScalar(config.ADJ_RATE)
		accelSeparation := separation.MultiplyScalar(config.ADJ_RATE)
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeparation)
	}
	return accel
}

func (b *Boid) borderBounce(pos, maxBorderPos float64) float64 {
	if pos < config.VIEW_RADIUS {
		return 1 / pos
	} else if pos > maxBorderPos-config.VIEW_RADIUS {
		return 1 / (pos - maxBorderPos)
	}
	return 0
}

func (b *Boid) moveOne(boidsMap *[config.SCREEN_WIDTH + 1][config.SCREEN_HEIGHT + 1]int, boids *[config.BOID_COUNT]*Boid, rWlock *sync.RWMutex) {
	acceleration := b.calcAcceleration(boidsMap, boids, rWlock)
	rWlock.Lock()
	b.velocity = b.velocity.Add(acceleration).Limit(-1, 1)
	boidsMap[int(b.position.GetX())][int(b.position.GetY())] = -1
	b.position = b.position.Add(b.velocity)
	boidsMap[int(b.position.GetX())][int(b.position.GetY())] = b.id
	rWlock.Unlock()
}

func (b *Boid) start(boidsMap *[config.SCREEN_WIDTH + 1][config.SCREEN_HEIGHT + 1]int, boids *[config.BOID_COUNT]*Boid, rWlock *sync.RWMutex) {
	for {
		b.moveOne(boidsMap, boids, rWlock)
		time.Sleep(5 * time.Millisecond)
	}
}

func NewBoid(bid int, boidsMap *[config.SCREEN_WIDTH + 1][config.SCREEN_HEIGHT + 1]int, boids *[config.BOID_COUNT]*Boid, rWlock *sync.RWMutex) {
	b := Boid{
		position: vector2d.NewVec(rand.Float64()*config.SCREEN_WIDTH, rand.Float64()*config.SCREEN_HEIGHT),
		velocity: vector2d.NewVec((rand.Float64()*2)-1.0, (rand.Float64()*2)-1.0),
		id:       bid,
	}
	boids[bid] = &b
	boidsMap[int(b.position.GetX())][int(b.position.GetY())] = b.id
	go b.start(boidsMap, boids, rWlock)
}
