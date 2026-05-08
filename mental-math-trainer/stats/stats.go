package stats

import "fmt"

type Stats struct {
	Correct      int
	Wrong        int
	CurrentLevel int
}

func NewStats() *Stats {
	return &Stats{CurrentLevel: 1}
}

func (s *Stats) SetLevel(level int) {
	s.CurrentLevel = level
}

func (s *Stats) AddCorrect() {
	s.Correct++
}

func (s *Stats) AddWrong() {
	s.Wrong++
}

func (s *Stats) Print() {
	fmt.Println("===== СТАТИСТИКА =====")
	fmt.Println("Текущий уровень сложности:", s.CurrentLevel)
	fmt.Println("Верных ответов:", s.Correct)
	fmt.Println("Неверных ответов:", s.Wrong)
}
