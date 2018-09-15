# spamcheck

检测聊天信息中潜在的`发广告`、`卖星币`等言论，后续也可以对违禁词以及暴恐词的检测进行支持

---

## DEPENDENCY

所需第三方依赖，可以通过 `go get -u `的方式进行安装


## EXPLAINATION

- `config.json` 服务运行配置文件

  - `SOCK_FILE`： 服务以**Unix Domain Socket** 方式运行的时候sock文件的存放位置
  - `WORKSPACE`： 代码根目录，spamcheck.go文件的位置
  - `CLASSES`： 分类类别，至少**2**个。每个分类需要有对应的先验概率词典文件，示例：分类为**adwords**，则*同级目录*下需要有**adwords.txt**，字典越大，服务的分类结果越准确。
  - `DICTIONARYFILE`：`sego`库分类字典，需要手动指定，文件路径为绝对路径；分词准确度可以通过自定义字典进行替换。
  - `SOCKET_BUFFER_SIZE`：服务以**Unix Domain Socket**的方式运行的时候，缓冲区大小，可以根据**数据量**的大小进行调试。

- `spamcheck.go`： 垃圾词检测主文件

- `README.md`：服务介绍文件

- `requirements.txt`: 项目依赖第三方库列表

- `*.txt`: 根据**config.json** 中**CLASSES** 指定的分类的先验概率字典。格式：一行一个，每行一个词汇

## USAGE

- **`Run as 'Unix Domain Socket'`**
```
go run spamcheck.go -c ./config.json -a sock
```

- **`Run as RPC`**
```
TODO
```
