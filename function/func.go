package function

//하나만 삽입
func Insert(slice []int, index int, value int) []int {
	slice = append(slice, 0) //slice에 공간 하나 추가
	// slice = append(slice[:index],append([]int{value},slice[index:]...)...) 원본 슬라이스가 가리키는 배열 외에 다른 슬라이스의 임시 배열을 만들기 때문에 그닥 효율적이지 않음
	copy(slice[index+1:], slice[index:])
	slice[index] = value

	return slice
}

//여러개 삽입하기
func Insert2(slice []int, index int, s ...int) []int {

	if cap(slice)-len(slice) < len(s) { //삽입 할 공간이 없을 경우
		tmp := make([]int, 0, cap(slice)+len(s))
		tmp = append(tmp, slice[:index]...)
		tmp = append(tmp, s...)
		tmp = append(tmp, slice[index:]...)
		return tmp
	} else { //삽입할 공간이 충분 할 경우
		copy(slice[index+len(s):], slice[index:])
		return append(slice[:index], s...)
	}

}

func Pop(slice []int) (int, []int) {
	pop := slice[len(slice)-1]
	slice = slice[:len(slice)-1]
	return pop, slice
}

func Remove(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func Copy(slice []int) []int {
	s := make([]int, len(slice))
	copy(s, slice)
	return s
}
