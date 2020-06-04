package main

import "fmt"

var s string
func main(){
    fmt.Scanf("%s", &s)
    at := false
    dot := false
    result := string(s[0])
    sz := len(s)
    for i:=1; i<sz-1; i++{
        if s[i]=='d' && i+2 != sz-1{
            dot = search(i+1, "ot")
            if dot{
                result += "."
                i+=2
                continue
            }
        }
        if s[i]=='a' && !at && i+1!= sz-1{
            at = search(i+1, "t")
            if at {
                i+=1
                result += "@"
                continue
            }
        }
        result += string(s[i])
    }
    result += string(s[sz-1])
    fmt.Println(result)
 }


 func search(i int, match string) bool{
     for j:=0; j<len(match); j++{
        if i >= len(s){
            return false
        }
        if s[i] != match[j]{
            return false
        }
        i++
     }
     return true
 }
