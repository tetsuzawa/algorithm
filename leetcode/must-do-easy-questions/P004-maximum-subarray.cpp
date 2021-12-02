#include <vector>
#include <iostream>

using namespace std;


void printSubArray(vector<int> &nums, int low, int high) {
    for (int ind = low; ind <= high; ind++)
        cout << nums[ind] << " ";
    cout << "\n";
}

class Solution {
public:
    void maxSubArray(vector<int> &nums) {
        int globalMaxSum = nums[0], currMaxSum = nums[0];
        int globalStart = 0, globalEnd = 0, currStart = 0;
        for (int ind = 1; ind < nums.size(); ind++) {
            if (currMaxSum < 0)    // same as: currMaxSum + nums[ind] < nums[ind]
                currMaxSum = nums[ind], currStart = ind;
            else currMaxSum += nums[ind];
            // use only below if cond. to find anyone of all subArrays with globalMaxSum
            if (globalMaxSum < currMaxSum)
                globalMaxSum = currMaxSum, globalStart = currStart, globalEnd = ind;
                // use below to find anyone of all maxLen subArrays with globalMaxSum
            else if (globalMaxSum == currMaxSum and globalEnd - globalStart < ind - currStart)
                globalStart = currStart, globalEnd = ind;
        }
        cout << "Following SubArray has maxSum: " << globalMaxSum;
        cout << " and has length: " << globalEnd - globalStart + 1 << " :\n";
        printSubArray(nums, globalStart, globalEnd);
    }
};

int main(int argc, char *argv[]) {
    Solution solution;
    vector<int> nums = {-2,1,-3,4,-1,2,1,-5,4};
    solution.maxSubArray(nums);
    return 0;
}