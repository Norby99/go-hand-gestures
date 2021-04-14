package volume

import "errors"

// SystemAudio represents the system volume
type SystemAudio struct {
	ID             int
	MicVolume      int // 0-100
	IsMicMuted     bool
	SpeakerVolume  int // 0-100
	IsSpeakerMuted bool
}

func (r *SystemAudio) IncreaseSpeaker(inc int) (int, error) {
	if outOfRange(inc) {
		return r.SpeakerVolume, errors.New("")
	}

	return 1, nil
}

func (r *SystemAudio) DecreaseSpeaker(inc int) (int, error) {
	if outOfRange(inc) {
		return r.SpeakerVolume, errors.New("")
	}

	return 1, nil
}

func outOfRange(val int) bool {
	return val < 0 || val > 100
}
