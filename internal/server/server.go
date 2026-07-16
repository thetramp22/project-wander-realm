package server

import "github.com/thetrame22/project-wander-realm/internal/simulation"

type Server struct {
	simulation *simulation.Simulation
}

func New() (*Server, error) {
	sim, err := simulation.New()
	if err != nil {
		return nil, err
	}

	srv := Server{
		simulation: sim,
	}

	return &srv, nil
}

func (s *Server) Start() error {

	return nil
}

func (s *Server) Stop() error {

	return nil
}
