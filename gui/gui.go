package gui

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"mm3d-save-editor/save"
)

var (
	mainWindow      *walk.MainWindow
	textField       *walk.LineEdit
	rupeesField     *walk.NumberEdit
	bankRupeesField *walk.NumberEdit

	greatBayCoastCheckBox   *walk.CheckBox
	zoraCapeCheckBox        *walk.CheckBox
	snowheadCheckBox        *walk.CheckBox
	mountainVillageCheckBox *walk.CheckBox
	clockTownCheckBox       *walk.CheckBox
	milkRoadCheckBox        *walk.CheckBox
	woodfallCheckBox        *walk.CheckBox
	southernSwampCheckBox   *walk.CheckBox
	ikanaCanyonCheckBox     *walk.CheckBox
	stoneTowerCheckBox      *walk.CheckBox

	filePath string
	saveFile *save.SaveFile
)

var teleportStatues = map[int]**walk.CheckBox{
	1:   &greatBayCoastCheckBox,
	2:   &zoraCapeCheckBox,
	4:   &snowheadCheckBox,
	8:   &mountainVillageCheckBox,
	16:  &clockTownCheckBox,
	32:  &milkRoadCheckBox,
	64:  &woodfallCheckBox,
	128: &southernSwampCheckBox,
	256: &ikanaCanyonCheckBox,
	512: &stoneTowerCheckBox,
}

func Open() {

	tip, err := walk.NewToolTip()
	if err != nil {
		fmt.Println("Error initializing tooltip:", err)
		return
	}
	defer tip.Dispose()

	_, err = MainWindow{
		AssignTo: &mainWindow,
		Title:    "MM3D Save Editor",
		Size:     Size{Width: 500, Height: 300},
		MinSize:  Size{Width: 500, Height: 300},
		MaxSize:  Size{Width: 500, Height: 300},
		Layout:   VBox{},
		Children: []Widget{
			Composite{
				Layout: HBox{},
				Children: []Widget{
					PushButton{
						Text: "Load Save",
						OnClicked: func() {
							dlg := new(walk.FileDialog)
							dlg.Title = "Select Save File"
							dlg.Filter = "Save Files (*.bin)|*.bin|All Files (*.*)|*.*"

							if ok, err := dlg.ShowOpen(mainWindow); err != nil {
								fmt.Println(err)
							} else if ok {
								filePath = dlg.FilePath
								loadSave()
							}
						},
					},
					PushButton{
						Text: "Save File",
						OnClicked: func() {
							saveFiles()
						},
					},
				},
			},
			Composite{
				Layout: VBox{},
				Children: []Widget{
					Label{
						Text: "Name:",
					},
					LineEdit{
						AssignTo: &textField,
						OnTextChanged: func() {
							text := textField.Text()
							if len(text) > 8 {
								err := textField.SetText(text[:8])
								if err != nil {
									return
								}
							}
						},
					},
					Composite{
						Layout: HBox{},
						Children: []Widget{
							Composite{
								Layout: VBox{},
								Children: []Widget{
									Label{
										Text: "Rupees:",
									},
									NumberEdit{
										AssignTo: &rupeesField,
										MinValue: 0,
										MaxValue: 999,
									},
								},
							},
							Composite{
								Layout: VBox{},
								Children: []Widget{
									Label{
										Text: "Bank Rupees:",
									},
									NumberEdit{
										AssignTo: &bankRupeesField,
										MinValue: 0,
										MaxValue: 65535,
									},
								},
							},
						},
					},
				},
			},
			Composite{
				Layout: VBox{},
				Children: []Widget{
					Label{
						Text: "Teleport Statue:",
					},
					Composite{
						Layout: VBox{},
						Children: []Widget{
							CheckBox{
								AssignTo: &greatBayCoastCheckBox,
								Text:     "Great Bay Coast",
							},
							CheckBox{
								AssignTo: &zoraCapeCheckBox,
								Text:     "Zora Cape",
							},
							CheckBox{
								AssignTo: &snowheadCheckBox,
								Text:     "Snowhead",
							},
							CheckBox{
								AssignTo: &mountainVillageCheckBox,
								Text:     "Mountain Village",
							},
							CheckBox{
								AssignTo: &clockTownCheckBox,
								Text:     "Clock Town",
							},
							CheckBox{
								AssignTo: &milkRoadCheckBox,
								Text:     "Milk Road",
							},
							CheckBox{
								AssignTo: &woodfallCheckBox,
								Text:     "Woodfall",
							},
							CheckBox{
								AssignTo: &southernSwampCheckBox,
								Text:     "Southern Swamp",
							},
							CheckBox{
								AssignTo: &ikanaCanyonCheckBox,
								Text:     "Ikana Canyon",
							},
							CheckBox{
								AssignTo: &stoneTowerCheckBox,
								Text:     "Stone Tower",
							},
						},
					},
				},
			},
		},
	}.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func saveFiles() {

	var statueIds []int
	for id, checkbox := range teleportStatues {
		if (*checkbox).Checked() {
			statueIds = append(statueIds, id)
		}
	}
	saveFile.TeleportStatueIds = statueIds

	saveFile.Name = textField.Text()
	saveFile.Money = int(rupeesField.Value())
	saveFile.Bank = int(bankRupeesField.Value())

	err := saveFile.SaveToFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	walk.MsgBox(mainWindow, "MM3D Save Editor", "Save file has been updated", walk.MsgBoxOK|walk.MsgBoxIconInformation)
}

func loadSave() {
	load, err := save.Load(filePath)
	saveFile = load
	if err != nil {
		fmt.Println(err)
		return
	}

	err = textField.SetText(load.Name)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = rupeesField.SetValue(float64(load.Money))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = bankRupeesField.SetValue(float64(load.Bank))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, id := range load.TeleportStatueIds {
		if checkbox, exists := teleportStatues[id]; exists {
			(*checkbox).SetChecked(true)
		}
	}
}
