package main

import (
	"fmt"
	"io"
	"os"
)

const N = 20

func NewGrid3D(size int) [][][]int {
	grid := make([][][]int, size)
	for i := 0; i < size; i++ {
		grid[i] = make([][]int, size)
		for j := 0; j < size; j++ {
			grid[i][j] = make([]int, size)
		}
	}
	return grid
}

func scanGrid(r io.Reader) [][][]int {
	grid := NewGrid3D(N)
	var x, y, z int
	_, err := fmt.Fscanf(r, "%d,%d,%d", &x, &y, &z)
	for err == nil {
		grid[x][y][z] = 1
		_, err = fmt.Fscanf(r, "%d,%d,%d", &x, &y, &z)
	}
	return grid
}

func surfaceArea(grid [][][]int, blockedCubes map[string]bool) int {
	seen := make(map[string]bool)
	allSides := 0
	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			for z := 0; z < N; z++ {
				if grid[x][y][z] == 1 {
					id := fmt.Sprintf("%d-%d-%d", x, y, z)
					if !seen[id] {
						visit := [][]int{{x, y, z}}
						for len(visit) > 0 {
							cube := visit[0]
							visit = visit[1:]
							cid := fmt.Sprintf("%d-%d-%d", cube[0], cube[1], cube[2])
							seen[cid] = true
							sides := 6
							for i := 0; i <= 2; i++ {
								for j := -1; j <= 1; j += 2 {
									conn := make([]int, 0)
									conn = append(conn, cube...)
									conn[i] += j
									if !(conn[0] < 0 || conn[1] < 0 || conn[2] < 0 || conn[0] >= N || conn[1] >= N || conn[2] >= N) {
										connId := fmt.Sprintf("%d-%d-%d", conn[0], conn[1], conn[2])
										if seen[connId] || blockedCubes[connId] {
											sides--
											continue
										}
										if grid[conn[0]][conn[1]][conn[2]] == 1 {
											sides--
											visit = append(visit, conn)
											seen[connId] = true
										}
									}
								}
							}
							allSides += sides
						}
					}
				}
			}
		}
	}
	return allSides
}

func blocked(grid [][][]int) map[string]bool {
	seen := make(map[string]bool)
	cubes := map[string]bool{}
	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			for z := 0; z < N; z++ {
				if grid[x][y][z] == 0 {
					id := fmt.Sprintf("%d-%d-%d", x, y, z)
					if !seen[id] {
						visit := [][]int{{x, y, z}}
						visited := map[string]bool{}
						isOpen := false
						for len(visit) > 0 {
							cube := visit[0]
							visit = visit[1:]
							cid := fmt.Sprintf("%d-%d-%d", cube[0], cube[1], cube[2])
							seen[cid] = true
							visited[cid] = true
							for i := 0; i <= 2; i++ {
								for j := -1; j <= 1; j += 2 {
									conn := append([]int{}, cube...)
									conn[i] += j
									if conn[0] < 0 || conn[1] < 0 || conn[2] < 0 || conn[0] >= N || conn[1] >= N || conn[2] >= N {
										isOpen = true
										continue
									}
									connId := fmt.Sprintf("%d-%d-%d", conn[0], conn[1], conn[2])
									if !seen[connId] && grid[conn[0]][conn[1]][conn[2]] == 0 {
										visit = append(visit, conn)
										seen[connId] = true
									}
								}
							}
						}
						if !isOpen {
							for c := range visited {
								cubes[c] = true
							}
						}
					}
				}
			}
		}
	}
	return cubes
}

func main() {
	r, _ := os.Open("./day18/data/input.txt")
	defer r.Close()

	grid := scanGrid(r)
	fmt.Println(surfaceArea(grid, map[string]bool{})) // part 1
	blockedCubes := blocked(grid)
	fmt.Println(surfaceArea(grid, blockedCubes)) // part 2
}
