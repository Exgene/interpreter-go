package repl

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/exgene/forge/internal/core"
	// "github.com/exgene/forge/internal/tokenizer"
)

func ReadValue() {
	fmt.Println("Running the interactive terminal. Press Ctrl+C to exit.")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	inputChan := make(chan string)
	done := false

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			if !done {
				continue
			}
			fmt.Printf("\n> ")
			text, err := reader.ReadString('\n')
			if err != nil {
				fmt.Printf("Error while reading from stdin: %v", err)
			}
			inputChan <- text
		}
	}()

	done = true

	for {
		select {
		case sig := <-sigChan:
			fmt.Printf("\nReceived signal: %v. Exiting.\n", sig)
			fmt.Println("Finished Running or interrupted")
			return
		case text := <-inputChan:
			done = false
			processedText := strings.TrimSpace(text)
			if processedText == "" {
				continue
			}
			parser := core.NewParser()
			astNodes := parser.ProduceAST(processedText)
			core.PrintNode(astNodes, "")
			fmt.Printf("AST Nodes: %v", astNodes)
			// tokenizerObj := tokenizer.NewTokenizer(processedText)
			// tokens := tokenizerObj.ScanTokens()
			// fmt.Printf("Tokens: %v\n", tokens)
			done = true
		default:
		}
	}
}
