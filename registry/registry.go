package registry

import "fmt"

type Part string

const (
	A Part = "a"
	B Part = "b"
)

type SolutionFunc func(input string, isExample bool) (string, error)

var sols = map[int]map[Part]SolutionFunc{}

func Register(day int, part Part, fn SolutionFunc) {
	if sols[day] == nil {
		sols[day] = map[Part]SolutionFunc{}
	}
	if _, exists := sols[day][part]; exists {
		panic(fmt.Sprintf("duplicate registration: day %02d part %s", day, part))
	}
	sols[day][part] = fn
}

func Get(day int, part Part) (SolutionFunc, bool) {
	m := sols[day]
	if m == nil {
		return nil, false
	}
	fn, ok := m[part]
	return fn, ok
}

func Days() []int {
	out := make([]int, 0, len(sols))
	for d := range sols {
		out = append(out, d)
	}
	return out
}

func Parts(day int) []Part {
	m := sols[day]
	out := make([]Part, 0, len(m))
	for p := range m {
		out = append(out, p)
	}
	return out
}
