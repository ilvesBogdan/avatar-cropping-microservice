# avatar-cropping-microservice

Этот репозиторий содержит проект микросервиса, который обрабатывает аватары пользователей. Микросервис получает изображения по протоколу gRPC и возвращает обрезанную версию изображения в квадратном формате и в произвольном формате изображения.

## Устновка

Для ручной сборки docker контенера, выполните команду
```bash
docker build -t avatar-cropping .
```
Для запуска выполните
```bash
docker run -p 50051:50051 avatar-cropping
```

## Запуск!

Для локального запуска можно создать виртуальное окружение Python
```bash
python -m venv .env
```
далее на необходимо активировать его командой
```bash
source .env/bin/activate
```
установим необходимые зависимости
```bash
pip install Pillow
pip install grpcio-tools
pip install grpcio
```
и теперь можно запустить
```bash
python server/src/main.py
```
микросервис будет запущен на порту `50051`.

## Запуск тестового клиента на golang

Для запуска клиента необходимо установить следующие зависимости
```bash
go mod init grpc_client
go mod vendor
go mod tidy
```
затем запускаем клиент командой
```bash
go run main/main.go
```
после этого перейдите по ссылке `http://localhost:7600`
и загрузите изображение через форму, после этого результат будет выведен в браузер.


## Компиляции файлов протокола буферизации

Для python
```bash
python -m grpc_tools.protoc -I. --python_out=./src --grpc_python_out=./src grpc.proto
```

Для тестового клиента на golang
необходимо использовать утилиту [protobuf](https://github.com/protocolbuffers/protobuf/releases)
```bash
protoc --go_out=. --go-grpc_out=. grpc.proto
```
