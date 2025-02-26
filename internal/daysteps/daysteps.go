package daysteps

import (
	"errors"
	"fmt"
	"spentCalories"
	"strconv"
	"strings"
	"time"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	// Разделяем строку на части
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("invalid package")
	}

	// Преобразуем первую часть в число (шаги)
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}

	// Преобразуем вторую часть в time.Duration (время)
	timeAction, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, err
	}

	// Возвращаем результат
	return steps, timeAction, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	// ваш код ниже
	steps, timeAction, err := parsePackage(data)
	if err != nil {
		return "parsePackage miss"
	}
	distance := float64(steps) * StepLength / 1000
	calories := WalkingSpentCalories(steps, height, weight, timeAction)
	info := fmt.Sprint(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		steps, distance, calories,
	)
	return info
}
