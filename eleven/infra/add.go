package infra

func Add(argOne int, argTwo int) int {
	return argOne + argTwo
}

func Pinyin(){

        var name string

        flag.StringVar(&name, "n", "谢伟", "") 
        flag.Parse()
        //var a = "谢伟" 
        pys := []string{}
        for _, arg := range flag.Args() {
                for _, v := range arg {
                        py, ok := infra.PinyinDict[v]
                        if ok {
                                //fmt.Println(py)
                                pyList := strings.Split(py, ",")
                                if len(pyList) > 0 { 
                                        pys = append(pys, pyList[0])
                                }   

                        }   
                }   

        }   

        fmt.Println(pys)

}
