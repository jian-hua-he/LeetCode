# Two Sum

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

## Solution 1

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

## Solution 2

In the previous solutions. The time complexity is `O(n^2)`. We can improve it by sorting the data and give two pointers from start and end.

```
target: 11
origin: 4, 6, 12, 5, 2
sorted: 2, 4, 5, 6, 12
{index: 4, data: 2}
{index: 0, data: 4}
{index: 3, data: 5}
{index: 1, data: 6}
{index: 2, data: 12}
```

```
2 4 5 6 12
i       j

11 < 2+12

Move j to previous
```

```
2 4 5 6 12
i     j

11 < 2+6

Move i to next
```

```
2 4 5 6 12
  i   j

11 < 4+6

Move i to next
```

```
2 4 5 6 12
    i j

11 == 5+6

Return index of data[i] and data[j]
```
