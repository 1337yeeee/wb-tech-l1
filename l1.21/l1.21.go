package main

import "fmt"

// Реализовать паттерн проектирования «Адаптер» на любом примере.
// Описание: паттерн Adapter позволяет сконвертировать интерфейс
// одного класса в интерфейс другого, который ожидает клиент.

// Продемонстрируйте на простом примере в Go:
// у вас есть существующий интерфейс (или структура) и другой,
// несовместимый по интерфейсу потребитель — напишите адаптер,
// который реализует нужный интерфейс и делегирует вызовы к встроенному объекту.

// Поясните применимость паттерна, его плюсы и минусы,
// а также приведите реальные примеры использования.

// Старая система (несовместимый интерфейс)
type LegacyAuthSystem struct{}

func (a *LegacyAuthSystem) Auth(username, password string) bool {
	fmt.Printf("Legacy auth: %s/%s\n", username, password)
	return true
}

// Новый интерфейс, который ожидает клиент
type ModernAuthenticator interface {
	Login(credentials map[string]string) error
}

type AuthAdapter struct {
	legacy *LegacyAuthSystem
}

func NewAuthAdapter(legacy *LegacyAuthSystem) ModernAuthenticator {
	return &AuthAdapter{legacy: legacy}
}

func (a *AuthAdapter) Login(creds map[string]string) error {
	// Явное преобразование интерфейсов:
	// 1. Извлекаем данные из нового формата
	username, ok1 := creds["user"]
	password, ok2 := creds["pass"]

	// 2. Проверяем наличие требуемых полей
	if !ok1 || !ok2 {
		return fmt.Errorf("invalid credentials format")
	}

	// 3. Адаптируем вызов к старому интерфейсу
	success := a.legacy.Auth(username, password)
	if !success {
		return fmt.Errorf("authentication failed")
	}
	return nil
}

func main() {
	// Клиентский код работает только с ModernAuthenticator
	var authSystem ModernAuthenticator = NewAuthAdapter(&LegacyAuthSystem{})

	// Новый формат credentials
	credentials := map[string]string{
		"user": "john",
		"pass": "qwerty",
	}

	if err := authSystem.Login(credentials); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Successfully authenticated")
	}
	// Legacy auth: john/qwerty
	// Successfully authenticated
}

// Применимость:
// - Когда нужно использовать существующий класс, но его интерфейс не соответствует требованиям (например, адаптация старых SOAP-сервисов к REST API)
// - Для интеграции сторонних библиотек (или обертка C-библиотек для использования в Go)
// - Для работы с устаревшим кодом (адаптация разных протоколов, например, AMQP -> HTTP)

// Плюсы:
// - Отделяет и скрывает преобразование интерфейсов от клиента (Клиентский код не знает о LegacyAuthSystem)
// - Позволяет повторно использовать существующий код (Старая система не требует изменений)
// - Реализует принцип открытости/закрытости (Адаптер инкапсулирует всю логику преобразования)

// Минусы:
// - Усложняет код дополнительными классами
// - Может снижать производительность из-за дополнительного слоя
