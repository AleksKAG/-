package ftracker

import (
	"fmt"
	"math"
)

// Константы для расчетов
const (
	minInH        = 60.0   // Минут в часе
	mInKM         = 1000.0 // Метров в километре
	factorRun     = 18.0
	factorRun2    = 1.79
	factorWalk1   = 0.035
	factorWalk2   = 0.029
	swimConstant1 = 1.1
	swimConstant2 = 2.0
)

// distance вычисляет дистанцию в километрах
func distance(speed, time float64) float64 {
	return speed * time
}

// meanSpeed вычисляет среднюю скорость для бега или ходьбы в км/ч.
func meanSpeed(distance, time float64) float64 {
	return distance / time
}

// swimmingMeanSpeed вычисляет среднюю скорость для плавания в км/ч.
func swimmingMeanSpeed(distance, time float64) float64 {
	return distance / time
}

// ShowTrainingInfo выводит информацию о тренировке на основе типа тренировки.
func ShowTrainingInfo(trainingType string, duration, distanceKM, weight, height float64) string {
	var avgSpeed float64
	var calories float64

	// Вычисляем среднюю скорость и сожженные калории в зависимости от типа тренировки
	switch trainingType {
	case "Бег":
		avgSpeed = meanSpeed(distanceKM, duration)
		calories = RunningSpentCalories(avgSpeed, weight, duration)
	case "Ходьба":
		avgSpeed = meanSpeed(distanceKM, duration)
		calories = WalkingSpentCalories(avgSpeed, weight, height, duration)
	case "Плавание":
		avgSpeed = swimmingMeanSpeed(distanceKM, duration)
		calories = SwimmingSpentCalories(avgSpeed, weight, duration)
	default:
		return "Неизвестный тип тренировки"
	}

	// Форматируем строку вывода с информацией о тренировке
	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nРасход калорий: %.2f ккал.",
		trainingType, duration, distanceKM, avgSpeed, calories)
}

// RunningSpentCalories рассчитывает количество сожженных калорий во время бега.
func RunningSpentCalories(avgSpeed, weight, duration float64) float64 {
	return ((factorRun * avgSpeed * factorRun2) * weight / mInKM * duration * minInH)
}

// WalkingSpentCalories рассчитывает количество сожженных калорий во время ходьбы.
func WalkingSpentCalories(avgSpeedKMH, weight, height, duration float64) float64 {
	avgSpeedMS := avgSpeedKMH * 1000 / 3600 // Конвертируем скорость из км/ч в м/с
	return ((factorWalk1*weight + (math.Pow(avgSpeedMS, 2)/height)*factorWalk2*weight) * duration * minInH)
}

// SwimmingSpentCalories рассчитывает количество сожженных калорий во время плавания.
func SwimmingSpentCalories(avgSpeed, weight, duration float64) float64 {
	return (avgSpeed + swimConstant1) * swimConstant2 * weight * duration
}
