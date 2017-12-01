package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

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

	fmt.Println("I am happy.")
	fmt.Println(ElizaResponce("I am happy."))
	fmt.Println()

	fmt.Println("I am not happy with your responses.")
	fmt.Println(ElizaResponce("I am not happy with your responses."))
	fmt.Println()

	fmt.Println("I am not sure that you understand the effect that your questions are having on me.")
	fmt.Println(ElizaResponce("I am not sure that you understand the effect that your questions are having on me."))
	fmt.Println()

	fmt.Println("I am supposed to just take what you’re saying at face value?")
	fmt.Println(ElizaResponce("I am supposed to just take what you’re saying at face value?"))
	fmt.Println()

}

func ElizaResponce(input string) string {

	if matched, _ := regexp.MatchString(`(?i).*\bfather\b.*`, input); matched {
		return "Why don’t you tell me more about your father?"
	}

	re := regexp.MustCompile(`(?i).*\bi am|i'm|im\b.*([^.?!]*)(?i)[.?!]?`)
	if matched := re.MatchString(input); matched {
		return re.ReplaceAllString(input, "How do know you are $1?")
	} /*FindStringSubMatch*/

	responce := [3]string{"I’m not sure what you’re trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?"}

	return responce[rand.Intn(len(responce))]

}

func Reflect(input string) string {
	boundries := regexp.MustCompile(`\b`)
	words := boundries.Split(input, -1)

	reflectedWords := [][]string{
		{`i`, `you`},
		{`you`, `i`},
		{`were`, `was`},
		{`am`, `are`},
		{`are`, `am`},
		{`was`, `were`},
		{`i'm`, `you are`},
		{`you're`, `I am`},
		{`you`, `me`},
		{`my`, `your`},
		{`me`, `you`},
		{`you've`, `I have`},
		{`you'll`, `I will`},
		{`your`, `my`},
		{`yours`, `mine`},
		{`myself`, `yourself`},
		{`any`, `some`},
		{`some`, `any`},
	}

	//Loop through each word and if there is a match do a reflection
	for i, word := range words {
		for _, reflectedWords := range reflectedWords {
			if matched, _ := regexp.MatchString(reflectedWords[0], word); matched {
				words[i] = reflectedWords[1]
				break
			}
		}
	}
	return strings.Join(words, ``)
}
