package main

import (
	"fmt"
	"bufio"
	"os"
)
	
var program []byte
var pc int
var memory []byte
var dc int

var inp []byte

func step(){
	
	switch program[pc] {
		case '>':
			dc++
		case '<':
			dc--
		case '+':
			memory[dc]++
		case '-':
			memory[dc]--
		case '.':
			fmt.Printf("%c", memory[dc])
		case ',':
			reader := bufio.NewReader(os.Stdin)
			
			_, err := reader.Read(inp)
			
			if err != nil {
				fmt.Println("Error reading input")
				return
			}
			
			memory[dc] = inp[0]
			
		case '[':
			
			if memory[dc] == 0 {
				// search for matching closing
				encounters := 1
				 for {
					pc++
					if program[pc] == '[' {
						encounters++
					} else if program[pc] == ']' {
						encounters--
						if encounters == 0 {
							break
						}
					}
				 }
			}
			
		case ']':
			
			if memory[dc] != 0 {
				// search for matching opening
				encounters := 1
				 for {
					pc--
					if program[pc] == ']' {
						encounters++
					} else if program[pc] == '[' {
						encounters--
						if encounters == 0 {
							break
						}
					}
				 }
			}
			
		default:
			fmt.Println("error")
	}
	
	// debug
	//fmt.Println("PC:", pc)
	//fmt.Printf("%q\n", program[pc])
	//fmt.Println("Data", memory[:10])
	//fmt.Println("===")
	
	// Increment program counter
	pc++
}

func main(){
	fmt.Println("Starting interpreter...")
	
	// Initialize
	//program = []byte("+[-[<<[+[--->]-[<<<]]]>>>-]>-.---.>..>.<<<<-.<+.>>>>>.>.<<.<-.")
	program = []byte("++++++++[>+>++++<<-]>++>>+<[-[>>+<<-]+>>]>+[-<<<[->[+[-]+>++>>>-<<]<[<]>>++++++[<<+++++>>-]+<<++.[-]<<]>.>+[>>]>+]")
	pc = 0
	memory = make([]byte, 2048)
	dc = 0
	
	inp = make([]byte, 1)
	
	for ; pc < len(program); {
		step()
	}
	
}