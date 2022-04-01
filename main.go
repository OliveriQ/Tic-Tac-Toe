package main

import (
	"fmt"
	"math/rand"
)

var win_masks = [8]uint16{
	0b000000111, 0b000111000, 0b111000000,
	0b100100100, 0b010010010, 0b001001001,
	0b100010001, 0b001010100,
}

const (
	infinity int = 99999
	x, o     int = 0, 1
)

func main() {
	var bb [2]uint16
	var turn int = 0
	for is_game_over(bb) == 2 {
		if turn == x {
			fmt.Println("\nYour turn")
			bb = get_player_move(bb)
		} else {
			fmt.Println("\nAI's turn")
			//bb = pick_random_move(bb)
			bb = search_best_move(bb)
		}
		print_board(bb)
		turn = next_turn(turn)
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

func negamax(bb [2]uint16, turn int) int {
	turn = next_turn(turn)
	result := is_game_over(bb)
	if result == 0 {
		return 0
	} else if result == 1 {
		return -1
	} else if result == -1 {
		return -1
	}

	max := -infinity
	move_list := get_empty_cells(bb)
	for i := 0; i < len(move_list); i++ {
		move := move_list[i]
		bb[turn] = make_move(bb[turn], move)
		score := -negamax(bb, turn)
		bb[turn] = unmake_move(bb[turn], move)

		if score > max {
			max = score
		}
	}

	return max
}

func search_best_move(bb [2]uint16) [2]uint16 {
	var best_move int
	max := -infinity
	turn := o
	move_list := get_empty_cells(bb)
	for i := 0; i < len(move_list); i++ {
		move := move_list[i]
		bb[turn] = make_move(bb[turn], move)
		score := -negamax(bb, turn)
		bb[turn] = unmake_move(bb[turn], move)

		if score > max {
			max = score
			best_move = move
		}
	}

	fmt.Println("best move: ", best_move)
	turn = o
	bb[o] = make_move(bb[o], best_move)
	return bb
}

func pick_random_move(bb [2]uint16) [2]uint16 {
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

	bb[x] = make_move(bb[x], move-1)
	return bb
}

func get_empty_cells(bb [2]uint16) []int {
	var empty_cells []int

	for i := 0; i < 9; i++ {
		if (get_bit(bb[x], i) == 0) && (get_bit(bb[o], i) == 0) {
			empty_cells = append(empty_cells, i)
		}
	}

	return empty_cells
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
		if (get_bit(bb[x], i) == 0) && (get_bit(bb[o], i) == 0) {
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

func make_move(bb uint16, move int) uint16 {
	bb = set_bit(bb, move)
	return bb
}

func unmake_move(bb uint16, move int) uint16 {
	bb = flip_bit(bb, move)
	return bb
}

func set_bit(bb uint16, index int) uint16 {
	return bb | (1 << index)
}

func flip_bit(bb uint16, index int) uint16 {
	return bb ^ (1 << index)
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
	if (occ & (1 << index)) > 0 {
		return true
	}
	return false
}

func print_board(bb [2]uint16) {
	for i := 0; i < 9; i++ {
		bit_x := get_bit(bb[x], i)
		bit_o := get_bit(bb[o], i)
		if bit_x == 1 {
			fmt.Print("X ")
		} else if bit_o == 1 {
			fmt.Print("O ")
		} else {
			fmt.Print("- ")
		}
		if i%3 == 2 {
			fmt.Print("\n")
		}
	}
}

func next_turn(turn int) int {
	if turn == x {
		return o
	} else {
		return x
	}
}
