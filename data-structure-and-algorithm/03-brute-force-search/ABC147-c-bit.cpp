
//   1 2 3
// 1   1
// 2 1
// 3   0

//   1 2 3
// 1   1 0
// 2 0   1
// 3 1 0

//   1 2
// 1   0
// 2 0


// Assume each combination is correct
// Prepare a table to note whether each combination was explored or not.
// Look for contradictions (has 0 and 1 on same column)
// If there are no contradictions
// Count the max number of righteous (marked as 1)


// 1 2
// each one never have 1 and 0
// length({1,2}) =2
//#include <iostream>
//#include <unordered_set>
//#include <vector>
//
//using namespace std;
//
//
//int main() {
//    int N;
//    cin >> N;
//    unorderd_set<int> combination_table;
//    vector <vector<int >> testimonies(N, vector<int>(N, -1));
//    for (int i = 0; i < N; i++) {
//        int A;
//        cin >> A;
//        for (int j = 0; j < A; j++) { ;
//            int x, y;
//            cin >> x >> y;
//            testimonies.at(i).at(x) = y
//        }
//    }
//
//    for (int bit = 1; bit < (1 << (N - 1)); ++bit) {
//        unorderd_set<int> combination_table2;
//        int tmp_bit = bit;
//        for (tmp_bit / 2) {
//            if (tmp_bit % 2) {
//                combination_table2.;
//                // deal with
//                for ()
//
//            }
//            tmp_bit /= 2
//        }
//    }
//
//    return 0;
//}


// j = between 0 and 2^N
// k = between 1 and N
// k is honest and j as bits 11101 k-1 th digit is same value

#include <iostream>

using namespace std;

int counter(int x) {
    if(x == 0) return 0;
    return counter(x >> 1) + (x & 1);
}

int main() {
    int N;
    cin >> N;

    vector<int> A_testimonies_address(N);
    vector <vector<int>> x_who_said_to_whom(N, vector<int>(N));
    vector <vector<int>> y_who_said_what(N, vector<int>(N));
    for (int i = 0; i < N; i++) {
        int A;
        cin >> A_testimonies_address.at(i)];
        for (int j = 0; j < A_testimonies_address.at(i); j++) { ;
            cin >> x_who_said_to_whom.at(i).at(j) >> y_who_said_what.at(i).at(j);
        }
    }

    int ans = 0
    for (int bit = 1; bit < (1 << N); bits++) {
        bool ok = true;
        for (int i = 1; i <= N; i++) {
            // assume that i-san is honest
            if (!bits & (1 << (i - 1))) continue;
            for (int j = 1; j <= A_testimonies_address.at(i); j++) {
                // iさんが正しいとして、証言がxさんにに行われているかつxさんも正しいと言われていて、食い違いがあったらfalse
                if (((bits >> (x_who_said_to_whom.at(i).at(j))) & 1) ^ y_who_said_what) ok = false;
            }
        }
        if(ok) ans = max(ans,counter(bits));
    }
    cout << ans << endl;
    return;
}