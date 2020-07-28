package queue

import "container/list"

// Queue 队列
type Queue interface {
	// Push 入列
	Push(items ...interface{})
	// Pop 出列
	Pop() (interface{})
	// MPop 批量出列
	MPop(n int) ([]interface{})
	// Length 长度
	Length() int
	// Clear 清空
	Clear()
}

type queue struct {
	// 使用链表结构实现
	list *list.List
}

func New() Queue {
	return &queue{list: list.New()}
}

func (q queue) Push(items ...interface{}) {
	for _, item := range items {
		q.list.PushBack(item)
	}
}

func (q queue) Pop() (interface{}) {
	item := q.list.Front()
	q.list.Remove(item)
	return item.Value
}

func (q queue) MPop(n int) ([]interface{}) {
	if n < 1 { // 当n小于1，返回nil
		return nil
	}
	var result []interface{}
	var items []*list.Element
	item := q.list.Front()

	for i := 0; i <n ; i++ {
		if item == nil { // 判断元素是否为空
			break
		}

		items = append(items, item)
		item = item.Next()
	}
	// 获取元素并删除队列里对应的数据
	for _, item := range items {
		q.list.Remove(item)
		result = append(result, item.Value)
	}
	return result
}

func (q queue) Length() int {
	return q.list.Len()
}

func (q queue) Clear() {
	q.list.Init()
}
