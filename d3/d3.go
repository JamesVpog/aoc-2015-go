package d3

import "fmt"

// set of coords, start at 0,0, when encounter direction add 1 or -1 to that direction
// return the length of that set
type twoDPoint struct {
	x int
	y int
}

func addToSet(set map[twoDPoint]struct{}, point twoDPoint) {
	_, ok := set[point]
	if !ok {
		set[point] = struct{}{}
	}
}
func P1(content string) {
	// read content, each char is >,^,v,<
	var start twoDPoint
	start.x = 0
	start.y = 0

	points := make(map[twoDPoint]struct{})

	addToSet(points, start)

	for _, r := range content {
		switch r {
		case '>':
			start.x += 1
		case 'v':
			start.y -= 1
		case '<':
			start.x -= 1
		case '^':
			start.y += 1
		}
		addToSet(points, start)
	}
	fmt.Printf("%d\n", len(points))
}

func moveAndAdd(r rune, point *twoDPoint, points map[twoDPoint]struct{}) {
	switch r {
	case '>':
		point.x += 1
	case 'v':
		point.y -= 1
	case '<':
		point.x -= 1
	case '^':
		point.y += 1
	}
	addToSet(points, *point)
}
func P2(content string) {
	santa := twoDPoint{x: 0, y: 0}
	robo := twoDPoint{x: 0, y: 0}

	turn := true

	points := make(map[twoDPoint]struct{})

	addToSet(points, santa)
	addToSet(points, robo)

	for _, r := range content {
		if turn {
			moveAndAdd(r, &santa, points)
		} else {
			moveAndAdd(r, &robo, points)
		}

		turn = !turn
	}
	fmt.Printf("%d\n", len(points))

}
