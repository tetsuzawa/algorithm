#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int main() {
    string s;
    cin >> s;
    int n;
    n = s.length();
    int l1, l2, l4;
    l1 = (n - 1) / 2;
    l2 = (n + 3) / 2;
    l4 = n / 2;
    if (n == 3) {
        if (s[0] == s[2]) {
            cout << "Yes" << endl;
            return 0;
        } else {
            cout << "No" << endl;
            return 0;
        }
    }
    rep(i, l4) {
        if ((s[i] != s[l1 - i - 1]) || (s[l2 + i - 1] != s[n - i - 1])) {
            cout << "No" << endl;
            return 0;
        }
    }
    rep(i, n) {
        if (s[i] != s[n - i - 1]) {
            cout << "No" << endl;
            return 0;
        }
    }
    cout << "Yes" << endl;
    return 0;
}