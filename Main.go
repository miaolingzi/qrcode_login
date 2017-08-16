package main

import (
	"fmt"
	"net/http"
	"github.com/skip2/go-qrcode"
	"strings"
	"github.com/satori/go.uuid"
)

var mAppID = "ifengfa7cf109aa3af1fa"
var mAppSecret = "d4624c36b6795d1d99dcf0547af5443d"

//缓存用户请求的appid，uuid相关信息
var mUUID_Map map[string]string

func init() {
	mUUID_Map = make(map[string]string)
	//mUUID_Map = map[string]string{}
}

/**
 * 每次重新请求扫码登录功能，用于为每个用户生成初始的uuid
 */
func qrcode_login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	appid := r.Form.Get("appid")
	appid = strings.TrimSpace(appid)
	if appid == "" {
		fmt.Printf("appid == null")
		//todo appid参数错误
	}
	//todo 检查appid
	//qrdb.isValidAppID
	// isValidAppID
	//产生uuid
	uuid := uuid.NewV4().String()
	mUUID_Map[appid] = uuid
	qrcode_generate(w, r)
}

/**
 * 生成二维码
 * params: appid,uuid,timestamp
 */
func qrcode_generate(w http.ResponseWriter, r *http.Request) {
	for k, v := range mUUID_Map {
		fmt.Printf("(key , value) = (%s , %s)\n", k, v)
	}
	var png []byte
	r.ParseForm()
	uuid := r.Form.Get("uuid")
	fmt.Printf("uuid = %s \n", uuid)
	//去掉多余的空格
	uuid = strings.TrimSpace(uuid)
	if uuid == "" {
		fmt.Printf("uuid == null\n")
		//todo uuid为空，不能直接请求获取二维码，需要重新请求qrcode_login接口
	}
	//todo 检查uuid
	fmt.Printf("uuid = %s\n", uuid)
	appid := r.Form.Get("appid")
	//去掉多余的空格
	appid = strings.TrimSpace(appid)
	if appid == "" {
		fmt.Printf("appid == null\n")
		//todo appid参数错误
	}
	//todo 检查appid是否正确
	fmt.Printf("appid = %s\n", appid)
	timestamp := r.Form.Get("timestamp")
	//去掉多余的空格
	timestamp = strings.TrimSpace(timestamp)
	if timestamp == "" {
		fmt.Printf("timestamp == null\n")
		//todo timestamp参数错误
	}
	//todo 检查时间戳
	fmt.Printf("timestamp = %s\n", timestamp)
	//对uuid进行检查等相关操作
	w.Header().Set("Content-Type", "image/png")
	png, err := qrcode.Encode("https://127.0.0.1:8082/connect/confirm", qrcode.High, 256)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Printf("qrcode is %d bytes long\n", len(png))
	}
	w.Write(png)

}

/**
 * 二维码信息
 */
func confirm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
<head>
<title>Hello Go 1.8</title>
</head>
<body>
    hello , welcome to 二维码信息
</body>
</html>
`)
}

/**
 * client确认／取消登录
 */
func confirm_replay(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
<head>
<title>Hello Go 1.8</title>
</head>
<body>
    hello , welcome to client确认／取消登录
</body>
</html>
`)
}

/**
 * server对接用户体系，rpc调用
 */
func query_user(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
<head>
<title>Hello Go 1.8</title>
</head>
<body>
    hello , welcome to server对接用户体系
</body>
</html>
`)
}

/**
 * web轮训server
 */
func qrconnect(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html>
<head>
<title>Hello Go 1.8</title>
</head>
<body>
    hello , welcome to web轮训server
</body>
</html>
`)
}

/**
 * 计算签名是否一致
 */

/**
 * 计算时间戳是否在规定区间
 */

/**
 * 查询确认登录前的扫码次数
 */

/**
 * 检查appid是否正确
 */

/**
 * 检查uuid绑定状态
 */

/**
 *
 */

/**
 *
 */

func main() {
	http.HandleFunc("/connect/qrcode_login", qrcode_login)
	http.HandleFunc("/connect/qrcode", qrcode_generate)
	http.HandleFunc("/connect/confirm", confirm)
	http.HandleFunc("/connect/confirm_replay", confirm_replay)
	http.HandleFunc("/connect/query_user", query_user)
	http.HandleFunc("/connect/qrconnect", qrconnect)
	http.ListenAndServeTLS(":8082", "./qrcode_login/cert.pem", "./qrcode_login/key.pem", nil)
}
