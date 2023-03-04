package main

import (
	"fmt"
)

type Student struct {
	name  string
	score []int
}

func (s *Student) addScore(score int) {
	s.score = append(s.score, score)
}

func (s *Student) averageScore() int {
	sum := 0
	for _, score := range s.score {
		sum += score
	}
	return sum / len(s.score)
}

func (s *Student) minScore() (string, int) {
	minScore := s.score[0]
	minName := s.name
	for i := 1; i < len(s.score); i++ {
		minScore = s.score[i]
		minName = s.name
	}
	return minName, int(minScore)
}

func (s *Student) maxScore() (string, int) {
	maxScore := s.score[0]
	maxName := s.name
	for i := 1; i < len(s.score); i++ {
		maxScore = s.score[i]
		maxName = s.name
	}
	return maxName, int(maxScore)
}

func main() {
	students := make([]Student, 0, 5)

	for i := 0; i < 5; i++ {
		var name string
		var score float64
		fmt.Printf("Masukkan nama siswa ke-%d: ", i+1)
		fmt.Scanln(&name)
		student := Student{name: name}
		for j := 0; j < 1; j++ {
			fmt.Printf("Masukkan nilai ke-%d untuk siswa %s: ", j+1, name)
			fmt.Scanln(&score)
			student.addScore(int(score))
		}
		students = append(students, student)
	}

	var totalScore int
	for _, s := range students {
		totalScore += s.averageScore()
	}
	fmt.Printf("Average score for all students: %v\n", totalScore/len(students))

	minScore := students[0].score[0]
	minName := students[0].name
	maxScore := students[0].score[0]
	maxName := students[0].name

	for _, s := range students {
		name, score := s.minScore()
		if score < int(minScore) {
			minScore = score
			minName = name
		}
		name, score = s.maxScore()
		if score > int(maxScore) {
			maxScore = score
			maxName = name
		}
	}

	fmt.Printf("Student with minimum score: %s (%v)\n", minName, minScore)
	fmt.Printf("Student with maximum score: %s (%v)\n", maxName, maxScore)
}
