package main

import (
	"bufio"
	"fmt"
	"github.com/VMETA3/sui-hackathon/common/gpt"
	"github.com/VMETA3/sui-hackathon/common/openai"
	oa "github.com/VMETA3/sui-hackathon/utils/openai"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	gpt := gpt.InitGPT()
	fmt.Println("What is your name?")
	inputReader := bufio.NewReader(os.Stdin)
	name, err := inputReader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	Completion(gpt, inputReader, name)
	// Edits(TCG, inputReader, name)
}

func Completion(chat *gpt.GPT, inputReader *bufio.Reader, name string) {
	for {
		fmt.Printf("Please enter > ")
		p, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		resp, err := chat.Continuous("Simulate", p)
		// resp, err := chat.Single(p)

		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(resp)
	}
}

func Completion3_5(chat []*oa.TurboMessage) {
	inputReader := bufio.NewReader(os.Stdin)
	OA := openai.Init()
	for {
		fmt.Printf("Please enter > ")
		p, err := inputReader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		chat = append(chat, &oa.TurboMessage{
			Role:    oa.USER,
			Content: p,
		})
		resp, err := OA.Client.Turbo().Chat(chat)

		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(resp.Content)

		chat = append(chat, &oa.TurboMessage{
			Role:    oa.ASSIS,
			Content: resp.Content,
		})
	}
}
