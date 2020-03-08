#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef unsigned long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int main() {
    ll N, A, B;
    cin >> N >> A >> B;
    ll ans = 0;
    ll sum = 0;
    sum = N / (A + B);
    ans = N / (A + B) * A + min(sum, A);

    cout << ans << endl;

    return 0;
}