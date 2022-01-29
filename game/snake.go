package game

type Snake struct {
	X         []Point
	direction Direction
}

func NewSnake() Snake {
	return Snake{
		X:         []Point{{X: 0, Y: 0}},
		direction: DIRECTION_NONE,
	}
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

func (s *Snake) canChangeDirection(direction Direction) bool {
	diff := s.direction - direction
	if diff < 0 {
		diff = -diff
	}

	return diff != 1
}

func (s *Snake) SetDirection(direction Direction) {
	if s.canChangeDirection(direction) {
		s.direction = direction
	}
}

func (s *Snake) GetDirection() Direction {
	return s.direction
}

func (s *Snake) Move() {
	head := s.Head()
	tail := s.Tail()

	switch s.direction {
	case DIRECTION_RIGHT:
		tail.X = head.X + 1
		tail.Y = head.Y
	case DIRECTION_LEFT:
		tail.X = head.X - 1
		tail.Y = head.Y
	case DIRECTION_UP:
		tail.X = head.X
		tail.Y = head.Y - 1
	case DIRECTION_DOWN:
		tail.X = head.X
		tail.Y = head.Y + 1
	}

	s.X = append(s.X[1:], tail)
}
