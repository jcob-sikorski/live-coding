def robotWithString(s: str) -> str:
    n = len(s)

    suffix_min = ["{"] * (n + 1)
    for i in range(n-1, -1, -1):
        suffix_min[i] = min(suffix_min[i+1], s[i])

    paper = []
    t = []
    i = 0

    while i < n:
        t.append(s[i])
        while t and t[-1] <= suffix_min[i+1]:
            paper.append(t.pop())

        i += 1

    while t:
        paper.append(t.pop())

    return "".join(paper)
    
print(robotWithString("bac"))