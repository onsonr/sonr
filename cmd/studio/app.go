package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"os"
	"path/filepath"
	r "runtime"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	libdir, _ = os.UserConfigDir()
	basedir   = filepath.Join(libdir, "sonr-studio")
	docsdir   = filepath.Join(basedir, "Documents")
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// Startup is called at application Startup
func (app *App) Startup(ctx context.Context) {
	app.ctx = ctx

	m := menu.NewMenu()

	if r.GOOS == "darwin" {
		m.Items = append(m.Items, menu.AppMenu())
	}

	m.Items = append(m.Items, menu.SubMenu("File", menu.NewMenuFromItems(
		menu.Text("Refresh", keys.CmdOrCtrl("r"), func(cd *menu.CallbackData) {
			runtime.EventsEmit(ctx, "shortcut.view.refresh")
		}),
		menu.Text("Hard Refresh", keys.Combo("r", keys.CmdOrCtrlKey, keys.ShiftKey), func(cd *menu.CallbackData) {
			runtime.EventsEmit(ctx, "shortcut.view.hard-refresh")
		}),
		menu.Separator(),
		menu.Text("Open Collection", keys.CmdOrCtrl("o"), func(cd *menu.CallbackData) {
			runtime.EventsEmit(ctx, "shortcut.collection.open")
		}),
		menu.Text("Save Collection", keys.CmdOrCtrl("s"), func(cd *menu.CallbackData) {
			runtime.EventsEmit(ctx, "shortcut.collection.save")
		}),
		menu.Separator(),
		menu.Text("Print...", keys.CmdOrCtrl("p"), func(cd *menu.CallbackData) {
			runtime.EventsEmit(ctx, "shortcut.collection.print")
		}),
	)))

	if r.GOOS == "darwin" {
		m.Items = append(m.Items, menu.EditMenu())
	}

	m.Items = append(m.Items, menu.SubMenu("Language", menu.NewMenuFromItems(
		menu.Text("🇺🇸 English", nil, func(cd *menu.CallbackData) {
			runtime.EventsEmit(ctx, "shortcut.language.english")
		}),
		menu.Text("🇪🇸 Español", nil, func(cd *menu.CallbackData) {
			runtime.EventsEmit(ctx, "shortcut.language.spanish")
		}),
	)))

	runtime.MenuSetApplicationMenu(ctx, m)
}

// DomReady is called after the front-end dom has been loaded
func (app *App) DomReady(ctx context.Context) {
	// Add your action here
}

// Shutdown is called at application termination
func (app *App) Shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (app *App) Title() string {
	if r.GOOS == "darwin" {
		return "🦄 Varly"
	}

	return "Varly"
}

func (app *App) OpenDirectoryDialog(title string) string {
	path, _ := runtime.OpenDirectoryDialog(app.ctx, runtime.OpenDialogOptions{
		Title:                      title,
		CanCreateDirectories:       true,
		TreatPackagesAsDirectories: true,
	})

	return path
}

func (app *App) OpenFileDialog() string {
	path, _ := runtime.OpenFileDialog(app.ctx, runtime.OpenDialogOptions{})

	return path
}

func (app *App) SaveFileDialog() string {
	path, _ := runtime.SaveFileDialog(app.ctx, runtime.SaveDialogOptions{})

	return path
}

func (app *App) EncodeImage(path string) string {
	image, err := os.ReadFile(path)
	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}

	encoded := base64.StdEncoding.EncodeToString(image)
	encoded = fmt.Sprintf("data:image/png;base64,%s", encoded)

	return encoded
}

func (app *App) SaveFile(file string, data string) bool {
	path := fmt.Sprintf("%s%s", docsdir, file)

	err := os.WriteFile(path, []byte(data), os.ModePerm)

	return err == nil
}

func (app *App) GetImageStats(path string) image.Config {
	reader, err := os.Open(path)
	if err != nil {
		return image.Config{}
	}
	defer reader.Close()
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		return image.Config{}
	}

	return img
}

func (app *App) MessageDialog(options runtime.MessageDialogOptions) string {
	res, _ := runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
		Type:          runtime.QuestionDialog,
		Title:         options.Title,
		Message:       options.Message,
		Buttons:       options.Buttons,
		DefaultButton: options.DefaultButton,
	})

	return res
}
