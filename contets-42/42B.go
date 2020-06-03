package main

import (
    "fmt"
    "strconv"
)

type pos struct{
    x int
    y int
}

var whites = [9][9]int{}

func mark_rock(board *[9][9]int, pos_r pos){
    for i:=pos_r.x+1; i<=8; i++{
        board[i][pos_r.y] += 1
        if whites[i][pos_r.y] != 0{
            break
        }
    }
    for i:=pos_r.x-1; i>=1; i--{
        board[i][pos_r.y] += 1
        if whites[i][pos_r.y] != 0{
            break
        }
    }
    for j:=pos_r.y+1; j<=8; j++{
        board[pos_r.x][j] += 1
        if whites[pos_r.x][j] != 0{
            break
        }
    }
    for j:=pos_r.y-1; j>=1; j--{
        board[pos_r.x][j] += 1
        if whites[pos_r.x][j] != 0{
            break
        }
    }
}

func is_valid(x int, y int) bool{
    if x>=1 && x<=8{
        return y>=1 && y<=8
    }
    return false
}

var move_x = [3]int{-1, 0, 1}
var move_y  = [3]int{-1, 0, 1}

func mark_king(board *[9][9]int, pos_k pos){
    for i:=0; i<3; i++{
        for j:=0; j<3; j++{
            x:= pos_k.x+move_x[i]
            y:= pos_k.y+move_y[j]
            if !is_valid(x,y){
                continue
            }
            board[x][y] += 1
        }
    }
}

func Result(wr1 pos, wr2 pos, bk pos, board [9][9]int) string{
    for i:= 0; i<3; i++{
        for j:=0; j<3; j++{
            x:= bk.x + move_x[i]
            y:= bk.y + move_y[j]
            if x==0 && y==0{
                continue
            }
            if !is_valid(x,y){
                continue
            }
            if board[x][y] == 0{
                return "OTHER"
            }
            if board[x][y] == 1 && whites[x][y] == 1{
                return "OTHER"
            } 
        }
    }
    return "CHECKMATE"
}

func main(){
    chess := map[string]pos{}
    toi := [8]byte{'a','b', 'c', 'd', 'e', 'f', 'g', 'h'}
    for k, letter := range toi{
        for j:=1; j<=8; j++{
            chess[string(letter) + strconv.Itoa(j)] = pos{k+1, j}
            whites[k+1][j] = 0
        }
    }
    board := [9][9]int{}
    var wr, wr2, wk, bk string
    fmt.Scanf("%s %s %s %s", &wr, &wr2, &wk, &bk)
    pos_wr, _ := chess[wr]
    board[pos_wr.x][pos_wr.y] = 1 //rock
    whites[pos_wr.x][pos_wr.y] = 1 //rock
    pos_wr2, _ := chess[wr2]
    board[pos_wr2.x][pos_wr2.y] = 1 //rock
    whites[pos_wr2.x][pos_wr2.y] = 1 //rock
    pos_wking, _ := chess[wk]
    board[pos_wking.x][pos_wking.y] = 1 //white king
    whites[pos_wking.x][pos_wking.y] = 2  //white king
    pos_bking, _ := chess[bk]

    mark_rock(&board, pos_wr)
    mark_rock(&board, pos_wr2)
    mark_king(&board, pos_wking)
    fmt.Println(Result(pos_wr, pos_wr2, pos_bking, board))
}

