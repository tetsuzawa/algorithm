#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

vector<vector<ll>> comb(int n, int r) {
    vector<vector<ll>> v(n + 1, vector<ll>(n + 1, 0));
    for (int i = 0; i < v.size(); i++) {
        v[i][0] = 1;
        v[i][i] = 1;
    }
    for (int j = 1; j < v.size(); j++) {
        for (int k = 1; k < j; k++) {
            v[j][k] = (v[j - 1][k - 1] + v[j - 1][k]);
        }
    }
    return v;
}

int main() {
    int n, m;
    cin >> n >> m;
    vector<vector<ll>> ncr, mcr;
    ncr = comb(n, 2);
    mcr = comb(m, 2);
    if (n <= 1 && m <= 1) {
        cout << 0 << endl;
    } else {
        cout << ncr[n][2] + mcr[m][2] << endl;
    }
    return 0;
}