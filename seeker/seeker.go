package seeker

import (
	"math/rand"

	"github.com/kaniusii/guessinggame/keeper"
)

type Seeker struct {
	Keeper *keeper.Keeper
	Turns  int
	Val    int
}

func (s *Seeker) BisectionBob() {
	s.Turns = 0
	min := s.Keeper.Min - 1
	max := s.Keeper.Max + 1

	found := false
	for !found {
		s.Val = (min + max) / 2
		s.Turns++
		res := s.Keeper.MakeGuess(s.Val)

		if res > 0 {
			min = s.Val
		} else if res < 0 {
			max = s.Val
		} else {
			found = true
		}
	}
}

func (s *Seeker) RandomRobert() {
	s.Turns = 0
	min := s.Keeper.Min
	max := s.Keeper.Max

	found := false
	for !found {
		s.Val = min + rand.Intn(max-min)
		s.Turns++

		res := s.Keeper.MakeGuess(s.Val)

		if res == 0 {
			return
		}
	}
}

func (s *Seeker) StartSam() {
	s.Turns = 0
	min := s.Keeper.Min
	max := s.Keeper.Max

	for i := min; i <= max; i++ {
		s.Turns++

		res := s.Keeper.MakeGuess(i)
		if res == 0 {
			s.Val = i
			return
		}
	}
}

func (s *Seeker) EndEmily() {
	s.Turns = 0
	min := s.Keeper.Min
	max := s.Keeper.Max

	for i := max; i >= min; i-- {
		s.Turns++

		res := s.Keeper.MakeGuess(i)
		if res == 0 {
			s.Val = i
			return
		}
	}
}
