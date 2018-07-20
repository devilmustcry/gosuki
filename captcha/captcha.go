package captcha

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	captchaString, result := generateCaptcha()
// 	fmt.Printf("%s = %d", captchaString, result)
// }

type Random struct {
	firstCaptcha    int
	secondCaptcha   int
	operationRandom int
}

var timeJa int64 = 10

func random(min, max int) int {
	timeJa = timeJa + 1
	rand.Seed(time.Now().Unix() + timeJa)
	return rand.Intn(max-min) + min
}

func (r Random) GenerateCaptcha() (string, int) {
	result := 0
	operator := ""
	switch r.operationRandom {
	case 0:
		result = r.firstCaptcha + r.secondCaptcha
		operator = "+"
	case 1:
		result = r.firstCaptcha - r.secondCaptcha
		operator = "-"
	case 2:
		result = r.firstCaptcha * r.secondCaptcha
		operator = "*"
	default:
		result = r.firstCaptcha / r.secondCaptcha
		operator = "/"
	}

	captchaString := fmt.Sprintf("%d %s %d", r.firstCaptcha, operator, r.secondCaptcha)
	return captchaString, result
}
