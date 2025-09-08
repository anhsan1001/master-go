package student

import (
	"fmt"
	"math/rand"
)

type Student struct {
	ID     int
	Name   string
	Class  string
	Scores []float64
}

func New(student Student) *Student {
	return &Student{
		ID:     rand.Int(),
		Name:   student.Name,
		Class:  student.Class,
		Scores: student.Scores,
	}
}

func (s *Student) GetStudentInfo() {
	fmt.Printf("ID: %d | Name: %s | Class: %s | Average score: %.2f\n", s.ID, s.Name, s.Class, s.CalculateAverageScore())
}
func (s *Student) CalculateAverageScore() float64 {

	var total float64

	for _, score := range s.Scores {

		total += score
	}
	return total / float64(len(s.Scores))
}

func (s *Student) Update(payload Student) {
	if payload.Name != "" {
		s.Name = payload.Name
	}

	if payload.Class != "" {
		s.Class = payload.Class
	}
	if len(payload.Scores) != 0 {
		copy(s.Scores, payload.Scores)
	}
	fmt.Println("Updated student successfully")
}
