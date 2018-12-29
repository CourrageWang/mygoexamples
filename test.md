# PA认证事件

- [程序安装](#install)
- [配置](#setting)
- [开始测试](#testing)
* [查看日志](#checkinglog)

## <span id="install">程序安装</span>

将附件内的`srun_intf_main` 程序下载，然后上传到服务器的 `/srun/bin `目录。同时确保可执行权。

```linux
chmod +x /srun3/bin/srun_intf_main
```

##<span id ="setting"> 配置</span>
1.配置srun_onlineintf.xml
程序安装好以后，需要预先开启一次，以生成最新的配置文件，命令如下
`[root@localhost ~]# service srun_onlineintfsvr restart `


执行完成后打开配置文件 `/srun3/etc/srun_onlineintf.xml`文件，在文件底部找到PaReceive节点，如下
```
    <PaReceive>
        <Description>此配置用于获取派网事件信息</Description>
        <ServerIp>0.0.0.0</ServerIp>
        <ServerPort>9527</ServerPort>
        <IsEnable>0</IsEnable>
    </PaReceive>
```
请配置`ServerIp`，`ServerPort`分别代表 IP地址、以及监听的端口号。
## <span id ="testing">开始测试</testing>

确保以上步骤成功完成后，即可开始测试。
1.重启在线接口。
`[root@localhost ~]# service srun_onlineintfsvr restart`
2.深澜系统账号认证。
3.查看8081/radius用户在线表，确保用户在线。
4.用认证后的设备开启WIFI热点。
5.用pc以及移动端连接上述的WIFI热点。
6.验证程序是否执行成功
登陆8081/Radius在线表 查看在线详情， 搜索pc_num、以及mobile_num，热点连接后如果有http访问操作会触发pa syslog 此时程序仍会将本次结果记录，因此程序中的pc_num，mobile_num数不一定和真实连接数一致。

## <span id="checkinglog">查看日志</span>
`[root@localhost ~]# tailf /srun3/log/srun_intf_main/srun_onlineintf.log`
