/*
 * @lc app=leetcode.cn id=52 lang=cpp
 *
 * [52] N皇后 II
 */

// @lc code=start
class Solution {
public:
    int total = 0;
    bool** vis = nullptr;

    inline bool check(int row, int col, int& n) {
        for(int i = 1; i <= row; ++i) {
            if(vis[row-i][col]) return false;
            if(col-i >= 0 && vis[row-i][col-i]) return false;
            if(col+i <  n && vis[row-i][col+i]) return false;
        }
        return true;
    }

    void dfs(int idx, int& n) {
        if(idx >= n) {
            ++total;
            return;
        }
        for(int i = 0; i < n; ++i) {
            if(!check(idx, i, n)) continue;
            vis[idx][i] = 1;
            dfs(idx+1, n);
            vis[idx][i] = 0;
        }
    }

    int totalNQueens(int n) {
        vis = new bool*[n];
        for (int i = 0; i < n; ++i) {
            vis[i] = new bool[n];
            memset(vis[i], 0, sizeof(bool)*n);
        }

        dfs(0, n);

        return total;
    }
};
// @lc code=end

