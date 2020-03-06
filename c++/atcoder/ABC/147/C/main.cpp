#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int main() {
    int N;
    cin >> N;
    vector<vector<int>> g(N, vector<int>(N, 0));
    int A[N];
    int x, y;
    for (int i = 0; i < (N); ++i) {
        cin >> A[i];
        for (int j = 0; j < (A[i]); ++j) {
            cin >> x >> y;
            if (y) {
                g[i][x] = y;
            }
        }
    }
    map<int, int> mp;
    rep(i, N) {
        rep(j, N) {
            if (g[j][i]) {
                mp[i] = 1;
            }
        }
    }
    cout << mp.size() << endl;

    return 0;
}