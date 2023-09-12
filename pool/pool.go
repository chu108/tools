package main

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type WorkerPool struct {
	taskQueue chan func()
	capacity  int
	wg        sync.WaitGroup
	taskCount int32
}

// NewPool 创建协程池
func NewPool(capacity int) *WorkerPool {
	wp := new(WorkerPool)
	if capacity == 0 {
		wp.capacity = runtime.NumCPU()
	} else {
		wp.capacity = capacity
	}
	wp.taskQueue = make(chan func(), wp.capacity)
	wp.execTask()
	return wp
}

// 执行任务
func (w *WorkerPool) execTask() {
	for i := 0; i < w.capacity; i++ {
		w.wg.Add(1)
		go func() {
			defer w.wg.Done()
			for fn := range w.taskQueue {
				if fn != nil {
					fn()
					atomic.AddInt32(&w.taskCount, -1)
				}
			}
		}()
	}
}

// Go 添加任务
func (w *WorkerPool) Go(fn func()) {
	w.taskQueue <- fn
	atomic.AddInt32(&w.taskCount, 1)
}

// Wait 等待所有任务执行完毕
func (w *WorkerPool) Wait() {
	w.wg.Wait()
}

// WaitClose 等待所有任务执行完毕
func (w *WorkerPool) WaitClose() {
	w.Close()
	w.wg.Wait()
}

// Close 关闭任务队列
func (w *WorkerPool) Close() {
	close(w.taskQueue)
}

// TaskCount 返回当前队列中的任务数量
func (w *WorkerPool) TaskCount() int32 {
	return atomic.LoadInt32(&w.taskCount)
}
