class Solution1 {
//using double index
    public int[] sortedSquares(int[] A) {
        int len = A.length;
        
        int i = 0;
        int j = len - 1;
        int n = len - 1;
        int res[] = new int[len];
        while(i <= j){
            if(Math.abs(A[i]) < Math.abs(A[j])){
                res[n--] = A[j] * A[j];
                j --;
            }
            else{
                res[n--] = A[i] * A[i];
                i ++;
            }
            
        }
        return res;
    }
}

class Solution2 {
//using HashHeap
    public int[] sortedSquares(int[] A) {
        int n = A.length;
        PriorityQueue<Integer> pq = new PriorityQueue<>((a,b) -> a - b);

        for (int i = 0; i < n; i++) {
            pq.add(Math.abs(A[i]));
        }
        int res[] = new int[n];
        for (int i = 0; i < n; i++) {
            int num = pq.poll();
            res[i] = num * num;
        }

        return res;

    }
}
