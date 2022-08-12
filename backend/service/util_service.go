package service

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
)

type UtilServiceApi struct{}

func NewUtilService() *UtilServiceApi {
	return &UtilServiceApi{}
}

func (u *UtilServiceApi) UploadFile(ctx iris.Context) {
	fileType := ctx.FormValue("file_type")
	if fileType != "pic" && fileType != "media" {
		_, _ = ctx.JSON(Response{ErrorCode, "文件类型错误", nil})
		return
	}

	file, info, err := ctx.FormFile("file")
	if err != nil {
		_, _ = ctx.JSON(Response{ErrorCode, "文件读取失败,请检查文件是否有效", nil})
		return
	}

	defer file.Close()
	filename := fmt.Sprintf("%d", time.Now().Unix()) + info.Filename
	out, err := os.OpenFile("./../file/"+fileType+"/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		_, _ = ctx.JSON(Response{ErrorCode, "目录不存在,请检查", nil})
		return
	}

	// 关闭上传之后的文件流
	defer out.Close()

	// 拷贝文件到指定位置
	_, err = io.Copy(out, file)
	if err != nil {
		_, _ = ctx.JSON(Response{ErrorCode, "文件上传失败, 请重试", nil})
		return
	}
	_, _ = ctx.JSON(Response{SuccessCode, "", map[string]string{"file_name": filename}})
}

type PicQuery struct {
	Quality int `url:"quality"`
}

func (u *UtilServiceApi) GetPic(ctx iris.Context) {
	picName := ctx.Params().Get("pic_name")
	picQuery := PicQuery{}
	ctx.ReadQuery(&picQuery)

	if picName == "" {
		return
	}

	ctx.Header("Content-Type", "image/png")
	file, err := ioutil.ReadFile("./../file/pic/" + picName)
	if err != nil {
		return
	}
	data :=compressImageResource(file, picQuery.Quality)
	_, _ = ctx.Write(data)
	out, err := os.OpenFile("./../file/pics/"+picName, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return
	}


	write := bufio.NewWriter(out)
		write.Write(data)
}

func (c *UtilServiceApi) GetMedia(ctx iris.Context) {
	fileName := ctx.Params().Get("media_name")
	filePath := "./../file/media/" + fileName
	f, err := ioutil.ReadFile(filePath) // 根据文件路径获取文件字节数组
	if err != nil {
		return
	}

	var reqRange = ctx.Request().Header.Get("Range") // 获取请求头的Range属性值
	var reqBlockRange []string
	if reqRange != "" && strings.HasPrefix(reqRange, "bytes=") {
		reqBlockRange = strings.Split(strings.Split(reqRange, "=")[1], "-")
		ctx.StatusCode(206) // 设置返回状态码，206表示为断点续传
		ctx.ResponseWriter().Header().Set("status", "206")
	}

	fileSize := len(f) // 计算文件字节长度
	startPosition := 0 // 开始读取的位置
	endPosition := 0   // 结束读取的位置
	if reqBlockRange != nil {
		startPosition, _ = strconv.Atoi(reqBlockRange[0])
		if len(reqBlockRange) > 1 && reqBlockRange[1] != "" { // 如果Range有截止位置，则将截止位置赋值
			tmp, _ := strconv.Atoi(reqBlockRange[1])
			endPosition = tmp
		} else { // 否则读取全部文件
			endPosition = fileSize - 1
		}
	}

	if reqBlockRange != nil { // 如果safari浏览器时
		if len(reqBlockRange) > 1 && reqBlockRange[1] != "" {
			ctx.ResponseWriter().Header().Set("accept-ranges", "bytes")
			ctx.ResponseWriter().Header().Set("access-control-allow-methods", "HEAD, GET, OPTIONS")
			ctx.ResponseWriter().Header().Set("cache-control", "public, max-age=30726563")
			ctx.ResponseWriter().Header().Set("Content-Type", "video/mp4") // 设置内容类型
			ctx.ResponseWriter().Header().Set("Last-Modified", time.Now().String())
			ctx.ResponseWriter().Header().Set("Connection", "keep-alive")
			ctx.ResponseWriter().Header().Set("content-range", fmt.Sprintf("bytes %d-%d/%d", startPosition, endPosition, fileSize)) // 设置本次读取的范围
			ctx.ResponseWriter().Header().Set("Content-Length", fmt.Sprintf("%d", endPosition-startPosition+1))                     // 设置返回的内容长度
			_, _ = ctx.ResponseWriter().Write(f[startPosition : endPosition+1])                                                     // 截取文件对应的字节位置写入response
		} else { // chrome 等
			ctx.StatusCode(200)
			_ = ctx.SendFile(filePath, fileName)
		}
	} else { // chrome 等
		ctx.StatusCode(200)
		_ = ctx.SendFile(filePath, fileName)
	}
}

type WeatherResult struct {
	Now NowResult `json:"now"`
}

type NowResult struct {
	FeelsLike string `json:"feelsLike"`
	Icon      string `json:"icon"`
	Text      string `json:"text"`
}

func (c *UtilServiceApi) GetWeather(ctx iris.Context) {
	url := "https://devapi.qweather.com/v7/weather/now?location=116.83,23.56&key=79d8f249259a45349d539b2d7e7d02b1"
	resp, err := http.Get(url)
	if err != nil {
		_, _ = ctx.JSON(Response{ErrorCode, "", nil})
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	var res WeatherResult
	json.Unmarshal(body, &res)
	text := "今天温度" + res.Now.FeelsLike + "°C," + res.Now.Text
	if strings.Contains(res.Now.Text, "晴") {
		text += ",请注意防晒噢"
	} else if strings.Contains(res.Now.Text, "雨") {
		text += ",请注意带伞噢"
	}
	res.Now.Text = text
	res.Now.Icon = res.Now.Icon + ".svg"
	_, _ = ctx.JSON(Response{SuccessCode, "", res.Now})
}

func compressImageResource(data []byte, quality int) []byte {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		return data
	}
	set := resize.Resize(480, 270, img, resize.Lanczos3)
	buf := bytes.Buffer{}
	if quality == 0 {
		quality = 60
	}
	err = jpeg.Encode(&buf, set, &jpeg.Options{Quality: quality})
	if err != nil {
		return data
	}
	if buf.Len() > len(data) {
		return data
	}
	return buf.Bytes()
}
