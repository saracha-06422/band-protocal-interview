package problem

import "sort"

func MaxChickensCovered(positions []int, n, k int) int {
	//ถ้า roof ไม่มีความยาว
	if k == 0 {
		return 1
	}
	var result int
	var chickensCovers []int
	var coverDistance int
	for i := 0; i < len(positions); i++ {
		coverDistance = positions[i] + k - 1 //-1 เพราะต้องเอาจุดเริ่มต้นด้วย

		for j := i; j < len(positions); j++ {
			if positions[j] > coverDistance { //j = 2
				chickensCovers = append(chickensCovers, j-i)
				break
			}

			//ตัวสุดท้ายแล้วแต่ coverDistance ยังมากกว่าหรือเท่ากับ
			if j == len(positions)-1 { //ตัวสุดท้าย
				chickensCovers = append(chickensCovers, (j-i)+1) //+1 เพราะต้องเอาตัวสุดท้ายด้วย
			}
		}
	}

	//หาจำนวนไก่ที่มามากที่สุดที่ช่วยได้
	sort.Ints(chickensCovers)
	result = chickensCovers[len(chickensCovers)-1]

	return result
}
