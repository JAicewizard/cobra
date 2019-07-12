package cobra

import (
	"strings"

	flag "github.com/spf13/pflag"
)

type (
	//TemplateData is to be used in templating. eg: the usage docs and the documentation generation
	TemplateData struct {
		Runnable                     bool
		HasHelpSubCommands           bool
		HasSubCommands               bool
		HasAvailableSubCommands      bool
		HasAvailableLocalFlags       bool
		HasAvailableInheritedFlags   bool
		HasExample                   bool
		IsAvailableCommand           bool
		IsAdditionalHelpTopicCommand bool
		DisableAutoGenTag            bool
		UseLine                      string
		CommandPath                  string
		Name                         string
		Short                        string
		Long                         string
		Example                      string
		Version                      string
		LocalFlags                   *flag.FlagSet
		InheritedFlags               *flag.FlagSet
		NonInheritedFlags            *flag.FlagSet
		Aliases                      []string
		commands                     []*TemplateData
		command                      *Command
	}
)

func newTemplateData(c *Command) *TemplateData {
	return &TemplateData{
		Runnable:                     c.Run != nil || c.RunE != nil,
		HasHelpSubCommands:           c.hasHelpSubCommands(),
		HasSubCommands:               len(c.commands) > 0,
		HasAvailableSubCommands:      c.hasAvailableSubCommands(),
		HasAvailableLocalFlags:       c.HasAvailableLocalFlags(),
		HasAvailableInheritedFlags:   c.HasAvailableInheritedFlags(),
		HasExample:                   len(c.Example) > 0,
		IsAvailableCommand:           c.isAvailableCommand(),
		IsAdditionalHelpTopicCommand: c.IsAdditionalHelpTopicCommand(),
		DisableAutoGenTag:            c.DisableAutoGenTag,
		UseLine:                      c.useLine(),
		CommandPath:                  c.CommandPath(),
		Name:                         c.Name(),
		Short:                        c.Short,
		Long:                         c.Long,
		Example:                      c.Example,
		Version:                      c.Version,
		LocalFlags:                   c.LocalFlags(),
		InheritedFlags:               c.InheritedFlags(),
		NonInheritedFlags:            c.NonInheritedFlags(),
		Aliases:                      c.Aliases,
		command:                      c,
	}
}

// VisitParents visits all parents of the command and invokes fn on each parent.
func (d *TemplateData) VisitParents(fn func(*Command)) {
	d.command.visitParents(fn)
}

//Parent returns the parent's TemplateData
func (d *TemplateData) Parent() *TemplateData {
	return d.command.parent.TemplateData()
}

//Command returns the original command
func (d *TemplateData) Command() *Command {
	return d.command
}

// Commands returns a sorted slice of child commands.
func (d *TemplateData) Commands() []*TemplateData {
	if len(d.commands) != 0 {
		return d.commands
	}

	commands := d.command.Commands()
	d.commands = make([]*TemplateData, len(commands))
	for k, v := range commands {
		d.commands[k] = newTemplateData(v)
	}
	return d.commands
}

// HasParent returns usage string.
func (d *TemplateData) HasParent() bool {
	return d.command.hasParent()
}

// UsageString returns usage string.
func (d *TemplateData) UsageString() string {
	return d.command.usageString()
}

// NameAndAliases returns a list of the command name and all aliases
func (d *TemplateData) NameAndAliases() string {
	return strings.Join(append([]string{d.command.Name()}, d.command.Aliases...), ", ")
}

var minUsagePadding = 25

// UsagePadding return padding for the usage.
// Should only be used inside a command function or template.
func (d *TemplateData) UsagePadding() int {
	if d.command.parent == nil {
		return minUsagePadding
	}
	len := d.command.parent.getMaxUsageLength()
	if len < minUsagePadding {
		return minUsagePadding
	}
	return len
}

var minCommandPathPadding = 11

// CommandPathPadding return padding for the command path.
// Should only be used inside a command function or template.
func (d *TemplateData) CommandPathPadding() int {
	if d.command.parent == nil {
		return minCommandPathPadding
	}
	len := d.command.parent.getMaxCommandPathLength()
	if len < minCommandPathPadding {
		return minCommandPathPadding
	}
	return len
}

var minNamePadding = 11

// NamePadding returns padding for the name.
// Should only be used inside a command function or template.
func (d *TemplateData) NamePadding() int {
	if d.command.parent == nil {
		return minNamePadding
	}
	len := d.command.parent.getMaxNameLength()
	if len < minNamePadding {
		return minNamePadding
	}
	return len
}
