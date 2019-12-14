package gbridge

import (
	"github.com/pborges/gbridge/proto"
)

// Command belongs to a device and handles the executing of functions. 
type Command interface {
	Name() string
	// passing maps in functions feels dirty
	Execute(ctx Context, args map[string]interface{}) proto.DeviceError
}

type OnOffCommand func(ctx Context, state bool) proto.DeviceError

// Execute checks if the arguments from the Intent Request are correct and passes them to the user defined execute handlers.
func (t OnOffCommand) Execute(ctx Context, args map[string]interface{}) proto.DeviceError {
	if val, ok := args["on"]; ok {
		if state, ok := val.(bool); ok {
			return t(ctx, state)
		}
		return proto.ErrorCodeNotSupported
	}
	return proto.ErrorCodeProtocolError
}

// Name returns the Identifier String of the Command
func (t OnOffCommand) Name() string {
	return "action.devices.commands.OnOff"
}

// DirectionalOpenCloseCommand defines how a function should handle this specific trait
type DirectionalOpenCloseCommand func(ctx Context, openPercent float64, openDirection OpenCloseTraitDirection) proto.DeviceError

// Execute validates the request and calls the user defined handler for the device trait
func (t DirectionalOpenCloseCommand) Execute(ctx Context, args map[string]interface{}) proto.DeviceError {
	// validate if our EXECUTE Request contains the "openPercent" object so the handler can actually handle something
	openDirection := OpenCloseTraitDirectionNone
	if argOpenPercent, ok := args["openPercent"]; ok {
		if openPercent, ok := argOpenPercent.(float64); ok {
			// if openDirection is specified, set it, otherwise lets use none
			if argDir, ok := args["openDirection"]; ok {
				if dir, ok := argDir.(string); ok {
					openDirection = OpenCloseTraitDirection(dir)
				}
			}
			return t(ctx, openPercent, openDirection)
		}
		return proto.ErrorCodeNotSupported
	}
	return proto.ErrorCodeProtocolError
}

// Name returns the intent for this Trait
func (t DirectionalOpenCloseCommand) Name() string {
	return "action.devices.commands.OpenClose"
}

// OpenCloseCommand defines how a function should handle this specific trait
type OpenCloseCommand func(ctx Context, openPercent float64) proto.DeviceError

// Execute validates the request and calls the user defined handler for the device trait
func (t OpenCloseCommand) Execute(ctx Context, args map[string]interface{}) proto.DeviceError {
	// validate if our EXECUTE Request contains the "openPercent" object so the handler can actually handle something
	if argOpenPercent, ok := args["openPercent"]; ok {
		if openPercent, ok := argOpenPercent.(float64); ok {
			return t(ctx, openPercent)
		}
		return proto.ErrorCodeNotSupported
	}
	return proto.ErrorCodeProtocolError
}

// Name returns the intent for this Trait
func (t OpenCloseCommand) Name() string {
	return "action.devices.commands.OpenClose"
}
