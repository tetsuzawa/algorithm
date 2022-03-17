#include <iostream>
#include <vector>

using namespace std;

// O(MN)
vector <vector<int>> memo;

bool func(int i, int w, vector<int> &a, int &counter) {
    // base case
    if (i == 0) {
        if (w == 0) {
            ++counter;
            return true;
        } else {
            return false;
        };
    }

    // check memo
    if (memo[i][w] != -1) return memo[i][w];

    // case: choose a[i-1]
    if (func(i - 1, w, a, counter))return memo[i][w] = 1;

    // case: not to choose a[i-1]
    if (func(i - 1, w - a[i - 1], a, counter))return memo[i][w] = 1;

    return memo[i][w] = 0;
}

int main() {
    vector<int> a = {2, 4, 6, 12, 4, 3, 1, 5};
    int N = a.size();
    int w = 16;
    int counter = 0;
    memo.assign(N + 1, vector<int>(w + 1, -1));
    if (!a.empty()) {
        cout << func(a.size() - 1, w, a, counter) << endl;
        cout << func(a.size() - 1, w, a, counter) << endl;
    }
    return 0;
}



