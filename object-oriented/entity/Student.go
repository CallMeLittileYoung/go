package entity

// Student /*

type Student struct {
	id     uint
	name   string
	age    uint8
	mael   bool
	scores float64
}

func (s *Student) Id() uint {
	return s.id
}

func (s *Student) SetId(id uint) {
	s.id = id
}

func (s *Student) Name() string {
	return s.name
}

func (s *Student) SetName(name string) {
	s.name = name
}

func (s *Student) Age() uint8 {
	return s.age
}

func (s *Student) SetAge(age uint8) {
	s.age = age
}

func (s *Student) Mael() bool {
	return s.mael
}

func (s *Student) SetMael(mael bool) {
	s.mael = mael
}

func (s *Student) Scores() float64 {
	return s.scores
}

func (s *Student) SetScores(scores float64) {
	s.scores = scores
}

func NewStudent(id uint, name string, age uint8, mael bool, scores float64) *Student {
	return &Student{id: id, name: name, age: age, mael: mael, scores: scores}
}
