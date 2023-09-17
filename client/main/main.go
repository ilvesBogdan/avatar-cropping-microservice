package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"encoding/base64"
	"log"
	"net/http"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"grpc_client/main/pb"
)

func main() {
	http.HandleFunc("/", uploadAndOutputFile)
	http.ListenAndServe(":7600", nil)
}

func uploadAndOutputFile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		// Получение файла из запроса
		file, _, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Проверка размера файла
		if 4194304 < r.ContentLength {
			http.Error(w, "Слишком большой размер файла", http.StatusBadRequest)
			return
		}

		// Подключение по grpc
		conn, err := grpc.Dial(
			"localhost:50051",
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Println("Ошибка подключения к grpc серверу", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer conn.Close()

		// Клиент grpc
		client := pb.NewImageServiceClient(conn)

		// Создание буфера для хранения байтов изображения
		buffer := bytes.NewBuffer(nil)
		_, err = io.Copy(buffer, file)
		if err != nil {
			log.Println("Ошибка чтения файла изображения", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Запрос
		// Раземер грани изображения и формат
		formats := []*pb.Format{
			{Size: 650, Format: "jpeg"},
			{Size: 200, Format: "jpeg"},
			{Size: 650, Format: "webp"},
			{Size: 200, Format: "webp"},
		}
		request := &pb.ImageRequest{
			RawImage: buffer.Bytes(),
			Formats:  formats,
		}

		// Установка ограничения на исполнение запроса
		ctx, cancel := context.WithTimeout(context.Background(), 4*time.Minute)
		defer cancel()

		// Отправка запроса и получение ответа
		stream, err := client.UploadImage(ctx, request)
		if err != nil {
			// Остальные ошибки
			log.Println("Ошибка отправки сообщения или получения ответа", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Получение изображений из итератора
		var sectionsImg []string
		var i = 0
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {

				// Ошибка получения изображения
				if st, ok := status.FromError(err); ok {
					switch st.Code() {
					case codes.InvalidArgument:
						switch st.Message() {
						case "Invalid file transfer format":
							http.Error(w, "Неверный формат файла", http.StatusBadRequest)
							return
						}
					}
				}

				// Остальные ошибки
				log.Println("Ошибка получения ответа", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			sectionsImg = append(sectionsImg, getImgTag(resp.ImageData, formats[i]))
			i = i + 1
		}

		// Вывод изображения
		fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>Вывод изображений</title>
</head>
<body>
	<h1>Вывод изображений</h1>
	%s
</body>
</html>`, strings.Join(sectionsImg, "\n"))

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

func getImgTag(imageData []byte, format *pb.Format) string {
	encodedString := base64.StdEncoding.EncodeToString(imageData)
	return fmt.Sprintf(`<figure>
	<figcaption>%v – %v</figcaption>
	<img src="data:image/jpeg;base64,%s">
</figure>`, format.Format, format.Size, encodedString)
}
