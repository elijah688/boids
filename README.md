

# Flocking Simulation with Boids Algorithm

This program simulates a flock of birds using the "Boids" algorithm. The flock is represented by an array of `Boid` structs stored in the `boids` variable. Each `Boid` is updated and moved by calling its `Update` function, which is not shown in the code snippet.

The positions of the `Boid` objects are stored in a 2D array called `boidMap`, and a read-write lock called `rWLock` is used to ensure safe access to the `boids` and `boidMap` variables in concurrent access scenarios.

## Dependencies

The program uses the [ebiten](https://github.com/hajimehoshi/ebiten) library to create a window and update and draw the flock in the window. Make sure that it is installed before running the program.

## Usage

To run the program, you can use the `go run` command:

```sh
go run main.go
```

Alternatively, you can use the included Makefile to run the program and run the tests. To run the program using the Makefile, use the `run` target:

```sh
make run
```

This will compile and run the program. To run the tests using the Makefile, use the `test` target:

```sh
make test
```

This will run the tests for the program using the `-race`, `-v`, and `-failfast` flags.

## Concurrency

The `boid` package uses goroutines to make the boids move concurrently. Specifically, each `Boid` is represented by a goroutine which continuously moves the `Boid` based on its velocity and acceleration vectors. The `start` method of the `Boid` struct is called as a goroutine, which calls the `moveOne` method of the `Boid` struct in a loop. The `moveOne` method calculates the acceleration vector for the current `Boid` by calling the `calcAcceleration` method, and then updates the `Boid`'s position and velocity vectors based on the acceleration vector.

In addition to the concurrent movement of individual boids, the `calcAcceleration` method also uses `RWMutex` to provide concurrent access to the shared `boidsMap` and `boids` data structures. Specifically, it uses a read lock to access the `boidsMap` data structure to find the other boids in the neighborhood of the current `Boid`, and then uses a write lock to update the `boidsMap` data structure after the `Boid`'s position is updated.
