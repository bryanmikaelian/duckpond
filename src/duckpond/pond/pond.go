package pond

import (
    "duckpond/duck"
)

func DuckAction(c chan string) {
    c <- duck.Quack()
}
