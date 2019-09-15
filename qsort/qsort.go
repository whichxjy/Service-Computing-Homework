package qsort

import (
    "math/rand"
    "time"
)

type Sortable interface {
    // Len(data) => the length of data
    Len() int
    // Less(i, j) => check if data[i] is less than data[j]
    Less(int, int) bool
    // Swap(i, j) => swap data[i] and data[j]
    Swap(int, int)
}

func Sort(data Sortable) {
    qsort(data, 0, data.Len() - 1)
}

func qsort(data Sortable, left int, right int) {
    if left < right {
        // pivot index
        pi := partition(data, left, right)
        qsort(data, left, pi - 1)
        qsort(data, pi + 1, right)
    }
}

// Returns an int in [min, max)
func randomInt(min int, max int) int {
    rand.Seed(time.Now().UnixNano())
    return min + rand.Intn(max - min)
}

func partition(data Sortable, left int, right int) int {
    // data[pivot] => element to be placed at right position
    randIndex := randomInt(left, right + 1)
    data.Swap(left, randIndex)
    pivot := left
    lower := left + 1
    upper := right
    for {
        for lower < right && data.Less(lower, pivot) {
            lower += 1
        }
        for upper > left && data.Less(pivot, upper) {
            upper -= 1
        }
        if lower < upper {
            data.Swap(lower, upper)
            lower += 1
            upper -= 1
        } else {
            break
        }
    }
    // data[upper] is always smaller than pivot
    data.Swap(left, upper)
    return upper
}