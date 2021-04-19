// +build !windows,!darwin

package volume

import (
	"errors"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var useAmixer = false
var pattern = regexp.MustCompile(`\d+%`)

func execCmd(args []string) ([]byte, error) {
	cmd := exec.Command(args[0], args[1:]...)
	out, err := cmd.Output()
	if err != nil {
		return nil, errors.New("Could not execute command" + strings.Join(args, " "))
	}
	return out, nil
}

func init() {
	useAmixer = false

	if _, err := execCmd([]string{"pactl", "info"}); err != nil {
		useAmixer = true
	}
}

func getVolumeCmd() []string {
	if useAmixer {
		return []string{"amixer", "get", "Master"}
	}
	return []string{"pactl", "list", "sinks"}
}

// GetSpeakerVolume Get the volume of the main speaker
func GetSpeakerVolume() (int, error) {
	out, err := execCmd(getVolumeCmd())
	if err != nil {
		return 0, err
	}
	return parseVolume(string(out))
}

func setVolumeCmd(volume int) []string {
	if useAmixer {
		return []string{"amixer", "set", "Master", strconv.Itoa(volume) + "%"}
	}
	return []string{"pactl", "set-sink-volume", "@DEFAULT_SINK@", strconv.Itoa(volume) + "%"}
}

// SetSpeakerVolume Sets the volume of the speaker to the provided value
func SetSpeakerVolume(v int) (int, error) {
	if outOfRange(v) {
		return 0, errors.New("Value out of range")
	}
	_, err := execCmd(setVolumeCmd(v))
	if err != nil {
		return 0, err
	}
	return v, nil
}

func parseVolume(v string) (int, error) {
	lines := strings.Split(v, "\n")
	for _, line := range lines {
		str := strings.TrimLeft(line, " \t")

		if (useAmixer && strings.Contains(str, "Playback") && strings.Contains(str, "%")) || !useAmixer && strings.HasPrefix(str, "Volume:") {
			volStr := pattern.FindString(str)
			return strconv.Atoi(volStr[:len(volStr)-1])
		}
	}
	return 0, errors.New("No volume found")
}

func changeVolumeCmd(v int, sign string) []string {
	if useAmixer {
		return []string{"amixer", "set", "Master", strconv.Itoa(v) + "%" + sign}
	}
	return []string{"pactl", "--", "set-sink-volume", "@DEFAULT_SINK@", sign + strconv.Itoa(v) + "%"}
}

// IncreaseVolume Increase the volume by the provided value
func IncreaseVolume(v int) (int, error) {
	if outOfRange(v) {
		return 0, errors.New("Value out of range")
	}
	_, err := execCmd(changeVolumeCmd(v, "+"))
	if err != nil {
		return 0, err
	}
	return v, nil
}

// DecreaseVolume Decrese the volume by the provided value
func DecreaseVolume(v int) (int, error) {
	if outOfRange(v) {
		return 0, errors.New("Value out of range")
	}
	_, err := execCmd(changeVolumeCmd(v, "-"))
	if err != nil {
		return 0, err
	}
	return v, nil
}

func outOfRange(val int) bool {
	return val < 0 || val > 100
}
