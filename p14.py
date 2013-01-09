def main():
    collatz = [None] * 3000000
    collatz[0] = collatz[1] = 1

    def calc(i):
        if i < 1000000 and collatz[i] is not None:
            c = collatz[i]
        else:
            if i%2 == 0:
                c = 1 + calc(i//2)
            else:
                c = 1 + calc(i*3+1)
        if i < 1000000:
            collatz[i] = c
        return c

    print(max((calc(i), i) for i in range(2, 1000000)))

main()
