/*
 * @lc app=leetcode.cn id=124 lang=cpp
 *
 * [124] 二叉树中的最大路径和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    pair<int, int> dfs(TreeNode* root) {
        int x = root->val;
        int y = root->val;
        if(!root->left) swap(root->left, root->right);
        if(root->left) {
            auto left = dfs(root->left);
            x = max(x, root->val + left.first);
            y = max(y, max(left.first, left.second));
            if (root->right) {
                auto right = dfs(root->right);
                x = max(x, root->val + right.first);
                y = max(y, max(right.first, right.second));
                y = max(y, left.first + right.first + root->val);
            }
        }
        return make_pair(x, y);
    }
    int maxPathSum(TreeNode* root) {
        if(!root) return INT_MIN;
        auto ans = dfs(root);
        return max(ans.first, ans.second);
    }
};
// @lc code=end

