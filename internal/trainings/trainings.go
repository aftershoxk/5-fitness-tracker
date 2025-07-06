package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var ErrIncorrectData = errors.New("некорректный формат данных, проверьте значения")
var ErrIncorrectTraining = errors.New("неизвестный тип тренировки")
var ErrIncorrectSteps = errors.New("неверное количество шагов")
var ErrIncorrectDuration = errors.New("")

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal

	Dataset []string
}

func (t *Training) Parse(datastring string) (err error) {
	splitData := strings.Split(datastring, ",")
	if len(splitData) != 3 {
		return fmt.Errorf("ошибка: %w", ErrIncorrectData)
	}
	t.Steps, err = strconv.Atoi(splitData[0])
	if err != nil {
		return fmt.Errorf("ошибка парсинга шагов: %w", err)
	} else if t.Steps <= 0 {
		return fmt.Errorf("ошибка: %w", ErrIncorrectSteps)
	}
	t.TrainingType = splitData[1]
	t.Duration, err = time.ParseDuration(splitData[2])
	if err != nil {
		return fmt.Errorf("ошибка парсинга времени: %w", err)
	} else if t.Duration <= 0.0 {
		return fmt.Errorf("ошибка: %w", ErrIncorrectDuration)
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	averageSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)
	var calories float64
	var err error
	switch {
	case t.TrainingType == "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("ошибка: %w", err)
		}
	case t.TrainingType == "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", fmt.Errorf("ошибка: %w", err)
		}
	default:
		return "", fmt.Errorf("ошибка: %w", ErrIncorrectTraining)
	}
	Info := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, t.Duration.Hours(), distance, averageSpeed, calories)
	return Info, nil
}
