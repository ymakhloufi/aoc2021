package vessels

import "fmt"

var _ Vessel = &AimedVessel{}

type AimedVessel struct {
	RegularVessel
	aim int
}

func (vessel *AimedVessel) Forward(steps int) error {
	if err := vessel.RegularVessel.Forward(steps); err != nil {
		return err
	}

	newY := vessel.Pos.Y + steps*vessel.aim
	if newY < 0 {
		return fmt.Errorf("depth turned negative")
	}

	vessel.Pos.Y = newY
	return nil
}

func (vessel *AimedVessel) Down(steps int) error {
	vessel.aim += steps
	return nil
}

func (vessel *AimedVessel) Up(steps int) error {
	return vessel.Down(-steps)
}
