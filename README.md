# Disks On Trail

Disks On Trail is a command-line application written in Go that allows you to manage and visualize trails and disks. The application supports printing trails with colorful disks, moving disks between trails, and provides a simple command interface.

## Features

- Print trails with colorful disks
- Move disks between trails
- Simple command interface

## Installation

1. Clone the repository:

```sh
git clone https://github.com/theristes/disks-on-trail.git
```

2. Navigate to the project directory:

```sh
cd disks-on-trail
```

3. Build the application:

```sh
go build . 
```

## Usage
Run the application:

```sh
./disks-on-trail
```
### Commands

P: Print the trails
M: Move a disk
H: Print help
Q: Quit the application

### Moving Disks

Enter the index of the trail you want to move from and press enter.
Enter the index of the trail you want to move to and press enter.


```sh
P
#1  T R A I L :         4
A  B I G G E R    D I S K 
T H I S   B I G   D I S K 
A   M E D I U M   D I S K 
A    S M A L L    D I S K 

#2  T R A I L :         0

#3  T R A I L :         0

M
Enter the index of the trail you want to move from and press enter
1
Enter the index of the trail you want to move to and press enter
2
```

### License
This project is licensed under the MIT License