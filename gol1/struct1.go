package main

import "fmt"

type car struct {
	gas_pedal unint16 //min 0 max 65535
	brake_pedal unint16
	steering_wheel int16 //  -32k - +32k
	top_speed_kmh float64
}

func main() {
	a_car := car{gas_pedal: 22341, 
	brake_pedal: 0, 
	steering_wheel: 12561, 
	top_speed_kmh: 225.0}
	
	fmt.Println(a_car.gas_pedal)




}
