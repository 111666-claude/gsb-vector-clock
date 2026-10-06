package vclock

import "testing"

func TestTickIncrementsOneSlot(t *testing.T) {
	c := New(3)
	c.Tick(1)
	if c.Vec[1] != 1 || c.Vec[0] != 0 {
		t.Fatalf("只应第 1 位加一：%v", c.Vec)
	}
}

func TestEqualVectorsAreConcurrent(t *testing.T) {
	if got := Order([]int{1, 1}, []int{1, 1}); got != "concurrent" {
		t.Fatalf("相等应算 concurrent，得到 %s", got)
	}
}

func TestBasicBefore(t *testing.T) {
	if got := Order([]int{1, 0}, []int{1, 1}); got != "before" {
		t.Fatalf("应判 before，得到 %s", got)
	}
}

func TestDedupKeepsDistinct(t *testing.T) {
	got := Dedup([]string{"a", "b", "c"}, &Counter{})
	if got != 3 {
		t.Fatalf("三条不同消息应剩 3 条，得到 %d", got)
	}
}

func TestNewClockIsZeroed(t *testing.T) {
	c := New(4)
	if len(c.Vec) != 4 || c.Vec[3] != 0 {
		t.Fatalf("新时钟应全零：%v", c.Vec)
	}
}
