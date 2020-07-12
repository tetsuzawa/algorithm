#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int main() {
    int N;
    cin >> N;
    int d[N];
    rep(i, N) { cin >> d[i]; }

    int sum = 0;
    rep(i, N) {
        for (int j = i + 1; j < N; ++j) {
            sum += d[i] * d[j];
        }
    }
    cout << sum << endl;

    return 0;
}