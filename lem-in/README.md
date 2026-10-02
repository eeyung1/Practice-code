# lem-in

A Go ant-farm solver using only the standard library. It reads a colony,
selects routes, assigns ants, and prints legal turn-by-turn moves.

## Run

```bash
cd lem-in
go run . example.txt
# Or build the executable:
go build -o lem-in .
./lem-in example.txt
```

Give exactly one filename. Successful output contains the original input,
a blank separator, and one line per turn, such as `L1-B L2-C`. Invalid maps
and unreadable files return a nonzero exit status and an error on stderr.

## Input

The first non-comment, nonblank line is a positive ant count. Rooms follow
as `name x y`, then tunnels as `name1-name2`. `##start` and `##end` each
mark the next room; intervening comments are allowed. Unknown commands
and comments are ignored. Blank lines and CRLF input are supported.

Room names cannot start with `L` or `#`, contain whitespace, or contain
`-` (the tunnel delimiter). Coordinates must be integers. Duplicate names,
coordinates, undirected tunnels, self-links, unknown endpoints, malformed
lines, repeated markers, rooms after tunnels, and absent start/end rooms
are rejected. A map without a route is rejected by the solver.

## Implementation

- `parser.go`: input validation and original-input preservation.
- `graph.go`: deterministic undirected adjacency lists.
- `bfs.go`: the original BFS remains available for learning.
- `paths.go`: room splitting and minimum-cost residual flow. Reverse edges
  can undo earlier choices. The solver evaluates feasible flow sizes and
  chooses the path set with the smallest allocated finishing time.
- `allocation.go`: assign each ant to the lowest projected finishing time,
  then build its `Ant` record. Path length means number of tunnels.
- `simulation.go`: move leading ants first; enforce intermediate room
  occupancy, tunnel capacity, and one move per ant per turn.
- `main.go`: filename argument, error reporting, and required output.

`FindPaths` keeps the earlier API but now returns the maximum available
set of internally disjoint routes. `BestPaths` selects routes for the
actual ant count. `Solve` connects selection, allocation, and simulation.
The direct start/end tunnel admits one ant per turn, matching the tunnel
capacity rule.

## Verification

```bash
go test -race -count=1 -timeout=120s ./...
go vet ./...
```

Tests include parser errors, original output, direct connections,
disconnected maps, balanced allocation, and movement legality. A seeded
small-graph test compares results with exhaustive internally disjoint path
sets for 150 maps and ant counts 1–6. GitHub Actions runs these checks,
formatting, a build, and CLI smoke tests.

The allocator costs O(ants × paths). Residual path search uses Bellman-Ford;
each augmentation costs O(vertices × edges), so very large colonies can
require more time than small teaching examples. All turn output is held
in memory before printing. The optional graphical visualizer is not part
of this core CLI implementation.
