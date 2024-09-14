import sys
from collections import defaultdict

sys.setrecursionlimit(200000)

def time_max(process, process_time, dependes, cache):
    # Если есть уже
    # время для процесса - вернем его
    if cache[process] != -1:
        return cache[process]
    
    # Время завершения процесса - его + дочки
    time_value = 0
    
    # Время всех дочерних процессов главного процесса
    for i in dependes[process]:
        time_value = max(
            time_value, 
            time_max(i, process_time, dependes, cache),
        )
    
    # Время завершения = макс время зависимостей + собственное
    cache[process] = time_value + process_time[process]
    return cache[process]

def main():
    n = int(input())
    
    exec_time  = []
    depends_on = defaultdict(list)
    
    for i in range(n):
        values = list(map(int, input().split()))
        exec_time.append(values[0])
        for j in values[1:]:
            depends_on[i].append(j - 1)

    cache      = [-1] * n
    total_time = 0
    
    for i in range(n):
        total_time = max(
            total_time, 
            time_max(i, exec_time, depends_on, cache),
        )
    print(total_time)

if __name__ == "__main__":
    main()
