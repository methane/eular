import sys
from pprint import pprint

def get_max_path(triangle):
    n = len(triangle)
    points = triangle[0]

    for y in range(n-1):
        #print(points)
        #print(triangle[y+1])
        nps = [0] * (y+2)
        for x in range(y+1):
            p = points[x]
            np = triangle[y+1][x] + p
            if nps[x] < np:
                nps[x] = np
            nps[x+1] = triangle[y+1][x+1] + p
        points = nps

    return max(points)


def main():
    lines = sys.stdin.readlines()
    triangle = [[int(x, 10) for x in line.split()] for line in lines]
    print(get_max_path(triangle))

main()
