# Engineering Journal

## Sprint 0

### Discovery

We established the following architectural principles:

- Server owns the simulation.
- Simulation owns time.
- World owns state.

Questions Raised

- Should dropped items be entities?
- How should time persist across shutdown?

Lessons Learned

Separating definitions from instances simplifies nearly every game system.

## Sprint 1

Implemented

- Server
- Simulation
- World

Lessons

Designing public APIs before implementations resulted in cleaner package boundaries.