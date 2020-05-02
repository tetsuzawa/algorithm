#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int dtob(int decimal) {
    int binary = 0;
    int base = 1;
    while (decimal > 0) {
        binary = binary + (decimal % 2) * base;
        decimal = decimal / 2;
        base = base * 10;
    }
    return binary;
}

int digsum(int n) {
    int res = 0;
    while (n > 0) {
        res += n % 10;
        n /= 10;
    }
    return res;
}

int pow(int x, int y) {
    int sum = 1;
    rep(i, y) { sum = sum * x; };
    return sum;
}

int main() {
    int N, M, Q;
    cin >> N >> M >> Q;
    int a[Q], b[Q], c[Q], d[Q];
    rep(i, Q) { cin >> a[i] >> b[i] >> c[i] >> d[i]; };
    int A[M];
    rep(i, M) { A[i] = i + 1; };

    int trueA[N];
    int presum = 0;
    for (int i = 1; i <= 10; ++i) {
        cout << N << " N:digsum " << digsum(dtob(i)) << endl;
        if (digsum(dtob(i)) == N) {
            int n = dtob(i);
            rep(j, N) {
                int cnt = 0;
                while (n > 0) {
                    trueA[j] = pow(2, cnt);
                    n /= 10;
                    cnt--;
                }
            }
            int sum = 0;
            rep(k, Q) {
                if (trueA[b[k]] - trueA[a[k]] == c[i]) {
                    sum += d[k];
                }
                return 0;
            }
            if (sum > presum) {
                presum = sum;
            }
        }
    }
    cout << presum << endl;
    return 0;
}