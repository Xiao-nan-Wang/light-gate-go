# 快速启动
## 1.下载发行版本

## 2.启动程序
```
./lg-linux -port 8080
```
-port用于指定light-gate的运行端口,默认运行于8080.

## 3.配置微服务
### 3.1 手动配置
向light-gate发送心跳,间隔5s发送一次
>POST  http://ip:port/heartbeat
> 
>请求体: (application/json)
> 
>{"name": "default","ip": 8080}
### 3.2 引入代码自动配置
#### A.springboot框架
新建HeartBeat.java文件
```
@Component
public class HeartBeat {
    @Value("${server.port}")
    public String port;
    @Value("${lightgate.serviceName}")
    public String serviceName;
    @Value("${lightgate.url}")
    public String url;
    @Autowired
    private RestTemplate restTemplate;

    @Scheduled(fixedRate = 5000)
    public void sendHeartBeat(){
        HeartBeatDto dto = new HeartBeatDto();
        dto.setPort(port);
        dto.setName(serviceName);
        String fullPath = "http://" + url + "/heartBeat";
        restTemplate.postForObject(fullPath, dto, String.class);
    }
    @Data
    private static class HeartBeatDto {
        String name;
        String port;
    }
}
```
并且在application.properties里加入
```
server.port=8080
lightgate.serviceName=default
lightgate.url=127.0.0.1:8080
```
B.其他框架(待补充)
## SSL支持
在light-gate程序所在位置创建conf文件夹,下载crt和key格式证书放入conf文件夹并改名server.

light-gate在启动时会检测是否存在ssl证书信息,存在则开启https模式,否则开启http模式.

目录结构如下:
```
...
│   lg_linux_linux
│   lg_windows.exe
│
└───conf
        server.crt
        server.key
```
# 项目说明
## 1.概述
是一个用gin框架构建的负载均衡器，服务通过向负载均衡器发送心跳来注册自己，负载均衡器会在活跃服务里选择一个处理请求.它运行只需要6MB内存!
## 2.注册策略
10s内负载均衡器接收不到心跳将会下线服务。
## 3.负载均衡策略
目前只支持轮询。
## 4.代理规范
代理路径规范：
```
http://ip:port/serviceName/url...
```

只转发形如/*/**的二层以上的路径。形如/*的单层路径是lightgate的保留路径。

在转发时，会去除服务名称，只转发上面的url...部分。
## 5.支持
（虽然大概率没人用）但是如果您需要支持，欢迎联系。besides，本项目也会持续更新。