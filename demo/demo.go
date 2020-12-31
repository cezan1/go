package main

func sum(Arr [10]float32) float32 {
	sum := float32(0.0)
	for _, val := range Arr {
		//累计求和
		sum += val
	}
	return sum
}
