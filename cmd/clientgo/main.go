package main

import (
	"os"

	"github.com/hold7door/clientgo"
	"github.com/rs/zerolog"
)

func main(){
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if err := clientgo.CreateCommand().Execute(); err != nil {
		switch e := err.(type){
		case clientgo.ReturnCodeError:
			os.Exit(e.Code())	
		default:
			os.Exit(1)
		}
	}
}