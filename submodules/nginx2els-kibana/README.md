## NginxのaccesslogをElasticsearchに投げる


## 準備
```
# template指定でindex, typeの作成
curl -XPOST 'localhost:9200/isucon-3' -d @elastic_template.json

# デフォルトのテンプレートでindex作成 (index名isucon-*)
curl -XPOST 'localhost:9200/_template/isucon' -d @set_template.json #templateの作成
curl -XPOST 'localhost:9200/isucon-3'

# テスト的にサンプルレコードを投げる
curl -XPOST 127.0.0.1:9200/isucon-3/access -d @test.json
``````

## 投げる

```
# accesslogを作成したindex,typeに投げる
go run ltsv2els.go /path/to/logfile host:port/index/type

# sample
go run ltsv2els.go access.log http://127.0.0.1:9200/isucon-3/access
```

## log形式サンプル
```
# This tool support only below items.

log_format ltsv
	"time:$time_local"
	"\thost:$remote_addr"
	"\tforwardedfor:$http_x_forwarded_for"
	"\treq:$request"
	"\tmethod:$request_method"
	"\turi:$request_uri"
	"\tstatus:$status"
	"\tsize:$body_bytes_sent"
	"\treferer:$http_referer"
	"\tua:$http_user_agent"
	"\treqtime:$request_time"
	"\truntime:$upstream_http_x_runtime"
	"\tapptime:$upstream_response_time"
	"\tcache:$upstream_http_x_cache"
	"\tvhost:$host"
	;
```

## Kibana
Dashbordから Settings -> Objects -> import で export_kibana.jsonを選択して読み込む

