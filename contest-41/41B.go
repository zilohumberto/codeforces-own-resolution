package main

import "fmt"

var n,b int
var dollars [2001]int

func main(){
    fmt.Scanf("%d %d\n", &n, &b)
    for i:=0; i<n; i++{
        fmt.Scanf("%d", &dollars[i])
    }
    max_profit := b
    for i:=0; i<n; i++{
        for j:=i+1; j<n; j++{
            profit:=get_profit(i,j, b)
            if profit > max_profit{
                max_profit = profit
            }
        }
    }    
    fmt.Println(max_profit)
}


func get_profit(i int, j int, b int) int {
    amount := b % dollars[i]
    to_buy := (b-amount) / dollars[i]
    b -= to_buy * dollars[i]
    b += to_buy * dollars[j]
    return b
}
