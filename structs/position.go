package main

import "math"

type position struct {
	x, y, z float64
}

func makePosition(x, y, z float64) position {
	return position{x, y, z}
}

func distance(p1, p2 position) float64 {
	return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2) + math.Pow(float64(p1.z-p2.z), 2))
}

func (p *position) setCoordinates(x, y, z float64) {
	p.x = x
	p.y = y
	p.z = z
}
