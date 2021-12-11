package vessels

type Position struct {
	X int
	Y int
}

type Vessel interface {
	GetPosition() Position
	Forward(int) error
	Up(int) error
	Down(int) error
}
