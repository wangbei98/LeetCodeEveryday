package com.dailyQue;

import java.lang.reflect.Array;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

/**
 * @program: jz_offer
 * @description: 找字符串里都出现的字符
 * @author: pablo
 * @create: 2020-10-14 23:47
 */
public class T1002commonChars {
    public List<String> commonChars(String[] A) {
        int[] minfreq = new int[26];
        Arrays.fill(minfreq,Integer.MAX_VALUE);
        for (String word:A){
            int[] freq = new int[26];
            int length = word.length();
            for(int i=0; i<length; i++){
                //语法
                char ch = word.charAt(i);
                freq[ch-'a']++;
            }
            for(int i=0; i<26; i++){
                minfreq[i] = Math.min(minfreq[i], freq[i]);
            }
        }
        List<String> ans = new ArrayList<String>();
        for(int i=0; i<26; i++){
            for(int j=minfreq[i]; j>0; j--){
                ans.add(String.valueOf((char)(i+'a')));
            }
        }
        return ans;

    }
}