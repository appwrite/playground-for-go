package main

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

var blue = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#56c5fd"))
var green = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#59f68c"))
var yellow = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#f1f89b"))
var red = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#fe5c56"))

func Info(str string) {
	fmt.Println(blue.Render(str))
}
func Danger(str string) {
	fmt.Println(red.Render(str))
}
func Success(str string) {
	fmt.Println(green.Render(str))
}
func Warning(str string) {
	fmt.Println(yellow.Render(str))
}
