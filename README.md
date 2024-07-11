# Go Dispatch

Go Dispatch is a library for dispatching system commands and http handlers.

## Using Go Dispatch

Simply add Go Dispatch to your project by importing it.

`import "dispatch github.com/clinto-bean/go-dispatch"`

If you want to run a command, simply pass its name to Command().

`dispatch.Command("help")`

If you are wanting to execute a handler, pass the name of the handler and its arguments to Handle().

`dispatch.Handle(handlerGet, {args})`

## Author & License

Original maintainer [clinto-bean](github.com/clinto-bean). License information can be found in [LICENSE](LICENSE) file.
