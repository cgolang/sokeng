package router

import (
    "fmt"
    "net/http"
    "regexp"
)

type SoKengMux struct{
}

func (p *SoKengMux) ServeHTTP(w http.ResponseWriter, r *http.Request) Handle{
      if r.URL.Path == "/"{ //当用户打开首页

	}

	if r.URL.Path =="/add/" { //当用户提交新坑
		add(w, r)
	}

	if r.URL.Path == "/ticket/" { //当用户审核随机坑
                ticket(w, r)
        } 

	if r.URL.Path == "/admin/" { //当用户是管理组
                admin(w, r)
        }

	if r.URL.Path == "/e/" { //当用户是企业用户
                e8(w, r)
        } 



          

	if regexp.MatchrString("^(\/s)(.*)",r.URL.Path) { //当用户点击开始搜索
		index(w, r)
		return
	}

	http.NotFound(w, r)
	return
}
