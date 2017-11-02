package main

import (
	"net/http"
	"fmt"
	"log"
	"html/template"
	"mszhe.me/gocaptcha"
)

const (
	dx = 150
	dy = 50
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/get/", Get)
	fmt.Println("服务已启动...")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("tpl/index.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, nil)
}

func Get(w http.ResponseWriter, r *http.Request) {
	captchaImage, err := gocaptcha.NewCaptchaImage(dx, dy, gocaptcha.RandLightColor())
	if err != nil {
		fmt.Println(err)
	}
	captchaImage.DrawNoise(gocaptcha.CaptchaComplexLower)
	captchaImage.DrawTextNoise(gocaptcha.CaptchaComplexLower)
	captchaImage.DrawText(gocaptcha.RandText(4))
	//captchaImage.Drawline(3);
	captchaImage.DrawBorder(gocaptcha.ColorToRGB(0x17A7A7A))
	captchaImage.DrawSineLine()
	//captchaImage.DrawHollowLine()

	//buf := &bytes.Buffer{}
	//err = captchaImage.SaveImage(buf, gocaptcha.ImageFormatJpeg)
	//fileUrl, err := write2file(buf)

	captchaImage.SaveImage(w, gocaptcha.ImageFormatJpeg)
}

//func write2file(buf *bytes.Buffer) (interface{}, error) {
//	filename := util.GenId() + ".jpg"
//
//	dir, err := os.Getwd()
//	if nil != err {
//		return nil, err
//	}
//
//	// 异步执行
//	go func() {
//		staticPath := dir + util.SEPARATOR + "static" + util.SEPARATOR
//		err = util.MkdirIfNotExist(staticPath)
//		if err != nil {
//			logs.Error(err)
//			return
//		}
//		filepath := staticPath + filename
//		file, err := os.OpenFile(filepath, os.O_APPEND, 0666) //打开文件
//		if err != nil {
//			logs.Error(err)
//			return
//		}
//		defer file.Close()
//		// 缓存->文件
//		ioutil.WriteFile(filepath, buf.Bytes(), 0666)
//	}()
//
//	return fmt.Sprintf(captchaDownUrl, filename), nil
//}
