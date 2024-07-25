package piper

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type TTS struct {
	onnxFn   string
	jsonFn   string
	piperExe string
	piperDir string
}

func (t *TTS) Synthesize(text string) ([]byte, error) {
	var output io.Writer
	outputFile := "-"

	if runtime.GOOS != "windows" {
		output = new(bytes.Buffer)
	} else {
		tmpDir, err := os.MkdirTemp("", "ab-piper.")
		if err != nil {
			return nil, fmt.Errorf("TTS.Synthesize: сreating temp directory: %w", err)
		}
		defer os.RemoveAll(tmpDir)
		outputFile = filepath.Join(tmpDir, "tts.wav")
	}

	cmd := exec.Command(t.piperExe,
		"--model", t.onnxFn,
		"--config", t.jsonFn,
		"--output_file", outputFile,
	)
	cmd.Dir = t.piperDir
	cmd.Stdin = strings.NewReader(text)
	cmd.Stdout = output
	cmd.Stderr = new(bytes.Buffer)
	cmd.SysProcAttr = sysProcAttr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("TTS.Synthesize: running piper: %w: %s", err, cmd.Stderr.(*bytes.Buffer).String())
	}

	if output != nil {
		return output.(*bytes.Buffer).Bytes(), nil
	}

	return os.ReadFile(outputFile)
}

func New(piperDir, voicesDir, voiceURL string) (*TTS, error) {
	exeFn, err := installPiper(piperDir)
	if err != nil {
		return nil, fmt.Errorf("piper.Install: cannot install piper binary: %w", err)
	}

	filePaths, err := installVoice(voiceURL, voicesDir)
	if err != nil {
		return nil, fmt.Errorf("piper.Install: cannot set piper voice: %w", err)
	}

	return &TTS{
		onnxFn:   filePaths[0],
		jsonFn:   filePaths[1],
		piperDir: filepath.Dir(exeFn),
		piperExe: exeFn,
	}, nil
}
