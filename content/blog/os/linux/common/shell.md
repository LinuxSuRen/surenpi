---
title: Shell
description: Linux Shell 的基本介绍
toc: true
---

## 执行最近的命令
有几种方法可以快速地执行最近执行过的命令：
* `!!` 或者 `!-1` 执行上一个命令
* `!u` 执行上一个以 `u` 开头的命令
* `Ctrl+R` 搜索命令
* `Ctrl+P Ctrl+O`

## 下载并执行

`curl http://test.com/test.sh | bash -s arg1 arg2`

## 文件路径

从文件路径中提取文件名和目录

```
path=/var/www/1.log
echo $(basename $path)
echo $(dirname $path)
```

## 文件是否存在

```
if [ ! -f "env.sh" ]; then
    echo "we need the env.sh to setup vars"
    exit -1
else
    echo "file exists"
fi
```

```
if [ ! -d "env.sh" ]; then
    echo "we need the env.sh to setup vars"
    exit -1
else
    echo "file exists"
fi
```

## 脚本相关

`${BASH_SOURCE[0]}` 脚本所在的相对路径

`$(dirname "${BASH_SOURCE[0]}")` 脚本所在的相对目录

## 参数

Shell 脚本可以像其他的命令行一样接收参数

|变量名称|解释|
|----|----|
|`$1`|第一个参数|
|`$2`|第二个参数|
|`$@`|完整的参数列表|
|`$#`|参数个数|

## 逻辑判断

```
if [[ "1" == "1" && "2" == "2" ]]; then
        echo "equal"
fi
```

## 循环

获取当前目录下的文件并循环打印出文件名：
```
for f in `ls`;
do
        echo $f
done
```

## 函数

我们可以在 `shell` 脚本中用下面的方式来定义函数：

```
function_name(){
    echo "hello"
}
function_name
```
