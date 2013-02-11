import math

def main():
    facts = [math.factorial(i) for i in range(101)]
    cnt = 0
    for n in range(1, 101):
        for r in range(2, n):
            ncr = facts[n] // (facts[r] * facts[n-r])
            if ncr > 1000000:
                cnt += 1
    print(cnt)

main()
