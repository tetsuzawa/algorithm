
// I will create a funciton
// S = string, length = N
// 239852
// what if S = 125
// 1 + 25 = 26, 12+5 = 17, 1 + 2+ 5 = 8
// the answer will be 176 = 125 + 26 + 17 = 8
// 00 = sum(s[0:len(s)])
// 01 = sum(s[0:1, 1:len(s)]
// 10 = sum(s[0:2, 2:len(s)]
// 11 = sum(s[0:1, 1:2, 2:3]


#include <iostream>
#include <bitset>

using namespace std;


// expected: 4237
// expected: 4 + 237
// expected: 42 + 37
// expected: 4 + 2 + 37
// expected: 423 + 7
// expected: 4 + 23 + 7
// expected: 42 + 3 + 7
// expected: 4 + 2 + 3 + 7
int sum_all(string &s) {
    int N = s.length();

    int64_t res = 0;
    for (int bit = 0; bit < (1 << (N - 1)); ++bit) {
        bitset<8> bits(bit);
        cout << bits << endl;

//        tmp: 4
//        tmp: 42
//        tmp: 4 + 2
//        tmp: 423
//        tmp: 4 + 23
//        tmp: 42 + 3
//        tmp: 4 + 2 + 3
        int64_t tmp = 0;
        for (int i = 0; i < N - 1; ++i) {
            tmp *= 10;
            tmp += s[i] - '0';

            if (bit & (1 << i)) {
                cout << "res += " << tmp << endl;
                res += tmp;
                tmp = 0;
            }
        }
// tmp: 4237
// tmp: 237
// tmp: 37
// tmp: 37
// tmp: 7
// tmp: 7
// tmp: 7
// tmp: 7
        tmp *= 10;
        tmp += s.back() - '0';
        cout << "tmp: " << tmp << endl;
        res += tmp;

    }
    return res;
}


int main() {
//    string s = "4723";
//    for (int i = 0; i < 4; ++i) {
//        int l = (int) (s[i] - '0');
//        cout << typeid(s[i]).name() << endl;
//        cout << l << endl;
//    }
    string s1 = "125";
    sum_all(s1);
    int64_t res = sum_all(s1);
    cout << "res=" << res << endl;
    cout << endl;
    string s2 = "4237";
    cout << s2.back() - '0' << endl;
    res = sum_all(s2);
    cout << "res=" << res << endl;
}

