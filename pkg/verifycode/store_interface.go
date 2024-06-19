package verifycode


// 验证码存储的通用接口：包括图形验证码、短信验证码都使用这个接口重写
type Store interface {
    // 保存验证码
    Set(id string, value string) bool

    // 获取验证码
    Get(id string, clear bool) string

    // 检查验证码
    Verify(id, answer string, clear bool) bool
}