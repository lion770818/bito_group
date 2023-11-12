
# 編譯後上傳到 dockerhub liuleo=dockerhub上面的帳號, 請自行帶入自己的
docker build -t liuleo/bito_group:v1 -f ./Dockerfile  ../../
# docker build -t liuleo/bito_group:v1 . --no-cache

docker buildx build --push -t liuleo/bito_group:v1 --platform linux/amd64,linux/arm64 -f ./Dockerfile  ../../


# 編譯後上傳到自己筆電的倉庫
#docker build -t test:v6 -f ./Dockerfile  ../../
