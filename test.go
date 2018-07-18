/**
administrator: Li yuan
datetime     : 2018/6/28 11:57
*/
package main
import (
	"net/http"
	"io/ioutil"
	"regexp"
	"fmt"
)
func main() {
	url:="https://movie.douban.com/subject/26925317/"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	web:=string(bytes)
	str1:=`<span\s*property="v:itemreviewed">(.*)</span>`
	str2:=`<strong\s*class="ll\s*rating_num"\s*property="v:average">(.*)</strong>`
	str3:=`<a.*rel="v:directedBy">(.*?)</a>`
	compil:= regexp.MustCompile(str1)
	compil2:= regexp.MustCompile(str2)
	compil3:= regexp.MustCompile(str3)
	pianming := compil.FindAllStringSubmatch(web, -1)
	pingfen := compil2.FindAllStringSubmatch(web, -1)
	daoyan := compil3.FindAllStringSubmatch(web, -1)
	fmt.Println("抓区的电影名："+pianming[0][1])
	fmt.Println("抓区的导演："+daoyan[0][1])
	fmt.Println("抓区的评分："+pingfen[0][1])
}


/*func main() {
	url := "http://tieba.baidu.com/千锋"
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	wangye := string(bytes)

	a:=`<i\s*class="icon-good"\s*alt="精品"\s*title="精品"></i>\s*
<a\s*rel="noreferrer"\s*href="(.*)"\s*title="软件测试免费公开课来啦？隔壁老王来测试哈哈"\s*target="_blank"\s*class="j_th_tit\s*">(.*)</a>`

	compil2 := regexp.MustCompile(a)


	pingfen := compil2.FindAllStringSubmatch(wangye, -1)

	fmt.Println("我抓的评分：" + pingfen[0][0])
}
*/
