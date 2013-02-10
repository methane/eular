import Data.Char
import Numeric

isParindrome x = xs == (reverse xs) && bs == (reverse bs)
    where xs = show x
          bs = showIntAtBase 2 intToDigit x ""

main = print $ sum [x | x <- [1..1000000], isParindrome x]
--main = print $ [x | x <- [1..10000], isParindrome x]
