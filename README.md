## 做一个简单的类似QQ交谈工具，带注册、登录、拉取房间列表、选择房间、进入房间、群聊、一对一聊天；  
#### 版本更新日志  
- 2018.8.24：完成客户端与服务器端的连接，完成协议的制定；目前尝试用Go语言完成使用协议的逻辑和对MySQL的操作。  
- 2018.8.24.23点晚上：完成Go语言操作MySQL数据库的初始化连接、插入数据和查找数据的逻辑，并在编译时碰到两个bug:  
Server\DBServ\InsertDB.go:6:12: undefined: db  
Server\DBServ\InsertDB.go:10:3: undefined: checkErr  
明天继续coding。  
