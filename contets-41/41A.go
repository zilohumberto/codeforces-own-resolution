package main

import "fmt"

var s,p,r string
func main(){
    fmt.Scanf("%s\n%s", &s, &p)
    ls := len(s)
    response:= "NO"
    if ls == len(p) {
        for i:=0; i < len(s); i++{
            r = string(s[i]) + r    
        }
        if (r==p){
            response = "YES"
        }
    }
    fmt.Println(response)
}
