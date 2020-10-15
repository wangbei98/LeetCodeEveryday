/*
 * @lc app=leetcode.cn id=116 lang=cpp
 *
 * [116] 填充每个节点的下一个右侧节点指针
 */

// @lc code=start
/*
// Definition for a Node.
class Node {
public:
    int val;
    Node* left;
    Node* right;
    Node* next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node* _left, Node* _right, Node* _next)
        : val(_val), left(_left), right(_right), next(_next) {}
};
*/

class Solution {
public:
    Node* connect(Node* root) {
        if(!root) return NULL;
        Node* last = NULL;
        Node* Next = root;
        while(Next) {
            if(!last) {
                last = Next;
                Next = last->left;
                continue;
            }
            Node* cur = last->left;
            while(last) {
                cur->next = last->left; cur = cur->next;
                cur->next = last->right; cur = cur->next;
                last = last->next;
            }
            cur->next = NULL;
        }
        return root;
    }
};
// @lc code=end

