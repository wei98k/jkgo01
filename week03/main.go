package main

import (
   "fmt" 
   "log"
   "net/http"
   "context"
   "os"
   "os/signal"
   "syscall"

   "github.com/pkg/errors"
   "golang.org/x/sync/errgroup"
)

func main() {
    eg, ctx := errgroup.WithContext(context.Background())
    
    eg.Go(func() error {
        return serveApp(ctx, 8080)
    })

    eg.Go(func() error{
        return serveApp(ctx, 8081)
    })

    eg.Go(func() error {
        return sysCmdSign(ctx)
    })

    if err := eg.Wait(); err != nil {
        //log.Printf("Error occurs: %v ", err)
        fmt.Printf("server stop: %+v\n", err)
    }
    
}

//
// 服务封装
//
func serveApp(ctx context.Context, port int) error {

    mux := http.NewServeMux()

    mux.HandleFunc("/hello", func(resp http.ResponseWriter, req *http.Request){
        fmt.Fprintln(resp, "are you ok ?")
    })
  
    addr := fmt.Sprintf("127.0.0.1:%d", port)
    srv := http.Server{Addr: addr, Handler: mux}

    log.Println("Start server ...")
    
    // 监听是否需要关闭服务
    go func() {
            <-ctx.Done()
            err := srv.Shutdown(ctx)
            log.Printf("Shutdown server:%d success, %v", port, err)
    }()

    return srv.ListenAndServe()
}

//
// 处理系统信号
//
func sysCmdSign(ctx context.Context) error {

    c := make(chan os.Signal, 0)
    signal.Notify(c, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
    
    select {
    case <-ctx.Done():
            return ctx.Err()
    case sig := <-c:
            return errors.Errorf("kill by signal: %v", sig)
        
    }
    
    //for {
    //    s := <-c
    //    switch s {
    //    case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
    //        log.Printf("get signal: %s\n", s.String())
    //        return errors.Errorf(fmt.Sprintf("kill by signal: %s", s.String()))
    //        //return fmt.Errorf("kill by signal: %s", s.String())
    //    case syscall.SIGHUP:

    //    default:
    //        return nil
    //    }
    //}
}
