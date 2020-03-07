/* ------------------- top ------------------- */
#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = １ << 30;
const ll MOD = 1000000007;

int main() { return 0; }
/* ------------------- top ------------------- */

/* ------------------- 最大公約数 ------------------- */
// O(log(N))
// 1個
int gcd(int a, int b) { return b ? gcd(b, a % b) : a; }

// n個
int ngcd(vector<int> a) {
    int res;
    res = a[0];
    for (int i = 1; i < a.size() && res != 1; i++) {
        res = gcd(a[i], res);
    }
    return res;
}

ll gcd(ll a, ll b) { return b ? gcd(b, a % b) : a; }

// n個
ll ngcd(vector<ll> a) {
    ll res;
    res = a[0];
    for (int i = 1; i < a.size() && res != 1; i++) {
        res = gcd(a[i], res);
    }
    return res;
}
/* ------------------- 最大公約数 ------------------- */

/* ------------------- 最小公倍数 ------------------- */
int lcm(int a, int b) { return a / gcd(a, b) * b; }

int nlcm(vector<int> numbers) {
    int res;
    res = numbers[0];
    for (int i = 1; i < numbers.size(); i++) {
        res = lcm(res, numbers[i]);
    }
    return res;
}

ll lcm(ll a, ll b) { return a / gcd(a, b) * b; }

ll nlcm(vector<int> numbers) {
    ll res;
    res = numbers[0];
    for (int i = 1; i < numbers.size(); i++) {
        res = lcm(res, numbers[i]);
    }
    return res;
}
/* ------------------- 最小公倍数 ------------------- */

/* ------------------- 素数 ------------------- */
bool isPrime(int x) {
    int i;
    if (x < 2)
        return 0;
    else if (x == 2)
        return 1;
    if (x % 2 == 0) return 0;
    for (i = 3; i * i <= x; i += 2)
        if (x % i == 0) return 0;
    return 1;
}
/* ------------------- 素数 ------------------- */

/* ------------------- 桁和 ------------------- */
int digsum(int n) {
    int res = 0;
    while (n > 0) {
        res += n % 10;
        n /= 10;
    }
    return res;
}
/* ------------------- 桁和 ------------------- */

/* ------------------- 桁数 ------------------- */
int ndig(int x) {
    if (!x) return 1;

    int cnt;
    while (x) {
        cnt++;
        x /= 10;
    }
    return cnt;
}
/* ------------------- 桁数 ------------------- */

/* ------------------- 約数全列挙 ------------------- */
vector<int> enum_div(int n) {
    vector<int> ret;
    for (int i = 1; i * i <= n; ++i) {
        if (n % i == 0) {
            ret.push_back(i);
            if (i != 1 && i * i != n) {
                ret.push_back(n / i);
            }
        }
    }
    return ret;
}
/* ------------------- 約数全列挙 ------------------- */

/* ------------------- 累乗 (二分累乗 or 繰りかえし二乗法) ------------------- */
// xのn乗%modを計算
ll mod_pow(ll x, ll n, ll mod) {
    ll res = 1;
    while (n > 0) {
        if (n & 1) res = res * x % mod;
        x = x * x % mod;
        n >>= 1;
    }
    return res;
}
/* ------------------- 累乗 (二分累乗 or 繰りかえし二乗法) ------------------- */

/* ------------------- 文字列中に存在する特定の文字の個数カウント ------------------- */
int stringcount(string s, char c) { return count(s.cbegin(), s.cend(), c); }
/* ------------------- 文字列中に存在する特定の文字の個数カウント ------------------- */

/* ------------------- セグメント木 ------------------- */
class Monoid {
   public:
    // 単位元
    int unit;

    Monoid() {
        // 単位元
        unit = 0;
    }

    // 演算関数
    int calc(int a, int b) { return a + b; }
};

class SegmentTree {
   public:
    // セグメント木の葉の要素数
    int n;

    // セグメント木
    vector<int> tree;

    // モノイド
    Monoid mono;

    SegmentTree(vector<int> &v) {
        n = 1 << (int)ceil(log2(v.size()));
        tree = vector<int>(n << 1);
        for (int i = 0; i < v.size(); i++) {
            update(i, v[i]);
        }
        for (int i = v.size(); i < n; i++) {
            update(i, mono.unit);
        }
    }

    // k番目の値(0-indexed)をxに変更
    void update(int k, int x) {
        k += n;
        tree[k] = x;
        for (k = k >> 1; k > 0; k >>= 1) {
            tree[k] = mono.calc(tree[k << 1 | 0], tree[k << 1 | 1]);
        }
    }

    // [l, r)の最小値(0-indexed)を求める．
    int query(int l, int r) {
        int res = mono.unit;
        l += n;
        r += n;
        while (l < r) {
            if (l & 1) {
                res = mono.calc(res, tree[l++]);
            }
            if (r & 1) {
                res = mono.calc(res, tree[--r]);
            }
            l >>= 1;
            r >>= 1;
        }
        return res;
    }
    int operator[](int k) {
        // st[i]で添字iの要素の値を返す
        if (k - n >= 0 || k < 0) {
            return -INF;
        }
        return tree[tree.size() - n + k];
    }

    void show() {
        int ret = 2;
        for (int i = 1; i < 2 * n; i++) {
            cout << tree[i] << " ";
            if (i == ret - 1) {
                cout << endl;
                ret <<= 1;
            }
        }
        cout << endl;
    }
};
/* ------------------- セグメント木 ------------------- */

/* ------------------- 組み合わせ ------------------- */
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
/* ------------------- 組み合わせ ------------------- */