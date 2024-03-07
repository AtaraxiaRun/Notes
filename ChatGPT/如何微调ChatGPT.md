# 微调ChatGPT

可以自己编写爬虫去心理学网站上面爬取数据，知乎，百度百科，维基百科，贴吧，小红书，抖音，然后对数据进行清洗

现在下载的数据集不能直接使用的，如果用的话，也需要数据清洗过后， 我想到的清洗过程就是把现有的数据，问题，答案分好类，然后用程序综合起来，用ChatGPT重新生成对应问题的答案，然后再训练到ChatGPT里面去，因为答案的质量可能一般，但是问题是对的，可能不需要这么多数据集，可以先拿1000条，看看训练效果，可能有用的噢，用GPT4来回答也很给力，用Free4，问题是引子，然后用Pandora的Token进行回答问题，爬取数据集下来，也是很有意思的，然后训练出来后，可以真正的帮助到大家

kaggle心理健康的数据集：https://www.kaggle.com/datasets/narendrageek/mental-health-faq-for-chatbot

直接上手微调好的心理学模型，并且提供测试：https://betterprogramming.pub/how-to-finetune-gpt-3-finetuning-our-virtual-mental-health-assistant-641c1f3b1ef3

ChatGLM-6B微调方法，，数据的来源，使用微信数据微调方法，有具体操作：https://mp.weixin.qq.com/s/0bQNLqlCSc97QGikse093w

数据集获取方法，数据集提供商：https://mp.weixin.qq.com/s/P_Uk3gri5LElbVzz_IoJ-w

Python教程：https://www.bilibili.com/video/BV1rN4y1F7AU/?spm_id_from=333.337.search-card.all.click&vd_source=2784fc6f010453d9e0bb14a36852e3dd

问题链接：https://chat.zhile.io/c/a1e80c97-1844-4589-8ffe-0ce07db98376

训练模型：https://github.com/imClumsyPanda/langchain-ChatGLM

官方训练微调链接：https://platform.openai.com/docs/guides/fine-tuning

使用ChatGPT进行微调：https://mp.weixin.qq.com/s/0DyUBf6V9jo6uhFqDRRwfQ

120美元： randyhillmr.0.5199.6@gmail.com	24438764 
  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

2023年5月15日  已使用7美元

2023年5月16日  已使用10.78美元

embedding可以理解成快速提供prompt和上下文，是个工程的优化方案。finetune是模型的优化方案，专业领域还是得靠finetune

**微调心理学方面的GPT**

参考资料

- 使用Python微调GPT-3以创建虚拟心理健康助理机器人 https://betterprogramming.pub/how-to-finetune-gpt-3-finetuning-our-virtual-mental-health-assistant-641c1f3b1ef3

1. Kaggle上找到了这个非常方便的数据集[-Chatbot的抑郁数据](https://www.kaggle.com/datasets/narendrageek/mental-health-faq-for-chatbot)、
2. 注意- OpenAI提供了一个有用的 [CLI（命令行界面）](https://platform.openai.com/docs/guides/fine-tuning/cli-data-preparation-tool)工具来准备我们的数据。它可以轻松地将CSV，TSV，XLSX和JSON转换为正确的JSONL格式 
3. 为了获得更好的准确性，Open AI建议至少有150-200个微调示例，数据集大小每增加一倍都会导致模型质量线性增加。
4. Railway是一个云平台，您可以轻松托管和部署任何应用程序。他们提供了一个慷慨的免费层，并拥有极快的部署速度。请确保首先将您的本地repo推送到Github，以便可以部署它。登录到您的铁路帐户，一旦在仪表板，去添加一个新的项目。然后选择upload from Github并选择你的git repo，就这样！您的应用现已上线。(Make确保更新Dialogflow代理中的Webhook URL）



# 微调ChatGPT

微调对数据集的要求很高，由于费用还有时间的关系，本次使用100条数据进行微调（花费1.75美金），100条数据对微调结果的影响是很小的，官方建议数据量最少要500以上，数据量越大越好，教程是为了给大家快速走一个微调流程，希望能够帮到感兴趣的朋友。



本教程使用工具：

Visual Studio Code  ： 所有代码在Vs Code运行

浏览器：Google浏览器

其他：OpenAI Key



微调步骤：

> 选择模型
>
> 制作训练数据
>
> 上传训练数据
>
> 训练专属的微调模型
>
> 使用微调模型

# 选择模型

微调只能选择基础模型，GPT-3.5有这些基础模型：

- Davinci
- Curie
- Babbage
- Ada

Davinci是最贵最好的模型，本次微调使用GPT-3.5 Davinci模型进行微调

# 制作训练数据

微调数据我们就不自己制作了，直接使用开源的心理咨询问答预料库，预料库由斯坦福大学，UCLA 和台湾辅仁大学临床心理学等心理学专业人士参与建设，并由 Chatopera 和诸多志愿者合作完成。

## 下载训练数据

[心理咨询问答预料库下载](https://github.com/702f851f-c126-4bda-bce1-b026f7c2ba20)，下载下来的格式不是符合OpenAI微调的文件格式内容，下面会用编写代码进行转换 

如果你怕麻烦，你也可以直接在下面下载用我转换好的微调数据，就不用转换格式了，直接上传训练数据，不过只有100条，上面的下载地址是原地址，有上万条。本次微调数据100多条:[点击下载](https://www.kdocs.cn/view/l/caK97wk36OOR) ![图片](https://mmbiz.qpic.cn/mmbiz_png/mOxT9icqKx7IKe3WM6wVkia2Lsn3YEiasx8jn2x2Licpx1IxDkTJZdnpwyicFpTfjfXGF8cU6QuVemr0Ry3JEWHjxyA/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

## 转换训练数据

```json
import json

CORPUS_DATA_PATH = "./efaqa-corpus-zh.utf8/EFA_Dataset_v20200314_latest.txt"

with open(CORPUS_DATA_PATH, encoding="utf-8") as f:
    output = []
    for line in f:
        co_data = json.loads(line)
        prompt = co_data["title"]
        completion = ''
        chats = co_data["chats"]
        for chat in chats:
            completion = chat["value"]
            completion = " " + completion.strip()
            line = {'prompt': prompt, 'completion': completion}
            output.append(line)
            break

with open("./data.jsonl", "w", encoding="utf-8") as outfile:
    for item in output:
        outfile.write(json.dumps(item, ensure_ascii=False) + '\n')

 

```

## 检查训练数据

![1684218898816](如何微调ChatGPT.assets/1684218898816.png)

`注：由于openai对jsonl的格式要求非常严格，最后一行多了一个空行要去掉，否则会出现“does not appear to be in valid jsonl format”的错误。`

## 转换官方格式

**安装openai模块**: pip install openai

![1684219536808](如何微调ChatGPT.assets/1684219536808.png)

**安装openai[datalib]模块** :pip install openai[datalib]

![1684219606915](如何微调ChatGPT.assets/1684219606915.png)

**开始转换格式**: openai tools fine_tunes.prepare_data -f './data.jsonl'   ,会有很多输入的，全部输入 y, yes ，一路yes

![1684219867259](如何微调ChatGPT.assets/1684219867259.png)

**检查转换好的训练数据**: 官方处理缩进，空格,加入END

![1684221347869](如何微调ChatGPT.assets/1684221347869.png)

# 上传训练数据

## 上传训练模型的数据 

数据量越大需要的费用，时间越多，本次微调中使用100条数据进行微调

```python
import openai

openai.api_key = "sk-TOXxCdi8cZaxxxxxxlbkFJIpngRRgd3nE1awPjtIXT"

# 上传文件
file = openai.File.create(
    file=open("./data_prepared.jsonl", "rb"),
    purpose='fine-tune'
)

# 输出文件信息
print(openai.File.retrieve(file.id))

#执行返回
{
  "bytes": 22019,
  "created_at": 1684222598,
  "filename": "file",
  "id": "file-ExYVR54RCcoLXeJ7TWYnNnCj",
  "object": "file",
  "purpose": "fine-tune",
  "status": "uploaded",
  "status_details": null
}
```

# 训练专属的微调模型

## 查看训练模型的信息

```python
import openai

openai.api_key = "sk-TOXxCdxxxxxgRRgd3nE1awPjtIXT"

model_info = openai.FineTune.create(
    training_file="file-ExYVR54RCcoLXeJ7TWYnNnCj", model="davinci")
print(model_info)

# 输出
{
  "created_at": 1684222815,
  "events": [
    {
      "created_at": 1684222815,
      "level": "info",
      "message": "Created fine-tune: ft-gm3VC33qcsOybFgV0Ra8ZXFu",
      "object": "fine-tune-event"
    }
  ],
  "fine_tuned_model": null,
  "hyperparams": {
    "batch_size": null,
    "learning_rate_multiplier": null,
    "n_epochs": 4,
    "prompt_loss_weight": 0.01
  },
  "id": "ft-gm3VC33qcsOybFgV0Ra8ZXFu",
  "model": "davinci",
  "object": "fine-tune",
  "organization_id": "org-h30PMzIKL28VraqLFD4fMPMG",
  "result_files": [],
  "status": "pending",
  "training_files": [
    {
      "bytes": 22019,
      "created_at": 1684222598,
      "filename": "file",
      "id": "file-ExYVR54RCcoLXeJ7TWYnNnCj",
      "object": "file",
      "purpose": "fine-tune",
      "status": "processed",
      "status_details": null
    }
  ],
  "updated_at": 1684222815,
  "validation_files": []
}
```

## 查看训练模型的状态

看返回的status,与message这两个属性的结果

```python
import openai

opena	i.api_key = "sk-TOXxxxxxxlbkFJIpngRRgd3nE1awPjtIXT"

model_info = openai.FineTune.retrieve(id="ft-gm3VC33qcsOybFgV0Ra8ZXFu")
print(model_info)

# 过了五分钟再看一下状态，看下面两个属性状态
1."message": "Uploaded model: davinci:ft-personal-2023-05-16-07-51-27" 微调好的模型名称
2."status": "succeeded",   已经微调成功了

{
  "created_at": 1684222815,
  "events": [
    {
      "created_at": 1684222815,
      "level": "info",
      "message": "Created fine-tune: ft-gm3VC33qcsOybFgV0Ra8ZXFu",
      "object": "fine-tune-event"
    },
    {
      "created_at": 1684223200,
      "level": "info",
      "message": "Fine-tune costs $1.54",
      "object": "fine-tune-event"
    },
    {
      "created_at": 1684223200,
      "level": "info",
      "message": "Fine-tune enqueued. Queue number: 0",
      "object": "fine-tune-event"
    },
    {
      "created_at": 1684223201,
      "level": "info",
      "message": "Fine-tune started",
      "object": "fine-tune-event"
    },
    {
      "created_at": 1684223335,
      "level": "info",
      "message": "Completed epoch 1/4",
      "object": "fine-tune-event"
    },
    {
      "created_at": 1684223373,
      "level": "info",
      "message": "Completed epoch 2/4",
      "object": "fine-tune-event"
    },
    {
      "created_at": 1684223412,
      "level": "info",
      "message": "Completed epoch 3/4",
      "object": "fine-tune-event"
    },
    {
      "created_at": 1684223451,
      "level": "info",
      "message": "Completed epoch 4/4",
      "object": "fine-tune-event"
    },
    {
      "created_at": 1684223488,
      "level": "info",
      "message": "Uploaded model: davinci:ft-personal-2023-05-16-07-51-27",
      "object": "fine-tune-event"
    },
    {
      "created_at": 1684223489,
      "level": "info",
      "message": "Uploaded result file: file-ado9OSxuQnLmLisvKF7gRmhz",
      "object": "fine-tune-event"
    },
    {
      "created_at": 1684223489,
      "level": "info",
      "message": "Fine-tune succeeded",
      "object": "fine-tune-event"
    }
  ],
  "fine_tuned_model": "davinci:ft-personal-2023-05-16-07-51-27",
  "hyperparams": {
    "batch_size": 1,
    "learning_rate_multiplier": 0.1,
    "n_epochs": 4,
    "prompt_loss_weight": 0.01
  },
  "id": "ft-gm3VC33qcsOybFgV0Ra8ZXFu",
  "model": "davinci",
  "object": "fine-tune",
  "organization_id": "org-h30PMzIKL28VraqLFD4fMPMG",
  "result_files": [
    {
      "bytes": 23324,
      "created_at": 1684223489,
      "filename": "compiled_results.csv",
      "id": "file-ado9OSxuQnLmLisvKF7gRmhz",
      "object": "file",
      "purpose": "fine-tune-results",
      "status": "processed",
      "status_details": null
    }
  ],
  "status": "succeeded",   已经微调成功了
  "training_files": [
    {
      "bytes": 22019,
      "created_at": 1684222598,
      "filename": "file",
      "id": "file-ExYVR54RCcoLXeJ7TWYnNnCj",
      "object": "file",
      "purpose": "fine-tune",
      "status": "processed",
      "status_details": null
    }
  ],
  "updated_at": 1684223490,
  "validation_files": []
}
```

# 使用微调模型

登录官网进行使用微调的模型，[登录链接](https://platform.openai.com/playground) 

![1684224265838](如何微调ChatGPT.assets/1684224265838.png)

找一个训练数据里面的问题让它回答一下看看

![1684224706708](如何微调ChatGPT.assets/1684224706708.png)

OpenAI自带模型的回复

![1684224743030](如何微调ChatGPT.assets/1684224743030.png)

微调的模型回答的结果

![1684224913078](如何微调ChatGPT.assets/1684224913078.png)



扩展阅读：[官方关于微调的教程](https://platform.openai.com/docs/guides/fine-tuning) ，[微调的费用信息](https://openai.com/pricing)



**为什么要微调ChatGPT？**

当我们使用ChatGPT等语言模型进行文本生成或对话等任务时，我们可能会发现模型在某些方面表现得不够理想，例如在某些特定的领域或语境下，模型的输出可能会出现不准确或不连贯的情况。

这时候，我们可以使用微调来针对这些问题对模型进行进一步的训练，以提高其在特定任务或领域中的表现。例如，在进行客服对话时，我们可以使用微调来优化ChatGPT模型，使其能够更好地理解用户的问题，并给出更准确和有用的回答。

另一个例子是，在进行机器翻译时，微调可以帮助我们优化模型的翻译能力，使其在特定语言对之间的翻译表现更佳。例如，当我们想要将某个产品或服务推广到不同的国家或地区时，我们可以使用微调来优化ChatGPT模型，使其能够更好地理解目标语言的语法和语言习惯，从而提高翻译的准确性和流畅度。

总之，微调可以帮助我们针对具体的任务或场景，优化ChatGPT等语言模型的表现，从而使其更加适用于实际应用。

 



# Python学习

## 启动项目

python 文件名称.py

![1684203965007](如何微调ChatGPT.assets/1684203965007.png)

## 变量和数据类型 int，string,list

在Python中，变量用于存储数据，数据可以是数字、字符串、列表、元组、字典等类型。Python是一种动态类型语言，这意味着你可以在运行时动态地给变量赋予不同类型的数据。

下面是Python中常见的数据类型：

1. 数字（int、float、complex）

   - `int`：整数，例如 `42`

   - `float`：浮点数，例如 `3.14`
   - `complex`：复数，例如 `1+2j`

2. 字符串（str）

   - 一组字符，用单引号、双引号或三引号括起来，例如 `'hello'` 或 `"world"`

3. 列表（list）

   - 一组有序的元素，可以包含不同类型的数据，例如 `[1, 'hello', 3.14]`

4. 元组（tuple）

   - 一组有序的元素，可以包含不同类型的数据，但是元组一旦创建就不能修改，例如 `(1, 'hello', 3.14)`，这个就相当于C#中的Array集合

5. 字典（dict）

   - 一组由键值对组成的数据，用花括号括起来，例如 `{'name': 'Tom', 'age': 18}` ,根据键值，查询整个类型的值，就相当于C#中的*Dictionary*

```python

在Python中声明变量时，可以直接给变量赋值，Python会自动根据所赋的值推断变量的类型。例如：
# 声明一个整数变量
num = 42

# 声明一个字符串变量
name = 'Tom'

# 声明一个列表变量
lst = [1, 2, 3]

# 声明一个字典变量
dic = {'name': 'Tom', 'age': 18}
在Python中，可以使用type()函数来获取变量的类型。例如：
num = 42
print(type(num))  # <class 'int'>

name = 'Tom'
print(type(name))  # <class 'str'>

lst = [1, 2, 3]
print(type(lst))  # <class 'list'>

dic = {'name': 'Tom', 'age': 18}
print(type(dic))  # <class 'dict'>
以上就是Python中常见的数据类型和变量的声明和赋值方法。希望这能帮助你开始学习Python。
```

## 类型转换(int)

- int(): 将一个值转换为整数类型。
- float(): 将一个值转换为浮点数类型。
- str(): 将一个值转换为字符串类型。
- list(): 将一个值转换为列表类型。
- tuple(): 将一个值转换为元组类型。
- dict(): 将一个值转换为字典类型。

```python
# 将42转换为int，如果转换不了会有异常
str_value = "42"
int_value = int(str_value)

```



## 控制结构 if,else,for,while

```python
# if elif else
x = 5

if x > 0:
    print("x is positive")
elif x == 0:
    print("x is zero")
else:
    print("x is negative")
# for 
fruits = ['apple', 'banana', 'orange']
for fruit in fruits:
    print(fruit)
# while    
i = 1
while i < 5:
    print(i)
    i += 1

```

## 异常处理try,catch,finally

```python
# 捕获所有异常
try:
    result = 1 / 0
except:
    print("除数不能为0！")
finally:
    print("程序执行完毕。")
# 捕获指定异常
try:
    result = 1 / 0
except ZeroDivisionError:
    print("除数不能为0！")
finally:
    print("程序执行完毕。")
```

## 比较运算符==,!=,>,<

比较运算符用于比较两个值，返回布尔值 True 或 False。常见的比较运算符有：

- `==`：等于
- `!=`：不等于
- `>`：大于
- `<`：小于
- `>=`：大于等于
- `<=`：小于等于

例如：

```python
x = 10
y = 5

print(x == y)   # False
print(x != y)   # True
print(x > y)    # True
print(x < y)    # False
print(x >= y)   # True
print(x <= y)   # False
```

## 逻辑运算符&&，||，！

逻辑运算符用于组合比较条件，返回布尔值 True 或 False。常见的逻辑运算符有：

- `and`：与，两个条件都为 True 时返回 True，否则返回 False。
- `or`：或，两个条件中至少一个为 True 时返回 True，否则返回 False。
- `not`：非，反转条件的值。

例如：

```python
x = 10
y = 5

print(x > 5 and y < 10)  # True
print(x > 5 or y > 10)   # True
print(not x == y)        # True
```

## 定义函数 public void SayHi()

```python
# 有返回值,那么 return result ,没有返回值就不写return
def add_numbers(a, b):
    result = a + b
    return result

sum = add_numbers(2, 3)
print(sum)  # 输出 5

```

## 函数传递方式 name,params	

### 参数传递方式

在 Python 中，参数传递可以通过位置、关键字、默认值、不定长等方式进行。下面是几个示例：

```python
# 位置参数传递
def greet(name):
    print(f"Hello, {name}!")

greet("John")  # 输出 "Hello, John!"

# 关键字参数传递
def greet(name, message):
    print(f"{message}, {name}!")

greet(message="How are you", name="John")  # 输出 "How are you, John!"

# 带默认值的参数
def greet(name, message="Hello"):
    print(f"{message}, {name}!")

greet("John")  # 输出 "Hello, John!"
greet("John", "How are you")  # 输出 "How are you, John!"

# 不定长参数,*代表任何类型，任何长度的参数，我们可以传字符串，数组进来，并且数量也可以不固定
def sum(*numbers):
    result = 0
    for n in numbers:
        result += n
    return result

print(sum(1, 2, 3, 4))  # 输出 10
```

## 内置函数和标准库模块

当然，我可以为您提供一些这些常用内置函数和标准库模块的简单用法示例，供您参考：

- 内置函数 - print(): 打印输出信息到控制台

```python
print("Hello, world!")
```

- 内置函数 - len(): 返回一个序列的长度

```python
my_list = [1, 2, 3, 4, 5]
print(len(my_list))
```

- 内置函数 - range(): 生成一个数列，用于循环操作

```python
for i in range(1, 6):
    print(i)
```

- 内置函数 - type(): 返回对象的类型

```python
x = 42
print(type(x))
```

- 内置函数 - input(): 获取用户输入的值

```python
name = input("Please enter your name: ")
print("Hello, " + name + "!")
```

- 标准库模块 - os: 提供访问操作系统功能的接口

```python
import os

current_dir = os.getcwd()
print("Current directory:", current_dir)
```

- 标准库模块 - sys: 提供对Python解释器的访问和控制

```python
import sys

print("Python version:", sys.version)
```

- 标准库模块 - math: 提供各种数学运算函数

```python
import math

x = 2.5
print("Square root of", x, "is", math.sqrt(x))
```

- 标准库模块 - datetime: 提供日期和时间处理功能

```python
import datetime

today = datetime.date.today()
print("Today's date is:", today)
```

- 标准库模块 - random: 提供生成随机数的功能

```python
import random

x = random.randint(1, 10)
print("Random number between 1 and 10:", x)
```



## 文件操作

- 读取文件-根目录的文件

```python
with open('file.txt', 'r', encoding='utf-8') as f:
    data = f.read()
    print(data)

```

- 读取文件，有异常处理

```python
try:
    with open('file.txt', 'r', encoding='utf-8') as f:
        data = f.read()
        print(data)
except FileNotFoundError:
    print('File not found!')

```

- 写入文件

```python
with open('file.txt', 'w', encoding='utf-8') as f:
    f.write('Hello, world2!')
```

在Python中，可以使用内置函数open()来打开文件，然后可以读取或写入文件的内容。open()函数的语法如下：

```
pythonCopy codefile = open(file_path, mode)
```

其中，file_path是要打开的文件的路径，mode是打开文件的模式。mode参数有以下几种：

- 'r': 只读模式，打开文件后只能读取，不能写入，如果文件不存在，则抛出FileNotFoundError异常。
- 'w': 写入模式，打开文件后只能写入，不能读取，如果文件不存在，则创建文件，如果文件已存在，则会覆盖文件原有的内容。
- 'a': 追加模式，打开文件后只能写入，不能读取，如果文件不存在，则创建文件，如果文件已存在，则在文件末尾追加内容。
- 'x': 独占模式，打开文件后只能写入，不能读取，如果文件不存在，则创建文件，如果文件已存在，则抛出FileExistsError异常。
- 'b': 二进制模式，与其他模式搭配使用，表示以二进制格式打开文件。
- '+': 读写模式，与其他模式搭配使用，表示既可以读取文件，也可以写入文件。

## 定义类 封装，继承，多态

面向对象编程（Object-Oriented Programming，简称OOP）是一种编程思想，它将现实中的对象抽象成计算机程序中的类，类可以包含数据（属性）和行为（方法），并且可以被其他类继承和重用。

### 基本语法

在Python中，我们可以使用class关键字来定义一个类，类的基本语法如下：

```python
class ClassName:
    # 类的属性和方法
```

其中，ClassName是类的名称，可以根据需要进行命名。

### 封装

在类中，可以定义属性和方法。属性是类中的数据，方法是类中的函数。下面是一个简单的类的例子：

```python
class Person:
    name = ''
    age = 0

    def say_hello(self):
        print('Hello, my name is', self.name, ', I am', self.age, 'years old.')

```

在这个例子中，我们定义了一个名为Person的类，它有两个属性：name和age，以及一个方法：say_hello()。方法中的self表示对象本身，可以通过它来访问对象的属性和方法。

类定义好之后，我们可以实例化一个对象，并调用它的方法。下面是一个实例化对象并调用方法的例子：

```python
person = Person()
person.name = 'Alice'
person.age = 18
person.say_hello()
```

在这个例子中，我们实例化了一个名为person的对象，并设置它的属性name和age。然后，我们调用了person的say_hello()方法，它会打印出对象的名称。

除了属性和方法之外，类还可以有其他的特性，例如继承和多态。

### 继承（Person）

继承是一种机制，它允许我们创建一个新的类，它可以继承现有类的属性和方法，并可以添加自己的属性和方法。下面是一个继承的例子：

```python
class Student(Person):
    grade = 0

    def say_hello(self):
        print('Hello, my name is', self.name, ', and I am a student')
```

在这个例子中，我们定义了一个名为Student的类，它继承了Person类，并添加了一个属性grade和一个方法say_hello()。方法中的super()函数可以调用父类的方法。

### 多态

多态是一种机制，它允许不同的对象对同一方法做出不同的响应。下面是一个多态的例子：

```python
class Animal:
    def say_hello(self):
        pass #这个是个空的占位符，表示没有具体的实现，这个是给子类进行实现的

class Dog(Animal):
    def say_hello(self):
        print('Woof!')

class Cat(Animal):
    def say_hello(self):
        print('Meow!')

def greet(animal):
    animal.say_hello()

dog = Dog()
cat = Cat()

greet(dog)
greet(cat)
```

在这个例子中，我们定义了一个名为Animal的类和两个子类Dog和Cat，它们都重写了父类的say_hello()方法。然后，我们定义了一个名为greet()的函数，它接受一个Animal类型的参数，并调用它的say_hello()方法。最后，我们实例化了一个Dog对象和一个Cat

## 异步编程 async,await

```python
import asyncio


async def hello():
    print("Hello")
    await asyncio.sleep(1)  # 模拟耗时操作，1是等待1秒钟的意思
    print("World")


async def main():
    await asyncio.gather(hello(), hello(), hello())

asyncio.run(main())

#输出
Hello
Hello
Hello
World
World
World
```

## 连接数据库

- 构造数据库脚本

```mysql
CREATE TABLE users  (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(50),
  age INT,
  gender VARCHAR(10)
);

INSERT INTO users  (name, age, gender) VALUES 
('Alice', 18, 'Female'),
('Bob', 20, 'Male'),
('Charlie', 22, 'Male'),
('David', 19, 'Male'),
('Emma', 21, 'Female');

```

- 安装模块:pip install mysql-connector-python  ,vscode中运行，记得安装后关闭一下vscode，然后重新打开

![1684211506356](如何微调ChatGPT.assets/1684211506356.png)

- 链接数据库，`%s` 是占位符的意思，按照占位符的顺序进行赋值

```python
import mysql.connector

# 连接 MySQL 数据库
cnx = mysql.connector.connect(user='root', password='123',
                              host='localhost',
                              database='testdb')

# 插入数据
add_user = ("INSERT INTO users "
            "(name, age) "
            "VALUES (%s, %s)")
user_data = ("John", 25)
cursor = cnx.cursor()
cursor.execute(add_user, user_data)
cnx.commit()

# 查询数据
query = ("SELECT id, name, age FROM users "
         "WHERE age > %s")
age_threshold = 18
cursor.execute(query, (age_threshold,))
for (id, name, age) in cursor:
    print(f"User {id}: {name}, {age} years old")

# 更新数据
update_user = ("UPDATE users SET age = %s WHERE name = %s")
user_data = (26, "John")
cursor.execute(update_user, user_data)
cnx.commit()

# 删除数据
delete_user = ("DELETE FROM users WHERE name = %s")
user_data = ("John",)
cursor.execute(delete_user, user_data)
cnx.commit()

# 关闭数据库连接
cursor.close()
cnx.close()

```

