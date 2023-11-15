
# Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
# export DOCKER_SCAN_SUGGEST=false

# 編譯後上傳到 dockerhub liuleo=dockerhub上面的帳號, 請自行帶入自己的
docker build -t liuleo/bito_group:v3 -f ./Dockerfile  ../../
# docker build -t liuleo/bito_group:v1 . --no-cache

# 支援多種格式
#docker buildx build --push -t liuleo/bito_group:v3 --platform linux/amd64,linux/arm64 -f ./Dockerfile  ../../


