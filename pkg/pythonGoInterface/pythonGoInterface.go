package pythonGoInterface

import(
	"os/exec"
	"fmt"
)

func tryPythonCode(str string) () {
	pythonPath := "import pythonGoInterface"	// I don't really know how i wrote this
	cmd := exec.Command("python.exe",  "-c", pythonPath+"; print(pythonInterface.draw_styled_landmarks("+str+"))")
	out, err := cmd.CombinedOutput()
	if err != nil { fmt.Println(err); }
	fmt.Println(string(out))

}