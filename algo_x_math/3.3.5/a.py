N = int(input())
A = map(int, input().split())

d = {1:0, 2:0, 3:0}
for a in A:
    d[a] += 1

ans = 0
for k in d.keys():
    ans += d[k]*(d[k] -1) / 2

print(int(ans))