package main

import (
	"fmt"
	"math/big"
	"sync"
	"time"
)

// Comienza la generación de números
func generateSeed() int {
	seed := getTime() / 16
	for true {
		if big.NewInt((int64)(seed)).ProbablyPrime(0) {
			break
		}
		seed += 1

	}
	return seed

}

// Genera una semilla con el reloj del sistema
func getTime() int {
	year := time.Now().Year()
	mes := time.Now().Month()
	dia := time.Now().Day()
	hour := time.Now().Hour()
	minutos := time.Now().Minute()
	segundos := time.Now().Second()
	seed := ((((year*100+int(mes))*100+dia)*100+hour)*100 + minutos) / 50 * segundos
	return seed
}

// Normaliza el numero dado en el rango de 0 a 31
func NormalizeRandom(randomNum float64, upperbound int16, lowerbound int16) int {
	RandomNum := 0 + (int)(randomNum*(((float64)(upperbound)-(float64)(lowerbound))+1))
	return RandomNum
}

// Genera un número aleatorio en el intervalo de [0,1[
func generateRandom(wg *sync.WaitGroup, channel chan int, arrayChannel chan []int, size int) {
	seed := generateSeed()
	randomArray := []int{}
	period := 8192
	constant := 23
	multiplicativeConstant := 3 + (8 * constant)
	for i := 0; i < size; i++ {
		num := ((seed * multiplicativeConstant) + constant) % int(period)
		seed = num
		normalizedNum := float64(num) / float64(period-1)
		randomNum := NormalizeRandom(normalizedNum, 31, 0)
		randomArray = append(randomArray, randomNum)
	} //For end
	arrayChannel <- randomArray
}

// Función que ejecuta el bubblesort
func BubbleSort(wg *sync.WaitGroup, randomArray []int, bubbleArray chan []int) {
	for true { //ciclo que atravieza el array multiples veces hasta no necesitar mas cambios
		var num1 = 0 //declaracion de variables
		var num2 = 1
		var numAux = 0
		var cambio = false
		for num1 < len(randomArray) && num2 < len(randomArray) { //ciclo de recorrida actual del array
			posiciones := []int{}
			if randomArray[num1] > randomArray[num2] {
				numAux = randomArray[num1]            //si el numero actual es x es mayor al numero x+1
				randomArray[num1] = randomArray[num2] //intercambian lugares
				randomArray[num2] = numAux
				cambio = true //se indica que en la recorrida actual hubo cambios
				posiciones = append(posiciones, num1)
				posiciones = append(posiciones, num2)
				bubbleArray <- posiciones
				<-bubbleArray
			}
			num1++
			num2++
		}
		if cambio == false { //si en la recorrida no hubo cambios se rompe el ciclo
			break
		}
	}
	posiciones := []int{}
	posiciones = append(posiciones, 0)
	posiciones = append(posiciones, 0)
	bubbleArray <- posiciones
	<-bubbleArray
	defer wg.Done()
	//arrChan <- randomArray
}

// Estructura del heap
type maxheap struct {
	arr []int
}

// Crea una estructura de heap
func newMaxHeap(arr []int) *maxheap {
	maxheap := &maxheap{
		arr: arr,
	}
	return maxheap
}

// Retorna el índice del hijo izquierdo de un nodo
func (m *maxheap) indiceIzquierdo(index int) int {
	return 2*index + 1
}

// Retorna el índice del hijo derecho de un nodo
func (m *maxheap) indiceDerecho(index int) int {
	return 2*index + 2
}

// Intercambia 2 elementos de un array entre si
func (m *maxheap) swap(primero, segundo int) {
	m.arr[primero], m.arr[segundo] = m.arr[segundo], m.arr[primero]
}

func (m *maxheap) leaf(index int, lenght int) bool {
	if index >= (lenght/2) && index <= lenght {
		return true
	}
	return false
}

// Se encarga de hacer un heap a partir de un array , esto comparando la raíz con sus hijos , si alguno de ellos es más grande que la raíz se intercambia su posición
func (m *maxheap) heapify(current int, lenght int) {
	if m.leaf(current, lenght) {
		return
	}
	mayor := current
	indiceIzquierdo := m.indiceIzquierdo(current)
	rightRightIndex := m.indiceDerecho(current)
	if indiceIzquierdo < lenght && m.arr[indiceIzquierdo] > m.arr[mayor] {
		mayor = indiceIzquierdo
	}
	if rightRightIndex < lenght && m.arr[rightRightIndex] > m.arr[mayor] {
		mayor = rightRightIndex
	}
	if mayor != current {
		m.swap(current, mayor)
		m.heapify(mayor, lenght)
	}
	return
}

// Recibe un array y lo convierte en un maxHeap
func (m *maxheap) buildMaxHeap(lenght int) {
	for index := ((lenght / 2) - 1); index >= 0; index-- {
		m.heapify(index, lenght)
	}
}

// Sortea el max heap , esto mediante el método de tomar la raíz , que es el mayor y lo envía al final
func (m *maxheap) sort(lenght int, start time.Time) {
	m.buildMaxHeap(lenght)
	for i := lenght - 1; i > 0; i-- {
		// Obtiene la raíz y la "elimina " del heap
		m.swap(0, i)
		m.heapify(0, i)
	}

	//fmt.Print(time.Since(start))
}

// Función para imprimir el heap
func (m *maxheap) print() {
	for _, val := range m.arr {
		fmt.Println(val)
	}
}

// Función que se encarga del proceso
func heapsort(array []int) {
	start := time.Now()
	fmt.Println("Heapsort \n")
	minHeap := newMaxHeap(array)
	minHeap.sort(len(array), start)

	// minHeap.print()
}

// Función encargada de ejecutar el insertionsort
func InsertionSort(wg *sync.WaitGroup, randomArray []int, insertionChannel chan []int) {
	var tiempoIni = time.Now()
	var tiempoFinal = time.Now()
	var num1 = 1
	var numAux = 0 //declaracion de variables
	var guardado = 0
	var ciclos = 0
	var condicionales = 0
	var intercambios = 0
	for num1 < len(randomArray) { //recorrida iteratiba sobre el array
		ciclos++
		posiciones := []int{}
		posiciones = append(posiciones, num1)
		numAux = num1
		guardado = randomArray[num1]
		var cambio = false
		for numAux > 0 { //iteracion para devolver un indice hasta donde sea necesario en el array
			ciclos++
			if guardado < randomArray[numAux-1] {
				condicionales++
				randomArray[numAux] = randomArray[numAux-1] //si el numero anterior es mayor inicia cambio de posiciones
				cambio = true                               //condicional usado para saber si el numero debe seguir siendo movido o nunca sera necesario
				intercambios++
			} else if cambio {
				condicionales++
				intercambios++
				randomArray[numAux] = guardado          // si en algun momento hubo algun cambio pero ya no es necesario
				posiciones = append(posiciones, numAux) // significa que es el lugar dode el numero debe permanecer
				posiciones = append(posiciones, ciclos)
				posiciones = append(posiciones, condicionales)
				posiciones = append(posiciones, intercambios)
				insertionChannel <- posiciones
				<-insertionChannel
				break
			} else {
				condicionales++
				numAux = -1 //de no haber habido un cambio nunca se cancela el ciclo para el numero actual
			}
			if numAux-1 == 0 {
				condicionales++
				intercambios++
				randomArray[numAux-1] = guardado
				posiciones = append(posiciones, numAux-1)
				posiciones = append(posiciones, ciclos)
				posiciones = append(posiciones, condicionales)
				posiciones = append(posiciones, intercambios)
				insertionChannel <- posiciones
				<-insertionChannel
			}
			numAux--
		}
		num1++
	}
	posiciones := []int{}
	tiempoFinal = time.Now()
	posiciones = append(posiciones, 0)
	posiciones = append(posiciones, 0)
	posiciones = append(posiciones, ciclos)
	posiciones = append(posiciones, condicionales)
	posiciones = append(posiciones, intercambios)
	print(tiempoIni.Format("2006-01-02 15:04:05"))
	print(tiempoFinal.Format("2006-01-02 15:04:05"))
	insertionChannel <- posiciones
	<-insertionChannel
	defer wg.Done()
	//arrChan <- randomArray
}

// Función encargade de realizar el quicksort
func QuickSort(wg *sync.WaitGroup, array []int, quickChannel chan []int, ultimoIndice int, wgEnd int, ciclos int, condicionales int, intercambios int) { //metodo recursivo con pivote al final
	var tiempoIni = time.Now()
	print(tiempoIni.Format("2006-01-02 15:04:05"))
	if len(array) < 2 {
		condicionales++
		//fmt.Println("Fin")		//condicion de salida
		return
	}
	wgEnd = wgEnd + 1
	//fmt.Println("Entra")
	ladoIzq := 0
	ladoDer := len(array) - 1 //declaracion de variables
	auxiliar := 0
	posiciones := []int{}

	for i, _ := range array {
		ciclos++
		if array[i] < array[ladoDer] { //Si el numero en el indice actual en menor al pivote
			condicionales++
			auxiliar = array[ladoIzq] //se mueve el numero del indice al final del subarray
			array[ladoIzq] = array[i] //de menores a la izquierda, y se actualiza la variable
			array[i] = auxiliar       //representando el final del mismo
			if (i + ultimoIndice) != (ladoIzq + ultimoIndice) {
				intercambios++
				posiciones = append(posiciones, i+ultimoIndice)
				posiciones = append(posiciones, ladoIzq+ultimoIndice)
				posiciones = append(posiciones, ciclos)
				posiciones = append(posiciones, condicionales)
				posiciones = append(posiciones, intercambios)
				quickChannel <- posiciones
				<-quickChannel
				time.Sleep(50 * time.Millisecond)
				posiciones = nil
			}
			ladoIzq++
		}
	}

	auxiliar = array[ladoIzq]
	array[ladoIzq] = array[ladoDer] //El pivote se mueve al final del subarray de la izquierda
	array[ladoDer] = auxiliar
	intercambios++
	posiciones = append(posiciones, ladoIzq+ultimoIndice)
	posiciones = append(posiciones, ladoDer+ultimoIndice)
	posiciones = append(posiciones, ciclos)
	posiciones = append(posiciones, condicionales)
	posiciones = append(posiciones, intercambios)
	quickChannel <- posiciones
	<-quickChannel
	time.Sleep(50 * time.Millisecond)
	posiciones = nil
	QuickSort(wg, array[:ladoIzq], quickChannel, ultimoIndice, wgEnd, ciclos, condicionales, intercambios) //Llamadas recursivas para ambos subarrays
	QuickSort(wg, array[ladoIzq+1:], quickChannel, ultimoIndice+ladoIzq+1, wgEnd, ciclos, condicionales, intercambios)
	if wgEnd == 1 {
		condicionales++
		var tiempoFinal = time.Now()
		print(tiempoFinal.Format("2006-01-02 15:04:05"))
		defer wg.Done()
	}
}

func TempGraficarBubble(bubbleChannel chan []int) {
	posicCamb := []int{}
	for true {
		posicCamb = <-bubbleChannel
		if posicCamb[0] == 0 && posicCamb[1] == 0 {
			bubbleChannel <- posicCamb
			break
		}
		fmt.Print(posicCamb[0])
		fmt.Print(" ")
		fmt.Print(posicCamb[1]) //Prueba, en el futuro va a ser el que obtiene para graficar
		fmt.Println(" " + "BubbleSort")
		bubbleChannel <- posicCamb
	}
}

func TempGraficarQuick(quickChannel chan []int) {
	posicCamb := []int{}
	for true {
		posicCamb = <-quickChannel
		// if posicCamb[0] == 0 && posicCamb[1] == 0 {
		// 	quickChannel <- posicCamb
		// 	break
		// }
		fmt.Print(posicCamb[0])
		fmt.Print(" ")
		fmt.Print(posicCamb[1]) //Prueba, en el futuro va a ser el que obtiene para graficar
		fmt.Print(" ")
		fmt.Print(posicCamb[2])
		fmt.Print(" ")
		fmt.Print(posicCamb[3])
		fmt.Print(" ")
		fmt.Print(posicCamb[4])
		fmt.Println(" " + "QuickSort")
		quickChannel <- posicCamb
	}
}

func TempGraficarInsertion(insertionChannel chan []int) {
	posicCamb := []int{}
	for true {
		posicCamb = <-insertionChannel
		if posicCamb[0] == 0 && posicCamb[1] == 0 {
			fmt.Print(posicCamb[2])
			fmt.Print(" ")
			fmt.Print(posicCamb[3])
			fmt.Print(" ")
			fmt.Print(posicCamb[4])
			insertionChannel <- posicCamb
			break
		}
		fmt.Print(posicCamb[0])
		fmt.Print(" ")
		fmt.Print(posicCamb[1]) //Prueba, en el futuro va a ser el que obtiene para graficar
		fmt.Print(" ")
		fmt.Print(posicCamb[2])
		fmt.Print(" ")
		fmt.Print(posicCamb[3])
		fmt.Print(" ")
		fmt.Print(posicCamb[4])
		fmt.Println(" " + "InsertionSort")
		insertionChannel <- posicCamb
	}
}

// Genera una copia de un array
func CopyArray(arrOri []int) []int {
	nuevoArray := []int{}
	var num = 0
	for i := 0; i < len(arrOri); i++ { //copiado de Array
		num = arrOri[i]
		nuevoArray = append(nuevoArray, num)
	}
	return nuevoArray
}

func main() {
	var size int
	for true {
		fmt.Println("Ingrese cuantos números desea [10,10000]")
		fmt.Scan(&size)
		if size >= 10 && size <= 10000 {
			break
		} else {
			fmt.Println("El número debe estar en el intervalo de [10,10000]")
		}
	}
	randomch := make(chan int, 1) //creacion de canales para fturo uso
	arrayChannel := make(chan []int, 1)
	bubbleChannel := make(chan []int, 1)
	insertionChannel := make(chan []int, 1)
	// //heapChannel := make(chan []int, 1)
	quickChannel := make(chan []int, 1)

	arr := []int{}
	arr2 := []int{}
	arr3 := []int{}
	arr4 := []int{}

	var waitGroup sync.WaitGroup
	go generateRandom(&waitGroup, randomch, arrayChannel, size) //genera el array

	fmt.Println("Unsorted Arrays")
	arr = <-arrayChannel //saca el array del channel
	fmt.Println(arr)
	arr2 = CopyArray(arr) //copia el array
	fmt.Println(arr2)
	arr3 = CopyArray(arr2) //copia el array
	fmt.Println(arr3)
	arr4 = CopyArray(arr3) //copia el array
	fmt.Println(arr4)

	waitGroup.Add(3)
	go TempGraficarBubble(bubbleChannel)
	go BubbleSort(&waitGroup, arr, bubbleChannel) //BubbleSort al primer Array
	go TempGraficarInsertion(insertionChannel)
	go InsertionSort(&waitGroup, arr2, insertionChannel) //InsetionSort al segundo Array
	go TempGraficarQuick(quickChannel)
	go QuickSort(&waitGroup, arr3, quickChannel, 0, 0, 0, 0, 0) //InsetionSort al segundo Array
	waitGroup.Wait()
	//heapsort(arr3)
	//fmt.Println("QuickSort")
	//arr5 := []int{2, 5, 12, 16, 20, 27, 28, 29, 29, 31}
	//QuickSort2(arr3, 0) //Quicksort sin corrutinas al tercer Array

	fmt.Println("\nSorted Arrays")
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println("Terminó") //verificacion
}
