package main

//偶数奇数の判定
//func main() {
//	for i := 1; i <= 100; i++ {
//		if i%2 == 0 {
//			message := "偶数"
//			fmt.Printf("%d-%s\n", i, message)
//		} else {
//			message := "奇数"
//			fmt.Printf("%d-%s\n", i, message)
//		}
//	}
//}

//func main() {
//	t := time.Now().UnixNano()
//	rand.Seed(t)
//	s := rand.Intn(5) + 1
//
//	if s >= 6 {
//		message := "大吉"
//		fmt.Printf("%d-%s\n", s, message)
//	} else if s >= 4 {
//		message := "中吉"
//		fmt.Printf("%d-%s\n", s, message)
//	} else if s >= 2 {
//		message := "小吉"
//		fmt.Printf("%d-%s\n", s, message)
//	} else {
//		message := "凶"
//		fmt.Printf("%d-%s\n", s, message)
//	}
//}

// 配列の初期化について

//func array() {
//	// ゼロ値で初期化
//	//var ns1 = [5]int
//
//	// 配列リテラルで初期化
//	//var ns2 = [5]int{10,20,30,40,50}
//
//	// 要素数を値から推論
//	//ns3 := [...]int{10, 20, 30, 40, 50}
//}
//
//func main() {
//	ns := [...]int{10, 20, 30, 40}
//	println(len(ns))
//	println(ns[3])
//	println(ns[1:2])
//}

//func main() {
//	//マップ
//
//	// ゼロ値はnil
//	//var m map[string]int
//
//	// makeで初期化
//	//m = make(map[string]int)
//
//	//リテラルで初期化
//	//m := map[string]int{"x":10, "y":20}
//}

//func main() {
//	m := map[string]int{"x": 20, "y": 50}
//
//	println(m["x"])
//
//	m["z"] = 30
//
//	n, ok := m["z"]
//	println(n, ok)
//
//	delete(m, "z")
//
//	n, ok = m["z"]
//	println(n, ok)
//}

//type User struct {
//}
//
//type Score struct {
//	UserID string
//	GameID int
//	Point int
//}

// 型のエイリアスを定義することができる
//type Applicant = http.Client
//
//func main() {
//	fmt.Printf("%T", Applicant{})
//}

//func add(x int, y int) int {
//	return x + y
//}
//
//func swap(x int, y int) (int, int) {
//	return y, x
//}

//関数の返り値を受け取りたくない時には _ ブランク変数を用いる

//名前付き戻り値

//func swap(x, y int) (x2, y2 int) {
//	y2, x2 = x, y
//	return
//}

// 無名関数

//func main() {
//	message := "Hello world"
//
//	func() {
//		// 関数外のmessageを参照できている
//		println(message)
//	}()
//}
