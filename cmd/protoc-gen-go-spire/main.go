package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const (
	pluginsdkPackage = protogen.GoImportPath("github.com/spiffe/spire-plugin-sdk/pluginsdk")
	grpcPackage      = protogen.GoImportPath("google.golang.org/grpc")
)

func main() {
	var flags flag.FlagSet
	var mode = flags.String("mode", "plugin", `generation mode (either "plugin" or "service"`)
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		var isPlugin bool
		switch *mode {
		case "plugin":
			isPlugin = true
		case "service":
		default:
			return fmt.Errorf(`invalid mode %q: expecting either "plugin" or "service"`, *mode)
		}

		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f, isPlugin)
		}
		return nil
	})
}

func generateFile(gen *protogen.Plugin, file *protogen.File, isPlugin bool) *protogen.GeneratedFile {
	if len(file.Services) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix
	if isPlugin {
		filename += "_spire_plugin.pb.go"
	} else {
		filename += "_spire_service.pb.go"
	}

	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-spire. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	for _, service := range file.Services {
		generateServiceBridges(g, service, isPlugin)
	}
	return g
}

func generateServiceBridges(g *protogen.GeneratedFile, service *protogen.Service, isPlugin bool) {
	kind := "Service"
	if isPlugin {
		kind = "Plugin"
	}

	serviceName := service.GoName
	serviceFullName := string(service.Desc.FullName())

	serverIntfName := serviceName + "Server"
	pluginServerCons := serviceName + kind + "Server"
	pluginServerType := unexport(pluginServerCons)
	pluginServerIdent := g.QualifiedGoIdent(pluginsdkPackage.Ident(kind + "Server"))

	clientIntfName := serviceName + "Client"
	pluginClientType := serviceName + kind + "Client"

	g.P()
	g.P("func ", pluginServerCons, "(server ", serverIntfName, ") ", pluginServerIdent, " {")
	g.P("return ", pluginServerType, "{", serverIntfName, ": server}")
	g.P("}")
	g.P()
	g.P("type ", pluginServerType, " struct {")
	g.P(serverIntfName)
	g.P("}")
	if isPlugin {
		g.P()
		g.P("func (s ", pluginServerType, ") Type() string {")
		g.P("return ", strconv.Quote(serviceName))
		g.P("}")
	}
	g.P()
	g.P("func (s ", pluginServerType, ") GRPCServiceName() string {")
	g.P("return ", strconv.Quote(serviceFullName))
	g.P("}")
	g.P()
	g.P("func (s ", pluginServerType, ") RegisterServer(server *", g.QualifiedGoIdent(grpcPackage.Ident("Server")), ") interface{} {")
	g.P("Register", serverIntfName, "(server , s.", serverIntfName, ")")
	g.P("return s.", serverIntfName)
	g.P("}")

	g.P()
	g.P("type ", pluginClientType, " struct {")
	g.P(clientIntfName)
	g.P("}")
	if isPlugin {
		g.P()
		g.P("func (s ", pluginClientType, ") Type() string {")
		g.P("return ", strconv.Quote(serviceName))
		g.P("}")
	}
	g.P()
	g.P("func (c *", pluginClientType, ") IsInitialized() bool {")
	g.P("return c.", clientIntfName, " != nil")
	g.P("}")
	g.P()
	g.P("func (c *", pluginClientType, ") GRPCServiceName() string {")
	g.P("return ", strconv.Quote(serviceFullName))
	g.P("}")
	g.P()
	g.P("func (c *", pluginClientType, ") InitClient(conn ", g.QualifiedGoIdent(grpcPackage.Ident("ClientConnInterface")), ") interface{} {")
	g.P("c.", clientIntfName, " = New", clientIntfName, "(conn)")
	g.P("return c.", clientIntfName)
	g.P("}")
}

func unexport(s string) string {
	return strings.ToLower(s[:1]) + s[1:]
}
