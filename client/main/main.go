package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"

	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"grpc_client/main/pb"
)

func main() {
	http.HandleFunc("/", uploadFile)
	http.ListenAndServe(":7600", nil)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// Получение Base64 строки из файла
		imageBase64Str, err := getImageBase64Str(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Подключение по grpc
		conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Println("Ошибка подключения к grpc серверу", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		// Клиент grpc
		client := pb.NewImageServiceClient(conn)

		// Запрос
		req := &pb.ImageRequest{Base64Image: imageBase64Str}

		// Отправка запроса и получение ответа
		resp, err := client.UploadImage(context.Background(), req)
		if err != nil {
			log.Println("Ошибка отправки сообщения или получения ответа", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Вывод изображения
		err = viewPicture(w, resp.ImageData)
		if err != nil {
			log.Println("Ошибка вывода картинки", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		fmt.Fprintln(w, `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>Загрузите файл изображения</title>
</head>
<body>
<form method="post" enctype="multipart/form-data">
    <p>
        <label>Выберите файл изображения: </label><br/>
        <input type="file" name="image"/>
    </p>
	<input type="submit" value="Загрузить"/>
</form>
</body>`)
	}
}

func getImageBase64Str(r *http.Request) (string, error) {
	file, _, err := r.FormFile("image")
	if err != nil {
		return "", err
	}
	defer file.Close()

	imageData := bytes.Buffer{}
	_, err = io.Copy(&imageData, file)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(imageData.Bytes()), nil
}

func viewPicture(w http.ResponseWriter, imageData []byte) error {
	w.Header().Set("Content-Type", "image/png")
	_, err := w.Write(imageData)
	if err != nil {
		return err
	}
	return nil
}
