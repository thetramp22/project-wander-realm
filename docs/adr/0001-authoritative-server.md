# Server Authority

### Decision:

The server is the authoritative source of all game state.

### Alternatives considered:

- Client authoritative
- Hybrid

### Reasons:

- Prevent cheating
- Single source of truth
- Simplified synchronization

### Consequences:

- Increased server responsibility
- Lower trust in clients