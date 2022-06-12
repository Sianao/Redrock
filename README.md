# Redrock
象棋客户端实现
##  采用 websocket 通过客户端与服务端建立连接实现通信
~~~go
注册 /register  
登录 /login
 // 额 由于那个客户端太折磨人了 所以 在测试的时候
//  需要 用http://tool.pfan.cn/websocket 吧
//  符合象棋下子的规律
//  可以开房间进入
// 可以在准备好与没准备好之间切换
 // 环境 的话 就用了个mysql redis 也没用上 那个将军将死的时候的判断有问题
 // 什么 docker grpc  一点没用上
 // (属实是我太垃圾了)
~~~
<img src="./utils/img.png" alt="li"/>