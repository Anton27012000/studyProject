package stats

import "fmt"

type Stats struct {
	Correct int
	Wrong   int
}

func NewStats() *Stats {
	fmt.Println("Сбор игровой статистики запущен")
	return &Stats{}
}

func (s *Stats) AddCorrect() {
	s.Correct++
}

func (s *Stats) AddWrong() {
	s.Wrong++
}

func (s *Stats) Print() {
	fmt.Println("===== СТАТИСТИКА =====")
	fmt.Println("Верных ответов:", s.Correct)
	fmt.Println("Неверных ответов:", s.Wrong)
}
