ints :: Char -> Int
ints c = read [c]

n :: Integer
n = 2 ^ 1000

main = putStrLn $ show $ sum $ map ints $ show n
