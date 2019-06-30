# Longest Palindromic Substring

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

```
Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
```

Example 2:

```
Input: "cbbd"
Output: "bb"
```

## Solution 1

The idea of solution 1 is simple. First, we need a function to identify palindrome. The function will check the first character and the last character is equal. If not then stop checking the string and return false. Otherwise, reduce the scope of the string.

```
abcba
^   ^

abcba
 ^ ^

abcba
  ^
  ^
```

Second, We start with the longest length of the given string to find the palindrome. If it is not matched. Reduce the length and check again.

```
Given String: faabcbax
```

```
Current length: 8
Loop 1:

Nested Loop 1:
faabcbax
^      ^
Not matched. Reduce the scope (current length - 1).
```

```
Current length: 7
Loop 2:

Nested Loop 1:
faabcbax
^     ^
Not matched. Go to next.

Nested Loop 2:
faabcbax
 ^     ^
Not matched. Reduce the scope (current length - 1).
```

```
Current length: 6
Loop 3:

Nested Loop 1:
faabcbax
^    ^
Not matched. Go to next.

Nested Loop 2:
faabcbax
 ^    ^
Not matched. Go to next.

Nested Loop 3:
faabcbax
  ^    ^
Not matched. Reduce the scope (current length - 1)

// Loop until finding the result...
```

It is not a good idea. We need a lot of nested loops to find palindrome, which introduces the performance issue.

## Solution 2

This idea comes up from Leetcode. But the description of the idea is not clear for me. I try to figure out in this document. Basically, the idea is trying to iterate the string from the start. And expanding the string try to find the longest palindrome.

```
Given String: faabcbax
```

```
Loop 1:

==========
Expand 1:

faabcbax
r
l

r == l, so continue expanding

 faabcbax
  r
l

l is out of range. Return the positions of the left pointer and the right pointer.

Left 1: 0
Right 1: 0
----------
Expand 2:

faabcbax
 r
l

r != l, Return the positions of the left pointer and the right pointer.

Left 2: 1
Right 2: 0
==========

Longest: 1 (Expand 1 > Expand 2)
Start Index: 0 (Left 1)
End Index: 0 (Right 1)
```

```
Loop 2:

==========
Expand 1:

faabcbax
 r
 l

r == l, so continue expanding.

faabcbax
  r
l

r != l. Return the positions of the left pointer and the right pointer.

Left 1: 1
Right 1: 1
----------
Expand 2:

faabcbax
  r
 l

r == l, so continue expanding.

faabcbax
   r
l

r != l. Return the positions of the left pointer and the right pointer.

Left 1: 1
Right 1: 2
==========

Longest: 2 (Expand 2 > Expand 1)
Because the current longest > previous longest. Change index.
Start Index: 1 (Left 1)
End Index: 2 (Right 1)

// And so on...
```
