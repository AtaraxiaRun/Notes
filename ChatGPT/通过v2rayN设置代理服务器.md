日期：2023年5月4日

作者：SayHi



# 通过v2rayN设置本地代理服务器

今天有一个需求需要连接OpenAi的接口，请求ChatGPT，功能都完成了，但是网络代理一直不通，有点郁闷，花了点时间，配置了一下v2rayN的本地代理，写了教程记录一下





# 下载v2rayN

[v2rayN-With-Core.zip](https://objects.githubusercontent.com/github-production-release-asset-2e65be/199570071/9e4c84ed-c74d-43b8-a0a9-d1aa8507ec93?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAIWNJYAX4CSVEH53A%2F20230504%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20230504T023033Z&X-Amz-Expires=300&X-Amz-Signature=ff8123a5f881db88ff7f6f088489188b64157bfabaed5467c09ce351aa5df5c8&X-Amz-SignedHeaders=host&actor_id=31299873&key_id=0&repo_id=199570071&response-content-disposition=attachment%3B%20filename%3Dv2rayN-With-Core.zip&response-content-type=application%2Foctet-stream)

下载后双击执行

![1683173926393](通过v2rayN.exe设置代理服务器.assets/1683173926393.png)





# 导入配置

我是通过vless协议进行请求的，我购买的代理供应商给了我一个链接，比如类似：vless://xxxxxxxx-afff-43f0-9532-52db8514d9ec@205.xxx.xxx.xxx:15336?encryption=none&security=none&type=ws&host=xx-xxx.xxx.xx&path=xxxxx

![1683174007096](通过v2rayN.exe设置代理服务器.assets/1683174007096.png)



粘贴后会提示导入成功（重要信息做了模糊处理，见谅）

![1683174329788](通过v2rayN.exe设置代理服务器.assets/1683174329788.png)

# 设置端口

设置本地需要对外的代理端口

![1683174407961](通过v2rayN.exe设置代理服务器.assets/1683174407961.png)

# 设置http代理

我这里设置的是socks:7889，那么我通过Http请求访问接口的时候就是设置代理为：**127.0.0.1:7890** 

![1683174501520](通过v2rayN.exe设置代理服务器.assets/1683174501520.png)