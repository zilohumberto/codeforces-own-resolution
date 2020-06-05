package main

import "fmt"

var n, m int
func main(){
    fmt.Scanf("%d %d", &n, &m)
    r := 0
    if n % 2 == 1 && m%2 == 1 || ((m == 1 || n == 1) && (m>2 || n>2)) {
        r = 1
    }
    fmt.Println(r)
    if r == 1 {
        fmt.Printf("%d %d 1 1\n", n, m)
        fmt.Printf("1 1\n")
        pathToEnd(n, m)
    }else{
        fmt.Println("1 1")
        pathToReturn(n, m)
    }
    fmt.Println("1 1")
}

func pathToEnd(x int, y int){
    b := 2
    for i:=1; i<=x; i+=2{
        for j:=b; j<=y; j++{
            fmt.Println(i, j)
        }
        if i+1>x{
            continue
        }
        for j:=y; j>0; j--{
            fmt.Println(i+1, j)
        }
        if b == 2{
            b = 1
        }
    }
}

func pathToReturn(n int, m int){
    x := n
    y := m
    if x % 2 == 1{
        for j:=1; j<=y; j+=2{
            for i:=2; i<=x; i++{
                fmt.Println(i, j)
            }
            if j+1>y{
                continue;
            }
            for i:=x; i>1; i--{
                fmt.Println(i, j+1)
            }
        }
        for j:=y; j>1; j--{
            fmt.Println(1,j)
        }
    }else{
        for i:=1; i<=x; i+=2{
            for j:=2; j<=y; j++{
                fmt.Println(i, j)
            }
            if i+1>x{
                continue
            }
            for j:=y; j>1; j--{
                fmt.Println(i+1, j)  
            }
        }
        for i:=x; i>1; i--{
            fmt.Println(i, 1)
        }
    }
}

