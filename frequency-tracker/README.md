## Frequency Tracker

### Intuition

`The FrequencyTracker` is designed to keep track of the frequency of numbers added to it. It provides methods to add a number, delete one occurrence of a number, and check if a specific frequency exists in the tracker.

### Approach

`The FrequencyTracker` uses two maps to store the count of numbers and the count of frequencies. When a number is added, its count in the `numCount` map is incremented. If the number had a previous count greater than zero, its frequency count is decremented and removed from the `frequencyCount` map if it reaches zero. The frequency count of the new count is then incremented.

When deleting one occurrence of a number, its count in the `numCount` map is decremented. If the number had a previous count greater than zero, its frequency count is decremented and removed from the `frequencyCount` map if it reaches zero. The frequency count of the new count is then incremented.

To check if a specific frequency exists, we simply check if the frequency exists in the `frequencyCount` map.

### Complexity

#### Time complexity: 

The time complexity of adding and deleting numbers is `O(1)` since map operations are constant time. The time complexity of checking a frequency is also `O(1)` since it involves a map lookup.

#### Space complexity: 

The space complexity is `O(n)`, where `n` is the number of distinct numbers added to the tracker. The `numCount` and `frequencyCount` maps will store counts for all distinct numbers, resulting in a space complexity proportional to the number of distinct numbers.