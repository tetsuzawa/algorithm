#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int main() {
    int A, B;
    cin >> A >> B;
    int ans;
    ans = A - 2 * B;
    if (ans <= 0) {
        cout << 0 << endl;
    } else {
        cout << ans << endl;
    }
    return 0;
}
