#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int main() {
    ll A, B, N;
    cin >> A >> B >> N;

    if (N > B) {
        int X = B * (N / B) - 1;
        cout << (A * X) / B - A * (X / B) << endl;
    } else {
        int X = N;
        cout << (A * X) / B - A * (X / B) << endl;
    }

    return 0;
}