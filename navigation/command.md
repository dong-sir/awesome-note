
---
title: ls
tags: Linux,排序,文件
url: https://mp.weixin.qq.com/s/g7PzLMUMJua4qNHGlJocDA
desc: ls 命令按大小对文件进行排序的各种方法,根据文件大小对文件进行排序(升序)
---

---
title: Linux 中查找大文件
tags: Linux,查找,find,du
url: https://mp.weixin.qq.com/s/NSoKzAOZr4R7X-rY6R7-Ag
desc: ls 命令按大小对文件进行排序的各种方法,根据文件大小对文件进行排序(升序)
---

在当前工作目录中要搜索大小超过100MB的文件
```
sudo find . -xdev -type f -size +100M
```
. 代表当前目录。如要搜索其它目录替换.为要搜索目录的路径


使用du命令查找大文件和目录
以下命令将打印最大的文件和目录
```
du -ahx . | sort -rh | head -5
```