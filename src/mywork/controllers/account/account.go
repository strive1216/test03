package account

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"encoding/json"
)

type user struct {
	Name string `form:"name" json:"name" xml:"name" binding:"required" `
	Password string `form:"password" json:"password" xml:"password"`
}


// func1: 处理最基本的GET
func Func1 (c *gin.Context)  {
	name :=c.DefaultQuery("name","中国");
	Id:= c.Query("Id")

	// 回复一个200OK,在client的http-get的resp的body中获取数据
	c.String(http.StatusOK, "hello %v %s",name,Id)
}
// func2: 处理最基本的POST
func Func2 (c *gin.Context) {
	// 回复一个200 OK, 在client的http-post的resp的body中获取数据
	name:=c.PostForm("name");
	password := c.DefaultPostForm("password","123456")

	var userinfo  user
	err := c.BindJSON(&userinfo)
	if err !=nil {
		c.JSON(http.StatusOK,gin.H{"code":"404"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"name":name,"password":password,"userinfo":userinfo})

	//struct 转 json 串
	jsons, errs :=json.Marshal(userinfo) //转换成JSON返回的是byte[]
	if errs != nil {
		fmt.Println(errs.Error())
	}
	fmt.Println(string(jsons))
}

func Func3(c *gin.Context)  {
	var person user
	if  err:= c.ShouldBindJSON(&person); err==nil {
		c.JSON(http.StatusOK,person)
	}else{
		fmt.Println(err.Error())
		c.JSON(http.StatusOK,gin.H{"code":404})
	}
}

func Func4(c *gin.Context)  {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	fmt.Println(string(buf[0:n]))
	//c.String(http.StatusOK, string(buf[0:n]))
	c.JSONP(http.StatusOK,string(buf[0:n]))
	//resp := map[string]string{"hello": "world"}
	//c.JSON(http.StatusOK, resp)

}