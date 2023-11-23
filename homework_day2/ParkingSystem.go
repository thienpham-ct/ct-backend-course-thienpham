package homework_day2

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
