#include<iostream>
#include<bitset>

using namespace std;

int count_bit_set(int num) {
    if (num == 0) return 0;
    bitset<8> numb(num);
    cout << num << (num&1) << " " << numb << endl;
    return (num & 1) + count_bit_set(num >> 1);
}

int main(){
    int a = 0b10110110;
    int cnt = count_bit_set(a);
    bitset<8> b(a);

    cout << b << " set count= " << cnt << endl;
}
