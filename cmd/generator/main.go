package main

import (
	"books-api/utils"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
	"html/template"
	"log"
	"os"
	"slices"
	"strings"
)

var (
	flags           = flag.NewFlagSet("generator", flag.ExitOnError)
	generatorType   = flags.String("type", "", "Type of file(s) to generate: model|controller|scaffold")
	generatorName   = flags.String("name", "", "Name of model or resource or controller or scaffold")
	generatorFields = flags.String("fields", "", "Fields for model: name:type,name:type,...")
	pluralizeClient = pluralize.NewClient()
)

type ResourceInfo struct {
	ResourceName   string
	SingleInstance string
	PluralInstance string
	ShortcutName   string
	Properties     []ModelProperty
}

type ModelProperty struct {
	Name string
	Type string
	Json string
}

func main() {
	if len(os.Args) == 1 {
		log.Fatalf("Type and name are required. Use --help for more info.")
	}
	flags.Parse(os.Args[1:])
	Generator()
}

func Generator() {
	rootPath := rootPath()
	if *generatorName != "" {
		fields := strings.Split(*generatorName, ",")
		for _, field := range fields {
			resourceInfo := ResourceInfo{
				ResourceName:   pluralizeClient.Singular(strcase.ToCamel(field)),
				SingleInstance: pluralizeClient.Singular(strcase.ToSnake(field)),
				PluralInstance: pluralizeClient.Plural(strcase.ToSnake(field)),
			}
			switch *generatorType {
			case "controller", "c":
				generateController(resourceInfo, rootPath)
			case "model", "m":
				generateModel(resourceInfo, rootPath)
			case "scaffold", "sc":
				generateScaffold(resourceInfo, rootPath)
			case "service", "s":
				generateService(resourceInfo, rootPath)
			default:
				log.Fatalf("Wrong Resource has Been Send: %s", os.Args[1])
			}
		}
	}
}

func rootPath() string {
	rootPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return rootPath
}

func generateController(resourceInfo ResourceInfo, rootPath string) {
	templateFilePath := fmt.Sprintf("%s/templates/controller.tmpl", rootPath)
	controllerDirectory := fmt.Sprintf("%s/app/controllers", rootPath)
	controllerName := strings.ReplaceAll(strings.ReplaceAll(utils.RemoveSpecialChars(strings.ToLower(resourceInfo.PluralInstance)), "controllers", ""), "controller", "")
	newFilePath := fmt.Sprintf("%s/%s.go", controllerDirectory, controllerName+"_controllers")
	resourceInfo.ResourceName = pluralizeClient.Plural(strcase.ToCamel(controllerName))
	resourceInfo.ShortcutName = string(controllerName[0]) + "c"
	fileGenerator(newFilePath, templateFilePath, &resourceInfo)
}

func generateModel(resourceInfo ResourceInfo, rootPath string) {
	templateFilePath := fmt.Sprintf("%s/templates/model.tmpl", rootPath)
	modelDirectory := fmt.Sprintf("%s/app/models", rootPath)
	newFilePath := fmt.Sprintf("%s/%s.go", modelDirectory, resourceInfo.SingleInstance)
	generateModelFields(&resourceInfo)
	fileGenerator(newFilePath, templateFilePath, &resourceInfo)
}

func generateService(resourceInfo ResourceInfo, rootPath string) {
	templateFilePath := fmt.Sprintf("%s/templates/service.tmpl", rootPath)
	serviceDirectory := fmt.Sprintf("%s/app/services", rootPath)
	serviceName := strings.ReplaceAll(strings.ReplaceAll(utils.RemoveSpecialChars(strings.ToLower(resourceInfo.PluralInstance)), "services", ""), "service", "")
	newFilePath := fmt.Sprintf("%s/%s.go", serviceDirectory, serviceName+"_services")
	resourceInfo.ResourceName = pluralizeClient.Plural(strcase.ToCamel(serviceName))
	fileGenerator(newFilePath, templateFilePath, &resourceInfo)
}

func generateScaffold(resourceInfo ResourceInfo, rootPath string) {
	generateController(resourceInfo, rootPath)
	generateModel(resourceInfo, rootPath)
}

func generateModelFields(resourceInfo *ResourceInfo) {
	if *generatorFields != "" {
		props := []ModelProperty{}
		fields := strings.Split(*generatorFields, ",")
		for _, field := range fields {
			pieces2 := strings.Split(field, ":")
			if len(pieces2) < 2 {
				log.Fatal("Each property needs a name and a type: ", field)
				continue
			}
			propType := pieces2[1]
			validTypes := []string{"string", "bool", "int8", "uint8", "int16", "uint16", "int32", "uint32", "int64", "uint64",
				"int", "uint", "uintptr", "float32", "float64", "complex64", "complex128"}
			if !slices.Contains(validTypes, propType) {
				log.Fatal("Invalid property type: ", field)
				continue
			}

			props = append(props, ModelProperty{
				strcase.ToCamel(pieces2[0]),
				pieces2[1],
				strcase.ToSnake(pieces2[0]),
			})
		}
		resourceInfo.Properties = props
	}
}

func fileGenerator(newFilePath string, templateFilePath string, resourceInfo *ResourceInfo) {
	if _, err := os.Stat(newFilePath); errors.Is(err, os.ErrNotExist) {
		tmpl, err := template.ParseFiles(templateFilePath)
		if err != nil {
			log.Printf("Error using template: %s", templateFilePath)
			log.Fatal(err)
		}
		writer := new(bytes.Buffer)
		err = tmpl.Execute(writer, resourceInfo)
		if err != nil {
			log.Printf("Error executing template: %s", templateFilePath)
			log.Fatal(err)
		}
		err = os.WriteFile(newFilePath, writer.Bytes(), 0644)
		if err != nil {
			log.Printf("Error writing file: %s", newFilePath)
			log.Fatal(err)
		}
	} else {
		log.Printf("File already exists: %s", newFilePath)
	}
}
