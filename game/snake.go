package game

type Point struct {
	X int
	Y int
}

type Snake struct {
	X []Point
}

func (s *Snake) IsIn(p Point) bool {
	for _, sp := range s.X {
		if p.X == sp.X && sp.Y == p.Y {
			return true
		}
	}

	return false
}

func (s *Snake) Grow() {
	// TODO remove loop and draw squares far away
	for i := 0; i < 5; i++ {
		s.X = append(s.X, s.X[len(s.X)-1])
	}
}

func (s *Snake) Head() Point {
	return s.X[len(s.X)-1]
}

func (s *Snake) Tail() Point {
	return s.X[0]
}

func (s *Snake) Move(direction string) {
	head := s.Head()
	tail := s.Tail()

	switch direction {
	case "R":
		tail.X = head.X + 1
		tail.Y = head.Y
	case "L":
		tail.X = head.X - 1
		tail.Y = head.Y
	case "U":
		tail.X = head.X
		tail.Y = head.Y - 1
	case "D":
		tail.X = head.X
		tail.Y = head.Y + 1
	}

	s.X = append(s.X[1:], tail)
}
