package slice

import "fmt"

//parçalamak - parça
//önemliiii
// yeni bir veri ekledigimizde dizinin ..
//yeniden boyutlandırılmasının ortadan kalkmasıdır
func Demo1() {
	isimler := make([]string, 3)
	//makele multi array gibi oldu

	isimler[0] = "hç"
	isimler[1] = "huri"
	isimler[2] = "berke"
	isimler = append(isimler, "larkin")
	//append(eklemek) yepyeni bir slice oluşturmuş oluyor
	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
