package runtime

import "unsafe"

func Malloc(size uintptr) unsafe.Pointer {
	return sysAllocOS(size)
}

func Calloc(num uintptr, size uintptr) unsafe.Pointer {
	fSize := num * size
	ptr := sysAllocOS(fSize)
	memclrNoHeapPointers(ptr, fSize)
	return ptr
}

func Free(ptr unsafe.Pointer, size uintptr) {
	sysFreeOS(ptr, size)
}
