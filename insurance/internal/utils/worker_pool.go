package utils

import (
	"context"
	"time"
)

type Pool interface {
	Add(task Task) error
	Start()
}

type Task func() error

func NewPool(ctx context.Context, size int) Pool {
	return &pool{
		ctx:   ctx,
		size:  size,
		tasks: make(chan Task, size),
	}
}

type pool struct {
	ctx  context.Context
	size int

	tasks chan Task
}

func (p pool) Add(task Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	select {
	case p.tasks <- task:
	case <-ctx.Done():
		return NewError("context timeout: failed to add task", Internal)
	}

	return nil
}

func (p pool) Start() {
	for i := 0; i < p.size; i++ {
		go func() {
			for {
				select {
				case task := <-p.tasks:
					_ = task()
				case <-p.ctx.Done():
					return
				}
			}
		}()
	}
}
