// This solution was either stolen, adapted or translated from another language by me.
// I keep track of tasks which I failed to solve and this is one of them.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Placement struct {
	rowMasks []uint64
	area     int
}

func rotate(shape [][]bool) [][]bool {
	height, width := len(shape), len(shape[0])
	out := make([][]bool, width)

	for i := range out {
		out[i] = make([]bool, height)
		for j := range out[i] {
			out[i][j] = shape[height-1-j][i]
		}
	}

	return out
}

func flip(shape [][]bool) [][]bool {
	height, width := len(shape), len(shape[0])
	out := make([][]bool, height)

	for i := range out {
		out[i] = make([]bool, width)
		for j := range out[i] {
			out[i][j] = shape[i][width-1-j]
		}
	}

	return out
}

func normalizeTrim(shape [][]bool) [][]bool {
	minI, minJ := 1<<30, 1<<30
	maxI, maxJ := -1, -1

	for i := range shape {
		for j := range shape[i] {
			if shape[i][j] {
				if i < minI {
					minI = i
				}
				if j < minJ {
					minJ = j
				}
				if i > maxI {
					maxI = i
				}
				if j > maxJ {
					maxJ = j
				}
			}
		}
	}

	if maxI < 0 || maxJ < 0 {
		return [][]bool{{}}
	}

	height := maxI - minI + 1
	width := maxJ - minJ + 1
	out := make([][]bool, height)
	for i := range out {
		out[i] = make([]bool, width)
	}

	for i := minI; i <= maxI; i++ {
		for j := minJ; j <= maxJ; j++ {
			if shape[i][j] {
				out[i-minI][j-minJ] = true
			}
		}
	}

	return out
}

func shapeKey(shape [][]bool) string {
	var b strings.Builder

	for _, r := range shape {
		for _, c := range r {
			if c {
				b.WriteByte('#')
			} else {
				b.WriteByte('.')
			}
		}
		b.WriteByte('\n')
	}

	return b.String()
}

func orientations(shape [][]bool) [][][]bool {
	var out [][][]bool
	seen := map[string]bool{}
	current := shape

	for range 4 {
		for _, v := range [][][]bool{current, flip(current)} {
			n := normalizeTrim(v)
			k := shapeKey(n)
			if !seen[k] {
				seen[k] = true
				out = append(out, n)
			}
		}
		current = rotate(current)
	}

	return out
}

func shapeArea(shape [][]bool) int {
	a := 0

	for i := range shape {
		for j := range shape[i] {
			if shape[i][j] {
				a++
			}
		}
	}

	return a
}

func genPlacementsForOrient(orient [][]bool, width, height int) []Placement {
	var cells [][2]int
	maxI, maxJ := 0, 0

	for i := range orient {
		for j := range orient[i] {
			if orient[i][j] {
				cells = append(cells, [2]int{i, j})
				if i > maxI {
					maxI = i
				}
				if j > maxJ {
					maxJ = j
				}
			}
		}
	}

	if len(cells) == 0 {
		return nil
	}

	var out []Placement
	for top := 0; top+maxI < height; top++ {
		for left := 0; left+maxJ < width; left++ {
			rowMasks := make([]uint64, height)
			for _, c := range cells {
				y := top + c[0]
				x := left + c[1]
				rowMasks[y] |= (uint64(1) << uint(x))
			}
			out = append(out, Placement{rowMasks: rowMasks, area: len(cells)})
		}
	}

	return out
}

func placementsForShape(shape [][]bool, width, height int) []Placement {
	var all []Placement

	for _, o := range orientations(shape) {
		ps := genPlacementsForOrient(o, width, height)
		all = append(all, ps...)
	}

	return all
}

func bitsCount64(x uint64) int {
	count := 0

	for x != 0 {
		x &= x - 1
		count++
	}

	return count
}

func countFreeCells(gridMasks []uint64, width int) int {
	height := len(gridMasks)
	fullMask := uint64(0)
	if width == 64 {
		fullMask = ^uint64(0)
	} else {
		fullMask = (uint64(1) << uint(width)) - 1
	}

	total := 0
	for i := 0; i < height; i++ {
		freeMask := ^gridMasks[i] & fullMask
		total += bitsCount64(freeMask)
	}

	return total
}

func solveRegionBit(width, height int, shapes [][][]bool, req []int) bool {
	if width > 64 {
		return false
	}

	P := len(shapes)
	shapePls := make([][]Placement, P)
	shapeAreas := make([]int, P)
	for i := 0; i < P; i++ {
		shapeAreas[i] = shapeArea(shapes[i])
		shapePls[i] = placementsForShape(shapes[i], width, height)
		if len(shapePls[i]) == 0 && req[i] > 0 {
			return false
		}
	}

	gridMasks := make([]uint64, height)
	totalPieces := 0
	remainingTotalArea := 0
	for i := 0; i < P; i++ {
		totalPieces += req[i]
		remainingTotalArea += req[i] * shapeAreas[i]
	}

	if totalPieces == 0 {
		return true
	}

	rem := make([]int, P)
	copy(rem, req)

	var dfs func([]uint64, []int, int, int) bool
	dfs = func(grid []uint64, remaining []int, placed int, remArea int) bool {
		if placed == totalPieces {
			return true
		}

		freeCells := countFreeCells(grid, width)
		if remArea > freeCells {
			return false
		}

		bestIdx := -1
		bestCount := 1 << 30
		feasibleLists := make([][]Placement, P)

		for i := 0; i < P; i++ {
			if remaining[i] == 0 {
				continue
			}

			ps := shapePls[i]
			var feas []Placement
			for _, pl := range ps {
				ok := true
				for r := 0; r < height; r++ {
					if (pl.rowMasks[r] & grid[r]) != 0 {
						ok = false
						break
					}
				}

				if ok {
					feas = append(feas, pl)
				}
			}
			feasibleLists[i] = feas
			if len(feas) == 0 {
				return false
			}

			if len(feas) < bestCount {
				bestCount = len(feas)
				bestIdx = i
			}
		}

		feas := feasibleLists[bestIdx]
		for _, pl := range feas {
			for r := 0; r < height; r++ {
				grid[r] |= pl.rowMasks[r]
			}

			remaining[bestIdx]--
			if dfs(grid, remaining, placed+1, remArea-shapeAreas[bestIdx]) {
				return true
			}

			remaining[bestIdx]++
			for r := 0; r < height; r++ {
				grid[r] &^= pl.rowMasks[r]
			}
		}

		return false
	}

	return dfs(gridMasks, rem, 0, remainingTotalArea)
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	parts := strings.Split(strings.Trim(string(data), "\n"), "\n\n")
	rawShapes := parts[:len(parts)-1]
	rawGrids := parts[len(parts)-1]

	var shapes [][][]bool
	for _, s := range rawShapes {
		lines := strings.Split(s, "\n")[1:]
		var sh [][]bool
		for _, l := range lines {
			var row []bool
			for _, c := range l {
				row = append(row, c == '#')
			}
			sh = append(sh, row)
		}
		shapes = append(shapes, sh)
	}

	answer := 0
	regionIdx := 0
	for _, l := range strings.Split(rawGrids, "\n") {
		if strings.TrimSpace(l) == "" {
			continue
		}
		p := strings.Split(l, " ")
		size := strings.Split(p[0][:len(p[0])-1], "x")
		W, _ := strconv.Atoi(size[0])
		H, _ := strconv.Atoi(size[1])

		req := make([]int, len(shapes))
		for i := 1; i < len(p); i++ {
			v, _ := strconv.Atoi(p[i])
			if i-1 < len(req) {
				req[i-1] = v
			}
		}

		ok := solveRegionBit(W, H, shapes, req)
		if ok {
			answer++
		}
		regionIdx++
	}

	fmt.Println(answer)
}
