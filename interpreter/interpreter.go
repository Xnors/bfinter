package interpreter

import (
	"bfinter/constants"
	"fmt"
)

// Brainfuck解释器
func Interpret(code string) {
	memory := make([]byte, constants.MemorySize) // 内存
	ptr := 0                                     // 指针位置
	pc := 0                                      // 程序计数器
	loopStep := 0                                // 循环几次了

	for pc < len(code) {
		switch code[pc] {
		case '>':
			ptr++ // 指针右移
			if ptr >= constants.MemorySize {
				ptr = 0 // 超出内存大小，回到起点
			}
		case '<':
			ptr-- // 指针左移
			if ptr < 0 {
				ptr = constants.MemorySize - 1 // 超出内存大小，回到末尾
			}
		case '+':
			memory[ptr]++ // 内存单元加一
		case '-':
			memory[ptr]-- // 内存单元减一
		case '.':
			fmt.Printf("%c", memory[ptr]) // 输出内存单元的字符
		case ',':
			// 在非交互模式下，我们简单地跳过输入指令
			// 如果需要，可以在这里实现其他逻辑，例如从字符串中读取输入
		case '[':
			if memory[ptr] == 0 {
				// 如果当前内存单元为0，跳过对应的']'

				bracketDepth := 1
				for pc++; pc < len(code) && bracketDepth > 0; pc++ {
					if code[pc] == '[' {
						bracketDepth++
					} else if code[pc] == ']' {
						bracketDepth--
					}
				}
			}
		case ']':
			if memory[ptr] != 0 {
				// 如果当前内存单元不为0，跳回对应的'['

				// 记录循环次数
				loopStep++
				if loopStep > constants.MaxLoop {
					fmt.Printf("ERROR: 有可能存在一个死循环!\nINDEX: %d\n\n", pc+1)
					return
				}

				bracketDepth := 1
				for pc--; pc >= 0 && bracketDepth > 0; pc-- {
					if code[pc] == '[' {
						bracketDepth--
					} else if code[pc] == ']' {
						bracketDepth++
					}
				}
			}
		}
		pc++
	}
}
