#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

string week[7] = {"SUN", "MON", "TUE", "WED", "THU", "FRI", "SAT"};

int search(string day) {
    rep(i, 7) {
        if (day == week[i]) return i;
    }
    return 0;
}

int main() {
    string S;
    cin >> S;
    cout << 7 - search(S) << endl;
    return 0;
}