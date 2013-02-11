def main():
    i = 0
    while True:
        i += 1
        if set(str(i)) == set(str(i*2)) == set(str(i*3)) == \
           set(str(i*4)) == set(str(i*5)) == set(str(i*6)):
               print(i)
               break
main()
