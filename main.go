package main

import (
    "log"
    "bioinformatics-as-a-service/internal/machinelearning"
    "bioinformatics-as-a-service/internal/data"
    "bioinformatics-as-a-service/internal/gui"
    "bioinformatics-as-a-service/internal/visualization"
    "bioinformatics-as-a-service/internal/storage"
)

func main() {
    // Connect to the cloud database
    err := cloud.Connect()
    if err != nil {
        log.Fatal(err)
    }
    defer cloud.Close()

    // Start the web GUI
    go gui.Start()

    // Start collecting data from the Apple Watches
    watchChan := make(chan *WatchData)
    go watch.Start(watchChan)

    for {
        select {
        case data := <-watchChan:
            // Process the data
            processedData, err := data.Process()
            if err != nil {
                log.Println(err)
                continue
            }

            // Train the machine learning model
            err = machinelearning.Train(processedData)
            if err != nil {
                log.Println(err)
                continue
            }

            // Visualize the data
            err = visualization.Visualize(processedData)
            if err != nil {
                log.Println(err)
                continue
            }

            // Store the data in the database
            err = storage.Store(processedData)
            if err != nil {
                log.Println(err)
                continue
            }
        }
    }
}
