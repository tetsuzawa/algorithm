#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int main() {
    ll X;
    cin >> X;
    ll money = 100;
    ll cnt = 0;
    while (money < X) {
        money = money + money / 100;
        cnt++;
    }
    cout << cnt << endl;
    return 0;
}