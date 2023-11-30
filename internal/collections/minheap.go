package collections

import "github.com/b5710546232/go-pathfinder/pathfinder/model"

type Item struct {
	value    model.Node
	priority float64
}

type MinHeap struct {
	array []Item
}

func NewMinHeap(cap int) *MinHeap {
	return &MinHeap{array: make([]Item, 0, cap)}
}

func (pq *MinHeap) Len() int {
	return len(pq.array)
}

func (pq *MinHeap) Push(value model.Node, priority float64) {
	item := Item{value: value, priority: priority}
	pq.array = append(pq.array, item)
	pq.heapifyUp(len(pq.array) - 1)
}

func (pq *MinHeap) Pop() model.Node {
	if len(pq.array) == 0 {
		panic("Cannot pop from an empty priority queue")
	}
	min := pq.array[0].value
	lastIdx := len(pq.array) - 1
	pq.array[0] = pq.array[lastIdx]
	pq.array = pq.array[:lastIdx]
	pq.heapifyDown(0)
	return min
}

func (pq *MinHeap) heapifyUp(idx int) {
	for {
		parentIdx := (idx - 1) / 2
		if parentIdx < 0 || pq.array[parentIdx].priority <= pq.array[idx].priority {
			break
		}
		pq.array[parentIdx], pq.array[idx] = pq.array[idx], pq.array[parentIdx]
		idx = parentIdx
	}
}

func (pq *MinHeap) heapifyDown(idx int) {
	size := len(pq.array)
	for {
		leftIdx := 2*idx + 1
		rightIdx := 2*idx + 2
		smallestIdx := idx
		if leftIdx < size && pq.array[leftIdx].priority < pq.array[smallestIdx].priority {
			smallestIdx = leftIdx
		}
		if rightIdx < size && pq.array[rightIdx].priority < pq.array[smallestIdx].priority {
			smallestIdx = rightIdx
		}
		if smallestIdx == idx {
			break
		}
		pq.array[idx], pq.array[smallestIdx] = pq.array[smallestIdx], pq.array[idx]
		idx = smallestIdx
	}
}
