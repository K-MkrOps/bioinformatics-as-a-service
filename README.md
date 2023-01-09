# bioinformatics-as-a-service

This project collects bioinformatics data from Apple Watches and performs machine learning on the data. The results are visualized and provided to users through a secured web GUI.

## Requirements

Go
Python
MySQL
SQLite
Usage

Clone the repository: 
```
git clone https://github.com/K-MkrOps/bioinformatics-as-a-service
```
### Install the requirements ###

### Run the program: ###
```
go run main.go
```

```
bioinformatics-as-a-service
├── main
│   ├── main.go # This will be the entry point for your program
│   ├── cloud.go # This file will handle communication with the cloud database
│   ├── watch.go # This file will handle communication with the Apple Watch
│   ├── auth.go # This file will handle user authentication
├── internal
│   ├── machinelearning # This directory will contain the machine learning models and code
│   │   ├── model.py # This file will contain the machine learning model
│   │   ├── train.py # This file will contain the code for training the model
│   ├── data # This directory will contain code for handling data
│   │   ├── process.py # This file will contain code for processing data
│   ├── gui # This directory will contain code for the web GUI
│   │   ├── gui.py # This file will contain the code for the web GUI
│   ├── visualization # This directory will contain code for data visualization
│   │   ├── visualize.py # This file will contain code for visualizing data
│   ├── storage # This directory will contain code for data storage
│   │   ├── store.py # This file will contain code for storing data in a database
```
