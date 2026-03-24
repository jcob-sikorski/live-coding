from collections import deque

def predictPartyVictory(senate: str) -> str:

    n = len(senate)
    banned = [False] * n

    r_sens = deque()
    d_sens = deque()
    sens = deque()

    for i in range(len(senate)):
        if senate[i] == 'R':
            r_sens.append(i)
        elif senate[i] == 'D':
            d_sens.append(i)
        sens.append(i)

    while len(sens) > 1:
        sen = sens.popleft()
            
        if banned[sen]: 
            continue

        if senate[sen] == 'D':
            if not r_sens: return "Dire"
            to_ban = r_sens.popleft()
            banned[to_ban] = True
            sens.append(sen)
        elif senate[sen] == 'R':
            if not d_sens: return "Radiant"
            to_ban = d_sens.popleft()
            banned[to_ban] = True
            sens.append(sen)

    return "Radiant" if senate[sens[0]] == 'R' else "Dire"

print(predictPartyVictory("DRRDRDRDRDDRDRDR"))