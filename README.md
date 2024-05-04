# mercury

## 微服务端口
bff 8080  
user 8091  
article 8092  
interactive 8093  
comment 8094  
captcha 8095  
sms 8096  
oauth2 8097  
ranking 8098  
crontask 8099  

## 微服务化进程
1. [x] bff
2. [x] user
3. [x] article
4. [x] interactive
5. [x] comment
6. [x] captcha
7. [x] sms
8. [x] oauth2
9. [x] ranking
10. [x] crontask

## TODO
1. [article] userrpc improve

## 启动顺序
1. user, sms, interactive, comment, ~~oauth2~~, ~~crontask~~  
2. article, captcha
3. ranking
4. bff