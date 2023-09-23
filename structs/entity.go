package main

type entity interface {
	getName() string
	getPosition() position
	setPosition(position)
	getDistance(entity) float64
}

const stepSize float64 = 0.2

func moveTowards(e entity, to position) {
	curPos := e.getPosition()
	curPos.x += (to.x - curPos.x) * stepSize
	curPos.y += (to.y - curPos.y) * stepSize
	curPos.z += (to.z - curPos.z) * stepSize
	e.setPosition(curPos)
}
