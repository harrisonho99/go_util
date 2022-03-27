package memory

type MemorySize int64

const (
	Byte     MemorySize = 1
	Kilobyte MemorySize = 1000 * Byte
	Megabyte MemorySize = 1000 * Kilobyte
	Gigabyte MemorySize = 1000 * Megabyte
)
