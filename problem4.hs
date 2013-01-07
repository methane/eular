isParindrome n =
    let digits = show n
    in digits == (reverse digits)

largestParindromeX x =
    let parindromes = [x * y | y <- [x,x-1..1], isParindrome (x*y)]
    in if null parindromes then 0 else head parindromes

largestParindrome x =
    let parindromes = [largestParindromeX y | y <- [x,x-1..1]]
    in maximum parindromes

main = putStrLn $ show $ largestParindrome 999
