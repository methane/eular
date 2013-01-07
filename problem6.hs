main = putStrLn $ show $ ((sum [1 .. 100]) ^ 2 - (sum [i^2 | i <- [1..100]]))
