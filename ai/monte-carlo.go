package ai

import (
	"fmt"
	"math/rand"

	"github.com/TonyChouteau/elrict3/engine"
)

/*
Enum color value
*/
const (
	NONE   int = 0
	CROSS  int = 1
	CIRCLE int = 2
)

func testMove(board engine.Matrix3x3, move int, color int) float64 {

	board, status := engine.Play(board, move, color)
	if status == color {
		return 1
	}
	if len(engine.ListLegal(board)) == 0 {
		return 0.5
	}

	backupBoard := engine.CopyBoard(board)

	var nbOfWins float64
	n := 5000

	for i := 0; i < n; i++ {
		testBoard := engine.CopyBoard(backupBoard)
		currentColor := engine.NextColor(color)

		p:=0
		finished := false
		for !finished {
			legalMoves := engine.ListLegal(testBoard)
			testBoard, status = engine.Play(testBoard, legalMoves[rand.Intn(len(legalMoves))], currentColor)
			if status == color {
				nbOfWins+=1
				finished = true
			} else if status == engine.NextColor(color) {
				if p==0 {
					nbOfWins-=float64(n)
				} else if {
					nbOfWins -= 1;	
				}
				finished = true
			} else if status == 3 {
				nbOfWins += 0.5
				finished = true
			}
			currentColor = engine.NextColor(currentColor)
			p++
		}
	}

	//fmt.Println(nbOfWins, n)
	probWins := nbOfWins / float64(n)
	return probWins
}

/*
GetBestMove : return best move
*/
func GetBestMove(inlineBoard string) int {

	// Decode URI board into a 3x3 Matrix
	if len(inlineBoard) == 9 {
		board := engine.Matrix3x3{}
		for i := range inlineBoard {
			board[i/3][i%3] = int(inlineBoard[i]) - 48
		}
		fmt.Println(board)

		l := engine.ListLegal(board)
		if len(l) == 0 {
			return -1
		}
		results := []float64{}
		for i := range l {
			results = append(results, testMove(board, l[i], CIRCLE))
		}
		maxIndex := 0
		for i := range results {
			if results[i] > results[maxIndex] {
				maxIndex = i
			}
		}
		fmt.Println(results)
		return l[maxIndex]
	}

	fmt.Println("ERROR : board URI must have 9 values.")
	return -1
}
