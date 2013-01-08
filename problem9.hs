productOfTriples a b =
    let
        c = 1000 - a - b
        ok = a < b && b < c && a^2 + b^2 == c^2
    in
        if ok
        then a*b*c
        else if b+1 >= 999-a-b
            then productOfTriples (a+1) (a+2)
            else productOfTriples a (b+1)

main = putStrLn $ show $ productOfTriples 1 2
