package vector2d

import (
	"testing"
)

func TestAdd(t *testing.T) {
	v1 := Vector2D{1.0, 2.0}
	v2 := Vector2D{3.0, 4.0}

	result := v1.Add(&v2)

	if result.x != 4.0 || result.y != 6.0 {
		t.Errorf("Expected %v, but got %v", Vector2D{4.0, 6.0}, result)
	}
}

func TestSubtract(t *testing.T) {
	v1 := Vector2D{1.0, 2.0}
	v2 := Vector2D{3.0, 4.0}

	result := v1.Subtract(&v2)

	if result.x != -2.0 || result.y != -2.0 {
		t.Errorf("Expected %v, but got %v", Vector2D{-2.0, -2.0}, result)
	}
}

func TestMultiply(t *testing.T) {
	v1 := Vector2D{1.0, 2.0}
	v2 := Vector2D{3.0, 4.0}

	result := v1.Multiply(&v2)

	if result.x != 3.0 || result.y != 8.0 {
		t.Errorf("Expected %v, but got %v", Vector2D{3.0, 8.0}, result)
	}
}

func TestAddScalar(t *testing.T) {
	v := Vector2D{1.0, 2.0}
	scalar := 3.0

	result := v.AddScalar(scalar)

	if result.x != 4.0 || result.y != 5.0 {
		t.Errorf("Expected %v, but got %v", Vector2D{4.0, 5.0}, result)
	}
}

func TestMultiplyScalar(t *testing.T) {
	v := Vector2D{1.0, 2.0}
	scalar := 3.0

	result := v.MultiplyScalar(scalar)

	if result.x != 3.0 || result.y != 6.0 {
		t.Errorf("Expected %v, but got %v", Vector2D{3.0, 6.0}, result)
	}
}

func TestDivisionScalar(t *testing.T) {
	v := Vector2D{3.0, 6.0}
	scalar := 3.0

	result := v.DivisionScalar(scalar)

	if result.x != 1.0 || result.y != 2.0 {
		t.Errorf("Expected %v, but got %v", Vector2D{1.0, 2.0}, result)
	}
}

func TestLimit(t *testing.T) {
	og := NewVec(5.0, 8.0)
	v := og.Limit(3.0, 7.0)

	if v.x != 5.0 || v.y != 7.0 {
		t.Errorf("Expected %v, but got %v", Vector2D{5.0, 7.0}, v)
	}
	v = og.Limit(2.0, 6.0)

	if v.x != 5.0 || v.y != 6.0 {
		t.Errorf("Expected %v, but got %v", Vector2D{5.0, 6.0}, v)
	}
	v = og.Limit(1.0, 4.0)

	if v.x != 4.0 || v.y != 4.0 {
		t.Errorf("Expected %v, but got %v", Vector2D{6.0, 6.0}, v)
	}

}
