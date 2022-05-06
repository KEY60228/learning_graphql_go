package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"gql/graph"
	"gql/graph/generated"
	m "gql/infrastructure/mongo"
	"gql/middleware"
	"gql/support"
)

const defaultPort = "8080"

func main() {
	dsn := support.MONGODB_URI
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dsn))
	if err != nil {
		panic(err)
	}
	repository, err := m.NewRepository(client)
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Auth(repository))
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3000"},
		AllowCredentials: true,
		Debug:            true,
		AllowedHeaders:   []string{"*"},
	}).Handler)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{Repo: repository}}))
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return r.Host == "localhost"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Println("connect to http://localhost:8080/ for GraphQL playground")
	log.Fatal(http.ListenAndServe(":8080", router))
}
