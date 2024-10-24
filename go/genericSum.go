package main

import "fmt"

func main() {
	intAvg := Avg[int]{}
	intAvg.Add(4).Add(3).Add(2)
	fmt.Println(intAvg.Val()) // 3

	floatAvg := Avg[float64]{}
	floatAvg.Add(4.0).Add(3.0)
	floatAvg.Add(2.0).Add(1.0)
	fmt.Println(floatAvg.Val()) // 2.5
}

type Avg[V int | float64] struct {
	values []V
}

func (a *Avg[V]) Add(val V) *Avg[V] {
	a.values = append(a.values, val)
	return a
}

func (a *Avg[V]) Val() V {
	sum := V(0)
	if a.values == nil {
		return sum
	}
	for _, val := range a.values {
		sum += val
	}
	return sum / V(len(a.values))
}
