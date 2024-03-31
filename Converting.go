package main

//Данная программа позволяет переводить Миллилитры, Литры и Галлоны между собой
import (
	"fmt"
	"keybord" //Данный пакет кастомный, позволяет считывать текст с клавиатуры и переводить в тип string
	"log"
	"strconv"
)

// Создаем типы для обозначения единиц
type liter float64
type gallon float64
type milliliter float64

// Методы для перевода
func (l *liter) toMilliliter() {
	*l *= 1000
}
func (g *gallon) toMilliliter() {
	*g *= 3785.41
}
func (l *liter) toGallon() {
	*l *= 0.264172
}
func (m *milliliter) toGallon() {
	*m *= 0.00026
}
func (m *milliliter) toLiter() {
	*m /= 1000
}
func (g *gallon) toLiter() {
	*g *= 3.79
}

func main() {
	var numL liter
	var numG gallon
	var numM milliliter
	fmt.Println("Привет! Ты хочешь перевести какие-либо единицы в другую единицу измерения?\n Какую единицу нужно перевести?\n(Литры, Миллилитры, Галлоны)")
	unit := keybord.Keybord()
	fmt.Println("Хорошо! В какую единицу нужно перевести?\n(Литры, Миллилитры или Галлоны?")
	unitConvert := keybord.Keybord()
	fmt.Println("Какое количество нужно перевести?")
	number := keybord.Keybord()
	numberF, err := strconv.ParseFloat(number, 64) //Переводим наше значение из формата string в float64
	if err != nil {                                //При вводе текста выдается ошибка

		log.Fatal("Пока!")
	}
	if unit == "Литры" { //Далее условия перевода в ту или иную единицу
		numL = liter(numberF)
		if unitConvert == "Миллилитры" {
			numL.toMilliliter()
			fmt.Printf("Получится %0.1f миллилитров", numL)
		} else if unitConvert == "Галлоны" {
			numL.toGallon()
			fmt.Printf("Получится %0.2f галлонов", numL)
		} else {
			fmt.Println("\n А зачем тебе это?") //Вывод вопроса в случае, если перевод нужен в ту же единицу или если при введении единиц сделали ошибку
		}
	}
	if unit == "Миллилитры" {
		numM = milliliter(numberF)
		if unitConvert == "Литры" {
			numM.toLiter()
			fmt.Printf("Получится %0.3f литров", numM)
		} else if unitConvert == "Галлоны" {
			numM.toGallon()
			fmt.Printf("Получится %0.4f галлонов", numM)
		} else {
			fmt.Println("\n А зачем тебе это?")
		}
	}
	if unit == "Галлоны" {
		numG = gallon(numberF)
		if unitConvert == "Литры" {
			numG.toLiter()
			fmt.Printf("Получится %0.2f литров", numG)
		} else if unitConvert == "Миллилитры" {
			numG.toMilliliter()
			fmt.Printf("Получится %0.1f миллилитров", numG)
		} else {
			fmt.Println("\n А зачем тебе это?")
		}
	}
	fmt.Println("\nРады были помочь! ")
}
