package main

import (
    "fmt"
    _ "fmt"
)

func commonChars(A []string) []string {
    n := len(A)
    if n == 0{
        return []string{}
    }
    resDict := string2map(A[0])
    if n == 1{
        return map2stringArray(resDict)
    }
    for i:=1;i<n;i++{
        cur_dict := string2map(A[i])
        resDict = mergeDict(resDict,cur_dict)
    }
    return map2stringArray(resDict)
}
func string2map(str string)map[string]int{
    dict := make(map[string]int)
    for i:=0;i<len(str);i++{
        cur := str[i:i+1]
        dict[cur]++
    }
    return dict
}
func map2stringArray(dict map[string]int)[]string{
    res := make([]string,0)
    for key,value := range dict{
        for i:=0;i<value;i++{
            res = append(res,key)
        }
    }
    return res
}
func mergeDict(dictA,dictB map[string]int)map[string]int{
    new_dict := make(map[string]int)
    for key,value := range dictB{
        new_dict[key] = min(value,dictA[key])
    }
    return new_dict
}
func min(x,y int) int{
    if x <y {
        return x
    }else {
        return y
    }
}

func main(){
    dict1 := string2map("bella")
    dict2 := string2map("roller")
    for key,value := range dict1{
        fmt.Println(key," " ,value)
    }
    fmt.Println("----------")
    for key,value := range dict2{
        fmt.Println(key," " ,value)
    }
    fmt.Println("----------")
    dict1 = mergeDict(dict1,dict2)
    for key,value := range dict1{
        fmt.Println(key," " ,value)
    }
    fmt.Println("----------")
    strArr := map2stringArray(dict1)
    fmt.Println(strArr)
}