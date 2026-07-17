# Simulation Owns Time

### Decision:

The Simulation advances time and instructs the World to update.

### Alternatives:

- World owns time.

### Reasons:

- Separates state from execution.
- Easier testing.
- Clearer responsibilities.