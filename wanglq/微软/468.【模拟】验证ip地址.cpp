/*
 * @lc app=leetcode.cn id=468 lang=cpp
 *
 * [468] 验证IP地址
 */

// @lc code=start
class Solution {
public:
    inline bool isIPV4(string& IP) {
        int idx = 0;
        for(int i = 0; i < 4; ++i) {
            if(i == 0 || IP[idx++] == '.') {
                if(!isdigit(IP[idx])) return false;
                if(IP[idx] == '0' && isdigit(IP[idx+1]))
                    return false;
                int num = 0;
                while(isdigit(IP[idx])) {
                    num = num*10 + IP[idx++] - '0';
                    if(num > 255) return false;
                }
            }
            else return false;
        }
        return idx == IP.length();
    }

    inline bool isIPV6(string& IP) {
        int idx = 0;
        auto number = [](char c) {
            if ('0' <= c && c <= '9') return true;
            c = tolower(c);
            if ('a' <= c && c <= 'f') return true;
            return false;
        };
        for (int i = 0; i < 8; ++i) {
            if(i == 0 || IP[idx++] == ':') {
                if(!number(IP[idx])) return false;
                int cnt = 0;
                do ++idx, ++cnt;
                while(number(IP[idx]));
                if(cnt > 4) return false;
            }
            else return false;
        }
        return idx == IP.length();
    }
    string validIPAddress(string IP) {
        if(isIPV4(IP)) return "IPv4";
        if(isIPV6(IP)) return "IPv6";
        return "Neither";
    }
};
// @lc code=end

