package main

import (
	"fmt"
)

func main() {
	// 맵, 래퍼런스 타입
	// var 맵 map[키_자료형]값_자료형

	/*	var a map[string]int = make(map[string]int)
		var b = make(map[string]int)
		c := make(map[string]int)

		fmt.Println(a, b, c)
	*/

	a := map[string]int{"hello": 10, "world": 20}
	fmt.Println(a)

	solarSystem := make(map[string]float32) // 키는 string, 값은 float32인 맵 생성 및 공간 할당

	solarSystem["Mercury"] = 87.969 // 맵[키] = 값
	solarSystem["Venus"] = 224.70069
	solarSystem["Earth"] = 365.25641
	solarSystem["Mars"] = 686.9600
	solarSystem["Jupiter"] = 4333.2867
	solarSystem["Saturn"] = 10756.1995
	solarSystem["Uranus"] = 30707.4896
	solarSystem["Neptune"] = 60223.3528

	fmt.Println(solarSystem["Earth"]) // 365.25641

	fmt.Println(solarSystem["Pluto"]) // 0

	value, ok := solarSystem["Pluto"] // 0 키값 , false 값의 유뮤
	fmt.Println(value, ok)

	if value, ok := solarSystem["Pluto"]; ok {
		fmt.Println(value, ok)
	}

	for index, value := range solarSystem {
		fmt.Println(index, value)
	}

	for _, value := range solarSystem {
		fmt.Println(value)
	}

	delete(a, "world")

	terrestrialPlanet := map[string]map[string]float32{
		"Mercury": map[string]float32{
			"meanRadius":    2439.7,
			"mass":          3.3022e+23,
			"orbitalPeriod": 87.969,
		},
		"Venus": map[string]float32{
			"meanRadius":    6051.8,
			"mass":          4.8676e+24,
			"orbitalPeriod": 224.70069,
		},
		"Earth": map[string]float32{
			"meanRadius":    6371.0,
			"mass":          5.97219e+24,
			"orbitalPeriod": 365.25641,
		},
		"Mars": map[string]float32{
			"meanRadius":    3389.5,
			"mass":          6.4185e+23,
			"orbitalPeriod": 686.9600,
		},
	}

	fmt.Println(terrestrialPlanet["Mars"]["mass"]) // 6.4185E+23
}
