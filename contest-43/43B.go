package main

import (
    "fmt"
    "bufio"
    "os"
)
func main(){
    reader := bufio.NewReader(os.Stdin)
    w, _ := reader.ReadString('\n')
    l, _ := reader.ReadString('\n')
    counter_w := make(map[byte]int)
    counter_l := make(map[byte]int)
    for i:=0; i<len(w); i++{
        if w[i] == ' '{
            continue
        }
        counter_w[w[i]] += 1
    }
    for i:=0; i<len(l); i++{
        if l[i] == ' '{
            continue
        }
        counter_l[l[i]] += 1
    }
    result := "YES"
    for letter, count := range counter_l{
        value, ok := counter_w[letter]
        if ok==false || value < count{
            result = "NO"
            break
        }
    }
    fmt.Println(result)
}
