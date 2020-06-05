package main

import "fmt"

func main(){
    var n int
    var s string
    goals := make(map[string]int)
    fmt.Scanf("%d\n", &n)
    for i:=0; i<n; i++{
        fmt.Scanf("%s\n", &s)
        goals[s] ++
    }
    won_goals  := 0
    won_key := ""
    for key, amount := range goals{
        if amount > won_goals{
            won_goals = amount
            won_key = key
        }
    }
    fmt.Println(won_key)
}
