package main

import "net/http"
import "os/exec"
import "log"
import "flag"

var httpFlag *string = flag.String("http", ":80", "default host/port to serve from")

func commandHandler(cmd string, arg ...string) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		e := exec.Command(cmd, arg...)
		log.Println(req.RemoteAddr, e.Args)
		bytes, _ := e.CombinedOutput()
		rw.Write(bytes)
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
	http.HandleFunc("/killall/vlc", commandHandler("killall", "vlc"))

	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir("web/")))

	err := http.ListenAndServe(*httpFlag, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
