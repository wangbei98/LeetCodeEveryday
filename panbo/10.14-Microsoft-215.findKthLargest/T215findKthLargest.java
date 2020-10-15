package microsoft;

import java.util.Random;

/**
 * @program: jz_offer
 * @description: 找到数组中第k大的元素
 * @author: pablo
 * @create: 2020-10-14 22:39
 */
public class T215findKthLargest {
//    Random random = new Random();
//    //方法一：快排(分治思想)，然后取倒数第k个元素。快排优化
//    //1、随机化 2、将原来两个递归区间变成1个区间
//    public int findKthLargest(int[] nums, int k) {
//        //注意快排从小到大，所以第k大需一个转换length-k
//        return quickSelect(nums, 0, nums.length-1, nums.length-k);
//    }
//    public int quickSelect(int[] nums, int l, int r, int index){
//        //快排--判断--单区间查找
//        int q = randomPartition(nums, l, r);
//        if(q==index){
//            return nums[q];
//        }else{
//            return q < index ? quickSelect(nums, q+1, r, index) : quickSelect(nums, l, q-1, index);
//        }
//
//    }
//
//    private int randomPartition(int[] nums, int l, int r) {
//        //左闭右开，所以要+1
//        int i = random.nextInt(r-l+1)+l;
//        //随机数跟最右边的元素换一下
//        swap(nums,i,r);
//        return partition(nums,l,r);
//    }
//
//    private int partition(int[] nums, int l, int r) {
//        //因为最右边是随机数，所以主要找最右边这个元素所在的位置
//        int x=nums[r],i=l-1;
//        for(int j=l; j<r; j++){
//            if(nums[j]<x){
//                //把小于随机数的都放在左边
//                swap(nums,++i,j);
//            }
//        }
//        //再把随机数调到其应当在的位置
//        swap(nums,i+1,r);
//        return i+1;
//    }
//
//    private void swap(int[] nums, int i, int r) {
//        int tmp = nums[i];
//        nums[i] = nums[r];
//        nums[r] = tmp;
//    }

    //方法二：建立最大堆，然后删除k-1个节点后，堆顶元素就是第k大的
    public int findKthLargest(int[] nums, int k) {
        int heapSize = nums.length;
        buildMaxHeap(nums, heapSize);
        //i表示最后还出堆的元素，一直出k-1个
        for(int i=nums.length-1; i>=nums.length-k+1; i--){
            //最大元素出堆
            swap(nums, 0 , i);
            heapSize--;
            //调整最大堆
            maxHeapify(nums,0, heapSize);
        }
        return nums[0];
    }

    private void buildMaxHeap(int[] nums, int heapSize) {
        //题解中没有减1
        for(int i=heapSize/2-1;i>=0;i--){
            maxHeapify(nums, i ,heapSize);
        }
    }

    private void maxHeapify(int[] nums, int i, int heapSize) {
        int l = i*2+1, r = i*2+2, largest=i;
        if(l<heapSize && nums[l]>nums[largest]){
            largest=l;
        }
        if(r<heapSize && nums[r]>nums[largest]){
            largest=r;
        }
        //这个条件很巧妙，既有条件作用，也起到递归终结的作用。
        if(largest!=i){
            swap(nums,i,largest);
            maxHeapify(nums, largest,heapSize);
        }

    }
    private void swap(int[] nums, int i, int r) {
        int tmp = nums[i];
        nums[i] = nums[r];
        nums[r] = tmp;
    }

}