package main

// This is built invisibly by hb proto
// Generate code for hud_services extensions
// https://github.com/golang/protobuf/issues/1260
import (
	"fmt"

	"github.com/dave/jennifer/jen"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		// The type information for all extensions is in the source files,
		// so we need to extract them into a dynamically created protoregistry.Types.
		extTypes := new(protoregistry.Types)
		for _, file := range gen.Files {
			if err := registerAllExtensions(extTypes, file.Desc); err != nil {
				panic(err)
			}
		}

		for _, f := range gen.Files {
			if !f.Generate || len(f.Services) < 1 {
				continue
			}

			if err := generateFile(gen, extTypes, f); err != nil {
				return err
			}
		}
		return nil
	})
}

func generateFile(gen *protogen.Plugin, extTypes *protoregistry.Types, file *protogen.File) error {
	fileJen := jen.NewFile(string(file.GoPackageName))

	for _, service := range file.Services {
		fileJen.Comment(service.GoName)

		// We need to re-marshal the service options with the dynamically loaded extTypes
		serviceOptions := service.Desc.Options().(*descriptorpb.ServiceOptions)
		b, err := proto.Marshal(serviceOptions)
		if err != nil {
			panic(err)
		}
		serviceOptions.Reset()
		err = proto.UnmarshalOptions{Resolver: extTypes}.Unmarshal(b, serviceOptions)
		if err != nil {
			panic(err)
		}

		// Now we can use reflection to pull hud options from the service
		var port uint64
		serviceOptions.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
			if !fd.IsExtension() {
				return true
			}

			switch fd.FullName() {
			case "hud.port":
				port = v.Uint()
			}

			return true
		})

		generateHudServiceDesc(fileJen, service, port)
	}

	filename := file.GeneratedFilenamePrefix + "_hud_services.pb.go"
	return fileJen.Render(gen.NewGeneratedFile(filename, file.GoImportPath))
}

const (
	grpcLib = "google.golang.org/grpc"
)

func generateHudServiceDesc(fileJen *jen.File, service *protogen.Service, port uint64) {
	hudServiceDescName := fmt.Sprintf("%s_HudServiceDesc", service.GoName)
	fileJen.Type().Id(hudServiceDescName).StructFunc(func(g *jen.Group) {
	})

	serviceDescName := fmt.Sprintf("%s_ServiceDesc", service.GoName)
	fileJen.Func().
		Params(jen.Id("h").Id(hudServiceDescName)).
		Id("ServiceDesc").
		Params().
		Params(jen.Op("*").Qual(grpcLib, "ServiceDesc")).
		Block(jen.Return(jen.Op("&").Id(serviceDescName)))

	fileJen.Func().
		Params(jen.Id("h").Id(hudServiceDescName)).
		Id("Port").
		Params().
		Params(jen.Id("uint64")).
		Block(jen.Return(jen.Lit(port)))
}

func registerAllExtensions(extTypes *protoregistry.Types, descs interface {
	Messages() protoreflect.MessageDescriptors
	Extensions() protoreflect.ExtensionDescriptors
}) error {
	// https://github.com/golang/protobuf/issues/1260
	mds := descs.Messages()
	for i := 0; i < mds.Len(); i++ {
		registerAllExtensions(extTypes, mds.Get(i))
	}
	xds := descs.Extensions()
	for i := 0; i < xds.Len(); i++ {
		if err := extTypes.RegisterExtension(dynamicpb.NewExtensionType(xds.Get(i))); err != nil {
			return err
		}
	}
	return nil
}
