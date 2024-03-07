# 高级班：ASP.NET CORE MVC

## 页面传值的5种方法：

```c#
    //控制器方法  
	public IActionResult Index()
    {
        //1.【知识点1】通过ViewBag静态属性传值
        base.ViewBag.User1 = "张三";
        //2.【知识点2】通过ViewData键值对传值
        base.ViewData["User2"] = "李四";
        //3.【知识点3】通过TempData传值：这里面的值只能获取一次
        base.TempData["User3"] = "王五";
     
        //4.【知识点4】1.通过Session传值，需要引用：Microsoft.AspNetCore.Http 命名空间
          //2.在Startup.cs中
          //在ConfigureServices写： services.AddSession();
         // 在Configure中写： app.UseSession(); 
         
          //2.打完HttpContext.Session.GetString(); 才可以引用这个命名空间
        if (HttpContext.Session.GetString("User5") == null)
        {
            HttpContext.Session.SetString("User5", "田七");
        }
        //5.【知识点5】通过Model对象 return view传值
        object user4 = "赵六";
        return View(user4);
    }

   //页面调用
        @model System.String;    //【知识点6】注意这里声明是小写，引用@Model是大写
        @using Microsoft.AspNetCore.Http;


        <h1>
            @base.ViewBag.User1
        </h1>
        <h1>
            @base.ViewData["User2"]
        </h1>
        <h1>
            @base.TempData["User3"]
        </h1>
        <h1>
        //【知识点7】引用的是@Model，大写M
           @Model
        </h1>

        <h1>
           //【知识点8】注意这里是Context，不是HttpContext
           @Context.Session.GetString("User5");
        </h1>
```

## Startup.cs文件解析：

**ConfigureServices方法**:注入接口实例：比如注入对象构造函数中的接口的实例，注册Sesion引用，IOC注册服务

**Configure方法**：配置中间件：比如配置Session，配置路由：Controller/Action

## Log4Net的使用

1.在NuGet中引用依赖包：Log4Net  用最新的稳定版本（2.0.12）

2.引用Log4Net在Core中的依赖包：Microsoft.Extensions.Logging.Log4Net.AspNetCore 用最新的稳定版本（5.0.4）

3.配置log4net.config文件，比如日志文件的大小，是否允许多线程写入，日志文件保存，日志保存的路径的命名规范

```c#
<?xml version="1.0" encoding="utf-8"?>
<log4net>
	<!-- Define some output appenders -->
	<appender name="rollingAppender" type="log4net.Appender.RollingFileAppender">
	    <!--保存文件的路径，以根目录为初始节点-->
		<file value="Customlog\log.txt" />
		<!--追加日志内容-->
		<appendToFile value="true" />

		<!--防止多线程时不能写Log,官方说线程非安全-->
		<lockingModel type="log4net.Appender.FileAppender+MinimalLock" />

		<!--可以为:Once|Size|Date|Composite-->
		<!--Composite为Size和Date的组合-->
		<rollingStyle value="Composite" />

		<!--当备份文件时,为文件名加的后缀-->
		<datePattern value="yyyyMMdd.TXT" />

		<!--日志最大个数,都是最新的-->
		<!--rollingStyle节点为Size时,只能有value个日志-->
		<!--rollingStyle节点为Composite时,每天有value个日志-->
		<maxSizeRollBackups value="20" />

		<!--可用的单位:KB|MB|GB-->
		<maximumFileSize value="3MB" />

		<!--置为true,当前最新日志文件名永远为file节中的名字-->
		<staticLogFileName value="true" />

		<!--输出级别在INFO和ERROR之间的日志-->
		<filter type="log4net.Filter.LevelRangeFilter">
			<param name="LevelMin" value="ALL" />
			<param name="LevelMax" value="FATAL" />
		</filter>
		<layout type="log4net.Layout.PatternLayout">
			<conversionPattern value="%date [%thread] %-5level %logger - %message%newline"/>
		</layout>
	</appender>
	<root>
		<priority value="ALL"/>
		<level value="ALL"/>
		<appender-ref ref="rollingAppender" />
	</root>
</log4net>

```

4.在Program.cs文件中编写Log4Net.config的读取代码

  **读取配置的第一种方法：**Program

```c#
                public static IHostBuilder CreateHostBuilder(string[] args) =>
            Host.CreateDefaultBuilder(args)
                 //配置log4Net文件的读取
                .ConfigureLogging(loggingBuilder =>
                {  //设置log4net.config的读取路径
                    loggingBuilder.AddLog4Net("CfgFile/log4net.Config");
                })
                .ConfigureWebHostDefaults(webBuilder =>
                {
                    webBuilder.UseStartup<Startup>();
                });
```

**读取配置的第二种方法：Startup.cs中**

```
 //添加ILoggerFactory loggerFactory参数，注意是来自命名空间：Microsoft.Extensions.Logging
 public void Configure(IApplicationBuilder app, IWebHostEnvironment env, ILoggerFactory loggerFactory)
        {
         //读取Log4Net的配置，放入配置文件的路径
        loggerFactory.AddLog4Net("CfgFile/log4net.Config");

        }
```

5.log4Net的日志写入：

```
    private readonly ILogger<HomeController> _logger;
     //通过构造函数注入Log4Net的实例
    public HomeController(ILogger<HomeController> logger)
    {
        _logger = logger;
        //向本地文件写入日志，每次写入都是向现有的日志进行追加，写入的信息有时间（精确到毫秒）+线程+日志内容
        _logger.LogWarning("HomeController被构造。。。");
    }
   
    public IActionResult Index()
    { 
       //向本地文件写入日志
        _logger.LogInformation("this is HomeController.Index");
    }
```

## 安装IIS

**安装步骤：**控制面板→程序→启用或关闭windows功能→Internet  Information Services（勾选全部点击确定）

**勾选全部提示错误：**参照的程序集没有安装在系统上，错误代码xxxx，

​                                  解决方案：取消部分内容的勾选后，重新安装成功！





## IIS发布 

**直接指向项目目录，不行：**通过IIS建立网站，目录直接指向项目根目录下的Debug/Net5文件，发现不行，网站不能正常访问（显示目录浏览权限，有个偏方可以解决:通过拷贝vs发布时生成的web.config到根目录，问题解决）

**通过VS发布项目，可以**：把项目发布以后，将IIS上面的网站目录指向项目发布目录 ，可以正常运行

**不行的原因**通过VS发布的项目会生成web.config文件，程序生成的项目问题下默认没有web.config文件：

```
<?xml version="1.0" encoding="utf-8"?>
<configuration>
  <location path="." inheritInChildApplications="false">
    <system.webServer>
      <handlers>
        <!--通过AspNetCoreModuleV2这个模块启动站点-->
        <add name="aspNetCore" path="*" verb="*" modules="AspNetCoreModuleV2" resourceType="Unspecified" />
      </handlers>
      <!--通过dotnet 命令，执行.\WebApplication1.dll这个文件，与下面通过脚本启动的结果一致-->
      <aspNetCore processPath="dotnet" arguments=".\WebApplication1.dll" stdoutLogEnabled="false" stdoutLogFile=".\logs\stdout" hostingModel="inprocess" />
    </system.webServer>
  </location>
</configuration>
```

如果IIS上面缺失AspnetcoreModuleV2模块，那么需要单独安装：

下载路径 ：**https://dotnet.microsoft.com/download/dotnet/5.0** 

下载安装：Hosting Bundle 需要根据自己电脑的系统位数来决定 

![1636858290079](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1636858290079.png)

下载Hosting Bundle 64位后,从VS中启动网站出现（没下载之前是正常的）：**HTTP 错误 500.31 - 无法加载 ASP.NET Core 运行时**

重新安装：**ASP.NET Core 5.0 Runtime (v5.0.12)** 解决：https://dotnet.microsoft.com/download/dotnet/thank-you/runtime-aspnetcore-5.0.12-windows-hosting-bundle-installer

猜测是安装的内容不全面，把原来配置好的环境进行了覆盖，导致的问题，重新安装后发现IIS发布的站点由错误：**无法访问请求的页面，因为该页的相关配置数据无效**。 重新安装后也变成可以访问了（需要在IIS站点中右键→编辑权限→安全→编辑→新增Everyone用户读取权限）。



ASPNET Core跨平台的原因：内置Kestrel（k（二声） s 肿）小型的服务器

![1634733922086](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1634733922086.png)

### 脚本启动 

1. 打开cmd命令

2. cd  指向项目文件路径

3. Cmd命令行执行：

   ```c#
   dotnet AspNetCore.Web.dll --urls="http://*:8079" 
   
   --ip="127.0.0.1"(本地IP) --port= 8079 
   （--ip,--port是参数，https://*:8079中*号代表本地IP,8079代表是站点的IP）
   ```

   也可以使用命令行缩写进行启动：

   ```c#
   //使用缩写进行启动
   dotnet run
   ```

    

4. 使用**ctrl+c** 在命令行窗口进行关闭启动，要快速按两下才会关闭



可以使用**dotnet help** 列出所有可用的命令行

## 读取静态文件：css,jquery

配置静态文件读取

startup.Configure写入读取静态文件的中间件代码（app.UseStaticFiles是原来已经存在startup.Configure中）：

```c#
//读取静态文件的中间件，引用命名空间：Microsoft.Extensions.FileProviders 
app.UseStaticFiles(new StaticFileOptions()
{
//读取根目录文件下的wwwroot文件夹内的文件
 FileProvider = new PhysicalFileProvider(Path.Combine(Directory.GetCurrentDirectory(), "wwwroot"))  
 }
 );
```





## 读取脚本参数:dotnet

1. 启动cmd命令，使用cd  项目根目录，进入Cd项目根目录

2. Cmd命令行执行：

   ```
   dotnet AspNetCore.Web.dll --urls="https://*:8079" --age=20 --name="张三" 
   ```

    在启动dll的脚本之后加入参数，格式：**--参数名称=参数值，--参数名称="参数值"**    注意-- 是必须的内容，参考上面执行的dotnet代码

3. 在控制器的构造函数中编写代码，我们可以通过构造函数进行注入IConfiguration

   ```c#
   private readonly IConfiguration _Configuratin;
    
    public FirstController( IConfiguration configuration)
    {
        _Configuratin = configuration;
    }
    
    // IConfiguration在Startup构造函数中已经默认注入
     public class Startup
     {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }
     }
   ```

4. 在控制器中action方法中，可以通过IConfiguration[参数名称]进行读取脚本中的参数

   ```
   public IActionResult Index()
   {
      //读取脚本参数age
       ViewBag.age = _Configuratin["age"]; 
       //读取脚本参数name
       ViewBag.name = _Configuratin["name"];
       return View();
   }
   ```

## 读取配置文件:appsettings.json

startup.ConfigureServices中读取appsettings.json配置文件，**需要注入Configuration（见上面脚本参数的注入过程）**，进行读取appsettings.json配置的参数

appsetings.json内容：

```
{
  "Logging": {
    "LogLevel": {
      "Default": "Information",
      "Microsoft": "Warning",
      "Microsoft.Hosting.Lifetime": "Information"
    }
  },
  "AllowedHosts": "*",
  "Id": "123456",
  "Name": "Richard老师",
  "TeachInfo": {
    "Id": 123456,
    "Name": "Richard001"
  },
  "ConnectionStrings": {
    "WriteConnection": "Server=LAPTOP-JU1DEJP1;Database=ZhaoxiDBSet;Trusted_Connection=True;",
    "ReadConnectionList": [
      "Server=LAPTOP-JU1DEJP1;Database=ZhaoxiDBSet01;Trusted_Connection=True;",
      "Server=LAPTOP-JU1DEJP1;Database=ZhaoxiDBSet02;Trusted_Connection=True;",
      "Server=LAPTOP-JU1DEJP1;Database=ZhaoxiDBSet03;Trusted_Connection=True;"
    ]
  } 
}
```

都可以通过注入Configuration类进行实现

**第一种方式：直接读取**

- 读取字符串配置，**如果读取的是乱码，更改Json的编码格式为UTF-8即可。**

  ```
     Console.WriteLine($"Id:{Configuration["Id"]}");
     Console.WriteLine($"Name:{Configuration["Name"]}");
  ```

  

- 读取类属性配置**（注意冒号）**

  ```
   Console.WriteLine($"TeachInfo.Id:{Configuration["TeachInfo:Id"]}");
   Console.WriteLine($"TeachInfo.Name:{Configuration["TeachInfo:Name"]}");
  ```

  

- 读取数组配置（**注意冒号与索引**，用索引appSettings.json里面会有警告是正常的）

```c#
        Console.WriteLine($"ConnectionStrings.WriteConnection:{Configuration["ConnectionStrings:WriteConnection"]}");
        Console.WriteLine($"ConnectionStrings.ReadConnectionList1:{Configuration["ConnectionStrings:ReadConnectionList:0"]}");
        Console.WriteLine($"ConnectionStrings.ReadConnectionList2:{Configuration["ConnectionStrings:ReadConnectionList:1"]}");
        Console.WriteLine($"ConnectionStrings.ReadConnectionList3:{Configuration["ConnectionStrings:ReadConnectionList:2"]}");
```

**第二种方式：通过实体类读取**

1.定义一个与配置文件格式内容一致的实体类

```
public class DbConnectionOptions
{
    public string WriteConnection { get; set; }

    public List<string> ReadConnectionList { get; set; }
}
```

2.在Startup.ConfiguServices中第一行（案例中是）写入：

```c#
//读取配置中间中ConnectionStrings节点下面的内容，并将读取到的值转换成实体对象类型(这里的Configuration实例还是用开始注入的那个对象进行读取)
 services.Configure<DbConnectionOptions>(Configuration.GetSection("ConnectionStrings"));
```

3.需要在控制器中使用的时候，通过IOptions<DbConnectionOptions>进行注入,构造函数中获取到的注入的**Options.Value** 就是我们获取的到配置文件的一个实体类型对象 

**碰到错误：**无法解析类型'ASP.NETCore1.Class的服务。当尝试活'ASP.NETCore1.Controllers.HomeController'时

注意控制的的入参是：**IOptions options**

```
   private readonly DbConnectionOptions _optionsCurrent;

    public FirstController(IOptions<DbConnectionOptions> options)
    {
        _optionsCurrent = options.Value;
    }
```

4.使用配置文件实体对象的值，序列化成Json对象进行读取使用

```
    public IActionResult Index()
    {
        object strResult= Newtonsoft.Json.JsonConvert.SerializeObject(_optionsCurrent); 
        return View(strResult);
    }
```

**注意：**放到View中进行返回的变量，在首页通过@model  对应类型（比如这里是object，也可以是实体类的名称）进行引用：

```c#
//cshtml顶部进行引用，注意@model是小写，object代表实体类类型（如果是object类型，那么就是object）
@model object 
//调用值，注意这里是大写@Model
<h1>@Model</h1> 
```







## MVC解读Razor混合编码：

### 什么是MVC？ 

- **M视图模型**：用作控制器和视图之间传递数据的载体 
- **V视图**：呈现给用户看到的内容（表现层） 
- **C控制器**：控制业务逻辑计算，调用服务，选择返回什么内容，可以返回视图，JSON,字符串，文件等等 

#### 控制器Action返回类型

1. **ViewResult ：**代表 HTML 和标记。用于输出视图内容：

   ![img](https://upload-images.jianshu.io/upload_images/13808716-f456df7debc14df3.png)

2. **EmptyResult** ： 代表无结果。

3. **RedirectToAction：**可以跳转到指定的action（可以是非本控制器）

    

   ![img](https://upload-images.jianshu.io/upload_images/13808716-a336b7cc69bf0cae.png)

4. **RedirectResult**  ：代表重定向页面,跳转到新的页面

   ![img](https://upload-images.jianshu.io/upload_images/13808716-46323f87b0771682.png)

5. **RedirectToRouteResult**  ： 代表重定向到新的控制器的

6. **

   ![img](https://upload-images.jianshu.io/upload_images/13808716-631e2b0f7b912f91.png)

7. **JsonResult**  ： 输出json字符串，一般前端通过ajax请求，返回的就是这个Json

   ![img](https://upload-images.jianshu.io/upload_images/13808716-03121b94be816382.png)

8. **ContentResult**  ：代表文本结果。

    

   ![img](https://upload-images.jianshu.io/upload_images/13808716-34b6d55f99985504.png)

    

    

   ![img](https://upload-images.jianshu.io/upload_images/13808716-2cd0b9c1e87bf143.png)

9. **IActionResult**：**所有Result的父类**



#### 控制器Action返回类型总结

1. 上面7种返回类型是MVC/API中返回的比较常见的几种类型，**但不管是ViewResult还是ContentResult他们所有的类型的基类都是IActionResult，我们在写方法返回值的时候可以不必写出具体的返回类型，直接写IActionResult也是可以的**。也可以写平常我们所见的int，string，或者void亦或者其他类型

2. 重定向的类型

   （1）Redirect（页面路径）

   （2）RedirectToAction（重定向到指定的控制器、动作方法）

   （3）RedirectToRoute（使用指定的路由值跳转）

   

3. **重点是下面文件上传下载中的知识点，也就是文件资源的输出内容**

   （1）直接使用项目中的路径资源，通过路径读取文件，输出文件。注意：需要使用虚拟（相对）路径。

   （2）使用字节数组，输出文件资源。需要使用物理（绝对）路径。

   （3）使用流，输出文件资源。需要使用物理（绝对路径）。

   

   

   #### 文件输出

   1.按指定的文件路径来输出文件：

   注意：用此方法，没有涉及到文件流，不需要读取硬盘上的文件，所以直接使用虚拟路径，即相对路径

    

   ![img](https://upload-images.jianshu.io/upload_images/13808716-e2fe36d89433d06e.png)

   测试结果如下：

    

    

   ![img](https://upload-images.jianshu.io/upload_images/13808716-0b00b1c30ef017be.png)

   2.使用字节数组输出文件：物理路径

    

   ![img](https://upload-images.jianshu.io/upload_images/13808716-ea501bb4e9f4df9e.png)

   测试结果如下：

    

   ![img](https://upload-images.jianshu.io/upload_images/13808716-2d77dfd2f829de2c.png)

   此为在线浏览方式，如若需要下载的方式，则在返回方法加上第三个参数，下载文件的命名：物理路径

    

   ![img](https://upload-images.jianshu.io/upload_images/13808716-a6c260d51fb193cf.png)

    

    

   ![img](https://upload-images.jianshu.io/upload_images/13808716-e8cb9ad3fc691ced.png)

   3.使用流方式输出文件：物理路径

   ![img](https://upload-images.jianshu.io/upload_images/13808716-2800bf54ae7aef8f.png)

   

    

   ###  

   ###  

#### AJAX语法

##### 添加

```c#
//添加         <input type="button" value="新增学生信息" onclick="AddPost()" />
function AddPost() {
    var paramerData = {
        studentID: $("#studentID").val(),
        Name: $("#Name").val()
    }
    $.ajax(
        {
            url: "/Students/AddStudent",
            type: "post",
            data: paramerData,
            success: function (data) {
                if (data.Succeed == 1) {
                    alert('添加学生信息成功');
                    SelectPost("ADD");
                } else {
                    alert(data.Message);
                }

            }

        });
}


```



##### 删除

~~~c#
//删除         <input type="button" value="删除学生信息" onclick="DeletePost()" />
function DeletePost() {
    $.ajax(
        {
            url: "/Students/DeleteStudent",
            type: "post",
            data: { Name: $("#Name").val(), studentID: $("#studentID").val()},
            success: function (data) {
                if (data.Succeed == 1) {
                    alert('删除学生信息成功');
                    SelectPost("delete");
                } else {
                    alert(data.Message);
                }
            }

```
    });
```

}
~~~



##### 查询

```c#
//查询         <input type="button" value="查询学生信息" onclick="SelectPost()" />
function SelectPost(obj) {
    var paramerData = null;
    if (typeof(obj) =="undefined") {
        paramerData = {
            studentID: $("#studentID").val(),
            Name: $("#Name").val()
        }
    }

    $.ajax(
        {
            url: "/Students/SelectStudent",
            type: "post",
            data: paramerData,
            success: function (data) {
                var studentDom = '';
                var rows = data.Data;
                $.each(rows, function (index, val) {
                    studentDom += '<p>' + '姓名：' + this.Name + '，学生编号：' + this.StudentID + '</p>';
                });
                $("#Studentxx").html(studentDom)
            }

        });
}
```

#### 调试的方法

![1636899054875](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1636899054875.png)

**调试的方法2**

查看IIS的日志，在站点→IIS→日志→浏览

![1636900365054](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1636900365054.png)

也可以通过windows的事件查看器，windows日志→应用程序

![1636900403049](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1636900403049.png)

#### Razor基础语法

cshtml其实是一个类文件，里面可以写html，也可以写后台代码

`见案例：Views\Second\Index.cshtml`

需要写入后台代码，记得写一个**艾特符号@**，我们可以在cshtml文件中

```
//顶部的using命名控件
@using Zhaoxi.NET5Demo.Project.Utility.RazorGrammar
@using Zhaoxi.NET5Demo.Interface;
```



- 继承接口，并且进行实现

  ```
  //声明接口
  @implements CustomInterface
  @functions {
      //实现接口
      public void Show()
      {
          Console.WriteLine("实现接口");
      }
  
  }
  //调用实现接口
  @inject ITestServiceA iITestServiceA
  @{
      iITestServiceA.Show();
      var intResult = Model;
  }
  ```

  

- 还可以实现依赖注入

- 增加特性：Attribute

  ```
  @*添加特性*@
  @attribute [CustomInfoAttribute]  //因为是个类文件，所以也可以增加特性
  ```

  

- 声明方法

  ```
  //声明方法
  @functions	{
      public string GetHello()
      {
          return "functions----Hello";
      }
  }
  //声明模板化方法：包含html元素
  @*在代码块中，使用标记将本地函数声明为用作模板化方法：*@
  @{
      void RenderName(string name)
      {
          <p>Name: <strong>@name</strong></p>
      }
  
  RenderName("模板化方法:Mahatma Gandhi");
  RenderName("模板化方法:Martin Luther King, Jr.");
  
  }
  ```

  

- switch

  ```
  @for (int l = 0; l < 5; l++)
  {
      @switch (l)
      {
          case 1:
              <a href="http://www.zhaoxiedu.net/">这里是朝夕官网链接 @l</a>
              break;
          case 2:
              <a href="http://www.zhaoxiedu.net/">这里是朝夕官网链接 @l</a>
              break;
          case 3:
              <a href="http://www.zhaoxiedu.net/">这里是朝夕官网链接 @l</a>
              break;
          case 4:
              <a href="http://www.zhaoxiedu.net/">这里是朝夕官网链接 @l</a>
              break;
          case 5:
              <a href="http://www.zhaoxiedu.net/">这里是朝夕官网链接 @l</a>
              break;
          default:
              break;
      }
      <br>
  }
  ```

  

- if else 

  ```
  @for (int l = 0; l < 10; l++)
  {
      <a href="zhaoxiedu.net">这是朝夕官网</a> //html
      if (l == 0)  //这个是后台代码
      {
          <a href="www.baidu.com">这里是百度链接 @l</a>
      }
      else if (l == 2)
      {
          <a href="http://www.zhaoxiedu.net/">这里是朝夕官网链接 @l</a>
      }
      else
      {
          <a href="http://www.zhaoxiedu.net/">这里是朝夕官网链接 @l</a>
      }
      <br>
  }
  ```

  

- 在cshtml中调用方法：

  ```
  @functions	{
      public string GetHello()
      {
          return "functions----Hello";
      }
  }
  <div>From method: @GetHello()</div>
  ```

  

- 使用DateTime.Now

  ```
  @*显式 Razor 表达式*@
  <P>显式 Razor 表达式: @(DateTime.Now - TimeSpan.FromDays(7))</p>
  ```

- 使用foreach循环

  ```
  @{
      List<int> intlist = new List<int>();
  }
  @foreach (var item in intlist)
  {
  
  }
  ```

  

- 使用for循环

  ```
  @for (int l = 0; l < 10; l++)
  {
      <a href="zhaoxiedu.net">这是朝夕官网</a> //html
      if (l == 0)  //这个是后台代码
      {
          <a href="www.baidu.com">这里是百度链接 @l</a>
      }
      else if (l == 2)
      {
          <a href="http://www.zhaoxiedu.net/">这里是朝夕官网链接 @l</a>
      }
      else
      {
          <a href="http://www.zhaoxiedu.net/">这里是朝夕官网链接 @l</a>
      }
      <br>
  }
  ```

  



标记的@符号就是后台代码，如果使用多行后台代码可以使用@{}作用域里面写多行的后台代码，如果@{}中包含html的元素那么在{作用域中的}后台代码仍然需要使用@符号

#### 动态渲染cshtml

改动cshtml中的代码内容，需要重新启动解决方案才可以作用到，不像改动aspx中的Html内容可以马上生效，日常开发中这个需求很重要，我们可以通过配置中间件进行支持：

1**.Nuget 安装：**

```
Microsoft.AspNetCore.Mvc.Razor.RuntimeCompilation
```

2.**在startup.ConfigureServices写入代码**：

```
   //解决修改视图内容，必须编译后方可生效的问题
   services.AddRazorPages().AddRazorRuntimeCompilation();
```

不只是可以渲染html页面的代码，增加了@的后台的代码也可以支持不重启解决方案就可以渲染生效

3.如果改动了的是后台代码，需要点击**这团火**进行生效![1636854286280](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1636854286280.png)



#### Razor布局

 Views.sharted _Layout



顶部

左边菜单+cshtml中部

底部

**暂时看其他的内容**

013.NET5_MVC_Razor布局
014.NET5_Razor扩展01
015.NET5_Razor扩展02
016.NET5_Razor局部视图

017.NET5_视图组件扩展定制

## **IOC容器IServiceCollection**

.NET5是内置IOC容器的； 

### 什么是IOC？ 

把对象的创建统一交给第三方容器来创建； 

### 什么是依赖注入呢？ 

控制器对象依赖于接口B，对象A继承了接口B，可以先构造接口B的属性，然后在程序执行的时候将对象A传递给接口B，控制器对象可以得到对象A的实例。

### IOC容器的三种注入方式：

#### IServiceCollection：构造函数注入

IServiceCollection（磕 l k 藓）

1.在Startup中的ConfigureServices 方法中注册服务，

```
         //注册抽象和具体的依赖关系，ITestServiceA是抽象，TestServiceA是具体的实现
       services.AddSingleton<ITestServiceA, TestServiceA>();
```

2.在需要使用实例的控制器中，通过构造函数入参定义实例的抽象实例，在运行时，Startup通过执行AddSingleton注入实例到抽象实例，在使用控制器时我们会自动拿到服务的实现实例； 

```
 private readonly ITestServiceA _ITestServicA = null; 
 
 public FifthController(ITestServiceA iTestServicA) 
 { 
  _ITestServicA = iTestServicA;  
 } 
```

3.调用通过抽象注入的实例方法.Show()

```
public IActionResult Index() 
{  
   _ITestServicA.Show();  
   return View(); 
} 
```

DI依赖注入：**IServiceCollection支持且仅支持构造函数注入** 

IServiceCollection：IServiceCollection是内置的IOC容器， 可以支持无线层级的依赖注入（构造函数中可以有多个接口，然后通过startup.configureServices中进行注入）， 前提是都要先注入服务（注册抽象和具体的映射关 

系）

#### 通过IServiceProvider进行注入

IServiceProvider（泼(浒厂) 歪（二声） 的）

1.在需要使用的控制器中，通过构造函数注入IServiceProvider （**奇怪的是视频中并没有在startup.ConfigureService中进行实例注入，可能是在下一个视频中讲**）

```
//注入的参数修饰符是readonly
private readonly IServiceProvider _ServiceProvider = null; 

public FifthController(IServiceProvider serviceProvider) 
{  
 _ServiceProvider=serviceProvider 
} 
```

3.通过_ServiceProvider获取到服务，然后通过服务实例调用服务内部的方法 

```c#
public IActionResult Index() 
{   
    //注意这里需要对获取到的接口对象进行(ITestServiceA)强制转换
   ITestServiceA testServiceA = (ITestServiceA)_ServiceProvider.GetService(typeof(ITestServiceA)); 
   testServiceA.Show(); 
} 
```

#### 通过Inject关键字注入cshtml

Inject（in jack t）

在注册服务后，视图中通过关键字@Inject 获取实例 

```
@inject ITestServiceA iTestServicA   ---获取到服务实例 
@{ 
   iTestServicA.Show(); 
 } 
```



### **IServiceCollection三种生命周期** 

那么在创建对象的时候，不同的情况下需要让对象有不同的作用，比如需要让对象单例，在整个应用程序中做一样的事情，或者每一次使用都创建新的对象实例，都有新的对象实例，或者A,B类使用对象这个实例，CD类使用对象另外的注入实例（拿写日志的场景来说，比如A,B类是敏感的交易类，需要往服务器写入日志，CD类是日常的模块操作，记录的日志写在客户本地就可以，所以它们需要注入不同的实现实例）。针对于应用的需要进行创建新的实例。 

**建议：**开发工作中，一般情况下，都是一次请求一个对象的实例，更多的是瞬时生命周期的使用； 

#### 瞬时生命周期

每一次getService获取的实例都是不同的实例 

```
   // 瞬时生命周期
    {
        IServiceCollection serviceCollection = new ServiceCollection();
        //AddTransient：瞬时生命周期，每一次getService获取的实例都是不同的实例
        serviceCollection.AddTransient<ITestServiceA, TestServiceA>(); 
        ServiceProvider serviceProvider = serviceCollection.BuildServiceProvider();
        ITestServiceA testServiceA = serviceProvider.GetService<ITestServiceA>();
        ITestServiceA testServiceA1 = serviceProvider.GetService<ITestServiceA>();
        bool isOK = object.ReferenceEquals(testServiceA, testServiceA1); //结果为false； 两次获取的对象不是同一个实例
    }
```

#### 单例生命周期

在整个进程中获取的都是同一个实例 

```
    //单例生命周期
    {
        IServiceCollection serviceCollection = new ServiceCollection();
        //AddSingleton:单例生命周期，在整个进程中获取的都是同一个实例
        serviceCollection.AddSingleton<ITestServiceA, TestServiceA>(); 
        
        ServiceProvider serviceProvider = serviceCollection.BuildServiceProvider();
        ITestServiceA testServiceA = serviceProvider.GetService<ITestServiceA>();
        ITestServiceA testServiceA1 = serviceProvider.GetService<ITestServiceA>();
        //结果为 true，是同一个引用，在整个进程中获取的都是同一个实例
        bool isOK = object.ReferenceEquals(testServiceA, testServiceA1); 
    }
```

#### 作用域生命周期

同一个作用域，获取的是同一个对象的实例；不同的作用域，获取的是不同的对象实例  

```
    // 作用域生命周期
    {
        IServiceCollection serviceCollection = new ServiceCollection();
        //AddScoped:作用域生命周期；同一个作用域，获取的是同一个对象的实例；不同的作用域，获取的是不同的对象实例
        serviceCollection.AddScoped<ITestServiceA, TestServiceA>(); 
        ServiceProvider serviceProvider = serviceCollection.BuildServiceProvider();
        ITestServiceA testServiceA = serviceProvider.GetService<ITestServiceA>();
        ITestServiceA testServiceA1 = serviceProvider.GetService<ITestServiceA>();

        bool isOK = object.ReferenceEquals(testServiceA, testServiceA1);

        ServiceProvider serviceProvider1 = serviceCollection.BuildServiceProvider();
        ITestServiceA testServiceA2 = serviceProvider1.GetService<ITestServiceA>();
        bool isOK1 = object.ReferenceEquals(testServiceA1, testServiceA2);
    }
```



​        

### **ConfigureServices的生命周期**

```c#
    public void ConfigureServices(IServiceCollection services)
    {
        //AddTransient ：瞬时生命周期
        services.AddTransient<IUserService, UserService>();
        //AddScoped ：作用域生命周期(我看到辉哥用来注册SqlSugar的服务对象，但是注册后并没有用对象去接收返回值，说明这个注册是声明全局使用SqlSugar的对象使用的是那种生命周期)
        services.AddScoped<IArticleService, ArticleService>();
        //AddSingleton：单例生命周期
        services.AddSingleton<IProductService, ProductService>();
    }
```

- `Transient`：每次请求 *Transient* 生命周期服务时都会创建它们。**此生命周期最适合轻量级、无状态的服务**。`ransient`类型的服务，每次注入的对象都是新的对象，即使在同一个类，同一个构造函数当中两个相同的对象也不一样。

```c#
    //AddTransient ：瞬时生命周期
    services.AddTransient<IUserService, UserService>();
```

- `Scoped`：在同一个`Scope`内只初始化一个实例 ，可以理解为在同一个作用域内的生命周期，重复创建相同的对象是一样的，**实体框架上下文（Entity Framework context）是使用 *Scoped* 服务的一个很好的场景**

```c#
    //AddScoped ：作用域生命周期
    services.AddScoped<IArticleService, ArticleService>();
```

- `Singleton`：创建服务类的单个实例，保存在内存中，可以在整个应用程序中重复使用。可**以用在实例化成本非常高的对象时使用**。使用Singleton单例模式创建的服务，在相同的控制器中，不同的控制器中调用都是一个类。

```c#
    //AddSingleton：单例生命周期
    services.AddSingleton<IProductService, ProductService>();
```

理解了生命周期的作用，我们在开发的时候就可以根据需要对不同的服务选择不同的生命周期

**如果单个对象没有多个实现的情况下是不是都使用单例的生命周期。**

答：不行，因为单例的生命周期，变量全部保存在内存当中，如果全部都放在内存当中损耗的性能是非常大的。

## Autofac

Autofac 是一个流行的第三方IOC容器，使用方法如下

1.NuGet中引用依赖包：

2.startup.ConfigureService实例化对象进行使用：

```
   //构造函数注入
   ContainerBuilder containerBuilder = new ContainerBuilder();
   containerBuilder.RegisterType<TestServiceA>().As<ITestServiceA>();
   containerBuilder.RegisterType<TestServiceB>().As<ITestServiceB>();
   containerBuilder.RegisterType<TestServiceC>().As<ITestServiceC>();
   //使用上面注册的组件，创建一个容器属性container：保存抽象与实例映射关系的属性容器
   IContainer container = containerBuilder.Build();iTestServiceA的实例也就是
   //获取容器中抽象ITestServiceB对象的实例，获取实例的时候TestServiceB的构造函数会获取到iTestServiceA的实例也就是TestServiceA
   ITestServiceB testServiceb = container.Resolve<ITestServiceB>();
   testServiceb.Show();
   
   //类TestServiceB的内容
   public class TestServiceB : ITestServiceB
   {
        public ITestServiceA _iTestServiceA = null;
     
         public TestServiceB(ITestServiceA iTestServiceA)
        {
            Console.WriteLine("container.Resolve<ITestServiceB>获取实例的时候，会获取到iTestServiceA的实例也就是TestServiceA");
        }
   }
```

IOC容器是保存对象抽象与实例关系映射的一个容器吗

### Autofac三种注入方式：方法，属性，构造函数

- 构造函数注入

  ```
  //构造函数注入
  ContainerBuilder containerBuilder = new ContainerBuilder();
  containerBuilder.RegisterType<TestServiceA>().As<ITestServiceA>();
  containerBuilder.RegisterType<TestServiceB>().As<ITestServiceB>();
  containerBuilder.RegisterType<TestServiceC>().As<ITestServiceC>();
  //使用上面注册的组件，创建一个容器属性container：保存抽象与实例映射关系的属性容器
  IContainer container = containerBuilder.Build();
  //获取容器中抽象ITestServiceB对象的实例
  ITestServiceB testServiceb = container.Resolve<ITestServiceB>();//获取服务
  testServiceb.Show();
  ```

- 属性注入

  ```c#
          ContainerBuilder containerBuilder = new ContainerBuilder();
          containerBuilder.RegisterType<TestServiceA>().As<ITestServiceA>();
          containerBuilder.RegisterType<TestServiceB>().As<ITestServiceB>();
          containerBuilder.RegisterType<TestServiceC>().As<ITestServiceC>();
          //调用PropertiesAutowired()方法，实例化类里面注册符合要求的全部属性，这些类里面的属性肯定是包含在IOC容器里面的
          containerBuilder.RegisterType<TestServiceD>().As<ITestServiceD>().PropertiesAutowired();
          IContainer container = containerBuilder.Build();
       //获取ITestServiceD接口在容器的注册的实例
          ITestServiceD testServiceD = container.Resolve<ITestServiceD>();
          testServiceD.Show();
  
      //类TestServiceD的内容
      public class TestServiceD : ITestServiceD
      {
          public string strin { get; set; }
          //获取TestServiceD实例的时候，_iTestServiceA，_iTestServiceB,_iTestServiceC这三个接口会全部注入成了实例对象，可以在下面Show方法中调用
          public ITestServiceA _iTestServiceA { get; set; }
          public ITestServiceB _iTestServiceB { get; set; } 
          public ITestServiceC _iTestServiceC { get; set; }
  
          public TestServiceD()
          {
              Console.WriteLine($"{this.GetType().Name}被构造。。。");
          }
          public void Show()
          {
              _iTestServiceA.Show();
              Console.WriteLine("D123456");
          }
      }
  ```

- 方法注入：Unity容器里面方法注入只需要对方法进行标记，获取抽象实例的时候会自动调用标记的防范，不用像Autofac主动调用。

  ```c#
               ContainerBuilder containerBuilder = new ContainerBuilder();
              containerBuilder.RegisterType<TestServiceA>().As<ITestServiceA>();
               //调用OnActivated(e => e.Instance.SetService(e.Context.Resolve<ITestServiceA>()))方法，调用SetService方法，获取到ITestServiceA接口的实例进行注入
              containerBuilder.RegisterType<TestServiceB>().OnActivated(e => e.Instance.SetService(e.Context.Resolve<ITestServiceA>())).As<ITestServiceB>();
              containerBuilder.RegisterType<TestServiceC>().As<ITestServiceC>();
              containerBuilder.RegisterType<TestServiceD>().As<ITestServiceD>().PropertiesAutowired();
              IContainer container = containerBuilder.Build();
          //获取ITestServiceB接口在容器中匹配的实例
              ITestServiceB testServiceb = container.Resolve<ITestServiceB>();
              testServiceb.Show();
   
     //TestServiceB类内容
     public class TestServiceB : ITestServiceB
      {
  
          public ITestServiceA _iTestServiceA = null;
           
          //通过方法进行注入_iTestServiceA的实例（感觉用实例直接调用这个方法也可以进行注入）
          public void SetService(ITestServiceA iTestServiceA)
          {
              _iTestServiceA = iTestServiceA;
          }
  
          public void Show()
          {
              _iTestServiceA.Show();
              Console.WriteLine($"This is TestServiceB B123456");
          }
      }
  ```

  

### Autofack生命周期

##### 瞬时生命周期（默认生命周期）

每次获取都是新的生命周期

```
                 //声明IOC容器Autofack的一个实例
                ContainerBuilder containerBuilder = new ContainerBuilder();
                //InstancePerDependency() 组件对TestServiceA实例获取的时候，每次都是一个新的实例
                containerBuilder.RegisterType<TestServiceA>().As<ITestServiceA>().InstancePerDependency();
                IContainer container = containerBuilder.Build();
                ITestServiceA testServiceA = container.Resolve<ITestServiceA>();//获取服务
                ITestServiceA testServiceA1 = container.Resolve<ITestServiceA>();//获取服务
                //对比对象的引用是否一致:false
                Console.WriteLine(object.ReferenceEquals(testServiceA, testServiceA1));
```

##### 单例生命周期

在当前生命周期内都是同一个实例 

```
                //声明IOC容器Autofack的一个实例
               ContainerBuilder containerBuilder = new ContainerBuilder();
               //SingleInstance 使用用单例模式获取TestServiceA的实例
                containerBuilder.RegisterType<TestServiceA>().As<ITestServiceA>().SingleInstance();
                  //使用上面注册的组件，创建一个容器属性container：保存抽象与实例映射关系的属性容器
                IContainer container = containerBuilder.Build();
                ITestServiceA testServiceA = container.Resolve<ITestServiceA>();//获取服务
                ITestServiceA testServiceA1 = container.Resolve<ITestServiceA>();//获取服务
                 //对比对象的引用是否一致:true
                Console.WriteLine(object.ReferenceEquals(testServiceA, testServiceA1));
```

##### 同生命周期范围内{}相同实例 

```c#
                //声明IOC容器Autofack的一个实例
               ContainerBuilder containerBuilder = new ContainerBuilder();
              //InstancePerLifetimeScope 在同一个生命周期内，获取到的实例都是相同的
                containerBuilder.RegisterType<TestServiceA>().As<ITestServiceA>().InstancePerLifetimeScope();
  //使用上面注册的组件，创建一个容器属性container：保存抽象与实例映射关系的属性容器
                IContainer container = containerBuilder.Build();
            
                ITestServiceA testServiceA15 = null;
                ITestServiceA testServiceA16 = null;
              //这里using{ container.BeginLifetimeScope()}是开启一个生命周期
                using (var scope1 = container.BeginLifetimeScope()) 
                {
                    ITestServiceA testServiceA11 = scope1.Resolve<ITestServiceA>();
                    ITestServiceA testServiceA12 = scope1.Resolve<ITestServiceA>();
                     //对比对象的引用是否一致:true
                    Console.WriteLine(object.ReferenceEquals(testServiceA11, testServiceA12));
                    testServiceA15 = testServiceA12;
                }

                using (var scope1 = container.BeginLifetimeScope())
                {
                    ITestServiceA testServiceA13 = scope1.Resolve<ITestServiceA>();
                    ITestServiceA testServiceA14 = scope1.Resolve<ITestServiceA>();
                    //对比对象的引用是否一致:true
                    Console.WriteLine(object.ReferenceEquals(testServiceA13, testServiceA14));
                    testServiceA16 = testServiceA14;
                }
               //对比对象的引用是否一致:false
                Console.WriteLine(object.ReferenceEquals(testServiceA15, testServiceA16));
```

##### 同一个key生命周期范围内{}相同实例

```c#
                //获取一个IOC容器Autofak的实例
               ContainerBuilder containerBuilder = new ContainerBuilder();
               //InstancePerMatchingLifetimeScope("Zhaoxi"); 通过Zhaoxi 这个key创建的生命周期
               //范围内的注入对象都是一个实例,即使生命周期内部开启新的生命周期，注入到对象也是同一个实例
                containerBuilder.RegisterType<TestServiceA>().As<ITestServiceA>().InstancePerMatchingLifetimeScope("Zhaoxi");
     //使用上面注册的组件，创建一个容器属性container：保存抽象与实例映射关系的属性容器
                IContainer container = containerBuilder.Build();

                ITestServiceA testServiceA15 = null;
                ITestServiceA testServiceA16 = null;
                using (var scope1 = container.BeginLifetimeScope("Zhaoxi"))
                {
                    ITestServiceA testServiceA11 = scope1.Resolve<ITestServiceA>();
                    using (var scope2 = scope1.BeginLifetimeScope())
                    {
                        ITestServiceA testServiceA12 = scope2.Resolve<ITestServiceA>();
                        //对比对象的引用是否一致:true
                        Console.WriteLine(object.ReferenceEquals(testServiceA11, testServiceA12));
                    }
                    testServiceA15 = testServiceA11;
                }

                using (var scope1 = container.BeginLifetimeScope("Zhaoxi"))
                {
                    ITestServiceA testServiceA13 = scope1.Resolve<ITestServiceA>();
                    using (var scope2 = scope1.BeginLifetimeScope())
                    {
                        ITestServiceA testServiceA14 = scope2.Resolve<ITestServiceA>();
                      //对比对象的引用是否一致:true
                        Console.WriteLine(object.ReferenceEquals(testServiceA13, testServiceA14));
                    }

                    testServiceA16 = testServiceA13;
                }
             //对比对象的引用是否一致:false
                Console.WriteLine(object.ReferenceEquals(testServiceA15, testServiceA16));
```

##### 每个网络请求一个实例

在每个web/http/api的请求时，共享一个组件实例

`下面的例子无法运行，因为不是通过web的请求进行访问`

```
                    //获取IOC容器Autofac的实例
                    ContainerBuilder containerBuilder = new ContainerBuilder();
                   //InstancePerRequest 在每个web/http/api的请求时，共享一个组件实例
                   containerBuilder.RegisterType<TestServiceA>().As<ITestServiceA>().InstancePerRequest();
           //使用上面注册的组件，创建一个容器属性container：保存抽象与实例映射关系的属性容器
                    IContainer container = containerBuilder.Build();
                    ITestServiceA testServiceA11 = container.Resolve<ITestServiceA>();
                    ITestServiceA testServiceA12 = container.Resolve<ITestServiceA>();
                    Console.WriteLine(object.ReferenceEquals(testServiceA11, testServiceA12));
```

### Autofak通过配置文件注入

**1.Nuget引入程序集：**

- Autofac 
- Autofac.Extensions.DependencyInjection 
- Autofac.Configuration 

**2.准备配置文件**

autofac.json配置文件内容

```c#
{
	"components": [
		{  
		    <!--实例带命名空间的类名，实例命名空间-->
			"type": "Zhaoxi.NET5Demo.Service.TestServiceUpdate,Zhaoxi.NET5Demo.Service",
			<!--接口带命名空间的名称，接口的命名空间。为什么是一个数组，我猜测可能是类可以继承多个接口的原因-->
			"services": [
				{
					"type": "Zhaoxi.NET5Demo.Interface.ITestServiceA,Zhaoxi.NET5Demo.Interface"
				}
			],
			"instanceScope": "single-instance", //生命周期类型，此处为单例生命周期
			"injectProperties": true // 是否支持属性注入，true代表支持，false代表不支持 
		},
		{
			"type": "Zhaoxi.NET5Demo.Service.TestServiceE,Zhaoxi.NET5Demo.Service",
			"services": [
				{
					"type": "Zhaoxi.NET5Demo.Interface.ITestServiceE,Zhaoxi.NET5Demo.Interface"
				}
			],
			"instanceScope": "single-instance", //生命周期
			"injectProperties": true // 属性注入 
		}
	]
}
```

**获取配置的代码：startup.ConfigureServices中添加**

```c#
           //获取IOC容器Autofac的实例
           ContainerBuilder containerBuilder = new ContainerBuilder();
            // 就可以在这里写入Autofac注入的各种 
            {
                //读取配置文件,把配置信息全部装载到ContainerBuilder
                IConfigurationBuilder config = new ConfigurationBuilder();
                IConfigurationSource autofacJsonConfigSource = new JsonConfigurationSource()
                {
                    Path = "CfgFile/autofac.json",
                    Optional = false,//boolean,默认就是false,可不写
                    ReloadOnChange = true,//同上
                };
                //添加全部的配置信息
                config.Add(autofacJsonConfigSource);
                ConfigurationModule module = new ConfigurationModule(config.Build());
                containerBuilder.RegisterModule(module);
            }
            //使用配置信息里面接口与类的映射关系，初始化容器container：保存抽象与实例映射关系的属性容器
            IContainer container = containerBuilder.Build();
            ITestServiceA testServiceA = container.Resolve<ITestServiceA>();
            //testServiceD里面的_iTestServiceA，_iTestServiceB,_iTestServiceC三个接口对象，会全部被属性进行注入
            ITestServiceD testServiceD = container.Resolve<ITestServiceD>();
            testServiceD.Show();
            
            
      //类TestServiceD的内容
    public class TestServiceD : ITestServiceD
    {
        public string strin { get; set; }
        //获取TestServiceD实例的时候，_iTestServiceA，_iTestServiceB,_iTestServiceC这三个接口会全部注入成了实例对象，可以在下面Show方法中调用
        public ITestServiceA _iTestServiceA { get; set; }
        public ITestServiceB _iTestServiceB { get; set; } 
        public ITestServiceC _iTestServiceC { get; set; }

        public TestServiceD()
        {
            Console.WriteLine($"{this.GetType().Name}被构造。。。");
        }
        public void Show()
        {
            _iTestServiceA.Show();
            Console.WriteLine("D123456");
        }
    }
```

### Autofac整合MVC（接管服务）

Autofac是一个第三方容器

1.,Program中进行指定 ，**使用Autofac工厂替换默认工厂**，替换了Autofac会接管IServiceCollection容器中所有的服务（NET5框架运行时很多服务都是通过IServiceCollection进行注入的，所以如果Autofac要接管系统的服务注册，它需要把IServiceCollection的原有功能也要满足）

```c#
public static IHostBuilder CreateHostBuilder(string[] args) => 
           Host.CreateDefaultBuilder(args)  
           .ConfigureWebHostDefaults(webBuilder => 
           { 
            webBuilder.UseStartup<Startup>(); 
           }) 
             //使用Autofac工厂替换默认工厂（这段代码加上去之前，原来是没有代码的，是新增，不是覆盖）
          .UseServiceProviderFactory(new AutofacServiceProviderFactory());
```

2.在Startup类增加ConfigureContainer 方法， **注册关系** 

```c#
//这个方法默认会被Autofac进行自动调用
public void ConfigureContainer(ContainerBuilder builder) 
       { 

          builder.RegisterType<TestServiceA>().As<ITestServiceA>(); 
          builder.RegisterType<TestServiceUpdate>().As<ITestServiceA>();  

      } 
```

3.通过控制器的构造函数注入，获取对象的实例进行使用

### Autofac注入控制器属性

通过构造函数注入控制器的对象实可以的，但是Autofac通过属性注入控制器默认是不可以，要经过配置才可以

1.startup.ConfigureServices中把控制器实例的权限移交给容器来创建

```
services.Replace(ServiceDescriptor.Transient<IControllerActivator, ServiceBasedControllerActivator>());
```

2. startup中注册所有控制器的关系+控制器实例化需要的所有组件

   ```c#
   //整个方法被Autofa自动调用,负责注册各种服务
   public void ConfigureContainer(ContainerBuilder containerBuilder)
   {
        // 注册所有控制器的关系+控制器实例化需要的所有组件
               Type[] controllersTypesInAssembly = typeof(Startup).Assembly.GetExportedTypes()
                .Where(type => typeof(ControllerBase).IsAssignableFrom(type)).ToArray();
         //根据配置的特性判断控制器的哪些属性需要注入（需要遍历所有的控制器与对应控制器的所有属性）,这个方法执行后，SixThController控制器中的接口iTestServiceAA属性将会注入到实例TestServiceA，属性可以正常使用
       containerBuilder.RegisterTypes(controllersTypesInAssembly).PropertiesAutowired(new CustomPropertySelector());
               
   }
      //控制器类
       public class SixThController : Controller
       {
           //添加这个特性代表对属性进行注入（特性是通过继承自定义的）
           [CustomPropertyAttribute]
           private ITestServiceA iTestServiceAA
           {
               get; set;
           }
           private ITestServiceB iTestServiceBB
           {
               get; set;
           }
           private ITestServiceC iTestServiceCC
           {
               get; set;
           }
           private ITestServiceD iTestServiceDD
           {
               get; set;
           }
           private ITestServiceE iTestServiceEE
           {
               get; set;
           }
       }
   
   //【新建两个类】
   
       //特性类：负责标记控制器中哪些属性需要注入
       [AttributeUsage(AttributeTargets.Property)]
       public class CustomPropertyAttribute:Attribute
       {
       }
       //判断类：判断控制器的类属性是否需要注入的方法类
       public class CustomPropertySelector : IPropertySelector
       {
            /// <summary>
            /// 根据入参的控制类中的属性，进行循环判断是否需要注入
            /// </summary>
            /// <param name="propertyInfo">控制器类所有的属性</param>
            /// <param name="instance">控制器类的实例</param>
            /// <returns>true 属性需要注入 false 属性不需要注入</returns>
           public bool InjectProperty(PropertyInfo propertyInfo, object instance)
           {
               //需要一个判断的维度；
               return propertyInfo.CustomAttributes.Any(it => it.AttributeType == typeof(CustomPropertyAttribute));
           }
       }
       
   
   
   
   ```

   

### Autofac单抽象多实现

#### 实现1：控制器构造函数注入声明对象的字段

1.一个接口的实现类，存在多个实现的实例，都注册了的情况下，最后面哪个类先注册，那么构造函数中注入的就是哪个类的实例，**后注册的会覆盖先注册的**

```c#
        //类TestServiceUpdate与类TestServiceA都实现了接口ITestServiceA，TestServiceA类后注册，那么会覆盖TestServiceUpdate的注册结果，Autofac容器默认会以TestServiceA的实例注入系统中
        containerBuilder.RegisterType<TestServiceUpdate>().As<ITestServiceA>();
        containerBuilder.RegisterType<TestServiceA>().As<ITestServiceA>();
```

2.当一个抽象存在多个实现的实例，并且都注册到了IOC容器当中,我们可以通过一个IEnumerable<抽象>,当做构造函数参数的，当调用控制器的时候，IOC容器就会把抽象存在的多个注册的实例，注入到IEnumerable<抽象>对象中进行调用（如果是同一个实现的实例，多次注册，那么也会存在多条记录）。

```c#
 private readonly IEnumerable<ITestServiceA> _iTestServiceAList;
 //如果同一个类实例向IOC容器中重复注册实现类，如果注册3次，那么iTestServiceAList中的数量也是3次，不会去除重复的实例
public SixThController(IEnumerable<ITestServiceA> iTestServiceAList)
{
    _iTestServiceAList = iTestServiceAList;
}
```

3.**【使用该方法进行单抽象多实现】**当控制器中根据需求需要注册一个抽象的多个实例对象进行调用，可以声明对象的实例字段， 在控制器的构造函数中，使用具体实现类型作为参数类型。

（这种情况的应用场景比如有一个日志接口，写入服务器日志类继承实现，写入本地日志类继承实现，这个时候需要动态获取类的实例，交易模块等类场景实现服务器日志类，往服务器写入日志，日常操作模块类记录到本地日志，实现本地日志类，所以**在一个应用程序对同接口有不同实例的注入要求**）

```
    
    
    
    private readonly ITestServiceA _iTestServiceA;

    private readonly IEnumerable<ITestServiceA> _iTestServiceAList;
    //实际类对象字段
    private readonly TestServiceA _iTestServiceAAA;
    //实际类对象字段
    private readonly TestServiceUpdate _iTestServiceUpdate;
    
    //通过构造函数注入实际的类字段
    public SixThController(
        ITestServiceA iTestServiceA, IEnumerable<ITestServiceA> iTestServiceAList, TestServiceA iTestServiceAAA, TestServiceUpdate iTestServiceUpdate)
    {
        _iTestServiceA = iTestServiceA;
        _iTestServiceAList = iTestServiceAList;
        _iTestServiceAAA = iTestServiceAAA;
        _iTestServiceUpdate = iTestServiceUpdate;
    }
```

4.通过IOC容器Autofac在Startup.ConfigureServer中注入实际的对象：**实例与抽象的方式注入，两种注入方式都需要，不能说有了实际的对象注入，就不对抽象进行注入了**

```c#
        // 注册单抽象多实例 :后注册后覆盖
        containerBuilder.RegisterType<TestServiceUpdate>().As<ITestServiceA>();
        containerBuilder.RegisterType<TestServiceA>().As<ITestServiceA>();
        #endregion

        //单抽象多实现 :通过系统方法实现
        containerBuilder.RegisterSource(new AnyConcreteTypeNotAlreadyRegisteredSource(t => t.IsAssignableTo<ITestServiceA>()));
       //单抽象多实现：通过封装系统方法到帮助类中实现（如果注入的单抽象多实现类过多，后面有改动直接改动帮助类即可）
        containerBuilder.RegisterModule(new AutofacModule());
       //单抽象多实现：泛型方法调用（与调用上面RegisterModule方法的作用一样，只是多一种调用方式）
        containerBuilder.RegisterModule<AutofacModule>();

    //声明单抽象多实现的帮助类：封装好了ITestServiceA接口的多实现（需要有写死接口名的入参，后期有改动可能不是很方便）
    public class AutofacModule: Module
    {
        protected override void Load(ContainerBuilder builder)
        {
            builder.RegisterSource(new AnyConcreteTypeNotAlreadyRegisteredSource(t => t.IsAssignableTo<ITestServiceA>()));
        }
    }
  
```

#### 实现2：注入容器IComponentContext组件的上下文

##### 1.通过构造函数注入

1.在Startup.ConfigureContainer注册的时候，指定一个标识key，然后可以在实例中进行读取（区分大小写）

```c#
//标志Key:TestServiceA  （可以自定义，只是此处与类名称一致）
containerBuilder.RegisterType<TestServiceA>().Named<ITestServiceA>("TestServiceA"); 
//标志Key:TestServiceUpdate （可以自定义，只是此处与类名称一致）
containerBuilder.RegisterType<TestServiceUpdate>().Named<ITestServiceA>("TestServiceUpdate"); 
```

2.在控制器中获取的时候获取一个Autofac的上下文IComponentContext，通过注入上下文的实例，然后调用方法＋标识，得到单个抽象的不同的实现的实例；

```c#
private readonly IComponentContext _ComponentContext = null; 

       public SixThController(IComponentContext componentContext) 

       { 

           _ComponentContext = componentContext; 

       } 

      public IActionResult Index() 

       { 
//通过标志key TestServiceA 获取TestServiceA类的实例
ITestServiceA testServiceA = _ComponentContext.ResolveNamed<ITestServiceA> 

("TestServiceA"); 
//通过标志key TestServiceUpdate 获取TestServiceUpdate类的实例
ITestServiceA testServiceUpdate = _ComponentContext.ResolveNamed<ITestServiceA> 

("TestServiceUpdate"); 

           iTestServiceAA.Show(); 

           return View(); 

       }


```

##### 2.通过属性注入

1.第一步与上面通过构造函数注入一样。

2.控制器代码中加入IOC容器上下文注入代码：

```
    //通过属性注入的对象
    [CustomPropertyAttribute]
    private IComponentContext _ComponentContextProp { get; set; }
    
    public IActionResult Index()
    { 
            ITestServiceA testServiceA = _ComponentContextProp.ResolveNamed<ITestServiceA>("TestServiceA");
            ITestServiceA testServiceUpdate = _ComponentContextProp.ResolveNamed<ITestServiceA>("TestServiceUpdate");
 
            testServiceA.Show();
            testServiceUpdate.Show();
            return View();
    }
```

### 抽象支持AOP

AOP: 不用更改代码，动态的改变类方法的行为，可以在方法执行前新增逻辑，也可以在方法结束后新增逻辑。比如我们可以在方法执行前增加权限判断，判断是否满足调用条件，也可以在方法执行完成后加入日志记录

#### 1.通过接口实现AOP

   1.Nuget引入

```
Castle（咳kei(四声)嗽）.Core;
Autofac.Extras.DynamicProxy;
```

  2.单独新建一个项目，项目中新建AOP切入类（**如果放到UI层建立这个类，会产生循环引用问题。因为接口层需要引用UI层的这个AOP切入类，UI层又需要引用接口层，所以会产生循环引用**），继承拦截器接口：IInterceptor（允许类拦截动态代理的权限）

```
 //拦截方法的执行，把方法执行的控制权放到我这里来
public class CustomAutofacAop : IInterceptor
{
    public void Intercept(IInvocation invocation)
    {
        {
            Console.WriteLine("加入方法执行前逻辑");
        }
        invocation.Proceed();//执行这句话就是去执行具体的实例的这个方法

        {
            Console.WriteLine("加入方法执行后逻辑");
        }
    }
}
```

3.标记接口，加入类的AOP拦截权限，只要获取到继承了这个接口的对象，是**通过抽象接口获取到类的实例**，都会**先进入类CustomAutofacAop**进行动态拦截

```
//加入这个特性，让AOP能够在当前接口生效
[Intercept(typeof(CustomAutofacAop))] 
public interface ITestServiceA
{
    void Show();
}
```

4.第一步startup.ConfigureContainer中加入注册类CustomAutofacAop 让它具有AOP权限的代码，第二步加入EnableClassInterceptors()方法代表接口ITestServiceA需要开启AOP的权限

```c#
 //向容器加入一个注册源  
containerBuilder.RegisterSource(new AnyConcreteTypeNotAlreadyRegisteredSource(t => t.IsAssignableTo<ITestServiceA>())); 
  //1.开启注册类CustomAutofacAop AOP权限的代码 
 containerBuilder.RegisterType(typeof(CustomAutofacAop));

 //2.EnableInterfaceInterceptors 通过接口进行拦截类的具体实现
 containerBuilder.RegisterType<TestServiceUpdate>().As<ITestServiceA>().EnableInterfaceInterceptors();  
    containerBuilder.RegisterType<TestServiceA>().As<ITestServiceA>().EnableInterfaceInterceptors(); ; 


```



#### 2通过类实现AOP





## ActionFilter：过滤器

**ActionFilter：**过滤器做权限认证，日志记录，异常捕获，缓存

下图为过滤器的执行顺序：

1.Resource Filter开始

2.ActionFilter开始

3.开始执行方法

4.ActionFilter结束

  T.开始Exception Filter（**存在没有捕获的异常才会执行这步**）

5.开始Result Filter

6.开始Result Filter后

7.执行Resource Filter结束

 

![1635345704593](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1635345704593.png)



### **ExceptionFilter** 异常过滤器

1.自定义一个CustomExceptionFilterAttribute ,实现IExceptionFilter接口， 

2.实现方法，先判断，异常是否被处理过，如果没有被处理过，就处理； 

3.分情况处理：1.如果是ajax请求，就返回JosnResult，如果不是Ajax请求，就返回错误页面 

```
 public class CustomExceptionFilterAttribute : Attribute, IExceptionFilter 

   { 

​       private IModelMetadataProvider _modelMetadataProvider =null; 

​       public CustomExceptionFilterAttribute(IModelMetadataProvider modelMetadataProvider) 

​       { 

​           _modelMetadataProvider = modelMetadataProvider; 

​       } 

​       /// <summary> 

​       /// 当异常发生的时候触发到这儿来 

​       /// </summary> 

​       /// <param name="context"></param> 

​       public void OnException(ExceptionContext context) 

​       { 

​           if (!context.ExceptionHandled) //异常是否被处理过 

​           { 

​               //在这里处理 如果是Ajax请求===返回Json 

​               if (this.IsAjaxRequest(context.HttpContext.Request))//header看看是不是 

XMLHttpRequest 

​               { 

​                   context.Result = new JsonResult(new 

​                   { 

​                       Result = false, 

​                       Msg = context.Exception.Message 

​                   });//中断式---请求到这里结束了，不再继续Action 

​               } 

​               else 

​               { 

​                   //跳转到异常页面 

​                   var result = new ViewResult { ViewName = "~/Views/Shared/Error.cshtml" }; 

​                   result.ViewData = new ViewDataDictionary(_modelMetadataProvider,  

context.ModelState); 

​                   result.ViewData.Add("Exception", context.Exception); 

​                   context.Result = result; //断路器---只要对Result赋值--就不继续往后了； 

​               } 

​               context.ExceptionHandled = true; 

​           }  

​       } 

​       private bool IsAjaxRequest(HttpRequest request) 

​       { 

​           string header = request.Headers["X-Requested-With"]; 

​           return "XMLHttpRequest".Equals(header); 

​       } 

   }
```

4.全局注册，在Starup中的ConfigureServices注册 

```

```

**ExceptionFilter能捕捉到哪些异常** 

1.控制器实例化异常 ----T 

2.异常发生在Try-cache中 ---F 

3.在视图中发生异常 ----F 

4.Service层发生异常 ---T 

5.在action中发生异常 ---T 

6.请求错误路径异常 ---可以使用中间件来支持，只要不是200的状态，就都可以处理； 

### ActionFilter基础过滤

 1.新建特性类，继承Attribute特性（**Asp.net core 可以直接集成ActionFilterAttribute特性实现里面的方法**），然后继承命名空间为Microsoft.AspNetCore.Mvc.Filters的IActionFilter  

```
//新建特性类，继承Attribute特性，然后继承命名空间为Microsoft.AspNetCore.Mvc.Filters的IActionFilter   
public class QrActionFilter : Attribute, IActionFilter
{
    public void OnActionExecuting(ActionExecutingContext context)
    {
        Console.WriteLine("方法执行前日志记录");
    }
    public void OnActionExecuted(ActionExecutedContext context)
    {
        Console.WriteLine("方法执行前日志记录");
    }
}
```

 2.将特性标记至控制器中的方法

```c#
    //标记控制器Action特性
    [QrActionFilter]
    public IActionResult MyView()
    {
        ViewBag.Name = "王锐";
        return View();
    }

```

3. Startup.ConfigureService全局注册

   //Startup.ConfigureService全局注册
   services.AddMvc(option =>
   {
       option.Filters.Add<QrActionFilter>(); //全局注册：
   });

### **ResultFilter视图渲染过滤器** 

双语言系统，其实就需要两个视图；要根据语言的不同，来选择不同的视图来渲染； 

因为在渲染视图之前，会进入到OnResultExecuting方法，就可以在这个方法中确定究竟使用哪一个视 

图文件； 

1.自定义一个类，继承Attribute,实现IResultFilter接口，实现方法 

2.标记在Action方法头上 

3.执行顺序：视图执行前，渲染视图，视图执行后 

### 匿名Filter避开Filter检查

如果全局注册，Filter会生效于所有控制器中的Acion，如果有部分Action我希望它不生效？要怎么办？

**使用匿名避开Filter的检查**

如何支持自定义的Filter,匿名 

1.自定义一个特性； 

```
public class NoActionFilter:Attribute
{
 
}
```

2.标记到方法上

```
 [NoActionFilter] //匿名
 public IActionResult Index()
 {
     return View();
 }
```

3.在需要匿名的Filter内部，检查是否需要匿名（检查是否标记的有匿名特性）,如果有就直接避开，方法开始前避开，如果需要方法开始后避开还是需要把这部分逻辑放到那里面

```
    //基础过滤器,方法开始前
   public void OnActionExecuting(ActionExecutingContext context)
    {
        if (context.ActionDescriptor.EndpointMetadata.Any(item=>item.GetType()==typeof(NoActionFilter))) //如果标记的有特殊的记号，就避开检查；
        {
            return;
        }
        
    }
```

### **Authorization Filters鉴权授权** 

为了拦截一些操作； 

传统的授权方式；session/Cookies来完成； 

1.在请求某个Action之前去做校验，验证当前操作者是否登录过，登录过就有权限 

2.如果没有权限就跳转到登录页中去 

AOP--Filter； ActionFilter: 

传统的登陆，需要匿名； 

中间件做：**权限管理，角色管理，中间件的引用

### Resource  Fileters 资源过滤器

资源过滤器 ，这一步用来判断用户的请求是否可以命中缓存，是否是从官网进来的，还是从第三方进来的（防盗链检查）

# 阿莱克斯ASP.NET CORE MVC学习

## 5章节：60分钟  --已完成

### 数据模型与仓库模式 

通过仓库模式获取数据，还有其他几种从数据库获取数据的方式比如ODBC,ADO.NET,ORM，数据仓库方式获取数据。通常我使用的sqlconnection,sqlcommand这个对象使用的是微软提供的（访问对象数据的组件）ADO.NET的方式提供。



### Demo创建model和数据仓库

**这章做了什么?**

新建兰州拉面类，新建接口（数据仓库），实现接口（获取所有兰州拉面信息，根据ID获取兰州拉面信息），创建兰州拉面视图进行展示数据



在Model文件夹新建class，命名为lazhoulamian（兰州拉面）

根据拉面ID获取兰州拉面，lamda表达式的返回值({左边是参数} =>{右边是表达式})

**为什么要声明一个Ilazhoulamian接口？**

如果数据源从目前的sqlserver更改成了oracle那么依靠这个接口，我们底层的代码也可以很方便的进行变动，这个叫**依赖注入**？

**IActionResult**:控制器中的action方法中的返回值**IActionResult**可以根据返回的类型（实体类，list,string,json,view）动态的进行转换，兼容所有的返回值

### 视图模型viewmodel

**这章做了什么？**

 视图里面还有模型？它与Model中的模型有什么关系？

**MVVM架构**：Model-View-View Model,可以简单的看做是MVC的扩展，VUE里面也应用了MVVM架构模式



### View Model

DTO实体类扩展的模型与视图模型的区别在于，DTO模型是针对于实体表进行扩展属性，而视图模型需要这怒地界面的展示信息将信息打包封装到视图模型中，方便页面进行提取展示。对应需要在控制器中对Model实体的数据进行封装到视图模型中进行返回到界面。

![1633840521771](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633840521771.png)

#### 依赖注入：面向接口编程

 依赖注入这种技巧让HomeController控制器与数据访问对象没有产生依赖，它们两个类之间没有接触到，所以是低耦合。使用接口抽象的行为进行操作，运行过程中依靠具体的类进行实现。控制器引用接口对象行为的时候会告诉类访问中有这个行为，至于到底怎么实现这个行为要等到运行的时候动态调用（进行注入）才知道，这种方式灵活性很强。

![1633841579821](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633841579821.png)

比如控制器里面有写入日志的方法，写入日志方法是一个抽象的实现，放到日志接口中，至于是往本地写入还是往服务器写入，还是写入到数据库，看用户的需要（根据配置）进行实现。

![1633857366047](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633857366047.png)

上面的截图中**有个兼容A与B注入的问题**，比如拿写入日志这个场景，电商中的交易场景比较私密只能往服务器上面写入日志，普通场景往本地写入日志就可以，如果日志的编程接口，注入的对象是往服务器写入日志的对象，那么我们**怎么同时兼容写入日志到服务器与本地？****



其实注入讲的太官方了，就是用的时候才告诉你我用的是哪个类的实例（这个时候控制器类依赖的是接口，而不是具体的实现），我们一般可以通过构造函数进行注入。其实注入讲的就是我现在要用这个声明的实例了哈，不管它在我们内部是声明为类（类也可以用多态）还是接口，我现在要用它的真实对象,可以通过构造函数入参，或者单独声明一个方法进行入参，参数可以是类或者是接口都可以





![1633841794553](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633841794553.png)

还有一点就是如果是写入日志的功能，有一种方法是写入到数据库的日志，还有是写入到服务器文件的日志， 现在写入日志是写入的数据库的日志，现在公司要修改成写入到服务器文件，打算怎么做？

![1633842637004](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633842637004.png)

这种方法的实现都是在类的内部直接声明，内部类高度依赖具体的实现类，像这种日志打印如果有100个类引用，如果后面有改动的话，我们在100个类中进行**全局替换**类的引用吗？这种全局替换的解决方案是比较**初级，风险性很高**的解决方案。



如果这个类是依赖注入的话（依靠外部的调用来确定实际引用的方法实现，把对象的创建交给类的外部），那么我们控制好注入的一处地方就可以。

#### 为什么要使用依赖注入

- - **方便改动**：传统的代码，每个对象负责管理与自己需要依赖的对象，导致如果需要切换依赖对象的实现类时，需要修改多处地方，如果是依靠注入实现实例化对象，我们只需要更改注入的位置就可以。
  - **松耦合：**依赖注入把对象的创造交给外部去管理,很好的解决了代码**紧耦合**的问题，是一种让代码实现**松耦合**的机制。
  - **灵活：**松耦合让代码更具灵活性，能更好地**应对需求变动**。

### 数据验证 

使用实体类的特性对界面数据进行验证,验证通过后保存数据，验证不通过提示不通过的原因

1. **特性验证：**验证数据是否为空，姓名字符长度，邮箱是否合法，支持使用正则表达式

2. **界面引用：**@addTagHelper![1633856356956](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633856356956.png)

   3.**HTML引用**：<span asp-validation-for="Name" class="text-danger"></span>

## 6章节：30分钟 --已完成

### 什么是Entity Framework 

- EF模型有事务验证，并发，数据迁移，数据增删改查，一级缓存（当向数据库查询一次后，第二次查询不会向数据库查询，会查询缓存里面的数据）
- 使用Linq操作实体对象进行操作数据库

### 添加Entity Framework Core

- 引用数据库链接组件后，新建面条实体类，新建content数据库访问类，构造IOC注入构造函数，新建json数据库访问配置，在startup.cs文件中构造IOC注入构造函数，并且在ConfigureServices方法中进行注入数据库的配置
- Nuget（妞get）中进行搜索entityframeworkcore,entityframeworkcore.sqlserver组件：

![1633856920221](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633856920221.png)

![1633856988986](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633856988986.png)



### 创建以及初始化数据库

- 使用entity framework core tools 组件，通过命令行的方式创建数据库,然后通过EF模型添加数据到数据库，并且进行读取展示
- NuGet（妞get）中引用entity framework core tools 组件：

![1633858325720](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633858325720.png)

### 修改数据模型

在实体类新增一个bool类型的属性后，通过entity framework core tools 组件，使用命令行的方式对应更新数据库的表，新增这个属性到表中：

![1633859336999](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633859336999.png)

## 7章节：30分钟

### Asp.net MVC 视图 详解

MVC默认使用cshtml页面，也可以使用aspx，html页面

### 添加视图与layout

通用模板页面有两种：一种需要在页面主动进行引用模板页面，如果引用的页面有几十页的话是比较繁琐。还有一种是直接作用到每页的模板页面

### 网站美化添加style

- 使用**lipman（来普man-三声）**管理页面的**bootstrap（布特丝拽普）**前端文件
- 给首页展示的兰州拉面信息添加样式
- 新增startup中显示兰州拉面的图片与展示样式的中间件

### 添加网站导航，完善细节

- 点击兰州拉面信息进行兰州拉面详情，展示兰州拉面信息与评论信息
- 设置模板页面展示顶部的标签信息

## 8章节：30分钟

### Asp.net Core identity

**identity（）**：身份认证框架，授权，身份认证，管理用户权限，给部分用户设置指定的权限，管理网站的角色与权限，可以使用框架迁移数据库，可以在不同的数据库中进行使用，可以与第三方登录权限（QQ,微信）进行对接关联到我们系统，这个框架也可以自动创建表，比如下面通过命令行可以直接创建相关的用户权限表

**四个场景：**注册，登录，注销，访问

### 添加Asp.net Core identity

Nuget（妞get）：引用identity.entity

![1633882531099](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633882531099.png)

集成由content更新为identity.entity框架的继承

![1633882772937](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633882772937.png)

![1633882818306](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633882818306.png)

### 用户认证

引用Identity框架后，根据框架内的模板新增登录，注册，注销三个页面

![1633883053253](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633883053253.png)

添加上面的页面后，在startup页面添加razor页面的中间件，可以正常访问这个页面

![1633883464540](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633883464540.png)





![](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633883431666.png)





### 用户授权

在startup页面注册identity框架权限控制的中间件

然后在需要控制权限的页面引用identity框架的命名空间：

![1633883731291](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633883731291.png)

在类的特性中新增：[Authorize]

![1633883719579](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633883719579.png)

### 课程总结

https://github.com/Yaduo/LanzhouBeefNoodles

![1633883919414](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633883919414.png)

![1633884351218](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1633884351218.png)

Mediator：中介者模式

Decorator：装饰器模式

Strategy：策略模式

Observer：观察者模式

Builder：建造者模式

Singleton：单例模式

facade：外观模式

factory：工厂模式

