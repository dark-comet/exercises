package vector

import "math"

type Vec2 struct {
	X, Y float64
}

func NewVec2(x, y float64) Vec2 {
	return Vec2{x, y}
}

func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v.X + v2.X, v.Y + v2.Y}
}

func (v Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{v.X - v2.X, v.Y - v2.Y}
}

func (v Vec2) Mul(f float64) Vec2 {
	return Vec2{v.X * f, v.Y * f}
}

func (v Vec2) Len() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2) Normalize() Vec2 {
	length := v.Len()
	return Vec2{v.X / length, v.Y / length}
}

func (v Vec2) Dot(v2 Vec2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vec2) Cross(v2 Vec2) float64 {
	return v.X*v2.Y - v.Y*v2.X
}

func (v Vec2) Angle() float64 {
	return math.Atan2(v.Y, v.X)
}
