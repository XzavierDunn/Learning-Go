package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
    "os"
    "os/signal"
    "syscall"
    "context"

    "gopkg.in/natefinch/lumberjack.v2"
)

func handler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    fmt.Println(query)
    name := query.Get("name")
    if name == "" {
        w.Write([]byte(fmt.Sprintf("Hello%s\n", name)))
    } else {
        w.Write([]byte(fmt.Sprintf("Hello, %s\n", name)))
    }
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handler)

    srv := &http.Server{
        Handler: mux,
        Addr: ":8080",
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 1 * time.Second,
    }

    // Configure Logging
    LOG_FILE_LOC := os.Getenv("LOG_FILE_LOC")
    if LOG_FILE_LOC != "" {
        log.SetOutput(&lumberjack.Logger{
            Filename: LOG_FILE_LOC,
            MaxSize: 500,
            MaxBackups: 3,
            MaxAge: 28,
            Compress: true,
        })
    }

    go func() {
        log.Println("Listening... http://localhost:8080")
        if err := srv.ListenAndServe(); err != nil {
            log.Fatal(err)
        }
    }()

    // Graceful shutdown
    waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
    interruptChan := make(chan os.Signal, 1)
    signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

    // Block until we recieve our signal
    <-interruptChan

    // Create a deadline to wait for
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
    defer cancel()
    srv.Shutdown(ctx)

    log.Println("Stopping")
    os.Exit(0)
}
