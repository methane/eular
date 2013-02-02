s = 1
i = 1
n = 2

#while n < 5:
while n < 1001:
    s += i*4 + n*10
    i += n*4
    n += 2

print(s)
