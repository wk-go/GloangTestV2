

来源url:https://medium.com/@mlowicki/http-s-proxy-in-golang-in-less-than-100-lines-of-code-6a51c2f2c38c

use chrome
```shell
Chrome --proxy-server=https://localhost:8888
```
use curl
```shell
curl -Lv --proxy http://127.0.0.1:8888 "https://httpbin.org"
```

use ssh:
```shell
ssh -o ProxyCommand="nc -X connect -x 127.0.0.1:8888 %h %p" username@server
```