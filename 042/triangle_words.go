package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func is_triangles(s string, cache *map[int]bool, max_cache int) (bool, int) {
    n := 0
    for _, r := range s {
        n += int(r - 'A' + 1)
    }
    for max_cache < n {
        max_cache++
        t := max_cache * (max_cache+1) / 2
        (*cache)[t] = true
    }
    return (*cache)[n], max_cache
}

func main() {
    cache := make(map[int]bool)
    cache[0] = true
    max_cache := 0
    cnt := 0
    f, e := os.Open("words.txt")
    r := bufio.NewReader(f)
    var ok bool

    for {
        var l string
        l, e = r.ReadString(',')
        l = strings.Trim(l, " \",")
        if len(l) > 0 {
            if ok, max_cache = is_triangles(l, &cache, max_cache); ok {
                cnt++
            }
        }
        if e != nil {
            break
        }
    }
    fmt.Println("count=", cnt)
}
