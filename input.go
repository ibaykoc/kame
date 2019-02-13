package kame

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Key int
type KeyAction int
type Input struct {
	keyStats    map[Key]KeyAction
	MouseX      float64
	MouseY      float64
	MouseDeltaX float64
	MouseDeltaY float64
	prevMouseX  float64
	prevMouseY  float64
	glfwWindow  *glfw.Window
}

const (
	Press KeyAction = iota
	JustPress
	Release
	JustRelease
)

const (
	KeyUnknown      Key = Key(glfw.KeyUnknown)
	KeySpace        Key = Key(glfw.KeySpace)
	KeyApostrophe   Key = Key(glfw.KeyApostrophe)
	KeyComma        Key = Key(glfw.KeyComma)
	KeyMinus        Key = Key(glfw.KeyMinus)
	KeyPeriod       Key = Key(glfw.KeyPeriod)
	KeySlash        Key = Key(glfw.KeySlash)
	Key0            Key = Key(glfw.Key0)
	Key1            Key = Key(glfw.Key1)
	Key2            Key = Key(glfw.Key2)
	Key3            Key = Key(glfw.Key3)
	Key4            Key = Key(glfw.Key4)
	Key5            Key = Key(glfw.Key5)
	Key6            Key = Key(glfw.Key6)
	Key7            Key = Key(glfw.Key7)
	Key8            Key = Key(glfw.Key8)
	Key9            Key = Key(glfw.Key9)
	KeySemicolon    Key = Key(glfw.KeySemicolon)
	KeyEqual        Key = Key(glfw.KeyEqual)
	KeyA            Key = Key(glfw.KeyA)
	KeyB            Key = Key(glfw.KeyB)
	KeyC            Key = Key(glfw.KeyC)
	KeyD            Key = Key(glfw.KeyD)
	KeyE            Key = Key(glfw.KeyE)
	KeyF            Key = Key(glfw.KeyF)
	KeyG            Key = Key(glfw.KeyG)
	KeyH            Key = Key(glfw.KeyH)
	KeyI            Key = Key(glfw.KeyI)
	KeyJ            Key = Key(glfw.KeyJ)
	KeyK            Key = Key(glfw.KeyK)
	KeyL            Key = Key(glfw.KeyL)
	KeyM            Key = Key(glfw.KeyM)
	KeyN            Key = Key(glfw.KeyN)
	KeyO            Key = Key(glfw.KeyO)
	KeyP            Key = Key(glfw.KeyP)
	KeyQ            Key = Key(glfw.KeyQ)
	KeyR            Key = Key(glfw.KeyR)
	KeyS            Key = Key(glfw.KeyS)
	KeyT            Key = Key(glfw.KeyT)
	KeyU            Key = Key(glfw.KeyU)
	KeyV            Key = Key(glfw.KeyV)
	KeyW            Key = Key(glfw.KeyW)
	KeyX            Key = Key(glfw.KeyX)
	KeyY            Key = Key(glfw.KeyY)
	KeyZ            Key = Key(glfw.KeyZ)
	KeyLeftBracket  Key = Key(glfw.KeyLeftBracket)
	KeyBackslash    Key = Key(glfw.KeyBackslash)
	KeyRightBracket Key = Key(glfw.KeyRightBracket)
	KeyGraveAccent  Key = Key(glfw.KeyGraveAccent)
	KeyWorld1       Key = Key(glfw.KeyWorld1)
	KeyWorld2       Key = Key(glfw.KeyWorld2)
	KeyEscape       Key = Key(glfw.KeyEscape)
	KeyEnter        Key = Key(glfw.KeyEnter)
	KeyTab          Key = Key(glfw.KeyTab)
	KeyBackspace    Key = Key(glfw.KeyBackspace)
	KeyInsert       Key = Key(glfw.KeyInsert)
	KeyDelete       Key = Key(glfw.KeyDelete)
	KeyRight        Key = Key(glfw.KeyRight)
	KeyLeft         Key = Key(glfw.KeyLeft)
	KeyDown         Key = Key(glfw.KeyDown)
	KeyUp           Key = Key(glfw.KeyUp)
	KeyPageUp       Key = Key(glfw.KeyPageUp)
	KeyPageDown     Key = Key(glfw.KeyPageDown)
	KeyHome         Key = Key(glfw.KeyHome)
	KeyEnd          Key = Key(glfw.KeyEnd)
	KeyCapsLock     Key = Key(glfw.KeyCapsLock)
	KeyScrollLock   Key = Key(glfw.KeyScrollLock)
	KeyNumLock      Key = Key(glfw.KeyNumLock)
	KeyPrintScreen  Key = Key(glfw.KeyPrintScreen)
	KeyPause        Key = Key(glfw.KeyPause)
	KeyF1           Key = Key(glfw.KeyF1)
	KeyF2           Key = Key(glfw.KeyF2)
	KeyF3           Key = Key(glfw.KeyF3)
	KeyF4           Key = Key(glfw.KeyF4)
	KeyF5           Key = Key(glfw.KeyF5)
	KeyF6           Key = Key(glfw.KeyF6)
	KeyF7           Key = Key(glfw.KeyF7)
	KeyF8           Key = Key(glfw.KeyF8)
	KeyF9           Key = Key(glfw.KeyF9)
	KeyF10          Key = Key(glfw.KeyF10)
	KeyF11          Key = Key(glfw.KeyF11)
	KeyF12          Key = Key(glfw.KeyF12)
	KeyF13          Key = Key(glfw.KeyF13)
	KeyF14          Key = Key(glfw.KeyF14)
	KeyF15          Key = Key(glfw.KeyF15)
	KeyF16          Key = Key(glfw.KeyF16)
	KeyF17          Key = Key(glfw.KeyF17)
	KeyF18          Key = Key(glfw.KeyF18)
	KeyF19          Key = Key(glfw.KeyF19)
	KeyF20          Key = Key(glfw.KeyF20)
	KeyF21          Key = Key(glfw.KeyF21)
	KeyF22          Key = Key(glfw.KeyF22)
	KeyF23          Key = Key(glfw.KeyF23)
	KeyF24          Key = Key(glfw.KeyF24)
	KeyF25          Key = Key(glfw.KeyF25)
	KeyKP0          Key = Key(glfw.KeyKP0)
	KeyKP1          Key = Key(glfw.KeyKP1)
	KeyKP2          Key = Key(glfw.KeyKP2)
	KeyKP3          Key = Key(glfw.KeyKP3)
	KeyKP4          Key = Key(glfw.KeyKP4)
	KeyKP5          Key = Key(glfw.KeyKP5)
	KeyKP6          Key = Key(glfw.KeyKP6)
	KeyKP7          Key = Key(glfw.KeyKP7)
	KeyKP8          Key = Key(glfw.KeyKP8)
	KeyKP9          Key = Key(glfw.KeyKP9)
	KeyKPDecimal    Key = Key(glfw.KeyKPDecimal)
	KeyKPDivide     Key = Key(glfw.KeyKPDivide)
	KeyKPMultiply   Key = Key(glfw.KeyKPMultiply)
	KeyKPSubtract   Key = Key(glfw.KeyKPSubtract)
	KeyKPAdd        Key = Key(glfw.KeyKPAdd)
	KeyKPEnter      Key = Key(glfw.KeyKPEnter)
	KeyKPEqual      Key = Key(glfw.KeyKPEqual)
	KeyLeftShift    Key = Key(glfw.KeyLeftShift)
	KeyLeftControl  Key = Key(glfw.KeyLeftControl)
	KeyLeftAlt      Key = Key(glfw.KeyLeftAlt)
	KeyLeftSuper    Key = Key(glfw.KeyLeftSuper)
	KeyRightShift   Key = Key(glfw.KeyRightShift)
	KeyRightControl Key = Key(glfw.KeyRightControl)
	KeyRightAlt     Key = Key(glfw.KeyRightAlt)
	KeyRightSuper   Key = Key(glfw.KeyRightSuper)
	KeyMenu         Key = Key(glfw.KeyMenu)
	KeyLast         Key = Key(glfw.KeyLast)
)

func newInput(glfwWindow *glfw.Window) Input {
	ks := make(map[Key]KeyAction)
	ks[KeyUnknown] = Release
	ks[KeySpace] = Release
	ks[KeyApostrophe] = Release
	ks[KeyComma] = Release
	ks[KeyMinus] = Release
	ks[KeyPeriod] = Release
	ks[KeySlash] = Release
	ks[Key0] = Release
	ks[Key1] = Release
	ks[Key2] = Release
	ks[Key3] = Release
	ks[Key4] = Release
	ks[Key5] = Release
	ks[Key6] = Release
	ks[Key7] = Release
	ks[Key8] = Release
	ks[Key9] = Release
	ks[KeySemicolon] = Release
	ks[KeyEqual] = Release
	ks[KeyA] = Release
	ks[KeyB] = Release
	ks[KeyC] = Release
	ks[KeyD] = Release
	ks[KeyE] = Release
	ks[KeyF] = Release
	ks[KeyG] = Release
	ks[KeyH] = Release
	ks[KeyI] = Release
	ks[KeyJ] = Release
	ks[KeyK] = Release
	ks[KeyL] = Release
	ks[KeyM] = Release
	ks[KeyN] = Release
	ks[KeyO] = Release
	ks[KeyP] = Release
	ks[KeyQ] = Release
	ks[KeyR] = Release
	ks[KeyS] = Release
	ks[KeyT] = Release
	ks[KeyU] = Release
	ks[KeyV] = Release
	ks[KeyW] = Release
	ks[KeyX] = Release
	ks[KeyY] = Release
	ks[KeyZ] = Release
	ks[KeyLeftBracket] = Release
	ks[KeyBackslash] = Release
	ks[KeyRightBracket] = Release
	ks[KeyGraveAccent] = Release
	ks[KeyWorld1] = Release
	ks[KeyWorld2] = Release
	ks[KeyEscape] = Release
	ks[KeyEnter] = Release
	ks[KeyTab] = Release
	ks[KeyBackspace] = Release
	ks[KeyInsert] = Release
	ks[KeyDelete] = Release
	ks[KeyRight] = Release
	ks[KeyLeft] = Release
	ks[KeyDown] = Release
	ks[KeyUp] = Release
	ks[KeyPageUp] = Release
	ks[KeyPageDown] = Release
	ks[KeyHome] = Release
	ks[KeyEnd] = Release
	ks[KeyCapsLock] = Release
	ks[KeyScrollLock] = Release
	ks[KeyNumLock] = Release
	ks[KeyPrintScreen] = Release
	ks[KeyPause] = Release
	ks[KeyF1] = Release
	ks[KeyF2] = Release
	ks[KeyF3] = Release
	ks[KeyF4] = Release
	ks[KeyF5] = Release
	ks[KeyF6] = Release
	ks[KeyF7] = Release
	ks[KeyF8] = Release
	ks[KeyF9] = Release
	ks[KeyF10] = Release
	ks[KeyF11] = Release
	ks[KeyF12] = Release
	ks[KeyF13] = Release
	ks[KeyF14] = Release
	ks[KeyF15] = Release
	ks[KeyF16] = Release
	ks[KeyF17] = Release
	ks[KeyF18] = Release
	ks[KeyF19] = Release
	ks[KeyF20] = Release
	ks[KeyF21] = Release
	ks[KeyF22] = Release
	ks[KeyF23] = Release
	ks[KeyF24] = Release
	ks[KeyF25] = Release
	ks[KeyKP0] = Release
	ks[KeyKP1] = Release
	ks[KeyKP2] = Release
	ks[KeyKP3] = Release
	ks[KeyKP4] = Release
	ks[KeyKP5] = Release
	ks[KeyKP6] = Release
	ks[KeyKP7] = Release
	ks[KeyKP8] = Release
	ks[KeyKP9] = Release
	ks[KeyKPDecimal] = Release
	ks[KeyKPDivide] = Release
	ks[KeyKPMultiply] = Release
	ks[KeyKPSubtract] = Release
	ks[KeyKPAdd] = Release
	ks[KeyKPEnter] = Release
	ks[KeyKPEqual] = Release
	ks[KeyLeftShift] = Release
	ks[KeyLeftControl] = Release
	ks[KeyLeftAlt] = Release
	ks[KeyLeftSuper] = Release
	ks[KeyRightShift] = Release
	ks[KeyRightControl] = Release
	ks[KeyRightAlt] = Release
	ks[KeyRightSuper] = Release
	ks[KeyMenu] = Release
	ks[KeyLast] = Release

	mX, mY := glfwWindow.GetCursorPos()
	glfwWindow.SetInputMode(glfw.CursorMode, glfw.CursorDisabled)
	return Input{
		keyStats:   ks,
		glfwWindow: glfwWindow,
		MouseX:     mX,
		MouseY:     mY,
		prevMouseX: mX,
		prevMouseY: mY,
	}
}

func (i *Input) glfwInputHandler(glfwKey glfw.Key, glfwAction glfw.Action) {
	prevStat := i.keyStats[Key(glfwKey)]
	if glfwAction == glfw.Press {
		if prevStat != Press {
			i.keyStats[Key(glfwKey)] = JustPress
		} else {
			i.keyStats[Key(glfwKey)] = Press
		}
	} else if glfwAction == glfw.Release {
		if prevStat != Release {
			i.keyStats[Key(glfwKey)] = JustRelease
		} else {
			i.keyStats[Key(glfwKey)] = Release
		}
	}
}

func (i *Input) glfwMousePosHandler(x, y float64) {
	i.MouseX = x
	i.MouseY = y
}

func (i *Input) update() {
	// Update Key Stat
	for k, v := range i.keyStats {
		if v == JustPress {
			i.keyStats[k] = Press
		} else if v == JustRelease {
			i.keyStats[k] = Release
		}
	}

	// Update Mouse Stat
	i.MouseDeltaX = i.MouseX - i.prevMouseX
	i.MouseDeltaY = i.MouseY - i.prevMouseY
	i.prevMouseX = i.MouseX
	i.prevMouseY = i.MouseY
}

func (i *Input) GetKeyStat(key Key) KeyAction {
	return i.keyStats[key]
}
