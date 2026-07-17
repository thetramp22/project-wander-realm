package server

import (
	"log"

	"github.com/thetrame22/project-wander-realm/internal/simulation"
)

type Server struct {
	simulation *simulation.Simulation
	running    bool
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
	log.Println("starting server...")
	s.running = true

	err := s.simulation.Start()
	if err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	log.Print("Stopping sever...")
	err := s.simulation.Stop()
	if err != nil {
		return err
	}
	s.running = false

	return nil
}
