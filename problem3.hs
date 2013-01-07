primes = 2:filter f [3,5..]
    where f n = all ((/= 0) . (mod n)) (takeWhile ((<= n) . (^ 2)) primes)

largestPrime n = 
    let
        checkNext n (p:ps) = if n == p
            then n
            else if n `mod` p == 0
                then checkNext (n `div` p) (p:ps)
                else checkNext n ps
    in
        checkNext n primes

main = putStrLn $ show $ largestPrime 600851475143
