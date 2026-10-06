// Package vclock 维护向量时钟并判定因果关系。
package vclock

// Counter 记录本轮检查过的历史消息条数。
type Counter struct {
	Scanned int
}

// Clock 是一个节点的本地向量时钟与已处理消息台账。
type Clock struct {
	Vec  []int
	Seen map[string]bool
}

// New 建一个长度为 n 的空时钟。
func New(n int) *Clock { return &Clock{Vec: make([]int, n), Seen: map[string]bool{}} }

// Tick 让第 node 位自增。
func (c *Clock) Tick(node int) {
	c.Vec[node]++
}

// Receive 收到一条消息，做合并再自增。
func (c *Clock) Receive(node int, msgID string, remote []int) {
	if c.Seen[msgID] {
		return
	}
	for i := range c.Vec {
		if i < len(remote) && remote[i] > c.Vec[i] {
			c.Vec[i] = remote[i]
		}
	}
	c.Seen[msgID] = true
	c.Tick(node)
}

// Before 判定 a 是否严格早于 b。
func Before(a, b []int) bool {
	less := false
	n := len(a)
	if len(b) > n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		av, bv := 0, 0
		if i < len(a) {
			av = a[i]
		}
		if i < len(b) {
			bv = b[i]
		}
		if av > bv {
			return false
		}
		if av < bv {
			less = true
		}
	}
	return less
}

// Order 给出两个时钟的关系：before、after、equal 或 concurrent。
func Order(a, b []int) string {
	if Before(a, b) {
		return "before"
	}
	if Before(b, a) {
		return "after"
	}
	return "concurrent"
}

// Dedup 按消息号去重，返回去重后的条数。
func Dedup(ids []string, c *Counter) int {
	seen := map[string]bool{}
	unique := 0
	for _, id := range ids {
		c.Scanned++
		if !seen[id] {
			seen[id] = true
			unique++
		}
	}
	return unique
}
