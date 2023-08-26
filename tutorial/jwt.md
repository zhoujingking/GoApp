# JWT
    package: golang-jwt
    reference: https://blog.logrocket.com/jwt-authentication-go/
## JWT的结构
JWT由三部分组成：

    Header
    Payload，声明  (Claims)
    Signature，签名
    
    Header由token的类型（"JWT"）和签名算法名称（比如HMAC-SHA256）组成；
    Payload包含着关于用户数据的各种声明；
    Signature就是密钥，用于验证消息在传递过程中有没有被更改，从而确保信息传输的安全。
