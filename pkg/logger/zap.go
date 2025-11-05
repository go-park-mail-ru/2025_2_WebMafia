package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	LevelDebug = "debug"
	LevelInfo  = "info"
	LevelWarn  = "warn"
	LevelError = "error"

	ModeDev  = "dev"
	ModeProd = "prod"
)

type Config struct {
	Level string
	Mode  string
}

type ZapLogger struct {
	sl *zap.SugaredLogger
}

func New(level, mode string) (Logger, error) {
	logLevel, err := zapcore.ParseLevel(level)
	if err != nil {
		return nil, fmt.Errorf("invalid log level: %w", err)
	}

	var encoder zapcore.Encoder

	if mode == ModeDev {
		encoderConfig := zap.NewDevelopmentEncoderConfig()
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoderConfig := zap.NewProductionEncoderConfig()
		encoderConfig.EncodeTime = zapcore.EpochTimeEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(logLevel),
	)

	z := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return &ZapLogger{
		sl: z.Sugar(),
	}, nil
}

func (l *ZapLogger) Debugf(template string, args ...interface{}) {
	l.sl.Debugf(template, args...)
}

func (l *ZapLogger) Infof(template string, args ...interface{}) {
	l.sl.Infof(template, args...)
}

func (l *ZapLogger) Warnf(template string, args ...interface{}) {
	l.sl.Warnf(template, args...)
}

func (l *ZapLogger) Errorf(template string, args ...interface{}) {
	l.sl.Errorf(template, args...)
}

func (l *ZapLogger) Fatalf(template string, args ...interface{}) {
	l.sl.Fatalf(template, args...)
}

func (l *ZapLogger) Sync() error {
	if l.sl != nil {
		return l.sl.Sync()
	}
	return nil
}

func (l *ZapLogger) With(args ...interface{}) Logger {
	newLogger := l.sl.With(args...)

	return &ZapLogger{sl: newLogger}
}
