package auth

import (
	v1 "easyweb/app/http/controllers/api/v1"
	"easyweb/app/requests"
	"easyweb/pkg/captcha"
	"easyweb/pkg/logger"
	"easyweb/pkg/response"
    "easyweb/pkg/verifycode"

	"github.com/gin-gonic/gin"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
    v1.BaseAPIController
}

// ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
    // 生成验证码
    id, b64s,answer, err := captcha.NewCaptcha().GenerateCaptcha()
    // 记录错误日志，因为验证码是用户的入口，出错时应该记 error 等级的日志
    logger.LogIf(err)
	//测试answer
	println(answer)
    response.JSON(c, gin.H{
        "captcha_id":    id,
        "captcha_image": b64s,
    })
}



// SendUsingPhone 发送手机验证码
func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {

    // 1. 验证表单
    request := requests.VerifyCodePhoneRequest{}
    if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
        return
    }

    // 2. 发送 SMS
    if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
        response.Abort500(c, "发送短信失败~")
    } else {
        response.Success(c)
    }
}