
# GoVoice - Speech Synthesis on Go

GoVoice converts the text into synthesized human speech. Uses [Piper](https://github.com/rhasspy/piper) to generate speech.

## Features

- Support for multiple languages for speech synthesis
- Automatic **translation of the input text** into the target synthesis language
- Local storage of language models for reuse
- Simple user interface

## Installation and use

1. Install all dependencies:
```
go mod tidy
```
2. Build the release version of the application:
```
go build -ldflags "-s -w"
```
3. Launch the application
4. Select the voice model
5. Enter the text


---
![preview](./preview.gif)