import euler


def main():
    primes = set([p for (_, p) in zip(range(10000), euler.primes())])
    max_n = 0
    aa = bb = 0
    for a in range(-999, 1000):
        for b in range(-999, 1000):
            n = 0
            while n*n + n*a + b in primes:
                n += 1
            if n > max_n:
                max_n = n
                aa = a
                bb = b
    print(aa, bb, aa*bb)

main()
