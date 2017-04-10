c := make(chan os.Signal, 1)
    signal.Notify(c, syscall.SIGHUP)

    go func(){
        for sig := range c {
            println(sig)
            fmt.Printf("Got A HUP Signal! Now Reloading Conf....\n")
        }
    }()
    for {
        time.Sleep(1000 * time.Millisecond)
        fmt.Printf(">>")
    }
