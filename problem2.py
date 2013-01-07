import itertools

def fibs(x=1, y=1):
    while True:
        yield x
        x, y = y, x+y

def main():
    fibs4m = itertools.takewhile(lambda x: x < 4000000, fibs(1, 2))
    print(sum(x for x in fibs4m if x % 2 == 0))

main()
