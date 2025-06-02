package main

import (
	"encoding/json"
	"fetchCourses/config"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func getMajor() {
	client := &http.Client{}
	var data = strings.NewReader(`rwlx=1&xklc=3&xkly=1&bklx_id=0&sfkkjyxdxnxq=0&kzkcgs=0&xqh_id=5&jg_id=21&njdm_id_1=2022&zyh_id_1=184&gnjkxdnj=0&zyh_id=184&zyfx_id=DB07C0DD2E3C7919E0530107010AE94E&njdm_id=2022&bh_id=FFB3FB0FF3369A4EE0530107010A61F1&bjgkczxbbjwcx=0&xbm=1&xslbdm=0&mzm=01&xz=4&ccdm=3&xsbj=8396804&sfkknj=1&sfkkzy=1&kzybkxy=0&sfznkx=0&zdkxms=0&sfkxq=1&sfkcfx=0&kkbk=0&kkbkdj=0&sfkgbcx=0&sfrxtgkcxd=1&tykczgxdcs=0&xkxnm=2024&xkxqm=12&kklxdm=01&bbhzxjxb=0&xkkz_id=2A7367754E7C0699E0630207010A42F6&rlkz=0&xkzgbj=0&kspage=1&jspage=10000&jxbzb=`)
	req, err := http.NewRequest("POST", "https://jwgl.whu.edu.cn/xsxk/zzxkyzb_cxZzxkYzbPartDisplay.html?gnmkdm=N253512", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("Origin", "https://jwgl.whu.edu.cn")
	req.Header.Set("Referer", "https://jwgl.whu.edu.cn/xsxk/zzxkyzb_cxZzxkYzbIndex.html?gnmkdm=N253512&layout=default")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36 Edg/133.0.0.0")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("dnt", "1")
	req.Header.Set("sec-ch-ua", `"Not(A:Brand";v="99", "Microsoft Edge";v="133", "Chromium";v="133"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("Cookie", config.Conf.Cookie)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("%s\n", bodyText)
	ioutil.WriteFile("jsons/major.json", bodyText, 0655)

}

func getAllMajorDetails() {
	file, _ := os.ReadFile("jsons/major.json")
	var allJson AllJson
	json.Unmarshal(file, &allJson)

	list := make([]OneJson, 0)
	for _, majorJson := range allJson.TmpList {
		res := getOneMajorDetail(majorJson.KchId)
		list = append(list, res)
		time.Sleep(2 * time.Second)
	}

	marshal, _ := json.Marshal(list)
	ioutil.WriteFile("jsons/major_details.json", marshal, 0655)
}

func getOneMajorDetail(kchId string) OneJson {
	client := &http.Client{}
	var data = strings.NewReader(`rwlx=1&xkly=1&bklx_id=0&sfkkjyxdxnxq=0&kzkcgs=0&xqh_id=5&jg_id=21&zyh_id=184&zyfx_id=DB07C0DD2E3C7919E0530107010AE94E&txbsfrl=1&njdm_id=2022&bh_id=FFB3FB0FF3369A4EE0530107010A61F1&xbm=1&xslbdm=0&mzm=01&xz=4&ccdm=3&xsbj=8396804&sfkknj=1&gnjkxdnj=0&sfkkzy=1&kzybkxy=0&sfznkx=0&zdkxms=0&sfkxq=1&sfkcfx=0&bbhzxjxb=0&kkbk=0&kkbkdj=0&xkxnm=2024&xkxqm=12&xkxskcgskg=0&rlkz=0&cdrlkz=0&rlzlkz=1&kklxdm=01&kch_id=` + kchId + `&jxbzcxskg=0&xklc=3&xkkz_id=2A7367754E7C0699E0630207010A42F6&cxbj=0&fxbj=0`)
	req, err := http.NewRequest("POST", "https://jwgl.whu.edu.cn/xsxk/zzxkyzbjk_cxJxbWithKchZzxkYzb.html?gnmkdm=N253512", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	req.Header.Set("Origin", "https://jwgl.whu.edu.cn")
	req.Header.Set("Referer", "https://jwgl.whu.edu.cn/xsxk/zzxkyzb_cxZzxkYzbIndex.html?gnmkdm=N253512&layout=default")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36 Edg/133.0.0.0")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("dnt", "1")
	req.Header.Set("sec-ch-ua", `"Not(A:Brand";v="99", "Microsoft Edge";v="133", "Chromium";v="133"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-gpc", "1")
	req.Header.Set("Cookie", config.Conf.Cookie)

	return doWebGetterWithRetry(client, req)

}
