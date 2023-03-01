package goha

/*
func mergef[T comparable](a, b *T) {
	if b != nil {
		*a = *b
	}
}

// TODO: Create per entity attribute structs and copy values into those.
// Use a generic optional package like https://github.com/Southclaws/opt for struct fields.
func (state *State) Merge(newState State) {
	mergef(&state.LastChanged, &newState.LastChanged)
	mergef(&state.LastUpdated, &newState.LastUpdated)
	mergef(&state.State, &newState.State)

	if len(newState.Attributes.RgbColor) > 0 {
		state.Attributes.RgbColor = newState.Attributes.RgbColor
	}

	mergef(&state.Attributes.ColorTemp, &newState.Attributes.ColorTemp)
	mergef(&state.Attributes.SupportedFeatures, &newState.Attributes.SupportedFeatures)

	if len(newState.Attributes.XyColor) > 0 {
		state.Attributes.XyColor = newState.Attributes.XyColor
	}

	mergef(&state.Attributes.Brightness, &newState.Attributes.Brightness)
	mergef(&state.Attributes.BrightnessPct, &newState.Attributes.BrightnessPct)
	mergef(&state.Attributes.WhiteValue, &newState.Attributes.WhiteValue)
	mergef(&state.Attributes.NextDawn, &newState.Attributes.NextDawn)
	mergef(&state.Attributes.NextDusk, &newState.Attributes.NextDusk)
	mergef(&state.Attributes.NextMidnight, &newState.Attributes.NextMidnight)
	mergef(&state.Attributes.NextNoon, &newState.Attributes.NextNoon)
	mergef(&state.Attributes.NextRising, &newState.Attributes.NextRising)
	mergef(&state.Attributes.NextSetting, &newState.Attributes.NextSetting)
	mergef(&state.Attributes.Elevation, &newState.Attributes.Elevation)
	mergef(&state.Attributes.Azimuth, &newState.Attributes.Azimuth)
	mergef(&state.Attributes.Rising, &newState.Attributes.Rising)
	mergef(&state.Attributes.FriendlyName, &newState.Attributes.FriendlyName)
	mergef(&state.Attributes.Source, &newState.Attributes.Source)
	mergef(&state.Attributes.Transition, &newState.Attributes.Transition)

	state.Context = newState.Context
}
*/

func OrZero[T comparable](value *T) T {
	if value != nil {
		return *value
	} else {
		return *new(T)
	}
}

func (state State) String() string {
	if state.State != nil {
		return *state.State
	} else {
		return ""
	}
}
