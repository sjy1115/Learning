package main

import (
	qr2 "learning/pkg/qr"
)

func main() {
	qr()

	//config.Init("")
	//
	//token, err := jwt.GenerateToken(1, 1, "xiaoming", "admin")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//userToken, err := jwt.ParseToken(token)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6InhpYW9taW5nIiwicGFzc3dvcmQiOiJhZG1pbiIsInJvbGUiOjEsImV4cCI6MTY1MDQ2NjgzNiwiaXNzIjoiZ2luLWxlYXJuaW5nIn0.Dn-i5THau0U1DNzMqPOCKKp8Q6NGq57yFMlWqHSMo_U
	//// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6InhpYW9taW5nIiwicGFzc3dvcmQiOiJhZG1pbiIsInJvbGUiOjEsImV4cCI6MTY1MDQ2NjY0NywiaXNzIjoiZ2luLWxlYXJuaW5nIn0.eewoPGqc6F0N_IwK5SVjomNfr_SoQBiInUuq4EuwKWE
	//
	//fmt.Println(token)
	//fmt.Println(userToken)
}

func qr() {
	if _, err := qr2.Generate(""); err != nil {
		panic(err)
	}
}
