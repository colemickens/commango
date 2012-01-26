package main

import "fmt"
import _ "flag"
import "net/http"
import "os/exec"
import "log"

func commandHandler(cmd string, arg ...string) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		e := exec.Command(cmd, arg...)
		fmt.Println(req.RemoteAddr, ":", cmd, arg)
		e.Run()
		http.Redirect(rw, req, "/", 302)
	}
}

func main() {
	http.HandleFunc("/blank", commandHandler("/etc/acpi/screenblank.sh"))
	http.HandleFunc("/mute", commandHandler("amixer", "set", "Master", "0"))
	http.HandleFunc("/voldown", commandHandler("amixer", "set", "Master", "5%-"))
	http.HandleFunc("/volup", commandHandler("amixer", "set", "Master", "5%+"))
	http.HandleFunc("/screen/left", commandHandler("/home/cole/Code/scripts/disper/disper-secondary"))
	http.HandleFunc("/screen/right", commandHandler("/home/cole/Code/scripts/disper/disper-primary"))
	http.HandleFunc("/screen/both", commandHandler("/home/cole/Code/scripts/disper/disper-dual"))

	http.Handle("/", http.FileServer(http.Dir("web/")))

	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}