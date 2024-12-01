package aoc2023

const STONE = 'O'
const GROUND = '.'
const OBSTACLE = '#'

func Day14part1() int {
	bytes := scanBytes("./aoc2023/input-day14.txt")
	total := 0

	var ground, obstacles []int
	for j := 0; j < len(bytes[0]); j++ {
		ground = []int{}
		obstacles = []int{}
		for i := 0; i < len(bytes); i++ {
			switch bytes[i][j] {
			case GROUND:
				ground = append(ground, i)
			case OBSTACLE:
				obstacles = append(obstacles, i)
			case STONE:
				//1. if no empty ground is available: do nothing and go in the next tile
				if len(ground) == 0 {
					//add to total
					total += len(bytes) - i
					continue
				}
				//2. if ground is available and there are no obstacles:
				//swap the first empty ground with the stone
				//cut the first ground from the slice
				//(the stone just rolls in the empty space)
				if len(ground) > 0 && len(obstacles) == 0 {
					//swap
					//ground -> stone
					bytes[ground[0]][j] = STONE
					//stone -> ground
					bytes[i][j] = GROUND
					//add to total
					total += len(bytes) - ground[0]
					//cut the first ground
					ground = ground[1:]
					//add space cleared by stone to ground
					ground = append(ground, i)
					continue
				}
				//3. if there is empty ground and obstacles:
				//find the first ground that is bigger than the last obstacle
				//swap the stone with this ground
				//remove this ground from the slice
				//(the stone rolls to the first emtpy space before obstacle or another stone)
				lastObstacleIndex := obstacles[len(obstacles)-1]
				swapped := false
				for index, groundIndex := range ground {
					if groundIndex > lastObstacleIndex && groundIndex < i {
						//swap
						swapped = true
						//ground -> stone
						bytes[groundIndex][j] = STONE
						//stone -> ground
						bytes[i][j] = GROUND
						//add to total
						total += len(bytes) - groundIndex
						//cut the ground
						if index == len(ground)-1 {
							ground = ground[:index]
						} else if index == 0 {
							ground = ground[1:]
						} else {
							ground = append(ground[:index], ground[index+1:]...)
						}
						//add space cleared by stone to ground
						ground = append(ground, i)
						break
					}
				}
				if !swapped {
					//no swaping so just add the stone to total
					total += len(bytes) - i
				}
			}
		}
	}

	return total
}

// https://en.wikipedia.org/wiki/Cycle_detection
func Day14part2() int {

	state := scanBytes("./aoc2023/input-day14.txt")
	tortoise := deepCopy(state)
	hare := deepCopy(tortoise)

	var t, h int = 0, 0

	for c := 1; c <= 1000000000; c++ {

		//hare moves 2 times faster
		t++
		h += 2

		tortoise = tiltCycle(tortoise)
		hare = tiltCycle(tiltCycle(hare))

		//compare with last state and break if not the same state
		if compare(tortoise, hare) {
			break
		}

	}

	//cycle length
	cl := h - 1 - t

	//fmt.Println(t, h, cl)

	//reset tortoise ???
	//yields the same result as the index of the iteration of first meeting above
	//means that the cycle starts at the same plase where it is first detected (looks like oddly specific coincidence ???)
	t = 0
	tortoise = deepCopy(state)
	for c := 1; c <= 1000000000; c++ {

		//same speed
		t++
		h++

		tortoise = tiltCycle(tortoise)
		hare = tiltCycle(tiltCycle(hare))

		//compare with last state and break if not the same state
		if compare(tortoise, hare) {
			break
		}
	}

	//fmt.Println(t, h)

	//the state iteration we need (which should be the last after the billionth opperation ??)
	i := t + cl // ???? TODO: this only true because of the input shape I guess??? calculate it properly

	//tot := (1000000000 - t) / cl
	//mod := (1000000000 - t) % cl

	//iterate the initial state to the desired state
	for c := 1; c <= i; c++ {
		state = tiltCycle(state)
		//fmt.Printf("%d -> %d\n", c, calcWeight(state))
	}

	//calculate weight and this is the end result!
	return calcWeight(state)
}

func calcWeight(state [][]byte) int {
	total := 0
	for i, l := range state {
		for _, b := range l {
			if b == STONE {
				total += len(state) - i
			}
		}
	}
	return total
}

func compare(arr1, arr2 [][]byte) bool {
	//we assume the 2D arrays has the same dimensions
	rows := len(arr1) - 1
	cols := len(arr1[0]) - 1
	for i := 0; i <= rows; i++ {
		for j := 0; j <= cols; j++ {
			if arr1[i][j] != arr2[i][j] {
				return false
			}
		}
	}
	return true
}

func tiltCycle(bytes [][]byte) [][]byte {
	//ROLL NORTH (UP)
	for j := 0; j < len(bytes[0]); j++ {
		ground := []int{}
		obstacles := []int{}
		for i := 0; i < len(bytes); i++ {
			switch bytes[i][j] {
			case GROUND:
				ground = append(ground, i)
			case OBSTACLE:
				obstacles = append(obstacles, i)
			case STONE:
				//1. if no empty ground is available: do nothing and go in the next tile
				if len(ground) == 0 {
					continue
				}
				//2. if ground is available and there are no obstacles:
				//swap the first empty ground with the stone
				//cut the first ground from the slice
				//(the stone just rolls in the empty space)
				if len(ground) > 0 && len(obstacles) == 0 {
					//swap
					//ground -> stone
					bytes[ground[0]][j] = STONE
					//stone -> ground
					bytes[i][j] = GROUND
					//cut the first ground
					ground = ground[1:]
					//add space cleared by stone to ground
					ground = append(ground, i)
					continue
				}
				//3. if there is empty ground and obstacles:
				//find the first ground that is bigger than the last obstacle
				//swap the stone with this ground
				//remove this ground from the slice
				//(the stone rolls to the first emtpy space before obstacle or another stone)
				lastObstacleIndex := obstacles[len(obstacles)-1]
				for index, groundIndex := range ground {
					if groundIndex > lastObstacleIndex {
						//ground -> stone
						bytes[groundIndex][j] = STONE
						//stone -> ground
						bytes[i][j] = GROUND
						//cut the ground
						if index == len(ground)-1 {
							ground = ground[:index]
						} else if index == 0 {
							ground = ground[1:]
						} else {
							ground = append(ground[:index], ground[index+1:]...)
						}
						//add space cleared by stone to ground
						ground = append(ground, i)
						break
					}
				}
			}
		}
	}

	//ROLL WEST (LEFT)
	for j := 0; j < len(bytes); j++ {
		ground := []int{}
		obstacles := []int{}
		for i := 0; i < len(bytes[0]); i++ {
			switch bytes[j][i] {
			case GROUND:
				ground = append(ground, i)
			case OBSTACLE:
				obstacles = append(obstacles, i)
			case STONE:
				//1. if no empty ground is available: do nothing and go in the next tile
				if len(ground) == 0 {
					continue
				}
				//2. if ground is available and there are no obstacles:
				//swap the first empty ground with the stone
				//cut the first ground from the slice
				//(the stone just rolls in the empty space)
				if len(ground) > 0 && len(obstacles) == 0 {
					//swap
					//ground -> stone
					bytes[j][ground[0]] = STONE
					//stone -> ground
					bytes[j][i] = GROUND
					//cut the first ground
					ground = ground[1:]
					//add space cleared by stone to ground
					ground = append(ground, i)
					continue
				}
				//3. if there is empty ground and obstacles:
				//find the first ground that is bigger than the last obstacle
				//swap the stone with this ground
				//remove this ground from the slice
				//(the stone rolls to the first emtpy space before obstacle or another stone)
				lastObstacleIndex := obstacles[len(obstacles)-1]
				for index, groundIndex := range ground {
					if groundIndex > lastObstacleIndex {
						//ground -> stone
						bytes[j][groundIndex] = STONE
						//stone -> ground
						bytes[j][i] = GROUND
						//cut the ground
						if index == len(ground)-1 {
							ground = ground[:index]
						} else if index == 0 {
							ground = ground[1:]
						} else {
							ground = append(ground[:index], ground[index+1:]...)
						}
						//add space cleared by stone to ground
						ground = append(ground, i)
						break
					}
				}
			}
		}
	}

	//ROLL SOUTH (DOWN)
	for j := 0; j < len(bytes[0]); j++ {
		ground := []int{}
		obstacles := []int{}
		for i := len(bytes) - 1; i > -1; i-- {
			switch bytes[i][j] {
			case GROUND:
				ground = append(ground, i)
			case OBSTACLE:
				obstacles = append(obstacles, i)
			case STONE:
				//1. if no empty ground is available: do nothing and go in the next tile
				if len(ground) == 0 {
					continue
				}
				//2. if ground is available and there are no obstacles:
				//swap the first empty ground with the stone
				//cut the first ground from the slice
				//(the stone just rolls in the empty space)
				if len(ground) > 0 && len(obstacles) == 0 {
					//swap
					//ground -> stone
					bytes[ground[0]][j] = STONE
					//stone -> ground
					bytes[i][j] = GROUND
					//cut the first ground
					ground = ground[1:]
					//add space cleared by stone to ground
					ground = append(ground, i)
					continue
				}
				//3. if there is empty ground and obstacles:
				//find the first ground that is bigger than the last obstacle
				//swap the stone with this ground
				//remove this ground from the slice
				//(the stone rolls to the first emtpy space before obstacle or another stone)
				lastObstacleIndex := obstacles[len(obstacles)-1]
				for index, groundIndex := range ground {
					if groundIndex < lastObstacleIndex {
						//ground -> stone
						bytes[groundIndex][j] = STONE
						//stone -> ground
						bytes[i][j] = GROUND
						//cut the ground
						if index == len(ground)-1 {
							ground = ground[:index]
						} else if index == 0 {
							ground = ground[1:]
						} else {
							ground = append(ground[:index], ground[index+1:]...)
						}
						//add space cleared by stone to ground
						ground = append(ground, i)
						break
					}
				}
			}
		}
	}

	//ROLL EAST (RIGHT)
	for j := 0; j < len(bytes); j++ {
		//fmt.Println("ROLL E first row")
		//fmt.Println(string(bytes[j]))
		ground := []int{}
		obstacles := []int{}
		for i := len(bytes[0]) - 1; i > -1; i-- {
			//fmt.Println("ROLL E")
			switch bytes[j][i] {
			case GROUND:
				ground = append(ground, i)
			case OBSTACLE:
				obstacles = append(obstacles, i)
			case STONE:
				//1. if no empty ground is available: do nothing and go in the next tile
				if len(ground) == 0 {
					continue
				}
				//2. if ground is available and there are no obstacles:
				//swap the first empty ground with the stone
				//cut the first ground from the slice
				//(the stone just rolls in the empty space)
				if len(ground) > 0 && len(obstacles) == 0 {
					//swap
					//ground -> stone
					bytes[j][ground[0]] = STONE
					//stone -> ground
					bytes[j][i] = GROUND
					//cut the first ground
					ground = ground[1:]
					//add space cleared by stone to ground
					ground = append(ground, i)
					continue
				}
				//3. if there is empty ground and obstacles:
				//find the first ground that is bigger than the last obstacle
				//swap the stone with this ground
				//remove this ground from the slice
				//(the stone rolls to the first emtpy space before obstacle or another stone)
				lastObstacleIndex := obstacles[len(obstacles)-1]
				for index, groundIndex := range ground {
					if groundIndex < lastObstacleIndex {
						//ground -> stone
						bytes[j][groundIndex] = STONE
						//stone -> ground
						bytes[j][i] = GROUND
						//cut the ground
						if index == len(ground)-1 {
							ground = ground[:index]
						} else if index == 0 {
							ground = ground[1:]
						} else {
							ground = append(ground[:index], ground[index+1:]...)
						}
						//add space cleared by stone to ground
						ground = append(ground, i)
						break
					}
				}
			}
		}
	}

	return bytes
	//return deepCopy(bytes)
}
