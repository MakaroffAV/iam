def main(input):
    res = []
    for i in input:
        if "-" not in i:
            res.append(int(i))
        else:
            s, e = map(int, i.split("-"))
            for j in range(s, e+1):
                res.append(j)
    print(" ".join(list(map(str, res))))


if __name__ == "__main__":
    main(input().split(","))