#include <string>
#include <stack>
#include <iostream>
using namespace std;

class Solution {
public:
    bool isValid(string s) {
        std::stack<char> stack = std::stack<char>();
        stack.push("");
        stack.pop("");
        cout << stack.top() <<endl;
        return false;
        for (char c : s) {
            switch (c) {
                case '(':
                case '[':
                case '{':
                    stack.push(c);
                    break;
                case ')':
                    if (stack.empty() || stack.top() != '(') {
                        return false;
                    }
                    stack.pop();
                    break;
                case ']':
                    if (stack.empty() || stack.top() != '[') {
                        return false;
                    }
                    stack.pop();
                    break;
                case '}':
                    if (stack.empty() || stack.top() != '{') {
                        return false;
                    }
                    stack.pop();
                    break;
                default:
                    return false;
            }
        }
        return stack.empty();
    }
};


int main(int argc, char *argv[]) {
    Solution sol = Solution();
    cout << sol.isValid(")") << endl;
}
