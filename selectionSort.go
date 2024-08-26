package main


func selectionSort(arr []int) []int {
    n := len(arr)
    newArr := make([]int, 0, n) // Создаем пустой слайс с емкостью n
    tempArr := make([]int, n)   // Создаем временную копию исходного массива
    copy(tempArr, arr)          // Копируем исходный массив во временный

    for i := 0; i < n; i++ {
        smallestIndex := findSmallest(tempArr)
        newArr = append(newArr, tempArr[smallestIndex]) // Добавляем наименьший элемент в новый массив
        // Удаляем наименьший элемент из временного массива
        tempArr = append(tempArr[:smallestIndex], tempArr[smallestIndex+1:]...)
    }

    return newArr
}