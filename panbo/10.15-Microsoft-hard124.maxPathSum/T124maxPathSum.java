package microsoft;

import com.pablo.TreeNode;

/**
 * @program: jz_offer
 * @description: 二叉树中的最大路径和
 * @author: pablo
 * @create: 2020-10-15 17:04
 */
public class T124maxPathSum {
    //类比 简单题 “最大连续子数组” 动态规划
    int maxSum = Integer.MIN_VALUE;
    public int maxPathSum(TreeNode root) {
        maxGain(root);
        return maxSum;
    }

    private int maxGain(TreeNode node) {
        if(node == null){
            return 0;
        }
        //递归计算左右节点的贡献值
        //只有在zuida贡献值大于0时，才会把路径加入进去
        int leftGain = Math.max(maxGain(node.left), 0);
        int rightGain = Math.max(maxGain(node.right), 0);

        //节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
        int value = node.val + leftGain + rightGain;
        //更新答案
        maxSum = Math.max(value, maxSum);
        return node.val + Math.max(leftGain, rightGain);
    }
}