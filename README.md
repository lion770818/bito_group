
# api 文件
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag@latest

swag init

go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files


# post man 測試api 腳本, 請匯入 postman
bito_group.postman_collection.json