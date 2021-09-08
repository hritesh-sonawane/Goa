package main

import (
	"fmt"
	"sort"
)

// O(nlog(n)) time | O(n) space | n: number of tasks
func TaskAssignment(k int, tasks []int) [][]int {
	pairedTasks := make([][]int, 0)
	taskDurationsToIndices := getTaskDurationToIndices(tasks)

	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i] < tasks[j]
	})

	var task1Index, task2Index int
	for idx := 0; idx < k; idx++ {
		task1Duration := tasks[idx]
		indicesWithTask1Duration := taskDurationsToIndices[task1Duration]

		task1Index, taskDurationsToIndices[task1Duration] = 
			indicesWithTask1Duration[len(indicesWithTask1Duration)-1],
			indicesWithTask1Duration[:len(indicesWithTask1Duration)-1]

		task2SortedIndex := len(tasks) - 1 - idx
		task2Duration := tasks[task2SortedIndex]
		
		indicesWithTask2Duration := taskDurationsToIndices[task2Duration]
		task2Index, taskDurationsToIndices[task2Duration] = 
			indicesWithTask2Duration[len(indicesWithTask2Duration)-1],
			indicesWithTask2Duration[:len(indicesWithTask2Duration)-1]

		pairedTasks = append(pairedTasks, []int{task1Index, task2Index})
	}
	return pairedTasks
}

func getTaskDurationToIndices(tasks []int) map[int][]int {
	taskDurationsToIndices := map[int][]int{}

	for idx := range tasks {
		taskDuration := tasks[idx]
		taskDurationsToIndices[taskDuration] = append(taskDurationsToIndices[taskDuration], idx)
	}
	return taskDurationsToIndices
}

func main() {
	tasks := []int{1, 3, 5, 3, 1, 4}
	fmt.Println(TaskAssignment(3, tasks))
}