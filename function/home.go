package function

import (
	"ascii_art_web/ascii"
	logger2 "ascii_art_web/logger"
	"html/template"
	"net/http"
)

var logger = logger2.GetLogger()

func GetHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("cmd/index.html"))

	err := tmpl.Execute(w, nil)
	if err != nil {
		logger.Info("some error 500")
	}
	logger.Info("Create page for home")
}

func GetAscii(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	tmpl := template.Must(template.ParseFiles("cmd/index.html"))
	logger.Info("Create page for Ascii")

	arg := r.FormValue("ascii_convert")

	if len(arg) > 400 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	banner := r.FormValue("fonts")
	allBanner := r.FormValue("allBanner")

	if allBanner == "all" {
		shadow, errB := ascii.Ascii(arg, "Shadow")
		if errB != nil {
			logger.Infof("error:%v", errB)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusInternalServerError)
			return
		}
		standard, errB := ascii.Ascii(arg, "Standard")
		if errB != nil {
			logger.Infof("error:%v", errB)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusInternalServerError)
			return
		}
		thinkertoy, errB := ascii.Ascii(arg, "Thinkertoy")
		if errB != nil {
			logger.Infof("error:%v", errB)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusInternalServerError)
			return
		}

		err := tmpl.Execute(w, standard+shadow+thinkertoy)
		if err != nil {
			logger.Infof("error:%v", err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusInternalServerError)
			return
		}
	} else {
		oneChoice, err := ascii.Ascii(arg, banner)
		if err != nil {
			logger.Infof("error:%v", err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, oneChoice)
		if err != nil {
			logger.Infof("error:%v", err)
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusInternalServerError)
			return
		}
	}

}
