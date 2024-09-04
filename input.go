package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"time"
)

// OpenInputForm launches GUI for standup data input
func OpenInputForm() {
	a := app.New()
	w := a.NewWindow("Daily Standup")

	yesterdayEntry := widget.NewMultiLineEntry()
	yesterdayEntry.SetPlaceHolder("What did you work on yesterday?")

	todayEntry := widget.NewMultiLineEntry()
	todayEntry.SetPlaceHolder("What are you planning to work on today?")

	blockersEntry := widget.NewMultiLineEntry()
	blockersEntry.SetPlaceHolder("Any blockers?")

	submitButton := widget.NewButton("Submit", func() {
		if yesterdayEntry.Text == "" || todayEntry.Text == "" {
			dialog.ShowError(fmt.Errorf("Please fill out all required fields."), w)
			return
		}

		standup := StandupData{
			Date:          time.Now().Format("2006-01-02"),
			YesterdayWork: yesterdayEntry.Text,
			TodayPlan:     todayEntry.Text,
			Blockers:      blockersEntry.Text,
		}

		err := StoreData(standup)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		dialog.ShowInformation("Success", "Standup data saved successfully!", w)
		a.Quit()
	})

	form := container.NewVBox(
		widget.NewLabel("Daily Standup"),
		widget.NewSeparator(),
		widget.NewLabel("What did you work on yesterday?"),
		yesterdayEntry,
		widget.NewLabel("What are you planning to work on today?"),
		todayEntry,
		widget.NewLabel("Any blockers?"),
		blockersEntry,
		submitButton,
	)

	w.SetContent(form)
	w.Resize(fyne.NewSize(500, 600))
	w.ShowAndRun()
}
