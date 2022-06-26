package captcha

import (
	"github.com/mojocn/base64Captcha"
)

var Base64CaptchaStore base64Captcha.Store = RedisStore{}

func init() {

}

// NewDriver 新建base64Captcha配置字符串
func NewDriver() *base64Captcha.DriverString {
	driver := new(base64Captcha.DriverString)
	driver.Height = 44
	driver.Width = 120
	driver.NoiseCount = 5
	driver.ShowLineOptions = base64Captcha.OptionShowSineLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowHollowLine
	driver.Length = 6
	driver.Source = "1234567890qwertyuipkjhgfdsazxcvbnm"
	driver.Fonts = []string{"wqy-microhei.ttc"}
	return driver
}

// GenerateCaptchaHandler 生成图形验证码
func GenerateCaptchaHandler() (string, base64Captcha.Item) {
	// 创建图形验证码的驱动字符串
	driver := NewDriver().ConvertFonts()
	// 创建图形码实例
	c := base64Captcha.NewCaptcha(driver, Base64CaptchaStore)
	// 生成图形验证码的问题和答案
	_, content, answer := c.Driver.GenerateIdQuestionAnswer()
	// 生成图形验证码
	item, _ := c.Driver.DrawCaptcha(content)
	return answer, item
}
