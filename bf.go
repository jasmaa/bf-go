// Interprets brainfuck programs
package main

import (
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
)
	
var program []byte
var pc int
var memory []byte
var dc int
var inp []byte

// Load program
func load(fname string){

	data, err := ioutil.ReadFile(fname)
	
	if err != nil {
		fmt.Println("ERROR: Could not read in file")
		return
	}
	
	// Initialize
	program = data
	pc = 0
	memory = make([]byte, 30000)
	dc = 0
	inp = make([]byte, 1)
}

// Run program
func run(){

	for ; pc < len(program); {
	
		// step
		switch program[pc] {
			case '>':
				dc++
				if dc >= len(memory) {
					fmt.Println("ERROR: Data pointer out-of-bounds")
					return
				}
				
			case '<':
				dc--
				if dc < 0 {
					fmt.Println("ERROR: Data pointer out-of-bounds")
					return
				}
				
			case '+':
				memory[dc]++
				
			case '-':
				memory[dc]--
				
			case '.':
				fmt.Printf("%c", memory[dc])
				
			case ',':
				reader := bufio.NewReader(os.Stdin)
				fmt.Print(">")
				_, err := reader.Read(inp)
				
				if err != nil {
					fmt.Println("ERROR: Could not read input")
					return
				}
				
				memory[dc] = inp[0]
				
			case '[':
				if memory[dc] == 0 {
					// search for matching closing
					encounters := 1
					 for {
						pc++
						
						if pc >= len(program) || pc < 0 {
							fmt.Println("ERROR: Program counter out-of-bounds")
							return
						}
						
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
						
						if pc >= len(program) || pc < 0 {
							fmt.Println("ERROR: Program counter out-of-bounds")
							return
						}
						
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
		}
		
		// Increment program counter
		pc++
	}
}

func main(){

	if len(os.Args) < 2 {
		fmt.Println("ERROR: No program specified")
		return
	}

	load(os.Args[1])
	run()
}