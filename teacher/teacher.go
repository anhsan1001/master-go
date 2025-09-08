package teacher

import (
	"fmt"
	"math/rand"
)

type Teacher struct {
	ID      int
	Name    string
	Subject string
	Class   []string
}

func New(teacher Teacher) *Teacher {
	return &Teacher{
		ID:      rand.Int(),
		Name:    teacher.Name,
		Subject: teacher.Subject,
		Class:   teacher.Class,
	}
}

func (t *Teacher) GetTeacherInfo() {
	fmt.Printf("ID: %d | Name: %s | Subject: %s |Class: %s \n", t.ID, t.Name, t.Subject, t.Class)
}
