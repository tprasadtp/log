package log

// Level Log levels
type Level uint32

const (
	// TraceLevel is the constant to use when setting the Trace level for loggers
	// provided by this library.
	TraceLevel Level = iota

	// DebugLevel is the constant to use when setting the Debug level for loggers
	// provided by this library.
	DebugLevel

	// InfoLevel is the constant to use when setting the Info level for loggers
	// provided by this library.
	InfoLevel

	// WarnLevel is the constant to use when setting the Warn level for loggers
	// provided by this library.
	WarnLevel

	// ErrorLevel is the constant to use when setting the Error level for loggers
	// provided by this library.
	ErrorLevel

	// PanicLevel is the constant to use when setting the Panic level for loggers
	// provided by this library.
	PanicLevel

	// FatalLevel is the constant to use when setting the Fatal level for loggers
	// provided by this library.
	FatalLevel
)
