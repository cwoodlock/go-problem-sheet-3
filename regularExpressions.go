package main

import "fmt"
import "math/rand"
import "time"

func main() {
	fmt.Println("People say I look like both my mother and father.")
	fmt.Println(ElizaResponce("People say I look like both my mother and father."))
	fmt.Println()

	fmt.Println("Father was a teacher.")
	fmt.Println(ElizaResponce("Father was a teacher."))
	fmt.Println()

	fmt.Println("I was my father’s favourite.")
	fmt.Println(ElizaResponce("I was my father’s favourite."))
	fmt.Println()

	fmt.Println("I’m looking forward to the weekend.")
	fmt.Println(ElizaResponce("I’m looking forward to the weekend."))
	fmt.Println()

	fmt.Println("My grandfather was French!")
	fmt.Println(ElizaResponce("My grandfather was French!"))
	fmt.Println()

}

func ElizaResponce(input string) string {
	rand.Seed(time.Now().UTC().UnixNano())
	responce := [3]string{"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?"}

	return responce[rand.Intn(len(responce))]

}
