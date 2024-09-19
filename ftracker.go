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
 
// distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Параметры:
//
// action int — количество совершенных действий (число шагов при ходьбе и беге, либо гребков при плавании).
func distance(action int) float64 {
    return float64(action) * lenStep / mInKm
}
 
// meanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки в часах.
func meanSpeed(action int, duration float64) float64 {
    if duration == 0 {
        return 0
    }
    distance := distance(action)
    return distance / duration
}
 
// ShowTrainingInfo возвращает строку с информацией о тренировке.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// trainingType string — вид тренировки(Бег, Ходьба, Плавание).
// duration float64 — длительность тренировки в часах.
func ShowTrainingInfo(action int, trainingType string, duration, weight, height float64, lengthPool, countPool int) string {
    // ваш код здесь
    switch {
    case trainingType == "Бег":
        distance := distance(action)                  // вызов функции для получения дистанции
        speed := meanSpeed(action, duration)                 // вызов функции для вычисления средней скорости
        calories := RunningSpentCalories(action, weight, duration) // вызов функции для расчета калорий
        return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
    case trainingType == "Ходьба":
        distance := distance(action)                  // вызов функции для получения дистанции
        speed := meanSpeed(action, duration)                 // вызов функции для вычисления средней скорости
        calories := WalkingSpentCalories(action, duration, weight, height) // вызов функции для расчета калорий
        return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
    case trainingType == "Плавание":
        distance := distance(action) // расчет дистанции на основе средней скорости и времени
        speed := swimmingMeanSpeed(lengthPool, countPool, duration)                      // вызов функции для вычисления средней скорости
        calories := SwimmingSpentCalories(lengthPool, countPool, duration, weight)         // вызов функции для расчета калорий
        return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", trainingType, duration, distance, speed, calories)
    default:
        return "неизвестный тип тренировки"
    }
}
 
// Константы для расчета калорий, расходуемых при беге.
const (
    runningCaloriesMeanSpeedMultiplier = 18   // множитель средней скорости.
    runningCaloriesMeanSpeedShift      = 1.79 // среднее количество сжигаемых калорий при беге.
)
 
// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// weight float64 — вес пользователя.
// duration float64 — длительность тренировки в часах.
func RunningSpentCalories(action int, weight, duration float64) float64 {
    return ((runningCaloriesMeanSpeedMultiplier * meanSpeed(action, duration) * runningCaloriesMeanSpeedShift) * weight / mInKm * duration * minInH)
}
 
// Константы для расчета калорий, расходуемых при ходьбе.
const (
    walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
    walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)
 
// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// action int — количество совершенных действий(число шагов при ходьбе и беге, либо гребков при плавании).
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
func WalkingSpentCalories(action int, duration, weight, height float64) float64 {
    // Проверка, что параметры имеют разумные значения
    if duration <= 0 || weight <= 0 || height <= 0 {
        return 0
    }
 
    // Средняя скорость в м/с
    meanSpeedSec := meanSpeed(action, duration) * kmhInMsec 
    heightInM := height / cmInM // перевод роста из см в метры
 
    // Формула для расчета потраченных калорий
    return ((walkingCaloriesWeightMultiplier * weight +
        (math.Pow(meanSpeedSec, 2) / heightInM) * walkingSpeedHeightMultiplier * weight) * duration * minInH)
}
 
// Константы для расчета калорий, расходуемых при плавании.
const (
    swimmingCaloriesMeanSpeedShift   = 1.1  // среднее количество сжигаемых колорий при плавании относительно скорости.
    swimmingCaloriesWeightMultiplier = 2    // множитель веса при плавании.
)
 
// swimmingMeanSpeed возвращает среднюю скорость при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
func swimmingMeanSpeed(lengthPool, countPool int, duration float64) float64 {
    if duration == 0 {
        return 0
    }
    return float64(lengthPool) * float64(countPool) / mInKm / duration
}
 
// SwimmingSpentCalories возвращает количество потраченных калорий при плавании.
//
// Параметры:
//
// lengthPool int — длина бассейна в метрах.
// countPool int — сколько раз пользователь переплыл бассейн.
// duration float64 — длительность тренировки в часах.
// weight float64 — вес пользователя.
func SwimmingSpentCalories(lengthPool, countPool int, duration, weight float64) float64 {
    return ((swimmingMeanSpeed(lengthPool, countPool, duration) + swimmingCaloriesMeanSpeedShift) * swimmingCaloriesWeightMultiplier * weight * duration)
}
