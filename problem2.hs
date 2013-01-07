fibs' :: Integer -> Integer -> [Integer]
fibs' x y = x + y : fibs' y (x+y)
fibs = [1, 2] ++ fibs' 1 2

main = putStrLn $ show $ sum $ filter (even) $ takeWhile (\x -> x < 4000000) fibs
