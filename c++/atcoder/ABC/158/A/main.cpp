#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int main() {
    string S;
    cin >> S;
    if (S[0] == S[1] && S[1] == S[2])
        cout << "No" << endl;
    else
        cout << "Yes" << endl;
    return 0;
}