# DoH Local
> 为不支持`DNS-over-HTTPS`的客户端提供`DNS-over-HTTPS`到普通`DNS`协议转换

仅支持 [RFC8484](https://datatracker.ietf.org/doc/html/rfc8484) 标准

## 使用
在本地或路由器/虚拟机上启动此服务，将终端DNS配置到该服务器上。

*不建议部署到公网然后内网访问，因为可能在公网路由的过程中数据就被污染了*

## DoH服务

### Google
- https://8.8.8.8/dns-query
- https://8.8.4.4/dns-query

### OpenDNS
- https://208.67.222.222/dns-query
- https://208.67.220.220/dns-query

### Cloudflare
- https://1.1.1.1/dns-query
- https://1.0.0.1/dns-query