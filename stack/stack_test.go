package stack

import (
	"testing"
)

type testCase[T any] struct {
	name         string
	in           []T
	defaultValue T
}

func initCases[T any](slice ...T) *[]testCase[T] {
	return &[]testCase[T]{
		{
			name: "empty",
			in:   []T{},
		},
		{
			name: "not empty",
			in:   slice,
		},
	}
}

func TestStackClear(t *testing.T) {
	cases := *initCases(2, 5, 8, 11, 14)

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			s := New(tCase.in...)
			s.Clear()

			if s.root != nil || s.len != 0 {
				t.Errorf("Case '%s' failed: s.root = %+v s.len = %d, want %+v %d",
					tCase.name, s.root, s.len, nil, 0)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	cases := *initCases(2.1, 5.2, 8.3, 11.4, 14.5)

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			s := New(tCase.in...)

			expect := len(tCase.in) == 0

			if s.IsEmpty() != expect {
				t.Errorf("Case '%s' failed: s.IsEmpty() = %t, want %t",
					tCase.name, s.IsEmpty(), expect)
			}
		})
	}
}

func TestLen(t *testing.T) {
	cases := initCases("2", "5", "8", "11", "14")

	for _, tCase := range *cases {
		t.Run(tCase.name, func(t *testing.T) {
			s := New(tCase.in...)

			if s.Len() != len(tCase.in) {
				t.Errorf("Case '%s' failed: s.Len() = %d, want %d",
					tCase.name, s.Len(), len(tCase.in))
			}
		})
	}
}

func TestPop(t *testing.T) {
	cases := *initCases('a', 'b', 'c', 'd', 'e')

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			s := New(tCase.in...)

			for i := s.Len() - 1; !s.IsEmpty(); i-- {
				if val, ok := s.Pop(); val != tCase.in[i] && ok == true {
					t.Errorf("Case '%s' failed: val = %+v ok = %t, want %+v %t",
						tCase.name, val, ok, tCase.in[i], true)
				}
			}

			// Pop from an empty stack
			if val, ok := s.Pop(); val != tCase.defaultValue || ok == true {
				t.Errorf("Case '%s' val = %+v ok = %t, want %+v %t",
					tCase.name, val, ok, tCase.defaultValue, false)
			}
		})
	}
}

func TestPeek(t *testing.T) {
	cases := *initCases(2, 5, 8, 11, 14)

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			s := New(tCase.in...)

			if len(tCase.in) > 0 {
				last := len(tCase.in) - 1
				if val, ok := s.Peek(); val != tCase.in[last] && ok == true {
					t.Errorf("Case '%s' failed: val = %+v ok = %t, want %+v %t",
						tCase.name, val, ok, tCase.in[last], true)
				}
			} else {
				// Peek from an empty stack
				if val, ok := s.Peek(); val != tCase.defaultValue || ok == true {
					t.Errorf("Case '%s' val = %+v ok = %t, want %+v %t",
						tCase.name, val, ok, tCase.defaultValue, false)
				}
			}
		})
	}
}

func TestPeekForSlice(t *testing.T) {
	cases := *initCases([]int{1, 2}, []int{5, 6}, []int{8, 9}, []int{11, 12}, []int{14, 15})

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			s := New(tCase.in...)

			if len(tCase.in) > 0 {
				last := len(tCase.in) - 1

				slice, ok := s.Peek()

				if len(slice) != len(tCase.in[last]) || ok != true {
					t.Errorf("Case '%s' failed: val = %+v ok = %t, want %+v %t",
						tCase.name, slice, ok, tCase.in[last], true)
				}

				for i, v := range slice {
					if v != tCase.in[last][i] {
						t.Errorf("Case '%s' failed: val = %+v, want %+v",
							tCase.name, v, tCase.in[last][i])
					}
				}
			} else {
				// Peek from an empty stack
				if val, ok := s.Peek(); len(val) != 0 || ok == true {
					t.Errorf("Case '%s' val = %+v ok = %t, want %+v %t",
						tCase.name, val, ok, 0, false)
				}
			}
		})
	}
}

func TestPush(t *testing.T) {
	cases := *initCases(2, 5, 8, 11, 14)

	for _, tCase := range cases {
		t.Run(tCase.name, func(t *testing.T) {
			s := New[int]()

			for _, v := range tCase.in {
				s.Push(v)

				if val, ok := s.Peek(); val != v || ok != true {
					t.Errorf("Case '%s' failed: val = %+v, want %+v",
						tCase.name, val, v)
				}
			}

			if len(tCase.in) != s.Len() {
				t.Errorf("Case '%s' failed: s.Len = %d, want %d",
					tCase.name, s.Len(), len(tCase.in))
			}
		})
	}
}
