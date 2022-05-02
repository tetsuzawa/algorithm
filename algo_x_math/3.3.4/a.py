N = int(input())
A = map(int, input().split())

d = {100:0, 200:0, 300:0, 400:0}
for a in A:
    d[a] += 1

print(d[100]*d[400] + d[200] * d[300])
