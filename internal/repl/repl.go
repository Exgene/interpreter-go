package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/exgene/forge/internal/tokenizer"
)

func ReadValue() {
	fmt.Println("Running the interactive terminal. Press Ctrl+C to exit.")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	inputChan := make(chan string)
	errChan := make(chan error, 1) 

	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("> ")
			text, err := reader.ReadString('\n')
			if err != nil {
				errChan <- err
				close(inputChan)
				return
			}
			inputChan <- text
		}
	}()

	for {
		select {
		case sig := <-sigChan:
			fmt.Printf("\nReceived signal: %v. Exiting.\n", sig)
			fmt.Println("Finished Running or interrupted")
			return

		case text, ok := <-inputChan:
			if !ok {
				select {
				case err := <-errChan:
					if err != io.EOF {
						fmt.Fprintf(os.Stderr, "\nInput error: %v\n", err)
					} else {
						fmt.Println("\nInput closed (EOF).")
					}
				default:
					fmt.Println("\nInput closed.")
				}
				fmt.Println("Finished Running or interrupted")
				return
			}
			processedText := strings.TrimSpace(text)
			if processedText == "" {
				continue
			}
			tokenizerObj := tokenizer.NewTokenizer(processedText)
			tokens := tokenizerObj.ScanTokens()
			fmt.Printf("Tokens: %v\n", tokens)

		case err := <-errChan:
			if err == io.EOF {
				fmt.Println("\nInput stream closed (EOF).")
			} else {
				fmt.Fprintf(os.Stderr, "\nError reading input: %v\n", err)
			}
			fmt.Println("Finished Running or interrupted")
			return
		}
	}
}
