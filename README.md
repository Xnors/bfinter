# 这是一个 brainfuck 语言解释器

## 构建:
切换到 `%GO_PATH%/src`目录
```shell
cd %GO_PATH%/src
```
克隆此项目
```shell
git clone https://github.com/fexcode/bfinter.git
```
切换到项目目录
```shell
cd bfinter
```
构建此项目
```shell
go build
```
或者, 运行此项目
```shell
go run .
```

## 用法：
`./bfinter run <filename>` 运行 brainfuck 文件

`./bfinter compile <filename>` 编译 brainfuck 文件成可执行文件

`./bfinter outc <filename>` 编译 brainfuck 文件成 C 语言文件

`./bfinter cmd <brainfuck-command>` 执行 brainfuck 命令(代码)

`./bfinter check <filename>` 检查 brainfuck 文件的可运行性

## 注意:
如果代码里面有无限循环, 检查器将不会有反应.
这样可能会占用极多内存资源.

所以,运行前请确认代码里面没有无用的无限循环

如果您希望我编写一个设置最大循环上线的功能(其实之前写过,但是最后删了),

请提交一个 issue

感谢!!!
