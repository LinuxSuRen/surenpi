---
title: Shell Example
description: Shell Example
toc: true
---

## 输出文件制定行

Solution 1
```
num=0 head -n 10 file.txt | while read line;
do
    if [ "$num" == "9" ]; then
        echo $line;
        break; 
    fi;
    num=$((num+1)); 
done
```

Solution 2
```
awk 'NR==10' file.txt
```
