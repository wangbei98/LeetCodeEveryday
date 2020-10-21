class Solution {
    public boolean isLongPressedName(String name, String typed) {

        int i = 0;
        int j = 0;
        if(name.charAt(0) != typed.charAt(0)) return false;
        while(i < name.length() && j < typed.length()){
            if(name.charAt(i) == typed.charAt(j)){
                i ++;
                j ++;
            }else if(typed.charAt(j) == typed.charAt(j - 1)){
                j ++;
            }else{
                return false;  
            }   
        }

        while(j < typed.length()){
            char c = typed.charAt(j - 1);
            if(typed.charAt(j ++) != c){
                return false;
            }
        }
        return i == name.length() ;
    }
}
