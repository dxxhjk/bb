# bb
## 安装
项目目录下使用
```sh
go install .
```
则编译好的可执行文件`bb`安装到了`$GOPATH/bin`下，将其添加到系统路径中
```sh
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
source ~/.bashrc
```
在 bb 下新建两个目录：`bb/adb_log` 和 `bb/result`

## 配置
配置文件是 `bb/congig` 目录下的config文件，将工作目录配置为 `"work_path": "your_path_to_bb"`
`bmc_port` 配置 bmc 的 ssh 端口
`base_ip` 配置工作系统的 ip
`local_port` 配置运行框架的机器的端口号
`soc_num` 配置系统中 soc 的数量（在使用中可以通过标志来手动指定执行任务的 soc 数目，但无法超过配置的范围）
`soc_base_port` 配置系统中 soc 的起始端口号（即全部 soc 端口号为：`soc_base_port + 1` ~ `soc_base_port + soc_num`）

## 使用
bb 命令需配合子命令使用，查看 bb 命令帮助手册：
```sh
bb -h
```
![在这里插入图片描述](https://img-blog.csdnimg.cn/321b837cbb90469facadf3ae88b6793f.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA6LS06LS05paw56eR5aiY,size_20,color_FFFFFF,t_70,g_se,x_16)
### bb upload_file
作用：生成一串命令，用户可以使用其将自己本机的文件上传到服务器的项目目录 `bb/file` 下，便于之后的文件分发。
![在这里插入图片描述](https://img-blog.csdnimg.cn/b385354fc69b4629b9641103e99c9071.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA6LS06LS05paw56eR5aiY,size_20,color_FFFFFF,t_70,g_se,x_16)

在获得命令行之后，用户只需要在本季复制粘贴就行并使用就行。
使用示例：
![在这里插入图片描述](https://img-blog.csdnimg.cn/e9cfc531e5194f8b82bf5f4944560f7b.png)
### bb distribute_file
作用：指定框架所在服务器上的文件或文件夹，将其下发到指定的 soc 上指定的路径中。
其中服务器文件路径必须要指定，其他 flag 的默认值如下图所示，-n 和 -s 的默认值由配置文件中读取。
![在这里插入图片描述](https://img-blog.csdnimg.cn/4c7a30fc4d0f4f99b9c53f5d7db17831.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA6LS06LS05paw56eR5aiY,size_20,color_FFFFFF,t_70,g_se,x_16)
使用示例：
![在这里插入图片描述](https://img-blog.csdnimg.cn/f264aaf071284d07821f93467200d41f.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA6LS06LS05paw56eR5aiY,size_20,color_FFFFFF,t_70,g_se,x_16)
### bb exec
作用：在指定的 soc 中执行指定的命令。

其中在 soc 中执行的命令必须要指定，其他 flag 的默认值如下图所示，-n 和 -s 的默认值由配置文件中读取。如果不使用 -e flag，则默认不开启能耗监控。
![在这里插入图片描述](https://img-blog.csdnimg.cn/ee1bcee4359b48c29c909571cb5352fa.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA6LS06LS05paw56eR5aiY,size_20,color_FFFFFF,t_70,g_se,x_16)
执行的 log 会记录在 `bb/adb_log` 下，每个 soc 有自己的文件夹，以端口号命名。
![在这里插入图片描述](https://img-blog.csdnimg.cn/ea91e35278f84926ab63d152f649fc91.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA6LS06LS05paw56eR5aiY,size_20,color_FFFFFF,t_70,g_se,x_16)
在命令执行失败的时候可以到 soc 文件夹下的 `stderr` 文件查看报错：
![在这里插入图片描述](https://img-blog.csdnimg.cn/70b7fa32aaed4c0a9ff5eec6025af391.png)
![在这里插入图片描述](https://img-blog.csdnimg.cn/e56c177867884acc8e9a8b3d45985d47.png)
使用示例：
![在这里插入图片描述](https://img-blog.csdnimg.cn/fe4137be30784897adeb38704d089c98.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA6LS06LS05paw56eR5aiY,size_20,color_FFFFFF,t_70,g_se,x_16)![在这里插入图片描述](https://img-blog.csdnimg.cn/2086cdcb68094c849cd3ecb37fe317ca.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA6LS06LS05paw56eR5aiY,size_20,color_FFFFFF,t_70,g_se,x_16)
**bmc 的能耗监控如果没有正常结束会一直进行工作，如果程序被 `ctrl+c` 非正常终止，一定要记得去 bmc 杀掉这个进程。**

### bb collect_result
作用：从指定的 soc 中将指定位置的 result 文件或文件夹上传到服器。

其中soc 上文件路径必须要指定，其他 flag 的默认值如下图所示，-n 和 -s 
的默认值由配置文件中读取。

取回的 result 会记录在 `bb/result` 下，每个 soc 有自己的文件夹，以端口号命名。
![在这里插入图片描述](https://img-blog.csdnimg.cn/acd4edf1d7334f4dba4203acd493735b.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA6LS06LS05paw56eR5aiY,size_20,color_FFFFFF,t_70,g_se,x_16)
使用示例：
![在这里插入图片描述](https://img-blog.csdnimg.cn/89baf30a3dc34c2c965bf2400a31f82c.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA6LS06LS05paw56eR5aiY,size_20,color_FFFFFF,t_70,g_se,x_16)
![在这里插入图片描述](https://img-blog.csdnimg.cn/6bf2ddbe54654b4885e59b843784ac87.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBA6LS06LS05paw56eR5aiY,size_20,color_FFFFFF,t_70,g_se,x_16)
