# Nginx使用HTTP3(QUIC)

简介:[维基百科中关于quic的定义](https://zh.m.wikipedia.org/zh-hans/QUIC)

# 1. 安装依赖

以Ubuntu64位为例,

```bash
sudo apt update && sudo apt install build-essential mercurial psmisc lsb-release cmake golang libunwind-dev git libpcre3-dev zlib1g-dev
```

# 2.编译boringssl

由于OpenSSL官方暂不支持nginx-quic,需要使用Google基于OpenSSL开发的BoringSSL分支来提供支持,[源码地址](https://github.com/google/boringssl)

```bash
cd ~
git clone https://github.com/google/boringssl
cd boringssl/
mkdir build && cd build
cmake ../
make
#期间可能会出现golang下载包超时的情况,需要设置goproxy
#执行 
go env -w GOPROXY="https://goproxy.cn,direct"
```

# 3.编译nginx

nginx-quic项目托管在[源码地址](https://hg.nginx.org/nginx-quic)

```bash
sudo apt install mercurial
hg clone -b quic https://hg.nginx.org/nginx-quic
 ./auto/configure --with-debug --prefix=/usr/local/nginx-quic --with-http_v3_module --with-http_ssl_module --with-http_v2_module      \
                       --with-cc-opt="-I../boringssl/include"     \
                       --with-ld-opt="-L../boringssl/build/ssl    \
                                      -L../boringssl/build/crypto"
make
make install
```

nginx-quic的相关目录在/usr/local/nginx-quic下

# 4.配置nginx.conf并测试

使用文本编辑软件编辑 /usr/local/nginx-quic/conf/nginx.conf

```bash
sudo vim /usr/local/nginx-quic/conf/nginx.conf
#将之前的server{}改为下面的,如果有多个网站,reuseport只能出现一个网站中
        server {
                #listen 80;
                listen 443 ssl http2;
                listen 443 http3 reuseport;
                server_name example.com;
                index index.html index.htm;
                root /web/static/example.com;

                charset koi8-r;

                #SSL-START
                ssl_certificate /web/cert/example.pem;
                ssl_certificate_key /web/cert/example.key;
                ssl_protocols TLSv1.1 TLSv1.2 TLSv1.3;
                #SSL-END

                add_header Alt-Svc 'h3=":443"; ma=86400;quic=":443"; h3-29=":443";h3-27=":443";h3-25=":443"; h3-T050=":443"; h3-Q050=":443";h3-Q049=":443";h3-Q048=":443"; h3-Q046=":443"; h3-Q043=":443"'; # Advertise that QUIC is available;;

                #limit_conn perserver 400;
                #limit_conn perip 15;
                #limit_rate 1024k;

                #ERROR-PAGE-START
                #error_page 404 /404.html;
                #error_page 502 /502.html;
                #ERROR-PAGE-END
                access_log /web/log/example.log;
                error_log /web/log/example.error.log;
        }

#退出vim,启动nginx
sudo ./usr/local/nginx-quic/sbin/nginx
```

# 5.测试是否开启http3

- 使用[http3check](https://http3check.net/)来进行测试,下面是我的网站测试截图

![](https://minio-upload.cybeor.com:443/images/202208032139627.png)

- 使用Google chrome测试

![](https://minio-upload.cybeor.com:443/images/202208032142307.png)

# 附上编译完成的nginx以及相关配置的[下载链接😀](https://pan.cybeor.com/%E9%98%BF%E9%87%8C%E4%BA%91%E7%9B%98/nginx-quic.7z)

