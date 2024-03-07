

# 让我进步的问题：

## PPS刊登系统的业务逻辑

Ebay管理：指的是某个电商平台的管理工作，Ebay管理指的就是

- Ebay刊登管理：刊登商品的管理工作，比如商品的图片，描述，关键字，价格，上架后的商品详细描述。如果没有的问题的，可以直接在这个模块提交到对应的平台，比如亚马逊，Ebay
- Ebay在售管理：管理在平台上面正在出售的商品信息，可以对正在出售的商品做价格的更改，信息的更改，下架的操作。 看商铺下拉的数量，公司在Ebay有几十家商铺，商铺可能专门出售母婴用品的商铺，专门卖服装的商铺
- Ebay库存明细：对商铺的库存进行管理，有些平台对发货的时间有要求，比如一日达，三日达，不同的平台的要求可能也不同，如果仓库的库存与线上的库存不一致，或者线上的库存快卖完，需要根据配置的进库逻辑，通知线下的仓库，比如义乌仓，佛山仓，美国仓进行进货，可以是按照上架数量进货100%，也可以是进货50件。保证线上有货卖，按时送到客户的手上。
- Ebay刊登模板：配置推送到Ebay平台商品展示信息的模板，比如图片怎么排版，文字是放在图片的左边还是图片的下面，售后服务，到货时间等，就是平常我们在淘宝购物的时候，最底部的那些信息。
- Ebay上架配置：看到配置有两个一个是Ebay，一个是亚马逊，忘记具体的功能了
- Ebay任务管理：配合Ebay在售管理进行使用，有时间批量更改上万个产品的信息，比如说价格，标题，如果在在售管理中一次去处理窗体卡顿，等待的时间会非常长，这个使用可以利用消息队列的形式，把更改的任务放到队列中，延后进行执行，通过菜单进行管理执行的结果，执行的商品，执行的时间，是否需要撤回等操作信息。

**PS:** 目前Ebay平台的模块建设的是比较完整的，Walmart亚马逊，*Fruugo*（英国电商平台）的模块建设稍微慢一些，这个从两个方便进行考量：第一从公司的业务发展，老板的战略推进进行建设相应的模块，而且现在开发这边人手不是很够，要把力气用在刀刃上面 第二 从平台开放的API来看，部分平台像Ebay，Walmart可能他们技术建设的比较完整，提供的功能很多，比如价格的更改，商品的撤回，具体要看平台开放了哪些API给我们，然后根据API进行分析系统的功能模块（一般是产品那边Eric进行分析整理后交给研发，也可以是我们进行分析整理，通过网站（这个是Ebay的开放API）：https://developer.ebay.com/docs#cat）

-   后面可能会开发商品的售卖监控，价格监控等模块，更好的服务于线上的电商售卖。

-   出售后的订单会通过易仓的API全部推送到易仓这个平台，对他们进行后续的订单管理操作

![1637207998692](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637207998692.png)

## Linq与Lambad表达式方法：

**ForEach：**循环

```
pageList.ForEach(p =>
{
    if (s.Id==p.Id)
    {
        p.Categories = string.Join(",", s.PathCategorie);
    }
});
```

**Where：** 筛选

```
Where(p => p.IsDeleted == false)
```

**OrderByDescending** ： 降序排序

```
OrderByDescending(p => p.Id)
```

**OrderBy ：**顺序排序

```
OrderBy(p => p.Id)
```

**Select ：**映射数据到新的对象,也可以映射到一个单独的属性 

```
 Select(p => new FyndiqListerDto
            {
                Id = p.Id,
                AccountId = p.AccountId
             }
        )
```



##  JSON对象,JSON数组,JSON字符串互转及取值

https://www.cnblogs.com/tangbang/p/tb.html

**前端处理：**



**后端处理：**

```c#
 //转换成Json数组(这样也可以)
 JArray jo1 = (JArray)JsonConvert.DeserializeObject(builderJson.ToString());
 //转换成JSON对象（保留后续使用）
 JObject jo1 = (JObject)JsonConvert.DeserializeObject(builderJson.ToString());
```



## 通过Mysql数据查询的值直接变成Json

```
SELECT CONCAT('[', GROUP_CONCAT(JSON_OBJECT( 
'id',id,'CategoryId',CategoryId,'CategoryName',CategoryName, 'ParentCategoryId', ParentCategoryId, 'IsLeaf', IsLeaf)), ']') as category FROM t_bi_fyndiq_category    where CategoryName='Entertainment'
```



## elementUI框架的 el-row el-col 与 el-table-column 用法区别

**Layout 布局**：el-row 一行24分，不论有多少el-col 格子，**:span**宽度加起来一般是24分（比较好控制，不是强制要求）。el-row通过**:gutter**设置el-col之间的间隔距离（间隔距离不影响:span加起来是24分的特点），el-col通过设置**:offset**进行分栏漂移（不需要设置:gutter间隔，我就可以偏离左边距多少多少位置）

**Table表格：** **el-table-column** 是Table表格的表头，比如姓名，年龄，性别

https://blog.csdn.net/acoolgiser/article/details/105891189  elementUI框架的 el-row el-col 与 el-table-column 用法区别

## vue：重置页面的数据

![image-20211201150522932](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20211201150522932.png)

Vue 重置页面data中数据：https://www.jianshu.com/p/f6c88d70a24b

## vue : 无法加载文件 D:\node-v14.12.0\vue.ps1，因为在此系统上禁止运行脚本

（1）以管理员身份运行VSCode

（2）执行命令：get-ExecutionPolicy（取得shell的当前执行策略，显示Restricted（表示状态是禁止的））

（3）执行命令：set-ExecutionPolicy RemoteSigned （开启远程下载权限）

（4）执行命令：get-ExecutionPolicy，显示RemoteSigned

版权声明：本文为CSDN博主「没有咸鱼的梦想」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/nineteenthdog/article/details/106415113

## 接口中引用的位置包括实现的实例：

如果有很多继承的话，滑动到最底下的那个就是实现

![1637496015723](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637496015723.png)

## Web First 代码生成工具的使用

**链接Mysql的字符串**：

```
server=gz-cdb-ezvn7csl.sql.tencentcdb.com;Port=58892;Database=bailun_bltpro;Uid=root;Pwd=shengye2021
```

[[MySQL\] - MySQL连接字符串总结](https://www.cnblogs.com/hcbin/archive/2010/05/04/1727071.html)



## js小技巧

js if判断多个条件_JS条件判断小技巧（一）https://blog.csdn.net/weixin_39999586/article/details/110745829?utm_term=js%E4%B8%AD%E4%B8%94%E6%9D%A1%E4%BB%B6&utm_medium=distribute.pc_aggpage_search_result.none-task-blog-2~all~sobaiduweb~default-0-110745829&spm=3001.4430

## Js数组转成字符串

```
var a, b;
a = new Array(0,1,2,3,4);
b = a.join("-");
```

[js数组与字符串的相互转换方法](https://www.cnblogs.com/asdyzh/p/9801784.html)

## MySql的使用

**查询建表时候的sql：**

```
show create table tableName
```

**生成数据库文档：**

[MySQL数据库生成数据库说明文档](https://www.cnblogs.com/hsybs/p/11196673.html)

**连接MySql:**

```
mysql -u root -p
```

**查询前10行数据**

```
select   * from t_bi_fyndiq_order limit 10
```

**根据数据库表直接生成实体类**：

```
select CONCAT('public ',CONCAT(CONCAT(' ',data_type),concat(CONCAT(' ',column_name),' {get;set;}'))) from information_schema.columns 
where table_name='t_bi_fruugo_order'  
```

**根据数据库表直接生成Dto实体类的返回内容：**

```
select  CONCAT(CONCAT(column_name,' = p.') ,CONCAT(column_name,',')) from information_schema.columns 
where table_name='t_bi_fruugo_order'  
```

**根据指定单词查询符合的属性：**

```
select column_name,data_type, column_comment  from information_schema.columns 
where table_name='t_bi_fruugo_order' and column_name like '%Price%'
```

**拼接值的语法**

```
select CONCAT(ShippingFirstName,ShippingLastName) as pjz from t_bi_fruugo_order 
```

**执行解释已选择的**：可以看到评估计划的结果

![1637222870887](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637222870887.png)

**查询创建工具：**可以查看所有的表字段

![1637222959918](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637222959918.png)

 **通过表名进行筛选：**

![1637223741329](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637223741329.png)

**拿到创建表的语句：**

![image-20211202190109345](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20211202190109345.png)

## 解决Mysql死锁

```
第一种：  
1.查询是否锁表  
show open tables where in_use > 0;  
2.查询进程（如果您有super权限，您可以看到所有线程。否则，您只能看到您自己的线程）  
        show processlist  
3.杀死进程id（就是上面命令的id列）  
kill id  
  
第二种：  
1.查看下在锁的事务   
select * from information_schema.innodb_trx;  
2.杀死进程id（就是上面命令的trx_mysql_thread_id列）  
kill 线程id  
例子：  
查出死锁进程：show processlist  
杀掉进程          kill 420821;  
其它关于查看死锁的命令：  
1：查看当前的事务  
select * from information_schema.innodb_trx;  
2：查看当前锁定的事务  
select * from information_schema.innodb_locks;  
3：查看当前等锁的事务  
select * from information_schema.innodb_lock_waits;  

```



## Git的使用

**下载代码**：git clone 地址

![1637210056074](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637210056074.png)

**拉取最新代码**：git pull

**提交代码**：git push

**resourcetree中的拉取和获取有什么区别** ：

**1、命令式解释：**
前者是 pull，后者是 fetch，pull 等于 fetch + merge。

**2、大白话解释：**
拉取：把你本地仓库没有 ，而远程仓库有的更新写入到你本地;

获取：用来查看对于你本地仓库的状态来说远程仓库是否有更新，仅此而已，并不会使你的本地仓库发生改变
————————————————
版权声明：本文为CSDN博主「zhongzunfa」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/zhongzunfa/article/details/80646346

## Navicat的使用

添加注释：**Ctrl+ /**

**激活地址**：https://www.cnblogs.com/wei9593/p/11907307.html  [Navicat Premium 12破解版激活（全新注册机）](https://www.cnblogs.com/wei9593/p/11907307.html)

## Vs Code的使用技巧

- **管理未完成的事件** ：TODO标签：下载插件：TODO Highlight，再下载：Todo Tree 进行管理，写大写的**TODO:**
- **碰到奇怪的错误：**ctrl+c 关闭一下vs code中的站点，新建一个终端: npm run dev
- **vue devtool插件** : 在谷歌浏览商城使用，Vue中AppMain/FyndiqListerAddEdit/Eldialog 展示当前窗体Open出来组件的使用的变量情况

界面的位置：

![image-20211126170844373](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20211126170844373.png)

根据变量名称过滤变量的值：

![image-20211126171335371](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20211126171335371.png)



- **空格警告的处理：**1.存在空格与制表符号，**解决**：到警告提示的行处，把前部的空格删除掉，然后再按ctrl+s 保存，系统自动会用没有问题的空格占满

```
 //504行：存在两个空格
 504:2  error  Mixed spaces and tabs  no-mixed-spaces-and-tabs

 505:2  error  Mixed spaces and tabs  no-mixed-spaces-and-tabs

 624:2  error  Mixed spaces and tabs  no-mixed-spaces-and-tabs

 639:2  error  Mixed spaces and tabs  no-mixed-spaces-and-tabs
```



- **点击直达教程：** 鼠标悬停到代码处，点击出现的v-if关键字可以直接到官网的教程。
- **绘制流程图**：插件商城搜索：**vscode-drawio**，左上角创建.drawio的文件，打开绘制流程图（打开后右上角地球可以调整语言）
- **返回上一个编辑的界面**：**alt+ 左右箭头**   ,文件→首选项→键盘快捷方式→搜索前进/后退
- **全局搜索** ：**ctrl+shift+f**


- **拖动窗体：**组件上下拆分的快捷方式，往右边拆分后，要怎么还原到拆分之前的样子：**直接拖动鼠标，放到那里不动的时候会出现阴影，出现阴影就说明可以了**

- **修改全局保存快捷键**：将 `saveAll` 绑定到 **Ctrl + S ＋Ctrl+S**上

- **快速执行html页面**: 先安装插件：**open in browser**，然后新建UI文件夹→新建index.html→打一个英文的**感叹号!**自动生成内容→输出快捷键：**Alt+B** 自动执行，如果需要引用**axios**进行后端接口Get请求，新建以下代码内容

```c#
<script>
axios.get("https://localhost:5001/Test/GetHello?userName=Ace&word=Hello")
.then(function(data){
           console.log(data);
        })
    </script>
  
```



## ABP框架的

![img](https://upload-images.jianshu.io/upload_images/2799767-f113cae025f4e844.png?imageMogr2/auto-orient/strip|imageView2/2/format/webp)

https://www.jb51.net/article/86688.htm --ABP框架的体系结构及模块系统讲解

https://www.jianshu.com/p/a6e9ace79345  --ABP入门系列目录——学习Abp框架之实操演练

**ABP—对象之间的双向映射** :

```c#
 //两个类对象之间的双向映射
 [AutoMap(typeof(Person))]
```

https://blog.csdn.net/wangwengrui40/article/details/86609037  --ABP—对象之间的映射

## Consul的执行

Consul的执行

```
consul agent -dev
```

执行Net

```
dotnet run
```



## Asp Net Core源码

https://github.com/dotnet/corefx 这个是.net core的 开源项目地址

https://github.com/aspnet 这个下面是asp.net core 框架的地址，里面有很多仓库。

https://github.com/aspnet/EntityFrameworkCore  EF Core源码

https://github.com/aspnet/Configuration 配置模块源码

https://github.com/aspnet/Routing 路由模块

https://github.com/aspnet/Security 认证及授权 

https://github.com/aspnet/DependencyInjection 依赖注入

https://github.com/aspnet/HttpAbstractions 有很多的一些HTTP管道的抽象接口都定义在这里

https://github.com/aspnet/Options   看名字

https://github.com/aspnet/Mvc 

https://github.com/aspnet/Hosting

# DDD领域驱动设计

DDD是一种程序分析设计方法，不关乎具体技术，具体代码实现依旧是OOP和AOP

**应用场景：**应对复杂软件开发，不熟悉，不明白的系统。

解决沟通问题，降低沟通成本。比如建立一个小区很简单，如果是建立一个区，一个城市呢，会很复杂

技术人员与需求人员由于各自缺乏彼此的专业知识，比较难站在对方的角度进行思考，现实中难以沟通。

双方的统一语言在哪里？

- **Domain（领域/业务）：**领域可以是一个单独的服务，也可以是一个项目，可以是一个模块，可以是一个BLL，甚至是一个问题（需求），领域就是用来分析需求的（需求就是业务），把一个大盒子，不断的划分成小盒子，小盒子里面不断的划分

- **Driven（驱动/实现）**：对外提供哪些功能，哪些人可以获取数据，权限验证，代码实现

- **Design（设计）**：把按照领域实现的模块在这里进行组装，组合成最终的产品

领域驱动设计以需求为主导，确定下来业务需求后再进行代码开发，避免开发过后的改动。可以**分为三步走**：第一先理解领域，再拆分领域，最后再细化领域，领域与领域之间的拆分是最复杂的，因为有些需求之间存在交集，怎么划分到正确的领域？有些模块是单独划分成一个领域还是合并到当前的一个领域当中（就是这个模块业务权重是怎么样的，比如秒杀模块），这个需要对业务有足够的理解。

- **理解领域：**

大的盒子长什么样子

- **拆分领域：**

把大盒子拆分成小盒子，小的模块

- **细化领域：** 

1. 梳理**领域概念**：领域中有哪些业务，业务的重点在哪里，有哪些交流词汇。比如SUK，易仓。这一步是理解宏观业务的环节，最快的理解方法就是和领域专家进行交流。
2. 梳理**业务规则**：领域中我们关于的业务规则，比如唯一性原则，余额不能小于零，国内物流不能超过5次。
3. 梳理**业务场景**：领域中核心业务场景，比如电商平台中加入购物车，提交订单，发起付款等核心业务场景
4. 梳理**业务流程**：领域中关键业务流程，比如订单处理流程，退款流程，这个阶段会画业务流程图。





《巨蟒与圣杯》的故事，袖子阴影是完美，喜剧效果是内核。

 

 面向对象的分析方法太过于细致，如果是小的项目对象不多的话，比如几十个对象，那么问题不大，如果是几百个，上千个表，依靠面向对象就不好进行控制。DDD是面向对象分析设计的延伸与扩展，教我们怎么去分析需求。



**边界：**分而治之

领域并不仅仅是模块



#### DDD领域驱动设计：大盒子里面套小盒子

DDD领域驱动设计是一种业务分析设计方法 ，把整个系统从业务层面进行从大到小的划分模块，由模块再划分成表，由表再清晰表里面具体的属性，这个时候考虑的不是一个类，考虑的是一个一个小盒子，比如电商平台：

**第一步：** 划分为仓储 物流 订单 交易 用户这些模块

**第二步：** 划分仓储里面有不同地点的仓库，仓库管理人员，仓库保存的货物等信息

**第三步：**进一步划分表里面的内容，或者单个模块技术业务上面的实现。



微服务架构落地的第一件事就是DDD，微服务落地时需要将业务进行分拆，把项目分解成多个低耦合的模块，而DDD刚好就可以进行拆分。



**领域：**大盒子，聚合根（可以包含多个对象）

## 组件的命名驼峰命名可以直接转中横线

下面声明：**ArticleDetail**  引用直接 **article-detail** 也可以

![1637199983719](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637199983719.png)

## Vue中Main执行的生命周期

**src/main.js**文件 入口文件，初始化vue实例并使用需要的插件（引入公共element组件等、路由拦截、start登录、鉴权可以在这里统一处理）

**router/index.js：**执行后应该是直接到路由这边跳转到登录界面 **答：老师说是通过webpack进行配置的**

## 创建Vue对象的生命周期详解

声明周期：简单就是打开页面加载的内容，渲染完数据的内容，加载完数据，窗体关闭的时间

https://www.jianshu.com/p/672e967e201c 

![11370083-f279314aef6741db](C:\Users\12987\Desktop\11370083-f279314aef6741db.jpg)



## 跨域问题如何解决？

在Asp.Net Core中可以进行跨域配置（使用**3、设置特定来源可以跨域** 进行解决） https://www.cnblogs.com/dotnet261010/p/10177166.html，日常的前后端分离开发中我们可以通过，跨域的配置进行解决跨域的问题，线上的环境可以通过Nginx部署方向代理进行解决跨域的问题

**ACE老师跨域问题解决方案：**

```c#
//startup.ConfigureServices中写在第一行
 services.AddCors(c => c.AddPolicy("any", p => p.AllowAnyMethod().AllowAnyOrigin().AllowAnyHeader()));

//startup.Configure中写在 app.UseRouting();下面
app.UseCors("any");

```

## VUE生成的Dist文件要怎么部署运行？

webpack打包vue项目之后生成的dist文件该怎么启动运行

https://www.cnblogs.com/zhujiabin/p/10557982.html

## Vs Code中快速创建Html代码

1. 选中文本为html，先输入一个 !
2. 点击 tab 键
3. 自动生成标准的html代码

 

##  PPS刊登系统代码中代码的结构

### Vue前端页面逻辑



- src --业务文件

 api --都是发起业务的请求，通过export 声明方法，传入请求的URL，请求的参数，请求的类型（get/post）进行请求数据

```
--API顶部引用通用的封装方法，比如这个是请求API的方法
 import request from '@/utils/request'  
 
```

  views  --展示页面

```
  --Es6语法，在View页面中引用API方法,导入括号{ }里面的方法权限到本页面进行使用
 import { login, getInfo, getUserInfo } from '@/api/user' 
```

assets  --静态文件的保存，比如css,js,images

 compoents--封装的组件，

```
--在View页面用引用组件的方法
import SimpleTable from '@/components/SimpleTable'
```

directive --管理类型的组件，在Fyndiq，Fruugo管理中有使用

filters --过滤的方法，通用的方法

 icons  --图标

 styles --样式

 utils --通用的方法

 vendor --导出excel文件操作类



 store --  管理组件（vue-x里面的东西）

router --路由  ，左边菜单栏的映射关系都是放在这里的，router/modules/

```
新品开发管理 --newproduct.js
Ebay管理   --ebaymanage.js
Walmart管理  --walmartmanage.js
刊登运营管理 --publishoperatemanage.js
Real管理   --realmanage.js
Fyndiq管理 --fyndiqmanage.js
```

layout--布局

- public  --不需要编译的静态文件
- package.json  --系统依赖的第三方库名单
- package-lock.json  --系统依赖第三方库的详细依赖关闭
- node_modules --第三方依赖的名单

- .env.development  --【开发环境】webapi请求的路径配置
- .env.production  --【生产环境】webapi请求的路径配置

点击左边菜单栏：顶部的名称就是在Vue页面上的名称，可以在文件夹中进行搜索确认页面的位置，也可以ctrl +shift +F 在解决方案中进行搜索

**解决跨域问题：** WalmartApi 在Vue前端页面如果请求的后端没有做跨域的处理，可以加上这个URL的前缀，比如

http://localhost:3000/WalmartApi/FyndiqOrderManage/GetFyndiqOrderList ，



可以在路由中进行查询：src\router\modules\newproduct.js

新增页面：先在router\modules 中新建Js文件，再把新建的文件放到index.js进行配置

![1637143162071](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637143162071.png)

### Asp Net MVC后端：

Ebay管理 →Ebay刊登管理

http://localhost:6031/Api   --这个是PPS刊登系统的接口，代码路径： web/Bailun.ProMgr.WebApi

**下面的这三个接口异常，无法启动：**

http://localhost:5002/platformfeeconfig/GetPlatforms  --VUE_APP_PROFIT

http://localhost:10120/SupplierManage/Supplier/GetSupplierList  --VUE_APP_PDS

http://localhost:3000/SupplierManage/Supplier/GetEnumList  --未知

Ebay管理→在售管理 ：**正常打开**

**依赖注入框架**:使用ABP框架中Castle.Windsor容器进行注入，代码位置如下：

Global.asax→Application_Start→IocManager

**注入类（应该是这里）**：

\bltpro\Bailun.ProMgr.Application\BltProDbApplicationModule.cs  

全局搜索没有找到容器对象：**WindsorContainer，IWindsorContainer **

![1637145523256](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637145523256.png)

通过什么注册的呢？ 配置文件注册还是反射？

下面是后端代码的结构：

publishApi：发布系统的API

infrastructure ：基础类库

整体感觉代码比较乱，各个依赖关系层次不强，里面的子系统，子API很多，而且好像还有定时的任务集成到这里面，实体类也不需要用t_开头，也没有放到统一的Model，Entity，Dto等文件夹中，散落在程序中（比如Dev，Account，SKU，Ebay前缀开头的可能是实体类，这些是根据业务进行命名的）,Service里面可能不是业务逻辑，包裹了接口信息（还好接口是通过I前缀声明的）。



**Bailun.ProMgr.WebApi** ：这个 PPS刊登系统的代码，还算比较清楚的，主要核心逻辑在Controllers这个文件夹，其他的文件夹大部分是空的，还有实体类与路由的实体类

![1637138440172](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637138440172.png)





## PPS刊登系统，新加一个功能

Fruugo在售管理  fruugo-listing/index
**Fruugo订单管理  fruugo-order/index**

--在售管理
http://localhost:3000/WalmartApi/FyndiqListingManage/GetFyndiqListingList

### **后端项目结构**

```c#
//访问数据库/数据库实体类/实体类访问层接口
Bailun.Walmart.Publish.Core/Domain/FyndiqListing
    
//实现实体的数据库接口的访问
Bailun.Walmart.Publish.EntityFrameworkCore\Fruugo
//映射EF的的关系
Bailun.Walmart.Publish.EntityFrameworkCore/BaseDbContext
   
//这里面写Dto实体类/接口/Service
Bailun.Walmart.Publish.Application 

//这里写控制器，WebApi的启动入口
service/Bailun.Walmart.Publish.Service.Publish   

//数据库的访问实现
Bailun.Walmart.Publish.Extension/Domain/RepositoryExtensions



```

**总结：**

1.先建立实体类，根据数据库表的结构,映射EF表关系

2.建立实体的数据库访问接口，并且在Bailun.Walmart.Publish.EntityFrameworkCore\Fruugo中新建实体数据接口的实现类

3.建立Dto类根据界面展示的属性
4.建立DtoFilter类，根据界面筛选的条件
5.建立Service的抽象接口，编写抽象的查询访问
6.编写Service的实现，继承抽象接口并且，调用声明实体的数据库访问接口进行实现。
7.声明控制器类，新建查询访问， 建立Service的抽象接口，通过Service实现的抽象接口进行
  访问。

**收获：**

**不能创建组件的接口的依赖：**Service接口的命名与实体的命名不一致，ABP中service的注入默认是按照接口对应的名字创建实例，而且可以通过

原文地址：https://www.cnblogs.com/sugarwxx/p/13650214.html

**通过Consul的请求也可以调试：**打开Consul的时候，通过前端请求接口也可以直接调试

vue通过变量绑定值**首字母需要小写**： scope.row.**ArticleTitle** 错误，需要scope.row.**articleTitle**



**1.在售管理新建控制器类**
FruugoListingManageController控制器的顶部增加访问特性：

```
    [Route("[controller]")]
    [Produces("application/json")]
```

继承父类：

```
: AbpController
```

**2.在售管理新建接口  IFruugoListingService** ，接口实现一个查询方法，并且继承

```
：IApplicationService
```

3.控制器新增查询方法：

```c#
//1.新增方法
public VueResult GetFruugoListingList：

//2.新增访问特性
[BailunAuthentication(LoginMode.Ignore)]
[HttpPost("GetFyndiqOrderList")]  //方法名称需要更改
//3.新增返回值
return new VueResult() { StatusCode = 200, Result = result };
//4.方法入参
[FromBody]FyndiqListingFilterDto input

```

 

4.新建Dto表实体类：一个是界面上展示的内容，一个是界面上查询的条件

**原表位置：**

```
Bailun.Walmart.Publish.Application/Fyndiq/Dto/FyndiqListingFilterDto.cs
```



## PPS刊登系统，开发商品刊登功能

在沃尔玛系统中的Fyndiq管理下新增：Fyndiq刊登管理功能

**业务：**

从商品库中通过审批的信息中，提取商品的信息填充到刊登模块，比如商品归属的店铺，商品的类别，尺寸，颜色，运输政策，退货政策，确认无误进行保存后，在窗体提供刊登信息查询功能，可以对待需要刊登的商品进行编辑，删除

**步骤：**

1. **设计实体表：**通过Fyndiq平台的API里面返回的信息类别设计表（注意对产品描述的解析是否支持富文本编辑器），根据每个信息类别的内容设计表的属性（每个类别的请求官方都提供的假数据Json，可以利用起来，官方文档中的属性还会有字典值属性，要做成下拉框，注意实体类中需要添加公司个性化的属性，比如：是否定时刊登，定时刊登时间，刊登状态），根据分析的信息类别与属性设计实体表

1. **设计界面：**根据实体表的类型与属性内容，进行设计窗体界面（目前先使用Walmart刊登范本模块进行Copy一份，然后逐步更改，然后重新设计左上角：添加的逻辑），界面需要有哪些模块（商品信息、分类属性、销售信息、退货政策、运输政策、物品所在地、产品描述、多属性信息），注意界面上的非空限制策略（根据平台API的要求进行设计），个性化业务策略（余额不能小于零，国内物流地址不能超过五行）。
2. **请求接口：**设计好界面与实体表信息后，可以通过PostMan请求Fyndiq平台API（有测试环境接口），进行返回（成功了就是成功了（只会成功一次），失败了我们可以改动），PostMan的右上角是有功能自动生成：C#-RestSharp代码的内容，可以参考
3. **创建定时任务：** 根据Fyndiq平台的选择API顶部的对接语言进行自动生成对接的代码，推荐使用平台的API语言（参考PostMan与Fyndiq平台生成的代码与目前生成的代码进行比对，看目前的生产环境的语法，内容是与哪个符合的）



**看API接口文档**

1. Rest风格的接口，通一个接口地址使用Post,Get,Delete,Put区分增加，获取，删除，更改订单信
2. Api接口里面第一个图片是页面所有的属性，下面的内容是对Json内部属性的抽象，Json属性的抽象可以看这个表格的内容![image-20211123154151287](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20211123154151287.png)

**碰到的问题：**

1. 界面一片空白问题的处理，注释了两个异步组件，新增了生命周期的日志打印，检查时发现浏览器的**控制台抛出了很多异常！**
2. 发现所有的界面点击都是一片空白，点击也没有触发输出日志，然后**重新启动**了站点，重启后点击模块列表正常展示数据（好像界面长时间不操作会出现假死，点击失效的情况）。
3. 直接拷贝所有的页面进行更改，对页面的组成不熟悉，比如Vue对象的属性，创建方法的作用，引用到的组件这几点都不了解，不利于学习。建议先创建一个Vue的基架，然后把对应的页面内容，方法，组件慢慢的拷贝到页面当中，这个过程是一个学习的过程
4. 开发过程中时刻看F12控制台的内容，配合**Vue devtools** 扩展程序，进行代码的调试。

**其他**

![image-20211122182000795](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20211122182000795.png)



###  定时任务

Net Core 3.1开发

![image-20211203094850600](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20211203094850600.png)

![image-20211203094424187](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20211203094424187.png)

![image-20211203094357807](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20211203094357807.png)

#### 项目结构

- Sunway.Task.Common  

1. **发起平台请求：**比如获取官方商品列表，获取类目列表（就是那个底部8700多个选项的东西），修改价格，修改库存，里面有Ebay，Fruugo，Fyndiq，Walmart等平台的请求接口
2. **通用类：**操作图片，操作JsonHelper,加解密方法，生成随机字符，Token操作，时间戳与日期的转换

- Sunway.Task.Data：

1. **Dto：** 特性类，用文件夹名称区分数据库名称
2. **Entity：**实体类 ，用文件夹名称区分数据库名称
3. **Enum：**通用的枚举，是否已经上传，是否上传成功

-   Sunway.Task.Job：

1. **ECCang：**同步易苍数据的同步，比如产品的数量，产品颜色数据，产品吃嘛数据，下单用户等请求信息
2. **PPS：**订单下载，范本刊登，更新刊登信息方法，分为订单，在售，报表信息类（Ebay,Fruugo,Fyndiq,Walmart四个平台）

- Sunway.Task.Service：业务逻辑层，用文件夹名称区分数据库名称

1. **Impl**：编写Service业务逻辑代码
2. **接口**：编写业务逻辑所需要实现的抽象方法

- Sunway.Task.Web ：控制器

​	编写通过Swagger请求的WebApi方法（实际不使用，实际的使用在Job里面直接调用Service层的方法），不过也是在Web层中的Startup中的服务器注册UseHangfireJob

做需求的路径：

1. 打开 Sunway.Task.Data 文件夹，在Entity文件夹下面找到对应的数据库，在库下面找到对应的平台， 然后建立实体类，我看里面的实体类是直接用WebFirst生成的，不用想EFCore在DbContext里面配置表的名字，Sugar是通过 [SugarTable("t_bi_fyndiq_lister")]确认表名，然后配置主键特性：      

   ```
   [SugarColumn(IsPrimaryKey = true, IsIdentity = true)]
   ```

   

2. 打开 Sunway.Task.Service 文件夹 ，打开对应的数据库文件夹，新建Service的抽象接口，然后打开Impl文件夹创建Service类，集成接口进行实现，接口与服务的注入是通过遍历程序集进行注入的（见Framework/Sunway.Framework/IocRegister.js RegisterAssembly方法）

   

```
1.先查询店铺的账号
2.根据请求的Json结构，声明对应的实体类接收后台的数据值，这个实体类就像界面保存数据的DTO，要保存全部的数据
3.四个表的数据赋值都要有不同的方法进行赋值，声明公共变量的Dto，每个方法的赋值都是引用这个公共的Dto进行赋值，而且方法要返回这个Dto
4.然后把对象转换成Json进行Post请求（参考钉钉的沃尔玛的请求进行处理）
```



1. 打开 Sunway.Task.Web 文件夹，找到或新建对应类的控制器，通过Swagger调用新增的方法进行测试，测试通过后进入4
2. 打开 Sunway.Task.Job 文件夹，打开PPS，找到或新建对应平台所属业务的Job类，调用新增的方法（这个是通过Hangfire进行调用的，只有正式环境才能调用）



![image-20211203094222993](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20211203094222993.png)

![image-20211203094006097](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20211203094006097.png)

# Vue高级班

Vue的作用 

- 响应式刷新界面的变量：比如界面绑定到Vue的变量，如果变量的值发生了改变，那么对应Html也会进行刷新改变。
- 可以模块化页面，
- 组件化页面

可以通过Vue把后端传递到界面的值（**可以是单个值，也可以是实体对象**）封装到.js文件中进行保存，程序界面通过script标签进行引用就可以

**以下内容见菜鸟教程** https://www.runoob.com/vue2/vue-template-syntax.html

## Vue基础结构

1.页面通过绑定的对象进行赋值，解构：html，引用script在线vue，<scrpit>new Vue对象（Vue对象里面有Vue的名字，Vue绑定的属性，Vue绑定的方法，而且我发现方法里面可以直接引用Vue绑定的属性）



```c#
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>Vue 测试实例 - 菜鸟教程(runoob.com)</title>
	//1.通过在线的方式引用Vue的环境
	<script src="https://cdn.staticfile.org/vue/2.4.2/vue.min.js"></script>
</head>
<body>
	<div id="vue_det">
		<h1>site : {{site}}</h1>
		<h1>url : {{url}}</h1>
		<h1>{{show()}}</h1>
	</div>
	//2.用script包裹住代码进行声明Vue的对象
	<script type="text/javascript">
	 
		var vm = new Vue({
		   //vue的名称
			el: '#vue_det',
            //界面渲染完后加载
            mounted()
            {
                
            },
		  //实体类下面绑定了哪些属性
			data: {
                a:1,
                b:2
			},
            //计算属性
            computed:
            {
                Add:function()
                {
                    reutrn a+b;
                }
            },
            //过滤器,比如把金额保留两位小数
          filters:{
	        //过滤器1：（使用时{{msg|filterA}}）
	      filterA(value){
	        return “￥”+value
	      },
         //过滤器2：
	      filterA(value){
	        return "$"+value
	      }
	    }
    })
			//方法列表
			methods: {
				show: function() {
					return  this.site + " - 学的不仅是技术，更是梦想！";
				}
			},
          //下面为创建Vue的生命周期函数
        beforeCreate: function () {
            console.log('beforeCreate 创建前状态===============》');
        },
        created: function () {
            console.log('created 创建完毕状态===============》');
        },
        beforeMount: function () {
            console.log('beforeMount 挂载前状态===============》');
        },
        mounted: function () {
            console.log('mounted 挂载结束状态===============》');
        },
        beforeUpdate: function () {
            console.log('beforeUpdate 更新前状态===============》');
        },
        updated: function () {
            console.log('updated 更新完成状态===============》');
        },
        beforeDestroy: function () {
            console.log('beforeDestroy 销毁前状态===============》');

        },
        destroyed: function () {
            console.log('destroyed 销毁完成状态===============》');
        },
		})
	</script>
</body>
</html>
```



## VsCode搭建Vue项目

如果执行项目出现：vue-cli-service 不是内部或外部命名。。。

执行:**npm install**

VsCode中新建终端，执行以下命令

- 安装vue-cli：

```
npm install -g vue-cli
```

- 安装webpack

```
npm install -g webpack 
```

- 创建保存项目的文件夹，cd到文件夹：

```
cd D:\VUE\01vue_vscode
```

- 创建项目

```
vue init webpack myvue
```

全部**按回车**： 创建时出现一大堆的配置：路由，作者，权限认证，npm依赖安装的配置 。

- 执行项目

```c#
//先cd 到项目目录
cd D:\VUE\01vue_vscode
//再执行运行命令
npm run dev
```

1. 通过http://localhost:8080/#/，可以访问生成后的地址，

2. 初始启动的页面是：**src/components/HelloWorld.vue** ，

3. 为了方便做逻辑，可以把HelloWorld.vue中的ul列表剔除（**注意尾部要保留一行空白，这是Vue的规范**）

4. 执行的时候如果启动的时候8080端口被占用了，比如开了多个vue启动项目，vscode会用空闲的端口去代替，

5. 可以用cls清空终端的命令（但是出现错误后就动不了了，出现错误的时候，命令行无法输出，那么使用Ctrl+C 快捷键退出当前会话，就可以重新输入变量值）

6. 执行的时候终端出现代码错误，看到提示的错误行，可以使用**ctrl+鼠标左键** 进入详细的错误

7. 使用自定义组件的时候， components:{  mycom  --左边需要空出四个空格，两个空格会抛异常} ，使用右键格式化文档，也可以使用快捷键：**Shift+Alt ＋F** （先按照文章进行配置（先配置好第一步的settings.json，然后再安装插件）：https://blog.csdn.net/zm_miner/article/details/94416776）

9. 使用有道截图翻译Vs Code的提示英文的时候有时候可以正常翻译（错误提示可以翻译，下载npm的时候也可以），有时候又不能，可以看一看

10. **字符串为单引号：**export default {name: **'**TestCom**'** }    

11. **设置快键键清空终端的日志：**先打开设置 - > keyboard Shortcuts -> 搜索 "workbench.action.terminal.clear" -> 目前设置：**ctrl+L ctrl＋k**

12. Class的命名用横杠：

    ```
    .sp-page-table .el-pagination {
      margin-top: 3px;
    }
    ```

    

13. **方法的结束不需要分号；  **： 

    ```
     data(){ 
     return {
        name: 'tina'
            }    这里注意是不需要这个分号的; 
    ```

    

- 发布项目

```
npm run build
```

项目根目录出现的dist文件夹就是发布部署的项目



**原文：**VSCode搭建Vue项目 https://www.cnblogs.com/zyskr/p/10609288.html



## Vue执行发布后的文件：

直接执行我们会发现页面是一片空白的，需要将.html文件中的引用由绝对路径更改为相对路径：

![1637024057513](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637024057513.png)

**修改前： ** 

> ```
> <script type=text/javascript src=/static/js/manifest.2ae2e69a05c33dfc65f8.js>
> ```

**修改后（去除/static，前面的左斜杠/）：**

> ```
> <script type=text/javascript src=/static/js/manifest.2ae2e69a05c33dfc65f8.js>
> ```



## Vue.js起步

- 绑定到Vue里面的变量是可以是外部的变量进行赋值data ：wbdata，这种赋值的方式是引用的赋值，如果两方任何一方对值进行改变，都会作用到对方的变量上。
- 直接通过VueName.Age可以引用data中的变量值，如果要引用VueName.data，可以使用VueName.$data，如果要引用html的div，可以使用VueName.$vue_student
- Html中引用Vue的变量可以先声明一个作用域,比如<div></div>, 然后通过ID进行绑定DIV：<div id="vue1"></div>，然后直接可以再div通过双倍的花括号{{}}进行引用vue中的变量：<div id="vue1">{{message}}</div>

**一个小的例子：**

```
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>Vue 测试实例 - 菜鸟教程(runoob.com)</title>
<script src="https://unpkg.com/vue/dist/vue.js"></script>
</head>
<body>
<div id="app">
  <p>{{ message }}</p>
</div>

<script>
new Vue({
  el: '#app',
  data: {
    message: 'Hello Vue.js!'
  }
})
</script>
</body>
</html>
```

## Vue模板语法

### v-html渲染html变量

vue中的变量是data_message:"<h1>渲染html变量</h1>"，div中可以调用<div v-html="data_message"></div> 进行渲染html变量的值

```
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>Vue 测试实例 - 菜鸟教程(runoob.com)</title>
<script src="https://cdn.staticfile.org/vue/2.2.2/vue.min.js"></script>
</head>
<body>
<div id="app">
    <div v-html="message"></div>
</div>
	
<script>
new Vue({
  el: '#app',
  data: {
    message: '<h1>菜鸟教程</h1>'
  }
})
</script>
</body>
</html>
```

### v-bind:class指令动态渲染界面元素

在vue中的变量，定义data_flag:false,在<div></div>中进行引用，可以是<div v-bind:class="{class_red:data_flag}"></div> （需要在最外层包一层<div id="vuename"></div>，绑定的id名称不需要使用#号）

```

```



### 使用v-model动态改变Vue中变量的值

声明Vue的变量，定义data_flag:false,可以通过v-model绑定变量到<input type="checkbox" v-model="data_flag" id="r1">中，通过点击checkbox可以很方便的改变data_flag变量的值，达到动态渲染界面的目的

### 使用Vue三元表达式

```c#
<div id="app">
    //直接可以计算5+5
    {{5+5}}<br>
    //可以使用三元表达式
    {{ ok ? 'YES' : 'NO' }}<br>
</div>
    
<script>
new Vue({
  el: '#app',
  data: {
    ok: true
  }
})
</script>
```



### 使用Vue使用javascript方法

```
<div id="app">
    {{ message.split('').reverse().join('') }}
</div>
    
<script>
new Vue({
  el: '#app',
  data: {
    message: 'RUNOOB'
  }
})
</script>
```



### 使用v-if进行判断

```c#
  //引用v-if="seen"
  <p v-if="seen">现在你看到我了</p>
  //声明
  new Vue(
  {
    //注意声明的时候是用井符号# 加上名字
   el:"#VueName",  
   data:
   {
      seen:true
   }
  })
```

### 使用v-bind绑定html元素的属性

```c#
  //使用v-bind:href绑定a标签的属性
  <a v-bind:href="url">菜鸟教程</a>
  //Vue属性赋值
  data: {
    url: 'http://www.runoob.com'
  }
```

v-bind代表我要操作页面元素的属性，可以是绑定href属性，可以是style属性，也可以是id属性，如果是绑定页面属性的事件比如click属性可以使用v-on。

**v-bind缩写**

```
<!-- 完整语法 -->
<a v-bind:href="url"></a>
<!-- 缩写 -->
<a :href="url"></a>
```

### 使用v-on绑定html元素的事件

```
  //引用Vue绑定的事件
   <input type="button" v-on:click="show"/>
  //在vue变量中声明事件
  methods:
  {
    show:function
    {
      return 'Method';
    }
  }
```

使用v-on负责监听dom中元素的事件

**v-on缩写**

```
<!-- 完整语法 -->
<a v-on:click="doSomething"></a>
<!-- 缩写 -->
<a @click="doSomething"></a>
```

### 使用v-model数据双向绑定

使用v-model绑定的数据，会随着input输入框的值的改变（一般用在值可以变化的变量上面），而改变引用到它值位置的值。

```
  //引用的值的位置
  <p>{{message}}</p>
  <input type="input" v-model="message"/>
  //赋值的位置
  var a=new Vue(
  {
    el:"#myvue",
    data:
    {
       message:"我是一个值"
    },
    methods:
    {
      show:function()
      {
        return '这是一个方法';
      }
    }
  
  });
  
```

**v-model** 指令用来在 input、select、textarea、checkbox、radio 等表单控件元素上创建双向数据绑定，根据表单上的值，自动更新绑定的元素的值。

##### **v-model修饰符**

  **.lazy**

在默认情况下，`v-model` 在每次 `input` 事件触发后将输入框的值与数据进行同步。可以添加 `lazy` 修饰符转为在 `change` 事件之后进行同步

```html
<!-- 在“change”时而非“input”时更新 -->
<input v-model.lazy="msg">
```

**.number**

如果想自动将用户的输入值转为数值类型，可以给 `v-model` 添加 `number` 修饰符：

```html
<input v-model.number="age" type="number">
```

**.trim**

自动过滤用户输入的首尾空白字符

```html
<input v-model.trim="msg">
```



### 使用v-show根据条件进行展示元素

感觉这个v-show与v-if的使用有点像，可能它使用起来更加规范

```
  //使用变量,放到v-show,if中的时候，不需要使用双井号进行引用{{}}
  <p v-show="flag">你看不见我<p>
  //声明变量
  data:{
    flag:false
  }
```



### 使用Vue的过滤器过滤Vue中的变量值

```c#
<!-- 在两个大括号中，多个过滤器之间用竖杠隔开 -->
{{ message | filterA | filterB }}
//过滤器可以接收参数，这里message是filetrC的第一个参数,bb,cc分别是第二与第三个参数
{{ message | filterC('bb', 'cc') }}

<!-- 在 v-bind 指令中 -->
<div v-bind:id="message | filterA | filterB "></div>

<!--声明过滤器filters，多个filter用逗号隔开-->
new Vue({
  el: '#app',
  data: {
    message: 'runoob'
  },
  filters: {
    filterA: function (value) {
      if (!value) return ''
      value = value.toString()
      return value.charAt(0).toUpperCase() + value.slice(1)
    },
    filterB: function (value) {
      if (!value) return ''
      value = value.toString()
      return value.charAt(0).toUpperCase() + value.slice(1)
    },
    filterC: function (value,b,c) {
      if (!value) return ''
      value = value.toString()
      return value.charAt(0).toUpperCase() + value.slice(1)+b,c
    },
  }
})
```

## Vue的 Ajax请求

### Vue的同步Ajax请求

请求的时候是用promise格式的方式进行请求，基于axios 进行ajax请求（使用axios需要在顶部进行引用<script>）

get(链接).then(成功).catch(失败)

这上面的内容写在Vue里面，格式是与微信小程序请求的格式一致。



### Vue的异步Ajax请求

使用Vue完成异步请求，需要单独引用vue-resource库，如下：

```
<script src="https://cdn.staticfile.org/vue-resource/1.5.1/vue-resource.min.js"></script>
```

## Vue自定义组件

### 创建组件

```
<!--
总结：
1.exports 去除s更改为 export。
-->
<template>
  <h1>{{ name }}</h1>
</template>

<script>
export default {
  name: 'TestCom',
  data () {
    return {
      name: 'dff'
    }
  }
}
</script>

```

### 调用组件

**App.vue文件中进行调用**：

```c#
//第三步：在html文件中调用
<mycom />
    
<script>
 //第一步：先引用组件
import mycom from './components/mucom.vue'
export default {
  name: 'App',
  //第二步：添加组件的声明
  components:
  {
    mycom
  }
}
</script>
```

### 动态组件

动态组件很优雅，直接用参数决定需要渲染的组件，如果是用IF ELSE代码非常冗余

<compoent>是没有实际意义的，不会渲染到页面，只是一个插槽，一个作用域，就像一个方便袋子，为了更好的控制，调用里面的内容而已，<template>,<keep-alitp>都是这样子。

### 异步组件

需要的时候再渲染里面的页面，而且会有缓存

```
()=>import(./compentname);
```



### 组件插槽

有默认的插槽，直接声明组件放到里面就可以使用，不限制div的数量，如果组件里面slot没有加名字，放在组件里面很多个地方，都可以渲染相同的内容



也可以声明slot的name然后在组件中就可以声明多个slot，可以在父组件的调用的时候通过name进行调用

### 组件缓存

可以通过<keep-alitp>声明在里面缓存的组件，可以通过name，正则表达式匹配需要缓存的组件，可以声明组件的数量，比如10个缓存的组件，name=‘wangrui，zhangsan’等，如果缓存的时候缓存的组件超过了11个，剔除最久数据没有变化的哪个，把它踢出去

### 组件通讯

#### 父组件通讯

父组件调用子组件的属性与方法

$refs：获取组件中有ref="a1",ref="a2"  名称的等组件，可以通过this.$refs.a1.age可以取到值。

$chilrent：获取当前组件的所有组件，可以通过this.$chilrent[0].age,如果后面在顶部增加了新的组件，索引的位置可能会变，代码会出问题，通过$refs获取内部的组件会比较可靠。

**推荐用：this.$refs比较可靠**

#### 子组件通讯

子组件引用父组件的属性与方法

- 通过事件：this.$limit 访问父组件的时间
- 通过this.$parent对象访问父组件

#### 子组件通讯：依赖注入

父类通过provide暴露实现的方法，属性提供给子组件使用，子组件可以在export default中通过injeck:[show]，进行调用父类的方法（传进去后子组件会自己创建这个变量与方法，与父组件的属性与方法是独立开的）

**PS:可以用于子组件通讯，让子组件来调用父组件的属性与方法**

### 循环引用

#### 递归组件

自己调用自己循环渲染数据，这样的好处是不用在每个调用的位置写for循环循环进行渲染，直接引用组件名称，把组件数据给我，我自己给你渲染好数据

**PS：判断好退出条件，不然可能会有内存溢出问题**

### 循环引用

A组件调用B组件，B组件还没有渲染好：

解决方法1：在B组件中使用Vue里面的生命周期函数BeforeCreate，渲染好才让调用

解决方法2：使用异步import引用组件也可以解决

## 动画

在组件插入，更新，移除的时候，使用过渡效果，比如要怎么显示，要怎么隐藏。一般不会自己写，都会使用第三方的组件进行使用，也可以使用钩子函数，监听过渡时候的生命周期，可以在里面写逻辑

## 混入(mixin)

**使用场景：**提取公共的业务逻辑到mixin中。比如数据校验的方法，数据是否为空，身份证是否正确。与业务有关的建立一个mixin进行管理，与业务无关的建立一个js进行管理，比如转换数字。

把公共数据，方法放到mixin对象里面，使用mixins导入给需要使用的页面与组件，但是使用的时候页面与组件对数据，属性，方法的改动是没有关系的。

**PS：这个应该有大的作用，管理公共的数据，属性，方法**



## 常用API

### this.$KeyCode

通过key键盘触发对应的事件

### this.$nextTick

确保Dom更新后拿到最新的数据，渲染结束后的回调函数

### this.$set

this.arrayname[0]=10; --不会同步更新渲染页面的dom

this.$set(arrayname,0,10);

--会同步更新渲染页面的dom

# WebApi

创建WebApi的时候**勾选OpenApi**自动引用Swagger

```
[Route("/[controller]/[action]")]
```

路由可以放在方法上面（单独作用于方法），也可以放到控制器上面（作用到控制器里面的所有方法），也可以配置的全局的路由上面（作用到所有的控制器）

​	

## HttpGet请求传参

**应用场景：**获取数据

不安全，明面上传递参数，拼接参数到URL尾部传递：

**?userName=wangrui&word=晚上好**

**前端请求：**

```
//传递单个参数（注意字符不需要单引号，双引用）
Test/GetHello?userName=wangrui
//传递多个参数
Test/GetHello?userName=wangrui&word=晚上好
```

**后端接收**：（**通过实际的参数接收**）：

```c#
[HttpGet]
public bool TryLogin(string userName,string word)
{
   return true
}
```

**Ps：**http返回请求 **2开头的返回**就是正确请求到了

## HttpPost请求传参

**应用场景：**新增数据 ，具有冥等操作的行为

安全，背地里打包参数进行传递

**前端调用：**

```c#
axios.post(
     "https://localhost:5001/Login/TryLogin",{userNo:this.userNo,password:this.password})
    .then( res => 
          if(res.status == 200){
                        localStorage["token"] =  res.data.token;
                        location.href = "./index.html"
                   }else{
                       that.errorMsg = "用户名或密码错误";
                    // alert("用户名或密码错误")
                   }
                })
```

**后端接收：**

-  **实体类接收值**：使用UserLoginViewModel实体配合（属性的名称可以忽略大小写）[ApiControl]特性(Net Core3.1加入的特性)对实体进行赋值。

-  **FromBody转换实体**：顶部的**[APIControl]** 特性加上的话，方法入参的时候就不用在前面加 **[FromBody]**

![1637480011382](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637480011382.png)

## HttpPut

**应用场景：**更新数据

**前端方法**

```
//发起请求
axios.put
```

**后端方法**

```
//声明方法
[HttpPut]
```

## HttpDelete

**应用场景：**删除数据

**前端方法**

```
//发起请求
axios.delete
```

**后端方法**

```
//声明方法
[HttpDelete]
```

## HTTP请求总结

**HTTP总结：**

- Resultful风格的HttpGet,HttpPost，HttpPut，HttpDelete在axios中使用get,post,put,delete这四个方法进行请求对应的接口链接，如果是resultful风格的接口，可以不传入具体的方法名称，axios会根据请求的类型（是get或者post）自动调用对应的方法。

- HTTP规范中HttpGet用于获取数据,HttpPost用于新增数据，HttpPut用于更新数据，HttpDelete用于删除数据，但是对资源的增，删，改，查操作，其实都可以通过GET/POST完成，不需要用到PUT和DELETE，很多人贪方便，更新资源时用了GET，因为用POST必须要到FORM（表单），这样会麻烦一些，但是这不符合Http的规范。

**PS：**Post与Get的区别：

- Get方式提交的数据最多是1024字节，浏览器做了限制，而Post不会限制
- Get方法请求的一般都会**冥等**，Post不会，需要特殊处理。
- Get方法请求把参数放到URL进行拼接（英文数字原样发送，中文转换为Base64进行发送），Post放到数据包中进行传输
- 通过GET提交数据，用户名和密码将明文出现在URL上，在浏览器的历史记录，别人就可以拿到你的账户和密码，或者其他关键信息。

**总结：**Get是向服务器发**索取**数据的一种请求，而Post是向服务器**提交**数据的一种请求，是标准，不是强制。

## 定义路由规则

**改变路由规则**：用于Resultful风格接口（HttpGet,HttpPost，HttpPut，HttpDelete）中命名相同的方法，可以用这个方法改变接口的命名去除冲突，声明很多HttpGet,HttpPost等方法。

```
//hahah 是调用展示的方法名称
[HttpGet("hahah")]
```

![1637491186741](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637491186741.png)

**强制必填参数：**

```
 //参数名要与方法入参名称一致
 [HttpGet("{userNo}/{password}")]
```

![1637490940087](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637490940087.png)

## Token存储位置

**LocalStorage：**本地存储，永久存储，刷新关闭浏览器都不会丢失,session关闭页面，关闭浏览器就丢失了，Cookie一般默认是浏览器关闭失效，可以自己设置失效时间。做登录操作后，从后台生成Token（生成GUID）进行返回，然后把Token放到LocalStorage本地缓存中，在vue对象中的data中新建一个属性**token:LocalStorage["token"]。**做权限操作的时候，进行验证是否允许操作，如果没有登录就跳转到登录界面。每次向后端发起请求的时候，都把Token传到参数中进行验证权限。

[cookie、sessionStorage和localStorage的区别]: https://blog.csdn.net/weixin_42614080/article/details/90706499

![1637494658957](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637494658957.png)

## Swagger Core 

Swagger Core的使用（创建WebApi的时候勾选OpenApi自动引用Swagger）  https://blog.csdn.net/weixin_42550800/article/details/94384296 

## 数据验证

在实体类中标记特性进行验证（好像是ApiControl特性的功能）

```
//最大长度是2
[MaxLength(2)]
```

## 常用Linq方法

**转换集合：**ToList

**投影到对象：**Select

**调用结果选择器投影到对象：**SelectMany

**去重：**Distinct

**是否存在：**Contains

**排序：**OrderByDescending

**根据条件筛选：**Where

**返回第一个元素如果没有值就返回默认值：**FirstOrDefault

**返回最小值：**Max

**返回最大值：**Min

**直接在对接上ForEach：**variationDtos.ForEach

## 常见的Http状态码 HTTP请求

![1637493721855](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637493721855.png)