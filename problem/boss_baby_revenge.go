package problem

func BossBabyRevenge(s string) string {
	//boss ยิงเด็กก่อน
	var result string = "Bad boy"
	if s[0] == 'R' {
		return result
	}

	//จะเหลือแค่ที่เป็น S นำ (เด็กยิงก่อน)
	var kidShootflag, shootNum = false, 0

	for _, v := range s {
		if v == 'S' {
			kidShootflag = true
			shootNum++
		}
		//revege
		if v == 'R' {
			kidShootflag = false
			//แก้แค่เกินไม่นับ
			if shootNum > 0 {
				shootNum--
			}
		}
	}

	if !kidShootflag && shootNum == 0 {
		//แก้แค้ครบแล้ว 555
		result = "Good boy"
	}

	return result
}
