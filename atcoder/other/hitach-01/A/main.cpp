#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int main() {
    string S;
    cin >> S;
    string ans = "No";
    if (S.size() % 2 != 0) {
        cout << ans << endl;
        return 0;
    }
    for (int i = 0; i < S.size() / 2; ++i) {
        if (S[2 * i] != 'h' || S[2 * i + 1] != 'i') {
            cout << ans << endl;
            return 0;
        }
    }
    ans = "Yes";
    cout << ans << endl;
    return 0;
};