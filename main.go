package main

import (
	"fmt"
	"strconv"
)

func main() {
	game := newGameSystem(NewChessBoard())
	game.Play()
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func min(x0, x1 int) int {
	if x1 < x0 {
		return x1
	} else {
		return x0
	}
}

func max(x0, x1 int) int {
	if x1 < x0 {
		return x0
	} else {
		return x1
	}
}

type Figure interface {
	ValidateTurn(x0, y0, x1, y1 int, matrix [8][8]Figure) bool
	ValidateAttack(x0, y0, x1, y1 int, matrix [8][8]Figure) bool
	String() string
	GetCode() int
}

type King struct {
	Name string
	Code int
}

func newKing(code int) *King {
	return &King{
		Name: "Ki",
		Code: code,
	}
}

func (k King) String() string {
	return k.Name
}

func (k King) GetCode() int {
	return k.Code
}

func (k King) ValidateTurn(x0, y0, x1, y1 int, matrix [8][8]Figure) bool {
	return abs(x0-x1) <= 1 && abs(y0-y1) <= 1
}

func (k King) ValidateAttack(x0, y0, x1, y1 int, matrix [8][8]Figure) bool {
	return abs(x0-x1) <= 1 && abs(y0-y1) <= 1
}

type Queen struct {
	Name string
	Code int
}

func (k Queen) GetCode() int {
	return k.Code
}

func (q Queen) String() string {
	return q.Name + strconv.Itoa(q.Code)
}

func newQueen(code int) *Queen {
	return &Queen{
		Name: "Q",
		Code: code,
	}
}

func (q Queen) ValidateTurn(x0, y0, x1, y1 int, matrix [8][8]Figure) bool {
	if !((x0-x1)*(y0-y1) == 0 || abs(x1-x0) == abs(y1-y0)) {
		return false
	}
	if abs(x1-x0) == abs(y1-y0) {
		if x0+y0 == x1+y1 {
			if x1 < x0 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0+i] != nil {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0-i] != nil {
						return false
					}
				}
			}
		} else if x0-y0 == x1-y0 {
			if x0 > x1 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0-i] != nil {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0+i] != nil {
						return false
					}
				}
			}
		}
	} else {
		if x1 == x0 {
			for i := min(y0, y1) + 1; i < max(y1, y0); i++ {
				if matrix[x0][i] != nil {
					return false
				}
			}
		} else if y0 == y1 {
			for i := min(x0, x1) + 1; i < max(x0, x1); i++ {
				if matrix[i][y0] != nil {
					return false
				}
			}
		}
	}
	return true
}
func (q Queen) ValidateAttack(x0, y0, x1, y1 int, matrix [8][8]Figure) bool {
	if !((x0-x1)*(y0-y1) == 0 || abs(x1-x0) == abs(y1-y0)) {
		return false
	}
	if abs(x1-x0) == abs(y1-y0) {
		if x0+y0 == x1+y1 {
			if x1 < x0 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0+i] != nil {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0-i] != nil {
						return false
					}
				}
			}
		} else if x0-y0 == x1-y0 {
			if x0 > x1 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0-i] != nil {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0+i] != nil {
						return false
					}
				}
			}
		}
	} else {
		if x1 == x0 {
			for i := min(y0, y1) + 1; i < max(y1, y0); i++ {
				if matrix[x0][i] != nil {
					return false
				}
			}
		} else if y0 == y1 {
			for i := min(x0, x1) + 1; i < max(x0, x1); i++ {
				if matrix[i][y0] != nil {
					return false
				}
			}
		}
	}
	return true
}

type Rook struct {
	Name string
	Code int
}

func (k Rook) GetCode() int {
	return k.Code
}

func (q Rook) String() string {
	return q.Name + strconv.Itoa(q.Code)
}

func newRook(code int) *Rook {
	return &Rook{
		Name: "R",
		Code: code,
	}
}

func (r Rook) ValidateTurn(x0, y0, x1, y1 int, matrix [8][8]Figure) bool {
	if !((x1-x0)*(y1-y0) == 0) {
		return false
	}
	if x1 == x0 {
		for i := min(y0, y1) + 1; i < max(y1, y0); i++ {
			if matrix[x0][i] != nil {
				return false
			}
		}
	} else if y0 == y1 {
		for i := min(x0, x1) + 1; i < max(x0, x1); i++ {
			if matrix[i][y0] != nil {
				return false
			}
		}
	}
	return true
}

func (r Rook) ValidateAttack(x0, y0, x1, y1 int, matrix [8][8]Figure) bool {
	if !((x1-x0)*(y1-y0) == 0) {
		return false
	}
	if x1 == x0 {
		for i := min(y0, y1) + 1; i < max(y1, y0); i++ {
			if matrix[x0][i] != nil {
				return false
			}
		}
	} else if y0 == y1 {
		for i := min(x0, x1) + 1; i < max(x0, x1); i++ {
			if matrix[i][y0] != nil {
				return false
			}
		}
	}
	return true
}

type Bishop struct {
	Name string
	Code int
}

func (k Bishop) GetCode() int {
	return k.Code
}

func (q Bishop) String() string {
	return q.Name + strconv.Itoa(q.Code)
}

func newBishop(code int) *Bishop {
	return &Bishop{
		Name: "B",
		Code: code,
	}
}

func (b Bishop) ValidateTurn(x0, y0, x1, y1 int, matrix [8][8]Figure) bool {
	if !(abs(x1-x0) == abs(y1-y0)) {
		return false
	}
	if abs(x1-x0) == abs(y1-y0) {
		if x0+y0 == x1+y1 {
			if x1 < x0 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0+i] != nil {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0-i] != nil {
						return false
					}
				}
			}
		} else if x0-y0 == x1-y0 {
			if x0 > x1 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0-i] != nil {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0+i] != nil {
						return false
					}
				}
			}
		}
	}
	return true
}

func (b Bishop) ValidateAttack(x0, y0, x1, y1 int, matrix [8][8]Figure) bool {
	if !(abs(x1-x0) == abs(y1-y0)) {
		return false
	}
	if abs(x1-x0) == abs(y1-y0) {
		if x0+y0 == x1+y1 {
			if x1 < x0 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0+i] != nil {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0-i] != nil {
						return false
					}
				}
			}
		} else if x0-y0 == x1-y0 {
			if x0 > x1 {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0-i][y0-i] != nil {
						return false
					}
				}
			} else {
				for i := 1; i < abs(x0-y0); i++ {
					if matrix[x0+i][y0+i] != nil {
						return false
					}
				}
			}
		}
	}
	return true
}

type Knight struct {
	Name string
	Code int
}

func (k Knight) GetCode() int {
	return k.Code
}

func (q Knight) String() string {
	return q.Name + strconv.Itoa(q.Code)
}

func newKnight(code int) *Knight {
	return &Knight{
		Name: "K",
		Code: code,
	}
}

func (k Knight) ValidateTurn(x0, y0, x1, y1 int, matrix [8][8]Figure) bool {
	return abs((x0-x1)*(y0-y1)) == 2
}

func (k Knight) ValidateAttack(x0, y0, x1, y1 int, matrix [8][8]Figure) bool {
	return abs((x0-x1)*(y0-y1)) == 2
}

type Pawn struct {
	Name string
	Code int
}

func (k Pawn) GetCode() int {
	return k.Code
}

func (q Pawn) String() string {
	return q.Name + strconv.Itoa(q.Code)
}

func newPawn(code int) *Pawn {
	return &Pawn{
		Name: "P",
		Code: code,
	}
}

func (p Pawn) ValidateTurn(x0, y0, x1, y1 int, matrix [8][8]Figure) bool {
	if x0 == 1 || x0 == 6 {
		if !(y0 == y1 || abs(x0-x1) > 2) {
			return false
		} else {
			if !(y0 == y1 || abs(x0-x1) > 1) {
				return false
			}
		}
	}
	return true
}

func (p Pawn) ValidateAttack(x0, y0, x1, y1 int, matrix [8][8]Figure) bool {
	return abs(x1-x0) == 1 && abs(y1-y0) == 1
}

type ChessBoard struct {
	ChessMatrix [8][8]Figure
}

func NewChessBoard() *ChessBoard {
	matrix := [8][8]Figure{}
	for i := 0; i < 8; i++ {
		matrix[1][i] = newPawn(1)
		matrix[6][i] = newPawn(2)
	}

	matrix[0][0] = newRook(1)
	matrix[0][7] = newRook(1)

	matrix[7][0] = newRook(2)
	matrix[7][7] = newRook(2)

	matrix[0][1] = newKnight(1)
	matrix[0][6] = newKnight(1)

	matrix[7][1] = newKnight(2)
	matrix[7][6] = newKnight(2)

	matrix[0][2] = newBishop(1)
	matrix[0][5] = newBishop(1)

	matrix[7][2] = newBishop(2)
	matrix[7][5] = newBishop(2)

	matrix[0][3] = newQueen(1)
	matrix[0][4] = newKing(1)

	matrix[7][3] = newQueen(2)
	matrix[7][4] = newKing(2)
	return &ChessBoard{
		ChessMatrix: matrix,
	}
}

func (c *ChessBoard) PrintMatrix() {
	fmt.Println("7 00 11 22 33 44 55 66 77")
	for i, d := range c.ChessMatrix {
		TmpString := strconv.Itoa(i) + " "
		for _, j := range d {
			if j == nil {
				TmpString += "00 "
			} else {
				TmpString += (j.String() + " ")
			}
		}
		fmt.Println(TmpString)
	}
}

func (c *ChessBoard) SelectForTurn(x, y int) bool {
	return !(x < 0 || x > 7 || y < 0 || y > 7 || c.ChessMatrix[x][y] == nil)
}

func (c *ChessBoard) MoveFigure(x0, y0, x, y int) bool {
	if c.SelectForTurn(x0, y0) {
		if c.ChessMatrix[x][y] != nil && c.ChessMatrix[x0][y0].GetCode() == c.ChessMatrix[x][y].GetCode() { // КОД КОД КОД
			fmt.Println("This is an incorrect turn")
			return false
		}
		if c.ChessMatrix[x][y] == nil && c.ChessMatrix[x0][y0].ValidateTurn(x0, y0, x, y, c.ChessMatrix) {
			tmp := c.ChessMatrix[x0][y0]
			c.ChessMatrix[x0][y0] = c.ChessMatrix[x][y]
			c.ChessMatrix[x][y] = tmp
			return true
		} else if c.ChessMatrix[x][y] != nil && c.ChessMatrix[x0][y0].ValidateAttack(x0, y0, x, y, c.ChessMatrix) {
			tmp := c.ChessMatrix[x0][y0]
			c.ChessMatrix[x0][y0] = nil
			c.ChessMatrix[x][y] = tmp
			return true
		}
	}
	fmt.Println("This is an incorrect turn")
	return false
}

type GameSystem struct {
	Chess      *ChessBoard
	Kings      []Figure
	King1Turns [][]int
	King2Turns [][]int
}

func newGameSystem(chess *ChessBoard) *GameSystem {
	return &GameSystem{
		Chess:      chess,
		Kings:      []Figure{chess.ChessMatrix[0][4], chess.ChessMatrix[7][4]},
		King1Turns: [][]int{{0, 4}},
		King2Turns: [][]int{{7, 4}},
	}
}

func (g *GameSystem) TurnInput() (int, int, int, int) {
	var x0, y0 int
	fmt.Println("choose the indexes of the figure you want to move")
	fmt.Scan(&x0, &y0)
	for !(g.Chess.SelectForTurn(x0, y0)) {
		fmt.Println("you cannot choose these indexes, try again")
		fmt.Scan(&x0, &y0)
	}
	var x, y int
	fmt.Println("choose the indexes of the cell where you want to put your figure")
	fmt.Scan(&x, &y)
	for min(x, y) < 0 || max(x, y) > 7 {
		fmt.Println("choose the indexes of the cell where you want to put your figure")
		fmt.Scan(&x, &y)
	}
	return x0, y0, x, y
}

func (g *GameSystem) CheckWin() int {
	x1 := g.King1Turns[len(g.King1Turns)-1][0]
	y1 := g.King1Turns[len(g.King1Turns)-1][1]

	x2 := g.King2Turns[len(g.King2Turns)-1][0]
	y2 := g.King2Turns[len(g.King2Turns)-1][1]

	if g.Chess.ChessMatrix[x1][y1] != g.Kings[0] {
		return 1 // Победил второй
	}
	if g.Chess.ChessMatrix[x2][y2] != g.Kings[1] {
		return 2
	}
	return 0
}

func (g *GameSystem) Play() {
	for {
		g.Chess.PrintMatrix()
		x0, y0, x, y := g.TurnInput()
		if !(g.Chess.MoveFigure(x0, y0, x, y)) {
			continue
		}
		if g.Chess.ChessMatrix[x][y] == g.Kings[0] {
			g.King1Turns = append(g.King1Turns, []int{x, y})
		}
		if g.Chess.ChessMatrix[x][y] == g.Kings[1] {
			g.King2Turns = append(g.King2Turns, []int{x, y})
		}
		if g.CheckWin() == 2 {
			fmt.Println("2 wins")
			return
		} else if g.CheckWin() == 1 {
			fmt.Println("1 wins")
			g.Chess.PrintMatrix()
			return
		}
	}
}
