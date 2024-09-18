package ftracker

import (
	"fmt"
	"math"
)

// Основные константы
const (
	lenStep   = 0.65  // Средняя длина шага
	mInKm     = 1000  // Количество метров в километре
	minInH    = 60    // Количество минут в часе
	kmhInMsec = 0.278 // Коэффициент для преобразования км/ч в м/с
	cmInM     = 100   // Количество сантиметров в метре
)

// distance возвращает дистанцию (в километрах), которую прошел пользователь.
func distance(action int) float64 {
	return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает среднюю скорость движения.
func meanSpeed(action int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return distance(action) / duration
}

// ShowTrainingInfo возвращает строку с информацией о тренировке.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
	switch trainingType {
	case "Бег":
		speed := meanSpeed(action, duration)
		calories := RunningSpentCalories(action, weight, duration)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance(action), speed, calories)
	case "Ходьба":
		speed := meanSpeed(action, duration)
		calories := WalkingSpentCalories(action, duration, weight, height)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance(action), speed, calories)
	case "Плавание":
		dist := float64(lengthPool*countPool) / mInKm
		speed := swimmingMeanSpeed(lengthPool, countPool, duration)
		calories := SwimmingSpentCalories(lengthPool, countPool, duration, weight)
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, dist, speed, calories)
	default:
		return "неизвестный тип тренировки"
	}
}

// Константы для расчета калорий при беге
const (
	runningCaloriesMeanSpeedMultiplier = 18   // Множитель средней скорости
	runningCaloriesMeanSpeedShift      = 1.79 // Сдвиг для расчета калорий при беге
)

// RunningSpentCalories возвращает количество потраченных калорий при беге.
func RunningSpentCalories(action int, weight, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	speed := meanSpeed(action, duration)
	return ((runningCaloriesMeanSpeedMultiplier * speed * runningCaloriesMeanSpeedShift) * weight / mInKm) * (duration * minInH)
}

// Константы для расчета калорий при ходьбе
const (
	walkingCaloriesWeightMultiplier = 0.035 // Множитель массы тела
	walkingSpeedHeightMultiplier    = 0.029 // Множитель роста
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
	if duration == 0 || height == 0 {
		return 0
	}
	avgSpeed := meanSpeed(action, duration)
	avgSpeedMS := avgSpeed * kmhInMsec // Конвертируем скорость из км/ч в м/с
	heightM := height / cmInM           // Конвертируем рост из см в метры
	return (walkingCaloriesWeightMultiplier*weight + (math.Pow(avgSpeedMS, 2)/heightM)*walkingSpeedHeightMultiplier*weight) * duration
}

// Константы для расчета калорий при плавании
const (
	swimmingCaloriesMeanSpeedShift   = 1.1  // Среднее количество сжигаемых калорий при плавании
	swimmingCaloriesWeightMultiplier = 2    // Множитель веса при плавании
)

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
	if duration == 0 {
		return 0
	}
	return float64(lengthPool*countPool) / mInKm / duration
}

// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
	if duration == 0 {
		return 0
	}
	speed := swimmingMeanSpeed(lengthPool, countPool, duration)
	return (speed + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration
}
