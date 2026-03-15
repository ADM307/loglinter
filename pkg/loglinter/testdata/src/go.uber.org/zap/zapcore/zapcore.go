package zapcore

type FieldType uint8

const (
	StringType FieldType = iota
)

type Field struct {
	Key    string
	String string
	Type   FieldType
}
