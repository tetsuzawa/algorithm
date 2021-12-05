#include <iostream>
#include <vector>

using namespace std;


void count_753_number_under_k(int64_t k, int64_t current_num, int combination_flag, int &count) {
    if (current_num > k) return;
    if (combination_flag == 0b111) {
        cout << current_num << endl;
        count++;
    }

    // choose 7
    count_753_number_under_k(k, current_num * 10 + 7, combination_flag | 0b100, count);

    // choose 5
    count_753_number_under_k(k, current_num * 10 + 5, combination_flag | 0b010, count);

    // choose 3
    count_753_number_under_k(k, current_num * 10 + 3, combination_flag | 0b001, count);
}

int main() {
    int64_t k = 10000;
    int count = 0;
    count_753_number_under_k(k, 0, 0, count);
    cout << count << endl;
}