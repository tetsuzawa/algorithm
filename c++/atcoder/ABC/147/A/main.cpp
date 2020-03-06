#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int main() {
    int a1, a2, a3;
    cin >> a1 >> a2 >> a3;
    int sum = a1 + a2 + a3;
    if (sum < 22) {
        cout << "win" << endl;
    } else {
        cout << "bust" << endl;
    }

    return 0;
}