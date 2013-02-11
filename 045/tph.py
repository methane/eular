def main():
    t = p = h = 0
    tn = pn = hn = 0

    while True:
        tn += 1
        t = tn * (tn + 1) / 2

        while p < t:
            pn += 1
            p = pn * (3 * pn - 1) / 2

        while h < t:
            hn += 1
            h = hn * (2 * hn - 1)

        if t == p == h:
            print(t)

main()
