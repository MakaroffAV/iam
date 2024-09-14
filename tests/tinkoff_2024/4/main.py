def is_prime(number):
    if number > 1:
        for num in range(2, number):
            if number % num == 0:
                return False
        return True
    return False
     
     
def div_number(number):
    d = 2
    for i in range(2, int(number/2)+1):
        if number % i == 0:
            d += 1
    return d   
        
        
def main(l, r):
    o = 0
    for i in range(l, r+1):
        # проверяем является
        # ли число составным
        if not is_prime(i) and i != 1:
            if is_prime(div_number(i)):
                o += 1
    print(o)


if __name__ == "__main__":
    main(*list(map(int, input().split(" "))))