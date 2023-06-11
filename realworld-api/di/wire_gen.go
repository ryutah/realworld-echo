// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/google/wire"
	"github.com/ryutah/realworld-echo/realworld-api/api/rest"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xerrorreport"
	"github.com/ryutah/realworld-echo/realworld-api/pkg/xtrace"
	"github.com/ryutah/realworld-echo/realworld-api/usecase"
	"go.opentelemetry.io/otel/sdk/trace"
)

// Injectors from wire.go:

func InitializeLocalRestExecuter(service xerrorreport.Service, version xerrorreport.Version) *rest.Extcuter {
	errorOutputPort := rest.NewErrorOutputPort()
	outputPort := rest.NewGetArticleOutputPort(errorOutputPort)
	errorReporter := xerrorreport.NewErrorReporter(service, version)
	errorHandler := usecase.NewErrorHandler(errorReporter, errorOutputPort)
	article := usecase.NewArticle(outputPort, errorHandler)
	restArticle := rest.NewArticle(article)
	server := rest.NewServer(restArticle)
	sampler := trace.NeverSample()
	initializer := xtrace.NewStdoutTracingInitializer(sampler)
	extcuter := rest.NewExecuter(server, initializer)
	return extcuter
}

func InitializeAppEngineRestExecuter(projectID xtrace.ProjectID, service xerrorreport.Service, version xerrorreport.Version) *rest.Extcuter {
	errorOutputPort := rest.NewErrorOutputPort()
	outputPort := rest.NewGetArticleOutputPort(errorOutputPort)
	errorReporter := xerrorreport.NewErrorReporter(service, version)
	errorHandler := usecase.NewErrorHandler(errorReporter, errorOutputPort)
	article := usecase.NewArticle(outputPort, errorHandler)
	restArticle := rest.NewArticle(article)
	server := rest.NewServer(restArticle)
	sampler := trace.AlwaysSample()
	initializer := xtrace.NewGoogleCloudTracingInitializer(projectID, sampler)
	extcuter := rest.NewExecuter(server, initializer)
	return extcuter
}

func InitializeCloudRunRestExecuter(projectID xtrace.ProjectID, service xerrorreport.Service, version xerrorreport.Version) *rest.Extcuter {
	errorOutputPort := rest.NewErrorOutputPort()
	outputPort := rest.NewGetArticleOutputPort(errorOutputPort)
	errorReporter := xerrorreport.NewErrorReporter(service, version)
	errorHandler := usecase.NewErrorHandler(errorReporter, errorOutputPort)
	article := usecase.NewArticle(outputPort, errorHandler)
	restArticle := rest.NewArticle(article)
	server := rest.NewServer(restArticle)
	sampler := trace.AlwaysSample()
	initializer := xtrace.NewGoogleCloudTracingInitializer(projectID, sampler)
	extcuter := rest.NewExecuter(server, initializer)
	return extcuter
}

// wire.go:

var (
	restSet       = wire.NewSet(rest.NewServer, rest.NewArticle, inputPortSet)
	inputPortSet  = wire.NewSet(usecase.NewArticle, usecase.NewErrorHandler, wire.Bind(new(usecase.GetArticleInputPort), new(*usecase.Article)), outputPortSet)
	outputPortSet = wire.NewSet(rest.NewErrorOutputPort, rest.NewGetArticleOutputPort)
)

var localTraceInitializerSet = wire.NewSet(xtrace.NewStdoutTracingInitializer, trace.NeverSample, xerrorreport.NewErrorReporter, wire.Bind(new(usecase.ErrorReporter), new(*xerrorreport.ErrorReporter)))

var gcpTraceInitializerSet = wire.NewSet(xtrace.NewGoogleCloudTracingInitializer, trace.AlwaysSample, xerrorreport.NewErrorReporter, wire.Bind(new(usecase.ErrorReporter), new(*xerrorreport.ErrorReporter)))