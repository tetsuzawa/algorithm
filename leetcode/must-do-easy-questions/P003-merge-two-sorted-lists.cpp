/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */

#include <vector>
#include <iostream>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;

    ListNode() : val(0), next(nullptr) {}

    ListNode(int x) : val(x), next(nullptr) {}

    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
public:
    ListNode *mergeTwoLists(ListNode *list1, ListNode *list2) {
        if (list1 == nullptr && list2 == nullptr) {
            return nullptr;
        } else if (list1 == nullptr) {
            return list2;
        } else if (list2 == nullptr) {
            return list1;
        }

        vector < ListNode * > v;

        ListNode *n = list1;
        v.push_back(list1);
        while (n->next != nullptr) {
            n = n->next;
            v.push_back(n);
        }

        n = list2;
        v.push_back(list2);
        while (n->next != nullptr) {
            n = n->next;
            v.push_back(n);
        }


        sort(v.begin(), v.end(), [](const ListNode *left, const ListNode *right) {
            return left->val < right->val;
        });

        for (int i = 0; i < v.size() - 1; i++) {
            v.at(i)->next = v.at(i + 1);
        }

        return v.at(0);
    }
}

int main(int argc, char *argv[]) {
    vector<int> v(3);
    cout << "v.size()=" << v.size() << endl;
    v.push_back(5);
    cout << "v.size()=" << v.size() << endl;
    cout << "v.back()=" << v.back() << endl;
    return 0;
}