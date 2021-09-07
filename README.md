## 一个go语言实现的博客 
  fork 自 https://github.com/zxysilent/blog
在其上做了二次开发。
增加了
1. markdown mathjax支持。
2. 对接了微信公众号。

修复了一些小bug。


预览地址： http://blog.go2live.cn/
# 补充说明
下载下来后。
`swag init` 生成swager 文档。直接运行会报错。

复制 config.dev.toml 为 config.toml 改成你自己的配置信息。 

密码直接登录后台改。账户得改数据库的num字段。

博客数据等，直接truncate table 清空吧。

默认的统计代码 是第三方的统计页面访问量的。
可以再加下百度的统计代码。
关于页 和 友链是 特殊的页面。
链接得用about和link。

录入文章 需要添加 <!--more-->
这个之前的是文章简介。
