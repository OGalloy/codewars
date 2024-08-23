//https://www.codewars.com/kata/57b06f90e298a7b53d000a86/train/go

package fairqueue

import "fmt"

/*There is a queue for the self-checkout tills
at the supermarket. Your task is write a function to calculate the total time
required for all the customers to check out!
input
    customers: an array of positive integers representing the queue. Each integer represents a customer, and its value is the amount of time they require to check out.
    n: a positive integer, the number of checkout tills.

output
	The function should return an integer, the total time required.*/

type Flows []int

func NewFlows(num int) Flows {
	return make(Flows, num)
}

func (f *Flows) GetMinFlow() (int, int) {
	var minIndex int = -1
	var minValue int = int(^uint(0) >> 1)
	for i, flow := range *f {
		if flow == 0 {
			minIndex = i
			minValue = flow
			return minValue, minIndex
		}
		if flow < minValue {
			minIndex = i
			minValue = flow
		}
	}
	return minValue, minIndex
}

func QueueTime(customers []int, n int) int {
	//create n flows.
	flows := make(Flows, n)

	for _, time := range customers {
		_, minIndex := flows.GetMinFlow()
		flows[minIndex] += time
		fmt.Println(flows)
	}
	var maxFlow int
	for _, flow := range flows {
		if flow > maxFlow {
			maxFlow = flow
		}
	}
	return maxFlow
}
