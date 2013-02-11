def main():
    ms = 0
    for a in range(1, 100):
        for b in range(1, 100):
            c = a ** b
            ms = max(ms, sum(map(int, str(c))))
    print(ms)
main()
