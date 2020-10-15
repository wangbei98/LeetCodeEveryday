package com.dailyQue;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.Queue;

/**
 * @program: jz_offer
 * @description: 填充每个节点的下一个右侧节点指针
 * @author: pablo
 * @create: 2020-10-15 15:32
 */
class Node {
    public int val;
    public Node left;
    public Node right;
    public Node next;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, Node _left, Node _right, Node _next) {
        val = _val;
        left = _left;
        right = _right;
        next = _next;
    }
};
//T115方法一：层次遍历(队列)+last指针
public class T116connect {
    public Node connect(Node root) {
        if(root==null){
            return null;
        }
        Queue<Node> queue = new LinkedList<>();
        queue.add(root);
        while (!queue.isEmpty()){
            //思路错了，应该只关注当前的一层
//            Node tmp = queue.poll();
//            for(Node p = tmp; p!=null; p=p.next){
//                if(p.left!=null){
//                    queue.add(p.left);
//                    if(last!=null){
//                        last = p.left;
//                    }
//                    last.next = p.left;
//                    p.left = last;
//                }
//                if(p.right!=null){
//                    queue.add(p.right);
//                    if(last!=null){
//                        last = p.right;
//                    }
//                    last.next = p.right;
//                    p.right = last;
//                }
//            }
            Node last =null;
            int nodeSize = queue.size();
            for(int i=0; i<nodeSize; i++){
                Node f = queue.poll();
                if(f.left!=null){
                    queue.add(f.left);
                }
                if(f.right!=null){
                    queue.add(f.right);
                }
                if (i!=0){
                    last.next = f;
                }
                last = f;
            }

        }
        return root;
    }
}

//T115方法二：用第i层节点去构造第i+1层节点
//public class T116connect {
//    Node last=null, nextStart=null;
//    public Node connect(Node root) {
//        if (root==null){
//            return root;
//        }
//        Node start = root;
//        while(start!=null){
//            last = null;
//            nextStart = null;
//            for (Node p=start; p!=null; p=p.next){
//                if(p.left!=null){
//                    handle(p.left);
//                }
//                if(p.right!=null){
//                    handle(p.right);
//                }
//            }
//            start = nextStart;
//        }
//        return root;
//    }
//
//    private void handle(Node p) {
//        //相当于构造单链表的操作，last往后移了一位
//        if(last!=null){
//            last.next = p;
//        }
//        if (nextStart==null){
//            nextStart = p;
//        }
//        last = p;
//    }
//}
//T116、方法一：层次遍历
//public class T116connect {
//    public Node connect(Node root) {
//        if(root==null){
//            return null;
//        }
//        return root;
//    }
//
//}
////T116、方法二
//public class T116connect {
//    public Node connect(Node root) {
//        if(root==null){
//            return null;
//        }
//        Node leftmost = root;
//        //这一行很重要，当第i+1层没有了，也就不需要在循环了
//        while (leftmost.left!=null){
//            Node head = leftmost;
//            while(head!=null){
//                head.left.next = head.right;
//
//                if(head.next!=null){
//                    head.right.next = head.next.left;
//                }
//                head = head.next;
//
//            }
//            leftmost = leftmost.left;
//        }
//        return root;
//    }
//
//}