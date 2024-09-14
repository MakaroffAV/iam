def main(n, levels):
    # добавим в начале массива нуль
    modf = [0] + levels
    
    # проверяем, что реальные данные
    # являются возрастающей последовательностью
    real_data = [
        i for i in modf if i != -1
    ]
    if real_data != sorted(real_data):
        print("NO"); return
    
    # попробуем восстановить данные
    res = modf[:]
    for i in range(1, n+1):
        
        # если потеря данных, 
        # то пробуем их восстановить
        if res[i] == -1:
            
            # найдем диапазон для
            # восстановления данных
            l = res[i-1]
            r = res[i+1] if i < n and res[i+1] != -1 else float("inf")
            if l + 1 >= r:
                print("NO"); return
            res[i] = l+1

    print("YES")
    print(" ".join(map(str, [res[i+1] - res[i] for i in range(n)])))


if __name__ == "__main__":
    main(int(input()), list(map(int, input().split())))
    