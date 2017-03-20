package querybuilder

import (
	"fmt"
	"strings"
)

type stack struct {
	data []string
}

func (s *stack) Push(val string) {
	s.data = append(s.data, val)
}

func (s *stack) Pop() string {
	if len(s.data) == 0 {
		return ""
	}
	retData := s.data[len(s.data)-1]
	s.data[len(s.data)-1] = ""
	s.data = s.data[:len(s.data)-1]

	return retData
}

func (s *stack) Top() string {
	if len(s.data) == 0 {
		return ""
	}
	return s.data[len(s.data)-1]
}

func (s *stack) Empty() bool {
	return len(s.data) == 0
}

func (s *stack) String() string {
	return fmt.Sprintf("%s, len: %v", strings.Join(s.data, ","), len(s.data))
}
