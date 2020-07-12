#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

template <class T>
void print(const T &value) {
    cout << value << endl;
}

bool isOdd(int x) { return x % 2 != 0; }

int main() {
    int N;
    string S;

    cin >> N >> S;

    if (isOdd(N)) {
        cout << "No" << endl;
        return 0;
    }

    rep(i, N / 2) {
        if (S[i] != S[i + N / 2]) {
            cout << "No" << endl;
            return 0;
        }
    }
    cout << "Yes" << endl;

    return 0;
}