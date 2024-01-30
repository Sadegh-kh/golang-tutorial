package structuers

import "fmt"

type Student struct {
	ID        uint
	FirstName string
	LastName  string
	Age       int
	isActive  bool
	Scores    []float32
}

func (s Student) ScoresAvg() float32 {
	var sum float32
	for _, value := range s.Scores {
		sum += float32(value)
	}
	return sum / float32(len(s.Scores))
}

func (s Student) PassedStudent() bool {
	if s.ScoresAvg() < 12 {
		return false
	}
	return true
}

func (s *Student) SetName(firstname, lastname string) {
	s.FirstName = firstname
	s.LastName = lastname
}
func (s Student) GetName() string {
	return fmt.Sprintf("firstname %s and lastname %s", s.FirstName, s.LastName)
}
