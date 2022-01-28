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
	switch direction {
	case "R":
		s.X[0].X += 1
	case "L":
		s.X[0].X -= 1
	case "U":
		s.X[0].Y -= 1
	case "D":
		s.X[0].Y += 1
	}

}
