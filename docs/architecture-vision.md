# Architecture Vision

## Project Overview

**Working Title:** *Project WanderRealm*

This project is a learning-focused multiplayer RPG backend written in Go with a Godot client.

The primary objective is not to create a commercially viable MMORPG, but to explore the architecture of persistent online worlds while strengthening backend engineering skills. The project emphasizes clean architecture, maintainability, and iterative learning over feature completeness or scale.

The game will draw inspiration from classic MMORPGs such as RuneScape and EverQuest, focusing on persistent progression, skill-based advancement, and a living shared world.

---

# Project Goals

The project should help answer questions such as:

* How does a multiplayer server maintain a persistent world?
* How are long-lived client connections managed?
* How can concurrent systems safely interact?
* How should game rules be represented and simulated?
* How can persistent player progression be modeled?
* How should a backend evolve as complexity increases?

Success is measured by the quality of the architecture and the lessons learned, not by the number of implemented features.

---

# Guiding Principles

## The Server Owns the Truth

The server is the authoritative source of all game state.

Clients express intentions.

The server validates those intentions, advances the simulation, and communicates the results.

Clients never decide game outcomes.

---

## The World Is a Simulation

The World is the authoritative simulation of the game.

It contains entities, tracks global state, advances time, and provides the environment in which gameplay occurs.

The world exists independently of any individual player.

---

## Definitions Are Separate from Runtime Instances

Static game definitions describe what things are.

Runtime instances represent things currently existing within the simulation.

Examples include:

* Goblin definition vs. Goblin #1837
* Bronze Sword definition vs. Scott's Bronze Sword
* Cooking skill definition vs. Scott's Cooking experience

This separation reduces duplication and simplifies future expansion.

---

## Store Primary Data

Persistent systems should store primary data and derive secondary values whenever practical.

Examples:

* Store experience.
* Derive levels.
* Store item definitions.
* Derive statistics from equipped items.

This minimizes duplicated information and reduces synchronization problems.

---

## Compose Systems Rather Than Hardcode Them

Complex behavior should emerge from combining simple reusable systems.

Examples include:

* Monster behaviors
* Item effects
* Status effects
* Character abilities

The project favors composition over rigid inheritance whenever practical.

---

## Build Vertical Slices

Development will proceed by implementing complete, playable slices rather than isolated systems.

Each milestone should leave the game in a working state.

Examples include:

* Connect to the server.
* Create a player.
* Move around.
* Disconnect.

Later milestones build upon these foundations.

---

# Core Concepts

## World

The authoritative simulation.

Responsibilities include:

* Managing entities
* Advancing time
* Tracking weather
* Managing global events
* Providing the environment in which gameplay occurs

---

## Entity

An object that independently exists within the running simulation.

Examples may include:

* Players
* Monsters
* NPCs
* Dropped items
* Interactive world objects

Not every object in the game is necessarily an entity.

---

## Player

A player is a controllable entity connected to a client.

A player has:

* Identity
* Connection
* Skills
* Inventory
* Equipment
* Resources
* Status effects
* State

Persistent progression belongs to the player.

---

## Monster

A monster is a non-player entity driven by behaviors.

A monster has:

* Identity
* Resources
* Status effects
* State
* Behaviors
* Possible rewards

Behavior is composed from reusable modules rather than hardcoded per monster.

---

## Skill

A skill defines a progression path.

A skill specifies:

* Experience curve
* Unlocks
* Level thresholds
* Stat progression

Player-specific progress is stored separately from the skill definition.

---

## Item

An item defines an object that may exist within the game.

An item definition includes:

* Identity
* Name
* Description
* Statistics
* Effects

Individual item instances may exist in inventories, containers, equipment, or the world.

---

## Container

A container holds item instances.

Examples include:

* Player inventory
* Bank
* Chest
* Ground storage
* Other future storage systems

Containers provide a common model for ownership and transfer of items.

---

# Communication Philosophy

Clients communicate intentions rather than outcomes.

Example:

Client:

"I would like to attack Goblin #42."

Server:

* Validate request
* Advance simulation
* Determine outcome
* Notify interested clients

Most gameplay interactions follow this pattern:

Intent

↓

Validation

↓

Simulation

↓

Result

↓

Broadcast

---

# Version 0.1 Goals

The initial milestone focuses on proving the architecture rather than implementing gameplay.

Version 0.1 should demonstrate:

* Server startup
* World creation
* Client connection
* Player creation
* Player entering the world
* Player disconnection
* Clean project architecture

No combat, persistence, inventories, or progression are required for Version 0.1.

The objective is simply to establish the smallest functioning multiplayer simulation upon which future systems can be built.
