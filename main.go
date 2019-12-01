package main

import (
	"fmt"
	"os/exec"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"github.com/go-vgo/robotgo"
)

func main() {
	// init COM
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	unknown, _ := oleutil.CreateObject("WbemScripting.SWbemLocator")
	defer unknown.Release()

	wmi, _ := unknown.QueryInterface(ole.IID_IDispatch)
	defer wmi.Release()

	// service is a SWbemServices
	serviceRaw, _ := oleutil.CallMethod(wmi, "ConnectServer")
	service := serviceRaw.ToIDispatch()
	defer service.Release()

	// result is a SWBemObjectSet
	resultRaw, _ := oleutil.CallMethod(service, "ExecNotificationQuery", "SELECT * FROM Win32_VolumeChangeEvent")
	result := resultRaw.ToIDispatch()
	defer result.Release()

	//done := make(chan bool)
	//go func() {

	for {
		// item is a SWbemObject, but really a Win32_Process
		itemRaw, _ := oleutil.CallMethod(result, "NextEvent")
		item := itemRaw.ToIDispatch()
		defer item.Release()

		asString, _ := oleutil.GetProperty(item, "DriveName")

		//Показ формы
		var answer int
		answer = robotgo.ShowAlert("VAS POSETILA CYBER POLICY ICFA", "Na diske "+
			asString.ToString()+"obnarujeno zapreschennoe anime!")

		if err := exec.Command("cmd", "/C", "logoff").Run(); err != nil {
			fmt.Println("Failed to logoff: ", err)
		}
		if answer == 0 { //Ответ ок

		} else { //Ответ отмена

		}

	}
	//}()

	//<-done
}
