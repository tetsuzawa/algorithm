// N switches, M bulbs,

// 5 switches
// bulb[1] 1, 2, 5: odd =>
//        5    2    1
//      (off, off, on)
//      (off, on, off)
//      (on, off, off)
//      (on, on, on)
// bulb[2] 2, 3: even =>
//        2   1
//      (on, on)
//      (off, off)
//      5, 4, 3, 2, 1 | res
//     ------------------
//      0  0  0  0  0 | 0
//      0  0  0  0  1 | 1
//      0  0  0  1  0 | 0
//      0  0  0  1  1 | 0
//      0  0  1  0  0 | 0
//      0  0  1  0  1 | 0
//      0  0  1  1  1 | 0
//      0  1  0  0  0 | 0
//      0  1  0  0  1 | 0
//      0  1  0  1  0 | 0
//      0  1  0  1  1 | 0
//      0  1  1  0  0 | 0
//      0  1  1  0  1 | 0
//      0  1  1  1  0 | 0
//      0  1  1  1  1 | 0
//      1  0  0  0  0 | 1
//      1  0  0  0  1 | 0
//      1  0  0  1  1 | 0
//      1  0  1  0  0 | 1
//      1  0  1  0  1 | 0
//      1  0  1  1  0 | 0
//      1  0  1  1  1 | 1
//      1  1  0  0  0 | 1
//      1  1  0  0  1 | 0
//      1  1  0  1  0 | 0
//      1  1  0  1  1 | 1
//      1  1  1  0  0 | 1
//      1  1  1  0  1 | 0
//      1  1  1  1  0 | 0
//      1  1  1  1  1 | 1


#include <iostream>
#include <bitset>
#include <vector>

using namespace std;


void print_vector(vector <vector<int>> vec) {
    for (int i = 0; i < vec.size(); ++i) {
        for (int j = 0; j < vec.at(i).size(); j++) {
            cout << vec.at(i).at(j) << " ";
        }
        cout << endl;
    }
}


int main() {
    int N, M;
    cin >> N >> M;
    vector <bitset<10>> bulb_switches;
    for (int i = 0; i < M; ++i) {
        int k;
        cin >> k;

        bitset<10> switches;
        for (int j = 0; j < k; ++j) {
            int switch_num;
            cin >> switch_num;
            switches |= (1 << (switch_num - 1));
        }
        bulb_switches.push_back(switches);
    }
    vector<int> bulb_light_up_condition;
    for (int i = 0; i < M; ++i) {
        int p;
        cin >> p;
        bulb_light_up_condition.push_back(p);
    }

    int lightening_bulb_count = 0;
    for (int bit = 0; bit < (1 << N); ++bit) {
        bool hasNonLighteningBulb = false;
        for (int i = 0; i < M; ++i) {
            bitset<10> bb (bit);
            if ((bulb_switches.at(i) & bb).count() % 2 != bulb_light_up_condition[i]) {
                hasNonLighteningBulb = true;
                break;
            }
        }
        if (hasNonLighteningBulb) {
            continue;
        }
        lightening_bulb_count++;
    }

    cout << lightening_bulb_count << endl;
}

