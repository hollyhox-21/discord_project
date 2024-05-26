package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/swaggest/swgui/v5emb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/hollyhox-21/discord_project/search_service/internal/app"
	pb "github.com/hollyhox-21/discord_project/search_service/pkg/search"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// init repo

	// Create a listener on TCP ports
	listenerHttp, listenerGrpc, listenerInfo := CreateListeners()

	// Create Implementation
	impl, err := app.NewImplementation()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	// Create a gRPC server object
	grpcSrv := grpc.NewServer()
	pb.RegisterSearchServiceServer(grpcSrv, impl)
	reflection.Register(grpcSrv)

	// Create an HTTP server object
	mux := runtime.NewServeMux()
	if err = pb.RegisterSearchServiceHandlerServer(ctx, mux, impl); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	httpRouter := RegisterRouter(mux)

	httpSrv := &http.Server{
		Handler: httpRouter,
	}

	// Create info server
	infoRouter := chi.NewRouter()

	swaggerRouter := initSwagger()
	infoRouter.Mount("/docs", swaggerRouter)

	infoSrv := &http.Server{
		Handler: infoRouter,
	}

	// Start the gRPC server
	go func() {
		log.Printf("gRPC server listening at %v", listenerGrpc.Addr())
		if err = grpcSrv.Serve(listenerGrpc); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Start the HTTP server
	go func() {
		log.Printf("HTTP server listening at %v", listenerHttp.Addr())
		if err = httpSrv.Serve(listenerHttp); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Start the info server
	go func() {
		log.Printf("INFO server listening at %v", listenerInfo.Addr())
		if err = infoSrv.Serve(listenerInfo); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	<-sigs

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		grpcSrv.GracefulStop()
		log.Fatalf("failed to gracefully stop the gRPC server: %v", err)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := httpSrv.Shutdown(ctx); err != nil {
			log.Fatalf("failed to gracefully stop the HTTP server: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := infoSrv.Shutdown(ctx); err != nil {
			log.Fatalf("failed to gracefully stop the INFO server: %v", err)
		}
	}()

	wg.Wait()

	log.Println("Server gracefully stopped")
}

func CreateListeners() (net.Listener, net.Listener, net.Listener) {
	listenerHttp, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	listenerGrpc, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	listenerInfo, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	return listenerHttp, listenerGrpc, listenerInfo
}

func RegisterRouter(mux *runtime.ServeMux) *chi.Mux {
	r := chi.NewRouter()

	initProbes(r)

	r.Mount("/", mux)

	return r
}

func initProbes(r *chi.Mux) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	r.Get("/ready", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
}

func initSwagger() http.Handler {
	r := chi.NewRouter()

	log.Println("swagger init")

	r.Mount("/", v5emb.NewHandler("API Definition",
		"/docs/search.swagger.json",
		"/docs"))

	r.Get("/search.swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "pkg/search/search.swagger.json")
	})

	return r
}
