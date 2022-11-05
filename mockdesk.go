package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fynedesk"
)

type mockDesk struct {
	fynedesk.Desktop
}

func (m *mockDesk) Screens() fynedesk.ScreenList {
	return &mockScreenList{
		[]*fynedesk.Screen{{
			"Primary",
			0, 0, 640, 480,
			1,
		}},
	}
}

func (m *mockDesk) ShowMenuAt(menu *fyne.Menu, pos fyne.Position) {
	log.Println("Failed to pop up menu:", menu.Label)
}

func (m *mockDesk) WindowManager() fynedesk.WindowManager {
	return &mockWM{}
}

type mockScreenList struct {
	screens []*fynedesk.Screen
}

func (m *mockScreenList) RefreshScreens() {

}

func (m *mockScreenList) AddChangeListener(func()) {

}

func (m *mockScreenList) Screens() []*fynedesk.Screen {
	return m.screens
}

func (m *mockScreenList) SetActive(*fynedesk.Screen) {
}

func (m *mockScreenList) Active() *fynedesk.Screen {
	return m.screens[0]
}

func (m *mockScreenList) Primary() *fynedesk.Screen {
	return m.screens[0]
}

func (m *mockScreenList) ScreenForWindow(fynedesk.Window) *fynedesk.Screen {
	return m.screens[0]
}

func (m *mockScreenList) ScreenForGeometry(x, y, width, height int) *fynedesk.Screen {
	return m.screens[0]
}

type mockWM struct {
	fynedesk.WindowManager
}

func (m *mockWM) ShowOverlay(w fyne.Window, s fyne.Size, p fyne.Position) {
	// TODO find a way to position
	w.Resize(s)
	w.Show()
}

