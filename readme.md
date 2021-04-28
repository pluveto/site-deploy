# site-deploy

一个用于打包发布静态网站的工具。你只需要在服务器上运行本程序，然后在本地 build 之后通过脚本压缩上传，服务器上的本程序会自动解压并覆盖到所要部署的静态位置。

**警告**：由于只是刚刚抽空完成开发，本程序仅用于测试环境。只使用了固定口令进行权限校验，并且错误提示信息中可能会暴露您的路径。

## 使用方法

### 服务端

下载服务器对应架构的可执行文件，重命名为 `site-deploy`（Windows：`site-deploy.exe`）。将其放到一个目录如 `~/deploy`。

创建下列目录：

```
conf/
log/
tmp/
```

在 `conf/` 下创建 `app.ini`：

```ini
[app]
Key=口令（必须大于32个字符）
TempPath=./tmp
SitePath=C:\doc\Projects\blog\public # 部署到的目录
LogPath=./log # 日志目录

[server]
RunMode=release
HttpPort=1848 # 端口号
```

你可以直接在浏览器 Console 执行下列代码获得一个比较安全的口令：

```js
function uuidv4() {
  return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
    var r = Math.random() * 16 | 0, v = c == 'x' ? r : (r & 0x3 | 0x8);
    return v.toString(16);
  });
}

console.log(uuidv4());
```

执行 `./site-deploy.exe` 后，程序开始运行。

### 本地端

通过 HTTP POST 到 `http://host:port/upload`，字段名为 `file`，内容是 zip 格式压缩的目录（注意：用 utf-8 编码压缩）。