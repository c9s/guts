
module Main where 
import System.Environment
import Text.ParserCombinators.Parsec hiding (spaces)

-- import CoffeePHP

main :: IO ()
main = do
    args <- getArgs
    putStrLn "Test"
-- main = coffeeMain
