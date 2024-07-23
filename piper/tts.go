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

func (t *TTS) Synthesize(text string) (wav []byte, err error) {
	stdoutFn := "-"
	var stdout io.Writer
	if runtime.GOOS != "windows" {
		stdout = bytes.NewBuffer(nil)
	} else {
		tmpDir, err := os.MkdirTemp("", "ab-piper.")
		if err != nil {
			return nil, fmt.Errorf("TTS.Synthesize: Cannot create temp file: %w", err)
		}
		defer os.RemoveAll(tmpDir)
		stdoutFn = filepath.Join(tmpDir, "tts.wav")
	}

	stdin := strings.NewReader(text)
	stderr := bytes.NewBuffer(nil)
	cmd := exec.Command(t.piperExe,
		"--model", t.onnxFn,
		"--config", t.jsonFn,
		"--output_file", stdoutFn)
	cmd.Dir = t.piperDir
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.SysProcAttr = sysProcAttr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("TTS.Synthesize: invalid UTF-8 path: %s: %s: %s", cmd, err, stderr.Bytes())
	}

	if stdout != nil {
		return stdout.(*bytes.Buffer).Bytes(), nil
	}

	wav, err = os.ReadFile(stdoutFn)
	if err != nil {
		return nil, fmt.Errorf("TTS.Synthesize: %s", err)
	}
	return wav, nil
}

func New(piperDir, voicesDir, voiceUrl string) (*TTS, error) {
	exeFn, err := installPiper(piperDir)
	if err != nil {
		return nil, fmt.Errorf("piper.Install: cannot install piper binary: %w", err)
	}

	filePaths, err := installVoice(voiceUrl, voicesDir)
	if err != nil {
		return nil, fmt.Errorf("piper.Install: cannot set piper voice: %w", err)
	}

	t := &TTS{
		onnxFn:   filePaths[0],
		jsonFn:   filePaths[1],
		piperDir: filepath.Dir(exeFn),
		piperExe: exeFn,
	}
	return t, nil
}
