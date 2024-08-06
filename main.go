package main

import (
	"C" // 必须引入, 否则编译时会报错, 我也不知道为什么
	"bfinter/analyzer"
	"bfinter/compiler"
	"bfinter/interpreter"
	"fmt"
	"os"
)

func main() {
	checkArgs()

	switch os.Args[1] {
	case "run":
		bfCode, err := readBrainfuckCode(os.Args[2])
		if err != nil {
			panic(err)
		}

		check(bfCode, false)
		run(bfCode)
	case "check":
		bfCode, err := readBrainfuckCode(os.Args[2])
		if err != nil {
			panic(err)
		}

		check(bfCode, true)
	case "cmd":
		// 直接执行命令
		cmd(os.Args[2])
	case "compile":
		// 检查代码
		check(os.Args[2], false)

		// 编译代码
		compiler.Compile(os.Args[2])

	case "outc":
		// 检查代码
		check(os.Args[2], false)

		// 转换代码
		compiler.CompileToC(os.Args[2])

	default:
		panic("Invalid command")
	}
}

func cmd(bfCode string) {
	check(bfCode, false)
	run(bfCode)
}

func check(bfCode string, printCheckLogs bool) string {

	if printCheckLogs {
		fmt.Println("静态检查中...")
		analyzer.StaticAnalyze(bfCode)

		fmt.Println("动态检查中...")
		// analyzer.DynamicAnalyze(bfCode)

		fmt.Println("代码检查完毕")

		return bfCode
	}

	analyzer.StaticAnalyze(bfCode)
	// analyzer.DynamicAnalyze(bfCode)
	return bfCode
}

func run(bfCode string) {
	interpreter.Interpret(bfCode)
}

// readBrainfuckCode 从指定的文件路径读取Brainfuck代码
func readBrainfuckCode(filePath string) (string, error) {
	// 读取文件内容
	codeBytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err // 如果读取文件出错，返回错误
	}

	// 将字节数组转换为字符串
	code := string(codeBytes)
	return code, nil
}

func checkArgs() {
	if len(os.Args) != 3 {
		panic("USAGE:\n\t<program> run <file>\t\tTo run brainfuck code from file\n\t<program> check <file>\tTo check brainfuck code from file\n")
	}
}
