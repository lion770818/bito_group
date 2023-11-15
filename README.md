# api 文件 製作方式

go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest

go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

go get github.com/swaggo/swag/example/celler/httputil
go get github.com/swaggo/swag/example/celler/model

go install github.com/swaggo/swag/cmd/swag@latest

swag init

swag 是執行檔 有問題就去設定
linux => PATH 例如: export PATH=$HOME/go/bin:$PATH
windos => 環境變數內設定

@Param：參數訊息，用空格分隔的參數。param name,param type,data type,is mandatory?,comment attribute(optional) 1.參數名稱 2.參數類型，可以有的值是 formData、query、path、body、header，formData 表示是 post 請求的數據， query 表示帶在 url 之後的參數，path 表示請求路徑上得參數，例如上面例子裡面的 key，body 表示是一個 raw 資料請求，header 表示帶在 header 資訊中得參數。 3.參數類型 4.是否必須 5.註釋
例如：

// @Param name query string true "用户姓名"

[常用註解格式]("https://blog.csdn.net/qq_38371367/article/details/123005909")

[swagger 教學]("https://igouist.github.io/post/2021/05/newbie-4-swagger/")

## 如果出現 Fetch error Internal Server Error http://localhost:8080/swagger/doc.json

請在專案上面 import \_ "bito_group/docs"

## 註解編譯 cannot find type definition: httputil.HTTPError

請 import 在編譯一次
"github.com/swaggo/swag/example/celler/httputil"
"github.com/swaggo/swag/example/celler/model"

# post man 測試 api 腳本, 請匯入 postman

bito_group.postman_collection.json
