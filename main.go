package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go/option"

	"github.com/openai/openai-go"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("😡:", err)
	}

	// Ollama
	chatURL := os.Getenv("OLLAMA_BASE_URL") + "/v1/"
	model := os.Getenv("LLM_CHAT")

	apiKey := "i-love-parakeets"

	client := openai.NewClient(
		option.WithBaseURL(chatURL),
		option.WithAPIKey(apiKey),
	)

	ctx := context.Background()

	var messages []openai.ChatCompletionMessageParamUnion

	systemInstruction := openai.SystemMessage("You are a useful AI assistant, you name is Seven of Nine")
	userQuestion := openai.UserMessage("Who is Jean-Luc Picard?")

	messages = append(messages, systemInstruction, userQuestion)

	param := openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    model,
	}

	stream := client.Chat.Completions.NewStreaming(ctx, param)

	fmt.Println("🔴 starting the completion...")
	fmt.Println("🟠 starting the completion...")
	fmt.Println("🟢 starting the completion...")

	for stream.Next() {
		chunk := stream.Current()
		// Stream each chunk as it arrives
		if len(chunk.Choices) > 0 && chunk.Choices[0].Delta.Content != "" {
			fmt.Print(chunk.Choices[0].Delta.Content)
		}
	}

	if err := stream.Err(); err != nil {
		fmt.Printf("Error in stream: %v\n", err)
		return
	}

	fmt.Println("\n🔵 done!")

}
