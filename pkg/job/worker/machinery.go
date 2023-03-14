package worker

import (
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	"github.com/RichardKnop/machinery/v2/example/tracers"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	"github.com/RichardKnop/machinery/v2/log"
	"github.com/RichardKnop/machinery/v2/tasks"
	"time"
)

type Config struct {
	BrokerDNS    string `yaml:"broker_dns"`
	BackendDNS   string `yaml:"backend_dns"`
	DefaultQueue string `yaml:"default_queue"`
}

type Server struct {
	*machinery.Server
}

func (server *Server) RegisterTask(name string, taskFunc interface{}) error {
	return server.Server.RegisterTask(name, taskFunc)
}

type Signature struct {
	UUID           string
	Name           string
	RoutingKey     string
	ETA            *time.Time
	GroupUUID      string
	GroupTaskCount int
	Args           []Arg
	Headers        Headers
	Priority       uint8
	Immutable      bool
	RetryCount     int
	RetryTimeout   int
}

func (server *Server) RegisterPeriodicTask(spec, name string, signature *Signature) error {
	var args []tasks.Arg
	for _, arg := range signature.Args {
		args = append(args, tasks.Arg{
			Name:  arg.Name,
			Type:  arg.Type,
			Value: arg.Value,
		})
	}
	return server.Server.RegisterPeriodicTask(spec, name, &tasks.Signature{
		UUID:           signature.UUID,
		Name:           signature.Name,
		RoutingKey:     signature.RoutingKey,
		ETA:            signature.ETA,
		GroupUUID:      signature.GroupUUID,
		GroupTaskCount: signature.GroupTaskCount,
		Args:           args,
		Headers:        signature.Headers.Headers,
		Priority:       signature.Priority,
		Immutable:      signature.Immutable,
		RetryCount:     signature.RetryCount,
		RetryTimeout:   signature.RetryTimeout,
	})
}

type Arg struct {
	Name  string      `bson:"name"`
	Type  string      `bson:"type"`
	Value interface{} `bson:"value"`
}

type Headers struct {
	tasks.Headers
}

func (server *Server) RunWorker(conf *Config) error {
	_ = conf
	consumerTag := "job_worker"

	cleanup, err := tracers.SetupTracer(consumerTag)
	if err != nil {
		log.FATAL.Fatalln("Unable to instantiate a tracer:", err)
	}
	defer cleanup()

	// The second argument is a consumer tag
	// Ideally, each worker should have a unique tag (worker1, worker2 etc)
	worker := server.NewWorker(consumerTag, 0)

	// Here we inject some custom code for error handling,
	// start and end of task hooks, useful for metrics for example.
	errorHandler := func(err error) {
		log.ERROR.Println("I am an error handler:", err)
	}

	preTaskHandler := func(signature *tasks.Signature) {
		log.INFO.Println("I am a start of task handler for:", signature.Name)
	}

	postTaskHandler := func(signature *tasks.Signature) {
		log.INFO.Println("I am an end of task handler for:", signature.Name)
	}

	worker.SetPostTaskHandler(postTaskHandler)
	worker.SetErrorHandler(errorHandler)
	worker.SetPreTaskHandler(preTaskHandler)

	return worker.Launch()
}

func NewServer(conf *Config) *Server {
	cnf := &config.Config{
		DefaultQueue:    conf.DefaultQueue,
		ResultsExpireIn: 3600,
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			IdleTimeout:            240,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
		},
	}

	// Create worker instance
	broker := redisbroker.NewGR(cnf, []string{conf.BrokerDNS}, 0)
	backend := redisbackend.NewGR(cnf, []string{conf.BackendDNS}, 0)
	lock := eagerlock.New()
	server := machinery.NewServer(cnf, broker, backend, lock)

	return &Server{server}
}
