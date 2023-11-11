package main

import (
	"fmt"
	"os"
)

type board [][]string

func main() {
	var s int
	fmt.Println("Chose One option \n 1. Start new game \n 2. Exit")
	switch fmt.Scanln(&s); s {
	case 1:
		play()
	case 2:
		os.Exit(0)
	}

}

func play() {
	fmt.Println("The game begin!")
	b := createNewBoard()
	for i := 0; b.checkGame(i); i++ {
		fmt.Println("Turn:", i+1)
		if i%2 == 0 {
			fmt.Println("It's X turn!")
			b.move("X")
		} else {
			fmt.Println("It's O turn!")
			b.move("O")
		}
		b.printTable()
	}
}

func createNewBoard() board {
	return board{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"}}
}

func (b board) printTable() {
	fmt.Println(b[0])
	fmt.Println(b[1])
	fmt.Println(b[2])
}

func (b *board) move(p string) {
	var r, c int
	fmt.Println("Choose one row:")
	for _, err := fmt.Scanln(&r); checkMove(err, r); {
		_, err = fmt.Scanln(&r)
	}
	fmt.Println("Now choose one column:")
	for _, err := fmt.Scanln(&c); checkMove(err, c); {
		_, err = fmt.Scanln(&c)
	}
	(*b)[r][c] = p
}

func (b board) checkGame(i int) bool {
	if i == 10 {
		fmt.Println("It's a draw!")
		return false
	}
	if ((b[0][0] == b[1][1] && b[1][1] == b[2][2]) || (b[0][2] == b[1][1] && b[1][1] == b[2][0])) && b[1][1] != "_" {
		if b[1][1] == "X" {
			fmt.Println("X Win!")
		} else if b[1][1] == "O" {
			fmt.Println("O Win")
		}
		return false
	}
	return true
}

func checkMove(err error, n int) bool {
	if err != nil || n > 2 {
		fmt.Println(err)
		fmt.Println("The move is invalid. Please chose a number between 0 and 2")
		return true
	}
	return false
}
