## 解析和求值的语言——Monkey

- 类C语法
- 变量绑定
- 类型：整型和布尔型
- 算术表达式
- 内置函数
- 头等函数和高阶函数
- 闭包
- 字符串数据结构
- 数组数据结构
- 哈希表数据结构

```
    // 绑定变量名称和值
    let age = 1;
    let name = "Monkey";
    let isFemale = false;
    let result = 10 * (20 / 2);
    let myArray = [1, 2, 3, 4, 5];
    let myHashtable = {"name": "Monkey", "age": 1};
    
    // 访问数组和哈希表元素
    myArray[2]; // 3
    myHashtable["name"]; // "Monkey"
    
    // 绑定变量和函数
    let add = fn(a, b) {return a + b;};
    // let add = fn(a, b) {a + b;}; // 隐式返回值
    // 调用函数
    fn(1, 2);
```