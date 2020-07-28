package queue

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := New()
	// test Push
	queue.Push(1,2,3,4)

	// test Length
	assert.Equal(t, queue.Length(), 4)

	// test Pop
	item := queue.Pop()
	assert.Equal(t,  item, 1)
	assert.Equal(t,  queue.Length(), 3)

	// test MPop
	multiItems := queue.MPop(2)
	assert.Equal(t, []interface{}{2,3}, multiItems)
	assert.Equal(t,  queue.Length(), 1)

	// test Clear
	queue.Clear()
	assert.Equal(t, queue.Length(), 0)

}

func BenchmarkQueue_Push(b *testing.B) {
	queue := New()
	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}
	fmt.Println(queue.Length())
}