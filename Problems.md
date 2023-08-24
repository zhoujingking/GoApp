## 1. 使用go get出现“连接超时”情况时，添加goproxy代理。

      go env -w GOPROXY=https://goproxy.cn
      go get ...    # 重新执行之前的命令