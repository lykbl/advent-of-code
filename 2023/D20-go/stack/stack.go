package stack

type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Len() int {
  return len(s.items)
}

func (s *Stack[T]) Shift() (T, bool) {
  if len(s.items) == 0 {
    var zeroValue T
    return zeroValue, false
  }

  shifted := s.items[0]
  s.items = s.items[1:]

  return shifted, true
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zeroValue T
        return zeroValue, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

func (s *Stack[T]) Peek() (T, bool) {
    if len(s.items) == 0 {
        var zeroValue T
        return zeroValue, false
    }
    return s.items[len(s.items)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
    return len(s.items)
}

