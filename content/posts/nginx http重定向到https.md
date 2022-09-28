---
title: "Nginx HTTP请求跳转HTTPS"
date: 2020-07-20
draft: false
---

# 1.使用nginx的rewrite功能将http修改为https
```conf
server {
    listen 80;
    #将http的URL rewrite为https
    rewrite ^(.*) https://$server_name$1 permanent; 
}

server {
    listen 443 ssl http2;
    #证书域名
    server_name www.xxx.com;
    ssl_certificate /web/cert/example.pem;
    ssl_certificate_key /web/cert/example.key;
    ssl_protocols TLSv1.1 TLSv1.2 TLSv1.3;
}
```

# 2.使用301重定向将http请求重定向到https
```conf
server {
    listen 80;
    #将http重定向为https
    return 301 https://$host$request_uri;
}

server {
    listen 443 ssl http2;
    #证书域名
    server_name www.xxx.com;
    ssl_certificate /web/cert/example.pem;
    ssl_certificate_key /web/cert/example.key;
    ssl_protocols TLSv1.1 TLSv1.2 TLSv1.3;
}
```

# 3.使用error_page转发
```conf
server {
    listen 80;
    listen 443 ssl;
    server_name www.xxx.com;
    ssl on;
    ssl_certificate     /etc/nginx/ssl/domain.com.crt; 
    ssl_certificate_key /etc/nginx/ssl/domain.com.crt;
    # other
    error_page 497 https://$server_name$request_uri;
}
```

## 建议使用301重定向，下面是我的网站截图
![](https://minio-upload.cybeor.com/images/202209270933238.png)