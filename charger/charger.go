package charger

import (
	"github.com/andig/evcc/api"
	"github.com/andig/evcc/provider"
)

// Charger is an api.Charger implementation with configurable getters and setters.
type Charger struct {
	statusG     provider.StringGetter
	enabledG    provider.BoolGetter
	enableS     provider.BoolSetter
	maxCurrentS provider.IntSetter
}

// NewConfigurableFromConfig creates a new configurable charger
func NewConfigurableFromConfig(log *api.Logger, other map[string]interface{}) api.Charger {
	cc := struct{ Status, Enable, Enabled, MaxCurrent *provider.Config }{}
	api.DecodeOther(log, other, &cc)

	charger := NewConfigurable(
		provider.NewStringGetterFromConfig(cc.Status),
		provider.NewBoolGetterFromConfig(cc.Enabled),
		provider.NewBoolSetterFromConfig("enable", cc.Enable),
		provider.NewIntSetterFromConfig("current", cc.MaxCurrent),
	)

	return charger
}

// NewConfigurable creates a new charger
func NewConfigurable(
	statusG provider.StringGetter,
	enabledG provider.BoolGetter,
	enableS provider.BoolSetter,
	maxCurrentS provider.IntSetter,
) api.Charger {
	return &Charger{
		statusG:     statusG,
		enabledG:    enabledG,
		enableS:     enableS,
		maxCurrentS: maxCurrentS,
	}
}

// Status implements the Charger.Status interface
func (m *Charger) Status() (api.ChargeStatus, error) {
	s, err := m.statusG()
	if err != nil {
		return api.StatusNone, err
	}

	return api.ChargeStatus(s), nil
}

// Enabled implements the Charger.Enabled interface
func (m *Charger) Enabled() (bool, error) {
	return m.enabledG()
}

// Enable implements the Charger.Enable interface
func (m *Charger) Enable(enable bool) error {
	return m.enableS(enable)
}

// MaxCurrent implements the Charger.MaxCurrent API
func (m *Charger) MaxCurrent(current int64) error {
	return m.maxCurrentS(current)
}
