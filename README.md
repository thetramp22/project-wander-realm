# WanderRealm

WanderRealm is an experimental multiplayer RPG server written in Go. The project is being built as both a learning exercise and a portfolio piece, with a focus on software architecture, simulation systems, and backend engineering rather than producing a complete game as quickly as possible.

The long-term vision is to build a persistent online world where an authoritative server manages the simulation, game state, and player interactions while lightweight clients communicate through well-defined intents.

Although inspired by classic MMORPGs, the primary goal of WanderRealm is to explore the engineering challenges behind persistent online worlds and to document the design decisions made along the way.

---

## Project Goals

This project is intended to demonstrate experience with:

* Go application architecture
* Long-running server processes
* Simulation and game loops
* Networking and client/server communication
* Concurrency
* State management
* Persistence
* Clean package design
* Incremental software development

Rather than attempting to build a full-featured MMORPG immediately, the project is developed in small, well-defined milestones that each produce a working application.

---

## Current Status

**Current Milestone:** Sprint 1 – Core Server Architecture

Implemented:

* Project structure
* Server package
* Simulation package
* World package
* Constructor-based dependency wiring
* Initial architectural documentation

In Progress:

* Simulation lifecycle
* Game loop
* World updates

Planned:

* Entity system
* Networking
* Player connections
* Persistence
* Gameplay systems

---

## Architecture Overview

The current architecture intentionally separates responsibilities into three core systems:

```text
Server
    ↓
Simulation
    ↓
World
```

### Server

The application's top-level coordinator.

Responsible for:

* Application lifecycle
* Initializing core systems
* Starting and stopping the simulation

### Simulation

Owns the passage of time.

Responsible for:

* Running the simulation loop
* Advancing the world
* Coordinating updates

### World

Represents the authoritative state of the game world.

Responsible for:

* Maintaining world state
* Updating entities
* Responding to simulation updates

This separation allows each system to evolve independently while maintaining clear ownership boundaries.

---

## Guiding Principles

WanderRealm follows several architectural principles throughout development.

### Server Authority

Clients communicate intentions.

The server validates those intentions, updates the simulation, and distributes the resulting state.

### Separation of Responsibilities

Each package should have a single, well-defined responsibility.

### Build the Simplest Thing

Features are added only when they solve a current problem.

Premature abstractions are intentionally avoided.

### Incremental Development

Every milestone should leave the project in a runnable state.

---

## Project Structure

```text
cmd/
    server/

internal/
    server/
    simulation/
    world/

docs/
```

Additional documentation can be found in the `docs/` directory.

---

## Documentation

The project documentation is maintained alongside the source code.

Current documents include:

* `architecture-vision.md` — Long-term architectural goals and system vision

Additional documentation, including an engineering journal, glossary, architectural decision records (ADRs), and roadmap, will be added as the project evolves.

---

## Running the Project

Requirements:

* Go 1.25+

Clone the repository:

```bash
git clone https://github.com/thetramp22/project-wander-realm.git
cd project-wander-realm
```

Run the server:

```bash
go run ./cmd/server
```

> **Note:** The project is in active development and does not yet implement gameplay or networking.

---

## Why This Project Exists

WanderRealm is intentionally being built as an engineering-first project.

Rather than focusing solely on creating a playable game, the project emphasizes thoughtful architecture, clear documentation, and incremental design. The repository is intended to showcase not only the finished software but also the reasoning behind its evolution.
