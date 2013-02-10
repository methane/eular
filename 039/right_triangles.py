from collections import Counter

def main():
    cnt = Counter()
    for a in range(1, 500):
        for b in range(a, 500):
            cq = a**2 + b**2
            c = int(cq ** 0.5)
            if c ** 2 == cq:
                if a == b:
                    cnt[a + b + c] += 1
                else:
                    cnt[a + b + c] += 2
    print(cnt.most_common(3))

main()
