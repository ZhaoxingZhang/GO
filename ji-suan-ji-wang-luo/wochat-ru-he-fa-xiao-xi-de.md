# wochat如何发消息的？



> 声明：是wochat（我信）不是wechat（微信），仅仅是帮助理解知识点，不要钻牛角尖哈。

## 背景

情人节赵兴下班比往常早了3个小时，因为在中午想起今天是情人节的时候，他就用wochat给晓薇发了留言“宝贝，今晚稍等我下，我去公司接你”。启动车子的时候，赵兴还在构想浪漫之夜的餐单....

wochat如何让赵兴给晓薇发了留言呢？

* 连接wifi，获得IP
  * 接入Wi-Fi：[NAT](https://www.jianshu.com/p/62028875d53e)
  * 获得IP：DHCP
* 登陆wochat: 
  * 解析服务器IP：DNS
  * 静态资源获取：[CDN](CDN是什么？使用CDN有什么优势？%20-%20阿里巴巴淘系技术的回答%20-%20知乎%20https://www.zhihu.com/question/36514327/answer/1604554133)
  * 发送登陆请求：HTTPS TLS TCP OPTION POST
  * 发送验证码：运营商直接提供的短信API
  * 设备注册：cookie，session，[ticket](https://leancloud.cn/docs/realtime-guide-senior.html#hash-1969632104)
  * 建立长连接：websocket
  * [网络状态检测](https://leancloud.cn/docs/realtime-guide-beginner.html#hash-455150906)：心跳机制
* 发消息
  * [创建会话、发消息](https://leancloud.cn/docs/realtime-guide-beginner.html#hash-1439227251)
* 消息通知
  * [离线推送](https://leancloud.cn/docs/realtime-guide-intermediate.html#hash-485620600)：iOS 和 Android 分别提供了内置的离线消息推送通知服务/应用被切换后台&杀死
  * 在线推送：应用在前台，通过websocket发送消息通知
  * [主动拉取](https://leancloud.cn/docs/realtime-guide-beginner.html#hash-2001347346)：进入会话后，获得没有及时获取到的消息（比如体现推送后进入会话，离线推送是新消息通知，不是消息本身）

**获得IP**

关键词：DHCP

客户机第一次还没有IP地址的时候和服务器使用广播地址进行通信。服务器端使用UDP67端口，客户端使用UDP68端口。

\(1\) 发现阶段：即DHCP客户端寻找DHCP服务器的阶段。

\(2\) 提供阶段：即DHCP服务器提供IP地址的阶段。

\(3\) 选择阶段：即DHCP客户端选择某台DHCP服务器提供的IP地址的阶段。

\(4\) 确认阶段：即DHCP服务器确认所提供的IP地址的阶段

![DHCP&#x8FC7;&#x7A0B;](https://img-blog.csdnimg.cn/20210612195055468.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXNtYW4y,size_16,color_FFFFFF,t_70)

## 登陆

### 域名解析

关键词：DNS，UDP，ARP [DNS域名解析时用的是UDP协议。整个域名解析的过程如下:](https://blog.csdn.net/u012862311/article/details/78753232) 

1. 浏览器向本机DNS模块发出DNS请求，DNS模块生成相关的**DNS报文**；

 2. DNS模块将生成的DNS报文传递给传输层的UDP协议单元；

 3. UDP协议单元将该数据封装成**UDP数据报**，传递给网络层的IP协议单元； 

4. IP协议单元将该数据封装成**IP数据包**，其目的IP地址为DNS服务器的IP地址； 

5. 封装好的IP数据包将传递给数据链路层的协议单元进行发送；

 6. 发送时在ARP缓存中查询相关数据，如果没有，就发送**ARP广播**（包含待查询的IP地址，收到广播的主机检查 自己的IP，符合条件的主机将含有自己MAC地址的ARP包发送给ARP广播的主机）请求，等待ARP回应；

 7. 得到ARP回应后，将**IP地址与路由的下一跳MAC地址**对应的信息写入**ARP缓存表**；

 8. 写入缓存后，以路由下一跳的地址填充目的MAC地址，以数据帧形式转发；

 9. 转发可能进行多次； 

10. DNS请求到达DNS服务器的数据链路层协议单元；

 11. DNS服务器的数据链路层协议单元解析数据帧，将内部的IP数据包传递给网络层IP协议单元；

 12. DNS服务器的IP协议单元解析IP数据包，将内部的UDP数据报传递给传输层UDP协议单元； 

13. DNS服务器的UDP协议单元解析收到的UDP数据报，将内部的DNS报文传递给DNS服务单元； 

14. DNS服务单元将域名解析成对应IP地址，产生**DNS回应报文；** 

15. DNS回应报文-&gt;UDP-&gt;IP-&gt;MAC-&gt;我的主机；

 16. 我的主机收到数据帧，将数据帧-&gt;IP-&gt;UDP-&gt;浏览器； 

17. 将域名解析结果以**域名和IP地址**对应的形式写入**DNS缓存表**。 

![DNS&#x57DF;&#x540D;&#x89E3;&#x6790;](https://img-blog.csdnimg.cn/20210613071532834.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXNtYW4y,size_16,color_FFFFFF,t_70)

### TLS握手

关键词：ssl，tls

### TCP握手

关键词：tcp

### HTTP请求

关键词：http\_code，option，post

### 发送短信验证码

关键词：运营商API

### 登陆，获得session

关键词：session，cookie

### 建立长连接

关键词：websocket

## 发消息

### 负载均衡

关键词：nginx，load balance

### 安全

关键词：SQL injection、XSS、REC

### 消息推送

关键词：voip

