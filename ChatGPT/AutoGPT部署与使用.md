日期：2023年4月27日

作者：SayHi

微信：www1298775182 ,添加备注来意 谢谢

Auto-GPT 是一种 AI 代理，它可以使用自然语言给出目标，并通过将其分解为子任务，使用互联网和其他工具进行自动循环来尝试实现它。它使用 OpenAI 的 GPT-4 或 GPT-3.5 API，是使用 GPT-4 执行自主任务的第一个例子之一⁽¹⁾。Auto-GPT 可以生成类似人类的文本、回答问题、翻译语言、总结文本并提供建议等任务⁽²⁾。

------

1. [en.wikipedia.org](https://en.wikipedia.org/wiki/Auto-GPT)
2. [zhuanlan.zhihu.com](https://zhuanlan.zhihu.com/p/623157739)



# 1.获取Open Ai Key

[登录OpenAi账号获取Key](https://platform.openai.com/account/usage )

**用上面的链接进行登录，点击登录**

![1682566175903](AutoGPT使用.assets/1682566175903.png)

**输入ChatGPT的账号**

![1682566218356](AutoGPT使用.assets/1682566218356.png)

**输入密码，点击Continue登录**

![1682566268228](AutoGPT使用.assets/1682566268228.png)

**这个就是目前你账号的剩余额度，我这里总额度是5美元，已经用了0.2美元，剩余的是4.8美元额度**

![1682566351184](AutoGPT使用.assets/1682566351184.png)

**生成Key注意事项**

![1682566622229](AutoGPT使用.assets/1682566622229.png)

**点击生成Key**

![1682566656976](AutoGPT使用.assets/1682566656976.png)

**生成后复制，找地方保存好，建议可以放到微信的收藏里面**

![1682566691402](AutoGPT使用.assets/1682566691402.png)

# 2.安装Python环境

**下载Python安装包**[官方下载地址](https://www.python.org/downloads/windows/)

我是windows 64位系统我下载这个，按照自己的系统版本进行下载安装包

![1682566852252](AutoGPT使用.assets/1682566852252.png)

安装的时候注意勾选添加到系统的环境变量

![1682567380828](AutoGPT使用.assets/1682567380828.png)

**安装python完成**

![1682567493106](AutoGPT使用.assets/1682567493106.png)

**输入`win键 + R` 输入cmd打开控制台（win键就是键盘左下角ctrl右边那个windows的符号键）**

![1682576098150](AutoGPT使用.assets/1682576098150.png)

**检查python是否安装成功**

![1682567772100](AutoGPT使用.assets/1682567772100.png)

# 3.干正事：部署AutoGPT

[Auto-GPT V0.2.2版本下载](https://github.com/Significant-Gravitas/Auto-GPT/archive/refs/tags/v0.2.2.zip)，今天是2023年4月27日，目前官方发布的AutoGPT稳定版本是：Auto-GPT v0.2.2

[Auto-GPT v0.2.2版本介绍](https://github.com/Significant-Gravitas/Auto-GPT/releases/tag/v0.2.2)

[AutoGPT开源项目地址](https://github.com/Significant-Gravitas/Auto-GPT)

**下载Auto-GPT**

![1682575808583](AutoGPT使用.assets/1682575808583.png)

**更改启动文件名称**：把`.env.template`文件名称 更改为 `.env`

![1682576167532](AutoGPT使用.assets/1682576167532.png)

![1682576163396](AutoGPT使用.assets/1682576163396.png)

**可以用记事本打开，也可以用其他的文本编辑工具打开,更改参数：`OPENAI_API_KEY` ,把手上的Key粘贴进去**

****

![1682576420747](AutoGPT使用.assets/1682576420747.png)

**完成粘贴，保存后退出，开始准备跑AutoGPT了**

![1682576457640](AutoGPT使用.assets/1682576457640.png)

**单击选中路径，输入：cmd**:

![1682576540131](AutoGPT使用.assets/1682576540131.png)

**输入: cmd **：

![1682576604921](AutoGPT使用.assets/1682576604921.png)

**进入cmd窗口**:

![1682576691856](AutoGPT使用.assets/1682576691856.png)

**开始安装AutoGPT需要的模块,输入： pip install -r requirements.txt**:  

![1682576743931](AutoGPT使用.assets/1682576743931.png)

**正在安装AutoGPT运行环境,预计20~60秒左右**:

![1682576769648](AutoGPT使用.assets/1682576769648.png)

**AutoGPT运行模块安装完成**

![1682576805674](AutoGPT使用.assets/1682576805674.png)

**开始执行AutoGPT,执行有两种模式： ** 

**执行模式分手动自动：**

- `手动模式`：AutoGPT去访问什么内容需要请示你，每次都需要用户输入：`y` 后才会访问  
- `自动模式`：只需要输入目标，全自动生成，任务请求都是AutoGPT自动的，风险高，对key的损耗大，可能会一直跑下去，不建议使用自动模式

- `手动模式命令`：python -m autogpt  --gpt3only      
- `自动模式命令`:   python -m autogpt --continuous --gpt3only，可以缩写： python -m autogpt --c  --gpt3only

- `查看参数的含义`: python -m autogpt --help

实际运行过程中有两种方法是比较稳妥的：

- `自动模式加允许次数`：python -m autogpt --c -3 --gpt3only ，这个命令的意思是让AutoGPT使用GPT3模型自动请求网页3次，请求3次后停止请求，等待命令

- `手动模式加允许次数`：采用手动模式加授权AutoGPT指定次数的方式运行，比如:y -20  让AutoGPT自动访问20次网站，访问20次后再请求授权是否还可以继续访问

实际使用过程开始我们可以用自动模式跑十几次，然后根据输出的情况,输入`y -n` 次 让AutoGPT进行请求

**接下来，正式开始**

我们开始用手动模式跑一下，输入：python -m autogpt  --gpt3only 回车

![1682577908721](AutoGPT使用.assets/1682577908721.png)

**输入Ai Name ，这个任务的名称是什么，可以随意，我们取一个：市场规划经理**

![1682578474547](AutoGPT使用.assets/1682578474547.png)

**市场规划经理任务的角色是用来干什么的？输入：帮助公司制定和实施市场营销计划以促进销售增长和品牌认知度提高 **

![1682578496969](AutoGPT使用.assets/1682578496969.png)

**接下来有五个Goal目标，Goal1,Goal2,Goal3,Goal4,Goal5,我们输入下面的答案：**

- `Goal 1`: 写一篇关于冰淇淋的市场调查报告
- `Goal 2`: 要求100字
- `Goal 3`: 写完报告后自动转换为中文
- `Goal 4`: 写完报告后自动停止程序
- `Goal 5`: xxx，你可以做补充，也可以不填，直接回车

**完成Goal目标填写后回车，输入:y - 3 ，代表允许AutoGPT自动请求3次网页，请求3次后它会停下来，等待指令**

![1682578738643](AutoGPT使用.assets/1682578738643.png)

**可以看到AutoGPT自动开始请求Google搜索引擎，去网络上检索数据了**

![1682578930955](AutoGPT使用.assets/1682578930955.png)

**跑了三次后，停了下来 ，输入：y -10 继续让它跑**

![1682579068110](AutoGPT使用.assets/1682579068110.png)

**查看阶段性成果，路径 \Auto-GPT-0.2.2\auto_gpt_workspace ,后续完成的结果也会保存在这里，可以自己检查**

![1682579409346](AutoGPT使用.assets/1682579409346.png)



后续跑的时间太长就不演示结果了，目前AutoGPT运行还是不太稳定，有时候AutoGPT会顺利跑出一个结果，但是我跑了大概有6,7次，只有一次会这样，跑完成功自动结束了。大部分场景是很简单的任务（比如我们的这个任务），跑4,5个小时跑出不了结果，生成的结果，也是堆了一堆文件在\Auto-GPT-0.2.2\auto_gpt_workspace路径下面，生成的结果不稳定，程序也不稳定，运行的过程中很多异常，后续官方出0.3版本出来可能会好一些，持续关注中。

# 4.可能会遇到的问题与解决方案

## 4.1访问OpenAi接口异常

**问题提示**：访问openai的api 偶尔提示异常：Remote end closed connection without response 远程端关闭连接，无响应 

**问题原因**：

我猜测这个问题有两种可能：

- Master版本的BUG：开始我用[Master](https://codeload.github.com/Significant-Gravitas/Auto-GPT/zip/refs/tags/v0.2.2) 这个版本会出现这个提示，我以为是公司的代理不稳定，后面我换了我自己的代理没有出现问题
- 使用Auto-GPT V0.2.2稳定版本：使用官方推荐的稳定版本后也没有出现过这个提示，[Auto-GPT V0.2.2版本下载](https://github.com/Significant-Gravitas/Auto-GPT/archive/refs/tags/v0.2.2.zip)

**解决方法**：

1. 换稳定VPN代理
2. 使用Auto-GPT V0.2.2稳定版本，[Auto-GPT V0.2.2版本下载](https://github.com/Significant-Gravitas/Auto-GPT/archive/refs/tags/v0.2.2.zip)

## 4.2[WinError 10060] 主机连接失败

**问题提示**： [WinError 10060] 由于连接方在一段时间后没有正确回答或连接的主机没有反应，连接试失败 

**问题原因**: 猜测是代理的问题，还是代理不给力

**解决方法**：

1. **设置代理请求参数**: 【！！！注意这个7890端口可能每个人使用的VPN代理端口都不同，看下面第三点找这个代理端口的方法】

```mysql
# 指定HTTP代理服务器的地址和端口号
set https_proxy=http://127.0.0.1:7890 #【！！！注意这个7890端口可能每个人使用的VPN代理端口都不同，看下面第三点找这个代理端口的方法】
#指定HTTPS代理服务器的地址和端口号
set http_proxy=http://127.0.0.1:7890 #【！！！注意这个7890端口可能每个人使用的VPN代理端口都不同，看下面第三点找这个代理端口的方法】
#指定使用SOCKS5代理服务器，代理HTTP、HTTPS等所有协议的网络访问
set all_proxy=socks5://127.0.0.1:7890 #【！！！注意这个7890端口可能每个人使用的VPN代理端口都不同，看下面第三点找这个代理端口的方法】
```

 ![1682585189181](AutoGPT部署与使用.assets/1682585189181.png)

**2.写入代理配置到：autogpt\commands\web_selenium.py ，有两处改动，记得看begin add提示**

```python
import autogpt.processing.text as summary
from autogpt.config import Config
from autogpt.processing.html import extract_hyperlinks, format_hyperlinks

FILE_DIR = Path(__file__).parent.parent
CFG = Config()

def browse_website(url: str, question: str) -> tuple[str, WebDriver]:
#【！！！begin add2】
    if 'http_proxy' in os.environ:
        del os.environ['http_proxy']
if 'https_proxy' in os.environ:
    del os.environ['https_proxy']

if 'all_proxy' in os.environ:
    del os.environ['all_proxy']
#【！！！end】
"""Browse a website and return the answer and links to the user

Args:
    url (str): The url of the website to browse
    question (str): The question asked by the user

Returns:
    Tuple[str, WebDriver]: The answer and links to the user and the webdriver
"""
driver, text = scrape_text_with_selenium(url)
add_header(driver)
summary_text = summary.summarize_text(url, text, question, driver)
links = scrape_links_with_selenium(driver, url)

# Limit links to 5
if len(links) > 5:
    links = links[:5]
close_browser(driver)
# 【！！！begin add1】
os.environ["http_proxy"] = "http://127.0.0.1:7890"
os.environ["https_proxy"] = "http://127.0.0.1:7890"
os.environ["all_proxy"] = "socks5://127.0.0.1:7890"
# 【！！！end】
return f"Answer gathered from website: {summary_text} \n \n Links: {links}", driver
```
**3.找到代理服务器端口的三个方法**

**方法1. 直接查看你梯子的软件，里面应该有设置，比如我的是Clash for Windows：**

![1682581135271](AutoGPT使用.assets/1682581135271.png)

**方法2. win10系统搜索代理，打开代理服务器设置**

![1682580722947](AutoGPT使用.assets/1682580722947.png)

**代理端口为：7890**

![1682580773813](AutoGPT使用.assets/1682580773813.png)

**方法3：系统左下角搜索 ie 浏览器**

![1682580848112](AutoGPT使用.assets/1682580848112.png)

**打开ie浏览器**

![1682581017673](AutoGPT使用.assets/1682581017673.png)

## 4.3 无法请求Google返回值

**问题描述：**browse_website函数（Selenium）不工作了，访问Google返回一个空的提示:Error 

**问题原因**: 代理的原因，AutoGPT代理需要设置

**解决方法：**写入代理配置到: `autogpt/workspace.py` ,注意有begin add2,1两处改动

```python
# 更新文件autogpt/workspace.py

from __future__ import annotations
#【！！！begin add2】
import os
#【 ！！！end】
from pathlib import Path

#【！！！begin add1】
os.environ["http_proxy"] = "http://127.0.0.1:7890"
os.environ["https_proxy"] = "http://127.0.0.1:7890"
os.environ["all_proxy"] = "socks5://127.0.0.1:7890"
# 【！！！end 】

from autogpt.config import Config
....
```

## 4.4 提示无法在容器中运行脚本

**问题描述**：在工作空间'\Auto-GPT-0.2.2\auto_gpt_workspace'中执行文件'analyze_smart_door_lock_market.py' 

无法在容器中运行脚本。如果还没有，请安装Docker https://docs.docker.com/get-docker/ 

系统:命令execute_python_file返回:Error:错误，而获取服务器API版本:(2，'CreateFile'， '系统找不到指定的文件. ')

**问题原因**： 可能是AutoGPT需要用到一些脚本需要在Docker环境下面执行，可以安装一下windows桌面版本的Docker

**解决方法**：安装[Docker](https://docs.docker.com/get-docker/ ) ,注意安装好了要进去用一下，配置一些默认的设置，让docker桌面端可以用

![1683268424068](AutoGPT部署与使用.assets/1683268424068.png)

## 4.5 Json转换不合法的问题

**问题描述**：File "D:\Auto-GPT-0.3.0\autogpt\llm\chat.py", line 166, in chat_with_ai
    agent.summary_memory = update_running_summary(
  File "D:\Auto-GPT-0.3.0\autogpt\memory_management\summary_memory.py", line 76, in update_running_summary
    content_dict = json.loads(event["content"])
  File "C:\Users\Administrator\AppData\Local\Programs\Python\Python310\lib\json\__init__.py", line 346, in loads
    return _default_decoder.decode(s)
  File "C:\Users\Administrator\AppData\Local\Programs\Python\Python310\lib\json\decoder.py", line 337, in decode
    obj, end = self.raw_decode(s, idx=_w(s, 0).end())
  File "C:\Users\Administrator\AppData\Local\Programs\Python\Python310\lib\json\decoder.py", line 355, in raw_decode
    **raise JSONDecodeError("Expecting value", s, err.value) from None**
**json.decoder.JSONDecodeError: Expecting value: line 1 column 1 (char 0)**

**问题原因：** 猜测是请求Google返回的Json格式不规范，然后转换的时候出现了问题，异常代码在：Auto-GPT-0.3.0\autogpt\memory_management\summary_memory.py 第76行，我们需要做一下异常json转换的兼容

**解决方案**:更改代码，注意begin add 2,1两处改动，可能会遇到缩进问题，最好直接用我的复制粘贴过去，[参考链接](https://blog.csdn.net/qq_51575088/article/details/126444314)

```python
    for event in new_events:
        if event["role"].lower() == "assistant":
            event["role"] = "you"

            # Remove "thoughts" dictionary from "content"
			#【begin add 2 用#符号注释这行代码】
            # content_dict = json.loads(event["content"])
            #【begin end】
            #【begin add 1 加入下面的代码】
            try:
                content_dict = json.loads(event["content"])
            except JSONDecodeError:
                print("Error: Invalid JSON string encountered. Skipping this input.")
                # Set content_dict to an empty dictionary or any default value you'd like to use.
                content_dict = {}
            #【begin end】
            if "thoughts" in content_dict:
                del content_dict["thoughts"]
            event["content"] = json.dumps(content_dict)

```

## 4.6 OpenAi Token长度超出限制

**问题描述**:openai.error.InvalidRequestError: This model's maximum context length is 4097 tokens. However, you requested 4181 tokens (3144 in the messages, 1037 in the completion). Please reduce the length of the messages or completion.

**问题原因**: GPT3.5的token最长能够支持4097，但是AutoGPT使用了4181的Token进行请求Openai的接口，所以造成了这个问题

**解决方法**: 进去.env配置文件更改FAST_TOKEN_LIMIT=3500，更改SMART_TOKEN_LIMIT=3500

![1683267763211](AutoGPT部署与使用.assets/1683267763211.png)

## 4.7命令返回未知命令

**问题描述**:SYSTEM:  Command returned: Unknown command ''. Please refer to the 'COMMANDS' list for available commands and only respond in the specified JSON format.

**问题原因**：

1.**GPT模型故障**:模型测试执行一个不存在的命令。根据系统提示应该知道。因此，如果仍然尝试，那就是“傻傻”或“AI模型的”局限制性”。我们对此无能为力。

2.**方法列表缺失**:“COMMAND”列表 缺失有效函数

**解决方案**:

1.**设置env参数**: 设置EXECUTE_LOCAL_COMMANDS=true，允许AutoGPT执行本地命令，可能存在未知风险谨慎使用

2.**添加COMMAND列表方法**:

```python
克隆存储库后，您在根目录中有一个名为autogpt.
里面有提到的文件app.py和文件夹commands。
在commands默认命令中，您的实例将可供使用。您创建一个将导出函数的
新文件。 在你添加行root/autogpt/commands/execute_code.pymy_function
root/autogpt/app.py

from autogpt.commands.execute_code import my_function
并编辑elif列出命令的位置。只需搜索以下形式的内容：

    elif command_name == "my_function":
        return my_function(arguments["some_returned_info"])
execute_code.py并且my_function只是示例名称。
```

[参考链接](https://github.com/Significant-Gravitas/Auto-GPT/issues/949)