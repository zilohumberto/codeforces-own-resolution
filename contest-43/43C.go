package main

import (
    "fmt"
)

func main(){
    n := 0
    fmt.Scanf("%d\n", &n)
    mods := make([]int, 4)
    aux := 0
    fmt.Println(n)
    for i:=0; i <n; i++{
        fmt.Scanf("%d", &aux)
        fmt.Println(aux)
        mods[aux%3] ++
    }
    x := mods[1]
    if x > mods[2]{
        x = mods[2]
    }
    fmt.Println(int(mods[0]/2) + x)
}
