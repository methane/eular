

main = do
    contents <- getContents
    let
        ls = lines contents
        triangle = [
            [(read w :: Int) | w <- words l]
            | l <- ls]
    in
        putStrLn triangle
