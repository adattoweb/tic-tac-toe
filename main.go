package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"strings"
)

func setSize() int {
	isValid := false
	size := 3
	for isValid == false {
		answer := "1";
		fmt.Println("Оберіть розмір поля:")
		fmt.Println("(1) 3x3")
		fmt.Println("(2) 4x4")
		fmt.Println("(3) 5x5")
		fmt.Scanln(&answer)
		if answer == "1" {
			size = 3
			isValid = true
		} else if answer == "2" {
			size = 4
			isValid = true
		} else if answer == "3" {
			size = 5
			isValid = true
		}
	}
	return size
}

func setSide() string {
	isValid := false
	side := "O"
	for isValid == false {
		answer := "1";
		fmt.Println("Оберіть сторону:")
		fmt.Println("(1) O")
		fmt.Println("(2) X")
		fmt.Scanln(&answer)
		if answer == "1" {
			side = "O"
			isValid = true
		} else if answer == "2" {
			side = "X"
			isValid = true
		}
	}
	return side
}
func setDifficult() int {
	isValid := false
	difficult := 1
	for isValid == false {
		answer := "1";
		fmt.Println("Оберіть складність:")
		fmt.Println("(1) Легка")
		fmt.Println("(2) Складна")
		fmt.Scanln(&answer)
		if answer == "1" {
			difficult = 1
			isValid = true
		} else if answer == "2" {
			difficult = 2
			isValid = true
		}
	}
	return difficult
}

func setWithBot() bool{
	isValid := false;
	withBot := true;
	for isValid == false {
		answer := "1"
		fmt.Println("Ви будете грати з ботом, чи з другом офлайн?")
		fmt.Println("(1) З ботом")
		fmt.Println("(2) З другом")
		fmt.Scanln(&answer)
		if answer == "1" {
			withBot = true;
			isValid = true
		} else if answer == "2" {
			withBot = false;
			isValid = true
		}
	}
	return withBot
}

func getField(size int, array []int, arrayLetters [5]string){
	fmt.Printf("  ")
	for i := 0; i < size; i++ { // Виводимо буквочкі
		fmt.Printf(" %s", arrayLetters[i])
	} 
	fmt.Printf("\n")
	for i := 0; i < len(array) / size; i++ { // виводимо поле
		fmt.Printf("%d:", i + 1) // циферкі по вертикалі
		for j := 0; j < size; j++ {
			print := "◻"
			if array[i * size + j] == 1 {
				print = "O"
			} else if array[i * size + j] == 2 {
				print = "X"
			}
			fmt.Printf(" %s", print)
		}
		fmt.Printf("\n")
	}
}

func checkWin(array []int, side int, size int) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == side {
			n := (i / size + 1)
			isLeft := i - i / size * size >= 2
			isRight := n * size - 1 - i >= 2 
			isBottom := size - n >= 2
			isTop := i / size >= 2

			// Рахуємо по діагоналі
			isTopRight := i / size >= 2 && n * size - 1 - i >= 2
			isTopLeft := i / size >= 2 && i - (n - 1) * size >= 2
			isBottomRight := size - n >= 2 && n * size - 1 - i >= 2
			isBottomLeft := size - n >= 2 && i - (n - 1) * size >= 2
			// (i / size + 1) -- знаходимо позицію вертикально
			if isLeft && array[i-1] == side && array[i-2] == side {
				return true
			}
			if isRight && array[i+1] == side && array[i+2] == side {
				return true
			}
			if isBottom && array[i+size] == side && array[i+size*2] == side {
				return true
			}
			if isTop && array[i-size] == side && array[i-size*2] == side {
				return true
			}
			if isTopRight && array[i-size+1] == side && array[i-size+1-size+1] == side {
				return true
			}
			if isTopLeft && array[i-size-1] == side && array[i-size-1-size-1] == side {
				return true
			}
			if isBottomRight && array[i+size+1] == side && array[i+size+1+size+1] == side {
				return true
			}
			if isBottomLeft && array[i+size-1] == side && array[i+size-1+size-1] == side {
				return true
			}
		}
	}
	return false
}

func botStep(difficult int, array []int, botValue int, playerValue int, size int) []int{
	emptyArray := []int{}
	for i := 0; i < len(array); i++ {
		if array[i] == 0 {
			emptyArray = append(emptyArray, i)
		}
	}
	if difficult == 1 { // логіка простого бота
		randN := rand.Intn(len(emptyArray))
		array[emptyArray[randN]] = botValue;
	}
	if difficult == 2 {
		for i := 0; i < len(array); i++ {
			if array[i] == 0 {
				array[i] = botValue
				if checkWin(array, botValue, size){
					return array
				}
				array[i] = 0;
			}
		}

		for i := 0; i < len(array); i++ {
			if array[i] == 0 {
				array[i] = playerValue
				if checkWin(array, playerValue, size){
					array[i] = botValue
					return array
				}
				array[i] = 0;
			}
		}

		randN := rand.Intn(len(emptyArray))
		array[emptyArray[randN]] = botValue;

	}
	return array
}

func main(){
		fmt.Println("Вітаю Вас у меню консольної гри хрестики ноліки!")
		withBot := setWithBot()
		difficult := 0;
		if withBot {
			difficult = setDifficult()
		}
		size := setSize()

		side := "O"
		if withBot {
			side = setSide()
		}
		playerValue := 1;
		botValue := 2
		if side == "X" {
			playerValue = 2
			botValue = 1
		}

		array := make([]int, size*size)
		isBotWin := false
		isPlayerWin := false;
		arrayLetters := [5]string{"A", "B", "C", "D", "E"}
		if botValue == 1 && withBot {
			if size == 3 {
				array[4] = botValue;
				fmt.Println("Бот здійснив крок на позицію: B2")
			} else if size == 4 {
				array[6] = botValue
				fmt.Println("Бот здійснив крок на позицію: C2")
			} else if size == 5 {
				array[12] = botValue
				fmt.Println("Бот здійснив крок на позицію: C3")
			}
		}
		for !isBotWin && !isPlayerWin{
			fmt.Println("Поле виглядає зараз так:")
			getField(size, array, arrayLetters)
			answer := "A1"
			num := 1;
			isValidAnswer := false
			for !isValidAnswer {
				fmt.Println("Напишіть клітинку куди будете ходити, наприклад, A1. Ваша сторона:", side)
				fmt.Scanln(&answer)
				if len(answer) > 1{
					num, _ = strconv.Atoi(string(answer[1]))
				}
				if num <= size && len(answer) == 2 {
					isValidAnswer = true;
				}
			}
			indexLetter := 0;
			answer = strings.ToUpper(answer)
			for i := 0; i < len(arrayLetters); i++ {
				if arrayLetters[i] == string(answer[0]) {				
					indexLetter = i;
					break;
				}
			}
			isValid := false
			if array[(num - 1) * size + indexLetter] == 0 {
				array[(num - 1) * size + indexLetter] = playerValue 
				isValid = true;
			}

			emptyCount := 0;
			for i := 0; i < len(array); i++ {
				if array[i] == 0 {
					emptyCount++
				}
			}

			if !withBot && isValid == true{
				if playerValue == 2 {
					playerValue = 1
					side = "O"
				} else if playerValue == 1 && isValid == true{
					playerValue = 2
					side = "X"
				}
			}

			// Перевірка перемоги, спочатку гравець потім бот.
			if !withBot {
				isPlayerWin = checkWin(array, 1, size)
				isBotWin = checkWin(array, 2, size)
			}
			if withBot {
				isPlayerWin = checkWin(array, playerValue, size)
			}
			if emptyCount == 0 && !isPlayerWin{
				isPlayerWin = true
				isBotWin = true
				break;
			} 
			if isValid == true && withBot && !isPlayerWin{
				array = botStep(difficult, array, botValue, playerValue, size)
				fmt.Println("Бот здійснив крок.")
			}
			if withBot {
				isBotWin = checkWin(array, botValue, size)
			}
		} 
		getField(size, array, arrayLetters)
		if isPlayerWin && isBotWin {
			fmt.Println("Нічия!")
		}else if isPlayerWin && withBot {
			fmt.Println("Вітаю, Ви перемогли!")
		} else if isBotWin && withBot{
			fmt.Println("Нажаль, Ви програли.")
		}
		if !withBot && side == "X" {
			side = "O"
		} else if !withBot && side == "O" {
			side = "X"
		}
		if !withBot {
			fmt.Println("Перемогла сторона:", side)
		}
		main()
}