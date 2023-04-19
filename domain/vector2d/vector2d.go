package vector2d

import "math"

type Vector2D struct {
	x, y float64
}

func NewVec(x, y float64) *Vector2D {
	return &Vector2D{x, y}
}

func (v *Vector2D) GetX() float64 {
	return v.x
}

func (v *Vector2D) GetY() float64 {
	return v.y
}

func (v *Vector2D) Add(otherVector *Vector2D) *Vector2D {
	return &Vector2D{
		x: v.x + otherVector.x,
		y: v.y + otherVector.y,
	}
}

func (v *Vector2D) Subtract(otherVector *Vector2D) *Vector2D {
	return &Vector2D{
		x: v.x - otherVector.x,
		y: v.y - otherVector.y,
	}
}

func (v *Vector2D) Multiply(otherVector *Vector2D) *Vector2D {
	return &Vector2D{
		x: v.x * otherVector.x,
		y: v.y * otherVector.y,
	}
}

func (v *Vector2D) AddScalar(d float64) *Vector2D {
	return &Vector2D{x: v.x + d, y: v.y + d}
}

func (v *Vector2D) MultiplyScalar(d float64) *Vector2D {
	return &Vector2D{x: v.x * d, y: v.y * d}
}

func (v *Vector2D) DivisionScalar(d float64) *Vector2D {
	return &Vector2D{x: v.x / d, y: v.y / d}
}

func (v1 *Vector2D) Limit(lower, upper float64) *Vector2D {
	return NewVec(math.Min(math.Max(v1.x, lower), upper),
		math.Min(math.Max(v1.y, lower), upper))
}

func (v *Vector2D) Distance(otherVector *Vector2D) float64 {
	return math.Sqrt(math.Pow((v.x-otherVector.x), 2) + math.Pow((v.y-otherVector.y), 2))

}
