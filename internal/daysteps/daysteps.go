package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

var ErrIncorrectSteps = errors.New("некорректное количество шагов")
var ErrIncorrectData = errors.New("некорректные данные")
var ErrIncorrectDuration = errors.New("некорректная продолжительность прогулки")
var ErrIncorrectDistance = errors.New("некорректная дистанция")

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal

	Weight float64
	Height float64
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	splitData := strings.Split(datastring, ",")
	if len(splitData) != 2 {
		return fmt.Errorf("ошибка: %w", ErrIncorrectData)
	}
	ds.Steps, err = strconv.Atoi(splitData[0])
	if err != nil {
		return fmt.Errorf("ошибка: %w", err)
	} else if ds.Steps <= 0 {
		return fmt.Errorf("ошибка: %w", ErrIncorrectSteps)
	}
	ds.Duration, err = time.ParseDuration(splitData[1])
	if err != nil {
		return fmt.Errorf("ошибка: %w", err)
	} else if ds.Duration <= 0 {
		return fmt.Errorf("ошибка: %w", ErrIncorrectDuration)
	}
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	if distance <= 0 {
		return "", fmt.Errorf("ошибка: %w", ErrIncorrectDistance)
	}
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка %w", err)
	}
	Info := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories)
	return Info, nil
}
