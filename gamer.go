package main

type Player struct {
	currPos   Vec2
	targetPos Vec2
	speed     float64
}

func (p *Player) Move(v Vec2) {
	p.targetPos = v
}

func (p *Player) Pos() Vec2 {
	return p.currPos
}

func (p *Player) IsArrived() bool {
	return p.currPos.Distance(p.targetPos) < p.speed
}

func (p *Player) Update() {
	if !p.IsArrived() {
		dir := p.targetPos.Sub(p.currPos).Normalize()
		newPos := p.currPos.Add(dir.Scale(p.speed))
		p.currPos = newPos
	}
}

func NewPlayer(speed float64) *Player {
	return &Player{
		speed: speed,
	}
}
