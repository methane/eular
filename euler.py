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

def fibs(a=1, b=1):
    while True:
        yield a
        a, b = b, a+b
