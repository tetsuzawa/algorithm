#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

string ABC = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";

int index(char letter) {
    rep(i, ABC.size()) {
        if (letter == ABC[i]) return i;
    }
    return 0;
}

int main() {
    int n;
    string s;
    cin >> n >> s;
    rep(i, s.size()) { s[i] = ABC[(index(s[i]) + n) % ABC.size()]; }
    cout << s << endl;
    return 0;
}