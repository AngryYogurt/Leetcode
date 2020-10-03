package main

import "fmt"

func main() {
	fmt.Println()
}

type ParkingSystem struct {
	b, m, s int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		b: big,
		m: medium,
		s: small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.b > 0 {
			this.b--
			return true
		}
		return false
	case 2:
		if this.m > 0 {
			this.m--
			return true
		}
		return false
	case 3:
		if this.s > 0 {
			this.s--
			return true
		}
		return false
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
