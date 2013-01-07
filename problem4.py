def isParindrome(n):
    s = str(n)
    return s == s[::-1]

def largestParindrome(x):
    largest = 0
    for y in range(x, 0, -1):
        for z in range(y, 0, -1):
            n = y * z
            if n > largest and isParindrome(n):
                largest = n
    return largest

print(largestParindrome(999))
