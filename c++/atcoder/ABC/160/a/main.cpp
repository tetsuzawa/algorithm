#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int main() {
    string s;
    cin >> s;
    if (s[2] == s[3] && s[4] == s[5]) {
        cout << "Yes" << endl;
    } else {
        cout << "No" << endl;
    }

    return 0;
}