#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;
const ll MAX = 1000000001;

ll A, B;

ll ndig(ll x) {
    if (!x) return 1;

    int cnt = 0;
    while (x) {
        cnt++;
        x /= 10;
    }
    return cnt;
}

ll price(ll A, ll B, ll N) { return A * N + B * ndig(N); }

bool isOK(ll N, ll key) {
    // cout << price(A, B, N) << endl;
    if (price(A, B, N) > key)
        return true;
    else
        return false;
}

ll binary_search(ll key) {
    ll ng = 0;
    ll ok = MAX;

    while (abs(ok - ng) > 1) {
        ll mid = (ok + ng) / 2;

        if (isOK(mid, key))
            ok = mid;
        else
            ng = mid;
    }

    return ng;
};

int main() {
    ll X, ans;
    cin >> A >> B >> X;
    if (X == MAX)
        ans = MAX;
    else
        ans = binary_search(X);

    cout << ans << endl;

    return 0;
}