package main

import (
	"fmt"
	"time"
)


type Job struct {
	Name     string
	Duration int
	Deadline int
}


func IterativeOptimization(jobs []Job) []Job {
	for i := 0; i < len(jobs); i++ {
		for j := 0; j < len(jobs)-i-1; j++ {
			if jobs[j].Deadline > jobs[j+1].Deadline {
				jobs[j], jobs[j+1] = jobs[j+1], jobs[j]
			}
		}
	}
	return jobs
}

func RecursiveOptimization(jobs []Job) []Job {
	if len(jobs) <= 1 {
		return jobs
	}
	

	mid := len(jobs) / 2
	left := RecursiveOptimization(jobs[:mid])
	right := RecursiveOptimization(jobs[mid:])


	return merge(left, right)
}


func merge(left, right []Job) []Job {
	result := []Job{}
	for len(left) > 0 && len(right) > 0 {
		if left[0].Deadline <= right[0].Deadline {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	return result
}

func main() {
	sizes := []int{1, 10, 20, 100, 1000, 10000}
	for _, size := range sizes {
		jobs := generateJobs(size)
		fmt.Printf("\nSize: %d\n", size)

		iterativeDuration := measureDuration(func() {
			for i := 0; i < 1000; i++ { 
				IterativeOptimization(jobs)
			}
		}) / 1000
		fmt.Printf("Iterative Duration: %v\n", iterativeDuration)


		recursiveDuration := measureDuration(func() {
			for i := 0; i < 1000; i++ { 
				RecursiveOptimization(jobs)
			}
		}) / 1000
		fmt.Printf("Recursive Duration: %v\n", recursiveDuration)
	}
}


func measureDuration(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}


func generateJobs(n int) []Job {
	jobs := make([]Job, n)
	for i := 0; i < n; i++ {
		jobs[i] = Job{
			Name:     fmt.Sprintf("Job-%d", i+1),
			Duration: i + 1,
			Deadline: n - i,
		}
	}
	return jobs
}
