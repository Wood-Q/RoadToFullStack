#!/bin/bash

name="woodq"
echo $name
url="jmcomic.com"
# 只读变量
#randonly url
# 删除变量
unset url

# 用declare声明变量，指定类型更好看懂
declare -i num=42
num+=1
echo $num

#声明数组
arr=(1 2 3 4 5)
echo ${arr[0]}
echo ${arr[*]}

declare -a dict
dict["name"]="woodqqqq"
echo ${dict["name"]}

# 获取环境变量
echo $PATH

# 字符串拼接
var1="hello"
sayhi='hello,'$var1''
echo $sayhi

# 获取字符串长度
echo ${#sayhi}

# 截取字符串
echo ${sayhi:1:4}