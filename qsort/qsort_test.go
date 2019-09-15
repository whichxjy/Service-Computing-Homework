package qsort

import (
    "testing"
    "reflect"
)

type MyInt struct {
    val int
}

type IntArr []MyInt

func (arr IntArr) Len() int {
    return len(arr)
}

func (arr IntArr) Less(i int, j int) bool {
    return arr[i].val < arr[j].val
}

func (arr IntArr) Swap(i int, j int) {
    arr[i], arr[j] = arr[j], arr[i]
}

func TestSort(t *testing.T) {
    cases := []struct {
        in, want IntArr
    } {
        {IntArr{}, IntArr{}},
        {IntArr{{1}}, IntArr{{1}}},
        {IntArr{{3}, {4}, {1}, {2}, {5}}, IntArr{{1}, {2}, {3}, {4}, {5}}},
        {IntArr{{123}, {-1}, {0}, {4}, {-10}}, IntArr{{-10}, {-1}, {0}, {4}, {123}}},
        {IntArr{{10}, {9}, {8}, {7}, {6}, {5}, {4}, {3}, {2}, {1}, {0}},
         IntArr{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10}}},
    }
    for _, c := range cases {
        Sort(c.in)
        if !reflect.DeepEqual(c.in, c.want) {
            t.Errorf("Fail to sort %v\n", c.in)
        }
    }
}