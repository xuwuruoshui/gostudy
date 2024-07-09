package main

import (
	"container/heap"
	"fmt"
)

// 定义一个元素结构体，其中包含值和优先级
type Item[T any] struct {
	value    T // 元素的值
	priority int    // 元素的优先级
	index    int    // 元素在堆中的索引
}

// 优先队列类型，内嵌一个元素切片
type PriorityQueue[T any] []*Item[T]

// 实现 heap.Interface 接口
func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	// 注意我们使用的是大顶堆，优先级越大，优先级越高
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 向优先队列中添加元素
func (pq *PriorityQueue[T]) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item[T])
	item.index = n
	*pq = append(*pq, item)
}

// Pop 从优先队列中移除并返回优先级最高的元素
func (pq *PriorityQueue[T]) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // 为了安全，清除索引
	*pq = old[0 : n-1]
	return item
}

// 更新优先队列中的元素
func (pq *PriorityQueue[T]) update(item *Item[T], value T, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	// 创建优先队列并初始化
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
	pq := make(PriorityQueue[string], len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &Item[string]{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)

	// 向优先队列中添加一个新元素
	item := &Item[string]{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)

	// 更新优先队列中的一个元素
	pq.update(item, item.value, 5)

	// 从优先队列中按优先级顺序取出元素
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item[string])
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
}
