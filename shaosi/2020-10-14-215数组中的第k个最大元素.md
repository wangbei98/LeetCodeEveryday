class Solution215{
    private int len;
    private int res=0;
    public int partition(int[] nums, int l, int r){
        Random random=new Random(r-l+1);
        int k=random.nextInt()+l;
        if(k>l && k<=r){
            int num=nums[k];
            nums[k]=nums[l];
            nums[l]=num;
        }
        
        int pos=nums[l];
        while(l<r){
            while(l<r && nums[r]>=pos)r--;
            nums[l]=nums[r];
            while(l<r && nums[l]<=pos)l++;
            nums[r]=nums[l];
        }
        nums[l]=pos;
        return l;
    }
    public void quicksort(int[] nums, int l, int r, int k){
        if(l<=r){
            int pos=partition(nums, l, r);
            if(pos==len-k){
                res=nums[pos];
                return;
            }
            else if(pos>len-k)quicksort(nums, l, pos-1, k);
            else quicksort(nums, pos+1, r, k);
        }
    }
    public int findKthLargest(int[] nums, int k) {
        len=nums.length;
        quicksort(nums, 0, len-1, k);
        return res;
    }
}