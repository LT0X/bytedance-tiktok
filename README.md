
极简版抖音（以互动接口为主）
====

技术选型与相关开发文档
----
### 技术选型：<br>
* 框架：<br>
     1、go 语言<br>
     首先由于来青训营就是为了学习go的，所以本次的主要开发语言选择了go语言。<br>
  2、Gin<br>
  对于本次的开发时间和人数，我们是选择了gin作为开发时候的框架，简单易用也是其特点。<br>
  3、Gorm<br>
  orm框架我们是选用了gorm，gorm作为一个优秀的orm框架，各方面的性能都不错，而且易用，同时文档部分也是很完整。<br>
* 中间件：<br>
  1、Redis<br>
  中间件选用了Redis，Redis作为一款高速的中间件，可以有效的为系统提供缓存，以此来提高系统的吞吐量。<br>
* 数据库：<br>
  1、Mysql<br>
  数据库是选用了MYSQL数据库，MySQL作为一个大家很常用的数据库，各方面的都很优秀，所以此次的数据库选择了MYSQL。<br>
* 其他工具依赖<br>
  引用了jwt来进行token的分发和鉴定，然后就是对视频的处理是使用了ffmpeg来对视频的封面处理，同时对密码的加密是用了，然后是运用了corn来对Redis的数据定时持久化到数据库中。<br>
  
### 文档参考：  
1、https://go.dev/ref/spec  
2、https://learnku.com/docs/gin-gonic/2018/gin-readme/3819  
3、https://gorm.io/docs/generic_interface.html  
4、https://redis.io/docs/  

架构设计
----
![](https://github.com/LT0X/bytedance-tiktok/blob/lhh/static/%E9%9D%92%E8%AE%AD%E8%90%A5%E6%9E%B6%E6%9E%84%E8%AE%BE%E8%AE%A1.jpg)  
整体的系统处理请求通过，进行路由的分发，首先请求会通过jwt进行鉴权，进行鉴权以后开始进入controller,进行参数的解析，controller进行参数解析以后，开始进入service进行处理，service层进行业务逻辑的处理，同时调用models层进行数据库的查询和更新，也和缓存配合使用，然后utils工具类进行服务全局。  

项目代码介绍
----
常规的代码实现:  
### 一、controllers层：  
1、解析请求得到相关参数。        
  示例代码如下：  
  ```
    //解析请求得到相关参数
  uid := c.Query("user_id")
  id, err := strconv.ParseInt(uid, 10, 64)
  ```
2、调用services层函数进行业务逻辑处理。  
  示例代码如下：  
  ```
    //开始业务逻辑，查找对应的UserInfo
  res, err := info_service.NewUserInfoService(handler.Uid).Do()
  ```  
3、返回相关响应数据。  
  示例代码如下：  
  ```
    // 开始返回响应数据
  handler.sendResponse(0, "success")
  ```
  
### 二、services层：
services层的函数主要进行两大类别的操作：  
1、对得到的参数进行业务逻辑处理，调用models层相关函数实现增删改。  
  示例代码如下：  
  ```
    //数据处理
    msg := models.Message{
        FromUserId: a.FromUserId,
        ToUserId:   a.ToUserId,
        Content:    a.Content,
        CreateTime: time.Now(),
     }
    //调用models层函数
    err := models.GetMessageDao().AddUserMessage(&msg)
  ```
  2、对得到的参数进行业务逻辑处理，调用models层相关函数实现查询操作，并返回相应数据。  
  示例代码如下：  
  ```
      //查询用户聊天录
     response := new(MessageResponse)
     messageList, err := models.GetMessageDao().QueryUserMessage(m.Uid, m.ToUserId)
      //返回数据
     response.MessageList = messageList
   ```
   
### 三、models层：  
根据不同模块提供增删改查方法。  


项目总结与反思
----
### 已经识别出的问题和技术改进：  
首先就是一些并发的问题，由于时间的紧迫，以及要应对学校的开学考试，以及团队成员的各个技术能力，加上之前没有处理并发的经验，所以这一部分并不是考虑了很多，往后的话可以通过进行一系列的测试，来进行并发的优化来提高吞吐量。  
然后的话就是另外的一块内容就是微服务，微服务是在上字节的内部课的时候了解到了，如果还要升级的话，可以往微服务的时候去升级。  
然后在架构方面，我们是选用常规的三层架构，controller–service–model ，这是我们最常用的架构，往后的优化可以寻找出更优的模型，然后就是设计模式的优化。  
可以用Nginx来进行负载均衡。  
可以用MQ来提高项目的高可用性和稳定性。  
### 总结反思：  
经过了寒假的假日的学习，也算是对go有了一定的了解，相比于之前学习的Java，这个go有着更轻量级的表现和不俗的开发开发效率，在这几天中，也是团队的小伙伴完成了一系列的开发，虽然其中发生了不少的事情，但也是成功开发出来但是由于其他事物的打扰以及时间的不足难免会有不足。  
这个项目完成以后，首先还是学到了很多，可以很好的发挥团队成员的主观能动性。作为我和团队的第一个go项目，也是极其的合适的，也认我对这个有了很多的体会，不像Java一样的臃肿，更多的是简洁，迅速，本次也是第一次和团队合作开发，虽然人数不多，但受益匪浅。  
最后，本次的青训营也进入了尾声，是一段不错的经历，也希望我们都能成为更好的自己。  
