package pattern

type MobileAlertStater interface {
	Alert() string
}

type MobileAlert struct {
	state MobileAlertStater
}

func (a *MobileAlert) Alert() string {
	return a.state.Alert()
}

func (a *MobileAlert) SetState(state MobileAlertStater) {
	a.state = state
}

func NewMobileAlert() *MobileAlert {
	return &MobileAlert{state: &MobileAlertVibration{}}
}

type MobileAlertVibration struct {
}

func (a *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr... Vrrr..."
}

type MobileAlertSong struct {
}

func (a *MobileAlertSong) Alert() string {
	return "No loyalty No royalty..."
}
