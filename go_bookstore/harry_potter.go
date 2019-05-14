package harry_potter

func Answer(n ...float64) float64 {

	n = append([]float64{0, 0, 0, 0, 0}, n...)
	n = n[len(n)-5:]

	return 8*5*n[0]*0.75 + 8*4*(n[1]-n[0])*0.8 + 8*3*(n[2]-n[1])*0.9 + 8*2*(n[3]-n[2])*0.95 + 8*1*(n[4]-n[3])*1
}
