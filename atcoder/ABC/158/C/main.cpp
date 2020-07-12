#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int main() {
    int A, B;
    cin >> A >> B;
    double xa, xb;
    xa = A * 100 / 8.0;
    xb = B * 100 / 10.0;
    double ans = 0;
    if (xa > xb)
        ans = (double)xa;
    else
        ans = (double)xb;

    if ((int)((int)ans * 0.08) != A || (int)((int)ans * 0.1) != B) {
        if ((int)(((int)ans + 1) * 0.08) == A && (int)(((int)ans + 1) * 0.1) == B) {
            cout << (int)ans + 1 << endl;
        } else
            cout << -1 << endl;
    } else
        cout << (int)ans << endl;
    return 0;
}