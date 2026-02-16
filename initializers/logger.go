package initializers

import "go.uber.org/zap"

var Log *zap.Logger

func Logger() {
	var err error

	Log, err = zap.NewDevelopment()
	if err != nil{
		panic("failed to initialize logger: " + err.Error()) 
	}

}
