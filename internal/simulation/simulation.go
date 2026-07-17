package simulation

import (
	"log"
	"time"

	"github.com/thetrame22/project-wander-realm/internal/world"
)

type Simulation struct {
	world   *world.World
	running bool
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
	log.Println("starting simulation...")
	s.running = true

	for s.running {
		err := s.world.Update()
		if err != nil {
			return err
		}

		time.Sleep(2 * time.Second)
	}

	return nil
}

func (s *Simulation) Stop() error {
	log.Print("Stopping simulation...")
	s.running = false

	return nil
}
