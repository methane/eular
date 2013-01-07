def primes():
    yield 2
    ps = [2]
    n = 3
    while True:
        for p in ps:
            if n % p == 0:
                break
            if p**2 >= n:
                ps.append(n)
                yield n
                break
        n += 2

def countFactor(n, p):
    c = 0
    while n % p == 0:
        c += 1
        n //= p
    return c

def maxFactors(n):
    for p in primes():
        if p > n:
            break
        yield p, max(countFactor(m, p) for m in range(1, n+1))

def main():
    prod = 1
    for i, j in maxFactors(20):
        prod *= i ** j
    print(prod)

main()
