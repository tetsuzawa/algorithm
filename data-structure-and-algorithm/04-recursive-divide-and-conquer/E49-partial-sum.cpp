#include <iostream>
#include <vector>

using namespace std;


// 2^N
bool func(int i, int w, vector<int> &a) {
    // base case
    if (i == 0) {
        return w == 0;
    }

    // case: choose a[i-1]
    if (func(i - 1, w, a))return true;

    // case: not to choose a[i-1]
    if (func(i - 1, w - a[i - 1], a))return true;

    return false;
}

int main() {
    vector<int> a = {2, 4, 6, 12, 4, 3, 1, 5};
    if (!a.empty()){
        cout << func(a.size()-1, 109, a);
    }
    return 0;
}



