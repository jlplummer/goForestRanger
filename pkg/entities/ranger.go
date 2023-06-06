package entities

type Ranger struct {
	X, Y, Health int
}

func (r *Ranger) Init(x, y int) {
	r.Health = 5
	r.X = x
	r.Y = y
}

func (r *Ranger) Draw() {}

func (r *Ranger) Update(x, y int) {

}
