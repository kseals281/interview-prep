package main

import (
	"reflect"
	"testing"
)

func TestMyQueue(t *testing.T) {
	tests := []struct {
		name     string
		inQueue  stack
		outQueue stack
		nodes    []*node
		want     []int
	}{
		{
			"Simple queue",
			stack{},
			stack{},
			[]*node{{num: 5, next: nil}, {num: 3, next: nil}, {num: 1, next: nil}},
			[]int{5, 3, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MyQueue{
				inQueue:  tt.inQueue,
				outQueue: tt.outQueue,
			}
			for _, n := range tt.nodes {
				m.add(n)
			}
			m.flush()
			for !m.outQueue.isEmpty() {
				n := m.remove()
				if n.num != tt.want[0] {
					t.Errorf("Got %d, want %d", n.num, tt.want[0])
				}
				tt.want = tt.want[1:]
			}
		})
	}
}

func TestMyQueue_flush(t *testing.T) {
	tests := []struct {
		name     string
		inQueue  stack
		outQueue stack
	}{
		{
			"One value",
			stack{
				&node{
					5,
					nil,
				},
				1,
			},
			stack{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := MyQueue{
				inQueue:  tt.inQueue,
				outQueue: tt.outQueue,
			}
			m.flush()
			if !reflect.DeepEqual(m.inQueue.size, 0) {
				t.Errorf("Enqueue size was above 0: %d", tt.inQueue.size)
			} else if reflect.DeepEqual(m.outQueue.size, 0) {
				t.Errorf("Got no values into output")
			}
		})
	}
}

func Test_stack_push(t *testing.T) {
	type fields struct {
		top  *node
		size int
	}
	type args struct {
		n *node
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantSize int
	}{
		{
			"One push",
			fields{
				top:  nil,
				size: 0,
			},
			args{&node{5, nil}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stack{
				top:  tt.fields.top,
				size: tt.fields.size,
			}
			s.push(tt.args.n)
			if !reflect.DeepEqual(s.size, tt.wantSize) {
				t.Errorf("Got stack of size %d, want %d", s.size, tt.wantSize)
			}
		})
	}
}

func Test_stack_min(t *testing.T) {
	type fields struct {
		top  *node
		size int
	}
	tests := []struct {
		name   string
		fields fields
		want   *node
	}{
		{
			"clear minimum",
			fields{
				top:  &node{5, nil},
				size: 1,
			},
			&node{5, nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := stack{
				top:  tt.fields.top,
				size: tt.fields.size,
			}
			if got := s.min(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}
