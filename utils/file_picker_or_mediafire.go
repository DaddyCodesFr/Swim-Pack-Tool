package utils

import (
	"io"
	"swim-pack-tool/mediafire"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type FilePickerOrMediafire struct {
	reader    io.ReadCloser
	button    *widget.Button
	linkEntry *widget.Entry

	isReady          bool
	readyStateChange func(isReady bool)

	allowedExtensions []string

	filename string
}

func NewFilePickerOrMediafire(readyStateChange func(isReady bool), allowedExtensions []string) *FilePickerOrMediafire {
	entry := widget.NewEntry()
	f := &FilePickerOrMediafire{linkEntry: entry, readyStateChange: readyStateChange, allowedExtensions: allowedExtensions}
	entry.OnChanged = func(s string) {
		valid := GetURLHost(s) == "mediafire.com" || f.reader != nil
		if valid != f.isReady {
			f.isReady = valid
			f.readyStateChange(f.isReady)
		}
	}
	return f
}

func (f *FilePickerOrMediafire) Show(w fyne.Window) fyne.CanvasObject {
	if f.reader != nil {
		f.reader.Close()
		f.reader = nil
		f.button.SetText("Choose file")
	}
	f.button = widget.NewButton("Choose file", func() {
		if f.reader != nil {
			f.reader.Close()
			f.reader = nil
			f.button.SetText("Choose file")
		}
		d := dialog.NewFileOpen(func(reader fyne.URIReadCloser, err error) {
			if reader == nil {
				return
			}
			f.reader = reader
			f.filename = reader.URI().Name()
			f.button.SetText("Choose file (" + f.filename + ")")
			f.isReady = true
			f.readyStateChange(f.isReady)
		}, w)
		d.Resize(fyne.Size{Width: 1000, Height: 700})
		d.SetFilter(storage.NewExtensionFileFilter(f.allowedExtensions))
		d.Show()
	})
	return container.NewVBox(
		f.button,
		widget.NewLabelWithStyle("or enter Mediafire link:", fyne.TextAlignCenter, fyne.TextStyle{}),
		f.linkEntry,
	)
}

func (f *FilePickerOrMediafire) Clear() {
	if f.reader != nil {
		f.reader.Close()
		f.reader = nil
	}
	f.linkEntry.SetText("")
	if f.isReady {
		f.readyStateChange(false)
		f.isReady = false
	}
	f.button.SetText("Choose file")
}

func (f *FilePickerOrMediafire) Reader() (io.ReadCloser, error) {
	if f.reader != nil {
		return f.reader, nil
	}
	data, _, name, err := mediafire.MediafireDownload(f.linkEntry.Text)
	if err != nil {
		return nil, err
	}
	f.filename = name
	return data, nil
}

func (f *FilePickerOrMediafire) Filename() string {
	return f.filename
}
