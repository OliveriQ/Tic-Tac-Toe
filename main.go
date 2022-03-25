package main

import (
	"fmt"
	"math/rand"
)

const x, o int = 0, 1

var turn int = 0

var win_masks = [8]uint16{
	0b000000111,
	0b000111000,
	0b111000000,

	0b100100100,
	0b010010010,
	0b001001001,

	0b100010001,
	0b001010100,
}

// implement make_move function

func main() {
	var bb [2]uint16
	for is_game_over(bb) == 2 {
		if turn == x {
			fmt.Println("\nYour turn")
			bb = get_player_move(bb)
		} else {
			fmt.Println("\nAI's turn")
			bb = get_computer_move(bb)
		}
		print_board(bb)
		next_turn()
	}

	var result int = is_game_over(bb)
	switch result {
	case 1:
		fmt.Println("X Wins!")
	case -1:
		fmt.Println("O Wins!")
	case 0:
		fmt.Println("Tie!")
	}
}

func get_computer_move(bb [2]uint16) [2]uint16 {
	var move int
	for {
		move = rand.Intn(9)
		if is_valid_move(bb, move) {
			break
		}
	}

	bb[o] = set_bit(bb[o], move-1)
	return bb
}

func get_player_move(bb [2]uint16) [2]uint16 {
	var move int
	for {
		fmt.Print("Enter move (1, 9): ")
		fmt.Scanln(&move)
		if is_valid_move(bb, move) {
			break
		}
	}

	bb[x] = set_bit(bb[x], move-1)
	return bb
}

func is_game_over(bb [2]uint16) int {
	for i := 0; i < 8; i++ {
		if (bb[x] & win_masks[i]) == win_masks[i] {
			return 1
		} else if (bb[o] & win_masks[i]) == win_masks[i] {
			return -1
		}
	}

	if is_board_full(bb) {
		return 0
	}

	return 2
}

func is_board_full(bb [2]uint16) bool {
	for i := 0; i < 9; i++ {
		if get_bit(bb[x], i) == 0 || get_bit(bb[o], i) == 0 {
			return false
		}
	}
	return true
}

func is_valid_move(bb [2]uint16, move int) bool {
	if move <= 9 && move >= 1 {
		if is_occupied(bb, move-1) == false {
			return true
		}
	}
	return false
}

func set_bit(bb uint16, index int) uint16 {
	return bb | (1 << (8 - index))
}

func get_bit(bb uint16, index int) int {
	if (bb & (1 << index)) > 0 {
		return 1
	}
	return 0
}

func get_occupied(bb [2]uint16) uint16 {
	return bb[x] | bb[o]
}

func is_occupied(bb [2]uint16, index int) bool {
	occ := get_occupied(bb)
	if (occ & (1 << (8 - index))) > 0 {
		return true
	}
	return false
}

func print_board(bb [2]uint16) {
	for i := 8; i >= 0; i-- {
		bit_x := get_bit(bb[x], i)
		bit_o := get_bit(bb[o], i)
		if bit_x == 1 {
			fmt.Print("X ")
		} else if bit_o == 1 {
			fmt.Print("O ")
		} else {
			fmt.Print("- ")
		}
		if i%3 == 0 {
			fmt.Print("\n")
		}
	}
}

func next_turn() {
	if turn == x {
		turn = o
	} else {
		turn = x
	}
}

/*
board representation:
=> 1 bitboard for naughts and 1 for crosses
=> compute occupied cells by ANDing naughs and crosses bitboards
=> compute empty cells by using occupied cells

*/
