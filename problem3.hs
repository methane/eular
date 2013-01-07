primes = 2:filter f [3,5..]
    where f n = all ((/= 0) . (mod n)) (takeWhile ((<= n) . (^ 2)) primes)

largestPrime n = checkNext n primes
    where
        checkNext n (p:ps) = if n == p
            then n
            else if n `mod` p == 0
                then checkNext (n `div` p) (p:ps)
                else checkNext n ps

main = putStrLn $ show $ largestPrime 600851475143
