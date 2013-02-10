def nth(r, k, n):
    o = n // k
    p = k - n % k - 1
    return (r[o] // (10**p)) % 10

def test_neth():
    assert nth(range(1, 10), 1, 0) == 1
    assert nth(range(1, 10), 1, 1) == 2
    assert nth(range(10, 100), 2, 0) == 1
    assert nth(range(10, 100), 2, 1) == 0
    assert nth(range(10, 100), 2, 2) == 1
    assert nth(range(10, 100), 2, 3) == 1

def main():
    res = 1
    k = 1
    p = 1
    d = 1
    while True:
        R = range(10**(k-1), 10**k)
        np = p + k * len(R)
        print("np=", np)
        while d < np:
            n = nth(R, k, d-p)
            res *= n
            print("d=", d, "n=", n)
            d *= 10
            if d > 1000000:
                return res
        p = np
        k += 1

print(main())
