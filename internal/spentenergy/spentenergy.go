package spentenergy

import (
	"errors"
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

var ErrInсorrectValue = errors.New("некорректные значения")

func CheckParameters(steps int, weight, height float64, duration time.Duration) bool {
	if steps > 0 && weight > 0.0 && height > 0.0 && duration > 0 {
		return true
	} else {
		return false
	}
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if !CheckParameters(steps, weight, height, duration) {
		return 0.0, fmt.Errorf("ошибка: %w", ErrInсorrectValue)
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	numberOfCalories := (weight * averageSpeed * durationInMinutes) / minInH
	walkSpentCalories := numberOfCalories * walkingCaloriesCoefficient
	return walkSpentCalories, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if !CheckParameters(steps, weight, height, duration) {
		return 0.0, fmt.Errorf("ошибка: %w", ErrInсorrectValue)
	}
	averageSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	runSpentCalories := (weight * averageSpeed * durationInMinutes) / minInH
	return runSpentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= 0 {
		return 0.0
	}
	distance := Distance(steps, height)
	meanSpeed := distance / duration.Hours()
	return meanSpeed
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distance := (stepLength * float64(steps)) / mInKm
	return distance
}
