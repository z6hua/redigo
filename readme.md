# Redigo
基于http的分布式非关系型数据库，其实是用go重构redis的主从复制功能，属于瞎几把写的

## 主从复制
就是分主服务器和从服务器，从机会复制主机的数据，两端进行交互达到数据库状态同步，这样访问每台机子效果是一样的

redis的主从复制涉及持久化的方法就是RDB和AOF可以百度一下

RDB是数据库快照（物理日志），AOF是数据库的命令追加日志（逻辑日志）

主从复制过程

1. slaveof命令 设置作为当前服务器的主机(即当前服务器变成从机)
2. psync命令 与主机同步数据库状态
3. 如果是第一次同步主机会发送RDB给从机
4. 如果是增量同步主机会根据从机的数据偏移量与自己服务器偏移量计算发送相差的数据命令给从机
5. 从机根据主机返回的同步类型是FullReSync(完整重同步) or PartReSync(部分重同步)进行处理
6. 完整重同步就主机的RDB数据载入从机数据库中
7. 部分重同步就按协议读取命令一条一条执行

## 项目使用方法
1. 启动，搞两个

    ` go run main.go -port 8080 `

   ` go run main.go -port 8081 `
2. 访问
    ```
    localhost:8080/set?key=a&val=a
   localhost:8080/getall
   localhost:8081/set?key=b&val=b
    localhost:8081/getall
    这时两个服务器有自己所存储的数据，是不一样的
   
   http://localhost:8080/slaveof?host=127.0.0.1&port=8081
   将8080作为从机 8081作为主机 再访问
   localhost:8080/getall
   数据与8081同步而且如果8081修改 8080也会同步数据
   ```

主要学 routers/repl里接口的过程

