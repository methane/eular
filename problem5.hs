primes = 2:filter f [3,5..]
    where f n = all ((/= 0) . (mod n)) (takeWhile ((<= n) . (^ 2)) primes)

countFactor n p =
    if n `mod` p == 0 then (1 + countFactor (n `div` p) p) else 0

smallestMultiple n =
        product $ zipWith (^) primesN (map (maxFactors n) primesN)
    where
        -- nまでの素数
        primesN :: [Integer]
        primesN = takeWhile (<n) primes
        -- [1..n] のうち、 p の最大の因数
        maxFactors :: Integer -> Integer -> Integer
        maxFactors n p = maximum [countFactor x p | x <- [1..n]]

main = putStrLn $ show $ smallestMultiple 20
