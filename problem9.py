#09:34:34
def main():
    for a in range(1, 333):
        for b in range(a+1, (1000-a)//2):
            c = 1000 - a - b
            if not (a < b < c):
                break
            if a**2 + b**2 == c**2:
                print(a*b*c)
                return
main()
#09:37:31
