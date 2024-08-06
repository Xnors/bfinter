import subprocess
import time
from os.path import getsize

cmds = []
subprocess.run("go build -a .", shell=True)  # 编译


# 定义一个函数来测量命令执行的时间
def measure_command_time(command):
    times = []
    for _ in range(2):
        print(f"正在执行命令 '{command}'...")
        start_time = time.time()  # 记录开始时间
        subprocess.run(command, shell=True)  # 执行命令
        end_time = time.time()  # 记录结束时间
        elapsed_time = (end_time - start_time) * 1000  # 转换为毫秒
        times.append(elapsed_time)
    min_time = min(times)  # 计算最短时间
    print(cmd := (f"命令 '{command}' 执行2次的最短用时: {min_time:.2f} 毫秒"))
    cmds.append(cmd)
    return min_time


# 定义要执行的命令列表
commands = [
    ".\\bfinter.exe compile test\\fiction.bf",
    ".\\bfinter.exe run test\\fiction.bf",
    ".\\test\\fiction.exe",
]


# 循环执行命令并测量时间
for cmd in commands:
    measure_command_time(cmd)

# 计算 test/fiction.exe 文件大小
size = getsize("test/fiction.exe") / 1024

print(f"test/fiction.exe 文件大小: {size}Kib")
cmds.insert(0, f"test/fiction.exe 文件大小: {size}Kib")

print("\n")
for cmd in cmds:
    print(cmd)
