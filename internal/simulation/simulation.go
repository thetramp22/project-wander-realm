package simulation

import "github.com/thetrame22/project-wander-realm/internal/world"

type Simulation struct {
	world *world.World
}

func New() (*Simulation, error) {
	wrld, err := world.New()
	if err != nil {
		return nil, err
	}

	sim := Simulation{
		world: wrld,
	}

	return &sim, nil
}

func (s *Simulation) Start() error {

	return nil
}

func (s *Simulation) Stop() error {

	return nil
}
