package goha

func mergef[T comparable](a, b *T) {
	if *b != *new(T) {
		*a = *b
	}
}

// TODO: Create per entity attribute structs and copy values into those.
// Use pointers for message structure and use mergef to determine which fields to copy over.
// Each entity type will then implement a merge function for its attributes.
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

	state.Context = newState.Context
}
