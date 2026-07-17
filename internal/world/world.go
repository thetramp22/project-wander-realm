package world

import "log"

type World struct {
	tickCount int
}

func New() (*World, error) {
	world := World{
		tickCount: 0,
	}

	return &world, nil
}

func (w *World) Update() error {
	w.tickCount++

	log.Printf("Updating world... tick %d", w.tickCount)

	return nil
}
