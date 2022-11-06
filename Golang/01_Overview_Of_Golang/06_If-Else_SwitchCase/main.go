package main

import "log"

func main() {
	myNum := 120
	isTrue := false

	if myNum == 100 && isTrue {
		log.Println("myNum is 100 and isTrue is True")
	} else if myNum > 100 && !isTrue {
		log.Println("Mynum is greater 100 and isTrue is False")
	} else {
		log.Println("myNum is not equal to 100 and isTrue is False")
	}

	// ------------------------------------- Switch Case -------------------------------------------------------

	myVar := "picture1"

	switch myVar {
	case "picture1":
		log.Println("Picture is set to JPEG")
	case "picture2":
		log.Println("Picture is set to PNG")

	default:
		log.Println("Invalid Picture")
	}
}
