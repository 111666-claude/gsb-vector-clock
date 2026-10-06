// Command vclock 是向量时钟的场景入口。
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"example.com/vector-clock/vclock"
)

func emit(v any) {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func main() {
	sample := flag.String("sample", "", "merge|concurrent|dup|work")
	flag.Parse()

	c := &vclock.Counter{}
	switch *sample {
	case "merge":
		clock := vclock.New(2)
		clock.Tick(0)
		clock.Receive(0, "m1", []int{0, 1})
		emit(map[string]any{"vec": clock.Vec})
	case "concurrent":
		emit(map[string]string{"order": vclock.Order([]int{1, 0}, []int{0, 1})})
	case "dup":
		clock := vclock.New(1)
		clock.Receive(0, "m1", nil)
		clock.Receive(0, "m1", nil)
		emit(map[string]any{"vec": clock.Vec})
	case "work":
		ids := []string{}
		for i := 0; i < 2000; i++ {
			ids = append(ids, fmt.Sprintf("m%04d", i))
		}
		got := vclock.Dedup(ids, c)
		emit(map[string]int{"ids": 2000, "unique": got, "scanned": c.Scanned})
	default:
		fmt.Fprintln(os.Stderr, "未知场景："+*sample)
		os.Exit(2)
	}
}
