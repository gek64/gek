# VIVYCORE
- `version 0.1.0`
- 是 `Project vivy` 项目群中使用的核心依赖包
  
## 编译说明
- 为了不依靠网络就能下载及编译,包名称改为`localhost/vivycore`,其他vivy项目均在`go.mod`中使用 `replace localhost/vivycore => ../vivycore`来获取core包内容,最大程度上避免因网络造成依赖问题,同时也不受制于某一家企业机构提供的网络服务
- 编译时需将`vivecore`与`project vivy`项目放在平级目录便可成功获取到本包内容