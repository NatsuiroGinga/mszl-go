package main

import "math"

type Vec2 struct {
	X, Y float64
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		v.X - other.Y,
		v.Y - other.Y,
	}
}

func (v Vec2) Scale(s float64) Vec2 {
	return Vec2{
		v.X * s,
		v.Y * s,
	}
}

func (v Vec2) Distance(other Vec2) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func (v Vec2) Normalize() Vec2 {
	mag := v.X*v.X + v.Y*v.Y
	if mag > 0 {
		otherOverMag := 1 / math.Sqrt(mag)
		return Vec2{
			v.X * otherOverMag,
			v.Y * otherOverMag,
		}
	}
	return Vec2{}
}
