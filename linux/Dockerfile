# build stage
FROM golang:1.21-alpine AS build-env

# 拷貝需要的專案
ADD bito_group /src/bito_group

# 編譯server執行檔
RUN cd /src/bito_group && go build -o bito_group

# final stage
#FROM alpine
FROM ubuntu
WORKDIR /app
# 拷貝編譯好的執行檔案 bito_group
COPY --from=build-env /src/bito_group/bito_group /app/

#bito_group 開放的port 
EXPOSE 8080


ENTRYPOINT ./bito_group
