package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
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

// https://www.reddit.com/r/golang/comments/18n78cd/text_to_speech_go_package/

func splitIntoSentences(text string) []string {
	var sentences []string
	var currentSentence string

	for _, word := range strings.Fields(text) {
		nextSentence := currentSentence + " " + word

		if len(nextSentence) > 300 {
			sentences = append(sentences, currentSentence)
			currentSentence = word
		} else {
			currentSentence = nextSentence
		}
	}

	sentences = append(sentences, currentSentence)
	return sentences
}

// func play(data []byte) error {

// 	fileBytesReader := bytes.NewReader(data)

// 	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
// 	if err != nil {
// 		return err
// 	}

// 	otoCtx, readyChan, err := oto.NewContext(decodedMp3.SampleRate(), 2, 2)
// 	if err != nil {
// 		return err
// 	}
// 	<-readyChan

// 	player := otoCtx.NewPlayer(decodedMp3)

// 	player.Play()

// 	for player.IsPlaying() {
// 		time.Sleep(time.Millisecond)
// 	}

// 	return player.Close()
// }

func sendRequest(text string) ([]byte, error) {

	url := fmt.Sprintf(
		"http://translate.google.com/translate_tts?client=tw-ob&q=%s&tl=%s",
		url.QueryEscape(text),
		voices.Russian,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err

	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("status code error: " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func translate(text, toLang string) (string, error) {
	params := url.Values{}
	params.Add("client", "gtx")
	params.Add("sl", "auto")
	params.Add("tl", toLang)
	params.Add("q", text)

	url := "https://translate.googleapis.com/translate_a/t?" + params.Encode()

	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var result [][]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	if len(result) > 0 && len(result[0]) > 0 {
		if translatedText, ok := result[0][0].(string); ok {
			return translatedText, nil
		}
	}

	return "", fmt.Errorf("unexpected response format")
}

func setupLogging() *os.File {
	file, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Failed to open log file: ", err)
	}

	mv := io.MultiWriter(file, os.Stdout)
	log.SetOutput(mv)

	return file
}

func playWav(data io.Reader) error {
	streamer, format, err := wav.Decode(data)
	if err != nil {
		return fmt.Errorf("playing audio: %s", err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
	return nil
}

func main() {

	logFile := setupLogging()
	defer logFile.Close()

	// text := "Привет мир!"

	// text := "Здравствуйте, это доставка с Ali-Express! Мы пытались доставить вам ваши посылки «Резиновый дилдо 25см, розовый», «Костюм Фурри лисицы XXL, детский», «Книга позы для зоо порно» но уже 5-й раз не застали никого дома. Так же мы звонили вам на мобильный, но все тщетно! Вынуждены побеспокоить вас в вашем стим профиле. Пожалуйста, позвоните нам на номер +375 (25) 228 1337. Мы заботимся о наших клиентах!"

	// for _, sentence := range splitIntoSentences(text) {

	// 	fmt.Println("Req")

	// 	body, err := sendRequest(sentence)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	if err := play(body); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// куник https://discord.com/channels/600695204965646346/699042990785822746/1250866684949168161
	// ask to play

	// path, err := filepath.Abs("./piperInstated")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for {

		url := voices.ChooseVoice()
		if url == "" {
			return
		}

		tts, err := piper.New("piperInstated", "cache", url)
		if err != nil {
			log.Fatal(err)
		}

		parts := strings.Split(url, "/")
		toLnag := parts[7]

		for {
			fmt.Println("Enter text (press 0 to go back):")
			reader := bufio.NewReader(os.Stdin)

			text, err := reader.ReadString('\n')
			if err != nil {
				log.Println(err)
				fmt.Scanln()
				break
			}
			text = strings.TrimSpace(text)

			if text == "0" {
				break
			}

			if res, err := translate(text, toLnag); err == nil {
				text = res
			}

			// fmt.Println(text, "\n", toLnag)

			wavBytes, err := tts.Synthesize(text)
			if err != nil {
				log.Println(err)
				fmt.Scanln()
				break
			}

			go func() {
				if err := os.WriteFile("bub.wav", wavBytes, 0644); err != nil {
					log.Println(err)
				}
			}()

			if err := playWav(bytes.NewReader(wavBytes)); err != nil {
				log.Println(err)
				fmt.Scanln()
				break
			}
		}

	}
}
