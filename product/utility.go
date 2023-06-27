package product

// Function to sort trainings by name
//
//	sortTrainings sorts a slice of TrainingStruct by name using quicksort.
//
// It has O(nlogn) time complexity.
func sortTrainings(trainings []TrainingStruct) []TrainingStruct {
	// Use quicksort for O(nlogn) time complexity
	quicksort(trainings, 0, len(trainings)-1)
	return trainings
}

// quicksort recursively partitions the slice around a pivot.
func quicksort(trainings []TrainingStruct, left, right int) {
	if left < right {
		// Partition the slice
		pivotIndex := partition(trainings, left, right)

		// Recursively call quicksort on left and right of pivot
		quicksort(trainings, left, pivotIndex-1)
		quicksort(trainings, pivotIndex+1, right)
	}
}

// partition partitions the slice around a pivot and returns the index of the pivot.
func partition(trainings []TrainingStruct, left, right int) int {
	pivot := trainings[right].Name
	i := left - 1

	for j := left; j <= right-1; j++ {
		if trainings[j].Name <= pivot {
			i++
			temp := trainings[i]
			trainings[i] = trainings[j]
			trainings[j] = temp
		}
	}
	temp := trainings[i+1]
	trainings[i+1] = trainings[right]
	trainings[right] = temp
	return i + 1
}
