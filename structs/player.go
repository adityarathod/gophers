package main

type player struct {
	name string
	position
}

func makePlayer(pos position) *player {
	return &player{"Player", pos}
}
func (p *player) getName() string {
	return p.name
}

func (p *player) getPosition() position {
	return p.position
}

func (p *player) setPosition(pos position) {
	p.position = pos
}

func (p *player) getDistance(e entity) float64 {
	return distance(p.position, e.getPosition())
}
