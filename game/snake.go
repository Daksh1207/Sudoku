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
	s.X = append(s.X, s.X[len(s.X)-1])
}

func (s *Snake) Move(direction string) {
	last := s.X[len(s.X)-1]
	first := s.X[0]

	switch direction {
	case "R":
		first.X = last.X + 1
		first.Y = last.Y
	case "L":
		first.X = last.X - 1
		first.Y = last.Y
	case "U":
		first.X = last.X
		first.Y = last.Y - 1
	case "D":
		first.X = last.X
		first.Y = last.Y + 1
	}

	s.X = append(s.X[1:], first)
}
