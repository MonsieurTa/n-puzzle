package state

import (
	"fmt"

	"github.com/MonsieurTa/n-puzzle/utils"
)

type State struct {
	Board          [][]int
	size           int
	BlankX, BlankY int
	Key            string
	Cache          map[int][2]int
}

func (s *State) CacheBoard() {
	s.Cache = map[int][2]int{}
	for y, row := range s.Board {
		for x, tile := range row {
			s.Cache[tile] = [2]int{x, y}
		}
	}
}

func (s *State) CachedPos(value int) (int, int) {
	return s.Cache[value][0], s.Cache[value][1]
}

func toString(state [][]int) string {
	base := ""
	for _, item := range state {
		base += fmt.Sprint(item)
	}
	return base
}

//NewState Instantiate new state
func NewState(board [][]int) *State {
	x, y := getBlankPos(board)
	return &State{
		Board:  board,
		size:   len(board),
		Key:    toString(board),
		BlankX: x, BlankY: y,
	}
}

//CompareState compare two state
func (s State) CompareState(state State) bool {
	return s.Key == state.Key
}

func getBlankPos(board [][]int) (int, int) {
	for y := range board {
		for x := range board[y] {
			if board[y][x] == 0 {
				return x, y
			}
		}
	}
	return -1, -1
}

func (s State) shiftBlank(x, y int) *State {
	if x < 0 || y < 0 || x >= s.size || y >= s.size {
		return nil
	}
	new := s
	new.Board = utils.DeepCopy(new.Board)
	new.BlankX = x
	new.BlankY = y
	new.Board[y][x] = 0
	new.Board[s.BlankY][s.BlankX] = s.Board[y][x]
	new.Key = toString(new.Board)
	return &new
}

//GetSurrounding get left, right, top, down state from current state
func (s State) GetSurrounding() []*State {
	surroundings := []*State{
		s.shiftBlank(s.BlankX-1, s.BlankY),
		s.shiftBlank(s.BlankX+1, s.BlankY),
		s.shiftBlank(s.BlankX, s.BlankY-1),
		s.shiftBlank(s.BlankX, s.BlankY+1),
	}
	ret := make([]*State, 0)
	for _, item := range surroundings {
		if item != nil {
			ret = append(ret, item)
		}
	}
	return ret
}
