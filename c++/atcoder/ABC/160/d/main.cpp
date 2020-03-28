#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const int INF = 1 << 30;
const ll MOD = 1000000007;

using Graph = vector<vector<int>>;

// 深さ優先探索
vector<bool> seen;
int deapth = 1;
set<pair<int, int>> st;
// vector<int> ks(MOD, 0);
int start;

void dfs(const Graph &G, int v, int k) {
    cout << "deapth: " << deapth << endl;
    seen[v] = true;  // v を訪問済にする
    deapth++;        // 深さをインクリメント

    // v から行ける各頂点 next_v について
    for (auto next_v : G[v]) {
        if (seen[next_v]) continue;  // next_v が探索済だったらスルー
        if (deapth > k) break;       // next_v が探索済だったらスルー
        if (deapth == k) {
            st.insert(make_pair(min(start, next_v), max(start, next_v)));
        }
        // TODO sokudo
        dfs(G, next_v, k);  // 再帰的に探索
    }
    seen[v] = false;
    deapth--;  // 深さをデクリメント
}

int main() {
    int n, x, y;
    cin >> n >> x >> y;

    // グラフ入力受取 (ここでは無向グラフを想定)
    Graph G(n);  // overを防ぐため+1
    rep(i, n - 1) {
        G[i].push_back(i + 1);
        G[i + 1].push_back(i);
    }
    G[x - 1].push_back(y - 1);
    G[y - 1].push_back(x - 1);

    for (int k = 1; k < n - 1; ++k) {
        for (int i = 0; i < n; ++i) {
            seen.assign(n, false);  // 全頂点を「未訪問」に初期化
            deapth = 0;
            start = i;
            dfs(G, i, k);
        }
        // cout << st.size() << "    sizeeeeee" << endl;
        // ks[k] = st.size();
        // cout << ks[k] << endl;
        cout << st.size() << endl;
        for (auto x : st) {
            cout << x.first << " " << x.second << "\n";  // 要素を順に表示
        }
        st.clear();
    }
    // for (int k = 1; k < n - 1; ++k) {
    //     cout << ks[k] << endl;
    // }

    return 0;
}