package main

import (
	"fmt"
	"github.com/dave/jennifer/jen"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"regexp"
	"strings"
	"time"
)

// These types should already be defined in the package
const (
	LogImportPath        = "github.com/edison-moreland/nreal-hud/sdk/log"
	MessageHeaderType    = "MessageHeader"
	ProxyType            = "Proxy"
	ErrInvalidOpType     = "ErrInvalidOp"
	UnmarshallInt32Func  = "UnmarshallInt32"
	UnmarshallUint32Func = "UnmarshallUint32"
	UnmarshallStringFunc = "UnmarshallString"
	UnmarshallArrayFunc  = "UnmarshallArray"
	UnmarshallFdFunc     = "UnmarshallFd"
	UnmarshallFixedFunc  = "UnmarshallFixed"
)

var toTitle = cases.Title(language.English)

func snakeToCamel(identifier string) string {
	components := strings.Split(identifier, "_")

	var camelIdentifier strings.Builder
	for _, c := range components {
		camelIdentifier.WriteString(toTitle.String(c))
	}

	return camelIdentifier.String()
}

func identifier(parts ...string) string {
	var id strings.Builder
	for _, p := range parts {
		id.WriteString(snakeToCamel(p))
	}

	return id.String()
}

type commenter interface {
	Commentf(format string, a ...interface{}) *jen.Statement
	Comment(str string) *jen.Statement
}

func typeDoc(c commenter, identifier, summary, description string) {
	if summary == "" {
		return
	}

	summary = strings.TrimPrefix(summary, strings.ToLower(identifier))

	c.Commentf("%s %s", identifier, summary)

	if description != "" {
		// The description usually has a bunch of extra whitespace that needs to be cleaned up
		re := regexp.MustCompile(`\n[\t ]*`)
		c.Comment("\t" + re.ReplaceAllString(strings.TrimSpace(description), "\n\t"))
	}
}

func genProtocol(p XMLProtocol, packageName string) *jen.File {
	f := jen.NewFile(packageName)

	f.Commentf(`
	DO NOT MODIFY!
	This file was auto-generated on %s
	protocol: %s
	`, time.Now(), p.Name).Line()

	for _, iface := range p.Interfaces {
		genInterface(f, iface)
	}

	return f
}

func genInterface(f *jen.File, i XMLInterface) {
	// Each wayland interface gets a:
	//   - Proxy struct (maybe just a type alias?)
	//   - Methods on the struct for each request
	//   - A listener interface type
	//   - All enums

	// Type
	interfaceIdentifier := identifier(i.Name)
	typeDoc(f, interfaceIdentifier, i.Description.Summary, i.Description.Value)
	f.Type().Id(interfaceIdentifier).StructFunc(func(g *jen.Group) {
		g.Id(ProxyType)
	})
	f.Func().
		Id(identifier("New", i.Name)).
		Params(jen.Id("p").Id(ProxyType)).
		Params(jen.Id(interfaceIdentifier)).
		Block(
			jen.Return(jen.Id(interfaceIdentifier).Values(jen.Dict{
				jen.Id(ProxyType): jen.Id("p"),
			})),
		)

	eventListenerIdentifier := identifier(i.Name, "Listener")
	f.Func().
		Params(jen.Id("w").Op("*").Id(interfaceIdentifier)).
		Id("AttachListener").
		Params(jen.Id("l").Id(eventListenerIdentifier)).
		Block(
			jen.Id("w").Dot(ProxyType).
				Dot("Dispatcher").Call().
				Dot("AttachListener").Call(jen.Id("l")),
		)

	genEnums(f, i)
	genRequests(f, i)
	genEvent(f, i)
}

func genEnums(f *jen.File, i XMLInterface) {
	for _, e := range i.Enums {
		enumIdentifier := identifier(i.Name, e.Name)
		typeDoc(f, enumIdentifier, e.Description.Summary, e.Description.Value)
		f.Type().Id(enumIdentifier).Uint32()
		f.Const().DefsFunc(func(g *jen.Group) {
			for _, c := range e.Entries {
				enumCaseIdentifier := fmt.Sprintf("%s_%s", enumIdentifier, snakeToCamel(c.Name))

				if c.Summary != "" {
					g.Commentf("%s %s", enumCaseIdentifier, c.Summary)
				}

				g.Id(enumCaseIdentifier).Id(enumIdentifier).Op("=").Id(c.Value)
			}
		})
	}
}

func genEvent(f *jen.File, i XMLInterface) {
	// Create event structs
	for _, e := range i.Events {
		eventIdentifier := identifier(i.Name, e.Name, "Event")
		typeDoc(f, eventIdentifier, e.Description.Summary, e.Description.Value)
		f.Type().Id(eventIdentifier).StructFunc(func(g *jen.Group) {
			genArguments(g, i, e.Arguments)
		})

		f.Func().
			Params(jen.Id("e").Op("*").Id(eventIdentifier)).
			Id("UnmarshallBinary").
			Params(jen.Id("d").Index().Id("byte")).
			Params(jen.Id("error")).
			BlockFunc(func(g *jen.Group) {
				if len(e.Arguments) == 0 {
					g.Return(jen.Id("nil"))
					return
				}

				g.Id("offset").Op(":=").Lit(0)
				for _, a := range e.Arguments {
					g.List(
						jen.Id("offset"),
						jen.Id("e").Dot(identifier(a.Name)),
					).Op("=").Do(func(s *jen.Statement) {
						switch a.Type {
						case "int":
							s.Id(UnmarshallInt32Func)
						case "uint", "object", "new_id", "enum":
							s.Id(UnmarshallUint32Func)
						case "string":
							s.Id(UnmarshallStringFunc)
						case "array":
							s.Id(UnmarshallArrayFunc)
						case "fixed":
							s.Id(UnmarshallFixedFunc)
						case "fd":
							s.Id(UnmarshallFdFunc)
						}
					}).Call(jen.Id("offset"), jen.Id("d"))
				}

				g.Return(jen.Id("nil"))
			})
	}

	// Create event listener interface
	eventListenerIdentifier := identifier(i.Name, "Listener")
	f.Type().Id(eventListenerIdentifier).InterfaceFunc(func(g *jen.Group) {
		for _, e := range i.Events {
			eventHandlerIdentifier := identifier(e.Name)
			eventIdentifier := identifier(i.Name, e.Name, "Event")
			g.Id(eventHandlerIdentifier).ParamsFunc(func(g *jen.Group) {
				g.Id(eventIdentifier)
			})
		}
	})

	// Create unimplemented event listener
	unimplementedEventListenerIdentifier := identifier("Unimplemented", i.Name, "Listener")
	f.Type().Id(unimplementedEventListenerIdentifier).Struct()
	for _, e := range i.Events {
		eventHandlerIdentifier := identifier(e.Name)
		eventIdentifier := identifier(i.Name, e.Name, "Event")
		f.Func().
			Params(jen.Id("e").Op("*").Id(unimplementedEventListenerIdentifier)).
			Id(eventHandlerIdentifier).
			Params(jen.Id("_").Id(eventIdentifier)).
			Block(jen.Return())
	}

	// Create event dispatcher
	eventDispatcherIdentifier := identifier(i.Name, "Dispatcher")
	f.Type().Id(eventDispatcherIdentifier).Struct(
		jen.Id(eventListenerIdentifier),
	)

	newEventDispatcherIdentifier := identifier("New", i.Name, "Dispatcher")
	f.Func().
		Id(newEventDispatcherIdentifier).
		Params().
		Params(jen.Op("*").Id(eventDispatcherIdentifier)).
		Block(
			jen.Return(jen.Op("&").Id(eventDispatcherIdentifier).Values(jen.Dict{
				jen.Id(eventListenerIdentifier): jen.Op("&").Id(unimplementedEventListenerIdentifier).Values(),
			})),
		)

	f.Func().
		Params(jen.Id("i").Op("*").Id(eventDispatcherIdentifier)).
		Id("Dispatch").
		Params(
			jen.Id("h").Id(MessageHeaderType),
			jen.Id("b").Index().Id("byte")).
		Params(jen.Id("error")).
		BlockFunc(func(g *jen.Group) {
			g.Switch(jen.Id("h").Dot("Opcode")).BlockFunc(func(g *jen.Group) {
				for op, e := range i.Events {
					eventIdentifier := identifier(i.Name, e.Name, "Event")
					eventHandlerIdentifier := identifier(e.Name)
					g.Case(jen.Lit(op)).Block(
						jen.Var().Id("e").Id(eventIdentifier),
						jen.Id("err").Op(":=").Id("e").Dot("UnmarshallBinary").Call(jen.Id("b")),
						jen.If(jen.Id("err").Op("!=").Id("nil")).Block(
							jen.Return(jen.Id("err")),
						),
						jen.Id("i").Dot(eventHandlerIdentifier).Call(jen.Id("e")),
					)
				}

				g.Default().Block(
					jen.Return(jen.Id(ErrInvalidOpType)),
				)
			})

			g.Return(jen.Id("nil"))
		})

	f.Func().
		Params(jen.Id("i").Op("*").Id(eventDispatcherIdentifier)).
		Id("AttachListener").
		Params(jen.Id("l").Id("interface").Values()).
		BlockFunc(func(g *jen.Group) {
			// Assert
			g.List(jen.Id("listener"), jen.Id("ok")).
				Op(":=").
				Id("l").Dot("").Parens(jen.Id(eventListenerIdentifier))

			g.If(jen.Op("!").Id("ok")).Block(
				jen.Qual(LogImportPath, "Panic").Call().
					Dot("Msg").Call(jen.Lit("listener is of wrong type!")),
			)

			g.Id("i").Dot(eventListenerIdentifier).Op("=").Id("listener")
		})
}

func genRequests(f *jen.File, i XMLInterface) {
	for _, r := range i.Requests {
		requestIdentifier := identifier(i.Name, r.Name, "Request")
		typeDoc(f, requestIdentifier, r.Description.Summary, r.Description.Value)
		f.Type().Id(requestIdentifier).StructFunc(func(g *jen.Group) {
			genArguments(g, i, r.Arguments)
		})

		//requestOpcodeIdentifier := identifier(i.Name, r.Name, "Request", "Op")
		//f.Const().Id(requestOpcodeIdentifier).Op("=").Lit(op)
	}
}

func genArguments(g *jen.Group, i XMLInterface, arguments []XMLArgument) {
	for _, a := range arguments {
		argumentIdentifier := snakeToCamel(a.Name)
		typeDoc(g, argumentIdentifier, a.Summary, "")
		g.Id(argumentIdentifier).
			Add(argumentType(i.Name, a)).
			Tag(map[string]string{"wayland": a.Type})
	}
}

func argumentType(parent string, a XMLArgument) *jen.Statement {
	var argType *jen.Statement

	switch a.Type {
	case "int":
		argType = jen.Int32()
	case "uint":
		argType = jen.Uint32()
		// TODO: Auto resolve enum?
		//if a.Enum != "" {
		//	components := strings.Split(a.Enum, ".")
		//	if len(components) == 1 {
		//		components = append([]string{parent}, components...)
		//	}
		//
		//	argType = jen.Id(identifier(components...))
		//} else {
		//	argType = jen.Uint32()
		//}
	case "fixed":
		argType = jen.Float64()
	case "object", "new_id":
		argType = jen.Uint32()
		// TODO: Auto resolve objects?
		//if a.Interface == "" {
		//	argType = jen.Uint32()
		//} else {
		//	argType = jen.Id(snakeToCamel(a.Interface))
		//}
	case "fd":
		argType = jen.Uintptr()
	case "string":
		argType = jen.String()
	case "array":
		argType = jen.Index().Byte()
	default:
		log.Panic().
			Str("interface", parent).
			Str("argument", a.Name).
			Str("argument_type", a.Type).
			Msg("Unknown argument type!")
	}

	if argType == nil {
		log.Panic().
			Str("interface", parent).
			Msg("This should not happen.")
	}

	return argType
}
