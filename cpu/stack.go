
package cpu

type StackFrame struct {
	arr []Value
}

func GetStackFrame() *StackFrame {
	return &StackFrame{
		arr:  make([]Value, 64),
	}
}

func (s *StackFrame) Push(i Value) {
	s.arr = append(s.arr, i)
}

func (s *StackFrame) Get() (i Value) {
	return s.arr[len(s.arr)-1]
}

func (s *StackFrame) Pop() (i Value) {
	i, s.arr = s.arr[len(s.arr)-1], s.arr[:len(s.arr)-1]
	return
}
