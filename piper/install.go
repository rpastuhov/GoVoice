package piper

import (
	"archive/tar"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/schollz/progressbar/v3"
)

const (
	DistArcName  = "dist.tzst"
	DistMetaName = "dist.json"
)

func extractTar(rootDir, dstNm string, h *tar.Header, r io.Reader) (retErr error) {
	rootDir = filepath.Clean(rootDir)
	dstFn := filepath.Join(rootDir, dstNm)
	rel, err := filepath.Rel(rootDir, dstFn)
	if err != nil {
		return fmt.Errorf("extract: rel(%s,%s): %w", rootDir, dstFn, err)
	}
	if filepath.Join(rootDir, rel) != dstFn {
		return fmt.Errorf("extract: `%s` appears to escaped root `%s`", dstFn, rootDir)
	}

	if h.Typeflag == tar.TypeDir {
		return nil
	}

	if h.Typeflag == tar.TypeSymlink && h.Linkname != "" {
		err := os.Symlink(h.Linkname, dstFn)
		if err != nil {
			return fmt.Errorf("extract: link(%s, %s): `%w`", h.Linkname, dstFn, err)
		}
		return nil
	}

	if h.Typeflag != tar.TypeReg {
		return fmt.Errorf("extract: unsupported file `%s`: type(%d) is not a dir(%d), symlink(%d) or regular file(%d)",
			dstFn, h.Typeflag, tar.TypeDir, tar.TypeSymlink, tar.TypeReg)
	}

	os.MkdirAll(filepath.Dir(dstFn), 0755)
	f, err := os.Create(dstFn)
	if err != nil {
		return fmt.Errorf("extract: create `%s`: %w", dstFn, err)
	}

	_, copyErr := io.Copy(f, r)
	closeErr := f.Close()
	if copyErr != nil {
		return fmt.Errorf("extract: copy `%s`: %w", dstFn, copyErr)
	}
	if closeErr != nil {
		return fmt.Errorf("extract: close `%s`: %w", dstFn, closeErr)
	}
	return nil
}

func rimraf(fn string) {
	if !filepath.IsAbs(fn) {
		panic(fn + ": is not absolute")
	}
	os.RemoveAll(fn)
}

func installArc(dstDir string, srcFS fs.FS) error {
	dstDir = filepath.Clean(dstDir)
	if !filepath.IsAbs(dstDir) {
		return fmt.Errorf("installArc: `%s` is not absolute", dstDir)
	}

	os.MkdirAll(filepath.Dir(dstDir), 0755)

	alreadyInstalled, tmpDir, err := installMeta(dstDir, srcFS)
	if err != nil {
		return fmt.Errorf("extract: Cannot create temp dir: %w", err)
	}
	if alreadyInstalled {
		return nil
	}
	defer rimraf(tmpDir)

	arcRd, err := srcFS.Open(DistArcName)
	if err != nil {
		return fmt.Errorf("installArc: open fs `%s`: %w", DistArcName, err)
	}
	defer arcRd.Close()

	arc, err := openTarZst(arcRd)
	if err != nil {
		return fmt.Errorf("installArc: open archive `%s`: %w", DistArcName, err)
	}
	defer arc.Close()

	for {
		h, err := arc.Next()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("installArc: iter `%s`: %w", DistArcName, err)
		}
		if err := extractTar(tmpDir, h.Name, h, arc); err != nil {
			return fmt.Errorf("installArc: extract `%s`: %w", DistArcName, err)
		}
	}

	bakDir := fmt.Sprintf("%s.%d.bak", dstDir, time.Now().UnixNano())
	if err := os.Rename(dstDir, bakDir); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("extractTarZst: Cannot rename `%s` to `%s` %w", dstDir, bakDir, err)
	}
	if err := os.Rename(tmpDir, dstDir); err != nil {
		os.Rename(bakDir, dstDir)
		return fmt.Errorf("extractTarZst: Cannot rename `%s` to `%s` %w", tmpDir, dstDir, err)
	}
	rimraf(bakDir)
	return nil
}

func installMeta(dstDir string, srcFS fs.FS) (alreadyInstalled bool, tmpDir string, err error) {
	srcMeta, err := fs.ReadFile(srcFS, DistMetaName)
	if err != nil {
		return false, "", fmt.Errorf("installMeta: Cannot read meta: %w", err)
	}

	dstMeta, err := os.ReadFile(filepath.Join(dstDir, DistMetaName))
	if err == nil && bytes.Equal(dstMeta, srcMeta) {
		return true, "", nil
	}

	tmpDir, err = os.MkdirTemp(filepath.Dir(dstDir), filepath.Base(dstDir))
	if err != nil {
		return false, "", fmt.Errorf("installMeta: Cannot create temp dir: %w", err)
	}

	if err := os.WriteFile(filepath.Join(tmpDir, DistMetaName), srcMeta, 0644); err != nil {
		rimraf(tmpDir)
		return false, "", fmt.Errorf("installMeta: write meta file `%s`: %w", tmpDir, err)
	}
	return false, tmpDir, nil
}

func installPiper(dataDir string) (exeFn string, err error) {
	// pkgDir := filepath.Join(dataDir, "piper-bin-"+piperAsset.Name)
	dataDir, err = filepath.Abs(dataDir)
	if err != nil {
		return "", fmt.Errorf("installArc: `%s` is not correct", dataDir)
	}

	defer os.Chmod(exeFn, 0755)
	if err := installArc(dataDir, piperAsset.FS); err != nil {
		return "", fmt.Errorf("installPiper: walk embedded fs: %w", err)
	}
	return filepath.Join(dataDir, piperExe), nil
}

//
//
//
//

func downloadWithBar(fileName, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("response during download: %w", err)
	}
	defer resp.Body.Close()

	if err := os.MkdirAll(filepath.Dir(fileName), 0755); err != nil {
		return fmt.Errorf("creating directories: %s", err)
	}

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("creating file: %s", err)
	}
	defer file.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
	)

	if _, err := io.Copy(io.MultiWriter(bar, file), resp.Body); err != nil {
		return fmt.Errorf("copying data: %w", err)
	}

	return nil
}

func installVoice(voiceURL, folderName string) ([]string, error) {
	extensions := []string{".onnx", ".onnx.json"}
	fileName := filepath.Base(voiceURL)

	absFolder, err := filepath.Abs(folderName)
	if err != nil {
		return nil, fmt.Errorf("getting absolute folder path: %w", err)
	}

	filePaths := make([]string, len(extensions))
	for i, ext := range extensions {
		targetPath := filepath.Join(absFolder, fileName+ext)

		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			if err := downloadWithBar(targetPath, voiceURL+ext); err != nil {
				return nil, fmt.Errorf("downloading file: %w", err)
			}
		}

		filePaths[i] = targetPath
	}

	return filePaths, nil
}
