/*
 * @lc app=leetcode.cn id=977 lang=cpp
 *
 * [977] 有序数组的平方
 */

// @lc code=start
class Solution {
public:
    vector<int> sortedSquares(vector<int>& A) {
        vector<int> ans(A.size());
        int L = 0, R = A.size()-1;
        int idx = 0;
        while(L <= R)
            ans[idx++] = -A[L] > A[R] ? A[L] * A[L++] : A[R] * A[R--];
        reverse(ans.begin(), ans.end());
        return ans;
    }
};
// @lc code=end

