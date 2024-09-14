from datetime import datetime, timedelta


def main(started, n, commands):
    """
    Если команда взламывает сервер, её счет увеличивается на один — а к штрафному
    времени прибавляется время в минутах, округленное вниз, которое прошло от начала 
    соревнования до времени взлома. 
    
    Если перед удачной попыткой взлома одного сервера
    команда совершает одну или несколько неудачных попыток взлома этого же сервера — то
    к штрафному времени прибавляется по двадцать минут за каждую такую неудачную попытку
    
    ACCESSED — сервер взломан;
    DENIED, FORBIDEN - неудачная попытка взлома;
    PONG - просто запрос к серверу, мы его никак не учитываем
    """
    
    teams = set()
    servers = set()
    for i in commands:
        if i[0] not in teams:
            teams.add(i[0])
        if i[2] not in servers:
            servers.add(i[2])
            
    rateboard = {
        i: {
            k: {
                "hack": False,
                "fine": 0,
            }
            for k in servers
        }
        for i in teams
    }
    
    for i in commands:
        if i[3] == "ACCESSED":
            # Если команда взломала сервер, отмечаем это
            # и добавляем время от начала хакатона к штрафу
            rateboard[i[0]][i[2]]["hack"] = True
            rateboard[i[0]][i[2]]["fine"] += get_minutes_from_start(started, convert_time(i[1]))
        if i[3] == "DENIED" or i[3] == "FORBIDEN":
            # Если команда не взломала сервер,
            # добавляем для команды штрафные балы по этому серверу
            rateboard[i[0]][i[2]]["fine"] += 20
            
    # создаем leaderboard
    leaderboard = []
    for i in rateboard:
        # имя, кол-во взломанных серверов, общее время
        leaderboard.append([i, total_cracked_servers(rateboard[i]), total_fine_minutes(rateboard[i])])
        
    # сортируем и выводим результат
    leaderboard.sort(key=lambda x: (-x[1], x[2], x[0]))
    rank = 1
    for i, (team_name, servers, fine_time) in enumerate(leaderboard):
        if i > 0 and (leaderboard[i - 1][1] != servers or leaderboard[i - 1][2] != fine_time):
            rank = i + 1
        print(f'{rank} {team_name} {servers} {fine_time}')

def total_fine_minutes(attempts):
    r = 0
    for i in attempts:
        r += attempts[i]["fine"]
    return r

def total_cracked_servers(attempts):
    r = 0
    for i in attempts:
        if attempts[i]["hack"]:
            r += 1
    return r

def convert_time(time_str):
    return datetime.strptime(time_str, "%H:%M:%S")
    
def get_minutes_from_start(start, current):
    if current < start:
        current += timedelta(days=1)
    return int((current - start).total_seconds() // 60)

if __name__ == "__main__":
    c = []
    s = input()
    n = int(input())

    for i in range(n):
        c.append(input().split(" "))

    main(convert_time(s), n, c)