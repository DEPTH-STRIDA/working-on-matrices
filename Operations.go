package main

import (
	"fmt"
)

func main() {
	mat1 := [][]float64{{2, 4, 1}, {5, 3, 1}, {7, 3, 1}}
	mat2 := [][]float64{{7, 5, 3}, {9, 7, 5}, {4, 7, 2}}
	mat11 := [][]float64{{1, 3}, {2, 4}}
	mat22 := [][]float64{{5, 7}, {6, 8}}
	a, b, c := summOfMatrices(mat1, mat2)
	fmt.Println("Сумма матриц", a, b, c)
	a, b, c = diffMatrices(mat1, mat2)
	fmt.Println("Разница матриц", a, b)
	a, b, c = multMatrixAtNumber(-2, mat1)
	fmt.Println("Умножение матрицы на число", a, b)
	a, b, c = multMatrixAtMatrix(mat1, mat2)
	fmt.Println("Умножение матрицы на матрицу", a, b)
	a, b, c = multMatrixAtMatrix(mat11, mat22)
	fmt.Println("Умножение матрицы на матрицу", a, b)
	a, b, c = expMatrices(0, mat2)
	fmt.Println("Возведение в степень матрицы", a, b)
	a, b, c = expMatrices(3, mat11)
	fmt.Println("Возведение в степень матрицы", mat11, "Ответ:", a, b)
	a, b, c = transMatrices(mat1)
	fmt.Println("Транспорнирование матрицы", mat1, "Ответ:", a, b)
}

func summOfMatrices(mat1 [][]float64, mat2 [][]float64) ([][]float64, string, bool) {
	toReturnMat := make([][]float64, len(mat1))
	len1 := len(mat1)
	for _, v := range mat1 {
		if len(v) != len1 {
			return toReturnMat, "Сумма не квадратных матриц невозможна.", false
		}
	}
	for i := range toReturnMat {
		toReturnMat[i] = make([]float64, (len(mat1)))
	}
	var sN, sV, s string
	for i := 0; i < len(mat1); i++ {
		for k := 0; k < len(mat1); k++ {
			toReturnMat[i][k] = mat1[i][k] + mat2[i][k]
			sN = "С(" + fmt.Sprint(k) + ";" + fmt.Sprint(i) + ") = A(" + fmt.Sprint(k) + ";" + fmt.Sprint(i) + ")" + " + B(" + fmt.Sprint(k) + ";" + fmt.Sprint(i) + ")"
			sV = "= " + fmt.Sprint(mat1[i][k]) + " + " + fmt.Sprint(mat2[i][k]) + " = " + fmt.Sprint(toReturnMat[i][k]) + ". "
			s += sN + sV
		}
	}
	return toReturnMat, s, true
}
func diffMatrices(mat1 [][]float64, mat2 [][]float64) ([][]float64, string, bool) {
	toReturnMat := make([][]float64, len(mat1))
	len1 := len(mat1)
	for _, v := range mat1 {
		if len(v) != len1 {
			return toReturnMat, "Разность не квадратных матриц невозможна.", false
		}
	}
	for i := range toReturnMat {
		toReturnMat[i] = make([]float64, (len(mat1)))
	}
	var sN, sV, s string
	for i := 0; i < len(mat1); i++ {
		for k := 0; k < len(mat1); k++ {
			toReturnMat[i][k] = mat1[i][k] - mat2[i][k]
			sN = "С(" + fmt.Sprint(k) + ";" + fmt.Sprint(i) + ") = A(" + fmt.Sprint(k) + ";" + fmt.Sprint(i) + ")" + " - B(" + fmt.Sprint(k) + ";" + fmt.Sprint(i) + ")"
			sV = "= " + fmt.Sprint(mat1[i][k]) + " - " + fmt.Sprint(mat2[i][k]) + " = " + fmt.Sprint(toReturnMat[i][k]) + ". "
			s += sN + sV
		}
	}
	return toReturnMat, s, true
}
func multMatrixAtNumber(num float64, mat1 [][]float64) ([][]float64, string, bool) {
	toReturnMat := make([][]float64, len(mat1))
	for i := range toReturnMat {
		toReturnMat[i] = make([]float64, (len(mat1)))
	}
	var sN, sV, s string
	for i := 0; i < len(mat1); i++ {
		for k := 0; k < len(mat1); k++ {
			toReturnMat[i][k] = num * mat1[i][k]
			sN = "C(" + fmt.Sprint(k) + ";" + fmt.Sprint(i) + ") = " + fmt.Sprint(num) + " * A(" + fmt.Sprint(k) + ";" + fmt.Sprint(i) + ")"
			sV = "= " + fmt.Sprint(num) + " * " + fmt.Sprint(mat1[i][k]) + " = " + fmt.Sprint(toReturnMat[i][k]) + ". "
			s += sN + sV
		}
	}
	return toReturnMat, s, true
}
func multMatrixAtMatrix(mat1 [][]float64, mat2 [][]float64) ([][]float64, string, bool) {
	toReturnMat := make([][]float64, len(mat1))
	if len(mat1) != len(mat2[0]) {
		return toReturnMat, "Количество строк первой матрицы отличается от количества столбцов второй матрицы.", false
	}
	for i := range toReturnMat {
		toReturnMat[i] = make([]float64, (len(mat1)))
	}
	var sS, sN, s string
	first := true
	for x := 0; x < len(mat1); x++ { //По горизонтали
		for y := 0; y < len(mat2[0]); y++ { //По вертикали
			sS = ". C(" + fmt.Sprint(y) + ";" + fmt.Sprint(x) + ") = "
			sN = " = "
			if first == true {
				first = false
				sS = " C(" + fmt.Sprint(y) + ";" + fmt.Sprint(x) + ") = "
			}
			for i := 0; i < len(mat1[0]); i++ { //Сложение произведений
				if i != 0 {
					sS += "+ "
					sN += "+"
				}
				toReturnMat[x][y] += mat1[i][x] * mat2[y][i]
				sS += "A(" + fmt.Sprint(x) + ";" + fmt.Sprint(i) + ")*B(" + fmt.Sprint(i) + ";" + fmt.Sprint(y) + ")"
				sN += fmt.Sprint(mat1[i][x]) + "*" + fmt.Sprint(mat2[y][i])
			}
			s += sS + sN
		}
	}
	return toReturnMat, s, true
}
func expMatrices(num int, mat1 [][]float64) ([][]float64, string, bool) {
	toReturnMat := make([][]float64, len(mat1))
	len1 := len(mat1)
	for _, v := range mat1 {
		if len(v) != len1 {
			return toReturnMat, "Только квадратную матрицу можно возводить в степень", false
		}
	}
	if num < 0 {
		return toReturnMat, "Степень не должна быть отрицательной.", false
	}
	for i := range toReturnMat {
		toReturnMat[i] = make([]float64, (len(mat1)))
	}
	if num == 0 {
		for i := 0; i < len(mat1); i++ {
			for j := 0; j < len(mat1); j++ {
				if i == j {
					toReturnMat[i][j] = 1
				} else {
					toReturnMat[i][j] = 0
				}
			}
		}
		return toReturnMat, "Матрица в нулевой степени равна единичной матрице.", true
	}
	if num == 1 {
		return mat1, "Матрица в первой степени равна этой матрице.", true
	}
	temp := mat1
	stroka, s := "Возведение в степень - это умножение матрицы n-ое количество раз по порядку", ""
	for i := 0; i < num-1; i++ {
		stroka += fmt.Sprint(temp) + "*" + fmt.Sprint(mat1)
		//fmt.Println(fmt.Sprint(temp) + "*" + fmt.Sprint(mat1))
		temp, s, _ = multMatrixAtMatrix(temp, mat1)
		stroka += s
	}
	return temp, stroka, true
}

func transMatrices(mat1 [][]float64) ([][]float64, string, bool) {
	toReturnMat := make([][]float64, len(mat1))
	for i := range toReturnMat {
		toReturnMat[i] = make([]float64, (len(mat1)))
	}
	s := ""
	for i := 0; i < len(mat1); i++ {
		for j := 0; j < len(mat1); j++ {
			s += "A(" + fmt.Sprint(i) + ";" + fmt.Sprint(j) + ") = " + "A(" + fmt.Sprint(j) + ";" + fmt.Sprint(i) + "). "
			toReturnMat[j][i] = mat1[i][j]
		}
	}
	return toReturnMat, s, true
}
