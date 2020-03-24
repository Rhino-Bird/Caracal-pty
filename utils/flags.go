package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/Rhino-Bird/structs"
	"github.com/urfave/cli/v2"

	"github.com/rhino-bird/caracal-pty/globals"
)

// GenerateFlags Generate command line flags.
func GenerateFlags(options ...interface{}) ([]cli.Flag, map[string]string) {
	mappings := make(map[string]string)
	flags := make([]cli.Flag, 0, 10)

	for _, strt := range options {
		o := structs.New(strt)
		for _, fld := range o.Fields() {
			flagName := fld.Tag("flagName")
			if flagName == "" {
				continue
			}

			envName := globals.ProcessName + "_" + strings.ToUpper(strings.Join(strings.Split(flagName, "-"), "_"))
			mappings[flagName] = fld.Name()

			flagShortName := fld.Tag("flagSName")
			if flagShortName != "" {
				flagName += ", " + flagShortName
			}

			flagDescription := fld.Tag("flagDescribe")

			switch fld.Kind() {
			case reflect.String:
				flags = append(flags, &cli.StringFlag{
					Name:    flagName,
					Value:   fld.Value().(string),
					Usage:   flagDescription,
					EnvVars: []string{envName},
				})
			case reflect.Bool:
				flags = append(flags, &cli.BoolFlag{
					Name:    flagName,
					Usage:   flagDescription,
					EnvVars: []string{envName},
				})
			case reflect.Int:
				flags = append(flags, &cli.IntFlag{
					Name:    flagName,
					Value:   fld.Value().(int),
					Usage:   flagDescription,
					EnvVars: []string{envName},
				})
			}
		}
	}

	return flags, mappings
}

// ApplyDefaultValues set the default value if no parameters are passed.
func ApplyDefaultValues(strt interface{}) (err error) {
	o := structs.New(strt)

	for _, fld := range o.Fields() {
		dv := fld.Tag("default")
		if dv == "" {
			continue
		}

		var val interface{}
		switch fld.Kind() {
		case reflect.String:
			val = dv
		case reflect.Bool:
			if dv == "true" || dv == "false" {
				val, _ = strconv.ParseBool(dv)
			} else {
				return fmt.Errorf("invalid bool expression: %v, use true/false", dv)
			}
		case reflect.Int:
			val, err = strconv.Atoi(dv)
			if err != nil {
				return err
			}
		default:
			val = fld.Value()
		}

		fld.Set(val)
	}
	return nil
}
