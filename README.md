# bytedance-tiktok
青训营抖音项目

//注意修改了一下middlerwares的一些bug。

更新：增加了cache包，把redis初始化移入cache
     增加 static包，修改了一些东西,
     规范了一下命名规划，
     然后大家在编写的时候，注意一下，命名的规范，包命名的规范，文件的命名规范，结构体的命名规范，
    可以提前去了解一下go语言的命名规范，然后注意各种大小写，大小写影响访问的权限.

注意：如果没有文件读取不到，可以切换绝对路径，

拉取代码后在goland右下角Terminal输入 go mod download 下载依赖包

或者直接在项目根目录控制台 输入 go mod download 下载依赖包..。
