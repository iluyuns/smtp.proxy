# smtp.proxy.client

## 超级简单的邮件代理服务端 给多个客户提供邮件发送服务，避免交付在客户的源码或服务层暴露自己的邮件服务器信息

* 未添加附件发送功能，可以按自己的需要增加 仅支持html格式邮件发送
* 使用 docker compose 运行 , 复制 .env.example为.env文件，修改其中的配置项即可
* 或自行构建镜像运行


## 认证方式
* 通过配置文件中的token进行认证，客户端请求时需要在header中添加Authorization字段：Bearer ${token}


// 完整的请求示例
```js
var api = 'http://localhost:3000/api/send'
var method = 'POST'
var token = 'your token'
var headers = {
    "Content-Type": "application/json",
    "Authorization: 'Bearer {token}'"
}
var body = {
    "from": "i@example.minis.app",
    "to": ["i@example.com"],
    "body":"<h1>这是一个代理测试邮件</h1>", // 目前仅支持html格式
    "subject": "这是一个代理测试邮件"
  }
fetch(api, {
    method: method,
    headers: headers,
    body: JSON.stringify(body)
}).then(res => res.json())
.then(res => console.log(res))
```


.env配置项
这里仅提供了一个 icloud的smtp服务器的配置示例，其他的smtp服务器配置项可以参考对应的smtp服务器配置
官方文档：https://support.apple.com/zh-cn/102525?caller=baiduansbx&cid=baiduansbx
```env
# 服务Token
SMTP_PROXY_TOKEN = AAH-1sadiasbdash7d8ahsudh2383298jd9ma89
# 服务Host 这里是 icloud的smtp服务器
SMTP_PROXY_HOST = smtp.mail.me.com
# 服务端口 这里是 icloud的smtp服务器端口
SMTP_PROXY_PORT = 587
# 服务用户名 这里是 icloud的用户名
SMTP_PROXY_USERNAME = i@icloud.com
# 服务密码 需要 icloud的 开启 两边验证 并且生成一个APP 专用 密码： https://support.apple.com/zh-cn/102525?caller=baiduansbx&cid=baiduansbx
SMTP_PROXY_PASSWORD = 11111

```
