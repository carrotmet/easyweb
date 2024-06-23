// Package routes 注册路由
package routes

import (
    "easyweb/app/http/controllers/api/v1/auth"
    "github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

    // 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
    v1 := r.Group("/v1")
    {
        // 注册一个路由
        authGroup := v1.Group("/auth")
        {
            suc := new(auth.SignupController)
            // 判断手机是否已注册
            authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
            // 判断邮箱是否已注册
            authGroup.POST("/signup/email/exist", suc.IsEmailExist)
            //登录路由
            authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)

            // 发送验证码
            vcc := new(auth.VerifyCodeController)
            // 图片验证码，需要加限流
            authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
            // 手机验证码
            authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)

            lgc := new(auth.LoginController)
            // 使用手机号，短信验证码进行登录
            authGroup.POST("/login/using-phone", lgc.LoginByPhone)

            // 支持手机号，Email 和 用户名
            authGroup.POST("/login/using-password", lgc.LoginByPassword)
            //刷新token
            authGroup.POST("/login/refresh-token", lgc.RefreshToken)
            

        }
    }
}