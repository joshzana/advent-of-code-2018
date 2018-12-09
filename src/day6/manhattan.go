package day6

func Dist(x1 int, y1 int, x2 int, y2 int) int {
	return abs(x2-x1) + abs(y2-y1)
}

func abs(dist int) int {
	if dist < 0 {
		return -dist
	} else {
		return dist
	}
}
