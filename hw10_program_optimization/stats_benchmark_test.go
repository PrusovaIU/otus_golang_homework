package hw10programoptimization

import (
	"bytes"
	"encoding/json"
	"strconv"
	"testing"
)

func BenchmarkGetUsers(b *testing.B) {
	// Создаём тестовые данные: JSON-строки с пользователями
	var sb bytes.Buffer
	n := 1000
	for i := 0; i < n; i++ {
		newUser := User{
			ID:       i,
			Name:     "User " + strconv.Itoa(i),
			Username: "user" + strconv.Itoa(i),
			Email:    "user" + strconv.Itoa(i) + "@example.com",
			Phone:    "phone" + strconv.Itoa(i),
		}
		data, err := json.Marshal(newUser)
		if err != nil {
			panic(err)
		}
		sb.Write(data)
		if i < n-1 {
			sb.WriteString("\n")
		}
	}
	reader := bytes.NewBuffer(sb.Bytes())

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = getUsers(reader)
		reader.Reset() // Перезаписываем reader для следующей итерации
		reader.Write(sb.Bytes())
	}
}
