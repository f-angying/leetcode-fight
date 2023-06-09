package main

import "fmt"

type FrequencyTracker struct {
	numCount       map[int]int
	frequencyCount map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{
		numCount:       make(map[int]int),
		frequencyCount: make(map[int]int),
	}
}

func (this *FrequencyTracker) Add(number int) {
	oldCount := this.numCount[number]
	this.numCount[number]++

	if oldCount > 0 {
		this.frequencyCount[oldCount]--
		if this.frequencyCount[oldCount] == 0 {
			delete(this.frequencyCount, oldCount)
		}
	}

	this.frequencyCount[oldCount+1]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if this.numCount[number] == 0 {
		return
	}

	oldCount := this.numCount[number]
	this.numCount[number]--

	this.frequencyCount[oldCount]--
	if this.frequencyCount[oldCount] == 0 {
		delete(this.frequencyCount, oldCount)
	}
	this.frequencyCount[oldCount-1]++
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	return this.frequencyCount[frequency] > 0
}

/**
 * Your FrequencyTracker object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(number);
 * obj.DeleteOne(number);
 * param_3 := obj.HasFrequency(frequency);
 */
func main() {
	frequencyTracker := Constructor()
	frequencyTracker.Add(3)
	frequencyTracker.Add(3)
	fmt.Println(frequencyTracker.HasFrequency(2)) // Expected: true

	frequencyTracker2 := Constructor()
	frequencyTracker2.Add(1)
	frequencyTracker2.DeleteOne(1)
	fmt.Println(frequencyTracker2.HasFrequency(1)) // Expected: false

	frequencyTracker3 := Constructor()
	fmt.Println(frequencyTracker3.HasFrequency(2)) // Expected: false
	frequencyTracker3.Add(3)
	fmt.Println(frequencyTracker3.HasFrequency(1)) // Expected: true
}
