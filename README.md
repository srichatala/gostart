# Go Lang Web Application

###Configure go environment on Windows

Download [golang] (https://golang.org/doc/install?download=go1.7.3.windows-amd64.msi)

####System level variables 

 - Add C:\Go by editing system variable path in Environment variable 
 - Add New user variable
    - Variable name : GOPATH
    - Variable value : D:\git\gostart
    
####Application folder structure:
path D:\git\gostart

    bin
    pkg 
    src
        main
            hello.go
            
       Sample code:
       
       package main
       
       import "fmt"
       
       func main(){
           fmt.Println("Hello world")
       }
 
####Environment Setup
  - To check environment command, run the following command
  
        D:\git\gostart>go env  
  - Set GOPATH
  
        D:git\gostart> set GOPATH=d:git\gostart
  - Set GOBIN
  
        D:git\gostart> set GOBIN=d:git\gostart\bin
      
####Run
  - Run command
  
        D:git\gostart> go run src/main/hello.go   
  - Output
  
        Hello world
  
  
  
  
 







