package analyzer

import (
	"fmt"
)

// 静态分析Brainfuck代码(会忽略非法字符)
// 检测不匹配的括号
func StaticAnalyze(code string) {
	// 检测不匹配的括号
	checkUnmatchedLoopLabel(code)
}

// detectDeadunmatchedParenthesesLoop 检查Brainfuck程序中是否存在未匹配的循环标签。
func checkUnmatchedLoopLabel(program string) bool {
	var loopStack []int // 初始化一个栈，用于存放循环开始标签 '[' 的位置

	// 遍历程序中的每个指令
	for index, cmd := range program {
		switch cmd {
		case '[':
			// 遇到 '['，将当前索引位置推入栈中
			loopStack = append(loopStack, index)
		case ']':
			// 遇到 ']'，检查栈是否为空
			if len(loopStack) == 0 {
				fmt.Printf("ERROR: 未匹配的']'\nINDEX: %d\n\n", index+1)
				return false
			}
			loopStack = loopStack[:len(loopStack)-1]
		}
	}

	// 遍历完整个程序后，如果栈不为空，说明有未匹配的 '['
	if len(loopStack) > 0 {
		fmt.Printf("ERROR: 未匹配的 '['\nINDEX: %d\n\n", loopStack[len(loopStack)-1]+1)
	}

	// 如果栈为空，说明循环标签匹配正确，返回未检测到明显的死循环
	return false
}

// 动态分析Brainfuck代码(会忽略非法字符)

/* 
func DynamicAnalyze(code string) {
	memory := make([]byte, constants.MemorySize) // 内存
	ptr := 0                                     // 指针位置
	pc := 0                                      // 程序计数器
	maxLoop := 10                                // 最大循环
	loopStep := 0                                // 循环几次了

	for pc < len(code) {
		switch code[pc] {
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
				if loopStep > maxLoop {
					fmt.Printf("ERROR: 有可能存在一个死循环!\nINDEX: %d\n", pc+1)
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
 */