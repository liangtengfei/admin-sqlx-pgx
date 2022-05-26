package captcha

import (
	"errors"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

// DataOfCaptcha 验证码信息
type DataOfCaptcha struct {
	Body string `json:"body"` // 验证码内容
	Id   string `json:"id"`   // 验证码编号
}

var store = base64Captcha.DefaultMemStore

func GenerateCaptcha(captchaType string) (DataOfCaptcha, error) {
	driverString := &base64Captcha.DriverString{
		Height:          60,
		Width:           240,
		NoiseCount:      0,
		ShowLineOptions: 2,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 125,
			G: 125,
			B: 0,
			A: 118,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driverChinese := &base64Captcha.DriverChinese{
		Height:          60,
		Width:           320,
		NoiseCount:      0,
		ShowLineOptions: 0,
		Length:          1,
		Source:          "设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,,不想要,的值",
		BgColor: &color.RGBA{
			R: 125,
			G: 125,
			B: 0,
			A: 118,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driverMath := &base64Captcha.DriverMath{
		Height:          40,
		Width:           120,
		NoiseCount:      2,
		ShowLineOptions: 0,
		BgColor: &color.RGBA{
			R: 125,
			G: 125,
			B: 0,
			A: 119,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	var driver base64Captcha.Driver
	if captchaType == "string" {
		driver = driverString.ConvertFonts()
	} else if captchaType == "chinese" {
		driver = driverChinese.ConvertFonts()
	} else if captchaType == "math" {
		driver = driverMath.ConvertFonts()
	}

	c := base64Captcha.NewCaptcha(driver, store)

	id, b64s, err := c.Generate()

	return DataOfCaptcha{
		Body: b64s,
		Id:   id,
	}, err
}

func VerifyCaptcha(id, code string) error {
	if store.Verify(id, code, true) {
		return nil
	}
	return errors.New("验证码错误")
}
