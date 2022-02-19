package parser

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/charlsonz/crawler/engine"
	"github.com/charlsonz/crawler/model"
)

var (
	priceTemplateRE = `<a href="/%s/baojia/".*>(\d+\.\d+)</a>`
	nameRE          = regexp.MustCompile(`<title>【(.*)报价_图片_参数】.*</title>`)
	carImageRE      = regexp.MustCompile(`<img class="color_car_img_new" src="([^"]+)"`)
	sizeRE          = regexp.MustCompile(`<li.*车身尺寸.*<em>(\d+[^\d]\d+[^\d]\d+mm)`)
	fuelRE          = regexp.MustCompile(`<li.*工信部油耗.*<em>(\d+\.\d+)L/100km`)
	transmissionRE  = regexp.MustCompile(`<li.*变\s*速\s*箱.*<em>(.+)</em>`)
	engineRE        = regexp.MustCompile(`发\s*动\s*机.*\s*.*<.*>(\d+kW[^<]*)<`)
	displacementRE  = regexp.MustCompile(`<li.*排.*量.*(\d+\.\d+)L`)
	maxSpeedRE      = regexp.MustCompile(`<td.*最高车速\(km/h\).*\s*<td[^>]*>(\d+)</td>`)
	accelerationRE  = regexp.MustCompile(`<td.*0-100加速时间\(s\).*\s*<td[^>]*>([\d\.]+)</td>`)
	urlRE           = regexp.MustCompile(`http://newcar.xcar.com.cn/(m\d+)/`)
)

func ParseCarDetail(contents []byte, url string) engine.ParseResult {
	id := extractString([]byte(url), urlRE)
	car := model.Car{
		Name:         extractString(contents, nameRE),
		ImageURL:     "http:" + extractString(contents, carImageRE),
		Size:         extractString(contents, sizeRE),
		Fuel:         extractFloat(contents, fuelRE),
		Transmission: extractString(contents, transmissionRE),
		Engine:       extractString(contents, engineRE),
		Displacement: extractFloat(contents, displacementRE),
		MaxSpeed:     extractFloat(contents, maxSpeedRE),
		Acceleration: extractFloat(contents, accelerationRE),
	}
	priceRE, err := regexp.Compile(fmt.Sprintf(priceTemplateRE, regexp.QuoteMeta(id)))
	if err == nil {
		car.Price = extractFloat(contents, priceRE)
	}

	result := engine.ParseResult{
		Items: []engine.Item{
			{
				Id:      id,
				Url:     url,
				Type:    "xcar",
				Payload: car,
			},
		},
	}
	carModelResult := ParseCarModel(contents, url)
	result.Requests = carModelResult.Requests

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}

	return ""
}

func extractFloat(contents []byte, re *regexp.Regexp) float64 {
	f, err := strconv.ParseFloat(extractString(contents, re), 64)
	if err != nil {
		return 0
	}

	return f
}
