package main

import (
	"fmt"
	"sync/atomic"
)

type RingBuffer struct {
	buffer []int
	size   int32
	read   int32 
	write  int32 
}

func NewRingBuffer(size int) *RingBuffer {
	return &RingBuffer{
		buffer: make([]int, size),
		size:   int32(size),
	}
}

func (rb *RingBuffer) Enqueue(value int) bool {
	writePos := atomic.LoadInt32(&rb.write)
	readPos := atomic.LoadInt32(&rb.read)

	if (writePos+1)%rb.size == readPos {
		return false 
	}

	rb.buffer[writePos] = value

	atomic.StoreInt32(&rb.write, (writePos+1)%rb.size)
	return true
}

func (rb *RingBuffer) Dequeue() (int, bool) {
	readPos := atomic.LoadInt32(&rb.read)
	writePos := atomic.LoadInt32(&rb.write)

	if readPos == writePos {
		return 0, false 
	}

	value := rb.buffer[readPos]

	atomic.StoreInt32(&rb.read, (readPos+1)%rb.size)
	return value, true
}

func (rb *RingBuffer) Size() int {
	writePos := atomic.LoadInt32(&rb.write)
	readPos := atomic.LoadInt32(&rb.read)

	if writePos >= readPos {
		return int(writePos - readPos)
	}
	return int(rb.size - readPos + writePos)
}

func (rb *RingBuffer) IsEmpty() bool {
	return atomic.LoadInt32(&rb.read) == atomic.LoadInt32(&rb.write)
}

func (rb *RingBuffer) IsFull() bool {
	writePos := atomic.LoadInt32(&rb.write)
	readPos := atomic.LoadInt32(&rb.read)
	return (writePos+1)%rb.size == readPos
}

func main() {
	rb := NewRingBuffer(5) 

	fmt.Println(rb.Enqueue(1)) 
	fmt.Println(rb.Enqueue(2)) 
	fmt.Println(rb.Enqueue(3)) 
	fmt.Println(rb.Enqueue(4)) 
	fmt.Println(rb.Enqueue(5)) 

	fmt.Println(rb.Dequeue()) 
	fmt.Println(rb.Dequeue()) 
	fmt.Println(rb.Dequeue()) 

	fmt.Println(rb.Enqueue(6))
	fmt.Println(rb.Enqueue(7))

	fmt.Println(rb.Dequeue())
	fmt.Println(rb.Dequeue())
	fmt.Println(rb.Dequeue())
	fmt.Println(rb.Dequeue())
}
