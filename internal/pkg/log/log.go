package log

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/vinoMamba/lazydoc/internal/pkg/known"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	Sync()
}

type zapLogger struct {
	z *zap.Logger
}

var _ Logger = &zapLogger{}

var (
	mu  sync.Mutex
	std = NewLogger(NewOpts())
)

func InitLogger(opts *Options) {
	mu.Lock()
	defer mu.Unlock()
	std = NewLogger(opts)
}

func NewLogger(opts *Options) *zapLogger {
	if opts == nil {
		opts = NewOpts()
	}
	var zapLevel zapcore.Level
	if err := zapLevel.UnmarshalText([]byte(opts.Level)); err != nil {
		zapLevel = zapcore.InfoLevel
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.MessageKey = "message"
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encoderConfig.EncodeDuration = func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendFloat64(float64(d) / float64(time.Millisecond))
	}

	cfg := &zap.Config{
		DisableCaller:     opts.DisableCaller,
		DisableStacktrace: opts.DisableStacktrace,
		Level:             zap.NewAtomicLevelAt(zapLevel),
		Encoding:          opts.Format,
		EncoderConfig:     encoderConfig,
		OutputPaths:       opts.OutputPaths,
		ErrorOutputPaths:  []string{"stderr"},
	}

	z, err := cfg.Build(zap.AddStacktrace(zapcore.PanicLevel), zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	logger := &zapLogger{z: z}
	zap.RedirectStdLog(z)

	return logger
}

func (l *zapLogger) Sync() {
	_ = l.z.Sync()
}
func Sync() {
	std.Sync()
}

func (l *zapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Debugw(msg, keysAndValues...)
}
func Debugw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Debugw(msg, keysAndValues...)
}

func (l *zapLogger) Infow(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Infow(msg, keysAndValues...)
}
func Infow(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Infow(msg, keysAndValues...)
}

func (l *zapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Warnw(msg, keysAndValues...)
}
func Warnw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Warnw(msg, keysAndValues...)
}

func (l *zapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Errorw(msg, keysAndValues...)
}
func Errorw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Errorw(msg, keysAndValues...)
}

func (l *zapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Panicw(msg, keysAndValues...)
}
func Panicw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Panicw(msg, keysAndValues...)
}

func (l *zapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	l.z.Sugar().Fatalw(msg, keysAndValues...)
}
func Fatalw(msg string, keysAndValues ...interface{}) {
	std.z.Sugar().Fatalw(msg, keysAndValues...)
}

func C(ctx context.Context) *zapLogger {
	return std.C(ctx)
}

func (l *zapLogger) C(ctx context.Context) *zapLogger {
	lc := l.clone()
	//打印ctx
	fmt.Println(ctx)

	if requestID := ctx.Value(known.XRequestIDKey); requestID != nil {
		lc.z = lc.z.With(zap.Any(known.XRequestIDKey, requestID))
	}

	return lc
}

func (l *zapLogger) clone() *zapLogger {
	lc := *l
	return &lc
}
