## Vue基础

### 一、创建项目

```
vue create my-app
```



### 二、模板语法

#### 1.插值语法 {{}}

**会将数据解释为普通文本，而非 HTML 代码**

```
<div id="app">{{name}}</div>
```

 **使用javascript表达式**

```
 <div id="app">{{'个人介绍：'+name}}</div>
```

```
 <div id="app">{{isShow ? name : ''}}</div>
```

```
 <div id="app">{{ name.split('').join('-') }}</div>
```



#### 2.v-text

使用指令的方式渲染普通文本

```
<div v-text="name" id="app"></div>
```



#### 3.v-html

会将内容转换成html

```
<div id="app" v-html="name"></div> 
```



#### 3.表单元素数据绑定

##### `v-model` 数据双向绑定

```
<input v-model="name">
```

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



#### 4.属性绑定

**v-bind**可以给元素绑定动态属性

```html
<!--绑定字符串-->
<div v-bind:id="'box'"></div>
<!--绑定变量-->
<div v-bind:id="idName"></div>
<!--js表达式-->
<div v-bind:id="idName+'1'"></div>
```

也可以通过` :   `简写完成属性绑定

```html
<!-- 完整语法 -->
<a v-bind:href="url">...</a>

<!-- 缩写 -->
<a :href="url">...</a>
```



#### 5.事件绑定

通过`v-on`给元素绑定事件

用在普通元素上时，只能监听原生DOM事件。用在自定义元素组件上时，也可以监听子组件触发的**自定义事件**

```HTML
<!-- 点击事件 -->
<button v-on:click="getName"></button>
```

也可以通过`@`简写事件绑定

```HTML
<!-- 点击事件 -->
<button @click="getName"></button>
```

**事件修饰符**

- `.stop` - 调用 `event.stopPropagation()`阻止冒泡事件传播。

- `.prevent` - 调用 `event.preventDefault()`组织默认事件执行。

- `.capture` - 添加事件侦听器时使用 capture 模式（影响冒泡执行顺序）。

    `当元素发生冒泡时，先触发带有capture 修饰符的元素。若有多个该修饰符，则由外而内触发。
    就是谁有该事件修饰符，就先触发谁`

- `.self` - 只当事件是从侦听器绑定的元素本身触发时才触发回调。

- `.{keyCode | keyAlias}` - 只当事件是从特定键触发时才触发回调。

- `.native` - 监听组件根元素的原生事件（用于给某个自定义组件绑定事件）。

- `.once` - 只会触发一次回调。

- `.left` - (2.2.0) 只当点击鼠标左键时触发。

- `.right` - (2.2.0) 只当点击鼠标右键时触发。

- `.middle` - (2.2.0) 只当点击鼠标中键时触发。



#### 6.条件渲染

#####  v-show

根据表达式之真假值，切换元素的 `display`属性

```html
 <div v-show="true" id="box"></div>
```

##### v-if & v-else-if & v-else

根据表达式的值来有条件地渲染元素。在切换时元素被销毁并重建

```html
 <div v-if="xxx">内容1</div> 
 <div v-else-if="xxx">内容2</div> 
 <div v-else >内容3</div>
```



#### 7.列表渲染（循环）

基于数据多次渲染某一段元素，并且为为当前遍历的元素提供别名,循环渲染数据时需要为元素添加`key`作为当前元素的唯一标识

另外也可以为数组索引指定别名 (或者用于对象的键)

```html
<div v-for="item in list" :key="i.id">
  {{ item.text }}
</div>

<div v-for="(item, index) in list"></div>
```



### 三、组件

#### 1.创建组件

**全局组件**  （全局注册的组件可以直接在各个页面直接调用）

```js
Vue.component('button-com', {
  data: function () {
    return {
      count: 0
    }
  },
  template: '<button v-on:click="count++">点击{{ count }} 次</button>'
})

//组件调用
<button-com></button-com>
```

**局部组件**  （局部注册的组件需要在调用的页面进行配置才可以使用）

```js
//组件
var Child = {
  template: '<h1>自定义组件!</h1>'
}

//组件调用
new Vue({
  el: '#app',
  components: {
    // <runoob> 将只在父模板可用
    'h1-com': Child
  }
})

<h1-com></h1-com>

```



####  2.数据data

data是一个函数（确保数据独立唯一），从该函数内返回一个对象作为当前Vue 实例的数据对象。Vue 会把 data 的 property 转换为 getter/setter，从而让 data 的 property 能够响应数据变化。data内返回的属性都可以在当前组件内通过`this`访问

```js
Vue.component('button-com', {
  data: function () {
    return {
      count: 0
    }
  },
  template: '<button v-on:click="count++">点击{{ count }} 次</button>'
})
```



#### 3.props

props 可以是数组或对象，用于接收来自父组件的数据。props 可以是简单的数组，或者使用对象作为替代，对象允许配置高级选项，如类型检测、自定义验证和设置默认值

```js
// 数组语法
Vue.component('demo-com', {
  props: ['size', 'myMessage']
})

// 对象语法，提供验证
Vue.component('demo-com', {
  props: {
    // 检测类型
    height: Number,
    // 检测类型 + 其他验证
    age: {
      type: Number,
      default: 0,
      required: true,
      validator: function (value) {
        return value >= 0
      }
    }
  }
})
```

基于对象的语法使用选项：

- `type`：规定当前属性接受的数据类型：`String`、`Number`、`Boolean`、`Array`、`Object`、`Date`、`Function`
- `default`：为当前属性指定一个默认值。如果该 prop 没有被传入，则换做用这个值。对象或数组的默认值必须从一个函数返回
- `required`：`Boolean`
    定义该 prop 是否是必填项
- `validator`：`Function`
    自定义验证函数验证属性值

所有的 prop 都使得其父子 prop 之间形成了一个**单向下行绑定**：父级 prop 的更新会向下流动到子组件中，但是反过来则不行



#### 4.计算属性

 当数据需要进行一些处理才可以正确显示的时候，就可以使用`computed`属性计算，计算属性的结果会被缓存，当计算的数据发生变化时才会重新计算

```js
 data: {
    message: 'Hello'
  },
  computed: {
    // 计算属性的 getter
    reversedMessage: function () {
      // `this` 指向 vm 实例
      return this.message.split('').reverse().join('')
    }
  }
```

**使用场景：**

- 当页面中有某些数据依赖其他数据进行变动

- 每个计算属性默认是踏他get方法的实现，但是添加set方法

- 计算属性可以一个方法监听多个属性变化

    

#### 5.数据监听

当需要随时监听一个属性发生的所有变化的时候，可以使用`watch`监听

```js
 data: {
    firstName: 'Foo',
    lastName: 'Bar',
    fullName: 'Foo Bar'
  },
  watch: {
    firstName: function (val) {
      this.fullName = val + ' ' + this.lastName
    },
    lastName: function (val) {
      this.fullName = this.firstName + ' ' + val
    }
  }
```

**可选参数：**

- handler：数据变化的回调
- immediate：规定是否进入页面后就会执行一次监听 （默认false）
- deep: 是否开启深度监听，即所有属性都加上监听器，如果其中一个发生改变了就执行handler函数

**使用场景：**

- 数据变化时执行异步或开销较大的操作
- 数据监听是以每个单独数据为基础进行的监听



#### 6.自定义事件

可以通过自定义事件让父子组件之间完成数据传递

```HTML
<!--事件绑定-->
<my-component @getName="getNameFun"></my-component>

<!--事件触发-->
this.$emit('getName')
```



#### 7. .sync修饰符

在有些情况下，我们可能需要对一个 prop 进行“双向绑定”，可以通过` .sync`修饰符修饰该属性为组建的双向绑定属性

```html
<!--父组件-->
<demo-com :title.sync="title"></demo-com>

<!--子组件-->
this.$emit('update:title', newTitle)

```



#### 8.插槽slot

在vue组件里，把部分代码结构写在组件双标签内，由vue归纳到耨一个标签内，传递到子组件，供给子组件放置到任意地方显示，这种使用方式叫**插槽**

 1.父组件在引用子组件时希望向子组价传递模板内容`<p>测试一下吧内容写在这里了能否显示</p>` 

 2.子组件让父组件传过来的模板内容在所在的位置显示 

 3.子组件中的`<slot>`就是一个槽，可以接收父组件传过来的模板内容，`<slot>` 元素自身将被替换 

##### 默认插槽

直接写在组件标签内的插槽

```html
<my-component>
	<span>我不知道自己会放在哪里</span>
</my-component>

<my-component>
      <template v-slot:default>
		不具名查抄也可以通过  v-slot:default 定义      
  	  </template>
</my-component>


<!--my-component组件-->
<div>
    <slot></slot>
</div>

<!--渲染结果-->
<div>
    <span>我不知道自己会放在哪里</span>
</div>
```

##### 具名插槽

带有自己名字的插槽

```html
<my-component>
     <template v-slot:header>
   		 <h1>页头</h1>
  	 </template>
      <main>
        内容
      </main>
    
      <template v-slot:footer>
   		 <h1>页头</h1>
  	 </template>
</my-component>
<!--具名插槽缩写-->
<my-component>
     <template #:header>
   		 <h1>页头</h1>
  	 </template>
      <main>
        内容
      </main>
    
      <template #:footer>
   		 <h1>页头</h1>
  	 </template>
</my-component>

<!--my-component组件-->
<div>
    <slot name="header"></slot>
    <slot name="main"></slot>
    <slot name="footer"></slot>
</div>

```

##### 作用域插槽

规定当前插槽可以使用子组件的部分数据

```HTML
<my-component>
    <template v-slot:default="slotProps">
        {{ slotProps.user.name }}
    </template>
</my-component>

<!--my-component组件-->
<div>
    <slot v-bind:userData="user">
        {{ user.name }}
      </slot>
</div>
```

##### 解构插槽

通过子组件传递的数据最终会被**slotProps**形参接收，所以我们可以通过解构语法直接拿到需要渲染的数据结构

```HTML
<template v-slot:default="{user}">
        {{user.name }}
</template>
```

也可以给解构的数据设定默认值

```html
<template v-slot:default="{user:{name:'tina'}}">
        {{user.name }}
</template>
```



#### 9.动态组件

vue提供了`is`语法扩展了在HTML标签上渲染组件

```html
//在component内渲染my-component组件 通过修改变量值为其他组件使得当前component渲染其他组件
<component :is="变量"></component>
<component :is="'my-component'"></component>
```

```html
<div :is="变量"></div>
<div :is="'my-component'"></div>
```



#### 10.组件缓存

在动态组件上使用 `keep-alive`可以实现组件切换的同时组件数据被缓存，避免重复渲染

`<keep-alive>` 是一个抽象组件：它自身不会渲染一个 DOM 元素

```html
<keep-alive>
  <component :is="变量"></component>
</keep-alive>
```

**Props**：

- `include` - 字符串或正则表达式。只有名称匹配的组件会被缓存。
- `exclude` - 字符串或正则表达式。任何名称匹配的组件都不会被缓存。
- `max` - 数字。最多可以缓存多少组件实例。



#### 11.异步组件

异步组件就是只有在这个组件需要被渲染的时候才会触发渲染，并且会把结果缓存起来供下一次渲染

```js
  components: {
    'my-component': () => import('./my-component')
  }
```



#### 12.【组件通讯】[]()

- **父组件访问子组件**

    - 通过props向子组件传递数据

    - 使用$children或$refs访问子组件

          通过$children访问子组件时，是一个数组类型，访问其中的子组件必须通过索引值，
          
          想明确获取其中一个特定的组件，这个时候就可以使用$refs

        ```html
        <!--给子组件传递props-->
        <demo-com :title.sync="title"></demo-com>
        
        <!--定义ref-->
        <demo-com ref="demoRef" :title.sync="title"></demo-com>
        console.log(this.$refs.demoRef) //获取到demo-com实例
        
        console.log(this.$children) //获取当前组件内的所有子组件
        ```

        

- **子组件访问父组件**

    - 通过事件向父组件发送消息

    - 使用$parent访问父组件

        ```HTML
        <!--给子组件绑定一个事件-->
        <demo-com @demoClick="getName"></demo-com>
        
        <!--子组件内-->
        this.$emit('demoClick', 参数)
         
        console.log(this.$parent) //获取父组件实例
        ```



####   13.依赖注入

​		**provide**：指定提供给后代组件的数据/方法

​		**inject**：在任何后代组件接受在祖先组件指定的数据/方法

```js
//祖先组件 
provide:function(){
    return {
      name:this.name,
      setName:this.setName
    }
 },
//子组件
inject:['name','setName']
```



#### 14.循环引用

- **递归组件**

      组件是可以在它们自己的模板中通过组件的 `name`调用自身的

    ```js
     name: 'HelloWorld',
      props: {
        msg: String
      },
      data(){
        return {
          name:'<span>tina</span>'
        }
      },
    ```

    ```html
    <template>
    	<div>
            <HelloWorld msg="Welcome to Your Vue.js App"/>
        </div>
    </template>
    ```

    **如果自调用组件没有设定渲染条件，就会导致无限循环**

    

- **组件循环引用**

    如果有两个组件相互引用，会出现A组件调用B组而B组件还未解析的问题

    可以通过2种方式解决：

    - **在生命周期钩子 `beforeCreate` 内注册子组件**

        ```js
        beforeCreate: function () {
          this.$options.components.domeCom = require('./domeC-com.vue').default
        }
        ```

    - **异步 `import`子组件**

        ```js
        components: {
          domeCom: () => import('./domeC-com.vue.vue')
        }
        ```

        **由于存在互相调用的关系，如果不设定渲染条件也会导致无限循环**



### 四、动画

#### 	1.过渡

​     包裹在`transition`内的组件在插入、更新或者移除时，都可以使用过渡效果

```HTML
  <transition name="fade">
    <p v-if="show">hello</p>
  </transition>

<style>
.fade-enter-active, .fade-leave-active {
  transition: opacity .5s;
}
.fade-enter, .fade-leave-to {
  opacity: 0;
}
</style>
```

当插入或删除包含在 `transition` 组件中的元素时，Vue 将会做以下处理：

1. 自动匹配 CSS 过渡或动画，在恰当的时机添加/删除 CSS 类名。
2. 过渡组件配置的钩子函数也会在相对的时机被调用。



####   2、过渡类名

​	在进入/离开的过渡中，会有 6 个 class 切换。

​        **显示：**

1. `v-enter`：起始状态

2. `v-enter-active`：过渡状态

3. `v-enter-to`：过渡结果 **2.1.8 版及以上**

    **隐藏：**

4. `v-leave`：起始状态

5. `v-leave-active`：过渡状态

6. `v-leave-to`：过渡结果 **2.1.8 版及以上**

![Transition Diagram](https://cn.vuejs.org/images/transition.png)

对于这些在过渡中切换的类名来说，如果你使用一个没有名字的 `<transition>`，则 `v-` 是这些类名的默认前缀。如果你使用了 `<transition name="my-transition">`，那么 `v-enter` 会替换为 `my-transition-enter`。



#### 3.动画

也可以使用css动画实现过渡效果

```HTML
 <transition name="bounce">
    <p v-if="show">内容内容</p>
  </transition>

<style>
	.bounce-enter-active {
  animation: bounce-in .5s;
}
.bounce-leave-active {
  animation: bounce-in .5s reverse;
}
@keyframes bounce-in {
  0% {
    transform: scale(0);
  }
  50% {
    transform: scale(1.5);
  }
  100% {
    transform: scale(1);
  }
}
</style>
```



#### 4.自定义过渡类名

适用于使用第三方动画库，直接指定每个过渡时机的类名

```HTML
 <transition
    name="custom-classes-transition"
    enter-active-class="animated tada"
    leave-active-class="animated bounceOutRight"
  >
    <p v-if="show">hello</p>
  </transition>
```



#### 5.过渡周期钩子

可以使用对应的钩子函数监听过渡

@before-enter=“beforeEnter“ ：进入过渡运行前

@enter=“enter“ ：进入过渡运行时

@after-enter=“afterEnter“ ：进入过渡运行后

@enter-cancelled=“enterCancelled“ ：进入过渡被打断时

@before-leave=“beforeLeave“ ：离开过渡运行前

@leave=“leave“ ：离开过渡运行时

@after-leave=“afterLeave“ ：离开过渡运行后

@leave-cancelled=“leaveCancelled“ ：离开过渡被打断时

```HTML
<transition
  v-on:before-enter="beforeEnterFun"
  v-on:enter="enterFun"
  v-on:after-enter="afterEnterFun"
  v-on:enter-cancelled="enterCancelledFun"

  v-on:before-leave="beforeLeaveFun"
  v-on:leave="leaveFun"
  v-on:after-leave="afterLeaveFun"
  v-on:leave-cancelled="leaveCancelledFun"
>
  <!-- ... -->
</transition>


```



### 五、混入（mixin）

