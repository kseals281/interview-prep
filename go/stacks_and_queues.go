package main

type stack struct {
	top  *node
	size int
}

type queue struct {
	head *node
	tail *node
	size int
}

type node struct {
	num  int
	next *node
}

func (s *stack) push(n *node) {
	n.next = s.top
	s.top = n
	s.size++
}

func (s *stack) pop() *node {
	n := s.top
	s.top = s.top.next
	s.size--
	return n
}

func (s *stack) isEmpty() bool {
	return s.size == 0 && s.top == nil
}

func (q *queue) add(n *node) {
	q.tail.next = n
	q.tail = n
	q.size++
}

func (q *queue) remove() *node {
	n := q.head
	q.head = q.head.next
	q.size--
	return n
}

func (q *queue) isEmpty() bool {
	return q.size == 0 && q.head == nil
}

// Implement a MyQueue class which implement a queue using two stacks
type MyQueue struct {
	inQueue  stack
	outQueue stack
}

func (m *MyQueue) add(n *node) {
	m.inQueue.push(n)
}

func (m *MyQueue) remove() *node {
	return m.outQueue.pop()
}

func (m *MyQueue) flush() {
	for !m.inQueue.isEmpty() {
		m.outQueue.push(m.inQueue.pop())
	}
}

// How would you design a stack which, in addition to push and pop, has a function min which returns the minimum element?
// Push, pop, and min should all operate in O(n) time
func (s stack) min() *node {
	var temp stack
	m := s.top
	for !s.isEmpty() {
		if s.top.num < m.num {
			m = s.top
		}
		temp.push(s.pop())
	}
	for !temp.isEmpty() {
		s.push(temp.pop())
	}
	return m
}
