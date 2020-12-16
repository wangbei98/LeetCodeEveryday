package 剑指offer.矩阵中的路径;

class Solution {
    public boolean exist(char[][] board, String word) {
        int m = board.length;
        if (m==0)return false;
        int n = board[0].length;
        boolean[][] visited = new boolean[m][n];
        for (int i = 0;i<m;i++){
            for (int j = 0;j<n;j++){
                visited[i][j] = false;
            }
        }
        return backtrack(board,visited,word,0);
    }
    // 回溯算法
    public boolean backtrack(char[][]board,boolean[][]visited,String word,int start){
        
    }
}