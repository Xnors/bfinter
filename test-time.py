import subprocess
import time

cmds = []


# 定义一个函数来测量命令执行的时间
def measure_command_time(command):
    times = []
    for _ in range(5):
        print(f"正在执行命令 '{command}'...")
        start_time = time.time()  # 记录开始时间
        subprocess.run(command, shell=True)  # 执行命令
        end_time = time.time()  # 记录结束时间
        elapsed_time = (end_time - start_time) * 1000  # 转换为毫秒
        times.append(elapsed_time)
    average_time = sum(times) / len(times)  # 计算平均时间
    print(cmd := (f"命令 '{command}' 执行5次的平均用时: {average_time:.2f} 毫秒"))
    cmds.append(cmd)
    return average_time


# 定义要执行的命令列表
commands = [
    # "go build -a .",
    ".\\bfinter.exe compile test\\fiction.bf",
    ".\\bfinter.exe run test\\fiction.bf",
    ".\\test\\fiction.exe",
]

# 循环执行命令并测量时间
for cmd in commands:
    measure_command_time(cmd)

print("\n".join(cmds))
