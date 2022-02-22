# DNS Sec
> 为不支持`DoT/DoH`的客户端提供`DoT/DoH``到普通`DNS`协议转换

支持 [RFC8484](https://datatracker.ietf.org/doc/html/rfc8484) [RFC7858](https://datatracker.ietf.org/doc/html/rfc7858) 标准

## 使用
在本地或路由器/虚拟机上启动此服务，将终端DNS配置到该服务器上。

DoT
```shell
./dnsec 208.67.222.222:853
```

DoH
```shell
./dnsec https://208.67.222.222/dns-query
```

*不建议部署到公网然后内网访问，因为可能在公网路由的过程中数据就被污染了*

## DoH/DoT服务

### Google
- https://8.8.8.8/dns-query
- https://8.8.4.4/dns-query
- 8.8.8.8:853
- 8.8.4.4:853

### OpenDNS
- https://208.67.222.222/dns-query
- https://208.67.220.220/dns-query
- 208.67.220.220:853
- 208.67.220.220:853

### Cloudflare
- https://1.1.1.1/dns-query
- https://1.0.0.1/dns-query
- 1.1.1.1:853
- 1.0.0.1:853