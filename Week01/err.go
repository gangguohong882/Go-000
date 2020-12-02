package main

import (
	"github.com/pkg/errors"
	"fmt"
)

func main() {
	err := biz()
	fmt.printf("service:%+v",err)
}

func dao() err {
	
      
      return nil
}


func biz()  err {

   if err := dao(); err != nil {
	return erros.wrap(err,"biz err")
}

   return nil	
}