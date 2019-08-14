package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

//正弦函数的运用

func main() {
	//图片大小
	const size = 300
	//创建灰度图
	pic := image.NewGray(image.Rect(0, 0, size, size))
	//遍历像素填充颜色
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			pic.SetGray(x, y, color.Gray{Y: 255})
		}
	}
	//从0到最大像素生成x坐标
	for x := 0; x < size; x++ {
		//让sin的值范围在0~2pi之间
		s := float64(x) * 2 * math.Pi / size
		//sin的幅度为一半的像素，向下偏移一半像素并翻转
		y := size/2 - math.Sin(s)*size/2
		pic.SetGray(x, int(y), color.Gray{Y: 0})
	}
	//创建文件
	file, err := os.Create("sin.png")
	if err != nil {
		log.Fatal(err)
	}
	_ = png.Encode(file, pic)
	//关闭文件流
	defer func() {
		file.Close()
	}()
}
