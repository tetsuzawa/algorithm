#include <iostream>
#include <vector>

using namespace std;

class Tribonacci {
public:
    vector<int> memo;

    Tribonacci(int n) {
        memo.assign(n + 10, -1);
    };

    int tribonacci(int n) {
        if (n == 0)return 0;
        if (n == 1)return 0;
        if (n == 2)return 1;

        if (memo.at(n) != -1) return memo.at(n);

        return memo.at(n) = tribonacci(n - 1) + tribonacci(n - 2) + tribonacci(n - 3);
    }
};

int main() {
    Tribonacci t(10);
    for (int i = 0; i <= 10; ++i) {
        cout << "t.toribonacci(10) = " << t.tribonacci(i) << endl;
    }
}