#include <iostream>
#include <vector>

using namespace std;


class Fibo {
public:
    vector<int> memo;

    Fibo(int n) {
        memo.assign(n, -1);
    }

    int fibo(int n) {
        if (n == 0) return 0;
        if (n == 1) return 1;
        if (memo.at(n) != -1) return memo.at(n);
        return memo.at(n) = fibo(n - 1) + fibo(n - 2);
    }
};


int main() {
    Fibo f = Fibo(500);
    int res = f.fibo(10);
    // must be 55
    cout << res << endl;
    return 0;
}



