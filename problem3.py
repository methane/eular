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

def largestPrime(n):
    for p in primes():
        while n % p == 0:
            n //= p
        if n == 1:
            return p

print(largestPrime(600851475143))
