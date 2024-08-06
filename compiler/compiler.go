package compiler

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
)

func Compile(filename string) {
	// 编译Brainfuck代码到可执行文件
	cFileName := CompileToC(filename)
	CompileCToExecutableFile(cFileName)

	//删除C语言代码文件
	os.Remove(cFileName)

}

func CompileToC(filename string) string {
	// 获取输入文件名
	inputFileName := filename

	// 去除后缀名
	ext := path.Ext(inputFileName)
	outputFileName := inputFileName[:len(inputFileName)-len(ext)] + ".c"

	// 打开输入文件
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Printf("打开文件时出错: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close() // 程序结束时关闭文件

	// 读取输入文件内容
	code, err := readEntireFile(inputFile)
	if err != nil {
		fmt.Printf("读取文件时出错: %v\n", err)
		os.Exit(1)
	}

	// 初始化C语言代码字符串
	cCode := "#include <stdio.h>\nint main() {\n    char tape[30000] = {0};\n    char *ptr = tape;\n"

	// 遍历Brainfuck代码，并转换为C语言
	var count int
	var lastOp rune
	for _, char := range string(code) {
		switch char {
		case '+', '-':
			if char == lastOp {
				// 如果当前操作符与前一个相同，增加计数
				count++
			} else {
				// 如果操作符改变或首次遇到操作符，处理之前的操作
				if lastOp != 0 {
					if count > 1 {
						if lastOp == '+' {
							cCode += fmt.Sprintf("    *ptr += %d;\n", count)
						} else {
							cCode += fmt.Sprintf("    *ptr -= %d;\n", count)
						}
					} else if count == 1 {
						if lastOp == '+' {
							cCode += "    ++*ptr;\n"
						} else {
							cCode += "    --*ptr;\n"
						}
					}
				}
				// 重置计数器和上一个操作符
				count = 1
				lastOp = char
			}
		default:
			// 如果操作符改变或首次遇到操作符，处理之前的操作
			if lastOp != 0 {
				if count > 1 {
					if lastOp == '+' {
						cCode += fmt.Sprintf("    *ptr += %d;\n", count)
					} else {
						cCode += fmt.Sprintf("    *ptr -= %d;\n", count)
					}
				} else if count == 1 {
					if lastOp == '+' {
						cCode += "    ++*ptr;\n"
					} else {
						cCode += "    --*ptr;\n"
					}
				}
				// 重置计数器和上一个操作符
				count = 0
				lastOp = 0
			}
			// 根据当前操作符添加C代码
			switch char {
			case '>':
				cCode += "    ++ptr;\n"
			case '<':
				cCode += "    --ptr;\n"
			case '.':
				cCode += "    putchar(*ptr);\n"
			case ',':
				cCode += "    *ptr = getchar();\n"
			case '[':
				cCode += "    while (*ptr) {\n"
			case ']':
				cCode += "    }\n"
			}
		}
	}

	// 处理最后一个操作符（如果有的话）
	if lastOp != 0 {
		if count > 1 {
			if lastOp == '+' {
				cCode += fmt.Sprintf("    *ptr += %d;\n", count)
			} else {
				cCode += fmt.Sprintf("    *ptr -= %d;\n", count)
			}
		} else if count == 1 {
			if lastOp == '+' {
				cCode += "    ++*ptr;\n"
			} else {
				cCode += "    --*ptr;\n"
			}
		}
	}

	// 添加C语言代码的结束部分
	cCode += "    return 0;\n}"

	// 写入C语言代码到输出文件
	err = writeToFile(outputFileName, []byte(cCode))
	if err != nil {
		fmt.Printf("写入文件时出错: %v\n", err)
		os.Exit(1)
	}

	return outputFileName
}

func CompileCToExecutableFile(filename string) {

	// 去除后缀名
	ext := path.Ext(filename)
	noExtFileName := filename[:len(filename)-len(ext)]

	// 获取输入文件名
	inputFileName := filename

	fmt.Printf("编译: 输入文件:  %s\n输出文件:  %s\n", inputFileName, noExtFileName)

	// 生成输出文件名
	var outputFileName string

	// windows下生成.exe文件, linux下生成.out文件
	if runtime.GOOS == "windows" {
		outputFileName = noExtFileName + ".exe"
	} else if runtime.GOOS == "linux" {
		outputFileName = noExtFileName + ".out"
	}

	// gcc 编译
	// 运行 gcc inputFileName -o outputFileName
	err := exec.Command("gcc", inputFileName, "-o", outputFileName).Run()
	if err != nil {
		log.Fatalf("编译失败: %v", err)
		panic(err)
	}
	fmt.Printf("编译成功: 输出文件:  %s\n", outputFileName)
}

// readEntireFile 读取整个文件内容
func readEntireFile(file *os.File) ([]byte, error) {
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// 根据文件大小分配缓冲区
	size := info.Size()
	buf := make([]byte, size)

	// 读取文件内容
	_, err = file.Read(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// writeToFile 将数据写入文件
func writeToFile(filename string, data []byte) error {
	// 打开文件用于写入
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close() // 程序结束时关闭文件

	// 写入数据到文件
	_, err = file.Write(data)
	return err
}
