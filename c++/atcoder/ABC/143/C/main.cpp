#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int main() {
    int N;
    cin >> N;
    string S;
    cin >> S;

    char c = S[0];
    int cnt = 1;
    for (int i = 1; i < (N); ++i) {
        if (S[i] != c) {
            cnt++;
        }
        c = S[i];
    }
    cout << cnt << endl;

    return 0;
}