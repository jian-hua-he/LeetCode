# Regular Expression Matching

Given an input string (`s`) and a pattern (`p`), implement regular expression matching with support for `'.'` and `'*'`.

```
'.' Matches any single character.
'*' Matches zero or more of the preceding element.
```

The matching should cover the **entire** input string (not partial).

Note:

- `s` could be empty and contains only lowercase letters `a-z`.
- `p` could be empty and contains only lowercase letters `a-z`, and characters like `.` or `*`.

Example 1:

```
Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
```

Example 2:

```
Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
```

Example 3:

```
Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
```

Example 4:

```
Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
```

Example 5:

```
Input:
s = "mississippi"
p = "mis*is*p*."
Output: false
```

# Solution 1

We can breakdown the issue. First, matches string without `*` character. It only needs to check the character is the same as the pattern or the character is equal `.`.

```
func isMatch(s string, p string) bool {
    if p == "" {
        return s == ""
    }

    firstMatch := s != "" && (s[0] == p[0] || string(p[0]) == ".")

    return firstMatch && isMatch(s[1:], p[1:])
}
```

Then handle `*` character. The length of the pattern should bigger than 2. Also, the second character should equal `*`. If the condition is matched then split string and pattern to two-part and continue matching.

The first part is to match `*`. So the first character of the string should match the first character of the pattern. And remove the first character continues the match with `*`. The second part matches the whole without `*` character. So remove the first two character in the pattern and continue matching.

```
func isMatch(s string, p string) bool {
    if p == "" {
        return s == ""
    }

    firstMatch := s != "" && (s[0] == p[0] || string(p[0]) == ".")

    // firstMatch && isMatch(s[1:], p) continue match with *
    // isMatch(s, p[2:]) skip * pattern and match again
    if len(p) >= 2 && string(p[1]) == "*" {
        return firstMatch && isMatch(s[1:], p) || isMatch(s, p[2:])
    }

    return firstMatch && isMatch(s[1:], p[1:])
}
```
