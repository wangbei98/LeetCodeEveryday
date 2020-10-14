class Solution1002 {
    public List<String> commonChars(String[] A) {
        int[][] dp=new int[A.length][26];
        for(int i=0; i<A.length; i++){
            for(int j=0; j<A[i].length(); j++){
                dp[i][A[i].charAt(j)-'a']++;
            }
        }
        List<String> list=new ArrayList<>();
        for(int j=0; j<26; j++){
            int Min=~(1<<31);
            for(int i=0; i<A.length; i++){
                Min=Math.min(Min, dp[i][j]);
            }
            char c=(char)(j+'a');
            String s=String.valueOf(c);
            while(Min>0){
                list.add(s);
                Min--;
            }
        }
        return list;
    }
}