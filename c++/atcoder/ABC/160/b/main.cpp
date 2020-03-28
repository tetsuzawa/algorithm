#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int main() {
    ll x;
    cin >> x;
    int n500 = 0, n5 = 0;
    if (x == 0) {
        cout << 0 << endl;
        return 0;
    }
    while (x - 500 >= 0) {
        x -= 500;
        n500++;
    }
    while (x - 5 >= 0) {
        x -= 5;
        n5++;
    }
    // while (x - 100 >= 0) {
    //     x -= 100;
    // }
    // while (x - 50 >= 0) {
    //     x -= 50;
    // }
    // while (x - 10 >= 0) {
    //     x -= 10;
    // }
    // while (x - 1 >= 0) {
    //     x -= 1;
    // }
    cout << n500 * 1000 + n5 * 5 << endl;

    return 0;
}