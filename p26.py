#!/usr/bin/env python3


def recurring_cycle(n):
    L = []
    x = 10
    while True:
        y = x % n
        if y in L:
            return len(L) - L.index(y)
        L.append(y)
        x *= 10

R = [(recurring_cycle(n), n) for n in range(2, 1000)]
R.sort(reverse=True)
print(R[:5])
