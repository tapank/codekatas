package school

import (
	"slices"
	"sort"
)

type Grade struct {
	level    int
	students []string
}

type School struct {
	levels []int
	grades map[int]*Grade
}

func New() *School {
	return &School{levels: []int{}, grades: map[int]*Grade{}}
}

func (s *School) Add(student string, grade int) {
	// create grade if needed
	if !slices.Contains((*s).levels, grade) {
		s.levels = append(s.levels, grade)
		sort.Ints(s.levels)
		s.grades[grade] = &Grade{grade, []string{}}
	}

	// ensure no duplicate student names in a grade
	if !slices.Contains((*s).grades[grade].students, student) {
		gr := (*s).grades[grade]
		gr.students = append(gr.students, student)
		sort.Strings(gr.students)
	}
}

func (s *School) Grade(level int) (students []string) {
	if slices.Contains((*s).levels, level) {
		students = (*s).grades[level].students
	}
	return
}

func (s *School) Enrollment() (grades []Grade) {
	for _, grade := range (*s).levels {
		grades = append(grades, *(*s).grades[grade])
	}
	return
}
