package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func API(c *gin.Context) {
	fmt.Printf("get: %+v\n", c.Request.URL)
	bs, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("post: %+v\n", string(bs))
	c.JSON(200, gin.H{
		"msg":  "ok",
		"code": 200,
		"data": string(bs),
	})
}

func main() {
	r := gin.Default()
	// r.GET("/index/api/x", API)
	r.POST("/api/x", API)
	r.StaticFS("/", http.Dir("./admin-null-vue-bulma-dashboard/dist"))
	

	r.Run(":8888")
	// s.HandleFunc("/api/x", func(rw http.ResponseWriter, r *http.Request) {
	// 	fmt.Printf("get: %+v\n", r.URL.Query())
	// 	r.ParseForm()
	// 	fmt.Printf("postform: %+v\n", r.PostForm)
	// 	fmt.Printf("postjson: %+v\n", r.)

	// 	type a struct {
	// 		Code int    `json:"code,omitempty"`
	// 		Msg  string `json:"msg,omitempty"`
	// 		Data string `json:"data,omitempty"`
	// 	}

	// 	json.NewEncoder(rw).Encode(a{200, "ok", "ok"})
	// })
	// s.Handle("/", http.FileServer(http.Dir("./admin-null-vue-bulma-dashboard/dist")))

	// http.ListenAndServe(":8888", s)
}
