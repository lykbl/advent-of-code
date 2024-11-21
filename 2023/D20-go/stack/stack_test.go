
package stack

import "testing"

func TestStack(t *testing.T) {
    s := Stack[int]{}

    s.Push(1)
    s.Push(2)

    if s.Size() != 2 {
        t.Errorf("expected size 2, got %d", s.Size())
    }

    top, ok := s.Peek()
    if !ok || top != 2 {
        t.Errorf("expected top 2, got %v", top)
    }

    item, ok := s.Pop()
    if !ok || item != 2 {
        t.Errorf("expected pop 2, got %v", item)
    }

    if s.IsEmpty() {
        t.Errorf("expected stack not to be empty")
    }
}
