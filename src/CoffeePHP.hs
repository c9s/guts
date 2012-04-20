module CoffeePHP
import Text.ParserCombinators.Parsec hiding (spaces)

coffeeMain :: IO ()
coffeeMain = do args <- getArgs
    putStrLn "Test"
