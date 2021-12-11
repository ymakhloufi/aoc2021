package vessels

import "fmt"

var _ Vessel = &RegularVessel{}

type RegularVessel struct {
	Pos Position
}

func (vessel *RegularVessel) GetPosition() Position {
	return vessel.Pos
}

func (vessel *RegularVessel) Up(steps int) error {
	return vessel.Down(-steps)
}

func (vessel *RegularVessel) Forward(steps int) error {
	newX := vessel.Pos.X + steps
	if newX < 0 {
		return fmt.Errorf("horizontal position turned negative")
	}

	vessel.Pos.X = newX
	return nil
}

func (vessel *RegularVessel) Down(steps int) error {
	newY := vessel.Pos.Y + steps
	if newY < 0 {
		return fmt.Errorf("depth turned negative")
	}

	vessel.Pos.Y = newY
	return nil
}
