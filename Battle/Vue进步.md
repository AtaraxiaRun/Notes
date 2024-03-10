 工具使用

## Vue

**Vue的注释**：

```
/**
 * @param {*} list 现有标签导航列表
 * @param {*} newRoute 新添加的路由原信息对象
 * @description 如果该newRoute已经存在则不再添加
 */
export const getNewTagList = (list, newRoute) => {
    const { name, path, meta } = newRoute
    let newList = [...list]
    if (newList.findIndex(item => item.name === name) >= 0) return newList
    else newList.push({ name, path, meta })
    return newList
}
```

mapVue的作用 

- 响应式刷新界面的变量：比如界面绑定到Vue的变量，如果变量的值发生了改变，那么对应Html也会进行刷新改变。
- 可以模块化页面，
- 组件化页面

可以通过Vue把后端传递到界面的值（**可以是单个值，也可以是实体对象**）封装到.js文件中进行保存，程序界面通过script标签进行引用就可以

**以下内容见菜鸟教程** https://www.runoob.com/vue2/vue-template-syntax.html

### 要点内容

#### Vue基础结构

1.开发环境基础架构：

```c#
export default
{
  el: '#app-4', //绑定到具体哪个组件（vue实例就是一个规范，要把这种规范依附到哪个容器中，一定要依附到具体的容器中才可以起到作用），这个容器一般是一个父的vue实例页面，也可以是一个<div id="app-4" />,span，一般layout路由中会绑定到你绑定的路由页面，然后路由页面就是父的实例，父的里面又可以容纳很多子的组件
  name: 'Lister',
    //指令：点击右键复制文字
  directives:{},  
   //定义自定义组件的属性,当组件在另一个页面被引用的时候，那么引用的组件为父组件，这个时候这里面的属性就是父组件往子组件里面传属性参数
  props: {
    id: {
      type: Number, // 规定当前属性接受的数据类型：`String`、`Number`、`Boolean`、`Array`、`Object`、`Date`、`Function`
      required: true,   //required代表是必填的
        default: 0, //默认值
        validator: function (value) {
        return value >= 0   //验证函数
      }
    },
    publishStatus: {
      type: Number,
      required: true
    }
  },
  //监听值的变化，就好像触发器一样，维持某一组业务数据的一致性
    //
  watch:{
      //监听属性里面的属性（注意用单引号括起来），（newVlue, oldValue），可以只放一个newvalue新的入参，如果不需要旧参的话	
     'formData.platformName'(newVlue, oldValue) {
      console.log("newVlue",newVlue)  //新的值
      console.log("oldValue",oldValue)  //旧的值
      this.displaySelect = this.selectPlatformName.indexOf(newVlue) > -1
    },
     //监听单个属性
     title(newVlue, oldValue) {
      console.log("newVlue",newVlue)  //新的值
      console.log("oldValue",oldValue)  //旧的值
     
    }
  }
  //组件
  components:{},
  //data是一个方法，data接上括号，然后加上大括号
 data () {
    //return { 左边的这个括号要在上面，不能在下面不然就会异常 }
    return {
      message: 1,
      messaget: 2
    }
  },
   //过滤器可以当作一个方法被页面上的变量所调用，调用方式： <div>{{name | nameFilter}}</div>，比如name是王锐，调用这个过滤器后，结果就是：我是王锐，并且可以声明多个过滤器进行叠加, <div>{{name | nameFilter |nameFilter2}}</div>
filters: {
  nameFilter: function (value) {
    return value="我是"+value
  }
}
      
  },
  //计算属性(基于它们依赖的响应式变量进行缓存，如果变量的值没有变动，那么直接返回原来计算好的值，如果在计算逻辑比较复杂，数据量比较大的情况下非常有用)
    //
  computed:{
    reversedMessage: function () {
            // getter 如果firstName或者lastName有变化，那么reversedMessage的变量会自动更新
            get: function () {
              return this.firstName + ' ' + this.lastName
            },
            // setter 根据reversedMessage的值回写,firstName与lastName的值
            set: function (newValue) {
              var names = newValue.split(' ')
              this.firstName = names[0]
              this.lastName = names[names.length - 1]
            }
    }
  },
    //在实例初始化之后，创建之前执行
    beforeCreated(){
        
    },
  //界面渲染完前调用，即通常初始化某些属性值，然后再渲染成视图
  created () {
      console.log("初始化页面中...")
  },
  //在挂载开始之前调用
  beforeMounted(){
      console.log("页面开始挂在")
  }
  //界面渲染完后执行，通常是初始化页面完成后，再对html的dom节点进行一些需要的操作
  mounted(){
      console.log("页面渲染完毕")
  },
    
   //方法
  methods:{}
   想看更详细的步骤下面有 
}
```

![img](Vue进步.assets/201833094430875.png)

```

```

2.页面通过绑定的对象进行赋值，解构：html，引用script在线vue，<scrpit>new Vue对象（Vue对象里面有Vue的名字，Vue绑定的属性，Vue绑定的方法，而且我发现方法里面可以直接引用Vue绑定的属性）

#### 钩子函数（结合上面图片看）

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



#### VsCode搭建Vue项目

如果执行项目出现：vue-cli-service 不是内部或外部命名。。。

执行:**npm install** 

解释：`npm` (Node Package Manager) 是 Node.js 的包管理器，就像 .NET 中的 `nuget` 一样。它允许开发者安装和管理他们项目所依赖的代码库。直接运行 `npm install`,npm 会查找当前目录下的 `package.json` 文件，然后安装文件中列出的所有依赖。如果你在 `npm install` 后面指定了一个或多个包的名称，比如 `npm install express`，npm 就会从 npm 的在线仓库下载并安装指定的包，同时也会安装该包的所有依赖，

VsCode中新建终端，执行以下命令

- 安装vue-cli：安装脚手架到系统中，只需要运行一次，后续可以重复使用

```
npm install -g vue-cli
```

- 安装webpack：打包压缩vue项目，打包优化后的项目浏览器访问时，性能更快

```
npm install -g webpack 
```

- 创建保存项目的文件夹，cd到文件夹：

```
cd D:\VUE\01vue_vscode
```

- **创建项目** ：

```c#
vue init webpack myvue   //可能会提示版本不对的警告信息，但是如果末尾有看到让你输入项目名称，就不用管，一直回车
```

全部**按回车**： 创建时出现一大堆的配置：路由，作者，权限认证，npm依赖安装的配置 。

- 安装/使用Element

```c#
//先cd 到项目目录
cd D:\VUE\01vue_vscode
//终端执行这一句：安装element-ui，   i 是install缩写, -s 是添加到package.json包管理器中的安装插件dependencies列表，就好像vs中的依赖插件
npm i element-ui -S
//在根目录中main.js 中写入以下内容：
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(ElementUI)

//使用element ui：src/components/HelloWorld.vue 中添加下面的代码
<el-button type="primary" :loading="true">加载中</el-button>
```

- 执行项目

```c#
//先cd 到项目目录
cd D:\VUE\01vue_vscode
//清除之前的npm包缓存
npm cache clean --force
//运行前先升级基础包的版本
npm install
//安装vue
npm install --save vue
//安装可能缺失的npm包
npm install ajv@^5.0.0 browserslist@>= 4.21.0
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

8. 使用有道截图翻译Vs Code的提示英文的时候有时候可以正常翻译（错误提示可以翻译，下载npm的时候也可以），有时候又不能，可以看一看

9. **字符串为单引号：**export default {name: **'**TestCom**'** }    

10. **设置快键键清空终端的日志：**先打开设置 - > keyboard Shortcuts -> 搜索 "workbench.action.terminal.clear" -> 目前设置：**ctrl+L ctrl＋k**

11. Class的命名用横杠：

    ```
    .sp-page-table .el-pagination {
      margin-top: 3px;
    }
    ```

    

12. **方法的结束不需要分号；  **： 

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

- 使用Nginx在Linux部署项目，Nginx部署看下Linux/Docker部署Nginx

```mysql
# 部署Nginx
docker run -d --name nginx \
-p 8088:80 \
-v $(pwd)/nginx/default.conf:/etc/nginx/conf.d/default.conf \
-v $(pwd)/nginx/logs:/var/log/nginx \
-v $(pwd)/dist:/usr/share/nginx/html \
nginx
# 配置是核心，过程中我想更改location / 更改为 /web/ ，但是总是不行，上网看了四五篇攻略也不行
cat > nginx/default.conf << EOL

server {
  listen 80;
  listen  [::]:80;
  server_name localhost; 

 location / {
    root /usr/share/nginx/html;  
    try_files $uri $uri/ /index.html;
}
}

EOL
#访问地址，成功
http://192.168.226.128:8088

#备用脚本 刷新Nginx配置
docker exec -it cbdf807df4b1 nginx -s reload
  
```

**原文：**VSCode搭建Vue项目 https://www.cnblogs.com/zyskr/p/10609288.html

Nginx部署参考地址：

https://segmentfault.com/a/1190000019960079 [Vue项目的部署简述](https://segmentfault.com/a/1190000019960079)

https://juejin.cn/post/6844904098622341133 前端vue打包后部署方式





#### 【快速新建页面：看这个】

**无注释版本**

```html
<template>

  <el-dialog
    :visible.sync="show"
    :close-on-click-modal="false"
    :close-on-press-escape="false"
    custom-class="add-dialog"
    top="5vh"
    width="50%"
    :title="title"
    append-to-body
    lock-scroll
  >

    <el-form ref="form" :model="formData" label-width:100px  size="small">
      <el-form-item prop="shopName">
        <PlatformShopSelect v-model="filter.shopId" platform="walmart" />
      </el-form-item>
    </el-form>
    <UploadFile
      ref="uploadinsertfiles"
      :template-down-url="templateInsertDownUrl"
      :upload-url="uploadInsertUrl"
      :auto-upload="false"
      @reloadTable="reloadTable"
    />
    <div
      slot="footer"
      class="dialog-footer"
    >
      <el-button size="small" type="danger" @click="excelImport()">确 定 </el-button>

      <el-button size="small" type="danger" @click="show=!show">关 闭</el-button>
    </div>
  </el-dialog>

</template>
<script>
import { importBrandList } from '@/api/walmart/walmartmonitor'
import PlatformShopSelect from '@/components/Common/PlatformShopSelect'

export default {
  name: 'AddBrand',
  components: {
    PlatformShopSelect
  },
  data() {
    return {
      templateInsertDownUrl: 'https://qiniufile.yuetenginc.com/importTemplate20220831162017.csv', // 实例模板下载网络路径（需要先上传到七牛云服务器）
      uploadInsertUrl: importBrandList.url, // 批量新增
      title: '添加',
      show: false,
      formData:
      {
        shopName: null,
        shopId: null
      },
      tableData: {}
    }
  },
  methods: {
    excelImport() {
      this.$refs.uploadinsertfiles.open()
    },
    open() {
      this.show = true
    }
  }
}
</script>


 
```

**有注释版本**

```c#
// 有注释版本
<template>
  <div>
     //label-width是里面el-form-item里面的label的宽度（越大，就距离左边的间距越大）
    <el-form ref="form" :model="formData" label-width:80px>
     //label是左边的文字,prop是需要做表单验证的文字,
      <el-form-item label="活动名称"  prop="name" >
        <el-input v-model="formData.name" />
      </el-form-item>
        <el-form-item label="内部品类">
        // v-model是下拉绑定的值,placeholder是提示的文字，clearable是下拉框右边清除值的选项
        <el-select
          v-model="formData.pl"
          placeholder="请选择内部品类"
          style="width:240px"
          clearable
        >
          <el-option
            label="一级品类"
            value="yjpl"
          />
        </el-select>
      </el-form-item>
    </el-form>
    <el-table
      :data="tableData"
      style="width:100%"
    >
      <el-table-column
        prop="data"
        label="日期"
        width="180"
      />

    </el-table>
  </div>
</template>
<script>
export default {
  name: 'CagegoryAdd',
  components: {},
  data() {
    return {
      formData: {},
      tableData: {}
    }
  },
  methods: {

  }
}
</script>

```



#### 界面完成检查清单

- 利用调试工具检查值的清空情况（调试工具中的data再重新加载后值是否被覆盖）

### template slot-scope="scope"插槽的作用

在 Vue.js 中，当您在模板中使用 `slot-scope` 属性定义插槽作用域时，`scope` 是一个对象，用于暴露父组件传递给子组件插槽内容的数据。这个对象包含了可以在插槽内部使用的数据。这里面的父组件就是el-table，子组件就是被包裹在里面的组件，比如el-table-column，el-select， 为什么用template包裹，不用div包裹？有两个原因：1.template不会渲染到界面上，不会占用界面的位置 ,div确会 2.div没有 slot-scope对象，无法定义定义作用域插槽，让子组件放到父组件的数据

```c#
    <el-table
        :data="formData.platformShopList"
        >
          <el-table-column
            prop="siteList"
            label="站点"
          >
            <template slot-scope="scope">
              <el-select v-model="scope.row.siteList" style="width:290px;" placeholder="全部" multiple :platform="scope.row.platform" />
            </template>
          </el-table-column>
</el-table>

```



### 调用全局类Utils中的方法,Main.js扩展Vue原型

`this.utils`似乎是在Vue组件中直接使用的一个对象，这可能是因为它被添加到了Vue实例的原型链上，使得它在组件中可以全局访问。

通常情况下，为了在Vue组件中实现全局访问，可以通过Vue的插件机制或原型链来定义和扩展属性。如果你在项目中看到`this.utils`可以直接使用，那么很可能是在某个地方对Vue的原型进行了扩展。

以下是一种可能的实现方式，将`this.utils`添加到Vue的原型中，使得它在所有Vue组件中都可以访问：

1. 首先，在适当的地方，比如一个单独的文件中（例如`utils.js`），定义你的`deepClone`方法和其他工具方法。

```c#
// utils.js

function deepClone(obj) {
  // 实现深拷贝逻辑
}

// 其他工具方法...

export default {
  deepClone,
  // 其他方法...
}

```

在你的Vue项目中，将这个`utils.js`文件导入，并将它添加到Vue的原型中，这样就可以在组件中使用`this.utils`来访问这些方法。

```c#
// main.js 或者类似的入口文件

import Vue from 'vue';
import utils from '@/path/to/utils.js';  // 根据你的文件路径进行调整

Vue.prototype.utils = utils;

new Vue({
  // ...
});

```

![1691997120922](Vue进步.assets/1691997120922.png)

![1691997151094](Vue进步.assets/1691997151094.png)

![1691997187995](Vue进步.assets/1691997187995.png)

### Element UI中组件属性Attributes，Method,Events,Slots的使用

> Slots：插槽，插槽就好像一个占位符一样，我们可以使用插槽，将我们想插入的元素放到指定的位置中,下面用el-select的empty插槽来说明当可选项为空时，用户将看到 "No options available" 的提示文本，这个内容是通过 `empty` 插槽自定义的。

```html
<template>
  <div>
    <el-select v-model="selectedOption" placeholder="Select an option">
      <!-- 使用 empty 插槽来自定义空状态时的显示内容 ,当可选项为空时，用户将看到 "No options available" 的提示文本，这个内容是通过 empty 插槽自定义的。-->
      <template #empty>
        <span style="color: #999;">No options available</span>
      </template>
      <el-option v-for="option in options" :key="option.value" :label="option.label" :value="option.value" />
    </el-select>
    <p>Selected Option: {{ selectedOption }}</p>
  </div>
</template>

<script>
export default {
  data() {
    return {
      selectedOption: '',
      options: [
        { value: 'option1', label: 'Option 1' },
        { value: 'option2', label: 'Option 2' },
        { value: 'option3', label: 'Option 3' }
      ]
    };
  }
};
</script>

```



> `Attributes`：代表的是组件的属性，比如type="primary",:loading="true" 是响应式属性，也是一样的意思

```html
<el-button type="primary" :loading="true">加载中</el-button>

```

> `Method` :Method代表的是组件的方法，可以直接调用Elementui中已经帮我们声明好的方法，下面用El-form组件举出一个例子，使用this.$refs.form.validate调用组件的validate方法，对整个表单的属性进行校验

```html
<template>
  <div>
    <el-form :model="formData" ref="form">
      <el-form-item label="Name">
        <el-input v-model="formData.name" />
      </el-form-item>
      <el-form-item label="Age">
        <el-input v-model.number="formData.age" type="number" />
      </el-form-item>
      <el-form-item>
        <el-button @click="submitForm">Submit</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      formData: { name: '', age: null }
    };
  },
  methods: {
    submitForm() {
      this.$refs.form.validate(valid =>
        valid
          ? console.log('Form data:', this.formData)
          : console.log('Validation failed')
      );
    }
  }
};
</script>

```

> `Events`: 代表的是组件的事件，可以通过@click="onExport"方式进行调用的意思，比如点击事件，鼠标焦点事件，下面用el-button按钮进行举出例子

```html
<template>
 <el-button @click="submitForm">Submit</el-button>
</template>

<script>
export default {
  data() {
    return {
      formData: { name: '', age: null }
    };
  },
  methods: {
    submitForm() {
     console.log('提交事件')
    }
  }
};
</script>

```



### Vue中的{{}},:loading="isLoading",  prop="isEnable",          @click="onExport"的区别

这三种赋值的方法涉及到Vue.js框架中数据绑定和渲染的不同方式。让我为您解释一下它们在什么情况下使用：

1. ` prop="isEnable"`： 这种方式通常用于HTML标签的属性绑定，或者`elmentui组件的属性(Attributes)绑定`，。在HTML标签中中，您可以将属性值直接绑定为一个字符串。在这种情况下，属性值将始终被解释为字符串，而不会被Vue的响应式系统所追踪。

   例如：

   ```html
                 <el-table-column
                   label="是否开启此尺码"
                   prop="isEnable"
                 >
                 </el-table-column>
   或者
      <el-button size="small">小型按钮</el-button>
   
   ```

2. `{{ }}`（双大括号语法）： 这是Vue.js中的插值语法，用于将数据绑定到模板中。在这种情况下，表达式会被解析并渲染为字符串，并且会受到Vue的响应式系统的追踪，当数据发生变化时，模板会自动更新。

   例如：

   ```html
   <span>{{ message }}</span>
   ```

3. `:loading=isLoading`： 这是Vue中用于绑定数据到组件属性或Vue实例属性的方式，通过该方式绑定的数据会受到Vue的响应式系统的追踪。这样，在数据发生变化时，相关的组件或模板也会相应地更新。

   例如（组件属性绑定）：

   ```html
   <!--
   :loading是v-bind:的缩写，下面的语句也可以是：v-bind:loading="isLoading"
   -->
   <template>
     <el-button :loading="isLoading">Submit</el-button>
   </template>
   
   <script>
   export default {
     data() {
       return {
         isLoading: true
       };
     }
   }
   </script>
   ```

4.`@click="onExport"` 是Vue中处理点击事件的一种方式，比如`ElementUI中的事件(Events)方法绑定`，其中`@click` 是一个指令，用于监听元素的点击事件，而 `"onExport"` 则是要在点击事件发生时执行的方法名称。

具体解释如下：

- `@click` 是Vue的事件绑定指令，用于监听元素的点击事件。它等效于原生HTML中的`onclick` 属性，但提供了更灵活的处理方式。
- `"onExport"` 是在点击事件触发时将要调用的方法的名称。这个方法通常定义在当前Vue组件的方法部分。

综上所述：

- `prop="isEnable"` 用于在HTML标签中绑定静态字符串属性，不涉及Vue的响应式系统。
- `{{ }}` 用于在模板中插入表达式的值，使其成为模板的一部分，并且受Vue的响应式系统追踪。
- `:loading=isLoading` 用于将数据绑定到Vue组件或实例的属性，使其受Vue的响应式系统追踪。
-  `@click="onExport"` 用于将方法绑定到点击事件中，是vue用来监听点击事件的指令

根据具体情况，您可以选择合适的绑定方式以实现所需的功能和效果。

### Vue扩展运算符...展开一个Json的所有属性

```javascript
let jsonData = {
  ...this.formData,
  userName: this.$store.state.user.name,
  userId: this.$store.state.user.userId
};
... 是 JavaScript 中的扩展运算符（Spread Operator）。它的主要功能是将一个数组或对象的所有元素或属性展开。

在对象中使用时，它可以将一个对象的所有属性复制到另一个新对象中。例如：
let obj1 = { a: 1, b: 2 };
let obj2 = { ...obj1, c: 3 }; // obj2 现在是 { a: 1, b: 2, c: 3 }
在数组中使用时，它可以将一个数组的所有元素复制到另一个新数组中。例如：
在数组中使用时，它可以将一个数组的所有元素复制到另一个新数组中。例如：



扩展运算符可以用于函数调用，将数组的元素作为函数的参数。例如：
let nums = [1, 2, 3];
console.log(...nums); // 输出 1 2 3
```



### Vue中Element Ui中的方法原来是这样引用的

1.找到组件的方法

![1688630019231](Vue进步.assets/1688630019231.png)

2.给组件一个别名

![1688630055262](Vue进步.assets/1688630055262.png)

3.引用组件的方法

![1688630106588](Vue进步.assets/1688630106588.png)

### 日期联动组件加一天与时间问题

```c#
//加这两个行代码解决两个问题
            value-format="yyyy-MM-dd HH:mm:ss"
            :default-time="['00:00:00', '23:59:59']" 
//设置为日期格式，不会再有16:00的时间限制，不会有加1天的情况，后台不需要再单独处理
//完整代码
        <el-form-item prop="syncEcTimes">
          <el-date-picker
            v-model="filter.syncEcTimes"
            type="daterange"
            range-separator="至"
            start-placeholder="同步开始日期"
            end-placeholder="同步结束日期"
            value-format="yyyy-MM-dd HH:mm:ss"
            :default-time="['00:00:00', '23:59:59']"
          >
            />
          </el-date-picker>
```



### 在父组件中调用子组件的方法



### npm run dev 的实质

让npm环境（nodejs环境）去运行**package.json**（前端项目([node](https://so.csdn.net/so/search?q=node&spm=1001.2101.3001.7020)项目)配置文件）中本项目下

的 scripts脚本区域dev字样属性对应指定的命令行：

![1664333485249](Vue进步.assets/1664333485249.png)

### Vue state mutations 状态管理（理解为全局静态变量）

[Vuex笔记-state与mutations扩展学习](https://blog.csdn.net/init_8y0/article/details/115289529)

**mutations作用域**: mutations的使用方法类似于vue实例的methods，可以在mutations作用域里面写方法体，如下所示：（mutations中的函数，方法钟默认第一个参数为state）

![1661939044539](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1661939044539.png)

**commit方法**:使用this.$store.commit()对mutaions函数进行提交

比如：this.$store.commit('decrement') 就是调用store中的decrement方法，我们可以在其他的脚本中引用store的Js，然后引用方法（好像不需要，是全局注册的）：

![1661939424688](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1661939424688.png)

### ES6 对象解析语法 getUserInfo({ commit, state }) 

对象解析语法就是一个Json变量赋值的语法糖，简化代码



```c#
var jsonA={name:'王锐',age:20}
//未简化前
var name=jsonA.name
var age=jsonA.age
 如果在方法中使用解析语法入参，那么就是getAge({age}){...},调用getAge(jsonA),直接就可以拿到jsonA中Age的值
//使用对象解析语法
const {name,age}=jsonA
//经验1 值为NULL自动转换默认值：有时候后端返回没有赋值的情况，也可以自动转换为默认值0，也可以手动赋值默认值
var jsonA={name:'王锐'} //没有 age:null
const {name,age}=jsonA  //age值为null不会异常
const {name,age=18}=jsonA //也可以在没有值的时候赋值默认值
//经验2 按需取值
const {name}=jsonA    //可以只取其中一个值
//经验3 嵌套取值
var jsonA={name:'王锐',age:20,school:{ name:'东山小学',address:'东山浒厂' }}
var schoolName=jsonA.school.name //历史取值
const {school：{name}}=jsonA //解析语法取值
//经验3 命名冲突解决
const { age: userAge } = jsonA; //如果当前命名空间已经有age这个命名了，那么可以用这种语法既可以解析出正确的值，又可以创建值的别名
```

[解析语法扩展阅读](https://blog.csdn.net/sunhuaqiang1/article/details/124889449)

### 怎么引用export default getters 后变量的值？

​    使用import 导入关键字，就像导入方法一样，注意：export与export default的区别？

### export与export default的区别？

1. 单个，只能有单个：export default 输出单个值，使用import关键字引用的时候不需要使用大括号{}，直接引用就可以
2. 多个，可以有多个：export 输出多个值，使用Import导入引用的时候需要使用大括号{} （就像平时引用方法一样）

![1661940930233](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1661940930233.png)

**一个js文件中，只能有一个export default； 但是，一个js文件中，可以有多个export。**

### 局部引用与全局引用

**局部引用**：按需要在单独的页面引用对应的组件

![1661941347771](Vue进步.assets/1661941347771.png)

**全局引用**：在main.js里面引入(下面引用的store与router都是全局引用文件夹，可以直接在全局调用里面的变量的值)

![1661941177423](Vue进步.assets/1661941177423.png)

- 怎么给state加一个变量值进行返回，并且在全局进行调用？

   

   ### 





### vuex中 this.$store.dispatch() 与 this.$store.commit()方法的区别

commit 和 dispatch 两个方法都是传值给vuex的mutation改变state，区别在于commit是同步调用，dispatch是异步的调用

```c#
//commit同步调用方法,写法：this.$store.commit(‘mutations方法名’,值)
 this.$store.commit('changeValue',name)  //存储
 this.$store.state.changeValue  //取值
   使用commit方法的时候，就是执行完方法就执行完了，没有什么其他的动作
// dispatch：异步调用方法,写法：this.$store.dispatch(‘action方法名’,值)
 this.$store.dispatch('getlists',name) //存储
 this.$store.getters.getlists //取值
   //使用dispatch：当调用方法结束后，我还想做点什么就用这个方法。 提交actions并且携带一个参数后，返回一个promise, 就可以执行异步操作，
      this.$store.dispatch('action1', 'Bootstrap').then(（） => {
        console.log('下班了，我想去美院跑步！');
      });
```

[扩展阅读](https://blog.csdn.net/qq_44317018/article/details/105784086)



### ...es6中的扩展运算符

扩展运算符是用来迭代集合使用的（也可以用于单个的字符串，或者对象解析语法）

```c#
//作用1: 数组赋值
var arrayA=[a,b,c]
var arrayB=[e,f,g]
arrayA.push(...arrayB)
console.log(arrayA) //a,b,c,e,f,g
//作用2：合并数组
var arrayA=[a,b,c]
var arrayB=[e,f,g]
var arrayUnion=[...arrayA,...arrayB]
console.log(arrayUnion) //a,b,c,e,f,g
//作用3：字符串拆分
var strA='abc'
var arrayA=[...strA]  
console.log(arrayA) //a,b,c
//作用4：对象解析
let obj = {a:1,b:2,c:3,d:4,e:5}
let {a,...rest} = obj;
console.log(a);//1
console.log(rest);//{b:2,c:3,d:4,e:5}
```

[扩展学习](https://blog.csdn.net/qq_44317018/article/details/105784086)

### 【新的Vue调试工具的使用】



**还要选中左边的App 1**

![1660790924718](Vue进步.assets/1660790924718.png)

调试时出现：三个页面，好像三个太阳

![1692428220069](Vue进步.assets/1692428220069.png)

### 100秒就会超时 ，请求BulkUpdate

![1654827291958](Vue进步.assets/1654827291958.png)

### 获取用户信息

这是新的系统：

```c#
userName: this.$store.state.user.name 
userId: this.$store.state.user.userId 

//获取角色编号
this.$store.getters.roles
```
这是旧的系统：

```c#
      
      userName: this.$store.state.user.userName,
      userId: this.$store.state.user.userId,
      userRole: this.$store.state.user.access,
```

代码路径

![1651808915900](Vue进步.assets/1651808915900.png)

![1654740896663](Vue进步.assets/1654740896663.png)

### 新建模块如何可以不在正式环境展示

访问地址：

[测试--产品资料管理](http://localhost:3000/center-mgr)

[正式-产品资料管理](https://admin.shengweiinc.com/center-mgr)

注释原链接

![1649215230764](Vue进步.assets/1649215230764.png)

注意要加斜杠/

![1649215341582](Vue进步.assets/1649215341582.png)



### export default是白色的怎么办

![1647414915867](Vue进步.assets/1647414915867.png)

放下来就好了

![1647414939300](Vue进步.assets/1647414939300.png)

### 更改ElementUI控件样式只在当前页面生效的方法

包一层DIV的class解决

```c#
// .upload-hint 是包住elementUI控件的DIV样式，这样就只作用到这个DIV下面
.upload-hint .el-upload-dragger{
    height:320px;
    width:340px;
    display: table-cell;
    vertical-align: middle;
}
```



### JSZip解压，文件名乱码问题解决



[JSZip解压，文件名乱码问题解决](https://blog.csdn.net/qq_40036799/article/details/117280941)

扩展学习： [GB2312、GBK、Big5汉字编码](https://blog.csdn.net/lengshine/article/details/5470545)，

- GB2312 是中文字符编码

- Big5 是台湾与香港使用的繁体字编码

- GBK 是中文字符编码与繁体字字符编码


### 可以马上调试的方法：debugger

![1647229477952](Vue进步.assets/1647229477952.png)

### 弹性布局（display:flex;）属性详解

[弹性布局（display:flex;）属性详解](https://www.cnblogs.com/hellocd/p/10443237.html)

### Vue引用当前目录下面的组件

 import job from './job'; 表示当前目录下导入job.vue

### V-For的使用经验（注意头部，注意Key）

```c#
<template>
  //1.注意头部：不允许在<div class="import-image"> 这个根目录使用v-for指令,模板根目录禁止使用“v-for”指令(可以再用div包一层就可以在头部使用)
  <div class="import-image">
    <el-image
      v-for="imageUrl in list"   
 // 2.注意Key：v-for与后端的foreach一样,先属性再集合，然后记得要加:key属性，这是（Eslint代码检查）  
      :key="imageUrl"  
      style="width: 100px; height: 100px"
      :src="imageUrl"
      :fit="fit"
    />
    <div style="padding: 14px;">
      <span>好吃的汉堡</span>

    </div>

  </div>
</template>
```



### 判断整形值的时候注意等于0的情况

如果isQC的值等于0那么值会等于false（ms.isQC本来是判断有没有值，而不是判断值是不是等于1），而不是等于1

![1646115764455](Vue进步.assets/1646115764455.png)

### 方法里面可以给属性赋值

show = false

![1646019778807](Vue进步.assets/1646019778807.png)

###  数组可以使用object?.label运算符

目前假设只有数组支持?.写法

![1643186881581](Vue进步.assets/1643186881581.png)

### 自定义属性快捷传参

![1643166334421](Vue进步.assets/1643166334421.png)

### 展示图片显示图片

```
      <el-table-column prop="productImage" label="图片" width="80">
        <template slot-scope="scope">
          <el-popover
            placement="right"
            trigger="hover"
          >
            <img :src="scope.row.mainImage" width="350" height="350">
            <img slot="reference" :src="scope.row.mainImage" width="65" height="65">
          </el-popover>
        </template>
      </el-table-column>
```



### ElementUI组件的使用

Table Attributes 属性
Table Events  事件
Table Methods 方法
Table Slot   插槽
Table-column Attributes 列的属性

Table-column Scoped Slot 列的插槽

### 新建页面需要注意的内容

simplehead



### 在子组件中更改父组件的值，发送响应式消息

比如this.selectIds是在父组件中的变量，我在子组件中更改了，Vue并不会发送响应式的消息给父组件，可以调用this.$emit('changed', this.selectIds)，传入新值给父组件的变量即可

```
      this.$emit('changed', this.selectIds)
```

单个使用

```c#
          //子窗体调用父窗体的事件：reloadTable 添加数据后，刷新页面的数据
          this.$emit('reloadTable')
              
          //父窗体引用：将刷新方法传入进去
   <upload-file
           ...省略其他参数
      @reloadTable="reloadTable"
    />

```



### ElementUI事件的绑定与使用

绑定属性用冒号：:Name ，绑定事件（方法）用艾特：@GetName

![1642499533270](Vue进步.assets/1642499533270.png)

s

### 前端三元表达式返回true/false简写：

```c#
//简写(返回true)
this.formDisabled = !!this.id
//通常写法
this.formDisabled = this.id?true:false
```



### 特性器：> > >三个箭头

 **!importanty：** 也可以使用这个关键字进行覆盖表格的样式

[特性器的解释](https://developer.mozilla.org/en-US/docs/Web/CSS/Specificity)

### 调试发布后Vue组件里面的脚本

![1640857899723](Vue进步.assets/1640857899723.png)

### 前端页面也可以使用三元表达式

? true:false ，代替if(有值){ 再做什么}

### 浏览器调试

使用debugger;关键字放到vue的script代码里面，就可以在前端进行调试了

也可以在浏览器的脚本里面直接断点就可以调试（这种需要先定位到错误的脚本行），实用性有缺陷（如果我需要主动定位错误行数呢）

### 不允许有两个相同名称，相同类型的请求方法

比如getList，可以分别有一个post,get方法，但是不能有两个post,get相同的方法名称，不同参数的方法，axios请求的时候出现：Multiple actions matched. The following actions matched route data and had all constraints satisfied: （匹配多个方法的提示）



### 前端进行值的重复校验

**如果编辑的时候不允许编辑这个属性，那么只需要在新增方法限制就可以**

   1.编写后端接口读取所有TemplateName值，转换成List集合  
   2.编辑方法返回的Dto新增AllTemplateName属性，把list集合赋值给它 
   3.open进来初始化的时候赋空值  
   4.前端新增data 属性 AllTemplateName  
   5.输入的时候前端进行验证是否重复 
   6.保存的时候前端进行验证是否重复 
   6.保存的时候在后台进行验证 TemplateName 是否重复  

### 检查属性是否没有重置为空

打开添加窗体，进行查看

![1640593654260](Vue进步.assets/1640593654260.png)

### linq Group by 

```c#
//group by 后通过key.属性获取值      
var groupbyDetails = details.GroupBy(l => new { l.SpecificName, l.ShowType, l.ListValue }).Select(p=> new ListerCategorySpecificMappingDetailDto() {
                SpecificName = p.Key.SpecificName,
                ShowType = p.Key.ShowType,
                ListValue = p.Key.ListValue.IsNullOrWhitespace() ? null : p.Key.ListValue.Split('|').ToList()
            }).ToList();

//group by 后拿到明细的属性信息，拿到颜色在款式中重复的颜色(原来的List内容，直接u=>u就可以调用了，都不用u.key)
var beEvenColor = oldAllProductPropertys.GroupBy(u => u.ProductSpu).Select(u => new ProductPropertySpecificModifyDto { ProductSpu = u.Key, ColorList = u.Select(s => s.Color).GroupBy(l => l).Where(w => w.Count() > 1).Select(l => l.Key).ToList() }).ToList();
```



### 启动程序swagger不出来

关闭黑窗体，重启consul，然后启动程序

### table列表操作的悬浮：

fixed="right"

### 未找到：InternalAsyncHelper.cs 



**错误提示：**未找到：InternalAsyncHelper.cs 

**异常代码：**
          

```
var categoryList = await _categoryBailunRepository.Query().AsNoTracking().ToListAsync();
```

**解决方案：**这个是异步的请求，所有会有未找到：InternalAsyncHelper.cs 这个提示，把断点打到下面一行，按F5就可以返回了

### 33万条类目数据写入到redis提示错误

问题：SocketFailure (ReadSocketError/ConnectionReset, last-recv: 5) on 127.0.0.1:6379/Interactive, Idle, last: SET, origin: ReadFromPipe 

解决：使用入参条件作为key，进行针对性的缓存（比如缓存美国站点，英国站点的类目数据）

### 请求其他项目的接口方法

```c#
        public bool ExistsAmazonCategory(string siteCode, long categoryId)
        {
            if (siteCode.IsNullOrEmpty() || categoryId.IsNullOrEmpty())
            {
                return false;
            }

            //var param = $"{productIds.JoinAsString("&")}&siteCode={siteCode}";
            var param = $"siteCode={siteCode}&categoryId={categoryId}";
            var apiUrl = $"{SaasServersConsts.SaasWalmartServerUrl}/AmazonLister/GetCategoryTree?{param}";
            string response = null;
            try
            {
                response = HttpClientManager.HttpGet(apiUrl);
            }
            catch (Exception ex)
            {

                throw;
            }

            if (response.IsNullOrEmpty())
            {
                return false;
            }

            var category = JsonHelper.FromJson<AmazonCategoryDto>(response);
            if (category == null || category.IsLeaf == false)
                return false;
            return true;
        }
```



 IConfigurationRoot _appConfiguration 是继承IConfiguration的子类

![1640314542322](Vue进步.assets/1640314542322.png)

新项目中请求的路径：

![1640314284483](Vue进步.assets/1640314284483.png)



walmart 配置的路径

![1640314053814](Vue进步.assets/1640314053814.png)

### 数组入参前后端怎么写

总结：后端用HttPost请求，使用[FromBody],前端传数组，包在data:array 这样赋值（注意FromBody是不允许在HttpGet里面使用的）

注意：被标记为[FromBody]的参数只允许出现一次， 如果有多个参数那么都放到一个实体类中进行入参

![1640251566980](Vue进步.assets/1640251566980.png)

![1640251579042](Vue进步.assets/1640251579042.png)

![1640251596187](Vue进步.assets/1640251596187.png)

### 找前端对应的接口

接口不指定baseURL就默认是VUE_APP_BASE_API，也就是前端的端口号

![1640249871162](Vue进步.assets/1640249871162.png)

### ToListAsync()异步方法

```
在命名空间： Microsoft.EntityFrameworkCore 下面
引用了命名空间就可以使用，可以复制可以马上用：

ToListAsync() 
```



### Ef去除跟踪的方法

```
//记得复制完整，要带上括号，这个在命名空间： Microsoft.EntityFrameworkCore 下面
AsNoTracking()
```



### 裕文的分页查询

![1640224732867](Vue进步.assets/1640224732867.png)

```c#
 private readonly IOriginalProductInfoRepository _productInfoRepository;
        private readonly IMapper _mapper;
        public OriginalProductInfoService(
            IMapper mapper,
            IOriginalProductInfoRepository OriginalProductInfoRepository
            )
        {
            _productInfoRepository = OriginalProductInfoRepository;
            _mapper = mapper;
        }
        /// <summary>
        /// 分页查询
        /// </summary>
        /// <param name="input"></param>
        /// <returns></returns>
        public async Task<PageResult<List<OriginalProductInfoPageDto>>> GetPages(OriginalProductInfoSearchInput input)
        {
            try
            {
                var query = BuildQuery(input);
                if (!input.SortField.IsNullOrEmpty())
                {
                    query = LambdaHelper.GetOrderByAndSort(query, input.SortField, input.Sort);
                }
              //这个也可以参考    query.Skip((input.PageIndex - 1) * input.PageSize).Take(input.PageSize);

                var data = query.Skip((input.PageIndex - 1) * input.PageNumber).Take(input.PageNumber).ToList();
                var resultList = _mapper.Map<List<OriginalProductInfoPageDto>>(data);
                var total = await query.CountAsync();
                return new PageResult<List<OriginalProductInfoPageDto>>().ToPageResult(input.PageIndex, total, resultList);
            }
            catch (Exception ex)
            {
                throw ex;
            }
        }
        private IQueryable<OriginalProductInfo> BuildQuery(OriginalProductInfoSearchInput input)
        {
            var query = _productInfoRepository.Query().AsNoTracking().Where(u => !u.IsDeleted);

            if (!input.ProSku.IsNullOrEmpty())
            {
                query = query.Where(u => u.ProSku == input.ProSku);
            }
            return query;
        }
```



### Excel的导入

项目名称：顺序数第四个，**Bailun.PublishListing**

后端控制器：CleanStockController

![1640225217011](Vue进步.assets/1640225217011.png)

### 表单的校验提示往左边缩

```c#
//1.给form添加  label-width="80px"属性
  <el-form
           ...省略
          label-width="80px"
        >
 //2.   给table添加样式
<el-table
           ...省略
            class="tbClass"
          >
  //3.  最底部新增样式
  <style scoped>
.tbClass >>> .el-form-item__content {
  margin-left: 0px !important;
}
</style>
          
```



### **用 form.resetFields清空表单**：

```c#
//检查  ref="form" 是否赋值
<el-form ref="form" :model="formData" :rules="formValidate" label-width="120px">

//清空表单值（有时候会有问题，可以尝试放到close事件里面进行清空，或者手动清空）
   this.$nextTick(() => {
        this.$refs.form.resetFields()
      })
    
//清空表单值进阶版本(删除):
  //初始也需要先赋值为空
   this.formData.bulletPoint = []   //可以用这个清空，我看辉哥就是用这个清空值的
// 清空form表单中的信息：清空时再进行判断，如果是新增进来，并且值不等于空，再清空一次
      this.$nextTick(() => {
        this.$refs.form.resetFields()
        if (!id && this.formData.bulletPoint.length > 0) {
          this.formData.bulletPoint = []
          for (let index = 0; index < 5; index++) {
            this.formData.bulletPoint.push(
              { value: null }
            )
          }
        }
      })
    
  //【清空校验信息】（有点问题，提示没有这个方法）
    this.$refs.form.clearValidate()
```



### VUE调试

![1640081269036](Vue进步.assets/1640081269036.png)

### 异步的查询与删除

ToListAsync是在命名空间：

![1639817260918](Vue进步.assets/1639817260918.png)



### 快速更改第三方组件样式

比如下面的就是更改控件的宽度

![1640052499199](Vue进步.assets/1640052499199.png)

### Vue中数组$set赋值

```c#
 //数组，改变的位置（从0开始），新的值
 this.$set(this.items,2,'newvalue')
 
 //清空数组(用[]Vue可以正常响应,用下面的是多一种方法，时候使用form.resetFields清空表单不知道为什么会有自动赋值的情况出现)
  this.items.splice(0,items.length)
     
```



### 首页的查询用Post进行查询

```

我怀疑Get查询会一直错误，用Post正常

有时候请求不进来可能是consul出现问题了， 关闭重新注册consul，然后发起接口请求
```



### 快速引用命名空间

```
ait+回车
```

### 引用PPS新版前端的Table

```c#
//引用Table组件
import SimpleTable from '@/components/SimpleTable'


// 导入复制指令
import clipboard from '@/directive/clipboard/index.js'
    
    
//使用组件,ref表格名称,:url赋值的请求URL,method 请求类型,filter 查询的条件，offset 距离顶部的高度
<simple-table
      ref="table"  
      :filter="filter"
      :url="url"
      method="post"
      :offset="14"
      @col-change="val=>tableCols=val"
    >
    //选择框
      <el-table-column
        type="selection"
        width="40"
      />
      //自定义组件，有那个下拉箭头的
      <el-table-column
        type="expand"
        width="32"
      >
        <template slot-scope="scope">
          <detailslist :data="scope.row" />
        </template>
      </emitel-table-column>
      //使用score处理行里面的值，循环，或者拼接，里面加入了复制的指令
      <el-table-column
        prop="orderId"
        label="订单号"
        width="140"
      >
        <template slot-scope="scope">
          <div>
            <span
              v-clipboard:copy="scope.row.orderId"
              v-clipboard:success="clipboardSuccess"
            >{{ scope.row.orderId }}</span>
          </div>
        </template>
      </el-table-column>
 </simple-table>
 
                  //  
    data() {
    return {
     //请求数据的地址
      url: getFruugoOrderList.url,
     //查询的条件
      filter: {
        shopId: null,
        attributeValue: null,
        attributeName: null
      }
    }
  }
                  
                  
    //复制成功提示的方法
  clipboardSuccess() {
      this.$message({
        message: '复制成功',
        type: 'success',
        duration: 1000
      })
    }
```

### FirstOrDefault和First的区别

FirstOrDefault ：查询出来的序列为空的时候不会抛异常

First：查询出来的序列为空的时候会抛异常（就像对象为空的时候如果你访问他的属性，那就会异常的意思一行）

[FirstOrDefault和First的区别](https://blog.csdn.net/qq_29660957/article/details/107385777)

### 保存的时候增删改模板

1. 读取页面信息：把页面传输的DTO变量绑定到实体对象上
2. 更新新增？： 根据传输过来的对象ID判断是更新还是新增
3. 新增：直接批量新增
4. 主表：主表使用Dto更新对象，直接更新。
5. 副表增删改：声明新增，删除的List对象
6. 删：旧新循环(新where)，通过唯一属性比对，如果新中不存在旧的唯一属性，则对象为删除，保存对象，批量删除。
7. 更（放到删里面处理）：旧新循环(新where)，通过唯一属性比对，如果新中存在旧的唯一属性则为更新，更新旧对象的值，批量更新
8. 增： 新旧循环(旧where)，通过唯一属性对比，如果旧中不存在新的唯一属性，则对象为新增，保存对象，批量新增

   **批量更新：**

```c#
//批量更新赋值
  oldDetails.ForEach(item =>
  {
 if (details.Where(l => l.PlatformName == item.PlatformName ).Any())
{
 item.MappingName = details.Where(l => l.PlatformName == item.PlatformName ).FirstOrDefault().MappingName;
item.LastModifierUserName = HttpContextHelper.Current?.User?.GetUserName();
item.LastModifierUserId = HttpContextHelper.Current?.User?.GetUid().ToStr();
 item.LastModificationTime = DateTime.Now;
                            }
                            else
                            {
                                deleteDetails.Add(item);
                            }
                        });
//批量更新
_listerColorMappingDetailRepository.BulkUpdate(oldDetails.ToArray(), P =>
                        {
                            P.BatchSize = 100;
                            P.ColumnInputExpression = c => new
                            {
                                c.Id,
                                c.MappingName,
                                c.LastModifierUserName,
                                c.LastModifierUserId,
                                c.LastModificationTime

                            };
                        });
```

**批量新增：**

```
  item.CreateUserName = HttpContextHelper.Current?.User?.GetUserName();
  item.CreatorUserId = HttpContextHelper.Current?.User?.GetUid().ToStr();
  item.CreationTime = DateTime.Now;
```



### Scope.$index无法正常取值

![1639811289034](Vue进步.assets/1639811289034.png)

### 异常的处理

![1639795832800](Vue进步.assets/1639795832800.png)

### Get与Post传参区别

post：数组，实体类

get:1至3个参数的请求

```c#
// get明细信息 (get的参数需要特殊构造data的格式)
export function getDetail(id) {
  return request({
    url: '/ListerColorMapping/GetDetail',
    method: 'get',
    params: {
      id: id
    },
    baseURL: process.env.VUE_APP_SW_SKUMS
  })
}

//post 保存 （post的参数直接传参data就可以）
export function saveColorMapping(data) {
  return request({
    url: '/ListerColorMapping/SaveColorMapping',
    method: 'post',
    data: data,
    baseURL: process.env.VUE_APP_SW_SKUMS
  })
}

```



### 引用PPS搜索的文本框：

```c#
//输入文本框右边的，搜索，重置组件
import SimpleHead from '@/components/SimpleHead'
//店铺的下拉
import PlatformShopSelect from '@/components/Common/PlatformShopSelect'
    
<simple-head
      :model="filter"
      @on-search="reloadTable(true)"
    >
      <template slot="button">
        <!-- 操作按钮 -->
        <el-button type="success" size="small" icon="el-icon-download" @click="onExport">导出订单</el-button>
      </template>
      <!-- 默认搜索项 -->
      <el-form-item prop="shopId">
        <PlatformShopSelect
          v-model="filter.shopId"
          platform="fruugo"
        />
      </el-form-item>
      <el-form-item prop="attributeValue">
        <el-input
          v-model="filter.attributeValue"
          placeholder="请输入内容"
          clearable
        >
          <el-select
            slot="prepend"
            v-model="filter.attributeName"
            placeholder="请选择"
            style="width:110px"
            clearable
          >
            <el-option
              label="平台编码"
              value="skuId"
            />
            <el-option
              label="产品标题"
              value="skuName"
            />
            <el-option
              label="订单号"
              value="orderId"
            />
          </el-select>
        </el-input>
      </el-form-item>
      <!-- 更多搜索项 -->
      <template slot="more" />
    </simple-head>
```



### 模板或渲染函数未定义

```c#
//解决：关闭站点，重新Npm启动站点（后面可能值不更新，那么可以单独在对应的页面更新里面的message的值进行保存，Vs code里面ctrl+s是保存当前页面的文件）
Failed to mount component: template or render function not defined.

```

### Dialog外部点击不关闭

```
 :close-on-click-modal="false"
```

### Form表单统一设置大小

```
 size="small"
```

### 浏览器重新加载与F5刷新区别

```
F5：刷新的时候会读取页面的数据
右键重新加载：会清理缓存，并重新加载页面
npm run dev ：重新加载配置，加载页面保存配置，代码是最新的，但是不会清除缓存
```



### 灵魂垫高的内容(放取消，确认)

在dialog弹窗组件里面使用，放到最下面使用

```c#
   //放到Dialog最下面使用
   <div slot="footer" class="dialog-footer">
      <el-button @click="show = false">取 消</el-button>
      <el-button type="primary" :loading="okLoading" @click="ok()">保存</el-button>
    </div>
  </el-dialog>
```



### 公共的平台，店铺组件

![1639649944406](Vue进步.assets/1639649944406.png)

### 导入Excel组件

```c#
 //引用无注释版本
  <upload-file
      ref="uploadFile"
      :template-down-url="templateDownUrl"  
      :upload-url="uploadUrl"   
      @importToTable="importToTable"   
      @reloadTable="reloadTable"  
    />
    
//引用 (注释版本)
 <upload-file
      ref="uploadFile"
      :template-down-url="templateDownUrl"  --Excel的下载模板
      :upload-url="uploadUrl"    --验证excel数据合法性的API
      @importToTable="importToTable"   --点击确定，触发的方法，入库的API
      @reloadTable="reloadTable"   --刷新数据的方法
    />

          
 //导入       
import UploadFile from '@/components/Upload/UploadFile2.vue'
```



### 快速新建页面（内部组件）

```c#
<template>
  <div class="app-container">
    <simple-head
      :model="filter"
      @on-search="reloadTable(true)"
    >
      <template slot="button">
        <!-- 操作按钮 -->
        <el-button type="success" size="small" icon="el-icon-download" @click="onImport">导入</el-button>
        <el-button type="success" size="small" icon="el-icon-download" @click="onAdd">添加</el-button>

      </template>
      <el-form-item prop="attributeValue">
        <el-input
          v-model="filter.attributeValue"
          placeholder="请输入颜色中文"
          clearable
        >
          <el-select
            slot="prepend"
            v-model="filter.attributeName"
            placeholder="请选择"
            style="width:110px"
            clearable
          >
            <el-option
              label="颜色中文"
              value="skuId"
            />
          </el-select>
        </el-input>
      </el-form-item>
    </simple-head>
    <simple-table
      ref="table"
      :filter="filter"
      :url="url"
      method="post"
      :offset="14"
      @col-change="val=>tableCols=val"
    >
      <el-table-column
        type="selection"
        width="40"
      />
      <el-table-column
        prop="totalNumberOfItems"
        label="中文"
        width="180"
      />
      <el-table-column
        prop="totalNumberOfItems"
        label="操作"
        width="180"
      >
        <template slot-scope="scope">
          <div>
            <el-button type="primary" plain @click="onEdit(scope.row)">编辑</el-button>
          </div>
        </template>
      </el-table-column>

      <el-table-column
        prop="totalNumberOfItems"
        label="添加人"
        width="180"
      >
        <template slot-scope="scope">
          <div>
            <span>{{ scope.row.CreateMan }}</span><br>
            <span>{{ scope.row.CreateDate }}</span>
          </div>
        </template>
      </el-table-column>
      <el-table-column
        prop="totalNumberOfItems"
        label="修改人"
        width="180"
      >
        <template slot-scope="scope">
          <div>
            <span>{{ scope.row.updateMan }}</span><br>
            <span>{{ scope.row.updateDate }}</span>
          </div>
        </template>
      </el-table-column>
    </simple-table>
    <ColorDetail ref="colorDetail" @reloadTable="reloadTable" />
    <!-- <ColorImport ref="ColorImport" @reloadTable="reloadTable" /> -->
    <upload-file
      ref="uploadFile"
      :template-down-url="templateDownUrl"
      :upload-url="uploadUrl"
      @importToTable="importToTable"
      @reloadTable="reloadTable"
    />
  </div>
</template>

<script>
import SimpleTable from '@/components/SimpleTable'
import SimpleHead from '@/components/SimpleHead'
import UploadFile from '@/components/Upload/UploadFile2.vue'

// 访问Api请求
// import { getFruugoOrderList, exportOrder } from '@/api/fruugoorder'

export default {
  name: 'ColorEnglish',
  components: {
    SimpleTable,
    SimpleHead,
    ColorDetail: () => import('./color-detail'),
    // ColorImport: () => import('./color-import')
    UploadFile
  },
  data() {
    return {
      url: '',
      templateDownUrl: null, // 实例模板下载网络路径（需要先上传到服务器）
      uploadUrl: null, // 验证模板数据合法性的WebApi
      filter: {
        shopId: null,
        attributeValue: null,
        attributeName: null
      }
    }
  },
  methods: {

    // 刷新表格
    reloadTable(again) {
      this.$refs.table.reload(again)
    },
    // 添加数据
    onAdd() {
      this.$refs.colorDetail.open()
    },
    // 打开Excel导入视图开始导入
    onImport() {
      this.$refs.uploadFile.open()
    },
    // 入库的方法
    importToTable(filePaths) {
    // importColorExcelToTable(filePaths[0]).then(res => {
    //     if (res.statusCode === 200) {
    //       if (res.result.isSucceeded) {
    //         this.$message.success('保存成功！')
    //         this.$refs.uploadFile.close()
    //         this.reloadTable()
    //       } else {
    //         this.$message.error(res.result.message)
    //         this.$refs.uploadFile.setLoading(false)
    //       }
    //     } else {
    //       this.$message.error(res.message)
    //       this.$refs.uploadFile.setLoading(false)
    //     }
    //   })
    },
    // 编辑数据
    onEdit(rows) {
      this.$refs.colorDetail.open(rows)
    }

  }
}
</script>

```





### Get请求传参

```c#
  //后端 (传属性名称就可以)
        [HttpGet("GetFyndiqListerVariations")]
        public List<FyndIqListerVariationDetailDto> GetFyndiqListerVariations(long listerId) 
        {
            //long listerId = 424;
            return _fyndiqListerService.GetFyndiqListerVariations(listerId);
        }
  //前端请求（  参数会这样拼接到后端?listerId=424）
export function getVariations(listerId) {
  return request({
    url: '/WalmartApi/FyndiqLister/GetFyndiqListerVariations',
    method: 'get',
    params: {
      listerId: listerId
    }
  })
}
```

### **引用组件**

```c#
import PlatformShopSelect from '@/components/Common/PlatformShopSelect'
```

### **引用API**

```c#
import { getFyndiqListingList, setRunStockType } from '@/api/fyndiqlisting'

```

### **成功的弹窗**

```c#
 this.$message.warning(res.message)
 //完整的成功：
 if (res.isSucceeded) {
    _this.$message.success(res.message)
    _this.$emit('reloadTable') // 成功响应事件
    _this.close()
   } else {
  _this.$message.warning(res.message)
  }
```

### **失败的弹窗**

```
 this.$message.error(err.message)
```

### **获取商家信息**

```c#
//引用组件
import PlatformShopSelect from '@/components/Common/PlatformShopSelect'
//使用组件
<PlatformShopSelect ref="shop" v-model="formData.accountId" platform="fyndiq" style="width:50%" />
  //编写逻辑读取商家名称（shop是ref的名字，可以自定义,ref就像一个桥梁，可以通过他访问子组件中的find方法）
  const _shop = this.$refs.shop.find(this.formData.accountId)
   if (_shop) {
     this.formData.accountName = _shop.account
  }
```

### **获取form选中行**

```c#
 //获取当前选中值     
const checksRow = this.$refs.table.getSelection()
      console.log('checks', checks)
      if (checksRow.length === 0) {
        this.$message.warning('请选择单条或多条数据操作')
        return
      }
//es6 map过滤
  this.listerIds = checksRow.map(u => u.id)
   //传统过滤方法
      const listerIds = []
      checks.forEach(e => {
          listerIds.push(e.id)
      })
  
```

重置页面的数据**

![image-20211201150522932](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20211201150522932.png)

Vue 重置页面data中数据：https://www.jianshu.com/p/f6c88d70a24b

**无法加载文件** D:\node-v14.12.0\vue.ps1，因为在此系统上禁止运行脚本

（1）以管理员身份运行VSCode

（2）执行命令：get-ExecutionPolicy（取得shell的当前执行策略，显示Restricted（表示状态是禁止的））

（3）执行命令：set-ExecutionPolicy RemoteSigned （开启远程下载权限）

（4）执行命令：get-ExecutionPolicy，显示RemoteSigned

### 驼峰自动加-

组件的命名驼峰命名可以直接转中横线

下面声明：**ArticleDetail**  引用直接 **article-detail** 也可以

![1637199983719](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637199983719.png)

### Vue中Main执行的生命周期

**src/main.js**文件 入口文件，初始化vue实例并使用需要的插件（引入公共element组件等、路由拦截、start登录、鉴权可以在这里统一处理）

**router/index.js：**执行后应该是直接到路由这边跳转到登录界面 **答：老师说是通过webpack进行配置的**

### Vue生命周期

声明周期：简单就是打开页面加载的内容，渲染完数据的内容，加载完数据，窗体关闭的时间

https://www.jianshu.com/p/672e967e201c 

![11370083-f279314aef6741db](C:\Users\12987\Desktop\11370083-f279314aef6741db.jpg)



### 后端跨域

在Asp.Net Core中可以进行跨域配置（使用**3、设置特定来源可以跨域** 进行解决） https://www.cnblogs.com/dotnet261010/p/10177166.html，日常的前后端分离开发中我们可以通过，跨域的配置进行解决跨域的问题，线上的环境可以通过Nginx部署方向代理进行解决跨域的问题

**ACE老师跨域问题解决方案：**

```c#
HttpContextHelper//startup.ConfigureServices中写在第一行
 services.AddCors(c => c.AddPolicy("any", p => p.AllowAnyMethod().AllowAnyOrigin().AllowAnyHeader()));

//startup.Configure中写在 app.UseRouting();下面
app.UseCors("any");

```

### VUE打包部署

[webpack打包vue项目之后生成的dist文件该怎么启动运行](https://www.cnblogs.com/zhujiabin/p/10557982.html)



## VS



### 我发现当我帮助别人的时候，比如杨魁

**我非常兴奋，感觉自己有能量，有价值**

裕文说采购链接为空，我处理了一下，也很开心

我发现我和人交流的时候我的大脑积极性很足够，一直坐在一个位置，精神状态很萎靡（这是不是我在佳禾做了五年运维技术支撑工作的原因）

当我心情沮丧的时候，碰到有人问我的问题，我的感受是激动的，兴奋的，马上就可以转换好自己的状态

![1662347038177](Vue进步.assets/1662347038177.png)

比如杨魁的这句话

我发现我适合和人交流的工作，比如之前在佳禾的工作中，我做研发工作，并且我也和人交流，我的精力很好，如果一直待在电脑前，工作内容也和其他人交流很少，我的精力就很少，得不到很多提升

**很多人问我问题的时候，我不会感觉烦躁，我会感觉很兴奋**

![1663291522848](Vue进步.assets/1663291522848.png)





### 【高质量答案的搜索渠道】

- learn.microsoft.com/zh-cn
- stackoverflow.com/
- blog.51cto.com
- cnblogs.com/
- cloud.tencent.com/

### VS2022查看项目的类图

https://www.cnblogs.com/douzi2/p/17002600.html

[Visual Studio 2022 查看类关系图](https://www.cnblogs.com/douzi2/p/17002600.html)



步骤1：勾选扩展开发

![1705650443163](Vue进步.assets/1705650443163.png)

![img](Vue进步.assets/7824f9ca7b5a4c1c86b34c8594f8677b.png)

 步骤2:  选择单个组件，搜索类，勾选类设计器

![img](Vue进步.assets/5f4e72a4e3954e4c948f3bddc448403c.png)



右键要查看的项目-“查看”-“查看类图”

![img](Vue进步.assets/ab930dbc6c2248f58f5575fd19b77509.png)

效果展示：

![img](Vue进步.assets/4034c87cb409414a98d0f3e7356df987.png)

 

作者:[宋桓公](http://www.cnblogs.com/douzi2/)
出处：http://www.cnblogs.com/douzi2/
如果您觉得阅读本文对您有帮助，请点一下“推荐”按钮，您的“推荐”将是我最大的写作动力！欢迎各位转载，但是未经作者本人同意，转载文章之后必须在文章页面明显位置给出作者和原文连接，否则保留追究法律责任的权利

### 新项目出现类库Nuget引用失败提示，怎么办？

这么大一个项目出现80个Nuget异常，太恐怖了，怎么办？

![1697707735022](Vue进步.assets/1697707735022.png)

下面的命令解决,参考：https://learn.microsoft.com/zh-cn/dotnet/core/tools/sdk-errors/netsdk1064 启用长路径内容

```
dotnet build 生成项目及依赖项

dotnet publish 输出编译指定依赖，发送到指定目录
```

![1697708730579](Vue进步.assets/1697708730579.png)

然后设置长路径支持，https://learn.microsoft.com/zh-cn/nuget/reference/cli-reference/cli-ref-long-path，记得启用后重启一下电脑生效

![1697784720320](assets/1697784720320.png)

### 操作Excel的工具列表

推荐MiniExcel，https://github.com/mini-software/MiniExcel/blob/master/README.zh-CN.md  MiniExcel傻瓜操作Excel，高性能处理Excel查、写、填充数据工具

![1697687175400](Vue进步.assets/1697687175400.png)

![1697687201061](Vue进步.assets/1697687201061.png)

![1697687211780](Vue进步.assets/1697687211780.png)

https://github.com/mini-software/MiniExcel/blob/master/README.zh-CN.md  MiniExcel傻瓜操作Excel，高性能处理Excel查、写、填充数据工具



### 写日志的方法，写到本地的文件,logs

> 有时候使用Console.WriteLine日志多了，会发现不了问题，我只需要看我出现异常的问题的时候就可以使用，LoggerUtility.WriteLog($"{account.AccountAlias} - 处理数据异常", ex.Message); 

```c#
        
try
          {
                   string fileName = Path.GetFileName(shop.PicCodingUrl);
                    var response = QiNiuCloudHelper.FruugoUploadStreamAsync("sw-fruugo-feed", filePath, $"file/{fileName}");
                    if (response.Code == 200)
                    {
                        //七牛云有CDN缓存，所有每次文件更新，需要刷新一次
                        QiNiuCloudHelper.RefreshFile(new string[] { shop.PicCodingUrl, fileUrl });
                        //更新范本状态
                        service.UpdateListerPublish(listers, variaitonList);
                    }
                    else
                    {
                        LoggerUtility.WriteLog($"{account.AccountAlias} -上传七牛云失败", Newtonsoft.Json.JsonConvert.SerializeObject(response));
                    }
                
            }
            catch (Exception ex)
            {
                LoggerUtility.WriteLog($"{account.AccountAlias} - 处理数据异常", ex.Message);
            }

//日志写入类
 public class LoggerUtility
    {
        //在网站根目录下创建日志目录
        public static string path = AppDomain.CurrentDomain.BaseDirectory + "log";

        /// <summary>
        /// 写入日志
        /// </summary>
        /// <param name="type"></param>
        /// <param name="className"></param>
        /// <param name="content"></param>
        public static void WriteLog(string type, string content)
        {
            if (!Directory.Exists(path))//如果日志目录不存在就创建
            {
                Directory.CreateDirectory(path);
            }

            string time = DateTime.Now.ToString("yyyy-MM-dd HH:mm:ss.fff");//获取当前系统时间
            string filename = path + "/" + DateTime.Now.ToString("yyyy-MM-dd") + ".txt";//用日期对日志文件命名

            //创建或打开日志文件，向日志文件末尾追加记录
            StreamWriter mySw = File.AppendText(filename);

            //向日志文件写入内容
            string write_content = time + " " + type + ": " + content;
            mySw.WriteLine(write_content);

            //关闭日志文件
            mySw.Close();
        }

        /// <summary>
        /// 写入日志
        /// </summary>
        /// <param name="type"></param>
        /// <param name="className"></param>
        /// <param name="content"></param>
        public static void WriteLog(string type, string className, string content)
        {
            if (!Directory.Exists(path))//如果日志目录不存在就创建
            {
                Directory.CreateDirectory(path);
            }

            string time = DateTime.Now.ToString("yyyy-MM-dd HH:mm:ss.fff");//获取当前系统时间
            string filename = path + "/" + DateTime.Now.ToString("yyyy-MM-dd") + ".txt";//用日期对日志文件命名

            //创建或打开日志文件，向日志文件末尾追加记录
            StreamWriter mySw = File.AppendText(filename);

            //向日志文件写入内容
            string write_content = time + " " + type + " " + className + ": " + content;
            mySw.WriteLine(write_content);

            //关闭日志文件
            mySw.Close();
        }

        /**
        * 向日志文件写入调试信息
        * @param className 类名
        * @param content 写入内容
        */
        public static void Debug(string className, string content)
        {
            WriteLog("DEBUG", className, content);
        }

        /**
        * 向日志文件写入运行时信息
        * @param className 类名
        * @param content 写入内容
        */
        public static void Info(string className, string content)
        {
            WriteLog("INFO", className, content);
        }

        /**
        * 向日志文件写入出错信息
        * @param className 类名
        * @param content 写入内容
        */
        public static void Error(string className, string content)
        {
            WriteLog("ERROR", className, content);
        }
    }


补充知识，日志级别就是上文的string type
在日志文件中，"[INF]" 是日志级别的缩写，代表 "Information"（信息）。日志级别是描述日志条目重要性的方式，它帮助开发者了解发生了什么事情，以及是否需要采取行动。不同的日志级别代表了不同的严重性，"Information" 是其中之一。

比如2023-10-18 17:50:10.681 +08:00 [INF] Hello World!. 就是一条普通日志


以下是常见的日志级别，以及它们的意义：

Trace（跟踪）: 这是最详细的日志级别，通常只在需要诊断问题时使用。它提供了应用程序运行的非常详细的信息。

Debug（调试）: 这个级别用于那些对调试过程有帮助的信息，比如变量的值、程序执行的特定步骤等。

Information（信息）: "[INF]" 就是这个级别的标识。它用于记录应用程序的正常操作过程，比如启动和关闭、用户已经采取的操作等。这些信息对于了解系统的状态和发生的操作是有用的，但它们通常不是紧急的。

Warning（警告）: 这个级别表示可能的问题，这些问题不会立即影响操作，但可能需要未来的注意。

Error（错误）: 这个级别表示严重的问题，这些问题已经影响了应用程序的某些功能，需要立即解决。

Critical（严重）: 这个级别用于非常严重的错误，可能导致应用程序崩溃或终止。
```



### BulkUpdate导致服务器内存泄漏，Gc

**事件：**在做Walmart比对价格功能的时候，需要先将比对的明细表的数据放到中间表中，中间表数据很大，于是用limit进行查询，在循环中查询加BulkInset中间表,BulkUpdate原明细表状态，发现这样操作会导致腾讯云内存飚高，容器频繁重启

**解决**：上面已经是分析出原因了，其实**分析的方法是一段代码一段代码的注释，然后找到了这个问题**，注释一段代码，然后把代码上传到腾讯云运行，然后再注释代码，一段段排查，知道的问题的原因，解决方法：

> - **（最终方法）使用后台任务解决**：MQ里面不适合执行，周期性并发执行时间很长的任务，MQ适合马上执行完，然后马上走的任务，适合短期的任务，但是我这个任务可能会跑十几个小时，并且MQ在腾讯云上面对资源的限制很大，所以放到公司内部的Linux服务器上面跑更加适合

- **减少并发**：比如开始使用Parallel.ForEach，使用了50个并发，那么把并发降低到10个，看看会不会还这样。`实际尝试后发现也会崩溃，只是崩溃的时间晚了几个小时，所以减少线程的数量只是减缓了内存崩溃的时间，但是更改不了结果`

```
//密集CPU计算性能高，网络请求，加载文件不推荐，推荐算法，计算，渲染逻辑
ParallelOptions options = new ParallelOptions();
//支持50个并发执行
options.MaxDegreeOfParallelism = 50;
Parallel.ForEach(proSpus, options,  spu =>
{
 {
   SyncChangeProduct(spu, accounts);
 }
});
```



- **批量提交**：更改提交方式，不每次进行提交，20000提交一次，比如原方法是2000提交一次，提交的太密集了，与数据库交互的次数多导致内存消耗很大。`这个有用，崩溃的晚一些，还是会奔溃`

  ```c#
         if (insertTasksRunDetails.Count >= 20000)
                      {
                          await _adjustPriceTaskRunDetailRepository.BulkInsertAsync(insertTasksRunDetails.ToArray());
                          insertTasksRunDetails = new List<AdjustPriceTaskRunDetail>();
                      }
  ```

  

- **Select选择获取**：ToList()查询的时候不查所有的属性，只查指定的属性值，避免内存过大`这个有用，崩溃的晚一些，还是会奔溃`

  

- **主动GC回收**: 使用GC.Collect()显式回收所有的0代，1代，2代对象，`推荐使用GC.Collect(1);  回收0代和1代的对象`,比如下面的案例：`感觉没什么用，性能好像更加差了`

```c#
                insertTasksRunDetails = null;
                updateComparisonDetailList = null;
                GC.Collect(1); //回收0代与1代短期对象
```

**原代码**



### 将对象赋值为NULL值与GC.Collect垃圾回收有什么区别？



### 怎么使用GC进行垃圾回收（内存回收？）

在.NET中，垃圾回收（GC）使用三个代来组织对象：0代、1代和2代，所以一共有三代对象。

1. **0代（Generation 0）**: 这是新对象所在的代。当这些对象被实例化后，它们首先被分配到0代。0代是经常进行垃圾回收的地方，因为很多对象在实例化后很快就不再被使用。
2. **1代（Generation 1）**: 如果在0代的垃圾回收后对象仍然存活，它们会被提升到1代。1代的回收频率低于0代。这是一个缓冲区，因为如果对象在0代中生存下来，它们可能有中等的生命周期。
3. **2代（Generation 2）**: 如果对象在1代的垃圾回收后仍然存活，它们会被提升到2代。2代的回收频率是最低的。通常，长时间运行的应用程序中的长生命周期对象最终会在这个代中。

这种分代的方式允许GC以不同的频率回收不同的对象，依赖于对象的生命周期。0代的回收是最频繁的，而2代的回收是最不频繁的，因为回收2代涉及的对象更多，因此成本也更高。

- `GC.Collect(0);` 只会回收0代对象。
- `GC.Collect(1);` 会回收0代和1代对象，0代与1代主要存放短命期的对象。
- `GC.Collect(2);` 会回收0代、1代和2代对象。

### PostMan通过Get方法进行请求

参考链接：https://blog.csdn.net/weixin_40608713/article/details/113542836

![1696653424901](Vue进步.assets/1696653424901.png)

![1696653453758](Vue进步.assets/1696653453758.png)

### 从网上下载excel模板进行导出数据

```c#
        /// <summary>
        /// 导出
        /// </summary>
        /// <param name="dto"></param>
        /// <returns></returns>
        public IWorkbook Export(FyndiqListingFilterDto dto)
        {
            var restclient = new RestClient(@"https://qiniufile.yuetenginc.com/Fyndiq-Template.xlsx");
            restclient.Timeout = 0;
            var fileData = restclient.DownloadData(new RestRequest(Method.GET));
            IWorkbook workbook = new XSSFWorkbook(new MemoryStream(fileData));

            var query = BuildQuery(dto);
            var listings = query.Select(p => new FyndiqListing
            {
                Id = p.Id,
                ParentSku = p.ParentSku,
                MainImage = p.MainImage,
                Images = p.Images,
                Brand = p.Brand,
                ArticleId = p.ArticleId,
                ProductId = p.ProductId,
                Markets = p.Markets,
                Sku = p.Sku,
                WarehouseSku = p.WarehouseSku,
                TitleEn = p.TitleEn,
                DescriptionEn = p.DescriptionEn,
                PriceAmount = p.PriceAmount,
                PriceVatAmount = p.PriceVatAmount,
                PriceVatRate = p.PriceVatRate,
                Currency = p.Currency,
                Quantity = p.Quantity,
                OriginalPriceAmount = p.OriginalPriceAmount,
                OriginalVatAmount = p.OriginalVatAmount,
                OriginalVatRate = p.OriginalVatRate,
                Status = p.Status,
                FyndiqStatus = p.FyndiqStatus,
                VariationalProperties = p.VariationalProperties,
                Properties = p.Properties,
                Categories = p.Categories,
                ShippingTimeMax = p.ShippingTimeMax,
                ShippingTimeMin = p.ShippingTimeMin
            }).ToList();

            if (listings.Count == 0)
                return workbook;
            SetExtendProperty(listings); //给扩展属性赋值
            ISheet sheet = workbook.GetSheetAt(1);
            //表头
            int rowIndex = 0;
            IRow row0 = sheet.GetRow(rowIndex); rowIndex++;
            IRow row1 = sheet.GetRow(rowIndex); rowIndex++;
            IRow row2 = sheet.GetRow(rowIndex); rowIndex++;
            //对其方式
            ICellStyle cellStyle = workbook.CreateCellStyle();
            cellStyle.Alignment = HorizontalAlignment.Left; //左对齐
            Action<string, object> setCellValue = (name, value) =>
            {
                var cellIndex = row0.Cells.FindIndex(c => c.StringCellValue == name);
                if (cellIndex != -1)
                {
                    IRow row = sheet.GetRow(rowIndex);
                    ICell cell = row.GetCell(cellIndex);
                    if (cell == null)
                    {
                        cell = row.CreateCell(cellIndex);
                    }
                    cell.SetCellValue(value?.ToString());
                    cell.CellStyle = cellStyle;
                }
            };

            for (int i = 0; i < listings.Count; i++)
            {
                try
                {
                    var listing = listings[i];
                    //创建行
                    sheet.CreateRow(rowIndex);
                    setCellValue("*分类", listing.Categories.Replace("|", ","));
                    setCellValue("*产品名称（英语）", listing.TitleEn);
                    setCellValue("*产品描述（英语）", listing.DescriptionEn);
                    setCellValue("*发货时间（最小）", listing.ShippingTimeMin);
                    setCellValue("*发货时间（最大）", listing.ShippingTimeMax);
                    setCellValue("产品编号", listing.ProductId);
                    setCellValue("链接编码", listing.ArticleId);
                    setCellValue("品牌", listing.Brand);
                    setCellValue("*状态", listing.Status);
                    setCellValue("刊登状态", listing.Status);
                    setCellValue("在线状态", listing.FyndiqStatus);
                    setCellValue("父SKU", listing.ParentSku);
                    setCellValue("*SKU", listing.Sku);
                    setCellValue("*数量", listing.Quantity);
                    setCellValue("*价格", listing.PriceAmount + listing.PriceVatAmount);
                    setCellValue("*原价", listing.OriginalPriceAmount + listing.OriginalVatAmount);
                    setCellValue("*主图", listing.MainImage);
                    if (!string.IsNullOrEmpty(listing.VariationalProperties))
                        setCellValue("variational-property", listing.VariationalProperties.Replace("|", ","));
                    if (!string.IsNullOrEmpty(listing.Properties))
                    {
                        var items = Newtonsoft.Json.JsonConvert.DeserializeObject<List<ProperyItemDto>>(listing.Properties);
                        foreach (var item in items)
                        {
                            if (item.Name == "color")
                                setCellValue("property:color:sv-SE", item.Value);
                            if (item.Name == "size")
                                setCellValue("property:size:sv-SE", item.Value);
                            if (item.Name == "material")
                                setCellValue("property:material:sv-SE", item.Value);
                            if (item.Name == "pattern")
                                setCellValue("property:pattern:sv-SE", item.Value);
                        }
                    }
                    if (!string.IsNullOrEmpty(listing.Images))
                    {
                        var images = listing.Images.Split('|');
                        for (int k = 1; k <= images.Length; k++)
                        {
                            setCellValue($"图片{k + 1}", images[k - 1]);
                        }
                    }
                    rowIndex++;
                }
                catch (Exception ex)
                {

                }
            }

            return workbook;
        }


        /// <summary>
        /// 导出
        /// </summary>
        /// <param name=""></param>
        /// <returns></returns>
        [BailunAuthentication(LoginMode.Ignore)]
        [HttpPost("Export")]
        public IActionResult Export([FromBody] FyndiqListingFilterDto input)
        {
            var workbook = _fyndiqListingService.Export(input);
            var ms = new NpoiMemoryStream();
            ms.AllowClose = false;
            workbook.Write(ms);
            ms.Flush();
            ms.Seek(0, SeekOrigin.Begin);
            ms.AllowClose = true;
            return File(ms, "application/octet-stream", "Fyndiq范本导出");
        }
```



### EFCore双表，两表查询

```c#
// 单个属性关联
            var amazonQuery = from a in _amazonListingRepository.GetAll().IgnoreQueryFilters()
                              join b in _amazonListingVariationRepository.GetAll().IgnoreQueryFilters()
                              on a.Asin
                          equals b.ParentAsin
                              where
                                     a.IsDeleted == false &&
                                     b.IsDeleted == false && productSkus.Contains(b.SkuCode) && productSpus.Contains(b.ProCode)
                              select new ProductSizeExportSkuDto
                              {
                                  ShopId = a.AccountId,
                                  ShopName = a.AccountName,
                                  PlatformSku = b.SellerSku,
                                  ProductSku = b.SkuCode,
                                  Platform = "Amazon",
                                  Status = b.Status,
                                  PublishTime = b.ListingDate,
                              };



var amazonListings = await amazonQuery.ToListAsync();
//多个属性关联【注意匿名类型中的别名：Asin = a.Asin】
 var amazonQuery = from a in _amazonListingRepository.GetAll().IgnoreQueryFilters()
                              join b in _amazonListingVariationRepository.GetAll().IgnoreQueryFilters()
                             on new { Asin = a.Asin, SellerSku = a.SellerSku, AccountId = a.AccountId }
                            equals new { Asin = b.ParentAsin, SellerSku = b.ParentSku, AccountId = b.AccountId }
                              where
                                     a.IsDeleted == false &&
                                     b.IsDeleted == false && productSkus.Contains(b.SkuCode) && productSpus.Contains(b.ProCode)
                              select new ProductSizeExportSkuDto
                              {
                                  ShopId = a.AccountId,
                                  ShopName = a.AccountName,
                                  PlatformSku = b.SellerSku,
                                  ProductSku = b.SkuCode,
                                  Platform = "Amazon",
                                  Status = b.Status,
                                  PublishTime = b.ListingDate,
                              };



            var amazonListings = await amazonQuery.ToListAsync();

```



### SqlSugar也可以直接查询DataTable

![1691134138686](Vue进步.assets/1691134138686.png)

### SqlSugar也可以直接执行Sql

task项目

![1691132848883](Vue进步.assets/1691132848883.png)

### 字符并不是只能转换成CSV，也可以转换成Excel

Task项目

![1691132466469](Vue进步.assets/1691132466469.png)

### 直接在后端通过Excel导出的方法

```c#
获取List集合逻辑略...              
using (var memoryStream = new MemoryStream())
                {
                    using (var writer = new StreamWriter(memoryStream, Encoding.UTF8, 1024, true))
                    {
                        var headers = "SPU,平台SKU,系统SKU,供应商原始值,偏码值,新映射的系统SKU,系统SKU数量,易仓账号,平台,站点,速牛店铺,开发员,主销售,附属销售,开发时间,销售状态";
                        await writer.WriteLineAsync(headers);

                        foreach (var item in exportList)
                        {
                            var line = $"{item.ProductSpu},{item.PlatformSku},{item.ProductSku},{item.Size},{item.OffsetSize},{item.OfficeProductSku},{item.ProductSkuCount},{item.EccangAccount},{item.Platform},{item.Site},{item.ShopName},{item.DeveloperUserName},{item.MainSaleUserName},{item.SubSaleUserName},{item.DeveloperCreationTime},{item.SaleStatus}";
                            await writer.WriteLineAsync(line);
                        }
                    }
                    var username = HttpContextHelper.Current?.User?.GetUserName();
                    var fileName = $"SKU关系映射{username}_{DatetimeHelper.GetDateNowTimeStamp()}.xlsx";
               //publishListing七牛云模板项目
                    var fileResult = await SWQiNiuFileHelper.UploadStreamAsync(fileName, memoryStream.ToArray());
                 //Skums项目
                        var fileResult = await Sunway.Skums.Domain.Util.QiNiuCloudHelper.UploadStreamAsync(fileName, memoryStream.ToArray());

                    if (fileResult.Code == 200)
                    {
                        //publishListing七牛云模板项目
                        string fileUrl = SWQiNiuFileHelper.CreateProductsUrl(fileName);
                        //skums项目
     string fileUrl =Sunway.Skums.Domain.Util.QiNiuCloudHelper.CreateProductsUrl(fileName); ;

                        Console.WriteLine($"下载地址:{fileUrl}");
                    }
                }
```



### VS2022设置下一条语句BUG

**问题描述**：当出现需要更改程序运行顺序， 设置下一条执行语句的时候 就有可能提示：Object reference not set to an instance of an object 未将对象引用设置到对象的实例

**解决**：按照顺序执行，不要拖动顺序就可以解决（我耍小聪明，想看看拖动左边的箭头会不会有这个问题，发现也是会有这个问题）

**疑问**：为什么不是所有的拖动执行顺序的场景会有这种情况，我遇到的这种场景有

![1689740417959](Vue进步.assets/1689740417959.png)

### Microsoft.Bcl程序集无法找到

**问题提示**：System.IO.FileNotFoundException: Could not load file or assembly 'Microsoft.Bcl.AsyncInterfaces, Version=7.0.0.0, Culture=neutral, PublicKeyToken=cc7b13ffcd2ddd51'. The system cannot find the file specified.（System.IO.FileNotFoundException:无法加载文件或程序集Microsoft.Bcl。AsyncInterfaces, Version=7.0.0.0, Culture=neutral, PublicKeyToken=cc7b13ffcd2ddd51'。系统找不到指定的文件。）

**解决方案**：NetGet安装Microsoft.Bcl.AsyncInterfaces，然后重新发布程序部署就可以了

**扩展阅读**：https://github.com/microsoft/playwright-dotnet/issues/1531

### **Obsolete属性启用标志**

```c#
[Obsolete("新版本已弃用，请使用ServiceName")]
		public string ApiName
		{
			get
			{
				return ConfigManagerConf.GetValue("ServiceDiscovery:ServiceName");
			}
		}
```

### Notepad如何替换换行符

1.显示所有不可见字符

![1688008015888](Vue进步.assets/1688008015888.png)

2.Ctrl+H 替换字符，替换过程中可以用\r,\n进行分别替换，`CRLF`（`\r\n`), \r是回车，\n是换行

![1688008091758](Vue进步.assets/1688008091758.png)

### Startup.cs创建的WebApi有这个文件的方法

使用Net5创建Web Api的方法

![1681876139188](Vue进步.assets/1681876139188.png)

### 网站不能上网了--代理

![1681369428059](Vue进步.assets/1681369428059.png)

### 清理浏览器缓存的方法

先F12打开开发者工具，然后看下图

![1679623024532](Vue进步.assets/1679623024532.png)

### CSRedisCore.pdb 包含查找模块 CSRedisCore.dll 的源文件所需的调试信息

设置一下解决：

![1679387621563](Vue进步.assets/1679387621563.png)

### 查看失效链接/网站的存档内容，这个真太牛了

[失效链接，查看存档内容，保存指定链接的内容](https://web.archive.org/web/20210401000000*/https://www.cnblogs.com/xiaobaozi-95/p/9328948.html)

[维基百科关于失效链接的说明](https://zh.wikipedia.org/zh-hans/Wikipedia:%E5%A4%B1%E6%95%88%E9%93%BE%E6%8E%A5)

### 调试情况下，结束所有的程序，多线程，方便测试，退出

- System.Environment.Exit(0);   这是最彻底的退出方式，不管什么[线程](http://wenwen.sogou.com/s/?w=线程&ch=w.search.intlink)都被强制退出，并且保留vs的执行状态

### 当前不会命中断点，还没有为该文档加载任何符号

![1697193430063](Vue进步.assets/1697193430063.png)

### 开发环境下面不进去断点的解决

![1678094816731](Vue进步.assets/1678094816731.png)

### 输入逗号,看方法重载的方法

![1677817680104](Vue进步.assets/1677817680104.png)

### Cron表达式的使用

![1677554376768](Vue进步.assets/1677554376768.png)

[Cron表达式](https://help.aliyun.com/document_detail/160666.html)

### EF中的子查询的功能，牛

```c#
            // 子查询：侵权词InfringeKeyword
            if (!string.IsNullOrWhiteSpace(input.InfringeKeyword))
            {
                var platformSourceIds = _infringeKeywordScanRepository.Query().AsNoTracking().Where(u => u.InfringeKeyword == input.InfringeKeyword&&u.PlatformSource=="Walmart Lister").Select(u => u.PlatformSourceId);
                query = query.Where(p => platformSourceIds.Contains(p.Id));
}
```



### C#10新的特性：等于号=简洁方法输入

```
public class Foo
{
    public IBar Bar{get;}
    public Foo(IBar bar) =>Bar = bar;
}
```



### 打小例子的时候，可以使用方法注入的方式

**[FromServices]**：直接注入使用，注意一个注入服务，一个[FromServices]

![img](https://ask.qcloudimg.com/http-save/7234372/kakzqt82y8.png?imageView2/2/w/1620)

### &看内存地址判断对象是否相同

- 也可以用GetHashCode()方法进行检查

  ![1675309217166](Vue进步.assets/1675309217166.png)

- 也可以用Object.ReferenceEquals方法

  ![1675309012452](Vue进步.assets/1675309012452.png)

- &符号进行判断

![1675307681786](Vue进步.assets/1675307681786.png)

### 无法解析类型“生命周期”的服务。 IndexModel” 

**问题背景：**

打生命周期的案例时，有个方法类是不需要接口抽象的，我要直接net IndexModel().Show()方法验证逻辑

**异常提示**:

 Unable to resolve service for type '生命周期.IndexModel'

无法解析类型“生命周期”的服务。 IndexModel” 

**解决方案**：

1. builder.Services.AddScoped<IndexModel>(); 需要在services中注入这个对象
2. 然后直接在构造函数中注入这个对象**(不需要新建接口进行继承)**：    IndexModel indexModel....        _indexModel = indexModel;

### 名称相同的属性路由'Show'必须有相同的模板

**问题背景：**

新建生命周期案例，有两个路由名称重复，方法名称也重复了

**异常提示：**

Attribute routes with the same name 'Show' must have the same template:

名称相同的属性路由'Show'必须有相同的模板

**解决方案：**

1. 更改路由名称：由**WeatherForecastController2** 更改为 **WeatherForecast2Controller**
2. 更改路由规格由：    **[Route("[controller]")]** 更改为 **[Route("[controller]/[action]")]**
3. 更改方法名称：由    **[HttpGet(Name ="Show")]**  更改为 **[HttpGet("Show")]**

### QPS与TPS

**TPS**：

Transactions Per Second，**意思是每秒事务数,一个事务可能会有多个QPS**。 可以理解为单个接口的请求代表一个事务，这个接口的请求里面可以包含了多个查询更新操作，或者接口内部又去调用接口，就是包含了多个QPS

**QPS**：

Queries Per Second，**意思是每秒查询并发率**。

**QPS 基本类似于 TPS，但是不同的是，对于一个事务访问，会形成一个 “ T ”；但一次 " T " 中，可能产生多次对服务器的请求，服务器对这些请求，就可计入 QPS 之中。**

**区别**：

如果是对一个查询接口压测，但是这个接口内部不需要请求其它接口，那么 TPS就是等于QPS，否则，TPS不等于QPS




### 程序包管理器下载

用命令下载快几十倍

```c#
-> Install-Package Newtonsoft.Json
```



### 字符串中用双引号 \（右斜杠）

```c#
            string strJson = "ABC\""; //如果这个值有换行符，用@艾特符号，就会提示错误，要把换行符去除，或者想办法将@艾特符号与换行符兼容
            Console.Write(strJson); //  ABC" 
```



### 请求接口返回不规则的Json解决



```c#
// 方法1使用dynamic接受，然后逐步读取里面的值，使用场景：整体全是不规则的
var response= JsonConvert.DeserializeObject<dynamic>(res);
 方法1有可能直接遍历出结果，有可能遍历不出结果，如果遍历不出方法可以把返回值复制出来，用单独的小例子慢慢分析字符串进行转换，慢慢试出来,还有一种情况是，如果返回值只有一个data,然后即使里面是数组也不会转换成数组，需要手动绑定两列冗余的属性值，比如增加message,Success这两列属性值：  response = $"{{ \n\"message\": \"Success\",\n\"data\":{response}}}"; //不合法Json转合法

// 【使用这个方法】方法2使用实体类接受，使用场景：整体是规则的只是里面某个对象是不规则的单独用dynamic声明进行接受
var response= JsonConvert.DeserializeObject<BulkSoldOuResponseDto>(res);
 // 类：PayloadErrors 是类 ：BulkSoldOuResponseDto 的子类
 public class PayloadErrors
 {
    [JsonProperty("bulk_payload")]
    public dynamic BulkPayload { get; set; } //试了一下用object但是不可以
 }
// 不规则json
            "bulk_payload":[
                {
                    "1":[
                        {
                            "id":[
                                "Must be a valid UUID 4 string"
                            ]
                        }
                    ],
                    "3":[
                        {
                            "id":[
                                "Must be a valid UUID 4 string"
                            ]
                        }
                    ],
                    "4":[
                        {
                            "body":[
                                "required field"
                            ],
                            "id":[
                                "Must be a valid UUID 4 string"
                            ]
                        }
                    ]
                }
            ]
```



### 异步方法没有用await调用引发的对象为空的BUG



方法内容：在断点的那一行抛出，抛出未将对象引用设置到对象的实例bug

![1666946456375](Vue进步.assets/1666946456375.png)

引用方法（错误）：没有加await标签

![1666946520316](Vue进步.assets/1666946520316.png)

引用方法（正确）：加了await标签（用foreach只是方便一些，也可以用spus.ForEach(async spu=>)）

![1666946500279](Vue进步.assets/1666946500279.png)

### 日期比对Compare

```c#
    //Compare 返回值：相等 0，小于 -1，大于 1
    DateTime dt1=new DateTime(2022,8,1);
    DateTime dt2=new DateTime(2022,10,1);
    Console.WriteLine(DateTime.Compare(dt2,dt1)); //1
    Console.WriteLine(DateTime.Compare(dt1,dt2)); //-1

```



### Random取随机数

```c#
    Random rd=new Random();
    int number = rd.Next(0,2); 
    Console.WriteLine(number);//取值：0,1，取不到2值
```



### ApiController：导致绑定参数File参数校验失败

ApiController ：有自动参数校验的功能，今天在控制器里面配置了这个参数，导致控制器File文件一直校验失败，在控制器顶部剔除了这个特性正常了

[扩展阅读](https://qa.1r1g.com/sf/ask/4658209181/)

### 匿名实体类临时用起来很香

```c#
//匿名类型 
var oldProducts = await _productRepository.Query().AsNoTracking().Where(l => allProductSpus.Contains(l.ProductSpu)).Select(u => new { ProductSpu = u.ProductSpu, CommonFinshStatuts = u.CommonFinshStatuts }).ToListAsync();
//读取匿名类型
var noProducts = allProductSpus.Except(oldProducts.Select(u => u.ProductSpu)).ToList();

```



### 如何预热WebApi，提高初始执行的效率

创建一个预热的API方法，然后在startUp中注入到默认启动的方法列表中，这样就可以在启动服务的时候自动调用某个API方法，达到初始化管道方法的目的

[如何预热Web API，减少初次执行时间](https://blog.csdn.net/sD7O95O/article/details/122227191)

### Tuple数据类型（元组）：自定义多个返回值类型

```c#
//普通的方法只能返回一个参数，如果想要返回多个参数，可以用实体类封装，用ref,用out方法进行返回，不过有一定的局限性，比如用实体类封装需要单独创建实体类，用ref ,out 就不能使用await 异步调用了，并且如果用字典Dictiary的话<key,value>的类型中，key是不允许重复的
//Tuple永远的神
//【经验】可以用Tuple.Item1,Tuple.Item2调用元组里面的值
public Tuple<string,object,int> GetStudent(int id)
{
var student=studentContext.Query().Where(u=>u.id==id).FirstOrDefault();
 return new Tuple<string,object,int>(student.name,student,student.age);

}
```



### vs循环中调用构造函数注入方法

需要规避在循环中调用构造函数注入的方法，可能会极大的影响到系统的性能，而且调用的方法越多，对性能影响的越大

例子：通过批量导入更新功能更新2700条数据，调用三个通过构造函数注入的方法，导入时间8分钟，将构造函数注入的方法更改为本地调用，5秒钟就循环结束了

![1662689051349](Vue进步.assets/1662689051349.png)

### vs怎么开多个消费者（并发）

```c#
//方法1：使用cmd命令
dotnet run --project ./Sunway.Skums.Service.csproj  --urls="https://*:5005"
dotnet run --project ./Sunway.Skums.Service.csproj  --urls="https://*:5006"
dotnet run --project ./Sunway.Skums.Service.csproj  --urls="https://*:5007"

//方法2(推荐使用3)
开一个vs 然后用PostMan开启多个请求
//方法3
磁盘复制两个vs 执行访问
//方法4 ：本地开多个消费者(推荐使用1)
使用MQ调用方法，黑窗口创建调用MQ的方法，不断的往里面发送，参考辉哥的脚本（还有裕文的环境变量参数，可以利用环境变量指定调用的MQ）
start cmd /k "D: && cd D:\Code\Sunway\Sunway.MQ.Consumer && dotnet run"
start cmd /k "D: && cd D:\Code\Sunway\Sunway.MQ.Consumer && dotnet run"
start cmd /k "D: && cd D:\Code\Sunway\Sunway.MQ.Consumer && dotnet run"
start cmd /k "D: && cd D:\Code\Sunway\Sunway.MQ.Consumer && dotnet run"
start cmd /k "D: && cd D:\Code\Sunway\Sunway.MQ.Consumer && dotnet run"
start cmd /k "D: && cd D:\Code\Sunway\Sunway.MQ.Consumer && dotnet run"
exit
//方法5：线上环境开多个消费者(推荐使用2)
开多个pod节点，结合服务器资源更新数量的个数
```



### source.dot.net看源码

**1.搜索方法**（通过关键字搜索的，也可以搜索：string.cs 关键字代表搜索类）

![1662520882680](Vue进步.assets/1662520882680.png)

![1662521110680](Vue进步.assets/1662521110680.png)

![1662521195972](Vue进步.assets/1662521195972.png)

### dynamic匿名类型序#列化接口返回值妙用

![1660103787043](Vue进步.assets/1660103787043.png)

```c#
//Json值
"product_address": [
                        {
                            "sp_product_address": "https://detail.1688.com/offer/587573801117.html"
                        }
                    ],
```

![1660011167166](Vue进步.assets/1660011167166.png)

### 转移字符转义& 转义字符 &amp;

```c#
//使用&amp;替换&
& 转义字符 &amp;
```

[参考链接](https://blog.csdn.net/feng8403000/article/details/120049745)

### 复制粘贴一段长的文本，到代码中

下面的方法不绝对，我刚刚试了一下它又不行了，有触发条件才能重现这个需求

![1658222602206](Vue进步.assets/1658222602206.png)

### 指定端口号启动后端服务

```
dotnet run --project ./Sunway.Skums.Service.csproj --urls="https://*:5004"
```



### Net计时方法

```c#
 //秒表方法一:
            Stopwatch sw = new Stopwatch();
            sw.Start();
            for (int i = 0; i < 10000; i++)
            {
 
            }
            sw.Stop();
// 结果以毫秒为单位
            MessageBox.Show(sw.Elapsed.TotalMilliseconds.ToString());
 
            //当前时间减去开始时间方法二:
            DateTime begintime = DateTime.Now;
            for (int i = 0; i < 10000; i++)
            {
 
            }
            TimeSpan ts = DateTime.Now.Subtract(begintime);
            MessageBox.Show(ts.TotalMilliseconds.ToString());
```



### Count(*)如何做优化

删除部分索引试试效果，有时候有索引也会有问题

[参考解决方案](https://blog.csdn.net/leeta521/article/details/119378768)

![1657513781228](Vue进步.assets/1657513781228.png)

### ASCII字符代码表

![1.png](https://img.php.cn/upload/image/158/313/978/1650266038660675.png)

### 新的东西：entity、bo、vo、po、dto、pojo区别

[entity、bo、vo、po、dto、pojo区别](https://baijiahao.baidu.com/s?id=1731226378554209465&wfr=spider&for=pc)

### 即时窗口里面可以声明变量进行使用

下面的temp就是我声明的变量可以在局部方法中进行使用（可以做一些骚操作）

![1656477610334](Vue进步.assets/1656477610334.png)

### 当Consul不动了的时候，用鼠标点一下，然后按Enter

consul不动了的时候，只会展示红框框住的那部分

![1656405736081](Vue进步.assets/1656405736081.png)

### 当出现Cmd黑窗口的时候，鼠标点击一下窗口

鼠标点击一下窗口，窗口就不会往下面继续执行，我们可以检查打印的日志，然后再按enter按键就可以马上继续执行

### 当Cmd黑窗口出现异常的时候，通过error关键字筛查

![1655886468937](Vue进步.assets/1655886468937.png)

### uow.Complete()事务不起作用的方法

```c#
//方法1 已验证
UpdateCommitAsync
//方法2 已验证
BulkInsertAsync
//方法3(已验证)
BulkUpdate
```



### DateTime使用指定时间初始化

```c#
//日期之间的比较也可以用这个      
DateTime now=new DateTime(2022,6,21);
```



### ++在前与++在后的区别

**++在前**：先运算后赋值【优先用这个】

**在后++**：表示先赋值后运算

### Token的学习

[Token登陆验证机制的原理及实现](https://www.jb51.net/article/230951.htm)

token比Session更好用，因为token可以支持分布式的系统，可以支持跨域请求，安全性也比session更高，Session因为是保存在服务器的内存中，所以比较耗内存，而且不支持跨域请求，安全性也不高（可能遭到CSRE攻击）,如果是分布式部署，那么还要想办法去处理Session信息共享的问题（比如可以用redis解决）

[JWT全面解读和详细使用步骤](https://www.jb51.net/article/230951.htm)

JWT是针对于Http无状态情况，对于用户权限认证的方法（不需要用户每次请求都发送用户名与密码信息给我），他的优点是通过token令牌就可以获取接口的信任，同时他的缺点也是如果有人伪造了用户令牌，那么接口同样也可以通过授权，默认情况下JWT是未加密的，任何人都可以解读其内容，因此如果一些敏感信息不要存放在此（我试了一下公司的token头部是没有加密的，中部有部分加密，有部分没有加密，尾部是全部加密的，如果需要加密可以生成原始 Token 以后，可以用一个秘钥再用算法再加密一次）

- 头部

  ```
  {"alg":"RS256","kid":"ad44ef585d03407f8ac05dad208dcbaa","typ":"JWT"}
  ```

  JWT第三部分是签名。是这样生成的，首先需要指定一个secret（秘钥），该secret仅仅保存在服务器中，保证不能让其他用户知道。然后使用Header指定的算法对Header和Payload进行计算，然后就得出一个签名哈希。也就是Signature。

  **那么Application Server如何进行验证呢？可以利用JWT前两段，用同一套哈希算法和同一个secret计算一个签名值，然后把计算出来的签名值和收到的JWT第三段比较，如果相同则认证通过。**

  **获取token:**算出签名以后，把 Header、Payload、Signature 三个部分拼成一个字符串，每个部分之间用"点"（`.`）分隔，就可以返回给用户

  ```c#
  // 验证的代码
      public Claims parseJWT(String jwtString) {
          return Jwts.parser().setSigningKey(secretKey)
                  .parseClaimsJws(jwtString)
                  .getBody();
      }
  ```

  [这个讲JWT三部分的内容讲的很清晰](https://www.zhihu.com/question/485758060/answer/2257869896)

[IdentityServer4实战 - 谈谈 JWT 的安全策略](https://www.cnblogs.com/stulzq/p/9678501.html)

identityServer有两种加密token的方法，一种是通过授权鉴权服务器，一种是通过采用非对称加密算法（Jwt token）进行加密，解密

### 通过GC回收内存

skums.ZipFile 方法

            GC.Collect();
            GC.Collect(1);
![1655457315533](Vue进步.assets/1655457315533.png)

### Net5与Net6与NetCore,NetFramework之间的区别

- .NetCore与Net5+

跨平台

开源

支持安卓，IOS,Linux,macOS

高性能

NetGet包的形式分发程序集

最小API部署

所有应用程序后端服务都使用相同的 .NET 库很方便共享代码

- Net FrameWork

闭源

仅支持windows系统

本地程序集引用

体积臃肿

性能低

不在维护

单一整体部署

![img](https://p9.itc.cn/q_70/images03/20220410/6f60d7836b934efa910c526332d6d71f.png)

### Net Core有时候卡顿严重，没有反应

问题： cmd窗口卡住没有，程序执行没有反应了

![1655260458466](Vue进步.assets/1655260458466.png)

解决：按回车键，就是用了有副作用，它可能会不间断的卡在那里，然后一直要按回车键

### Any可以用来看对象中哪些属性存在值

```c#
//判断字典值是否有值，在整个对象List集合中
if (importDtoList.Any(p => !string.IsNullOrWhiteSpace(p.ZZJG)))
            {
                 //如果字典值有值的话，再读取转换字典的逻辑
            }
```
### 拖动代码顺序导致，CS$<>8__locals1 是 null。问题提示

问题描述：

拖动代码顺序出现：CS$<>8__locals1 是 null ，未将对象引用设置到对象的实例错误提示

解决方法：

不拖动代码运行，就可以解决

扩展链接：

[NullReferenceException: CS$<>8__locals0 was null](https://www.codebye.com/nullreferenceexception-cs8__locals0-was-null.html)

### Skywalking介绍

[Skywalking介绍](https://www.jianshu.com/p/ffa7ddcda4ab)

Skywaling是用来监测微服务的性能，粒度可以细化到方法的级别，还可以监测出各个服务之间的依赖关系

### VS打开诊断工具

![1653634043545](Vue进步.assets/1653634043545.png)

### Mysql保存不了负数值的问题

今天用Ef保存数据发现一直保存不了负数值，开始在后端保存保存不了，然后想到我去数据库试一试，发现数据库也更新不了-1负数值，有个设置的问题：

![1653361614402](Vue进步.assets/1653361614402.png)



### 登录腾讯云容器服务

![1651818943494](Vue进步.assets/1651818943494.png)

选择容器

![1651819356063](Vue进步.assets/1651819356063.png)

选择集群

![1651818949432](Vue进步.assets/1651818949432.png)



点击日志

![1651819320967](Vue进步.assets/1651819320967.png)

SKUMS后端的日志

![1651897910731](Vue进步.assets/1651897910731.png)

Web前端

admin web

![1652340902587](Vue进步.assets/1652340902587.png)

查看镜像版本

![1652340804521](Vue进步.assets/1652340804521.png)

### 注入services.AddRedis出现的未注册的IRedisManager提示



错误提示：

Castle.MicroKernel.Handlers。不能创建组件Sunway.Skums.Application.Lister。ListerProductPropertyService'，因为它有依赖需要满足。



“Sunway.Skums.Application.Lister。ListerProductPropertyService'正在等待以下依赖项:

——服务”Sunway.ServiceFabric.Redis。未注册的IRedisManager。

![1651199297117](Vue进步.assets/1651199297117.png)



解决方案：

![1651199284865](Vue进步.assets/1651199284865.png)

### 向上取整，向下取整

Math.Round：四舍六入五取整

Math.Ceiling：向上取整，只要有小数都加1

Math.Floor：向下取整，总是舍去小数

### 通过Select可以new 声明匿名类型

```c#
//分组，选择
var userCategory = listUserDistributions.Select(u => new { CategoryId = u.Key, List = u.ToList() }).ToList();
```



###  出现两个属性名称，无法保存

有个表设计中有两个ID,保存的时候不确定是把哪个ID保存到数据库

错误提示：

Ambiguous match found 模糊匹配发现

解决方案：

在实体类中去除其中一个ID

### 深度拷贝的方法错误提示

**提示**：

导入失败，原因：The property 'Int32 Count' has no 'set' accessor
Parameter name: member

**中文：**

 属性'Int32 Count'没有'set'访问器r\nParameter name:成员

```c#
//引用代码：
ExpressionDeepCloneObject.TransExpClone(oldProductProperty);

//改进： 灵魂代码：item.CanWrite
  // 1.如果标记了NotMapped，代表不需要保存到数据库 2.如果不可写也直接到下一个
                    if (item.GetCustomAttribute(typeof(NotMappedAttribute)) !=null&&!item.CanWrite)
                    {
                        continue;
                    }
```



### 按照100条一行，进行拆分数组

如果是一个1000行的数组，那么会拆分成十个，有100个值的集合

```
     var pushProductGroups = pushProductList.ToArray().SplitArray(100);
```



### Vs调试启动编辑并且继续的功能

更改代码后并不能即使作用到源码中，可以更改后，配合即时窗口的功能进行验证程序

**调试-》选项-》去除勾选：** 去除勾选！是去除勾选！是去除勾选

![1648794071982](Vue进步.assets/1648794071982.png)



![1648794134007](Vue进步.assets/1648794134007.png)

[VS调试启动"编辑并继续"功能](https://www.cnblogs.com/liuhuimh/p/11346775.html)

### 如果标记了特性代表不需要转换

```c#
//反射获取  
foreach (var item in typeof(TIn).GetProperties())
 {
   // 如果标记了NotMapped，代表不需要保存到数据库
  if (item.GetCustomAttribute(typeof(NotMappedAttribute)) !=null)
    {
    continue;
     }
 }
//使用特性 
 [NotMappedAttribute]
  public int MyProperty { get; set; }
 //声明特性   
 public class NotMappedAttribute : Attribute 
 {
 
 }
```



### AsNoTracking去除跟踪，不读取缓存

![1648453868022](Vue进步.assets/1648453868022.png)

### 通过数据链接对象查看具体的数据库链接

![1648452980755](Vue进步.assets/1648452980755.png)

### 实体类的继承，如果没有表里面的属性

```
 : FullAuditedEntity<long>  --有全部的属性
 :Entity<long>    --不需要其他的属性
```



### 忽略这个属性的序列化：JsonIgnoreAttribute



用户后端向前端返回参数的时候，忽略一些不需要返回的属性，可能是哪些大的属性值

![1648190388502](Vue进步.assets/1648190388502.png)

### ABP框架更新指定单个属性的方式

并且熟悉了一波表达式类目树的入参方式

UpdateFormQueryAsync

![1648084604906](Vue进步.assets/1648084604906.png)

### ABP框架获取所有List的信息

​              GetAllListAsync ABP框架异步原生方法

```
   result = await _productSkuECCInfoRepository.GetAllListAsync(p => p.SkuCode == skuCode);

```

 

### Git查看提交的历史记录

![1647837221878](Vue进步.assets/1647837221878.png)

### 显示状态栏，Typora

![1647403737605](Vue进步.assets/1647403737605.png)

### 更改GitLab代码管理路径到新的Coding代码仓库



第一步：选择项目

![1647400386428](Vue进步.assets/1647400386428.png)

第二步：点击代码仓库箭头

![1647400447159](Vue进步.assets/1647400447159.png)

第三步：点击项目名称

![1647400532006](Vue进步.assets/1647400532006.png)

第四步：复制下载链接

![1647400594607](Vue进步.assets/1647400594607.png)

第五步：替换Git代码管理工具链接

![1647400713275](Vue进步.assets/1647400713275.png)

第六步：粘贴新的Coding链接

![1647400783032](Vue进步.assets/1647400783032.png)

第七步：检验成果

![1647400854100](Vue进步.assets/1647400854100.png)

第八步：拉取成功

![1647401023318](Vue进步.assets/1647401023318.png)

### ABP框架表连接查询

```c#
        data = await _productSkuRepository.Query().AsNoTracking()
            .Join(_productSkuStateRepository.Query().AsNoTracking(), x => x.Id, y => y.SkuId, (x, y) => new
            {
                x,y
            })
            .Join(_productExploitRepository.Query().AsNoTracking(), x => x.x.Id, y => y.SkuId, (x, y) => new
            {
               x,y
            })
            .Select(x => new WebProductSkuDto
            {
                Id = x.x.x.Id,
                ProductId = x.x.x.ProductId,
                CategoryBailunId = x.x.x.CategoryBailunId,
                CategoryShopifyId = x.x.x.ShopifyCategoryId ?? 0,
                CategoryBailunName = x.x.x.CategoryBailunName,
                Color = x.x.x.Color,
                Ctitle = x.x.x.Ctitle,
                CreationTime = x.x.x.CreationTime.ToString("yyyy-MM-dd HH:ss"),
                      ....
            })
            .Where(x => ids.Contains(x.Id))
            .OrderBy(p=>p.Sku)
            .ToListAsync();
```
### WebApi去掉烦人的FormBoby

[ASP.NET Core API 接收参数去掉烦人的 [FromBody] ](https://www.cnblogs.com/zhan520g/p/10292903.html)

### 流程图工具

processon，亿图，visio

### 删除数据库的属性名称时，需要对应更新实体类的属性

EFCore框架的特性，使用对象操作实体类，数据库属性的新增，减少。需要跟着实体类走，不能在正式使用的时候新增，减少属性的数量

![1645516354573](Vue进步.assets/1645516354573.png)

### 上传模板到七牛云

DataAnalysis.ImportExcel

![1645091019987](Vue进步.assets/1645091019987.png)

### Vs开发使用的技巧

开发新功能的时候把菜单的模块都悬浮在上面，因为做起来难免会一直复制，粘贴

![1645070465820](Vue进步.assets/1645070465820.png)

###  Parallel并行执行循环

适用场景：方法执行不需要顺序，循环的执行没有顺序，执行内容耗时的操作（耗时要抵消系统创建委托的性能） （skip,tolist也有并行的办法）

```c#
 static void Main(string[] args)
        {
            //每次五个线程一起跑
            Parallel.For(0, 1000000, (i) => { Thread.Sleep(3000); Console.WriteLine(i); }); 
            Console.ReadKey();
            //每次三个线程一起跑，下面的主线程代码会等待Parallel.Invoke执行完后再进行执行
            Parallel.Invoke(() => Show1(), () => Show2(), () => Show3());  
            for (int i = 0; i < 10000; i++)
            {
                Console.WriteLine($"【主线程】{Thread.CurrentThread.ManagedThreadId},数字：{i}");
                Thread.Sleep(1000);

            }
            Console.Read();
        }
        public static void Show1()
        {
            for (int i = 0; i < 10000; i++)
            {
                Console.WriteLine($"【子线程1】{Thread.CurrentThread.ManagedThreadId},数字：{i}");
                Thread.Sleep(1000);

            }
        }
        public static void Show2()
        {
            for (int i = 0; i < 10000; i++)
            {
                Console.WriteLine($"【子线程2】{Thread.CurrentThread.ManagedThreadId},数字：{i}");
                Thread.Sleep(1000);

            }
        }
        public static void Show3()
        {
            for (int i = 0; i < 10000; i++)
            {
                Console.WriteLine($"【子线程3】{Thread.CurrentThread.ManagedThreadId},数字：{i}");
                Thread.Sleep(1000);

            }
        }
```

**Paraller.ForEach的语法**

![1643363511653](Vue进步.assets/1643363511653.png)

遍历System.Data.DataRow需要使用Cast转换：   

```
 Parallel.ForEach(dt.Rows.Cast<System.Data.DataRow>(), row => { Console.WriteLine(1); });
```



### 安装Java JDK环境

[安装Java JDK环境](https://blog.csdn.net/weixin_31159999/article/details/114789851)

导出数据库文档： 设置好了数据库路径，表的信息没有展示出来，无法导出

官方文档1:https://www.oschina.net/p/chiner

官方文档2：https://www.yuque.com/chiner/docs/manual

配置数据库：jdbc:mysql://localhost:3306/mydb?characterEncoding=UTF-8&useSSL=false&useUnicode=true&serverTimezone=UTC

用户名密码：uid=root;pwd=123;

![1643267130259](Vue进步.assets/1643267130259.png)

### Efcore注入的过滤器：

_productRepository.GetAll().IgnoreQueryFilters()   :IgnoreQueryFilters 忽略过滤器

[配置软删除过滤器](https://blog.csdn.net/weixin_30888707/article/details/95342137)

[配置软删除过滤器（这个人讲的详细一些）](https://www.cnblogs.com/willick/p/13358580.html)

### 多表链接用&&

```

            var query2 = from a in _productRepository.GetAll()
                         join b in _productPropertyRepository.GetAll() on a.ProductSpu equals b.ProductSpu &&a.pname=b.pname
                 
                         select a.Id;
```



### 拼接表达式目录树条件

```c#
   //拼接条件             
input.SelfCategoryList.ForEach(item =>
                {
                    query = query.Where(l => l.SelfCategoryList.Contains(item));
                });
 //上下两种方法都可以，不过上面的更优雅            
Expression<Func<ProductPropertyDto, bool>> predicate = null;
            input.SelfCategoryList.ForEach(item =>
            {
                if (predicate == null)
                {
                    predicate = l => l.SelfCategoryList.Contains(item);
                }
                else
                {
                    predicate = predicate.AndExpression<ProductPropertyDto>(l => l.SelfCategoryList.Contains(item));
                }
            });
            query = query.Where(predicate);
```
### 忽略Dto中指定的属性（一般是大字段）返回前端：

 **[JsonIgnoreAttribute]** ：添加这个特性就可以

```c#
        [JsonIgnoreAttribute]
        public List<ProductPropertySpecific> Specifics
        {
            get
            {
                if (!SpecificJson.IsNullOrEmpty())
                {
                    return JsonHelper.FromJsonList<ProductPropertySpecific>(SpecificJson);
                }
                return new List<ProductPropertySpecific>();
            }
        }
```



### 拼接字符串的快捷方式

```c#
//JoinAsString("|")是拼接字符串的快捷方式，这个方法是ABP框架的方法,命名空间：Abp.Collections.Extensions.EnumerableExtensions
lister.TitleAdjective = listStr.JoinAsString("|");

```



### Efcore使用异步方法优化

![1642840043702](Vue进步.assets/1642840043702.png)

### 表达式目录树赋值

```c#
//查询的属性：单个
Expression<Func<ListerCategoryMapping, bool>> where = (l => l.Id == dto.Id);
//更新的属性:数组
Expression<Func<ListerCategoryMapping, object>>[] updatedProperties = new Expression<Func<ListerCategoryMapping, object>>[]
{
                    l=> l.CategoryId, l=> l.CategoryName, l=> l.ToCategoryId, l=> l.ToCategoryName, l=> l.ToCategory, l=> l.LastModifierUserName, l=> l.LastModifierUserId, l=> l.LastModificationTime
  };
```



### 提示无权限访问地址

后端接口方法去除特性：        [BailunAuthentication(LoginMode.Enforce)]

![1642642420590](Vue进步.assets/1642642420590.png)

### 使用NPOI生成的文件打不开

注意下面打开读取流的配置(open,create)，用两个FileStream进行操作就不会出现

```c#
           IWorkbook workbook = null;
           //打开读取流:FileMode.Open, FileAccess.Read
            using (var fs = new FileStream(path, FileMode.Open, FileAccess.Read))
            {
                try
                {
                    workbook = new XSSFWorkbook(fs);
                }
                catch
                {
                    workbook = new HSSFWorkbook(fs);
                }
            }
           //创建写入流:FileMode.Create, FileAccess.Write
            using (FileStream fout = new FileStream(path, FileMode.Create, FileAccess.Write))
            {
                workbook.Write(fout);//写入文件
                fout.Flush();
                workbook = null;
                fout.Close();
            }
```



### 使用NPOI设置单元格背景颜色

[C#使用NPOI设置Excel单元格背景颜色(xls,xlsx)](https://blog.csdn.net/qq_34452824/article/details/119379559?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.pc_relevant_default&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.pc_relevant_default&utm_relevant_index=2)

```

ICellStyle style = workbook.CreateCellStyle();
 
//使用NPOI已经有的颜色创建
style.FillForegroundColor = NPOI.HSSF.Util.HSSFColor.Yellow.Index;
style.FillPattern = FillPattern.SolidForeground;
```



### 自定义读取appsettings中的配置信息

可以在start.up里面进行注入，进行测试访问是否有效

```c#
//读取类属性配置**（注意冒号）
Console.WriteLine($"TeachInfo.Id:{_configuration["TeachInfo:Id"]}");
```

如果是需要在静态类中进行读取，可以通过这个方法进行调用 [ASP.NET Core-静态类的访问配置](https://www.it1352.com/1945270.html)

如果是使用封装好的读取方法（注意使用的命名空间，如果有冲突的话使用这个，如果是其他的可能读取不到配置的信息）：

![1642391955621](Vue进步.assets/1642391955621.png)

### 通过Swagger上传文件参数不存在

其实通过postMan也可以上传文件

下面代码配置不了很奇怪 [[Net Core小技巧 - 使用Swagger上传文件 ]()](https://www.cnblogs.com/Erik_Xu/p/8904854.html)，链接地址：

![1642156206884](Vue进步.assets/1642156206884.png)

### NuGet这里安装包

![1642155694393](Vue进步.assets/1642155694393.png)

**Install-Package** Furion.Template.Api  ： 下载指定的项目

**dotnet new --install Furion.Template.Api::4.8.4.5**  下载成功 (乱码的话，在cmd命令下面执行)，不带版本号总是安装最新的版本，这里安装的是在vs里面官方目录的模板（`Visual Studio` 集成），真的牛逼

![1672800675997](Vue进步.assets/1672800675997.png)

### Nuget安装检查到包降级问题解决

**错误提示：**

```c#
NU1605	检测到包降级: Swashbuckle.AspNetCore 从 6.4.0 降级到 6.2.3。直接从项目引用包以选择不同版本。 
 ORMWebApi -> Furion 4.6.6 -> Swashbuckle.AspNetCore (>= 6.4.0) 
 ORMWebApi -> Swashbuckle.AspNetCore (>= 6.2.3)	ORMWebApi	D:\坚果云同步\我的坚果云\Web进步\ORMWebApi\ORMWebApi.csproj	1	

```

**问题原因：**

Furion框架 4.6.6版本中引用的Swashbuckle.AspNetCore类库版本是6.4.0的版本，但是我本地的Swashbuckle.AspNetCore版本是6.2.3所以下载框架的代码就会出现问题

**解决方案：**

把 官方Swashbuckle.AspNetCore类库在本地的解决方案中的版本6.2.3 升级到 6.4.0

![1673401638306](Vue进步.assets/1673401638306.png)

### 公司的项目管理Nuget

**5. 创建自己的 NuGet Server。**

如果你是一个大公司的程序设计人员，你设计了解决某一问题的 解决方案即一个良好的会被同事频繁使用的动态链接库中NuGet.Demo.dll，你将它放置在公司共享的某个磁盘文件夹下，并email通知大家最新的更新等。

这个dll一直有你维护和更新，以往大家也就会得到你的email通知, 根据自己的需求，到上文说的文件夹下copy并删除当前旧的dll.

如今我们可以通过NuGet Server提供的手段，帮我们更加自动和人性化的简化并更加方便这一工作场景：

那么，如何创建这个自己公司的NuGet Server。

我们新创建一个 工程：Infoshare.Projects.

然后通过上面的方法，打开Package Manager Console, 并输入如下命令并回车：

...... 

**PM> Install-Package NuGet.Server** 

...... 

这时，你会发现工程的Refernces 下添加了一个NuGet.Core的引用，它将为我们实现这一方案。

同时，添加了两个文件夹： “DataServices”  ；“Packages” 和若干文件 以及一个WebForm "Default.aspx"，如下图所示：

![img](https://images.cnblogs.com/cnblogs_com/dlonghow/%E6%9C%AC%E7%AB%99%E9%9A%8F%E7%AC%94%E6%8F%90%E4%BE%9B/6.png) 

运行，并查看Default.aspx页面：

 ![img](https://images.cnblogs.com/cnblogs_com/dlonghow/%E6%9C%AC%E7%AB%99%E9%9A%8F%E7%AC%94%E6%8F%90%E4%BE%9B/7.png)

 点击here，进入Packages页面， 你应该可以看到Package下包含的nuget包:

 ![img](https://images.cnblogs.com/cnblogs_com/dlonghow/%E6%9C%AC%E7%AB%99%E9%9A%8F%E7%AC%94%E6%8F%90%E4%BE%9B/8.png)

但是，如上面显示并没有任何Feed 记录。这是因为在刚才的Infoshare.Projects的Packages文件夹下，并没有任何nuget包。

我们将上面通过NuGet的Packages Manager Cnsole添加的 NuGet 包，放到当前这个Infoshare.Projects的Packages文件夹下，

添加完成后如下图： 

![img](https://images.cnblogs.com/cnblogs_com/dlonghow/%E6%9C%AC%E7%AB%99%E9%9A%8F%E7%AC%94%E6%8F%90%E4%BE%9B/9.png) 

在运行查看：

 ![img](https://images.cnblogs.com/cnblogs_com/dlonghow/%E6%9C%AC%E7%AB%99%E9%9A%8F%E7%AC%94%E6%8F%90%E4%BE%9B/10.png)

我们在项目中创建一个Class Library 命名为：**Infoshare.Core**, 将作为公司普用的关键dll，通过NuGet Server进行管理。

 ![img](https://images.cnblogs.com/cnblogs_com/dlonghow/%E6%9C%AC%E7%AB%99%E9%9A%8F%E7%AC%94%E6%8F%90%E4%BE%9B/12.png)

 现在的问题是如何才能生成类似Newtonsoft.Json可供NuGet.Server使用的nuget包，即将自己的公司的dll打包成nuget包？

 这要用到NuGet提供的工具，你需要到 http://nuget.codeplex.com/上面下载：

![img](https://images.cnblogs.com/cnblogs_com/dlonghow/%E6%9C%AC%E7%AB%99%E9%9A%8F%E7%AC%94%E6%8F%90%E4%BE%9B/11.png) 

 下载nuget.exe 后，你可以使用下面的方式进行打包：

 使用nuget.exe pack将工程文件打包， （命令："nuget pack"）：

 将下载的NuGet.exe放到Infoshare.Core的工程文件夹下：

 ![img](https://images.cnblogs.com/cnblogs_com/dlonghow/%E6%9C%AC%E7%AB%99%E9%9A%8F%E7%AC%94%E6%8F%90%E4%BE%9B/13.png)

执行如下命令：

 ![img](https://images.cnblogs.com/cnblogs_com/dlonghow/%E6%9C%AC%E7%AB%99%E9%9A%8F%E7%AC%94%E6%8F%90%E4%BE%9B/14.png)

 将生成的 **Infoshare.Core.1.0.0.0.nupkg** 文件和相应的Infoshare.Core.dll 拷贝到 Infoshare.Projects下的Packages文件夹下：

 ![img](https://images.cnblogs.com/cnblogs_com/dlonghow/%E6%9C%AC%E7%AB%99%E9%9A%8F%E7%AC%94%E6%8F%90%E4%BE%9B/15.png)

 在运行查看defalut.aspx:

![img](https://images.cnblogs.com/cnblogs_com/dlonghow/%E6%9C%AC%E7%AB%99%E9%9A%8F%E7%AC%94%E6%8F%90%E4%BE%9B/16.png) 

那么如何像NuGet Server那样，可以通过“Add Library Package Reference...”，弹出Library Reference 搜索和安装对话框？

这就要向 NuGet 添加更多程序包源， 即添加你的NuGet Server程序包源。 

通过“Tools”（工具）|“Library Package Manager”（库程序包管理器）|“Package Manager Settings”（程序包管理器设置）菜单选项，单击“Package Sources”（程序包源）节点，

![img](https://images.cnblogs.com/cnblogs_com/dlonghow/%E6%9C%AC%E7%AB%99%E9%9A%8F%E7%AC%94%E6%8F%90%E4%BE%9B/19.png)

今天突然看到网上又一遍更加系统和全面的讲解阅读： [使用NuGet 管理项目库](http://msdn.microsoft.com/zh-cn/magazine/hh547106.aspx)

转载自：http://www.cnblogs.com/Dlonghow/archive/2012/03/16/2399993.html

### **上传七流云的逻辑**

![1642149466284](Vue进步.assets/1642149466284.png)

### 多条件判断的逻辑

**存在可以：**用多个包含Contains去判断存在里面的数据是可以的

**不存在：**用多个包含Contains去筛选不存在里面的数据，是筛选不到的

```c#
//代码粘贴就可以用           

List<Plat> plat = new List<Plat>() {
              new Plat(){  pt="平A", xl="下1",vl="ABC"},
              new Plat(){  pt="平A", xl="下1",vl="ABC"},
              new Plat(){  pt="平B", xl="下1",vl="123"},
              new Plat(){  pt="平B", xl="下1",vl="123"},
              new Plat(){  pt="平C", xl="下1",vl="DDD"},
            };
            //判断公共属性逻辑：看看是不是全平台为公共属性
            var commDetails = plat.GroupBy(l => new { l.xl, l.vl }).ToList();
            if (commDetails.Count == 1)  //全平台公共属性
            {


            }
            else if (commDetails.Where(l => l.Count() > 1).Any()) // 大于1代表明细中部分平台有公共属性
            {
                //细化平台公共属性：加入平台分组
                var commPlatformDetails = plat.GroupBy(l => new { l.pt, l.xl, l.vl }).Where(l => l.Count() > 1).ToList();
                // 【多个包含筛选信息，是可以的】筛选公共的明细信息，结果：4条数据
                var detailsA = plat.Where(l => commPlatformDetails.Select(c => c.Key.pt).Contains(l.pt) && commPlatformDetails.Select(c => c.Key.xl).Contains(l.xl) && commPlatformDetails.Select(c => c.Key.vl).Contains(l.vl)).ToList();
                
                //【多个不包含筛选信息是不行的，结果为空】从主明细中减去这部分公共的属性,结果：0条数据
                var detailsB = plat.Where(l => !commPlatformDetails.Select(c => c.Key.pt).Contains(l.pt) && !commPlatformDetails.Select(c => c.Key.xl).Contains(l.xl) && !commPlatformDetails.Select(c => c.Key.vl).Contains(l.vl)).ToList();
            }

public class Plat
{

    public string pt { get; set; }
    public string xl { get; set; }
    public string vl { get; set; }

}
```
### VS内存耗尽的提示

---------------------------
Microsoft Visual Studio

.NET 调试服务已耗尽内存并导致调试终止。如果此问题定期出现，请参阅 https://aka.ms/DotNetDebugServicesOutOfMemory 了解如何排查内存问题。



**接口中引用的位置包括实现的实例：**

如果有很多继承的话，滑动到最底下的那个就是实现

![1637496015723](C:\Users\12987\AppData\Roaming\Typora\typora-user-images\1637496015723.png)

**WebApi的入参：**

```c#
一、方法的入参
//1.默认 获取拼接的参数，简单参数，可以去掉[FromUri]标记
public string GetStudent(string name,string age){  } 
      //上面的方法等于 [FromUri]特性   
void Action([FromUri] int id, [FromUri] string name)
 //2.复杂参数(参数个数多，url值过大时（get的Url有8190字节的长度）) 用 [FromBody]
public string SaveStudent([FromBody]Student student){  } 
 //实际使用【记住，Long[]这种类型的数组，也是需要FromBody的入参的】
 1.[FromBody] long[] listerIds  --数组   [HttpPost("SubmitPubish")]
 2.long  listerid  --单个参数    [HttpGet("ListerRemove")]
 3.long  listerid  --单个参数  [HttpGet("GetFyndiqListerVariations")] 见本文档：Vue使用/Get请求传参：

二、方法的特性
		[NonAction]  //不被外部应用程序调用，标记了特性，swagger里面不会显示这个方法
		[AcceptVerbs("GET","POST")] //表示既可以Get请求，又可以Post请求
        [HttpPost("SaveFyndiqOrder")]  //新增操作，括号中定义方法别名
        [HttpDelete("DeleteFyndiqOrder")] //删除操作，括号中定义方法别名
	    [HttpPut("PutFyndiqOrder")] //更新操作，括号中定义方法别名
        [HttpGet("GetFyndiqOrder")] //查询操作，括号中定义方法别名
```

数组传参与FromBody传参与数组的关系

![1642414614013](Vue进步.assets/1642414614013.png)

### 批量更新数据的方法

小数据量，100条以内用：**UpdateCommitAsync** ，大于100条用：**BulkUpdate**

```c#
//使用【UpdateCommitAsync】碰到一种情况，数据量达到5500条的时候，直接卡死，程序debug不动了
// 【UpdateCommitAsync】 不传更新属性名称的话，代表全部更新
   _ListerCategorySpecificMappingDetailRepository.UpdateCommitAsync(
       updateDetails.ToArray(), u => u.ShowType, u => u.ListValue, u => u.LastModifierUserName, u => u.LastModifierUserId, u => u.LastModificationTime);
//改为使用【BulkUpdate】，使用时注意C.Id，一定要传，这个是主键!!!,不然就会有提示：发生错误，无法找到或解决主节点（An error occured, no primary could be found or resolved）
                                       _ListerCategorySpecificMappingDetailRepository.BulkUpdate(updateDetails.ToArray(), P =>
                    {

                        P.BatchSize = 100;
                        P.ColumnInputExpression = c => new
                        {
                            c.Id,
                            c.ShowType,
                            c.ListValue,
                            c.LastModifierUserName,
                            c.LastModifierUserId,
                            c.LastModificationTime

                        };
                    });

//根据条件直接更新数据结果，不需要传实体对象的值
   _productPropertyPushEcRepository.BulkUpdate(new[] { new ProductPropertyPushEc { EcangSyncStatus = 0, ProductSkuID = productSkuID } }, P =>
                 {
                     P.BatchSize = 100;
                     P.ColumnPrimaryKeyExpression = o => o.ProductSkuID;
                     P.ColumnInputExpression = c => new
                     {
                         c.EcangSyncStatus

                     };
                 });
//也可以根据Dto转换进行更新
       _adjustPriceTaskRunDetailRepository.BulkUpdate(taskDetailList.Select(u => new AdjustPriceTaskRunDetail() { Id = u.Id, SyncStatus = "已保存核查表" }).ToArray(), P =>
                    {
                        P.BatchSize = 100;
                        P.ColumnInputExpression = c => new
                        {
                            c.Id,
                            c.SyncStatus
                        };
                    });
//用sqlsugar的update方法可以批量，并且获取指定的属性
```



### 调试WebApi的方法

![1642039137380](Vue进步.assets/1642039137380.png)

### abp框架5.1，在方法传递的参数为datatable时过慢问题解决方案

应该是拦截器问题，方法上加上特性 [DisableValidation] 就可以关闭了

此问题应该会在6.1版本完善

[ https://github.com/aspnetboilerplate/aspnetboilerplate/issues/5875](https://github.com/aspnetboilerplate/aspnetboilerplate/issues/5875)

错误提示：

Method arguments are not valid! See ValidationErrors for details 

扩展博客：

[ ABP方法校验](https://blog.csdn.net/sinat_27535209/article/details/102835488)

### 访问特性的解析LoginMode

现在用网关做了限制，外部是访问不了我们的站点的

```
    public enum LoginMode
    {
        Enforce = 0,   --实施，执行，现在用的最多的
        Ignore = 1,  --忽略忽视
        MockAdmin = 2, --模拟管理员权限
        PublicEnforce = 3  --公共执行的方法
    }
```



### Lambad表达式使用多个OrderBy的方法

用OrderBy() 和 OrderByDescending()两个方法排序，如果后面两项排序顺序相同，就可以使用 ThenBy()和 ThenByDescending  ()方法继续排序。

```c#
//LINQ 按多个字段排序（orderby、thenby）
 var racers =ArrayList.OrderBy(r => r.Country). ThenBy(r => r.LastName). ThenBy(r => r.FirstName). Take(10);

//LINQ 按多个字段排序（OrderByDescending、ThenByDescending  ）
 var racers = ArrayList.OrderByDescending(r => r.Country).ThenByDescending(r => r.LastName).ThenByDescending(r => r.FirstName).ToList();
```



### 导出数据时Excel与CSV的区别

**CSV格式比Excel格式具备的优势：**

- **节省内存：**CSV是纯文本文件，支持追加模式写入，节省内存。Excel是结构复杂的二进制文件，只支持一次性写入，较费内存。
- **行数限制：**CSV的文件行数没有限制，在实际项目中我们已输出过上千万行的CSV文件；32位系统下Excel单个Sheet最多支持65535行。
- **编辑简单：**CSV是纯文本文件，可以使用任何文本编辑器进行编辑，因此可以在Linux终端下对其进行修改。Excel是二进制文件，目前已知的编辑工具有Office，OpenOffice，WPS，都为GUI工具，不支持在终端下编辑。

### 后端返回两个对象的快捷方法

```
        /// <summary>
        /// 产品性质和环球币种（web）
        /// </summary>
        /// <returns></returns>
        [HttpGet("GetcharacterAndcurrency")]
        public async Task<object> GetcharacterAndcurrency()
        {
            var characters = await _categoryAppService.GetSkumsCharacters();
            var currencys = await _categoryAppService.GetCategoryCurrencys();
            return new { characters, currencys };
        }
```



### 链接Redis的类

![1641348262632](Vue进步.assets/1641348262632.png)

### 入参的参数过大导致的故障记录

**错误原因：** **因为入参过大导致的问题，有四万多的数据作为参数，**可能会产生方法通讯的问题，感觉不对，引用对象应该只是传参数的地址进去，这个方法

**解决方法：**把代码移动到外部进行处理，问题解决：

```c#
  // 移动前代码：判断平台品类是否合法(可以只读取子级类目进行优化)
                    //         if (!_listerCommPlatformrCategoryService.CheckCategoryId(mapping.PlatformName, mapping.SiteName, mapping.PlatformCategoryId, amazonCategorys, ebayCategorys,
                    //warmartCategorys))
                    //         {
                    //             continue;
                    //         }
//移动后代码
                    var platform = (ListingPlatform)Enum.Parse(typeof(ListingPlatform), mapping.PlatformName);
                    if (ListingPlatform.Amazon == platform && !amazonCategorys.Where(m => m.SiteCode == mapping.SiteName && m.CategoryId == mapping.PlatformCategoryId && m.IsLeaf == true).Any())
                    {
                        continue;
                    }
                    else if (ListingPlatform.Walmart == platform && !warmartCategorys.Where(m => m.Id == mapping.PlatformCategoryId).Any())
                    {
                        continue;
                    }
                    else if (ListingPlatform.Ebay == platform && !ebayCategorys.Where(m => m.Site == mapping.SiteName && m.CategoryID == mapping.PlatformCategoryId && m.IsLeaf == true).Any())
                    {
                        continue;
                    }

```

**错误提示：** http请求错误， System.Net.Http.HttpRequestException:当发送请求时发生错误。System.IO.IOException:服务器返回一个无效或无法识别的响应。

   Error making http request, exception: System.Net.Http.HttpRequestException: An error occurred while sending the request. ---> System.IO.IOException: The server returned an invalid or unrecognized response.
   at System.Net.Http.HttpConnection.FillAsync()
   at System.Net.Http.HttpConnection.ReadNextResponseHeaderLineAsync(Boolean foldedHeadersAllowed)
   at System.Net.Http.HttpConnection.SendAsyncCore(HttpRequestMessage request, CancellationToken cancellationToken)
   --- End of inner exception stack trace ---
   at System.Net.Http.HttpConnection.SendAsyncCore(HttpRequestMessage request, CancellationToken cancellationToken)
   at System.Net.Http.HttpConnectionPool.SendWithNtConnectionAuthAsync(HttpConnection connection, HttpRequestMessage request, Boolean doRequestAuth, CancellationToken cancellationToken)
   at System.Net.Http.HttpConnectionPool.SendWithRetryAsync(HttpRequestMessage request, Boolean doRequestAuth, CancellationToken cancellationToken)
   at System.Net.Http.HttpClient.FinishSendAsyncBuffered(Task`1 sendTask, HttpRequestMessage request, CancellationTokenSource cts, Boolean disposeCts)
   at Ocelot.Requester.HttpClientHttpRequester.GetResponse(DownstreamContext context)

**错误代码：**

```
// 判断平台品类是否合法(可以只读取子级类目进行优化)，入参的参数最大的有4万多导致的问题，但是第一次读取商品信息又没有问题，是第二次读取Ebay的类目导致的问题，匪夷所思
 if (!_listerCommPlatformrCategoryService.CheckCategoryId(mapping.PlatformName, mapping.SiteName, mapping.PlatformCategoryId, amazonCategorys, ebayCategorys,
           warmartCategorys))
 {
       continue;
 }
 
 
```

### 指定方法某个参数赋值

命名参数：方法名：参数  给指定的方法参数赋值

![1640862927485](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1640862927485.png)

### 获取当前操作用户信息

```
 	//创建
    CreateUserName = HttpContextHelper.Current?.User?.GetUserName(),
     CreatorUserId =  HttpContextHelper.Current?.User?.GetUid().ToString() ?? "0",
     CreationTime = DateTime.Now 
    //更新
     item.LastModifierUserName = HttpContextHelper.Current?.User?.GetUserName();
     item.LastModifierUserId = HttpContextHelper.Current?.User?.GetUid().ToString() ?? "0";
     item.LastModificationTime = DateTime.Now;
     // 删除
     item.DeletionTime = DateTime.Now;
     item.DeleterUserId = HttpContextHelper.Current?.User?.GetUid().ToString() ?? "0";
     item.DeleteUserName = HttpContextHelper.Current?.User?.GetUserName();
     item.IsDeleted = true;
 
          _listerKeywordTemplateRepository.UpdateCommit(mapping.ToArray(), u => u.IsDeleted, u => u.DeleteUserName, u => u.DeleterUserId, u => u.DeletionTime);
```



### EfCore，Linq，Lambad通用方法

Append ：两个序列的值累加到A序列，返回新的序列

Concat：把A，B两个序列放到一起，形成C序列，返回新的序列

Cast ：把序列类型，转换成另一个序列，返回转换后的序列

Contains ：判断序列中是否包含指定的元素，返回true或者false

Count：返回序列的函数，返回Int类型的值

DefaultIfEmpty ：如果序列为空，则返回序列的默认值，返回序列的为空的默认值（这个返回的是list，firstorDefault返回的是单个的实体）

ElementAtOrDefault：入参序列的index索引，返回单个序列的值，如果为空则返回单个序列的默认值

Empty：返回序列的空值，单个的对象（就像firstOrDefault）

Join:对两个序列进行表关联的查询，返回匹配的序列

OfType：根据指定的类型过滤序列的值，是不是可以返回string,int类型的值

Repeat：生成重复的序列值，输入个数，就复制出一个list集合的序列进行返回

**FirstOrDefault:**返回序列的第一个元素, 如果序列不包含任何元素, 则返回默认值。

```c#
//返回单个实体的方法
_fyndiqListerRepository.Query().Where(s => s.Contains("color")).FirstOrDefault()
```

**Intersect**：交集 

```
ListA.Intersect(ListB).ToList();
```

**Union：并集**

```
ListA.Union(ListB).ToList();  
```

### **Any：**判断

```
_fyndiqListerRepository.Query().Where(l => l.AccountId ==AccountId ).Any()
```

### **EfCore：**事务

```
    private readonly IUnitOfWorkManager _unitOfWorkManager;
    
    using (var uow = _unitOfWorkManager.Begin(TransactionScopeOption.RequiresNew))
    {
       uow.Complete();
    }

```

### **ForEach：**循环

```
pageList.ForEach(p =>
{
    if (s.Id==p.Id)
    {
        p.Categories = string.Join(",", s.PathCategorie);
    }
});
```

### **Where：** 筛选

```c#
//基础的筛选
Where(p => p.IsDeleted == false)
//多属性商品里面的Skucode如果全部不是空的，就不用循环了
variations.Where(v => string.IsNullOrEmpty(v.SkuCode)).Count() == 0
```

### **OrderByDescending** ：排序

降序排序

```
OrderByDescending(p => p.Id)
```

### **OrderBy ** ：排序

顺序排序

```
OrderBy(p => p.Id)
```

### Select ：投影

映射数据到新的对象,也可以映射到一个单独的属性 

```c#
 // 映射到对象（注意不要用{} 大括号把对象包起来了）
Select(p => new FyndiqListerDto
            {
                Id = p.Id,
                AccountId = p.AccountId
             }
        )
   // 映射到字符串， 把集合中的一列(s.Value)转换成一个字符串
List<string> listStr=listSpecific.Select(s => s.Value);
```

### **string.Join** ：拼接

用指定的字符把集合拼接成字符串

```
string strValue= string.Join("|",listSpecific.Select(s => s.Value));
```

### Except：差异(A-B)

![1654133808254](Vue进步.assets/1654133808254.png)



### Intersect：交集

# 工具使用



## 经验

### 【置顶】编码的经验

编写代码的时候问自己的一些问题：

1.如果往表里面加一个属性，我的程序改动大吗，怎么让改动最小化？

2.如果有失败的情况，如何在后台进行弥补？

### 【置顶】记录关键字经验

1.比如ServiceProvider 开始直接丢怎么用，概念放到最下面，金字塔原理，结论先行

### Google浏览器页面变图片

打开网页——输入ctrl+shift+i或F12——输入ctrl+shift+p­——输入full回车——浏览器会自动保存整个网页为png图片

这个方法目前2023年8月31日已经失效了

### 编写接口分享规范

方法描述：根据中文名称获取walmart属性映射名称
请求路径：ListerColorMapping/GetMappingNames  
请求类型：HttpGet
入参格式：IEnumerable<string> nameCns  字符串集合
返回对象：List<ListerColorMappingDetailMappingNameDto>
返回格式：{ mappingName,nameCn}
请求参数： ['军绿','黑色2','米色']
返回值：

{
  "result": [
    {
      "mappingName": "Green",
      "nameCn": "军绿"
    },
    {
      "mappingName": "Beige",
      "nameCn": "米色"
    }
  ],
  "statusCode": 200,
  "message": null
}

### 正式库Redis链接地址

  "redis": {
    "HostName": "192.168.255.235",
    "Port": "6379",
    "Password": "",
    "Defaultdatabase": 2
  },

### 返回对象

OperationOutput

### typepo卡顿

1.不用微软原生的输入法输入，使用搜狗输入法进行输入试试

2. 开源代码模式

//调试的需求

//接口的权限请求
    public enum LoginMode
    {
        Enforce = 0,
        Ignore = 1,   --忽略这个接口的权限验证
        MockAdmin = 2,
        PublicEnforce = 3
    }

### 沃尔玛请求的路径：

--公网路径
https://admin.shengweiinc.com/WalmartApi/Common/GetWalmartCategorys
--内部路径
http://192.168.255.133:8000/WalmartApi/Common/GetWalmartCategorys



### 天翼云

1298775182@qq.com
~WR1436062434

//正式的网关服务
http://192.168.254.42:8500 --正式的

### 为什么要学习Docker

公司后面发展壮大，自建机房 肯定是要自己搞的，自己会搞，不求人

### 公司现有的项目

项目就是用的腾讯云一整套部署体系，集群+k8s+docker

### 观看开源经验

1.开源社区就像分类不同的工具，有前端，后端，移动开发，桌面开发，前沿技术，人工智能模块

//编码的含义： 7位，15位
商品编码/仓库SKU：指的一个款的商品，即绿色的 RBC0002
SKU/仓库编码：指的一款商品里面的不同颜色的每个个体。RBC000200001     
平台编码/平台SKU/刊登SKU：使用SKU上架到平台的编码，是用仓库SKU+#+账号简称+随机2-3位数字  RBC0002#WAOkk-03
SPU集合：商品的集合，比如苹果13是一个商品，红色的外壳，白色的外壳，黑色的外壳，就是SPU集合
listing :listing就是我们说的链接，一条listing有多个SKU
表的命名：1.有关销售平台的表，以listing结尾的，都是在售表，lister结尾都是范本表  2.t_bi_平台_业务


//其他编码
仓库编码 proCode  --界面展示的
平台编码 skuCode

### 控制器的返回值

控制器注意下，只有列表才返回VueResult这个结果，其他就正常发货List<FyndiqListerDto>就行了（搞明白是什么意思）  --只有查询列表才需要返回VueResult这个结果（isSecceed，messsage,result(object)）,其他的查询直接return实体类，或者List<实体类>的集合

//设置最后一列固定：fixed="right"  ，前面一列如果不设置宽度可以动态改变宽度
<el-table-column label="操作" align="center" width="180" fixed="right" />
	

### appsettings配置解析

Development  开发环境
Production 正式环境

### 引用图片 

 RBA0773 这个有图片

### &&符号的含义

当 && 连接两个值，第一个值为true,则返回第二个值；若第一个值为false,则直接返回一个值；

Undefined、Null、 String、NaN、false、0 转换为布尔值的时候，都为false；其他所有值都为true;
//||符号的含义
与|| ,碰到真，就返回；如果两者都为假，就返回后面的一个假【

### 刊登模板的

### 转换数据库实体到Dto实体的

**方法1**：通过Select方法

_fruugoOrderRepository.Query().OrderByDescending(p => p.Id).Select(p => new FruugoOrderListDto
{
    Id = p.Id,
    ShopId = p.ShopId
	
}
 **方法2**：AutoMap

通过AutoMap对象的双向映射 Map，     

需要在dto实体中配置：**[AutoMap(typeof(FyndiqLister))]** ，比如需要转换FyndiqListerDto，那么就在实体类中配置，引用的命名空间是：**Abp.AutoMapper**
  var dtoList= Mapper.Map<List<FyndiqListerDto>>(query);

 //如果转换出现问题会有具体的属性提示，比如CreatorOrganizeId， An exception occurred while reading a database value for property 'FyndiqLister.CreatorOrganizeId'. The expected type was 'System.String' but the actual value was of type 'System.Int64'.

### 弹出弹窗

this.$message.warning('Fyndiq平台限制只允许选择两个属性')

--设置 <el-input>的宽度
 style="width:30%;"

--form表单的结构
<el-form-item></el-form-item>

  

### 快速执行Net程序

dotnet run

### Mysql字符占字数

在mysql中，如果是latin1字符集下，一个中文汉字占2个字节数；如果是utf8字符集下，一个中文汉字占3个字节数；如果是gbk字符集下，一个中文汉字占2个字节数。
show variables like '%character%';  --查看数据库字符集 （目前都是utf8,一个中文汉字占3个字节数）
show variables like 'collation%';  --查看服务器字符集（目前都是utf8,一个中文汉字占3个字节数）

### Navicat 常用的数据类型

int
decimal
longtext
datetime
longtext
Navicat

### Navicat添加注释

使用：Ctrl+ /


--共享盘符信息
访问路径：\\192.168.11.246\sunway
用户名：it111
密码：it@#654321

--数据库密码
server=gz-cdb-ezvn7csl.sql.tencentcdb.com;Port=58892;Database=bailun_bltpro;Uid=root;Pwd=shengye2021
host=gz-cdb-ezvn7csl.sql.tencentcdb.com;port=58892;database=bailun_bltpro;uid=root;pwd=shengye2021;Convert Zero Datetime=True;

host=localhost;port=3306;database=mydb;uid=root;pwd=123;Convert Zero Datetime=True;

gz-cdb-ezvn7csl.sql.tencentcdb.com --服务器
58892 端口
root 用户名 shengye2021 密码
bailun_bltpro 数据库

### 关键单词

Vat  --增值税
Tax  --税
Shipping --运费，地址，国家相关
Customer --顾客

--typora目录快捷键 ctrl+1,2,3...
 ctrl+1,2,3  对应一级目录，二级目录，三级目录

###  集群

 集群通常用来提升速度、可靠性。
 高可用性集群：一般是指当集群中有某个节点失效的情况下，其上的任务会自动转移到其他正常的节点上。还指可以将集群中的某节点进行离线维护再上线，该过程并不影响整个集群的运行。

（Nginx是负载均衡集群）负载均衡集群：负载均衡集群运行时，一般通过一个或者多个前端负载均衡器，将工作负载分发到后端的一组服务器上，从而达到整个系统的高性能和高可用性。

高性能计算集群：高性能计算集群采用将计算任务分配到集群的不同计算节点而提高计算能力，因而主要应用在科学计算领域。


--创建对象的时候，假象它新增了一个对象怎么办？
这个时候我们的大脑会去思考怎么更好的进行设计？



### 常用对象的解决

1.this.$refs 
2.this.$emit('input', val)
3.this.$message.success('保存成功！')
4.<template slot-scope="scope">

### 电商平台

Real 德国电商平台
Fruugo 英国电商平台



--通过git下载代码 Git Bash Here
git clone http://159.75.212.51:8888/sunway/sunway.admin.web.git  --这是前端代码


--gitlab 公司自建代码托管平台
http://159.75.212.51:8888/    账号：1298775182@qq.com   wr1436062434

### 下载Mysql的地址

https://downloads.mysql.com/archives/community/

### 大黄蜂截图的方法

shift＋win＋S

### 公司的wifi信息

wifi名称：GZSW1 或者 GZSW2020

连接密码：Shengwei888#

验证账号：wangrui
验证密码：wr@#123456

### 给label的变量赋值

$("#jg").text("结果是:" + data);

### 基础的html架构

```
<html>

<head>一个案例
<title>这是一个小例子</title>
<script src="https://cdn.staticfile.org/vue/2.4.2/vue.min.js"></script>
<script src="https://cdn.staticfile.org/vue-resource/1.5.1/vue-resource.min.js"></script>
<script src="https://unpkg.com/axios/dist/axios.min.js"></script>
</head>

<body>

</body>
</html>
```

## 流程图

https://zhuanlan.zhihu.com/p/630168369  直接通过ChatGPT画流程图 牛逼

![扫描逻辑侵权图.drawio (Vue进步.assets/扫描逻辑侵权图.drawio (2).png)](C:/Users/Administrator/Downloads/扫描逻辑侵权图.drawio (2).png)

## 原型图

使用GPT自动画原型图：https://zhuanlan.zhihu.com/p/624638476

## 易仓

**易仓可以看订单销售**

![1690252257867](Vue进步.assets/1690252257867.png)

**拉取类目方法：SyncCategory**



**易仓先新增，速牛后新增**：在易仓新增ID后（1级，2级，3级）通过调用Task接口**SyncCategory**方法自动同步新增的1级，2级，3级类目的信息到速牛系统的类目，如下图所属：

- 需要先提供大类ID,大类ID由桂芬提供
- 需要注释更新逻辑的代码
- 需要添加断点功能

![1660642812716](Vue进步.assets/1660642812716.png)

**速牛新增，易仓再新增**：通过PostMan手动调用接口，手动将ID值赋值到类目表的易仓ID属性中，通过调用类目接口categotyList获得ID， 如下图所示：

![1660700190588](Vue进步.assets/1660700190588.png)

## ElementUI

**style：**改变ElementUi控件样式

```
style="width:500px;" 
```

**Dialog 居中弹窗** ：居中弹出，可以放表单提交，用table展示信息

**select下拉选项：** 文本框样式的下拉，适合放在form表单中，支持单选，多选

**Dropdown 下拉菜单**：lable样式的下拉，适合放在手机端设计下拉查询

**form表单**：表单元素

```
      // 清空form表单中的信息
      this.$nextTick(() => {
        this.$refs.form.resetFields()
      })
```

**PlatformShopSelect 店铺名称组件：** 引用店铺名称

```c#
//引用组件
import PlatformShopSelect from '@/components/Common/PlatformShopSelect'

//使用组件（注意后面的:multiple="true"，是bool变量那么属性的名称要赋值冒号：）
<PlatformShopSelect v-model="formData.acctionId" platform="fyndiq" :multiple="true" />

```

### Layout 布局：el-row el-col

el-row 一行24分，不论有多少el-col 格子，下面是一个例子：（注意如果碰到el-select可以设置： style="width:100%" 让他的长度与el-input的长度一致）

**gutter是指栅格间间隔（左右间隔），offset是指栅格左侧的间隔格数**

```c#
 //:span宽度加起来一般是24分（比较好控制，不是强制要求）。
 <el-row> // el-row通过:gutter设置el-col之间的间隔距离(通用的设置，就不用每个el-col进行设置)
  <el-col :span="12"></el-col> //el-col通过设置:offset进行分栏漂移
 <el-col :span="12"> </el-col> //（不需要设置:gutter间隔，就可以偏离左边距多少多少位置）
  </el-row>
```

一个实际的例子与效果： **一行两列**

```
              <el-row>
                <el-col :span="10">
                  <el-form-item prop="proCode" label="产品颜色">
                    <el-input v-model="formData.proCode" type="text" placeholder="请输入产品颜色" />
                  </el-form-item>
                </el-col>
                <el-col :span="10" :offset="4">
                  <el-form-item prop="proCode" label="产品规格">
                    <el-input v-model="formData.proCode" type="text" placeholder="产品规格" />
                  </el-form-item>
                </el-col>
              </el-row>
```

**效果：**

![1643251525305](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1643251525305.png)

**快速两列模板**

```c#
  <el-row>
  <el-col :span="10">
  
  </el-col>
  <el-col :span="10" :offset="4">
  
  </el-col>
</el-row>
```

**一行三列**

```c#
              <el-row>
                <el-col :span="6">
                  <el-form-item prop="proCode" label="产品重量：">
                    <el-input v-model="formData.proCode" type="text" placeholder="产品重量：" />
                  </el-form-item>
                </el-col>
                <el-col :span="6" :offset="3">
                  <el-form-item prop="proCode" label="包装重量：">
                    <el-input v-model="formData.proCode" type="text" placeholder="包装重量：" />
                  </el-form-item>
                </el-col>
                <el-col :span="6" :offset="3">
                  <el-form-item prop="proCode" label="重量允许浮动值">
                    <el-input v-model="formData.proCode" type="text" placeholder="重量允许浮动值" />
                  </el-form-item>
                </el-col>

              </el-row>
```

**效果**

![1643252155292](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1643252155292.png)

**快速三列模板：**

```c#
            <el-row>
                  <el-col :span="6">
                  
                  </el-col>
                   <el-col :span="6" :offset="3">
                   
                   </el-col>
                  <el-col :span="6" :offset="3">
                  
                  </el-col>
            </el-row>
```



### el-table-column

**Table表格：** **el-table-column** 是Table表格的表头，比如姓名，年龄，性别

[elementUI框架的 el-row el-col 与 el-table-column 用法区别](https://blog.csdn.net/acoolgiser/article/details/105891189  )

## CMD

**如何进入D盘**

进入Cmd内容，输入: D: 进入D盘

![1665625876764](Vue进步.assets/1665625876764.png)

也可以/D 目标路径 直接跳转

![1665626149686](Vue进步.assets/1665626149686.png)

## Makdown语法

**箭头符号**>:

> '>'这个可以有特殊格式

**```c#**

```c#
//c#代码格式
```





## ABP



## AutoGPT

```mysql
#1.部署AutoGPT


```

​	https://mp.weixin.qq.com/s/gipZcQjyr6IE0bpgxf-pyw 环境搭建配置 

1.python环境的安装注意同时安装好环境变量，使用python --version判断是否安装正常

2.	搜索引擎的创建，注意上面攻略的位置，参考下面，搜索引擎的API，可以限制搜索哪些网站，不搜索哪些网站，还可以限制指定国家，这是很智能的
	
	https://blog.csdn.net/weixin_45178729/article/details/130255001 创建Google搜索引擎

​    3.pip install -r requirements.txt，这个命令在cmd命令下面执行，先到D:，再CD D:\GSCODE\Auto-GPT-master

![1682391702243](Vue进步.assets/1682391702243.png)

4.运行python -m autogpt，在cmd窗口执行,输入：manual  进入手动模式

![1682392290274](Vue进步.assets/1682392290274.png)

5.进入授权模式，输入y

![1682392410294](Vue进步.assets/1682392410294.png)

6.执行遇到错误了输入Y继续让他跑：

![1682392599751](Vue进步.assets/1682392599751.png)

7.输入 N 退出,我感觉我输入错了，他执行了GPT-4的API,我开始应该让AutoGPT执行ChatGPT3.5的API![1682392744681](Vue进步.assets/1682392744681.png)

8.python -m autogpt 重新开始

![1682393290391](Vue进步.assets/1682393290391.png)

9.python -m autogpt --help 查看一下帮助的提示

![1682393548102](Vue进步.assets/1682393548102.png)

10.python -m autogpt --gpt3only  ，使用gpt3only跑起来 这是手动的命令，如果要自动的话，用python -m autogpt --continuous --gpt3only，全自动生成命令，应该默认就是手动，可以输入python -m autogpt --help查看完整命令，输出文件位置：D:\GSCODE\Auto-GPT-master\autogpt\auto_gpt_workspace，回答的时候可以输入：

**1.更改配置**：根目录下面.env文件，配置三个配置：



OpenAI：OPENAI_API_KEY=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

 GoogleAPI: GOOGLE_API_KEY=AIzaSyDu6BSxviS1emdC3GSgKLCTp1LvAc2G_pU
 Google搜索引擎：CUSTOM_SEARCH_ENGINE_ID=9737d2a9e5e62406d

 用来保存结果的key PINECONE_API_KEY=40f3a3df-bee2-40ab-ac6f-0ca2376b7de7
 用来保存结果的环境名称 PINECONE_ENV=asia-southeast1-gcp

**2.安装依赖库**：pip install -r requirements.txt

3.**设置代理**: 我用的代理端口号是7890，如果不确定可以问网络管理员代理的端口号是多少

![1682472519375](Vue进步.assets/1682472519375.png)

set https_proxy=http://127.0.0.1:7890
set http_proxy=http://127.0.0.1:7890

set all_proxy=socks5://127.0.0.1:7890


4.**更改源码**:修改browse_website函数

1.`写入代理配置到:autogpt/workspace.py`



```python
# 更新文件autogpt/workspace.py
from __future__ import annotations
#begin add2
import os
# end
from pathlib import Path

#begin add1
os.environ["http_proxy"] = "http://127.0.0.1:7890"
os.environ["https_proxy"] = "http://127.0.0.1:7890"
os.environ["all_proxy"] = "socks5://127.0.0.1:7890"
# end 

from autogpt.config import Config
....
```

2.`写入代理配置到：autogpt\commands\web_selenium.py`

```python
import autogpt.processing.text as summary
from autogpt.config import Config
from autogpt.processing.html import extract_hyperlinks, format_hyperlinks

FILE_DIR = Path(__file__).parent.parent
CFG = Config()


def browse_website(url: str, question: str) -> tuple[str, WebDriver]:
    #【begin add2】
    if 'http_proxy' in os.environ:
        del os.environ['http_proxy']

    if 'https_proxy' in os.environ:
        del os.environ['https_proxy']

    if 'all_proxy' in os.environ:
        del os.environ['all_proxy']
    #【end】
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
    # 【begin add1】
    os.environ["http_proxy"] = "http://127.0.0.1:7890"
    os.environ["https_proxy"] = "http://127.0.0.1:7890"
    os.environ["all_proxy"] = "socks5://127.0.0.1:7890"
    # 【end】
    return f"Answer gathered from website: {summary_text} \n \n Links: {links}", driver

```

[原文链接](https://github.com/Significant-Gravitas/Auto-GPT/issues/1467)

5.**执行源码**:python -m autogpt --gpt3only ，自动跑：python -m autogpt --continuous --gpt3only

6.**问题**:

Goal1:写一篇门锁市场的研究报告

Goal2:字数是100

Goal3:全部输出后翻译成中文保存

Goal4:once done,shutdown





PLAN:计划是什么
CRITICISM：这个目标的执行需要注意什么
NEXT ACTION：下一个方法
SYSTEM：系统提示,执行结果，或者遇到了什么问题

THOUGHTS：想法，做这件事情目标的细化拆解

REASONING：推理， 做的时候会遇到什么问题

提问的标准：结尾加上一句，如果完成后马上结束程序

![1682394761108](Vue进步.assets/1682394761108.png)

11.答案出来了，但是没有找到答案，打算换一种方式找出来,输入y,yes

![1682394873216](Vue进步.assets/1682394873216.png)

12.它说2023年的奥斯卡颁奖典礼还没有开始，然后我让他找2022年奥斯卡最佳女主角是谁,输入问题后，输入y，yes

![1682394993946](Vue进步.assets/1682394993946.png)

错误日志：

1.**启动自用代理后解决**： 访问openai的api 偶尔提示异常：Remote end closed connection without response 远程端关闭连接，无响应 

2.**没有好的解决方案，重新启动** :  I'm sorry, but I cannot provide a `return` value without knowing the contents of `<full_code_string>` and `<list_of_focus_areas>`.  Please provide me with the necessary information. 对不起，在不知道`&lt;full_code_string&gt;`和`&lt;list_of_focus_areas&gt;`内容的情况下，我无法提供`return`值。请提供给我必要的信息。

3.**重新启动**：卡在那里thinking... 一直不动，按ctrl+c也没有反应

4.**重新启动**:openai.error.APIConnectionError: Error communicating with OpenAI 远程主机强迫关闭了一个现有的连接



5.**设置代理端口，更改源码autogpt/workspace.py**: [WinError 10060] 由于连接方在一段时间后没有正确回答或连接的主机没有反应，连接试失败 

6.**设置代理端口，更改源码autogpt\commands\web_selenium.py**:browse_website函数（Selenium）不工作了，抛出一个空错误，访问google返回不了结果

7.**安装docker桌面端解决**:注意安装好了要进去用一下，配置一些默认的设置，让docker桌面端可以用

在工作空间'D:\GSCODE\Auto-GPT-0.2.2\auto_gpt_workspace'中执行文件'analyze_smart_door_lock_market.py' 

无法在容器中运行脚本。如果还没有，请安装Docker https://docs.docker.com/get-docker/ 

系统:命令execute_python_file返回:Error:错误，而获取服务器API版本:(2，'CreateFile'， '系统找不到指定的文件. ')

## Cursor

总体来说效果是不错的，有一些小技巧：

1. 和ChatGPT、NewBing一样，如果你直接给他一个复杂任务，它可能处理得不一定好，但你如果帮它拆解好，或者通过对话的方式让他自行拆解，再分段输出，效果就很棒
2. 如果输出的是英文注释你可以选中对应部分，按Ctrl+K要求输出中文，它会帮你重新写
3. 写到一半断掉的话，选中断掉的部分按Ctrl+K，输入“继续”或者“Continue”，它会帮你继续写

## GRPC

> 扩展阅读：
>
> **如何实现RPC远程服务调用？** https://learn.lianglianglee.com/%e4%b8%93%e6%a0%8f/%e4%bb%8e0%e5%bc%80%e5%a7%8b%e5%ad%a6%e5%be%ae%e6%9c%8d%e5%8a%a1/06%20%e5%a6%82%e4%bd%95%e5%ae%9e%e7%8e%b0RPC%e8%bf%9c%e7%a8%8b%e6%9c%8d%e5%8a%a1%e8%b0%83%e7%94%a8%ef%bc%9f.md 
>
> .NET Core微服务之路：（纯干货）基于gRPC服务发现与服务治理的方案 https://developer.aliyun.com/article/665171
>
> 【从0-1 千万级直播项目实战】gRpc落地 | 可能是全网最全实战篇 https://juejin.cn/post/7234543902400397373

### RPC基础概念


RPC 是 "Remote Procedure Call"（远程过程调用）的缩写。它是一种通信协议，允许运行在一台计算机上的程序（客户端）通过网络请求另一台计算机（服务器）上的服务，就像调用本地函数或方法一样。简单地说，它允许你执行一个位于另一台机器上的函数，并接收返回结果。

RPC 有多种实现和协议，例如：

- **XML-RPC**: 使用 XML 来编码消息，通常运行在 HTTP 上。
- **JSON-RPC**: 使用 JSON 来编码消息。
- **SOAP**: 一种更复杂，使用 XML 编码的协议，通常运行在 HTTP 或 SMTP 上。
- **gRPC**: `Google 开发的一个高性能、开源和通用的 RPC 框架，使用 Protocol Buffers 作为接口描述语言。`

### RPC 的主要优点包括：

1. **强类型**: 由于使用接口描述语言（如 Protocol Buffers），RPC 能确保客户端和服务器端的强类型一致性。
2. **高性能**: `尤其是像 gRPC 这样的现代 RPC 框架，由于使用了 HTTP/2 和 Protocol Buffers，所以性能通常优于传统的 RESTful API。`
3. **多语言支持**: 基于 IDL 自动生成的代码可以支持多种编程语言，这样可以更容易地创建多语言的客户端和服务器。
4. **功能丰富**: 现代 RPC 框架通常提供了高级功能，如流处理、错误传播、负载均衡等。

### RPC 的主要缺点：

1. **复杂性**: 特别是像 SOAP 这样的老式 RPC 协议，其复杂性可能会让人望而却步。
2. **不是标准的 Web 协议**: 除了 HTTP 基础的 RPC（如某些 RESTful 服务）外，大多数 RPC 框架不能通过标准的 Web 浏览器直接访问。
3. **调试和缓存**: RPC 调用往往更难以缓存和调试，特别是与 RESTful API 相比。
4. **不适合对外**：因为要提供.proto文件给对方生成调用方法的格式，每个服务都需要提供这个文件，还是很麻烦

因为有多种 RPC 的风格和实现，不同的 RPC 技术可能会有不同的特点和限制。总体而言，RPC 是一种强大的分布式计算模型，但也需要适当的场景和专业知识来最大化其优势。

### 当考虑使用何种技术时：

- **内部服务间通信**: gRPC 是一个很好的选择，因为它高效且强类型。
- **公开 API**: RESTful 通常更受欢迎，因为它更简单、更广泛支持。
- **混合使用**: 你也可以在一个项目中同时使用 gRPC 和 RESTful API。例如，内部服务使用 gRPC 进行通信，而对外则提供 RESTful API。

选择哪一种技术取决于你的具体需求，两者都有一系列成功的使用案例。



gRPC 和 RESTful Web API 都有各自的优缺点，适用于不同的场景。

### gRPC 的优点：

1. **性能**: gRPC 使用 Protocol Buffers 和 HTTP/2，通常比 JSON over HTTP/1.1 更高效。
2. **强类型**: 通过 `.proto` 文件，gRPC 为 API 和数据模型提供了强类型支持。
3. **流**: gRPC 支持服务器和客户端流，以及双向流。
4. **多语言支持**: 自动代码生成支持多种语言。
5. **内置特性**: 如认证、负载均衡、重试和超时等。

### gRPC 的缺点：

1. **复杂性**: 对于简单的 API，gRPC 可能显得有些过度。
2. **不易阅读**: Protocol Buffers 是二进制格式，不如 JSON 易于人工阅读和调试。
3. **限制**: 需要 HTTP/2 支持，而一些老旧的设备和浏览器可能不支持。
4. **跨域请求**: 在 Web 场景中，跨域请求（CORS）的支持不如 RESTful API。

### Web API（RESTful）的优点：

1. **简单易用**: 对于简单的 CRUD 操作，RESTful API 更为直观。
2. **广泛支持**: 几乎所有设备和浏览器都支持 HTTP/1.1。
3. **可读性**: JSON 格式易于阅读和调试。
4. **灵活性**: RESTful API 通常更加灵活，可以更容易地进行版本控制。

### Web API（RESTful）的缺点：

1. **性能开销**: 通常使用 JSON 格式和 HTTP/1.1，这可能比使用 Protocol Buffers 和 HTTP/2 的 gRPC 稍慢。
2. **弱类型**: RESTful API 通常不提供强类型的 schema（尽管可以通过 OpenAPI、JSON Schema 等工具进行补充），这可能会导致客户端和服务器之间的接口不匹配。
3. **版本管理**: RESTful API 的版本管理通常更为复杂，因为它没有内建的服务定义或类型检查。
4. **错误处理**: RESTful API 的错误处理和状态代码不如 gRPC 的丰富。
5. **批量操作和流**: 虽然有 WebSockets 和 HTTP/2，但 RESTful API 在批量和流式操作方面的支持通常不如 gRPC。
6. **多语言支持**: 尽管很多语言都支持 HTTP 和 JSON，但你通常需要手动编写客户端和服务器代码，而不是像 gRPC 那样自动生成。
7. **功能丰富性**: 虽然 HTTP/1.1 和 HTTP/2 提供了基础的请求/响应模型，但它们没有像 gRPC 那样提供诸如超时、取消、流等高级功能。

### GRPC实战

> 看了概念开始实战内容，新建两个项目：1.服务端：定义服务接口 2.客户端：调用服务接口

#### 服务端

Net 7  Web Api

![1693558926071](Vue进步.assets/1693558926071.png)

`1.引用GRPC类库`

- Google.Protobuf
- Grpc.AspNetCore
- Grpc.Tools

`2.根目录新建：GreeterTest.proto 文件`

```protobuf
syntax = "proto3";

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply);
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}

以下为注释内容
syntax = "proto3"; 使用Protocol Buffers语言的版本，这里是第三版（proto3）。

定义了一个名为Greeter的gRPC服务，服务有一个SayHello方法
service Greeter { 
  SayHello 方法，入参HelloRequest，返回值：HelloReply
  rpc SayHello (HelloRequest) returns (HelloReply);
}

SayHello 方法 入参HelloRequest的实体内容，有一个string 类型的字符串
message HelloRequest {
  string name = 1;
}
SayHello 方法 返回值HelloReply的实体内容，有一个string 类型的字符串
message HelloReply {
  string message = 1;
}
```

`3.在.csproj解决方案配置中添加引用`

```cs
<PackageReference Include="Grpc.AspNetCore" Version="2.41.0" />
   <PackageReference Include="Grpc.Tools" Version="2.41.0" PrivateAssets="All" />
   <PackageReference Include="Google.Protobuf" Version="3.19.1" />
   <Protobuf Include="GreeterTest.proto" GrpcServices="Server" />
```

![1693559248739](Vue进步.assets/1693559248739.png)

`4.配置好上面的路径后，右键重新生成解决方案，自动根据.proto配置生成服务类`

![1693559103250](Vue进步.assets/1693559103250.png)

`5.创建GreeterService.cs服务类`

```c#
using Grpc.Core;
using Microsoft.AspNetCore.Mvc;

namespace WebApplicationGRPCW.Controllers
{
    /// <summary>
    /// Greeter.GreeterBase 是通过.proto文件自动生成的基类
    /// </summary>
    public class GreeterService : Greeter.GreeterBase
    {
        /// <summary>
        /// 这个方法就是我们在.proto文件中定义的那个方法，入参，返回值都是一样的
        /// ServerCallContext是服务上下文的请求信息，比如请求头，授权，客户端是否取消
        /// </summary>
        /// <param name="request"></param>
        /// <param name="context"></param>
        /// <returns></returns>
        public override Task<HelloReply> SayHello(HelloRequest request, ServerCallContext context)
        {
            return Task.FromResult(new HelloReply
            {
                Message = "Hello " + request.Name
            });
        }
    }

}
```

`6.注入GRPC服务到Program.cs`

```c#
using Microsoft.AspNetCore.Server.Kestrel.Core;
using WebApplicationGRPCW.Controllers;

var builder = WebApplication.CreateBuilder(args);

// Add services to the container.
builder.Services.AddControllers();
builder.Services.AddGrpc();
...
builder.WebHost.ConfigureKestrel(options =>
{
     //这个7109端口看你服务的端口地址
    options.ListenAnyIP(7109, listenOptions =>
    {
        listenOptions.Protocols = HttpProtocols.Http2;
        listenOptions.UseHttps(); // 使用默认或开发证书
    });
});

var app = builder.Build();

...
app.UseRouting();

app.UseEndpoints(endpoints =>
{
    endpoints.MapGrpcService<GreeterService>();
});
...

app.Run();

```



#### 客户端

客户端就是调用上面的那个SayHello方法，有五个步骤：

`1.引用GRPC类库`

- Google.Protobuf
- Grpc.AspNetCore
- Grpc.Tools

`2.根目录新建：GreeterTest.proto 文件，注意这个文件的内容要和服务端的一样`

```protobuf
syntax = "proto3";

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply);
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}

```

`3.在.csproj解决方案配置中添加引用，注意这里的GrpcServices="Client"`

```cs
<PackageReference Include="Grpc.AspNetCore" Version="2.41.0" />
   <PackageReference Include="Grpc.Tools" Version="2.41.0" PrivateAssets="All" />
   <PackageReference Include="Google.Protobuf" Version="3.19.1" />
   <Protobuf Include="GreeterTest.proto" GrpcServices="Client" />
```

![1693559994814](Vue进步.assets/1693559994814.png)

`4.配置好上面的路径后，右键重新生成解决方案，自动根据proto配置生成调用类`

![1693560075123](Vue进步.assets/1693560075123.png)

`5.定义SayHelloAsync方法，调用`

```c#
using Grpc.Net.Client;
using Microsoft.AspNetCore.Mvc;

namespace WebApplicationGRPC.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class WeatherForecastController : ControllerBase
    {

        private readonly ILogger<WeatherForecastController> _logger;

        public WeatherForecastController(ILogger<WeatherForecastController> logger)
        {
            _logger = logger;
        }

        [HttpGet(Name = "SayHelloAsync")]
        public async Task SayHelloAsync()
        {
            var httpHandler = new HttpClientHandler();
            httpHandler.ServerCertificateCustomValidationCallback = HttpClientHandler.DangerousAcceptAnyServerCertificateValidator;
            // 不验证服务端证书（仅用于开发和测试！）
            var channel = GrpcChannel.ForAddress("https://localhost:7109", new GrpcChannelOptions { HttpHandler = httpHandler });
            var client = new Greeter.GreeterClient(channel);

 
            var reply = await client.SayHelloAsync(
                              new HelloRequest { Name = "World" });
            Console.WriteLine("Greeting: " + reply.Message);

        }


    }
}
```

`6.调用`

> ```
> 注意客户端与服务端两边的服务都是用Https的方式启动的
> ```

![1693560198587](Vue进步.assets/1693560198587.png)

返回：Hello World 成功

![1693560295837](Vue进步.assets/1693560295837.png)

#### Vue调用

使用 `grpc-web` 生成 JavaScript 代码的过程通常涉及以下几个步骤：

##### 1. 安装必要的工具

确保你已经安装了 `protoc`（Protocol Buffers 编译器）和 `protoc-gen-grpc-web`（gRPC-Web 插件）。

- `protoc`: [下载页面](https://github.com/protocolbuffers/protobuf/releases)

- `protoc-gen-grpc-web`: 可以通过以下命令安装：

  ```
  bashCopy codesudo apt-get install protoc-gen-grpc-web
  ```

  或者

  ```
  bashCopy codego get -u github.com/grpc/grpc-web
  ```

##### 2. 定义 `.proto` 文件

创建一个 `.proto` 文件（例如，`helloworld.proto`）来定义你的 gRPC 服务和消息类型。

```
syntax = "proto3";

package helloworld;

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply);
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

##### 3. 生成代码

使用 `protoc` 和 `protoc-gen-grpc-web` 插件生成 JavaScript 代码。

在命令行中运行：

```
protoc -I=. helloworld.proto \
  --js_out=import_style=commonjs:. \
  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.
```

这将生成两个文件：

- `helloworld_pb.js`: 包含消息类型的 JavaScript 代码。
- `helloworld_grpc_web_pb.js`: 包含 gRPC-Web 客户端的 JavaScript 代码。

##### 4. 在前端项目中使用生成的代码

首先，确保你已经安装了 `grpc-web` 的 JavaScript 客户端库：

```
npm install grpc-web
```

然后，在你的前端代码中导入生成的 JavaScript 文件和 `grpc-web` 库，然后使用它们来调用 gRPC 服务：

```javascript
const {HelloRequest, HelloReply} = require('./path/to/helloworld_pb.js');
const {GreeterClient} = require('./path/to/helloworld_grpc_web_pb.js');

const client = new GreeterClient('https://localhost:7109');

const request = new HelloRequest();
request.setName('World');

client.sayHello(request, {}, (err, response) => {
  console.log(response.getMessage());
});
```

这样，你就可以在前端代码中通过 gRPC-Web 调用你的 gRPC 服务了。注意替换代码中的路径和 URL，以便与你的项目设置匹配。

## HTTP

`RestSharp` 和 `HttpClient` 都是用于.NET环境中进行HTTP请求的库，但它们之间存在一些区别。以下是一些主要区别：

### 设计和使用场景

- **HttpClient**: 是.NET Framework 和 .NET Core 的一部分，更为通用和底层。适用于多种HTTP请求场景。
- **RestSharp**: 更专注于RESTful API调用，提供了更多高级特性，例如自动序列化和反序列化JSON、XML等。

### 可配置性和灵活性

- **HttpClient**: 提供了高度可定制和灵活的API，但这通常意味着需要写更多的代码来实现特定功能。
- **RestSharp**: 为常见任务提供了更多的“开箱即用”的功能，使得代码更为简洁，但这可能会牺牲一些灵活性。

### 异步支持

- **HttpClient**: 在.NET Core 和 .NET 5/6 中，有很强的异步编程支持。
- **RestSharp**: 也支持异步请求，但其API设计更偏向于同步或基于事件的异步模型。

### 依赖和版本

- **HttpClient**: 作为.NET标准库的一部分，维护周期和.NET运行时绑定。
- **RestSharp**: 是一个独立的第三方开源库，需要单独安装和更新。

### 社区和文档

- **HttpClient**: 由于是官方库，文档通常更全面，社区支持也更广泛。
- **RestSharp**: 虽然有良好的社区支持和文档，但可能没有`HttpClient`那么广泛。

### 兼容性

- **HttpClient**: 在多个.NET运行时版本和平台上有很好的支持。
- **RestSharp**: 主要用于.NET Framework 和 .NET Core，也有对其他平台的支持，但可能不如`HttpClient`全面。

### 示例代码

使用 `HttpClient` 发起一个GET请求：

```c#
HttpClient httpClient = new HttpClient();
var response = await httpClient.GetAsync("https://api.example.com/items");
```

使用 `RestSharp` 发起一个GET请求：

```c#
var client = new RestClient("https://api.example.com/items");
var request = new RestRequest(Method.GET);
IRestResponse response = client.Execute(request);
```

选择哪一个库取决于你的具体需求，包括你愿意编写多少代码，需要哪些特定功能，以及你的项目依赖。

### 示例 1: 发起 GET 请求并获取 JSON 响应

#### 使用 HttpClient：

```c#
using System.Net.Http;
using Newtonsoft.Json;
using System.Threading.Tasks;

class Program
{
    static async Task Main(string[] args)
    {
        HttpClient httpClient = new HttpClient();
        HttpResponseMessage response = await httpClient.GetAsync("https://jsonplaceholder.typicode.com/todos/1");
        string json = await response.Content.ReadAsStringAsync();
        var todoItem = JsonConvert.DeserializeObject<TodoItem>(json);
    }
}

class TodoItem
{
    public int UserId { get; set; }
    public int Id { get; set; }
    public string Title { get; set; }
    public bool Completed { get; set; }
}
```

这里，你需要手动进行反序列化，通常使用如 `Newtonsoft.Json` 的库。

#### 使用 RestSharp：

```c#
using RestSharp;

class Program
{
    static void Main(string[] args)
    {
        var client = new RestClient("https://jsonplaceholder.typicode.com");
        var request = new RestRequest("todos/1", Method.GET);
        var response = client.Execute<TodoItem>(request);
        TodoItem todoItem = response.Data;
    }
}

class TodoItem
{
    public int UserId { get; set; }
    public int Id { get; set; }
    public string Title { get; set; }
    public bool Completed { get; set; }
}
```

在这个示例中，RestSharp 自动进行了 JSON 的反序列化，不需要额外的代码。

------

### 示例 2: 发起 POST 请求并发送 JSON 数据

#### 使用 HttpClient：

```c#
using System.Net.Http;
using Newtonsoft.Json;
using System.Threading.Tasks;
using System.Text;

class Program
{
    static async Task Main(string[] args)
    {
        HttpClient httpClient = new HttpClient();
        TodoItem todoItem = new TodoItem { UserId = 1, Id = 2, Title = "task", Completed = false };
        string json = JsonConvert.SerializeObject(todoItem);
        var content = new StringContent(json, Encoding.UTF8, "application/json");
        HttpResponseMessage response = await httpClient.PostAsync("https://jsonplaceholder.typicode.com/todos", content);
    }
}

class TodoItem
{
    public int UserId { get; set; }
    public int Id { get; set; }
    public string Title { get; set; }
    public bool Completed { get; set; }
}
```

这里，你需要手动进行序列化，并创建一个 `StringContent` 对象。

#### 使用 RestSharp：

```c#
using RestSharp;

class Program
{
    static void Main(string[] args)
    {
        var client = new RestClient("https://jsonplaceholder.typicode.com");
        var request = new RestRequest("todos", Method.POST);
        TodoItem todoItem = new TodoItem { UserId = 1, Id = 2, Title = "task", Completed = false };
        request.AddJsonBody(todoItem);
        var response = client.Execute(request);
    }
}

class TodoItem
{
    public int UserId { get; set; }
    public int Id { get; set; }
    public string Title { get; set; }
    public bool Completed { get; set; }
}
```

在这个示例中，RestSharp 提供了一个 `AddJsonBody` 方法，它自动进行了 JSON 的序列化。

------

这些示例应该能帮助你更好地理解两者之间的不同。`HttpClient` 更为底层和通用，而 `RestSharp` 提供了一些高级功能，如自动序列化和反序列化，这在与 RESTful APIs 交互时可能会更方便。

当然，这次我将给出一个示例，展示如何通过`HttpClient`自定义HTTP请求，以便发送特定的HTTP头和对响应进行流式处理。这种自定义的需求在一些高级用例下是非常有用的。

### 示例 3: 发起自定义HTTP请求

#### 使用 HttpClient 

在这个示例中，我们将使用 `HttpClient` 来发送一个包含自定义 `User-Agent` 和 `Authorization` 头的 `GET` 请求，并且在获取响应后，直接对其进行流式处理。

```c#
using System;
using System.IO;
using System.Net.Http;
using System.Threading.Tasks;

class Program
{
    static async Task Main(string[] args)
    {
        HttpClient httpClient = new HttpClient();

        // 自定义HTTP请求头
        HttpRequestMessage request = new HttpRequestMessage(HttpMethod.Get, "https://jsonplaceholder.typicode.com/todos/1");
        request.Headers.Add("User-Agent", "MyApp");
        request.Headers.Add("Authorization", "Bearer " + "your-token-here");

        HttpResponseMessage response = await httpClient.SendAsync(request, HttpCompletionOption.ResponseHeadersRead);  // ResponseHeadersRead使我们能在收到响应头后立即开始处理响应体

        // 流式处理响应数据
        Stream responseStream = await response.Content.ReadAsStreamAsync();
        StreamReader reader = new StreamReader(responseStream);
        while (!reader.EndOfStream)
        {
            string line = await reader.ReadLineAsync();
            Console.WriteLine(line);
        }
    }
}
```

这个示例演示了几个关键点：

1. 我们可以轻易地为HTTP请求添加任意自定义头。
2. 使用`HttpCompletionOption.ResponseHeadersRead`可以更有效地处理大型响应，因为我们不需要等待整个响应都读入内存。

这种高度的灵活性和自定义能力是 `HttpClient` 的一个明显优势，尤其在处理更为复杂和特定的HTTP需求时。




#### 使用 RestSharp

```c#
using System;
using RestSharp;

class Program
{
    static void Main(string[] args)
    {
        var client = new RestClient("https://jsonplaceholder.typicode.com");
        var request = new RestRequest("todos/1", Method.GET);

        // 自定义 HTTP 请求头
        request.AddHeader("User-Agent", "MyApp");
        request.AddHeader("Authorization", "Bearer " + "your-token-here");

        IRestResponse response = client.Execute(request);

        // 输出响应数据
        Console.WriteLine(response.Content);
    }
}
```

#### 选择 `RestSharp` 还是 `HttpClient`

1. **如果你需要高度定制HTTP请求，或者需要对HTTP响应进行底层处理**（如上面的流式处理示例），那么 `HttpClient` 可能是更好的选择。
2. **如果你主要与RESTful APIs交互，并希望尽快开发出可工作的解决方案，而不需要太多的HTTP细节和自定义选项**，那么 `RestSharp` 可能会更方便。
3. **对于简单到中等复杂度的REST API交互，RestSharp 通常足够了**，而且因为其简洁性，你可能会更快地开发出解决方案。
4. **在高度并发的长时间运行的应用程序（如Web服务器）中，HttpClient 的性能通常优于 RestSharp**，尤其是在.NET Core和.NET 5/6中，因为它支持HTTP/2和其他高级特性。

最终，哪一个更适合你，取决于你的具体需求和你想如何权衡复杂性、灵活性和开发速度。



## 微服务

### 微服务架构跨数据库，跨服务怎么保证数据一致性？

比如我有一个场景，是库存服务，与订单服务，还有销售额的新增，当收到消费者的订单后，我库存减去了1，然后订单加1的时候出现了异常，我需要回滚我的状态，这个时候库存服务，订单服务，销售额新增服务是三个接口，我要怎么统一这三个接口的状态，并且在三个不同的库里面，他们的状态怎么保证一致？

我想的是两个方法：

1. **Saga模式**：try catch 库存 ,try catch 订单,try catch 销售额 ，定义三个变量捕获错误，如果有错误则回滚业务状态，其他的则不执行。将一个分布式事务拆分为多个本地事务，并按照一定的顺序执行。每个本地事务都对应一个操作，并且有一个对应的补偿操作。如果某个操作失败，就执行之前所有操作的补偿操作，以此来保证数据的一致性。Saga模式可以避免长时间的锁定，提高系统的并发性能

2.不使用分布式接口，使用一个接口进行调用存储过程进行库存，订单，销售额的数据库事务管理，这个方法虽然可以，不过如果数据库服务器不在一个网段呢？那么就会有问题



经过交流分析有两个解决思路：

1.ChatGPT:分布式事务协调器Seata:可以使用分布式事务管理框架来协调多个服务之间的事务。例如，可以使用基于消息队列的分布式事务解决方案，如基于本地消息表的TCC（Try-Confirm-Cancel）模式或者使用分布式事务协调器（如Seata）。

2.Eleven讲过 CAP:广州-12453-56-男-朝夕同学和橙子分享CAP最适合我这个场景 https://github.com/dotnetcore/CAP/blob/master/README.zh-cn.md 开源地址![1692156158807](Vue进步.assets/1692156158807.png)



> 参考链接：
>
> https://seata.io/zh-cn/blog/seata-at-tcc-saga.html 分布式事务 Seata 及其三种模式详解 | Meetup#3 回顾
>
> https://www.cnblogs.com/wzh2010/p/15311142.html分布式：分布式事务（CAP、两阶段提交、三阶段提交）
>
> https://developer.aliyun.com/article/1136581 【基础CAP的案例】.NET Core 事件总线,分布式事务解决方案：CAP
>
> https://www.liujiajia.me/?p=5257 【不知道NET CORE有没有，不过这个要求是在一个事务里面连接多个数据库，并不支持多个微服务】使用 TransactionScope 实现 SqlServer 的分布式事务

> 设计的时候需要注意的地方：

1. **定时任务和补偿机制：** 可以设置定时任务来检查和修复数据不一致的情况，例如，定期对库存和订单进行对账。看看库存的剩余量与订单的销售数量是否一致

1. **监控和报警：** 建立监控和报警系统，及时发现并处理数据不一致的情况，及时接入处理，针对于不能由程序自动补偿的问题，要由人工接入解决。
2. **幂等性操作：** 所有服务的接口应当设计成幂等的，即多次执行相同操作不会产生额外影响。这样，即使由于网络问题或其他原因导致操作重复，也不会破坏数据一致性。比如如果提交库存-1的时候就已经失败了，那么回滚的时候库存就不应该+1，多次提交也不会有问题
3. **重试与补偿机制：** 当涉及到跨服务的操作时，可以实现重试与补偿机制。例如，在库存减少操作成功后，如果订单增加操作失败，可以触发一个重试的机制，比如重试3次，重试3次后还是失败，可以触发一个补偿操作，回滚库存的减少。

## ORM

### EFCore建立仓储，统一数据访问接口

仓储模式是一种领域驱动设计（Domain Driven Design, DDD）中的模式，它的目标是创建一个清晰的抽象层，将数据访问逻辑隔离，使得业务逻辑可以不依赖具体的数据访问技术。

在C#的.NET开发中，仓储模式通常和Entity Framework等ORM工具一起使用，以提供一个统一的API来对数据库进行操作。这种方式可以提高代码的可维护性和可测试性，并且当底层数据库技术需要改变时，上层的业务逻辑代码不需要进行大的改动。

**我的疑惑**

```c#
    //问：EFDbContext实体类中DbContext，DbSet 这两个对象分别是什么意思
    ...代码省略
        private readonly DbContext context;
        private DbSet<T> entities;
        
    //答：DbContext可以看作是代表操作数据库的对象(多个库可以有多个DbContext)，DbSet<T>可以看作是操作具体哪个实体（具体表）的对象
	//以下是在仓储中直接使用DbContext对象操作Products
    public class ProductRepository : Repository<Product>, IProductRepository
    {
        private readonly EFDbContext _dbContext;
        public ProductRepository(EFDbContext context) : base(context)
        {
            _dbContext = context;
        }
        public Product GetProductById(int id)
        {
            return _dbContext.Products.Find(id);
            //return GetById(id);
        }

        public List<Product> GetProducts()
        {
            return _dbContext.Products.AsEnumerable().ToList();
            //return GetAll().ToList();
        }
    }
	
```



**第一步：引用类库：** 

```
Pomelo.EntityFrameworkCore.MySql：(支持 MySql 5.x +)
MySql.EntityFrameworkCore：支持 (MySql 8.x +)
```

以下是一个使用EF Core的简单仓储模式的例子：

首先，我们创建一个仓储接口（IRepository），这个接口定义了一些基本的数据访问操作：

```c#
public interface IRepository<T> where T : class
{
    IEnumerable<T> GetAll();
    T GetById(int id);
    void Add(T entity);
    void Update(T entity);
    void Delete(T entity);
}

```

然后我们创建一个EFRepository类来实现这个接口。这个类将具体的数据访问逻辑封装起来：

```c#
public class Repository<T> : IRepository<T> where T : class
{
    private readonly DbContext context;
    private DbSet<T> entities;

    public Repository(DbContext context)
    {
        this.context = context;
        entities = context.Set<T>();
    }

    public IEnumerable<T> GetAll()
    {
        return entities.AsEnumerable();
    }

    public T GetById(int id)
    {
        return entities.Find(id);
    }

    public void Add(T entity)
    {
        entities.Add(entity);
        context.SaveChanges();
    }

    public void Update(T entity)
    {
        entities.Update(entity);
        context.SaveChanges();
    }

    public void Delete(T entity)
    {
        entities.Remove(entity);
        context.SaveChanges();
    }
}

```

新建EFDbContext类，数据库请求上下文

```c#
    // EFDbContext 继承自 DbContext，是Entity Framework的核心部分
    public class EFDbContext : DbContext
    {
        // 这个构造函数允许将选项传递给DbContext（例如连接字符串、数据库提供商）
        public EFDbContext(DbContextOptions<EFDbContext> options) : base(options)
        {
        }

        // 假设你有一个名为Product的实体,通过context.Products.Where(p => p.Price > 1000)操作数据
        public DbSet<Product> Products { get; set; }

        // 如果有其他实体，你应该在这里也定义它们，例如：
        // public DbSet<OtherEntity> OtherEntities { get; set; }

        // 在此重写OnModelCreating方法，如果需要自定义模型创建过程
        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {
            modelBuilder.Entity<Product>(pc => pc.ToTable("product").HasKey(k => k.Id));

            base.OnModelCreating(modelBuilder);
        }
    }
```

appsettings.json文件中加入数据库配置

```json
   "ConnectionStrings": {
    "DefaultConnection": "Server=localhost;port=3306;database=mydb;uid=root;password=123456;Convert Zero Datetime=True;Allow User Variables=True;pooling=true"
  }
```

假设我们有一个名为`Product`的实体类，我们可能会定义一个名为`IProductRepository`的接口：

```c#
//创建实体类
 public class Product
 {
  public int Id { get; set; }
  public string Title { get; set; }
 
 }
//创建实体类仓储接口
//问：为什么要在这里继承IRepository<Product>接口呢？ProductRepository仓储类直接继承Repository类不久可以使用Repository仓库里面操作数据库的方法吗
  答：因为ProductContrller控制器中是使用IProductRepository抽象仓储去操作数据库的
public interface IProductRepository : IRepository<Product>
{
    Product GetProductById(int id);
    List<Product> GetProducts();
}


```

然后我们的仓储类（`ProductRepository`）可以实现这个接口：

```c#
    public class ProductRepository : Repository<Product>, IProductRepository
    {
        public ProductRepository(EFDbContext context) : base(context) { }

        /// <summary>
        /// 获取所有产品信息
        /// </summary>
        /// <returns></returns>
        public List<Product> GetProducts()
        {
            return GetAll().ToList();
        }

        /// <summary>
        /// 根据ID获取产品信息
        /// </summary>
        /// <param name="id"></param>
        /// <returns></returns>
        public Product GetProductById(int id)
        {
            return GetById(id);
        }
    }

```

创建表Product

```mysql
CREATE TABLE `product` (
  `Id` bigint NOT NULL AUTO_INCREMENT,
  `Title` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

# 构造数据
INSERT INTO `product` (`Title`) VALUES ('产品1');
INSERT INTO `product` (`Title`) VALUES ('产品2');
INSERT INTO `product` (`Title`) VALUES ('产品3');


```

`ProductRepository`类继承了`Repository<Product>`类，并实现了`IProductRepository`接口。这样我们就可以在业务逻辑中使用`IProductRepository`，从而操作Product数据，而无需关心具体的数据访问逻辑。

> 在.NET 6中，我们通常在Startup.cs文件的ConfigureServices方法中进行DbContext的依赖注入。首先，你需要安装Pomelo.EntityFrameworkCore.MySql包，它是用于.NET Core和EF Core的MySQL数据库提供程序。
>
> 以下是如何在.NET 6的ASP.NET Core项目中配置服务以使用MySQL数据库的示例：

```c#
public class Startup
{
    public Startup(IConfiguration configuration)
    {
        Configuration = configuration;
    }

    public IConfiguration Configuration { get; }

    public void ConfigureServices(IServiceCollection services)
    {
        services.AddDbContext<MyContext>(options =>
            options.UseMySql(Configuration.GetConnectionString("DefaultConnection"),
                new MySqlServerVersion(new Version(8, 0, 21))));

        services.AddScoped<IProductRepository, ProductRepository>();

        // 其他服务配置...
    }

    // 其他方法...
}

```

Net7中的语法

```c#
        public static void Main(string[] args)
        {
            var builder = WebApplication.CreateBuilder(args);

            // Add services to the container.

            builder.Services.AddControllers();
            // Learn more about configuring Swagger/OpenAPI at https://aka.ms/aspnetcore/swashbuckle
            builder.Services.AddEndpointsApiExplorer();
            builder.Services.AddSwaggerGen();
            //add by 从 builder 对象获取 Configuration 实例
            var configuration = builder.Configuration;
            builder.Services.AddDbContext<EFDbContext>(options =>
            options.UseMySql(configuration.GetConnectionString("DefaultConnection"),
                new MySqlServerVersion(new Version(8, 0, 21))));
             // add by 注册你的仓储，确保接口和实现类的对应关系被添加到了服务中
            builder.Services.AddScoped<IProductRepository, ProductRepository>(); 

            var app = builder.Build();

            // Configure the HTTP request pipeline.
            if (app.Environment.IsDevelopment())
            {
                app.UseSwagger();
                app.UseSwaggerUI();
            }

            app.UseHttpsRedirection();

            app.UseAuthorization();


            app.MapControllers();

            app.Run();
        }
    }
```

在上述代码中，我们首先在ConfigureServices方法中添加了一个MyContext的DbContext服务。UseMySql方法用于指定使用MySQL数据库，并配置了连接字符串。连接字符串通常在appsettings.json文件中定义，并通过IConfiguration接口获取。

然后，我们添加了一个仓储服务，这样在控制器或其他需要使用仓储的地方，就可以通过依赖注入的方式获取到仓储对象。

这里假定你有一个名为MyContext的DbContext类，以及IProductRepository和ProductRepository的接口和实现类。默认的数据库版本被设置为8.0.21，但这应根据你的MySQL服务器版本进行调整(**通过select @@version查询**)。

在控制器或其他类中，你可以通过构造函数注入的方式来使用DbContext和仓储：

```c#
using Microsoft.AspNetCore.Mvc;
using WebApplication11.Model;
using WebApplication11.Service;
// 不需要额外引入 'WebApplication11.Service.impl'，因为具体实现应该在 Startup.cs 中通过依赖注入来处理。
// using WebApplication11.Service.impl; 

namespace WebApplication11.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class ProductController : ControllerBase
    {
        private readonly IProductRepository _productRepository; // 一般使用 'private' 而非 'public'，因为这是一个注入的依赖，不应该从外部访问。

        public ProductController(IProductRepository productRepository)
        {
            _productRepository = productRepository;
        }

        // 这个方法对应 'GET /product'
        [HttpGet] // 注意：这里不需要 'Name' 属性，而是直接定义路由模板。
        public IActionResult GetProducts() // 建议使用 IActionResult 作为返回类型，这样可以更灵活地返回各种HTTP响应。
        {
            var products = _productRepository.GetProducts();
            return Ok(products); // 使用 'Ok' 方法返回一个 200 OK 响应，带有产品列表。
        }

        // 这个方法对应 'GET /product/{id}'
        [HttpGet("{id}")] // 路由模板包括参数 'id'，这样 Swagger 和路由系统可以知道如何处理它。
        public IActionResult GetProductById(int id)
        {
            var product = _productRepository.GetProductById(id);

            if (product == null) 
            {
                return NotFound(); // 如果没有找到产品，返回一个 404 Not Found 响应。
            }

            return Ok(product); // 如果找到了产品，使用 'Ok' 方法返回一个 200 OK 响应，带有产品数据。
        }
    }
}

```

[EfCore仓储建立参考](https://www.cnblogs.com/Mamba8-24/p/16098812.html)

### SqlSugar仓储建立，统一数据访问接口

引用类库：

```
Install-Package SqlSugarCore
```

**建立ISqlSugarRepository仓储接口**

..代码太长略，详情看skums代码,[也可以点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/ISqlSugarRepository.cs)

**建立SqlSugarRepository仓储实现类**

..代码太长略,详情看skums代码，[也可以点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/SqlSugarRepository.cs)

**建立IServiceCollection扩展类注入SqlSugar**

```c#
        /// <summary>
        /// SqlSugar配置
        /// </summary>
        /// <param name="services"></param>
        public static void AddSqlSugarConfigure(this IServiceCollection services, IConfiguration _appConfiguration)
        {
            services.AddScoped<ISqlSugarClient>(u =>
            {
                var sqlSugarClient = new SqlSugarClient(new ConnectionConfig()
                {
                    DbType = DbType.MySql,
                    ConnectionString = _appConfiguration.GetConnectionString("DefaultConnection"),
                    IsAutoCloseConnection = true,
                });
                return sqlSugarClient;
            });

            // 注册非泛型仓储
            services.AddScoped<ISqlSugarRepository, SqlSugarRepository>();

            // 注册 SqlSugar 仓储
            services.AddScoped(typeof(ISqlSugarRepository<>), typeof(SqlSugarRepository<>));
        }
```

**新建实体类**

```c#
    [SugarTable("Product")]
    public class Product
    {
        public int Id { get; set; }
        public string Title { get; set; }
    }
```

**调用注入方法**

```c#
builder.Services.AddSqlSugarConfigure(builder.Configuration);
var app = builder.Build();
```



**新建控制器调用**

```c#
using Microsoft.AspNetCore.Mvc;

namespace SqlSugar1.Controllers
{
    [ApiController]
    [Route("[controller]/[action]")]
    public class ProductController : ControllerBase
    {
        private readonly ISqlSugarRepository<Product> _productRepository;

        public ProductController(ISqlSugarRepository<Product> productRepository)
        {
            _productRepository = productRepository;
        }

        [HttpGet]
        public IActionResult GetProductById(int id)
        {
            var product = _productRepository.Context.Queryable<Product>().First(u => u.Id == id);
            if (product is null)
            {
                return NotFound();
            }
            return Ok(product);
        }

        [HttpGet]
        public IActionResult GetProducts()
        {
            var products = _productRepository.Context.Queryable<Product>().ToList();
            return Ok(products);
        }
    }
}

```

### SqlSugar多仓储模式，统一数据访问接口

这个项目是原先盛为Task项目中用到的SqlSugar的仓储，我写一个小例子进行应用，实际的代码可以点击下面的内容

- IBaseService<T> ：操作数据库的抽象接口，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/IBaseService.cs)
- BaseService<T> :操作数据库的实现类，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/BaseService.cs)
- IDBContext ：数据源，切换仓储接口，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/IDBContext.cs)
- SqlSugarImpl ：获取DbContext的实现类，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/SqlSugarImpl .cs)
- DataSourceAttribute ：数据库标志特性类，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/DataSourceAttribute.cs)
- IProductService ：产品服务接口抽象方法，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/IProductService.cs)
- ProductService：产品接口实现，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/ProductService.cs)
- Product: 产品实体类，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/Product.cs)
- ProductController ：产品控制器，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/ProductController.cs)
- Program：服务注入，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/https://github.com/AtaraxiaRun/Hello-Word/blob/master/Program.cs.cs)

学习到的点在于三点：

- **typeof注入泛型类**：通过builder.Services.AddScoped(typeof(IBaseService<>), typeof(BaseService<>)); 进行注入
- **标志特性标志数据库**： 通过特性标志在服务类中标志单个服务类使用的是哪个数据库
- **IServiceProvider获取注入类**： 通过IServiceProvider获取IDBContext的注入类

后续优化的方向可以是

1.**静态注入**：加入静态注入的方法，不需要在IproductService中注入IServiceProvider对象，简化服务类的使用，2.**规范文件夹**：然后规范文件夹的命名，将数据库仓储，服务类，控制器，设置好

3.**接入服务自动注入**： 自动注入新增的服务，AutoFac

## 设计模式

小明用积木的**结构** **创造**了一栋大楼，这个**行为**得到了父母的认可。

- **创建型模式**：提供创建对象的方法，想象你要养宠物，你可以外出森林抓一个，也可以去宠物院直接找店员买一个，常用创建型模式：单例模式、建造者模式、工厂方法模式、抽象工厂模式、原型模式。
- **结构型模式**：这些模式关注类和对象的组合，通过继承或组合来形成更大的结构（想象拼积木，把不同的积木拼凑在一起，组合成复杂的高楼大厦）。结构型模式：适配器模式、桥接模式、组合模式、装饰模式、门面模式、享元模式、代理模式。
- **行为型模式**：这些模式专注于对象之间的通信，帮助对象协作，分配职责（想想足球比赛，前锋，后卫需要配合才能赢得比赛）。常见的行为型模式包括：责任链模式、命令模式、解释器模式、迭代器模式、中介者模式、备忘录模式、观察者模式、状态模式、策略模式、模板方法模式、访问者模式。

### 单例模式（Singleton Pattern）

单例模式是为了提高类访问的性能，节约初始化开销，让实例在整个应用程序只有一个实例，并且有三种情况，1.简单单例模式，直接用 2.多线程单例模式，用Lock应对并发情况 3.分布式单例模式

> 简单单例模式：简单场景使用，没有单进程多线程并发，分布式并发的情况

```c#
using System;

// 单例类
public class Singleton
{
    //1.私有对象
    private static Singleton instance;
    //2.私有构造函数
    private Singleton() { }
    //3.通过方法获取
    public static Singleton GetInstance()
    {
        if (instance == null)
        {
            instance = new Singleton();
        }
        return instance;
    }

    public void SomeMethod()
    {
        Console.WriteLine("单例模式：执行了某个方法");
    }
}

// 客户端代码
class Client
{
    public void Main()
    {
        Singleton singleton = Singleton.GetInstance();
        singleton.SomeMethod();  // 输出：单例模式：执行了某个方法
    }
}

```

> 多线程单例模式：懒汉式双重检查锁定（Double-Checked Locking）,要注意这个只是获取的时候保证单例，如果你多个线程并发获取这个数据库对象保证是一个对象，但是你同一个数据库对象并发同时向数据库发起增/删/改/查请求也是有问题，数据库链接开启，关闭的状态都会有问题

```c#
using SqlSugar;

namespace SqlSugarThread
{
    /// <summary>
    /// 单例模式SqlSugar实例,双重检查锁定模式,适用于多线程场景
    /// </summary>
    public class SqlSugarSingletionClient
    {
        private static DbConnection instance;
        private static readonly object _lock = new object();

        /// <summary>
        /// 私有的构造函数，防止在外部被创建
        /// </summary>
        private SqlSugarSingletionClient() { }

        public static SqlSugarScope GetInstance()
        {
            //速记：空，锁，空，走
            if (instance == null)//步骤1：进来先看看有没有没有再加锁，防止不必要的性能损耗
            {
                lock (_lock) //步骤2：加锁
                {
                    if (instance == null) //步骤3：因为多个线程一般都卡在步骤2等lock释放锁，这里加instance == null的判断是为了，如果上一个A线程创建好了实例对象走了，释放了_lock，后面的B线程进来判断一下，如果实例不为空直接拿实例走就好，没必要再进行创建了，如果B再创建一个对象，就不是单例模式了
                    {
                        instance = new DbConnection();   
                    }
                }
            }

            return instance;
        }

    }
}



```

> 分布式高并发的单例模式

```c#
using StackExchange.Redis;

public class DistributedSingleton
{
    // Redis连接的懒加载实例
    private static readonly Lazy<ConnectionMultiplexer> lazyConnection = new Lazy<ConnectionMultiplexer>(() =>
    {
        return ConnectionMultiplexer.Connect("localhost");
    });

    // 获取Redis连接实例
    private static ConnectionMultiplexer RedisConnection => lazyConnection.Value;
    // 获取Redis数据库实例
    private static IDatabase RedisDb => RedisConnection.GetDatabase();

    // 私有构造函数，避免外部直接实例化
    private DistributedSingleton() { }

    // 获取单例实例的静态方法
    public static DistributedSingleton GetInstance()
    {
        // 尝试获取分布式锁，等待10秒钟,如果在这个等待时间内成功获取到锁，则会返回 true 表示获取锁成功，否则会返回 false 表示获取锁失败。
        if (TryGetDistributedLock("distributed_lock_key", TimeSpan.FromSeconds(10)))
        {
            try
            {
                // 成功获取锁，创建并返回单例实例
                return new DistributedSingleton();
            }
            finally
            {
                // 释放分布式锁
                ReleaseDistributedLock("distributed_lock_key");
            }
        }
        else
        {
            // 未获取到分布式锁，返回 null 或抛出异常，根据需求来决定
            return null;
        }
    }

    // 尝试获取分布式锁的方法
    private static bool TryGetDistributedLock(string lockKey, TimeSpan expirationTime)
    {
        // 使用Redis的LockTake方法获取分布式锁
        // lockKey：锁的键
        // "lock_value"：锁的值（可以是任意非空字符串）
        // expirationTime：锁的过期时间
        return RedisDb.LockTake(lockKey, "lock_value", expirationTime);
    }

    // 释放分布式锁的方法
    private static void ReleaseDistributedLock(string lockKey)
    {
        // 使用Redis的LockRelease方法释放分布式锁
        // lockKey：锁的键
        RedisDb.LockRelease(lockKey, "lock_value");
    }
}


```

### 简单工厂-自定义（Simple Factory）

工厂可以理解为获取类的实例的工厂,简单工厂比较适用于简单的场景，业务不会有什么变动的情况下

- **解耦**：比如我需要汉堡，薯条，可以，我不需要自己去厨房弄（New）,我只需要直接问工厂方法要，比如直接输入汉堡，薯条，可乐的参数我就可以得到对应的东西（对象），而不用看具体的实现。 （**这一块的参数也可以通过枚举进行完成**）
- **扩展性** ：如果其中一天，汉堡的调料变了，这个工厂方法被10个类引用了，如果没有工厂方法，类全部是依赖具体的new 汉堡()对象，那么我要更改10个类，那太恐怖了，可是现在工厂方法里面有创建汉堡的实例，我只需要更改工厂方法里面的类就可以了
- **新增功能**: 如果要吃鸡腿，那么新建一个鸡腿类，继承食物接口，然后在工厂方法中加入new 鸡腿()对象就可以了（其实这里违反了开闭原则）

![1705655679936](Vue进步.assets/1705655679936.png)

[简单工厂源码](https://github.com/AtaraxiaRun/Hello-Word/tree/master/Factory.Method/Factory.Method)

### 工厂方法 （Factory Method）

[工厂方法模式](https://www.zhihu.com/search?q=工厂方法模式&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra={"sourceType"%3A"answer"%2C"sourceId"%3A688147580})的工厂是创建出**一种**产品，而抽象工厂是创建出**一类**产品。

工厂方法中新增一个新的类别时不需要像简单工厂那样要在Factory工厂中新增if else方法，直接新建一个新的类别：比如数据库日志，直接通过继承基类的抽象工厂方法就可以创建实例。

![1705656468624](Vue进步.assets/1705656468624.png)

[抽象工厂源码](https://github.com/AtaraxiaRun/Hello-Word/tree/master/Abstract.Factory/Abstract.Factory)

### 抽象工厂（Abstract Factory）

说不通的是代表返回不同类型的对象，不看具体的

![1705658724288](Vue进步.assets/1705658724288.png)

**工厂方法总结** ：

> 简单工厂：一个工厂类，多个产品类，需要通过if else新增产品
>
> 抽象工厂：多个工厂类，多个产品类，通过新增工厂类新增产品类，产品工厂

## DDD

https://mp.weixin.qq.com/s/q2JVpGpnc7FvBX_a9QWbvA Quartz、Schedule、XXL-Job 3种任务的极简使用教程 - Docker 自动化配置自动导入库表！

## 定时调度

总结：

### Hangfire

**自动出现的Hangfire提示**

```c#
//提示原因未知（辉哥提醒如果是用dotnet run 命令启动的vs，是使用正式环境的配置进行启动的，这个时候需要注意hangfire配置是否正常）
info: Hangfire.Processing.BackgroundExecution[0]
      Execution RecurringJobScheduler recovered from the Faulted state after 00:00:16.0833009 and is in the Running state now
```

![1660179891470](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1660179891470.png)

### Coravel 

盛为目前在用的，很好用，有以下有点：

- **轻量级** :ConfiguraService注入一下就可以使用
- **语法简单**： 语法简单，学习成本低
- **功能丰富**: 支持cron表达式，时，分，秒，按小时，按天调度,可以按照管道线程分来任务

### BackgroundService

简洁版本的定时任务，使用微软BackgroundService创建

> 可以简单的启动定时任务，并且也可以通过写逻辑解决任务重叠问题（上一次任务没有启动，这一次不启动），也可以启动后立即执行，然后按照指定间隔执行，不过不支持Cron表达式，比如Hangfire就支持Cron表达式，并且有仪盘表可以看到作业的执行，可以很方便的管理，调度指定时间段的作业执行，不会导致某个时间段作业执行的很臃肿

```c#
//HelloWorldJob类
using Microsoft.Extensions.Hosting;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace ConsoleApp13
{
    public class HelloWorldJob : BackgroundService
    {

        protected override async Task ExecuteAsync(CancellationToken stoppingToken)
        {
            while (!stoppingToken.IsCancellationRequested)
            {
                var msg = $"CurrentTime:{DateTime.Now}, Hello World!";

                Console.WriteLine(msg);
                //1秒钟执行一次
                await Task.Delay(1000, stoppingToken);
            }
        }
    }
}
//调用类
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using System;
using System.Threading.Tasks;

namespace ConsoleApp13
{
    class Program
    {
        static async Task Main(string[] args)
        {
            // 创建和配置宿主
            var host = Host.CreateDefaultBuilder(args)
                .ConfigureServices((context, services) =>
                {
                    // 添加 HelloWorldJob 作为托管服务
                    services.AddHostedService<HelloWorldJob>();
                })
                .Build();

            // 启动宿主
            await host.StartAsync();

            Console.WriteLine("后台服务正在运行。按 'Enter' 键停止服务...");
            Console.ReadLine();

            // 开始关闭服务和宿主
            await host.StopAsync(TimeSpan.FromSeconds(5));
            Console.WriteLine("后台服务已停止。");
        }
    }
}


```

### Timer

简洁版本的定时任务，使用Timer类创建

> 需求足够的简单，不需要使用Hangfire,Coravel 这些可能都太重了，还有一点就是方便...
>
> ThreadPool`还可以用于执行定时任务。你可以使用`ThreadPool.RegisterWaitForSingleObject`方法来安排一个方法在特定的时间间隔后被调用。

```c#
//Program调用
using System;
using System.Collections.Generic;
using System.Linq;
using System.ServiceProcess;
using System.Text;
using System.Threading.Tasks;

namespace Sunway.Task.Win.FruugoListerPublish
{
    internal static class Program
    {
        /// <summary>
        /// 应用程序的主入口点。
        /// </summary>
        static void Main()
        {
            ServiceBase[] ServicesToRun;
            ServicesToRun = new ServiceBase[]
            {
                new Service1()
            };
            ServiceBase.Run(ServicesToRun);
        }
    }
}

//定时服务类
using System.ServiceProcess;

namespace Sunway.Task.Win.FruugoListerPublish
{
    public partial class Service1 : ServiceBase
    {
        ListerService service = new ListerService();
        public Service1()
        {
            InitializeComponent();
        }

        protected override void OnStart(string[] args)
        {
            service.Start();
        }

        protected override void OnStop()
        {
            service.Stop();
        }
    }
}
//服务实现类
using CsvHelper;
using MiniExcelLibs;
using MySqlX.XDevAPI.Common;
using Sunway.Task.Win.Common.Data;
using Sunway.Task.Win.Common.Entity.BailunBltpro;
using Sunway.Task.Win.Common.Enum;
using Sunway.Task.Win.Common.Service.BailunAccount;
using Sunway.Task.Win.Common.Service.BailunBltpro;
using Sunway.Task.Win.Common.Utils;
using System;
using System.Collections.Generic;
using System.Configuration;
using System.IO;
using System.Linq;
using System.Net;
using System.Text;
using System.Threading;

namespace Sunway.Task.Win.FruugoListerPublish
{
    public class ListerService
    {
        private Timer _timer;
        private TimeSpan intervalMinutes;
        private readonly static object lockObject = new object();

        private readonly Common.ThreadPool pool = new Common.ThreadPool(1, ActionExecute);

        public ListerService()
        {
            int minutes = Convert.ToInt32(ConfigurationManager.AppSettings["Interval"]); // N分钟
            intervalMinutes = new TimeSpan(0, minutes, 0);
        }

        public void ThreadAction(object obj)
        {
            try
            {
                //初始化下架数据
                new FruugoListerService().InitEndItemData();
                // 得到所有账号
                List<PlatformAccount> accounts = PlatformAccountService.Instance.GetAccountList((int)PlatformEnum.Fruugo);
                accounts = accounts.Where(a => a.AccountAlias == "caihui").ToList();  //测试使用
                foreach (var item in accounts)
                {
                    //pool.AddNewAction(item);
                    ActionExecute(item); //测试使用
                }
            }
            catch (Exception ex)
            {
                LoggerUtility.WriteLog($"初始化异常", ex.Message);
            }
        }
        public void Start()
        {
            if (null == _timer)
                _timer = new Timer(new TimerCallback(ThreadAction));
            _timer.Change(new TimeSpan(0), intervalMinutes);
        }

        public void Stop()
        {
            _timer.Dispose();
            _timer = null;

            pool.StopAll();
        }
}

```

### Quartz.NET

写了一个小Demo，官网地址：https://www.quartz-scheduler.net/documentation/quartz-3.x/quick-start.html#trying-out-the-application,引用类库：

Quartz.Plugins
Quartz
Quartz.Serialization.Json

```c#
using Quartz.Impl;
using Quartz.Logging;
using LogLevel = Quartz.Logging.LogLevel;

namespace Quartz.NET.Task
{
    public class Program
    {
        public static async System.Threading.Tasks.Task Main(string[] args)
        {
            // 配置日志提供者
            LogProvider.SetCurrentLogProvider(new ConsoleLogProvider());

            // 从工厂获取调度器实例
            StdSchedulerFactory factory = new StdSchedulerFactory();
            IScheduler scheduler = await factory.GetScheduler();

            // 启动调度器
            await scheduler.Start();

            // 定义工作任务，并将其与我们的HelloJob类关联
            IJobDetail job = JobBuilder.Create<HelloJob>()
                .WithIdentity("job1", "group1")
                .Build();
            // 配置触发器，现在立即执行工作任务，然后每10秒重复执行一次
            ITrigger trigger = TriggerBuilder.Create()
                .WithIdentity("trigger1", "group1")
                .StartNow()
                .WithSimpleSchedule(x => x
                    .WithIntervalInSeconds(10)
                    .RepeatForever())
                .Build();


            // 1.定义基础工作任务，按秒执行,也可以按小时，按分钟，按年月日执行
            IJobDetail job2 = JobBuilder.Create<HelloJob2>()
                .WithIdentity("job2", "group2")
                .Build();

            ITrigger trigger2 = TriggerBuilder.Create()
               .WithIdentity("trigger2", "group2")
               .StartNow()
               .WithSimpleSchedule(x => x
                   .WithIntervalInSeconds(5)
                   .RepeatForever())
               .Build();
           // 2.根据Cron表达式触发
            var example = TriggerBuilder.Create()
                          .WithIdentity("trigger-name", "trigger-group")
                          .ForJob("job-name", "job-group")
                          .WithCronSchedule("45 23 * * 6")
                          .Build();
            // 告诉Quartz使用我们的触发器来安排工作任务
            await scheduler.ScheduleJob(job, trigger);
            await scheduler.ScheduleJob(job2, trigger2);

            // 这里等待一段时间，以展示正在发生的情况
            await System.Threading.Tasks.Task.Delay(TimeSpan.FromSeconds(60));

            // 最后，当你准备关闭程序时，关闭调度器
            await scheduler.Shutdown();

            Console.WriteLine("按任意键关闭应用程序");
            Console.ReadKey();
        }

        public class ConsoleLogProvider : ILogProvider
        {
            public Logger GetLogger(string name)
            {
                return (level, func, exception, parameters) =>
                {
                    if (level >= LogLevel.Info && func != null)
                    {
                        Console.WriteLine("[" + DateTime.Now.ToLongTimeString() + "] [" + level + "] " + func(), parameters);
                    }
                    return true;
                };
            }

            public IDisposable OpenNestedContext(string message)
            {
                throw new NotImplementedException();
            }

            public IDisposable OpenMappedContext(string key, object value, bool destructure = false)
            {
                throw new NotImplementedException();
            }
        }

        public class HelloJob : IJob
        {
            public async System.Threading.Tasks.Task Execute(IJobExecutionContext context)
            {
                await Console.Out.WriteLineAsync("Greetings from HelloJob!");
            }
        }

        public class HelloJob2 : IJob
        {
            public async System.Threading.Tasks.Task Execute(IJobExecutionContext context)
            {
                await Console.Out.WriteLineAsync("Wang Rui Good!Greetings from HelloJob!");
            }
        }

    }
}
```



## Net线程并发

乐观锁，悲观锁

1. **线程同步和竞态条件：** 多个线程同时访问共享资源可能导致竞态条件和数据不一致问题。需要使用同步机制如互斥锁、信号量、事件等来保护共享资源的访问，以避免数据损坏或不一致性。
2. **死锁：** 死锁是指两个或多个线程相互等待对方释放资源，导致程序无法继续执行。正确地设计和管理锁定顺序是避免死锁的重要考虑因素。
3. **线程安全性：** 确保多个线程可以安全地访问共享数据结构是一个挑战。需要使用线程安全的数据结构或实施自定义的线程安全控制。
4. **性能问题：** 虽然多线程可以提高程序的性能，但如果处理不当，也可能导致性能问题，例如线程频繁切换、资源争夺等。需要进行适当的性能测试和优化。
5. **取消和异常处理：** 在多线程环境中，当一个线程取消操作或抛出异常时，需要确保其他线程能够正确地响应和处理这些情况，以避免应用程序崩溃或处于不稳定状态。
6. **调试困难：** 多线程程序中的bug可能会比单线程程序更难以调试和定位。并发问题可能是不稳定的，依赖于特定的运行时条件，使得问题难以复现。
8. **正确的线程协作：** 多线程程序中，线程之间需要正确地进行协作和通信。错误的线程通信可能导致死锁或数据一致性问题。

为了克服这些难点，开发者应该深入了解多线程编程概念和.NET提供的线程相关功能。同时，采用一些最佳实践，如避免共享状态、尽量使用不可变数据等，可以帮助减少多线程编程中的复杂性和问题。

扩展阅读：

面试必备：C#多线程技术 https://mp.weixin.qq.com/s/OyAOeZyuOcpE0Mu-Cu41Bw

### 使用多线程的五种方法

#### Task.WhenAll

> 方法1：所有任务同时并发， 消耗大量的性能，适合Http请求等，IO任务
>
> 方法2：使用SemaphoreSlim 信号量来限制并发，适合资源有限的情况，CPU,内存的消耗低

```c#
//方法1
List<Task> tasks = new List<Task>();

foreach (var detail in comparisonDetailList)
{
    tasks.Add(ProcessDetail(detail, sdkDict, oldComparisonDetailList, runDetailList, accountsAutoList, shopIds));
}

await Task.WhenAll(tasks);

//方法2
SemaphoreSlim semaphore = new SemaphoreSlim(20); // 允许20个任务并发执行
List<Task> tasks = new List<Task>();

foreach (var detail in comparisonDetailList)
{
    tasks.Add(Task.Run(async () =>
    {
        await semaphore.WaitAsync();
        try
        {
            await ProcessDetail(detail, sdkDict, oldComparisonDetailList, runDetailList, accountsAutoList, shopIds);
        }
        finally
        {
            semaphore.Release();
        }
    }));
}

await Task.WhenAll(tasks);
```

#### Parallel.ForEach

> 密集CPU计算性能高，网络请求，加载文件不推荐，推荐算法，计算，渲染逻辑
>
> [官方说明](https://learn.microsoft.com/zh-cn/previous-versions/msp-n-p/ff963552(v=pandp.10))

```c#

ParallelOptions options = new ParallelOptions();
//支持4个并发执行
options.MaxDegreeOfParallelism = 4; //可以用Environment.ProcessorCount 获取当前进程可用的处理器数量
Parallel.ForEach(proSpus, options,  spu =>
{
 {
   SyncChangeProduct(spu, accounts); //传入Spu，与accounts商品信息
 }
});

//如果是循环也可以直接使用Linq版本的ForAll
IEnumerable<MyObject> myEnumerable = ...

myEnumerable.AsParallel().ForAll(obj => DoWork(obj));
```

#### Task.StartNew(不分页)

> 简单粗暴，直接数据库几个task,就StartNew多少个任务，不过不好控制不推荐使用，也可以通过take控制数据库单次查询数，达到控制线程数量的目的，也可以用下面分页控制线程数量的方法

```c#
/// <summary>
/// 开始同步对比价格
/// </summary>
/// <returns></returns>
public void SyncTaskComparisonPrice()
{
    var tasks = DBClient.Queryable<AdjustPriceTaskConfig>().Where(u => !u.IsDel
             && (u.TaskType == AdjustPriceTaskTypeEnum.CleanStock
                 || u.TaskType == AdjustPriceTaskTypeEnum.RegularAdjustPrice
                 || u.TaskType == AdjustPriceTaskTypeEnum.NewProductAdjustPrice)
             && u.Platforms.Contains("Walmart")
             && u.ComparisonFlag == 0
             && u.CreateTime >= sevenDaysAgo)
         .ToList();

    foreach (var task in tasks)
    {
        System.Threading.Tasks.Task.Factory.StartNew(() => ComparisongPriceTaskMq(task));
        System.Threading.Thread.Sleep(2000);
    }
}
```



#### Task.StartNew(分页)

> 对任务明细进行分页，进行并发处理，适合数据量大，没有分组属性的项目   

```c#
    
public void DelSkuNew()
        {
            var shopList = _platformAccountService.GetAccounts((int)PlatformEnum.Walmart);
            var account=shopList.FirstOrDefault(s=>s.AccountAlias=="WLU");
            var dels = DBClient.Queryable<SpWalmartListingDel>().Where(l => l.IsDelete == false && l.ShopName == "WLU").Take(24000).ToList();
            int count = dels.Count / 60;
            if (count == 0) count = 1;
            var comShopGroup = dels.ToArray().SplitArray(count);
            foreach (var comShops in comShopGroup)
            {
                System.Threading.Tasks.Task.Factory.StartNew(() => DelSkuNew(comShops, account));
                System.Threading.Thread.Sleep(200);
            }
        }

        public void DelSkuNew(SpWalmartListingDel[] dels, PlatformAccount account)
        {
            var authObj = JsonConvert.DeserializeObject<dynamic>(account.AuthJson);
            string accessToken = authObj.accessToken;
            WalmartSdk sdk = new WalmartSdk(account.ClientId, account.ClientSecret, accessToken);
            foreach (var item in dels)
            {
                try
                {
                    var res = sdk.RetireAnItem(item.Sku);
                    if (res == null) continue;
                    if (res.Error != null)
                    {
                        item.Message = res.Error[0].Info;
                    }
                    item.IsDelete = true;
                    item.DeleteDate = DateTime.Now;
                    if (!string.IsNullOrEmpty(item.Message) && item.Message.Contains("Unauthorized token"))
                        break;
                    DBClient.Updateable(item).ExecuteCommand();
                }
                catch(Exception ex) { 
                    Console.WriteLine(ex.Message);
                }
            }
        }

    public static class ArrayExtend
    {
        public static List<T[]> SplitArray<T>(this T[] data, int size)
        {
            if (data == null)
            {
                return null;
            }
            List<T[]> list = new List<T[]>();
            for (int i = 0; i < data.Length / size; i++)
            {
                T[] r = new T[size];
                Array.Copy(data, i * size, r, 0, size);
                list.Add(r);
            }
            if (data.Length % size != 0)
            {
                T[] r = new T[data.Length % size];
                Array.Copy(data, data.Length - data.Length % size, r, 0, data.Length % size);
                list.Add(r);
            }
            return list;
        }
    }
```

#### 线程池ThreadPool

> 适合多个短小精悍方法的同步，线程池可以重用线程，达到节省CPU与内存开销的作用，并且可以方便控制线程的数量,Sunway.Task.Win项目中

```c#
//线程池中声明8个线程，ActionExecute是需要执行的方法
private readonly Common.ThreadPool pool = new Common.ThreadPool(8, ActionExecute);

        public void ThreadAction(object obj)
        {
            try
            {
                // 得到所有账号
                List<PlatformAccount> accounts = PlatformAccountService.Instance.GetAccountList((int)PlatformEnum.Fruugo);
                foreach (var item in accounts)
                {
                    pool.AddNewAction(item);
                    //ActionExecute(item); //测试使用
                }
            }
            catch (Exception ex)
            {
                LoggerUtility.WriteLog($"初始化异常", ex.Message);
            }
        }
        public static void ActionExecute(PlatformAccount account)
        {
            SyncListing(account);
        }

//Common.ThreadPool类
   public  delegate void DoThreadAction(PlatformAccount item);

    public class ThreadPool
    {
        private int threadNum = 10;
        private List<Thread> threads = new List<Thread>();
        private ThreadActivity activity = new ThreadActivity();
        private bool running = false;
        private object lockObj = new object();
        

        private DoThreadAction action;

        public ThreadPool(int threadNum, DoThreadAction action)
        {
            this.threadNum = threadNum;
            this.action = action;
            this.running = true;

            // 创建线程
            for (int i = 0; i < this.threadNum; i++)
            {
                Thread thread = new Thread(new ParameterizedThreadStart(ThreadAction));
                thread.IsBackground = true;
                thread.Priority = ThreadPriority.Normal;
                threads.Add(thread);
                thread.Start();
            }
        }

        public void StopAll()
        {
            lock (lockObj)
            {
                this.running = false;
            }
            for (int i = 0; i < this.threadNum; i++)
            {
                Thread thread = this.threads[i];
                thread.Interrupt();
            }

            threads.ForEach(x => x.Join());
            threads.ForEach(x => x.Abort());
            threads = null;
        }

        public void AddNewAction(PlatformAccount item)
        {
            activity.Add(item);
        }

        private void ThreadAction(object obj)
        {
            while (true) 
            {
                lock (lockObj) {
                    if (!running) break;
                }

                PlatformAccount item = activity.NextItem();
                if (null != item)
                {
                    this.action(item);
                    activity.Remove(item);
                }
                else {
                    Thread.Sleep(30 * 1000);
                }
            }
        }
    }
```



###  SqlSugar线程安全问题CopyNew()

**问题描述**：保存数据的时候，用多线程DBClient.Utilities.PageEach，提示 1.Object reference not set to an instance of an object. 2.Cannot Open when State is Connecting

**问题提示**： 1.  HResult=0x80004003
  Message=Object reference not set to an instance of an object.
  Source=SqlSugar   2.Error 保存数据出现异常：中文提示 :  连接数据库过程中发生错误，检查服务器是否正常连接字符串是否正确，错误信息：Cannot Open when State is Connecting.DbType="MySql";ConfigId="".
English Message : Connection open error . Cannot Open when State is Connecting.DbType="MySql";ConfigId=""

**问题解决：**

1.**更改DB获取**:

![1691374429112](Vue进步.assets/1691374429112.png)



![1691374469502](Vue进步.assets/1691374469502.png)

2.**CopyNew**:

![1691374507640](Vue进步.assets/1691374507640.png)

3.**SqlSugarScope **:未验证

![1691374548564](Vue进步.assets/1691374548564.png)

### 如何使用ApacheBench工具进行压测实现一秒破万的并发？

> ApacheBench：简单易用，可以迅速获得服务器的性能指标，如请求处理速度（每秒请求数），服务器的连接时间，传输速率

```mysql
#安装ApacheBench：你可以使用以下命令来安装ApacheBench：在CentOS系统中，ApacheBench是作为httpd-tools的一部分来提供的，可以通过yum包管理器来进行安装。
sudo yum update httpd-tools
sudo yum install httpd-tools
#验证安装
ab -V
#压测：对http://localhost:5000/api进行压力测试，发送总共5000个请求，每次并发100个请求，你可以使用以下命令：
ab -n 5000 -c 100 http://192.168.226.128:7078/WeatherForecast 

ApacheBench是一个基础的压测工具，优点是简单好用，如果需要更深入的分析或者更复杂的压力测试，考虑使用JMeter
```

> **Requests per second**: 2637.22。这是服务器的吞吐量，也就是每秒可以处理的请求数。

![1691224367723](Vue进步.assets/1691224367723.png)

### SqlSugar爆炸多线程十并发操作数据

> 并发推荐 SqlSugarScope 单例模式下db.CopyNew()使用，看案例GetExamplesNameScopeNewSingletionThread这个方法,其实GetExamplesNameScopeSingletionThread也行，但是官方推荐db.CopyNew()这个方法，其他方法不行的原因是多个线程同时在打开/关闭同一个数据库链接会有异常，不允许这样操作。
>
> Net6+SqlSugarCore

```c#
using Microsoft.AspNetCore.Mvc;
using SqlSugar;

namespace SqlSugarThread.Controllers
{
    /// <summary>
    /// 1.GetExamplesNameSingletionThread 不支持并发，支持单例
    /// 2.GetExamplesNameModelThread 不支持并发
    /// 3.GetExamplesNameNoSingletionThread 可以并发，但是官方不推荐，非线程安全
    /// 4.GetExamplesNameScopeSingletionThread 可以并发，但是官方不推荐，官方用第五种db.CopyNew()
    /// 5.【并发推荐 GetExamplesNameScopeNewSingletionThread 推荐用此类并发，单是不是无限制用，可能会爆内存】
    /// </summary>
    [ApiController]
    [Route("[controller]/[action]")]
    public class TextController : ControllerBase
    {
        private readonly ISqlSugarRepository<example_table> _example_table;
        public TextController(ISqlSugarRepository<example_table> example_table)
        {

            _example_table = example_table;
        }
        int count = 10;
        [HttpGet]
        public async Task GetExamplesNameSingletionThread()
        {
            var db = SqlSugarSingletionClient.GetInstance();
            var tasks = Enumerable.Range(0, count).Select(u => db.Queryable<example_table>().Take(100).ToListAsync()).ToArray();
            var listResult = await Task.WhenAll(tasks);
            var result = listResult.SelectMany(u => u).ToList();
            foreach (var item in result)
            {
                Console.WriteLine($"ScopeSingletionClient.GetInstance Name：{item.name}");
            }

        }
        [HttpGet]
        public async Task GetExamplesNameModelThread()
        {
            var tasks = Enumerable.Range(0, count).Select(u => _example_table.Context.Queryable<example_table>().Take(100).ToListAsync()).ToArray();
            var listResult = await Task.WhenAll(tasks);
            var result = listResult.SelectMany(u => u).ToList();
            foreach (var item in result)
            {
                Console.WriteLine($"_example_table.Context.Queryable Name：{item.name}");
            }

        }

        [HttpGet]
        public async Task GetExamplesNameNoSingletionThread()
        {
            var tasks = Enumerable.Range(0, count).Select(u =>
            {
                var db = new SqlSugarClient(new ConnectionConfig()
                {
                    ConnectionString = "Server=127.0.0.1;port=3306;database=mydb;uid=root;password=123456;Convert Zero Datetime=True;Allow User Variables=True;pooling=true",
                    DbType = DbType.MySql,
                    IsAutoCloseConnection = true
                });
                return db.Queryable<example_table>().Take(100).ToListAsync();
            }).ToArray();
            var listResult = await Task.WhenAll(tasks);
            var result = listResult.SelectMany(u => u).ToList();
            foreach (var item in result)
            {
                Console.WriteLine($"New_Client_Thread.GetInstance Name：{item.name}");
            }
        }
        [HttpGet]
        public async Task GetExamplesNameScopeSingletionThread()
        {
            var tasks = Enumerable.Range(0, count).Select(u =>
            {
                var db = new SqlSugarScope(new ConnectionConfig()
                {
                    ConnectionString = "Server=127.0.0.1;port=3306;database=mydb;uid=root;password=123456;Convert Zero Datetime=True;Allow User Variables=True;pooling=true",
                    DbType = DbType.MySql,
                    IsAutoCloseConnection = true
                });
                return db.Queryable<example_table>().Take(100).ToListAsync();
            }
            ).ToArray();
            var listResult = await Task.WhenAll(tasks);
            var result = listResult.SelectMany(u => u).ToList();
            foreach (var item in result)
            {
                Console.WriteLine($" New_Scope_Thread.GetInstance Name：{item.name}");
            }
        }

        [HttpGet]
        public async Task GetExamplesNameScopeNewSingletionThread()
        {
            var db = SqlSugarSingletionClient.GetInstance();

            var tasks = Enumerable.Range(0, count).Select(u =>
            {
                return db.CopyNew().Queryable<example_table>().Take(100).ToListAsync();
            }
            ).ToArray();
            var listResult = await Task.WhenAll(tasks);
            var result = listResult.SelectMany(u => u).ToList();
            foreach (var item in result)
            {
                Console.WriteLine($" New_Scope_Thread.GetInstance Name：{item.name}");
            }
        }
    }
    /// <summary>
    /// 单例模式SqlSugar实例,双重检查锁定模式,适用于多线程场景
    /// </summary>
    public class SqlSugarSingletionClient
    {
        private static SqlSugarScope? instance;
        private static readonly object _lock = new object();

        /// <summary>
        /// 私有的构造函数，防止在外部被创建
        /// </summary>
        private SqlSugarSingletionClient() { }

        public static SqlSugarScope GetInstance()
        {
            if (instance == null)//步骤1：进来先看看有没有没有再加锁，防止不必要的性能损耗
            {
                lock (_lock) //步骤2：加锁
                {
                    if (instance == null) //步骤3：因为多个线程一般都卡在步骤2等lock释放锁，这里加instance == null的判断是为了，如果上一个A线程创建好了实例对象走了，释放了_lock，后面的B线程进来判断一下，如果实例不为空直接拿实例走就好，没必要再进行创建了，如果B再创建一个对象，就不是单例模式了
                    {
                        instance = new SqlSugarScope(new ConnectionConfig()
                        {
                            ConnectionString = "Server=127.0.0.1;port=3306;database=mydb;uid=root;password=123456;Convert Zero Datetime=True;Allow User Variables=True;pooling=true",
                            DbType = DbType.MySql,
                            IsAutoCloseConnection = true
                        });
                    }
                }
            }

            return instance;
        }

    }
}

```



### 再次震惊的处理速度Parallel.ForEach百万数据1分钟

> 比下面的更加让我惊艳：
>
> 1.Environment.ProcessorCount当spus.Count小于这个数的时候注意 
>
> 2.productPropertyDict.TryGetValue 字典筛选性能比Where好一百万倍 （数据库越大，性能约好，过百万的话）
>
> 3.timer.Elapsed时间的妙用，绝了

```c#
        /// <summary>
        /// 根据Sku转换所有款号的销售状态
        /// </summary>
        public void ConvertProductSalestatus()
        {
            Stopwatch watch = new Stopwatch();
            watch.Start();
            Console.WriteLine($"开始转换销售状态:{DateTime.Now}");
            var saleStatusList = GetOptionDictionary().Result.Where(u => u.DicName == "产品销售状态").ToList();
            var outSaleStatusStr = GetOptionDictionary().Result.Where(u => u.DicName == "销售状态转换产品状态").Where(u => u.Value == "不可用").FirstOrDefault()?.Label;
            var outSaleStatusList = GetValueByDictionary(saleStatusList, outSaleStatusStr)?.Split(',').Select(u => u?.ToInt32()).ToList();
            // 从数据库中一次性加载所有需要处理的数据
            var spus = _productRepository.Context.Queryable<TProduct>()
                .Where(u => u.IsDeleted == false)
                .Select(u => new TProduct { Id = u.Id, ProductSpu = u.ProductSpu })
                .ToList();
            var skus = _productPropertyRepository.Context.Queryable<TProductProperty>()
                .Where(u => u.IsDeleted == false)
                .Select(u => new TProductProperty { Id = u.Id, ProductSpu = u.ProductSpu, SaleStatus = u.SaleStatus })
                .ToList();
            var productPropertyDict = skus.GroupBy(p => p.ProductSpu).ToDictionary(g => g.Key, g => g.ToList());
            // 初始化已处理的计数为0
            int processedCount = 0;
            // 计算报告阈值
            int reportThreshold = spus.Count / 10;
            // 创建一个定时器，每隔一段时间检查进度并打印
            System.Timers.Timer timer = new System.Timers.Timer(10000); // 每10秒检查一次
            timer.Elapsed += (sender, e) =>
            {
                // 计算进度
                double progress = processedCount / (double)spus.Count * 100;
                Console.WriteLine($"当前已经完成百分之{progress}%");
            };
            timer.Start();
            if (spus.Count > Environment.ProcessorCount)
            {
                int parallelDegree = Environment.ProcessorCount;
                Parallel.ForEach(Partitioner.Create(0, spus.Count, spus.Count / parallelDegree), range =>
                {
                    for (int i = range.Item1; i < range.Item2; i++)
                    {
                        var spu = spus[i];
                        if (productPropertyDict.TryGetValue(spu.ProductSpu, out var matchedProductPropertys))
                        {
                            ConvertProductSaleStatus(outSaleStatusList, spu, matchedProductPropertys);
                        }
                        // 增加已处理计数
                        Interlocked.Increment(ref processedCount);
                    }
                });
            }
            else
            {
                foreach (var spu in spus)
                {
                    if (productPropertyDict.TryGetValue(spu.ProductSpu, out var matchedProductPropertys))
                    {
                        ConvertProductSaleStatus(outSaleStatusList, spu, matchedProductPropertys);
                    }
                }
            }
            // 停止定时器
            timer.Stop();

            // 筛选出需要更新的数据并批量更新
            if (spus.Any())
            {
                const int batchSize = 2000;
                for (int i = 0; i < spus.Count; i += batchSize)
                {
                    var batchToUpdate = spus.Skip(i).Take(batchSize).ToList();
                    _productRepository.Context.Updateable(batchToUpdate)
                        .UpdateColumns(it => new { it.SaleStatus })
                        .ExecuteCommand();
                }

            }
            Console.WriteLine($"转换完成:{DateTime.Now},总耗时：{watch.ElapsedMilliseconds / 1000 / 60}分");
        }
```



### 惊艳我的Parallel.ForEach百万数据30秒

> 有时候要上线新功能的时候需要处理历史数据，如果不用并发，多线程的话预计要30分钟，并发加上一次性读取数据，汇总到一起提交，只需要30秒搞定100万数据
>

```c#
        
/// <summary>
        /// 转换历史数据，属于带电产品的数据
        /// </summary>
        public void ConvertAllProductIsBattery()
        {
            Stopwatch watch = new Stopwatch();
            watch.Start();
            Console.WriteLine($"开始转换带电产品:{DateTime.Now}");
            var dics = GetOptionDictionary().Result.Where(u => u.DicName == "是否含电池" || u.DicName == "易仓自定义分类").ToList();
            var selfCategoryDiclist = dics.Where(u => u.DicName == "易仓自定义分类").ToList();
            var batterySelfCategoryStr = dics.Where(u => u.DicName == "是否含电池").FirstOrDefault()?.Label;
            var batterySelfCategoryList = GetValueByDictionary(selfCategoryDiclist, batterySelfCategoryStr)?.Split(',').Select(u => u?.ToInt32()).ToList();

            // 从数据库中一次性加载所有需要处理的数据
            var skus = _productPropertyRepository.Context.Queryable<TProductProperty>()
                .Where(u => u.SelfCategoryList != null)
                .Select(u => new TProductProperty { Id = u.Id, SelfCategoryList = u.SelfCategoryList })
                .ToList();

            // 并行处理数据，根据 CPU 核心数调整并行度
            int parallelDegree = Environment.ProcessorCount;
            Parallel.ForEach(Partitioner.Create(0, skus.Count, skus.Count / parallelDegree), range =>
            {
                for (int i = range.Item1; i < range.Item2; i++)
                {
                    var sku = skus[i];
                    if (!sku.SelfCategoryList.IsNullOrWhitespace())
                    {
                        sku.IsBattery = ConvertProductBattery(sku.SelfCategoryList, batterySelfCategoryList);
                    }
                }
            });

            // 筛选出需要更新的数据并批量更新
            var skusToUpdate = skus.Where(u => u.IsBattery == true).ToList();
            if (skusToUpdate.Any())
            {
                const int batchSize = 2000;
                for (int i = 0; i < skusToUpdate.Count; i += batchSize)
                {
                    var batchToUpdate = skusToUpdate.Skip(i).Take(batchSize).ToList();
                    _productPropertyRepository.Context.Updateable(batchToUpdate)
                        .UpdateColumns(it => new { it.IsBattery })
                        .ExecuteCommand();
                }

            }
            Console.WriteLine($"转换完成:{DateTime.Now},总耗时：{watch.ElapsedMilliseconds / 1000 / 60}分");
        }
```



### Parallel.ForEach 并发操作数据库问题

解决：**计算逻辑用多线程，数据库提交用单线程一次提交**

比如Real平台的下架功能，平台没有提供批量下架的接口，我的解决方案就是通过整合目前的需求，通过Paraller并发50个线程+Bulk批量处理，将原来需要5天推送10万的性能，缩短至10个小时，性能提升了10倍。

所以**计算逻辑用多线程，数据库提交用单线程一次提交**，并且其中注意要问自己如果是大批量的数据流量进入，数据库一次读取的数据是多少（我一般控制1000读取一次），加上Take数量限制的前提下，还可以加入Select限定指定的属性，如果再加上Select（属性覆盖索引）+查询条件有索引，那性能更是爆表。还有一点就是并发循环的时候最好用主表做为循环的上游，下游用主表的ID作为查询条件然后while+take循环，这样非常好操作，比如款号表+SKU表用这种没事，Config下架表与Run明细表用这种模式循环，也很好操作里面的数据。



```c#
        ParallelOptions options = new ParallelOptions();
        //支持4个并发执行
        options.MaxDegreeOfParallelism = 4;
        Parallel.ForEach(proSpus, options, async spu =>
         {
             {
                 await SyncChangeProduct(spu, accounts);
             }
         });
```

![1653012094396](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1653012094396.png)

#### async/await，Task,Parallel.For使用技巧

**结论：**  I/O-bound任务用 异步方法，日常推荐用：`async`/`await` ，Task.WhenAll。CPU-bound密集型任务（计算型任务）用Parallel.For 和 Parallel.ForEach, I/O-bound型任务（网络请求，读写硬盘）用Task.Run 和 Task.Factory.StartNew（Task.Run 和 Task.Factory.StartNew都会新建线程消耗性能很大，一般使用async /await使用异步操作）

**官方资料**：https://learn.microsoft.com/zh-cn/dotnet/csharp/asynchronous-programming/async-scenarios

```c#
public static async Task ConcurrentCall()
{
    var tasks = Enumerable.Range(0, 3).Select(_ => GetEans()).ToArray();
    var results = await Task.WhenAll(tasks);
}

```

- `【async`/`await】`： 日常I/O-bound异步使用这个，不会创建新的线程，并且不会阻塞主线程
- `Task.WhenAll` ：多个异步任务用这个处理，创建多个Task对象，放进来就可以，不会阻塞主线程，除非你Task.Wait()
- `Task.Run` ：创建新的线程异步，具有IO密集+CPU密集的任务，比 `Task.Factory.StartNew`  相对于简单一些，推荐日常使用  
- `Task.Factory.StartNew` ：创建新的线程异步，具有IO密集+CPU密集的任务，与`Task.Run`的区别在于可以控制取消，线程优先级
- `Parallel.For 和 Parallel.ForEach` ：耗时长的CPU密集操作，无或者少IO密集操作（如果有IO密集HTTP请求，又有CPU计算需求可以用Task.Run{ Parallel.For} +Task.WhenAll进行处理）
- `ThreadPool ` ：耗时短的CPU密集操作，比如写入日志，推送消息

四种并发方法的分析：

1. `Task.Run`： 

   **描述**：它在默认情况下总是在线程池上调度一个任务。此方法在处理CPU-bound工作（需要大量计算的任务）或在确保不阻塞主线程时非常有用。这是因为它在后台创建一个新的线程来处理工作。

   **案例**：假设你有一个电商网站，你希望在每个用户完成购买后生成和发送一个包含其购买详情的报告。报告生成是CPU密集型的操作，因为它需要收集和处理大量数据。在这种情况下，你可以使用`Task.Run`来在后台线程中生成报告，这样就**不会阻塞主线程**，用户仍然可以继续在你的网站上浏览。

   **优点：**对于CPU-bound的工作，可以有效地利用多核处理器的计算资源，提高应用程序的响应速度。 

   **缺点**：对于I/O-bound的工作，`Task.Run`的效率较低，因为它无法有效地利用异步I/O的优点。

2. `Task.Factory.StartNew`： 

   **描述**：这个方法更加灵活，它提供了更多的选项，如任务调度、任务创建选项（长时间运行、偏好公平调度等）等。然而，它的使用相对复杂，对于简单的场景（如在线程池上运行一个操作），更推荐使用`Task.Run`。

   **案例**：如果你需要更高级的控制，如任务的取消，延迟或优先级调度，这种方法可能会有所帮助。例如，如果你有一个重要的后台任务需要立即运行，但已经有很多任务在等待运行，你可以使用`Task.Factory.StartNew`设置高优先级来确保这个任务优先执行。

   **优点**：提供了对任务的更多控制，例如，可以设置任务的优先级，或者取消任务。 

   **缺点**：使用复杂，很容易被误用。例如，如果忽略了默认的任务调度器，可能会在不恰当的线程上执行任务。

3. `Parallel.For，Parallel.ForEach`：

    **适用场景**：`Parallel.For`和`Parallel.ForEach`则是专门设计用于并行处理的情况，它们会将操作并行分配到所有可用的CPU核心上。它们用于处理CPU-bound任务，因为它们可以将任务分割到多个线程中，并行执行以达到最快的处理速度。 

   **案例**：假设你需要分析大量的产品数据，例如，你可能需要计算一系列产品的平均评分。这是一个可以并行的CPU密集型任务，因为每个产品的计算都是独立的。在这种情况下，`Parallel.For`可以被用来并行执行计算，从而充分利用多核处理器。

   **优点**：`Parallel.For`能自动地将循环的迭代分配给线程池中的多个线程，从而充分利用多核处理器的计算资源。

    **缺点**：`Parallel.For`并不支持异步操作。在你的示例中，`GetEans`是一个异步方法，但是在`Parallel.For`的循环体中，异步方法的执行并不会被正确地等待。

4. `ThreadPool`： 

   **适用场景**：当你需要执行大量的并发工作时，可以使用`ThreadPool`。`ThreadPool`会自动地管理线程的创建和销毁，使得线程可以被复用，减少了线程创建和销毁的开销。 

   **案例**：如果你的电商平台有大量的图片需要处理（例如，缩放或裁剪），那么ThreadPool是个不错的选择。ThreadPool可以自动处理线程创建和销毁的细节，对于大量并发工作非常有用。然而，因为它无法有效地利用异步I/O，所以对于I/O-bound的工作（例如，从远程服务器下载图片），它的效率可能较低。

   **优点**：自动地管理线程，使得线程可以被复用，减少了线程创建和销毁的开销。 

   **缺点**：对于长时间运行的任务，如果直接在线程池中运行，可能会阻塞线程池中的其他工作。对于I/O-bound的工作，使用`ThreadPool`的效率较低，因为它无法有效地利用异步I/O的优点。

在大多数场景下，如果你的任务是CPU-bound的，你应该使用`Parallel.For`或`Parallel.ForEach`。如果你的任务是I/O-bound的，或者你需要创建一些相对独立的后台任务，你应该使用`Task.Run`或`Task.Factory.StartNew`。

在某些情况下，你可能需要更细粒度的控制，例如任务调度、任务优先级、取消任务等，这时你可能会选择使用`Task.Factory.StartNew`。但是，由于这个方法的使用相对复杂，并且有可能被误用，所以在没有特殊需求的情况下，一般推荐使用`Task.Run`。

最后需要注意的是，`Task.Run`和`Task.Factory.StartNew`主要是用于异步执行任务，而`Parallel.For`和`Parallel.ForEach`则是用于并行执行任务。异步执行主要是为了提高应用程序的响应性，而并行执行则是为了利用多核处理器的计算能力。

在你的情况中，你正在执行的是一个I/O-bound的工作（HTTP请求），并且这个工作是异步的。所以，你应该使用异步方法，并且尽量避免使用`Task.Run`，`Task.Factory.StartNew`，`Parallel.For`和`ThreadPool`。你可以使用`Task.WhenAll`来并发执行多个异步操作。用上面那个案例的结果。

**问题**：**Parallel.For和Parallel.ForEach则是专门设计用于并行处理的情况 ，为什么不用这两个，要用Task.Run和Task.Factory.StartNew**

`Parallel.For`和`Parallel.ForEach`确实是专门设计用于并行处理的情况，这两个方法能够自动将迭代的工作分配给多个线程，从而充分利用多核处理器的计算能力。这对于CPU-bound任务来说是非常有效的，尤其是当这些任务可以自然地划分为许多独立的、可以并行执行的子任务时。

然而，`Parallel.For`和`Parallel.ForEach`并不适用于所有的并行处理场景，特别是以下两种情况：

1. **异步任务**：`Parallel.For`和`Parallel.ForEach`并不支持异步操作，如果在`Parallel.For`或`Parallel.ForEach`的循环体中使用异步方法，这些异步方法的执行将不会被正确地等待。这就意味着，如果你的任务是I/O-bound的，例如网络请求或文件操作，使用`Parallel.For`或`Parallel.ForEach`就会有问题，因为这些操作通常需要使用异步方法。
2. **长时间运行的任务**：如果你有一些可能会运行很长时间的任务，那么直接在`Parallel.For`或`Parallel.ForEach`中运行这些任务可能会阻塞其他并行任务的执行。这是因为`Parallel.For`和`Parallel.ForEach`默认会在所有的并行任务完成后才返回，这意味着，如果有一个任务运行时间很长，那么它将阻塞`Parallel.For`或`Parallel.ForEach`的返回。

而`Task.Run`和`Task.Factory.StartNew`则更加灵活，它们不仅支持异步操作，而且可以自己控制任务的完成状态。因此，对于一些需要更精细控制的场景，或者需要异步操作的场景，通常会选择使用`Task.Run`或`Task.Factory.StartNew`。

**问题**：**Task.Run`或`Task.Factory.StartNew`的使用场景是什么**

`Task.Run`或`Task.Factory.StartNew`主要适用于CPU-bound任务，也就是任务的主要工作在于执行计算或处理数据，这种工作主要花费在CPU上的时间。以下是一些使用这两个方法的场景：

1. **长时间运行的计算密集型任务**：例如图像处理，复杂的数学计算等，这些任务通常需要占用大量的CPU时间，通过`Task.Run`或`Task.Factory.StartNew`启动这样的任务，可以将任务放到后台线程中运行，避免阻塞主线程，提高应用程序的响应性。
2. **需要并行处理的任务**：对于那些可以并行处理的任务，例如处理一个大的数据集，可以将任务分解为多个小任务，然后使用`Task.Run`或`Task.Factory.StartNew`将这些小任务并行的放到线程池的多个线程中运行，这样可以充分利用多核处理器的计算资源，提高计算效率。
3. **需要转移的任务**：有时候，为了避免阻塞当前线程，可能需要将任务转移到另一个线程上执行。例如在UI线程中，不能执行长时间运行的任务，否则会导致UI失去响应，这时就可以使用`Task.Run`或`Task.Factory.StartNew`将任务转移到后台线程执行。

请注意，对于I/O-bound任务，比如文件操作，网络请求等，应当使用.NET的异步API（即以`Async`结尾的方法），这样可以避免阻塞线程，提高系统的效率。

**问题：ThreadPool有适用的场景吗，上面IO,CPU密集的任务，异步同步的任务都没有ThreadPool的作用，被async/await,Task.WhenAll,Parallel.For,ForEach平分了**

`ThreadPool`（线程池）是用于在.NET应用程序中处理并发任务的一个有效工具。它为在后台处理任务提供了一个简单的方式，同时管理了线程的生命周期。

以下是`ThreadPool`适用的一些常见场景：

1. **短时间的后台任务**：`ThreadPool`非常适合处理短时间运行的后台任务，因为这样可以避免频繁地创建和销毁线程，从而减少了开销。
2. **处理大量并发的工作**：当你需要执行大量的并发工作时，可以使用`ThreadPool`。`ThreadPool`会自动地管理线程的创建和销毁，使得线程可以被复用，减少了线程创建和销毁的开销。
3. **计划定时任务**：`ThreadPool`还可以用于执行定时任务。你可以使用`ThreadPool.RegisterWaitForSingleObject`方法来安排一个方法在特定的时间间隔后被调用。

然而，和其他并发处理工具一样，`ThreadPool`并不适用于所有的并发处理场景。如果任务的执行时间较长，或者任务需要频繁地等待一些资源（例如，I/O操作，网络请求等），那么直接在线程池中运行这些任务可能会导致线程池中的线程被阻塞，影响线程池的性能。

此外，`ThreadPool`也并不支持异步操作，因此如果任务是I/O-bound的，或者涉及到网络请求或文件操作等异步操作，你应该考虑使用异步编程模型，例如使用`Task.Run`，`Task.Factory.StartNew`或`Task.WhenAll`等方法。

**问题:为什么不用Parallel.For和Parallel.ForEach 用ThreadPool相对于有什么优点和缺点**

`Parallel.For` 和 `Parallel.ForEach` 是专门为数据并行设计的，它们会在多个线程上执行循环体中的操作。这在处理 CPU-bound 工作（例如计算密集型任务）时非常有用，因为它可以充分利用多核处理器的计算能力。

然而，`Parallel.For` 和 `Parallel.ForEach` 的一个主要限制是，它们不支持异步操作。也就是说，如果你需要在循环体中执行 I/O-bound 的操作（例如网络请求、文件读写等），并且这些操作是异步的，那么 `Parallel.For` 和 `Parallel.ForEach` 就无法正确地等待这些异步操作的完成。

`ThreadPool` 不同于 `Parallel.For` 和 `Parallel.ForEach`，它更像一个通用的并发处理工具。`ThreadPool` 会自动地管理线程的创建和销毁，从而使线程可以被复用。**这对于处理大量短时间运行的后台任务来说非常有用(大量短时间后台执行的任务，使用ThreadPool是比Parallel.For,Parallel.ForEach性能更好的，比如写入日志，消息推送，短，快的适合ThreadPool，Parallel.For,Parallel.ForEach适合并行，但是慢，没有异步的IO)**，因为它可以避免频繁地创建和销毁线程，从而减少了开销。

然而，`ThreadPool` 也有其限制。如果任务的执行时间较长，或者任务需要频繁地等待一些资源（例如，I/O操作，网络请求等），那么直接在线程池中运行这些任务可能会导致线程池中的线程被阻塞，影响线程池的性能。此外，和 `Parallel.For`、`Parallel.ForEach` 一样，`ThreadPool` 也并不支持异步操作。

所以，选择使用哪种方法取决于具体的应用场景和任务类型。在处理 CPU-bound 工作时，可以考虑使用 `Parallel.For` 和 `Parallel.ForEach`。在处理 I/O-bound 工作，尤其是异步操作时，应该使用异步编程模型，例如使用 `Task.Run`，`Task.Factory.StartNew` 或 `Task.WhenAll` 等方法。如果需要处理大量的短时间运行的后台任务，可以考虑使用 `ThreadPool`。

**I/O-bound 和 CPU-bound 是两种用于描述计算任务特性的术语**。

1. **I/O-bound（输入输出限制）**：此类任务的执行效率主要受限于输入/输出操作，例如文件读写，网络请求等。这些任务在等待 I/O 完成的过程中，CPU 可能不会有太多的计算工作，可能会在大部分时间内处于空闲状态。对于这样的任务，提升 I/O 速度或使用异步操作能有效地改善性能，**比如关键字：Task.WhenAll，**。例如，从数据库查询数据、从网络下载文件等。
2. **CPU-bound（CPU限制）**：此类任务的执行效率主要受限于 CPU 的计算速度。这类任务在执行过程中，CPU 一直处于忙碌状态，对 CPU 计算能力的需求很高。对于这类任务，提升 CPU 速度或使用更有效的算法能有效地改善性能,**比如关键字：Task.Run 和 Task.Factory.StartNew ,Parallel.For 和 Parallel.ForEach，ThreadPool**。例如，数学计算、图像处理、数据分析等。

在进行并发编程时，了解任务的特性（是 I/O-bound 还是 CPU-bound）能帮助我们更好地选择合适的工具和方法。例如，对于 I/O-bound 的任务，我们通常会使用异步编程来避免浪费 CPU 资源在等待 I/O 完成时处于空闲状态。而对于 CPU-bound 的任务，我们则可能会使用多线程或者并行编程来充分利用多核 CPU 的计算能力。

**当在并发线程中操作同一个数组对象使用LOCK：**

```c#
            var listerEans4 = new List<string>();
            var listTask4 = new List<Task>();
            for (int i = 0; i < 3; i++)
            {
                ThreadPool.QueueUserWorkItem(async _ =>
                {
                    var eans = await GetEans();
                    lock (listerEans4)
                    {
                        listerEans4.AddRange(eans);
                    }
                });
            }
//为什么要用lock锁住呢，不锁住会有什么问题?
在并发编程中，当多个线程同时访问和修改同一块内存区域（例如，一个列表或数组）时，可能会出现数据被覆盖，写入数据少了，结果输出不对的混乱情况。这是因为一个线程在读取或修改数据的同时，另一个线程可能也在对同一数据进行操作。

在你的例子中，listerEans是一个列表，可能被多个线程同时添加元素。如果不使用lock关键字来确保线程安全，可能会出现以下情况：

1.两个线程几乎同时调用AddRange方法，一个线程读取了列表长度并计划在列表末尾添加元素，但在它完成添加操作之前，另一个线程也读取了同样的列表长度并尝试在同一位置添加元素。这种情况下，一个元素可能会覆盖另一个元素，导致数据丢失。

2.列表内部可能需要重新分配内存来容纳新的元素。如果两个线程同时触发这个操作，可能导致内存损坏，甚至引发异常。

lock关键字可以确保一次只有一个线程进入临界区（也就是被lock保护的代码块），从而避免上述问题。当一个线程进入临界区时，其他试图访问该临界区的线程将被阻塞，直到第一个线程退出临界区。这样可以确保每次只有一个线程在修改listerEans，从而避免数据竞争和数据不一致。
```

##### 下面是对你提供的四种并发方法的分析：

##### Task.Run

1. `Task.Run`： 适用场景：在.NET中，`Task.Run`方法常被用来将CPU-bound的工作转移到线程池中的一个线程上。 优点：对于CPU-bound的工作，可以有效地利用多核处理器的计算资源，提高应用程序的响应速度。 缺点：对于I/O-bound的工作，`Task.Run`的效率较低，因为它无法有效地利用异步I/O的优点

   ##### Task.Factory.StartNew

2. `Task.Factory.StartNew`： 适用场景：类似于`Task.Run`，这个方法常被用来将工作转移到线程池中的一个线程上。和`Task.Run`相比，`Task.Factory.StartNew`提供了更多的控制选项，例如任务调度和任务状态。 优点：提供了对任务的更多控制，例如，可以设置任务的优先级，或者取消任务。 缺点：使用复杂，很容易被误用。例如，如果忽略了默认的任务调度器，可能会在不恰当的线程上执行任务。

   ##### Parallel.For

3. `Parallel.For`： 适用场景：当你有一个可以并行执行的循环，并且每个循环的迭代都是CPU-bound的工作时，可以使用`Parallel.For`。 优点：`Parallel.For`能自动地将循环的迭代分配给线程池中的多个线程，从而充分利用多核处理器的计算资源。 缺点：`Parallel.For`并不支持异步操作。在你的示例中，`GetEans`是一个异步方法，但是在`Parallel.For`的循环体中，异步方法的执行并不会被正确地等待。

   ##### ThreadPool

4. `ThreadPool`： 适用场景：当你需要执行大量的并发工作时，可以使用`ThreadPool`。`ThreadPool`会自动地管理线程的创建和销毁，使得线程可以被复用，减少了线程创建和销毁的开销。 优点：自动地管理线程，使得线程可以被复用，减少了线程创建和销毁的开销。 缺点：对于长时间运行的任务，如果直接在线程池中运行，可能会阻塞线程池中的其他工作。对于I/O-bound的工作，使用`ThreadPool`的效率较低，因为它无法有效地利用异步I/O的优点。

在你的情况中，你正在执行的是一个I/O-bound的工作（HTTP请求），并且这个工作是异步的。所以，你应该使用异步方法，并且尽量避免使用`Task.Run`，`Task.Factory.StartNew`，`Parallel.For`和`ThreadPool`。你可以使用`Task.WhenAll`来并发执行多个异步操作。例如：

​     

```
public static async Task ConcurrentCall()
{
    var tasks = Enumerable.Range(0, 3).Select(_ => GetEans()).ToArray();
    var results = await Task.WhenAll(tasks);
    // 在这里处理结果
}

```

**解决方式：**

E神方案1：

![1653012201159](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1653012201159.png)



   ![1653012219373](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1653012219373.png)

![1653012253193](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1653012253193.png)

 方案2： 百度文档
   https://wenku.baidu.com/view/87b718c8adaad1f34693daef5ef7ba0d4a736ddb.html 
 方案3：  CSDN
 https://blog.csdn.net/xiaobaoTX/article/details/119828894

方案4：杨魁

![1655796935357](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1655796935357.png)

方案5：dotnet run 多个黑窗口（尝试失败，原因：无法访问新加的链接，提示:**localhost** 拒绝了我们的连接请求。）

dotnet run --project ./Sunway.Skums.Service.csproj --urls="https://*:5004"

方案6：使用不同链接进行调用，并发（尝试失败，提示：执行命令时遇到致命错误）



http://localhost:5003/ecproductpropertycompare/convertselfproperty
http://localhost:5003/ecproductpropertycompare/convertcolor
http://localhost:5003/ecproductpropertycompare/convertsize
http://localhost:5003/ecproductpropertycompare/convertselfcategory
http://localhost:5003/ecproductpropertycompare/convertsubsaleuserids

### Action与Func有什么区别

`Func` 和 `Action` 是 C# 中的两种内置委托类型，它们都可以封装方法，但是主要的区别在于，**`Func` 类型的委托有返回值，而 `Action` 类型的委托不返回值。**

详细来说：

- `Func`： `Func` 委托类型期望有一个返回值，并且它可以有 0 到 16 个输入参数。例如 `Func<int>` 表示一个无参数并返回整数的方法，而 `Func<int, string>` 表示一个接收整数并返回字符串的方法。
- `Action`： `Action` 委托类型不返回值，但是它可以有 0 到 16 个输入参数。例如，`Action` 表示一个无参数的方法，而 `Action<int, string>` 表示一个接收整数和字符串作为参数的方法。

以下是一个 `Func` 的例子：

```c#
Func<int, int, int> add = (x, y) => x + y; //最后一个参数是返回值的参数,有0~16个输入参数，一个返回值参数
int result = add(2, 3); // 结果是 5

//Task代表是异步的返回值
Func<Task> myFunc = async () =>
{
    await Task.Delay(1000);
    Console.WriteLine("Hello from Func!");
};

await myFunc();  // 调用异步的 Func

//Task有返回值的
Func<Task<int>> myFunc = async () =>
{
    await Task.Delay(1000);
    Console.WriteLine("Hello from Func!");
    return 42;  // 返回一个结果
};

int result = await myFunc();  // 调用异步的 Func
Console.WriteLine($"The result is {result}");

```

以下是一个 `Action` 的例子：

```c#
Action<int, int> printAdd = (x, y) => Console.WriteLine(x + y); //两个参数都是入参的参数,有0~16个参数
printAdd(2, 3); // 输出：5

//Task案例
Action myAction = () =>
{
    Task.Run(async () =>
    {
        await Task.Delay(1000);
        Console.WriteLine("Hello from Action!");
    });
};

myAction();  // 调用 Action，异步任务在后台运行

```

总的来说，你会根据你需要封装的方法是否有返回值来选择使用 `Func` 还是 `Action`。如果有返回值，你会使用 `Func`，如果没有，你会使用 `Action`。

### await Task.Delay与Thread.Sleep有什么区别

>  await Task.Delay 不会阻塞主线程，Thread.Sleep会阻塞主线程
>

```c#
Console.WriteLine($"第一种：Thread.Sleep，阻塞主线程,{DateTime.Now}");
ProcessTasks();
for (int i = 0; i < 5; i++)
{
    Console.WriteLine($"Thread.Sleep：{i},{DateTime.Now}");
}
Console.WriteLine("**************************");
Console.WriteLine($"第二种：Task.Delay，不会阻塞主线程，但是因为await关键字会等待,{DateTime.Now}");
await UpdateScreenAsync(); //由于使用了await关键字，所有主线程会等待await Task.Delay的处理结果（等25秒）后再输出，如果把await关键字拿掉，那么主线程就不会等待UpdateScreenAsync方法的处理结果，UpdateScreenAsync方法会马上执行
for (int i = 0; i < 5; i++)
{
    Console.WriteLine($"Task.Delay：{i+1},{DateTime.Now}");
}

void ProcessTasks()
{
    for (int i = 0; i < 5; i++)
    {
        // Do some processing...
        Console.WriteLine($"Processing task {i + 1}...,{DateTime.Now}");

        // Sleep for 5 seconds
        Thread.Sleep(5000);
    }
}


async Task UpdateScreenAsync()
{
    for (int i = 0; i < 5; i++)
    {
        // Update the screen...
        Console.WriteLine($"Updating screen {i + 1}...,{DateTime.Now}");

        // Wait for 5 seconds without blocking the UI thread
        await Task.Delay(TimeSpan.FromSeconds(5));
    }
}
输出结果：
第一种：Thread.Sleep，阻塞主线程,2023/7/24 15:00:38
Processing task 1...,2023/7/24 15:00:38
Processing task 2...,2023/7/24 15:00:43
Processing task 3...,2023/7/24 15:00:48
Processing task 4...,2023/7/24 15:00:53
Processing task 5...,2023/7/24 15:00:58
Thread.Sleep：0,2023/7/24 15:01:03
Thread.Sleep：1,2023/7/24 15:01:03
Thread.Sleep：2,2023/7/24 15:01:03
Thread.Sleep：3,2023/7/24 15:01:03
Thread.Sleep：4,2023/7/24 15:01:03
**************************
第二种：Task.Delay，不会阻塞主线程，但是因为await关键字会等待,2023/7/24 15:01:03
Updating screen 1...,2023/7/24 15:01:03
Updating screen 2...,2023/7/24 15:01:08
Updating screen 3...,2023/7/24 15:01:13
Updating screen 4...,2023/7/24 15:01:18
Updating screen 5...,2023/7/24 15:01:23
Task.Delay：1,2023/7/24 15:01:28
Task.Delay：2,2023/7/24 15:01:28
Task.Delay：3,2023/7/24 15:01:28
Task.Delay：4,2023/7/24 15:01:28
Task.Delay：5,2023/7/24 15:01:28

```



### Task.WhenAny配合White的妙用

所有的下载任务在 `Task.WhenAny(downloadTasks)` 之后会同时进行。`Task.WhenAny(downloadTasks)` 只是返回最先完成的任务，但这并不会停止或者暂停其他的任务。这就是我们能够在循环中移除已完成的任务并等待下一个完成的任务，而其他的任务在此期间还在继续执行。

所以，你可以把 `Task.WhenAny` 看作是一种方式，让你可以在多个任务同时进行的时候，立即处理那些已经完成的任务，而不需要等待所有的任务都完成。

> 相对于await Task.WhenAll(downloadTasks); 这个是完成所有任务之后才能开始后续的处理，Task.whenAny可以让我们体会到异步并发的效率，同时又可以马上去将结果保存到数据库

```c#
public async Task DownloadWebPagesAsync(List<string> urls)
{
    var downloadTasks = urls.Select(url => DownloadWebPageAsync(url)).ToList(); // start all downloads

       // 只要还有未完成的任务就继续循环
    while (downloadTasks.Any())
    {
        // 等待任何一个任务完成，并获取这个已完成的任务
        var downloadTask = await Task.WhenAny(downloadTasks); 
        // 从任务列表中移除已经完成的任务
        downloadTasks.Remove(downloadTask); 
        // 获取已完成任务的结果
        var webPageContent = await downloadTask; 
        // 处理已完成任务的结果
        ProcessWebPageContent(webPageContent);
    }
}

```

## NetCore

#### 架构知识体系

[课程目录](https://wechat.zhaoxiedu.net/#/my/mycourse/vip)

![1693473024747](Vue进步.assets/1693473024747.png)

![1693467875725](Vue进步.assets/1693467875725.png)

#### 高级班知识体系

[课程链接](https://ke.qq.com/course/453949#term_id=106163181)

![1693472583922](Vue进步.assets/1693472583922.png)

![1693472625463](Vue进步.assets/1693472625463.png)

#### AOP

我开始认为AOP只是通过Filter过滤器还有中间件来实现的，没想到还可以通过特性（Attributes），拦截器（Interceptors），依赖注入（DI）容器（Autofac）进行实现。

其实理解这些实现的话，我们可以从AOP的概念来看，AOP是为了将哪些不是业务主要的逻辑，但是影响多个模块的东西抽离出业务本身，通过切面的方式，将这些公共的实现，比如日志记录，权限验证，授权，异常处理，结果返回，缓存等内容集中到外部进行统一实现，即使后续有改动也很方便管理

**一般AOP有下面这三种实现方式：**

1. **.NET Core内建特性**：
   - **Filters**：[主要用于ASP.NET](http://xn--asp-g88dxot62q103a.net/) Core中，对MVC和Web API的请求进行拦截，实现诸如异常处理、授权、日志记录等。
   - **Middleware**：[适用于ASP.NET](http://xn--asp-xi9dr87onx7a.net/) Core的HTTP请求管道，用于构建请求处理流程，如身份验证、错误处理、日志等。
2. **第三方框架或库**（提供三个最流行的）：
   - **Castle DynamicProxy**：通过创建对象的动态代理来实现拦截器，可以在方法调用之前或之后插入额外的行为，而不影响原有类的代码。
   - **PostSharp**：利用编译时代码注入技术实现AOP，开发者可以通过在代码中使用特性标记来定义切面逻辑。
   - **AspectCore Framework**：[专为.NET](http://xn--ohqpd.net/) Core设计的轻量级AOP框架，允许通过依赖注入和属性标记来应用拦截逻辑。
3. **手动实现（不依赖框架）**：
   - **使用拦截器（Interceptors）**：手动创建拦截器，结合像Castle DynamicProxy这样的动态代理库来实现AOP。
   - **使用特性（Attributes）**：创建自定义特性并应用于方法或类，结合反射和动态代理技术在运行时注入切面逻辑。
   - **装饰器模式**：通过手动封装对象来增加额外的行为，虽然不是AOP，但可以提供类似的功能。
   - **反射和表达式树**：使用.NET的反射API和表达式树来动态地创建对象并且注入横切逻辑。

这种分类方式提供了一个清晰的视图，[区分了基于.NET](http://xn--ykqqa67xg9a13m.net/) Core内建功能的AOP实现、依赖流行第三方库的实现，以及不依赖特定框架的手动实现方式。这样的分类有助于开发者根据实际项目需求和技术栈偏好来选择合适的AOP实现策略。

#### 系统内置注入类

内置的依赖注入（DI）系统自动注入的类，直接可以使用，比如直接使用ILogger类：

```c#
    public class LoggingMiddleware
    {
        private readonly ILogger _logger;

        public LoggingMiddleware(ILogger<LoggingMiddleware> logger)
        {
            _logger = logger;
        }

        public async Task InvokeAsync(HttpContext context)
        {
           _logger.LogInformation("Finished handling request.");
        }
    }
```



##### 日志记录和配置

1. **ILogger<T>**: 用于应用程序的日志记录。允许记录不同级别的日志信息，如调试、信息、警告、错误等。
2. **IConfiguration**: 为应用程序提供配置数据的访问。可以从多个来源读取配置，如 JSON 文件、环境变量、命令行参数等。

##### 环境和上下文信息

1. **IWebHostEnvironment / IHostingEnvironment**: 提供应用程序运行环境的信息，如环境名称（Development、Staging、Production）、应用程序根目录等。
2. **IHttpContextAccessor**: 提供对当前 HTTP 请求上下文的访问。可以用于获取请求、响应、用户、Session 等信息。

##### 依赖注入和服务配置

1. **IServiceProvider**: 作为 DI 容器的基础接口，用于获取服务实例,获取IServiceProvider的四种方法：
   -  var servicesProvider=services.BuildServiceProvider();  ，**不看版本，通用**
   - IServiceProvider 通过构造函数注入  ，**不看版本，通用**
   - IApplicationBuilder app ，            var provider = app.ApplicationServices;  **Net Core 3.1**
   - var app = builder.Build();   var provider = app.Services ;  **NET5,6,7,8** 
2. **IServiceCollection**: 用于在应用程序启动时注册和配置服务。
3. **IOptions<T>**: 提供对强类型配置类的访问，允许从配置文件中读取并映射设置。

##### 缓存和资源管理

1. **IMemoryCache**: 提供内存中的缓存机制，用于存储和检索数据。
2. **IDistributedCache**: 提供分布式缓存服务，适用于多服务器或微服务环境。

##### HTTP 客户端管理

1. **IHttpClientFactory**: 用于创建和配置 `HttpClient` 实例，管理 HTTP 连接的生命周期。

##### 用户和身份管理（ASP.NET Core Identity）

1. **UserManager<TUser>**: 管理用户的创建、删除、搜索、角色等操作。
2. **SignInManager<TUser>**: 处理用户的登录、登出、认证等操作。

##### 数据库和模型绑定

1. **DbContext**: 在使用 Entity Framework Core 时，用于数据库的 CRUD 操作，管理数据库上下文。
2. **IModelBinder**: 负责将请求数据（如表单数据、路由数据）绑定到模型。

##### MVC 和视图帮助程序

1. **IActionContextAccessor**: 提供对当前 MVC 控制器动作的上下文信息的访问。
2. **ITempDataDictionary**: 用于在连续的请求之间存储临时数据。
3. **IUrlHelper**: 用于构建 URL，包括路由 URL 和动作方法的 URL。
4. **IViewComponentHelper**: 用于在视图中呈现和管理视图组件。
5. **IApplicationBuilder**: 用于配置和构建应用程序的请求处理管道。

##### 其他

1. **IHttpClientFactory**: 提供一种方式来创建和配置 `HttpClient` 实例，适用于与外部 HTTP 服务的交互。

这些服务为 ASP.NET Core 应用程序提供了从前端到后端、从基础设施到业务逻辑的全面支持，帮助开发者构建功能丰富、高效可靠的 Web 应用程序。

#### 系统内置中间件

内置中间件是微软提供的，可以直接使用

##### 示例代码

```
public class Startup
{
    public void ConfigureServices(IServiceCollection services)
    {
        // 配置服务
        services.AddControllers();
        services.AddResponseCaching(); // 服务注册示例
    }

    public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
    {
        // 环境判断
        if (env.IsDevelopment())
        {
            app.UseDeveloperExceptionPage();
        }
        else
        {
            app.UseExceptionHandler("/Error");
            app.UseHsts();
        }

        app.UseHttpsRedirection();
        app.UseStaticFiles();

        app.UseRouting();

        app.UseAuthentication();
        app.UseAuthorization();

        app.UseEndpoints(endpoints =>
        {
            endpoints.MapControllers(); // 设置控制器路由
        });

        // 其他中间件的配置
        app.UseCors();
        app.UseResponseCaching(); // 启用响应缓存
        // 继续配置其他中间件...
    }
}
```

[ASP.NET](http://asp.net/) Core 默认中间件列表

##### 安全性相关中间件

1. **Authentication Middleware** (`UseAuthentication`)
   - 启用认证功能，确保用户在访问资源之前被认证。
2. **Authorization Middleware** (`UseAuthorization`)
   - 用于授权，确保用户具有对请求资源的访问权限。
3. **CORS Middleware** (`UseCors`)
   - 配置跨源资源共享（CORS），允许或拒绝来自不同源的请求。
4. **Antiforgery Middleware** (`UseAntiforgery`)
   - 防止跨站请求伪造（CSRF）攻击。

##### 错误处理和诊断中间件

1. **Developer Exception Page Middleware** (`UseDeveloperExceptionPage`)
   - 在开发环境中显示详细的异常信息页面。
2. **Exception Handler Middleware** (`UseExceptionHandler`)
   - 在生产环境中捕获并处理异常，重定向到错误页面。
3. **HTTP Strict Transport Security Middleware** (`UseHsts`)
   - 强制实施HTTP严格传输安全（HSTS），增加安全性。
4. **StatusCode Pages Middleware** (`UseStatusCodePages`)
   - 提供简单的错误处理页面，适用于不同的HTTP状态代码。

##### 性能优化中间件

1. **Response Caching Middleware** (`UseResponseCaching`)
   - 缓存HTTP响应，减少服务器负载和响应时间。
2. **Response Compression Middleware** (`UseResponseCompression`)
   - 压缩HTTP响应内容，提高响应速度。

##### 请求处理中间件

1. **Routing Middleware** (`UseRouting`)
   - 启用路由。
2. **Static Files Middleware** (`UseStaticFiles`)
   - 提供静态文件服务，如图片、CSS和JavaScript文件。
3. **Directory Browser Middleware** (`UseDirectoryBrowser`)
   - 允许浏览服务器的文件目录。

##### 会话和数据处理中间件

1. **Session Middleware** (`UseSession`)
   - 提供会话支持，存储用户数据。
2. **Cookie Policy Middleware** (`UseCookiePolicy`)
   - 确保应用程序遵守Cookie政策。

##### 本地化和全球化中间件

1. Request Localization Middleware

    

   (

   ```
   UseRequestLocalization
   ```

   )

   - 支持本地化和全球化，处理不同文化的日期、数字格式等。

##### 日志记录和监控中间件

1. **HTTP Logging Middleware** (`UseHttpLogging`)
   - 记录HTTP请求和响应数据，帮助调试。
2. **Request Timing Middleware** (`UseRequestTiming`)
   - 用于监控请求处理时间。

##### 其他中间件

1. **URL Rewriting Middleware** (`UseRewriter`)
   - 重写URL请求，用于SEO或遗留系统兼容。
2. **Forwarded Headers Middleware** (`UseForwardedHeaders`)
   - 用于处理由反向代理设置的X-Forwarded-For, X-Forwarded-Proto, X-Forwarded-Host头。



#### `Startup` 两个重要的方法:

其实这个文件名称叫Startup是不重要的（需要遵守Startup中的函数的编写规范），并且需要在Program中进行UseStartup注册这个对象，代表引用程序启动时，需要注册哪些服务与中间件

**中间件与服务的区别是什么？** 

中间件好像是一个程序一个功能，你所有的请求过来我都需要看一遍，看看你符不符合往管道深入请求的权限，比如可能到这一步URL会被转发出去，会出现跨域问题，会对请求进行限流

服务的话就是用别人的功能，我需要初始化一下，激活一下（然后才能正常在应用程序里面使用），配置一些参数阿，比如数据库链接，用户名，密码。还有就是引用一下，就像Using 命名空间一样，告诉应用程序我需要在我这个框架里面用到已有的哪些服务，要往Services里面加入哪些服务，比如services.AddAbp,AddCors,AddRedis 引用这些服务的时候需要加入方法的参数配置（比如跨域的范围），加入config文件中的数据库配置



服务是工具类一样，我在ConfigureServices里面注册了这些服务类，工具类，那么我在应用程序全局都可以被使用，但是中间件里面注册的是管道类型的应用，偏更加底层的东西，所有的请求都会进入到每一个注册的中间件中，然后判断是否有访问权限，是否需要转发服务等，如果没有问题再访问具体的功能方法

##### ConfigureServices

   

- `ConfigureServices`：配置应用所需服务，在该方法中可以添加应用所需要的功能或服务，这个方法注入IServiceCollection services对象，比如可以通过services添加跨域处理功能，添加MVC的服务，添加Redis的服务 ，需要注意的是主机启动时，必须把Startup.cs通过 `.UseStartup<>` 进行注册（一般放在Program里面），比如下图：

  **添加功能/服务：** 跨域，MVC,跨域请求

  ![1672889662430](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1672889662430.png)

  **添加功能/服务**：就是添加Services对象的一个扩展方法

  ![1672889511502](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1672889511502.png)

 **AddHostedService**: 定时任务功能

##### Configure

- `Configure`：配置应用请求处理管道，中间件，管道的意思就是每个Http请求都会先进入到我们配置的管道中，然后再进入到具体的请求API，比如可以配置跨域的管道，Swagger的管道

  

​     管道也是配置的扩展方法，我们需要也可以自己写管道的扩展方法



**Startup中ConficeServices与Configure的思考：**

`3.1.2.1 如何配置 AppStartup`
AppStartup 是一个抽象的空类，没有任何定义成员。正是因为这样，才提供更加灵活的配置方式。

（AppStartup是Furion框架配置的）

`3.1.2.2 AppStartup 约定``

``AppStartup` 派生类只有两个小约定：

- 任何公开、非静态、返回值为 `void` 且方法第一个参数是 `IServiceCollection` 类型，那么他就是一个 `ConfigureServices` 方法
- 任何公开、非静态、返回值为 `void` 且方法第一个参数是 `IApplicationBuilder` 类型，第二个参数是 `IWebHostEnvironment` 类型，那么他就是一个 `Configure` 方法

所以，我们可以自由的编写方法，只要遵循约定即可，如：



```c#
using Microsoft.Extensions.DependencyInjection;

namespace Furion.EntityFramework.Core
{
    public class MyStartup : AppStartup
    {
    //标准的官方案例
    public void ConfigureServices(IServiceCollection services)
    {
       services.AddDataValidation();
    }
    //AppStartup 针对 Configure 方法提供了参数解析注入功能，也就是只要在方法中声明接口参数即可自动注入,如：app 和 env 会自动注入
   public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
    {
        app.UseSwagger();
    }

    // 可以随意定义名字和方法，
    public void XXXXName(IServiceCollection services)
    {
    }

    // 可以随意定义名字和方法，如：
    public void ZZZName(IApplicationBuilder app, IWebHostEnvironment env)
    {
    }
}
}
```

并且可以定义很多Startup，只需要集成AppStartup这个类就可以，然后按照上面的规范进行定义方法就可以，通过定义特性`[AppStartup(order)]`来决定框架调用的顺序，多个Startup其实并不冲突，[有兴趣可以看](https://furion.baiqian.ltd/docs/appstartup)



#### Startup重要的方法

**services.AddDynamicApiControllers**: 添加动态接口控制器服务，就是不需要写控制器的那个功能

```
 services.AddDynamicApiControllers(); //ConfigureServices中
```

**services.AddControllers**：注册服务功能到每个控制器上面

```c#
 //比如注册全局异常处理到控制器服务上（ExceptionWangRui是自定义的异常过滤类）
 services.AddControllers(u=>u.Filters.Add(typeof(ExceptionWangRui)));
//应该有一些服务注册的逻辑需要在AddControllers后面	，比如下面的
services.AddControllers().AddFriendlyException() // 友好异常服务处理逻辑，需在 services.AddControllers() 之后注册。


```

**app.UseStaticFiles()** ：实现对静态文件的访问（css,jquery） ，这个是网站部署后，如果不启用这个那么应用程序相关的脚本，css，图片都无法正常访问展示到界面

**builder.UseMiddleware** :构建自定义中间件

```c#
//调用
 app.UseRequestCulture();
//中间件拓展类，MiddlewareExtensions后缀，通过 UseMiddleware添加到中间件管道中
public static class RequestCultureMiddlewareExtensions
{
    public static IApplicationBuilder UseRequestCulture(this IApplicationBuilder builder)
    {
        return builder.UseMiddleware<RequestCultureMiddleware>();
    }
}
//定义中间件：Middleware 后缀
public class RequestCultureMiddleware
{
    
    private readonly RequestDelegate _next;
    //默认只需要注入一个RequestDelegate 请求委托服务，也可以注入多个对象，比如日志，配置对象（注入方式与普通注入一致）
    public RequestCultureMiddleware(RequestDelegate next)
    {
        _next = next;
    }

    public async Task InvokeAsync(HttpContext context)
    {
       // 中间件内容

        // 调用下一个中间件
        await _next(context);
    }
}
```

**app.Run** :所有的请求返回同一个结果

```c#
app.Run(async context =>
{
    await context.Response.WriteAsync("Hello world!");
});
```

**app.Use()**: 拦截请求，添加一些东西，Use开头的，比如UseCors（添加跨域支持），UseAbpService（Abp服务支持），表达的意思都是所有的请求进入的时候，都会进入我这个管道中间件中

```c#
app.Use(async (context, next) =>
{
    // 比如设置统一头
    context.Response.Headers["framework"] = "Furion";
    // 执行下一个中间件
    await next.Invoke();
});

// 多个
app.Use(...);
```

**app.Map** :添加路由中间件,当**访问指定路由(/hello)**的时候，每个请求都需要先进入到我这里面,app.Run写在最外层是拦截了所有的请求，让他们都返回同样的一个结果

```c#
app.Map("/hello", app => {
    app.Run(async context =>
    {
        await context.Response.WriteAsync("Map Test 1");
    });
});

app.Map("/hello/say", app => {
    // ....
});
```



#### Startup相关的对象

##### ConfigurationBuilder/配置容器

**ConfigurationBuilder**： 用来管理键值对的配置对象

![1672901963609](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1672901963609.png)

**读取ConfigurationBuilder中读取的配置：**

![1672903667554](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1672903667554.png)

##### Configuration/配置

**Configuration**：读取配置

[博客园读取配置](https://oschina.gitee.io/learn-git-branching/)

[微软读取配置](https://learn.microsoft.com/zh-cn/aspnet/core/fundamentals/configuration/?view=aspnetcore-7.0#getsection-getchildren-and-exists)

通过依赖注入注入实例有几种方式：

```c#

//方法1： 构造函数注入方式,net6之前用这种方式，一般在服务中注入
private readonly IConfiguration _configuration;
public DefaultController(IConfiguration configuration)
{
    _configuration = configuration;
}
//方法2 从 builder 对象获取 Configuration 实例，一般在Program.cs中使用
            var configuration = builder.Configuration;
            builder.Services.AddDbContext<EFDbContext>(options =>
            options.UseMySql(configuration.GetConnectionString("DefaultConnection"),
                new MySqlServerVersion(new Version(8, 0, 21))));
            var app = builder.Build();
也可以直接读取
 builder.Configuration.GetSection("ConnectionStrings")["Default"]，或者简写
 builder.Configuration["ConnectionStrings:Default"] 效果是一样的
 
//方法3 静态方法中调用，方便全局使用  Program.cs中调用初始化方法，作为方法参数传递配置参数
    //比如Startup中调用
    var configuration = builder.Configuration; // 获取配置
    services.SqlSugarConfigure(configuration); // 传递配置
    // ... 其他服务注册


    //声明SqlSugar静态方法扩展，添加配置
        public static class ServiceCollectionExtension
    {
        /// <summary>
        /// SqlSugar配置
        /// </summary>
        /// <param name="services"></param>
        private void SqlSugarConfigure(this IServiceCollection services,IConfiguration _appConfiguration)
        {
            services.AddScoped<ISqlSugarClient>(u =>
              ...
                    DbType = DbType.MySql,
                    ConnectionString = _appConfiguration.GetConnectionString("Default"),
				.....
 
        }
    }
//方法4 声明一个静态配置类，在startup构造函数中赋值，然后直接在其他的类中进行调用，类似百伦的ConfigManagerConf静态对象
                                    
public static class ApplicationConfig
{
    public static IConfiguration Configuration { get; set; }
}
 startup中给configuration对象赋值
public class Startup
{
    public Startup(IConfiguration configuration)
    {
        ApplicationConfig.Configuration = configuration;
    }
    // ...
}
 然后可以很方便的在其他的方法中进行调用
public static class ServiceCollectionExtension
{
    public static void SqlSugarConfigure(this IServiceCollection services)
    {
        services.AddScoped<ISqlSugarClient>(u =>
        {
           
                DbType = DbType.MySql,
                ConnectionString = ApplicationConfig.Configuration.GetConnectionString("Default"),
                
        // ...
    }
}
                                                
/*************下面的知识点可以扩展***********************/
// 参数注入方式 [FromServices]
public string Get([FromServices] IConfiguration configuration)
{
}

// 属性注入方式（Net Core默认不支持属性注入的方式，需要依赖IOC的其他框架）
public IConfiguration Configuration { get; set; }
// 读取配置
var data = Configuration.GetSection("配置节点").Get<类型>();


```

##### WebHostBuilder/浏览器

```c#
    static void Main(string[] args)
    {
new WebHostBuilder()  // 声明一个Web主机
            .UseKestrel() //声明一个Kestrel服务器，负责对HTTP请求的监听、接收和最终的响应
            .UseStartup<Startup>() //使用配置的服务与中间件注入服务器
            .Build() // 初始化Web主机对象
            .Run();  //启用这个Web主机
    }
```

##### ServiceDescriptor/服务对象

**是的，ServiceCollection妈妈的肚子里面装的都是我们**,保存了服务的名称，以及服务注册的方法，我也可以直接创建一个服务对象，然后Add到ServiceCollection对象中去

![1684829756435](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1684829756435.png)

**代码案例**

```c#
using System;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.DependencyInjection.Extensions;

public interface IMyService
{
    string GetName();
}

public class MyService : IMyService
{
    public string GetName() => "MyService";
}


public class Program
{
    static void Main(string[] args)
    {
        // 使用ServiceDescriptor和ServiceCollection的Insert方法
        var services = new ServiceCollection();
        var descriptor = ServiceDescriptor.Singleton<IMyService, MyService>();
        services.Insert(0, descriptor);

        var provider = services.BuildServiceProvider();
        using (var scope= provider.CreateScope())
        {
            var scopeProvider = scope.ServiceProvider;
            var myservice = scopeProvider.GetService<IMyService>();
            Console.WriteLine($"Insert:{myservice?.GetName()}");  //输出： Insert:MyService
        }
        // 使用AddSingleton方法，这种方法更简洁
        services.RemoveAll<IMyService>();
        services.AddSingleton<IMyService, MyService>();
        using (var scope = provider.CreateScope())
        {
            var scopeProvider = scope.ServiceProvider;
            var myservice = scopeProvider.GetService<IMyService>();
            Console.WriteLine($"AddSingleton:{myservice?.GetName()}");  //输出： AddSingleton:MyService
        }

    }
}
```



##### ServiceCollection/服务注册器（重要）

**是的，我就是保存服务的容器，服务宝宝都在我的肚子里面**

往里面添加服务的容器，并且可以指定添加服务的生命周期，很多服务的添加都是声明了静态扩展方法，然后通过AddSingleton/AddScoped/AddTransient三个其中之一的方法，将服务添加到服务管道中，比如Services.AddRedis...，就是调用AddSingleton注入服务对象，原理就是扩展方法

###### 服务注册案例

```c#
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.DependencyInjection.Extensions;
using System;

// 定义服务接口
public interface IMyService
{
    void DoSomething();
}

// 定义服务的实现
public class MyService : IMyService
{
    public void DoSomething()
    {
        Console.WriteLine("Doing something...");
    }
}

class Program
{
    static void Main(string[] args)
    {
        // 创建一个新的 ServiceCollection
        var serviceCollection = new ServiceCollection();

        // 创建一个 ServiceDescriptor
        var serviceDescriptor = new ServiceDescriptor(typeof(IMyService), typeof(MyService), ServiceLifetime.Scoped);

        // 添加 ServiceDescriptor 到 ServiceCollection
        serviceCollection.Add(serviceDescriptor);
   //也可以把上面两步和成一步：        serviceCollection.AddScoped(typeof(IMyService),typeof(MyService));，并且这里有个技巧，如果实例没有接口，那么直接注册实例本身也是可以的：serviceCollection.AddScoped(typeof(MyService),typeof(MyService));
        
        // 从 ServiceCollection 构建 ServiceProvider
        var serviceProvider = serviceCollection.BuildServiceProvider();

        // 使用 ServiceProvider 创建 ServiceScope ,看这里返回一个作用域的服务提供者
        using (var scope = serviceProvider.CreateScope())
        {
            // 从 ServiceScope 获取 ServiceProvider，然后从这里拿到服务提供者
            var scopedProvider = scope.ServiceProvider;

            // 使用 ServiceProvider 解析服务，拿到对应的服务
           // var myService = scopedProvider.GetRequiredService<IMyService>();
            //GetRequiredService如果没有获取到对应的服务会产生异常，但是GetService没有获取到对应的服务会返回null值
             var myService = scopedProvider.GetService<IMyService>();

            // 使用服务
            myService?.DoSomething();
        }
    }
}

```

###### 服务注册Scope权限

如果 `services.BuildServiceProvider(true)` ，那么直接用顶级服务提供者获取Scoped对象就会有异常，但是用作用域对象Scope获取这个服务就没有问题，所以顶级服务提供者是用来获取单例与瞬时的服务的

```c#
using System;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.DependencyInjection.Extensions;

public interface IMyService
{
    string GetName();
}

public class MyService : IMyService
{
    public string GetName() => "MyService";
}


public class Program
{
    static void Main(string[] args)
    {
        // 使用Scope镜像ServiceProvider获取服务
        var services = new ServiceCollection();
        services.TryAddScoped<IMyService, MyService>(); //Scoped作用域类型对象
        var provider = services.BuildServiceProvider(true);
        using (var scope= provider.CreateScope())
        {
            var scopeProvider = scope.ServiceProvider;
            var myservice = scopeProvider.GetService<IMyService>();
            Console.WriteLine($"Insert:{myservice?.GetName()}");  //输出： Insert:MyService
        }
        //使用顶级服务提供者,services.BuildServiceProvider(true)无法获取Scope方式注册的对象
        var myservice2 = provider.GetService<IMyService>();
        Console.WriteLine($"Insert:{myservice2?.GetName()}");  //异常：无法从根提供商解析限定范围的服务` IMyService `

    }
}
```

[扩展阅读ServiceCollection与ServiceDescriptor的关系](https://cloud.tencent.com/developer/article/1373968)

###### 服务原生方法

- ServiceCollection的原生方法，就把它当做一个集合看待，提供集合的添加，移除，根据索引读取，

  ![1684829390296](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1684829390296.png)

​      

- ServiceCollectionServiceExtensions.ServiceCollection的扩展方法，提供注入服务的方式：单例注入，瞬时注入，作用域注入

  ![1684808128080](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1684808128080.png)



###### 扩展：（常用）可重复的服务注册 Add

- ServiceCollectionServiceExtensions中注册服务，即使服务存在也会注册进去

  ![1684810053443](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1684810053443.png)

**问题:**.ServiceCollectionServiceExtensions与ServiceCollectionDescriptorExtensions ServiceCollection有两个扩展类，里面的扩展方法都是用来注册服务的，有什么区别吗?

`ServiceCollectionServiceExtensions`和`ServiceCollectionDescriptorExtensions`都提供了一些扩展方法，用于在`ServiceCollection`中注册服务。这些方法提供了一种方便的方式来注册服务，而不需要直接创建`ServiceDescriptor`实例。

下面是一些这两个类中扩展方法的相似之处和不同之处：

相似之处：

- 它们都是为`IServiceCollection`接口提供的扩展方法，可以被`IServiceCollection`的任何实现（例如`ServiceCollection`）使用。
- 它们都提供了一种方便的方式来注册服务，例如单例服务、作用域服务和瞬态服务。

不同之处：

- `ServiceCollectionServiceExtensions`提供了一些最常用的扩展方法，例如`AddSingleton`，`AddScoped`和`AddTransient`。这些方法可以接收服务类型和实现类型，或者直接接收一个服务的实例或工厂方法。
- `ServiceCollectionDescriptorExtensions`提供了一些更复杂的扩展方法，例如`Add`，`TryAdd`，`Replace`等。这些方法可以接收一个`ServiceDescriptor`实例，或者接收服务类型、实现类型和生命周期，然后内部创建`ServiceDescriptor`实例。

在实际编程中，你可以根据需要选择使用哪种扩展方法。如果你只需要注册一些简单的服务，那么`ServiceCollectionServiceExtensions`提供的方法就足够了。如果你需要进行更复杂的操作，例如替换已注册的服务，或者避免重复注册，那么你可能需要使用`ServiceCollectionDescriptorExtensions`提供的方法。

**问题**：ServiceCollectionServiceExtensions中不管服务存不存在都会重复注册，那么如果同一个接口，注册了两个服务实例，他也不会拒绝，那么我获取的时候是获取的哪个服务实例呢？

**答案**：确实，如果你在`ServiceCollection`中对同一个接口注册了两个服务实例，`ServiceCollection`将会保存所有的注册。然后，当你通过`IServiceProvider`的`GetService`方法获得服务时，将会发生什么那么就取决于你取取服务的方式：

1. 如果你使用`GetService`方法，它将返回最后登记册的服务实例。也就是说，后登记册的服务会“覆盖”先登记册的服务，至于使用获得取`GetService`单个服务实例时是这样的。
2. 如果你使用`GetServices`方法（注意是复数形式），它将返回一个包含所有注册表的服务实例的集合。这样你就可以获得所有注册表的服务实例，包括那些被“覆盖”的服务。

这就是为什么在注册服务时需要谦虚，尤其是在注册单例服务或使用域服务时，因为这可能会影响到你的应用程序的行为。 如果你不想覆盖已经注册的服务，你可以使用系列`TryAdd`方法，它们只有在没有登记相关服务时才会添加服务。

这也是为什么`IServiceCollection`提供了多种注册服务的方法，以方便你可以根据你的需要选择合适的方法。

```c#
using System;
using Microsoft.Extensions.DependencyInjection;

public interface IMyService
{
    string GetName();
}

public class MyService1 : IMyService
{
    public string GetName() => "MyService1";
}

public class MyService2 : IMyService
{
    public string GetName() => "MyService2";
}

public class Program
{
    static void Main(string[] args)
    {
        // 创建一个新的 ServiceCollection
        var services = new ServiceCollection();

        // 向 ServiceCollection 中注册两个实现同一接口的服务
        services.AddSingleton<IMyService, MyService1>();
        services.AddSingleton<IMyService, MyService2>();

        // 构建服务提供者
        var serviceProvider = services.BuildServiceProvider();

        // 使用 GetService 获取服务
        var myService = serviceProvider.GetService<IMyService>();
        Console.WriteLine(myService.GetName());  // 输出: "MyService2"

        // 使用 GetServices 获取所有服务
        var myServices = serviceProvider.GetServices<IMyService>();
        foreach (var service in myServices)
        {
            Console.WriteLine(service.GetName());  // 输出: "MyService1" 和 "MyService2"
        }
    }
}

```

###### 扩展：不重复的服务注册 TryAdd

- ServiceCollectionDescriptorExtensions中注册服务，如果服务存在的话，那么就不注册

![1684809262660](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1684809262660.png)

###### 扩展：创建服务提供者 BuildServiceProvider

`ServiceCollectionContainerBuilderExtensions`: ServiceCollection对象创建ServiceProvider对象,validateScopes默认是ture创建作用域的对象，就好像把根容器的ServiceProvider复制了一份给到请求的服务，每一个服务彼此之间使用的服务提供者都是独立的。

![1684817125578](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1684817125578.png)

当涉及到并发、状态共享和安全性时，使用作用域可以帮助我们管理服务的生命周期和共享行为。下面是关于作用域的总结：

**作用域的概念**：在.NET Core的依赖注入容器中，作用域是一种用于管理服务生命周期和共享行为的机制。它定义了一段代码中服务实例的可见范围。

**作用域的创建**：每个作用域都是从根容器派生出来的子容器。通过创建作用域，我们可以创建一个独立的服务解析环境，与其他作用域隔离开来。

**服务解析优先级**：在作用域中解析服务时，解析器会首先查找作用域自身的服务注册。如果找到了匹配的服务，将使用作用域级别的服务实例。如果在作用域中没有找到匹配的服务，解析器会回退到根容器中的服务注册。

**并发问题**：使用作用域可以避免并发问题，不同作用域中的服务实例相互独立，不会相互干扰。每个作用域有自己独立的服务实例，因此在多个线程或并发请求的情况下，每个作用域都能够安全地进行服务解析和操作。

**状态共享问题**：作用域的设计使得每个作用域都有自己的服务实例，可以避免状态共享问题。每个作用域的服务实例是独立的，因此在一个作用域中的服务状态不会被其他作用域或根容器中的服务实例影响或修改。

**安全性问题**：通过使用作用域，我们可以实现对服务的访问控制和限制。在不同的作用域中可以注册不同的服务实现，从而实现安全性要求。只有在适当的作用域范围内才能访问和使用特定的服务实例。

总之，作用域是用于管理服务的生命周期和共享行为的机制。它通过隔离不同作用域的服务实例，解决了并发问题、状态共享问题和安全性问题。通过创建独立的作用域，我们可以确保每个作用域具有自己的服务实例，从而实现灵活的服务管理和控制。

**故事内容：**

在一个小学图书馆中，每个学生都可以借阅图书。这个图书馆使用.NET Core来构建一个图书馆管理系统，并使用依赖注入来管理图书馆的服务。

**作用域的概念**：在图书馆中，每个学生都有一个图书馆卡（作用域），它定义了学生借阅图书的可见范围。每个学生的借阅行为都在自己的作用域内进行。

**作用域的创建**：当学生进入图书馆时，他们会得到一张图书馆卡（作用域）。这张卡是从图书馆的根容器派生出来的，它为学生创建了一个独立的借阅环境。

**服务解析优先级**：当学生想要借阅图书时，他们会使用自己的图书馆卡去解析服务。解析器首先会查找学生作用域内的服务注册信息，如果找到匹配的服务，就会使用学生作用域级别的服务实例。如果在学生作用域内没有找到匹配的服务，解析器会回退到图书馆的根容器中的服务注册。

**并发问题**：在图书馆中，可能会有多个学生同时借阅图书。每个学生都有自己的作用域和服务实例，因此不同学生之间的借阅行为是相互独立的，不会相互干扰。即使多个学生同时访问图书馆，它们都能够在自己的作用域内安全地进行图书借阅。

**状态共享问题**：在图书馆中，每个学生的借阅状态是私有的，不会被其他学生或根容器中的服务实例影响或修改。每个学生都有自己独立的服务实例，借阅状态是与自己的作用域相关联的。这样，即使多个学生同时借阅相同的图书，它们的借阅状态也是相互独立的。

**安全性问题**：在图书馆中，可能会有一些图书是限制借阅的，只有年级较高的学生才能借阅。通过在适当的作用域范围内注册特定的服务实现，我们可以实现对图书的访问控制和限制。只有符合条件的学生在其作用域范围内才能够借阅相应的图书。

 

###### 扩展：配置操作 Configure

`OptionsServiceCollectionExtensions`

是一个扩展方法类，它提供了一些用于配置和注册选项（Options）的便捷方法。在 ASP.NET Core 中，选项模式（Options Pattern）用于将配置信息以类型安全的方式注入到应用程序中。

![1684829017433](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1684829017433.png)

###### SqlSqlSugar注入仓储经典

SqlSqlSugar的代码注入非常经典，下面我们来解释一下这段代码

```c#
   /// <summary>
        /// SqlSugar配置
        /// </summary>
        /// <param name="services"></param>
        private void SqlSugarConfigure(IServiceCollection services)
        {
            services.AddScoped<ISqlSugarClient>(u =>
            {
                var sqlSugarClient = new SqlSugarClient(new ConnectionConfig()
                {
                    DbType = DbType.MySql,
                    ConnectionString = _appConfiguration.GetConnectionString("Default"),
                    IsAutoCloseConnection = true,
                });
                return sqlSugarClient;
            });

            // 注册非泛型仓储
            services.AddScoped<ISqlSugarRepository, SqlSugarRepository>();

            // 注册 SqlSugar 仓储
            services.AddScoped(typeof(ISqlSugarRepository<>), typeof(SqlSugarRepository<>));
        }
```

这段代码在ASP.NET Core的服务容器中注册了三个服务。下面解释一下为什么需要这样做。

1. `ISqlSugarClient` 是SqlSugar库中的一个接口，它代表了一个到数据库的连接。在这段代码中，它被配置为每次HTTP请求时创建一个新的`SqlSugarClient`实例，并且这个实例在请求结束时会自动关闭连接。
2. `ISqlSugarRepository` 和 `SqlSugarRepository` 代表了一个非泛型的仓储。仓储是数据访问层的一个常见模式，用于抽象数据库操作。非泛型仓储是对所有类型的实体执行相同操作的场合，例如获取所有的实体或者删除特定的实体。
3. `ISqlSugarRepository<>` 和 `SqlSugarRepository<>` 则是泛型仓储，它们是针对特定类型的实体的。例如，你可能有一个`ISqlSugarRepository<User>` 来执行与用户实体相关的数据库操作，然后有另一个`ISqlSugarRepository<Product>` 来执行与产品实体相关的数据库操作。这样的设计可以帮助我们把类型相关的代码（例如，访问User表的代码和访问Product表的代码）分开，使得代码更易于理解和维护。

关于这三种服务的区别：

- `ISqlSugarClient` 提供了最基础的数据库操作能力，它是底层的数据库连接客户端，当需要更细粒度的控制，比如需要自定义查询的超时时间，或者你需要控制连接的生命周期可以操作这个对象
- `ISqlSugarRepository` 和 `SqlSugarRepository` 提供了对所有类型实体的数据库操作能力，它们是一种高级别的抽象，用于隐藏数据库操作的具体细节。
- `ISqlSugarRepository<>` 和 `SqlSugarRepository<>` 则提供了针对特定类型实体的数据库操作能力，它们在`ISqlSugarRepository`的基础上添加了类型安全性，使得在编写和阅读代码时更不易出错。

注入这三种服务的原因主要是为了分离关注点和提高代码的可维护性。具体来说，它们可以帮助我们把底层的数据库操作代码、泛型的数据访问代码和类型特定的数据访问代码分开，每种代码只关注自己的事务，不需要关心其他的细节。

以下是一些实际的使用案例和代码片段：

1. `ISqlSugarClient` 的使用案例： 例如你可以在开始的时候判断某个表是否存在，如果不存在就新建或者删除，或者查询的时候需要设置底层的超时时间，这些都是很有用的

```c#
public class MyService
{
    private readonly ISqlSugarClient _client;

    public MyService(ISqlSugarClient client)
    {
        _client = client;
    }

    public async Task CreateTableIfNotExistsAsync()
    {
        if (!_client.DbMaintenance.IsAnyTable("User"))
        {
            await _client.CodeFirst.InitTablesAsync(typeof(User));
        }
    }

    
    public async Task<List<User>> GetUsersWithTimeout(string query)
    {
        // 设置查询的超时时间为50秒
        _client.Ado.CommandTimeOut = 50;

        return await _client.SqlQueryable<User>(query).ToListAsync();
    }
}
```

1. `ISqlSugarRepository` 的使用案例： 如果你希望执行一些通用的数据库操作，例如获取所有的实体或者删除一个实体，你可以使用`ISqlSugarRepository`。

```c#
public class MyService
{
    private readonly ISqlSugarRepository _repository;

    public MyService(ISqlSugarRepository repository)
    {
        _repository = repository;
    }

    public async Task<List<User>> GetAllUsers()
    {
        return await _repository.GetListAsync<User>();
    }

    public async Task<List<Product>> GetAllProducts()
    {
        return await _repository.GetListAsync<Product>();
    }

    public async Task<bool> DeleteUser(int id)
    {
        return await _repository.DeleteAsync<User>(id);
    }

    public async Task<bool> DeleteProduct(int id)
    {
        return await _repository.DeleteAsync<Product>(id);
    }
}

```

1. `ISqlSugarRepository<>` 的使用案例： 如果你希望执行一些特定类型的实体相关的操作，你可以使用`ISqlSugarRepository<>`。

```
public class MyService
{
    private readonly ISqlSugarRepository<User> _userRepository;
    private readonly ISqlSugarRepository<Product> _productRepository;

    public MyService(ISqlSugarRepository<User> userRepository, ISqlSugarRepository<Product> productRepository)
    {
        _userRepository = userRepository;
        _productRepository = productRepository;
    }

    public async Task<List<User>> GetAllUsers()
    {
        return await _userRepository.GetListAsync();
    }

    public async Task<List<Product>> GetAllProducts()
    {
        return await _productRepository.GetListAsync();
    }

    public async Task<bool> DeleteUser(int id)
    {
        return await _userRepository.DeleteAsync(id);
    }

    public async Task<bool> DeleteProduct(int id)
    {
        return await _productRepository.DeleteAsync(id);
    }
}
```

这些例子中，`ISqlSugarClient` 提供最基础的数据库操作，`ISqlSugarRepository` 和 `ISqlSugarRepository<>` 在其基础上提供了一些更方便的方法，以便你能更简洁、更直观地执行一些常见的数据库操作。

**魔法餐厅的故事**
想象你是一个宇航员，被送到了一个奇妙的星球上，这个星球叫做“魔法星球”。在这个星球上，一切都是用代码构建的。这个星球的一部分看起来就像一个巨大的餐厅，我们把它叫做"魔法餐厅"。

在这个餐厅中，ISqlSugarClient 就像餐厅的厨房。这个厨房不是普通的厨房，它充满了奇妙的机器，会自动制作各种食物（数据）。不过，这个厨房的操作方式很复杂，你需要知道如何调整机器，如何控制温度，甚至还需要知道如何编程才能使用它。因此，只有厨师（高级程序员）才能直接操作厨房。

ISqlSugarRepository 就像餐厅的服务员。他们穿着闪亮的制服，总是笑容满面。你只需要告诉他们你想吃什么，他们就会去厨房为你准备。他们可以为你准备任何你想要的食物，无论是炒饭还是炒面。这就像是在你的程序中，你只需要告诉ISqlSugarRepository你需要什么数据，它就会帮你从数据库中取出来。

ISqlSugarRepository<> 是一种特别的服务员，他们是专门负责某一种食物的。比如，你有一个专门负责披萨的服务员，有一个专门负责汉堡的服务员。他们知道所有关于这种食物的秘密，知道如何制作出最好吃的披萨或汉堡。在你的程序中，ISqlSugarRepository<> 就是专门处理某一种数据的，他们知道所有关于这种数据的细节，知道如何最有效地从数据库中取出和保存这种数据。

这就是在"魔法星球"的"魔法餐厅"中，ISqlSugarClient、ISqlSugarRepository 和 ISqlSugarRepository<> 的故事。在这个星球上，它们就像是餐厅的厨房和服务员，帮助你制作和服务数据，让你能够更容易地编写和理解你的代码。

###### 骚操作，在控制台访问Web Api

为什么想在控制台运行Web Api呢，是因为这样可以更加体会WebApplication,ServicesCollection,ServiceProvider这些对象的作用，它们几个对象一起组成了Web Api，有人负责管理整个Http管道，有人负责管理Ioc注入。

**引用两个包**：

- Install-Package Swashbuckle.AspNetCore
- Install-Package Microsoft.Extensions.DependencyInjection

```c#

using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;

var builder = WebApplication.CreateBuilder(args);

// 设置配置
builder.Configuration.SetBasePath(Directory.GetCurrentDirectory())
                    .AddJsonFile("appsettings.json", optional: true);

// 注入服务
builder.Services.AddScoped<IMyService, MyService>();

// 创建WebApplication
var app = builder.Build();

// 配置路由
app.MapGet("/", ([FromServices] IMyService myService) => myService.SayHello());

app.Run();
public interface IMyService
{
    string SayHello();
}

public class MyService : IMyService
{
    public string SayHello()
    {
        return "Hello from MyService!";
    }
}
// 测试浏览器访问：http://localhost:5000/,https://localhost:5001/
```



##### ServiceProvider/服务提供者（根容器）

`ServiceProvider` 是.NET Core 中的一个类，用于提供依赖注入容器的服务解析功能。它是一个用于注册、存储和解析服务的容器。

具体来说，`ServiceProvider` 的主要功能包括：

1. **注册服务**：使用 `ServiceCollection` 将服务描述符（`ServiceDescriptor`）添加到容器中，然后通过 `BuildServiceProvider` 方法构建 `ServiceProvider` 实例。注册的服务可以包括接口和实现类之间的映射、生命周期信息和其他配置。
2. **解析服务**：通过调用 `ServiceProvider` 实例的 `GetService` 或 `GetRequiredService` 方法，根据服务类型获取服务实例。解析服务时，容器会根据注册的服务描述符和解析规则，实例化并返回相应的服务对象。
3. **管理生命周期**：`ServiceProvider` 负责跟踪和管理注册服务的生命周期。根据服务描述符中的生命周期设置（如瞬时、作用域、单例），容器在需要时创建新的服务实例、重用现有的服务实例或在不再需要时销毁服务实例。
4. **支持作用域**：`ServiceProvider` 可以创建和管理作用域（`ServiceScope`）。作用域提供了一个临时的服务解析环境，它可以在需要时创建，并在使用完毕后被释放。通过作用域，可以实现更细粒度的服务隔离和生命周期控制。

总之，`ServiceProvider` 是一个依赖注入容器，负责注册、存储和解析服务。它是依赖注入框架的核心组件，用于管理应用程序中的服务依赖关系，并提供灵活的服务解析和生命周期管理功能。



当涉及到服务的注册、解析和生命周期管理时，`ServiceProvider` 提供了相应的功能。下面是相关代码案例和解释：

**故事：**

从前，在一个小镇上有一个神奇的书屋。这个书屋里有各种各样的书籍，供小朋友们阅读和学习。书屋的老板为了方便小朋友们借书，决定引入一个新的系统来管理书籍的借阅和归还。

为了实现这个系统，老板决定使用一个特殊的工具——ServiceProvider。ServiceProvider就像是一个帮助管理书屋的工具箱，里面有各种各样的功能。

首先，老板使用ServiceProvider来注册书屋的服务。他创建了一个叫做ServiceCollection的列表，里面包含了书屋的各种服务，比如借书服务、还书服务等。然后，他使用ServiceProvider的方法将这些服务注册到ServiceProvider中。

接着，当小朋友们来借书时，ServiceProvider就派上了用场。小朋友们向ServiceProvider提出借书的请求，ServiceProvider会根据注册的服务描述符和规则，在书屋中找到对应的服务，然后将这个服务提供给小朋友们。这样，小朋友们就可以借到自己想要的书籍了。

而且，ServiceProvider还能管理书籍的生命周期。每一本书都有一个生命周期，比如有的书是永久保存的，有的书是限时借阅的。ServiceProvider会根据这些设置来管理书籍的状态，确保每一本书都能在适当的时候借出和归还。

此外，ServiceProvider还支持创建和管理作用域。作用域就像是一个小的书屋分部，里面有一些特定的书籍和服务。当小朋友们进入作用域时，ServiceProvider会提供一个专门的环境给他们，让他们可以借阅作用域内的书籍。而当小朋友们离开作用域时，ServiceProvider会释放这个环境，让其他小朋友继续使用。

所以，ServiceProvider就像是书屋的魔法工具，它帮助书屋管理各种服务，让借书变得更加方便和有序。有了ServiceProvider，小朋友们可以愉快地阅读和学习，书屋也变得更加高效和有组织。

总之，ServiceProvider是一个神奇的工具，它帮助书屋注册、存储和提供各种服务，管理书籍的生命周期，并支持创建和管理作用域。有了它，书屋变得更加方便和有序，小朋友们可以更好地享受阅读的乐趣。

**注册服务：**

```c#
// 创建一个 ServiceCollection 实例
var services = new ServiceCollection();

// 注册服务描述符
services.AddSingleton<IMyService, MyService>();
```

在这个示例中，我们创建了一个 `ServiceCollection` 实例，并使用 `AddSingleton` 方法注册了一个服务描述符。这里使用 `AddSingleton` 表示我们要注册一个单例服务，每次解析该服务时都会返回同一个实例。

**解析服务：**

```c#
// 构建 ServiceProvider
var serviceProvider = services.BuildServiceProvider();

// 解析服务实例
var myService = serviceProvider.GetService<IMyService>();
```

在这个示例中，我们使用 `BuildServiceProvider` 方法构建了一个 `ServiceProvider` 实例。然后通过调用 `GetService` 方法，根据服务类型 `IMyService` 解析了服务实例。`GetService` 方法会返回已注册的服务实例，如果没有找到匹配的服务则返回 `null`。

**生命周期管理：**

```c#
// 注册一个作用域服务描述符
services.AddScoped<IMyScopedService, MyScopedService>();

// 在作用域中解析服务
using (var scope = serviceProvider.CreateScope())
{
    var scopedProvider = scope.ServiceProvider;
    var scopedService = scopedProvider.GetService<IMyScopedService>();
    // 使用作用域内的服务实例
}
```

在这个示例中，我们使用 `AddScoped` 方法注册了一个作用域服务描述符。然后，通过创建作用域 `ServiceScope`，我们可以从作用域的服务提供者中解析作用域服务实例。在作用域内解析的服务实例具有相同的生命周期，当作用域结束时，服务实例会被销毁。

**支持作用域：**

```c#
// 创建一个作用域
using (var scope = serviceProvider.CreateScope())
{
    var scopedProvider = scope.ServiceProvider;
    // 在作用域中解析和使用服务
}
```

在这个示例中，我们使用 `CreateScope` 方法创建了一个作用域 `ServiceScope`。通过作用域的服务提供者，我们可以在作用域内解析和使用服务。作用域提供了一个临时的服务解析环境，可以在需要时创建，并在使用完毕后释放。

总结：`ServiceProvider` 是用于注册、解析和管理服务的容器。通过注册服务描述符，我们可以将服务添加到容器中。然后，通过调用 `GetService` 方法可以解析服务实例。使用作用域，我们可以实现更细粒度的服务隔离和生命周期控制。

###### 获取服务 GetService

**HttpContext**： 这里也可以直接获取IServiceProvider

```c#
        var cache = context.HttpContext.RequestServices.GetService<IMemoryCache>();
```

**IApplicationBuilder**: 这里也可以直接获取IServiceProvider

```c#
        var serviceProviderIsService = app.ApplicationServices.GetService<IServiceProviderIsService>();
//也可以用app.Services.GetService<IServiceProviderIsService>();

```



##### ServiceScope/获得服务作用域（子容器）

引用于ServiceProvider，通过CreateScope/CreateAsyncScope创建服务作用域，每个作用域ServiceScope都有一个自己的ServiceProvider（引用于根ServiceProvider），进行创建用完马上就释放的瞬时对象，如果直接用根ServiceProvider对象GetService那么对象会一直存在应用程序当中，只有应用程序停止后内存才会被GC回收

[多看几遍，闲聊.NET Core 依赖注入](https://zhuanlan.zhihu.com/p/113001805)

[蒋金楠大佬讲依赖注入，这是一个系列(整个asp net core 系统的系列)，都可以读一读](https://www.cnblogs.com/artech/p/inside-asp-net-core-03-08.html)

`2023年2月3日10:32:29 花了一个半小时看了一下系列前4的文章，讲的非常深入，用代码实例实现了asp net core的IOC容器，17年写的到现在都在流行，只要做asp net core就必须学习的东西`

**韬韬韬你羞得无礼大佬笔记：**

![1675328045331](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1675328050224.png)

[【韬韬韬你羞得无礼分享依赖注入】](https://www.kdocs.cn/l/crMl10mZAmuB)

看的过程中一些问题：

1.根容器创建的服务是不是一直不会销毁的，只有应用程序关机了才会销毁

![img](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/a255a8ad892f2378b0e230656e70537a.webp)

**比如以下代码：**

```c#
[HttpGet]
public Void JudgeScope([FromServices]IServiceProvider scopeProvider)
{
    // 使用using Scope用完直接释放（这种方式获取的对象是新的对象，即使是同一个请求，在下面using作用域声明的作用域对象都是不同于using外面声明的对象，正常应该是一样的）
    await using (AsyncServiceScope scope = scopeProvider.CreateAsyncScope())
    {
        var operation= scope.ServiceProvider.GetService<IOperationScoped>();
    }
    
    //根对象scopeProvider获取:对象会一直存在应用程序当中，只有应用程序停止后内存才会被GC回收
   var operation= scopeProvider.GetService<IOperationScoped>()
}

```

##### WebApplication /管家

在.NET 6中，Microsoft引入了新的WebApplication类，为应用启动和配置提供了更简洁的API。这个类集成了常见的Web开发功能，并且更好地整合了Configuration和依赖注入（DI）容器，使得开发变得更加便利。WebApplication类简化了ASP.NET Core应用程序的启动和配置流程，使开发者能更专注于应用程序的业务逻辑。

WebApplication类的主要功能有：

1. **集成配置**：WebApplication类封装了Configuration，让开发者更方便地访问和设置应用程序的配置。你可以使用`WebApplication.Configuration`属性来访问应用程序的配置，并且可以通过`WebApplication.CreateBuilder(args)`创建并初始化一个`WebApplicationBuilder`对象，使用其`Configuration`属性来配置应用程序。

2. **注入服务**：WebApplication类封装了依赖注入容器，使得开发者能更方便地注册和解析服务。你可以使用`WebApplication.Services`属性来访问应用程序的服务提供器，并通过`WebApplicationBuilder.Services`属性来注册你的服务。

3. **中间件和路由配置**：WebApplication类提供了直观的API来配置中间件和路由。例如，你可以使用`WebApplication.MapGet`，`WebApplication.MapPost`等方法来设置路由，无需创建新的控制器。

4. **应用程序生命周期管理**：WebApplication类提供了`WebApplication.Run`方法来启动应用程序，并提供了`WebApplication.Dispose`方法来清理资源。

   ![1684898574703](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1684898574703.png)

下面是一个使用WebApplication的示例代码：

```c#
using Microsoft.AspNetCore.Mvc;

var builder = WebApplication.CreateBuilder(args);

// 设置配置
builder.Configuration.SetBasePath(Directory.GetCurrentDirectory())
                    .AddJsonFile("appsettings.json", optional: false);

// 注入服务
builder.Services.AddScoped<IMyService, MyService>();

// 创建WebApplication
var app = builder.Build();

// 配置路由
app.MapGet("/", ([FromServices] IMyService myService) => myService.SayHello());

app.Run();
public interface IMyService
{
    string SayHello();
}

public class MyService : IMyService
{
    public string SayHello()
    {
        return "Hello from MyService!";
    }
}
// 测试浏览器访问：https://localhost:7101/
```

在这个示例中，我们首先创建并配置了一个`WebApplicationBuilder`，然后注册了一个服务。接下来，使用`WebApplicationBuilder.Build`方法创建了一个`WebApplication`对象，并配置了路由。最后，使用`WebApplication.Run`方法启动了应用程序。

总的来说，WebApplication类简化了.NET应用程序的配置和管理流程，使开发者能够更高效、便捷地进行应用开发和维护。

**乐高的故事**：

假设你正在玩一个叫做 "WebApplication乐高" 的游戏，这个游戏里，你可以建造各种你想要的乐高城堡。首先，你需要一个魔法盒子，我们叫它 "WebApplication"。这个盒子里面装了所有你需要的乐高积木。

首先，你要创建你的乐高城堡。但是每个城堡都有它自己的设计蓝图，我们可以把这个蓝图看作是"配置"。这个蓝图会告诉你如何建造城堡，哪个积木应该放在哪里，什么时候使用窗户积木，什么时候使用门积木等等。在我们的魔法盒子中，有一个神奇的工具叫做"Configuration"，它就像是我们的蓝图，告诉我们如何建造我们的乐高城堡。例如，它可以告诉我们，门需要在城堡的正面，窗户需要在城堡的两侧，或者你需要在什么时间打开城堡的大门等等

接下来，我们要开始建造城堡了。在建造过程中，我们可能需要各种各样的工具和帮手，例如小矮人工人帮我们搬运积木，魔法妖精帮我们粘合积木，甚至有时候我们可能需要巨人的力量帮我们堆叠高高的城堡墙壁。这些工具和帮手就像是我们的"服务"。在我们的魔法盒子中，有一个神奇的工具箱叫做"Services"，里面装满了各种我们需要的工具和帮手，比如小矮人工人，魔法妖精，巨人的力量，这些都可以帮助我们建造城堡。

在建造城堡的过程中，我们需要一条路来指引我们如何建造，这条路就像是我们的"路由"。例如，我们需要知道从哪个方向开始建造城堡，以及在哪里设置城堡的大门等等。在我们的魔法盒子中，有一个神奇的地图叫做"MapGet"或者"MapPost"，它就像是我们的路，"MapGet"可以告诉我们如何从一堆乱糟糟的积木中找到我们需要的积木，"MapPost"可以告诉我们如何把我们的城堡展示给游客看。

最后，当我们建造完成城堡，我们需要一个方法来打开城堡的大门，让游客们可以参观我们的城堡。这个方法就是"Run"。当我们想要关闭城堡，清理城堡里面的垃圾时，我们可以使用另一个方法，那就是"Dispose"。

所以，这个魔法盒子就是我们的WebApplication类。它提供了我们需要的所有工具和服务，帮助我们建造和管理我们的乐高城堡。希望你现在可以更好地理解这个概念了。

**问题**：我看很多类都有一个Builder的对象，比如WebApplication与WebApplicationBuilder，IConfiguration与IConfigurationBuilder，ILogger与ILoggingBuilder  ，我不知道这是为什么，你可以告诉我为什么要这样设计吗

**答案**：

**说明内容：**
这是一种常见的设计模式称为建造者模式（Builder Pattern）。Builder Pattern被广泛用于创建复杂对象的构建和配置过程。

在.NET中，通常会有一个主要的类，如WebApplication，IConfiguration，ILogger等，这些类用于执行主要的操作，如运行Web应用程序，访问配置信息，记录日志等。然后，还有与之对应的"Builder"类，如WebApplicationBuilder，IConfigurationBuilder，ILoggingBuilder等，这些Builder类用于在实例化主类之前配置和创建它们。

使用Builder Pattern的原因主要有以下几点：

- **分离构建和表示**：Builder类负责构建过程，主类负责业务逻辑。这使得主类的实现更加简洁，更关注于业务逻辑，而构建过程的复杂性则被封装在Builder类中。
- **更好的可读性和易用性**：如果要通过构造函数传递许多参数来构建对象，代码会变得难以阅读和维护。使用Builder Pattern，可以使代码更具可读性和易用性。
- **更好的灵活性**：使用Builder Pattern，可以很容易地为对象添加新的配置选项，而不需要修改对象的构造函数或其他方法。这使得对象的配置更加灵活。
- **控制对象的构建过程**：有时，对象的构建过程可能需要多个步骤，并且步骤之间可能有特定的顺序。在这种情况下，使用Builder Pattern可以更好地控制对象的构建过程。

总的来说，Builder Pattern是一种设计模式，它提供了一种构建复杂对象的有效方法。在.NET中，许多类都使用了这种模式，以简化构建过程并提供更好的灵活性和易用性。

**故事：**
想象你是一个小朋友，你有一盒乐高积木。这些乐高积木可以被用来搭建各种东西，比如房子，汽车，飞机等。这就像我们的主类，比如WebApplication，IConfiguration，ILogger等，它们都是用来做某件事情的，比如运行Web应用程序，访问配置信息，记录日志等。

但是，你不能直接开始搭建你的乐高积木，你需要一个搭建计划或者说明书，告诉你怎么去搭建这个东西，比如哪些积木应该在哪个位置，应该按什么顺序搭建等。这就像我们的"Builder"类，比如WebApplicationBuilder，IConfigurationBuilder，ILoggingBuilder等，这些Builder类是用来帮助你构建和配置你的乐高积木的。

那么为什么需要一个Builder呢？首先，Builder可以帮助你把搭建过程和最后的结果分开。你不需要亲自去记住每一个步骤，你只需要按照Builder的方法搭建好就可以了，然后不用管具体的搭建细节，直接拿到最后积木的结果，然后用这个结果去分享给朋友们，或者参加比赛都可以，取决你用这个结果干什么。

其次，使用Builder会让你的搭建过程更容易理解和使用。如果没有Builder，你可能需要记住所有的步骤和细节，这会让你感到很困难。但是如果有了Builder，你只需要按照它的指导一步步去做就可以了。

另外，使用Builder还可以提供更好的扩展性。比如，你可以按照你的想法来改变一些步骤，而不需要修改整个搭建过程。这就像你可以很容易地为你的乐高积木添加新的部件，而不需要从头开始搭建。

最后，使用Builder可以帮助你更好地控制搭建过程。有时，搭建乐高积木可能需要多个步骤，并且步骤之间可能有特定的顺序。使用Builder，你可以按照正确的顺序和步骤来搭建你的乐高积木。

所以，就像你使用说明书来搭建乐高积木一样，程序员也会使用Builder模式来帮助他们构建和配置他们的程序。

**代码案例：**

以下是一个例子说明如何在没有使用建造者模式和使用建造者模式的情况下创建一个复杂的“汽车”对象。

**案例一：没有使用建造者模式**

```c#
public class Car
{
    public string Engine { get; set; }
    public string Transmission { get; set; }
    public string Body { get; set; }
    public List<string> Accessories { get; set; }
}

// 使用代码
var car = new Car
{
    Engine = "V8",
    Transmission = "Automatic",
    Body = "Sedan",
    Accessories = new List<string> { "Leather Seats", "Sunroof", "Navigation System" }
};
```

在这个例子中，我们直接创建了一个汽车对象，并在创建时为其设置了属性。当属性数量增加，构造函数参数增多时，这种方式会导致代码难以阅读和理解。

**案例二：使用建造者模式**

```c#
public class Car
{
    public string Engine { get; set; }
    public string Transmission { get; set; }
    public string Body { get; set; }
    public List<string> Accessories { get; set; }
}

public class CarBuilder
{
    private readonly Car _car = new Car();

    public CarBuilder WithEngine(string engine)
    {
        _car.Engine = engine;
        return this;
    }

    public CarBuilder WithTransmission(string transmission)
    {
        _car.Transmission = transmission;
        return this;
    }

    public CarBuilder WithBody(string body)
    {
        _car.Body = body;
        return this;
    }

    public CarBuilder WithAccessories(List<string> accessories)
    {
        _car.Accessories = accessories;
        return this;
    }

    public Car Build()
    {
        return _car;
    }
}

// 使用代码
var car = new CarBuilder()
    .WithEngine("V8")
    .WithTransmission("Automatic")
    .WithBody("Sedan")
    .WithAccessories(new List<string> { "Leather Seats", "Sunroof", "Navigation System" })
    .Build();
```

在这个例子中，我们使用了建造者模式来创建汽车对象。这种方式使得代码更具可读性和易用性，同时也使得我们可以更灵活地配置汽车对象。当对象的构建过程需要多个步骤或者步骤之间有特定的顺序时，建造者模式能更好地控制对象的构建过程。

##### WebHostBuilder /旧管家

在 .NET Core 3.1 中，`WebHostBuilder` 是用于构建和配置 Web 主机的工具类，Net6中使用`WebApplication`。它提供了一种简单而灵活的方式来配置 Web 应用程序的启动和运行环境。

`WebHostBuilder` 类的主要功能如下：

1. **配置服务器**：通过使用 `UseKestrel()`、`UseIIS()` 等方法，你可以选择要使用的服务器类型，并配置服务器选项。例如，你可以指定监听的端口、启用 HTTPS、配置 SSL 证书等。
2. **配置应用程序**：使用 `Configure()` 方法，你可以配置应用程序的行为和中间件。通过提供一个委托，你可以指定应用程序的路由、中间件管道和其他自定义行为。在这里，你可以注册路由、配置身份验证、添加全局异常处理等。
3. **配置日志**：使用 `ConfigureLogging()` 方法，你可以配置应用程序的日志记录。你可以指定日志记录器的提供程序、日志级别和其他相关设置。例如，你可以添加控制台日志记录器、文件日志记录器等。
4. **配置服务**：使用 `ConfigureServices()` 方法，你可以配置应用程序的依赖注入（DI）容器。你可以注册服务、配置依赖关系和设置其他 DI 相关的选项。在这里，你可以添加数据库上下文、注入自定义服务等。



![1684911916177](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1684911916177.png)

以下是一个完整的代码示例，展示了如何使用 `WebHostBuilder` 构建和配置一个基本的 Web 应用程序：

```c#
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;

public class Program
{
    public static void Main(string[] args)
    {
        CreateHostBuilder(args).Build().Run();
    }

    public static IHostBuilder CreateHostBuilder(string[] args) =>
        Host.CreateDefaultBuilder(args)
            .ConfigureWebHostDefaults(webBuilder =>
            {
                webBuilder.ConfigureServices((hostContext, services) =>
                {
                    // 配置服务
                    // ...
                })
                .Configure(app =>
                {
                    var env = app.ApplicationServices.GetService<IWebHostEnvironment>();

                    if (env.IsDevelopment())
                    {
                        // 在开发环境执行特定的逻辑
                        // ...
                    }
                    else if (env.IsProduction())
                    {
                        // 在生产环境执行特定的逻辑
                        // ...
                    }

                    // 配置中间件和路由
                    app.UseRouting();
                    app.UseEndpoints(endpoints =>
                    {
                        endpoints.MapGet("/", async context =>
                        {
                            await context.Response.WriteAsync("Hello World!");
                        });
                    });
                })
                .ConfigureLogging((hostContext, logging) =>
                {
                    // 配置日志记录器
                    // ...
                });
            });
}

```



###### 扩展 WebHostBuilderExtensions

![1684911767647](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1684911767647.png)

##### IWebHostEnvironment/地图

`IWebHostEnvironment`接口在ASP.NET Core中是非常重要的一部分，主要用于提供Web主机的信息。特别地，它提供了以下几种信息：

1. 应用程序的根路径：`IWebHostEnvironment.ContentRootPath`属性指定了应用程序的根目录。这通常是项目的根目录，所有相对路径都是基于这个目录的。
2. Web根路径：`IWebHostEnvironment.WebRootPath`属性指定了web的根目录，这个目录通常用于存储静态文件，如css，js，images等。
3. 环境名称：`IWebHostEnvironment.EnvironmentName`属性提供了当前运行环境的名称，如"Development"，"Staging"，"Production"等。你可以根据这个属性来改变应用程序的行为。例如，在开发环境中，你可能会启用一些调试特性，而在生产环境中，你可能会关闭这些特性。
4. 内容文件提供器：`IWebHostEnvironment.ContentRootFileProvider`和`IWebHostEnvironment.WebRootFileProvider`属性提供了文件提供器，这些提供器可以用于读取文件。

这些信息是在应用程序启动时由主机设置的，然后通过依赖注入提供给应用程序的其他部分。

![1684898658456](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1684898658456.png)

下面是一个简单的示例，说明了如何在Nett 6中使用`IWebHostEnvironment`：

```c#
using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;

var builder = WebApplication.CreateBuilder(args);

// 设置配置
builder.Configuration.SetBasePath(Directory.GetCurrentDirectory())
    .AddJsonFile("appsettings.json", optional: false);

// 创建主机环境
var hostEnvironment = builder.Environment;

// 获取当前环境的名称
var environmentName = hostEnvironment.EnvironmentName;
Console.WriteLine($"当前环境：{environmentName}");

// 获取应用程序的根路径
var appRootPath = hostEnvironment.ContentRootPath;
Console.WriteLine($"应用程序根路径：{appRootPath}");

// 获取 Web 根路径
var webRootPath = hostEnvironment.WebRootPath;
Console.WriteLine($"Web 根路径：{webRootPath}");

// 注册服务
builder.Services.AddSingleton<IHostEnvironment>(hostEnvironment);

// 创建 WebApplication
var app = builder.Build();

// 配置路由和中间件等

app.Run();

```

在这个示例中，`IWebHostEnvironment`在构造函数中注入到`Startup`类中，然后在`Configure`方法中使用。特别地，我们使用`Environment.IsDevelopment()`方法来检查当前是否为开发环境，然后根据环境不同，使用不同的错误处理中间件。

**小故事：** 

想象你是一个寻宝者，你在一个神秘的城堡里寻找宝藏。城堡里有很多房间，但你并不知道每个房间的位置。在这种情况下，你需要一张城堡的地图。`IWebHostEnvironment`就像这张地图，它告诉你每个房间（在这里，房间是指应用程序的不同部分，如根目录，web根目录等）的位置。

另外，城堡可能有多个区域，如公共区域，私人区域，宴会区域等。每个区域可能有不同的规则。例如，在私人区域，你可能需要保持安静，而在宴会区域，你可以尽情欢乐。同样，你的应用程序可能运行在不同的环境中，如开发环境，测试环境，生产环境等。每个环境可能有不同的设置和行为。例如，在开发环境中，你可能想要更多的调试信息，而在生产环境中，你可能想要更多的安全和性能。`IWebHostEnvironment`的`EnvironmentName`属性就像是区域的标识，告诉你现在在哪个区域，你应该遵守什么规则。

在寻宝的过程中，你可能需要读取一些文件，如日记，地图等。城堡可能有一些特定的地方存储这些文件。`IWebHostEnvironment`的`ContentRootFileProvider`和`WebRootFileProvider`属性就像是指向这些地方的指南针，帮助你找到和读取这些文件。

所以，`IWebHostEnvironment`就像是你的地图，指南针，和规则手册，帮助你在复杂的应用程序中找到你需要的信息，遵守正确的规则，和读取正确的文件。它是你在开发ASP.NET Core应用程序时的重要伙伴。

##### IHostingEnvironment /旧地图

`IHostingEnvironment` 接口是 ASP.NET Core 3.1中用于获取和管理应用程序环境的接口(Net6中请使用`IWebHostEnvironment`)。它提供了一种方便的方式来获取应用程序所处的环境信息，例如开发环境、测试环境或生产环境，并且可以根据不同的环境执行不同的逻辑。

通过 `IHostingEnvironment` 接口，你可以获取当前应用程序的环境名称、应用程序根目录、Web 根目录等信息，以及判断当前是否处于开发环境或生产环境。

![1684909706045](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1684909706045.png)

以下是一个示例代码，演示了如何使用 `IHostingEnvironment` 接口：

```c#
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using System;

public class Startup
{
    public void ConfigureServices(IServiceCollection services)
    {
        // 配置服务
        // ...
    }

    public void Configure(IApplicationBuilder app, Microsoft.AspNetCore.Hosting.IHostingEnvironment env)
    {
        if (env.IsDevelopment())
        {
            // 在开发环境执行特定的逻辑
            // ...
        }
        else if (env.IsProduction())
        {
            // 在生产环境执行特定的逻辑
            // ...
        }

 
        // 配置路由
        app.Map("/info", (app) =>
        {
            app.Run(async (context) =>
            {
                // 获取应用程序环境信息
                await context.Response.WriteAsync($"EnvironmentName: {env.EnvironmentName}\n");
                await context.Response.WriteAsync($"ContentRootPath: {env.ContentRootPath}\n");
                await context.Response.WriteAsync($"ApplicationName: {env.ApplicationName}\n");
            });
        });
        // 测试，访问：https://localhost:7101/info
    }
}

public class Program
{
    public static void Main(string[] args)
    {
        Host.CreateDefaultBuilder(args)
            .ConfigureWebHostDefaults(webBuilder =>
            {
                webBuilder.UseStartup<Startup>();
            })
            .Build()
            .Run();
    }
}

```

`IHostingEnvironment` 和 `IWebHostEnvironment` 是用于管理应用程序环境的接口，它们在功能和使用上有一些相同点和不同点。

相同点：

1. **环境信息获取**：两个接口都提供了获取应用程序环境信息的方法和属性，例如获取当前环境名称、应用程序根目录、Web 根目录等。
2. **环境判断**：两个接口都提供了判断当前环境的方法，例如判断是否处于开发环境、生产环境等。
3. **环境配置**：两个接口都可以根据不同的环境执行特定的配置，例如根据环境设置数据库连接字符串、日志级别、调试选项等。

不同点：

1. **命名空间和程序集**：`IHostingEnvironment` 属于 `Microsoft.AspNetCore.Hosting` 命名空间和程序集，而 `IWebHostEnvironment` 属于 `Microsoft.Extensions.Hosting` 命名空间和程序集。在新版本的 ASP.NET Core 中，推荐使用 `IWebHostEnvironment`，因为它是对 `IHostingEnvironment` 的更现代化的替代。
2. **扩展性**：`IWebHostEnvironment` 是对 `IHostingEnvironment` 的扩展，提供了一些额外的功能和属性，例如获取 Web 主机的地址、端口号等。因此，如果使用的是 ASP.NET Core 的 Web 主机（例如 Kestrel），建议使用 `IWebHostEnvironment`。
3. **泛化接口**：`IWebHostEnvironment` 是 `IHostingEnvironment` 的泛化接口，可以用于扩展和替代 `IHostingEnvironment`。这意味着在新的 ASP.NET Core 版本中，可以直接使用 `IWebHostEnvironment` 来访问和管理应用程序环境。

使用场景：

- 使用 `IWebHostEnvironment` 的场景：
  - 当使用 ASP.NET Core 的 Web 主机（例如 Kestrel）作为应用程序的承载时，建议使用 `IWebHostEnvironment`。
  - 针对 Web 主机的配置和管理，例如设置静态文件目录、获取 Web 主机的地址和端口号等。
- 使用 `IHostingEnvironment` 的场景：
  - 当使用较旧版本的 ASP.NET Core 或者不涉及 Web 主机的应用程序时，可以使用 `IHostingEnvironment`。
  - 针对一般应用程序环境的配置和管理，例如根据不同环境设置数据库连接字符串、配置文件路径等。

总结来说，`IWebHostEnvironment` 是对 `IHostingEnvironment` 的扩展和泛化，适用于使用 ASP.NET Core 的 Web 主机的场景。而如果你在使用新版本的 ASP.NET Core，建议使用 `IWebHostEnvironment` 来获取和管理应用程序环境信息。

##### ILogger / 日志

`ILogger` 是 .NET Core 中的一个重要接口，它用于创建日志。日志是开发和维护应用程序过程中不可或缺的一部分，可以帮助你了解应用程序的运行情况，比如跟踪应用程序的活动、调试应用程序、诊断问题，甚至在发生错误时记录详细的错误信息。

![1684900844206](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1684900844206.png)

下面是一个使用 `ILogger` 的简单代码示例：

```c#
public class MyClass
{
    private readonly ILogger _logger;

    public MyClass(ILogger<MyClass> logger)
    {
        _logger = logger;
    }

    public void DoSomething()
    {
        _logger.LogInformation("Doing something...");
        
        try
        {
            // 尝试执行一些操作
        }
        catch (Exception ex)
        {
            _logger.LogError(ex, "An error occurred while doing something.");
        }
    }
}
```

在这个例子中，我们首先在类的构造函数中注入了 `ILogger<MyClass>`。然后在 `DoSomething` 方法中，我们使用 `ILogger` 的 `LogInformation` 方法记录了一条信息级别的日志，表示我们正在执行某项操作。如果在尝试执行操作时出现了异常，我们则使用 `ILogger` 的 `LogError` 方法记录了一条错误级别的日志，包含了异常信息和一条错误消息。

现在，让我们通过一个小故事来理解 `ILogger`：

假设你是一个太空船的飞行员，正在驾驶太空船前往火星。在这个长途旅行中，你需要不断地了解太空船的运行状态。你需要知道燃料的消耗情况、生命维持系统的运行情况、当前的位置等等。这时，太空船上的一个神奇装置 —— `ILogger` 就派上用场了。

当你启动发动机时，`ILogger` 就会记录一条信息，告诉你发动机已经启动；当你飞行一段时间后，`ILogger` 会记录下你已经飞行了多少时间，消耗了多少燃料；当你的太空船遇到了问题，比如生命维持系统出现了故障，`ILogger` 会记录下这个错误，并提供详细的故障信息，帮助你解决问题。

通过这个故事，我们可以看到 `ILogger` 就像是太空船上的黑匣子，记录下了太空船的所有运行情况，使得你能在飞行过程中及时了解太空船的状态，并在出现问题时迅速定位问题，采取正确的措施。

###### 扩展 LoggerExtensions 

![1684900884704](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1684900884704.png)

一个小的例子

```c#
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Logging;
using System;

var builder = WebApplication.CreateBuilder(args);

// 设置配置
builder.Configuration.SetBasePath(Directory.GetCurrentDirectory())
                    .AddJsonFile("appsettings.json", optional: false);

// 添加日志服务
builder.Logging.AddConsole(); // 添加控制台输出作为日志提供程序

// 创建WebApplication
var app = builder.Build();

// 获取ILogger对象实例
var logger = app.Services.GetRequiredService<ILogger<object>>();

// 配置路由
app.MapGet("/", () =>
{
    logger.LogInformation("Processing request...");  //一般的操作日志，这个不会在控制台输出，可以在生产环境用

    logger.LogWarning("Invalid input detected. Please check the input data."); //警告日志，这个不会在控制台输出，可以在生产环境用

    logger.LogError("An error occurred while processing the request."); //错误日志,只有这个会在控制台输出
    return "Hello World!";
});

app.Run();

```

##### IConfiguration / 配置

当在.NET中提到"Configuration属性"时，一般指的是应用程序的配置对象。该属性代表了应用程序的配置，包括从不同配置源加载的键值对。

在.NET中，配置通常是使用`IConfiguration`接口来表示。这个接口提供了访问和查询配置值的方法。`Configuration`属性我们可以自定义为一个静态属性，提供了对`IConfiguration`实例的访问，方便我们在应用程序的任何地方获取配置值。

通过`Configuration`属性，我们可以使用索引器语法来获取配置值，例如`Configuration["SettingKey"]`。这里的`SettingKey`是配置文件或其他配置源中的键名。

配置属性的使用使得在应用程序的各个组件中访问配置值变得非常方便。无需在每个需要访问配置的地方都显式创建和传递`IConfiguration`实例，而是可以直接通过`Configuration`属性来获取配置值。

这种方式可以简化代码并提高可维护性，因为我们只需要关注配置的读取和使用，而无需关心具体的配置加载和创建过程。

总之，实际使用中我们可以通过配置`Configuration`属性成一个静态属性，代表应用程序的配置对象（`IConfiguration`），通过它我们可以方便地获取应用程序的配置值。

**一、通过注入使用**

```c#
using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using System;

public class Program
{
    public static void Main(string[] args)
    {
        var builder = WebApplication.CreateBuilder(args);
        builder.Configuration.SetBasePath(Environment.CurrentDirectory)
            .AddJsonFile("appsettings.json");

        var app = builder.Build();
        var configuration = app.Services.GetRequiredService<IConfiguration>();

        var settingValue = configuration["Logging:LogLevel:Default"];
        Console.WriteLine($"SettingValue: {settingValue}"); //输出： Information

        app.Run();
    }
}

```

在这个示例中，我们使用 `WebApplication.CreateBuilder()` 创建一个 `WebApplicationBuilder` 对象，并通过 `SetBasePath()` 和 `AddJsonFile()` 方法来配置应用程序的配置。

然后，我们通过 `builder.Build()` 方法构建一个 `WebApplication` 对象，并使用 `app.Services.GetRequiredService<IConfiguration>()` 获取到注入的 `IConfiguration` 实例。

最后，我们可以使用 `configuration["SettingKey"]` 访问配置值，并打印出来。

通过依赖注入，我们可以将 `IConfiguration` 实例注入到需要访问配置的类中，使得配置的访问变得更加方便和可控。这种方式适用于 .NET 6 的 Web 应用程序。

**二、通过静态属性使用**

```c#
using Microsoft.Extensions.Configuration;
using System;

public class Program
{
    public static void Main(string[] args)
    {
        var settingValue = Configuration["Logging:LogLevel:Default"];
        Console.WriteLine($"SettingValue: {settingValue}"); //输出：Information
    }

    public static IConfigurationRoot Configuration { get; } = new ConfigurationBuilder()
        .SetBasePath(Environment.CurrentDirectory)
        .AddJsonFile("appsettings.json")
        .Build();
}

```

**问题**： Configuration，IConfigurationRoot,ConfigurationBuilder
这三个对象是用来干什么的，彼此之间有什么关系，请讲细一些，并且配合代码案例还有小故事

**回答**：.NET Core中，`Configuration`、`IConfigurationRoot`和`ConfigurationBuilder`是用于读取和管理应用程序配置的对象和类。它们之间有以下关系：

1. `ConfigurationBuilder`：`ConfigurationBuilder`是用于构建`IConfigurationRoot`实例的建造者类。它提供了一系列方法来配置和加载配置源，例如JSON文件、环境变量、命令行参数等。通过链式调用这些方法，我们可以根据需要构建一个包含所有配置的`IConfigurationRoot`对象。
2. `IConfigurationRoot`：`IConfigurationRoot`是应用程序配置的根对象，它表示整个配置的集合。它继承自`IConfiguration`接口，提供了访问和查询配置值的方法。通过`IConfigurationRoot`，我们可以以统一的方式读取配置，无论配置值来自于哪个配置源。
3. `Configuration`：`Configuration`是一个静态类，它提供了对`IConfigurationRoot`实例的访问。通过`Configuration`类，我们可以在应用程序的任何地方方便地访问配置值，而不需要显式地传递`IConfigurationRoot`对象的引用。

下面是一个示例代码，展示了如何使用`ConfigurationBuilder`来构建和访问配置：

```c#
using Microsoft.Extensions.Configuration;
using System;

public class Program
{
    public static void Main(string[] args)
    {
        // 返回IConfigurationRoot对象，提供对配置的访问
        var configuration = new ConfigurationBuilder()
            .SetBasePath(Environment.CurrentDirectory)
            .AddJsonFile("appsettings.json")
            .Build();

        var settingValue = configuration["Logging:LogLevel:Default"];

        Console.WriteLine($"SettingValue: {settingValue}"); // Information
    }
}

```

##### IApplicationBuilder /管道

`IApplicationBuilder`属性是在.NET Core中用于配置应用程序的中间件管道的对象。它允许我们按照特定的顺序添加、配置和组织中间件组件，以处理传入的HTTP请求并生成响应。

通过调用`IApplicationBuilder`上的各种扩展方法，我们可以实现以下功能：

1. **添加中间件组件**：使用`UseMiddleware`方法将中间件组件添加到管道中，按照添加的顺序依次处理请求。
2. **配置中间件组件**：通过在中间件组件的构造函数或`InvokeAsync`方法中编写逻辑来配置中间件的行为，例如身份验证、日志记录、异常处理等。
3. **定义中间件的顺序**：通过添加中间件组件的顺序来定义它们在管道中的执行顺序，确保每个中间件按照预期的顺序处理请求。
4. **调用下一个中间件组件**：在中间件组件的逻辑中，通过调用`_next`委托来调用下一个中间件组件，实现中间件的链式调用。

通过使用`IApplicationBuilder`属性，我们可以在`Startup`类的`Configure`方法中构建和配置中间件管道，以满足应用程序的需求。这种模块化的设计使得应用程序的处理逻辑更可控、可扩展，并且能够轻松地添加、移除或修改中间件组件，以适应不同的场景和功能。



以下是一个简单的示例代码，展示了如何使用`IApplicationBuilder`来配置中间件管道：

```c#
using Microsoft.AspNetCore.Mvc;

var builder = WebApplication.CreateBuilder(args);

// 设置配置
builder.Configuration.SetBasePath(Directory.GetCurrentDirectory())
                    .AddJsonFile("appsettings.json", optional: false);


// 创建WebApplication
var app = builder.Build();
// 添加中间件组件,每一个请求都会进入到我的中间件管道当中
app.UseMiddleware<LoggingMiddleware>();
app.UseMiddleware<AuthenticationMiddleware>();
app.UseMiddleware<RoutingMiddleware>();
app.Run();

public class LoggingMiddleware
{
    private readonly RequestDelegate _next;

    public LoggingMiddleware(RequestDelegate next)
    {
        _next = next;
    }

    public async Task InvokeAsync(HttpContext context)
    {
        // 执行日志记录逻辑
        // ...
        Console.WriteLine("执行日志记录逻辑");
        // 调用下一个中间件组件
        await _next(context);
    }
}

public class AuthenticationMiddleware
{
    private readonly RequestDelegate _next;

    public AuthenticationMiddleware(RequestDelegate next)
    {
        _next = next;
    }

    public async Task InvokeAsync(HttpContext context)
    {
        // 执行身份验证逻辑
        // ...
        Console.WriteLine("执行身份验证逻辑");

        // 调用下一个中间件组件
        await _next(context);
    }
}

public class RoutingMiddleware
{
    private readonly RequestDelegate _next;

    public RoutingMiddleware(RequestDelegate next)
    {
        _next = next;
    }

    public async Task InvokeAsync(HttpContext context)
    {
        // 执行路由匹配逻辑
        // ...
        Console.WriteLine("执行路由匹配逻辑");

        // 调用下一个中间件组件
        await _next(context);
    }
}
// 测试浏览器访问：https://localhost:7101/

```

在上面的示例中，`Startup`类的`Configure`方法接收一个`IApplicationBuilder`实例作为参数。通过调用`IApplicationBuilder`的`UseMiddleware`方法，我们可以按照指定顺序添加中间件组件。每个中间件组件都是一个独立的类，实现了`RequestDelegate`委托。

在每个中间件组件的`InvokeAsync`方法中，我们可以编写具体的逻辑来处理请求和生成响应。然后，通过调用`_next`委托来调用下一个中间件组件，实现中间件的链式调用。

通过使用`IApplicationBuilder`属性，我们可以轻松地定义和配置中间件管道，以实现请求的处理和响应生成。这样的设计使得应用程序的处理逻辑更清晰、模块化，并且能够轻松地添加、移除或修改中间件组件，以满足不同的需求。

**冰淇淋工厂的故事**;

在一个神奇的乐园里，有一个特别的工厂，它可以制造各种美味的冰淇淋。这个工厂的名字叫做".NET Core冰淇淋工厂"。

工厂里有一个重要的机器，它叫做"中间件管道机"。这台机器可以接收来自顾客的请求，然后根据请求的特点，添加、配置和组织不同的冰淇淋口味，最终生成一份美味的冰淇淋响应。

工厂的老板告诉小学生们，这台中间件管道机就是我们的IApplicationBuilder属性。它像是一条生产线，顺序地把各种工序串在一起，最终完成一杯杯美味的冰淇淋。

首先，顾客会将自己的口味请求传递给中间件管道机。管道机会使用"UseMiddleware"方法，按照添加的顺序依次添加不同的中间件组件，比如巧克力酱、草莓酱、坚果等。这些中间件组件就像是不同的冰淇淋配料，每一个都负责加入独特的味道和特点。

然后，中间件管道机会根据顾客的请求配置冰淇淋的行为。比如，如果顾客要求添加糖果粒，管道机就会在中间件组件的构造函数或InvokeAsync方法中设置糖果粒的逻辑，确保冰淇淋有甜甜的口感。

接下来，中间件管道机会根据添加中间件组件的顺序来定义它们在生产线上的执行顺序。这样，每个中间件组件都会按照预期的顺序处理请求，确保每一杯冰淇淋都有正确的配料。

最后，中间件管道机会调用下一个中间件组件，实现冰淇淋生产线的链式调用。每个中间件组件都会处理自己的任务，然后将冰淇淋传递给下一个中间件组件，最终完成一杯完美的冰淇淋。

小学生们听了老板的讲解后，兴奋地跑到中间件管道机旁边，看着冰淇淋生产线的忙碌场景。他们想象着中间件组件们如何协作，一起制作出各种口味的冰淇淋，让人们享受甜蜜的味道。

通过使用IApplicationBuilder属性，.NET Core冰淇淋工厂能够根据不同的顾客需求，灵活地构建和配置中间件管道，以生成定制化的冰淇淋。这种模块化的设计让工厂能够应对不同场景和口味的需求，让每个顾客都能够得到满意的冰淇淋体验。

最后，小学生们品尝了工厂制作的各种口味的冰淇淋，他们感受到了IApplicationBuilder属性的魔力，它带给工厂无限创意和灵活性，让每一杯冰淇淋都独一无二，让顾客们回味无穷。

于是，小学生们也开始动手制作自己的冰淇淋工厂，使用IApplicationBuilder属性来构建和配置自己的中间件管道。他们展示着创意和想象力，制作出了各种奇妙的冰淇淋口味，为大家带来了快乐和惊喜。

这就是IApplicationBuilder属性的故事，它像是一个冰淇淋工厂中的中间件管道机，通过添加、配置和组织中间件组件，实现了定制化的冰淇淋生产过程。它让我们的应用程序变得灵活、可扩展，并且能够满足不同用户的需求。无论是制作冰淇淋还是开发应用程序，中间件管道机都是一个重要的工具，让我们能够创造出美味和精彩的体验。

###### 中间件的作用

ASP.NET Core的中间件是用于处理HTTP请求和响应的组件，它们可以用于执行各种任务和功能。下面是一些中间件可以用来完成的任务：

- 身份验证和授权：中间件可以用于处理用户身份验证和授权，例如使用ASP.NET Core内置的Authentication中间件，支持各种身份验证方案，如Cookie身份验证、JWT（JSON Web Token）身份验证等。
- 路由和URL重写：通过使用中间件，可以实现自定义的URL路由和重写，以匹配和转发请求到正确的处理程序或控制器。
- 请求处理和响应：中间件可以用于执行各种请求处理操作，例如请求日志记录、请求压缩、请求重定向、响应缓存等。
- 异常处理：中间件可以捕获应用程序中的异常并提供自定义的异常处理逻辑，例如记录异常、返回自定义错误页面等。
- 静态文件服务：ASP.NET Core中的Static Files中间件可用于提供静态文件，例如CSS、JavaScript、图像等。
- 日志记录：中间件可以用于记录请求和响应的日志，以便进行应用程序的监控和故障排除。
- 性能优化：中间件可以实现各种性能优化技术，如HTTP请求缓存、Gzip压缩、响应缓存等，以提高应用程序的性能和响应速度。

这只是中间件的一些常见用途，实际上，您可以编写自定义的中间件来满足特定的需求。通过组合和配置不同的中间件，可以构建出灵活、可扩展和高度定制的ASP.NET Core应用程序。

**用中间件处理授权鉴权，异常处理，日志写入 和用Filter实现有什么区别？**

**总结：**第一点 中间件是作用于整个HTTP管道模型的，作用于整个应用程序全局，Filter的粒度更为小，可以作用于特定的方法与控制器上面，第二点 中间件执行的更为靠前，Filter在中间件后面进行执行，第三点 中间件的功能更加强大，更加灵活，可以更改Http请求的请求头，响应头

中间件和过滤器都可以用于处理授权鉴权、异常处理和日志写入，但它们在实现方式和使用场景上存在一些区别。

1. 授权鉴权：
   - 中间件：可以使用中间件来处理身份验证和授权。例如，使用ASP.NET Core内置的Authentication中间件来验证用户身份并授权访问受限资源。
   - 过滤器：可以使用授权过滤器（Authorization Filters）来实现授权鉴权逻辑。授权过滤器可以应用于整个控制器或特定的控制器动作方法，用于验证用户权限和访问控制。

2. 异常处理：
   - 中间件：通过编写自定义的异常处理中间件，可以捕获应用程序中的异常并执行相应的处理逻辑，如记录异常、返回自定义错误页面等。
   - 过滤器：可以使用异常过滤器（Exception Filters）来处理控制器动作方法中发生的异常。异常过滤器可以提供全局或局部的异常处理逻辑。

3. 日志写入：
   - 中间件：可以编写自定义的日志中间件，用于记录请求和响应的日志信息。中间件可以在请求管道中的适当位置记录日志。
   - 过滤器：过滤器本身不直接用于日志记录，但可以与日志库或日志中间件配合使用，通过在过滤器中调用日志记录功能来实现日志写入。

区别：
- 中间件是针对整个请求管道的处理，它们以特定的顺序执行，并可以对请求和响应进行全局的操作。而过滤器是针对控制器动作方法的执行过程进行的操作，可以在控制器的不同阶段插入逻辑。
- 中间件是按照添加到管道中的顺序逐个执行的，而过滤器可以通过设置优先级来控制执行顺序。
- 中间件可以在请求管道的较低级别阶段进行操作，例如处理HTTP请求和响应的最初阶段。而过滤器是在控制器动作执行过程中的特定阶段调用的。
- 中间件可以具有更广泛的应用范围，可以处理各种请求和响应，而过滤器主要用于控制器动作方法的处理。

综上所述，中间件和过滤器都可以用于处理授权鉴权、异常处理和日志写入。选择使用哪种方式取决于具体需求和场景，以及对请求处理的粒度和控制级别的要求。

###### 如何自定义中间件

**问题**：你是一个C#软件开发专家，请问如何使用NET自定义一个中间件，让我们一步一步进行思考

**回答**：

当然，我很愿意帮助你理解如何在.NET中创建自定义的中间件。这里，我将通过C#和ASP.NET Core来演示这个过程，因为在.NET Core和.NET 5+版本中，中间件的概念已经变得非常重要。

中间件在ASP.NET Core应用程序中扮演着非常重要的角色，它们定义了应用程序如何响应HTTP请求，以及请求/响应如何在处理管道中流转。你可以创建自定义的中间件来处理一些特定的功能。

以下是创建自定义中间件的基本步骤：

1. 首先，你需要创建一个新的类来表示你的中间件。这个类需要一个构造函数，该构造函数接受一个指向下一个中间件的引用，并且还需要一个Invoke或InvokeAsync方法来处理请求/响应。

例如，你可以创建如下的自定义中间件：

```c#
public class CustomMiddleware
{
    private readonly RequestDelegate _next;

    public CustomMiddleware(RequestDelegate next)
    {
        _next = next;
    }

    public async Task InvokeAsync(HttpContext context)
    {
        // 在这里执行你的中间件逻辑
        // ...

        // 调用下一个中间件
        await _next(context);
    }
}
```

​     2.接下来，你需要在你的Startup类中注册这个中间件。这可以在Configure方法中使用UseMiddleware方法来完成：

```c#
public class Startup
{
    public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
    {
        app.UseMiddleware<CustomMiddleware>();

        // ...其它中间件和配置
    }
}
```

​     3.最后，你就可以在你的中间件中执行你想要的逻辑了。例如，你可能想要记录每个请求的一些信息，或者你可能想要修改请求或响应的某些部分。

请记住，中间件的顺序非常重要。在Startup类的Configure方法中，你应该按照你想要的顺序注册你的中间件。请求会按照这个顺序通过中间件，响应则是相反的顺序。因此，你需要确保你的中间件在正确的位置。

以上就是在.NET中创建自定义中间件的基本步骤。希望这个例子可以帮助你理解这个过程。如果你有更具体的问题，或者你想要实现某些特定的功能，请随时向我提问。



###### 使用扩展方法自定义中间

这段代码实现了一个自定义中间件，并在 `Startup` 类中将该中间件添加到应用程序的中间件管道中。以下是对代码的一步一步的思考总结：

1. 首先，我们定义了一个静态扩展方法 `UseCustomMiddleware`，用于向 `IApplicationBuilder` 添加自定义中间件。这个中间件会在处理请求之前和之后执行一些逻辑，包括输出一条自定义的消息。
2. 在 `Startup` 类中，我们使用 `app.UseCustomMiddleware` 方法将自定义中间件添加到应用程序的中间件管道中。然后，我们使用 `app.Run` 方法添加一个最终的中间件，用于返回 "Hello from the final middleware" 的响应。
3. 在 `Program` 类的 `Main` 方法中，我们使用 `WebApplication.CreateBuilder` 创建一个 `WebApplication` 的构建器。然后，我们调用 `builder.Build` 方法构建应用程序的主机。
4. 接下来，我们调用 `app.Configure` 方法，并传入一个委托方法 `ConfigureApp`。在 `ConfigureApp` 方法中，我们创建一个 `Startup` 实例，并调用其 `Configure` 方法来配置应用程序的中间件管道。
5. 最后，我们调用 `app.RunAsync` 方法来运行应用程序。

通过以上步骤，我们实现了一个基本的自定义中间件，并在应用程序中使用了该中间件。这个例子展示了如何使用 `IApplicationBuilder` 和 `Startup` 类来配置和管理中间件管道，以及如何构建和运行应用程序。这种方式使得应用程序的配置和启动过程更加灵活和可扩展。

```c#
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.AspNetCore.Http;
using Microsoft.Extensions.Hosting;
using System;
using System.Threading.Tasks;

public static class CustomMiddlewareExtensions
{
    public static IApplicationBuilder UseCustomMiddleware(this IApplicationBuilder app, string message)
    {
        return app.Use(async (context, next) =>
        {
            // 在处理请求之前执行的逻辑
            Console.WriteLine($"Custom Middleware: {message}");

            // 调用下一个中间件
            await next.Invoke();

            // 在处理完请求之后执行的逻辑
            Console.WriteLine("Custom Middleware: Request processed");
        });
    }
}

public class Startup
{
    public void Configure(IApplicationBuilder app)
    {
        app.UseCustomMiddleware("Hello from Custom Middleware");

        app.Run(async (context) =>
        {
            await context.Response.WriteAsync("Hello from the final middleware");
        });
    }
}

public class Program
{
    public static async Task Main(string[] args)
    {
        var builder = WebApplication.CreateBuilder(args);

        var app = builder.Build();

        ConfigureApp(app);

        await app.RunAsync();
    }

    private static void ConfigureApp(WebApplication app)
    {
        var startup = new Startup();
        startup.Configure(app);
    }
}


```

###### MVC的路由与MapGet 和MapPost的差异

大型应用程序中，你可能会发现使用传统的 MVC 风格更合适，其中每个控制器类负责处理一组相关的 HTTP 请求，这样可以更好地组织代码，MapGet与MapPost更适合简单的应用场景路由下，它们的相同点与不同点如下：

相同之处：

1. 他们都是定义请求的 URL 如何映射到处理该请求的代码。
2. 他们都可以接受 HTTP 请求并将其路由到相应的控制器操作。

**不同之处：**

1. 传统的 MVC 路由通常在控制器类和动作方法上使用特性（如 `Route`、`HttpGet`、`HttpPost` 等）来定义。而在这段代码中，路由配置是在 `WebApplication` 实例上直接使用 `MapGet` 方法来定义的。
2. 在 MVC 中，路由配置通常会涉及到一个或多个控制器类，每个类可能有多个动作方法，每个方法都可以处理特定的 HTTP 方法（GET、POST 等）。而在这段代码中，每个 `MapGet` 调用直接定义了一个处理特定 HTTP GET 请求的端点。
3. 在 MVC 中，动作方法通常返回一个 `ActionResult`，它可以表示各种类型的 HTTP 响应，如 HTML 视图、JSON 数据、重定向、错误状态等。而在这段代码中，`MapGet` 调用的 lambda 表达式直接返回一个字符串，这个字符串将作为 HTTP 响应的内容发送。
4. 在 MVC 中，动作方法的参数通常由模型绑定系统从 HTTP 请求中提取。而在这段代码中，lambda 表达式的参数直接由依赖注入系统提供。

所以，这种新的路由配置方式是一种更为简洁、更直接的方法，特别适合构建小型的 API。然而，对于更复杂的应用程序，你可能仍然需要使用 MVC 风格的路由和控制器。

`MapGet` 和 `MapPost` 方法主要应用于简化应用程序或者创建小型 HTTP API 的情况。下面是一些具体的使用案例：

1. **简单的 HTTP 服务：** 如果你正在创建一个非常简单的 HTTP 服务，例如一个只返回某些固定数据或者简单处理数据的微服务，那么使用 `MapGet` 和 `MapPost` 方法可能是个好选择。因为你可以直接在启动类（比如 `Program.cs`）中定义所有的路由和处理函数，而无需创建任何控制器类。

   ```c#
   var app = builder.Build();
   
   app.MapGet("/", () => "Hello World");
   app.MapPost("/api/items", (Item item) => $"Received item with ID: {item.Id}");
   
   app.Run();
   ```

2. **功能测试或者原型开发：** 如果你正在进行功能测试或者原型开发，需要快速创建一些 HTTP 端点用于测试或者展示，那么使用 `MapGet` 和 `MapPost` 方法会非常方便。因为你可以直接在启动类中定义路由和处理函数，而无需创建额外的控制器类。

   ```c#
   var app = builder.Build();
   
   app.MapGet("/api/test", () => "This is a test endpoint");
   app.MapPost("/api/test", (TestObject testObject) => $"Received test object: {testObject.Name}");
   
   app.Run();
   ```

3. **HTTP API 的微调：** 在大型的 HTTP API 中，虽然主要的路由和处理函数通常会在控制器中定义，但是你可能仍然想要使用 `MapGet` 和 `MapPost` 方法来微调或者扩展 API。例如，你可能想要在所有请求之前或者之后添加一些自定义的处理逻辑，或者添加一些只在特定路由上有效的中间件。

   ```c#
   var app = builder.Build();
   
   app.UseMiddleware<MyCustomMiddleware>();
   
   app.MapControllers(); // 使用 MVC 风格的路由和控制器
   
   app.MapGet("/api/custom", () => "This is a custom endpoint"); // 添加一个自定义的路由
   
   app.Run();
   ```

上述只是一些基本的示例，实际使用情况可能会根据应用程序的需求和复杂性进行调整。



#### launchSettings

这里面配置的是启动配置，IIS 启动的端口，路径，默认页或者Kestrel服务器的启动配置

![1673405086654](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1673405086654.png)

## RabbitMQ

刚才有个服务我考虑了一下可以用定时服务和MQ实现，思考了一下决定用mq，理由如下：

**动态伸缩：**MQ可以根据流量增加POD节点的数量，在不改动代码的情况下提高数据的吞吐量

**性能瓶颈：**定时服务消费，我不确定一个周期的执行时间是多久，比如10分钟同步一个周期，10分钟消费1000sku，能不能消费完，如果消费完了，剩余的时间不就浪费了性能吗，如果10分钟加到10000sku，但是这个周期同步不完怎么办，下个周期继续同步吗？两个任务之间会不会有冲突产生？有安全隐患，但是mq每个消费者之间的进程都是独立的没有这种隐患

**心跳监控**：可以通过在前端看mq的剩余消费行数，看服务是否正常在执行，如果是定时服务，只能到pod节点下面去找日志，这个也比较麻烦，因为不只你一个人的日志

**与服务端解耦**： 消费端如果逻辑有变动可以快速更新，不需要更新主业务端

**异步处理**： 主业务方在界面上操作后，不需要等待事件处理完成，直接可以做其他的操作，等处理完了再回来看就可以

**容错性与可靠性**： Rabbitmq消息处理失败了也不会丢失，系统故障了也不会丢失，如果用线程或者异步处理会有这种情况

### RabbitMQ安装-Windows

[RabbitMQ安装](https://www.rabbitmq.com/install-windows.html#installer) ，点击这个链接到RabbitMQ官网，准备先安装Erlang环境

![1705543703206](Vue进步.assets/1705543703206.png)

进入Erlang下载界面安装包进行安装，安装好后配置环境变量

![1705544256447](Vue进步.assets/1705544256447.png)

 输入Erl回车，出现下面的提示代表安装成功

![1705544299048](Vue进步.assets/1705544299048.png)

回到开始RabbitMQ下载的链接，安装windows下载包进行下载安装

![1705544464663](Vue进步.assets/1705544464663.png)

安装好后检查服务列表是不是存在RabbitMQ服务

![1705544530543](Vue进步.assets/1705544530543.png)

进入RabbitMQ的安装目录，输入： rabbitmq-plugins enable rabbitmq_management 安装后台管理页面,提示下面的内容说明成功

![1705545519525](Vue进步.assets/1705545519525.png)

成功后重启RabbitMQ的服务

![1705545620496](Vue进步.assets/1705545620496.png)

点击 http://localhost:15672/#/ 进入后台，默认用户名/密码：guest/guest

![1705545684243](Vue进步.assets/1705545684243.png)

### RabbitMQ仓储注入使用



本仓储的建立有两个目的，第一个目的是为了熟悉在一个新的项目中引用RabbitMQ所涉及到的步骤方法，加深自己对RabbitMQ的理解，第二个目的是建立仓储，为了后续方便使用RabbitMQ

**引用类库**

> Install-Package RabbitMQ.Client -Version 5.2.0   --类库有版本兼容问题
>
> Install-Package Newtonsoft.Json

RabbitMQ使用类：
1.RabbitMQ1项目为发送者的项目:[点击](https://github.com/AtaraxiaRun/Hello-Word/tree/master/RabbitMQ/RabbitMQ1/RabbitMQ1)
2.RabbitMQ2项目为消费者项目，使用BackgroundService[点击](https://github.com/AtaraxiaRun/Hello-Word/tree/master/RabbitMQ/RabbitMQ2/RabbitMQ2)

注意控制台程序获取配置的时候注意选择始终复制


![1705572717608](Vue进步.assets/1705572717608.png)

本项目实现了通过BackgroundService 实现MQ消费者，并且可以支持随意新增消费者：

- **BackgroundService** ：定时任务，后台任务的使用
- **AddHostedService**： services.AddHostedService<SyncProductWork>(); 注册跟随程序一起启动的服务
- **配置始终复制**： 控制台配置配置成始终复制

## Redis

### Redis仓储，声明仓储接口

使用StackExchange.Redis类库进行封装仓储接口，盛为速牛的项目就是用的这个仓储接口

[Redis Windows版本下载（只有3.0版本，不过会自动注册Redis服务）](https://github.com/microsoftarchive/redis) ，参考链接：https://www.cnblogs.com/liuqingzheng/p/9831331.html

**引用类库**

> Install-Package StackExchange.Redis
>
> Install-Package Newtonsoft.Json

**Program**：服务注入，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/Redis%E4%BB%93%E5%82%A8/Program.cs)
**BaseLock**：锁基类，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/Redis%E4%BB%93%E5%82%A8/BaseLock.cs)
**HashLock**：锁实现类，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/Redis%E4%BB%93%E5%82%A8/HashLock.cs)
**Lock**：锁实体类，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/Redis%E4%BB%93%E5%82%A8/Lock.cs)
**Redlock**：锁实现类，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/Redis%E4%BB%93%E5%82%A8/Redlock.cs)
**ILock**：锁接口，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/Redis%E4%BB%93%E5%82%A8/ILock.cs)
**IRedisManager**：Redis仓储增删改查接口类，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/Redis%E4%BB%93%E5%82%A8/IRedisManager.cs)
**RedisManager**：Redis仓储增删改查实现类，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/Redis%E4%BB%93%E5%82%A8/RedisManager.cs)
**Product**：产品信息实体类，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/Redis%E4%BB%93%E5%82%A8/Product.cs)
**RedisOptions**：Redis用户信息实体类，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/Redis%E4%BB%93%E5%82%A8/RedisOptions.cs)
**ServiceCollectionExtensions**：IServiceCollection服务注入扩展方法，[点击](https://github.com/AtaraxiaRun/Hello-Word/blob/master/Redis%E4%BB%93%E5%82%A8/ServiceCollectionExtensions.cs)

这个项目实现了使用 StackExchange.Redis 增删改查的功能，并且使用了锁机制，可以在后续的项目中进行使用，并且学习到了Set类型保存集合信息

### Redis并发分布式互斥锁

> **.Net 下基于Redlock redis 分布式锁实现**
>
> 本文使用【RedLock.net】(C#/.NET implementation). Includes async and lock extension support. https://github.com/samcook/RedLock.net
>
> Redlock-cs (C#/.NET implementation). https://github.com/kidfashion/redlock-cs
>
> ScarletLock(C# .NET implementation with configurable datastore) https://github.com/aidapsibr/scarletlock
> Redlock4Net(C# .NET implementation) https://github.com/LiZhenNet/Redlock4Net

`基于RedLock.net`

```c#
//调用
using RedisLock;
var lockEans = await UniqueKeyGenerator.ConcurrentLockCall();
var noLockEans =await UniqueKeyGenerator.ConcurrentCall();
var ean3 = string.Join(",", noLockEans);
var ean4 = string.Join(",", lockEans);


Console.WriteLine($"不加锁的Ean：{ean3}");
Console.WriteLine($"加锁的Ean：{ean4}");

Console.WriteLine("********************************************************");
var ean1 = string.Join(",", noLockEans.GroupBy(i => i)
                     .Where(g => g.Count() > 1)
                     .Select(g => g.Key));
var ean2 = string.Join(",", lockEans.GroupBy(i => i)
                     .Where(g => g.Count() > 1)
                     .Select(g => g.Key));


Console.WriteLine($"不加锁的输出：{ean1}");
Console.WriteLine($"加锁的输出：{ean2}");


//代码
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using StackExchange.Redis;
using RedLockNet.SERedis;
using RedLockNet.SERedis.Configuration;
using System.Net;
using Newtonsoft.Json;

namespace RedisLock
{
    /// <summary>
    /// RedLock.net
    /// </summary>
    public static class UniqueKeyGenerator
    {
        private static Lazy<ConnectionMultiplexer> lazyConnection = new Lazy<ConnectionMultiplexer>(() =>
        {
            return ConnectionMultiplexer.Connect("localhost");
        });

        public static ConnectionMultiplexer Connection
        {
            get
            {
                return lazyConnection.Value;
            }
        }

        public static async Task<List<string>> GenerateUniqueKeyAsync()
        {
            string key = "Real_Publish_Ean_Lock";
            //用这个方法，有Bug，会导致redLockFactory.CreateLockAsync一直获取不到锁
            //var redLockFactory = RedLockFactory.Create(new List<RedLockEndPoint>
            // {
            //    new DnsEndPoint("localhost", 6379)
            // });
            var existingConnectionMultiplexer1 = ConnectionMultiplexer.Connect("127.0.0.1:6379");
            var multiplexers = new List<RedLockMultiplexer>
                {
                    existingConnectionMultiplexer1,
                };
            var redLockFactory = RedLockFactory.Create(multiplexers);

            //定义锁参数，过期时间为一分钟，等待时间为一分钟，重试时间为10秒
            var expiry = TimeSpan.FromMinutes(1);
            var wait = TimeSpan.FromMinutes(1);
            var retry = TimeSpan.FromSeconds(5);
            try
            {
                await using (var redLock = await redLockFactory.CreateLockAsync(key, expiry, wait, retry))
                {
                    // 确保我们取得了锁

                    if (redLock.IsAcquired)
                    {
                        // 获取Ean
                        var uniqueKey = await GetEans();
                        return uniqueKey;
                    }
                    else
                    {
                        Console.WriteLine($"{DateTime.Now},无法获取Ean，锁在一分钟内没有被释放");
                    }
                }
            }
            catch (Exception ex)
            {

                throw new Exception($"无法获取Ean ,Error：{ex.Message}");
            }
            finally
            {
                redLockFactory.Dispose();
            }

            throw new Exception("无法获取Ean，锁在一分钟内没有被释放");
        }

        private static async Task<List<string>> GetEans()
        {
            var url = "https://admin.shengweiinc.com/WalmartApi/UpcManage/GetUpcManagenByName?numberPoolName=天津项目组EAN&&num=10";
            using (HttpClient client = new HttpClient())
            {
                var response = await client.GetAsync(url);
                response.EnsureSuccessStatusCode(); // 抛出异常如果响应状态码不是成功的
                var responseBody = await response.Content.ReadAsStringAsync();
                if (responseBody == null)
                {
                    return new List<string>();
                }
                return JsonConvert.DeserializeObject<List<string>>(responseBody);
            }
        }

        /// <summary>
        /// 不带锁的并发调用
        /// </summary>
        /// <returns></returns>
        public static async Task<List<string>> ConcurrentCall()
        {
            var tasks = Enumerable.Range(0, 3).Select(_ => GetEans()).ToArray();
            var results = await Task.WhenAll(tasks);
            return results.SelectMany(x => x).ToList();
        }

        /// <summary>
        /// 带锁的并发调用
        /// </summary>
        /// <returns></returns>
        public static async Task<List<string>> ConcurrentLockCall()
        {
            var tasks = Enumerable.Range(0, 3).Select(_ => GenerateUniqueKeyAsync()).ToArray();
            var results = await Task.WhenAll(tasks);
            return results.SelectMany(x => x).ToList();

        }
    }
}

```

`基于CSRedis`

## Devops

DevOps可以帮助软件团队通过协作、自动化、监控和持续改进的方式，更好地管理和交付软件项目，提高质量、效率和用户满意度

**Gitlab Git Docker  k8s jenkins harbor Rancher （管理k8s）** 

**为什么要学习Despov?**

**故事：**

有一个小村庄里的农民叫小明。小明有一片农田，他想要种植一些作物。他发现每次他种植作物时，总会遇到一些问题。有时候他的庄稼得不到足够的水分，有时候他的作物会遭受虫害，还有时候他的庄稼会被严重的干旱所影响。

小明决定寻求帮助，于是他请来了两位专家：小花和小杰。小花是一位植物学家，她了解植物的需求和成长过程。小杰是一位水利工程师，他懂得如何管理水资源。

他们三个人决定共同合作，他们开始实施**DevOps**的原则和方法。

首先，他们建立了一个团队，每天定期开会讨论庄稼的情况。他们分享彼此的知识，小花解释了作物的需求，小杰提供了关于水资源的信息。

然后，他们使用自动化工具来帮助他们更高效地完成工作。他们购买了一台自动灌溉系统，这样就可以自动为庄稼浇水，确保作物得到足够的水分，而不需要小明每次都手动浇水。

接着，他们开始使用监控工具来观察庄稼的生长情况。他们安装了一个天气传感器，可以监测天气情况和土壤湿度。当传感器检测到干旱时，系统会自动启动灌溉系统，确保庄稼不会因为干旱而受到损害。

最后，他们使用持续改进的方法来不断优化他们的工作流程。他们定期回顾他们的农田管理过程，寻找改进的机会。他们发现通过合理轮作作物，使用有机肥料等措施，可以进一步提高作物的产量和质量。

通过小明、小花和小杰的合作，他们的庄稼茁壮成长，产量也有了显著的提高。他们的工作变得更高效、更自动化，他们不再担心庄稼的生长问题，可以专注于其他重要的事情。

这个故事告诉我们，**DevOps就像小明、小花和小杰的合作一样，通过协作、自动化、监控和持续改进，可以帮助团队更好地管理和发展软件项目**。

在软件开发中，就像小明的农田一样，团队面临许多挑战，例如代码错误、部署问题和性能瓶颈。通过应用DevOps原则，团队可以实现以下几点：

1. **协作与沟通**：团队成员之间的合作与沟通至关重要。就像小明、小花和小杰一样，软件开发人员、测试人员和运维人员需要定期开会，分享知识和经验，以便更好地理解项目需求和解决问题。
2. **自动化**：DevOps鼓励使用自动化工具来提高效率和一致性。就像小明购买了自动灌溉系统一样，软件团队可以使用自动化构建工具、测试工具和部署工具，以减少人为错误并提高交付速度。
3. **监控与反馈**：类似于小明使用天气传感器监测庄稼生长情况，团队可以使用监控工具来观察软件系统的运行状态。通过监控关键指标和日志数据，团队可以及时发现问题并采取措施解决。
4. **持续交付与部署**：就像小明的庄稼需要定期浇水和及时施肥一样，团队可以使用持续集成和持续交付的实践，确保软件的频繁交付和快速部署。这有助于减少错误、降低风险，并提供更好的用户体验。
5. **持续改进**：与小明、小花和小杰反思农田管理过程一样，软件团队应定期回顾项目流程，并寻找改进的机会。通过反馈和数据驱动的决策，团队可以不断优化开发流程和提高软件质量。

总而言之，DevOps可以帮助软件团队通过协作、自动化、监控和持续改进的方式，更好地管理和交付软件项目，提高质量、效率和用户满意度。就像小明、小花和小杰一起努力，DevOps促进了整个团队的成功。

### Portainer

Portainer Docker图形化工具，很给力，扩展阅读：https://www.jianshu.com/p/e76054edabe1

[微服务课程链接](https://www.tutorialspoint.com/microservices-net-core-masterclass/index.asp)，我在思考，为什么我不知道这个工具，开始我以为这个是个新的工具，但是一看才发现，这个工具

![1688617776138](Vue进步.assets/1688617776138.png)

### Rancher：国外用得容器管理

Rancher是一款开源的企业级容器管理平台。它的主要功能和用途包括：

1. **容器管理：** Rancher 提供了一种方便的方式来运行和管理 Docker 和 Kubernetes。通过 Rancher 的用户界面，你可以快速部署、管理和扩展容器。
2. **Kubernetes 集群管理：** Rancher 提供了创建和管理 Kubernetes 集群的完整解决方案。Rancher 支持多种云服务提供商，也可以在现有的裸机服务器上部署 Kubernetes。Rancher 还提供了许多用于管理 Kubernetes 集群的工具，如多集群管理、集群导入、集群模板等。
3. **多租户和访问控制：** Rancher 提供了多租户和访问控制的功能，使你可以为每个用户或团队创建和管理独立的 Kubernetes 名称空间。
4. **集成 CI/CD：** Rancher 可以集成 Jenkins、GitLab CI/CD、CircleCI 等 CI/CD 工具，使你可以更容易地实现持续集成和持续交付。
5. **服务发现和负载均衡：** Rancher 提供了服务发现和负载均衡的功能，使你可以更容易地在 Kubernetes 集群中部署和扩展应用。
6. **持久化存储：** Rancher 支持多种存储插件，你可以根据需要选择最适合你的存储解决方案。

综上所述，如果你正在寻找一个企业级的容器管理平台，能够管理和运行 Docker 和 Kubernetes，提供多租户管理、集成 CI/CD、服务发现、负载均衡以及持久化存储等功能，那么 Rancher 就是一个非常好的选择。

```
# 查找镜像资源
docker search rancher
#拉取镜像文件
docker pull rancher/server:stable
#启动rancher容器
docker run -d --restart=unless-stopped \
  -p 8080:8080 \
  --privileged \
  --name rancher  rancher/server:stable
  
docker run -d --restart=unless-stopped -p 8080:8080 rancher/rancher:latest
  
```

参考链接：

Rancher是一个开源的企业级全栈化容器部署及管理平台。https://www.cnblogs.com/naixin007/p/11009741.html 

Rancher 使用介绍（可以通过界面管理 K8s 平台） https://blog.csdn.net/weixin_46902396/article/details/122433622

如何在 Kubernetes 中使用 Rancher  https://www.trendmicro.com/en_us/devops/21/b/rancher-and-trend-micro-container-security.html

### Kubernetes



### Harbor：容器管理

Harbor 是一个开源的企业级 Docker Registry 项目，起初是由 VMware 公司开发并开源，后来成为了 CNCF（Cloud Native Computing Foundation，云原生计算基金会）的孵化项目。它主要用于存储、签名、扫描和分发 Docker 镜像。

下面是 Harbor 的主要功能和用途：

1. **镜像存储：** Harbor 提供了一个安全的、策略严格的 Docker 镜像存储库。它可以存储企业或开发者创建的 Docker 镜像，以供开发、测试和生产环境使用。
2. **镜像分发：** Harbor 支持镜像的复制和分发，这对于多数据中心或者边缘计算场景非常有用。它可以将镜像从一个 Harbor 实例复制到另一个 Harbor 实例，从而减少跨数据中心的网络延迟。
3. **镜像扫描：** Harbor 集成了 Clair，一个开源的漏洞扫描项目。这使得 Harbor 能够扫描存储在其上的 Docker 镜像，找出可能存在的安全漏洞。
4. **镜像签名：** Harbor 支持 Docker Content Trust，可以对 Docker 镜像进行签名。这保证了 Docker 镜像的完整性和来源可追溯性。
5. **访问控制：** Harbor 提供了详细的访问控制策略，可以为不同的用户或团队分配不同的访问权限。
6. **多租户：** Harbor 支持多租户管理，每个用户可以有自己的私有项目和镜像。

因此，如果你的企业或项目需要使用 Docker，并且需要一个安全、可靠的镜像存储和分发解决方案，那么 Harbor 就是一个非常好的选择。

```mysql

#下载安装包
#为什么不通过Docker安装Harbor呢？因为Harbor的镜像包是独立的，分为权限，日志，数据库，缓存这么多个，直接安装这个包，会帮我们自动用镜像启动Harbor服务
wget https://github.com/goharbor/harbor/releases/download/v2.0.0/harbor-online-installer-v2.0.0.tgz
wget https://github.com/goharbor/harbor/releases/download/v2.8.2/harbor-online-installer-v2.8.2.tgz
#解压缩
tar xvf harbor-online-installer-v2.8.2.tgz
#解压之后，会有一个 harbor 的目录。
#配置 Harbor：进入 harbor 目录，你会看到一个名为 harbor.yml.tmpl 的配置模板文件，复制这个文件并重命名为 harbor.yml。
cd harbor
cp harbor.yml.tmpl harbor.yml

#方括号包住的内容，直接双击进去编辑barbor.yml的内容就可以,记得注释下面的内容
# https related config
#https:
  # https port for harbor, default is 443
#  port: 443
  # The path of cert and key files for nginx
#  certificate: /your/certificate/path
#  private_key: /your/private/key/path
{
# 删除原有的https验证方式，并且重新添加端口映射为：8082端口
sed -i '/https:/,+4d' harbor.yml   
echo "http:
  port: 8082" >> harbor.yml

#更新主机名hostname: reg.mydomain.com是旧值/hostname: 192.128.226.129 这个是新值，就是服务器的IP地址
sed -i 's/hostname: reg.mydomain.com/hostname: 192.128.226.129/' harbor.yml
#更新管理员密码，
sed -i 's/harbor_admin_password: Harbor12345/harbor_admin_password: admin123/' harbor.yml
#更新数据库密码，Harbor 默认会在 Docker 容器中启动一个 PostgreSQL关系数据库实例，保存项目、用户、角色、权限、镜像信息
sed -i 's/password: root123/password: db123/' harbor.yml

}

#扩大docker-compose的超时时间，默认是60秒，设置成300秒，因为./install.sh里面要下载权限，数据库，日志等镜像模块，要很久时间
export COMPOSE_HTTP_TIMEOUT=3000
#更改配置后，更新其他组件的配置，比如nginx,log,db,service等组件,每次更改配置后都需要执行这个命令
#【这一步如果有访问docker hub镜像超时错误，那么去看左边目录下面：Linux/安装Docker 模块中最下面设置代理的内容，设置代理后就好了】
./prepare     //这一步是根据harbor.yml的内容生成docker-compose配置的步骤

#启动服务，这个脚本会自动使用 Docker Compose 下载 Harbor 相关的所有的镜像，然后启动容器服务
./install.sh

#验证是否登录成功
192.128.226.129:8080
上面的访问不了，进度卡在这里了，继续努力

#使用harbor管理镜像
docker login -u admin -p admin123 192.168.226.129:8082   # 默认用户名是admin（看下面/Jenkins/Jenkins Push Har/更新Docker限制Http执行环境的限制内容设置后才可以访问）
docker tag myimage 192.168.226.129:8082/testproject/myimage:v1.0
docker push 192.168.226.129:8082/testproject/myimage:v1.0

#【备用脚本1】:解决提示Redis版本问题,用下面的解决方案
 #1.下载符合版本要求的镜像
docker pull goharbor/redis-photon:v2.5.2
 #2.更新/root/harbor/docker-compose.yml中的版本号
 redis:
    image: goharbor/redis-photon:v2.5.2  //直接进去把这个版本号更改成符合镜像要求的版本号
    container_name: redis
    restart: always
  #3.重启所有的docker-compose中的镜像
  docker-compose up -d 
  
#【备用脚本2】：解决 postgresql 数据库登录，一直提示密码不对问题
    #【直接重新设置密码就可以】
  docker exec -it 17355661c338 bash -c "psql -U postgres -c \"ALTER USER postgres PASSWORD 'root123';\""

  #1.使用root权限更新goharbor/harbor-db:v2.0.0容器中的yum，安装postgresql客户端
  docker exec -u 0 -it 71d2c968a649 /bin/bash -c "yum update && yum update postgresql"
  #2.验证是否正常安装
  docker exec -it 71d2c968a649 bash
  psql --version   // 看到了版本号就可以
  #3.登录一下，默认的数据库：registry  ，默认的用户名：postgres ，默认的用户密码：root123
  psql -U postgres -d registry    //如果你看到 "registry=#" 提示符，那么这意味着你已经成功连接到了数据库，你可以尝试运行一些命令，比如查询所有的表：\dt,或者退出数据库：\q
  #4.如果在容器外部登录，那么需要更新docker-compose.yml,新增ports 5432:5432 端口映射的内容
   postgresql:
    image: goharbor/harbor-db:v2.0.0
    container_name: harbor-db
    restart: always
    ports:
      - 5432:5432    //这两行代码是新增的
  #5.重新启动Harbor服务
  docker-compose up -d 
  #6.重新登录,localhost本机IP:192.168.226.129,默认的数据库：registry  ，默认的用户名：postgres ，默认的用户密码：root123
  psql -h 192.168.226.129 -U postgres -d registry  //回车后会要求你输入密码，默认密码:root123
 
 
  
#Harbor镜像解释
1.goharbor/harbor-jobservice:v2.0.0: 这是 Harbor 的 job service 组件，负责处理和调度一些长时间运行的后台任务，例如图像扫描，垃圾收集和复制。

2.goharbor/nginx-photon:v2.0.0: 这是 Harbor 的反向代理服务器，由 Nginx 提供。它处理并转发到 Harbor 服务的所有 HTTP 请求。

3.goharbor/harbor-core:v2.0.0: 这是 Harbor 的核心服务组件，负责处理核心业务逻辑，例如用户认证，事件处理和项目管理。

4.goharbor/harbor-db:v2.0.0: 这是 Harbor 的数据库服务组件，由 PostgreSQL 提供。它存储了所有的 Harbor 数据。

5.goharbor/redis-photon:v2.0.0: 这是 Harbor 的 Redis 服务组件，负责提供缓存和其他可能需要的功能。

6.*goharbor/harbor-portal:v2.0.0*: 这是 Harbor 的前端 Web 用户界面。用户可以通过此界面管理 Harbor 实例。

7.goharbor/registry-photon:v2.0.0: 这是 Docker Registry 服务组件，存储并分发 Docker 容器镜像。

8.goharbor/harbor-registryctl:v2.0.0: 这是 Harbor 的 registry 控制服务，负责管理和控制 Docker registry。

9.goharbor/harbor-log:v2.0.0: 这是 Harbor 的日志收集服务，用于收集和管理 Harbor 中所有组件的日志。

这些镜像一起工作，提供了一个完整的 Harbor 服务，用于存储，分发和扫描 Docker 容器镜像。

```



### Jenkins

Jenkins 是一款开源的持续集成（Continuous Integration，简称 CI）和持续交付（Continuous Delivery，简称 CD）工具，起源于 Hudson 项目，它由 Java 编写并运行在 Servlet 容器中（如 Apache Tomcat）。

以下是 Jenkins 的主要功能和用途：

1. **持续集成/持续交付：** Jenkins 最核心的功能就是持续集成和持续交付。在软件开发中，开发人员将代码变更集成到主分支的频率越高，问题就越能早期发现、早期解决。这种实践被称为持续集成。持续交付则是持续集成的下一步，它确保了你可以随时将你的应用发布到生产环境中。
2. **构建自动化：** Jenkins 可以执行各种构建任务，包括编译源代码、运行自动化测试、打包应用等。这些任务可以通过 Jenkinsfile 配置，并且可以以管道的形式进行。
3. **测试自动化：** Jenkins 可以集成各种测试工具，如 JUnit、Selenium 等，自动执行测试任务，并提供详细的测试报告。
4. **部署自动化：** Jenkins 可以自动部署应用到各种环境中，包括开发环境、测试环境和生产环境。
5. **插件系统：** Jenkins 有一个丰富的插件生态，几乎所有的构建、测试和部署工具都有对应的 Jenkins 插件。通过插件，你可以扩展 Jenkins 的功能，并将其集成到你的工具链中。
6. **分布式构建：** Jenkins 支持分布式构建，可以在多台机器上并行执行构建任务，从而加快构建速度。

因此，如果你的团队正在寻找一个能够自动化软件构建、测试和部署流程的工具，那么 Jenkins 是一个很好的选择



```mysql
{
经验分享
1.jenkins/jenkins_home与/jenkins/jenkins_home路径，要注意
2.扫描失败了，检查一下https://github.com/AtaraxiaRun/Hello-Word.git项目的权限是不是私有
3.新建项目是倒数第二个那个，多分支管道

}




# 创建Dockerfile，拉取Jenkins的镜像，下面是Dockerfile文件的解释
# 1.下载Jenkins的镜像
# 2.替换apt-get update的软件镜像源为清华大学的镜像源，这样的话用国内的线路会快很多，不然真的更新一下要花十几分钟
# 3.更新好后，下载docker-ce-cli软件包，以便可以在jenkins容器中运行dockerfile命令生成镜像推送到harbor
# 4.然后将jenkins用户组，加入到docker用户组获取访问宿主机/var/run/docker.sock的权限，groupadd docker && \ usermod -aG docker jenkins （在容器内部用jenkins用户获取操作所有docker容器的权限，其实还是有一定的风险的，开发环境可以用，生产环境不建议用）
# 5.并且需要传入一个DOCKER_GID的参数，是宿主机的docker用户组的ID，然后让Jenkins绑定到这个组，就可以使用/var/run/docker.sock的权限
cat > Dockerfile << EOF
FROM jenkins/jenkins:2.410-jdk11

ARG DOCKER_GID

USER root

RUN echo "deb http://mirrors.tuna.tsinghua.edu.cn/debian/ buster main contrib non-free\n\
deb http://mirrors.tuna.tsinghua.edu.cn/debian/ buster-updates main contrib non-free\n\
deb http://mirrors.tuna.tsinghua.edu.cn/debian/ buster-backports main contrib non-free\n\
deb http://mirrors.tuna.tsinghua.edu.cn/debian-security buster/updates main contrib non-free" > /etc/apt/sources.list && \
    apt-get update && \
    apt-get -y install apt-transport-https \
         ca-certificates curl gnupg2 \
         software-properties-common && \
    curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add - && \
    add-apt-repository \
       "deb [arch=amd64] https://download.docker.com/linux/debian \
       buster \
       stable" && \
    apt-get update && \
    apt-get -y install docker-ce-cli && \
    groupadd -g $DOCKER_GID docker_host && \
    usermod -aG docker_host jenkins

USER jenkins
EOF
#【！！！执行完上面的命令后，手动打开dockerfile检查一下groupadd -g $DOCKER_GID docker_host中的$DOCKER_GID参数是否正常赋值了】


# 生成镜像，有一个docker_gid的参数是动态的，需要动态入参，这两行命令需要一起运行的，linux命令还可以这样声明参数入参，真的牛
DOCKER_GID=$(getent group docker | cut -d: -f3)
docker build --build-arg DOCKER_GID=$DOCKER_GID -t myjenkins . //上下这两句要一起执行



#新建挂载路径
mkdir -p /jenkins/jenkins_home 

#给挂载的路径开在容器里面的操作权限,不然运行容器后，里面挂载的路径Jenkins容器没有操作写入权限
sudo chmod 777 /jenkins/jenkins_home

# 运行容器,挂载/var/run/docker.sock,Jenkins容器与宿主机的Docker守护进程进行通信
docker run -d --name=jenkins.service --restart=always -p 8081:8080 -p 50000:50000 -v /var/run/docker.sock:/var/run/docker.sock -v /jenkins/jenkins_home:/var/jenkins_home -e JENKINS_ADMIN_ID=admin -e JENKINS_ADMIN_PASSWORD=123 myjenkins

/* 2.410版本没有密码了，好像2.22之前老版本会有,查找登录密码(一般都在日志最下面) :Please use the following password to proceed to installation 下面的就是密码，比如：Jenkins initial setup is required.  An admin user has been created and a password generated.
Please use the following password to proceed to installation:

1d8381cf02c34438b887a1ece2d2a6e0    //这个就是密码
*/
docker logs -f jenkins.service

#进去容器看看，能不能执行docker命令，成功生成镜像
docker exec -it 17bb7e8a63fcf266adf3 bash
# 验证docker客户端是否正常安装
docker version

#登录验证
curl -u admin:1d8381cf02c34438b887a1ece2d2a6e0 http://192.168.226.129:8081/

#浏览器访问
http://192.168.226.129:8081/

以下为备用脚本信息
#Github访问令牌,参考地址：https://docs.cloudbees.com/docs/cloudbees-ci-kb/latest/client-and-managed-masters/github-user-scopes-and-organization-permissions-overview
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


#判断正常访问宿主机的docker version返回值
jenkins@2162b161c157:/$ docker version
Client: Docker Engine - Community
 Version:           24.0.2
 API version:       1.41 (downgraded from 1.43)
 Go version:        go1.20.4
 Git commit:        cb74dfc
 Built:             Thu May 25 21:52:12 2023
 OS/Arch:           linux/amd64
 Context:           default

Server: Docker Engine - Community
 Engine:
  Version:          20.10.21
  API version:      1.41 (minimum version 1.12)
  Go version:       go1.18.7
  Git commit:       3056208
  Built:            Tue Oct 25 18:02:38 2022
  OS/Arch:          linux/amd64
  Experimental:     false
 containerd:
  Version:          1.6.12
  GitCommit:        a05d175400b1145e5e6a735a6710579d181e7fb0
 runc:
  Version:          1.1.4
  GitCommit:        v1.1.4-0-g5fd4c4d
 docker-init:
  Version:          0.19.0
  GitCommit:        de40ad0

```

#### **Jenkins界面操作**

启动Jenkins后，输入docker logs -f 容器ID，这个就是登录用户名的密码

![1687914544457](Vue进步.assets/1687914544457.png)

输入http://192.168.226.128:8081/ 登录Jenkins，这里Administrator password 输入上面的密码：fad3b0e747174c1cb83b7ec73912b620

![1687914575212](Vue进步.assets/1687914575212.png)

安装Github插件支持，多分支管道构建支持等插件，还可以安装与钉钉交流的插件，大约5~10分钟
![1687914804210](Vue进步.assets/1687914804210.png)

这个用户是用来后面登录Jenkins的，毕竟用长的GUID码登录还是很不好记

![1687915223916](Vue进步.assets/1687915223916.png)

保存后续登录与其他应用程序交互的网址

![1687915272632](Vue进步.assets/1687915272632.png)

开始使用Jenkins

![1687915301467](Vue进步.assets/1687915301467.png)

创建新的项目

![1687915402322](Vue进步.assets/1687915402322.png)

输入项目名称，选择项目构建方式

![1687915498131](Vue进步.assets/1687915498131.png)

设置项目名称，描述，登录凭证

![1687916042122](Vue进步.assets/1687916042122.png)

输入Github凭证的用户名与密码，点击Add 添加

![1687916135189](Vue进步.assets/1687916135189.png)

选择登录凭证，输入存储库的地址

![1687916242737](Vue进步.assets/1687916242737.png)

设置Jenkinsfile脚本文件路径，里面有生成镜像，发布镜像到镜像仓库从镜像仓库拉取部署到目标服务器的脚本，一般像Dockerfile文件一样，放到项目的根目录下面

![1687917203943](Vue进步.assets/1687917203943.png)

Jenkinsfile就放在我Github程序的根目录,注意还有Dockerfile，因为我Jenkinsfile里面调用了Dockerfile

![1687917461591](Vue进步.assets/1687917461591.png)

Jenkinsfile文件内容**此处是推送到DockerHub案例**，下面的jenkinsfile的文件是有三个步骤：1.构建 2.推送 3.部署 ，实际生成过程中我们可以按我们自己的需求新增或者减少其中的步骤，比如我们加入测试步骤，自动生成Api文档的步骤，只需要加一个stage('test name')节点就可以了，jenkinsfile脚本是按照顺序执行的，我们可以添加任意数量的stage步骤方法，里面的方法名称也是可以自定义的



```c#
**此处是推送到DockerHub案例**
pipeline {
    agent any

    stages {
        stage('Build') { // 1.构建步骤：在此处进行您的构建操作，生成镜像
            steps {
                
                // 例如，使用 Dockerfile 构建您的项目镜像【注意这个Dockerfile也是要提前准备在项目的根目录】
                sh 'docker build -t consoleapp6-image:1.0 .'
            }
        }

        stage('Publish') { // 2.推送步骤：推送构建好的镜像到镜像仓库，比如DockerHub,Harbor仓库
            steps {
                script {
                    // 登录到 Docker Hub
                    sh 'docker login -u libaibaihi -p qSNz9ofsnKb00A4c'

                    // 标记镜像
                    sh 'docker tag consoleapp6-image:1.0 libaibaihi/consoleapp6:1.0'

                    // 推送镜像到 Docker Hub
                    sh 'docker push libaibaihi/consoleapp6:1.0'
                }
            }
        }

        stage('Deploy') { //3.部署步骤（下面为案例，实际是无法执行的）：将镜像仓库的镜像部署到目标服务器，腾讯云中 
			steps {
				// 将镜像保存为tar文件
				sh 'docker save -o consoleapp6.tar libaibaihi/consoleapp6:1.0'

				// 使用SSH传输tar文件到目标服务器
				sshagent(['your-credentials-id']) {
					sh "scp -i /path/to/ssh/private/key consoleapp6.tar root@192.168.226.128:/path/to/destination"
				}

				// 在目标服务器上加载镜像
				sshagent(['your-credentials-id']) {
					sh "ssh -i /path/to/ssh/private/key root@192.168.226.128 'docker load -i /path/to/destination/consoleapp6.tar'"
				}

				// 在目标服务器上运行容器
				sshagent(['your-credentials-id']) {
					sh "ssh -i /path/to/ssh/private/key root@192.168.226.128 'docker run -d -p 7071:7071 -p 7072:7072 --restart=always --name consoleapp6 -v /usr/local/webapp/RedisDemo/:/app libaibaihi/consoleapp6:1.0 --urls=\"http://*:7071;http://*:7072\"'"
				}
			}
		}

    }
}

```



提交，提交后会自动扫描代码库，看是否能正常拉取Github的代码

![1687918103284](Vue进步.assets/1687918103284.png)

扫描存储库成功

![1687918322084](Vue进步.assets/1687918322084.png)

点击左上角的仪盘表回到主界面，简单介绍一下基本的功能按钮作用

![1687918556469](Vue进步.assets/1687918556469.png)

点击具体的项目，进入项目，进去以后可以构建项目

![1687918638266](Vue进步.assets/1687918638266.png)

构建的步骤与方法

![1687918935824](Vue进步.assets/1687918935824.png)

查看构建是否成功，查看构建的详细信息

![1687919066820](Vue进步.assets/1687919066820.png)

查看控制台信息

![1687919163864](Vue进步.assets/1687919163864.png)

访问https://hub.docker.com/repositories/libaibaihi，查看镜像是否正常推送到dockerhub,看结果是推送成功

![1687919217368](Vue进步.assets/1687919217368.png)

#### Jenkins Push Harbor

更新Docker限制Http执行环境的限制

```mysql
#编辑/etc/docker/daemon.json文件
vi /etc/docker/daemon.json
#在后面加一行,D让ocker客户端不对指定的Harbor服务器应用HTTPS限制
"insecure-registries": ["192.168.226.130:8082"]
#保存后重启docker服务，刷新配置文件
sudo systemctl restart docker

```

新建项目，推送的时候需要设置一个项目名称

![1687936344586](Vue进步.assets/1687936344586.png)

新建Jenkinsfile,可以推送到Harbor中去，这个是项目根目录下面的Jenkinsfile，替换**Jenkins界面操作**里面推送到Docker Hub里面的代码

```c#
pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                // 在此处进行构建操作，生成镜像
                // 例如使用 Dockerfile 构建项目镜像
                sh 'docker build -t consoleapp6-image:1.0 .'
            }
        }

        stage('Publish') {
            steps {
                script {
                    // 登录到 Harbor
                    withDockerRegistry([credentialsId: 'Harbor_User_003', url: 'http://192.168.226.130:8082']) {
                        // 标记镜像
                        sh 'docker tag consoleapp6-image:1.0 192.168.226.130:8082/text/consoleapp6-image:1.0'

                        // 推送镜像到 Harbor
                        sh 'docker push 192.168.226.130:8082/text/consoleapp6-image:1.0'
                    }
                }
            }
        }
    }
}
```

**问题**：为什么需要让Jenkins容器与宿主机的Docker守护进程进行通行，不是内部安装了 apt-get -y install docker-ce-cli，直接用Jenkins容器的这个docker生成镜像不就可以了吗？

**回答**：这个问题的答案涉及到Docker的运行方式和架构。

在Linux中，Docker的运行实际上是由两部分组成的：**Docker客户端和Docker守护进程**。Docker客户端是你在命令行中运行的"docker"命令，而Docker守护进程则是在后台运行的，负责实际的镜像管理、容器创建、容器运行等任务。

当你在Docker容器内部运行"apt-get -y install docker-ce-cli"命令时，你实际上只是在容器内安装了Docker客户端，而并没有安装Docker守护进程。这是因为Docker设计为只能在宿主机上运行Docker守护进程，而不能在容器内运行Docker守护进程。这是由Docker的安全模型和容器的隔离性决定的。

因此，即使你在Jenkins容器内部安装了Docker客户端，但如果你没有将宿主机的Docker套接字映射到Jenkins容器内，那么Jenkins容器内的Docker客户端将无法与Docker守护进程进行通信，也就无法执行任何Docker操作。

所以，我们需要将宿主机的Docker套接字映射到Jenkins容器内，这样Jenkins容器内的Docker客户端才能通过这个套接字与宿主机的Docker守护进程进行通信，从而执行Docker操作，比如构建镜像、运行容器等。

另外，虽然我们可以通过这种方式让Jenkins容器执行Docker操作，但这同时也意味着Jenkins容器能够控制宿主机上的所有Docker资源，这可能带来安全风险。因此，在生产环境中，我们通常会采取更复杂的方案来实现安全的Docker操作，比如使用Docker-in-Docker，或者使用专门的Docker操作服务等。

**参考链接**

初始化项目文件到到Github https://www.cnblogs.com/babywin/p/6073124.html

K8S基于Jenkins的自动化部署：https://bs.julywhj.cn/archives/k8s-ji-yu-jenkins-de-zi-dong-hua-bu-shu

基于jenkins、gitlab、docker、k8s 实现自动部署微服务：https://www.cnblogs.com/qlqwjy/p/15911579.html

【Docker+Jenkins+Harbor+Rancher持续集成部署分布式项目】 [Docker+Jenkins+Harbor+Rancher持续集成部署分布式项目](https://blog.csdn.net/qq_57756904/article/details/126316152?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522168687820316800185879075%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fblog.%2522%257D&request_id=168687820316800185879075&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~blog~first_rank_ecpm_v1~rank_v31_ecpm-3-126316152-null-null.268^v1^koosearch&utm_term=Jenkins&spm=1018.2226.3001.4450)

Docker安装Jenkins，安装的过程很细节 [Docker安装Jenkins，安装的过程很细节](https://blog.csdn.net/qq_57756904/article/details/125697440?ops_request_misc=%257B%2522request%255Fid%2522%253A%2522168687820316800185879075%2522%252C%2522scm%2522%253A%252220140713.130102334.pc%255Fblog.%2522%257D&request_id=168687820316800185879075&biz_id=0&utm_medium=distribute.pc_search_result.none-task-blog-2~blog~first_rank_ecpm_v1~rank_v31_ecpm-1-125697440-null-null.268^v1^koosearch&utm_term=Jenkins&spm=1018.2226.3001.4450)

GitHub 令牌如何使用，如何配置的文章 https://www.cnblogs.com/cuitang/p/16313140.html 
从零开始搭建JENKINS+GITHUB持续集成环境【多图】 https://developer.aliyun.com/article/896733

## Linux



### Docker常用命令

**强制删除容器，即使被守护进程占用也删除（但是如果有容器还在运行的话还是不能删除，如果停止运行了可以删除，镜像被干掉后，重启启动容器后也可以正常运行，说明容器在启动的时候会Copy镜像文件到自身里面）**： 

docker rmi -f d5a11facb68f  ，-f是精髓

**查找镜像被哪些容器占用**:docker ps -a --filter "ancestor=139f84bcb226"

**查看指定容器的信息** ：docker ps -f "name=my-container"

**查看指定端口是否被占用**:sudo netstat -tlnp | grep 8083  ，查看8083端口是否被占用

**从宿主机copy/覆盖文件到容器中**:docker cp /root/nginx/default.conf 3ce5b0c8d745:etc/nginx/conf.d/default.conf

**从容器中拷贝文件到宿主机**:docker cp 3ce5b0c8d745:etc/nginx/conf.d/default.conf /root/nginx/

**根据镜像创建一个临时的容器**: docker run --rm -it --entrypoint /bin/bash libaibaihi/repertory.web.api:1.0 ,libaibaihi/repertory.web.api:1.0是镜像的名称

**查看容器使用的镜像类型**: cat /etc/os-release , Alpine Linux或者Debian/Ubuntu基础镜像，决定用什么命令要执行,比如apt-get update是在Debian/Ubuntu系统中执行的,apk update是在Alpine Linux系统中执行的

**容器中安装常用命令**（Debian/Ubuntu）: 

1.apt-get update ，可以用docker exec -it -u 0 e2528f9263f4 apt update 替换

 2.apt-get install vim       安装vim

 3.apt-get install  telnet   安装telnet
 4.apt-get install  net-tools   安装ifconfig 
 5.apt install iputils-ping   安装ping

6.apt-get install nmap 安装看rabbitmq服务端口是不是通的插件

7.apt-get install curl 安装curl

8.RUN apt-get update && apt-get install -y docker.io  在容器里面安装容器环境，当和另一个容器需要通信的时候使用，比如Consul Template和nginx进行通信

9.docker exec -it e2528f9263f4 /bin/bash usermod -aG docker jenkins ，开通指定的程序访问docker守护进程的权限，在容器内部没有root权限是执行不了这句代码的



**容器中安装常用命令(Alpine Linux)**:

1.apk update 更新,docker exec -it -u 0 f2d3e594b074fb612 apk update 直接以管理员执行（不用进去先下载各种环境）

2.apk add curl 下载curl ，,docker exec -it -u 0 f2d3e594b074fb612 apk add curl  直接以管理员执行

3.RUN apk update && \
    apk add docker  在容器里面安装容器环境，当和另一个容器需要通信的时候使用，比如Consul Template和nginx进行通信

4.docker exec -u 0 -it 71d2c968a649 /bin/bash -c "yum update && yum update postgresql" ， 使用root命令，在容器中运行需要权限才能执行的多行命令（-u 0 ，里面的0代表root），这个命令很给力

**更改容器中的配置参数**: 更改容器rabbitMQ的用户名与密码：docker container update --env RABBITMQ_DEFAULT_USER=newusername --env RABBITMQ_DEFAULT_PASS=newpassword container_name，更改容器的启动方式： docker update --restart=always 80a3ea94880d

**设置容器自启动**:docker update --restart=always 80a3ea94880d ,  在docker run 的时候设置： --restart=always

**列出包含指定名称的镜像**:docker ps | grep rabbit   , rabbit 就是需要包含的镜像名称

**删除容器与镜像**:docker stop $(docker ps -a -q) && docker rm $(docker ps -a -q),   这个单独执行： docker rmi -f $(docker images -a -q) ，再清理Docker缓存： docker builder prune -a，清理未被使用的数据卷：docker volume prune ，清理构建缓存：docker builder prune，**清理未使用的资源网络挂载卷**：docker system prune -a --volumes --force，**删除所有挂载的数据卷**：docker volume rm $(docker volume ls -q)

**清理Dockre空间** : 先看磁盘情况：docker system df，再清理没有使用的镜像、容器、卷和网络：docker system prune -a

**检查容器的所有信息（ip,路径）**:docker inspect dee3d6ba29fa

**导出容器的日志(查找容器的路径用下面的脚本)**:docker logs --since='2023-03-14T00:00:00' --until='2023-03-14T09:00:00'  dee3d6ba29fa   >> /var/lib/docker/overlay2/14ab758b48d6fbb0f2162ce317c6af59e12adabba25f948612aca2593a59d527/merged/2023-03-14.log

**查看容器在宿主机的路径**:docker inspect --format='{{.GraphDriver.Data.MergedDir}}' dee3d6ba29fa

**安装telnet命令**: sudo yum install telnet -y ，然后可以使用telnet 172.17.0.2 15672

**检查容器与主机的映射关系**:docker port some-rabbit ,输出：5672/tcp -> 0.0.0.0:5672
5672/tcp -> :::5672
15672/tcp -> 0.0.0.0:15672
15672/tcp -> :::15672

**查看容器的ip**:docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' e4a6903ec395 ,输出：172.17.0.2



**显示容器正在执行的进程**:docker top 256c4bbd19c4 ,输出：dotnet Sunway.TaskCoravel.Web.dll

**查看最近100条日志**：docker logs --tail=100 256c4bbd19c4

**查看最近30分钟日志**: docker logs --since 30m 256c4bbd19c4

**查看某时间之后日志**：docker logs -t --since="2023-03-07T00:00:00"  256c4bbd19c4

**查看时间范围日志**:docker logs -t --since="2023-03-07T00:00:00" --until "2023-03-07T03:00:00"  256c4bbd19c4

**查看指定容器占用的内存**（静态）：docker stats --no-stream e399225c7d2c    ，--no-trunc 是不间断输出，好像Logs -f

**查看所有容器占用cpu内存**（动态）：docker stats   ,CPU:占用系统CPU百分比,MEM USAGE / LIMIT 当前占用内存/最大占用内容

**容器内部查看宿主机Ip**：docker inspect sunway.taskcoravel.service --format {{.NetworkSettings.Gateway}} ，

输出：172.17.0.1，也可以通过：docker inspect some-rabbit | grep "IPAddress" 获取

**实时展示Docker日志**:   docker logs -f  c2cbd9f35c23

**删除所有docker镜像**：docker rmi -f $(docker images -aq)

**删除所有docker服务** : docker rm -f $(docker ps -aq)

**查看启动容器的服务信息** ：docker inspect c2cbd9f35c23 （容器ID）,展示的结果中变量：WorkingDir 的值是 容器服务安装的目录。

**查看容器镜像的制作过程,制作日志**:**docker history kibana:7.9.3** ,**docker history kibana:7.9.3 --no-trunc** 

**查看指定容器的日志**：docker logs c2cbd9f35c23 （容器ID）/docker logs -f kibana(容器名称) , 这两种方法都可以用来看程序输出的日志(比如console.writeline输出的日志)，也可以用来看外部程序部署产生的启动中的日志，启动情况的日志

**启动一个容器**：docker start c2cbd9f35c23 （容器ID）

**关闭容器**：docker stop c2cbd9f35c23 （容器ID）

**重启容器**： docker restart  c2cbd9f35c23 （容器ID）

**进入一个容器**(输入exit退出)：docker exec -it mysql bash  ，也可以是：docker exec -it mysql sh ，进去了就可以做登录操作了，比如Mysql进去了就可以mysql -u root -p登录，然后输入数据库密码123456，就可以操作数据库,redis也是这样，进入到了redis，我们直接redis-cli就可以操作redis服务，下面连接redis服务的方法是第二个方法

**【设置容器开机自启动】**：docker update --restart=always 容器名称 ，比如docker update --restart=always mysql

**查看镜像的版本号**:docker inspect -f '{{index .Config.Labels "org.opencontainers.image.version"}}' consul

**查看pull的镜像详细信息(可以看到版本信息)**：docker inspect  mysql    ，注意这里的mysql 名字是docker ps -a 中第一列（REPOSITORY ）镜像的名称，比如mysql镜像的详细信息中就有这一列：MYSQL_VERSION 代表的是版本信息

**sudo开头执行意思是管理员执行**：sudo docker pull mysql （可以指定版本号拉取镜像：sudo docker pull mysql:5.7），如果不用sudo 开头，就会提示：在unix上试图连接到Docker守护进程套接字时，权限被拒绝，也可以在开始用su输入管理员密码直接进入

**使用Redis镜像创建一个Redis容器服务：**docker run -p 6380:6380 --name redis-text -d redis

-p 左边是容器使用的端口，右边是需要映射到本地的端口；--name redis-test （容器服务的名称）,最后面的 redis 代表是容器的名称

**连接Redis服务**：docker exec -it redis-text redis-cli         

- -it redis-text 代表redis容器服务的名称，是我们启动的时候自定义的

**停止一个容器服务**：先执行：docker ps ,拿到第一列：ONTAINER ID，然后执行： docker kill ffc1792471da(如果有多个需要停止，后面加空格继续放ID就可以)

**删除一个容器服务**：先执行：docker ps ,拿到第一列：ONTAINER ID，然后执行： docker rm ffc1792471da(如果有多个需要删除，后面加空格继续放ID就可以)

**查看正在执行的容器**：docker ps 

**查看所有的容器：**docker ps  -a

**启动docker服务**：systemctl start docker

**停止docker服务**：systemctl stop docker

**重启docker服务**：systemctl restart docker 

**查看docker状态：**systemctl status docker，看到命令那里有显眼的：Active: active (running)  就是启动中

**查看docker版本：**docker -v

**查看服务器镜像列表** ：docker images  （这个就是看我们装了多少软件的镜像，比如mysql,rabbitMq，redis软件的docker镜像），镜像可以到

**【拉取镜像】**：docker pull 镜像名 比如 ：docker pull redis ，这种拉取代表拉取镜像仓库中redis最新的一个docker版本，下载完成后，该镜像被默认保存到 `/var/lib/docker` 目录下。

**拉取镜像（指定版本号）**：docker pull 镜像名:tag （版本号） 比如： docker pull redis:latest ，这个也是拉取最新版本,latest就是执行版本的名称/号码

**保存镜像**：docker save 镜像名/镜像ID -o 镜像保存的名字（这个是内部本地保存的名称） 镜像保存的tag（这个是下架的名称），例如：docker save -o redis.tar redis:latest

**加载镜像（保存后就可以加载了）**：docker load -i 镜像保存文件位置  例如： docker load -i redis.tar

**删除镜像**：docker rmi -f 镜像名/镜像ID ，比如：docker rmi -f redis:latest

**根据名称删除指定的容器**：

```mysql
cat > stop.sh << EOF
#!/bin/bash
# 查找所有镜像名称中包含 "goharbor" 的容器
containers=\$(docker ps -a --format="{{.ID}} {{.Image}}" | grep "goharbor" | awk '{print \$1}')

# 遍历找到的容器，将它们停止
for container in \$containers
do
  echo "Stopping container \$container"
  docker stop \$container
done
EOF

#执行
sudo ./stop.sh


```

Docker镜像下架地址

**help类库的使用**:![1678151371822](Vue进步.assets/1678151371822.png)



**帮助按钮**:docker --help



### Linux常用命令

**更换 yum镜像源: **1.curl -o /etc/yum.repos.d/CentOS-Base.repo https://mirrors.aliyun.com/repo/Centos-7.repo 2.curl -o /etc/yum.repos.d/epel.repo http://mirrors.aliyun.com/repo/epel-7.repo 3.yum makecache 4.yum makecacheyum repolist  5.yum -y update   完成5后就可以正常下载你要下载的软件了

**查看内核版本**:uname -r ,输出：3.10.0-514.el7.x86_64

**只查看文件的内容(一般docker内部都默认安装了)**:cat /etc/nginx/nginx.conf

**检查防火墙状态** ：有两个防火墙，sudo firewall-cmd --state ，还有一个是sudo iptables -L

**显示操作系统的信息**:cat /etc/os-release

**查看网卡信息**： ifconfig

**查看网卡,网关名称**:netstat -rn , 看第一排Iface的属性名就是网关名，这里是：Iface：ens33， Gateway：192.168.226.2  ，也可以通过：ip link show chaka

**查找文件是否存在占用**：fuser keepalived/keepalived.conf，输出为空代表不存在占用

**查看文件夹与文件列表** :ls 查看文件与文件夹信息，ls -a 查看文件与文件夹列表(包括隐藏文件),ls -l 查看文件夹的详细信息,查看指定文件的详细生成信息：ls -l /root/rabbitmq-node-data/.erlang.cookie，查看文件的内容信息：cat /root/rabbitmq-node-data/.erlang.cookie，查看内容信息，并且可以编辑：vi /root/rabbitmq-node-data/.erlang.cookie

**查询系统时间**：date

**查看是否有程序在监听5000,5001端口**: ss -an | grep LISTEN | grep -E "(:5000|:5001)\s"  ,如果输出出来有结果，那么代表有程序监听，输出结果为空代表没有程序在监听这个两个端口，没有程序监听，自然也就无法访问

**创建文件，并且编辑**:vi Dockerfile

**新建文件夹**:mkdir RedisDemo ,也可以递归创建文件夹： mkdir RedisDemo/Data (这样会创建RedisDemo和Data两个文件夹)

**只查看开机自启动的服务**: systemctl list-unit-files --state=enabled

**查看指定关键字的服务**:systemctl list-units | grep elk , elk我这里是启动用户的名称，也可以是服务的名称

**查看指定服务的状态**: systemctl status kibana

**查看系统启动的所有服务**：systemctl

**判断服务是否开机自启动**:systemctl is-enabled kibana,输出enabled代表开机自启动，输出disabled代表不会开机自启动

**执行的进程服务根据指定关键字查找**：ps auxw | grep elk  ，elk 是用户名称，也可以是kibana应用程序名称

**查看系统占用CPU,内存**: top

**查看内存使用情况**: free -h,   total:总内存,used：占用内容,free空闲内存

**判断IP与端口是不是通**: tcping 127.0.0.1 9200   /  telnet 127.0.0.1 9200

**curl带用户名密码请求**:curl -u elastic:VZTbGu6nQ7nR7E*hlHh- https://14.29.220.179:9200 -v -k

**curl发起http get请求**： curl http://localhost:8070/weatherforecast  ，会拿到get请求的报文
**解压压缩包**:unzip elasticsearch.zip
**删除压缩包**：**rm -rf** elasticsearch.zip ，也可以用来指定删除文件:rm -rf dockerfile (-rf代表的是不进行确认，直接删除)

**复制压缩包到指定路径**： cp /root/elasticsearch.zip es:/usr/share/elasticsearch/plugins

**移动文件到指定位置** ： mv docker.zip(文件名)  /root/dockertext/(目标位置)， 另一个目录

**移动文件夹到指定位置：** ：

**到达指定目录** ： cd /root/   ， 一般拷贝文件到了linux系统都是在这个目录下面的，所以需要操作一下

**新增文件夹**：mkdir dockertext

**返回目录**：**cd** (返回到上一步的目录) ,**cd ..** （返回到上一级的目录）,**cd~** （到用户所在的家目录）

**查看目录信息**:   **ls**   (后面的命令可以加上的条件)，-d：只看当前目录的信息，-l：看详细信息，-a：显示所有 任何一个文件前面加上“.”，表示隐藏文件。-h：显示文件大小

**查看当前目录路径**: pwd

**查看IP**: ifconfig  (看ens33(inet),也不一定完全是ens33:inet属性，也可能是virbr0:inet属性，或者其他的属性，反正找不到上面的两个就找个192.168开头的怼，比如192.168.226.128)

**查看已经启动的服务**：systemctl list-units --type=service    ，docker服务的名字叫docker.service

**查看是否已经设置开机启动(根据docker服务名称查询)**: systemctl list-unit-files | grep docker   （enable是设置好了开机启动，dis开头的是没有设置开机启动）

**设置服务开机启动**：systemctl enable docker.service

**冒号:q退出当前窗口**：q 

**查看端口使用情况**：netstat -nultp （看Local Address这列后面的端口号）

**查看指定端口是否被占用/解除占用**： **netstat -anp |grep 82** ，如果存在占用就会输出一行ip信息的数据，如果没有占用就不会输出一行IP信息的命令，如果占用可以执行：**kill -9 进程的pid**

**查看文件内容** : cat 文件名 ， 比如cat  kibana.yml , 也可以是带路径名称，比如cat /opt/kibana/config/kibana.yml

**【使用管理员权限】** ：su ,输入su 然后输入用户密码

**vi编辑器保存文件（错误）**: **:w！** (强制存盘) , 然后输入 **:qa!** --- 解决这个错误：    No write since last change (add ! to override) ，自上次更改(添加!覆盖)

**vi编辑器编辑文件**: vi /opt/kibana/config/kibana.yml  ，输入 I 进入到Insert编辑模式
输入**esc** 退出编辑,输入:**wq 保存当前文件**后退出,输入**:q 直接退出**，不保存，**:qa!** 强制退出不保存

![1670903703106](Vue进步.assets/1670903703106.png)





### Ubuntu常用命令

**sudo 名称无效**：apt-get update，apt-get install sudo

**安装vim编辑器**: 1.apt-get update 更新一波基础配置 2.apt-get install vim 下载vim编辑器

**查看系统版本**:lsb_release -a

**安装ssr脚本包没有安装**：E: Package 'python' has no installation candidate，E:包“python”没有安装候选

```c#
# apt-get update
# apt-get upgrade
# apt-get install python-is-python3
```



### DockerIgnore

忽略指定后缀，哪些开头的文件

### 安装Docker

```mysql
# 命令1: 安装docker
yum install -y yum-utils \
           device-mapper-persistent-data \
           lvm2 --skip-broken

# 命令2：更新本地镜像源
# 设置docker镜像源
yum-config-manager \
    --add-repo \
    https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
    
#单独执行    
sed -i 's/download.docker.com/mirrors.aliyun.com\/docker-ce/g' /etc/yum.repos.d/docker-ce.repo

# 配置镜像缓存
yum makecache fast

# 命令3：【安装docker-ce】 社区免费版 docker，这一步时间比较长要下载100多MB，看网速，我在公司10-20分钟左右
yum install -y docker-ce

# 命令4：关闭防火墙：Docker应用需要用到各种端口，逐一去修改防火墙设置。非常麻烦，因此建议直接关闭防火墙
	-- 关闭
	systemctl stop firewalld
	-- 禁止开机启动防火墙
	systemctl disable firewalld
	-- 查看是否关闭防火墙
	systemctl status firewalld
# 命令5：启动/关闭Docker
    systemctl start docker  # 启动docker服务

    systemctl stop docker  # 停止docker服务

    systemctl restart docker  # 重启docker服务
# 查看Docker版本
  docker -v
# 配置阿里云镜像加速，docker官方镜像仓库网速较差，我们需要设置国内镜像服务：参考阿里云的镜像加速文档：https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors
  -- 创建文件夹
sudo mkdir -p /etc/docker
> --  在文件夹内新建一个daemon.json文件（下面的命令直接在cmd执行）
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://g117omq3.mirror.aliyuncs.com"]
}
EOF

	-- 重新加载文件
	sudo systemctl daemon-reload
	-- 重启docker
	sudo systemctl restart docker
# 更改镜像下载地址
  -- 查看原来的镜像地址
  cat /etc/docker/daemon.json
  -- 更改镜像地址，重复上面的命令就可以（只是把registry-mirrors：这里面的值更改一下）
  sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://g117omq3.mirror.aliyuncs.com"]
}
EOF

# 监测镜像源地址是否有效：
curl -I https://g117omq3.mirror.aliyuncs.com
返回HTTP/1.1 200 OK就是有效，返回其他就是有问题
# 排查Docker的一些问题
1.重启docker    sudo systemctl restart docker
2.查看docker状态 systemctl status docker.service
3.查看系统日志，获取更多的信息 journalctl -xe
4.重新刷新配置文件 sudo systemctl daemon-reload
5.设置多个镜像源地址：sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://g117omq3.mirror.aliyuncs.com","https://dockerproxy.com"]
}
EOF
6.IP 104.18.123.25:443是https://hub.docker.com/的IP地址，dial tcp 104.18.122.25:443: i/o timeout 超时问题解决思路：
  1.运行sudo yum install -y docker-ce docker-ce-cli containerd.io 安装缺失依赖 
  2.设置dns为公共的dns：8.8.8.8 
  3.设置vpn代理 
  4.ipconfig /flushdns 更新dns 
  原文地址：https://stackoverflow.com/questions/48042184/how-to-solve-i-o-timeout-error-in-docker-pull

7.上面设置设置VPN代理
 7.1.设置linux中docker vpn代理的方法：
sudo mkdir -p /etc/systemd/system/docker.service.d/
echo '[Service]
Environment="HTTP_PROXY=http://127.0.0.1:7890"
Environment="HTTPS_PROXY=http://127.0.0.1:7890"' | sudo tee /etc/systemd/system/docker.service.d/http-proxy.conf
 ，设置后刷新配置：sudo systemctl daemon-reload ，重启docker服务 sudo systemctl restart docker
 7.2 设置linux的代理（本次会话）
 export HTTP_PROXY=http://192.168.226.222:7890 ,设置unset HTTP_PROXY可以取消代理
 export HTTPS_PROXY=http://192.168.226.222:7890，设置unset HTTPS_PROXY可以取消代理
 7.3 设置windows的代理（本次会话）
 setx HTTP_PROXY "http://127.0.0.1:7890"
 setx HTTPS_PROXY "http://127.0.0.1:7890"
 7.4 怎么判断代理是可用的？
     1.telnet 127.0.0.1 7890 ,windows和linux都可以用这个，可以访问的话看2 
     2.curl -x 127.0.0.1:7890 http://www.google.com 可以访问就说明代理是通的
     3.我发现与本机与虚拟机通讯的IP,cmd中用ipconfig查到VMnet8网卡中本机与虚拟机网段相同的ip端：192.168.226.222，然后使用curl -x 127.0.0.1:7890 http://www.google.com 正常联通，参考链接：https://www.zhihu.com/question/495148700/answer/2496921147

```

[镜像拉取/docke安装/常用命令原文链接](https://blog.csdn.net/m0_59196543/article/details/124749175)



### **界面与CMD模式切换**

**右键**

![1670567296660](Vue进步.assets/1670567296660.png)

### **Linux远程工具**

- MobaXterm远程服务器 -- 这个是yyds

  打开多页面分割

  ![1681284758994](Vue进步.assets/1681284758994.png)

打开多终端同时输入

![1681285464132](Vue进步.assets/1681285464132.png)

退出多终端同时输入

![1681285417344](Vue进步.assets/1681285417344.png)

-  SecureCRT 远程服务器
-  SSH
- Winscp
- finalshell:国产最好用的ssh工具

### Linux文件传输工具

**在官网下载的时候记得在所有下载那里选择个人版本与家庭版本的免费版本，就没有一个月的失效时间**

**XSHELL**：连接到Linux，安装过程下一步下一步，连接过程看下图

![1670896365122](Vue进步.assets/1670896365122.png)

**XFTP**：传输文件到Linux，安装过程下一步下一步，在xshell中使用看下图，也可以直接在XFTP中输入IP地址，用户名密码进行连接

![1670896662862](Vue进步.assets/1670896662862.png)



[原文教程](https://blog.csdn.net/m0_63004677/article/details/126503333)

### Linux知识

**CentOS和Ubuntu的比较，以及各自适用的场景：**

| 特性         | CentOS（适合服务器正式用）                           | Ubuntu（适合个人电脑，开发机）                           |
| ------------ | ---------------------------------------------------- | -------------------------------------------------------- |
| 发行版本     | RHEL的社区克隆版本                                   | Debian的社区克隆版本                                     |
| 发行周期     | 稳定性优先，长周期（一般5年或更久）                  | 新特性优先，短周期（一般6个月）                          |
| 软件包管理   | yum / rpm，更适合安装较稳定的软件版本                | apt / dpkg，更适合安装新特性的软件版本                   |
| 软件源       | 官方源和第三方源都有，但第三方源不如Ubuntu丰富       | 官方源和第三方源都有，且第三方源较为丰富                 |
| 软件包数量   | 较少，但都比较稳定                                   | 较多，但部分较新的软件包可能不够稳定                     |
| 默认桌面环境 | 没有默认桌面环境，适合作为服务器操作系统             | GNOME，适合作为桌面操作系统                              |
| 内核版本     | 内核版本较旧，但已经过测试，较为稳定                 | 内核版本较新，包含一些新特性，但稳定性可能稍逊一些       |
| 适用场景     | 服务器、稳定性要求高的环境，如企业服务器、云服务器等 | 桌面、开发环境、新特性需求高的环境，如个人电脑、开发机等 |

对于企业服务器、云服务器等稳定性要求较高的环境，CentOS是更为合适的选择，因为其发行周期较长，软件包较稳定，适合长时间运行的稳定业务。此外，由于CentOS的默认桌面环境为无界面状态，更适合服务器操作系统的需求。

对于个人电脑、开发机等新特性需求较高的环境，Ubuntu是更为合适的选择，因为其发行周期较短，软件包更新较快，适合安装较新的软件包和尝试新特性。此外，Ubuntu的默认桌面环境为GNOME，更适合桌面操作系统的需求。

总之，在选择操作系统时，需要根据实际的需求和场景来选择合适的操作系统，不能仅仅看重某一个特性而忽略其他的因素。

### **Linux安装教程**

[Linux安装教程](https://blog.csdn.net/weixin_52799373/article/details/124324077)

[失效链接，查看存档内容](https://web.archive.org/web/20210401000000*/https://www.cnblogs.com/xiaobaozi-95/p/9328948.html)

[维基百科关于失效链接的说明](https://zh.wikipedia.org/zh-hans/Wikipedia:%E5%A4%B1%E6%95%88%E9%93%BE%E6%8E%A5)

[Linux安装教程2，上面的失效了，看这个](https://blog.csdn.net/m0_37479246/article/details/84944002)

### Linux部署Net Core

#### Docker 容器化部署

##### 新建Docker解决方案

```mysql
选择项目的时候可以选择webapi也可以选择控制台程序，重点是需要勾选 启用 Docker支持，操作系统选择Linux，选择好后直接确定新建项目，进入下一步在Linux系统安装Docker Compose
 
```

[参考内容](http://www.qb5200.com/article/476577.html)

[参考内容2：这个更多学习的内容](https://jw123.com.cn/tech/shownews.php?id=11762)

#####  Linux安装Docker Compose

```mysql
#执行命令
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

#执行命令授权
sudo chmod +x /usr/local/bin/docker-compose

# 检查是否安装成功
docker-compose --version

# 新建docker-compose.yml文件
version: '3.3'
services:
  net7web:  
    build:
      context: ./WebApplication5/  -- 需要更改项目名称：WebApplication5（这是文件根目录整个解决方案的首页面，双击进去才是.sln文件的位置）
      dockerfile: ./WebApplication5/Dockerfile -- 需要更改项目名称：WebApplication5
    image: net7web:2022  -- 镜像名
    container_name: "net5webContainer"    --  容器名称
    ports:
      - 8070:80/tcp  -- 可以不需要更改，通过8070端口访问部署的webapi,将容器的8070端口映射到外部的80端口
    command: ["dotnet", "/app/WebApplication5","--urls=http://*:80","--ip=127.0.0.1"]
    --  需要更改项目名称：WebApplication5
```

##### Linux执行Docker Compose

```mysql
# 拷贝整个解决方案文件到linux服务器，
  pwd, cd 到 解决方案文件路径，单独赋值docker compose文件到解决方案路径一起
# 执行Docker Compose文件
docker-compose up
# 访问以下链接 ：虚拟机ip+映射端口/方法名称  下面两个链接都可以在虚拟机访问
curl http://192.168.226.128:8070/weatherforecast  
curl http://localhost:8070/weatherforecast  -- curl命令在cmd窗口发起http get请求

小插曲：火狐浏览器访问链接提示，您尝试查看的页面无法显示，因为已收到数据的可靠性无法证实，把访问链接的https更改为http就可以 
```

##### 【发布部署，暂不使用，服务部署启动后，访问不了】发布项目到文件夹部署到Linux成功

```mysql
# 1.选项项目的时候可以选择webapi也可以选择控制台程序，重点是需要勾选 启用 Docker支持，操作系统选择Linux，不勾选swagger支持
  -- 看上面的注意事项
# 2.发布好文件夹后，写Dockerfile
FROM mcr.microsoft.com/dotnet/core/aspnet:3.1-buster-slim AS base
FROM mcr.microsoft.com/dotnet/core/sdk:3.1-buster -- -buster 这个需要注意，系统自动生成的是build
WORKDIR /app
EXPOSE 5000

COPY . .

ENTRYPOINT ["dotnet", "MyDockerApi.dll"]
# 3.使用Xftp软件拷贝发布的项目文件和Dockerfile文件到Linux服务器

# 4.使用Dockerfile生成镜像文件
docker build -t webapplication3 .    -- 注意后面这个点，这个点是精华
# 5.生成容器
docker run --name=aspnetcoredocker6 -p 7070:80 -d webapplication3
# 6.请求验证
curl http://localhost:8081/weatherforecast/MyDocker

/***************分割线************************/

# 1.细节1，dockerfile文件

FROM mcr.microsoft.com/dotnet/core/sdk:3.1-buster -- -buster 这个需要注意，系统自动生成的是build

# 2.细节2 ：访问链接 ，注意是Http不是Https
curl http://localhost:8081/weatherforecast/MyDocker
有错误提示，访问提示：载入页面时到 localhost:8081 的连接被中断 待解决
```

[Linux系统docker部署.net core3.1](https://www.cnblogs.com/saltlight-wangchao/archive/2022/09/01/16646005.html)

#### Dotnet 直接部署

##### **1.Linux上面安装Net SDK**

```mysql 
# 1.先下载脚本，通过https://dot.net/v1/dotnet-install.sh 进行下载脚本
 -- 再通过Xftp将dotnet-install.sh脚本上传到linux
# 2.切换到对应的脚本路径，安装对应的Net Sdk
   -- 1. pwd   看看初始路径
   -- 2. cd /root/   跳到dotnet-install.sh保存的路径
   -- 3. ls   看看dotnet-install.sh脚本是否存在
   -- 4. sudo chmod +x dotnet-install.sh   获取脚本管理员权限
   -- 5. ./dotnet-install.sh --channel 5.0    下载5.0的sdk，也可以是7.0/6.0等版本，根据需求来（这一步安装可能会提示网络问题，可以联系网管看一下，如果安装不了，可以手动下载安装包放到linux里面通过命令行进行安装）
   -- 6. dotnet --list-sdks   看一下安装是否成功
      -- 6.1 提示"bash: dotnet: 未找到命令..."错误的解决：配置环境变量，因为.sh脚本不会自动将dotnet添加到环境变量中，需要我们手动设置环境变量
      -- 6.2  echo $PATH  检查系统目前的环境变量，路径用:隔开
      -- 6.3  export PATH="$PATH:$HOME/.dotnet" 配置系统的环境变量
      -- 6.4  echo $PATH  检查是否安装成功,看最末尾那里是不是有dotnet
      -- 6.5  dotnet --list-sdks   看一下安装是否成功
# 3.手动安装sdk 镜像环境
docker pull mcr.microsoft.com/dotnet/aspnet:3.1   -- pull后面的路径可以在系统自动生成的dockerfile里面拿

 如果不docker部署需要服务自己启动需要【linux守护程序部署】，需要的技术是worker service
```

[安装Linux的NetSdk环境：看通过脚本如何安装内容](https://learn.microsoft.com/zh-cn/dotnet/core/install/linux-scripted-manual#scripted-install)

##### 2.Dotnet Dll部署服务

```mysql
# 1.将发布后的文件拷贝到linux目录下,注意不要勾选openapi支持，不知道为啥访问不了（我不知道访问方法）
# 2.dotnet dllname  直接执行
# 3.测试访问：http://localhost:5001/weatherforecast  
```



#### Docker容器化部署定时任务

```mysql
#1.拷贝解决方案文件到linxu服务器
 cd /root
# 2.解压缩
unzip Sunway.TaskCoravel.Web.zip
# 3.删除压缩文件
rm -rf Sunway.TaskCoravel.Web.zip
# 4.使用Dockerfile生成镜像文件
docker build -t sunway.taskcoravel.web .    -- 注意后面这个点，这个点是精华，并且项目的名称需要是小写
# 5.发现没有自动拉取dotnet镜像，手动拉取
docker pull mcr.microsoft.com/dotnet/core/aspnet:3.1-buster-slim
# 5.生成容器,80端口默认有程序使用，最好不用80端口,可以用7070,8010等等
docker run -e TZ=Asia/Shanghai -d -p 7070:7070 --restart=always --name sunway.taskcoravel.service  -v /usr/local/webapp/Sunway.TaskCoravel.Web/:/app -v /etc/localtime:/etc/localtime sunway.taskcoravel.web
# 6.请求验证
curl http://localhost:44323/weatherforecast

备用脚本，部署时出现问题，用来排除异常
1.dotnet Sunway.TaskCoravel.Web.dll --urls "http://0.0.0.0:7070;"  --直接执行
```

#### 通过发布执行Docker

发布

![1677578164635](Vue进步.assets/1677578164635.png)

进入到项目文件里面执行

![1677577894671](Vue进步.assets/1677577894671.png)



### Docker部署Redis

```mysql
# 拉取镜像
docker pull redis
# 下载配置文件，选择：Download 6.0.6 这个下载，然后解压拿到根目录下面：redis.conf配置文件
http://www.redis.cn/download.html

# 更改配置，打开redis.config
 bind 127.0.0.1 #注释掉这部分，使redis可以外部访问
 requirepass wr123 #配置密码（不配置外部有权限限制用不了）
# 拷贝redis.conf配置文件到linux下面redis容器中
 一般默认我是放在 /usr/local/webapp/Redis路径 下面
#启动容器
docker run -p 6379:6379 --name redis -v /usr/local/webapp/Redis/redis.conf:/etc/redis/redis.conf  -v /usr/local/webapp/Redis:/data -d redis redis-server /etc/redis/redis.conf --appendonly yes
#检查日志
docker logs redis
#本地linux 安装redis客户端
sudo yum install epel-release  -- 安装软件仓库（就像苹果的app store）
sudo yum update  -- 更新仓库中的软件包
sudo yum install redis  -- 安装Redis客户端
#内部linux调用，外部windows调用都是用这个命令
redis-cli -h 192.168.226.128  -p 6379 -a wr123
```

[Docker部署Redis原文链接](https://cloud.tencent.com/developer/article/1670205)

### Linux部署Net Core守护进程

- Docker部署：

  **优点**：1.异常自动重启 2.随着Linux服务启动  3.日志方便查看 4.通过k8s进行可视化管理 5.可以动态伸缩 6.隔离性：容器之间互相隔离 7.跨平台性：windows,linux，macos都可以 8.可移植性：可以打包镜像快速移植共享

  **缺点：** 1.学习曲线复杂 2.网络配置复杂 3.运维成本高：可能要有专门的运维团队维护

- 守护进程部署：

  **优点**：1.异常自动重启 2.随着Linux服务启动  3.日志方便查看  

  **缺点**：1.没有可视化管理界面 2.不方便动态伸缩 3.不方便移植 4.配置守护进程相对于复杂

```mysql
# 新建解决方案
新建linux web api,net 6
# 注释只在开发环境使用swagger的配置
app.Environment.IsDevelopment()
# 发布到文件夹
发布到本地的文件夹
# 将文件夹拷贝到Linux服务器，通过MobaXterm客户端
一般默认我是放在 /usr/local/webapp/RedisApi路径 下面
# dotrun 测试访问情况，配置0.0.0.0 是为了外部可以通过虚拟机ip访问，默认是localhost/127.0.0.1只能在虚拟机内部进行访问
dotnet RedisDemo.dll --urls "http://0.0.0.0:5000;https://0.0.0.0:5001"
# linux内部测试 ,-k 忽略https证书验证,-v 查看详细的信息（如果有错误的话会很详细）
curl -k -v https://192.168.226.128:5001/swagger/index.html

# windows外部测试，直接访问
https://192.168.226.128:5001/swagger/index.html

# 部署守护进程，自动启动
# 以编辑模式创建redisapi.service服务文件（没有任何后缀就是一个文件）
sudo vi /etc/systemd/system/RedisDemo.service


#里面的内容：
#【注意】1.配置正确dotnet命令位置： /root/.dotnet/dotnet 是dotnet安装的位置，通过which dotnet命令获得
#2.urls符号没有双引号：正确：--urls=http://0.0.0.0:5000;https://0.0.0.0:5001 ，没有双引号，外部执行是这样 ，--urls="http://0.0.0.0:5000;https://0.0.0.0:5001",外部执行是对的，这里守护进程里面是错误的
[Unit]
Description=RedisDemo

[Service]
WorkingDirectory=/usr/local/webapp/RedisDemo
ExecStart=/root/.dotnet/dotnet /usr/local/webapp/RedisDemo/RedisDemo.dll --urls=http://0.0.0.0:5000;https://0.0.0.0:5001
Restart=always
RestartSec=10
SyslogIdentifier=RedisDemo
User=root
Environment=ASPNETCORE_ENVIRONMENT=Production

[Install]
WantedBy=multi-user.target

#刷新配置
sudo systemctl daemon-reload
#启动服务
sudo systemctl start RedisDemo
#检查状态
sudo systemctl status RedisDemo
#设置开机自启动
sudo systemctl enable RedisDemo
#停止服务
sudo systemctl stop RedisDemo
```



### Linux部署Net Core容器

```mysql
# 新建Linux WebApi ，net 6.0 ,新建Docker支持
项目名：RedisDemo，记得注释测试环境也可以看swagger
app.Environment.IsDevelopment()

# 发布项目到文件夹，通过MobaXterm拷贝文件到Linux

# cd到项目的目录下面，新建Dockerfile文件，把下面的内容copy进来
vi Dockerfile

# Dockerfile内容，参考系统生成的进行改写
FROM mcr.microsoft.com/dotnet/aspnet:6.0
WORKDIR /app
EXPOSE 7071
COPY . /app
ENTRYPOINT ["dotnet", "RedisDemo.dll"]
# 生成镜像
docker build -t redisdemo.
# 执行
# 【注意】
# 1. -p 是用来容器的端口与外部环境的端口的，左边是宿主机的端口，右边是容器的端口，并且一次可以映射多个端口出去，比如我映射了7071,7072端口
# 2. --urls 很重要是用来监听应用程序的端口的，我这里urls配置的7071，7072刚好与-p容器与外部宿主机的端口7071，7072相呼应
# 3. 如果提示确实SDK环境，确认一下-v /usr/local/webapp/RedisDemo/的映射路径是不是到了dll级别，没有到这个级别的映射就会有问题
docker run -d -p 7071:7071 -p 7072:7072 --restart=always --name redisApi -v /usr/local/webapp/RedisDemo/:/app redisdemo --urls="http://*:7071;http://*:7072"

```

### Docker部署RabbitMQ

```mysql
# 拉取镜像 ,我发现docker pull rabbitmq:latest ，这个版本没有后台管理的插件，不知道是镜像源有问题，还是就是这样子
docker pull rabbitmq:3.7.7-management
# 启动服务 ，RABBITMQ_DEFAULT_VHOST=my_vhost 注意这个虚拟主机的名称，使用的时候配置完ip，端口号还需要配置这个
docker run -d --name rabbitmq3.7.7 -p 5672:5672 -p 15672:15672  --restart=always -v `pwd`/data:/var/lib/rabbitmq --hostname myRabbit -e RABBITMQ_DEFAULT_VHOST=my_vhost  -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=123 2888deb59dfc
# 访问
http://192.168.226.128:15672/
```

### Docker部署RabbitMQ普通集群(单服务器)

这种集群如果docker stop  rabbitmq-node1,那么剩余的集群全部不能访问了，直接挂掉，虽然可以使用代码轮训的方式访问，但是没有高可用

```mysql
# 拉取镜像 ,我发现docker pull rabbitmq:latest ，这个版本没有后台管理的插件，不知道是镜像源有问题，还是就是这样子
docker pull rabbitmq:3.7.7-management

# 配置网络，-link的连接方式docker已经弃用了
docker network create rabbitmq_network
# 创建.erlang.cookie文件
echo "mysecretcookie" > rabbitmq-node1data/.erlang.cookie
echo "mysecretcookie" > rabbitmq-node2data/.erlang.cookie
echo "mysecretcookie" > rabbitmq-node3data/.erlang.cookie


# 启动服务 ，RABBITMQ_DEFAULT_VHOST=my_vhost 注意这个虚拟主机的名称，使用的时候配置完ip，端口号还需要配置这个
docker run -d --hostname rabbitmq-node1 --name rabbitmq-node1 -p 5672:5672 -p 15672:15672 --restart=always -v `pwd`/rabbitmq-node1data:/var/lib/rabbitmq -v `pwd`/rabbitmq-node1data/.erlang.cookie:/var/lib/rabbitmq/.erlang.cookie -e RABBITMQ_DEFAULT_VHOST=my_vhost -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=123 --network rabbitmq_network rabbitmq:3.7.7-management

docker run -d --hostname rabbitmq-node2 --name rabbitmq-node2 -p 5673:5672 -p 15673:15672 --restart=always -v `pwd`/rabbitmq-node2data:/var/lib/rabbitmq -v `pwd`/rabbitmq-node2data/.erlang.cookie:/var/lib/rabbitmq/.erlang.cookie -e RABBITMQ_DEFAULT_VHOST=my_vhost -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=123 --network rabbitmq_network rabbitmq:3.7.7-management

docker run -d --hostname rabbitmq-node3 --name rabbitmq-node3 -p 5674:5672 -p 15674:15672 --restart=always -v `pwd`/rabbitmq-node3data:/var/lib/rabbitmq -v `pwd`/rabbitmq-node3data/.erlang.cookie:/var/lib/rabbitmq/.erlang.cookie -e RABBITMQ_DEFAULT_VHOST=my_vhost -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=123 --network rabbitmq_network rabbitmq:3.7.7-management
# 将节点加入到集群 ,注意，下面的命令要一行一行的执行
docker exec -it rabbitmq-node2 rabbitmqctl stop_app
docker exec -it rabbitmq-node2 rabbitmqctl join_cluster rabbit@rabbitmq-node1
-- docker exec rabbitmq2 rabbitmqctl join_cluster --ram rabbit@rabbitmq1


docker exec -it rabbitmq-node2 rabbitmqctl start_app

docker exec -it rabbitmq-node3 rabbitmqctl stop_app
docker exec -it rabbitmq-node3 rabbitmq


# 访问，检查集群Overview-Nodes 列表
http://192.168.226.128:15672/
http://192.168.226.128:15673/
http://192.168.226.128:15674/

 
```

### Docker部署RabbitMQ普通集群(多服务器)

```mysql
一、每台服务器上面执行
# 拉取镜像 
docker pull rabbitmq:3.7.7-management
# 创建网络
docker network create --driver bridge rabbitmq_cluster_network
# 创建.erlang.cookie文件
mkdir rabbitmq-node-data
echo "mysecretcookie" > rabbitmq-node-data/.erlang.cookie
#打开容器内部操作权限
chmod 600 `pwd`/rabbitmq-node-data/.erlang.cookie

#更新host文件
sudo vi /etc/hosts
#host文件内容
192.168.226.128 rabbitmq-node1.local
192.168.226.129 rabbitmq-node2.local
192.168.226.130 rabbitmq-node3.local



二、单独服务器执行


#服务器1执行
docker run -d --hostname rabbitmq-node1.local --add-host=rabbitmq2:192.168.226.129  --add-host=rabbitmq3:192.168.226.130  --name rabbitmq-node1 -p 5672:5672 -p 15672:15672 -p 4369:4369  -p 25672:25672 --restart=always -v `pwd`/rabbitmq-node-data:/var/lib/rabbitmq -v `pwd`/rabbitmq-node-data/.erlang.cookie:/var/lib/rabbitmq/.erlang.cookie  -v /etc/hosts:/tmp/hosts -v `pwd`/rabbitmq-node-data/logs:/var/lib/rabbitmq/logs -e RABBITMQ_DEFAULT_VHOST=my_vhost -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=123 -e RABBITMQ_NODENAME=rabbit@rabbitmq-node1.local -e RABBITMQ_USE_LONGNAME=true --network rabbitmq_cluster_network rabbitmq:3.7.7-management

#服务器2执行 -v /etc/hosts:/tmp/hosts，这个挂载配置是核心
docker run -d --hostname rabbitmq-node2.local --name rabbitmq-node2 --add-host=rabbitmq1:192.168.226.128 --add-host=rabbitmq3:192.168.226.130 -p  5672:5672 -p 15672:15672 -p 4369:4369  -p 25672:25672 --restart=always -v `pwd`/rabbitmq-node-data:/var/lib/rabbitmq -v `pwd`/rabbitmq-node-data/.erlang.cookie:/var/lib/rabbitmq/.erlang.cookie  -v /etc/hosts:/tmp/hosts -v `pwd`/rabbitmq-node-data/logs:/var/lib/rabbitmq/logs -e RABBITMQ_DEFAULT_VHOST=my_vhost -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=123 -e RABBITMQ_NODENAME=rabbit@rabbitmq-node2.local -e RABBITMQ_USE_LONGNAME=true --network rabbitmq_cluster_network rabbitmq:3.7.7-management

#服务器3执行 
docker run -d --hostname rabbitmq-node3.local --name rabbitmq-node3 --add-host=rabbitmq1:192.168.226.128 --add-host=rabbitmq2:192.168.226.129 -p  5672:5672 -p 15672:15672 -p 4369:4369  -p 25672:25672 --restart=always -v `pwd`/rabbitmq-node-data:/var/lib/rabbitmq -v `pwd`/rabbitmq-node-data/.erlang.cookie:/var/lib/rabbitmq/.erlang.cookie  -v /etc/hosts:/tmp/hosts -v `pwd`/rabbitmq-node-data/logs:/var/lib/rabbitmq/logs -e RABBITMQ_DEFAULT_VHOST=my_vhost -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=123 -e RABBITMQ_NODENAME=rabbit@rabbitmq-node3.local -e RABBITMQ_USE_LONGNAME=true --network rabbitmq_cluster_network rabbitmq:3.7.7-management

#进入容器rabbitmq-node2拷贝配置
docker exec -it rabbitmq-node2 bash -c "cp /tmp/hosts /etc/hosts; exit"

#进入容器rabbitmq-node3拷贝配置
docker exec -it rabbitmq-node3 bash -c "cp /tmp/hosts /etc/hosts; exit"


# 服务器2执行,一行一行的执行
docker exec -it rabbitmq-node2 rabbitmqctl stop_app
docker exec -it rabbitmq-node2 rabbitmqctl join_cluster --longnames rabbit@rabbitmq-node1.local
docker exec -it rabbitmq-node2 rabbitmqctl start_app

#服务器3执行,一行一行的执行
docker exec -it rabbitmq-node3 rabbitmqctl stop_app
docker exec -it rabbitmq-node3 rabbitmqctl join_cluster --longnames rabbit@rabbitmq-node1.local
docker exec -it rabbitmq-node3 rabbitmqctl start_app
 
# 更新为镜像队列
#进入容器
docker exec -it rabbitmq-node1 bash
#设置镜像队列，在配置中设置，下面的设置是设置目前的队列为镜像队列，新增的队列也都是镜像队列,-p my_vhost是虚拟主机的名称,如果设置默认的虚拟主机的话，用一个左斜杠/
rabbitmqctl set_policy ha-all "" '{"ha-mode":"all","ha-sync-mode":"automatic"}' --priority 1 -p my_vhost

注意,备用脚本：
1.检查容器中erlang.cookie的值：
docker exec rabbitmq-node1 cat /var/lib/rabbitmq/.erlang.cookie
2.查看指定节点中的集群状态
docker exec rabbitmq-node1 rabbitmqctl cluster_status
3.检查容器中hosts的值
docker exec rabbitmq-node1 cat /etc/hosts
```

### Docker部署RabbitMQ普通集群(多服务器2：查看镜像队列即将写入的数据)

```mysql
这个教程是在部署好集群，设置好镜像队列的基础上确定镜像队列，集群的数据同步，是不是持久化到了磁盘上面

#进入容器
docker exec -it rabbitmq-node1 bash
#设置镜像队列，也可以在代码中设置，也可以在配置中设置，下面的设置是设置目前的队列为镜像队列，新增的队列还不是镜像队列，要新增的队列是镜像队列，需要新增rabbitmq的配置
rabbitmqctl set_policy ha-all "" '{"ha-mode":"all","ha-sync-mode":"automatic"}'
# 安装RabiitMQ查看数据插件bsdmainutils
apt-get update
apt-get install -y bsdmainutils
# 进入到容器中的rabbitmq服务
su rabbitmq
#cd到数据目录:
#1.rabbit@rabbitmq-node1.local是-e RABBITMQ_NODENAME容器的名称 
#2.9CXUC1GMRBKT0XK28FYKNS76M 是数据磁盘的名称，可能不同服务器不同，我看集群上面这个标志在三台服务器上面都是一样的
#3.F02V7J2ATNNJQ2GJ7ZNEMP0NG 队列的信息，一个队列有很多这种字符的文件夹，一个队列的数据保存在磁盘的多个文件夹下面
/var/lib/rabbitmq/mnesia/rabbit@rabbitmq-node1.local/msg_stores/vhosts/9CXUC1GMRBKT0XK28FYKNS76M/queues/F02V7J2ATNNJQ2GJ7ZNEMP0NG
#查看数据文件，journal.jif里面保存的是10进制的数据，下面的命令可以看到左边是十进制，右边是队列中保存的数据内容
hexdump -C journal.jif
```



### Docker部署RabiitMq+Haproxy+Keepalived高可用集群

Keepalived：生成虚拟的静态IP,可以通过它访问集群的信息，并且如果服务器IP变动了，程序的配置也不需要更改，还有心跳检查功能，高可用

Haproxy：高并发，高可用的反向代理服务器，有负载均衡，心跳检查

```mysql

# 拉取镜像
docker pull haproxy:2.0
docker pull osixia/keepalived:2.0.20

# 创建配置文件目录
mkdir -p haproxy
mkdir -p keepalived

#查看网关名称, 看第一排Iface的属性名就是网关名，这里是：Iface：ens33， Gateway：192.168.226.2
netstat -rn
 

# 创建Keepalived配置文件，更改interface ens33， 192.168.226.200/24 dev ens33中的网关名称,虚拟ip根据上面的网关:192.168.226.2进行配置,只要没有被占用就可以，我这里是用了200
cat > keepalived/keepalived.conf << EOL
global_defs {
   router_id haproxy_router
}

vrrp_script haproxy_check {
   script "killall -0 haproxy"
   interval 1
   weight 2
}

vrrp_instance VI_1 {
   interface ens33
   virtual_router_id 51
   priority 100
   advert_int 1
   state MASTER
   virtual_ipaddress {
       192.168.226.200/24 dev ens33
   }
   track_script {
       haproxy_check
   }
}
EOL

# 创建HAProxy配置文件
cat > haproxy/haproxy.cfg << EOL
global
    log stdout local0
    maxconn 4096
    user haproxy
    group haproxy
    daemon

defaults
    log global
    mode    tcp
    option  tcplog
    option  dontlognull
    retries 3
    timeout connect 5000
    timeout client  50000
    timeout server  50000

listen stats
    bind *:9000
    mode http
    stats enable
    stats hide-version
    stats realm Haproxy\ Statistics
    stats uri /
    stats auth admin:admin

frontend rabbitmq_front
    bind *:5672
    default_backend rabbitmq_back

backend rabbitmq_back
    balance roundrobin
    server rabbitmq-node1 192.168.226.128:5672 check inter 5000 rise 2 fall 3
    server rabbitmq-node2 192.168.226.129:5672 check inter 5000 rise 2 fall 3
    server rabbitmq-node3 192.168.226.130:5672 check inter 5000 rise 2 fall 3
EOL



# 在第二台服务器上，打开keepalived/keepalived.conf文件，将priority设置为80，并将state设置为BACKUP：
vi keepalived/keepalived.conf

#更改内容
vrrp_instance VI_1 {
   ...
   priority 80
   state BACKUP
   ...
}
# 在第三台服务器上，打开keepalived/keepalived.conf文件，将priority设置为60，并将state设置为BACKUP：
vrrp_instance VI_1 {
   ...
   priority 60
   state BACKUP
   ...
}

#更改内容
vrrp_instance VI_1 {
   ...
   priority 60
   state BACKUP
   ...
}

#每台服务器执行：启动keepalived，【 --copy-service 】解决sed: can't move '/container/service/keepalived/assets/keepalived. confgNngGN' to '/container/service/keepalived/assets/keepalived.conf': Resource busy （Sed:不能移动'/container/service/keepalived/assets/keepalived。confgNngGN' to '/container/service/keepalive /assets/keepalive .conf':资源繁忙） 问题
docker run -d --name keepalived --cap-add=NET_ADMIN --net=host --restart=always -v `pwd`/keepalived/keepalived.conf:/container/service/keepalived/assets/keepalived.conf osixia/keepalived:2.0.20 --copy-service

#验证keepalived ，ping的通说明对了
ping 192.168.226.200

#每台服务器执行：启动haproxy,通过左边的7070端口访问负载均衡集群，9000是后台界面
docker run -d --name haproxy -p 7070:5672 -p 9000:9000 --restart=always -v `pwd`/haproxy/haproxy.cfg:/usr/local/etc/haproxy/haproxy.cfg haproxy:2.0

#访问haproxy后台界面验证
http://192.168.226.128:9000/
http://192.168.226.129:9000/
http://192.168.226.130:9000/

#验证通过虚拟Ip 7070访问RabbitMQ集群
telnet 192.168.226.200 7070

注意：
后台通过192.168.226.200，7070 虚拟ip访问集群的时候，可以通过docker logs -f haproxy 监控RabbitMQ集群调用情况 (这里是主节点的haproxy，keepalived配置是 priority 100，state MASTER 的就是主节点)
```

### Docker部署Nginx

```mysql
# 拉取Nginx镜像
docker pull  nginx

# 创建名为 nginx 的文件夹
mkdir nginx
# 创建日志目录
mkdir nginx/logs

# 创建 Nginx 的配置文件：1.upstream 代表的是负载均衡服务组 2.listen  [::]:80; 代表可以通过IP:80进行访问 3.location /api2/  代表可以通过Nginx IP:8088/api2访问 4.root   /usr/share/nginx/html; root代表本地的路径，nginx内部的路径（如果是容器中就是容器内部的路径）
cat > nginx/default.conf << EOL

upstream product {
    server 192.168.226.128:7073;
    server 192.168.226.128:7074;
}

upstream repertory {
    server 192.168.226.129:7073;
    server 192.168.226.129:7074;
}

upstream consul {
    server 192.168.226.128:8500;
    server 192.168.226.129:8500;
    server 192.168.226.130:8500;
}

server {
    listen       80;
    listen  [::]:80;
    server_name  localhost; # localhost是域名的意思

    location /api1/ {
        proxy_pass http://product/;
    }

    location /api2/ {
        proxy_pass http://repertory/;
    }
     
    location / {
        proxy_pass http://consul/;
    }
     
    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }
}

EOL



# 启动 Nginx 容器
docker run -d --name nginx \
-p 8088:80 \
-v $(pwd)/nginx/default.conf:/etc/nginx/conf.d/default.conf \
-v $(pwd)/nginx/logs:/var/log/nginx \
nginx

#访问测试
192.168.226.128:8088

注意：
1.查看错误日志：
docker exec -it nginx cat /var/log/nginx/error.log
2.查看运行日志
docker exec -it nginx cat /var/log/nginx/access.log
3.空nginx.conf备用：

user  nginx;
worker_processes  auto;

error_log  /var/log/nginx/error.log notice;
pid        /var/run/nginx.pid;


events {
    worker_connections  1024;
}


http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;

    log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                      '$status $body_bytes_sent "$http_referer" '
                      '"$http_user_agent" "$http_x_forwarded_for"';

    access_log  /var/log/nginx/access.log  main;

    sendfile        on;
    #tcp_nopush     on;

    keepalive_timeout  65;

    #gzip  on;

    include /etc/nginx/conf.d/*.conf;
}

4.空的default.conf备用
server {
    listen       80;
    listen  [::]:80;
    server_name  localhost;
    
    location / {
        root   /usr/share/nginx/html;
        index  index.html index.htm;
    }

    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }
}

```

### Docker上传镜像到Docker Hub

```mysql
#登录
docker login
#用户名：这里注意要是小写
libaibaihi
#密码：
qSNz9ofsnKb00A4c
#更新镜像版本号名称 repertory.web.api:latest是目前的名称,libaibaihi/repertory.web.api:1.0是现在的名称,libaibaihi是Docker Hub的用户名
docker tag repertory.web.api:latest libaibaihi/repertory.web.api:1.0
#推送到Docker Hub
docker push libaibaihi/repertory.web.api:1.0
# 查看我的仓库
https://hub.docker.com/repositories/libaibaihi

#拉取，不同服务器拉取不用登录，因为推送到的是公共的仓库，所以直接可以拉取，但是要输入用户名
docker pull libaibaihi/repertory.web.api:1.0

#启动容器，注意-v /root/Repertory.Web.Api/:/app ，其实是不需要挂载的，从本地镜像执行，从docker hub上面拉取都是不需要挂载的，挂载的好处就是更新的时候方便一些，应该不挂载就是用镜像里面的项目文件，挂载了就是用挂载的项目文件
docker run -d -p 7075:7075 --restart=always --name Repertory.Web.Api2 libaibaihi/repertory.web.api:1.0 --urls="http://*:7075"
```

### Docker部署Consul

```mysql
一、单独部署
# 下载Consul
docker pull consul
# 运行
docker run -d --name=consul \
  -p 8500:8500 \
  -p 8600:8600/udp \
  consul agent -server -ui -bootstrap-expect=1 -client=0.0.0.0
  
# 访问
http://192.168.226.129:8500/ui/dc1/services

二、部署集群
！！！注意：只要后端的服务部署在Consul同一台虚拟机，访问的时候就有不知道这样的主机No such host is known的问题
！！！注意：只要后端的服务部署在Consul同一台虚拟机，访问的时候就有不知道这样的主机No such host is known的问题
！！！注意：只要后端的服务部署在Consul同一台虚拟机，访问的时候就有不知道这样的主机No such host is known的问题
注意：这里有一个天坑的地方，如果部署Consul集群的虚拟机上面，同样部署了后端接口，并且是同一个ip地址，就会出现Host节点与ip映射不了的问题，如果后端服务不在Consul集群的ip节点上，就不会有这个解析的问题,我发现是这个配置：  -advertise=192.168.226.128 \ 导致的，单节点中加入这个配置，访问对应后端部署服务的时候就有这个：不知道这样的主机No such host is known的问题，单节点中可以不加这个属性  -advertise=192.168.226.128 \配置，但是集群节点中需要加入这个属性，不然主节点没有办法与集群进行通信


# 下载Consul镜像
docker pull consul
# 创建文件夹，每台服务器执行
sudo mkdir -p /consul/data
# 服务器1（192.168.226.128）：
 docker run -d \
  --name consul-server-1 \
  -p 8301:8301 -p 8301:8301/udp \
  -p 8302:8302 -p 8302:8302/udp \
  -p 8300:8300 \
  -p 8500:8500 \
  --net=host \
  -e CONSUL_BIND_INTERFACE=ens33 \
  -e CONSUL_CLIENT_INTERFACE=ens33 \
  -v /consul/data:/consul/data \
  consul agent -server \
  -bootstrap-expect=3 \
  -ui \
  -data-dir=/consul/data \
  -node=server-1 \
  -advertise=192.168.226.128 \
  -retry-join=192.168.226.129 \
  -retry-join=192.168.226.130
  
# 服务器2（192.168.226.129）：
docker run -d \
  --name consul-server-2 \
  -p 8301:8301 -p 8301:8301/udp \
  -p 8302:8302 -p 8302:8302/udp \
  -p 8300:8300 \
  -p 8500:8500 \
  --net=host \
  -e CONSUL_BIND_INTERFACE=ens33 \
  -e CONSUL_CLIENT_INTERFACE=ens33 \
  -v /consul/data:/consul/data \
  consul agent -server \
  -ui \
  -data-dir=/consul/data \
  -node=server-2 \
  -advertise=192.168.226.129 \
  -retry-join=192.168.226.128 \
  -retry-join=192.168.226.130

# 服务器3（192.168.226.130）：
docker run -d \
  --name consul-server-3 \
  -p 8301:8301 -p 8301:8301/udp \
  -p 8302:8302 -p 8302:8302/udp \
  -p 8300:8300 \
  -p 8500:8500 \
  --net=host \
  -e CONSUL_BIND_INTERFACE=ens33 \
  -e CONSUL_CLIENT_INTERFACE=ens33 \
  -v /consul/data:/consul/data \
  consul agent -server \
  -ui \
  -data-dir=/consul/data \
  -node=server-3 \
  -advertise=192.168.226.130 \
  -retry-join=192.168.226.128 \
  -retry-join=192.168.226.129
# 查询Consul集群成员列表
docker exec -it consul-server-1 consul members -http-addr="192.168.226.128:8500"


# 访问集群
# 1.服务通过主节点注册后，也会在129,130这两个从节点上面进行访问
# 2.可以通过nginx代理，实现负载均衡，调用的时候直接访问nginx就可以，不过一般实际
http://192.168.226.128:8500

Tips：
代码解析：
docker run -d \
  --name consul-server-1 \ # 容器名称
  --net=host \ # 使用宿主机网络
  -e CONSUL_BIND_INTERFACE=ens33 \ # 绑定网卡
  -e CONSUL_CLIENT_INTERFACE=ens33 \ # 客户端网卡
  -v /consul/data:/consul/data \ # 将宿主机上的/consul/data目录映射到容器中的/consul/data，只要挂载了里面注册好的服务，重启了后目录也不会丢失
  consul agent -server \
  -bootstrap-expect=3 \ # 预期的服务器节点数量
  -ui \ # 启用Consul UI
  -data-dir=/consul/data \ # 在容器内使用/consul/data目录存储数据
  -node=server-1 \ # 节点名称
  -advertise=192.168.226.128 \ # 广播地址
  -retry-join=192.168.226.129 \ # 加入集群的其他节点
  -retry-join=192.168.226.130

```

#### Net Core注册Consul

**引用包， 使用Consul 官方操作Net类库**
Install-Package Consul

```c#
using System;
using Consul;

namespace ConsulDemo
{
    class Program
    {
        static void Main(string[] args)
        {
            // 创建Consul客户端对象
             var consul = new ConsulClient(new ConsulClientConfiguration() { Address = new Uri("http://192.168.226.129:8500") });
            // 注册服务A
            var registration = new AgentServiceRegistration()
            {
                ID = "Product.Web.Api1",
                Name = "Product.Web.Api",
                Address = "192.168.226.129",
                Port = 7073,
                Check = new AgentServiceCheck()
                {
                    HTTP = "http://192.168.226.129:7073/weatherforecast",
                    Interval = TimeSpan.FromSeconds(1)
                }
            };
            consul.Agent.ServiceRegister(registration).Wait();

            // 注册服务B
            var registration2 = new AgentServiceRegistration()
            {
                ID = "Product.Web.Api2",
                Name = "Product.Web.Api",
                Address = "192.168.226.128",
                Port = 7073,
                Check = new AgentServiceCheck()
                {
                    HTTP = "http://192.168.226.128:7073/weatherforecast",
                    Interval = TimeSpan.FromSeconds(1)
                }
            };
            consul.Agent.ServiceRegister(registration2).Wait();

            // 查询服务
            var services = consul.Agent.Services().Result.Response;
            foreach (var service in services.Values)
            {
                if (service.Service.Equals("Product.Web.Api", StringComparison.OrdinalIgnoreCase))
                {
                    Console.WriteLine($"Found service: {service.ID}, {service.Service}, {service.Address}, {service.Port}");
                }
            }

            Console.WriteLine("Press any key to exit...");
            Console.ReadKey();

            // 注销服务
            consul.Agent.ServiceDeregister("my-service-1").Wait();
        }
    }
}


```

#### Net Core请求Consul(轮询负载均衡)

```c#
using System;
using System.Collections.Generic;
using System.Linq;
using Consul;

namespace ConsulDemo
{
    class Program
    {
        static void Main(string[] args)
        {
            // 创建Consul客户端对象
            var consul = new ConsulClient(new ConsulClientConfiguration() { Address = new Uri("http://192.168.226.129:8500") });

            // 服务名称
            var serviceName = "Product.Web.Api";

            // 获取所有服务实例
            var queryResult = consul.Health.Service(serviceName, null, true).Result;
            var serviceInstances = queryResult.Response.Select(s => s.Service).ToList();

            // 初始化轮询索引
            var index = 0;

            // 循环调用服务
            while (true)
            {
                // 如果没有可用的服务，输出提示信息并退出循环
                if (serviceInstances.Count == 0)
                {
                    Console.WriteLine("No service instances available.");
                    break;
                }

                // 选择下一个服务实例
                var serviceInstance = serviceInstances[index];
                //这行代码是精华：取模运算实现了一个简单的轮询算法
                index = (index + 1) % serviceInstances.Count;

                // 构建服务地址
                var serviceAddress = $"http://{serviceInstance.Address}:{serviceInstance.Port}";

                // 发起HTTP请求
                using (var client = new System.Net.Http.HttpClient())
                {
                    var response = client.GetAsync($"{serviceAddress}/weatherforecast").Result;
                    if (response.IsSuccessStatusCode)
                    {
                        var result = response.Content.ReadAsStringAsync().Result;
                        Console.WriteLine($"Service instance: {serviceInstance.ID}, {serviceInstance.Address}, {serviceInstance.Port}, Response: {result}");
                        //   break;
                        Thread.Sleep(1000);
                    }
                    else
                    {
                        Console.WriteLine($"Failed to call service instance: {serviceInstance.ID}, {serviceInstance.Address}, {serviceInstance.Port}, StatusCode: {response.StatusCode}");
                    }
                }
            }

            Console.WriteLine("Press any key to exit...");
            Console.ReadKey();
        }
    }
}

```

```c#


```

### Docker部署Ocelot-Consul

> Ocelot不需要部署，直接引用类在Net框架中进行使用

上下游路径的三个好处：

- 路径解耦：保留真实的下游路径访问路径，这样当下游服务路径发生改变时，上游路径可以不受影响，便于维护和更新。
- 简化客户端请求：上游路径通常较短，易于配置和使用，对客户端来说更容易理解和使用。
- 聚合请求：可以将单个请求，通过同时请求三个api将结果聚合起来进行访问给客户端，需要实现Ocelot聚合接口，方便管理聚合请求，也可以解耦

上游路径就是一个大业务的模块，下游路径是具体的请求，比如大的版块请求路径是：https://admin.shengweiinc.com/products/CategoryId/GetID，下游路径具体的是：192.168.226.128:8080/api/v1/products/CategoryId/GetID,

其中https://admin.shengweiinc.com是网关映射的域名地址，无论请求哪个微服务都是这个域名，然后根据Skums映射到对应的微服务（也就是上游路径），然后上游路径会匹配下游路径中的请求，也就是将https://admin.shengweiinc.com/products/CategoryId/GetID的请求转换为：192.168.226.128:8080/api/v1/products/CategoryId/GetID，可以参考一下配置：

```json
{
  "Routes": [
    {
      "DownstreamPathTemplate": "/api/v1/orders/{everything}",
      "DownstreamScheme": "http",
      "DownstreamHostAndPorts": [
        {
          "Host": "192.168.226.128",
          "Port": 8081
        }
      ],
      "UpstreamPathTemplate": "/orders/{everything}",
      "UpstreamHttpMethod": [ "GET", "POST", "PUT", "DELETE" ]
    },
    {
      "DownstreamPathTemplate": "/api/v1/products/{everything}",
      "DownstreamScheme": "http",
      "DownstreamHostAndPorts": [
        {
          "Host": "192.168.226.128",
          "Port": 8080
        }
      ],
      "UpstreamPathTemplate": "/products/{everything}",
      "UpstreamHttpMethod": [ "GET", "POST", "PUT", "DELETE" ]
    }
  ],
  "GlobalConfiguration": {
    "BaseUrl": "https://admin.shengweiinc.com"
  }
}

```

正式Consul配置,新建文件：ocelot.json

```json
{
  "Routes": [
    {
      "UpstreamPathTemplate": "/Api/{controller}/{action}", // 上游请求路径模板
      "DownstreamPathTemplate": "/{controller}/{action}", // 下游请求路径模板
      "UpstreamHttpMethod": [ "GET", "POST", "DELETE", "PUT" ], // 允许的上游 HTTP 请求方法
      "DownstreamScheme": "http", // 下游请求的协议（HTTP 或 HTTPS）
      "UseServiceDiscovery": true, // 是否使用服务发现功能
      "ServiceName": "Apigetway.Web.Api", // 要请求的服务名称
      "LoadBalancerOptions": { // 负载均衡选项
        // RoundRobin: 轮询到每个服务，LeastConnection: 选择当前连接数最少的那个，NoLoadBalance: 不进行负载均衡，一直请求到同一个服务，CookieStickySessions: 基于会话 Cookie 将请求路由到同一个后端服务实例。
        "Type": "RoundRobin" // 负载均衡类型
      }
    }
  ],
  "GlobalConfiguration": { // 全局配置
    "BaseUrl": "http://localhost:5000", // Ocelot 网关的基本 URL，就是启动网关的api路径
    "ServiceDiscoveryProvider": { // 服务发现提供者配置
      "Host": "192.168.226.129", // 服务发现提供者的主机名或 IP 地址,如果是nginx代理Consul集群，那么这里配置Nginx的地址，下面配置Nginx的端口
      "Port": 8500, // 服务发现提供者的端口
      "Type": "Consul" // 服务发现提供者的类型（如 Consul）
    }
  }
}


```

Program.cs

```c#
using Ocelot.DependencyInjection;
using Ocelot.Middleware;
using Ocelot.Provider.Consul;

var builder = WebApplication.CreateBuilder(args);

builder.Configuration.AddJsonFile("ocelot.json", false, true);
builder.Services.AddOcelot().AddConsul();

var app = builder.Build();
app.UseOcelot();
app.Run();


```

验证

```c#
// 访问

```

### Docker部署Consul Template动态服务更新

```dockerfile
# 创建自定义 Docker 网络
docker network create consul-nginx-network

# 拉取镜像
docker pull hashicorp/consul-template

# 创建文件夹
#模板配置：consul-template 
#consul-template生成的文件：consul 
sudo mkdir -p /etc/consul-template  
sudo mkdir -p /etc/nginx/conf.d/consul 

#开放容器内操作文件权限
sudo chmod -R 777 /etc/consul-template
sudo chmod -R 777 /etc/nginx/conf.d/consul


# 创建 Consul 模板文件
sudo cat > /etc/consul-template/nginx.conf.ctmpl << EOF
{{ $serviceName1 := "Product1" }}
{{ $serviceName2 := "Product2" }}

upstream {{ $serviceName1 }} {
    {{range service $serviceName1}}
    server {{.Address}}:{{.Port}};
    {{end}}
}

upstream {{ $serviceName2 }} {
    {{range service $serviceName2}}
    server {{.Address}}:{{.Port}};
    {{end}}
}

server {
    listen 80;
    listen  [::]:80;
    
    location /{{ $serviceName1 }}/ {
        proxy_pass http://{{ $serviceName1 }}/;
    }

    location /{{ $serviceName2 }}/ {
        proxy_pass http://{{ $serviceName2 }}/;
    }
}
EOF

#加入用户到docker安全组
#先查看用户姓名，linux下查看
whoami   #root
sudo usermod -aG docker root  #加入安全组

#查看docker 安装路径，配置到command中
which docker #输出：/usr/bin/docker
!!![]我的进度在这里：docker -H unix:///var/run/docker.sock exec nginx nginx -s reload 在consul-template中调用这个命令，尝试一下，可以的话就成功用docker部署了consul-template

# 创建 Consul Template 配置文件，/usr/bin/docker请用上面的命令进行配置
sudo cat > /etc/consul-template/config.hcl << EOF
consul {
  address = "192.168.226.128:8500"
}

template {
  source = "/etc/consul-template/nginx.conf.ctmpl"
  destination = "/etc/nginx/conf.d/consul/myapp.conf"
  command = "su docker -H unix:///var/run/docker.sock exec nginx nginx -s reload"
}
EOF

# Nginx中的配置nginx/nginx.conf中是主配置，配置线程等全局配置，default.conf是配置负载均衡等一些映射关系，这下面的内容在nginx.conf中进行配置
http {
    ...
    include /etc/nginx/conf.d/*.conf; #conf.d目录下面所有的.conf配置包含进来
    include /etc/nginx/conf.d/consul/*.conf;  #consul-template自动生成的配置
    ...
}

#启动consul-template
#挂载 Docker 守护程序套接字，直接可以在docker中执行docker命令，比如在consul-template中刷新nginx配置（/usr/bin/docker exec nginx nginx -s reload）,挂载Docker守护程序套接字：-v /etc/nginx/conf.d/consul:/etc/nginx/conf.d/consul \ 


docker run -d \
  --name consul-template \
  --network consul-nginx-network \
  -v /etc/consul-template/nginx.conf.ctmpl:/etc/consul-template/nginx.conf.ctmpl \
  -v /etc/consul-template/config.hcl:/etc/consul-template/config.hcl \
  -v /etc/nginx/conf.d/consul:/etc/nginx/conf.d/consul \ 
  hashicorp/consul-template

#启动Nginx
docker run -d \
  --name nginx \
  --network consul-nginx-network \
  -p 8088:80 \
  -v /etc/nginx/conf.d:/etc/nginx/conf.d \
  nginx:latest


#命令行执行Consul Template，先下载最新版本
wget https://releases.hashicorp.com/consul-template/0.27.0/consul-template_0.27.0_linux_amd64.zip
#安装压缩包
sudo yum install unzip
#解压
unzip consul-template_0.27.0_linux_amd64.zip
#移到执行目录
sudo mv consul-template /usr/local/bin/
#检查安装版本
consul-template -version
#启动consul-template，注意更新路径
consul-template -config=/etc/consul-template/config.hcl

```



### Docker部署Nginx+Ocelot+Consul+Keepalived

- 这个架构中这三个工具分别有什么作用？


**Nginx(负载均衡，反向代理)**: 负责高并发的场景，负责负载均衡，静态文件处理（首页，异常页面），反向代理（隐藏真实的IP），判断请求如果不是访问的缓存(比如缓存统计页面API的结果，缓存静态的首页面不会变化)，不是访问静态页面后将请求转发给Ocelot

**Ocelot(鉴权，认证，服务路由转发)**：负责鉴权（是不是合法的用户，用户的权限能不能访问这个接口），熔断(某个关键服务出现了3次异常，为了防止服务堆积造成所有服务瘫痪，停止访问关键服务1分钟，等待服务恢复后继续请求)，限流（每秒允许多少流量进入），请求通过鉴权等步骤后，分析请求的路径，转发到Consul对应的服务上面，Ocelot自身是不能感知到Consul中的哪些服务是正常的哪些是已经挂掉，这个需要Consul服务自己做判断，Ocelot中的Consul配置情况需要Consul自身去进行维护，Ocelot只负责转发到Consul当中

**Consul(服务注册，服务发现)**:服务注册，服务发现，接收Ocelot网关通过鉴权，限流后的流量到这边进行服务发现，可以把Consul看做一个Key,List<value>的一个集合，负责维护微服务的列表，不过Consul支持心跳检查，可以检测微服务节点的动态，Cousul还可以当做配置中心，保存数据库配置，功能开关配置，还可以用分布式锁，很牛

**总结：**

刚看完大佬的文章有新的体会，[链接](https://www.cnblogs.com/Leo_wl/p/14043862.html)

- 用户开始访问的时候是访问Keepalived集群，产生的虚拟Ip，然后访问到Nginx集群上面，然后Nginx集群求到Ocelot的集群,Ocelot转发到Consul,开始一关使用Nginx是为了提高Ocelot的高可用性，和负载均衡，并且利用Nginx的高并发能力与反向代理能力

- 进入到Ocelot后，通过了鉴权后进入到Consul，这个时候负载均衡是Ocelot来做的，然后配置到Ocelot配置中的地址也是Nginx中配置的Consul集群地址

  架构图

  ![](Vue进步.assets/架构图.png)

  

### Docker Compose

```mysql
#下载docker-compose ，1.29.2 是版本号，可以调整
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
#添加执行权限
sudo chmod +x /usr/local/bin/docker-compose
#判断是否成功安装
docker-compose --version

一、单服务部署

#创建yml文件,保存为docker-compose.yml文件，放到linux中，libaibaihi/repertory.web.api:1.0镜像可以不提前拉取，在docker-compose up -d的时候会自动拉取 ，version代表的是docker-compose的版本号，表示遵守这个版本的语法与书写规范
cat <<EOF > docker-compose.yml
version: '3'
services:
  repertory:
    container_name: repertory-service
    image: libaibaihi/repertory.web.api:1.0
    ports:
      - 7075:7075
    restart: always
    command: --urls="http://*:7075"
EOF

#原命令转换：docker run -d -p 7075:7075 --restart=always --name Repertory.Web.Api2 libaibaihi/repertory.web.api:1.0 --urls="http://*:7075"


#run起来，要加-d，不然ctrl+c 就会退出容器
docker-compose up -d

#日志
docker-compose logs

#测试
curl http://192.168.226.128:7075/weatherforecast

二、【Compose】多服务部署，比如下面的命令是模仿启动三个后端程序，一个前端程序，一个数据库程序
#注释版本（下面有没有注释的版本）
cat <<EOF > docker-compose.yml
version: '3.8'
services:
  backend-api-service1:
    container_name: backend-api-service1  # 容器名称
    image: libaibaihi/repertory.web.api:1.0  # 后端 API 镜像
    ports:
      - 7075:7075  # 端口映射，将主机的7075端口映射到容器的7075端口
    restart: always  # 容器重启策略
    command: --urls="http://*:7075"  # 启动命令
    depends_on:
      - database  # 依赖于 database 服务
    resources:  # 资源限制和配额
      limits:
        cpus: '0.5'  # CPU限制
        memory: '512M'  # 内存限制
      reservations:
        cpus: '0.2'  # CPU预留值
        memory: '256M'  # 内存预留值
    volumes:
      - backend-data:/app/data  # 数据卷

  backend-api-service2:
    container_name: backend-api-service2  # 容器名称
    image: libaibaihi/repertory.web.api:1.0  # 后端 API 镜像
    ports:
      - 7076:7075  # 端口映射，将主机的7076端口映射到容器的7075端口
    restart: always  # 容器重启策略
    command: --urls="http://*:7075"  # 启动命令
    depends_on:
      - database  # 依赖于 database 服务
    resources:  # 资源限制和配额
      limits:
        cpus: '0.5'  # CPU限制
        memory: '512M'  # 内存限制
      reservations:
        cpus: '0.2'  # CPU预留值
        memory: '256M'  # 内存预留值
    volumes:
      - backend-data:/app/data  # 数据卷

  backend-api-service3:
    container_name: backend-api-service3  # 容器名称
    image: libaibaihi/repertory.web.api:1.0  # 后端 API 镜像
    ports:
      - 7077:7075  # 端口映射，将主机的7077端口映射到容器的7075端口
    restart: always  # 容器重启策略
    command: --urls="http://*:7075"  # 启动命令
    depends_on:
      - database  # 依赖于 database 服务
    resources:  # 资源限制和配额
      limits:
        cpus: '0.5'  # CPU限制
        memory: '512M'  # 内存限制
      reservations:
        cpus: '0.2'  # CPU预留值
        memory: '256M'  # 内存预留值
    volumes:
      - backend-data:/app/data 
      frontend-web-service:
    container_name: frontend-web-service  # 容器名称
    image: your-frontend-image:tag  # 前端镜像
    ports:
      - 80:80  # 端口映射，将主机的80端口映射到容器的80端口
    restart: always  # 容器重启策略
    depends_on:
      - backend-api-service1
      - backend-api-service2
      - backend-api-service3
    volumes:
      - frontend-logs:/app/logs  # 数据卷

  database-service:
    container_name: database-service  # 容器名称
    image: your-database-image:tag  # 数据库镜像
    restart: always  # 容器重启策略
    environment:
      - POSTGRES_USER=your-username  # 数据库用户名
      - POSTGRES_PASSWORD=your-password  # 数据库密码
      - POSTGRES_DB=your-database-name  # 数据库名称

    volumes:
      backend-data:  # 后端数据卷
      frontend-logs:  # 前端日志卷

    networks:
      your-network:  # 自定义网络
EOF
#没有注释的版本
cat <<EOF > docker-compose.yml
version: '2.4'
services:
  backend-api-service1:
    container_name: backend-api-service1
    image: libaibaihi/repertory.web.api:1.0
    ports:
      - 7075:7075
    restart: always
    command: --urls="http://*:7075"
    volumes:
      - backend-data:/app/data
    cpus: '0.5'
    mem_limit: '512M'

  backend-api-service2:
    container_name: backend-api-service2
    image: libaibaihi/repertory.web.api:1.0
    ports:
      - 7076:7075
    restart: always
    command: --urls="http://*:7075"
    volumes:
      - backend-data:/app/data
    cpus: '0.5'
    mem_limit: '512M'

  backend-api-service3:
    container_name: backend-api-service3
    image: libaibaihi/repertory.web.api:1.0
    ports:
      - 7077:7075
    restart: always
    command: --urls="http://*:7075"
    volumes:
      - backend-data:/app/data
    cpus: '0.5'
    mem_limit: '512M'

volumes:
  backend-data:

networks:
  your-network:
EOF


 

#三、常用命令
docker-compose up -d  通过后台启动的方式进行启动服务
docker-compose up -d backend-api-service  只启动我指定的服务，更新的时候比较适合

docker-compose restart backend-api-service 只重启我指定的服务，也可以配合start,stop,logs进行查看   

docker-compose down: 停止并移除 Docker Compose 文件中定义的服务的容器、网络和卷等资源。这将停止并清理掉整个应用程序的容器环境。

docker-compose start: 启动已创建的 Docker Compose 文件中定义的服务的容器。与 up 命令不同，start 命令仅启动已创建但当前处于停止状态的容器，而不会重新创建新的容器。

docker-compose stop: 停止服务

docker-compose restart: 重新启动 

docker-compose ps: 服务状态

docker-compose logs: 日志

```



## VPS

### AWS:VPS

**注册与配置**

1.这个教程前面的都是简单，地址用英文Hong Kong，信用卡卡号这些需要看，后面就是设置root的用户登录权限这个要注意

2.第二个ubuntu版本选择20版本,选择22版本不知道为什么python总是安装失败

3.第三个放开所有入站流量



![1678960479076](Vue进步.assets/1678960479076.png)

4.观察流量

![1678961706449](Vue进步.assets/1678961706449.png)

[AWS免费服务器申请教程](https://www.landafu.com/29528.html)

**SSH登录超时提示**

Network error: Software caused connection abort

网络错误:软件导致连接中断

[登录超时解决](https://blog.csdn.net/FriendshipTang/article/details/113851764)

## Github



**新手该如何使用 GitHub？**

[新手该如何使用 GitHub？](https://zhuanlan.zhihu.com/p/388451230)

**为什么要逛Github?**

走马观花的看项目，学习到的也是浅显的东西，**重点应该在把注意力下沉到计算机基础，网络通信层面**，Github上面的项目可以激发我的好奇心（计算机基础，网络通信的基础是比较枯燥的），并且根据这些项目我可以加入到对应的社群，可能就有未来一起合作开源Github的伙伴（**为什么不主动要求对应的开源作者，我想帮助他们呢**）



**搜索团队使用框架**：github是用来逛的吗？我一直按需搜索来用。比如上次操作Excel内存爆了的问题，使用MiniExcel操作，消耗内存就小很多，我感觉很多大的框架，技术，大家基本上都是用的那些，比如vue,react,不过也需要去学习，比如consul,nginx,jmeter这些也需要看视频才知道，学习了才知道有这些框架，有些特别的也需要知道存在，并且怎么选型（比如MiniExcel）

**代码片段**：就像博客园，CSDN一样把github当作一个代码仓库来使用，也可以当作博客园来使用，这里有着是成套的解决方案，也可以化整为零，只拿对方成套解决方案其中的一部分代码

**学习资料库**: 还可以把它当作学习编程语言的学习库，里面有课程，文档，学习教材

**小工具**：比如京东抢秒杀的助手，还有ChatGPT聊天机器人等，有意思的工具

**启发大脑**： 启发大脑，原来编程还可以做这么有意思的事情！

**技术前沿项目**：比如算法，人工智能，谷歌，微软大公司开源的项目

**学习开源代码**：比如学习开源的net源码，学习开源作者的设计模式，编码的思路

**开源项目**:开源项目可以提高自己的动手能力，链接到彼此志向相同的伙伴，共同进步

**成就感**：这一路日夜兼程，嘲讽批评常伴眼耳，即便辛苦无奈、想过放弃，但为了那微不足道的成就感依然努力着。当然，也收获了不少... 越来越多**拥趸者**，越发**精湛技术能力**，更**高层次思维模式**，还有许多**跨界跨行朋友**。

### 创建一个新的开源项目进行管理

第一步 新建项目

![1709110935975](Vue进步.assets/1709110935975.png)

第二步 描述项目信息

![1709111020753](Vue进步.assets/1709111020753.png)

第三步 拉取项目下来进行管理

![1709111369770](Vue进步.assets/1709111369770.png)

### 查看一个项目开源的时间

linux命令：curl  -k https://api.github.com/repos/tensorflow/hub | grep created_at

![1688618515941](Vue进步.assets/1688618515941.png)

原文链接：https://blog.csdn.net/xinghaikongmeng/article/details/127529654

### 单独下载Github指定文件/文件夹的方法

**多个文件/单个文件下载**：

**使用谷歌插件：GitZip For Github,然后看下面的攻略** 

![1675340170221](Vue进步.assets/1675340170221.png)

**单个文件下载：**

点击 Raw ，如果是不能预览的文件就会自动下载，如果是文件或者代码什么的，会在浏览器中显示，但可以复制浏览器中链接，使用任何一个下载工具，新建下载任务，粘贴地址即可下载单个文件。



![img](https://picx.zhimg.com/v2-f96c95ee8260a99142174c6a3278ea76_r.jpg?source=1940ef5c)

[原攻略地址](https://www.zhihu.com/question/25369412)

### 有趣的一些词汇

**PR**：pr 是 pull request，用直白的话来说就是我这里有一些改动，求求（request）你 pull 到你的分支里吧，还有一个有趣的解释：一个请对方拉取自己的代码的请求（我改了你们的代码，请拉回去看看吧）

**MR** : mr是merge request，合并代码的意思，Gitlab管合并代码叫MR,Github上面叫PR,两种都是一样的作业：推送代码到分支上面，只是叫法不同

**fork** ：一般改动开源的代码，都会从对方的仓库拉取一个分支下来，就是fork拉取一个新的分支，然后自己本地改动好后，在提交到自己的个人仓库，然后等待开源 作者合并到主分支（**PR就是合并到主分支的意思**）

**GVP** ：**Gitee 最有价值开源项目** 是开源中国Gitee综合评定出的优秀开源项目

**Issue**： 发现了开源项目的一个问题，或者有疑问想要作者解答或者需要和大家讨论，可以新建一个Issue

![1672819512313](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1672819512313.png)

![1672820077401](Vue进步.assets/1672820077401.png)： 代表被Gitee官方推荐使用的项目

![1672820102556](Vue进步.assets/1672820102556.png)

**贡献指南**：每个开源代码都有贡献指南，指的是如果我想参与到这个项目中那么需要遵守的一些规范，比如你新建的子分支需要遵守的一些规范，还有你提交后我们会怎么处理你提交的代码，需要一个或者两个以上的维护者通过才可以合并到主分支。

![1672819768832](Vue进步.assets/1672819768832.png)



**热门项目的解释**

**分支数**代表的是依据这个吧版本check下来了多少分支，有两种情况：第一种是单独下载下来的分支，是为了为主分支贡献代码，还有一种情况就是单立门派，比如[MariaDB](https://www.laozuo.org/24402.html)数据库就是Mysql开源下面的一个分支

 

![1672736329755](Vue进步.assets/1672736329755.png)

**按照时间周期看热门项目**： [热门项目](https://github.com/trending)

![1672372886134](Vue进步.assets/1672372886134.png)

**找项目：**直接搜索：中文可以直接搜（中英文切换一起搜）

![1672372413640](Vue进步.assets/1672372413640.png)

**最后更新时间：**

![1672372548279](Vue进步.assets/1672372548279.png)

### 搜索开源技巧

-  **人,社群**:关注Net圈的大佬（百小僧，），还有Net圈子（比如[Net Core核心社区](https://docs.github.com/en/search-github/searching-on-github/searching-wikis)），微软MVC看他们关注的项目，他们关注的人是谁
- **在组织中搜索**: org:dotnetcore sqlsugar

- **name: 项目名**

  ```makefile
  in:name React  搜索`项目名`里面包含`React`的项目:
  in:name React stars:>5000 并且下载数大于5000 (一般设置100)
  in:name React stars:>5000 forks:>3000 并且分支数大于3000 (一般设置)
  ```

- **description: 项目的简要描述**

  搜索`Description` 项目描述里面里面包含`微服务/Microservice`的项目:

  ```
  in:description 微服务微信 （为什么加微信，因为我发现这里的搜索是用了分词的功能的，它会自动拆分词语）
  in:description Microservice
  in:description 微服务 language:c#  查找项目描述里面有微服务，并且开发语言是c#
  in:description 微服务 language:c# pushed:>2023-01-01 并且更新时间大于23年1月1日
  ```

- **README.md: 项目的详细情况的介绍**

  搜索`README.md`里面包含`React`的项目:

  ```
   in:readme React
  ```

[原文地址](https://www.zhihu.com/question/19637879)

[官方地址](https://docs.github.com/en/search-github/searching-on-github/searching-wikis)

## Git

 [Git命令行在线学习](https://oschina.gitee.io/learn-git-branching/)

### Git 大全

[Git大全](https://gitee.com/all-about-git)，使用git的工具，搭建git的服务，在线使用git的站点（比如github,码云），git教程（入门，进阶，菜鸟，博客园,csdn写的好的入门教程，非常不错），还有git的常用命令

### 设置Git下载的代理

把cmd调出来，直接执行，执行成功后，在git cmd那里直接下载github的项目

```
git config --global http.proxy http://127.0.0.1:14008
git config --global https.proxy https://127.0.0.1:14008
```



### 【注意解决冲突后要重新生成啊！！！】拉取代码出现冲突的解决方式

![1672713891593](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1672713891593.png)



**第二种情况**：准确，可以直接使用，我已经和杨魁核实过，这个方法，准确有效

![1676281039597](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1676281039597.png)

### 回滚版本到指定的版本（整个版本）

[回滚整个版本](https://www.jianshu.com/p/07a52781e6e2)

### 合并测试分支到Master分支忽略指定文件

问题：因为我本地的分支拉取了staing分支的最新代码，然后合并到了master分支导致了配置被覆盖。

解决：方案1:本地代码应该从master分支进行拉取，不拉取测试分支

​           方案2：合并代码的时候忽略指定的配置文件，不进行合并

根目录运行命令：

```c#
git config merge.ours.driver true    //开启忽略文件配置
```

[扩展阅读](https://blog.csdn.net/weixin_57653722/article/details/126630275)

### 新建分支

![1665994051499](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1665994051499.png)

### 提交到测试环境

![1665994273847](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1665994273847.png)

### 提交到正式环境

![1665994176916](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1665994176916.png)

### 下载指定分支的代码：

![1646984046098](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1646984046098.png)

![1686539461231](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1686539461231.png)

### **下载代码（新的都是用这个）**：git clone 地址

```c#
//例如
git clone https://shengwei-gz.coding.net/p/shengwei/d/Bailun.Security.Authorize/git
```

**拉取最新代码**：git pull

**提交代码**：git push

**拉取指定分支代码：** git cone -b 分支名称 地址

案例：optimize-eccang-push 是分支名称，后面是代码路径

git clone -b optimize-eccang-push https://e.coding.net/shengwei-gz/shengwei/Bailun.Skums.git

![1651912782431](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1651912782431.png)

**resourcetree中的拉取和获取有什么区别** ：

**1、命令式解释：**
前者是 pull，后者是 fetch，pull 等于 fetch + merge。

**2、大白话解释：**
拉取：把你本地仓库没有 ，而远程仓库有的更新写入到你本地;

获取：用来查看对于你本地仓库的状态来说远程仓库是否有更新，仅此而已，并不会使你的本地仓库发生改变



##  Swagger-Editor

前后端分离开发中，自动生成接口文档，请求的方法，入参的参数，返回值都自动生成，不需要和前端争吵了

## MiniProfiler

第三方性能组件，通过该组件可以方便查看请求性能、异常堆栈、数据库操作等信息。默认在 `Swagger` 首页左上角显示。

## MiniExcel

使用底层的方式操作Excel，可以使用几M内存，操作几百M的Excel

## Exceptionless

Net Core集成Exceptionless[分布式](https://so.csdn.net/so/search?q=分布式&spm=1001.2101.3001.7020)日志功能以及全局异常过滤



## K8S

**K8S**:kubernetes，简称K8s，是用8代替名字中间的8个字符“ubernete”而成的缩写，是用来管理容器的一个工具，由Google公司2014年开源，特点可以支持自动部署，自动重启，模块化部署应用

**容器**：容器就像是简易版的虚拟机，每个容器可以设置不同的操作系统，不同的开发语言环境，不同的应用程序，不同容器之间运行的内容是相互隔离的（就是如果部署负载均衡的话，即使同一个服务器的容器崩溃了，同服务器的另一个容易也不会受到影响）

但是我们打包部署就方便了很多。。jenkins+harbor+k8s

## Furion

### 关于框架的思考

[高校学习开源框架](https://www.zhihu.com/question/19637879)

开始看这个框架的时候处于不了解的阶段，看了作者的更新历史记录，关注量从几千到几万，再到几十万，上百万，再到国家级别的花木兰社区重点培育的开源项目，真的牛，而且看到作业在入门指南中写的一些关于建立自己的框架注意的一些经验，还有开发上一些经验与便捷的地方（我想他肯定有很多话和经验想要分享出来，但是又不能写成入门科普文档，应该这里面的内容是经过精挑细选后才选择写到这个位置的）

昨天的学习过程中学习到了将项目脚手架集成到VS当中，还有可以把自己的项目（公司的项目）发布到Nuget方便大家的使用与更新，还有可以根据公司的实际情况搭建自己项目的脚手架，还了解到了github一些关键字的用法，与作者目前框架建设的成就与影响力

这些框架这么完善，完全可以看github上面成熟的start的框架进行学习，或者到现有的Furion这种类似的框架里面进行学习框架的一部分东西，比如里面的定时调度的解决方案，好的异常处理方式的解决方案，这个也可以利用开源框架

**看开源框架要以一个点做目标，去攻击，这个项目哪个点吸引你，你就去看它的源码是怎么实现的，比如MiniExcel可以通过字节码的方式操作Excel，如果我要写操作的工具，我就可以去借鉴它的算法是怎么样子的，微软的进程之间通信的类库，可以跨进程的方式，通过内存共享不同进程的数据，那我也可以去借鉴它的代码。**

看框架的源码可以逐步逐步，一个模块，一个模块，一个功能，一个功能，一个类，一个类的啃下去，虽然关联关系中还是很坑爹，但是就是需要花时间的（写在2023年1月12日看Furion源码Efcore看奔溃了），刚才看了下大家关于读开源代码学习的思路，总结一下：

**向框架所在的社群帮助**: 一般你碰到的问题，可能其他的人也遇到过，加入Furion框架学习的社群和大家一起学习，这个比在代码中找答案省心

**在Issus中找答案**：可能已经被作者或者其他人发现，并且已经修复

**给框架代码弄一个思维导图**: 这样对框架的整体源码会有一个整体的了解，知道遇到哪个问题要到哪一块的代码去找（也可以根据相关模块的关键字找到对应的源码位置）

**项目解决了什么问题？**：我觉得Furion解决了什么问题？ 好像真的是**工具包的存在**，Swagger中的自动生成控制器，方便读取配置/强类型配置的方法，好的异常提示功能

**项目的缺点是什么？** ：项目的实现中有哪些缺点？比如Furion架构是工具包的存在，我们可以利用这个特点去里面找可以重用的代码，可以在架构设计思路上面它可能比ABP的思想要低很多

**有没有同类的项目**： 他们的相同点与不同点在哪里？



**权限管理,用户登录,高并发，安全，运维**

**业务了解，技术选型**

我发现当你学习了一些框架的知识后，体系后，内心就会很兴奋，很激动，想通过这些现有的概念加入自己的价值观与理解上去，创造出属于自己的开发框架。

我可不可以把开发框架看做是一个巨大的通用的类库工具箱，这个类库工具箱里面集成了权限认证，异常处理，日志管理,Excel操作等通用的东西，然后我可以在现有的项目里面引用这个框架，然后使用里面部分的功能这个是一种，我也可以单独把这个开源模块的代码拷贝出来，然后复制到已有的框架中，集成这部分功能。

**应用启动**：昨天学习了通用的启用方法，不管Net的版本，使用作者统一的App.Run()入口进行注册服务，功能，管道与中间件，并且在官方的基础上面优化了startup.cs的功能，比如可以自定义启动类，可以声明多个启动类（只需要按照规范继承Startup父类与声明规范的ConfigureServices与Configure方法），还可以配合服务之间的依赖关系，可能会有疑问，这么多服务注册类与管道注册类，程序启动的时候怎么放到一起注册呢？Furion框架会自动汇总所有的服务/管道注册类（并且也会自动汇集所有的配置），并且根据配置的依赖关系按顺序注册这些管道服务。刚才看了一下文档一点组件化启动内容，  组件化启动就是我们可以监控服务注册的时候注册前，注册中，注册后的动态，有三种组件：

- **服务组件**（就是在ConfigureServices中对象注册的组件，继承：IServiceComponent）

- **中间件组件**(Configure中注册的组件)

- **Web组件**：(未知的组件)

组件启动的方式可以解决：

1. 组件启动时，启动中，启动后的监控 
2. 后期服务注册臃肿，注册的方法超级长，注册的服务超级多，代码堆在一起难以理解 
3. 3.组件相互依赖的问题，还有有些组件需要在前面注册，有些组件需要在某个组件注册后再进行注册（比如ABP注册的那个坑），这些依赖关系如果换了一个开发人员，那就是天坑

我们可以依靠配置注册顺序解决以上三个问题（ABP框架中也有DependsOn配置服务之间的依赖关系）。

**配置与选项**：其中还有配置相关的读取，比如ConfigurationBuilder类，读取所有的配置初始化到Builder对象中进行读取，也有按需要读取的Configuration类，我们可以根据实际需要通过构造函数注入，参数注入，属性注入的方式进行读取，这种短时需要使用配置的场景可以使用Configuration,同时也有符合长时间，高频率使用配置场景的强类型对象，我们可以通过瞬时，范围，全局三种方法进行调用强类型的对象（事先需要声明对象），并且Furion框架也提供了实时监控配置更新的方法，可以进行配置热更新的操作

- 单独读取App.Configuration["ConnectionStrings:Default"]

- 强类型（选项读取）：比如数据库配置，Mysql配置，Redis配置都需要在全局读取的

  注册：builder.Services.AddConfigurableOptions<ConnectionStringsOptions>()

  读取：App.GetOptions<ConnectionStringsOptions>().Default（如果是在startup中用App.GetConfig<ConnectionStringOptions>("ConnectionStrings", true).Default进行读取）



**动态WebApi**：还有动态API的功能，这个功能太实用了，我们可以直接在服务层里面写控制器的方法，只需要继承控制器的父类就可以了，并且我们可以配置不同方法的版本信息，可以在方法的特性中新增[HttpGet],[HttpPost]等特性，Furion框架会帮我们自动生成相关的两个方法（感觉很鸡肋），并且可以通过V1_1,V1_2等标签去生命方法的历史版本，还可以对Swagger方法进行分类型，比如实际工作中我们可以将服务类里面的方法进行分类，更新数据的放在一起，删除数据的放在一起，用标签进行管理，也可以按照业务进行分类，在Swagger右上角对类型进行选择，并且Furion框架会根据我们的方法自动生成对应的控制器（平常测试再也不用去写接口还有控制器了），自动生成，并且他还有对应的生成规范，比如方法名称是以Get开头，那么会生成HttpGet请求方法，以Post,Add等前缀词开头会自动标签[HttpPost]的方法等，这些配置是可以自定义配置与更新的，使用上也很方便

**HttpContext：**还有读取HttpContext对象，原来在ASP.NET的版本中这个对象是一个静态的对象，可以直接的网页/WebApi的环境直接读取，但是在ASP.NET COR的版本中HttpContext对象需要通过注入IHttpContextAccessor对象来进行读取，我们在HttpContext中可以读取到用户请求的参数，用户请求的表单（Form）,Session信息，Response(返回的响应信息)，Request（请求的Token，Hearder，浏览器信息等）信息，有了这个对象就可以操控整个网页的对象。



**筛选器/拦截器/过滤器/AOP**: 有授权鉴权过滤器，资源过滤器，ActionFilter(行为过滤器)，异常过滤器，结果过滤器

- 授权鉴权过滤器：看看用户的访问有没有权限
- 资源过滤器：看看用户的请求能不能请求缓存
- ActionFilter：记录用户操作日志
- 异常过滤器：记录应用程序的异常信息（根据指定的格式保存到本地文件或者入库）
- 结果过滤器：对返回的结果进行格式化，比如日期，金额，属性大小写



**审计日志**: 这里只是一个小例子后面有更加详细的内容

- 是保存在数据库还是保存在服务器文件？

 在ActionFilter中拿到请求的日志后，我们可以拿它写入到本地日志也可以写入到指定的数据库中

- 为什么要用这个审计日志？

**感觉这里定义了一个规范（需要哪些属性，还有日志格式）**，使用的是ActionFilter过滤器进行实现，读取登录用户姓名，请求的参数，操作的模块信息

- 我们现有的审计日志是什么？

没有日志，只有k8s上面的执行日志



**中间件**：中间件是Startup中Configure方法中的第一个参数ApplicationBuilder app,这个Builder就是管理中间件的管道，请求进入Http管道后会进入ConfigureServices管道，然后再进入中间件管道，中间件中的管道内容是对所用的Http请求有效的，比如通过app.Run() 可以拦截所有的Http请求到这个管道，然后后面的管道就会短路，也可以通过app.use对指定的控制器管道进行拦截（实现限流，流量转发的功能），同时我们也可以定制自己的中间件（可以看上面的Startup中Configure目录的内容）



**Vue/React/Angular 接口代理**: 前后端分离的场景下面，如果后端的请求接口发生变化，或者返回值打乱了原有的返回逻辑，需要重新写文档与前端进行交流，但是使用Swagger-Editor可以自动生成前端请求接口代码，import引用代码等代码，所以只要后端发生改动前端可以直接拷贝Swagger-Editor生成的代码到前端项目中直接使用不需要再互相扯皮。。。我感觉这样有以下需要注意：

- **旧项目的逻辑代码覆盖**:比如如果我现在想在我的项目里面使用这种模式，不可能把整个项目的代码调用方式进行覆盖吧？不过我想到这个工具可以单独执行我这个模块的代码进行覆盖，我只生成与这个前端对应的模块代码就可以了，这样也行，感觉还是挺不错的





**规范化接口文档**: 通过分组（模块），分版本（V1,V2）的方式管理Swagger接口的展示。我小看了这章节的内容，上面我讲的两点规范是Swagger的功能，还有一些通过的接口规范比如接口的命名（Get，Post,Query）这些通用的规范命名方式 ，还有接口请求的路径，比如Skums就是新的后端模块，或者Walmart是沃尔玛模块，请求的Api地址（顶部的控制器），可能就会生成不同的请求路径/skums/product/getProduct, /Walmart/xxxx/xxxx,还有就是返回状态码，返回状态信息，返回具体数据（这里的状态码可以是业务的状态码与HTTP的状态做区分，也可以相同的逻辑定义（比如返回2xx，就是请求成功,4xx就是请求失败等等，业务的内部逻辑与http的逻辑也可以做一定的区分），http状态与业务状态共存），Furion的Swaggre还有以下四种优点：

- **权限管理**：可以通过配置登录账户，密码的方式提供给人进行登录（安全保障）
- **性能监测**：`MiniProfiler` 第三方性能组件，查看性能、异常堆栈、数据库操作等信息
- **授权进行记住** ：token请求的时候需要提交header，token等请求参数,swaggre也可以支持
- **规范化返回结果** ：默认有返回操作是否成功,业务状态码，业务提示，数据的方法，不需要单独实现OperationOutput方法,直接用YourRESTfulResultProvider.OnSucceeded等方法



**友好异常处理**:

- **直接抛出Throw new exception("产生异常")与现有Furion框架的区别是什么？**

  原有的提示异常主动Throw new exception("产生异常")，与Furion :Oops.Oh异常其实是一样的效果，不过Furion框架有以下优点：

  **状态码管理**:可以将异常通过统一的状态码进行管理，比如必须为数字，必须身份证号码等逻辑，通过状态码提示的异常返回的异常提示都是一致的，并且支持方法入参，也可以先声明状态码后续阶段通过配置的方式配置到系统中

  

- **现在抛出的异常不友好吗？ （没有捕获到的异常，不优化，通过全局异常捕获解决）**

  有些没有处理的异常，会抛出英文代码（甚至在winform的情况下会提示异常的堆栈信息，全局异常捕获可以解决这个问题，设置一个默认的异常提示：系统异常，请联系管理员）

  try catch finally  不推荐（提升代码阅读难度，打乱业务逻辑）

- **解决的问题是什么？**

  状态码管理

  异常提示后期自定义配置

  自动重试异常方法（这个牛Retry）

  异常全局捕获（解决没有捕获的异常http状态码，堆栈码的问题）

  默认异常页面（鸡肋）

  全局异常捕获（鸡肋）



**【数据校验】**： 数据校验就是通过绑定实体类添加校验逻辑，比如校验数字，校验字符长度，校验是否是身份证号码，校验是否是电话（甚至可以根据业务逻辑，自定义校验逻辑），校验逻辑可以达到简化实际代码逻辑，实现一处注册，到处使用。Furion的数据校验功能与原有的net core的校验逻辑不冲突，**百小僧老师在原有的逻辑之上扩展了校验的范围与逻辑**，并且可以支持自定义校验逻辑（不然百小僧的校验逻辑是怎么扩展的。。。） ；刚才我又去看了一下关于数据校验的内容，内心无比的震撼，我下面表达一下。原有官方MVC的校验只支持在控制器方法与实体类中进行校验，但是官方提供的特性只能在Control控制器中进行使用方法参数的校验，但是**Furion的校验与自定义校验，可以在任何类的任何方法中进行使用**



**SaaS 多租户** ：公司的系统就是通过Sass多租户的方式，通过后台的权限系统进行控制不同部门，不同角色的权限管理，不同用户，不同部门登录的时候有不同的权限，相互的数据之间是隔离的。

看完整个文档内容后，我推翻上面的结论，这个Saas 多租户模式是用产品化的系统通过互联网向用户提供服务使用的，比如易仓系统 ，中大的Wiki系统，佳禾的项目管理系统。他们就是用这种模式，我也有一种想法，多租户是多个用户使用，我们公司不同部门，不同角色也是不同用户，我们用部门ID进行区分，那么也属于多租户的最后一种租户模式：

- **共享数据库**：多租户模式是用户租赁某个公司现有的产品进行使用，通过互联网进行租赁一样，根据使用的需求决定使用系统的模式，**使用公共的数据库但是使用单独的用户进行访问数据（每个公司一个自己的sa数据库用户）**

- **单独租赁数据库使用**：单独部署一个数据库使用，成本最高 

-  **共享数据库2** ：使用相同的数据库，相同的用户sa，但是用单独的用户编号进行数据隔离（公司目前的模式，不管，我就是这样理解的）

**多租户大白话**：在一条服务器上面部署了一个应用程序，然后它被多个租户（用户）使用

![1672821561713](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1672821561713.png)

#### 分布式缓存

 **为什么需要用缓存？**

- **降低数据库压力**：将常用的数据读取放到缓存当中，可以减少很多数据库的压力

- **提高程序的并发** ：缓存是从内存中读取的，吞吐要比关系型数据库要好，并发也更加大，同样的并发去访问数据库可能会把数据库搞奔溃（比如大批量推送数据的时候，Mysql数据库就经常阻塞），mysql 4 核8G单机并发一秒2000QPS可能就是巅峰了，但是redis单机的QPS可以达到1秒10万

  

  **用了缓存之后会有什么不良后果？**

  - 缓存与数据库双写不一致
  - **缓存雪崩**：缓存在同一个时间点全部失效，造成数据库压力冲击过大，设置缓存的失效时间点要分散，可以设置随机的时间，不要集中在某个时间同时生效，同时生效。
  - **缓存失效（击穿）**： 某个热点的缓存在一个时间突然失效，刚好那个时间节点是访问的高发期，一秒钟可能上万QPS造成数据库CPU压力过大，热点缓存可以分散到多个缓存，设置不同的时间段才失效
  - **缓存穿透**：查询的数据再数据库和Redis中都不存在（也有可能是恶意查询攻击），然后用户一直查，反复查，QPS一高的话，那么数据库就容易芭比Q. 可以在查询的时候通过缺省值进行限制，比如查询了三次都没有结果，那么下一次的查询不访问数据库，直接返回空值

  [Redis 缓存击穿（失效）、缓存穿透、缓存雪崩怎么解决？](https://www.51cto.com/article/703396.html)

  

**内存缓存**：存储在服务器本地的内容，可以通过注入 `IMemoryCache` 进行操作内存，但是服务器内存是及其昂贵的，并且不适合分布式服务的读取。

**分布式缓存：**

- **Redis缓存**：内存数据库，吞吐量大，延迟低，性能高，可持久化，缺点只能存储字符类型

- **Sql Server缓存**：微软官方支持的数据库（其实就是建立了一个专门的表进行存储缓存数据，我不太明白这种和直接读取表有啥区别。。。可能微软读取的API中做了优化），关系型数据库读取的性能要比Redis性能低

- **Ncache 缓存**：不支持持久化，关机数据就挂，并且是收费的，优点是可以保存图片，视频文件,适合NET与Java程序使用，性能比Redis更好

  

#### 依赖注入/控制反转

**依赖注入**：就是对象的注册不是由类本身进行注册，而是依赖于外部的需求进行注入，动态灵活

- **构造函数注入**：**推荐的注入方式**，一看构造函数就知道这个类依赖哪些服务，缺点是方法参数太臃肿了，并且会破坏初始的无参构造函数

- **方法注入**：在控制器中的Aation中进行注入，使用优点是根据方法的使用进行注入，缺点是一旦新增或者改动了注入对象，其他所有引用到方法的位置都需要重新编码

  ![1675311912719](Vue进步.assets/1675311912719.png)

- **属性注入**：**Asp Net Core默认是不支持属性注入的**，我们可以用Autofac进行实现，通过get,set属性的方式进行注入，如果构造函数注入与属性注入一起使用，可能会显得引用到的服务不显得清晰明了

  

  

**控制反转**：程序本身的行为依赖于抽象的接口，实际执行的行为是可以根据注入对象的不同，反转为不同的行为。

**对象的生命周期**

- **瞬时生命周期（Transient）**： 每次注入的对象都是获取的新的对象，Web服务中的同一次请求中，声明的的多个对象都是属于不同的对象（**可以通过&符号或者GetHashCode()方法进行判断**），通常编写的服务都是使用的瞬时生命周期

- **作用域生命周期（Scoped）**：在一个作用域里面使用的都是同一个对象，在Web请求中一个请求就是一个作用域，在这个作用域里面生命多个Scoped的对象，都是属于同一个对象

- **单例生命周期（Singleton ）**：在Services里面注入Singleton对象后，程序在所有位置的引用都是使用同一个对象，优点是节省性能，缺点是不能用于并发（同一个对象在一个时间内，只能被一个线程引用）

  

  **Tips:汤建华-朝夕：**其实。把serviceCollection  serviceProvider  serviceScope  ServiceDescriptor 这几个理解了。.net 中的容器注入应该就明白了 

### 框架索引

![img](https://furion.baiqian.ltd/img/functions.png)



## VMware 

局域网访问虚拟机设备

[手把手教你设置局域网访问虚拟机内服务器](https://cloud.tencent.com/developer/article/1599108)



## 压力测试

Fio压力测试工具：  压测  iops 和吞吐量

## 数据传输协议

| 名词   | 特点                                                         |
| ------ | ------------------------------------------------------------ |
| RPC    | 远程过程调用（分布式、微服务间的方法调用）                   |
| HTTP   | 无状态，每次请求都要发送一个request，服务器响应之后就断掉（http header中的keep-alive指的是tcp） |
| TCP    | 面向连接，三次握手保证通信可靠                               |
| UDP    | 非面向连接，不可靠，速度快（可以手动对数据收发进行验证，IM系统多采用，QQ） |
| socket | TCP协议的接口实现，面向传输层进行网络编程                    |

[扩展阅读](https://www.jianshu.com/p/959030de7f1c)

## 架构

无论技术如何演进变化，无论我们使用何种架构和技术，初心不可变，我们最终的目的是为了**简单高效的实现业务需求，以需求为核心找到适合自己和企业的技术才是最明智的选择**

**软件架构的目标：**第一：加快软件发布；第二，减少整个软件生命周期（设计，实现，持续迭代、线上发布维护等）中的资源投入，包括人力资源、软硬件资源。**第三：使用合适的规范，工具包提高的效率**

权限认证：idnetityserver4

日志：log4net,log

网关：consul,coelet

文件服务器：七牛云，本地服务器

数据库：mysql

缓存数据库：redis

文件搜索：

消息队列: rabbitMQ

微服务架构图1：

![img](https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fwww.pianshen.com%2Fimages%2F316%2F87bbe9003239e25a9e714eb51807203c.png&refer=http%3A%2F%2Fwww.pianshen.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1658297836&t=7c3dafc015b1532b058158392d2df033)

微服务架构图2：

![img](https://gimg2.baidu.com/image_search/src=http%3A%2F%2Fprocesson.com%2Fchart_image%2Fid%2F60a5dca6e401fd7ce8d670cf.png&refer=http%3A%2F%2Fprocesson.com&app=2002&size=f9999,10000&q=a80&n=0&g=0n&fmt=auto?sec=1658297836&t=41290e830797967d1d45ac3fad099df4)

 **低代码平台：**微搭

## 腾讯云

**网关**：

![1666079832017](Vue进步.assets/1666079832017.png)

新建pod节点的技巧：

**pod数量：**指的就是后端服务的数量，比如说我推送易仓的服务容器开了6个就是6个pod数量，开了四个就是4个pod数量

**limit:**是最大使用资源上限，

**request:**是最小使用资源上限，点击确定马上会占用的节点

![1655447720688](Vue进步.assets/1655447720688.png)

**查看剩余的内存资源：**

确认服务属于哪个节点:（点击单独的子服务进去就可以看到）

![1655447843123](Vue进步.assets/1655447843123.png)

查看指定节点的资源： CPU与内存资源要留多点最少要留50%，如果节点重启是先申请新的一个pod，再关掉旧的 。所以崩溃的话，或者重启的话，资源是翻倍使用的，要做最坏的打算

![1655447824464](Vue进步.assets/1655447824464.png)

## Newtonsoft

**反序列化** （转换Json为实体对象）：  

```c#
var models=Newtonsoft.Json.JsonConvert.DeserializeObject<GetProductListResponse>(json)
```



## XMIND

先点击Tab,再点击Enter，再点击Space空格进行操作

![1642996001241](Vue进步.assets/1642996001241.png)

## AWS

https://www.landafu.com/29528.html   注册aws免费服务器

https://www.cnblogs.com/miwujun/p/14106525.html)

## MySql

### 知识体系

感觉锁和事务的内容比较重点，然后大表优化的内容，高性能的内容

https://learn.lianglianglee.com/%e4%b8%93%e6%a0%8f/MySQL%e5%ae%9e%e6%88%98%e5%ae%9d%e5%85%b8/27%20%20%e5%88%86%e5%b8%83%e5%bc%8f%e4%ba%8b%e5%8a%a1%ef%bc%9a%e6%88%91%e4%bb%ac%e5%88%b0%e5%ba%95%e8%a6%81%e4%b8%8d%e8%a6%81%e4%bd%bf%e7%94%a8%202PC%ef%bc%9f.md 姜老师博客，捡到宝了

![img](Vue进步.assets/640.jpg)

### 实战内容

#### 如何优化Amazon在售表大字段大属性？

今天在线上做推码表的导出的时候，发现一个很诡异的问题，导出的时候Http请求一直提示An error occurred while sending the request.  The server returned an invalid or unrecognized response（发送请求时发生错误。服务器返回了无效或无法识别的响应）错误，又找不到问题点在哪里，于是开始深入挖掘，我单个单个导出发现问题在Amazon导出上，Amazon平台导出的时候一直提示异常，于是找到了数据库语句，查询发现问题在数据库瓶颈上面

![1692929748063](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1692929748063.png)

```mysql
# 这样一条语句，查询的结果才几百条，有索引 ，t_bi_amazon_listing_variation表数据400万，但是时间竟然要27秒，我感到震惊
SELECT `a`.`AccountId` AS `ShopId`, `a`.`AccountName` AS `ShopName`, `b`.`SellerSku` AS `PlatformSku`, `b`.`SkuCode` AS `ProductSku`, `b`.`Status`, `b`.`ListingDate` AS `PublishTime`
FROM `t_bi_amazon_listing` AS `a`
INNER JOIN `t_bi_amazon_listing_variation` AS `b` ON (((`a`.`Asin` = `b`.`ParentAsin`) ) AND ((`a`.`SellerSku` = `b`.`ParentSku`) )) AND (`a`.`AccountId` = `b`.`AccountId`)
WHERE (((`a`.`IsDeleted` = 0) AND (`b`.`IsDeleted` = 0)) AND `b`.`SkuCode` IN ('RCB29170001', 'RCB29170006', 'RCB29170011', 'RCB29170016', 'RCB29170021', 'RCB29170002', 'RCB29170007', 'RCB29170012', 'RCB29170017', 'RCB29170022', 'RCB29170003', 'RCB29170008', 'RCB29170013', 'RCB29170018', 'RCB29170023', 'RCB29170004', 'RCB29170009', 'RCB29170014', 'RCB29170019', 'RCB29170024', 'RCB29170005', 'RCB29170010', 'RCB29170015', 'RCB29170020', 'RCB29170025')) AND `b`.`ProCode` IN ('RCB2917', 'RBJ2219', 'RAA5427')

```

罪魁祸首是这两个大字段

![1692929952663](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1692929952663.png)

```mysql
#最好的方法当然是拆分表格，但是这样对程序的改动太大，所以我想了一下，看看能不能调整下这两个列的大小，可能这两个列的最大值根本就没有这么大，设置到这么大是完全冗余的
select Max(CHAR_LENGTH(StockSyncInfo)) as character_count, Max(LENGTH(StockSyncInfo)) as byte_count from t_bi_amazon_listing_variation    -- character_count 171 byte_count 419
果然字节最大字节长度419，字符长度才171，优化空间很大，我们设置的字符长度是2500，我感觉可以设置成1000，绰绰有余

#StockSyncInfo 同步成功:原数量:0,更新后数量:100,【库存数量大于等于1按比例100%补货，补货数量有小数时进位取整库存数量小于1按数量0补货】;佛山仓,二部分销大件云仓--美国虚拟仓,二部万邑通美南仓,二部万邑通美东仓,二部万邑通美西仓,二部谷仓美东仓176,二部谷仓美西仓176,二部美国自有仓,Walmart美国虚拟仓,二部Amazon美国虚拟仓	
#171 character_count	 
#419 byte_count  字节长度，就是varchar 2500中的字节，不过要注意不同数据库的字符集,如果是UTF-8中文字符通常占用 3 个字节，而英文字符通常占用 1 个字节
  select StockSyncInfo,CHAR_LENGTH(StockSyncInfo) AS character_count, LENGTH(StockSyncInfo) as byte_count from t_bi_amazon_listing_variation where  id=38295339


用同样的方式看了一下ImageUrl的长度，最大字符长度是580，于是我把ImageUrl，StockSyncInfo这两个属性都设置成了800长度
#更新后重置
alter table t_bi_amazon_listing_variation engine=innodb

#补充知识：
　　1、【varchar（可变长度）】： 【如果值的字节是419，那么即使设置成了2500字符，那么也是419字节空间】
　　可变长度的非 Unicode 数据，最长为 8,000 个字符。 数字，符号，字母，中文都是占用一个字符（2500字符中的一个字符，如果你设置成varchar(2) ，那么只能存储王锐，不能存储王锐锐）
　　
　　2、（注意Mysql并没有nvarchar数据类型）nvarchar（固定长度）： 【如果值的字节是419，那么设置成了2500字符，缺少了那2081会用英文空格补齐，】
　　可变长度 Unicode 数据，其最大长度为 4,000 字符。 
　　
  
 
```



#### 如何优化大表Walmart刊登表？

总结：到了一定的体量，分库分表分区表折腾人厉害，并且业务使用上不友好（无论是分区表还分物理表都需要强制用户选择或者后台查询一次逻辑得到分区列），mysql就是不支持过亿，过千万的数据库，有这么多的数据量就需要用搜索引擎了，要么就是在架构开始的时候就开始设计好表，分库负债均衡的架构，按业务的年月日进行分表的底层架构，好好设计表的属性数据类型，建立好索引，后期再来搞这些本末倒置，成本太高。

```mysql
# 1.看看表的基本信息，数据量有多少
select table_name,table_rows as '表行数',avg_row_length as '平均一行字节', concat(truncate(data_length/1024/1024,2),'MB') as '数据大小',concat(truncate(index_length/1024/1024,2),'MB') as '索引大小',concat(truncate(data_free/1024/1024,2),'MB') as '数据碎片/索引碎片'
from information_schema.tables where table_name in('t_bi_walmart_lister_variation')
# 2.看看表结构：46个属性，不算多，但是这下面7个大字段，会把表搞死的
desc  t_bi_walmart_lister_variation
 MainImageUrl	varchar(2000)	YES			
 SpecificJson	longtext	YES			
 DeleteUserName	longtext	YES			
 DeleterUserId	longtext	YES			
 IngestionErrors	longtext	YES			
 UnpublishedReasons	varchar(4000)	YES		
# 3.看看表索引 ，10个索引，除去IsDeleted这个属性区分度2个值，这两个属性PublishedStatus，属性IngestionStatus，其他的7个索引区分度都是挺高的
show index FROM t_bi_walmart_lister_variation
-- 区分度(性能太差用简易版本)：select 1318/20697875 区分度：0.0001，越接近1的越好，小于0.2不建议用索引，不过具体也需要分析业务的使用场景，才能决定取不取消
select COUNT(distinct IngestionStatus) ,ROUND(count(distinct IngestionStatus)/COUNT(IngestionStatus),4) from bailun_bltpro.t_bi_walmart_lister_variation 
-- 区分度(性能太差用简易版本)：select 1803/20697875 区分度：0.0001
select COUNT(distinct PublishedStatus) ,ROUND(count(distinct PublishedStatus)/COUNT(PublishedStatus),4) from bailun_bltpro.t_bi_walmart_lister_variation 
# 4.看看Procedure Analyse优化表结构：2000w数据执行300s,与desc 表结构与实际业务需要，调整Optimal_fieldtype列给出来的建议
  -- 执行返回中的Optimal_fieldtype列是mysql建议采用的列。
    select * from t_bi_walmart_lister_variation  PROCEDURE ANALYSE(1) 
# 5.清理索引碎片,数据碎片,清理完后可以再看看数据的大小

 alter table t_bi_walmart_lister_variation engine=innodb
  
# 5.优化建议：（推荐1-4方案成本低，局部起作用，6-8全库起作用，成本大，作用的范围很广）
 1.垂直分表：将7个大字拆分到另外的数据表中，也可以把其他不重要的属性字段拆分到另外的数据表中，在程序中做关联的查询，虽然成本也是比较大的
 2.根据索引的区分度，删除不使用的索引
 3.根据Procedure analyse优化表结构的建议，根据实际的业务进行调整表结构
 4.表分区，不过分区需要提供分区表的分区因子键，比如按照日期分区，那么需要用户在检索的时候强制选择日期条件，然后再进一步选择条件进行检索（作者经验：首先根据月份把上网记录表RANGE分区了12份，查询效率提高6倍左右，效果不明显，故：换id为HASH分区，分了64个分区，查询速度提升显著。问题解决）
  5.水平分表（按照行分表，Hash分表，范围分表，键值分表，列表分表都是这种分表），讲表根据年，月，日或者主键ID进行拆分，不过这个适合在设计初期就考虑到分表的情况，程序底层做了分表的数据，已经正在使用的阶段做分表的操作，成本应该是很高的，这个效果比分区表好，但是成功很高
  6.换分布式数据库，加钱就可以解决，阿里云POLARDB，OcenanBase 数据库性能比Mysql好5-10倍，100%兼容mysql，腾讯云DCDB数据库也是100%兼容mysql，成本比阿里云更低，可以尝试。
  7.冷数据拆分+用搜索引擎Es（ElasticSearch）（或者也可以加一个table_old表，把数据放进去，或者也可以加一张table_new，新建一个新业务表，达到分表的目的），整张表都放在Es，Es功能重，环境复杂是瞎扯，装一个java的jdk就可以玩了，功能一点儿也不重ES，这种需要确认业务方，看看热数据的时效性有多久，冷数据一概从正式环境移除，弄一份全量数据放es或其他地方，专门用于后续查询，Es几千万几亿那完全是小case，数据是放在硬盘+内存上面的，Es因为有倒排索引，他的检索效率就是比mysql要高
  8.用aws opensearch （亚马逊基于es开源分支开发的在线收费版本），不用自己搭环境，直接同步过去用，好用是好用，就是贵，没有预算的话，自己物理机搭一套，缺点是需要维护，维护成本高（装个机器，装个软件，搭个集群，别给自己设限，服务器部署开发全家桶，我上我也行），想省事，有效率一点就是用收费的
   
   
```



### 【索引的优化】

**type:** **需要记住一个性能的排序**：从好到坏依次是：null > system  > const > eq_ref > ref  > ref_or_null > index_merge > unique_subquery > index_subquery > range > index > ALL

在执行explain 后我们还可以执行：**show warnings**  （这两个单词单独执行）

会返回Level、Code、Message三个字段，当Code 为1003时会给出SQL的优化建议，可以作为sql语句优化的参考。

[辉哥推荐-对EXPLAN 的说明比较详细](https://www.cnblogs.com/yidengjiagou/p/16527697.html)

[mysql索引优化](https://mp.weixin.qq.com/s/l_pwKrJWdJUs52Vy3I21cw) :对EXPLAN 的进行说明
[索引优化1](https://mp.weixin.qq.com/s/wWtcFcDbSUIdmu6f3gv9CQ) ：Mysql的存储引擎
[索引优化2](https://mp.weixin.qq.com/s/065Qpfn7QoAjV0EagxfamA) ：B+树解析
[索引优化3](https://mp.weixin.qq.com/s/00WAd-x15YpvNHUPzwNrLw) ：B+树索引的高度
[索引优化4](https://mp.weixin.qq.com/s/NcxeYahWzvL1s-8-fxPJiA)
[索引优化5](https://mp.weixin.qq.com/s/lVZrppwAAIiHdXYA_Ji7BA)
[索引优化7](https://mp.weixin.qq.com/s/hf5ilxZF7iuCjL2EDzJswQ)
[MySQL连载：为什么使用B+树做索引](https://mp.weixin.qq.com/s?__biz=MzIyMzEyNzIyNg==&mid=2247484149&idx=1&sn=ac45e9778aff9c0ccce094f5f2a52df1&chksm=e823b120df543836cb9fdd5572bccffae4deb4caad3e2fbea102ec19d9431e609b9664822b68&cur_album_id=2151576090794852356&scene=189#wechat_redirect)

#### 用了EXPLAIN优化没有办法，怎么用SHOW PROFILES进一步分析？

`SHOW PROFILE` 是一个用于分析查询性能的 MySQL 命令。它可以用来显示有关特定查询执行期间所发生事件的详细信息，从而帮助您识别查询执行中的瓶颈和性能问题。

当您运行一个查询并使用 `SHOW PROFILES` 命令时，MySQL 会返回一系列与查询执行相关的事件，包括查询的各个阶段花费的时间。这些事件可以包括连接数据库、解析查询、优化查询、执行查询、发送结果等。**每个事件都有一个唯一的编号，您可以根据这些编号来进一步剖析每个事件的时间消耗。**

```mysql
-- 查看是否开启
show variables like "%profiling%";


-- 启用查询分析（会话级别，只对当前会话有作用，保存最近15次的操作结果）
SET profiling = 1;

-- 执行您的查询
select * from bailun_skums.t_lister_product_size_mapping  

-- 查看查询执行事件的详细信息
SHOW PROFILES;
Query_ID	Duration	Query
145	0.0874555	SHOW STATUS
146	0.087306	SHOW STATUS
147	0.0834845	select * from bailun_skums.t_lister_product_size_mapping

-- 我们可以拿到Query_ID进行进一步的分析这个查询的0.08秒的时间是在哪里消耗的
show profile for query 147;
Status	Duration  

starting	0.000357  查询开始执行的起始时间。
checking permissions	0.000235 验证用户权限的耗时。
Opening tables	0.000182  打开相关表的时间。
init	0.000157 初始化阶段的耗时。
System lock	0.000142 获取系统级锁的耗时。
optimizing	0.000136 查询优化阶段的耗时。
statistics	0.000135 收集统计信息的耗时。
preparing	0.000133 准备查询执行的耗时。
executing	0.000137 实际执行查询的耗时。
【Sending data	0.080463 将查询结果发送给客户端的耗时。】
end	0.000466  查询结束的时间。
query end	0.000335 查询结束的时间。
closing tables	0.000185 关闭相关表的时间。
freeing items	0.000218 释放查询中使用的内部项的时间。
cleaning up	0.000207  执行清理工作的时间。

通过分析每个查询步骤进行针对性的优化

#测试
select count(1) from t_product_property 耗费时间57s
#看看哪里最耗费时间	
show profile for query 166;
Status	Duration
...
Sending data	56.308916  基本时间都消耗在这里
...
优化的还是有分页，加索引的方式优化，感觉这个方法很鸡肋阿，不过看到其他人的案例看到也有| Copying to tmp table |1.916551 | 复制数据到临时表消耗的时间比较多，这个方法也有参考的意义
# 关闭查询分析
SET profiling = 0;

扩展信息：
可以在 SHOW PROFILE 命令中使用不同的属性，以获取有关查询执行的更多详细信息。在 MySQL 中，您可以通过在 SHOW PROFILE 命令后添加不同的属性，来查看不同方面的性能分析。以下是一些常用的属性：

ALL: 显示所有可用的性能属性。
BLOCK IO: 显示与块 I/O 操作相关的信息，包括读取和写入操作。
CONTEXT SWITCHES: 显示上下文切换的次数。
CPU: 显示与 CPU 使用相关的信息。
IPC: 显示与进程间通信 (IPC) 相关的信息。
MEMORY: 显示与内存使用相关的信息。
例如，您使用的命令 SHOW PROFILE CPU, BLOCK IO FOR QUERY query_id; 会显示查询的 CPU 使用情况和块 I/O 操作的信息。

这些属性可以帮助您更深入地了解查询执行期间发生的各种事件和资源使用情况。您可以根据需要选择适当的属性来分析性能问题，找出潜在的瓶颈，并进行相应的优化。不同的属性将提供不同类型的信息，帮助您全面评估查询的性能。

#案例：新增CUP与IO的使用信息，
show profile cpu,block io for query 238;
Status	Duration	CPU_user	CPU_system	Block_ops_in	Block_ops_out
starting	0.000346	0.000000	0.000000	0	0
checking permissions	0.000217	0.000000	0.000000	0	0
Opening tables	0.000144	0.000000	0.000000	0	0

#案例：显示所有属性列信息
show profile All for query 238;
Status	Duration	CPU_user	CPU_system	Context_voluntary	Context_involuntary	Block_ops_in	Block_ops_out	Messages_sent	Messages_received	Page_faults_major	Page_faults_minor	Swaps	Source_function	Source_file	Source_line

starting	0.000269	0.004000	0.000000	4	0	128	0	0	0	0	0	0			
checking permissions	0.000159	0.000000	0.000000	2	0	96	0	0	0	0	0	0	check_access	sql_authorization.cc	1023
Opening tables	0.000180	0.000000	0.000000	4	0	64	0	0	0	0	0	0	open_tables	sql_base.cc	5910

扩展阅读：
https://juejin.cn/post/6850418111930171399 MySQL性能分析show profiles详解
```



####  为什么复合索引查询时没有走覆盖索引？

> 1.查询复合索引中包含的列的时候，如果不加Where条件会走覆盖索引
>
> 2.查询符合索引中包含的列的数据，如果加上Where条件(即使条件是复合索引第一列)，并不总是会走覆盖索引，如果数据量小，Mysql用不到索引直接返回结果，数据量大的时候，我新建表，新建索引，然后构造100w数据，第一次查询'EXPLAIN SELECT id, name FROM example_table   WHERE id > 10000 AND id < 20000;' 走覆盖索引，后面删除索引，然后新建索引不会走覆盖索引，后面Drop表，重新新建表，新建索引，然后构造100w数据，第一次查询'EXPLAIN SELECT id, name FROM example_table   WHERE id > 10000 AND id < 20000;' 还是不会走覆盖索引
>
> 3.查询复合索引中包含的列的数据，如果加上Where条件强制FORCE INDEX可以让Mysql走覆盖索引，但是这不是最好的解释，推荐还是用Mysql执行计划自己的计划

```mysql
#创建表
CREATE TABLE example_table (
    id INT PRIMARY KEY,
    name VARCHAR(50),
    age INT,
    address VARCHAR(100)
);
#创建索引 idx_id_name
CREATE INDEX idx_id_name ON example_table (id, name);
#构造3条测试数据
INSERT INTO example_table (id, name, age, address) VALUES (1, 'John', 30, 'New York');
INSERT INTO example_table (id, name, age, address) VALUES (2, 'Alice', 25, 'Los Angeles');
INSERT INTO example_table (id, name, age, address) VALUES (3, 'Bob', 28, 'Chicago');

#type=const where条件中代表查询一次就找到结果了，extra=Using Where不会走覆盖索引
EXPLAIN SELECT id, name FROM example_table   WHERE id =1

#【走覆盖索引1】extra=Using Index 会走覆盖索引，如果不加条件，只查覆盖索引的列，但是实际查询中不可能不用条件
EXPLAIN SELECT id, name FROM example_table  

#type=range,key=primary,extra=Using where  ,没有走覆盖索引
SELECT id, name  FROM example_table WHERE id >1 and id<3  

#加入age索引,type:ref 正常,key=idx_age 正确,extra=null 很奇怪，没有走覆盖索引
CREATE INDEX idx_age ON example_table (age);
EXPLAIN SELECT id, name FROM example_table where age=10

数据量只有三条的情况下，不走覆盖索引，数据量太少了，Msql评估计划索引都不会去用，构造100w数据进行测试

#构造100w测试数据
CREATE PROCEDURE InsertExampleData()
BEGIN
    DECLARE counter INT DEFAULT 1;
    WHILE counter <= 1000000 DO
        INSERT INTO example_table (id, name, age, address) VALUES (counter, CONCAT('Name', counter), FLOOR(RAND()*100), CONCAT('Address', counter));
        SET counter = counter + 1;
    END WHILE;
END 
#执行
call InsertExampleData()

#【会走覆盖索引】extra=Using Index 会走覆盖索引
EXPLAIN SELECT id, name FROM example_table   WHERE id > 10000 AND id < 20000;

但是神奇的是我如果把索引idx_id_name删除掉，然后再重新新建，我再查询上面的语句他又不会走覆盖索引，即使我用了下面四个清理的语句，Mysql也不会走覆盖索引，总结就是：


#更新表格的统计信息。MySQL会重新计算表格的大小、行数、索引的分布等统计数据
ANALYZE TABLE example_table;
#清除优化器的缓存。MySQL的查询优化器可能会在查询过程中缓存执行计划或成本估计
FLUSH TABLE example_table;
#【可能会锁表】InnoDB存储引擎，这个命令会重新组织表格的物理存储，释放空间碎片，优化索引结构，并更新统计信息
OPTIMIZE TABLE example_table;
#重新计算表格的大小、行数、索引的分布等统计数据，并将其存储在数据字典
ANALYZE TABLE example_table;
```



#### 如何重置索引扫描行数让评估计划更加准确？



```mysql
ANALYZE TABLE tableName; #执行时会有写锁，支持读，不支持写。
运行ANALYZE TABLE命令，MySQL会更新表的统计信息，包括表的行数、索引统计信息等，以便查询优化器能够更准确地估计查询的成本和选择更合适的执行计划。

ANALYZE LOCAL TABLE tableName;#如果不想对表进行锁定，可以考虑使用 ANALYZE LOCAL TABLE 命令，它会更轻量级地更新统计信息，但也不会获得表级别的锁，因此在某些情况下可能不够准确。


```



####  怎么使用Show Index From查看表有哪些索引？

```mysql
# 查看表有哪些索引
SHOW INDEX from  t_infringe_keyword_scan
```



#### Mysql优化的经典顺序是怎么样的？

[原文链接](http://www.yaotu.net/biancheng/9971.html)

![](Vue进步.assets/b0c60c8dd8c941e1816e07946f40bf8a.png)



#### Mysql有哪九个存储引擎？

Mysql的存储引擎是表级别的，我们可以根据我们业务的特点动态对表使用不同的引擎，存储引擎是表级别的，并不是库级别的，比如日常的业务表需要保障事务，ACID数据一致性，日志表的话可以使用MyISAM存储引擎，因为读很快等。不过腾讯的数据库貌似只支持InnoDB引擎，其他的引擎是不支持的

| **引擎名称** |                                                              | **索引类型**            |
| ------------ | ------------------------------------------------------------ | ----------------------- |
| InnoDB       | 事务，行级锁，并发高                                         | BTREE（不支持Hash索引） |
| MyISAM       | 表锁（写入的时候锁表），并发低[MyIsAm官方解释](https://dev.mysql.com/doc/refman/5.7/en/myisam-storage-engine.html) | BTREE                   |

MySQL 支持多种不同的存储引擎，每个存储引擎都有不同的特性和适用场景。以下是一些常见的 MySQL 存储引擎：

1. **【InnoDB】**：InnoDB 是 MySQL 的默认存储引擎，它支持事务、ACID（原子性、一致性、隔离性和持久性）属性，提供了行级锁定，适用于要求数据完整性和并发性能的应用程序。
2. **MyISAM**：MyISAM 适用于读密集的应用，支持全文搜索索引和表级锁定，但不支持事务和行级锁定。在 MySQL 5.5 及更高版本中，MyISAM 已不再是默认存储引擎。
3. **Memory (Heap)**：Memory 存储引擎将表数据保存在内存中，适用于临时表或缓存数据。但由于依赖内存，数据不会持久保存，数据库重启后数据会丢失。
4. **Archive**：Archive 存储引擎用于存储大量历史数据，支持高度压缩，但不支持索引和实时查询。适合归档和数据仓库用途。
5. **CSV**：CSV 存储引擎将表数据保存为 CSV 格式的文件，适用于导入导出数据。但不支持索引，性能较低。
6. **Blackhole**：Blackhole 存储引擎会将写入的数据丢弃，适用于将数据写入日志或同步到其他服务器的场景。
7. **NDB (MySQL Cluster)**：NDB 存储引擎用于 MySQL 集群环境，支持分布式、高可用性、事务和实时查询。
8. **TokuDB**：TokuDB 存储引擎提供了高压缩率和高性能的数据存储，适用于大量写入和查询的场景。
9. **RocksDB**：RocksDB 存储引擎基于 Google 的 RocksDB 开源项目，适用于高度写入密集的场景。

不同存储引擎适合不同的用途，选择正确的存储引擎取决于应用程序的需求和性能特点。需要根据数据一致性、并发性能、数据量、读写比例等因素来选择合适的存储引擎。

#### Mysql如何根据表的特点设置Hash索引？

Mysql8.0数据库中InnoDB 在内部利用哈希索引来实现其自适应哈希索引功能，不能显试的设置Innodb引擎表新增Hash索引，Mysql5.7的版本可以。(扩展阅读：https://dev.mysql.com/doc/refman/8.0/en/innodb-introduction.html)

```mysql
本次建立Hash索引的引擎为：Memory引擎 ，Mysql8.0数据库中InnoDB 在内部利用哈希索引来实现其自适应哈希索引功能
# 查看表的存储引擎:Innodb
show create table student;
# 更改表的存储引擎（我发现只要把表的存储引擎一更改成Memory引擎，表的主键索引类型就自动变成了hash，因为Memory引擎，默认的存储引擎是hash引擎）
ALTER TABLE student ENGINE = Memory;
# 新增索引为hash索引
alter table student add index ix_hash_name using hash (NAME);  -- 也可以using btree
# 查看表的索引内容
show index from student
```



#### 我们设计Mysql一个表的时候要注意什么？

**设计表时要注意：**

- 表字段**避免null值出现**，null值很难查询优化且占用额外的索引空间，**推荐默认数字0代替null，如果是字符可以用空字符默认值**。并且程序的实体类中使用 {get;set;}=true; 默认值这种设置，保证程序没有空值的出
- **尽量使用INT而非BIGINT，如果非负则加上UNSIGNED（这样数值容量会扩大一倍）**，当然能使用TINYINT、SMALLINT、MEDIUM_INT更好。
- 使用枚举Enum数据类型或整数数字代替字符串类型 ，存储空间更小，检索更加快，像刊登状态，请求状态如果是中文的就可以用枚举类型进行读取
- 尽量使用TIMESTAMP而非DATETIME：TIMESTAMP 类型通常需要更少的存储空间，因为它只需要 4 个字节，而 DATETIME 需要 8 个字节。这可能在存储大量时间戳数据时节省存储空间，并且TIMESTAMP可以自动根据表的更改时间自动更新



**索引选择**

- 索引并不是越多越好，要根据查询有针对性的创建，考虑在WHERE和ORDER BY命令上涉及的列建立索引，可根据EXPLAIN来查看是否用了索引还是全表扫描


- 应尽量避免在WHERE子句中对字段进行NULL值判断，否则将导致引擎放弃使用索引而进行全表扫描


- 值分布很稀少的字段不适合建索引，例如"性别"这种只有两三个值的字段


- 使用多列索引时主意顺序和查询条件保持一致，同时删除不必要的单列索引


**简言之就是使用合适的数据类型，选择合适的索引**

**选择合适的数据类型**

（1）**使用可存下数据的最小的数据类型，整型 < date,time < char,varchar < blob**

（2）使用简单的数据类型，整型比字符处理开销更小，因为字符串的比较更复杂。如，int类型存储时间类型，bigint类型转ip函数

（3）使用合理的字段属性长度，固定长度的表会更快。使用enum、char而不是varchar

（4）尽可能使用not null定义字段

（5）**尽量少用text，非用不可最好分表**

**选择合适的索引列**

（1）查询频繁的列，在where，group by，order by，on从句中出现的列

（2）where条件中<，<=，=，>，>=，between，in，以及like 字符串+通配符（%）出现的列

（3）长度小的列，索引字段越小越好，因为数据库的存储单位是页，一页中能存下的数据越多越好

（4）离散度大（不同的值多）的列，放在联合索引前面。查看离散度，通过统计不同的列值来实现，count越大，离散程度越高：

#### FORMAT=JSON通过Json分享Explain 

查看优化结果Json文字版本，可以方便复制，粘贴给其他人

```mysql
# JSON格式是四种格式里面输出信息最详尽的格式，里面包含了执行的成本信息
EXPLAIN FORMAT=JSON select * from  temp

-- 返回结果
	{
  "query_block": {
    "select_id": 1,
    "cost_info": {
      "query_cost": "421859.29"
    },
    "table": {
      "table_name": "temp",
      "access_type": "ALL",
      "rows_examined_per_scan": 4186341,
      "rows_produced_per_join": 4186341,
      "filtered": "100.00",
      "cost_info": {
        "read_cost": "3225.19",
        "eval_cost": "418634.10",
        "prefix_cost": "421859.29",
        "data_read_per_join": "3G"   估计的扫描行数数据大小，可能比表大小更大
      },
      "used_columns": [  -- 本次查询的列，如果是*，那么这个显示的是表的所有列属性名称
        "ID",
        "Name"
      ]
    }
  }
}
```



#### 怎么通过开启Mysql慢查询来监控数据库状态？

```mysql
# 查看是否开启慢查询：show VARIABLES like '%系统变量名%' 可以看目前的变量状态(下面设置日志参数中的参数名称都可以用来查询)
show VARIABLES like 'slow_query%';  开启慢查询日志要更新数据库的配置，要登录mysql服务器进去更新配置文件，具体可以百度，或者看下面的：设置日志参数

-- slow_query_log	ON   #表示开启，on(开启)/off(关闭)
-- slow_query_log_file	/data1/mysql_root/data/20394/slow_query.log #慢查询保存位置

# 查看是表保存，还是文件保存日志
show VARIABLES like 'log_output%';

# 查看慢查询的个数，Slow_queries看这个属性，	48465 个慢查询
show global status like '%Slow_queries%'; 

# 设置日志参数(这些参数都是重启失效，如果需要永久有效需要在Ini文件中设置)
set global slow_query_log=on;  -- 临时开启慢查询日志，数据库重启后失效
set global long_query_time=3;  -- SQL语句执行时间超过3就被认为是慢查询语句，被记录到慢查询日志中(long_query_time的默认值为10秒，目前bailun_skums就是设置的10秒)
set global log_queries_not_using_indexes=on -- 把查询语句没有使用到索引的记录到慢查询日志中 on(开启)/off(关闭);
set global log_output='TABLE'; -- 设置日志输出格式为表还是文件 TABLE/FILE
#慢查询日志同时输出到文本和表中（担心影响性能，不敢开）
set global log_output='FILE,TABLE';
 -- 查询表里面的慢查询日志 ，表里面的sql是加密了的，需要使用CONVERT 进行转换
   select   CONVERT(sql_text USING utf8) as sql_text from mysql.slow_log

 # 慢查询日志分析工具：mysqldumpslow
 如果要手工分析日志，查找、分析SQL，显然是个体力活，MySQL提供了日志分析工具mysqldumpslow。
 mysqldumpslow --help  -- 查看工具的帮助信息(cmd命令下使用)
```



#### InnoDB引擎存储数据的内部结构是怎么样的？

表空间(tableSace):由段（segment），区(extent),页(page)组成

![img](https://img2018.cnblogs.com/blog/1465200/201903/1465200-20190313211422354-1043302132.png)



[InnoDB逻辑存储结构：段，区，页](https://www.cnblogs.com/GrimMjx/p/10526821.html)

#### 怎么根据数据量计算索引树的高度？

[计算索引树的高度](https://zhuanlan.zhihu.com/p/113311980)

#### 怎么知道Mysql查询扫描磁盘的时候扫描了多少磁盘页？

可以比对不同的语句的执行数值，来选择最优的查询语句

```mysql
# 查看查询的sql
select ProductSku from t_product_property where ProductSpu='NDB2576'  -- 查询指定的属性
select * from t_product_property where ProductSpu='NDB2576' -- 查询全部属性

上面两个不同的查询，对磁盘页的扫描是一致的，说明都是回表从聚簇索引中查询返回结果的，返回的字节流可能是少的，但是扫描的磁盘数是一样的，有没有办法优化？答：可以用覆盖索引，这样查就会少磁盘扫描了
# 查看需要扫描多少磁盘页（执行一句查询，执行一下下面的）
show status like 'last_query_cost'; 
12.499000,代表需要扫描12个磁盘页才能返回结果
```

示例中的结果表示优化器认为大概需要做**12.499个数据页的随机查找**才能完成上面的查询。这个结果是根据一些列的统计信息计算得来的，这些统计信息包括：每张表或者索引的页面个数、索引的基数、索引和数据行的长度、索引的分布情况等等。

  (1)它是作为比较各个查询之间的开销的一个依据。

  (2)它只能检测比较简单的查询开销，对于包含子查询和[union](https://so.csdn.net/so/search?q=union&spm=1001.2101.3001.7020)的查询是测试不出来的。

  (3)当我们执行查询的时候，MySQL会自动生成一个执行计划，也就是query  plan，而且通常有很多种不同的实现方式，它会选择最低的那一个，而这个cost值就是开销最低的那一个。

  (4)它对于比较我们的开销是非常有用的，特别是我们有好几种查询方式可选的时候。

#### 怎么使用FORCE key(PRIMARY) 强制走索引？

强制**FORCE key(PRIMARY)** 索引，所损耗的性能还不如全表扫描就不会走索引，而是直接全表扫描

```mysql
# 不走索引： IsDeleted 没有建立索引，不会走索引
SELECT COUNT(*) FROM `t_bi_real_listing` AS `u` WHERE `u`.`IsDeleted` = 0;
# 不一定走索引：可能FORCE key(PRIMARY) 的性能比全表扫描更大就不会走索引
SELECT COUNT(*) FROM `t_bi_real_listing` AS `u` FORCE key(PRIMARY) WHERE `u`.`IsDeleted` = 0;
# 肯定走索引：因为有 ID<10000000 条件
SELECT COUNT(*) FROM `t_bi_real_listing` AS `u` WHERE `u`.`IsDeleted` = 0 and ID<10000000;
```



#### mysql单表可以保存上亿的数据

今天写了一个统计sql，在一个近亿条数据的表上执行，200s都查不出结果。SQL如下：

```text
select customer,count(1) c
from upv_**
where created between "2015-07-06" and "2015-07-07"
group by customer 
having c > 20
order by c desc
```

执行explain,发现这个sql**扫描了8000W条记录到磁盘上。然后再进行筛选**。**type=index说明整个索引树都被扫描了**，效果显然不理想。

![img](https://pic3.zhimg.com/80/v2-bcf084c284fe88f5569884fb4992a536_1440w.webp)

拿着这个SQL去请教项目组的数据库大牛，仅仅加了一个force index，花了1s多就出结果了。修改后的SQL如下：

```text
select customer,count(1) c
from upv_** force index(idx_created)
where created between "2015-07-06" and "2015-07-07"
group by customer 
having c > 15
order by c desc
```

同样执行以下explain命令，这个SQL仅仅扫描了磁盘的110W行记录。也就是上一个SQL的80分之一。大家都知道，扫描磁盘是很耗时的IO操作，比内存操作慢几个数量级。**type=range,说明索引树仅仅被部分扫描**，要优于前面那个SQL.

![img](https://pic1.zhimg.com/80/v2-8022fb6f7a4b9a50c9aec6d75c2a0134_1440w.webp)

**除了磁盘扫描的行数的不一样，还有采用的索引的不同，上面的sql用的是联合索引，而下面的是单纯的created字段的索引。由于用的是created的索引，驱动条件就是created的区间，需要扫描的数据就立刻变小了，因为时间区间小。后面的SQL的key_len要远远小于前面的SQL，也就意味着要扫描的磁盘上的索引数据量要远远小于前面的SQL**。

第一个sql使用的是错误的索引，带来低效的查询。然后每条SQL只可能使用一个索引。通过上面的分析就可以发现，force index()指令可以指定本次查询使用哪个索引！一条sql只会用到一个索引，mysql优化器会计算出一个合适的索引，但是这个索引不一定是最好的。force index()指令可以避免MySql优化器用到了一个低效的索引。

#### 为什么索引建议建立数字类型，不建议建立字符类型

尽量使用数字型字段（如性别，男：1 女：2），若只含数值信息的字段尽量不要设计为字符型，这会降低查询和连接的性能，并会增加存储开销。**这是因为引擎在处理查询和连接时会 逐个比较字符串中每一个字符，而对于数字型而言只需要比较一次就够了**。并且如果是数字类型的话，索引树的非叶子节点上，可以在单个页（16k）上面保存更多的值，这样磁盘扫描一次就可以在单个页上面扫描尽可能多的行

#### Or走索引与不走索引的情况

**走索引：两边都有索引**

- where  左 or 右：or条件中的属性都建立了索引会走索引，建立复合索引也会走索引

**不走索引：有一个没有索引都不走索引**

- where  左 or 右：or条件中 无论左右那个条件建立索引，只要存在一个不建立索引，那么不走走索引，都是全表扫描

```mysql
# 不会走索引：建立了IX_ProductSpu索引列，但是productsku 没有建立索引  
select * from t_product_property_bak20221028   where ProductSpu='EDA0005'  or productsku='EDA00051103'
# 用Or走索引的情况：两边都有索引
1.两边都有索引：单独建立IX_ProductSku索引就会走索引（index会提示Merge合并索引） 
2.两边都有索引（复合）：建立IX_ProductSpu_ProductSku复合索引，也会走索引

```



#### 单独查询索引列不加条件也会用到索引

```mysql
# 会触发到索引，有IX_ProductSpu_ProductSku复合索引
select ProductSku,id,ProductSpu from t_product_property_bak20221028  
```

![1668418032535](Vue进步.assets/1668418032535.png)

#### 前缀索引（字符不长不使用，比如大于500）

```mysql
 --（也有可能是不够长影响的检索效率） 【前缀索引IX_ProductSku(11)性能没有这么好】 15s,12s,9s
 select ProductSku from t_product_property_bak20221028  force index(ix_productsku) where     ProductSku
  in (select * from ( select ProductSku from t_product_property_bak20221028 LIMIT 1000000) a );
 -- 【普通索引IX_ProductSku性能更好】12s,9s,6s
 select ProductSku from t_product_property_bak20221028 force index(ixproductsku2)   where      ProductSku
  in (select * from ( select ProductSku from t_product_property_bak20221028 LIMIT 1000000) a );
 
```

**当字段类型为字符串（varchar，text等）时,并且字符串很长的时候，可以使用前缀索引**。

有时候需要索引很长的字符串，这会让索引变得很大，查询时浪费大量的磁盘IO，影响查询效率。此时可以只将字符串的一部分前缀，建立索引，这样可以大大节约索引空间，从而提高索引效率。
语法：

![1668416328198](Vue进步.assets/1668416328198.png)

```mysql
#一般SKU都是11位，创建11位的前缀索引，创建11位的索引可以节省索引的磁盘空间，挺高索引检索的效率
create index ix_productsku on t_product_property_bak20221028(productsku(11));
#即使用长度大于11位的SKU去查询，也是会用到长度等于11位的前缀索引
select ProductSku,id,ProductSpu from t_product_property_bak20221028   where ProductSku='NDB25760001UK1232' 
```

**前缀索引会导致覆盖索引失效：**

![1668419648224](Vue进步.assets/1668419648224.png)

**前缀索引的前缀区分度：**
可以根据索引的选择性来决定，二选择性是指不重复的索引值和数据表的记录总数的比值，索引选择性越高则查询效率越高，唯一索引的选择性是1，这是最好的索引选择性，性能也是最好的。
选择性计算：

#### 回表查询

**通过二级索引查询到主键，再根据主键去聚集索引中查询一行的数据，称为回表查询**

在InnoDB中，二级索引叶子节点存储的不是记录的指针，而是主键的值。所以，对于二级索引的查询，会查两颗B+树，先在二级索引定位主键，再去主键索引的B+树上找到最终记录。由于主键索引构建的B+树其叶子节点存放的是所有的数据行，即整张表的数据，所以称之为回表查询。

#### B-Tree（高瘦树）

B-Tree树（B树，平衡二叉树）：1.树节点高度很高，单个节点保存一个数据（所以树的高度很高） 2.单个节点是保存数据，并没有保存左右节点的指针（所以不支持范围查找）

#### B+Tree（矮胖树）

B+Tree（B+树）：1.因为单个节点上面可以保存多个值（单个节点一般保存一个页的值），所以树的节点低（3层的树可以容纳1千万的值）2.并且单个叶子节点上面不只保存了值，还会保存左右两边

在Innodb中，默认定义的B+tree的节点大小是**16kb**，这就是说，假如一个key是8个字节，那么一个节点大约存2048个key，同时Innodb一次磁盘IO,读取的都是16kb的整数倍数据，这样，一次性能读取很多节点的内容出来。

 

[一篇神文，看完直说牛逼](https://zhuanlan.zhihu.com/p/59818056)

首先，数据库索引使用树来存储，因为**树的查询效率高**，而且**二叉查找树（B+Tree）还可以保持数据的有序**。

为了减少内存的占用，**数据库索引是存储在外部磁盘**上。**Mysql衡量查询效率的标准就是磁盘IO次数**。

如果利用二叉树作为索引结构，那么磁盘的IO次数和索引树的高度是相关的。为了提高查询效率，就需要减少磁盘IO数。为了**减少磁盘IO的次数，就需要尽量降低树的高度，需要把原来“瘦高”的树结构变的“矮胖”，树的每层的分叉越多越好，B树正好符合我们的要求**，这也是B树的特征之一。



B+Tree树每一个叶子节点（Leaf Node ）上面出了**存放索引键值(Name:张三,Age:20)和主键（ID:20）**的相关信息之外，B+Tree还存储了指向与该 Leaf Node 相邻的后一个 LeafNode 的指针信息，这主要是为了加快检索多个相邻 Leaf Node 的效率考虑（也就是范围扫描）

B树（B类树）的特点就是每层节点数目非常多，层数很少，目的就是为了减少磁盘IO次数。
B树是一种多路平衡查找树，它的每一个节点最多包含K个孩子（整数型值比较小，所以一个节点上面能放足够多的孩子），k称为B树的阶。k的大小取决于磁盘页的大小，
————————————————
版权声明：本文为CSDN博主「七慕卿」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/qq_32320337/article/details/109487925

**B+Tree 的优点**

- 每个节点中的叶子节点都保存了访问左右的指针信息，每个大节点（下图5-20是一个大节点）代表磁盘页大小的整数倍（许多操作系统中，页的大小通常为4K）

  ![img](https://pic2.zhimg.com/v2-fd647a4bec0ff4b2ae99c76f440927b5_r.jpg)

[B+Tree扩展阅读（叶分裂的内容讲的绝了，查找关键字：最后插入一个节点95）](https://zhuanlan.zhihu.com/p/59818056)

**B+Tree的页分裂**

- **左边有空位，但是位置不对**的拆分：

  插入一个节点70，在Index Page中查询后得知应该插入到50 - 70之间的叶子节点，但叶子节点已满，这时候就需要进行也分裂的操作，当前的叶子节点起点为50，所以根据中间值来拆分叶子节点，如下图所示。

  ![img](https://pic4.zhimg.com/v2-1df17e064a508b6943520a46568102df_r.jpg)

- **Leaf Page拆分**

  最后插入一个节点95，这时候Index Page和Leaf Page都满了，就需要做两次拆分，如下图所示

  ![img](https://pic4.zhimg.com/v2-1cb51fa0a8b059bad907ad44ba157233_r.jpg)

  Leaf Page与Index Page拆分

  拆分后最终形成了这样一颗树。

  ![img](https://pic3.zhimg.com/80/v2-4da3f307dfaa75274da6b290dc66b77e_1440w.webp)

- **左右没有空位的分裂**

**B+Tree的左旋**（也可以右旋）: 看图会发现，新增70到索引中，65与75位置已经是满的，但是左边30还是有一个位置，这个时候Mysql会触发左旋的操作，让50,55,60,65往左平移一位，然后把70放到原来65的右边，这样不会产生页分裂，提高了保存/新增/删除的性能

总结：当操作值的时候，主节点上没有节点了，但是兄弟节点上面有节点，就会触发左旋的操作（会不会右旋？会）



![img](https://pic2.zhimg.com/80/v2-a66ddaf688e1c21c4bc6aebf9512e23d_1440w.webp)

B+Tree为了保持平衡，对于新插入的值需要做大量的拆分页操作，而页的拆分需要I/O操作，为了尽可能的减少页的拆分操作，B+Tree也提供了类似于平衡二叉树的旋转功能。**当Leaf Page已满但其左右兄弟节点没有满的情况下，B+Tree并不急于去做拆分操作，而是将记录移到当前所在页的兄弟节点上**。通常情况下，左兄弟会被先检查用来做旋转操作。就比如上面第二个示例，当插入70的时候，并不会去做页拆分，而是左旋操作。

#### 平衡二叉树与二叉查找树

**红黑树为什么叫红黑树**：红黑树之所以叫红黑树，是因为红黑树的取名者喜欢红色和黑色这两种颜色。

**平衡二叉树**：见图一，左边叶子节点总是比右边的叶子节点小，如果要查找6的话，只需要找一次，如果要找5的话需要找三次，先找到根节点：6，然后5比六小，找左边，然后找到3,3比5小，找右边，找到5。查询性能虽然好，但是维护的成本也会很高，体现在新增，插入，删除的时候，平衡二叉树的叶子节点会做变动。

**二叉查找树**:见图2，支持顺序查询，但是数的节点太高了，不是平衡的二叉树，实际查询中IO遍历的次数太多

B+Tree: 感觉就像是结合了平衡二叉树的扁平，又支持了二叉查找树的范围查找



![1668482099356](Vue进步.assets/1668482099356.png)

#### 索引下推机制

索引下推是联合索引里面（IX_ProductSpu_ProductSku），第一个筛选条件找到的是范围的值(ProductSpu)，然后在这个筛选条件的基础上，继续利用联合索引上面的ProductSku属性值进行筛选符合要求的值，筛选一遍后再用ID进行回表查询

explain：extra列等于：**Using index condition代表使用到了索引下推**

```mysql
 # ProductSpu='NDB2576' 查询到结果后
  SELECT  * FROM bailun_skums.t_product_property_bak20221028  where ProductSpu like 'NDB%' and ProductSku='NDB25760001UK'


```

如果不使用索引下推的逻辑，比如两个单独的索引，name，phone分别建立了索引，当查询的时候,name是第一列,name等于张三筛选出了三条结果，这个时候需要回表查询phone的值，然后再到phone索引进行回表查询吗

[五分钟搞懂MySQL索引下推](https://blog.csdn.net/tuning_optmization/article/details/119945967)

#### where和order by的索引问题

​    mysql查询只使用一个索引，因此如果where子句中已经使用了索引的话，那么order by中的列是不会使用索引的。(**可以用联合索引解决这个问题**)

```mysql
 # 只能使用到IX_ProductSku的索引，使用不到IX_ProductSku的索引，只能建立复合索引（IX_ProductSpu_ProductSku）
 SELECT ProductSpu,ID,ProductSku FROM bailun_skums.t_product_property_bak20221028  where ProductSpu='NDB2576' order by ProductSku desc 

```

​    因此数据库默认排序可以符合要求的情况下不要使用排序操作；尽量不要包含多个列的排序，如果需要最好给这些列创建复合索引。

#### IS NULL与IS NOT NULL索引失效的情况

属性IS NOT NULL,  WHERE IS NOT NULL(只在select覆盖索引有效)/ IS NULL 索引都失效

属性IS NULL ,WHERE IS NULL(只在select覆盖索引有效) /IS NOT NULL 索引都失效

```mysql
 # 属性ProductSku 不允许为空，Is null 索引失效,Is not null 索引有效（覆盖索引的情况下有效）
 select ProductSku （换成*索引失效） from t_product_property_bak20221028   where     ProductSku IS   NULL 
 # 属性ProductNameEn 允许为空，Is null 索引生效（覆盖索引的情况下有效）,IS NOT NULL 索引失效
select ProductSku （换成*索引失效） from t_product_property_bak20221028   where     ProductNameEn IS   NULL 
# 属性不允许为空语句可以走索引例外的情况：使用 Or 
# 注意：productsku 属性还是不允许走索引，但是可以走属性ProductSpu的索引
select ProductSku from t_product_property_bak20221028   where ProductSpu='EDA0005'  or productsku is   null 

```



#### 覆盖索引(USING INDEX)

复合索引建立原则：

- 最窄的字段放在键的左边
- 尽量不超过5个属性值
- 前导列一定是使用最频繁的列
- 

复合索引的优点：

- 巧妙使用覆盖索引提升性能
- 单个索引可以变成多个索引，提高索引使用效率
- 索引的较少，提高了新增，更新和删除的性能
- 降低了磁盘空间

见第三条：**【使用复合索引巧妙使用覆盖索引提升性能】**

```mysql
# 使用二级索引: ID（主键）存在所有的索引当中，因为当select *  中的属性不存在索引当中的时候，Mysql回表查询的时候是使用主键回表查询不在索引中的列的值，如果仅仅是查询ProductSpu,ID那直接就是从索引中返回结果，走覆盖索引，不需要回表查询数据
 SELECT ProductSpu,ID FROM bailun_skums.t_product_property  where ProductSpu='NDB2576'
# 需要回表查询 ProductSku 这个属性需要回表查询（即使ProductSku列存在表t_product_property的索引当中，但是ProductSpu的索引中没有存储这个列的值）
 SELECT ProductSpu,ID,ProductSku FROM bailun_skums.t_product_property  where ProductSpu='NDB2576'
# 【使用复合索引巧妙使用覆盖索引提升性能】
 #建立了IX_PorductSpu_ProductSku 复合索引,直接使用覆盖索引返回值，不需要回表
 SELECT ProductSpu,ID,ProductSku FROM bailun_skums.t_product_property_bak20221028  where ProductSpu in ('NDB25761', 'NDB25762', 'NDB25763');
 #只建立了IX_ProductSpu 单个索引，需要回表查询ProductSku的属性值
 SELECT ProductSpu,ID,ProductSku FROM bailun_skums.t_product_property where  ProductSpu in  ('NDB25761', 'NDB25762', 'NDB25763');
```



#### 全文索引（FULLTEXT） ****

全文索引用在text,nvarchar,varchar等字符串的大批量数据的模糊查询，性能比like高N倍

[全文索引的使用](https://blog.csdn.net/tuning_optmization/article/details/119945967)

#### 提高分页查询效率的方法：延迟关联，书签加载

两个相同的表，通过主键进行关联后查询，性能就会提升 ，这可太牛了

[原文地址](https://blog.csdn.net/a6636656/article/details/124129367)

![1668073944258](Vue进步.assets/1668073944258.png)



**书签方式**
书签方式说白了就是找到limit第一个参数对应的主键值，再根据这个主键值再去过滤并limit
例如

```mysql
select * from table where id > (select * from table where type = 2 and level = 9 order by id asc limit 190289, 1) limit 10;
```

**扩展内容**

**使用合理的分页方式以提高分页效率**
使用合理的分页方式以提高分页效率 针对展现等分页需求，合适的分页方式能够提高分页的效率。

案例1：

```mysql
select * from t where thread_id = 10000 and deleted = 0 
   order by gmt_create asc limit 0, 15;
```

上述例子通过一次性根据过滤条件取出所有字段进行排序返回。数据访问开销=索引IO+索引全部记录结果对应的表数据IO。因此，该种写法越翻到后面执行效率越差，时间越长，尤其表数据量很大的时候。

适用场景：当中间结果集很小（10000行以下）或者查询条件复杂（指涉及多个不同查询字段或者多表连接）时适用。

案例2：

```mysql
select t.* from (select id from t where thread_id = 10000 and deleted = 0
   order by gmt_create asc limit 0, 15) a, t 
      where a.id = t.id;
```

上述例子必须满足t表主键是id列，且有覆盖索引**(thread_id, deleted, gmt_create)**。通过先根据过滤条件利用覆盖索引取出主键id进行排序，再进行join操作取出其他字段。数据访问开销=索引IO+索引分页后结果（例子中是15行）对应的表数据IO。因此，该写法每次翻页消耗的资源和时间都基本相同，就像翻第一页一样。

**适用场景：**当查询和排序字段（即where子句和order by子句涉及的字段）有对应覆盖索引时，且中间结果集很大的情况时适用。
————————————————
版权声明：本文为CSDN博主「一名陈序员」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/m0_56864757/article/details/125217302

#### 复合索引查询范围条件放在最后查询

```mysql
# 中间有范围查询会导致后面的索引列全部失效(按照name,age,pos，索引创建的顺序，age后面的POS会失效)
EXPLAIN SELECT * FROM staffs WHERE NAME = 'July' and age >22 and pos='manager' #（索引只会用到NAME 和AGE列）
————————————————
版权声明：本文为CSDN博主「柚子茶1990」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/a6636656/article/details/124129367
```



#### 检查全表的索引，看区分度，Null类型

```mysql
-- 看Null这个列，看表里面是不是有NULL值,【重复值很少的属性考虑剔除索引】
show index from t_product_property 
-- 根据指定的索引名称筛查
show index from t_product_property where key_name='IX_ProductType'
```

| 参数         | 说明                                                         |
| :----------- | :----------------------------------------------------------- |
| Table        | 表示创建索引的数据表名，这里是 tb_stu_info2 数据表。         |
| Non_unique   | 表示该索引是否是唯一索引。若不是唯一索引，则该列的值为 1；若是唯一索引，则该列的值为 0。 |
| Key_name     | 表示索引的名称。                                             |
| Seq_in_index | 表示该列在索引中的位置，如果索引是单列的，则该列的值为 1；如果索引是组合索引，则该列的值为每列在索引定义中的顺序。 |
| Column_name  | 表示定义索引的列字段。                                       |
| Collation    | **表示列以何种顺序存储在索引中。在 MySQL 中，升序显示值“A”（升序），若显示为 NULL，则表示无分类。** |
| Cardinality  | **索引中唯一值数目的估计值。基数根据被存储为整数的统计数据计数，所以即使对于小型表，该值也没有必要是精确的。基数越大，当进行联合时，MySQL 使用该索引的机会就越大。** |
| Sub_part     | 表示列中被编入索引的字符的数量。若列只是部分被编入索引，则该列的值为被编入索引的字符的数目；若整列被编入索引，则该列的值为 NULL。 |
| Packed       | 指示关键字如何被压缩。若没有被压缩，值为 NULL。              |
| Null         | **用于显示索引列中是否包含 NULL。若列含有 NULL，该列的值为 YES。若没有，则该列的值为 NO。** |
| Index_type   | 显示索引使用的类型和方法（BTREE、FULLTEXT、HASH、RTREE）。   |
| Comment      | 显示评注。                                                   |

#### Mysql优化器是怎么选择索引的

mysql选择索引是根据预估的扫描行数进行的，当使用的索引不按照我们的与其进行的时候有三种解决方案：

- **重写sql**：比如想让他走主键索引，可以加上id<10000000 ，对最终结果影响不大，但是可以走主键索引
- **强制索引**：force key(ix_name)
- **重置索引扫描行数**：analyze table tableName，可以用来重新统计索引信息(不要在上班的时候设置，会加上只读锁)

[mysql之优化器选择索引](https://blog.csdn.net/qq_41820066/article/details/127214633)

[MySQL之ANALYZE TABLE](https://blog.csdn.net/weixin_50518271/article/details/109498885)

#### 巧妙使用ID<1000000使用主键索引的方法

综合来看查询总行数，查询行数最优的性能方法是：**id<100000000**

目前在线上所有的条件上面都加入了**id<100000000**这个逻辑，这个查询逻辑**数据量越大对性能的提升约明显**，比如查询时间周期长，筛选平台下架，筛选在线产品的条件对性能的提升很大，原来查询速度要6-8秒，现在1-2秒就可以返回，但是缺点在于数据量小的情况，比如查询条件只有：NDB2576,原来是66-100ms毫秒返回结果，现在要300-500ms毫秒返回结果，对性能还是有一定的损耗。**总体来说推荐使用：id<100000000**

```mysql
 
 
 # 查询总行数
 -- 【0.846s】
 SELECT count(0) FROM t_product_property   WHERE id<100000000 and  ( `ISDELETED` = 0 )  AND  USERORGANIZATIONID IN ( 1,24,27,26,25,30,29,32,33,34,13,14,15,1001,1002,1003,1004,1005,1006,1008,1009,31)   ORDER BY `CREATIONTIME` DESC;
 
 -- 0.746s
  SELECT count(0) FROM t_product_property force key(PRIMARY)  WHERE      ( `ISDELETED` = 0 )  AND  USERORGANIZATIONID IN ( 1,24,27,26,25,30,29,32,33,34,13,14,15,1001,1002,1003,1004,1005,1006,1008,1009,31)   ORDER BY `CREATIONTIME` DESC;
 
 -- 1.964s
   SELECT count(0) FROM t_product_property   WHERE      ( `ISDELETED` = 0 )  AND  USERORGANIZATIONID IN ( 1,24,27,26,25,30,29,32,33,34,13,14,15,1001,1002,1003,1004,1005,1006,1008,1009,31)   ORDER BY `CREATIONTIME` DESC;
 
 # 查询数据
 --  【0.014s】
  SELECT * FROM t_product_property   WHERE id<100000000 and  ( `ISDELETED` = 0 )  AND  USERORGANIZATIONID IN ( 1,24,27,26,25,30,29,32,33,34,13,14,15,1001,1002,1003,1004,1005,1006,1008,1009,31)   ORDER BY `CREATIONTIME` DESC  LIMIT 0,20;
 --  0.886s
  SELECT * FROM t_product_property force key(PRIMARY)  WHERE      ( `ISDELETED` = 0 )  AND  USERORGANIZATIONID IN ( 1,24,27,26,25,30,29,32,33,34,13,14,15,1001,1002,1003,1004,1005,1006,1008,1009,31)   ORDER BY `CREATIONTIME` DESC LIMIT 0,20;
 
 --  0.014s
   SELECT * FROM t_product_property   WHERE      ( `ISDELETED` = 0 )  AND  USERORGANIZATIONID IN ( 1,24,27,26,25,30,29,32,33,34,13,14,15,1001,1002,1003,1004,1005,1006,1008,1009,31)   ORDER BY `CREATIONTIME` DESC  LIMIT 0,20;
```



#### 临时表关联不会使用索引

如果A表与B表进行关联，A或B表中存在一个临时表，那么A表或者B表的索引就会失效

```mysql
-- 易仓仓库 [这个不走索引：20秒]
		  	 update tmp_property  a   force index(IX_DefaulBuyWarehouseCode)
			left join (select dic_text,dic_text_En as dic_value from t_product_dictionary where dic_name='易仓仓库' and dic_status=0 ) ck on ck.dic_value=a.DefaulBuyWarehouseCode
	 set a.DefaulBuyWarehouseCodeCh=ck.dic_text;
-- 换一种写法 【走索引：1秒】
	 update tmp_property a inner join t_product_dictionary ck on    ck.dic_value=a.DefaulBuyWarehouseCode
	 	 set a.DefaulBuyWarehouseCodeCh=ck.dic_text
		 where ck.dic_name='易仓仓库' and ck.dic_status=0
```



#### 【禁用缓存，禁用某个索引】

```c#
//强制使用某个索引：
select * from table force index(idx_user);
//禁止使用某个索引：
select * from table ignore index(idx_user);
//禁用缓存(在测试时去除缓存的影响)： SQL_NO_CACHE
select SQL_NO_CACHE * from table;
//使用某个缓存 use index：使用某个索引（暂时不确定与 force index(idx_user)）有什么区别，放在此处做为补充知识内容，还是使用force index使用索引
select * from t_product_property use index(ix_ProductSku) where ProductSpu='NDB2576' AND ProductSku='NDB25760001UK' ;

```

- **USE INDEX** 在你查询语句中表名的后面，添加 USE INDEX 来提供希望 MySQL 去参考的索引列表，就可以让 MySQL 不再考虑其他可用的索引。例子: SELECT col1 FROM table USE INDEX (mod_time, name)...
- **IGNORE INDEX** 如果只是单纯的想让 MySQL **忽略一个或者多个索引**，可以使用 IGNORE INDEX 作为 Hint。例子: SELECT col1 FROM table IGNORE INDEX (priority) ...
- **FORCE INDEX** 为强制 MySQL 使用一个特定的索引，可在查询中使用FORCE INDEX 作为Hint。例子: SELECT col1 FROM table FORCE INDEX (mod_time) ...


[扩展阅读](https://zhuanlan.zhihu.com/p/59694085)

#### 【强制使用索引：Force  index失效的情况】

**结论**：堆积量太大  MySQL内部计算cost 走索引不合理（因为： 正常的数据量超过20%的全表数据对于走索引反而是更不好的）

**解决方案**：删除isdelete无效的索引

```mysql
#索引失效：因为IsDeleted` = 0 等于扫描全部的结果（group by 发现只有一个结果），但是IX_IsValid结果中又不会包含全部数据（IsDeleted` = 0是all data，这种情况走索引没有意义）
SELECT COUNT(*)
FROM `t_bi_real_listing` AS `u`
FORCE index(IX_IsValid)
WHERE `u`.`IsDeleted` = 0
#进一步分析原因
-- 1.看了大字段的情况，没有发现大字段
DESC t_bi_real_listing;
-- 2.看了属性的个数情况，53个属性，属性个数也不多，productProperty表有65个属性，不过字符型个数差不多,real在售表有28个，sku表有30个属性，具体占用多大我就没有看
-- 3.【看看表占用数据空间，与索引占用空间大小】,real的行数虽然查了60w，但是数据大小是基本一致的【所以优化的时候不能简单的看数据行数，需要看表结构，数据存储的内容（可能real在售有很多标题文字很大，保存满了空间），都会影响数据返回的速度】，索引占用情况倒是不意外，sku表建立了13个索引，但是real表只建立了6个索引
   -- t_bi_real_listing	813.98MB(747.98MB)	224.39MB(139.34MB)  81w数据 count(1) 0.9秒
   -- t_product_property	791.98MB	514.76MB  140w数据 count(1) 1.7秒
   -- 【下面的语句可以夸库查询表】
select table_name,concat(truncate(data_length/1024/1024,2),'MB') as data_size,concat(truncate(index_length/1024/1024,2),'MB') as index_size,concat(truncate(data_free/1024/1024,2),'MB') as '数据碎片/索引碎片'
from information_schema.tables where table_name in('t_bi_real_listing','t_product_property')
   -- 清理碎片
    alter table t_bi_real_listing engine=innodb
   -- 再次查看，表空间，索引占用空间大小，碎片的空间要看是不是频繁更新，写入数据，索引也是，是不是经常建立索引，然后再删除，可能碎片空间就比较多会影响性能
     表名                    数据空间清理前（数据空间清理后）  索引空间清理前（索引空间清理后）
     t_bi_real_listing  	813.98MB->(747.98MB)	    224.39MB(139.34MB)

-- 4.使用show warnings看看真实的执行语句，结果也是没有多少有用的信息
   -- 0.3秒 
     -- show warnings 前
    SELECT COUNT(*) 
    FROM `t_bi_real_listing` AS `u` force index (PRIMARY)   where `u`.`IsDeleted` = 0;
     -- show warnings 后（除了标志性的count(*)转换成了count(0) 其他的条件不变）
     select count(0) AS `COUNT(*)` from `bailun_bltpro`.`t_bi_real_listing` `u` FORCE INDEX (PRIMARY) where (`bailun_bltpro`.`u`.`IsDeleted` = 0)

  -- 0.6秒 （不进行show warnings，结果和上面一样）
    select count(*)
    from bailun_skums.t_product_property force index (PRIMARY);
-- 5.看看索引的区分度，可以干掉区分度低的
    -- 看所有的索引情况，cardinality 列代表索引的重复数，如果只有1~10 可能要考虑废除索引（结合实际的业务场景）
   show index from t_bi_real_listing 
   -- 根据指定的索引名称筛查
show index from t_product_property where key_name='ix_IsValid'
-- 6.使用force index (primary) 或者id<100000000强制使用主键索引扫描，提升性能 

-- 7.查看索引的使用情况
	-- 1. 查询冗余索引
    select * from sys.schema_redundant_indexes where table_name='t_ec_product';
    -- 2. 查询未使用过的索引
    select * from sys.schema_unused_indexes; 
    -- 3. 查询索引的使用情况，如果查询不出数据，看看是否开了pfs之类的，sys里面的数据时从其他元数据库里面来的
    select index_name,rows_selected,rows_inserted,rows_updated,rows_deleted
    from sys.schema_index_statistics where table_schema='t_ec_product' ;
-- 8.使用Procedure Analyse优化表数据结构，找到更好的数据类型替换现有的数据类型
   -- 执行返回中的Optimal_fieldtype列是mysql建议采用的列。
    select * from t_product_property_bak20221028  PROCEDURE ANALYSE(1) 
-- 9.重置索引扫描行数：
analyze table tableName -- (不要在上班的时候设置，会加上读锁，并且无法对表进行更新)可以用来重新统计索引信息,比如show index中的散列值cardinality，初始化查询优化器explain执行优化时将数据的行数与散列数进行对比，如果差异太大可能会不走索引，初始化了的话就不会出现这个问题
```

**执行计划：**

![1668734341671](Vue进步.assets/1668734341671.png)

#### 索引合并的情况：自动合并为复合索引

**索引合并(index merge)**：对多个索引分别进行条件扫描，然后将它们各自的结果进行合并(intersect/union)。一般用OR会用到，如果是AND条件，考虑建立复合索引。EXPLAIN显示的索引类型会显示index_merge，EXTRA会显示具体的合并算法和用到的索引

```mysql
explain   select  productsku,productspu from t_product_property where ProductSpu='NDB2576' or productsku='NDB25760001UK'

```

![1667362169738](Vue进步.assets/1667362169738.png)

#### 使用show warnings查看评估计划实际执行的语句

```mysql
explain   select  * from t_product where ProductSpu='NDB2576'; #注意分号
show warnings #查看实际运行sql，在 结果2 列表中有sql值



```

评估器实际执行SQL:

![1667361517456](Vue进步.assets/1667361517456.png)

#### 表索引占用大小查询（可查全库）

```mysql
# 查看单个表每个索引占用的大小
SELECT

*

FROM

mysql.innodb_index_stats

WHERE

table_name = 't_product_property'

AND database_name = 'bailun_skums'

AND stat_description LIKE 'Number of pages in the index'

GROUP BY

table_name, index_name 
order by stat_value desc 

# 查看单个表占用索引的大小
select table_name,concat(truncate(data_length/1024/1024,2),'MB') as data_size,concat(truncate(index_length/1024/1024,2),'MB') as index_size

from information_schema.tables where table_name='t_product_property'

order by index_size desc;
```



#### 全局索引扫描与当前会话索引扫描

```
show status like 'Handler_read%' 
show global status like 'Handler_read%';
```

[扩展阅读](https://blog.csdn.net/Leon_Jinhai_Sun/article/details/121527970)

#### 主键索引对count(*)的影响

默认用主键查询速度为：0.7s，用查：IX_ProductSpu索引2s，走主键索引性能最好

```mysql
# 走其他索引count(*)很慢：只有一个索引的情况下,count(1),count(*) 选择的是IX_ProductSpu索引,这时候需要2s出结果
# 走主键索引count(*)很快： 如果把IX_ProductSpu删除，那么走的key是主键索引，0.7s出现结果
#【count(*)默认是找一个长度最小的索引做为选择，比如t_product_property会选择ix_isdeleted索引做为扫描，删除ix_Isdeleted索引后，会选择IX_SaleStatus索引】
SELECT COUNT(*) FROM t_product_property_bak20221028  FORCE INDEX(PRIMARY)

```



#### 大字段对count(*)的影响

默认用主键查询速度为：0.7s，用空的大字段查：2s，用有值的大字段查：10s ，有值的大字段对性能的影响有十几倍

```mysql

# 【大字段会影响count(*)的性能，即使值是空的，性能差2倍】：加了一个5000的大字段，值是空的，count(*)，慢了700ms至1秒的性能，要花2s出结果，所以只要字段大小在那里，他就会影响count(*)的性能
# 大字段有值对性能影响更大，性能差10倍： 把大字段的值赋值，看看对性能的影响情况,赋值后，查询性能直线下降：查询速度达到：10s，差了10倍的性能，把字段大小由5000缩小至50，查询性能又到了0.7s

SELECT COUNT(*) FROM t_product_property_bak20221028  FORCE INDEX(PRIMARY)

```



#### 大字段对索引的影响

3500长度的大字段SpecificJson，删除后性能提升了12倍

```mysql
# SpecificJson 3500长度
# 未删除前count(1)性能： 9s,9s,9s 查询三次
# 删除后count(1)性能：0.8,0.8,0.8
select count(1) from t_product_property_bak20221028  #1477496

# 保险的话清理后用语句整理一下空间碎片
alter table t_product_property_bak20221028 engine=innodb
```



#### 复合索引与单列索引的场景

**复合索引建立场景：**

**覆盖索引**：当查询的列比较少的时候，该语句存在使用频率高，并且回表查询性能不好的情况下，可以通过冗余索引的方式，把查询的列冗余到索引当中（并不一定要在where条件中，但是索引要结合where条件符合最左匹配原则）

**复合查询where**：当查询的where有多个条件的时候，一般可以建立复合索引，这样这个语句在一个索引里面就可以检索出结果（而不是检索完索引1，再拿索引1返回的结果集放到索引2里面去找），不用mysql将单个索引的结果查询出来后，然后再拿这个结果集到另一个索引中去扫描索引树检索数据，综合情况要看业务上面的逻辑是如何（有个观点是：**对同一个列建单个索引与复合索引是没有关系的**，因为每个索引之间是单独管理存储的，查询的时候使用哪个索引取决于你的Sql，以及哪个索引是最优的（由优化器算法决定）），有个技巧是业务允许的情况下：**将区分度高的列放到最左匹配原则的最前面**（比如主键就是区分度最高的，它不会重复，SKU的区分度就比SPU的区分度要高，不过最终要结合业务的使用情况来决定）

Tips:而且复合索引可以提高写入，新增的性能，因为每新增数据后，对应的索引页mysql都会重新排列他们的位置，如果表上面的索引非常多的话，一定程度上也会影响mysql写入，更新的性能

**单个索引（二级索引）**

单个条件就是单个索引咯，不过有个技巧是，**二级索引里面存储了主键的ID**，查询的时候我们可以通过索引属性列查询出索引属性列+主键值，这两个属性是不会回表查询的性能会非常快，我们可以利用这一点做一些骚操作，比如将大查询操作查两个，通过索引属性查询结果拿到主键，然后通过主键去查询最终的结果，性能可能也会有提升。

#### 聚簇索引

猜想：一个表只能有一个聚簇索引，聚簇索引有两部分内容：B-Tree索引（主键值）+数据行（放在叶子页上面），聚簇索引上面保存的是主键的ID，叶子页上面保存整个行的数据（所有的列），我们检索表数据的时候其实不是从磁盘那里去查询，而是查询的聚簇索引上面的数据

**大胆猜测**：是不是通过主键去查询，就可以利用到聚簇索引上面你的行，就不需要回表查询了

**验证结果**：查询值的时候，Extra是空的，结果显示可能仍然是回表查询了（也有可能聚簇索引不是我想象的这样查询的）

![1666749231758](Vue进步.assets/1666749231758.png)

[扩展阅读](https://zhuanlan.zhihu.com/p/102157197)

#### 复合索引与单个索引的区别

我猜测就是根据业务逻辑看，查询的时候需不需要回表，因为复合索引的值都存在索引里面，只查询索引的值，就不用回表查其他的属性值了

#### Exists走索引，In不走索引

```mysql
-- 不走索引
select * from A where A.id in (select id from B);
-- 走索引
select * from A where exists (select * from B where B.id = A.id);
```



#### 用or与in的区别

**结论**：in和or的效率，取决⽬标条件列**是否有索引**或者是否是主键，如果**有索引或者主键性能没啥差别，如果没有索引，in的性能要远远优于or.**

测试结论一致

![1666324634961](Vue进步.assets/1666324634961.png)

union all 不行，性能几乎一致

![1666324812107](Vue进步.assets/1666324812107.png)

[扩展阅读](https://wenku.baidu.com/view/b52d63d1adf8941ea76e58fafab069dc50224702.html)

#### 强制使用索引（有索引不走索引）

```c#
// 建立索引默认走索引，5.70s
select  count(1) from bailun_bltpro.t_bi_ebay_category_specific   where  usageconstraint='Required';
// 使用force index(sex) 主动调用索引 5.72s
select  count(1) from bailun_bltpro.t_bi_ebay_category_specific  force index(IX_UsageConstraint) where  usageconstraint='Required';
// 使用ignore index(sex) 不调用索引，2.741s
select count(1) from bailun_bltpro.t_bi_ebay_category_specific  ignore index(IX_UsageConstraint) where  usageconstraint='Required';
//【技巧】:count(*)的时候，强制使用主键索引，会快一些，FORCE INDEX(PRIMARY)
SELECT COUNT(*) FROM t_product_property FORCE INDEX(PRIMARY)

```

#### 区分度：两个列用哪个列做索引效果最好

**答案：选择区分度高的列建立索引，比如易仓推送状态（EcangSyncStatus）的区分度是：0.0012，SKU属性的区分度是：0.9667，前者区分度小于0.3，理论上不建议建立索引的，具体经过性能测试后决定**

**区分度公式：COUNT(DISTINCT 列名)/COUNT(\*)**

区分度字段不重复的比例，比例越大扫描的记录数越少，唯一键的区分度是1（**值越接近1越好，等于1代表只有一个值，越小于1索引的效果越差**）

**Ebay平台属性必填选项示例：**

```c#
//3 ,0.00000141

// 总行数 212w
select COUNT(distinct UsageConstraint) ,ROUND(count(distinct UsageConstraint)/COUNT(UsageConstraint),4) from bailun_bltpro.t_bi_ebay_category_specific 

```

示例表中性别有三种值：

- Optional 可选
- Recommended 推荐
- Required 要求

从结果可得知区分度为0.00000141，即根据UsageConstraint查询1条数据需要扫描约10W行

**款号属性查询示例：**

```c#
// 70136 ,	0.0551
select COUNT(distinct ProductSpu),ROUND(COUNT(distinct ProductSpu)/COUNT(ProductSpu),4) from t_product_property
```

说明该字段在满足其它索引条件的情况下，比较适合加上索引

**注：无法从区分度准确计算出扫描行数，但可以确定区分度越大，扫描行数越少**

#### 为什么性别男女列不建议加索引（见上面的数据属性区分度）

答：**用索引的选择性来回答，惯用的百分比界线是"30%" ，当值超过了这个分界线，不介意用索引，比如男，女这两个值可能都接近50%这个分界线（超过了百分之30%），那么不介意使用**

**索引的选择性**是指索引列中不同值的数目和表的记录数的比值。假如表里面有1000条数据，表索引列有980个不同的值，这时候索引的选择性就是980/1000=0.98 。索引的选择性越接近1，这个索引的效率很高。

性别可以认为是3种，男，女，其他。如果创建索引，查询语句 性别=‘男'的数据，索引的选择性就是3/1000=0.003。索引的选择性值很低，对查询提升不大，所以性别建索引意义不大

#### 普遍的优化逻辑

- 大字段：调整大字段的大小，剔除空格，换行字符，根据最大值进行缩减

- 字段大小：根据实际值的情况调整字段值的大小

- 首页面优化：1.获取指定属性，不获取全部*（ORM框架这种情况很普遍，如果想提高一些速度这个是不错的方法） 2.一次获取行数与Limit：SQL_CALC_FOUND_ROWS 3.使用下一页不使用翻页 （就不用获取总行数）

  



#### 不走索引可能的情况

- 跨库查询的情况可能出现A库与B库的字符集不一样，不同库的字符集不一样不会走索引
- 两个表的数据类型不一样，比如varchar直接和int类型相比较

### 优化内容

#### 想针对于某个点看看Mysql的源码，源码在哪里？

源码的下载可以在：https://dev.mysql.com/downloads/mysql/
在网页上“SelectOperating System”选择“SourceCode”，在“Select OS Version”选项中使用“All”，即可出现下面的列表：

![img](https://img-blog.csdnimg.cn/20210504155811835.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZwY2M=,size_16,color_FFFFFF,t_70)

mysql_source

点击“Download”，进入新页面后选择“No thanks, just start my download.”，直接下载即可，根据网速，二百多兆的数据不大也不小。下载后，解压缩，即可出现代码，基本结构如下：

![img](https://img-blog.csdnimg.cn/20210504155823823.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZwY2M=,size_16,color_FFFFFF,t_70)

**mysql_dic**

代码的主要功能如下：
boost:这个版本是自带Boost的库相关文件的，放在此处，如果是其它的版本就没有这个文件夹
client:客户端相关的软件和工具代码
cmake:CMAKE相关的脚本命令文件
components:组件工具
Docs:文档文件夹
doxyen_resources:doxyen工具相关资源
extra:引入的一些其它包，如网络消息谁的SSL包以及一些小工具。
include:源码用的相关的头文件放置的文件夹，但不包括存储引擎的头文件。
libbinlogevents:解析Binlog的lib服务，5.7后提供。
libbinlogstandalone:脱机配置CMAKE
libmysql:可嵌入式的客户端API
libservices:动态服务插件管理
man:帮助文档
mysql-test:服务端mysqlid的测试工具。
mysys:MySql自己实现的数据结构和一些基本算法。如数组和链表等。
packaging:打包相关
plugin:插件管理文件夹，包括一些动态加入的插件。
router:集群路由
scripts:系统工具运行的脚本。
share:共享信息，err和字符集
source_downloads:
sql:服务端的主要代码，包括main函数。
sql-common:服务端和客户端通用的一些代码。
storage:存储引擎相关文件。
strings:字符串库
support-files:.conf的示例文件和相关工具。
testclients:客户框架测试。
unittest:单元测试，这个搞程序的都知道。
utilities:公用的一些文件，有ZLIB等
vio:虚拟网络IO处理系统，不同平台或不同协议的网络通信API的二次封装。

三、体系结构
MySql的体系结构主要如下：



mysql_construct

其主要分为以下几层：
1、连接管理层，包括连接池支持的各种组件
2、服务管理层，含管理服务和相关工具、接口、分析、优化等组件。
3、存储引擎，含各种以插件提供的数据存储引擎。
4、文件存储层，含各种实际落到硬盘的物理文件。
————————————————
版权声明：本文为CSDN博主「fpcc」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/fpcc/article/details/116401551



#### 文件存储格式：Compact



文件存储格式选择是否正确？

MySQL默认的行存储格式是Compact，但大部分情况并不是最优的选择，可以根据表中数据特点为表选择合适的存储格式。例如：在表中存在大量重复，compressed格式可能是更好的选择。对于text、BLOB等文本数据，dynamic格式可能更适合。

文件存储格式指的是表级别的格式还是数据库级别的格式？表级别的话我感觉日志层面也可以进行更改

表中存在大量重复值，指的是行全部相同算重复，还是说同一行，列属性相同也算重复？

```mysql
#Row_format 代表使用的文件存储格式
SHOW TABLE STATUS LIKE 't_product_property_bak20221028' -- 本地这个表使用的是Dynamic格式
```



#### 使用Procedure Analyse优化表数据结构

```mysql
# 执行返回中的Optimal_fieldtype列是mysql建议采用的列。
select * from t_product_property_bak20221028  PROCEDURE ANALYSE(1) 
```

[原文地址](http://t.zoukankan.com/duanxz-p-3968639.html)

#### 开源数据搜索引擎

**Solr**:Java web 项目

**ES（Elasticsearch）**:性能很重，应用复杂场景

**meilisearch**:小型搜索引擎，应用简单场景

https://www.meilisearch.com/
刚才我说的全文搜索是这个。这个目前据了解数据量超过100W级别索引数据会很慢，但是一些简单的业务场景是完全可以满足了。

#### 更改表的存储引擎

mysql的搜索引擎是**针对于表**的，我们可以针对于不同的表更换不同的搜索引擎：

```mysql
# 更改搜索引擎：MyISAM，InnoDB，MEMORY
ALTER TABLE t_product_property_log ENGINE = MyISAM;#5.7.18-txsql-log 修改不了，mysql 8 可以修改成功， 但是腾讯官方做了限制，自动会转换为InnoDB引擎
# 查看存储引擎：在第二个create table的脚本最后面看是哪个存储引擎
show create table t_product_property_log;
```

**腾讯的数据库更改 MyISAM后会强制自动转换 InnoDB为存储引擎** 真恶心阿

[ MyISAM 自动转换为 InnoDB 引擎限制](https://cloud.tencent.com/document/product/236/45043)

![1668587622769](Vue进步.assets/1668587622769.png)

#### MySql的存储引擎

- MyISAM存储引擎：数据表主要用来插入和查询记录，这种场景效率很高，count(0)直接返回结果不需要查询，可以用作日志表

- InnoDB存储引擎：提供事务管理，回滚，崩溃回复能力，并且需要实现并发控制

- MEMORY存储引擎：临时存放数据，数据量不大，不需要安全，**数据存放在内存，读取速度快**

  ```mysql
  #查看目前的存储引擎 support='DEFAULT' 等于目前默认存储引擎
  show engines  ; #mysql 默认的是InnoDB存储引擎
  # 设置数据库默认的存储引擎（注意：数据库重启后，会回复到原始默认的存储引擎）
  SET default_storage_engine=MyISAM;
  ```

  

#### 尽可能的使用 varchar/nvarchar 代替 char/nchar 

**因为char/nchar长度是不可变的，如果允许为NULL 值，Null也是占用声明时的长度值，对磁盘空间，检索效率还是有影响，如果用varchar/nvarchar就没有这个烦恼，因为它的长度是可变的长度，并且NULL不占用空间**

因为首先变长字段存储空间小，可以节省存储空间，其次对于查询来说，在一个相对较小的字段内搜索效率显然要高些。不要以为 NULL 不需要空间，比如：char(100) 型，在字段建立时，空间就固定了， 不管是否插入值（NULL也包含在内），都是占用 100个字符的空间的，如果是varchar这样的变长字段， null 不占用空间。



#### 优化union查询:union all比union高效

MySQL通过创建并填充临时表的方式来执行union查询。除非确实要消除重复的行，否则建议使用union all。原因在于**如果没有all这个关键词，MySQL会给临时表加上distinct选项，这会导致对整个临时表的数据做唯一性校验，这样做的消耗相当高**。

```mysql
#高效：

SELECT COL1, COL2, COL3 FROM TABLE WHERE COL1 = 10 

UNION ALL 

SELECT COL1, COL2, COL3 FROM TABLE WHERE COL3= 'TEST'; 
#低效：

SELECT COL1, COL2, COL3 FROM TABLE WHERE COL1 = 10 

UNION 

SELECT COL1, COL2, COL3 FROM TABLE WHERE COL3= 'TEST';


```



#### 使用group by 关闭默认的order by 排序

默认情况下，MySQL 会对GROUP BY分组的所有值进行排序，如 “GROUP BY col1，col2，....;” 查询的方法如同在查询中指定 “ORDER BY col1，col2，...;” 如果显式包括一个包含相同的列的 ORDER BY子句，MySQL 可以毫不减速地对它进行优化，尽管仍然进行排序。

因此，**如果查询包括 GROUP BY 但你并不想对分组的值进行排序，你可以指定 ORDER BY NULL禁止排序**。例如：


```mysql
SELECT col1, col2, COUNT(*) FROM table GROUP BY col1, col2 ORDER BY NULL ;

```



#### 多表关联查询时，小表在前，大表在后。

在MySQL中，执行 from 后的表关联查询是**从左往右执行的**（Oracle相反），**第一张表会涉及到全表扫描**，所以将小表放在前面，先扫小表，扫描快效率较高，在扫描后面的大表，或许只扫描大表的前100行就符合返回条件并return了。

**例如：**表1有50条数据，表2有30亿条数据；如果全表扫描表2，你品，那就先去吃个饭再说吧是吧。

如果难以判断哪个是大表，哪个是小表，可以用inner join连接，MySQL会自动选择**小表去驱动大表**

#### 【Select执行顺序】

**执行过程**

![img](https://pic1.zhimg.com/v2-3158800935bdbd30a57c2263ac8b5eb4_r.jpg)

**Select语法顺序：**

```mysql
SELECT

DISTINCT <select_list>

FROM <left_table>

<join_type> JOIN <right_table>

ON <join_condition>

WHERE <where_condition>

GROUP BY <group_by_list>

HAVING <having_condition>

ORDER BY <order_by_condition>

LIMIT <limit_number>
```

**执行顺序：**

```mysql
FROM<表名> # 选取表，将多个表数据通过笛卡尔积变成一个表。

ON<筛选条件> # 对笛卡尔积的虚表进行筛选

JOIN <join, left join, right join...>

<join表> # 指定join，用于添加数据到on之后的虚表中，例如left join会将左表的剩余数据添加到虚表中

WHERE<where条件> # 对上述虚表进行筛选

GROUP BY<分组条件> # 分组

<SUM()等聚合函数> # 用于having子句进行判断，在书写上这类聚合函数是写在having判断里面的

HAVING<分组筛选> # 对分组后的结果进行聚合筛选

SELECT<返回数据列表> # 返回的单列必须在group by子句中，聚合函数除外

DISTINCT数据除重

ORDER BY<排序条件> # 排序

LIMIT<行数限制>
```

————————————————
版权声明：本文为CSDN博主「一名陈序员」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/m0_56864757/article/details/125217302

#### 产品资料管理使用了哪些方法进行优化

- 根据实际使用长度缩减不必要的属性大小
- 剔除大属性Json那个 （这个最关键）
- 删除区分度低的索引，比如有些状态值，重复率太高，用索引反而影响性能
- count(*) 强制用主键索引扫描
- 剔除冗余的consul实例

#### 表数据量过千万怎么处理？

短期内不好优化，看你业务是怎样的，建议**分库分表**，比如把**业务完结的数据放到关单表**，**业务表只保留业务操作的数据**

关单表提供查询功能，关单表数据量大就上**tidb es**这样的处理大数据的数据库

我们项目中是这样的，月数据流水上亿也没问题

#### 是否开启了查询缓存

担心会影响库存同步的内容，如果读取的是缓存的值，不是结果的值就不太好，暂时不开启查询缓存

```mysql
show variables like '%query_cache%'
/*
have_query_cache	YES  # YES 数据库支持缓存功能，No不支持缓存功能
query_cache_limit	1048576  #单位kb，缓存的结果集大小最多是多少kb(默认为1mb，这里就是1mb),超过了1mb就不会被缓存
query_cache_min_res_unit	4096
query_cache_size	0        //给多少内存留给缓存，如果为0，及时下面的缓存开关打开，也不会有缓存
query_cache_type	OFF    # ON 启用 、OFF 禁用、DEMAND 按需缓存
query_cache_wlock_invalidate	OFF  #有锁的情况等锁释放返回值，还是返回缓存值
*/
```

[查询缓存扩展阅读](https://www.zsythink.net/archives/1111)

#### 多表更新如何优化

```mysql
#原语句：超时严重
-- tmp_end_sku 50w
update tmp_end_sku set Stock=0
 where ProCode not  in (
 -- 查询37w
  select ProductSku from bailun_skums.t_ec_product_sku_stock as s
  where s.WarehouseId in (2,45)
 )
#使用： not exists优化
#使用： left join 优化

```

#### 分页的时候也可以一次查分页数与总行数

```mysql
# 查询中加入 SQL_CALC_FOUND_ROWS 代表读取总行数(主要时间消耗在这里)
SELECT  SQL_CALC_FOUND_ROWS  *  FROM `t_product_property`  WHERE ( `IsDeleted` = 0 )  AND  UserOrganizationId in ( 1,24) LIMIT 20;
SELECT FOUND_ROWS() as total_count; #FOUND_ROWS()读取到总行数，这一条语句不花时间0.00s 

#对数量准确度不高可以用query_rewrite快速返回
```

  

#### count(*)与count(1)，count(name)的选择

推荐使用count(*)性能最优，也可以使用： 

show table status where NAME='t_product_property'  获取表的近视值进行使用，结果：1329418

```mysql
select count(*) from t_product_property; -- 22s, 1409751,count(*)等于count(0)，性能与count(1)没有区别
select count(1) from t_product_property; -- 23s,1409751,性能与count(*)一致
select count(id) from t_product_property; -- 22s ,1409751, 走主键索引
select count(productsku) from t_product_property; -- 22, 1409751 productsku几乎没有空值，没有问题
select count(upc) from t_product_property; -- 10s ,1100995,只查不为空的行，所以最快
```





### 【Mysql使用变量：牛逼】

```mysql
 select @p:= ProductSku from t_product_property_bak20221028 where USERORGANIZATIONID='30' LIMIT 1; #RBB23730001 (1.注意分号; 2.注意变量的冒号: 3.结果在结果1里面看)
 select @p; #RBB23730001(1.注意分号; 2.结果在结果2里面看)
 
 #还有另一种办法给变量赋值（这种方法在存储过程，在单个sql中通用，上面的@符号在存储过程中用不了）
 DECLARE _variabletype int; 
 SELECT variableType into _variableType from t_adjust_price_task_config where id=task_id;
```

### 怎么不通过Count(*)快速知道表的总行数？

有时候千万级别的大表，count(*)的速度非常慢，可能要好几分钟，这个时候我们可以用快速的方法，得到一个估计的数值，接近正确值的数值：

```mysql
SHOW TABLE STATUS LIKE 't_adjust_price_task_run_detail';
21372418  实际：22309564  差异 100万
```



### 这么大一个美化SQL的按钮在这里

![1679021811600](Vue进步.assets/1679021811600.png)

### mysql使用反引号进行转义

```mysql
#不使用反引号抛异常，因为：Condition是mysql关键字
select Condition from t_bi_amazon_lister    LIMIT 1
#使用反引号，正常查询
select `Condition` from t_bi_amazon_lister    LIMIT 1

反引号在键盘左上角，Esc键的下面，按一下就可以，不用加alt,shift辅助键辅助
```



### 两张千万级大表查询如何进行优化

最近去面试，面试官提了一个问题：现在有2张表，一张是user表，一张是order表，数据都在千万级以上的，order表有个user_id字段是user表的主键，现在查出用户订单情况的sql是：

```c#
SELECT
    `user`. user_name,
    `order`.order_number
FROM
    `user`
LEFT JOIN `order` ON`user`.id = `order`.user_id		
```

**我的思路：**

​            1.alter order add index ix_user_id(user_id)
​			2.name与number建立索引，不回表查询，直接覆盖索引 
​			3.冗余属性order表里面扩展出用户名称，就不用查两个表了
​			4.页面上查询的情况，先查询订单信息，然后异步返回用户信息，根据用户编号查询20行数据很快
​			5.检查两个表建立索引的情况，检查explan的情况，看索引的使用情况
​			

**论坛的思路：**

​            1.学习：建立专门的中间表，存储用户名与订单编号，提供给业务人员使用
​			2.学习：对用户编号进行分区，建立分区表的概念提高查询的性能
​			3.学习：对订单表进行存量数据冗余处理，一般历史的订单用户基本很少看，可以进行分表，具体范围看业务那边

​            4.学习：将结果放到es中，维护es中的数据，然后进行使用，也可以很大提升性能

### 高性能更新大数据量两表数据

[快速比较mysql中两个表中的数据是否有差异并找出差异项](https://blog.csdn.net/usualheart/article/details/107403759)

**分区表不只是能解决查询的问题，也可以解决更新的问题**

**场景：**
两个大表联表更新，性能低，执行速度慢。
条件：两个表的数据量都大于2000W

**案例：**
`bailun_bltpro`.`t_bi_walmart_listing`表数据量2800W
`bailun_publish_walmart`.`sp_walmart_listing`表数据量2000W+

**结果：**

优化前执行十分钟，优化后执行三分钟

### 原始更新语句

```sql
UPDATE `bailun_bltpro`.`t_bi_walmart_listing` as l join `bailun_publish_walmart`.`sp_walmart_listing` as t
SET l.`WpId` = t.`WpId`, 
  l.`Gtin` = t.`Gtin`, 
  l.`Upc` = t.`Upc`, 
  l.`Sku` = t.`Sku`, 
  l.`SalePrice` = t.`SalePrice`,
  l.`SellerBuyboxItemPrice` = t.`BuyBoxPrice`, 
  l.`SellerBuyboxShipPrice` = t.`BuyBoxShippingPrice`, 
  l.`MSRP` = t.`MSRP`, 
  l.`ShippingWeight` = t.`ShippingWeight`, 
  l.`ShippingWeightUnit` = t.`ShippingWeightUnit`, 
  l.`Brand` = t.`Brand`,  
  l.`PublishedStatus` = t.`PublishedStatus`, 
  l.`LifeCycleStatus` = t.`LifeCycleStatus`, 
  l.`StatusChangeReason` = t.`StatusChangeReason`, 
  l.`OfferStartDate` = t.`OfferStartDate`, 
  l.`OfferEndDate` = t.`OfferEndDate`, 
  l.`ItemCreationDate` = t.`ItemCreationDate`, 
  l.`ItemLastUpdated` = t.`ItemLastUpdated`, 
  l.`Reviews` = t.`Reviews`, 
  l.`AverageRating` = t.`AverageRating`, 
  l.`FulfillmentLagTime` = t.`FulfillmentLagTime`, 
  l.`FulfillmentType` = t.`FulfillmentType`, 
  l.`ShipMethod` = t.`ShipMethod`, 
  l.`PrimaryImageUrl` = t.`PrimaryImageUrl`, 
  l.`ProductName` = t.`ProductName`, 
  l.`IsVariation` = t.`IsVariation`, 
  l.`VariantGroupId` = t.`VariantGroupId`, 
  l.`IsPrimary` = t.`IsPrimary`, 
  l.`GroupingAttributes` = t.`GroupingAttributes`
WHERE l.ShopId=t.ShopId
and l.ItemId=t.ItemId
and t.ShopId=2342;
```

**此语句根据ShopId的数据量，一般执行时间超过10分钟**

### 优化更新语句

```sql
## step_1 查询两个表差异数据(根据字段排查)，目的是减少结果集
DROP TABLE IF EXISTS tmp_listing_diff;
CREATE TEMPORARY TABLE tmp_listing_diff(
  SELECT * 
  FROM(
    SELECT Id,ShopId,ItemId,Sku,SalePrice,SellerBuyboxItemPrice,SellerBuyboxShipPrice,Brand,PublishedStatus,LifeCycleStatus,StatusChangeReason,OfferStartDate,OfferEndDate,ItemCreationDate,ItemLastUpdated,FulfillmentType,ShipMethod,ProductName,VariantGroupId,'bltpro' as 'FromTable' 
    FROM `bailun_bltpro`.`t_bi_walmart_listing`
    WHERE shopid=2341
    UNION ALL 
    SELECT Id,ShopId,ItemId,Sku,SalePrice,BuyBoxPrice as SellerBuyboxItemPrice,BuyBoxShippingPrice as SellerBuyboxShipPrice,Brand,PublishedStatus,LifeCycleStatus,StatusChangeReason,OfferStartDate,OfferEndDate,ItemCreationDate,ItemLastUpdated,FulfillmentType,ShipMethod,ProductName,VariantGroupId,'publish_walmart' as 'FromTable' 
    FROM `bailun_publish_walmart`.`sp_walmart_listing`
    WHERE shopid=2341
  )  t 
  GROUP BY Sku,SalePrice,SellerBuyboxItemPrice,SellerBuyboxShipPrice,Brand,PublishedStatus,LifeCycleStatus,StatusChangeReason,OfferStartDate,OfferEndDate,ItemCreationDate,ItemLastUpdated,FulfillmentType,ShipMethod,ProductName,VariantGroupId
  HAVING COUNT(*) = 1
);

## step2 获取待更新表的原始数据
DROP TABLE IF EXISTS tmp_listing_update;
CREATE TEMPORARY TABLE tmp_listing_update(
  SELECT l.* FROM `bailun_publish_walmart`.`sp_walmart_listing` as l
  JOIN tmp_listing_diff as t on l.Id=t.Id
  where t.FromTable='publish_walmart'
);

CREATE INDEX ix_shopid_itemid on tmp_listing_update(ShopId,ItemId);

## step3 将最终结果集的数据更新到原始表
UPDATE `bailun_bltpro`.`t_bi_walmart_listing` as l join tmp_listing_update as t
    SET l.`WpId` = t.`WpId`, 
      l.`Gtin` = t.`Gtin`, 
      l.`Upc` = t.`Upc`, 
      l.`Sku` = t.`Sku`, 
      l.`SalePrice` = t.`SalePrice`,
      l.`SellerBuyboxItemPrice` = t.`BuyBoxPrice`, 
      l.`SellerBuyboxShipPrice` = t.`BuyBoxShippingPrice`, 
      l.`MSRP` = t.`MSRP`, 
      l.`ShippingWeight` = t.`ShippingWeight`, 
      l.`ShippingWeightUnit` = t.`ShippingWeightUnit`, 
      l.`Brand` = t.`Brand`,  
      l.`PublishedStatus` = t.`PublishedStatus`, 
      l.`LifeCycleStatus` = t.`LifeCycleStatus`, 
      l.`StatusChangeReason` = t.`StatusChangeReason`, 
      l.`OfferStartDate` = t.`OfferStartDate`, 
      l.`OfferEndDate` = t.`OfferEndDate`, 
      l.`ItemCreationDate` = t.`ItemCreationDate`, 
      l.`ItemLastUpdated` = t.`ItemLastUpdated`, 
      l.`Reviews` = t.`Reviews`, 
      l.`AverageRating` = t.`AverageRating`, 
      l.`FulfillmentLagTime` = t.`FulfillmentLagTime`, 
      l.`FulfillmentType` = t.`FulfillmentType`, 
      l.`ShipMethod` = t.`ShipMethod`, 
      l.`PrimaryImageUrl` = t.`PrimaryImageUrl`, 
      l.`ProductName` = t.`ProductName`, 
      l.`IsVariation` = t.`IsVariation`, 
      l.`VariantGroupId` = t.`VariantGroupId`, 
      l.`IsPrimary` = t.`IsPrimary`, 
      l.`GroupingAttributes` = t.`GroupingAttributes`
    WHERE l.ShopId=t.ShopId
    and l.ItemId=t.ItemId
    and t.ShopId=2342
```

**优化后，执行时间3分钟左右**

### pt-osc 在线DDL工具

[扩展阅读](https://cloud.tencent.com/developer/article/1702932)

### Docker安装Mysql一行命令

```mysql
 # 端口3306,用户名:root ,密码：123456，--name mysql（这个mysql代表的是我设置容器的名字），后面的那个mysql是pull下来的镜像的名字
 sudo docker run -d -p 3306:3306 -v /usr/local/mysql/conf:/etc/mysql/conf.d -v /usr/local/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 --name  mysql mysql
```

[原文地址](https://blog.csdn.net/weixin_43830765/article/details/123849821)

### Mysql 分区表怎么用？（方便管理，无法提高性能）

2023年9月1日12:02:06 分析过后发现分区表不是用来提升性能的，是为了方便数据的管理，那我没有研究的必要了。如果使用分区表的时候没有使用到分区列，那速度比全表更慢

原文链接：https://cloud.tencent.com/developer/article/1877218

> 许多人会认为分区表将一个大表划分为多个小表，因此MySQL数据库的性能将大大提高。

> 这是错误的理解！**分区表技术并不是用来提高MySQL数据库的性能，而是为了方便数据管理。**

> 分区表的创建需要主键包含分区列；在分区表中唯一索引仅在当前分区文件唯一，而不是全局唯一；分区表唯一索引推荐使用类似 UUID 的全局唯一实现；

> **分区表不解决性能问题，如果使用非分区列查询，性能反而会更差；推荐分区表用于数据管理、速度快、日志小**。
>
> 原文地址：https://juejin.cn/post/7164406180243046436



分区表是物理分区不是逻辑分区（物理上是多个表，逻辑上是一个表），根据我们的具体业务，决定使用哪个属性进行分区，比如如果是电商行业，我们可以拿店铺作为分区属性，每个店铺一个分区表进行管理，我们也可以每月放一个分区表（**时间范围分区是最常见的**），按月份进行分表，也可以根据ID主键的范围1~10w一个分区表

**分区数量**：5.7版本单服务器单表分区最大支持1024个（如果是集群的话可以最大到8192个），并且会话第一次访问分区表，都需要打开全部的分区表

**分区属性**：根据具体业务，与表的数据量分布尽可能取一个平均参数，比如电商行业中订单创建时间的月份，订单属于哪个店铺，不动产系统中这个属性可能是地区编号，创建时间。**创建时间按月份进行使用还是挺好的，数据分布均匀，每个分区表的数据量都相似，不过后续地区的数据加进来不会在一起，这个倒是挺麻烦**

**分区函数**：根据分区函数的逻辑，决定使用哪个分区表，比如使用日期分区，使用哈希取余分区

**分区类型：**

- **范围分区**：**日志表，流水表这种按日期类型进行操作的**。根据时间范围年，月，日，ID进行划分区域,**范围分区只能是整型值作为分区列（多列分区除外）**，比如先来100万，放到分区A表，后面的101~200万，放到分区B表（需要注意分区的键必须是主键的其中之一，比如b键是分区列， PRIMARY KEY （a, b）需要把b加入到主键中，`MySQL 中的分区表是把一张大表拆成了多张表，每张表有自己的索引，从逻辑上看是一张表，但物理上存储在不同文件中`。）
- **列表分区**：根据指定的值进行分区，比如城市，不同的城市划分不同的分区表，广州在一个分区，深圳在一个分区，不同的渠道不同的类型进行分区。
- **哈希分区**：**对于分布规则不明显的数据，并没有明显的范围查找等特征，可以使用HASH分区，比如按用户ID类型的进行操作的表，uuid操作的表**。哈希取余分区，比如1~10取模，哈希一下取余，1-10分区表选择其中一个分区表（哈希分区的属性值必须是整型或者是返回整型值的函数值）
- **键值分区**：键值分区与哈希分区使用场景一致，区别在于这个是mysql内部实现的哈希算法帮助我们分区，并且使用主键/唯一索引键作为分区列，对数据类型的兼容度大一些（Key分区可以是整形也可以是字符串(除text和BLOB之外的所有数据类型)，只要是唯一值就可以,）

[原文地址/四大分区/子分区/分区维护内容](https://blog.csdn.net/frostlulu/article/details/122304238)

**范围分区：**

例：定义一个员工表，根据员工ID分区，1~10号员工一个分区，11~20号员工一个分区，依次类推，共建立4个分区：

```mysql
#1.创建表
create table employees (
    id int not null primary key,
    first_name varchar(30),
    last_name varchar(30))
    partition by range(id)( #根据ID分区，分区的值需要连续，并且不允许重复，partition by range(属性名)定义分区的属性
    partition p0 values less than (11), #1~10一个分区，分区的范围通过partition (分区名)p0 values less than(11)(分区范围最大值，id<11)
    partition p1 values less than (21), #11~20一个分区
    partition p2 values less than (31), #21-30一个分区
    partition p3 values less than (41) #31-40一个分区
    PARTITION p4 VALUES LESS THAN MAXVALUE #不在范围的值的放在这个分区
    );
    
#2.保存到p0分区(i 符合id<11分区p0条件)
insert into employees values(1,'Vincent','Chen');

#3.指定分区查询
select * from employees partition(p0);  -- 查询p0分区
select * from employees partition(p0,p1);  -- 查询p0和p1分区

#4.删除指定分区： 分区p1被删除后，所有p1分区的【数据都已丢失】，此时原p1分区的范围将由p2覆盖。
alter table employees drop partition p1;

#5.【追加分区：在已有分区的最大范围上面追加分区】
alter table employees add partition (partition p4 values less than(51));
 -- 范围分区中，不允许通过新增分区的方式追加之前被删除的分区
alter table employees add partition (partition n1 values less than(21));
 -- 上面试图对ID10~20的员工新增一个分区，系统会返还错误。

# 6.分区重组
如果一定要在分区之间插入新的分区，则可以采用重组织的方式，将已有分区的数据重新划分，达到创建新分区的效果：见图1：
 -- 将P2分区重新组织，拆分成两个分区：p1,p2
alter table employees reorganize partition p2 into (
partition p1 values less than(21),
partition p2 values less than(31)
);

# 7.定义上限分区：所有不包含在id范围值的分区都放在这个分区里面:maxvalue代表的是当前分区的上限值
alter table employees add partition(partition pmax values less than(maxvalue));

# 8.技巧：比如函数分区，日期分区
1.范围分区的条件除了直接用值，还可以用函数来定义。
2.一个常用的场景就是按时间分区，例如：**在create table中使用partition by range(year(hire_date))**，可以按照年份来进行分区。这种分区方式在需要定期清理过期数据的场景会非常方便。
3.Range partition分区表与null值的关系:
对range分区来说，如果插入数据分区键为null，是可以成功的，数据会被放到第一个分区中。

# 9.多列分区技巧:可以使用日期(非年月日也可以)，字符串，多个列进行分区
  -- 日期分区/字符串分区
CREATE TABLE members (
    firstname VARCHAR(25) NOT NULL,
    lastname VARCHAR(25) NOT NULL,
    username VARCHAR(16) NOT NULL,
    joined DATE NOT NULL
)
PARTITION BY RANGE COLUMNS(joined) (
    PARTITION p0 VALUES LESS THAN ('1960-01-01'),
    PARTITION p1 VALUES LESS THAN ('1970-01-01'),
    PARTITION p2 VALUES LESS THAN ('1980-01-01'),
    PARTITION p3 VALUES LESS THAN ('1990-01-01'),
    PARTITION p4 VALUES LESS THAN MAXVALUE);
 -- 多列分区，a=0,b=10 那么就到分区p0表
 CREATE TABLE rc2 (
    a INT,
    b INT
)
-- 使用RANGE COLUMNS(列1，列2) 创建多列分区/不限制字符类型的分区
PARTITION BY RANGE COLUMNS(a,b) (
    PARTITION p0 VALUES LESS THAN (0,10),
    PARTITION p1 VALUES LESS THAN (10,20),
    PARTITION p2 VALUES LESS THAN (10,30),
    PARTITION p3 VALUES LESS THAN (MAXVALUE,MAXVALUE) );
```

 **列表分区：**

```mysql
#创建列表分区：1,3,5,7,9的值在一个分区中，其余的值在其他的分区中，如果不在范围的值保存，会出现异常（innodb引擎异常出现后会自动回滚（数据不在），myisam引擎异常出现后不会回滚（数据还在））
CREATE TABLE employees (
    id INT NOT NULL,
    fname VARCHAR(30),
    lname VARCHAR(30)
)
PARTITION BY LIST(id) (
    PARTITION p0 VALUES IN (1,3,5,7,9), #也可以是城市，广州，深圳，或者性别男，女
    PARTITION p1 VALUES IN (2,4,6,8,10)
);
# 新增分区: p2
ALTER TABLE employees ADD PARTITION(PARTITION p2 VALUES IN (11,12,13,14,15));
# 删除分区：p0
ALTER TABLE employees DROP PARTITION p0;
# List columns多列分区：

和range分区类似，list partition也有一个变种的list columns 分区，此分区类型可以使用多个列同时作为分区条件，且不再限制整数。并且可以使用string,date等数据类型作为分区条件。

# 列表分区与null值的关系:

对于列表分区来说，必须有1个分区显示的指定可以包含null，否则插入会失败。

 
```

**哈希分区：**

```mysql
# 创建hash分区
CREATE TABLE employees (
    id INT NOT NULL,
    first_name VARCHAR(30),
    last_name VARCHAR(30)
)
PARTITION BY HASH(id) -- 对应的hash列必须是整数类型，或者必须是返回整数类型返回值的表达式
PARTITIONS 4;  -- 指定分区的数量4个，忽略partitions创建数量，默认创建1个分区
# 新增分区：指定新增多少个分区
ALTER TABLE employees ADD PARTITION PARTITIONS 3;  -- 新增3个hash分区

# 删除分区：无法指定删除哪个分区，这一点和范围分区与列表分区不一致
ALTER TABLE employees COALESCE PARTITION 1;  -- 移除一个分区

# 哈希分区与null值:
哈希分区与键值分区中的NULL，都会被当做返回为0，放到第一个分区当中
# 哈希分区技巧
哈希分区有两种分区，一种是普通分区使用哈希值取余进行获取分区地址，第二种是线性哈希分区，使用更复杂的算法来进行哈希分区，性能比普通分区要好（TB级别数据存储）
```

**键值分区：**

```mysql
# 创建键值分区：使用主键
CREATE TABLE k1 (
    id INT NOT NULL PRIMARY KEY,
    name VARCHAR(20)
)
PARTITION BY KEY()  -- 未指定分区列，自动使用主键
PARTITIONS 2;  -- 声明两个分区
# 创建键值分区 ：使用唯一索引
CREATE TABLE k1 (
    id INT NOT NULL,  -- 如果未定义not null，创建表会失败
    name VARCHAR(20),
    UNIQUE KEY (id) -- 创建唯一索引
)
PARTITION BY KEY()  -- 未定义主键，自动使用unique key
PARTITIONS 2;
# 创建键值分区：使用字符类型主键
CREATE TABLE tm1 (
    s1 CHAR(32) PRIMARY KEY   -- 字符型主键，同时作为partition key
)
PARTITION BY KEY(s1) -- 指定主键名称
PARTITIONS 10;
# 创建线性哈希分区
CREATE TABLE tk (
    col1 INT NOT NULL,
    col2 CHAR(5),
    col3 DATE
)
PARTITION BY LINEAR KEY (col1) -- Liner key  创建线性哈希分区
PARTITIONS 3;

# 删除/新增操作与哈希分区一致
```

**图1：**

![1669863363439](Vue进步.assets/1669863363439.png)



### Mysql 分区表怎么用好？（方便管理，无法提高性能）

**为什么要做表的分区？**

我脑子首先想到的是日志表的分区，根据月份进行分区放到不同的表中

**提高表的并发**：怎么个提高并发？mysql的锁是索引的读锁，也不会影响其他的行数据，能不能举个例子？

解决后台大数据量使用Like 查询出现卡顿严重的现象，如果知道数据在某个分区上，直接查询某个分区，再进行like

解决检索没有索引，或者索引失效的情况下的数据检索，知道对应的分区，就可以检索，和上面情况下Like索引失效有点类似



**为什么是日志表呢？**

日志表只提供查询，新增功能，没有更新的操作



将普通表更改为分区表，应用层面是不需要进行改动的，数据聚合操作数据库层面也会帮我们直接包括全部的分区，查询的时候需要加入分区的条件吗？



```mysql
show plugins   ;  -- NAME='partition' 查看是否支持分区表
# 一、判断某个表是否为分区表(partition_name属性不等于空就代表有分区表)：可以查看表具有哪几个分区、分区的方法、分区中数据的记录数等信息
 SELECT * FROM information_schema.PARTITIONS WHERE TABLE_SCHEMA=SCHEMA()  and table_name='t_adjust_price_task_run_detail'
# 二、查询所有分区表：有三个分区表，bi_wms_delivery
SELECT * FROM information_schema.PARTITIONS WHERE PARTITION_NAME IS NOT NULLGROUP BY TABLE_NAME ORDER BY CNT DESC LIMIT 50;
# 二、查询表有多少个分区
SELECT TABLE_NAME, COUNT(*) AS CNT
FROM information_schema.PARTITIONS WHERE PARTITION_NAME IS NOT NULL
GROUP BY TABLE_NAME ORDER BY CNT DESC LIMIT 50;
# 三、分析分区执行语句
 -- 查看partitions查看这个列看分区的调用情况，没有条件的情况下就是全分区扫描
explain  select * from bi_wms_delivery  
 -- 有条件的情况下就是扫描指定的分区,分区的数量根据分区的范围进行决定
 explain  select * from bi_wms_delivery   where  id>10 and id <100
# 四、分区排查
-- 查看某⼀特定分区列值属于哪个分区
select M2.$partition.Part_mediasec_func('20150325') as partition_num;
-- 查看某⼀特定分区的记录
select * from SecureUsbLog where M2.$partition.Part_mediasec_func(logTime) = 3
-- 查看各分区所包含的记录数
select $partition.Part_mediasec_func(logTime) as partition_num,
count(*) as record_num
from SecureUsbLog
group by $partition.Part_mediasec_func(logTime)
order by $partition.Part_mediasec_func(logTime)
```



### Mysql分区表实战（方便管理，无法提高性能）

#### 创建分区数据

```mysql
#这张表有2000w数据，首先复制一下表的数据，花费20分钟
create table partition_t_adjust_price_task_run_detail select * from bailun_bltpro.t_adjust_price_task_run_detail

```

#### 已有表分区

**实际业务场景中有可能是现有的表过大，然后导致需要人为拆分成多个分区表进行使用**，如果是开始新建的表马上做分区就简单很多了，但这通常是一个繁琐且可能风险较高的过程。根据你的具体需求和表的大小，你可能需要进行一些特殊的步骤来确保数据的完整性和可用性。

下面是一些基本的方法（**推荐第二种方法**）：

1. **使用ALTER TABLE 语句**

   你可以使用`ALTER TABLE`语句来为已经存在的表添加分区。但请注意，这个操作可能会耗费很多时间，并且在此过程中表通常会被锁定。

   ```mysql
   ALTER TABLE your_existing_table
   PARTITION BY RANGE (your_column) (
       PARTITION p0 VALUES LESS THAN (value1),
       PARTITION p1 VALUES LESS THAN (value2),
       ...
   );
   ```

   在这里，`your_existing_table`是你想要分区的表的名称，`your_column`是你想要用于分区的列，`value1`、`value2`等是分区的边界值。

2. **创建一个新分区表并迁移数据**

   另一种常见的做法是创建一个新的分区表，然后将旧表的数据迁移到新表中。

   ```mysql
   #创建新的分区表
   CREATE TABLE new_partitioned_table (...);
   INSERT INTO new_partitioned_table SELECT * FROM your_existing_table;
   ```

   然后，重命名旧表，重命名新表。

   ```mysql
   RENAME TABLE your_existing_table TO your_existing_table_bak;
   RENAME TABLE new_partitioned_table TO your_existing_table;
   ```

在进行这样的操作之前，务必：

- **备份数据**：总是在执行此类操作之前备份你的数据。
- **测试**：在一个测试环境中进行操作，确保一切按预期进行。
- **考虑性能和可用性**：这样的操作可能会耗费很多时间，并影响到正在使用这个表的其他用户或系统。

请务必谨慎行事，并在进行任何更改之前与数据库管理员或其他相关的专业人士进行咨询。

> 开始实际操作

```mysql
CREATE TABLE `partition_t_adjust_price_task_run_detail` (
  `Id` bigint(20) NOT NULL DEFAULT '0' COMMENT '主键',
  `TaskId` bigint(20) NOT NULL COMMENT '任务Id',
   ...
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```



#### 新建表分区

### desc查看指定表结构

```mysql
# 查看表结构,可以看到所有的属性名称，属性大小
DESC t_product_property_bak20221028;
```



### mysql IndexOf的使用

```mysql
 # 找得到，返回从1开始的索引 (注意小在前，大在后)
select LOCATE('def', 'abcdef'); # 4
 # 找不到，返回0 
 select LOCATE('gg', 'abcdef'); # 0
```



### Mysql动态执行Sql

```mysql
 set @exportSql='select * from t_product limit 1';
 # 预处理sql
 PREPARE stmt FROM @exportSql; 
 # 执行sql
 EXECUTE stmt ;
 #释放预处理命令
 DEALLOCATE PREPARE stmt;
```



### Mysql硬件的优化

MySQL 对硬件的要求主要体现在三个方面：磁盘、网络和内存

**磁盘**

磁盘应该尽量**使用有高性能读写能力的磁盘，比如固态硬盘**，这样就可以减少 I/O 运行的时间，从而提高了 MySQL 整体的运行效率

磁盘也可以**尽量使用多个小磁盘而不是一个大磁盘**，因为磁盘的转速是固定的，有多个小磁盘就相当于拥有多个并行运行的磁盘一样

**网络**

保证网络带宽的通畅（低延迟）以及够大的网络带宽是 MySQL 正常运行的基本条件，**如果条件允许的话也可以设置多个网卡**，以提高网络高峰期 MySQL 服务器的运行效率

**内存**

MySQL 服务器的**内存越大，那么存储和缓存的信息也就越多**，而内存的性能是非常高的，从而提高了整个 MySQL 的运行效率

### Mysql事务与锁定表对性能的影响

**锁定表：**尽管事务是维护数据库完整性的一个非常好的方法，但却因为**事务的独占性**，有时会影响数据库的性能，尤其是在很大的应用系统中。由于在事务执行的过程中，数据库将会被锁定，因此其它的用户请求只能暂时等待直到该事务结束。如果一个数据库系统只有少数几个用户

来使用，事务造成的影响不会成为一个太大的问题；但假设有成千上万的用户同时访问一个数据库系统，例如访问一个电子商务网站，就会产生比较严重的响应延迟。

其实，有些情况下我们可以**通过锁定表代替事务的独占性来获得更好的性能**。下面的例子就用锁定表的方法来完成前面一个例子中事务的功能。

**疑问：**事务也是把表锁住了，排它锁也是把表锁住了，为什么说提高了性能，我不能理解

```mysql
LOCK TABLE inventory WRITE #给指定表加入写锁（排它锁），排它锁状态下只有当前会话可以读，可以写 当前的锁住范围数据，其他的会话都不允许读写当前锁住的范围
SELECT Quantity FROM inventory
WHEREItem='book';
...
UPDATE inventory SET Quantity=11
WHEREItem='book';
UNLOCK TABLES #释放被当前会话锁住的任何锁


```

这里，我们用一个 SELECT 语句取出初始数据，通过一些计算，用 UPDATE 语句将新值更新到表中。包含有 WRITE 关键字的 LOCK TABLE 语句可以保证在 UNLOCK TABLES 命令被执行之前，不会有其它的访问来对 inventory 进行插入、更新或者删除的操作。

### Mysql的事务

```mysql
# ROLLBACK命令可以回滚
# 只要A,B其中一个失败，那么都不能提交
BEGIN;
INSERT INTO salesinfo SET CustomerID=14; #操作A
UPDATE inventory SET Quantity=11
WHERE item='book'; #操作B
COMMIT;
#可以接着上面锁定表的逻辑看
```



### While的使用

```mysql
	 while @count>4 DO 
	  set @exportSql='';
	  set @exportSql =CONCAT('alter table tmp_property add column SpecificName',@count,' varchar(20);');
	 END WHILE;
```



### IF的使用

```mysql
if @maxSecificCount is null or  @maxSecificCount<=4 THEN   
 	  set @maxSecificCount=4;  
	END IF;
# If等于0的情况 （单独执行的情况下，If总是无法单独执行）
IF setIsRuning=0 THEN
SET sync_status=0;
END IF;	
#其他
IF search_condition THEN 
			 
RETURN()
			 
ELSEIF search_condition THEN 
			 
RETURN()
			 
ELSE #存储过程里面没有ELSE
			 
RETURN()
			 
END IF;
# 存储过程中出现的错误代码
IF LOCATE(',',@SALESTATUS)>0 THEN  -- 我猜测是存储过程中不允许写 if 函数逻辑 then
   SET @EXCUTESQL=CONCAT(@EXCUTESQL,' AND B.SALESTATUS IN (6,16) LIMIT 10');
   ELSEIF  LOCATE(',',@SALESTATUS)=0 THEN
   SET @EXCUTESQL=CONCAT(@EXCUTESQL,' AND B.SALESTATUS=@SALESTATUS LIMIT 10');
   END IF;
```



### Mysql获取unix时间戳

```mysql
#获取unix时间戳:unix_timestamp()= 1668675009=	2022-11-17 16:50:09
select  unix_timestamp() ; #1668675009
```



### 查看Mysql版本

```mysql
# 5.7.18-txsql-log
select version();
```





### 创建调用存储过程

```mysql
#创建存储过程（没有参数就把括号里面的参数去除就可以了）
DROP PROCEDURE IF EXISTS proc_product_property_export;
CREATE   PROCEDURE `proc_product_property_export`(in exportSql LONGTEXT)
BEGIN
 	# 产品资料管理导出
       SELECT exportSql;
END
#调用存储过程
call proc_product_property_export('3123');
# 存储过程经验
1.不允许有特殊字符@,这样创建变量是不合法的，declare @VariableType int；可以将@符号更改成_.(注意，如果是外部的变量，单独执行的语法块，可以使用@艾特符号)
2.先创建变量，再创建select语句：顶部是select开头，不允许后面接declare @VariableType int;创建变量，比如select now;开头，后面不允许接declare @VariableType int;，可以先declare声明变量，然后再用select
3.# 不允许用 return ;,下面的抛异常
		IF _taskid IS NULL THEN
	   select  '没有找到符合要求的任务,请检查后重试' as Message;
		return; 
		END IF;
4.给变量赋值 不能用:= ，要用Into 进行赋值
5.入参变量的命名不要和表的属性名称一致，可以用Para前缀开始，就是不要一样
6.存储过程里面可以使用if else
7.存储过程里面可以使用: IF IFNULL(batid,'') ='' THEN

```

[扩展阅读](https://blog.csdn.net/weixin_48530729/article/details/123832523)

### 两表删除新的写法

```
连表删除：DELETE user FROM user,black WHERE user.id=black.id

	# 删除重复数据
	delete a from  t_infringe_keyword_scan a, (select MIN(ID) as id from t_infringe_keyword_scan  where PlatformSource=source  group by InfringeId,InfringeProperty,InfringeKeyword HAVING count(1)>1) as b where a.id=b.id;
```



### 两表更新新的写法

```mysql
#新连表更新：
  UPDATE user a, room b SET a.num=a.num+1 WHERE a.room_id=b.id
# 旧连表更新：
  update tmp_runDetailSku a 
  inner join bailun_skums.t_product  b on a.ProductSpu=b.ProductSpu 
  set a.OrganizationId=b.UserOrganizationId,a.CategoryId=b.CategoryId
  where  b.IsDeleted=0;		
# 多表更新
update t_product a inner join t_product_property b  on a.ProductSpu=b.ProductSpu
   inner join sheet1 c on b.ProductSku=c.sku
	 set a.UserOrganizationId=27,b.UserOrganizationId=27
where a.isdeleted=0 and b.isdeleted=0 

```



### 表字节的占用情况

utf-8下面varchar(60) 代表60个字节，可以保存60个汉字，60个字母，60个数字

utf8mb4是utf8的一个升级扩展:varchar(60) 代表60个字节，可以保存60个汉字，60个字母，60个数字,**60个表情**

### MySql锁表的情况

**结论：**锁表发生在于update,delete,insert当中，（当两表关联更新的时候，即使B表没有更新，只更新A表，**表B是关联表也会锁表**，这个有点有趣）

**锁的类型：**

![1666607823313](Vue进步.assets/1666607823313.png)

**案例：**比如有：update set ...where id between 1 and 10 执行中就会产生独占锁（有多个值），当update set ...where id between 1 and 5 语句2进行更新，那就会产生阻塞锁。如果Update set where id=1,只有一条数据不会产生独占锁，如果id=1有多条就会产生独占锁

**观点：**在 update 语句的 where 条件没有使用索引，就会全表扫描，如果表中有索引的话只会锁住索引里面的范围值（范围锁），如果没有范围锁走的是全表扫描的话，那么就会锁住全表，等语句更新完毕，数据库释放事务后才会解锁。（这个地方要注意一点，并不是where里面的条件加入了索引列就一定不会锁住全表，因为加了索引列，但可能查询优化器他不走索引，所以要看执行计划看查询优化器是不是走了索引，**如果没有走索引，需要使用force index(IX_IndexName)强制mysql走索引**,如果走的是主键索引只有一条数据，那么Mysql只会锁住那一条数据）

**解决：**

1.操作的条件尽量唯一

2.操作语句的时间尽量缩小，减少锁表的时间

3.update与delete 要使用索引作为条件，不然会锁住全表（执行时间再长的话，等着挨怼吧）

**扩展：**

**共享锁（读锁）**：读也会有锁，只不过是共享锁，可以共享给多个读的事务，但是不可以分享给排它锁，简单说就是读的时候可以允许多个直接读，这个时候有共享锁，但是我读的范围有共享锁，如果有更新语句要给这个范围进行更新，要想加上排它锁，那么是不允许的，只有等共享锁释放后才可以加排它锁

**排它锁（写锁）**：独占锁，在更新，删除语句里面触发的，当有个事务触发了这个锁，其他的事务就不能给他锁了

**锁粒度：**锁分为表锁与行锁，表锁就是给整个表加锁，比如上面的更新语句没有where条件，或者where条件没有索引就会锁住表，行锁就是比如where条件是主键索引就是加独占锁，锁住一行数据，如果是走范围索引就是锁住这个索引里面的范围值

**相关语句：**

```mysql
# 释放被当前会话锁住的任何锁
unlock tables
# 给指定表加入写锁（排它锁），排它锁状态下只有当前会话可以读，可以写 当前的锁住范围数据，其他的会话都不允许读写当前锁住的范围
lock tables tableName write
# 给指定表加入读锁（共享锁），共享锁状态下其他的会话可以拿到当前锁住范围的读的权限，但是其他的会话不可以拿到写的权限（共享锁可以共享给多个会话一起读，但是如果要更新共享锁里面的数据，需要等共享锁释放后才可以更新）
lock tables tableName read
# 给指定会话指定表加入读锁：local 关键字
lock tables tableName read local
# 一次锁定多个表的读锁
lock tables tableNameA read, tableNameB read
# 【查看哪些表锁住了】:in_use表示有多少个事务当前在锁住这个表
show open tables where in_use>1  
# 查看当前会话ID:每个窗口（会话）查询结果都不一样
select connection_id() # 2002469
# 【查看所有的事务，可以看到语句，可以看到锁定的语句】
SELECT * FROM INFORMATION_SCHEMA.INNODB_TRX;
# 查看正在等待锁定的语句（前面有锁在排队，我在等它释放资源然后抢占锁）
SELECT * FROM INFORMATION_SCHEMA.INNODB_LOCK_WAITS;
#给整个表加排他锁（写锁），加入了写锁的话只有我当前的会话（事务）可以读，写，删这个表（如果没有指定索引的话就是全表锁，不然就是锁索引）的数据
SELECT * from t_product_property  FOR UPDATE;
#给整个表加共享锁（读锁），除了我可以读，其他的事务也可以进来一起读，但是其他的事务都不允许【写】
SELECT* from t_product_property LOCK IN SHARE MODE;
```



### 获取指定时间年月日

```c#
 SELECT DATE_FORMAT('2021-08-24 05:56:18', '%Y-%m-%d'); //2021-08-24
```



### 更改Mysql密码

**1.停止MySQL服务，设置无密码登录**
1、停止 ：mysql8 net stop mysql8
2、无密码启动： mysqld --console --skip-grant-tables --shared-memory
3、前面窗口不能关闭，在开启列外一个新窗口进行无密码登陆 mysql -u root -p

4、清空密码： update mysql.user set authentication_string='' where user='root' and
host='localhost'
5、刷新 ： FLUSH privileges; 

6、重新启动mysql服务（net start mysql），再以无密码登陆mysql
**2.登录后使用SQL修改密码**
alter user root@localhost identified by '123';

### SQL语句中 where 和 on 的区别

[SQL语句中 where 和 on 的区别](https://cloud.tencent.com/developer/article/1455864)

### join 与innner join 的区别

```c#
//没有区别,join 是Inner join 的缩写,下面这两句的含义是一样的
select * from tableA  a inner join tableB b on a.id=b.id 
select * from tableA  a join       tableB b on a.id=b.id
```



### MySql分布式数据库

[MySql分布式数据库](https://zhuanlan.zhihu.com/p/403410971)

### 删除大批量数据的方法

有时候truncate table 会出现删除很慢的情况（超过20秒就可以重启Navicat），重启后删除就会很快（truncate table一般是毫秒级别的删除）

![1660095052511](Vue进步.assets/1660095052511.png)

### 赋值0值为默认值

**按照这个设置：**

![1679019009988](Vue进步.assets/1679019009988.png)

**存在两种情况**

1. **属性已存在**：如果是对已经存在的属性做这个操作，并且属性中存在null,0,1值，那么属性原有的null值不会自动更新为0值

​      这种情况下我发现不需要设置原有的值为0，不需要设置参数为not null ，查询的时候也会用到索引，即使ContainInfringe这个属性里面是有null值的，但是有Null值，用IFNULL(ContainInfringe,0)的时候扫描行数会非常大

2. **属性不存在**：如果是在新增属性的时候做这个操作，那么新增后所有的值都会赋值为默认值0，如果



**注意：**不要自作聪明在，默认：null那里更改成0，与下面的填充零冲突了，有需要设置为默认值1，或者2的时候才需要设置





### 获取字符串长度与字节长度函数

```c#
	// 获取字符长度:单位为字符，不管汉字还是数字或者是字母都算是一个字符
	select char_length('在线产品[2]') -- 7
	// 获取字节长度 :utf8编码下,一个汉字三个字节，一个数字或字母或一个符号一个字节。gbk编码下,一个汉字两个字节，一个数字或字母一个字节
	select length('在线产品[2]') -- 15
```



### 停止正在执行的SQL

```c#
// 查询所有正在执行的进程
show full processlist
// 找到对应正在执行的语句后，拿到语句的ID列，执行KILL关键字，停止SQL
KILL 3165257
```



### 格式化日期的年月时间

```c#
//获取日期的年月 202208
select date_format('2022-08-03 11:21:06, ‘%Y%m’ )

```



### 清理索引碎片

- 索引碎片对频繁删除，更新的表来说碎片会很大，最好一周，或者一个月清理一次频繁更新，删除的表的碎片，提高索引扫描的性能

- alter table语句可以触发表重建 ，比如：alter table test.userc modify column name varchar(20)  not null default '';

```c#
//清理索引碎片（这条语句不只是可以整理碎片空间大小，还可以减少索引空间的大小，并且mysql 5.6版本后，alter table ...是不会加入共享锁与排他锁的，讲大白话就是不会影响线上的使用（但是最好还是在下班后操作））
alter table t_product_property engine=innodb
//查找所有得碎片表
SELECT table_schema, TABLE_NAME, concat(data_free/1024/1024, 'M') as data_free
FROM `information_schema`.tables
WHERE data_free > 3 * 1024 * 1024
  AND ENGINE = 'innodb'  
ORDER BY data_free DESC

```

[索引锁片扩展阅读：几种碎片整理的方法](https://blog.csdn.net/userfrank/article/details/124711275)

[Mysql中删除数据造成的碎片使用alter table进行修复](https://wenku.baidu.com/view/2081673b5aeef8c75fbfc77da26925c52dc5915b.html)



### INNER JOIN 的表链接使用的发现

![1658903897637](Vue进步.assets/1658903897637.png)

### Case When Then的使用

![1658732602386](Vue进步.assets/1658732602386.png)

### 默认填充日期的选项

![1657596594494](Vue进步.assets/1657596594494.png)

### 填充零的选项

填充零后，已存在表里面的历史数据填充零的值不会自动更新为0，新增的数据会自动更新为0

（我记得有个选项会影响值的赋值，是不是无符号的选项，待定）

![1657591233516](Vue进步.assets/1657591233516.png)

### MySql转换数据类型

```c#
select CAST('0005' as  SIGNED)
select CONVERT('0005',SIGNED)
```

[mysql中Cast与Convert转换数据类型](https://baijiahao.baidu.com/s?id=1731226378554209465&wfr=spider&for=pc)

### MySql添加索引

```
ALTER TABLE `table_name` ADD INDEX index_name ( `column` ) 
```

### MySql创建临时表

```mysql
//在同一个query语句中，你只能查找一次临时表。	
create TEMPORARY  TABLE tmp_badlink2 as 
select a.ProductSku from tmp_loselink_sku 
```
### MySql回滚数据

```
show binary logs;  --查看回滚日志
```

[参考链接](https://blog.csdn.net/sinat_38371226/article/details/120924032)

### Mysql Substr函数



[参考链接](https://www.w3cschool.cn/mysql/mysql-3d9x27pc.html)

### Mysql IndexOf函数

- ## LOCATE(substr,str)

  返回子串substr在字符串str第一个出现的位置，如果substr不是在str里面，返回0.

  ```
      mysql> select LOCATE('bar', 'foobarbar');  
      -> 4  
      mysql> select LOCATE('xbar', 'foobar');  
      -> 0
  ```

### Mysql日期相减，获得小时

```
//结果：10
select TIMESTAMPDIFF(HOUR,'2022-05-18 23:30:00','2022-05-19 09:30:00')


```



### 监控腾讯云的慢SQL

[监控腾讯云的慢SQL,看数据库的负载](https://console.cloud.tencent.com/dbbrain/performance/analysis?instId=cdb-ezvn7csl)

### MYSQL表关联的删除

```mysql
#需要嵌套一层B表才能正常更新，让数据库认为你不是查同一表的数据作为同一表的更新数据源 （临时表中这种写法不支持，临时表不允许在单行语句中打开两次）

delete from temp_product_push_ec where ProductSku in ( select ProductSku from (select ps_sku as ProductSku from t_product_sku a INNER join temp_product_push_ec b 
on a.ps_sku=b.productsku where is_deleted=1) b )
-- 删除order表的数据，注意：通过别名删除
DELETE o 
FROM `order` o
	LEFT JOIN customer c ON o.customer_id = c.id 
WHERE o.customer_id =1

-- 删除customer表数据，注意：通过别名删除
DELETE c 
FROM `order` o
	LEFT JOIN customer c ON o.customer_id = c.id 
WHERE o.customer_id =1

-- 同时删除order、customer表数据，注意：通过别名删除
DELETE o,c 
FROM `order` o
	LEFT JOIN customer c ON o.customer_id = c.id 
WHERE o.customer_id =1
-- 同时更新两个表
UPDATE wx_customer wc
LEFT JOIN `order` o ON wc.id = o.wx_customer_id 
SET o.wx_customer_id = 12, wc.customer_id = 13
WHERE	wc.id = "937680";

```

[MySQL之多表关联删除/更新](https://blog.csdn.net/u011072936/article/details/124616996)

### MYSQL IsNULL事件

 

```
//返回 1
IFNULL(1,0) 
//返回0
IFNULL(null,0) 
```



### 【Mysql解决死锁的情况】

```
解除正在死锁的状态有两种方法：

第一种：

1.查询是否锁表

show OPEN TABLES where In_use > 0;

2.查询进程（如果您有SUPER权限，您可以看到所有线程。否则，您只能看到您自己的线程）

show processlist

3.杀死进程id（就是上面命令的id列）

kill id

 

第二种：

1.查看下在锁的事务 

SELECT * FROM INFORMATION_SCHEMA.INNODB_TRX where trx_query like '%t_infringe_keyword_scan_temp%';


2.杀死进程id（就是上面命令的trx_mysql_thread_id列）

kill 线程ID
```

[查看解决死锁的方法](https://www.cnblogs.com/zery/p/11801168.html)

### Mysql雪花ID

**雪花ID**:前后不连续不重复无规律的id:一串18位长度的long值,优点：全局唯一和有序递增

**使用场景：**

结合业务进行分析
一般主流的会用自增ID与GUID生成主键，自增的优点是有序，插入索引的时候也是有序，缺点是移植旧库数据的时候可能会重复，扩展性也不好，分库分表的数据主键ID可能会重复   ，而且对于数据敏感场景不宜使用（被竞争对手公司抓取到数据 很容易根据ID的增长分析你的业务增长情况），且不适合于分布式场景
GUID的优点是不重复，扩展性好，分库分表，移植数据随便移植，缺点是无序，索引也是无序，插入的时候效率不好
雪花ID的优点是不重复，而且有序，不过如果电脑的时间不稳定，可能也会重复

**总结：**

如果是数据量量大的表可能上千万，或者主业务表，为了扩展性，可以用雪花ID,如果是普通的业务表，数据几百条，几十万，几百万条，完全可以用自增的int（8位数就可以，8位数数据已经过了千万）

雪花id是数字，是不断增长的，既能做聚集索引，也能满足分布式唯一,而且还避免了连续增长，电商领域，你也不希望自己的真实数据暴露了,自增的ID  如果被竞争对手公司抓取到数据 很容易根据ID的增长分析你的业务增长情况

[雪花ID扩展](https://zhuanlan.zhihu.com/p/374667160)

[雪花ID扩展](https://baijiahao.baidu.com/s?id=1695003123690346223&wfr=spider&for=pc)

### Mysql优化Order by哪个表的主键ID

```
SELECT `a`.`ReferenceLink`, `a`.`Id` AS `Id0`, `b`.`Id` AS `PropertyId`, `b`.`SelfCategoryList`, `b`.`ProductSpu`, `b`.`ProductSku`, `b`.`ProductNameCn`, `b`.`ProductNameEn`, `b`.`ProductStatus`, `b`.`SaleStatus`, `b`.`ProductType`, `b`.`ReCostPrice`, `b`.`CostPrice`, `b`.`CostPriceCurrency`, `b`.`ReWeight`, `b`.`ProductWeight`, `b`.`ProductNetWeight`, `b`.`ProductLength`, `b`.`ProductWidth`, `b`.`ProductHeight`, `b`.`PackageWeight`, `b`.`PackageLength`, `b`.`PackageWidth`, `b`.`PackageHeight`, `b`.`DefaultImage`, `b`.`LastModificationTime`, `b`.`LastModifierUserName`, `b`.`CreateUserName`, `b`.`CreationTime`, `b`.`SelfPropertyList`, `a`.`CategoryId`, `a`.`CategoryNamePath`, `a`.`EcPushTime`, `a`.`SubSaleUserName`, `a`.`MainSaleUserName`, `a`.`BuyerUserName`, `a`.`DeveloperUserName`, `a`.`UserOrganizationId`, `a`.`DefaultSupplierName`, `a`.`DefaultSupplierCode`, `a`.`IsQc`, `a`.`DefaulBuyWarehouseCode`, `b`.`Delivery`, `b`.`Moq`, `b`.`SpecificJson`, `a`.`MainSaleUserId`, `a`.`DeveloperUserId`, `a`.`SubSaleUserIds`
FROM `t_product` AS `a`
INNER JOIN `t_product_property` AS `b` ON `a`.`ProductSpu` = `b`.`ProductSpu`
WHERE ((`a`.`IsDeleted` = FALSE) AND (`b`.`IsDeleted` = FALSE))  
ORDER BY `b`.`Id` DESC
LIMIT 20 OFFSET 1
 
```

![1648182088905](Vue进步.assets/1648182088905.png)

### MySql优化记录，表数据大于80万很慢

初次加载前20条数据，使用Limit



我现在有这些优化方向：

- 使用count(id) 比count(1)要快
- 把表insert into copy 到新表 ，同等的数据，效率直线上升（但是把ID变成主键后，效率直线下降，不知道为什么）
- 把SELECT COUNT(1) FROM TABLE 换成 SELECT cOUNT(1) FROM table where pk>0 要快
- 根据实际情况使用int，设置长度8（8刚好是千万级别的数据，因为数据过千万的话，肯定就是分表查询）
- 表里面的大属性也会影响表的查询速度，可以根据实际的使用调整属性表中的字段所占用的空间

efcore的Count是用：count(*)

```
SELECT COUNT(*)
FROM `t_product` AS `a`
INNER JOIN `t_product_property` AS `b` ON `a`.`ProductSpu` = `b`.`ProductSpu`
WHERE `a`.`IsDeleted` = FALSE
```

查询结果每次都不一样，有可能是腾讯云内部的缓存

建立索引会影响查询的结果，不一定，我看到t_product_sku这个表里面的索引有30个左右，但是我的表里面只有6个索引，也是这样慢





### MySql Case When

和sql server一样，注意QuotedClothingStyleFlag=0（单个等号，不是多个）

```
select case  when QuotedClothingStyleFlag is   null or QuotedClothingStyleFlag=0 then '否' else '是' end  as 是否引用ClothingStyleFlag from t_lister_category_specific_default  limit 10 order  by categoryid,PlatformName,SiteName
```



### MySql语句跟踪工具

[mysql监控工具sqlprofiler，类似sqlserver的profiler工具](https://www.cnblogs.com/huangshuqiang/p/12964071.html)

### 获取当前时间

```
select now();
```

### 获取UUID

```
select uuid()
```



###   【复制copy表的数据到新表】

```c#
//创建物理表：一直存在，不会删除
create table dust select * from student;
// 创建临时表：存在会话中，自动删除
CREATE TEMPORARY TABLE dust select * from student;
```



### 使用数量小的表进行order by 

用数据量小的那个表进行Order by id（如果可以选择的话） 

```c#
EXPLAIN SELECT `t`.`Id`, `t`.`SelfCategoryList`, `t`.`ProductSpu`, `t`.`ProductSku`, `t`.`ProductNameCn`, `t`.`ProductNameEn`, `t`.`ProductStatus`, `t`.`SaleStatus`, `t`.`ProductType`, `t`.`ReCostPrice`, `t`.`CostPrice`, `t`.`CostPriceCurrency`, `t`.`ReWeight`, `t`.`ProductWeight`, `t`.`ProductNetWeight`, `t`.`ProductLength`, `t`.`ProductWidth`, `t`.`ProductHeight`, `t`.`PackageWeight`, `t`.`PackageLength`, `t`.`PackageWidth`, `t`.`PackageHeight`, `t`.`DefaultImage`, `t`.`LastModificationTime`, `t`.`LastModifierUserName`, `t`.`CreateUserName`, `t`.`CreationTime`, `t`.`SelfPropertyList`, `e`.`CategoryId`, `e`.`CategoryNamePath`, `e`.`EcPushTime`, `e`.`SubSaleUserName`, `e`.`MainSaleUserName`, `e`.`BuyerUserName`, `e`.`DeveloperUserName`, `e`.`UserOrganizationId`, `e`.`DefaultSupplierName`, `e`.`DefaultSupplierCode`, `e`.`IsQc`, `e`.`DefaulBuyWarehouseCode`, `t`.`Delivery`, `t`.`Moq`, `t`.`SpecificJson`, `e`.`MainSaleUserId`, `e`.`BuyerUserId`, `e`.`SubSaleUserIds`
FROM `t_product` AS `e` 
INNER JOIN (
    SELECT `e0`.*
    FROM `t_product_property` AS `e0`
    WHERE (@__ef_filter__IsSoftDeleteFilterEnabled_1 = FALSE) OR (`e0`.`IsDeleted` = FALSE)
) AS `t` ON `e`.`ProductSpu` = `t`.`ProductSpu`
WHERE (@__ef_filter__IsSoftDeleteFilterEnabled_0 = FALSE) OR (`e`.`IsDeleted` = FALSE)
ORDER BY `e`.`Id` DESC  // 【表2数据量44,585条，order by 查询速度0.093秒】
ORDER BY `e0`.`Id` DESC  //【表e0数据量724,586条，order by 查询速度15秒】
LIMIT 20 OFFSET 1

-- 44585
select count(1) from t_product
-- 724586
select count(1) from t_product_property
```





### limit的含义

```c#
LIMIT 20 OFFSET 1   --从第一条数据开始读，读取20条数据
LIMIT 2,1      --从第二条数据开始读，读取一条数据
```



### 字符串包含

```c#
// 如果字符串 string 包含 subStr
locate(subStr,string) > 0
 
// 如果字符串 string 不包含 subStr
locate(subStr,string) = 0
    
//使用方法
SELECT * 
FROM party_course_study
WHERE LOCATE(findCode, '00001') >0
 
// 注：Mybatis使用场景，需要加 <![CDATA[ ]]>
SELECT * 
FROM party_course_study
WHERE <![CDATA[ LOCATE(findCode, '00001') > 0 ]]>
```

 [MySQL-locate()函数](https://blog.csdn.net/m0_37524661/article/details/87859423)

### 【Mysql两表数据的更新】

```

UPDATE student a

INNER JOIN temp b ON a.ID = b.ID

SET b.Name = a.Name 

```



### 通过SQL生成下拉的内容：

```c#
//拼接字符串1
select CONCAT('<el-option label="',CONCAT(CONCAT(' ',a.dic_text),concat(CONCAT('" value="',a.dic_value),'" />'))), a.dic_text,a.dic_value,a.* from bailun_skums.t_product_dictionary a where dic_name ='易仓自定义属性'
//拼接字符串2 ：拼接多个字符
    select  supplierCode,SupplierName,CONCAT(CONCAT('[',supplierCode,']'),SupplierName) from bailundb.scm_supplier;

```



### 查看Mysql版本： 

```c#
//公司使用的版本：5.7.18-txsql-log
select version() ; 
```



### Sql优化是小表驱动大表：

**inner join ：** 坐标是

**left join ：**

**right join ：**右边表要是小表，左表是大表

**使用语句检验：**EXPLAIN select * from A a left join B b on a.code=b.code，通过EXPLAIN分析来判断在sql中谁是驱动表，EXPLAIN语句分析出来的第一行的表即是驱动表

[mysql驱动表与被驱动表及join优化](https://blog.csdn.net/qq_34452824/article/details/119379559?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.pc_relevant_default&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.pc_relevant_default&utm_relevant_index=2)



**查询非空值**： LENGTH(trim(a.name))>0   



 定义复合主键（这个方法不行）

```c#
//脚本
alter table t_lister_category_specific_mapping add primary key (CategoryId,SpecificNameCn);
//执行结果
alter table t_lister_category_specific_mapping add primary key (CategoryId,SpecificNameCn)
> 1068 - Multiple primary key defined (英文含义是： 出现了错误，定义了多个主键)
> 时间: 0.007s


```



## Navicat

### 导入Excel信息失败，提示

错误提示：所使用的表类型的最大行大小(不包括blob)是65535。这包括存储开销，请查看手册。您必须将一些列更改为TEXT或blob

错误原因：因为有大字段，超过 了导入的255的字符值，所以导入失败

解决方案 ：更改大字段的存储类型为text即可，或者剔除不需要的那些大字段

![1656918761035](Vue进步.assets/1656918761035.png)

### 展示/显示隐藏的数据库

![1644891004235](Vue进步.assets/1644891004235.png)

**自动联想属性名称：**

![1639729251981](C:/Users/Administrator/Desktop/Vue%E8%BF%9B%E6%AD%A5.assets/1639729251981.png)

**筛选替换Where** ：筛选表后，打开表，查询值的时候，直接通过左上角的筛选功能，不用写where条件 ，查询数据更加快捷

**添加注释**：**Ctrl+ /**

**激活地址**：https://www.cnblogs.com/wei9593/p/11907307.html  [Navicat Premium 12破解版激活（全新注册机）](

**生成数据库文档：**

[MySQL数据库生成数据库说明文档](https://www.cnblogs.com/hsybs/p/11196673.html)

**获取当前时间：**

```
select  now()
```

**建立索引：** 列上面有了单独的索引，不需要建立复合索引

![1638929987511](D:\坚果云同步\我的坚果云\Rolan菜单\Typora图片文件夹\1638929987511.png)

**查询字节数：**

```c#
//长度为6位
select  CHAR_LENGTH（"中国时代00"） 
//字节为18位,length中文3字节，数字，字符，英文一个字节
select LENGTH（"中国时代00"）    
```

**连接MySql:**

```
mysql -u root -pmap
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





## Web First

**链接Mysql的字符串**：

```
server=gz-cdb-ezvn7csl.sql.tencentcdb.com;Port=58892;Database=bailun_bltpro;Uid=root;Pwd=shengye2021
```

[[MySQL\] - MySQL连接字符串总结](https://www.cnblogs.com/hcbin/archive/2010/05/04/1727071.html)

## JS

### JS创建数组

```
var array=[0,1,2,3]
```



### JS浅拷贝（只要值，不需要引用）

```
var newJson=JSON.parse(JSON.stringify(obj))
```



### 6种For循环

[6种For循环的使用](https://www.jb51.net/article/244117.htm)

### vs编辑器的方法提示

![1655694554102](Vue进步.assets/1655694554102.png)

![1655694762264](Vue进步.assets/1655694762264.png)

### Object.keys

```c#
//拿到json所有的key
let obj = { name:"张三", age:25, address:"深圳"} 
console.log( Object.keys(obj) ) //{name,age,address}
```



### Object.values

```c#
//拿到Json所有的values
let obj = { name:"张三", age:25, address:"深圳"} 
console.log( Object.values(obj) ) //{"张三",25,"深圳"}
```



### 判断属性为空的方法

```c#
//这种可以判断空字符串，null,undefined
if (!this.multipleCondition) {
        this.$message.warning('请选中下拉条件后再输入值')
        return
      }
  //上面等于下面的条件  
 if(this.multipleCondition== null || this.multipleCondition === undefined||this.multipleCondition==='')
 // 如果是Int类型的话等于0也是不存在的值，比如下面
 var total = 0  //注意整型的值，默认值没有null只有0, 0就是不存在
 // 【total等于0的时候用if判断total是不会进来的】
 if (total) {
 this.$message.warning('值存在')
 return
 }
  var total = 1
 //这样判断就会进来了
 if (total) {
 this.$message.warning('值存在')
 return
 }  
```



### 字符串排序的方法

```c#
//使用中文进行排序 --升序
_this.organizationNameList = response.result.sort((a,b) => { return a.text.localeCompare(b.text,"zh"); });
//升序
a.localCompare(b) 
//降序
b.localCompare(a) 
//默认排序 ,数字，字母
array.sort()
```

[排序参考链接](https://wenku.baidu.com/view/2081673b5aeef8c75fbfc77da26925c52dc5915b.html)

### JS中把字符串转换成为Int

```
var j = parseInt("11")
```



### JS中问号？的使用

? 符号 可以在数组中使用，可以在js(script)作用域里面使用，不能在Vue的UI界面中用

script作用域：可以用

![1645166504120](Vue进步.assets/1645166504120.png)

UI界面：不能用

![1645167996309](Vue进步.assets/1645167996309.png)

### JS里面==，与===的区别

**==符号：**先检查两个操作数的数据类型是否相同
如果相同，则比较两个数是否相等
如果不同，则先将两个数转换为相同数据类型，再进行比较

![这里写图片描述](https://img-blog.csdn.net/20180908151605795?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dlaXhpbl80MjQyMDcwMw==/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

**===符号**：先检查两个操作数的数据类型是否相同
不一样，直接返回false
一样，再比较两个值是否相等

![1643178906481](Vue进步.assets/1643178906481.png)



```c#
//有趣的例子：
null==undefined //true
null===undefined //false
//【注意NaN与任何值都不相等，包括NaN自身】而NaN（not a number）是一个特殊的数值，这个数值用于表示一个本来要返回数值的操作数未返回数值的情况，避免报错。NaN与任何值都不相等，包括NaN自身。
NaN==NaN //false
NaN===NaN //false
```



### 把数组转换为字符串（互换）

```
array = new Array(0,1,2,3,4);
var temp = array.toString()
console.log(temp) ; // 0,1,2,3,4
var temp2 = array.join('-')
console.log(temp2) ; // 0-1-2-3-4
```



### 获取数组最后一位数（返回新的数组）

不会改变原来的数组

```
 arrayNumber.slice(-1);
```

**检查值是否为 NaN**

```c#
isNaN(123) //false
isNaN(-1.23) //false
isNaN(5-2) //false
isNaN(0) //false
isNaN('123') //false
isNaN('Hello') //true
isNaN('2005/12/12') //true
isNaN('') //false
isNaN(true) //false
isNaN(undefined) //true
isNaN('NaN') //true
isNaN(NaN) //true
isNaN(0 / 0) //true
isNaN(null) //false
```

**javascript如何设置字符串首字母大写？**

```c#
//首字母大写
function titleCase(str) {
newStr = str.slice(0,1).toUpperCase() +str.slice(1).toLowerCase();
    return newStr;
}

titleCase("hello World！"); //Hello World！

//所有单词首字母大写
function titleCase(str) {    
    var newStr = str.split(" ");   
    for(var i = 0; i<newStr.length; i++){
        newStr[i] = newStr[i].slice(0,1).toUpperCase() + 		        newStr[i].slice(1).toLowerCase();
    }    return newStr.join(" ");
}

titleCase("I'm a little tea pot"); //I'm A Little Tea Pot
```

[javascript如何设置字符串首字母大写？](https://m.html.cn/qa/javascript/11395.html)

**JS转换对象成字符串：**

比如转换数组为字符串，与formData进行浅拷贝

let data = JSON.stringify(this.formData);

合并使用：

```c#
      //先转换成字符串，然后转成JSON完成浅拷贝
      const data = JSON.parse(JSON.stringify(_this.formData))
       //再对值进行处理
      data.productCharacter = data.productCharacter.join(',')
```

**JS转换字符串为Json:**

let json=JSON.parse(jsonString)

**JS转换字符串为数组：**
let array = JSON.parse(ArrayStr);

js if判断多个条件_JS条件判断小技巧（一）https://blog.csdn.net/weixin_39999586/article/details/110745829?utm_term=js%E4%B8%AD%E4%B8%94%E6%9D%A1%E4%BB%B6&utm_medium=distribute.pc_aggpage_search_result.none-task-blog-2~all~sobaiduweb~default-0-110745829&spm=3001.4430

**Js数组转成字符串**

```
var a, b;
a = new Array(0,1,2,3,4);
b = a.join("-");
```

[js数组与字符串的相互转换方法](https://www.cnblogs.com/asdyzh/p/9801784.html)

**Map** ：返回新的数组，像select的功能

```c#
//返回一个新的数组,一个属性，直接返回
var Listings =this.rows.map()(r => r.id)
//返回一个新的数据，多个属性，需要增加1.{return} 2.左右的括号 {} 3.属性的别名
var Listings =this.rows.map()(r => { return { ListingId: r.id, Sku: r.sku } })
//根据条件更改原数组的值 。map里面写if else {} 大括号是精髓 ，还有尾部的：return u
var specificNumberList = specificNumberList.map(u => { if (u.specificValue === null || u.specificValue === '') { u.specificValue = '无' } return u})


```



**JSON对象,JSON数组,JSON字符串互转及取值**

**Json字符串转换成Json前端处理：**

```
JSON.parse(jsonString)
```

**后端处理：**

```c#
 //转换成Json数组(这样也可以)
 JArray jo1 = (JArray)JsonConvert.DeserializeObject(builderJson.ToString());
 //转换成JSON对象（保留后续使用）
 JObject jo1 = (JObject)JsonConvert.DeserializeObject(builderJson.ToString());
```

原文地址：[JSON对象,JSON数组,JSON字符串互转及取值](https://www.cnblogs.com/tangbang/p/tb.html)





## Vs Code

- **管理未完成的事件** ：TODO标签：下载插件：TODO Highlight，再下载：Todo Tree 进行管理，写大写的**TODO:**

- **碰到奇怪的错误：**ctrl+c 关闭一下vs code中的站点，新建一个终端: npm run dev

- **vue devtool插件** : 在谷歌浏览商城使用，Vue中AppMain/FyndiqListerAddEdit/Eldialog 展示当前窗体Open出来组件的使用的变量情况

  ## Vs Code中快速创建Html代码

  1. 选中文本为html，先输入一个 !
  2. 点击 tab 键
  3. 自动生成标准的html代码

界面的位置：

![image-20211126170844373](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20211126170844373.png)

根据变量名称过滤变量的值：

![image-20211126171335371](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20211126171335371.png)

第二次查询界面的位置：

![1638928008299](D:\坚果云同步\我的坚果云\Rolan菜单\Typora图片文件夹\1638928008299.png)

- **空格警告的处理：**1.存在空格与制表符号，**解决**：到警告提示的行处，把前部的空格删除掉，然后再按ctrl+s 保存，系统自动会用没有问题的空格占满

```
 //504行：存在两个空格
 504:2  error  Mixed spaces and tabs  no-mixed-spaces-and-tabs

 505:2  error  Mixed spaces and tabs  no-mixed-spaces-and-tabs

 624:2  error  Mixed spaces and tabs  no-mixed-spaces-and-tabs

 639:2  error  Mixed spaces and tabs  no-mixed-spaces-and-tabs
```



- **点击直达教程：** 鼠标悬停到代码处，点击出现的v-if关键字可以直接到官网的教程。
- **返回上一个编辑的界面**：**alt+ 左右箭头**   ,文件→首选项→键盘快捷方式→搜索前进/后退
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



## SqlSugar

### 数据库单例，仓储，泛型仓储

- ISqlSugarClient ：获取数据库的对象，可以从里面拿到所有的数据库对象，可以查询所有的库
- ISqlSugarRepository：指定的数据库对象仓储
- ISqlSugarRepository<>（ 数据库上下文对象定位器） ：指定了泛型，指定了对象的仓储，很方便操作对象，比如更新，删除结果

### 忽略指定属性的数据库映射

```c#
        /// <summary>
        /// Desc:属性json（从数据库实体中忽略该值的读取）
        /// Default:
        /// Nullable:True
        /// </summary>           
        [SugarColumn(IsIgnore = true)] //有可选的选项：IsOnlyIgnoreInsert(可以插入) 和IsOnlyIgnoreUpdate(可以更新) 两个属性
        public string SpecificJson { get; set; }
```



### 使用参数化方式查询

```c#
 query.Where("ProductSpu in (SELECT ProductSpu FROM `t_product` WHERE `CreationTime` >= @CreationTimeStart and `CreationTime`<=@CreationTimeEnd and IsDeleted=0)", new { CreationTimeStart = input.CreationTimes[0].Value, CreationTimeEnd = creationTimesAddDays });
```



### 使用SqlSugar

注入SqlSugar

![1665628946150](Vue进步.assets/1665628946150.png)

### 使用官方文档中的实例

![1661410451065](Vue进步.assets/1661410451065.png)

### 快速打印当前sql

​            var sql = query.ToSql();

**异常捕获问题：**数据表属性类型不匹配不会抛异常

![1638782723851](D:\坚果云同步\我的坚果云\Rolan菜单\Typora图片文件夹\1638782723851.png)

### SqlSugar输出Sql

```c#
//执行语句
    DBClient.Updateable<FyndiqListerVariation>()
            .SetColumns(it => new FyndiqListerVariation() { PublishStatus = 4, PublishTime = DateTime.Now, PublishMsg = PublishMsg })
            .Where(it => it.Id == variationId.Value)
            .ExecuteCommand();
//方法1执行完成后，输出日志与参数
 DBClient.Aop.OnLogExecuted = (sql, pars) =>
  {
     //下面两个输出方法推荐用Fail,WriteLine不清楚为什么有时候输出不了日志
   System.Diagnostics.Debug.WriteLine(sql);
   System.Diagnostics.Debug.Fail(sql);
  };



            //方法2：执行完成后，输出日志与参数
            _productPropertyRepository.Context.Aop.OnLogExecuted = (sql, pars) =>
            {
                //下面两个输出方法推荐用Fail,WriteLine不清楚为什么有时候输出不了日志
                System.Diagnostics.Debug.WriteLine(sql);
                System.Diagnostics.Debug.Fail(sql);
            };

        /// <summary>
        /// 【方法3】打印所有执行的sql
        /// </summary>
        /// <param name="sqlSugarRepository"></param>
        /// <param name="db"></param>
        public SqlSugarRepository(ISqlSugarRepository sqlSugarRepository
            , ISqlSugarClient db)
        {
            _sqlSugarRepository = sqlSugarRepository;

            DynamicContext = Context = _db = (SqlSugarClient)db;
            Ado = _db.Ado;
            //方法3：打印执行sql
            _db.Aop.OnLogExecuting = (sql, pars) =>
            {
                System.Diagnostics.Debug.Fail(sql);
                Console.WriteLine("********打印执行sql*************");
                Console.WriteLine(sql);
            };

        }
```





## Nginx

公司应该是把Nginx作为反向代理的服务器，做为入口的转发，公司请求路径：

发起请求->nginx(内，外)->内网关（Ocelot）->内consul(服务注册)->请求对应服务

**代理服务器**：A服务器到B服务器中间的跳板服务器，可以看做婚姻介绍中的媒婆

- 正向代理(隐藏请求客户端)：内部的服务器A，通过Nginx服务器B，请求互联网服务器C
- 反向代理（隐藏接收请求客户端）：客户通过浏览器访问服务接口C,服务访问Nginx服务器B，Nginx通过负载均衡访问内部服务器A

**负载均衡**（缺点不能动态伸缩）：三台服务器A,B,C, 请求三次，分别落在A,B,C三台服务器上面，如果其中B服务器挂了，那么A,C服务器会正常承担请求

nginx有三种配置请求的方法：

- **普通轮询**：一种是轮询就是先请求A,然后下一次请求
- **配置权重轮训**：比如有些服务器配置比较好，那么权重配置高一些，承担百分之50的请求，有些服务器配置比较低，那么配置承担百分之20的请求等
- **请求数** ：根据服务器当前的请求数量进行判断，将当前请求放到请求少的服务器上面（因为有些请求特别耗时，可能会积累在某台服务器上面）
- ip Hash（特殊作用）：有状态服务，当服务挂了需要手动剔除，单个用户的请求一直落在同一台服务器，适用于用户分片上传文件到服务器，需要在同一个服务器合并的情况

### 基本使用

**什么是Nginx**

Nginx是一个高性能的http和反向代理服务器

**为什么要用Nginx**

- 多台服务器可以用Nginx负载均衡，将请求均匀的发送到多台服务器上面

- 可以当做静态文件服务器进行使用

- 代理服务器

  

**Nginx特点**:

- 高并发连接：官方称单节点支持5万并发连接数，实际生产环境能够承受2-3万并发。

- 开源免费

- 不吃服务器配置

- 配置使用简单

  

  **Consul与Nginx 有什么区别？**

  nginx是代理服务器，高并发，高性能，高稳定性

  consul是接收请求后，将请求转发到对应的服务请求，专门做服务注册与发现，可以使用代码进行动态注册，动态注销，负载均衡，配置心跳检查

  **不同点：**
  consul可以动态的去注册服务，发现服务，nginx只能支持配置的方式更改
  nginx性能比consul高，单点并发达到了5w，并且nginx消耗的服务器硬件低

  **相同点**
  consul与nginx都可以做负载均衡

  [微服务架构-Consul集群、Ocelot服务和Nginx网关服务](https://www.cnblogs.com/Leo_wl/p/14043862.html)

[好的Nginx入门文章](https://blog.51cto.com/u_14820531/2512102)

## CMD

nslookup：判断域名是否可以正常访问，是不是被墙了

## **Ocelot**

网关可以做**限流，超时，网关，缓存，聚合**请求。在config配置里面配置超时，缓存，限流的信息，可以配置Consul的访问链接由它进行转发到具体的WebApi，配合应用程序实现**动态伸缩服务器资源的功能**；

Ocelog是Net开源项目，可以很方便的进行扩展，微软，腾讯都在用的

## ES:ElasticSearch

[原文部署攻略](https://blog.csdn.net/qq_49059667/article/details/120355033)

### Docker安装kibana

```mysql
# 从Docker拉取kibana
docker pull kibana
# 启动容器（下面命令行没有换行符），-link es 连接elasticsearch镜像，es是你运行elasticsearch时候起的名字，注意elasticsearch.hosts【中的IP地址保存成你电脑的ip地址】，我发现在latest7.17.8的版本中可以直接用http://elasticsearch:9200，但是在7.13.1用这个地址就不可以了，要用具体的ip地址：http://192.168.226.128/:9200
docker run -p 5601:5601 --name kibana --link es -e "elasticsearch.hosts=http://192.168.226.128/:9200" -d kibana:latest 
#测试访问,ifconfig看ip
输入：http://192.168.226.128:5601/    -- 5601 这个Ip是看最上面-p 9200 左边的这个5601
看到kibana图形化界面表示安装成功
# 部署过程还需要更改 kibana/kibana.yml 文件中配置的路径，由http://elasticsearch:9200 更改为：http://192.168.226.128/:9200
 -- 1.先启动kibana容器服务 docker start kibana
 -- 2.然后进入kibana容器 docker exec -it kibana bash
 -- 3.进入config文件夹 cd config
 -- 4.更改kibana.yml文件 :使用 ls,cd config,vi kibana.yml,:wq进行保存 ，加入i18n.locale: "zh-CN"这一行，放到hosts下面那行就可以，设置为中文配置
 -- 5.重启服务：docker restart kibana
【有一个插曲是从阿里云镜像下载的kibana与Es的镜像使用latest不是最新的，最好是指定版本下载对应的镜像，或者下载后使用命令检查一下下载镜像的版本，以免出现未知问题】
```

[rm更改kibana异常排查](https://www.jianshu.com/p/b413df10fbe8)

### 使用守护进程启动kibana

```mysql
#原来启动命令：启动kibana
cd /usr/local/elastic/kibana-8.6.1
启动：su elk -c ./bin/kibana
#使用守护进程自动启动
#以编辑模式创建kibana.service服务文件（没有任何后缀就是一个文件）
 sudo vi /etc/systemd/system/kibana.service
 -- 粘贴这里面的内容，到linux窗口中，注意粘贴后检查一下是不是一致，尤其是头部的[Unit]
[Unit]
Description=Kibana

[Service]
Type=simple
ExecStart=/usr/local/elastic/kibana-8.6.1/bin/kibana
Restart=always
User=elk

[Install]
WantedBy=multi-user.target
/*
解释：
[Unit]：这是一个systemd单元文件的标头，用于定义一个系统服务。
Description=Kibana：描述Kibana服务的名称或简单描述。

[Service]：这个标头定义了要管理的服务的设置。
Type=simple：定义了如何管理服务的进程。simple类型表示服务的进程是systemd管理的主进程。

ExecStart=/path/to/kibana/bin/kibana：定义了服务的启动命令，即启动Kibana的命令。你需要将/path/to/kibana替换为实际的Kibana安装路径。

Restart=always：定义了当服务进程异常退出时如何处理。这里是将服务进程自动重启。
User=elk：定义了服务进程运行的用户。在这里，Kibana服务将作为elk用户运行。

[Install]：定义了如何安装这个服务，包括安装到哪个运行级别并在哪些服务之后启动。

WantedBy=multi-user.target：定义了在哪个运行级别（或多个运行级别）中安装此服务。这里是安装到multi-user.target运行级别，这是Linux系统默认的图形用户界面运行级别，它在启动时加载多个用户服务。
*/
# 输入:wq保存配置文件，然后使用以下命令刷新linux系统配置
sudo systemctl daemon-reload
# 设置kibana服务的自启动（如果失败看下面）
sudo systemctl enable kibana.service
# 启动/重启kibana服务
sudo systemctl start/restart kibana 
# 检查服务状态
sudo systemctl status kibana

# 注意：如果启动失败，或者设置自启动失败：
sudo journalctl -u kibana.service  # 查看日志
sudo journalctl -u kibana.service -n 50  #查看最近50条日志


比如：Assignment outside of section. Ignoring.（节外的分配。无视。）
这个就是配置文件的格式不对，检查一下配置文件，然后重新刷新 linux系统的配置，走下面的流程
# 查看系统的服务启动日志，可能有更加详细的日志记录
sudo journalctl -xe


```



### Docker安装Es

```mysql
# 从Docker拉取Es,发现拉取指定版本的时候出现超时的问题，拉取默认latest版本的时候成功，或者可以多试两下
docker pull elasticsearch  -- 成功
# 启动容器，配置账户 --name  es 代表启动容器的名字，-d  代表使用的镜像， -e 代表参数,xms256m最小内存，后面的mx256m是最大内存,dis...=single-node 代表的是单节点集群
docker run --name es -p 9200:9200 -p 9300:9300 -e ES_JAVA_OTPS="-Xms256m -Xmx256m" -e "discovery.type=single-node"   -d elasticsearch:latest
# 测试访问,ifconfig看ip
输入：http://192.168.226.128:9200/    -- 9200 这个Ip是看最上面-p 9200 左边的这个9200
{
  "name" : "k2IQVK6",
  "cluster_name" : "elasticsearch",
  .... -- 部署成功
```

### Docker安装IK分词器

下载分词器的安装包后，将zip安装包放到linux服务器里面,IK分词是针对中文,英文检索不会分词

![1670897419796](Vue进步.assets/1670897419796.png)

看上面部署的链接，看怎么安装IK分词器

 

### Es介绍

**ELK**：一般说Es都是说这三种：e指的就是es搜索引擎,L（Logstash）指的就是信息收集模块,K（Kibana）指的就是数据展示，分析的模块。分布式日志分析整体解决方案，解决日志分析，大数据量搜索

- **什么场景下面用Es？**

1. **高并发**：写入居多，改动比较少

2. **大数据**：数据量过3千万，过亿级别，并且不断在新增

3. **高速度**：对数据返回速度的要求很高

   对数据的及时性要求不高的（1秒之外），比如常用的解决方案有：日志模块，报表模块，搜索模块使用,更普遍应用于生产日志的查询分析中，比如用来收集nginx的转发日志，分析系统的性能故障

* **为什么要用Es？**
mysql数据库达到千万级别，三千万级别，表在使用上就非常不方便：
1. 前端首页，不可能提供count聚合查询的操作。
2. 后台查询的时候，一定要提供索引去查询
3. 没有索引的情况下，基本不能用like作为查询条件
4. 密集的io并发操作，对数据库的整体性能非常拉垮
5. 使用聚合统计时基本不可能或者性能非常非常慢。

* **Es的优点？**
1. noSql数据库Es的性能非常好，*百亿级别的存储，查询，秒级别返回*，采用的是分布式保存，分布式搜索的能力
2. 系统提供了一些常用的聚合操作，报表统计的操作，以及近三天近七天的数据。
3. 使用倒推索引性能非常好，比b+树的收益更好，性能更快百亿级别的数据可以在秒级别返回。
4. 使用数据压缩原理，使用更少的内存保存更多的数据。（最大32G内存）
5. 全文索引，分词功能
6. 配套环境，有保存数据到es，有分析查询数据的这种查询插件
7. 支撑高并发的写入，读取环境
8. 负载均衡，容错率强，主片挂了从片自动顶上
9. 倒推索引，全文索引，分词，同义词，数据压缩，自动添加索引，聚合操作友好

- **Es的缺点**
  1. **数据实时性**：及时性不好，秒级别。如果新增了一条数据，那么最少要在一秒后才可以操作这条数据。（每当有新增的数据时，就将其先写入到内存中，在内存和磁盘之间是文件系统缓存，当达到默认的时间（1秒钟）或者内存的数据达到一定量时，会触发一次刷新（Refresh），将内存中的数据生成到一个新的段上并缓存到文件缓存系统 上，稍后再被刷新到磁盘中并生成提交点。因此，从Index请求到对外可见能够被搜到，最少要1秒钟的数据延时）
  2. 不支持事务
  3. 不支持表连接 join 
  4. 字段类型建立后就不允许进行改动（一般也不会进行改动）
  5. 写入性能低
  6. **硬件资源消耗大**
  7. 表的管控没有权限管理（这个应该也不成问题，我们现在的表就没有权限管控的）

-  **其他**

  ### ES数据架构的主要概念（与关系型数据库mysql相比）

  （1）关系型数据库中的数据库（DataBase），等价于ES中的索引（Index） 

  （2）一个数据库下面有N张表（Table），等价于1个索引Index下面有N多类型（Type）

  （3）一个数据库表（Table）下的数据由多行（ROW）多列（column，属性）组成，等价于1个Type由多个文档（Document）和多Field组成。 

  （4）在一个关系型数据库里面，schema定义了表、每个表的字段，还有表和字段之间的关系。

  　　 与之对应的，在ES中：Mapping定义索引下的Type的字段处理规则，即索引如何建立、索引类型、是否保存原始索引JSON文档、是否压缩原始JSON文档、是否需要分词处理、如何进行分词处理等。 

  （5）在数据库中的增insert、删delete、改update、查search操作等价于ES中的增PUT/POST、删Delete、改_update、查GET.



### ES的使用



#### **1.安装Kibana**

参考上面Docker/安装Kibana

#### **2.进入管理面板-开发工具**

[访问Kibana](http://192.168.226.128:5601/app/dev_tools#/console)

![1675652994118](Vue进步.assets/1675652994118.png)

![1675674879069](Vue进步.assets/1675674879069.png)

#### 3.创建表/create/PUT

```c#
// 删除原来的index/表
DELETE /student
// 创建index/表
PUT /student
{
    "mappings":{ 
    "properties":{  //声明创建属性
      "name":   //属性名称
      {
        "type":"keyword"  // 数据类型
      },
      "address":   
      {
        "type":"text"
      },
      "age":
      {
        "type":"integer"
      }
    }
  }
}
// 输出成功
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "student"
}
```

#### 4.添加数据/insert/POST

```c#
// 往student  index/表中新增一条数据
POST /student/_doc
{
  "name":"张三",
  "address":"广州天河",
  "age":30
}
// 添加数据有双引号的情况，用右斜杠进行转义
POST /student/_doc
{
  "name":"张三",
  "address":"\"fda\"",
  "age":30
}
//批量导入语法
POST /student/_bulk
{"index":{"_id":1}}
{"name":"张三","address":"广州天河","age":30}  //要在一行
{"index":{"_id":2}}
{"name":"张三","address":"\"fda\"","age":30}

// 输出成功
{
  "_index" : "student",  // index表的名称
  "_type" : "_doc",  
  "_id" : "Lxq-JIYBkhLM5sijqoxY",  //默认创建唯一ID
  "_version" : 1,
  "result" : "created",  //返回创建成功
  "_shards" : {
    "total" : 2,
    "successful" : 1,
    "failed" : 0
  },
  "_seq_no" : 0,  //表里面从0生成的序列号，如果顺序是唯一的可以看作表的行数
  "_primary_term" : 1
}

```



#### 5.无条件查询数据/select/GET

```c#
// 查询索引/表全量数据  ，默认只返回十条数据
GET /student/_doc/_search

//返回结果
{
  "took" : 3,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 7,    // 总行数
      "relation" : "eq"  // 等于查询
    },
    "max_score" : 1.0,
    "hits" : [
      {
        "_index" : "student",
        "_type" : "_doc",
        "_id" : "Lxq-JIYBkhLM5sijqoxY",
        "_score" : 1.0,
        "_source" : {
          "name" : "张三",
          "address" : "广州天河",
          "age" : 30
        }
      },
      {
        "_index" : "student",
        "_type" : "_doc",
        "_id" : "MBrEJIYBkhLM5sijJ4yX",
        "_score" : 1.0,
        "_source" : {
          "name" : "王锐",
          "address" : "广州海珠",
          "age" : 28
        }
      },
      {
        "_index" : "student",
        "_type" : "_doc",
        "_id" : "MRrEJIYBkhLM5sijN4z7",
        "_score" : 1.0,
        "_source" : {
          "name" : "王五",
          "address" : "广州白云",
          "age" : 40
        }
      }
      ..... 剩余结果省略
      
// 查询前10000数据，最大只能单次查询10000条数据，可以通过更改最大限制index进行剔除这个限制
GET student/_doc/_search
{
  "query": {
    "match_all": {}
  },
  "size": 10000
}
// 返回结果与上面一致
```

#### 6.排序查询/order/sort

```c#
//排序查询
GET student/_doc/_search
{
  "sort":[
    {"age":{"order":"desc"}},
    {"name":{"order":"asc"}}
    ]
}
// 返回结果
{
  "took" : 2,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 7,
      "relation" : "eq"
    },
    "max_score" : null,
    "hits" : [
      {
        "_index" : "student",
        "_type" : "_doc",
        "_id" : "MRrEJIYBkhLM5sijN4z7",
        "_score" : null,
        "_source" : {
          "name" : "王五",
          "address" : "广州白云",
          "age" : 40
        },
        "sort" : [
          40,
          "王五"
        ]
      },
      {
        "_index" : "student",
        "_type" : "_doc",
        "_id" : "Lxq-JIYBkhLM5sijqoxY",
        "_score" : null,
        "_source" : {
          "name" : "张三",
          "address" : "广州天河",
          "age" : 30
        },
        "sort" : [
          30,
          "张三"
        ]
      }
      ...其余结果省略
```

#### 7.【筛选结果and/where/must/term】

```c#


//=/term  注意term是大小写敏感的
GET student/_doc/_search
{
  "query":{
     "term": {
      "name": "王锐"
    }
  }
}
//and/must 只查看广州白云的人员信息 
GET student/_doc/_search
{
  "query":{
    "bool":{
      "must":[
       {
          "match":{
          "address":"广州白云"
        }   //如果是多个条件，那么加入 ,"match":{ "属性名称":"筛选属性值"}
       }
      ]
    }
  }
}


// 返回结果：发现白云的在最上面，其余的结果也会出现，为什么？因为分词了
{
  "took" : 2,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 7,
      "relation" : "eq"
    },
    "max_score" : 3.4770298,
    "hits" : [
      {
        "_index" : "student",
        "_type" : "_doc",
        "_id" : "MRrEJIYBkhLM5sijN4z7",
        "_score" : 3.4770298,
        "_source" : {
          "name" : "王五",
          "address" : "广州白云",
          "age" : 40
        }
      },
      {
        "_index" : "student",
        "_type" : "_doc",
        "_id" : "Lxq-JIYBkhLM5sijqoxY",
        "_score" : 0.12907705,
        "_source" : {
          "name" : "张三",
          "address" : "广州天河",
          "age" : 30
        }
      }
      ... 省略其他搜索结果
```

#### 8.筛选结果or/where/should

```c#
 // 筛选结果 or/should
GET student/_doc/_search
{
   "query":{
      "bool": {
        "should": [
          {
     "match": {
       "address": "广州天河"
     }},{
            "match": {
       "name": "王五"
     }
          }
        ]
      }
  }
}
```



#### 9.分页查询/limit/from/size

```c#
 // 分页查询
GET student/_doc/_search
{
  "from":0,  //从第一行开始
  "size":5  //查询5条数据
}
// 结果略
```

#### 10.只查看指定属性/*/_source

```c#
// 只查询name,age
GET student/_doc/_search
{
  "_source":["name","age"]
}
//返回结果，顶部略，底部略
 {
        "_index" : "student",
        "_type" : "_doc",
        "_id" : "Lxq-JIYBkhLM5sijqoxY",
        "_score" : 1.0,
        "_source" : {
          "name" : "张三",
          "age" : 30
        }
 }
```

#### 11.筛选范围值/where(filter)/>(gt)/<(lt)/>=(gte)/<=(lte )

```c#
// 筛选姓名等于张三，年龄在20到30之间的
GET student/_doc/_search
{
   "query":{
     
    "bool": {
      "must": [
        {
           "match": {
       "name": "张三"
        
     }
        }
      ],
      "filter": { //注意filter是写在bool里面的
        "range": {
          "age": {
            "gte": 20,  //也可以写一个条件，也可以写两个条件
            "lte": 30
          }
        }
      }
    }
    
  }
}
// 返回结果 省略顶部，底部结果
"hits" : [
      {
        "_index" : "student",
        "_type" : "_doc",
        "_id" : "Lxq-JIYBkhLM5sijqoxY",
        "_score" : 1.0498221,
        "_source" : {
          "name" : "张三",
          "address" : "广州天河",
          "age" : 30
        }
      },
      {
        "_index" : "student",
        "_type" : "_doc",
        "_id" : "Nhr7JIYBkhLM5sij1owj",
        "_score" : 1.0498221,
        "_source" : {
          "name" : "张三",
          "address" : "广州天河",
          "age" : 20
        }
      }
    ]
```

#### 12.排除not in /must_not

```c#
//查询名称不是王锐的信息
GET student/_doc/_search
{
  "query":{
    "bool":{
      "must_not":[
        {
          "match":{
            "name":"王锐"
          }
        }
        ]
    }
  }
}
// 查询结果，头尾，剩余结果略
 {
        "_index" : "student",
        "_type" : "_doc",
        "_id" : "Lxq-JIYBkhLM5sijqoxY",
        "_score" : 0.0,
        "_source" : {
          "name" : "张三",
          "address" : "广州天河",
          "age" : 30
        }
      },
      {
        "_index" : "student",
        "_type" : "_doc",
        "_id" : "MRrEJIYBkhLM5sijN4z7",
        "_score" : 0.0,
        "_source" : {
          "name" : "王五",
          "address" : "广州白云",
          "age" : 40
        }
      }
```

#### 13.【筛选结果in/match /match_phrase】

```c#
//match(对于经过分词器的结果):匹配分词的结果，不会百分百吻合 查询地址是广州天河，广州白云的人员信息
GET student/_doc/_search
{
  "query":{
    "match":{"address":"广州天河 广州白云"}
  }
}
//match_phrase ：对于经过分词器的，也百分百吻合结果
GET student/_doc/_search
{
  "query":{
    "match_phrase":{"address":"广州天河 广州白云"}
  }
}
//查询单个值
GET student/_doc/_search
{
  "query":{
     "match_phrase" : {
            "address" : "rock climbing"
        }
  }
}



//返回结果，头尾剩余结果略
"hits" : [
      {
        "_index" : "student",
        "_type" : "_doc",
        "_id" : "MRrEJIYBkhLM5sijN4z7",
        "_score" : 3.9994133,
        "_source" : {
          "name" : "王五",
          "address" : "广州白云",
          "age" : 40
        }
      },
      {
        "_index" : "student",
        "_type" : "_doc",
        "_id" : "Lxq-JIYBkhLM5sijqoxY",
        "_score" : 2.3048174,
        "_source" : {
          "name" : "张三",
          "address" : "广州天河",
          "age" : 30
        }
      },
      {
        "_index" : "student",
        "_type" : "_doc",
        "_id" : "Nhr7JIYBkhLM5sij1owj",
        "_score" : 2.3048174,
        "_source" : {
          "name" : "张三",
          "address" : "广州天河",
          "age" : 20
        }
      },
```

#### 14.更新update/_update

```c#

 // 方法一：在原有PUT创建索引时直接进行覆盖操作
PUT /my_index/user/1
{
  "name":"朝夕教育-123"
}
//#缺陷：相当于再次创建没有传入的参数会为空
​
// 【用这个方法二】
POST /my_index/user/1/_update
{
  "doc":{
    "name": "朝夕教育-123"
  }
}
//没有传入的参数不会为空
```

#### 15.求最大值max/aggs

```c#
// 求最大值
GET /student/_doc/_search
{
    "size":0,
    "aggs":{
        "max_age":{
            "max":{
                "field":"age"
            }
        }
    }
} 
// 输出：公共属性省略
 "aggregations" : {
    "max_age" : {
      "value" : 40.0
    }
```

#### 16.分组查询group by /aggs

```c#
// 查询刊登数据导入情况
GET /platform_lister/_search
{
    "size" : 0,
    "aggs" : { 
        "platformSource_count" : { 
            "terms" : { 
              "size": 1000,  //默认显示10行，加上这个就可以展示全部了
              "field" : "platformSource"
            }
        }
    }
}
//结果
"buckets": [
        {
          "key": "Ebay刊登",
          "doc_count": 268748
        },
        {
          "key": "Walmart刊登",
          "doc_count": 98080
        },
        {
          "key": "Amazon刊登",
          "doc_count": 24083
        }
      ]
```

#### 18.清空所有数据(保留表结构) truncate

```c#
//清空（清空四千万walmart_listing的时候耗费时间大约1小时，删除后提示超时，提示版本不一致不要紧，点了几下删除，不用管，后台慢慢在那里删除）
POST platform_lister/_delete_by_query
{
  "query": { 
    "match_all": {
    }
  }
}
//返回
{
  "took": 15608,
  "timed_out": false,
  "total": 205448,
  "deleted": 205448,
  "batches": 206,
  "version_conflicts": 0,
  "noops": 0,
  "retries": {
    "bulk": 0,
    "search": 0
  },
  "throttled_millis": 0,
  "requests_per_second": -1,
  "throttled_until_millis": 0,
  "failures": []
}
```

#### 19.模糊查询加等于查询like += /wildcard +must

```
//查询
GET platform_lister/_search
{
  "query": {
    "bool": {
      "must": [
        {
          "match": {
            "platformSource": "Amazon刊登"
          }
        },
        {
          "wildcard": {
            "title": "*Womens*"
          }
        }
      ]
    }
  }
}

```

#### 20.根据条件删除数据 delete /_delete_by_query

```c#
//删除数据来源是沃尔玛刊登的数据
POST platform_lister/_delete_by_query
{
 "query":{
    "match":{
      "platformSource":"Walmart刊登"  
    }
  }

}
// 删除可能会有超时，并且要不断执行，要不断删除好几次才能全部删除完,而且间隔一段时间按一下删除的会多一些
{
  "took": 54,
  "timed_out": false,
  "total": 0,
  "deleted": 0,
  "batches": 0,
  "version_conflicts": 0,
  "noops": 0,
  "retries": {
    "bulk": 0,
    "search": 0
  },
  "throttled_millis": 0,
  "requests_per_second": -1,
  "throttled_until_millis": 0,
  "failures": []
}
```

#### 21.查询表数据行数 count/_count

```c#
//查询行数
GET /student/_doc/_count
//返回结果
{
  "count" : 1,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  }
}
```

#### 22.查询表结构 columns/ get

```c#
//查询表结构
GET /student
//返回结果
{
  "student" : {
    "aliases" : { },
    "mappings" : {
      "properties" : {
        "address" : {
          "type" : "text"
        },
        "age" : {
          "type" : "integer"
        },
        "name" : {
          "type" : "keyword",
          "normalizer" : "lowercase_normalizer"
        }
      }
    },
    "settings" : {
      "index" : {
        "routing" : {
          "allocation" : {
            "include" : {
              "_tier_preference" : "data_content"
            }
          }
        },
        "number_of_shards" : "1",
        "provided_name" : "student",
        "creation_date" : "1676428316854",
        "analysis" : {
          "normalizer" : {
            "lowercase_normalizer" : {
              "filter" : [
                "lowercase",
                "asciifolding"
              ],
              "type" : "custom",
              "char_filter" : [ ]
            }
          }
        },
        "number_of_replicas" : "1",
        "uuid" : "-LuVoBeNQzWsxFfhcw7tfA",
        "version" : {
          "created" : "7130199"
        }
      }
    }
  }
}

```

#### 23.查询属性经过分词后的结果 分词/_analyze

**es的分词器analyzer**：[千变万化的分词器](https://www.cnblogs.com/xiaobaozi-95/p/9328948.html)

```c#
//分词器
GET _analyze
{
  "analyzer": "standard",  //keyword不分词，完整句子，standard 默认分词，转换小写
  "text": "this is a test"
}
//还可以这样用,输出：I'm so happy!
POST _analyze
{
  "tokenizer": "keyword",
  "char_filter": ["html_strip"],
  "text": ["<p>I&apos;m so <b>happy</b>!</p>"]
}
//调用索引上面的自定义分词器
POST t_index/_analyze
{
  "analyzer": "lower_html_analyzer",
  "text":"字符"
}

{
  "tokens" : [
    {
      "token" : "this",
      "start_offset" : 0,
      "end_offset" : 4,
      "type" : "<ALPHANUM>",
      "position" : 0
    },
    {
      "token" : "is",
      "start_offset" : 5,
      "end_offset" : 7,
      "type" : "<ALPHANUM>",
      "position" : 1
    },
    {
      "token" : "a",
      "start_offset" : 8,
      "end_offset" : 9,
      "type" : "<ALPHANUM>",
      "position" : 2
    },
    {
      "token" : "test",
      "start_offset" : 10,
      "end_offset" : 14,
      "type" : "<ALPHANUM>",
      "position" : 3
    }
  ]
}
```

#### 24.更改最大的返回长度设置

```mysql
//设置所有index，最多可以返回两百万的结果
PUT /_all/_settings
{
  "index.max_result_window":2000000
}
//设置指定的index
PUT /ebay_listing/_settings
{
  "index.max_result_window":2000000
}
//创建表
PUT /platform_listing
{
	"settings": {
	  "index":{
      "max_result_window": 2000000
       }
	}
	 "mappings":{ 
    "properties":{   
      "id":    
      {
        "type":"keyword"
      }
    }
}      
```

[ElasticSearch 查询条数限制](https://learnku.com/articles/68372)

#### 25.match phrase替换wildcard

```mysql
#match phrase性能yyds
#短语查询：match_phrase，
#https://www.elastic.co/guide/cn/elasticsearch/guide/current/_phrase_search.html
GET ebay_listing/_search 
{
    "query" : {
        "match_phrase" : {
            "specifics" : "Bob"
        }
    }
}

经验：
1.match_phrase只能作用在有分词的属性上面，比如text属性，如果属性是keyword，默认是没有设置分词，并且也没有自定义分词，那么不能使用match_phrase
2.默认不区分大小写，也就是说我们可以直接用大写，或者小写的结果查询一遍就可以拿到所有匹配的结果，不需要设置大小写分词器进行转换
3.match phrase适合的是短语的查询，比如New York City，How Are You，并且会自动过滤“a”、“an”、“the”、“and”、“or”、“but”这些停用词，Wildcard更适合查询单独的单词进行使用，比如Stitched，damage，Dance这些，我们的需求侵权词基本都是单独的单词，所以使用wildcard进行查询
4.match_phrase查询的是单词短语的匹配，也就是说不需要在左右，前面，后面加空格的那种匹配，但是如果wildcard左右，前面，后面不加空格，会查询出十几万，几十万的重复数据，比如trippy侵权词，实际只有9000条数据，只like '%trippy%'查询出来了31万条记录，如果加上空格的逻辑只有10000条结果
5.match_phrase开12个线程，4核的cpu才到25%,wildcard占用的cpu到了100%，代表可以开更多的线程
```

#### 26.异步在后台更新属性为指定值

```c#
//在后台异步更新所有数据（四千万），并且返回一个任务编号
POST walmart_listing/_update_by_query?wait_for_completion=false
{
  "script": {
    "source": "ctx._source.createTime = '2023-08-05T11:25:20.1786388+08:00'",
    "lang": "painless"
  }
}
返回结果
{
  "task": "BxBy4ICRTnWUnojjjVXG3w:188454569"
}

//查询任务编号
GET _tasks/BxBy4ICRTnWUnojjjVXG3w:188454569

//验证结果:查询总行数
 GET /walmart_listing/_count   //  "total": 40417268,
// 查询所有行:  
GET /walmart_listing/_count
{
  "query": {
    "range" : {
      "createTime": {
        "gte": "2023-08-05T11:25:20.1786388+08:00",
        "lte": "2023-08-05T11:25:20.1786388+08:00"
      }
    }
  }
}

```

#### 27.异步在后台删除指定索引数据

```c#
//异步删除数据，不会超时
POST /walmart_listing/_delete_by_query?wait_for_completion=false
{
  "query": {
    "match_all": {}
  }
}

```



### ES的经验

**查看es执行日志**:

![1679454442298](Vue进步.assets/1679454442298.png)

**打印日志**：

                       settings.DisableDirectStreaming(true);  //调试开关，正式环境关闭
    
           var requestStr = System.Text.Encoding.Default.GetString(response.ApiCall.RequestBodyInBytes);
            var responseStr = System.Text.Encoding.Default.GetString(response.ApiCall.ResponseBodyInBytes);
            Console.WriteLine($"请求语句：{requestStr}");
            Console.WriteLine($"返回结果：{responseStr}");
**keyword/text: **keyword :32766最大字符长度，text:text类型是没有长度限制

**es的分词器analyzer**：[千变万化的分词器](https://www.cnblogs.com/xiaobaozi-95/p/9328948.html)

**text类型为什么要小写查询**:Text类型默认会进行分词，es默认是用的analyzer:`standard`分词器，它在分词后会将单词转成小写输出，所以使用`Oppo`大写查不出文档，使用小写`oppo`可以查出来

**为什么keyword模糊查询用大写**： 因为keyword类型，不会走分词器，所以不会转换成大小写，并且模糊查询的时候不会分词

**程序创建的index，字符默认是text**: 类库创建的index 字符类型默认是text类型，存在分词情况，要先用脚本创建，然后在进行新增数据

**elasticsearch-head使用**:google 扩展程序中搜索下载

![1679303703480](Vue进步.assets/1679303703480.png)

**通过服务启动kibana**:

```mysql
#windows下Kibana 可以从命令行启动，如下：

.\bin\kibana
#通过服务启动kibana
https://blog.csdn.net/weixin_44150634/article/details/118796246
```

**ES数据保存位置**:C:\elasticsearch\elasticsearch-8.6.1\data\indices,保存文件详情：

C:\elasticsearch\elasticsearch-8.6.1\data\indices\rQenuYAVTsaksei3vUqczg\0\index，

![1676253604013](Vue进步.assets/1676253604013.png)



**文件解释**

![img](Vue进步.assets/20134346_62d796122d6ee64201.webp)



**查询并发**：**几十到几百之间**，Elasticsearch的查询并发能力取决于多个因素，如硬件配置、索引大小、查询复杂度等。一般来说，单个节点的查询并发能力在几十到几百之间。要提高查询并发能力，可以通过扩展集群、优化索引结构和查询语句等方式来实现。

**并发问题**:

1.Pa...Foreach 中使用async await，控制不了执行顺序，并且有一个ForeachAsync方法可以使用

2.并发过程中用Thtread.sheep(100)，可以减少很多问题，CPU频繁切换可能会产生未知的问题

3.78个店铺一起并发，沃尔玛出现了异常提示：操作被取消了（The operation was canceled），使用了五个线程一起并发 （检测中）



4.Task.Factory.StartNew 加async await控制不到执行顺序，Task.Run可以控制执行顺序，有点坑

![1675830148133](Vue进步.assets/1675830148133.png)

## Consul

**Consul（服务注册发现，支持动态伸缩）:**支持负载均衡，支持动态伸缩，可以自定义访问策略，可以根据配置的权重访问，随机访问，轮训访问（按照1,2,3...）顺序访问一遍。

 		**权重访问** ：这个使用的比较频繁，可能服务器的配置是不一样的，有些配置好可以承担多的并发，那么就让它的权重高一些

​		**随机访问**：由程序使用Random生成随机数（范围是1至服务数）， 然后使用随机数的索引访问服务

​		**轮训访问** ：按照顺序依次循环访问，比如从1,2,3...，循环完毕后然后再从1,2,3...

**删除Consul单个实例**

每次重启服务（服务崩溃，更新pod节点的时候）都会往Consul注册新的实例，可能会影响性能

进入指定的服务列，复制实例ID

删除单个实例(后面是实例名称)：http://localhost:8500/v1/agent/service/deregister/2D69F26094044544B73112F128C69E9E

请求所有已失效节点（Vs关了就可以）：http://localhost:8500/v1/health/state/critical

请求所有节点（有效，无效都在）：http://localhost:8500/v1/agent/services

![1666151944469](Vue进步.assets/1666151944469.png)

删除指定的实例ID

![1666151908173](Vue进步.assets/1666151908173.png)

**Consul卡死**

卡死其实是鼠标点击了cmd界面，再次点击后，按一下Enter就可以

![1642142831787](Vue进步.assets/1642142831787.png)

**Consul的执行**

```
consul agent -dev
```

**执行Net**

```
dotnet run
```





## ABP

ABP全称：**Aspnet Boiler Plate** （ Asp Net项目下面通用的项目模板)

`卡在下载项目下来需要引用很多类库，难道要一个个自己Nuget中引用吗`



### ABP框架：控制台

1.在指定的文件夹新建文件夹名称

![1697622452775](Vue进步.assets/1697622452775.png)

2.执行：abp new Acme.MyConsoleApp -t console 创建项目

![1697622482788](Vue进步.assets/1697622482788.png)

3.双击打开项目

![1697622507717](Vue进步.assets/1697622507717.png)

4.项目列表

![1697622588248](Vue进步.assets/1697622588248.png)

5.启动展示

![1697623008376](Vue进步.assets/1697623008376.png)

### ABP框架：应用程序

> 下一步行动：配置数据库链接，ABP默认是使用Sql Server数据库的，我本地只有Mysql数据库

第一步先下载基础环境：dotnet tool install -g Volo.Abp.Cli --version=6.0

![1690968590630](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1690968590630.png)

然后新建一个vs Web Api项目，自动生成ABP脚手架，文件结构解释：https://docs.abp.io/zh-Hans/abp/latest/Startup-Templates/Application ,abp new Acme.BookStore

![1690968802082](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1690968802082.png)

然后再引用这个NetGet：Volo.Abp.Cli.Core 包 ，好像还要引用Volo.Abp.Identity等一系列引用到的包，这也太难了吧

![1690969331276](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1690969331276.png)

打算从Github下载一个案例看一看：https://github.com/abpframework/abp-samples/tree/master/TodoApp/Angular-EfCore，下载下来发现要Net7，安装Net7试一试

![1691055394826](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1691055394826.png)

https://dotnet.microsoft.com/zh-cn/download/dotnet/thank-you/sdk-7.0.306-windows-x64-installer 在这里安装

![1697525082052](Vue进步.assets/1697525082052.png)

项目创建成功，怎么执行，可以看教程：https://www.abp.io/project-created-success?ui=Mvc&db=EntityFrameworkCore&tiered=no

![1697525068866](Vue进步.assets/1697525068866.png)

发现教程里面还是Abp官网的内容

![1697525484147](Vue进步.assets/1697525484147.png)

我还是研究上面Github下载下来的代码看看，VS 2022+NET7直接跑

![1697525525819](Vue进步.assets/1697525525819.png)

**AbpTimer组件**

![1679642055761](../../../%E5%9D%9A%E6%9E%9C%E4%BA%91%E5%90%8C%E6%AD%A5/%E6%88%91%E7%9A%84%E5%9D%9A%E6%9E%9C%E4%BA%91/Rolan%E8%8F%9C%E5%8D%95/Vue%E8%BF%9B%E6%AD%A5.assets/1679642055761.png)

### ABP框架：安装模块

1.安装命令：abp new Acme.IssueManagement -t module

![1697624130008](Vue进步.assets/1697624130008.png)

2.出现了需要安装Yarn环境的提示，从官网：https://classic.yarnpkg.com/lang/en/进行安装，**Yarn是一个JavaScript的包管理工具，和NPM一样，不过它比NPM更快，安全性更高**

**安装Yarn**：您需要确保已经在您的系统上安装了Yarn，并且它已经被添加到了环境变量中。您可以访问[Yarn官网](https://classic.yarnpkg.com/lang/en/)了解如何安装。安装完成后，您可能需要重新启动命令行或计算机来更新环境变量。

3.安装环境：npm install --global yarn

![1697624320588](Vue进步.assets/1697624320588.png)

4.检测安装结果：yarn --version 

![1697624356958](Vue进步.assets/1697624356958.png)

5.重新执行：abp new Acme.IssueManagement -t module

![1697624686823](Vue进步.assets/1697624686823.png)

### abp中linq的应用

[abp中linq的应用](https://blog.csdn.net/qq_34452824/article/details/119379559?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.pc_relevant_default&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.pc_relevant_default&utm_relevant_index=2)

[执行原生sql的方式](https://blog.csdn.net/qq_34452824/article/details/119379559?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.pc_relevant_default&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.pc_relevant_default&utm_relevant_index=2)

```
            var query = from a in _productRepository.GetAll()
                        join b in _productPropertyRepository.GetAll() on a.ProductSpu equals b.ProductSpu
                        where b.IsDeleted == false
                        select new ProductPropertyDto
                        {
                            ProductSpu = b.ProductSpu,
                            ProductSku = b.ProductSku,
                            CategoryID = a.CategoryId,
                            CategoryNamePath = a.CategoryNamePath,
                        };
```



### ABP框架原来还有界面的

看到有Bootstrap 还是响应式的布局

![1642815736601](Vue进步.assets/1642815736601.png)

### ABP框架与ASP.NET CORE的关系

**两者都是模块化的：**ASP.NET Core引入了许多内置解决方案（扩展库），用于**依赖注入**，**日志记录**，**缓存**，**本地化**，**配置**等。它们实际上独立于ASP.NET Core，可用于任何类型的应用程序（**原来Asp.net Core这些组件化的应用都是可以拆分进行使用的**）

目前ABP是模块化的而且包含了很多粒度很小的nuget包，Abp里面由许多功能，比如DDD类，审计，授权，后台作业，事件总线，json序列化，本地化，多租户，线程，时间等......这些功能都是拆分到各自独立的Nuget包中，我们可以独立使用，选择性的使用他们

### ABP约定服务注册

ABP框架是通过命名约定的方式由ABP框架自动便利程序集进行依赖注入的（可以重写`PreConfigureServices`方法,设置`SkipAutoServiceRegistration`为`true`跳过自动注册，那么就需要在statup.cs中像Autofac样进行手动注册）

[ABP源码分析 - 约定注册](https://blog.csdn.net/qq_34452824/article/details/119379559?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.pc_relevant_default&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.pc_relevant_default&utm_relevant_index=2)

[ABP官网对于依赖注入的解析](https://blog.csdn.net/qq_34452824/article/details/119379559?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.pc_relevant_default&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7Edefault-1.pc_relevant_default&utm_relevant_index=2)

**ABP—对象之间的双向映射** :

```c#
 //两个类对象之间的双向映射
 [AutoMap(typeof(Person))]
```

https://blog.csdn.net/wangwengrui40/article/details/86609037  --ABP—对象之间的映射



### **ABP框架结构**

![img](https://upload-images.jianshu.io/upload_images/2799767-f113cae025f4e844.png?imageMogr2/auto-orient/strip|imageView2/2/format/webp)

https://www.jb51.net/article/86688.htm --ABP框架的体系结构及模块系统讲解

https://www.jianshu.com/p/a6e9ace79345  --ABP入门系列目录——学习Abp框架之实操演练

### DDD领域驱动设计1

DDD是一种程序分析设计方法，不关乎具体技术，具体代码实现依旧是OOP和AOP

**应用场景：**应对复杂软件开发，不熟悉，不明白的系统。

解决沟通问题，降低沟通成本。比如建立一个小区很简单，如果是建立一个区，一个城市呢，会很复杂

技术人员与需求人员由于各自缺乏彼此的专业知识，比较难站在对方的角度进行思考，现实中难以沟通。

双方的统一语言在哪里？

- **Domain（领域/业务）：**领域可以是一个单独的服务，也可以是一个项目，可以是一个模块，可以是一个BLL，甚至是一个问题（需求），领域就是用来分析需求的（需求就是业务），把一个大盒子，不断的划分成小盒子，小盒子里面不断的划分
- **Driven（驱动/实现）**：对外提供哪些功能，哪些人可以获取数据，权限验证，代码实现
- **Design（设计）**：把按照领域实现的模块在这里进行组装，组合成最终的产品

领域驱动设计以需求为主导，确定下来业务需求后再进行代码开发，避免开发过后的改动。可以**分为三步走**：第一先理解领域，再拆分领域，最后再细化领域，领域与领域之间的拆分是最复杂的，**因为有些需求之间存在交集，怎么划分到正确的领域**？**有些模块是单独划分成一个领域还是合并到当前的一个领域当中**（就是这个模块业务权重是怎么样的，比如秒杀模块是放到交易的场景后端服务中，还是单独部署一个服务），这个需要对业务有足够的理解。

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



### DDD领域驱动设计2

DDD领域驱动设计：大盒子里面套小盒子



DDD领域驱动设计是一种业务分析设计方法 ，把整个系统从业务层面进行从大到小的划分模块，由模块再划分成表，由表再清晰表里面具体的属性，这个时候考虑的不是一个类，考虑的是一个一个小盒子，比如电商平台：

**第一步：** 划分为仓储 物流 订单 交易 用户这些模块

**第二步：** 划分仓储里面有不同地点的仓库，仓库管理人员，仓库保存的货物等信息

**第三步：**进一步划分表里面的内容，或者单个模块技术业务上面的实现。



微服务架构落地的第一件事就是DDD，微服务落地时需要将业务进行分拆，把项目分解成多个低耦合的模块，而DDD刚好就可以进行拆分。



**领域：**大盒子，聚合根（可以包含多个对象）



## Typora

放大缩小内容（图片）：shift+ctrl ＋＋加号 (放大)，shift+ctrl ＋-减号 （缩小）

## Jira：自动化管理

**jira卖点就是自动化 ，每一个任务 什么时候提交的 什么时候发布的 发布到哪些生产环境都有记录** 

![1668651182295](Vue进步.assets/1668651182295.png)

**如何准确估计组员的工作量复杂度**：这个挺复杂的，我简单说下。你有一个**组复杂度取值**，通常是一个斐波那契数列，比如2 3 5 8 13 21。然后在你所有任务里面选一个**最简单**的(要考虑任务确定性，需要做哪些事等等)，把它当作复杂度为2的任务，**其他任务基于这个任务给它一个相对的复杂度取值**。一般来说，跑完几个任务之后，一个团队每个任务能解决的复杂度综合是在一个范围里面的，那么下次再给团员分任务的时候就可以参考这个值来预测本次sprint可以做到哪些了。
如果一个任务复杂度超过13，就表示这个任务太大了，还需要斩小一点。
这些事，最好队伍里面有个scrum master（敏捷管理助手）配合你做，容易很多。没有的话就只有自己啃了

jira里面有个插件叫**magic estimations可以辅助我们做复杂度预测的**

[jira官网](https://www.atlassian.com/zh/software/jira)

## 面试技术

今天逛了下前程无忧，发现一下几点，比如Linux与容器化部署要会，还有数据库优化，索引，b+树,orm框架需要会

## 人工智能

自动机器学习（AutoML）：让机器学会学习这件事情，元学习

“他们将拳击胜负分数和一场场的数据输入程序，然后不到12秒的时间里，机器就给出了四十年里所有重要的比赛历程，然后迅速地对应给出了一个比赛短评，辅以最优配图和最佳标题”。

要用人工只能生成什么？

聊天服务：

*对话系统目前有两种主要类型：面向目标（Siri，Alexa，Cortana等）和一般会话（Microsoft Tay bot）。*

在面对某一个领域的机器对话系统已经有应用的地方了，但是面向生活使用的一般会话，目前还没有应用

`可以尝试找一个懂机器学习的人指导一下，或者`

