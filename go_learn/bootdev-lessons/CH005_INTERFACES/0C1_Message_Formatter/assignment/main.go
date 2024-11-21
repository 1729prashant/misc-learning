/*
my solution
*/

package main
import "fmt"

type Formatter interface {
	Format() string
}

type PlainText struct {
	message string
}

type Bold struct {
	message string
}

type Code struct {
	message string
}


func (pt PlainText) Format() string {
	return fmt.Sprintf("%s", pt.message)
}

func (b Bold) Format() string {
	return fmt.Sprintf("**%s**", b.message)
}

func (c Code) Format() string {
	return fmt.Sprintf("`%s`", c.message)
}
// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}



/*
provided solution

package main

type Formatter interface {
	Format() string
}

type PlainText struct {
	message string
}

func (p PlainText) Format() string {
	return p.message
}

type Bold struct {
	message string
}

func (b Bold) Format() string {
	return "**" + b.message + "**"
}

type Code struct {
	message string
}

func (c Code) Format() string {
	return "`" + c.message + "`"
}

// Don't Touch below this line

func SendMessage(formatter Formatter) string {
	return formatter.Format() // Adjusted to call Format without an argument
}


*/
