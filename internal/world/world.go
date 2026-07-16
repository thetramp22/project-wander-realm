package world

type World struct{}

func New() (*World, error) {
	world := World{}

	return &world, nil
}

func (w *World) Update() error {

	return nil
}
