package main

import (
	"encoding/json"
	"github.com/go-gl/glfw/v3.2/glfw"
	"log"
	"os"
)

const filePath = "Sources/KeyBinds.json"

var (
	Bindings = make(map[string]glfw.Key)
	Active   = make(map[glfw.Key]bool)
	CursorX  = 400.0
	CursorY  = 300.0
	ScrollX  = 0.0
	ScrollY  = 0.0
)

func init() {
	raw, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	if json.Unmarshal(raw, &Bindings) != nil {
		log.Fatal(err)
	}
}

func KeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if action == glfw.Press {
		Active[key] = true
	} else if action == glfw.Release {
		Active[key] = false
	}
}
func CursorPosCallback(w *glfw.Window, x float64, y float64) {
	CursorX = x
	CursorY = y
}
func ScrollCallBack(w *glfw.Window, x float64, y float64) {
	ScrollX = x
	ScrollY = y
}
