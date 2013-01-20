import Data.Char

main = putStrLn $ show $ sum $ map digitToInt $ show $ product [1..100]
