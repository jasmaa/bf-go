// Interprets brainfuck programs
package main

import (
	"fmt"
	"bufio"
	"os"
	"io/ioutil"
)

type BF struct {
	pc, dc int
	program, memory, inp []byte
}

// Load program
func (bf *BF) Load(fname string){

	data, err := ioutil.ReadFile(fname)
	
	if err != nil {
		fmt.Println("ERROR: Could not read in file")
		return
	}
	
	// Initialize
	bf.program = data
	bf.pc = 0
	bf.memory = make([]byte, 30000)
	bf.dc = 0
}

// Run program
func (bf *BF) Run(){

	for ; bf.pc < len(bf.program); {
	
		// step
		switch bf.program[bf.pc] {
			case '>':
				bf.dc++
				if bf.dc >= len(bf.memory) {
					fmt.Println("ERROR: Data pointer out-of-bounds")
					return
				}
				
			case '<':
				bf.dc--
				if bf.dc < 0 {
					fmt.Println("ERROR: Data pointer out-of-bounds")
					return
				}
				
			case '+':
				bf.memory[bf.dc]++
				
			case '-':
				bf.memory[bf.dc]--
				
			case '.':
				fmt.Printf("%c", bf.memory[bf.dc])
				
			case ',':
				reader := bufio.NewReader(os.Stdin)
				inp := make([]byte, 1)
				
				fmt.Print(">")
				_, err := reader.Read(inp)
				
				if err != nil {
					fmt.Println("ERROR: Could not read input")
					return
				}
				
				bf.memory[bf.dc] = inp[0]
				
			case '[':
				if bf.memory[bf.dc] == 0 {
					// search for matching closing
					encounters := 1
					 for {
						bf.pc++
						
						if bf.pc >= len(bf.program) || bf.pc < 0 {
							fmt.Println("ERROR: Program counter out-of-bounds")
							return
						}
						
						if bf.program[bf.pc] == '[' {
							encounters++
						} else if bf.program[bf.pc] == ']' {
							encounters--
							if encounters == 0 {
								break
							}
						}
					 }
				}
				
			case ']':
				if bf.memory[bf.dc] != 0 {
					// search for matching opening
					encounters := 1
					 for {
						bf.pc--
						
						if bf.pc >= len(bf.program) || bf.pc < 0 {
							fmt.Println("ERROR: Program counter out-of-bounds")
							return
						}
						
						if bf.program[bf.pc] == ']' {
							encounters++
						} else if bf.program[bf.pc] == '[' {
							encounters--
							if encounters == 0 {
								break
							}
						}
					 }
				}
		}
		
		// Increment program counter
		bf.pc++
	}
}

func main(){

	if len(os.Args) < 2 {
		fmt.Println("ERROR: No program specified")
		return
	}

	bf := BF{}
	bf.Load(os.Args[1])
	bf.Run()
}