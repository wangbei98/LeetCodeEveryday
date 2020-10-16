/*
 * @lc app=leetcode.cn id=215 lang=cpp
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        auto swiftdown = [&](int idx) {
            while(idx < k) {
                int lc = (idx << 1) | 1;
                int rc = (idx << 1) + 2;
                int cur = idx;
                if(lc < k && nums[cur] > nums[lc])
                    cur = lc;
                if(rc < k && nums[cur] > nums[rc])
                    cur = rc;
                if(cur != idx) {
                    swap(nums[idx], nums[cur]);
                    idx = cur;
                }
                else break;
            }
        };
        auto init = [&]() {
            for(int i = k-1; i > 0; --i) {
                int o = (i - 1) >> 1;
                if(nums[o] > nums[i])
                    swiftdown(o);
            }
        };
        init();
        for(int i = k; i < nums.size(); ++i) {
            if(nums[i] <= nums[0]) continue;
            swap(nums[i], nums[0]);
            swiftdown(0);
        }
        return nums[0];
    }
};
// @lc code=end

