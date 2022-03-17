// Find the maximum value of the sum of the values of up to the first i items, chosen so that their weights do not exceed w.
// There are two cases to consider
// case 1: pick i-th item
// case 2: do not pick i-th item


#include <iostream>
#include <vector>

using namespace std;

template<class T>
void chmax(T &a, T b) {
    if (a < b) {
        a = b;
    }
}

void print2vec(vector <vector<int64_t>> &v) {
    cout << endl;
    for (auto &vv: v) {
        for (auto &vvv:vv) {
            cout << vvv << " ";
        }
        cout << endl;
    }
    cout << endl;
}

int main() {
    int N;
    int64_t W;
    cin >> N >> W;
    vector <int64_t> weight(N), value(N);
    for (int i = 0; i < N; ++i) cin >> weight[i] >> value[i];

    vector <vector<int64_t >> dp(N + 1, vector<int64_t>(W + 1, 0));

    for (int i = 0; i < N; ++i) {
        for (int w = 0; w <= W; ++w) {
            // case 1
            if (w - weight[i] >= 0) {
                chmax(dp[i + 1][w], dp[i][w - weight[i]] + value[i]);
            }
            chmax(dp[i + 1][w], dp[i][w]);
            print2vec(dp);
        }
    }
    cout << dp[N][W] << endl;
}