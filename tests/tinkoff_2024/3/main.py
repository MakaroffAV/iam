def check_char_rule(chars, must_chars):
    return set(must_chars) == set(chars)


def sliding_window(chars, password_len):
    if len(chars) <= password_len:
        return chars
    for i in range(len(chars)-password_len+1):
        yield chars[i:i+password_len]


def main(chars, must_chars, password_len):
    # сгенерируем все
    # возможные пароли
    # через алгоритм скользящего окна
    possible_password = ""
    for i in sliding_window(chars, password_len):
        if check_char_rule(i, must_chars):
            possible_password = i
    
    # если не найден ни один
    # возможный пароль - выводим -1
    print(possible_password) if len(possible_password) != 0 else print(-1)

    
if __name__ == "__main__":
    main(input(), input(), int(input()))
