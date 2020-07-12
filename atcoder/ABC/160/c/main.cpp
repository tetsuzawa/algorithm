#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int main() {
    int k, n;
    cin >> k >> n;
    int A[n];
    int d = 0;
    rep(i, n) {
        cin >> A[i];
        if (i == 0) {
            continue;
        }
        if (A[i] - A[i - 1] > d) {
            d = A[i] - A[i - 1];
        }
    }
    if (A[0] + k - A[n - 1] > d) {
        d = A[0] + k - A[n - 1];
    }

    cout << k - d << endl;
    return 0;
}

/*
int main() {
    int k, n;
    cin >> k >> n;
    int A[n];
    int d = 0, top = 0, sum = 0;
    rep(i, n) {
        cin >> A[i];
        if (i == 0) {
            continue;
        }
        sum += A[i] - A[i - 1];
        if (A[i] - A[i - 1] > d) {
            d = A[i] - A[i - 1];
            top = i;
        }
    }
    sum += abs(A[n] - A[0]);
    if (A[0] + k - A[n - 1] > d) {
        d = A[0] + k - A[n - 1];
        top = 0;
    }

    // cout << sum - d << endl;
    cout << k - d << endl;

    return 0;
}
*/