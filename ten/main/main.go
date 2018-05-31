package main

import "go-example-for-live/ten/app"

func main() {
	NewApp := app.NewConfigInstance("2")
	NewApp.Instance("git@github.com:wuxiaoxiaoshen/Resume.git")


        var dictMap = map[int32]string{}

        dictMap[0x8c22] = "xiè"
        dictMap[0x4f1f] = "wěi"

        var a = "谢伟" 
        pys := [][]string{}
        for _, v := range a { 
                py, ok := dictMap[v]
                pyr := []string{}
                if ok {
                        fmt.Println(py)
                        pyr = strings.Split(py, ",")
                }   
                pys = append(pys, pyr)
        }   

        fmt.Println(pys)




}
