package apis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
	"wetime-go/config"
	"wetime-go/db"
	"wetime-go/models"

	"github.com/gin-gonic/gin"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddUserApi(c *gin.Context) {
	userID := c.Request.FormValue("userID")
	password := c.Request.FormValue("password")
	age := c.Request.FormValue("age")
	sex := c.Request.FormValue("sex")
	url := c.Request.FormValue("avatorURL")

	user := models.User{
		UserID:    userID,
		Age:       age,
		Sex:       sex,
		Password:  password,
		AvatorURL: url,
	}
	err := db.User.SaveUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "E11000") {
			c.JSON(http.StatusOK, gin.H{"respCode": "02", "msg": "账号重复错误"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"respCode": "01", "msg": "InternalServerError"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"respCode": "00", "msg": "ok"})

}

func LoginUserApi(c *gin.Context) {
	userID := c.Request.FormValue("userID")
	password := c.Request.FormValue("password")

	user := models.User{
		UserID:   userID,
		Password: password,
	}
	res, err := db.User.FindUser(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"respCode": "01", "msg": "not found"})
		return
	}
	if res.Password == password {
		c.JSON(http.StatusOK, gin.H{"respCode": "00", "msg": "success", "data": res})
	} else {
		c.JSON(http.StatusOK, gin.H{"respCode": "02", "msg": "fail"})
	}
}

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Print("controller - admin - upload - SaveUploadedFile - file not attached.")
		c.JSON(http.StatusOK, fmt.Sprintf("'%s' uploaded", file.Filename))
		return
	}

	filepath := path.Join("/home/wetime-src", file.Filename)
	err = c.SaveUploadedFile(file, filepath)
	if err != nil {
		log.Printf("controller - admin - upload - SaveUploadedFile:%s", err)
		c.JSON(http.StatusOK, gin.H{"respCode": "02", "msg": err.Error()})
		return
	}
	//文件服务器必须是NginxPort
	c.JSON(http.StatusOK, gin.H{"respCode": "00", "msg": "success", "data": gin.H{"avatorURL": "http://" + config.Conf.Server + ":" + config.Conf.NginxPort + "/" + file.Filename}})
}

func PostFeedApi(c *gin.Context) {
	userID := c.Request.FormValue("userID")
	avatorURL := c.Request.FormValue("avatorURL")
	postURL := c.Request.FormValue("postURL")

	post := models.Post{
		UserID:    userID,
		AvatorURL: avatorURL,
		PostURL:   postURL,
		PostTime:  time.Now().Format("2006-01-02 15:04:05"),
	}
	err := db.Post.SavePost(post)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"respCode": "01", "msg": "InternalServerError"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"respCode": "00", "msg": "ok"})

}

func GetUserApi(c *gin.Context) {
	userID := c.Param("id")

	user := models.User{
		UserID: userID,
	}
	res, err := db.User.FindUser(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"respCode": "01", "msg": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"respCode": "00", "msg": "success", "data": gin.H{"userID": res.UserID, "age": res.Age, "sex": res.Sex, "avatorURL": res.AvatorURL}})

}

func FollowUserApi(c *gin.Context) {
	followID := c.Param("id")
	userID := c.Param("uid")

	follow := models.Follow{
		UserID:     userID,
		FollowerID: followID,
	}
	err := db.Follow.SaveFollow(follow)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"respCode": "01", "msg": "fail"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"respCode": "00", "msg": "success"})

}

func GetFollowApi(c *gin.Context) {
	userID := c.Param("id")

	user := models.Follow{
		UserID: userID,
	}
	res, err := db.Follow.FindFollow(user)

	follows, _ := json.Marshal(res)
	fmt.Printf("follows:%s\n", string(follows))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"respCode": "01", "msg": "not found"})
		return
	}
	userIDs := []string{}
	for _, value := range res {
		userIDs = append(userIDs, value.FollowerID)
	}
	users, err := db.User.FindUsers(userIDs)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"respCode": "01", "msg": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"respCode": "00", "msg": "success", "datas": users})

}

func GetPostApi(c *gin.Context) {
	userID := c.Param("id")

	user := models.Follow{
		UserID: userID,
	}
	res, err := db.Follow.FindFollow(user)

	follows, _ := json.Marshal(res)
	fmt.Printf("follows:%s\n", string(follows))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"respCode": "01", "msg": "not found"})
		return
	}
	userIDs := []string{}
	for _, value := range res {
		userIDs = append(userIDs, value.FollowerID)
	}
	// 包含自己
	userIDs = append(userIDs, userID)
	posts, err := db.Post.FindPost(userIDs)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"respCode": "01", "msg": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"respCode": "00", "msg": "success", "posts": posts})

}
func GetRCTokenApi(c *gin.Context) {
	userID := c.Param("id")

	var baseURL = "http://api-cn.ronghub.com/user/getToken.json"
	var requestURL = fmt.Sprintf("%s?userId=%s", baseURL, userID)
	request, err := http.NewRequest("POST", requestURL, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"respCode": "01", "msg": "fail"})
		return
	}
	request.Header.Set("App-Key", "vnroth0kvoe5o") //R6v0tEcazSo4
	request.Header.Set("Nonce", "henry")
	request.Header.Set("Timestamp", "1644942751")
	request.Header.Set("Signature", "cd7a2eafd21eceacd137402417f5cee8d7642c53") // 上面三个一起sha1运算

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"respCode": "01", "msg": "fail"})
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("http.Do failed,[err=%s][url=%s]", err, baseURL)
		return
	}
	fmt.Printf("rc resp:%s\n", string(body))
	var token models.Token       //这是一个model，struct类型
	json.Unmarshal(body, &token) //解析二进制json，把结果放进TOKEN中
	if token.Code == 200 {
		c.JSON(http.StatusOK, gin.H{"respCode": "00", "msg": "success", "token": token.Token})
		return
	}
	c.JSON(http.StatusOK, gin.H{"respCode": "02", "msg": token.ErrorMessage})

}
