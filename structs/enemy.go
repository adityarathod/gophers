package main

type enemy struct {
	name string
	position
}

func makeEnemy(pos position) *enemy {
	return &enemy{"Enemy", pos}
}

func (e *enemy) getName() string {
	return e.name
}

func (e *enemy) getPosition() position {
	return e.position
}

func (e *enemy) setPosition(pos position) {
	e.position = pos
}

func (e *enemy) getDistance(p entity) float64 {
	return distance(e.position, p.getPosition())
}
