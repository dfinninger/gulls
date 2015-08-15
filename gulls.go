package main

import (
  //"github.com/goamz/goamz/s3"
  //"github.com/goamz/goamz/aws"
  "flag"
)

var (
  awsSecretKey string
  awsAccessKey string
  s3Path string
)

func init() {
  flag.StringVar(&awsSecretKey, "secret-key", "", "AWS Secret Key")
  flag.StringVar(&awsAccessKey, "access-key", "", "AWS Access Key")
  flag.StringVar(&s3Path, "path", "", "Queue to monitor")
}

func check(err error) {
  if err != nil {
    panic(err)
  }
}

// Main ----------------------------------------------------------------
func main() {
  flag.Parse()

  //fmt.Printf("Messsages in [%s]: %v", sqsQueue, count)
}
