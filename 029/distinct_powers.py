def main():
    distinct_numbers = set()
    for a in range(2, 101):
        for b in range(2, 101):
            distinct_numbers.add(a**b)
    print(len(distinct_numbers))

main()
