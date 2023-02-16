package Utils

import (
	"encoding/json"
	"github.com/go-gl/glfw/v3.1/glfw"
	"log"
	"os"
)

const filePath = "Sources/KeyBinds.json"

type KeyStatus struct {
	Pressed    bool
	Registered bool
}

var Bindings map[string]glfw.Key
var Status = make(map[glfw.Key]KeyStatus)

func init() {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		log.Print(err)
	}

	if json.Unmarshal(bytes, &Bindings) != nil {
		log.Print(err)
	}

	for _, key := range Bindings {
		Status[key] = KeyStatus{}
	}

}

func KeyListener(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if _, ok := Status[key]; ok {

		if action == glfw.Press {
			Status[key].Pressed = true
		} else if action == glfw.Release {
			Status[key].Pressed = false

		}

	}
}

func exit() {
	marshal, err := json.Marshal(Bindings)
	if err != nil {
		log.Print(err)
	}

	if os.WriteFile(filePath, marshal, 0666) != nil {
		log.Print(err)
	}
}
