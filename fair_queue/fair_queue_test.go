package fairqueue

import "testing"

func TestGetMinFlow(t *testing.T) {
	tests_flows := []Flows{
		{1, 3, 4, 5, 6},
		{2, 10, 50, 1},
		{2, 2, 2, 2, 2},
		{10, 10, 9, 10, 10},
	}
	indexes := []int{0, 3, 0, 2}
	for i, flow := range tests_flows {
		t.Run("test", func(t *testing.T) {
			_, index := flow.GetMinFlow()
			want := indexes[i]
			if index != want {
				t.Errorf("got index: %d, but want index %d", index, want)
			}
		})
	}
}
func TestQueueTime(t *testing.T) {
	got := QueueTime([]int{2, 3, 4, 5}, 3)
	want := 7

	if got != want {
		t.Errorf("have got %d, but want %d", got, want)
	}
}
