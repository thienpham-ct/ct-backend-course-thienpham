package main

import "fmt"

type ParkingSystem struct {
	Big    int
	Medium int
	Small  int
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	a := ParkingSystem{
		Big:    big,
		Medium: medium,
		Small:  small,
	}
	return a
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		added := this.big < this.Big
		this.big += 1
		return added
	case 2:
		added := this.medium < this.Medium
		this.medium += 1
		return added
	case 3:
		added := this.small < this.Small
		this.small += 1
		return added
	}
	return false
}

func main() {
	carpark1 := Constructor(1, 1, 0)
	ans := make([]bool, 2)

	ans[0] = carpark1.AddCar(1)
	ans[1] = carpark1.AddCar(1)
	fmt.Println(ans)

	//func (this *ParkingSystem) AddCar(carType int) bool {
	//
	//}

	/**
	 * Your ParkingSystem object will be instantiated and called as such:
	 * obj := Constructor(big, medium, small);
	 * param_1 := obj.AddCar(carType);
	 */
}
