package main

import "git_bak/imc/src/floger"

func main() {
	floger.SetLevel(5)
	floger.Info("info")
	floger.Error("error")
	floger.Warn("warn")
	floger.Debug5("debug5")
	floger.Debug4("debug4")
	floger.Debug3("debug3")
	floger.Debug2("debug2")
	floger.Debug1("debug1")

}
