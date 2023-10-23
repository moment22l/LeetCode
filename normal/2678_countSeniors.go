package main

func countSeniors(details []string) (ans int) {
	for _, detail := range details {
		age := int(detail[11]-'0')*10 + int(detail[12]-'0')
		if age > 60 {
			ans++
		}
	}
	return
}
