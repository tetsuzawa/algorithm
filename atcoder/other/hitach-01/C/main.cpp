#include <bits/stdc++.h>
#define rep(i, n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;
const ll MOD = 1000000007;

using random_type = uint_fast32_t;
random_type xor128() {
    static random_type x = 123456789, y = 362436069, z = 521288629, w = time(0);
    random_type t = x ^ (x << 11);
    x = y;
    y = z;
    z = w;
    w = (w ^ (w >> 19)) ^ (t ^ (t >> 8));
    return w;
}

using key_type = int;
const key_type INF = 1000000000;
const int MAX_N = 1000010;  // 挿入の回数 + 少し

struct Node {
    key_type key;
    random_type priority;
    Node *left, *right;
    int size;
    key_type min;

    static Node *const nil;
    static int node_count;

    Node() {}

    Node(const key_type &x) : Node(x, xor128(), nil, nil, 1, x) {}

    Node(const key_type &key_, random_type priority_, Node *left_, Node *right_, int size_, const key_type &min_)
        : key(key_), priority(priority_), left(left_), right(right_), size(size_), min(min_) {}

    void *operator new(std::size_t) {
        static Node pool[MAX_N];
        return pool + node_count++;
    }

    static void delete_all() { Node::node_count = 1; }
};

using np = Node *;
using npair = std::pair<np, np>;

int Node::node_count = 0;
Node *const Node::nil = new Node(key_type(), std::numeric_limits<random_type>::min(), nullptr, nullptr, 0, INF);

void verify(np n, int depth = 0) {
    if (n == Node::nil) return;
    assert(n->size == 1 + n->left->size + n->right->size);
    assert(n->min == std::min({n->left->min, n->key, n->right->min}));
    assert(n->left != nullptr);
    assert(n->right != nullptr);
    assert(n->priority >= n->left->priority);
    verify(n->left, depth + 1);
    assert(n->priority >= n->right->priority);
    verify(n->right, depth + 1);
}

int get_depth(np n) {
    if (n == Node::nil) return 0;
    return 1 + std::max(get_depth(n->left), get_depth(n->right));
}

np update(np n) {
    n->size = 1 + n->left->size + n->right->size;
    n->min = std::min({n->left->min, n->key, n->right->min});
    return n;
}

np make_tree(key_type *left, key_type *right) {
    int sz = right - left;
    if (sz == 0) return Node::nil;
    key_type *mid = left + sz / 2;
    np lc = make_tree(left, mid);
    np rc = make_tree(mid + 1, right);
    return new Node(*mid, -1, lc, rc, sz, std::min({lc->min, *mid, rc->min}));
}

np make_treap(key_type *left, key_type *right) {
    np root = make_tree(left, right);
    int n = right - left;
    std::vector<random_type> ps(n);
    std::generate(ps.begin(), ps.end(), xor128);
    std::sort(ps.begin(), ps.end());
    std::queue<np> que;
    que.push(root);
    while (que.size()) {
        np n = que.front();
        que.pop();
        random_type pri = ps.back();
        ps.pop_back();
        n->priority = pri;
        if (n->left != Node::nil) que.push(n->left);
        if (n->right != Node::nil) que.push(n->right);
    }
    return root;
}

key_type range_min(np n, int l, int r) {
    l = std::max<int>(l, 0);
    r = std::min(r, n->size);
    if (l >= r) return INF;
    if (l == 0 && r == n->size) return n->min;
    key_type res = INF;
    int sl = n->left->size;
    res = std::min(res, range_min(n->left, l, r));
    if (l <= sl && sl < r) res = std::min(res, n->key);
    res = std::min(res, range_min(n->right, l - sl - 1, r - sl - 1));
    return res;
}

np merge(np l, np r) {
    if (l == Node::nil) return r;
    if (r == Node::nil) return l;
    if (l->priority > r->priority) {
        l->right = merge(l->right, r);
        return update(l);
    } else {
        r->left = merge(l, r->left);
        return update(r);
    }
}

npair split(np n, int k) {
    if (n == Node::nil) return {Node::nil, Node::nil};
    if (k <= n->left->size) {
        npair p = split(n->left, k);
        n->left = p.second;
        return npair(p.first, update(n));
    } else {
        npair p = split(n->right, k - n->left->size - 1);
        n->right = p.first;
        return npair(update(n), p.second);
    }
}

void set_tree(np n, int k, const key_type &x) {
    if (k < n->left->size)
        set_tree(n->left, k, x);
    else if (n->left->size == k)
        n->key = x;
    else
        set_tree(n->right, k - n->left->size - 1, x);
    update(n);
}

key_type get(np n, int k) {
    if (k < n->left->size)
        return get(n->left, k);
    else if (n->left->size == k)
        return n->key;
    else
        return get(n->right, k - n->left->size - 1);
}

np lower_bound(np n, const key_type &x) {
    if (n == Node::nil)
        return Node::nil;
    else if (n->key >= x) {
        np res = lower_bound(n->left, x);
        return res != Node::nil ? res : n;
    } else
        return lower_bound(n->right, x);
}

np upper_bound(np n, const key_type &x) {
    if (n == Node::nil)
        return Node::nil;
    else if (n->key > x) {
        np res = upper_bound(n->left, x);
        return res != Node::nil ? res : n;
    } else
        return upper_bound(n->right, x);
}

int main() {
    int N;
    cin >> N;
    vector<int> a(N - 1), b(N - 1);
    rep(i, N - 1) cin >> a[i] >> b[i];
    rep(i, N - 1) cout << a[i] << endl;

    return 0;
}