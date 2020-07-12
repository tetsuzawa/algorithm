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
    ll n;
    cin >> n;
    ll a[n], kmax = 0, vmax = 0;
    unordered_map<ll, ll> bucket;

    rep(i, n) { bucket[i] = 0; }
    rep(i, n) {
        cin >> a[i];
        bucket[a[i]]++;
        if (a[i] > kmax) {
            kmax = a[i];
        }
    }
    rep(i, n) {
        for (auto itr = bucket.begin(); itr != bucket.end(); ++itr) {
            if (itr->second > vmax) {
                vmax = itr->second;
            }
        }
    }
    vector<vector<ll>> ncr;
    ncr = comb(n, vmax);
    vector<ll> sumncr;
    for (auto itr = bucket.begin(); itr != bucket.end(); ++itr) {
        sumncr[itr->first] = ncr[itr->second][itr->second - 1];
    }
    ll sum = accumulate(sumncr.begin(), sumncr.end(), 0);
    for (auto itr = bucket.begin(); itr != bucket.end(); ++itr) {
        cout << sum - ncr[itr->second][itr->second - 1] + ncr[itr->second - 1][itr->second - 2] << endl;
    }

    return 0;
}