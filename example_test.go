package disjoint_test

import (
	"fmt"
	"github.com/spakin/disjoint/v2"
	"math/rand"
)

// Draw a maze.  This is my favorite use of disjoint-set forests.  The
// algorithm works by repeatedly knocking down walls between two sets
// of rooms that are not part of the same union and merging those
// rooms into a single union.  A union represents connected
// components—rooms in the maze that are mutually reachable.  A single
// union implies that every room is reachable from every other room.
//
// This is a fairly long example, but note that only half of it relates to
// generating the maze.  The other half renders the maze using Unicode
// box-drawing characters.
func Example_maze() {
	const width = 50  // Width of maze in rooms (must be > 1)
	const height = 10 // Height of maze in rooms (must be > 1)

	// A room is identified by its walls and by the other rooms it can reach.
	type Room struct {
		N       bool              // North side of room is a wall
		S       bool              // South side of room is a wall
		E       bool              // East side of room is a wall
		W       bool              // West side of room is a wall
		Reaches *disjoint.Element // Element in a set of reachable rooms
	}

	// Initialize the maze data structure.
	maze := make([][]Room, height)
	for y := range maze {
		maze[y] = make([]Room, width)
		for x := range maze[y] {
			// Start with all walls present and no other rooms reachable.
			maze[y][x].N = true
			maze[y][x].S = true
			maze[y][x].E = true
			maze[y][x].W = true
			maze[y][x].Reaches = disjoint.NewElement()
		}
	}

	// Repeatedly remove walls until a single connected component remains.
	rand.Seed(5552368)
	for cc := width * height; cc > 1; {
		// Because of symmetry, we need only connect to the right or
		// down rather than in all four directions.
		x0 := rand.Intn(width)
		y0 := rand.Intn(height)
		x1 := x0
		y1 := y0
		dir := rand.Intn(2)
		if dir == 0 && x0 < width-1 {
			x1++ // Go right.
		} else if dir == 1 && y0 < height-1 {
			y1++ // Go down.
		} else {
			continue // Can't go in the desired direction
		}
		if maze[y0][x0].Reaches.Find() == maze[y1][x1].Reaches.Find() {
			continue // Already connected
		}

		// Tear down the wall.
		if dir == 0 {
			// Right/left
			maze[y0][x0].E = false
			maze[y1][x1].W = false
		} else {
			// Down/up
			maze[y0][x0].S = false
			maze[y1][x1].N = false
		}
		maze[y0][x0].Reaches.Union(maze[y1][x1].Reaches)
		cc--
	}

	// Punch holes for an entry (UL) and exit (LR).
	maze[0][0].W = false
	maze[height-1][width-1].E = false

	// Convert the grid of rooms to a grid of walls.  Walls are staggered
	// spatially by half a room vertically and horizontally.
	type Walls struct {
		N bool // Northbound wall from cell center
		S bool // Southbound wall from cell center
		E bool // Eastbound wall from cell center
		W bool // Westbound wall from cell center
	}
	walls := make([][]Walls, height+1)
	for y := range walls {
		walls[y] = make([]Walls, width+1)
	}
	for y := 0; y < height+1; y++ {
		for x := 0; x < width+1; x++ {
			if y < height {
				if x < width {
					walls[y][x].E = maze[y][x].N
					walls[y][x].S = maze[y][x].W
				}
				if x > 0 {
					walls[y][x].W = maze[y][x-1].N
					walls[y][x].S = maze[y][x-1].E
				}
			}
			if y > 0 {
				if x > 0 {
					walls[y][x].W = maze[y-1][x-1].S
					walls[y][x].N = maze[y-1][x-1].E
				}
				if x < width {
					walls[y][x].E = maze[y-1][x].S
					walls[y][x].N = maze[y-1][x].W
				}
			}
		}
	}

	// Define a map from wall types to Unicode box-drawing characters.
	wallsToGlyph := []rune{' ', '╴', '╶', '─', '╷', '┐', '┌', '┬', '╵', '┘', '└', '┴', '│', '┤', '├', '┼'}

	// Define a map from Booleans to integers.
	boolToInt := map[bool]int{false: 0, true: 1}

	// Output the glyph corresponding to each cell of walls.
	for _, row := range walls {
		for _, cell := range row {
			val := boolToInt[cell.N]<<3 | boolToInt[cell.S]<<2 | boolToInt[cell.E]<<1 | boolToInt[cell.W]
			fmt.Printf("%c", wallsToGlyph[val])
		}
		fmt.Println("")
	}

	// Output:
	// ╶───┬─────┬┬──┬────┬────┬─┬──┬┬─┬───┬─┬─┬───┬───┬─┐
	// ╷╷╶┐├┐┌┐╷╶┤│╶─┤╷╶┬─┘╶┐╷╷│╶┴─┐╵╵╶┴─╴╷│╷└╴├┬─╴│╷╶┬┴╴│
	// ├┴╴├┤╵│└┘┌┘╵╶┬┼┘╶┴┬╴╷│├┤├─╴┌┴┐╶┐╶──┤└┤╶┬┘├╴┌┴┤╷└╴┌┤
	// ├╴╷│└┬┼┐┌┤╷╶┬┘│╷╷╷╵╶┼┴┘╵├╴╶┼╴╵╷├─╴┌┘╶┤╷│┌┘╷│┌┤│╶┐╵│
	// │┌┘╵╶┤╵╵│││┌┴╴│││└┐╷├╴╷╶┴╴┌┼─╴│└┐╶┼─┐││╵├╴│╵╵└┘╶┼┬┤
	// ├┼┬╴╷└─┐╵│├┘╷╶┴┴┤╷││├─┴╴┌┐│╵╶┬┴┐├─┤┌┘└┘╶┤┌┘╶┐┌─╴╵╵│
	// │╵│╶┴┐╷╵╷╵╵┌┴┐╶┐╵└┼┼┤┌┐╷│││╶─┼╴└┘┌┘├┬┐┌┐└┘╶┐├┘╶┬┐╶┤
	// ├╴├┐┌┴┤╷├╴╶┴┐╵╶┤╷┌┘╵├┘╵├┤│╵╷┌┼┐┌┐├┐╵│└┤╵╷┌╴└┴┐╶┤│╷│
	// │╷╵╵│╶┘├┼╴┌─┴──┼┼┤╷╷├╴╶┘││╶┤╵│╵╵││╵╶┘╷│╷├┴┐╶┬┼╴│└┴┤
	// ├┴╴┌┴╴╷│╵╷╵╷╷┌╴│╵╵│└┘╶┐╶┘│╶┴─┘╷╷├┴╴╷╶┤╵└┤╶┴╴╵└┐╵╶┐╵
	// └──┴──┴┴─┴─┴┴┴─┴──┴───┴──┴────┴┴┴──┴─┴──┴─────┴──┴╴
}
