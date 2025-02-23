package utils

import (
	"bytes"
	"errors"
	"image"
	"path"
	"swim-pack-tool/rar"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"github.com/swim-services/swim_porter/utils"

	_ "image/jpeg"
	_ "image/png"

	_ "golang.org/x/image/webp"
)

func ReadImage(in []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(in))
	return img, err
}

func LoadArchive(filename string, data []byte) (*utils.MapFS, error) {
	ext := path.Ext(filename)
	switch ext {
	case ".zip", ".mcpack":
		unzipped, err := utils.Unzip(data)
		if err != nil {
			return nil, err
		}
		return utils.NewMapFS(unzipped), nil
	case ".rar":
		unrared, err := rar.Unrar(data)
		if err != nil {
			return nil, err
		}
		return utils.NewMapFS(unrared), nil
	default:
		return nil, errors.New("unsupported archive type")
	}
}

func SaveFile(data []byte, name string, w fyne.Window) {
	saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err == nil && writer != nil {
			writer.Write(data)
			writer.Close()
		}
	}, w)
	saveDialog.SetFileName(name)
	saveDialog.Resize(fyne.Size{Width: 1000, Height: 700})
	saveDialog.Show()
}

func SaveMapFsAsZip(fs *utils.MapFS, name string, w fyne.Window) error {
	data, err := utils.Zip(fs.RawMap())
	if err != nil {
		return err
	}
	SaveFile(data, name, w)
	return nil
}
