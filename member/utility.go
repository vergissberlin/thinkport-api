package member

// Function to sort members by name
//
//	sortMembers sorts a slice of MemberStruct by name using quicksort.
//
// It has O(nlogn) time complexity.
func sortMembers(members []MemberStruct) []MemberStruct {
	// Use quicksort for O(nlogn) time complexity
	quicksort(members, 0, len(members)-1)
	return members
}

// quicksort recursively partitions the slice around a pivot.
func quicksort(members []MemberStruct, left, right int) {
	if left < right {
		// Partition the slice
		pivotIndex := partition(members, left, right)

		// Recursively call quicksort on left and right of pivot
		quicksort(members, left, pivotIndex-1)
		quicksort(members, pivotIndex+1, right)
	}
}

// partition partitions the slice around a pivot and returns the index of the pivot.
func partition(members []MemberStruct, left, right int) int {
	pivot := members[right].Name
	i := left - 1

	for j := left; j <= right-1; j++ {
		if members[j].Name <= pivot {
			i++
			temp := members[i]
			members[i] = members[j]
			members[j] = temp
		}
	}
	temp := members[i+1]
	members[i+1] = members[right]
	members[right] = temp
	return i + 1
}
