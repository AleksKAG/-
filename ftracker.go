package ftracker

import (
    "fmt"
    "math"
)

// Основные константы, необходимые для расчетов.
const (
    lenStep   = 0.65  // средняя длина шага.
    mInKm     = 1000  // количество метров в километре.
    minInH    = 60    // количество минут в часе.
    kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
    cmInM     = 100   // количество сантиметров в метре.
)

// distance возвращает дистанцию (в километрах), которую преодолел пользователь за время тренировки.
func distance(action int) float64 {
    return float64(action) * lenStep / mInKm
}

// meanSpeed возвращает значение средней скорости движения во время тренировки.
func meanSpeed(action int, duration float64) float64 {
    if duration == 0 {
        return 0
    }
    dist := distance(action)
    return dist / duration
}

// ShowTrainingInfo выводит информацию о тренировке.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
    var dist, speed, calories float64

    switch trainingType {
    case "Бег":
        dist = distance(action)
        speed = meanSpeed(action, duration)
        calories = RunningSpentCalories(action, weight, duration)
    case "Ходьба":
        dist = distance(action)
        speed = meanSpeed(action, duration)
        calories = WalkingSpentCalories(action, duration, weight, height)
    case "Плавание":
        speed = swimmingMeanSpeed(lengthPool, countPool, duration)
        dist = float64(lengthPool*countPool) / mInKm
        calories = SwimmingSpentCalories(lengthPool, countPool, duration, weight)
    default:
        return "Неизвестный тип тренировки"
    }

    return fmt.Sprintf(
        "Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожжено калорий: %.2f ккал.",
        trainingType, duration, dist, speed, calories)
}

// Константы для расчета калорий, расходуемых при беге.
const (
    runningCaloriesMeanSpeedMultiplier = 18   // множитель средней скорости.
    runningCaloriesMeanSpeedShift      = 1.79 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories рассчитывает количество сожженных калорий при беге.
func RunningSpentCalories(action int, weight, duration float64) float64 {
    if duration <= 0 {
        return 0
    }
    speed := meanSpeed(action, duration)
    return ((runningCaloriesMeanSpeedMultiplier * speed * runningCaloriesMeanSpeedShift) * weight / mInKm * duration * minInH)
}

// Константы для расчета калорий, расходуемых при ходьбе.
const (
    walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
    walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories рассчитывает количество сожженных калорий при ходьбе.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
    if duration <= 0 || height == 0 {
        return 0
    }
    avgSpeed := meanSpeed(action, duration)
    avgSpeedMS := avgSpeed * 1000 / 3600 // Конвертируем скорость из км/ч в м/с
    return (walkingCaloriesWeightMultiplier*weight + (math.Pow(avgSpeedMS, 2)/height)*walkingSpeedHeightMultiplier*weight) * duration * minInH
}

// Константы для расчета калорий, расходуемых при плавании.
const (
    swimmingCaloriesMeanSpeedShift   = 1.1  // среднее количество сжигаемых калорий при плавании относительно скорости.
    swimmingCaloriesWeightMultiplier = 2    // множитель веса при плавании.
)

// swimmingMeanSpeed возвращает среднюю скорость при плавании.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
    if duration == 0 {
        return 0
    }
    return float64(lengthPool*countPool) / mInKm / duration
}

// SwimmingSpentCalories рассчитывает количество сожженных калорий при плавании.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
    if duration <= 0 {
        return 0
    }
    speed := swimmingMeanSpeed(lengthPool, countPool, duration)
    return (speed + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration
}
