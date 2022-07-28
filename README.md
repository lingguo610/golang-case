<!--
 * @Author: error: git config user.name && git config user.email & please set dead value or install git
 * @Date: 2022-07-28 17:01:00
 * @LastEditors: error: git config user.name && git config user.email & please set dead value or install git
 * @LastEditTime: 2022-07-28 17:50:41
 * @FilePath: \goprojecte:\d\README.md
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
-->
# golang-case

## 1、case_plugin 插件的例子
### 描述：
        在主程序 的main.main函数中加载一个插件， 插件的main包和common 包都定义了init函数
### 步骤：
        生成插件 ， 进入 case_plugin\goplug， 执行 go build -buildmode=plugin -o plugin1.so plugin1.go ， 把 plugin1.so 放入到 case_plugin\useplug下
        使用插件， 进入case_plugin\useplug， 执行 go build, 执行 main，打印结果如下：

### 执行结果：
        2022/07/28 17:17:21 main package init function called     #主程序的init函数
        2022/07/28 17:17:21 main function stared                  #主程序的main函数
        2022/07/28 17:17:21 plugin common init                    #插件的common包的init函数
        2022/07/28 17:17:21 plugin1 init                          #插件的main包的init函数
        plugin common f2                                          #插件的F2函数， 因为在 插件main包的init函数里面调用了F2
        2022/07/28 17:17:21 plugin opened                         #主程序打开插件后的打印

