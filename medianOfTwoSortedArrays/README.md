# Median of Two Sorted Arrays

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

```
nums1 = [1, 3]
nums2 = [2]

The median is 2.0
```

Example 2:

```
nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
```

## Solution

The median is the middle of the numbers. For example:

```
3, 5, 9, 11, 200
Median is 9
```

If the length of the numbers is even. Then adding two numbers in the middle and divide by two is median. For example:

```
3, 5, 9, 11, 51, 200
Median is (9 + 11) / 2 = 10
```

In this question. We have two sorted arrays of numbers. We can combine them together and find the median in the combined array.

```
nums1 = [1, 2, 5]
         ^
nums2 = [3, 4]
         ^

1 < 3, put 1 into the array and slice nums1.

arr = [1]
```

```
nums1 = [2, 5]
         ^
nums2 = [3, 4]
         ^

2 < 3, put 2 into the array and slice nums1.

arr = [1, 2]
```

```
nums1 = [5]
         ^
nums2 = [3, 4]
         ^

3 < 5, put 3 into the array and slice nums2.

arr = [1, 2, 3]
```

```
nums1 = [5]
         ^
nums2 = [4]
         ^

4 < 5, put 4 into the array and slice nums2.

arr = [1, 2, 3, 4]
```

```
nums1 = [5]
         ^
nums2 = []

The nums2 is empty. Put remaining element into the array.

arr = [1, 2, 3, 4, 5]
```

Then find the median:

```
index = length of array / 2 + 1 = 5 / 2 + 1 = int 3
result: arr[3] = 3
```
