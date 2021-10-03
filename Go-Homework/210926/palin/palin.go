package palin

func Palin(n string) string {
	if n != "0" {
		l := len(n)
		for x := 0; x < l; x++ {
			if n[x] != n[l-1-x] {
				return "no"
			}
		}
		return "yes"
	}
	return ""
}

// func Palin() {
// 	var n string
// 	fmt.Scanln(&n)
// 	if n != "0" {
// 		l := len(n)
// 		for x := 0; x < l; x++ {
// 			if n[x] != n[l-1-x] {
// 				fmt.Println("no")
// 				Palin()
// 				return
// 			}
// 		}
// 		fmt.Println("yes")
// 		Palin()
// 		return
// 	}
// 	return
// }

// func main() {
// 	Palin()
// }
