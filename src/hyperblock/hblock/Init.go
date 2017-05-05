package hblock

import (
	"fmt"

	"libguestfs.org/guestfs"

	"log"
)

func create_empty_template(obj InitParams, logger *log.Logger) (int, error) {

	//output := obj.name
	print_Log("Init hb directory.", logger)
	_, err := hb_Init()
	if err != nil {
		print_Error(err.Error(), logger)
		return FAIL, err
	}
	g, errno := guestfs.Create()
	if errno != nil {
		return FAIL, errno
	}
	//	defer
	//fmt.Println(size)
	if errCreate := g.Disk_create(obj.name, "qcow2", obj.size, nil); errCreate != nil {
		//return FAIL, errCreate
		g.Close()
		print_Panic(errCreate.Errmsg, logger)
	}
	g.Close()
	msg := fmt.Sprintf("Create template '%s' finished.", obj.name)

	print_Log(format_Success(msg), logger)
	print_Log("Creating volume named "+obj.output, logger)
	checkoutObj := CheckoutParams{layer: "", output: obj.output, template: obj.name}
	return volume_checkout(checkoutObj, logger)
}