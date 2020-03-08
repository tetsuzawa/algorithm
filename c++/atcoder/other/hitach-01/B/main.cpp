#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
#define ARRAY_LENGTH(array) (sizeof(array) / sizeof(array[0]))
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

int minElement(int array[], size_t size) {
    assert(array != NULL);
    assert(size >= 1);

    int min = array[0];
    for (size_t i = 1; i < size; ++i) {
        if (min > array[i]) {
            min = array[i];
        }
    }

    return min;
}
template <class T>
void print(const T &value) {
    cout << value << endl;
}

int main() {
    int A, B, M;
    cin >> A >> B >> M;
    int a[A], b[B], x[M], y[M], c[M];
    rep(i, A) cin >> a[i];
    rep(i, B) cin >> b[i];
    rep(i, M) cin >> x[i] >> y[i] >> c[i];

    int min_a = minElement(a, ARRAY_LENGTH(a));
    int min_b = minElement(b, ARRAY_LENGTH(b));
    int min_val = min_a + min_b;
    int tmp;

    rep(i, M) {
        tmp = a[x[i] - 1] + b[y[i] - 1] - c[i];
        if (tmp < min_val) min_val = tmp;
    }
    cout << min_val << endl;

    return 0;
}