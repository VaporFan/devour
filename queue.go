/*
 * mtStats Devour - Queue
 */

package main

import (
	"container/heap"
)

type Job struct {
	task     string
	value    string
	priority int
	index    int
}

type Queue []*Job

func (q Queue) Len() int {

	return len(q)

}

// Order by priority, highest first.
func (q Queue) Less(i, j int) bool {

	return q[i].priority > q[j].priority

}

func (q Queue) Swap(i, j int) {

	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j

}

// Add Job to Queue
func (q *Queue) Push(x interface{}) {

	n := len(*q)
	job := x.(*Job)
	job.index = n
	*q = append(*q, job)

}

// Remove Job from Queue
func (q *Queue) Pop() interface{} {

	old := *q
	n := len(old)
	job := old[n-1]
	job.index = -1 // for safety
	*q = old[0 : n-1]

	return job

}

// Modifies the Priority and Task of a Job.
func (q *Queue) update(job *Job, task string, value string, priority int) {

	job.task = task
	job.value = value
	job.priority = priority
	heap.Fix(q, job.index)

}
