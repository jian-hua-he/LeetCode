# Two Sum

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

## Solution

We can iterate the numbers and minus by the target number.

```
target = 9
nums = 2 7 11 15

2 7 11 15
i

n = target - nums[i] = 7

2 7 11 15
j

data[j] != 7, index j go to next.

2 7 11 15
  j

data[j] == 7, return i, j
```
