// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Tue, 05 Sep 2017 19:49:49 MSK.
// By https://git.io/cgogen. DO NOT EDIT.

package portaudio

/*
#cgo pkg-config: portaudio-2.0
#include <portaudio.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"
import (
	"sync"
	"unsafe"
)

// cgoAllocMap stores pointers to C allocated memory for future reference.
type cgoAllocMap struct {
	mux sync.RWMutex
	m   map[unsafe.Pointer]struct{}
}

var cgoAllocsUnknown = new(cgoAllocMap)

func (a *cgoAllocMap) Add(ptr unsafe.Pointer) {
	a.mux.Lock()
	if a.m == nil {
		a.m = make(map[unsafe.Pointer]struct{})
	}
	a.m[ptr] = struct{}{}
	a.mux.Unlock()
}

func (a *cgoAllocMap) IsEmpty() bool {
	a.mux.RLock()
	isEmpty := len(a.m) == 0
	a.mux.RUnlock()
	return isEmpty
}

func (a *cgoAllocMap) Borrow(b *cgoAllocMap) {
	if b == nil || b.IsEmpty() {
		return
	}
	b.mux.Lock()
	a.mux.Lock()
	for ptr := range b.m {
		if a.m == nil {
			a.m = make(map[unsafe.Pointer]struct{})
		}
		a.m[ptr] = struct{}{}
		delete(b.m, ptr)
	}
	a.mux.Unlock()
	b.mux.Unlock()
}

func (a *cgoAllocMap) Free() {
	a.mux.Lock()
	for ptr := range a.m {
		C.free(ptr)
		delete(a.m, ptr)
	}
	a.mux.Unlock()
}

// allocHostApiInfoMemory allocates memory for type C.PaHostApiInfo in C.
// The caller is responsible for freeing the this memory via C.free.
func allocHostApiInfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfHostApiInfoValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfHostApiInfoValue = unsafe.Sizeof([1]C.PaHostApiInfo{})

// unpackPCharString represents the data from Go string as *C.char and avoids copying.
func unpackPCharString(str string) (*C.char, *cgoAllocMap) {
	h := (*stringHeader)(unsafe.Pointer(&str))
	return (*C.char)(unsafe.Pointer(h.Data)), cgoAllocsUnknown
}

type stringHeader struct {
	Data uintptr
	Len  int
}

// packPCharString creates a Go string backed by *C.char and avoids copying.
func packPCharString(p *C.char) (raw string) {
	if p != nil && *p != 0 {
		h := (*stringHeader)(unsafe.Pointer(&raw))
		h.Data = uintptr(unsafe.Pointer(p))
		for *p != 0 {
			p = (*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) // p++
		}
		h.Len = int(uintptr(unsafe.Pointer(p)) - h.Data)
	}
	return
}

// RawString reperesents a string backed by data on the C side.
type RawString string

// Copy returns a Go-managed copy of raw string.
func (raw RawString) Copy() string {
	if len(raw) == 0 {
		return ""
	}
	h := (*stringHeader)(unsafe.Pointer(&raw))
	return C.GoStringN((*C.char)(unsafe.Pointer(h.Data)), C.int(h.Len))
}

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *HostApiInfo) Ref() *C.PaHostApiInfo {
	if x == nil {
		return nil
	}
	return x.ref2fd9b1fc
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *HostApiInfo) Free() {
	if x != nil && x.allocs2fd9b1fc != nil {
		x.allocs2fd9b1fc.(*cgoAllocMap).Free()
		x.ref2fd9b1fc = nil
	}
}

// NewHostApiInfoRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewHostApiInfoRef(ref unsafe.Pointer) *HostApiInfo {
	if ref == nil {
		return nil
	}
	obj := new(HostApiInfo)
	obj.ref2fd9b1fc = (*C.PaHostApiInfo)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *HostApiInfo) PassRef() (*C.PaHostApiInfo, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref2fd9b1fc != nil {
		return x.ref2fd9b1fc, nil
	}
	mem2fd9b1fc := allocHostApiInfoMemory(1)
	ref2fd9b1fc := (*C.PaHostApiInfo)(mem2fd9b1fc)
	allocs2fd9b1fc := new(cgoAllocMap)
	var cstructVersion_allocs *cgoAllocMap
	ref2fd9b1fc.structVersion, cstructVersion_allocs = (C.int)(x.StructVersion), cgoAllocsUnknown
	allocs2fd9b1fc.Borrow(cstructVersion_allocs)

	var c_type_allocs *cgoAllocMap
	ref2fd9b1fc._type, c_type_allocs = (C.PaHostApiTypeId)(x.Type), cgoAllocsUnknown
	allocs2fd9b1fc.Borrow(c_type_allocs)

	var cname_allocs *cgoAllocMap
	ref2fd9b1fc.name, cname_allocs = unpackPCharString(x.Name)
	allocs2fd9b1fc.Borrow(cname_allocs)

	var cdeviceCount_allocs *cgoAllocMap
	ref2fd9b1fc.deviceCount, cdeviceCount_allocs = (C.int)(x.DeviceCount), cgoAllocsUnknown
	allocs2fd9b1fc.Borrow(cdeviceCount_allocs)

	var cdefaultInputDevice_allocs *cgoAllocMap
	ref2fd9b1fc.defaultInputDevice, cdefaultInputDevice_allocs = (C.PaDeviceIndex)(x.DefaultInputDevice), cgoAllocsUnknown
	allocs2fd9b1fc.Borrow(cdefaultInputDevice_allocs)

	var cdefaultOutputDevice_allocs *cgoAllocMap
	ref2fd9b1fc.defaultOutputDevice, cdefaultOutputDevice_allocs = (C.PaDeviceIndex)(x.DefaultOutputDevice), cgoAllocsUnknown
	allocs2fd9b1fc.Borrow(cdefaultOutputDevice_allocs)

	x.ref2fd9b1fc = ref2fd9b1fc
	x.allocs2fd9b1fc = allocs2fd9b1fc
	return ref2fd9b1fc, allocs2fd9b1fc

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x HostApiInfo) PassValue() (C.PaHostApiInfo, *cgoAllocMap) {
	if x.ref2fd9b1fc != nil {
		return *x.ref2fd9b1fc, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *HostApiInfo) Deref() {
	if x.ref2fd9b1fc == nil {
		return
	}
	x.StructVersion = (int32)(x.ref2fd9b1fc.structVersion)
	x.Type = (HostApiTypeId)(x.ref2fd9b1fc._type)
	x.Name = packPCharString(x.ref2fd9b1fc.name)
	x.DeviceCount = (int32)(x.ref2fd9b1fc.deviceCount)
	x.DefaultInputDevice = (DeviceIndex)(x.ref2fd9b1fc.defaultInputDevice)
	x.DefaultOutputDevice = (DeviceIndex)(x.ref2fd9b1fc.defaultOutputDevice)
}

// allocHostErrorInfoMemory allocates memory for type C.PaHostErrorInfo in C.
// The caller is responsible for freeing the this memory via C.free.
func allocHostErrorInfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfHostErrorInfoValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfHostErrorInfoValue = unsafe.Sizeof([1]C.PaHostErrorInfo{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *HostErrorInfo) Ref() *C.PaHostErrorInfo {
	if x == nil {
		return nil
	}
	return x.ref5c7a77d7
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *HostErrorInfo) Free() {
	if x != nil && x.allocs5c7a77d7 != nil {
		x.allocs5c7a77d7.(*cgoAllocMap).Free()
		x.ref5c7a77d7 = nil
	}
}

// NewHostErrorInfoRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewHostErrorInfoRef(ref unsafe.Pointer) *HostErrorInfo {
	if ref == nil {
		return nil
	}
	obj := new(HostErrorInfo)
	obj.ref5c7a77d7 = (*C.PaHostErrorInfo)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *HostErrorInfo) PassRef() (*C.PaHostErrorInfo, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref5c7a77d7 != nil {
		return x.ref5c7a77d7, nil
	}
	mem5c7a77d7 := allocHostErrorInfoMemory(1)
	ref5c7a77d7 := (*C.PaHostErrorInfo)(mem5c7a77d7)
	allocs5c7a77d7 := new(cgoAllocMap)
	var chostApiType_allocs *cgoAllocMap
	ref5c7a77d7.hostApiType, chostApiType_allocs = (C.PaHostApiTypeId)(x.HostApiType), cgoAllocsUnknown
	allocs5c7a77d7.Borrow(chostApiType_allocs)

	var cerrorCode_allocs *cgoAllocMap
	ref5c7a77d7.errorCode, cerrorCode_allocs = (C.long)(x.ErrorCode), cgoAllocsUnknown
	allocs5c7a77d7.Borrow(cerrorCode_allocs)

	var cerrorText_allocs *cgoAllocMap
	ref5c7a77d7.errorText, cerrorText_allocs = unpackPCharString(x.ErrorText)
	allocs5c7a77d7.Borrow(cerrorText_allocs)

	x.ref5c7a77d7 = ref5c7a77d7
	x.allocs5c7a77d7 = allocs5c7a77d7
	return ref5c7a77d7, allocs5c7a77d7

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x HostErrorInfo) PassValue() (C.PaHostErrorInfo, *cgoAllocMap) {
	if x.ref5c7a77d7 != nil {
		return *x.ref5c7a77d7, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *HostErrorInfo) Deref() {
	if x.ref5c7a77d7 == nil {
		return
	}
	x.HostApiType = (HostApiTypeId)(x.ref5c7a77d7.hostApiType)
	x.ErrorCode = (int)(x.ref5c7a77d7.errorCode)
	x.ErrorText = packPCharString(x.ref5c7a77d7.errorText)
}

// allocDeviceInfoMemory allocates memory for type C.PaDeviceInfo in C.
// The caller is responsible for freeing the this memory via C.free.
func allocDeviceInfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfDeviceInfoValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfDeviceInfoValue = unsafe.Sizeof([1]C.PaDeviceInfo{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *DeviceInfo) Ref() *C.PaDeviceInfo {
	if x == nil {
		return nil
	}
	return x.refb60de337
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *DeviceInfo) Free() {
	if x != nil && x.allocsb60de337 != nil {
		x.allocsb60de337.(*cgoAllocMap).Free()
		x.refb60de337 = nil
	}
}

// NewDeviceInfoRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewDeviceInfoRef(ref unsafe.Pointer) *DeviceInfo {
	if ref == nil {
		return nil
	}
	obj := new(DeviceInfo)
	obj.refb60de337 = (*C.PaDeviceInfo)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *DeviceInfo) PassRef() (*C.PaDeviceInfo, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.refb60de337 != nil {
		return x.refb60de337, nil
	}
	memb60de337 := allocDeviceInfoMemory(1)
	refb60de337 := (*C.PaDeviceInfo)(memb60de337)
	allocsb60de337 := new(cgoAllocMap)
	var cstructVersion_allocs *cgoAllocMap
	refb60de337.structVersion, cstructVersion_allocs = (C.int)(x.StructVersion), cgoAllocsUnknown
	allocsb60de337.Borrow(cstructVersion_allocs)

	var cname_allocs *cgoAllocMap
	refb60de337.name, cname_allocs = unpackPCharString(x.Name)
	allocsb60de337.Borrow(cname_allocs)

	var chostApi_allocs *cgoAllocMap
	refb60de337.hostApi, chostApi_allocs = (C.PaHostApiIndex)(x.HostApi), cgoAllocsUnknown
	allocsb60de337.Borrow(chostApi_allocs)

	var cmaxInputChannels_allocs *cgoAllocMap
	refb60de337.maxInputChannels, cmaxInputChannels_allocs = (C.int)(x.MaxInputChannels), cgoAllocsUnknown
	allocsb60de337.Borrow(cmaxInputChannels_allocs)

	var cmaxOutputChannels_allocs *cgoAllocMap
	refb60de337.maxOutputChannels, cmaxOutputChannels_allocs = (C.int)(x.MaxOutputChannels), cgoAllocsUnknown
	allocsb60de337.Borrow(cmaxOutputChannels_allocs)

	var cdefaultLowInputLatency_allocs *cgoAllocMap
	refb60de337.defaultLowInputLatency, cdefaultLowInputLatency_allocs = (C.PaTime)(x.DefaultLowInputLatency), cgoAllocsUnknown
	allocsb60de337.Borrow(cdefaultLowInputLatency_allocs)

	var cdefaultLowOutputLatency_allocs *cgoAllocMap
	refb60de337.defaultLowOutputLatency, cdefaultLowOutputLatency_allocs = (C.PaTime)(x.DefaultLowOutputLatency), cgoAllocsUnknown
	allocsb60de337.Borrow(cdefaultLowOutputLatency_allocs)

	var cdefaultHighInputLatency_allocs *cgoAllocMap
	refb60de337.defaultHighInputLatency, cdefaultHighInputLatency_allocs = (C.PaTime)(x.DefaultHighInputLatency), cgoAllocsUnknown
	allocsb60de337.Borrow(cdefaultHighInputLatency_allocs)

	var cdefaultHighOutputLatency_allocs *cgoAllocMap
	refb60de337.defaultHighOutputLatency, cdefaultHighOutputLatency_allocs = (C.PaTime)(x.DefaultHighOutputLatency), cgoAllocsUnknown
	allocsb60de337.Borrow(cdefaultHighOutputLatency_allocs)

	var cdefaultSampleRate_allocs *cgoAllocMap
	refb60de337.defaultSampleRate, cdefaultSampleRate_allocs = (C.double)(x.DefaultSampleRate), cgoAllocsUnknown
	allocsb60de337.Borrow(cdefaultSampleRate_allocs)

	x.refb60de337 = refb60de337
	x.allocsb60de337 = allocsb60de337
	return refb60de337, allocsb60de337

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x DeviceInfo) PassValue() (C.PaDeviceInfo, *cgoAllocMap) {
	if x.refb60de337 != nil {
		return *x.refb60de337, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *DeviceInfo) Deref() {
	if x.refb60de337 == nil {
		return
	}
	x.StructVersion = (int32)(x.refb60de337.structVersion)
	x.Name = packPCharString(x.refb60de337.name)
	x.HostApi = (HostApiIndex)(x.refb60de337.hostApi)
	x.MaxInputChannels = (int32)(x.refb60de337.maxInputChannels)
	x.MaxOutputChannels = (int32)(x.refb60de337.maxOutputChannels)
	x.DefaultLowInputLatency = (Time)(x.refb60de337.defaultLowInputLatency)
	x.DefaultLowOutputLatency = (Time)(x.refb60de337.defaultLowOutputLatency)
	x.DefaultHighInputLatency = (Time)(x.refb60de337.defaultHighInputLatency)
	x.DefaultHighOutputLatency = (Time)(x.refb60de337.defaultHighOutputLatency)
	x.DefaultSampleRate = (float64)(x.refb60de337.defaultSampleRate)
}

// allocStreamParametersMemory allocates memory for type C.PaStreamParameters in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStreamParametersMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStreamParametersValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStreamParametersValue = unsafe.Sizeof([1]C.PaStreamParameters{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *StreamParameters) Ref() *C.PaStreamParameters {
	if x == nil {
		return nil
	}
	return x.reff0e97a84
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *StreamParameters) Free() {
	if x != nil && x.allocsf0e97a84 != nil {
		x.allocsf0e97a84.(*cgoAllocMap).Free()
		x.reff0e97a84 = nil
	}
}

// NewStreamParametersRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewStreamParametersRef(ref unsafe.Pointer) *StreamParameters {
	if ref == nil {
		return nil
	}
	obj := new(StreamParameters)
	obj.reff0e97a84 = (*C.PaStreamParameters)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *StreamParameters) PassRef() (*C.PaStreamParameters, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff0e97a84 != nil {
		return x.reff0e97a84, nil
	}
	memf0e97a84 := allocStreamParametersMemory(1)
	reff0e97a84 := (*C.PaStreamParameters)(memf0e97a84)
	allocsf0e97a84 := new(cgoAllocMap)
	var cdevice_allocs *cgoAllocMap
	reff0e97a84.device, cdevice_allocs = (C.PaDeviceIndex)(x.Device), cgoAllocsUnknown
	allocsf0e97a84.Borrow(cdevice_allocs)

	var cchannelCount_allocs *cgoAllocMap
	reff0e97a84.channelCount, cchannelCount_allocs = (C.int)(x.ChannelCount), cgoAllocsUnknown
	allocsf0e97a84.Borrow(cchannelCount_allocs)

	var csampleFormat_allocs *cgoAllocMap
	reff0e97a84.sampleFormat, csampleFormat_allocs = (C.PaSampleFormat)(x.SampleFormat), cgoAllocsUnknown
	allocsf0e97a84.Borrow(csampleFormat_allocs)

	var csuggestedLatency_allocs *cgoAllocMap
	reff0e97a84.suggestedLatency, csuggestedLatency_allocs = (C.PaTime)(x.SuggestedLatency), cgoAllocsUnknown
	allocsf0e97a84.Borrow(csuggestedLatency_allocs)

	var chostApiSpecificStreamInfo_allocs *cgoAllocMap
	reff0e97a84.hostApiSpecificStreamInfo, chostApiSpecificStreamInfo_allocs = *(*unsafe.Pointer)(unsafe.Pointer(&x.HostApiSpecificStreamInfo)), cgoAllocsUnknown
	allocsf0e97a84.Borrow(chostApiSpecificStreamInfo_allocs)

	x.reff0e97a84 = reff0e97a84
	x.allocsf0e97a84 = allocsf0e97a84
	return reff0e97a84, allocsf0e97a84

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x StreamParameters) PassValue() (C.PaStreamParameters, *cgoAllocMap) {
	if x.reff0e97a84 != nil {
		return *x.reff0e97a84, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *StreamParameters) Deref() {
	if x.reff0e97a84 == nil {
		return
	}
	x.Device = (DeviceIndex)(x.reff0e97a84.device)
	x.ChannelCount = (int32)(x.reff0e97a84.channelCount)
	x.SampleFormat = (SampleFormat)(x.reff0e97a84.sampleFormat)
	x.SuggestedLatency = (Time)(x.reff0e97a84.suggestedLatency)
	x.HostApiSpecificStreamInfo = (unsafe.Pointer)(unsafe.Pointer(x.reff0e97a84.hostApiSpecificStreamInfo))
}

// allocStreamCallbackTimeInfoMemory allocates memory for type C.PaStreamCallbackTimeInfo in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStreamCallbackTimeInfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStreamCallbackTimeInfoValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStreamCallbackTimeInfoValue = unsafe.Sizeof([1]C.PaStreamCallbackTimeInfo{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *StreamCallbackTimeInfo) Ref() *C.PaStreamCallbackTimeInfo {
	if x == nil {
		return nil
	}
	return x.ref20d551f2
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *StreamCallbackTimeInfo) Free() {
	if x != nil && x.allocs20d551f2 != nil {
		x.allocs20d551f2.(*cgoAllocMap).Free()
		x.ref20d551f2 = nil
	}
}

// NewStreamCallbackTimeInfoRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewStreamCallbackTimeInfoRef(ref unsafe.Pointer) *StreamCallbackTimeInfo {
	if ref == nil {
		return nil
	}
	obj := new(StreamCallbackTimeInfo)
	obj.ref20d551f2 = (*C.PaStreamCallbackTimeInfo)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *StreamCallbackTimeInfo) PassRef() (*C.PaStreamCallbackTimeInfo, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.ref20d551f2 != nil {
		return x.ref20d551f2, nil
	}
	mem20d551f2 := allocStreamCallbackTimeInfoMemory(1)
	ref20d551f2 := (*C.PaStreamCallbackTimeInfo)(mem20d551f2)
	allocs20d551f2 := new(cgoAllocMap)
	var cinputBufferAdcTime_allocs *cgoAllocMap
	ref20d551f2.inputBufferAdcTime, cinputBufferAdcTime_allocs = (C.PaTime)(x.InputBufferAdcTime), cgoAllocsUnknown
	allocs20d551f2.Borrow(cinputBufferAdcTime_allocs)

	var ccurrentTime_allocs *cgoAllocMap
	ref20d551f2.currentTime, ccurrentTime_allocs = (C.PaTime)(x.CurrentTime), cgoAllocsUnknown
	allocs20d551f2.Borrow(ccurrentTime_allocs)

	var coutputBufferDacTime_allocs *cgoAllocMap
	ref20d551f2.outputBufferDacTime, coutputBufferDacTime_allocs = (C.PaTime)(x.OutputBufferDacTime), cgoAllocsUnknown
	allocs20d551f2.Borrow(coutputBufferDacTime_allocs)

	x.ref20d551f2 = ref20d551f2
	x.allocs20d551f2 = allocs20d551f2
	return ref20d551f2, allocs20d551f2

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x StreamCallbackTimeInfo) PassValue() (C.PaStreamCallbackTimeInfo, *cgoAllocMap) {
	if x.ref20d551f2 != nil {
		return *x.ref20d551f2, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *StreamCallbackTimeInfo) Deref() {
	if x.ref20d551f2 == nil {
		return
	}
	x.InputBufferAdcTime = (Time)(x.ref20d551f2.inputBufferAdcTime)
	x.CurrentTime = (Time)(x.ref20d551f2.currentTime)
	x.OutputBufferDacTime = (Time)(x.ref20d551f2.outputBufferDacTime)
}

func (x StreamCallback) PassRef() (ref *C.PaStreamCallback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if streamCallback615600DBFunc == nil {
		streamCallback615600DBFunc = x
	}
	return (*C.PaStreamCallback)(C.PaStreamCallback_615600db), nil
}

func NewStreamCallbackRef(ref unsafe.Pointer) *StreamCallback {
	return (*StreamCallback)(ref)
}

//export streamCallback615600DB
func streamCallback615600DB(cinput unsafe.Pointer, coutput unsafe.Pointer, cframeCount C.ulong, ctimeInfo *C.PaStreamCallbackTimeInfo, cstatusFlags C.PaStreamCallbackFlags, cuserData unsafe.Pointer) C.int {
	if streamCallback615600DBFunc != nil {
		input615600db := (unsafe.Pointer)(unsafe.Pointer(cinput))
		output615600db := (unsafe.Pointer)(unsafe.Pointer(coutput))
		frameCount615600db := (uint)(cframeCount)
		timeInfo615600db := NewStreamCallbackTimeInfoRef(unsafe.Pointer(ctimeInfo))
		statusFlags615600db := (StreamCallbackFlags)(cstatusFlags)
		userData615600db := (unsafe.Pointer)(unsafe.Pointer(cuserData))
		ret615600db := streamCallback615600DBFunc(input615600db, output615600db, frameCount615600db, timeInfo615600db, statusFlags615600db, userData615600db)
		ret, _ := (C.int)(ret615600db), cgoAllocsUnknown
		return ret
	}
	panic("callback func has not been set (race?)")
}

var streamCallback615600DBFunc StreamCallback

func (x StreamFinishedCallback) PassRef() (ref *C.PaStreamFinishedCallback, allocs *cgoAllocMap) {
	if x == nil {
		return nil, nil
	}
	if streamFinishedCallbackD2BB8F82Func == nil {
		streamFinishedCallbackD2BB8F82Func = x
	}
	return (*C.PaStreamFinishedCallback)(C.PaStreamFinishedCallback_d2bb8f82), nil
}

func NewStreamFinishedCallbackRef(ref unsafe.Pointer) *StreamFinishedCallback {
	return (*StreamFinishedCallback)(ref)
}

//export streamFinishedCallbackD2BB8F82
func streamFinishedCallbackD2BB8F82(cuserData unsafe.Pointer) {
	if streamFinishedCallbackD2BB8F82Func != nil {
		userDatad2bb8f82 := (unsafe.Pointer)(unsafe.Pointer(cuserData))
		streamFinishedCallbackD2BB8F82Func(userDatad2bb8f82)
		return
	}
	panic("callback func has not been set (race?)")
}

var streamFinishedCallbackD2BB8F82Func StreamFinishedCallback

// allocStreamInfoMemory allocates memory for type C.PaStreamInfo in C.
// The caller is responsible for freeing the this memory via C.free.
func allocStreamInfoMemory(n int) unsafe.Pointer {
	mem, err := C.calloc(C.size_t(n), (C.size_t)(sizeOfStreamInfoValue))
	if err != nil {
		panic("memory alloc error: " + err.Error())
	}
	return mem
}

const sizeOfStreamInfoValue = unsafe.Sizeof([1]C.PaStreamInfo{})

// Ref returns the underlying reference to C object or nil if struct is nil.
func (x *StreamInfo) Ref() *C.PaStreamInfo {
	if x == nil {
		return nil
	}
	return x.reff696a1d0
}

// Free invokes alloc map's free mechanism that cleanups any allocated memory using C free.
// Does nothing if struct is nil or has no allocation map.
func (x *StreamInfo) Free() {
	if x != nil && x.allocsf696a1d0 != nil {
		x.allocsf696a1d0.(*cgoAllocMap).Free()
		x.reff696a1d0 = nil
	}
}

// NewStreamInfoRef creates a new wrapper struct with underlying reference set to the original C object.
// Returns nil if the provided pointer to C object is nil too.
func NewStreamInfoRef(ref unsafe.Pointer) *StreamInfo {
	if ref == nil {
		return nil
	}
	obj := new(StreamInfo)
	obj.reff696a1d0 = (*C.PaStreamInfo)(unsafe.Pointer(ref))
	return obj
}

// PassRef returns the underlying C object, otherwise it will allocate one and set its values
// from this wrapping struct, counting allocations into an allocation map.
func (x *StreamInfo) PassRef() (*C.PaStreamInfo, *cgoAllocMap) {
	if x == nil {
		return nil, nil
	} else if x.reff696a1d0 != nil {
		return x.reff696a1d0, nil
	}
	memf696a1d0 := allocStreamInfoMemory(1)
	reff696a1d0 := (*C.PaStreamInfo)(memf696a1d0)
	allocsf696a1d0 := new(cgoAllocMap)
	var cstructVersion_allocs *cgoAllocMap
	reff696a1d0.structVersion, cstructVersion_allocs = (C.int)(x.StructVersion), cgoAllocsUnknown
	allocsf696a1d0.Borrow(cstructVersion_allocs)

	var cinputLatency_allocs *cgoAllocMap
	reff696a1d0.inputLatency, cinputLatency_allocs = (C.PaTime)(x.InputLatency), cgoAllocsUnknown
	allocsf696a1d0.Borrow(cinputLatency_allocs)

	var coutputLatency_allocs *cgoAllocMap
	reff696a1d0.outputLatency, coutputLatency_allocs = (C.PaTime)(x.OutputLatency), cgoAllocsUnknown
	allocsf696a1d0.Borrow(coutputLatency_allocs)

	var csampleRate_allocs *cgoAllocMap
	reff696a1d0.sampleRate, csampleRate_allocs = (C.double)(x.SampleRate), cgoAllocsUnknown
	allocsf696a1d0.Borrow(csampleRate_allocs)

	x.reff696a1d0 = reff696a1d0
	x.allocsf696a1d0 = allocsf696a1d0
	return reff696a1d0, allocsf696a1d0

}

// PassValue does the same as PassRef except that it will try to dereference the returned pointer.
func (x StreamInfo) PassValue() (C.PaStreamInfo, *cgoAllocMap) {
	if x.reff696a1d0 != nil {
		return *x.reff696a1d0, nil
	}
	ref, allocs := x.PassRef()
	return *ref, allocs
}

// Deref uses the underlying reference to C object and fills the wrapping struct with values.
// Do not forget to call this method whether you get a struct for C object and want to read its values.
func (x *StreamInfo) Deref() {
	if x.reff696a1d0 == nil {
		return
	}
	x.StructVersion = (int32)(x.reff696a1d0.structVersion)
	x.InputLatency = (Time)(x.reff696a1d0.inputLatency)
	x.OutputLatency = (Time)(x.reff696a1d0.outputLatency)
	x.SampleRate = (float64)(x.reff696a1d0.sampleRate)
}
