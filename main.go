package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"govorilka/piper"
	"govorilka/voices"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

func translate(text, toLang string) (string, string, error) {
	params := url.Values{
		"client": {"gtx"},
		"sl":     {"auto"},
		"tl":     {toLang},
		"q":      {text},
	}

	url := "https://translate.googleapis.com/translate_a/t?" + params.Encode()

	res, err := http.Get(url)
	if err != nil {
		return "", "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", fmt.Errorf("reading response body: %w", err)
	}

	var result [][]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", "", fmt.Errorf("JSON unmarshaling failed: %w", err)
	}

	if len(result) == 0 || len(result[0]) < 2 {
		return "", "", fmt.Errorf("unexpected response format")
	}

	translatedText, ok1 := result[0][0].(string)
	detectedLang, ok2 := result[0][1].(string)

	if !ok1 || !ok2 {
		return "", "", fmt.Errorf("unexpected data types in response")
	}

	return translatedText, detectedLang, nil
}

func playWav(data io.Reader) error {
	streamer, format, err := wav.Decode(data)
	if err != nil {
		return fmt.Errorf("playing audio: %s", err)
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return err
	}

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
	return nil
}

func setupLogging() (*os.File, error) {
	logFile, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}
	log.SetOutput(io.MultiWriter(logFile, os.Stdout))
	return logFile, nil
}

func main() {
	logFile, err := setupLogging()
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()

	for {
		url := voices.ChooseVoice()
		if url == "" {
			return
		}

		tts, err := piper.New("piperInstated", "downloadedVoices", url)
		if err != nil {
			log.Println(err)
			continue
		}

		parts := strings.Split(url, "/")
		toLang := parts[7]

		if err := processUserInput(tts, toLang); err != nil {
			log.Printf("Error processing user input: %v", err)
		}
	}
}

func processUserInput(tts *piper.TTS, toLang string) error {
	reader := bufio.NewReader(os.Stdin)
	previousLang := ""

	for {
		fmt.Println("Enter text (press 0 to go back):")
		text, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("reading input: %w", err)
		}
		text = strings.TrimSpace(text)

		if text == "0" {
			return nil
		}

		if previousLang != toLang {
			translatedText, detectedLang, err := translate(text, toLang)
			if err != nil {
				return fmt.Errorf("translate: %w", err)
			}
			text = translatedText
			previousLang = detectedLang
			fmt.Println("translate")
		}

		wavBytes, err := tts.Synthesize(text)
		if err != nil {
			return fmt.Errorf("synthesizing text: %w", err)
		}

		go func() {
			if err := os.WriteFile("synthesized.wav", wavBytes, 0644); err != nil {
				log.Println(err)
			}
		}()

		if err := playWav(bytes.NewReader(wavBytes)); err != nil {
			return fmt.Errorf("playing WAV: %w", err)
		}
	}
}
