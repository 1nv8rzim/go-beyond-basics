//go:build linux && !prod

package build

func OS() string {
	return "Linux Dev"
}
