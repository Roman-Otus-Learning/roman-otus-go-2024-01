package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type TaskRunnerInterface interface {
	Run() error
}

type TaskRunner struct {
	wg             *sync.WaitGroup
	tasks          []Task
	workersCount   int
	maxErrorsCount int
	errorCounter   int32
}

func (taskRunner *TaskRunner) Run() error {
	taskChannel := taskRunner.pushTasksToChannel()
	taskRunner.runTasks(taskChannel)

	if taskRunner.isErrorsLimitExceeded() {
		return ErrErrorsLimitExceeded
	}

	return nil
}

func (taskRunner *TaskRunner) pushTasksToChannel() <-chan Task {
	taskChannel := make(chan Task)
	go func() {
		defer close(taskChannel)

		for _, task := range taskRunner.tasks {
			if taskRunner.isErrorsLimitExceeded() {
				return
			}

			taskChannel <- task
		}
	}()

	return taskChannel
}

func (taskRunner *TaskRunner) runTasks(taskChannel <-chan Task) {
	for i := 0; i < taskRunner.workersCount; i++ {
		taskRunner.wg.Add(1)
		go func() {
			defer taskRunner.wg.Done()

			for task := range taskChannel {
				if taskRunner.isErrorsLimitExceeded() {
					return
				}

				if err := task(); err == nil {
					continue
				}

				taskRunner.increaseErrorCounter()
			}
		}()
	}

	taskRunner.wg.Wait()
}

func (taskRunner *TaskRunner) increaseErrorCounter() {
	atomic.AddInt32(&taskRunner.errorCounter, 1)
}

func (taskRunner *TaskRunner) isErrorsLimitExceeded() bool {
	return atomic.LoadInt32(&taskRunner.errorCounter) >= int32(taskRunner.maxErrorsCount) &&
		taskRunner.maxErrorsCount > 0
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, workersCount, maxErrorsCount int) error {
	taskRunner := createTaskRunner(tasks, workersCount, maxErrorsCount)
	if err := taskRunner.Run(); err != nil {
		return err
	}

	return nil
}

func createTaskRunner(tasks []Task, workersCount int, maxErrorsCount int) TaskRunnerInterface {
	return &TaskRunner{
		wg:             &sync.WaitGroup{},
		tasks:          tasks,
		workersCount:   workersCount,
		maxErrorsCount: maxErrorsCount,
		errorCounter:   0,
	}
}
