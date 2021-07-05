## 文件模块初始化

+ 先判断是否有导包`import`，若有按顺序依次进入到该文件，并循环判断，直到无导包为止
+ 先初始化常量（或枚举）`const`
+ 初始化全局变量`var`
+ 执行初始化函数`init()`
+ 返回上一个入口，如此往返，直到最初的入口
+ 执行主函数`main()`

![img.png](docs/img/img.png)

## 指针

### 值传递

+ 先对形参开辟内存 并赋默认值
+ 将实参的值赋值给形参 （仅改变行参的值）
+ 在调用函数内对形参进行操作
  ![img.png](docs/img/img2.png)

### 地址传递

+ 形式参数是某一个类型的地址
+ 实际参数为地址 使用`&`可以获取变量地址
+ 形式参数储存的是实际参数的地址
+ 在调用函数内实际上是对地址指向的内存进行操作(使用`*`可以获取地址对应内存的变量)

**指针可以嵌套，可以获取指针的指针，多个`*`也可以分别获取低级指针的地址或实际值**
![img.png](docs/img/img3.png)