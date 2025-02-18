package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"time"
)

const (
	WALL   string = "#"
	END    string = "v"
	BLOCK  string = " "
	CURSOR string = "x"
)

var totalLines int

type Node struct {
	x     int
	y     int
	value string
}

type Walker struct {
	x       int
	y       int
	queue   []*Node
	visited []*Node
}

func pop(list *[]*Node) *Node {
	size := len((*list)) - 1
	removed := (*list)[size]
	(*list) = (*list)[0:size]
	return removed
}

func findDirections(maze [][]*Node, current *Node) []*Node {
	return []*Node{
		maze[current.y-1][current.x],
		maze[current.y][current.x+1],
		maze[current.y+1][current.x],
		maze[current.y][current.x-1],
	}

}

func printMaze(maze *[][]*Node) {

	fmt.Print("\033[H\033[J")

	for _, v := range *maze {
		for _, k := range v {
			fmt.Print(k.value)
		}
	}
	time.Sleep(50 * time.Millisecond)
}

func findShortest(parents map[*Node]*Node, end *Node, maze *[][]*Node) {

	for i := end; i != nil; i = parents[i] {
		(*maze)[i.y][i.x].value = "*"
	}
}

func (w *Walker) visit(maze [][]*Node) {
	parents := make(map[*Node]*Node)
	parents[w.queue[0]] = nil
	for len(w.queue) > 0 {
		printMaze(&maze)

		removed := pop(&w.queue)

		maze[removed.y][removed.x].value = "."

		w.visited = append(w.visited, removed)

		directions := findDirections(maze, removed)

		for _, v := range directions {

			if v.value == "v" {
				printMaze(&maze)
				findShortest(parents, removed, &maze)
				return
			}

			if v.value == " " && slices.Index(w.visited, v) < 0 {
				parents[v] = removed
				// DFS
				// w.queue = append(w.queue, v)

				// BFS
				w.queue = append([]*Node{v}, w.queue...)
			}
		}

	}
}

func main() {
	// load maze

	maze := [][]*Node{}
	var walker Walker
	file, err := os.ReadFile("maze.txt")

	if err != nil {
		log.Fatalln("Erron on reading the file", err)
	}

	x := 0
	y := 0
	line := []*Node{}
	for _, v := range file {
		sprite := string(v)

		if sprite == "\n" {
			y++
			x = 0
			maze = append(maze, line)
			line = []*Node{}
		}

		n := Node{x: x, y: y, value: sprite}

		if sprite == CURSOR {
			walker = Walker{x: x, y: y}
			walker.queue = append(walker.queue, &n)
		}

		line = append(line, &n)
		x++
	}
	totalLines = y

	maze = append(maze, line)
	walker.visit(maze)
	printMaze(&maze)

}
