package protocol

import log "github.com/edison-moreland/nreal-hud/sdk/log"

/*

	DO NOT MODIFY!
	This file was auto-generated on 2023-02-23 18:09:24.820761528 -0800 PST m=+0.004816623
	protocol: wayland

*/

// WlDisplay core global object
/*
	The core global object.  This is a special singleton object.  It
	is used for internal Wayland protocol features.
*/
type WlDisplay struct {
	Proxy
}

func NewWlDisplay(p Proxy) WlDisplay {
	return WlDisplay{Proxy: p}
}
func (w *WlDisplay) AttachListener(l WlDisplayListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

// WlDisplayError global error values
/*
	These errors are global and can be emitted in response to any
	server request.
*/
type WlDisplayError uint32

const (
	// WlDisplayError_InvalidObject server couldn't find object
	WlDisplayError_InvalidObject WlDisplayError = 0
	// WlDisplayError_InvalidMethod method doesn't exist on the specified interface or malformed request
	WlDisplayError_InvalidMethod WlDisplayError = 1
	// WlDisplayError_NoMemory server is out of memory
	WlDisplayError_NoMemory WlDisplayError = 2
	// WlDisplayError_Implementation implementation error in compositor
	WlDisplayError_Implementation WlDisplayError = 3
)

// WlDisplaySyncRequest asynchronous roundtrip
/*
	The sync request asks the server to emit the 'done' event
	on the returned wl_callback object.  Since requests are
	handled in-order and events are delivered in-order, this can
	be used as a barrier to ensure all previous requests and the
	resulting events have been handled.

	The object returned by this request will be destroyed by the
	compositor after the callback is fired and as such the client must not
	attempt to use it after that point.

	The callback_data passed in the callback is the event serial.
*/
type WlDisplaySyncRequest struct {
	// Callback  object for the sync request
	Callback uint32 `wayland:"new_id"`
}

// WlDisplayGetRegistryRequest get global registry object
/*
	This request creates a registry object that allows the client
	to list and bind the global objects available from the
	compositor.

	It should be noted that the server side resources consumed in
	response to a get_registry request can only be released when the
	client disconnects, not when the client side proxy is destroyed.
	Therefore, clients should invoke get_registry as infrequently as
	possible to avoid wasting memory.
*/
type WlDisplayGetRegistryRequest struct {
	// Registry global registry object
	Registry uint32 `wayland:"new_id"`
}

// WlDisplayErrorEvent fatal error event
/*
	The error event is sent out when a fatal (non-recoverable)
	error has occurred.  The object_id argument is the object
	where the error occurred, most often in response to a request
	to that object.  The code identifies the error and is defined
	by the object interface.  As such, each interface defines its
	own set of error codes.  The message is a brief description
	of the error, for (debugging) convenience.
*/
type WlDisplayErrorEvent struct {
	// ObjectId object where the error occurred
	ObjectId uint32 `wayland:"object"`
	// Code error code
	Code uint32 `wayland:"uint"`
	// Message error description
	Message string `wayland:"string"`
}

func (e *WlDisplayErrorEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.ObjectId = UnmarshallUint32(offset, d)
	offset, e.Code = UnmarshallUint32(offset, d)
	offset, e.Message = UnmarshallString(offset, d)
	return nil
}

// WlDisplayDeleteIdEvent acknowledge object ID deletion
/*
	This event is used internally by the object ID management
	logic. When a client deletes an object that it had created,
	the server will send this event to acknowledge that it has
	seen the delete request. When the client receives this event,
	it will know that it can safely reuse the object ID.
*/
type WlDisplayDeleteIdEvent struct {
	// Id deleted object ID
	Id uint32 `wayland:"uint"`
}

func (e *WlDisplayDeleteIdEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Id = UnmarshallUint32(offset, d)
	return nil
}

type WlDisplayListener interface {
	Error(WlDisplayErrorEvent)
	DeleteId(WlDisplayDeleteIdEvent)
}
type UnimplementedWlDisplayListener struct{}

func (e *UnimplementedWlDisplayListener) Error(_ WlDisplayErrorEvent) {
	return
}
func (e *UnimplementedWlDisplayListener) DeleteId(_ WlDisplayDeleteIdEvent) {
	return
}

type WlDisplayDispatcher struct {
	WlDisplayListener
}

func NewWlDisplayDispatcher() *WlDisplayDispatcher {
	return &WlDisplayDispatcher{WlDisplayListener: &UnimplementedWlDisplayListener{}}
}
func (i *WlDisplayDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	case 0:
		var e WlDisplayErrorEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Error(e)
	case 1:
		var e WlDisplayDeleteIdEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.DeleteId(e)
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlDisplayDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlDisplayListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlDisplayListener = listener
}

// WlRegistry global registry object
/*
	The singleton global registry object.  The server has a number of
	global objects that are available to all clients.  These objects
	typically represent an actual object in the server (for example,
	an input device) or they are singleton objects that provide
	extension functionality.

	When a client creates a registry object, the registry object
	will emit a global event for each global currently in the
	registry.  Globals come and go as a result of device or
	monitor hotplugs, reconfiguration or other events, and the
	registry will send out global and global_remove events to
	keep the client up to date with the changes.  To mark the end
	of the initial burst of events, the client can use the
	wl_display.sync request immediately after calling
	wl_display.get_registry.

	A client can bind to a global object by using the bind
	request.  This creates a client-side handle that lets the object
	emit events to the client and lets the client invoke requests on
	the object.
*/
type WlRegistry struct {
	Proxy
}

func NewWlRegistry(p Proxy) WlRegistry {
	return WlRegistry{Proxy: p}
}
func (w *WlRegistry) AttachListener(l WlRegistryListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

// WlRegistryBindRequest bind an object to the display
/*
	Binds a new, client-created object to the server using the
	specified name as the identifier.
*/
type WlRegistryBindRequest struct {
	// Name unique numeric name of the object
	Name uint32 `wayland:"uint"`
	// Id bounded object
	Id uint32 `wayland:"new_id"`
}

// WlRegistryGlobalEvent announce global object
/*
	Notify the client of global objects.

	The event notifies the client that a global object with
	the given name is now available, and it implements the
	given version of the given interface.
*/
type WlRegistryGlobalEvent struct {
	// Name numeric name of the global object
	Name uint32 `wayland:"uint"`
	// Interface  implemented by the object
	Interface string `wayland:"string"`
	// Version interface version
	Version uint32 `wayland:"uint"`
}

func (e *WlRegistryGlobalEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Name = UnmarshallUint32(offset, d)
	offset, e.Interface = UnmarshallString(offset, d)
	offset, e.Version = UnmarshallUint32(offset, d)
	return nil
}

// WlRegistryGlobalRemoveEvent announce removal of global object
/*
	Notify the client of removed global objects.

	This event notifies the client that the global identified
	by name is no longer available.  If the client bound to
	the global using the bind request, the client should now
	destroy that object.

	The object remains valid and requests to the object will be
	ignored until the client destroys it, to avoid races between
	the global going away and a client sending a request to it.
*/
type WlRegistryGlobalRemoveEvent struct {
	// Name numeric name of the global object
	Name uint32 `wayland:"uint"`
}

func (e *WlRegistryGlobalRemoveEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Name = UnmarshallUint32(offset, d)
	return nil
}

type WlRegistryListener interface {
	Global(WlRegistryGlobalEvent)
	GlobalRemove(WlRegistryGlobalRemoveEvent)
}
type UnimplementedWlRegistryListener struct{}

func (e *UnimplementedWlRegistryListener) Global(_ WlRegistryGlobalEvent) {
	return
}
func (e *UnimplementedWlRegistryListener) GlobalRemove(_ WlRegistryGlobalRemoveEvent) {
	return
}

type WlRegistryDispatcher struct {
	WlRegistryListener
}

func NewWlRegistryDispatcher() *WlRegistryDispatcher {
	return &WlRegistryDispatcher{WlRegistryListener: &UnimplementedWlRegistryListener{}}
}
func (i *WlRegistryDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	case 0:
		var e WlRegistryGlobalEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Global(e)
	case 1:
		var e WlRegistryGlobalRemoveEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.GlobalRemove(e)
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlRegistryDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlRegistryListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlRegistryListener = listener
}

// WlCallback callback object
/*
	Clients can handle the 'done' event to get notified when
	the related request is done.
*/
type WlCallback struct {
	Proxy
}

func NewWlCallback(p Proxy) WlCallback {
	return WlCallback{Proxy: p}
}
func (w *WlCallback) AttachListener(l WlCallbackListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

// WlCallbackDoneEvent done event
//
//	Notify the client when the related request is done.
type WlCallbackDoneEvent struct {
	// CallbackData request-specific data for the callback
	CallbackData uint32 `wayland:"uint"`
}

func (e *WlCallbackDoneEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.CallbackData = UnmarshallUint32(offset, d)
	return nil
}

type WlCallbackListener interface {
	Done(WlCallbackDoneEvent)
}
type UnimplementedWlCallbackListener struct{}

func (e *UnimplementedWlCallbackListener) Done(_ WlCallbackDoneEvent) {
	return
}

type WlCallbackDispatcher struct {
	WlCallbackListener
}

func NewWlCallbackDispatcher() *WlCallbackDispatcher {
	return &WlCallbackDispatcher{WlCallbackListener: &UnimplementedWlCallbackListener{}}
}
func (i *WlCallbackDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	case 0:
		var e WlCallbackDoneEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Done(e)
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlCallbackDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlCallbackListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlCallbackListener = listener
}

// WlCompositor the compositor singleton
/*
	A compositor.  This object is a singleton global.  The
	compositor is in charge of combining the contents of multiple
	surfaces into one displayable output.
*/
type WlCompositor struct {
	Proxy
}

func NewWlCompositor(p Proxy) WlCompositor {
	return WlCompositor{Proxy: p}
}
func (w *WlCompositor) AttachListener(l WlCompositorListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

// WlCompositorCreateSurfaceRequest create new surface
//
//	Ask the compositor to create a new surface.
type WlCompositorCreateSurfaceRequest struct {
	// Id the new surface
	Id uint32 `wayland:"new_id"`
}

// WlCompositorCreateRegionRequest create new region
//
//	Ask the compositor to create a new region.
type WlCompositorCreateRegionRequest struct {
	// Id the new region
	Id uint32 `wayland:"new_id"`
}
type WlCompositorListener interface{}
type UnimplementedWlCompositorListener struct{}
type WlCompositorDispatcher struct {
	WlCompositorListener
}

func NewWlCompositorDispatcher() *WlCompositorDispatcher {
	return &WlCompositorDispatcher{WlCompositorListener: &UnimplementedWlCompositorListener{}}
}
func (i *WlCompositorDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlCompositorDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlCompositorListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlCompositorListener = listener
}

// WlShmPool a shared memory pool
/*
	The wl_shm_pool object encapsulates a piece of memory shared
	between the compositor and client.  Through the wl_shm_pool
	object, the client can allocate shared memory wl_buffer objects.
	All objects created through the same pool share the same
	underlying mapped memory. Reusing the mapped memory avoids the
	setup/teardown overhead and is useful when interactively resizing
	a surface or for many small buffers.
*/
type WlShmPool struct {
	Proxy
}

func NewWlShmPool(p Proxy) WlShmPool {
	return WlShmPool{Proxy: p}
}
func (w *WlShmPool) AttachListener(l WlShmPoolListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

// WlShmPoolCreateBufferRequest create a buffer from the pool
/*
	Create a wl_buffer object from the pool.

	The buffer is created offset bytes into the pool and has
	width and height as specified.  The stride argument specifies
	the number of bytes from the beginning of one row to the beginning
	of the next.  The format is the pixel format of the buffer and
	must be one of those advertised through the wl_shm.format event.

	A buffer will keep a reference to the pool it was created from
	so it is valid to destroy the pool immediately after creating
	a buffer from it.
*/
type WlShmPoolCreateBufferRequest struct {
	// Id buffer to create
	Id uint32 `wayland:"new_id"`
	// Offset buffer byte offset within the pool
	Offset int32 `wayland:"int"`
	// Width buffer width, in pixels
	Width int32 `wayland:"int"`
	// Height buffer height, in pixels
	Height int32 `wayland:"int"`
	// Stride number of bytes from the beginning of one row to the beginning of the next row
	Stride int32 `wayland:"int"`
	// Format buffer pixel format
	Format uint32 `wayland:"uint"`
}

// WlShmPoolDestroyRequest destroy the pool
/*
	Destroy the shared memory pool.

	The mmapped memory will be released when all
	buffers that have been created from this pool
	are gone.
*/
type WlShmPoolDestroyRequest struct{}

// WlShmPoolResizeRequest change the size of the pool mapping
/*
	This request will cause the server to remap the backing memory
	for the pool from the file descriptor passed when the pool was
	created, but using the new size.  This request can only be
	used to make the pool bigger.

	This request only changes the amount of bytes that are mmapped
	by the server and does not touch the file corresponding to the
	file descriptor passed at creation time. It is the client's
	responsibility to ensure that the file is at least as big as
	the new pool size.
*/
type WlShmPoolResizeRequest struct {
	// Size new size of the pool, in bytes
	Size int32 `wayland:"int"`
}
type WlShmPoolListener interface{}
type UnimplementedWlShmPoolListener struct{}
type WlShmPoolDispatcher struct {
	WlShmPoolListener
}

func NewWlShmPoolDispatcher() *WlShmPoolDispatcher {
	return &WlShmPoolDispatcher{WlShmPoolListener: &UnimplementedWlShmPoolListener{}}
}
func (i *WlShmPoolDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlShmPoolDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlShmPoolListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlShmPoolListener = listener
}

// WlShm shared memory support
/*
	A singleton global object that provides support for shared
	memory.

	Clients can create wl_shm_pool objects using the create_pool
	request.

	On binding the wl_shm object one or more format events
	are emitted to inform clients about the valid pixel formats
	that can be used for buffers.
*/
type WlShm struct {
	Proxy
}

func NewWlShm(p Proxy) WlShm {
	return WlShm{Proxy: p}
}
func (w *WlShm) AttachListener(l WlShmListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

// WlShmError wl_shm error values
//
//	These errors can be emitted in response to wl_shm requests.
type WlShmError uint32

const (
	// WlShmError_InvalidFormat buffer format is not known
	WlShmError_InvalidFormat WlShmError = 0
	// WlShmError_InvalidStride invalid size or stride during pool or buffer creation
	WlShmError_InvalidStride WlShmError = 1
	// WlShmError_InvalidFd mmapping the file descriptor failed
	WlShmError_InvalidFd WlShmError = 2
)

// WlShmFormat pixel formats
/*
	This describes the memory layout of an individual pixel.

	All renderers should support argb8888 and xrgb8888 but any other
	formats are optional and may not be supported by the particular
	renderer in use.

	The drm format codes match the macros defined in drm_fourcc.h, except
	argb8888 and xrgb8888. The formats actually supported by the compositor
	will be reported by the format event.

	For all wl_shm formats and unless specified in another protocol
	extension, pre-multiplied alpha is used for pixel values.
*/
type WlShmFormat uint32

const (
	// WlShmFormat_Argb8888 32-bit ARGB format, [31:0] A:R:G:B 8:8:8:8 little endian
	WlShmFormat_Argb8888 WlShmFormat = 0
	// WlShmFormat_Xrgb8888 32-bit RGB format, [31:0] x:R:G:B 8:8:8:8 little endian
	WlShmFormat_Xrgb8888 WlShmFormat = 1
	// WlShmFormat_C8 8-bit color index format, [7:0] C
	WlShmFormat_C8 WlShmFormat = 0x20203843
	// WlShmFormat_Rgb332 8-bit RGB format, [7:0] R:G:B 3:3:2
	WlShmFormat_Rgb332 WlShmFormat = 0x38424752
	// WlShmFormat_Bgr233 8-bit BGR format, [7:0] B:G:R 2:3:3
	WlShmFormat_Bgr233 WlShmFormat = 0x38524742
	// WlShmFormat_Xrgb4444 16-bit xRGB format, [15:0] x:R:G:B 4:4:4:4 little endian
	WlShmFormat_Xrgb4444 WlShmFormat = 0x32315258
	// WlShmFormat_Xbgr4444 16-bit xBGR format, [15:0] x:B:G:R 4:4:4:4 little endian
	WlShmFormat_Xbgr4444 WlShmFormat = 0x32314258
	// WlShmFormat_Rgbx4444 16-bit RGBx format, [15:0] R:G:B:x 4:4:4:4 little endian
	WlShmFormat_Rgbx4444 WlShmFormat = 0x32315852
	// WlShmFormat_Bgrx4444 16-bit BGRx format, [15:0] B:G:R:x 4:4:4:4 little endian
	WlShmFormat_Bgrx4444 WlShmFormat = 0x32315842
	// WlShmFormat_Argb4444 16-bit ARGB format, [15:0] A:R:G:B 4:4:4:4 little endian
	WlShmFormat_Argb4444 WlShmFormat = 0x32315241
	// WlShmFormat_Abgr4444 16-bit ABGR format, [15:0] A:B:G:R 4:4:4:4 little endian
	WlShmFormat_Abgr4444 WlShmFormat = 0x32314241
	// WlShmFormat_Rgba4444 16-bit RBGA format, [15:0] R:G:B:A 4:4:4:4 little endian
	WlShmFormat_Rgba4444 WlShmFormat = 0x32314152
	// WlShmFormat_Bgra4444 16-bit BGRA format, [15:0] B:G:R:A 4:4:4:4 little endian
	WlShmFormat_Bgra4444 WlShmFormat = 0x32314142
	// WlShmFormat_Xrgb1555 16-bit xRGB format, [15:0] x:R:G:B 1:5:5:5 little endian
	WlShmFormat_Xrgb1555 WlShmFormat = 0x35315258
	// WlShmFormat_Xbgr1555 16-bit xBGR 1555 format, [15:0] x:B:G:R 1:5:5:5 little endian
	WlShmFormat_Xbgr1555 WlShmFormat = 0x35314258
	// WlShmFormat_Rgbx5551 16-bit RGBx 5551 format, [15:0] R:G:B:x 5:5:5:1 little endian
	WlShmFormat_Rgbx5551 WlShmFormat = 0x35315852
	// WlShmFormat_Bgrx5551 16-bit BGRx 5551 format, [15:0] B:G:R:x 5:5:5:1 little endian
	WlShmFormat_Bgrx5551 WlShmFormat = 0x35315842
	// WlShmFormat_Argb1555 16-bit ARGB 1555 format, [15:0] A:R:G:B 1:5:5:5 little endian
	WlShmFormat_Argb1555 WlShmFormat = 0x35315241
	// WlShmFormat_Abgr1555 16-bit ABGR 1555 format, [15:0] A:B:G:R 1:5:5:5 little endian
	WlShmFormat_Abgr1555 WlShmFormat = 0x35314241
	// WlShmFormat_Rgba5551 16-bit RGBA 5551 format, [15:0] R:G:B:A 5:5:5:1 little endian
	WlShmFormat_Rgba5551 WlShmFormat = 0x35314152
	// WlShmFormat_Bgra5551 16-bit BGRA 5551 format, [15:0] B:G:R:A 5:5:5:1 little endian
	WlShmFormat_Bgra5551 WlShmFormat = 0x35314142
	// WlShmFormat_Rgb565 16-bit RGB 565 format, [15:0] R:G:B 5:6:5 little endian
	WlShmFormat_Rgb565 WlShmFormat = 0x36314752
	// WlShmFormat_Bgr565 16-bit BGR 565 format, [15:0] B:G:R 5:6:5 little endian
	WlShmFormat_Bgr565 WlShmFormat = 0x36314742
	// WlShmFormat_Rgb888 24-bit RGB format, [23:0] R:G:B little endian
	WlShmFormat_Rgb888 WlShmFormat = 0x34324752
	// WlShmFormat_Bgr888 24-bit BGR format, [23:0] B:G:R little endian
	WlShmFormat_Bgr888 WlShmFormat = 0x34324742
	// WlShmFormat_Xbgr8888 32-bit xBGR format, [31:0] x:B:G:R 8:8:8:8 little endian
	WlShmFormat_Xbgr8888 WlShmFormat = 0x34324258
	// WlShmFormat_Rgbx8888 32-bit RGBx format, [31:0] R:G:B:x 8:8:8:8 little endian
	WlShmFormat_Rgbx8888 WlShmFormat = 0x34325852
	// WlShmFormat_Bgrx8888 32-bit BGRx format, [31:0] B:G:R:x 8:8:8:8 little endian
	WlShmFormat_Bgrx8888 WlShmFormat = 0x34325842
	// WlShmFormat_Abgr8888 32-bit ABGR format, [31:0] A:B:G:R 8:8:8:8 little endian
	WlShmFormat_Abgr8888 WlShmFormat = 0x34324241
	// WlShmFormat_Rgba8888 32-bit RGBA format, [31:0] R:G:B:A 8:8:8:8 little endian
	WlShmFormat_Rgba8888 WlShmFormat = 0x34324152
	// WlShmFormat_Bgra8888 32-bit BGRA format, [31:0] B:G:R:A 8:8:8:8 little endian
	WlShmFormat_Bgra8888 WlShmFormat = 0x34324142
	// WlShmFormat_Xrgb2101010 32-bit xRGB format, [31:0] x:R:G:B 2:10:10:10 little endian
	WlShmFormat_Xrgb2101010 WlShmFormat = 0x30335258
	// WlShmFormat_Xbgr2101010 32-bit xBGR format, [31:0] x:B:G:R 2:10:10:10 little endian
	WlShmFormat_Xbgr2101010 WlShmFormat = 0x30334258
	// WlShmFormat_Rgbx1010102 32-bit RGBx format, [31:0] R:G:B:x 10:10:10:2 little endian
	WlShmFormat_Rgbx1010102 WlShmFormat = 0x30335852
	// WlShmFormat_Bgrx1010102 32-bit BGRx format, [31:0] B:G:R:x 10:10:10:2 little endian
	WlShmFormat_Bgrx1010102 WlShmFormat = 0x30335842
	// WlShmFormat_Argb2101010 32-bit ARGB format, [31:0] A:R:G:B 2:10:10:10 little endian
	WlShmFormat_Argb2101010 WlShmFormat = 0x30335241
	// WlShmFormat_Abgr2101010 32-bit ABGR format, [31:0] A:B:G:R 2:10:10:10 little endian
	WlShmFormat_Abgr2101010 WlShmFormat = 0x30334241
	// WlShmFormat_Rgba1010102 32-bit RGBA format, [31:0] R:G:B:A 10:10:10:2 little endian
	WlShmFormat_Rgba1010102 WlShmFormat = 0x30334152
	// WlShmFormat_Bgra1010102 32-bit BGRA format, [31:0] B:G:R:A 10:10:10:2 little endian
	WlShmFormat_Bgra1010102 WlShmFormat = 0x30334142
	// WlShmFormat_Yuyv packed YCbCr format, [31:0] Cr0:Y1:Cb0:Y0 8:8:8:8 little endian
	WlShmFormat_Yuyv WlShmFormat = 0x56595559
	// WlShmFormat_Yvyu packed YCbCr format, [31:0] Cb0:Y1:Cr0:Y0 8:8:8:8 little endian
	WlShmFormat_Yvyu WlShmFormat = 0x55595659
	// WlShmFormat_Uyvy packed YCbCr format, [31:0] Y1:Cr0:Y0:Cb0 8:8:8:8 little endian
	WlShmFormat_Uyvy WlShmFormat = 0x59565955
	// WlShmFormat_Vyuy packed YCbCr format, [31:0] Y1:Cb0:Y0:Cr0 8:8:8:8 little endian
	WlShmFormat_Vyuy WlShmFormat = 0x59555956
	// WlShmFormat_Ayuv packed AYCbCr format, [31:0] A:Y:Cb:Cr 8:8:8:8 little endian
	WlShmFormat_Ayuv WlShmFormat = 0x56555941
	// WlShmFormat_Nv12 2 plane YCbCr Cr:Cb format, 2x2 subsampled Cr:Cb plane
	WlShmFormat_Nv12 WlShmFormat = 0x3231564e
	// WlShmFormat_Nv21 2 plane YCbCr Cb:Cr format, 2x2 subsampled Cb:Cr plane
	WlShmFormat_Nv21 WlShmFormat = 0x3132564e
	// WlShmFormat_Nv16 2 plane YCbCr Cr:Cb format, 2x1 subsampled Cr:Cb plane
	WlShmFormat_Nv16 WlShmFormat = 0x3631564e
	// WlShmFormat_Nv61 2 plane YCbCr Cb:Cr format, 2x1 subsampled Cb:Cr plane
	WlShmFormat_Nv61 WlShmFormat = 0x3136564e
	// WlShmFormat_Yuv410 3 plane YCbCr format, 4x4 subsampled Cb (1) and Cr (2) planes
	WlShmFormat_Yuv410 WlShmFormat = 0x39565559
	// WlShmFormat_Yvu410 3 plane YCbCr format, 4x4 subsampled Cr (1) and Cb (2) planes
	WlShmFormat_Yvu410 WlShmFormat = 0x39555659
	// WlShmFormat_Yuv411 3 plane YCbCr format, 4x1 subsampled Cb (1) and Cr (2) planes
	WlShmFormat_Yuv411 WlShmFormat = 0x31315559
	// WlShmFormat_Yvu411 3 plane YCbCr format, 4x1 subsampled Cr (1) and Cb (2) planes
	WlShmFormat_Yvu411 WlShmFormat = 0x31315659
	// WlShmFormat_Yuv420 3 plane YCbCr format, 2x2 subsampled Cb (1) and Cr (2) planes
	WlShmFormat_Yuv420 WlShmFormat = 0x32315559
	// WlShmFormat_Yvu420 3 plane YCbCr format, 2x2 subsampled Cr (1) and Cb (2) planes
	WlShmFormat_Yvu420 WlShmFormat = 0x32315659
	// WlShmFormat_Yuv422 3 plane YCbCr format, 2x1 subsampled Cb (1) and Cr (2) planes
	WlShmFormat_Yuv422 WlShmFormat = 0x36315559
	// WlShmFormat_Yvu422 3 plane YCbCr format, 2x1 subsampled Cr (1) and Cb (2) planes
	WlShmFormat_Yvu422 WlShmFormat = 0x36315659
	// WlShmFormat_Yuv444 3 plane YCbCr format, non-subsampled Cb (1) and Cr (2) planes
	WlShmFormat_Yuv444 WlShmFormat = 0x34325559
	// WlShmFormat_Yvu444 3 plane YCbCr format, non-subsampled Cr (1) and Cb (2) planes
	WlShmFormat_Yvu444 WlShmFormat = 0x34325659
	// WlShmFormat_R8 [7:0] R
	WlShmFormat_R8 WlShmFormat = 0x20203852
	// WlShmFormat_R16 [15:0] R little endian
	WlShmFormat_R16 WlShmFormat = 0x20363152
	// WlShmFormat_Rg88 [15:0] R:G 8:8 little endian
	WlShmFormat_Rg88 WlShmFormat = 0x38384752
	// WlShmFormat_Gr88 [15:0] G:R 8:8 little endian
	WlShmFormat_Gr88 WlShmFormat = 0x38385247
	// WlShmFormat_Rg1616 [31:0] R:G 16:16 little endian
	WlShmFormat_Rg1616 WlShmFormat = 0x32334752
	// WlShmFormat_Gr1616 [31:0] G:R 16:16 little endian
	WlShmFormat_Gr1616 WlShmFormat = 0x32335247
	// WlShmFormat_Xrgb16161616f [63:0] x:R:G:B 16:16:16:16 little endian
	WlShmFormat_Xrgb16161616f WlShmFormat = 0x48345258
	// WlShmFormat_Xbgr16161616f [63:0] x:B:G:R 16:16:16:16 little endian
	WlShmFormat_Xbgr16161616f WlShmFormat = 0x48344258
	// WlShmFormat_Argb16161616f [63:0] A:R:G:B 16:16:16:16 little endian
	WlShmFormat_Argb16161616f WlShmFormat = 0x48345241
	// WlShmFormat_Abgr16161616f [63:0] A:B:G:R 16:16:16:16 little endian
	WlShmFormat_Abgr16161616f WlShmFormat = 0x48344241
	// WlShmFormat_Xyuv8888 [31:0] X:Y:Cb:Cr 8:8:8:8 little endian
	WlShmFormat_Xyuv8888 WlShmFormat = 0x56555958
	// WlShmFormat_Vuy888 [23:0] Cr:Cb:Y 8:8:8 little endian
	WlShmFormat_Vuy888 WlShmFormat = 0x34325556
	// WlShmFormat_Vuy101010 Y followed by U then V, 10:10:10. Non-linear modifier only
	WlShmFormat_Vuy101010 WlShmFormat = 0x30335556
	// WlShmFormat_Y210 [63:0] Cr0:0:Y1:0:Cb0:0:Y0:0 10:6:10:6:10:6:10:6 little endian per 2 Y pixels
	WlShmFormat_Y210 WlShmFormat = 0x30313259
	// WlShmFormat_Y212 [63:0] Cr0:0:Y1:0:Cb0:0:Y0:0 12:4:12:4:12:4:12:4 little endian per 2 Y pixels
	WlShmFormat_Y212 WlShmFormat = 0x32313259
	// WlShmFormat_Y216 [63:0] Cr0:Y1:Cb0:Y0 16:16:16:16 little endian per 2 Y pixels
	WlShmFormat_Y216 WlShmFormat = 0x36313259
	// WlShmFormat_Y410 [31:0] A:Cr:Y:Cb 2:10:10:10 little endian
	WlShmFormat_Y410 WlShmFormat = 0x30313459
	// WlShmFormat_Y412 [63:0] A:0:Cr:0:Y:0:Cb:0 12:4:12:4:12:4:12:4 little endian
	WlShmFormat_Y412 WlShmFormat = 0x32313459
	// WlShmFormat_Y416 [63:0] A:Cr:Y:Cb 16:16:16:16 little endian
	WlShmFormat_Y416 WlShmFormat = 0x36313459
	// WlShmFormat_Xvyu2101010 [31:0] X:Cr:Y:Cb 2:10:10:10 little endian
	WlShmFormat_Xvyu2101010 WlShmFormat = 0x30335658
	// WlShmFormat_Xvyu1216161616 [63:0] X:0:Cr:0:Y:0:Cb:0 12:4:12:4:12:4:12:4 little endian
	WlShmFormat_Xvyu1216161616 WlShmFormat = 0x36335658
	// WlShmFormat_Xvyu16161616 [63:0] X:Cr:Y:Cb 16:16:16:16 little endian
	WlShmFormat_Xvyu16161616 WlShmFormat = 0x38345658
	// WlShmFormat_Y0l0 [63:0]   A3:A2:Y3:0:Cr0:0:Y2:0:A1:A0:Y1:0:Cb0:0:Y0:0  1:1:8:2:8:2:8:2:1:1:8:2:8:2:8:2 little endian
	WlShmFormat_Y0l0 WlShmFormat = 0x304c3059
	// WlShmFormat_X0l0 [63:0]   X3:X2:Y3:0:Cr0:0:Y2:0:X1:X0:Y1:0:Cb0:0:Y0:0  1:1:8:2:8:2:8:2:1:1:8:2:8:2:8:2 little endian
	WlShmFormat_X0l0 WlShmFormat = 0x304c3058
	// WlShmFormat_Y0l2 [63:0]   A3:A2:Y3:Cr0:Y2:A1:A0:Y1:Cb0:Y0  1:1:10:10:10:1:1:10:10:10 little endian
	WlShmFormat_Y0l2 WlShmFormat = 0x324c3059
	// WlShmFormat_X0l2 [63:0]   X3:X2:Y3:Cr0:Y2:X1:X0:Y1:Cb0:Y0  1:1:10:10:10:1:1:10:10:10 little endian
	WlShmFormat_X0l2        WlShmFormat = 0x324c3058
	WlShmFormat_Yuv4208Bit  WlShmFormat = 0x38305559
	WlShmFormat_Yuv42010Bit WlShmFormat = 0x30315559
	WlShmFormat_Xrgb8888A8  WlShmFormat = 0x38415258
	WlShmFormat_Xbgr8888A8  WlShmFormat = 0x38414258
	WlShmFormat_Rgbx8888A8  WlShmFormat = 0x38415852
	WlShmFormat_Bgrx8888A8  WlShmFormat = 0x38415842
	WlShmFormat_Rgb888A8    WlShmFormat = 0x38413852
	WlShmFormat_Bgr888A8    WlShmFormat = 0x38413842
	WlShmFormat_Rgb565A8    WlShmFormat = 0x38413552
	WlShmFormat_Bgr565A8    WlShmFormat = 0x38413542
	// WlShmFormat_Nv24 non-subsampled Cr:Cb plane
	WlShmFormat_Nv24 WlShmFormat = 0x3432564e
	// WlShmFormat_Nv42 non-subsampled Cb:Cr plane
	WlShmFormat_Nv42 WlShmFormat = 0x3234564e
	// WlShmFormat_P210 2x1 subsampled Cr:Cb plane, 10 bit per channel
	WlShmFormat_P210 WlShmFormat = 0x30313250
	// WlShmFormat_P010 2x2 subsampled Cr:Cb plane 10 bits per channel
	WlShmFormat_P010 WlShmFormat = 0x30313050
	// WlShmFormat_P012 2x2 subsampled Cr:Cb plane 12 bits per channel
	WlShmFormat_P012 WlShmFormat = 0x32313050
	// WlShmFormat_P016 2x2 subsampled Cr:Cb plane 16 bits per channel
	WlShmFormat_P016 WlShmFormat = 0x36313050
	// WlShmFormat_Axbxgxrx106106106106 [63:0] A:x:B:x:G:x:R:x 10:6:10:6:10:6:10:6 little endian
	WlShmFormat_Axbxgxrx106106106106 WlShmFormat = 0x30314241
	// WlShmFormat_Nv15 2x2 subsampled Cr:Cb plane
	WlShmFormat_Nv15 WlShmFormat = 0x3531564e
	WlShmFormat_Q410 WlShmFormat = 0x30313451
	WlShmFormat_Q401 WlShmFormat = 0x31303451
	// WlShmFormat_Xrgb16161616 [63:0] x:R:G:B 16:16:16:16 little endian
	WlShmFormat_Xrgb16161616 WlShmFormat = 0x38345258
	// WlShmFormat_Xbgr16161616 [63:0] x:B:G:R 16:16:16:16 little endian
	WlShmFormat_Xbgr16161616 WlShmFormat = 0x38344258
	// WlShmFormat_Argb16161616 [63:0] A:R:G:B 16:16:16:16 little endian
	WlShmFormat_Argb16161616 WlShmFormat = 0x38345241
	// WlShmFormat_Abgr16161616 [63:0] A:B:G:R 16:16:16:16 little endian
	WlShmFormat_Abgr16161616 WlShmFormat = 0x38344241
)

// WlShmCreatePoolRequest create a shm pool
/*
	Create a new wl_shm_pool object.

	The pool can be used to create shared memory based buffer
	objects.  The server will mmap size bytes of the passed file
	descriptor, to use as backing memory for the pool.
*/
type WlShmCreatePoolRequest struct {
	// Id pool to create
	Id uint32 `wayland:"new_id"`
	// Fd file descriptor for the pool
	Fd uintptr `wayland:"fd"`
	// Size pool size, in bytes
	Size int32 `wayland:"int"`
}

// WlShmFormatEvent pixel format description
/*
	Informs the client about a valid pixel format that
	can be used for buffers. Known formats include
	argb8888 and xrgb8888.
*/
type WlShmFormatEvent struct {
	// Format buffer pixel format
	Format uint32 `wayland:"uint"`
}

func (e *WlShmFormatEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Format = UnmarshallUint32(offset, d)
	return nil
}

type WlShmListener interface {
	Format(WlShmFormatEvent)
}
type UnimplementedWlShmListener struct{}

func (e *UnimplementedWlShmListener) Format(_ WlShmFormatEvent) {
	return
}

type WlShmDispatcher struct {
	WlShmListener
}

func NewWlShmDispatcher() *WlShmDispatcher {
	return &WlShmDispatcher{WlShmListener: &UnimplementedWlShmListener{}}
}
func (i *WlShmDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	case 0:
		var e WlShmFormatEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Format(e)
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlShmDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlShmListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlShmListener = listener
}

// WlBuffer content for a wl_surface
/*
	A buffer provides the content for a wl_surface. Buffers are
	created through factory interfaces such as wl_shm, wp_linux_buffer_params
	(from the linux-dmabuf protocol extension) or similar. It has a width and
	a height and can be attached to a wl_surface, but the mechanism by which a
	client provides and updates the contents is defined by the buffer factory
	interface.

	If the buffer uses a format that has an alpha channel, the alpha channel
	is assumed to be premultiplied in the color channels unless otherwise
	specified.
*/
type WlBuffer struct {
	Proxy
}

func NewWlBuffer(p Proxy) WlBuffer {
	return WlBuffer{Proxy: p}
}
func (w *WlBuffer) AttachListener(l WlBufferListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

// WlBufferDestroyRequest destroy a buffer
/*
	Destroy a buffer. If and how you need to release the backing
	storage is defined by the buffer factory interface.

	For possible side-effects to a surface, see wl_surface.attach.
*/
type WlBufferDestroyRequest struct{}

// WlBufferReleaseEvent compositor releases buffer
/*
	Sent when this wl_buffer is no longer used by the compositor.
	The client is now free to reuse or destroy this buffer and its
	backing storage.

	If a client receives a release event before the frame callback
	requested in the same wl_surface.commit that attaches this
	wl_buffer to a surface, then the client is immediately free to
	reuse the buffer and its backing storage, and does not need a
	second buffer for the next surface content update. Typically
	this is possible, when the compositor maintains a copy of the
	wl_surface contents, e.g. as a GL texture. This is an important
	optimization for GL(ES) compositors with wl_shm clients.
*/
type WlBufferReleaseEvent struct{}

func (e *WlBufferReleaseEvent) UnmarshallBinary(d []byte) error {
	return nil
}

type WlBufferListener interface {
	Release(WlBufferReleaseEvent)
}
type UnimplementedWlBufferListener struct{}

func (e *UnimplementedWlBufferListener) Release(_ WlBufferReleaseEvent) {
	return
}

type WlBufferDispatcher struct {
	WlBufferListener
}

func NewWlBufferDispatcher() *WlBufferDispatcher {
	return &WlBufferDispatcher{WlBufferListener: &UnimplementedWlBufferListener{}}
}
func (i *WlBufferDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	case 0:
		var e WlBufferReleaseEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Release(e)
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlBufferDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlBufferListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlBufferListener = listener
}

// WlDataOffer offer to transfer data
/*
	A wl_data_offer represents a piece of data offered for transfer
	by another client (the source client).  It is used by the
	copy-and-paste and drag-and-drop mechanisms.  The offer
	describes the different mime types that the data can be
	converted to and provides the mechanism for transferring the
	data directly from the source client.
*/
type WlDataOffer struct {
	Proxy
}

func NewWlDataOffer(p Proxy) WlDataOffer {
	return WlDataOffer{Proxy: p}
}
func (w *WlDataOffer) AttachListener(l WlDataOfferListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

type WlDataOfferError uint32

const (
	// WlDataOfferError_InvalidFinish finish request was called untimely
	WlDataOfferError_InvalidFinish WlDataOfferError = 0
	// WlDataOfferError_InvalidActionMask action mask contains invalid values
	WlDataOfferError_InvalidActionMask WlDataOfferError = 1
	// WlDataOfferError_InvalidAction action argument has an invalid value
	WlDataOfferError_InvalidAction WlDataOfferError = 2
	// WlDataOfferError_InvalidOffer offer doesn't accept this request
	WlDataOfferError_InvalidOffer WlDataOfferError = 3
)

// WlDataOfferAcceptRequest accept one of the offered mime types
/*
	Indicate that the client can accept the given mime type, or
	NULL for not accepted.

	For objects of version 2 or older, this request is used by the
	client to give feedback whether the client can receive the given
	mime type, or NULL if none is accepted; the feedback does not
	determine whether the drag-and-drop operation succeeds or not.

	For objects of version 3 or newer, this request determines the
	final result of the drag-and-drop operation. If the end result
	is that no mime types were accepted, the drag-and-drop operation
	will be cancelled and the corresponding drag source will receive
	wl_data_source.cancelled. Clients may still use this event in
	conjunction with wl_data_source.action for feedback.
*/
type WlDataOfferAcceptRequest struct {
	// Serial  number of the accept request
	Serial uint32 `wayland:"uint"`
	// MimeType mime type accepted by the client
	MimeType string `wayland:"string"`
}

// WlDataOfferReceiveRequest request that the data is transferred
/*
	To transfer the offered data, the client issues this request
	and indicates the mime type it wants to receive.  The transfer
	happens through the passed file descriptor (typically created
	with the pipe system call).  The source client writes the data
	in the mime type representation requested and then closes the
	file descriptor.

	The receiving client reads from the read end of the pipe until
	EOF and then closes its end, at which point the transfer is
	complete.

	This request may happen multiple times for different mime types,
	both before and after wl_data_device.drop. Drag-and-drop destination
	clients may preemptively fetch data or examine it more closely to
	determine acceptance.
*/
type WlDataOfferReceiveRequest struct {
	// MimeType mime type desired by receiver
	MimeType string `wayland:"string"`
	// Fd file descriptor for data transfer
	Fd uintptr `wayland:"fd"`
}

// WlDataOfferDestroyRequest destroy data offer
//
//	Destroy the data offer.
type WlDataOfferDestroyRequest struct{}

// WlDataOfferFinishRequest the offer will no longer be used
/*
	Notifies the compositor that the drag destination successfully
	finished the drag-and-drop operation.

	Upon receiving this request, the compositor will emit
	wl_data_source.dnd_finished on the drag source client.

	It is a client error to perform other requests than
	wl_data_offer.destroy after this one. It is also an error to perform
	this request after a NULL mime type has been set in
	wl_data_offer.accept or no action was received through
	wl_data_offer.action.

	If wl_data_offer.finish request is received for a non drag and drop
	operation, the invalid_finish protocol error is raised.
*/
type WlDataOfferFinishRequest struct{}

// WlDataOfferSetActionsRequest set the available/preferred drag-and-drop actions
/*
	Sets the actions that the destination side client supports for
	this operation. This request may trigger the emission of
	wl_data_source.action and wl_data_offer.action events if the compositor
	needs to change the selected action.

	This request can be called multiple times throughout the
	drag-and-drop operation, typically in response to wl_data_device.enter
	or wl_data_device.motion events.

	This request determines the final result of the drag-and-drop
	operation. If the end result is that no action is accepted,
	the drag source will receive wl_data_source.cancelled.

	The dnd_actions argument must contain only values expressed in the
	wl_data_device_manager.dnd_actions enum, and the preferred_action
	argument must only contain one of those values set, otherwise it
	will result in a protocol error.

	While managing an "ask" action, the destination drag-and-drop client
	may perform further wl_data_offer.receive requests, and is expected
	to perform one last wl_data_offer.set_actions request with a preferred
	action other than "ask" (and optionally wl_data_offer.accept) before
	requesting wl_data_offer.finish, in order to convey the action selected
	by the user. If the preferred action is not in the
	wl_data_offer.source_actions mask, an error will be raised.

	If the "ask" action is dismissed (e.g. user cancellation), the client
	is expected to perform wl_data_offer.destroy right away.

	This request can only be made on drag-and-drop offers, a protocol error
	will be raised otherwise.
*/
type WlDataOfferSetActionsRequest struct {
	// DndActions actions supported by the destination client
	DndActions uint32 `wayland:"uint"`
	// PreferredAction action preferred by the destination client
	PreferredAction uint32 `wayland:"uint"`
}

// WlDataOfferOfferEvent advertise offered mime type
/*
	Sent immediately after creating the wl_data_offer object.  One
	event per offered mime type.
*/
type WlDataOfferOfferEvent struct {
	// MimeType offered mime type
	MimeType string `wayland:"string"`
}

func (e *WlDataOfferOfferEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.MimeType = UnmarshallString(offset, d)
	return nil
}

// WlDataOfferSourceActionsEvent notify the source-side available actions
/*
	This event indicates the actions offered by the data source. It
	will be sent right after wl_data_device.enter, or anytime the source
	side changes its offered actions through wl_data_source.set_actions.
*/
type WlDataOfferSourceActionsEvent struct {
	// SourceActions actions offered by the data source
	SourceActions uint32 `wayland:"uint"`
}

func (e *WlDataOfferSourceActionsEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.SourceActions = UnmarshallUint32(offset, d)
	return nil
}

// WlDataOfferActionEvent notify the selected action
/*
	This event indicates the action selected by the compositor after
	matching the source/destination side actions. Only one action (or
	none) will be offered here.

	This event can be emitted multiple times during the drag-and-drop
	operation in response to destination side action changes through
	wl_data_offer.set_actions.

	This event will no longer be emitted after wl_data_device.drop
	happened on the drag-and-drop destination, the client must
	honor the last action received, or the last preferred one set
	through wl_data_offer.set_actions when handling an "ask" action.

	Compositors may also change the selected action on the fly, mainly
	in response to keyboard modifier changes during the drag-and-drop
	operation.

	The most recent action received is always the valid one. Prior to
	receiving wl_data_device.drop, the chosen action may change (e.g.
	due to keyboard modifiers being pressed). At the time of receiving
	wl_data_device.drop the drag-and-drop destination must honor the
	last action received.

	Action changes may still happen after wl_data_device.drop,
	especially on "ask" actions, where the drag-and-drop destination
	may choose another action afterwards. Action changes happening
	at this stage are always the result of inter-client negotiation, the
	compositor shall no longer be able to induce a different action.

	Upon "ask" actions, it is expected that the drag-and-drop destination
	may potentially choose a different action and/or mime type,
	based on wl_data_offer.source_actions and finally chosen by the
	user (e.g. popping up a menu with the available options). The
	final wl_data_offer.set_actions and wl_data_offer.accept requests
	must happen before the call to wl_data_offer.finish.
*/
type WlDataOfferActionEvent struct {
	// DndAction action selected by the compositor
	DndAction uint32 `wayland:"uint"`
}

func (e *WlDataOfferActionEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.DndAction = UnmarshallUint32(offset, d)
	return nil
}

type WlDataOfferListener interface {
	Offer(WlDataOfferOfferEvent)
	SourceActions(WlDataOfferSourceActionsEvent)
	Action(WlDataOfferActionEvent)
}
type UnimplementedWlDataOfferListener struct{}

func (e *UnimplementedWlDataOfferListener) Offer(_ WlDataOfferOfferEvent) {
	return
}
func (e *UnimplementedWlDataOfferListener) SourceActions(_ WlDataOfferSourceActionsEvent) {
	return
}
func (e *UnimplementedWlDataOfferListener) Action(_ WlDataOfferActionEvent) {
	return
}

type WlDataOfferDispatcher struct {
	WlDataOfferListener
}

func NewWlDataOfferDispatcher() *WlDataOfferDispatcher {
	return &WlDataOfferDispatcher{WlDataOfferListener: &UnimplementedWlDataOfferListener{}}
}
func (i *WlDataOfferDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	case 0:
		var e WlDataOfferOfferEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Offer(e)
	case 1:
		var e WlDataOfferSourceActionsEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.SourceActions(e)
	case 2:
		var e WlDataOfferActionEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Action(e)
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlDataOfferDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlDataOfferListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlDataOfferListener = listener
}

// WlDataSource offer to transfer data
/*
	The wl_data_source object is the source side of a wl_data_offer.
	It is created by the source client in a data transfer and
	provides a way to describe the offered data and a way to respond
	to requests to transfer the data.
*/
type WlDataSource struct {
	Proxy
}

func NewWlDataSource(p Proxy) WlDataSource {
	return WlDataSource{Proxy: p}
}
func (w *WlDataSource) AttachListener(l WlDataSourceListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

type WlDataSourceError uint32

const (
	// WlDataSourceError_InvalidActionMask action mask contains invalid values
	WlDataSourceError_InvalidActionMask WlDataSourceError = 0
	// WlDataSourceError_InvalidSource source doesn't accept this request
	WlDataSourceError_InvalidSource WlDataSourceError = 1
)

// WlDataSourceOfferRequest add an offered mime type
/*
	This request adds a mime type to the set of mime types
	advertised to targets.  Can be called several times to offer
	multiple types.
*/
type WlDataSourceOfferRequest struct {
	// MimeType mime type offered by the data source
	MimeType string `wayland:"string"`
}

// WlDataSourceDestroyRequest destroy the data source
//
//	Destroy the data source.
type WlDataSourceDestroyRequest struct{}

// WlDataSourceSetActionsRequest set the available drag-and-drop actions
/*
	Sets the actions that the source side client supports for this
	operation. This request may trigger wl_data_source.action and
	wl_data_offer.action events if the compositor needs to change the
	selected action.

	The dnd_actions argument must contain only values expressed in the
	wl_data_device_manager.dnd_actions enum, otherwise it will result
	in a protocol error.

	This request must be made once only, and can only be made on sources
	used in drag-and-drop, so it must be performed before
	wl_data_device.start_drag. Attempting to use the source other than
	for drag-and-drop will raise a protocol error.
*/
type WlDataSourceSetActionsRequest struct {
	// DndActions actions supported by the data source
	DndActions uint32 `wayland:"uint"`
}

// WlDataSourceTargetEvent a target accepts an offered mime type
/*
	Sent when a target accepts pointer_focus or motion events.  If
	a target does not accept any of the offered types, type is NULL.

	Used for feedback during drag-and-drop.
*/
type WlDataSourceTargetEvent struct {
	// MimeType mime type accepted by the target
	MimeType string `wayland:"string"`
}

func (e *WlDataSourceTargetEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.MimeType = UnmarshallString(offset, d)
	return nil
}

// WlDataSourceSendEvent send the data
/*
	Request for data from the client.  Send the data as the
	specified mime type over the passed file descriptor, then
	close it.
*/
type WlDataSourceSendEvent struct {
	// MimeType mime type for the data
	MimeType string `wayland:"string"`
	// Fd file descriptor for the data
	Fd uintptr `wayland:"fd"`
}

func (e *WlDataSourceSendEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.MimeType = UnmarshallString(offset, d)
	offset, e.Fd = UnmarshallFd(offset, d)
	return nil
}

// WlDataSourceCancelledEvent selection was cancelled
/*
	This data source is no longer valid. There are several reasons why
	this could happen:

	- The data source has been replaced by another data source.
	- The drag-and-drop operation was performed, but the drop destination
	did not accept any of the mime types offered through
	wl_data_source.target.
	- The drag-and-drop operation was performed, but the drop destination
	did not select any of the actions present in the mask offered through
	wl_data_source.action.
	- The drag-and-drop operation was performed but didn't happen over a
	surface.
	- The compositor cancelled the drag-and-drop operation (e.g. compositor
	dependent timeouts to avoid stale drag-and-drop transfers).

	The client should clean up and destroy this data source.

	For objects of version 2 or older, wl_data_source.cancelled will
	only be emitted if the data source was replaced by another data
	source.
*/
type WlDataSourceCancelledEvent struct{}

func (e *WlDataSourceCancelledEvent) UnmarshallBinary(d []byte) error {
	return nil
}

// WlDataSourceDndDropPerformedEvent the drag-and-drop operation physically finished
/*
	The user performed the drop action. This event does not indicate
	acceptance, wl_data_source.cancelled may still be emitted afterwards
	if the drop destination does not accept any mime type.

	However, this event might however not be received if the compositor
	cancelled the drag-and-drop operation before this event could happen.

	Note that the data_source may still be used in the future and should
	not be destroyed here.
*/
type WlDataSourceDndDropPerformedEvent struct{}

func (e *WlDataSourceDndDropPerformedEvent) UnmarshallBinary(d []byte) error {
	return nil
}

// WlDataSourceDndFinishedEvent the drag-and-drop operation concluded
/*
	The drop destination finished interoperating with this data
	source, so the client is now free to destroy this data source and
	free all associated data.

	If the action used to perform the operation was "move", the
	source can now delete the transferred data.
*/
type WlDataSourceDndFinishedEvent struct{}

func (e *WlDataSourceDndFinishedEvent) UnmarshallBinary(d []byte) error {
	return nil
}

// WlDataSourceActionEvent notify the selected action
/*
	This event indicates the action selected by the compositor after
	matching the source/destination side actions. Only one action (or
	none) will be offered here.

	This event can be emitted multiple times during the drag-and-drop
	operation, mainly in response to destination side changes through
	wl_data_offer.set_actions, and as the data device enters/leaves
	surfaces.

	It is only possible to receive this event after
	wl_data_source.dnd_drop_performed if the drag-and-drop operation
	ended in an "ask" action, in which case the final wl_data_source.action
	event will happen immediately before wl_data_source.dnd_finished.

	Compositors may also change the selected action on the fly, mainly
	in response to keyboard modifier changes during the drag-and-drop
	operation.

	The most recent action received is always the valid one. The chosen
	action may change alongside negotiation (e.g. an "ask" action can turn
	into a "move" operation), so the effects of the final action must
	always be applied in wl_data_offer.dnd_finished.

	Clients can trigger cursor surface changes from this point, so
	they reflect the current action.
*/
type WlDataSourceActionEvent struct {
	// DndAction action selected by the compositor
	DndAction uint32 `wayland:"uint"`
}

func (e *WlDataSourceActionEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.DndAction = UnmarshallUint32(offset, d)
	return nil
}

type WlDataSourceListener interface {
	Target(WlDataSourceTargetEvent)
	Send(WlDataSourceSendEvent)
	Cancelled(WlDataSourceCancelledEvent)
	DndDropPerformed(WlDataSourceDndDropPerformedEvent)
	DndFinished(WlDataSourceDndFinishedEvent)
	Action(WlDataSourceActionEvent)
}
type UnimplementedWlDataSourceListener struct{}

func (e *UnimplementedWlDataSourceListener) Target(_ WlDataSourceTargetEvent) {
	return
}
func (e *UnimplementedWlDataSourceListener) Send(_ WlDataSourceSendEvent) {
	return
}
func (e *UnimplementedWlDataSourceListener) Cancelled(_ WlDataSourceCancelledEvent) {
	return
}
func (e *UnimplementedWlDataSourceListener) DndDropPerformed(_ WlDataSourceDndDropPerformedEvent) {
	return
}
func (e *UnimplementedWlDataSourceListener) DndFinished(_ WlDataSourceDndFinishedEvent) {
	return
}
func (e *UnimplementedWlDataSourceListener) Action(_ WlDataSourceActionEvent) {
	return
}

type WlDataSourceDispatcher struct {
	WlDataSourceListener
}

func NewWlDataSourceDispatcher() *WlDataSourceDispatcher {
	return &WlDataSourceDispatcher{WlDataSourceListener: &UnimplementedWlDataSourceListener{}}
}
func (i *WlDataSourceDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	case 0:
		var e WlDataSourceTargetEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Target(e)
	case 1:
		var e WlDataSourceSendEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Send(e)
	case 2:
		var e WlDataSourceCancelledEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Cancelled(e)
	case 3:
		var e WlDataSourceDndDropPerformedEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.DndDropPerformed(e)
	case 4:
		var e WlDataSourceDndFinishedEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.DndFinished(e)
	case 5:
		var e WlDataSourceActionEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Action(e)
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlDataSourceDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlDataSourceListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlDataSourceListener = listener
}

// WlDataDevice data transfer device
/*
	There is one wl_data_device per seat which can be obtained
	from the global wl_data_device_manager singleton.

	A wl_data_device provides access to inter-client data transfer
	mechanisms such as copy-and-paste and drag-and-drop.
*/
type WlDataDevice struct {
	Proxy
}

func NewWlDataDevice(p Proxy) WlDataDevice {
	return WlDataDevice{Proxy: p}
}
func (w *WlDataDevice) AttachListener(l WlDataDeviceListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

type WlDataDeviceError uint32

const (
	// WlDataDeviceError_Role given wl_surface has another role
	WlDataDeviceError_Role WlDataDeviceError = 0
)

// WlDataDeviceStartDragRequest start drag-and-drop operation
/*
	This request asks the compositor to start a drag-and-drop
	operation on behalf of the client.

	The source argument is the data source that provides the data
	for the eventual data transfer. If source is NULL, enter, leave
	and motion events are sent only to the client that initiated the
	drag and the client is expected to handle the data passing
	internally. If source is destroyed, the drag-and-drop session will be
	cancelled.

	The origin surface is the surface where the drag originates and
	the client must have an active implicit grab that matches the
	serial.

	The icon surface is an optional (can be NULL) surface that
	provides an icon to be moved around with the cursor.  Initially,
	the top-left corner of the icon surface is placed at the cursor
	hotspot, but subsequent wl_surface.attach request can move the
	relative position. Attach requests must be confirmed with
	wl_surface.commit as usual. The icon surface is given the role of
	a drag-and-drop icon. If the icon surface already has another role,
	it raises a protocol error.

	The current and pending input regions of the icon wl_surface are
	cleared, and wl_surface.set_input_region is ignored until the
	wl_surface is no longer used as the icon surface. When the use
	as an icon ends, the current and pending input regions become
	undefined, and the wl_surface is unmapped.
*/
type WlDataDeviceStartDragRequest struct {
	// Source data source for the eventual transfer
	Source uint32 `wayland:"object"`
	// Origin surface where the drag originates
	Origin uint32 `wayland:"object"`
	// Icon drag-and-drop icon surface
	Icon uint32 `wayland:"object"`
	// Serial  number of the implicit grab on the origin
	Serial uint32 `wayland:"uint"`
}

// WlDataDeviceSetSelectionRequest copy data to the selection
/*
	This request asks the compositor to set the selection
	to the data from the source on behalf of the client.

	To unset the selection, set the source to NULL.
*/
type WlDataDeviceSetSelectionRequest struct {
	// Source data source for the selection
	Source uint32 `wayland:"object"`
	// Serial  number of the event that triggered this request
	Serial uint32 `wayland:"uint"`
}

// WlDataDeviceReleaseRequest destroy data device
//
//	This request destroys the data device.
type WlDataDeviceReleaseRequest struct{}

// WlDataDeviceDataOfferEvent introduce a new wl_data_offer
/*
	The data_offer event introduces a new wl_data_offer object,
	which will subsequently be used in either the
	data_device.enter event (for drag-and-drop) or the
	data_device.selection event (for selections).  Immediately
	following the data_device.data_offer event, the new data_offer
	object will send out data_offer.offer events to describe the
	mime types it offers.
*/
type WlDataDeviceDataOfferEvent struct {
	// Id the new data_offer object
	Id uint32 `wayland:"new_id"`
}

func (e *WlDataDeviceDataOfferEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Id = UnmarshallUint32(offset, d)
	return nil
}

// WlDataDeviceEnterEvent initiate drag-and-drop session
/*
	This event is sent when an active drag-and-drop pointer enters
	a surface owned by the client.  The position of the pointer at
	enter time is provided by the x and y arguments, in surface-local
	coordinates.
*/
type WlDataDeviceEnterEvent struct {
	// Serial  number of the enter event
	Serial uint32 `wayland:"uint"`
	// Surface client surface entered
	Surface uint32 `wayland:"object"`
	// X surface-local x coordinate
	X float64 `wayland:"fixed"`
	// Y surface-local y coordinate
	Y float64 `wayland:"fixed"`
	// Id source data_offer object
	Id uint32 `wayland:"object"`
}

func (e *WlDataDeviceEnterEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Serial = UnmarshallUint32(offset, d)
	offset, e.Surface = UnmarshallUint32(offset, d)
	offset, e.X = UnmarshallFixed(offset, d)
	offset, e.Y = UnmarshallFixed(offset, d)
	offset, e.Id = UnmarshallUint32(offset, d)
	return nil
}

// WlDataDeviceLeaveEvent end drag-and-drop session
/*
	This event is sent when the drag-and-drop pointer leaves the
	surface and the session ends.  The client must destroy the
	wl_data_offer introduced at enter time at this point.
*/
type WlDataDeviceLeaveEvent struct{}

func (e *WlDataDeviceLeaveEvent) UnmarshallBinary(d []byte) error {
	return nil
}

// WlDataDeviceMotionEvent drag-and-drop session motion
/*
	This event is sent when the drag-and-drop pointer moves within
	the currently focused surface. The new position of the pointer
	is provided by the x and y arguments, in surface-local
	coordinates.
*/
type WlDataDeviceMotionEvent struct {
	// Time stamp with millisecond granularity
	Time uint32 `wayland:"uint"`
	// X surface-local x coordinate
	X float64 `wayland:"fixed"`
	// Y surface-local y coordinate
	Y float64 `wayland:"fixed"`
}

func (e *WlDataDeviceMotionEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Time = UnmarshallUint32(offset, d)
	offset, e.X = UnmarshallFixed(offset, d)
	offset, e.Y = UnmarshallFixed(offset, d)
	return nil
}

// WlDataDeviceDropEvent end drag-and-drop session successfully
/*
	The event is sent when a drag-and-drop operation is ended
	because the implicit grab is removed.

	The drag-and-drop destination is expected to honor the last action
	received through wl_data_offer.action, if the resulting action is
	"copy" or "move", the destination can still perform
	wl_data_offer.receive requests, and is expected to end all
	transfers with a wl_data_offer.finish request.

	If the resulting action is "ask", the action will not be considered
	final. The drag-and-drop destination is expected to perform one last
	wl_data_offer.set_actions request, or wl_data_offer.destroy in order
	to cancel the operation.
*/
type WlDataDeviceDropEvent struct{}

func (e *WlDataDeviceDropEvent) UnmarshallBinary(d []byte) error {
	return nil
}

// WlDataDeviceSelectionEvent advertise new selection
/*
	The selection event is sent out to notify the client of a new
	wl_data_offer for the selection for this device.  The
	data_device.data_offer and the data_offer.offer events are
	sent out immediately before this event to introduce the data
	offer object.  The selection event is sent to a client
	immediately before receiving keyboard focus and when a new
	selection is set while the client has keyboard focus.  The
	data_offer is valid until a new data_offer or NULL is received
	or until the client loses keyboard focus.  Switching surface with
	keyboard focus within the same client doesn't mean a new selection
	will be sent.  The client must destroy the previous selection
	data_offer, if any, upon receiving this event.
*/
type WlDataDeviceSelectionEvent struct {
	// Id selection data_offer object
	Id uint32 `wayland:"object"`
}

func (e *WlDataDeviceSelectionEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Id = UnmarshallUint32(offset, d)
	return nil
}

type WlDataDeviceListener interface {
	DataOffer(WlDataDeviceDataOfferEvent)
	Enter(WlDataDeviceEnterEvent)
	Leave(WlDataDeviceLeaveEvent)
	Motion(WlDataDeviceMotionEvent)
	Drop(WlDataDeviceDropEvent)
	Selection(WlDataDeviceSelectionEvent)
}
type UnimplementedWlDataDeviceListener struct{}

func (e *UnimplementedWlDataDeviceListener) DataOffer(_ WlDataDeviceDataOfferEvent) {
	return
}
func (e *UnimplementedWlDataDeviceListener) Enter(_ WlDataDeviceEnterEvent) {
	return
}
func (e *UnimplementedWlDataDeviceListener) Leave(_ WlDataDeviceLeaveEvent) {
	return
}
func (e *UnimplementedWlDataDeviceListener) Motion(_ WlDataDeviceMotionEvent) {
	return
}
func (e *UnimplementedWlDataDeviceListener) Drop(_ WlDataDeviceDropEvent) {
	return
}
func (e *UnimplementedWlDataDeviceListener) Selection(_ WlDataDeviceSelectionEvent) {
	return
}

type WlDataDeviceDispatcher struct {
	WlDataDeviceListener
}

func NewWlDataDeviceDispatcher() *WlDataDeviceDispatcher {
	return &WlDataDeviceDispatcher{WlDataDeviceListener: &UnimplementedWlDataDeviceListener{}}
}
func (i *WlDataDeviceDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	case 0:
		var e WlDataDeviceDataOfferEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.DataOffer(e)
	case 1:
		var e WlDataDeviceEnterEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Enter(e)
	case 2:
		var e WlDataDeviceLeaveEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Leave(e)
	case 3:
		var e WlDataDeviceMotionEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Motion(e)
	case 4:
		var e WlDataDeviceDropEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Drop(e)
	case 5:
		var e WlDataDeviceSelectionEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Selection(e)
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlDataDeviceDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlDataDeviceListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlDataDeviceListener = listener
}

// WlDataDeviceManager data transfer interface
/*
	The wl_data_device_manager is a singleton global object that
	provides access to inter-client data transfer mechanisms such as
	copy-and-paste and drag-and-drop.  These mechanisms are tied to
	a wl_seat and this interface lets a client get a wl_data_device
	corresponding to a wl_seat.

	Depending on the version bound, the objects created from the bound
	wl_data_device_manager object will have different requirements for
	functioning properly. See wl_data_source.set_actions,
	wl_data_offer.accept and wl_data_offer.finish for details.
*/
type WlDataDeviceManager struct {
	Proxy
}

func NewWlDataDeviceManager(p Proxy) WlDataDeviceManager {
	return WlDataDeviceManager{Proxy: p}
}
func (w *WlDataDeviceManager) AttachListener(l WlDataDeviceManagerListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

// WlDataDeviceManagerDndAction drag and drop actions
/*
	This is a bitmask of the available/preferred actions in a
	drag-and-drop operation.

	In the compositor, the selected action is a result of matching the
	actions offered by the source and destination sides.  "action" events
	with a "none" action will be sent to both source and destination if
	there is no match. All further checks will effectively happen on
	(source actions ∩ destination actions).

	In addition, compositors may also pick different actions in
	reaction to key modifiers being pressed. One common design that
	is used in major toolkits (and the behavior recommended for
	compositors) is:

	- If no modifiers are pressed, the first match (in bit order)
	will be used.
	- Pressing Shift selects "move", if enabled in the mask.
	- Pressing Control selects "copy", if enabled in the mask.

	Behavior beyond that is considered implementation-dependent.
	Compositors may for example bind other modifiers (like Alt/Meta)
	or drags initiated with other buttons than BTN_LEFT to specific
	actions (e.g. "ask").
*/
type WlDataDeviceManagerDndAction uint32

const (
	// WlDataDeviceManagerDndAction_None no action
	WlDataDeviceManagerDndAction_None WlDataDeviceManagerDndAction = 0
	// WlDataDeviceManagerDndAction_Copy copy action
	WlDataDeviceManagerDndAction_Copy WlDataDeviceManagerDndAction = 1
	// WlDataDeviceManagerDndAction_Move move action
	WlDataDeviceManagerDndAction_Move WlDataDeviceManagerDndAction = 2
	// WlDataDeviceManagerDndAction_Ask ask action
	WlDataDeviceManagerDndAction_Ask WlDataDeviceManagerDndAction = 4
)

// WlDataDeviceManagerCreateDataSourceRequest create a new data source
//
//	Create a new data source.
type WlDataDeviceManagerCreateDataSourceRequest struct {
	// Id data source to create
	Id uint32 `wayland:"new_id"`
}

// WlDataDeviceManagerGetDataDeviceRequest create a new data device
//
//	Create a new data device for a given seat.
type WlDataDeviceManagerGetDataDeviceRequest struct {
	// Id data device to create
	Id uint32 `wayland:"new_id"`
	// Seat  associated with the data device
	Seat uint32 `wayland:"object"`
}
type WlDataDeviceManagerListener interface{}
type UnimplementedWlDataDeviceManagerListener struct{}
type WlDataDeviceManagerDispatcher struct {
	WlDataDeviceManagerListener
}

func NewWlDataDeviceManagerDispatcher() *WlDataDeviceManagerDispatcher {
	return &WlDataDeviceManagerDispatcher{WlDataDeviceManagerListener: &UnimplementedWlDataDeviceManagerListener{}}
}
func (i *WlDataDeviceManagerDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlDataDeviceManagerDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlDataDeviceManagerListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlDataDeviceManagerListener = listener
}

// WlShell create desktop-style surfaces
/*
	This interface is implemented by servers that provide
	desktop-style user interfaces.

	It allows clients to associate a wl_shell_surface with
	a basic surface.

	Note! This protocol is deprecated and not intended for production use.
	For desktop-style user interfaces, use xdg_shell. Compositors and clients
	should not implement this interface.
*/
type WlShell struct {
	Proxy
}

func NewWlShell(p Proxy) WlShell {
	return WlShell{Proxy: p}
}
func (w *WlShell) AttachListener(l WlShellListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

type WlShellError uint32

const (
	// WlShellError_Role given wl_surface has another role
	WlShellError_Role WlShellError = 0
)

// WlShellGetShellSurfaceRequest create a shell surface from a surface
/*
	Create a shell surface for an existing surface. This gives
	the wl_surface the role of a shell surface. If the wl_surface
	already has another role, it raises a protocol error.

	Only one shell surface can be associated with a given surface.
*/
type WlShellGetShellSurfaceRequest struct {
	// Id shell surface to create
	Id uint32 `wayland:"new_id"`
	// Surface  to be given the shell surface role
	Surface uint32 `wayland:"object"`
}
type WlShellListener interface{}
type UnimplementedWlShellListener struct{}
type WlShellDispatcher struct {
	WlShellListener
}

func NewWlShellDispatcher() *WlShellDispatcher {
	return &WlShellDispatcher{WlShellListener: &UnimplementedWlShellListener{}}
}
func (i *WlShellDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlShellDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlShellListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlShellListener = listener
}

// WlShellSurface desktop-style metadata interface
/*
	An interface that may be implemented by a wl_surface, for
	implementations that provide a desktop-style user interface.

	It provides requests to treat surfaces like toplevel, fullscreen
	or popup windows, move, resize or maximize them, associate
	metadata like title and class, etc.

	On the server side the object is automatically destroyed when
	the related wl_surface is destroyed. On the client side,
	wl_shell_surface_destroy() must be called before destroying
	the wl_surface object.
*/
type WlShellSurface struct {
	Proxy
}

func NewWlShellSurface(p Proxy) WlShellSurface {
	return WlShellSurface{Proxy: p}
}
func (w *WlShellSurface) AttachListener(l WlShellSurfaceListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

// WlShellSurfaceResize edge values for resizing
/*
	These values are used to indicate which edge of a surface
	is being dragged in a resize operation. The server may
	use this information to adapt its behavior, e.g. choose
	an appropriate cursor image.
*/
type WlShellSurfaceResize uint32

const (
	// WlShellSurfaceResize_None no edge
	WlShellSurfaceResize_None WlShellSurfaceResize = 0
	// WlShellSurfaceResize_Top top edge
	WlShellSurfaceResize_Top WlShellSurfaceResize = 1
	// WlShellSurfaceResize_Bottom bottom edge
	WlShellSurfaceResize_Bottom WlShellSurfaceResize = 2
	// WlShellSurfaceResize_Left left edge
	WlShellSurfaceResize_Left WlShellSurfaceResize = 4
	// WlShellSurfaceResize_TopLeft top and left edges
	WlShellSurfaceResize_TopLeft WlShellSurfaceResize = 5
	// WlShellSurfaceResize_BottomLeft bottom and left edges
	WlShellSurfaceResize_BottomLeft WlShellSurfaceResize = 6
	// WlShellSurfaceResize_Right right edge
	WlShellSurfaceResize_Right WlShellSurfaceResize = 8
	// WlShellSurfaceResize_TopRight top and right edges
	WlShellSurfaceResize_TopRight WlShellSurfaceResize = 9
	// WlShellSurfaceResize_BottomRight bottom and right edges
	WlShellSurfaceResize_BottomRight WlShellSurfaceResize = 10
)

// WlShellSurfaceTransient details of transient behaviour
/*
	These flags specify details of the expected behaviour
	of transient surfaces. Used in the set_transient request.
*/
type WlShellSurfaceTransient uint32

const (
	// WlShellSurfaceTransient_Inactive do not set keyboard focus
	WlShellSurfaceTransient_Inactive WlShellSurfaceTransient = 0x1
)

// WlShellSurfaceFullscreenMethod different method to set the surface fullscreen
/*
	Hints to indicate to the compositor how to deal with a conflict
	between the dimensions of the surface and the dimensions of the
	output. The compositor is free to ignore this parameter.
*/
type WlShellSurfaceFullscreenMethod uint32

const (
	// WlShellSurfaceFullscreenMethod_Default no preference, apply default policy
	WlShellSurfaceFullscreenMethod_Default WlShellSurfaceFullscreenMethod = 0
	// WlShellSurfaceFullscreenMethod_Scale scale, preserve the surface's aspect ratio and center on output
	WlShellSurfaceFullscreenMethod_Scale WlShellSurfaceFullscreenMethod = 1
	// WlShellSurfaceFullscreenMethod_Driver switch output mode to the smallest mode that can fit the surface, add black borders to compensate size mismatch
	WlShellSurfaceFullscreenMethod_Driver WlShellSurfaceFullscreenMethod = 2
	// WlShellSurfaceFullscreenMethod_Fill no upscaling, center on output and add black borders to compensate size mismatch
	WlShellSurfaceFullscreenMethod_Fill WlShellSurfaceFullscreenMethod = 3
)

// WlShellSurfacePongRequest respond to a ping event
/*
	A client must respond to a ping event with a pong request or
	the client may be deemed unresponsive.
*/
type WlShellSurfacePongRequest struct {
	// Serial  number of the ping event
	Serial uint32 `wayland:"uint"`
}

// WlShellSurfaceMoveRequest start an interactive move
/*
	Start a pointer-driven move of the surface.

	This request must be used in response to a button press event.
	The server may ignore move requests depending on the state of
	the surface (e.g. fullscreen or maximized).
*/
type WlShellSurfaceMoveRequest struct {
	// Seat  whose pointer is used
	Seat uint32 `wayland:"object"`
	// Serial  number of the implicit grab on the pointer
	Serial uint32 `wayland:"uint"`
}

// WlShellSurfaceResizeRequest start an interactive resize
/*
	Start a pointer-driven resizing of the surface.

	This request must be used in response to a button press event.
	The server may ignore resize requests depending on the state of
	the surface (e.g. fullscreen or maximized).
*/
type WlShellSurfaceResizeRequest struct {
	// Seat  whose pointer is used
	Seat uint32 `wayland:"object"`
	// Serial  number of the implicit grab on the pointer
	Serial uint32 `wayland:"uint"`
	// Edges which edge or corner is being dragged
	Edges uint32 `wayland:"uint"`
}

// WlShellSurfaceSetToplevelRequest make the surface a toplevel surface
/*
	Map the surface as a toplevel surface.

	A toplevel surface is not fullscreen, maximized or transient.
*/
type WlShellSurfaceSetToplevelRequest struct{}

// WlShellSurfaceSetTransientRequest make the surface a transient surface
/*
	Map the surface relative to an existing surface.

	The x and y arguments specify the location of the upper left
	corner of the surface relative to the upper left corner of the
	parent surface, in surface-local coordinates.

	The flags argument controls details of the transient behaviour.
*/
type WlShellSurfaceSetTransientRequest struct {
	// Parent  surface
	Parent uint32 `wayland:"object"`
	// X surface-local x coordinate
	X int32 `wayland:"int"`
	// Y surface-local y coordinate
	Y int32 `wayland:"int"`
	// Flags transient surface behavior
	Flags uint32 `wayland:"uint"`
}

// WlShellSurfaceSetFullscreenRequest make the surface a fullscreen surface
/*
	Map the surface as a fullscreen surface.

	If an output parameter is given then the surface will be made
	fullscreen on that output. If the client does not specify the
	output then the compositor will apply its policy - usually
	choosing the output on which the surface has the biggest surface
	area.

	The client may specify a method to resolve a size conflict
	between the output size and the surface size - this is provided
	through the method parameter.

	The framerate parameter is used only when the method is set
	to "driver", to indicate the preferred framerate. A value of 0
	indicates that the client does not care about framerate.  The
	framerate is specified in mHz, that is framerate of 60000 is 60Hz.

	A method of "scale" or "driver" implies a scaling operation of
	the surface, either via a direct scaling operation or a change of
	the output mode. This will override any kind of output scaling, so
	that mapping a surface with a buffer size equal to the mode can
	fill the screen independent of buffer_scale.

	A method of "fill" means we don't scale up the buffer, however
	any output scale is applied. This means that you may run into
	an edge case where the application maps a buffer with the same
	size of the output mode but buffer_scale 1 (thus making a
	surface larger than the output). In this case it is allowed to
	downscale the results to fit the screen.

	The compositor must reply to this request with a configure event
	with the dimensions for the output on which the surface will
	be made fullscreen.
*/
type WlShellSurfaceSetFullscreenRequest struct {
	// Method  for resolving size conflict
	Method uint32 `wayland:"uint"`
	// Framerate  in mHz
	Framerate uint32 `wayland:"uint"`
	// Output  on which the surface is to be fullscreen
	Output uint32 `wayland:"object"`
}

// WlShellSurfaceSetPopupRequest make the surface a popup surface
/*
	Map the surface as a popup.

	A popup surface is a transient surface with an added pointer
	grab.

	An existing implicit grab will be changed to owner-events mode,
	and the popup grab will continue after the implicit grab ends
	(i.e. releasing the mouse button does not cause the popup to
	be unmapped).

	The popup grab continues until the window is destroyed or a
	mouse button is pressed in any other client's window. A click
	in any of the client's surfaces is reported as normal, however,
	clicks in other clients' surfaces will be discarded and trigger
	the callback.

	The x and y arguments specify the location of the upper left
	corner of the surface relative to the upper left corner of the
	parent surface, in surface-local coordinates.
*/
type WlShellSurfaceSetPopupRequest struct {
	// Seat  whose pointer is used
	Seat uint32 `wayland:"object"`
	// Serial  number of the implicit grab on the pointer
	Serial uint32 `wayland:"uint"`
	// Parent  surface
	Parent uint32 `wayland:"object"`
	// X surface-local x coordinate
	X int32 `wayland:"int"`
	// Y surface-local y coordinate
	Y int32 `wayland:"int"`
	// Flags transient surface behavior
	Flags uint32 `wayland:"uint"`
}

// WlShellSurfaceSetMaximizedRequest make the surface a maximized surface
/*
	Map the surface as a maximized surface.

	If an output parameter is given then the surface will be
	maximized on that output. If the client does not specify the
	output then the compositor will apply its policy - usually
	choosing the output on which the surface has the biggest surface
	area.

	The compositor will reply with a configure event telling
	the expected new surface size. The operation is completed
	on the next buffer attach to this surface.

	A maximized surface typically fills the entire output it is
	bound to, except for desktop elements such as panels. This is
	the main difference between a maximized shell surface and a
	fullscreen shell surface.

	The details depend on the compositor implementation.
*/
type WlShellSurfaceSetMaximizedRequest struct {
	// Output  on which the surface is to be maximized
	Output uint32 `wayland:"object"`
}

// WlShellSurfaceSetTitleRequest set surface title
/*
	Set a short title for the surface.

	This string may be used to identify the surface in a task bar,
	window list, or other user interface elements provided by the
	compositor.

	The string must be encoded in UTF-8.
*/
type WlShellSurfaceSetTitleRequest struct {
	// Title surface title
	Title string `wayland:"string"`
}

// WlShellSurfaceSetClassRequest set surface class
/*
	Set a class for the surface.

	The surface class identifies the general class of applications
	to which the surface belongs. A common convention is to use the
	file name (or the full path if it is a non-standard location) of
	the application's .desktop file as the class.
*/
type WlShellSurfaceSetClassRequest struct {
	// Class surface class
	Class string `wayland:"string"`
}

// WlShellSurfacePingEvent ping client
/*
	Ping a client to check if it is receiving events and sending
	requests. A client is expected to reply with a pong request.
*/
type WlShellSurfacePingEvent struct {
	// Serial  number of the ping
	Serial uint32 `wayland:"uint"`
}

func (e *WlShellSurfacePingEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Serial = UnmarshallUint32(offset, d)
	return nil
}

// WlShellSurfaceConfigureEvent suggest resize
/*
	The configure event asks the client to resize its surface.

	The size is a hint, in the sense that the client is free to
	ignore it if it doesn't resize, pick a smaller size (to
	satisfy aspect ratio or resize in steps of NxM pixels).

	The edges parameter provides a hint about how the surface
	was resized. The client may use this information to decide
	how to adjust its content to the new size (e.g. a scrolling
	area might adjust its content position to leave the viewable
	content unmoved).

	The client is free to dismiss all but the last configure
	event it received.

	The width and height arguments specify the size of the window
	in surface-local coordinates.
*/
type WlShellSurfaceConfigureEvent struct {
	// Edges how the surface was resized
	Edges uint32 `wayland:"uint"`
	// Width new width of the surface
	Width int32 `wayland:"int"`
	// Height new height of the surface
	Height int32 `wayland:"int"`
}

func (e *WlShellSurfaceConfigureEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Edges = UnmarshallUint32(offset, d)
	offset, e.Width = UnmarshallInt32(offset, d)
	offset, e.Height = UnmarshallInt32(offset, d)
	return nil
}

// WlShellSurfacePopupDoneEvent popup interaction is done
/*
	The popup_done event is sent out when a popup grab is broken,
	that is, when the user clicks a surface that doesn't belong
	to the client owning the popup surface.
*/
type WlShellSurfacePopupDoneEvent struct{}

func (e *WlShellSurfacePopupDoneEvent) UnmarshallBinary(d []byte) error {
	return nil
}

type WlShellSurfaceListener interface {
	Ping(WlShellSurfacePingEvent)
	Configure(WlShellSurfaceConfigureEvent)
	PopupDone(WlShellSurfacePopupDoneEvent)
}
type UnimplementedWlShellSurfaceListener struct{}

func (e *UnimplementedWlShellSurfaceListener) Ping(_ WlShellSurfacePingEvent) {
	return
}
func (e *UnimplementedWlShellSurfaceListener) Configure(_ WlShellSurfaceConfigureEvent) {
	return
}
func (e *UnimplementedWlShellSurfaceListener) PopupDone(_ WlShellSurfacePopupDoneEvent) {
	return
}

type WlShellSurfaceDispatcher struct {
	WlShellSurfaceListener
}

func NewWlShellSurfaceDispatcher() *WlShellSurfaceDispatcher {
	return &WlShellSurfaceDispatcher{WlShellSurfaceListener: &UnimplementedWlShellSurfaceListener{}}
}
func (i *WlShellSurfaceDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	case 0:
		var e WlShellSurfacePingEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Ping(e)
	case 1:
		var e WlShellSurfaceConfigureEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Configure(e)
	case 2:
		var e WlShellSurfacePopupDoneEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.PopupDone(e)
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlShellSurfaceDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlShellSurfaceListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlShellSurfaceListener = listener
}

// WlSurface an onscreen surface
/*
	A surface is a rectangular area that may be displayed on zero
	or more outputs, and shown any number of times at the compositor's
	discretion. They can present wl_buffers, receive user input, and
	define a local coordinate system.

	The size of a surface (and relative positions on it) is described
	in surface-local coordinates, which may differ from the buffer
	coordinates of the pixel content, in case a buffer_transform
	or a buffer_scale is used.

	A surface without a "role" is fairly useless: a compositor does
	not know where, when or how to present it. The role is the
	purpose of a wl_surface. Examples of roles are a cursor for a
	pointer (as set by wl_pointer.set_cursor), a drag icon
	(wl_data_device.start_drag), a sub-surface
	(wl_subcompositor.get_subsurface), and a window as defined by a
	shell protocol (e.g. wl_shell.get_shell_surface).

	A surface can have only one role at a time. Initially a
	wl_surface does not have a role. Once a wl_surface is given a
	role, it is set permanently for the whole lifetime of the
	wl_surface object. Giving the current role again is allowed,
	unless explicitly forbidden by the relevant interface
	specification.

	Surface roles are given by requests in other interfaces such as
	wl_pointer.set_cursor. The request should explicitly mention
	that this request gives a role to a wl_surface. Often, this
	request also creates a new protocol object that represents the
	role and adds additional functionality to wl_surface. When a
	client wants to destroy a wl_surface, they must destroy this 'role
	object' before the wl_surface.

	Destroying the role object does not remove the role from the
	wl_surface, but it may stop the wl_surface from "playing the role".
	For instance, if a wl_subsurface object is destroyed, the wl_surface
	it was created for will be unmapped and forget its position and
	z-order. It is allowed to create a wl_subsurface for the same
	wl_surface again, but it is not allowed to use the wl_surface as
	a cursor (cursor is a different role than sub-surface, and role
	switching is not allowed).
*/
type WlSurface struct {
	Proxy
}

func NewWlSurface(p Proxy) WlSurface {
	return WlSurface{Proxy: p}
}
func (w *WlSurface) AttachListener(l WlSurfaceListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

// WlSurfaceError wl_surface error values
//
//	These errors can be emitted in response to wl_surface requests.
type WlSurfaceError uint32

const (
	// WlSurfaceError_InvalidScale buffer scale value is invalid
	WlSurfaceError_InvalidScale WlSurfaceError = 0
	// WlSurfaceError_InvalidTransform buffer transform value is invalid
	WlSurfaceError_InvalidTransform WlSurfaceError = 1
	// WlSurfaceError_InvalidSize buffer size is invalid
	WlSurfaceError_InvalidSize WlSurfaceError = 2
	// WlSurfaceError_InvalidOffset buffer offset is invalid
	WlSurfaceError_InvalidOffset WlSurfaceError = 3
)

// WlSurfaceDestroyRequest delete surface
//
//	Deletes the surface and invalidates its object ID.
type WlSurfaceDestroyRequest struct{}

// WlSurfaceAttachRequest set the surface contents
/*
	Set a buffer as the content of this surface.

	The new size of the surface is calculated based on the buffer
	size transformed by the inverse buffer_transform and the
	inverse buffer_scale. This means that at commit time the supplied
	buffer size must be an integer multiple of the buffer_scale. If
	that's not the case, an invalid_size error is sent.

	The x and y arguments specify the location of the new pending
	buffer's upper left corner, relative to the current buffer's upper
	left corner, in surface-local coordinates. In other words, the
	x and y, combined with the new surface size define in which
	directions the surface's size changes. Setting anything other than 0
	as x and y arguments is discouraged, and should instead be replaced
	with using the separate wl_surface.offset request.

	When the bound wl_surface version is 5 or higher, passing any
	non-zero x or y is a protocol violation, and will result in an
	'invalid_offset' error being raised. To achieve equivalent semantics,
	use wl_surface.offset.

	Surface contents are double-buffered state, see wl_surface.commit.

	The initial surface contents are void; there is no content.
	wl_surface.attach assigns the given wl_buffer as the pending
	wl_buffer. wl_surface.commit makes the pending wl_buffer the new
	surface contents, and the size of the surface becomes the size
	calculated from the wl_buffer, as described above. After commit,
	there is no pending buffer until the next attach.

	Committing a pending wl_buffer allows the compositor to read the
	pixels in the wl_buffer. The compositor may access the pixels at
	any time after the wl_surface.commit request. When the compositor
	will not access the pixels anymore, it will send the
	wl_buffer.release event. Only after receiving wl_buffer.release,
	the client may reuse the wl_buffer. A wl_buffer that has been
	attached and then replaced by another attach instead of committed
	will not receive a release event, and is not used by the
	compositor.

	If a pending wl_buffer has been committed to more than one wl_surface,
	the delivery of wl_buffer.release events becomes undefined. A well
	behaved client should not rely on wl_buffer.release events in this
	case. Alternatively, a client could create multiple wl_buffer objects
	from the same backing storage or use wp_linux_buffer_release.

	Destroying the wl_buffer after wl_buffer.release does not change
	the surface contents. Destroying the wl_buffer before wl_buffer.release
	is allowed as long as the underlying buffer storage isn't re-used (this
	can happen e.g. on client process termination). However, if the client
	destroys the wl_buffer before receiving the wl_buffer.release event and
	mutates the underlying buffer storage, the surface contents become
	undefined immediately.

	If wl_surface.attach is sent with a NULL wl_buffer, the
	following wl_surface.commit will remove the surface content.
*/
type WlSurfaceAttachRequest struct {
	// Buffer  of surface contents
	Buffer uint32 `wayland:"object"`
	// X surface-local x coordinate
	X int32 `wayland:"int"`
	// Y surface-local y coordinate
	Y int32 `wayland:"int"`
}

// WlSurfaceDamageRequest mark part of the surface damaged
/*
	This request is used to describe the regions where the pending
	buffer is different from the current surface contents, and where
	the surface therefore needs to be repainted. The compositor
	ignores the parts of the damage that fall outside of the surface.

	Damage is double-buffered state, see wl_surface.commit.

	The damage rectangle is specified in surface-local coordinates,
	where x and y specify the upper left corner of the damage rectangle.

	The initial value for pending damage is empty: no damage.
	wl_surface.damage adds pending damage: the new pending damage
	is the union of old pending damage and the given rectangle.

	wl_surface.commit assigns pending damage as the current damage,
	and clears pending damage. The server will clear the current
	damage as it repaints the surface.

	Note! New clients should not use this request. Instead damage can be
	posted with wl_surface.damage_buffer which uses buffer coordinates
	instead of surface coordinates.
*/
type WlSurfaceDamageRequest struct {
	// X surface-local x coordinate
	X int32 `wayland:"int"`
	// Y surface-local y coordinate
	Y int32 `wayland:"int"`
	// Width  of damage rectangle
	Width int32 `wayland:"int"`
	// Height  of damage rectangle
	Height int32 `wayland:"int"`
}

// WlSurfaceFrameRequest request a frame throttling hint
/*
	Request a notification when it is a good time to start drawing a new
	frame, by creating a frame callback. This is useful for throttling
	redrawing operations, and driving animations.

	When a client is animating on a wl_surface, it can use the 'frame'
	request to get notified when it is a good time to draw and commit the
	next frame of animation. If the client commits an update earlier than
	that, it is likely that some updates will not make it to the display,
	and the client is wasting resources by drawing too often.

	The frame request will take effect on the next wl_surface.commit.
	The notification will only be posted for one frame unless
	requested again. For a wl_surface, the notifications are posted in
	the order the frame requests were committed.

	The server must send the notifications so that a client
	will not send excessive updates, while still allowing
	the highest possible update rate for clients that wait for the reply
	before drawing again. The server should give some time for the client
	to draw and commit after sending the frame callback events to let it
	hit the next output refresh.

	A server should avoid signaling the frame callbacks if the
	surface is not visible in any way, e.g. the surface is off-screen,
	or completely obscured by other opaque surfaces.

	The object returned by this request will be destroyed by the
	compositor after the callback is fired and as such the client must not
	attempt to use it after that point.

	The callback_data passed in the callback is the current time, in
	milliseconds, with an undefined base.
*/
type WlSurfaceFrameRequest struct {
	// Callback  object for the frame request
	Callback uint32 `wayland:"new_id"`
}

// WlSurfaceSetOpaqueRegionRequest set opaque region
/*
	This request sets the region of the surface that contains
	opaque content.

	The opaque region is an optimization hint for the compositor
	that lets it optimize the redrawing of content behind opaque
	regions.  Setting an opaque region is not required for correct
	behaviour, but marking transparent content as opaque will result
	in repaint artifacts.

	The opaque region is specified in surface-local coordinates.

	The compositor ignores the parts of the opaque region that fall
	outside of the surface.

	Opaque region is double-buffered state, see wl_surface.commit.

	wl_surface.set_opaque_region changes the pending opaque region.
	wl_surface.commit copies the pending region to the current region.
	Otherwise, the pending and current regions are never changed.

	The initial value for an opaque region is empty. Setting the pending
	opaque region has copy semantics, and the wl_region object can be
	destroyed immediately. A NULL wl_region causes the pending opaque
	region to be set to empty.
*/
type WlSurfaceSetOpaqueRegionRequest struct {
	// Region opaque region of the surface
	Region uint32 `wayland:"object"`
}

// WlSurfaceSetInputRegionRequest set input region
/*
	This request sets the region of the surface that can receive
	pointer and touch events.

	Input events happening outside of this region will try the next
	surface in the server surface stack. The compositor ignores the
	parts of the input region that fall outside of the surface.

	The input region is specified in surface-local coordinates.

	Input region is double-buffered state, see wl_surface.commit.

	wl_surface.set_input_region changes the pending input region.
	wl_surface.commit copies the pending region to the current region.
	Otherwise the pending and current regions are never changed,
	except cursor and icon surfaces are special cases, see
	wl_pointer.set_cursor and wl_data_device.start_drag.

	The initial value for an input region is infinite. That means the
	whole surface will accept input. Setting the pending input region
	has copy semantics, and the wl_region object can be destroyed
	immediately. A NULL wl_region causes the input region to be set
	to infinite.
*/
type WlSurfaceSetInputRegionRequest struct {
	// Region input region of the surface
	Region uint32 `wayland:"object"`
}

// WlSurfaceCommitRequest commit pending surface state
/*
	Surface state (input, opaque, and damage regions, attached buffers,
	etc.) is double-buffered. Protocol requests modify the pending state,
	as opposed to the current state in use by the compositor. A commit
	request atomically applies all pending state, replacing the current
	state. After commit, the new pending state is as documented for each
	related request.

	On commit, a pending wl_buffer is applied first, and all other state
	second. This means that all coordinates in double-buffered state are
	relative to the new wl_buffer coming into use, except for
	wl_surface.attach itself. If there is no pending wl_buffer, the
	coordinates are relative to the current surface contents.

	All requests that need a commit to become effective are documented
	to affect double-buffered state.

	Other interfaces may add further double-buffered surface state.
*/
type WlSurfaceCommitRequest struct{}

// WlSurfaceSetBufferTransformRequest sets the buffer transformation
/*
	This request sets an optional transformation on how the compositor
	interprets the contents of the buffer attached to the surface. The
	accepted values for the transform parameter are the values for
	wl_output.transform.

	Buffer transform is double-buffered state, see wl_surface.commit.

	A newly created surface has its buffer transformation set to normal.

	wl_surface.set_buffer_transform changes the pending buffer
	transformation. wl_surface.commit copies the pending buffer
	transformation to the current one. Otherwise, the pending and current
	values are never changed.

	The purpose of this request is to allow clients to render content
	according to the output transform, thus permitting the compositor to
	use certain optimizations even if the display is rotated. Using
	hardware overlays and scanning out a client buffer for fullscreen
	surfaces are examples of such optimizations. Those optimizations are
	highly dependent on the compositor implementation, so the use of this
	request should be considered on a case-by-case basis.

	Note that if the transform value includes 90 or 270 degree rotation,
	the width of the buffer will become the surface height and the height
	of the buffer will become the surface width.

	If transform is not one of the values from the
	wl_output.transform enum the invalid_transform protocol error
	is raised.
*/
type WlSurfaceSetBufferTransformRequest struct {
	// Transform  for interpreting buffer contents
	Transform int32 `wayland:"int"`
}

// WlSurfaceSetBufferScaleRequest sets the buffer scaling factor
/*
	This request sets an optional scaling factor on how the compositor
	interprets the contents of the buffer attached to the window.

	Buffer scale is double-buffered state, see wl_surface.commit.

	A newly created surface has its buffer scale set to 1.

	wl_surface.set_buffer_scale changes the pending buffer scale.
	wl_surface.commit copies the pending buffer scale to the current one.
	Otherwise, the pending and current values are never changed.

	The purpose of this request is to allow clients to supply higher
	resolution buffer data for use on high resolution outputs. It is
	intended that you pick the same buffer scale as the scale of the
	output that the surface is displayed on. This means the compositor
	can avoid scaling when rendering the surface on that output.

	Note that if the scale is larger than 1, then you have to attach
	a buffer that is larger (by a factor of scale in each dimension)
	than the desired surface size.

	If scale is not positive the invalid_scale protocol error is
	raised.
*/
type WlSurfaceSetBufferScaleRequest struct {
	// Scale positive scale for interpreting buffer contents
	Scale int32 `wayland:"int"`
}

// WlSurfaceDamageBufferRequest mark part of the surface damaged using buffer coordinates
/*
	This request is used to describe the regions where the pending
	buffer is different from the current surface contents, and where
	the surface therefore needs to be repainted. The compositor
	ignores the parts of the damage that fall outside of the surface.

	Damage is double-buffered state, see wl_surface.commit.

	The damage rectangle is specified in buffer coordinates,
	where x and y specify the upper left corner of the damage rectangle.

	The initial value for pending damage is empty: no damage.
	wl_surface.damage_buffer adds pending damage: the new pending
	damage is the union of old pending damage and the given rectangle.

	wl_surface.commit assigns pending damage as the current damage,
	and clears pending damage. The server will clear the current
	damage as it repaints the surface.

	This request differs from wl_surface.damage in only one way - it
	takes damage in buffer coordinates instead of surface-local
	coordinates. While this generally is more intuitive than surface
	coordinates, it is especially desirable when using wp_viewport
	or when a drawing library (like EGL) is unaware of buffer scale
	and buffer transform.

	Note: Because buffer transformation changes and damage requests may
	be interleaved in the protocol stream, it is impossible to determine
	the actual mapping between surface and buffer damage until
	wl_surface.commit time. Therefore, compositors wishing to take both
	kinds of damage into account will have to accumulate damage from the
	two requests separately and only transform from one to the other
	after receiving the wl_surface.commit.
*/
type WlSurfaceDamageBufferRequest struct {
	// X buffer-local x coordinate
	X int32 `wayland:"int"`
	// Y buffer-local y coordinate
	Y int32 `wayland:"int"`
	// Width  of damage rectangle
	Width int32 `wayland:"int"`
	// Height  of damage rectangle
	Height int32 `wayland:"int"`
}

// WlSurfaceOffsetRequest set the surface contents offset
/*
	The x and y arguments specify the location of the new pending
	buffer's upper left corner, relative to the current buffer's upper
	left corner, in surface-local coordinates. In other words, the
	x and y, combined with the new surface size define in which
	directions the surface's size changes.

	Surface location offset is double-buffered state, see
	wl_surface.commit.

	This request is semantically equivalent to and the replaces the x and y
	arguments in the wl_surface.attach request in wl_surface versions prior
	to 5. See wl_surface.attach for details.
*/
type WlSurfaceOffsetRequest struct {
	// X surface-local x coordinate
	X int32 `wayland:"int"`
	// Y surface-local y coordinate
	Y int32 `wayland:"int"`
}

// WlSurfaceEnterEvent surface enters an output
/*
	This is emitted whenever a surface's creation, movement, or resizing
	results in some part of it being within the scanout region of an
	output.

	Note that a surface may be overlapping with zero or more outputs.
*/
type WlSurfaceEnterEvent struct {
	// Output  entered by the surface
	Output uint32 `wayland:"object"`
}

func (e *WlSurfaceEnterEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Output = UnmarshallUint32(offset, d)
	return nil
}

// WlSurfaceLeaveEvent surface leaves an output
/*
	This is emitted whenever a surface's creation, movement, or resizing
	results in it no longer having any part of it within the scanout region
	of an output.

	Clients should not use the number of outputs the surface is on for frame
	throttling purposes. The surface might be hidden even if no leave event
	has been sent, and the compositor might expect new surface content
	updates even if no enter event has been sent. The frame event should be
	used instead.
*/
type WlSurfaceLeaveEvent struct {
	// Output  left by the surface
	Output uint32 `wayland:"object"`
}

func (e *WlSurfaceLeaveEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Output = UnmarshallUint32(offset, d)
	return nil
}

type WlSurfaceListener interface {
	Enter(WlSurfaceEnterEvent)
	Leave(WlSurfaceLeaveEvent)
}
type UnimplementedWlSurfaceListener struct{}

func (e *UnimplementedWlSurfaceListener) Enter(_ WlSurfaceEnterEvent) {
	return
}
func (e *UnimplementedWlSurfaceListener) Leave(_ WlSurfaceLeaveEvent) {
	return
}

type WlSurfaceDispatcher struct {
	WlSurfaceListener
}

func NewWlSurfaceDispatcher() *WlSurfaceDispatcher {
	return &WlSurfaceDispatcher{WlSurfaceListener: &UnimplementedWlSurfaceListener{}}
}
func (i *WlSurfaceDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	case 0:
		var e WlSurfaceEnterEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Enter(e)
	case 1:
		var e WlSurfaceLeaveEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Leave(e)
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlSurfaceDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlSurfaceListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlSurfaceListener = listener
}

// WlSeat group of input devices
/*
	A seat is a group of keyboards, pointer and touch devices. This
	object is published as a global during start up, or when such a
	device is hot plugged.  A seat typically has a pointer and
	maintains a keyboard focus and a pointer focus.
*/
type WlSeat struct {
	Proxy
}

func NewWlSeat(p Proxy) WlSeat {
	return WlSeat{Proxy: p}
}
func (w *WlSeat) AttachListener(l WlSeatListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

// WlSeatCapability seat capability bitmask
/*
	This is a bitmask of capabilities this seat has; if a member is
	set, then it is present on the seat.
*/
type WlSeatCapability uint32

const (
	// WlSeatCapability_Pointer the seat has pointer devices
	WlSeatCapability_Pointer WlSeatCapability = 1
	// WlSeatCapability_Keyboard the seat has one or more keyboards
	WlSeatCapability_Keyboard WlSeatCapability = 2
	// WlSeatCapability_Touch the seat has touch devices
	WlSeatCapability_Touch WlSeatCapability = 4
)

// WlSeatError wl_seat error values
//
//	These errors can be emitted in response to wl_seat requests.
type WlSeatError uint32

const (
	// WlSeatError_MissingCapability get_pointer, get_keyboard or get_touch called on seat without the matching capability
	WlSeatError_MissingCapability WlSeatError = 0
)

// WlSeatGetPointerRequest return pointer object
/*
	The ID provided will be initialized to the wl_pointer interface
	for this seat.

	This request only takes effect if the seat has the pointer
	capability, or has had the pointer capability in the past.
	It is a protocol violation to issue this request on a seat that has
	never had the pointer capability. The missing_capability error will
	be sent in this case.
*/
type WlSeatGetPointerRequest struct {
	// Id seat pointer
	Id uint32 `wayland:"new_id"`
}

// WlSeatGetKeyboardRequest return keyboard object
/*
	The ID provided will be initialized to the wl_keyboard interface
	for this seat.

	This request only takes effect if the seat has the keyboard
	capability, or has had the keyboard capability in the past.
	It is a protocol violation to issue this request on a seat that has
	never had the keyboard capability. The missing_capability error will
	be sent in this case.
*/
type WlSeatGetKeyboardRequest struct {
	// Id seat keyboard
	Id uint32 `wayland:"new_id"`
}

// WlSeatGetTouchRequest return touch object
/*
	The ID provided will be initialized to the wl_touch interface
	for this seat.

	This request only takes effect if the seat has the touch
	capability, or has had the touch capability in the past.
	It is a protocol violation to issue this request on a seat that has
	never had the touch capability. The missing_capability error will
	be sent in this case.
*/
type WlSeatGetTouchRequest struct {
	// Id seat touch interface
	Id uint32 `wayland:"new_id"`
}

// WlSeatReleaseRequest release the seat object
/*
	Using this request a client can tell the server that it is not going to
	use the seat object anymore.
*/
type WlSeatReleaseRequest struct{}

// WlSeatCapabilitiesEvent seat capabilities changed
/*
	This is emitted whenever a seat gains or loses the pointer,
	keyboard or touch capabilities.  The argument is a capability
	enum containing the complete set of capabilities this seat has.

	When the pointer capability is added, a client may create a
	wl_pointer object using the wl_seat.get_pointer request. This object
	will receive pointer events until the capability is removed in the
	future.

	When the pointer capability is removed, a client should destroy the
	wl_pointer objects associated with the seat where the capability was
	removed, using the wl_pointer.release request. No further pointer
	events will be received on these objects.

	In some compositors, if a seat regains the pointer capability and a
	client has a previously obtained wl_pointer object of version 4 or
	less, that object may start sending pointer events again. This
	behavior is considered a misinterpretation of the intended behavior
	and must not be relied upon by the client. wl_pointer objects of
	version 5 or later must not send events if created before the most
	recent event notifying the client of an added pointer capability.

	The above behavior also applies to wl_keyboard and wl_touch with the
	keyboard and touch capabilities, respectively.
*/
type WlSeatCapabilitiesEvent struct {
	// Capabilities  of the seat
	Capabilities uint32 `wayland:"uint"`
}

func (e *WlSeatCapabilitiesEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Capabilities = UnmarshallUint32(offset, d)
	return nil
}

// WlSeatNameEvent unique identifier for this seat
/*
	In a multi-seat configuration the seat name can be used by clients to
	help identify which physical devices the seat represents.

	The seat name is a UTF-8 string with no convention defined for its
	contents. Each name is unique among all wl_seat globals. The name is
	only guaranteed to be unique for the current compositor instance.

	The same seat names are used for all clients. Thus, the name can be
	shared across processes to refer to a specific wl_seat global.

	The name event is sent after binding to the seat global. This event is
	only sent once per seat object, and the name does not change over the
	lifetime of the wl_seat global.

	Compositors may re-use the same seat name if the wl_seat global is
	destroyed and re-created later.
*/
type WlSeatNameEvent struct {
	// Name seat identifier
	Name string `wayland:"string"`
}

func (e *WlSeatNameEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Name = UnmarshallString(offset, d)
	return nil
}

type WlSeatListener interface {
	Capabilities(WlSeatCapabilitiesEvent)
	Name(WlSeatNameEvent)
}
type UnimplementedWlSeatListener struct{}

func (e *UnimplementedWlSeatListener) Capabilities(_ WlSeatCapabilitiesEvent) {
	return
}
func (e *UnimplementedWlSeatListener) Name(_ WlSeatNameEvent) {
	return
}

type WlSeatDispatcher struct {
	WlSeatListener
}

func NewWlSeatDispatcher() *WlSeatDispatcher {
	return &WlSeatDispatcher{WlSeatListener: &UnimplementedWlSeatListener{}}
}
func (i *WlSeatDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	case 0:
		var e WlSeatCapabilitiesEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Capabilities(e)
	case 1:
		var e WlSeatNameEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Name(e)
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlSeatDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlSeatListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlSeatListener = listener
}

// WlPointer pointer input device
/*
	The wl_pointer interface represents one or more input devices,
	such as mice, which control the pointer location and pointer_focus
	of a seat.

	The wl_pointer interface generates motion, enter and leave
	events for the surfaces that the pointer is located over,
	and button and axis events for button presses, button releases
	and scrolling.
*/
type WlPointer struct {
	Proxy
}

func NewWlPointer(p Proxy) WlPointer {
	return WlPointer{Proxy: p}
}
func (w *WlPointer) AttachListener(l WlPointerListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

type WlPointerError uint32

const (
	// WlPointerError_Role given wl_surface has another role
	WlPointerError_Role WlPointerError = 0
)

// WlPointerButtonState physical button state
/*
	Describes the physical state of a button that produced the button
	event.
*/
type WlPointerButtonState uint32

const (
	// WlPointerButtonState_Released the button is not pressed
	WlPointerButtonState_Released WlPointerButtonState = 0
	// WlPointerButtonState_Pressed the button is pressed
	WlPointerButtonState_Pressed WlPointerButtonState = 1
)

// WlPointerAxis axis types
//
//	Describes the axis types of scroll events.
type WlPointerAxis uint32

const (
	// WlPointerAxis_VerticalScroll vertical axis
	WlPointerAxis_VerticalScroll WlPointerAxis = 0
	// WlPointerAxis_HorizontalScroll horizontal axis
	WlPointerAxis_HorizontalScroll WlPointerAxis = 1
)

// WlPointerAxisSource axis source types
/*
	Describes the source types for axis events. This indicates to the
	client how an axis event was physically generated; a client may
	adjust the user interface accordingly. For example, scroll events
	from a "finger" source may be in a smooth coordinate space with
	kinetic scrolling whereas a "wheel" source may be in discrete steps
	of a number of lines.

	The "continuous" axis source is a device generating events in a
	continuous coordinate space, but using something other than a
	finger. One example for this source is button-based scrolling where
	the vertical motion of a device is converted to scroll events while
	a button is held down.

	The "wheel tilt" axis source indicates that the actual device is a
	wheel but the scroll event is not caused by a rotation but a
	(usually sideways) tilt of the wheel.
*/
type WlPointerAxisSource uint32

const (
	// WlPointerAxisSource_Wheel a physical wheel rotation
	WlPointerAxisSource_Wheel WlPointerAxisSource = 0
	// WlPointerAxisSource_Finger finger on a touch surface
	WlPointerAxisSource_Finger WlPointerAxisSource = 1
	// WlPointerAxisSource_Continuous continuous coordinate space
	WlPointerAxisSource_Continuous WlPointerAxisSource = 2
	// WlPointerAxisSource_WheelTilt a physical wheel tilt
	WlPointerAxisSource_WheelTilt WlPointerAxisSource = 3
)

// WlPointerSetCursorRequest set the pointer surface
/*
	Set the pointer surface, i.e., the surface that contains the
	pointer image (cursor). This request gives the surface the role
	of a cursor. If the surface already has another role, it raises
	a protocol error.

	The cursor actually changes only if the pointer
	focus for this device is one of the requesting client's surfaces
	or the surface parameter is the current pointer surface. If
	there was a previous surface set with this request it is
	replaced. If surface is NULL, the pointer image is hidden.

	The parameters hotspot_x and hotspot_y define the position of
	the pointer surface relative to the pointer location. Its
	top-left corner is always at (x, y) - (hotspot_x, hotspot_y),
	where (x, y) are the coordinates of the pointer location, in
	surface-local coordinates.

	On surface.attach requests to the pointer surface, hotspot_x
	and hotspot_y are decremented by the x and y parameters
	passed to the request. Attach must be confirmed by
	wl_surface.commit as usual.

	The hotspot can also be updated by passing the currently set
	pointer surface to this request with new values for hotspot_x
	and hotspot_y.

	The current and pending input regions of the wl_surface are
	cleared, and wl_surface.set_input_region is ignored until the
	wl_surface is no longer used as the cursor. When the use as a
	cursor ends, the current and pending input regions become
	undefined, and the wl_surface is unmapped.

	The serial parameter must match the latest wl_pointer.enter
	serial number sent to the client. Otherwise the request will be
	ignored.
*/
type WlPointerSetCursorRequest struct {
	// Serial  number of the enter event
	Serial uint32 `wayland:"uint"`
	// Surface pointer surface
	Surface uint32 `wayland:"object"`
	// HotspotX surface-local x coordinate
	HotspotX int32 `wayland:"int"`
	// HotspotY surface-local y coordinate
	HotspotY int32 `wayland:"int"`
}

// WlPointerReleaseRequest release the pointer object
/*
	Using this request a client can tell the server that it is not going to
	use the pointer object anymore.

	This request destroys the pointer proxy object, so clients must not call
	wl_pointer_destroy() after using this request.
*/
type WlPointerReleaseRequest struct{}

// WlPointerEnterEvent enter event
/*
	Notification that this seat's pointer is focused on a certain
	surface.

	When a seat's focus enters a surface, the pointer image
	is undefined and a client should respond to this event by setting
	an appropriate pointer image with the set_cursor request.
*/
type WlPointerEnterEvent struct {
	// Serial  number of the enter event
	Serial uint32 `wayland:"uint"`
	// Surface  entered by the pointer
	Surface uint32 `wayland:"object"`
	// SurfaceX surface-local x coordinate
	SurfaceX float64 `wayland:"fixed"`
	// SurfaceY surface-local y coordinate
	SurfaceY float64 `wayland:"fixed"`
}

func (e *WlPointerEnterEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Serial = UnmarshallUint32(offset, d)
	offset, e.Surface = UnmarshallUint32(offset, d)
	offset, e.SurfaceX = UnmarshallFixed(offset, d)
	offset, e.SurfaceY = UnmarshallFixed(offset, d)
	return nil
}

// WlPointerLeaveEvent leave event
/*
	Notification that this seat's pointer is no longer focused on
	a certain surface.

	The leave notification is sent before the enter notification
	for the new focus.
*/
type WlPointerLeaveEvent struct {
	// Serial  number of the leave event
	Serial uint32 `wayland:"uint"`
	// Surface  left by the pointer
	Surface uint32 `wayland:"object"`
}

func (e *WlPointerLeaveEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Serial = UnmarshallUint32(offset, d)
	offset, e.Surface = UnmarshallUint32(offset, d)
	return nil
}

// WlPointerMotionEvent pointer motion event
/*
	Notification of pointer location change. The arguments
	surface_x and surface_y are the location relative to the
	focused surface.
*/
type WlPointerMotionEvent struct {
	// Time stamp with millisecond granularity
	Time uint32 `wayland:"uint"`
	// SurfaceX surface-local x coordinate
	SurfaceX float64 `wayland:"fixed"`
	// SurfaceY surface-local y coordinate
	SurfaceY float64 `wayland:"fixed"`
}

func (e *WlPointerMotionEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Time = UnmarshallUint32(offset, d)
	offset, e.SurfaceX = UnmarshallFixed(offset, d)
	offset, e.SurfaceY = UnmarshallFixed(offset, d)
	return nil
}

// WlPointerButtonEvent pointer button event
/*
	Mouse button click and release notifications.

	The location of the click is given by the last motion or
	enter event.
	The time argument is a timestamp with millisecond
	granularity, with an undefined base.

	The button is a button code as defined in the Linux kernel's
	linux/input-event-codes.h header file, e.g. BTN_LEFT.

	Any 16-bit button code value is reserved for future additions to the
	kernel's event code list. All other button codes above 0xFFFF are
	currently undefined but may be used in future versions of this
	protocol.
*/
type WlPointerButtonEvent struct {
	// Serial  number of the button event
	Serial uint32 `wayland:"uint"`
	// Time stamp with millisecond granularity
	Time uint32 `wayland:"uint"`
	// Button  that produced the event
	Button uint32 `wayland:"uint"`
	// State physical state of the button
	State uint32 `wayland:"uint"`
}

func (e *WlPointerButtonEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Serial = UnmarshallUint32(offset, d)
	offset, e.Time = UnmarshallUint32(offset, d)
	offset, e.Button = UnmarshallUint32(offset, d)
	offset, e.State = UnmarshallUint32(offset, d)
	return nil
}

// WlPointerAxisEvent axis event
/*
	Scroll and other axis notifications.

	For scroll events (vertical and horizontal scroll axes), the
	value parameter is the length of a vector along the specified
	axis in a coordinate space identical to those of motion events,
	representing a relative movement along the specified axis.

	For devices that support movements non-parallel to axes multiple
	axis events will be emitted.

	When applicable, for example for touch pads, the server can
	choose to emit scroll events where the motion vector is
	equivalent to a motion event vector.

	When applicable, a client can transform its content relative to the
	scroll distance.
*/
type WlPointerAxisEvent struct {
	// Time stamp with millisecond granularity
	Time uint32 `wayland:"uint"`
	// Axis  type
	Axis uint32 `wayland:"uint"`
	// Value length of vector in surface-local coordinate space
	Value float64 `wayland:"fixed"`
}

func (e *WlPointerAxisEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Time = UnmarshallUint32(offset, d)
	offset, e.Axis = UnmarshallUint32(offset, d)
	offset, e.Value = UnmarshallFixed(offset, d)
	return nil
}

// WlPointerFrameEvent end of a pointer event sequence
/*
	Indicates the end of a set of events that logically belong together.
	A client is expected to accumulate the data in all events within the
	frame before proceeding.

	All wl_pointer events before a wl_pointer.frame event belong
	logically together. For example, in a diagonal scroll motion the
	compositor will send an optional wl_pointer.axis_source event, two
	wl_pointer.axis events (horizontal and vertical) and finally a
	wl_pointer.frame event. The client may use this information to
	calculate a diagonal vector for scrolling.

	When multiple wl_pointer.axis events occur within the same frame,
	the motion vector is the combined motion of all events.
	When a wl_pointer.axis and a wl_pointer.axis_stop event occur within
	the same frame, this indicates that axis movement in one axis has
	stopped but continues in the other axis.
	When multiple wl_pointer.axis_stop events occur within the same
	frame, this indicates that these axes stopped in the same instance.

	A wl_pointer.frame event is sent for every logical event group,
	even if the group only contains a single wl_pointer event.
	Specifically, a client may get a sequence: motion, frame, button,
	frame, axis, frame, axis_stop, frame.

	The wl_pointer.enter and wl_pointer.leave events are logical events
	generated by the compositor and not the hardware. These events are
	also grouped by a wl_pointer.frame. When a pointer moves from one
	surface to another, a compositor should group the
	wl_pointer.leave event within the same wl_pointer.frame.
	However, a client must not rely on wl_pointer.leave and
	wl_pointer.enter being in the same wl_pointer.frame.
	Compositor-specific policies may require the wl_pointer.leave and
	wl_pointer.enter event being split across multiple wl_pointer.frame
	groups.
*/
type WlPointerFrameEvent struct{}

func (e *WlPointerFrameEvent) UnmarshallBinary(d []byte) error {
	return nil
}

// WlPointerAxisSourceEvent axis source event
/*
	Source information for scroll and other axes.

	This event does not occur on its own. It is sent before a
	wl_pointer.frame event and carries the source information for
	all events within that frame.

	The source specifies how this event was generated. If the source is
	wl_pointer.axis_source.finger, a wl_pointer.axis_stop event will be
	sent when the user lifts the finger off the device.

	If the source is wl_pointer.axis_source.wheel,
	wl_pointer.axis_source.wheel_tilt or
	wl_pointer.axis_source.continuous, a wl_pointer.axis_stop event may
	or may not be sent. Whether a compositor sends an axis_stop event
	for these sources is hardware-specific and implementation-dependent;
	clients must not rely on receiving an axis_stop event for these
	scroll sources and should treat scroll sequences from these scroll
	sources as unterminated by default.

	This event is optional. If the source is unknown for a particular
	axis event sequence, no event is sent.
	Only one wl_pointer.axis_source event is permitted per frame.

	The order of wl_pointer.axis_discrete and wl_pointer.axis_source is
	not guaranteed.
*/
type WlPointerAxisSourceEvent struct {
	// AxisSource source of the axis event
	AxisSource uint32 `wayland:"uint"`
}

func (e *WlPointerAxisSourceEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.AxisSource = UnmarshallUint32(offset, d)
	return nil
}

// WlPointerAxisStopEvent axis stop event
/*
	Stop notification for scroll and other axes.

	For some wl_pointer.axis_source types, a wl_pointer.axis_stop event
	is sent to notify a client that the axis sequence has terminated.
	This enables the client to implement kinetic scrolling.
	See the wl_pointer.axis_source documentation for information on when
	this event may be generated.

	Any wl_pointer.axis events with the same axis_source after this
	event should be considered as the start of a new axis motion.

	The timestamp is to be interpreted identical to the timestamp in the
	wl_pointer.axis event. The timestamp value may be the same as a
	preceding wl_pointer.axis event.
*/
type WlPointerAxisStopEvent struct {
	// Time stamp with millisecond granularity
	Time uint32 `wayland:"uint"`
	// Axis the axis stopped with this event
	Axis uint32 `wayland:"uint"`
}

func (e *WlPointerAxisStopEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Time = UnmarshallUint32(offset, d)
	offset, e.Axis = UnmarshallUint32(offset, d)
	return nil
}

// WlPointerAxisDiscreteEvent axis click event
/*
	Discrete step information for scroll and other axes.

	This event carries the axis value of the wl_pointer.axis event in
	discrete steps (e.g. mouse wheel clicks).

	This event is deprecated with wl_pointer version 8 - this event is not
	sent to clients supporting version 8 or later.

	This event does not occur on its own, it is coupled with a
	wl_pointer.axis event that represents this axis value on a
	continuous scale. The protocol guarantees that each axis_discrete
	event is always followed by exactly one axis event with the same
	axis number within the same wl_pointer.frame. Note that the protocol
	allows for other events to occur between the axis_discrete and
	its coupled axis event, including other axis_discrete or axis
	events. A wl_pointer.frame must not contain more than one axis_discrete
	event per axis type.

	This event is optional; continuous scrolling devices
	like two-finger scrolling on touchpads do not have discrete
	steps and do not generate this event.

	The discrete value carries the directional information. e.g. a value
	of -2 is two steps towards the negative direction of this axis.

	The axis number is identical to the axis number in the associated
	axis event.

	The order of wl_pointer.axis_discrete and wl_pointer.axis_source is
	not guaranteed.
*/
type WlPointerAxisDiscreteEvent struct {
	// Axis  type
	Axis uint32 `wayland:"uint"`
	// Discrete number of steps
	Discrete int32 `wayland:"int"`
}

func (e *WlPointerAxisDiscreteEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Axis = UnmarshallUint32(offset, d)
	offset, e.Discrete = UnmarshallInt32(offset, d)
	return nil
}

// WlPointerAxisValue120Event axis high-resolution scroll event
/*
	Discrete high-resolution scroll information.

	This event carries high-resolution wheel scroll information,
	with each multiple of 120 representing one logical scroll step
	(a wheel detent). For example, an axis_value120 of 30 is one quarter of
	a logical scroll step in the positive direction, a value120 of
	-240 are two logical scroll steps in the negative direction within the
	same hardware event.
	Clients that rely on discrete scrolling should accumulate the
	value120 to multiples of 120 before processing the event.

	The value120 must not be zero.

	This event replaces the wl_pointer.axis_discrete event in clients
	supporting wl_pointer version 8 or later.

	Where a wl_pointer.axis_source event occurs in the same
	wl_pointer.frame, the axis source applies to this event.

	The order of wl_pointer.axis_value120 and wl_pointer.axis_source is
	not guaranteed.
*/
type WlPointerAxisValue120Event struct {
	// Axis  type
	Axis uint32 `wayland:"uint"`
	// Value120 scroll distance as fraction of 120
	Value120 int32 `wayland:"int"`
}

func (e *WlPointerAxisValue120Event) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Axis = UnmarshallUint32(offset, d)
	offset, e.Value120 = UnmarshallInt32(offset, d)
	return nil
}

type WlPointerListener interface {
	Enter(WlPointerEnterEvent)
	Leave(WlPointerLeaveEvent)
	Motion(WlPointerMotionEvent)
	Button(WlPointerButtonEvent)
	Axis(WlPointerAxisEvent)
	Frame(WlPointerFrameEvent)
	AxisSource(WlPointerAxisSourceEvent)
	AxisStop(WlPointerAxisStopEvent)
	AxisDiscrete(WlPointerAxisDiscreteEvent)
	AxisValue120(WlPointerAxisValue120Event)
}
type UnimplementedWlPointerListener struct{}

func (e *UnimplementedWlPointerListener) Enter(_ WlPointerEnterEvent) {
	return
}
func (e *UnimplementedWlPointerListener) Leave(_ WlPointerLeaveEvent) {
	return
}
func (e *UnimplementedWlPointerListener) Motion(_ WlPointerMotionEvent) {
	return
}
func (e *UnimplementedWlPointerListener) Button(_ WlPointerButtonEvent) {
	return
}
func (e *UnimplementedWlPointerListener) Axis(_ WlPointerAxisEvent) {
	return
}
func (e *UnimplementedWlPointerListener) Frame(_ WlPointerFrameEvent) {
	return
}
func (e *UnimplementedWlPointerListener) AxisSource(_ WlPointerAxisSourceEvent) {
	return
}
func (e *UnimplementedWlPointerListener) AxisStop(_ WlPointerAxisStopEvent) {
	return
}
func (e *UnimplementedWlPointerListener) AxisDiscrete(_ WlPointerAxisDiscreteEvent) {
	return
}
func (e *UnimplementedWlPointerListener) AxisValue120(_ WlPointerAxisValue120Event) {
	return
}

type WlPointerDispatcher struct {
	WlPointerListener
}

func NewWlPointerDispatcher() *WlPointerDispatcher {
	return &WlPointerDispatcher{WlPointerListener: &UnimplementedWlPointerListener{}}
}
func (i *WlPointerDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	case 0:
		var e WlPointerEnterEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Enter(e)
	case 1:
		var e WlPointerLeaveEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Leave(e)
	case 2:
		var e WlPointerMotionEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Motion(e)
	case 3:
		var e WlPointerButtonEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Button(e)
	case 4:
		var e WlPointerAxisEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Axis(e)
	case 5:
		var e WlPointerFrameEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Frame(e)
	case 6:
		var e WlPointerAxisSourceEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.AxisSource(e)
	case 7:
		var e WlPointerAxisStopEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.AxisStop(e)
	case 8:
		var e WlPointerAxisDiscreteEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.AxisDiscrete(e)
	case 9:
		var e WlPointerAxisValue120Event
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.AxisValue120(e)
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlPointerDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlPointerListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlPointerListener = listener
}

// WlKeyboard keyboard input device
/*
	The wl_keyboard interface represents one or more keyboards
	associated with a seat.
*/
type WlKeyboard struct {
	Proxy
}

func NewWlKeyboard(p Proxy) WlKeyboard {
	return WlKeyboard{Proxy: p}
}
func (w *WlKeyboard) AttachListener(l WlKeyboardListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

// WlKeyboardKeymapFormat keyboard mapping format
/*
	This specifies the format of the keymap provided to the
	client with the wl_keyboard.keymap event.
*/
type WlKeyboardKeymapFormat uint32

const (
	// WlKeyboardKeymapFormat_NoKeymap no keymap; client must understand how to interpret the raw keycode
	WlKeyboardKeymapFormat_NoKeymap WlKeyboardKeymapFormat = 0
	// WlKeyboardKeymapFormat_XkbV1 libxkbcommon compatible, null-terminated string; to determine the xkb keycode, clients must add 8 to the key event keycode
	WlKeyboardKeymapFormat_XkbV1 WlKeyboardKeymapFormat = 1
)

// WlKeyboardKeyState physical key state
//
//	Describes the physical state of a key that produced the key event.
type WlKeyboardKeyState uint32

const (
	// WlKeyboardKeyState_Released key is not pressed
	WlKeyboardKeyState_Released WlKeyboardKeyState = 0
	// WlKeyboardKeyState_Pressed key is pressed
	WlKeyboardKeyState_Pressed WlKeyboardKeyState = 1
)

// WlKeyboardReleaseRequest release the keyboard object
type WlKeyboardReleaseRequest struct{}

// WlKeyboardKeymapEvent keyboard mapping
/*
	This event provides a file descriptor to the client which can be
	memory-mapped in read-only mode to provide a keyboard mapping
	description.

	From version 7 onwards, the fd must be mapped with MAP_PRIVATE by
	the recipient, as MAP_SHARED may fail.
*/
type WlKeyboardKeymapEvent struct {
	// Format keymap format
	Format uint32 `wayland:"uint"`
	// Fd keymap file descriptor
	Fd uintptr `wayland:"fd"`
	// Size keymap size, in bytes
	Size uint32 `wayland:"uint"`
}

func (e *WlKeyboardKeymapEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Format = UnmarshallUint32(offset, d)
	offset, e.Fd = UnmarshallFd(offset, d)
	offset, e.Size = UnmarshallUint32(offset, d)
	return nil
}

// WlKeyboardEnterEvent enter event
/*
	Notification that this seat's keyboard focus is on a certain
	surface.

	The compositor must send the wl_keyboard.modifiers event after this
	event.
*/
type WlKeyboardEnterEvent struct {
	// Serial  number of the enter event
	Serial uint32 `wayland:"uint"`
	// Surface  gaining keyboard focus
	Surface uint32 `wayland:"object"`
	// Keys the currently pressed keys
	Keys []byte `wayland:"array"`
}

func (e *WlKeyboardEnterEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Serial = UnmarshallUint32(offset, d)
	offset, e.Surface = UnmarshallUint32(offset, d)
	offset, e.Keys = UnmarshallArray(offset, d)
	return nil
}

// WlKeyboardLeaveEvent leave event
/*
	Notification that this seat's keyboard focus is no longer on
	a certain surface.

	The leave notification is sent before the enter notification
	for the new focus.

	After this event client must assume that all keys, including modifiers,
	are lifted and also it must stop key repeating if there's some going on.
*/
type WlKeyboardLeaveEvent struct {
	// Serial  number of the leave event
	Serial uint32 `wayland:"uint"`
	// Surface  that lost keyboard focus
	Surface uint32 `wayland:"object"`
}

func (e *WlKeyboardLeaveEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Serial = UnmarshallUint32(offset, d)
	offset, e.Surface = UnmarshallUint32(offset, d)
	return nil
}

// WlKeyboardKeyEvent key event
/*
	A key was pressed or released.
	The time argument is a timestamp with millisecond
	granularity, with an undefined base.

	The key is a platform-specific key code that can be interpreted
	by feeding it to the keyboard mapping (see the keymap event).

	If this event produces a change in modifiers, then the resulting
	wl_keyboard.modifiers event must be sent after this event.
*/
type WlKeyboardKeyEvent struct {
	// Serial  number of the key event
	Serial uint32 `wayland:"uint"`
	// Time stamp with millisecond granularity
	Time uint32 `wayland:"uint"`
	// Key  that produced the event
	Key uint32 `wayland:"uint"`
	// State physical state of the key
	State uint32 `wayland:"uint"`
}

func (e *WlKeyboardKeyEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Serial = UnmarshallUint32(offset, d)
	offset, e.Time = UnmarshallUint32(offset, d)
	offset, e.Key = UnmarshallUint32(offset, d)
	offset, e.State = UnmarshallUint32(offset, d)
	return nil
}

// WlKeyboardModifiersEvent modifier and group state
/*
	Notifies clients that the modifier and/or group state has
	changed, and it should update its local state.
*/
type WlKeyboardModifiersEvent struct {
	// Serial  number of the modifiers event
	Serial uint32 `wayland:"uint"`
	// ModsDepressed depressed modifiers
	ModsDepressed uint32 `wayland:"uint"`
	// ModsLatched latched modifiers
	ModsLatched uint32 `wayland:"uint"`
	// ModsLocked locked modifiers
	ModsLocked uint32 `wayland:"uint"`
	// Group keyboard layout
	Group uint32 `wayland:"uint"`
}

func (e *WlKeyboardModifiersEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Serial = UnmarshallUint32(offset, d)
	offset, e.ModsDepressed = UnmarshallUint32(offset, d)
	offset, e.ModsLatched = UnmarshallUint32(offset, d)
	offset, e.ModsLocked = UnmarshallUint32(offset, d)
	offset, e.Group = UnmarshallUint32(offset, d)
	return nil
}

// WlKeyboardRepeatInfoEvent repeat rate and delay
/*
	Informs the client about the keyboard's repeat rate and delay.

	This event is sent as soon as the wl_keyboard object has been created,
	and is guaranteed to be received by the client before any key press
	event.

	Negative values for either rate or delay are illegal. A rate of zero
	will disable any repeating (regardless of the value of delay).

	This event can be sent later on as well with a new value if necessary,
	so clients should continue listening for the event past the creation
	of wl_keyboard.
*/
type WlKeyboardRepeatInfoEvent struct {
	// Rate the rate of repeating keys in characters per second
	Rate int32 `wayland:"int"`
	// Delay  in milliseconds since key down until repeating starts
	Delay int32 `wayland:"int"`
}

func (e *WlKeyboardRepeatInfoEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Rate = UnmarshallInt32(offset, d)
	offset, e.Delay = UnmarshallInt32(offset, d)
	return nil
}

type WlKeyboardListener interface {
	Keymap(WlKeyboardKeymapEvent)
	Enter(WlKeyboardEnterEvent)
	Leave(WlKeyboardLeaveEvent)
	Key(WlKeyboardKeyEvent)
	Modifiers(WlKeyboardModifiersEvent)
	RepeatInfo(WlKeyboardRepeatInfoEvent)
}
type UnimplementedWlKeyboardListener struct{}

func (e *UnimplementedWlKeyboardListener) Keymap(_ WlKeyboardKeymapEvent) {
	return
}
func (e *UnimplementedWlKeyboardListener) Enter(_ WlKeyboardEnterEvent) {
	return
}
func (e *UnimplementedWlKeyboardListener) Leave(_ WlKeyboardLeaveEvent) {
	return
}
func (e *UnimplementedWlKeyboardListener) Key(_ WlKeyboardKeyEvent) {
	return
}
func (e *UnimplementedWlKeyboardListener) Modifiers(_ WlKeyboardModifiersEvent) {
	return
}
func (e *UnimplementedWlKeyboardListener) RepeatInfo(_ WlKeyboardRepeatInfoEvent) {
	return
}

type WlKeyboardDispatcher struct {
	WlKeyboardListener
}

func NewWlKeyboardDispatcher() *WlKeyboardDispatcher {
	return &WlKeyboardDispatcher{WlKeyboardListener: &UnimplementedWlKeyboardListener{}}
}
func (i *WlKeyboardDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	case 0:
		var e WlKeyboardKeymapEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Keymap(e)
	case 1:
		var e WlKeyboardEnterEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Enter(e)
	case 2:
		var e WlKeyboardLeaveEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Leave(e)
	case 3:
		var e WlKeyboardKeyEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Key(e)
	case 4:
		var e WlKeyboardModifiersEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Modifiers(e)
	case 5:
		var e WlKeyboardRepeatInfoEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.RepeatInfo(e)
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlKeyboardDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlKeyboardListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlKeyboardListener = listener
}

// WlTouch touchscreen input device
/*
	The wl_touch interface represents a touchscreen
	associated with a seat.

	Touch interactions can consist of one or more contacts.
	For each contact, a series of events is generated, starting
	with a down event, followed by zero or more motion events,
	and ending with an up event. Events relating to the same
	contact point can be identified by the ID of the sequence.
*/
type WlTouch struct {
	Proxy
}

func NewWlTouch(p Proxy) WlTouch {
	return WlTouch{Proxy: p}
}
func (w *WlTouch) AttachListener(l WlTouchListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

// WlTouchReleaseRequest release the touch object
type WlTouchReleaseRequest struct{}

// WlTouchDownEvent touch down event and beginning of a touch sequence
/*
	A new touch point has appeared on the surface. This touch point is
	assigned a unique ID. Future events from this touch point reference
	this ID. The ID ceases to be valid after a touch up event and may be
	reused in the future.
*/
type WlTouchDownEvent struct {
	// Serial  number of the touch down event
	Serial uint32 `wayland:"uint"`
	// Time stamp with millisecond granularity
	Time uint32 `wayland:"uint"`
	// Surface  touched
	Surface uint32 `wayland:"object"`
	// Id the unique ID of this touch point
	Id int32 `wayland:"int"`
	// X surface-local x coordinate
	X float64 `wayland:"fixed"`
	// Y surface-local y coordinate
	Y float64 `wayland:"fixed"`
}

func (e *WlTouchDownEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Serial = UnmarshallUint32(offset, d)
	offset, e.Time = UnmarshallUint32(offset, d)
	offset, e.Surface = UnmarshallUint32(offset, d)
	offset, e.Id = UnmarshallInt32(offset, d)
	offset, e.X = UnmarshallFixed(offset, d)
	offset, e.Y = UnmarshallFixed(offset, d)
	return nil
}

// WlTouchUpEvent end of a touch event sequence
/*
	The touch point has disappeared. No further events will be sent for
	this touch point and the touch point's ID is released and may be
	reused in a future touch down event.
*/
type WlTouchUpEvent struct {
	// Serial  number of the touch up event
	Serial uint32 `wayland:"uint"`
	// Time stamp with millisecond granularity
	Time uint32 `wayland:"uint"`
	// Id the unique ID of this touch point
	Id int32 `wayland:"int"`
}

func (e *WlTouchUpEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Serial = UnmarshallUint32(offset, d)
	offset, e.Time = UnmarshallUint32(offset, d)
	offset, e.Id = UnmarshallInt32(offset, d)
	return nil
}

// WlTouchMotionEvent update of touch point coordinates
//
//	A touch point has changed coordinates.
type WlTouchMotionEvent struct {
	// Time stamp with millisecond granularity
	Time uint32 `wayland:"uint"`
	// Id the unique ID of this touch point
	Id int32 `wayland:"int"`
	// X surface-local x coordinate
	X float64 `wayland:"fixed"`
	// Y surface-local y coordinate
	Y float64 `wayland:"fixed"`
}

func (e *WlTouchMotionEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Time = UnmarshallUint32(offset, d)
	offset, e.Id = UnmarshallInt32(offset, d)
	offset, e.X = UnmarshallFixed(offset, d)
	offset, e.Y = UnmarshallFixed(offset, d)
	return nil
}

// WlTouchFrameEvent end of touch frame event
/*
	Indicates the end of a set of events that logically belong together.
	A client is expected to accumulate the data in all events within the
	frame before proceeding.

	A wl_touch.frame terminates at least one event but otherwise no
	guarantee is provided about the set of events within a frame. A client
	must assume that any state not updated in a frame is unchanged from the
	previously known state.
*/
type WlTouchFrameEvent struct{}

func (e *WlTouchFrameEvent) UnmarshallBinary(d []byte) error {
	return nil
}

// WlTouchCancelEvent touch session cancelled
/*
	Sent if the compositor decides the touch stream is a global
	gesture. No further events are sent to the clients from that
	particular gesture. Touch cancellation applies to all touch points
	currently active on this client's surface. The client is
	responsible for finalizing the touch points, future touch points on
	this surface may reuse the touch point ID.
*/
type WlTouchCancelEvent struct{}

func (e *WlTouchCancelEvent) UnmarshallBinary(d []byte) error {
	return nil
}

// WlTouchShapeEvent update shape of touch point
/*
	Sent when a touchpoint has changed its shape.

	This event does not occur on its own. It is sent before a
	wl_touch.frame event and carries the new shape information for
	any previously reported, or new touch points of that frame.

	Other events describing the touch point such as wl_touch.down,
	wl_touch.motion or wl_touch.orientation may be sent within the
	same wl_touch.frame. A client should treat these events as a single
	logical touch point update. The order of wl_touch.shape,
	wl_touch.orientation and wl_touch.motion is not guaranteed.
	A wl_touch.down event is guaranteed to occur before the first
	wl_touch.shape event for this touch ID but both events may occur within
	the same wl_touch.frame.

	A touchpoint shape is approximated by an ellipse through the major and
	minor axis length. The major axis length describes the longer diameter
	of the ellipse, while the minor axis length describes the shorter
	diameter. Major and minor are orthogonal and both are specified in
	surface-local coordinates. The center of the ellipse is always at the
	touchpoint location as reported by wl_touch.down or wl_touch.move.

	This event is only sent by the compositor if the touch device supports
	shape reports. The client has to make reasonable assumptions about the
	shape if it did not receive this event.
*/
type WlTouchShapeEvent struct {
	// Id the unique ID of this touch point
	Id int32 `wayland:"int"`
	// Major length of the major axis in surface-local coordinates
	Major float64 `wayland:"fixed"`
	// Minor length of the minor axis in surface-local coordinates
	Minor float64 `wayland:"fixed"`
}

func (e *WlTouchShapeEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Id = UnmarshallInt32(offset, d)
	offset, e.Major = UnmarshallFixed(offset, d)
	offset, e.Minor = UnmarshallFixed(offset, d)
	return nil
}

// WlTouchOrientationEvent update orientation of touch point
/*
	Sent when a touchpoint has changed its orientation.

	This event does not occur on its own. It is sent before a
	wl_touch.frame event and carries the new shape information for
	any previously reported, or new touch points of that frame.

	Other events describing the touch point such as wl_touch.down,
	wl_touch.motion or wl_touch.shape may be sent within the
	same wl_touch.frame. A client should treat these events as a single
	logical touch point update. The order of wl_touch.shape,
	wl_touch.orientation and wl_touch.motion is not guaranteed.
	A wl_touch.down event is guaranteed to occur before the first
	wl_touch.orientation event for this touch ID but both events may occur
	within the same wl_touch.frame.

	The orientation describes the clockwise angle of a touchpoint's major
	axis to the positive surface y-axis and is normalized to the -180 to
	+180 degree range. The granularity of orientation depends on the touch
	device, some devices only support binary rotation values between 0 and
	90 degrees.

	This event is only sent by the compositor if the touch device supports
	orientation reports.
*/
type WlTouchOrientationEvent struct {
	// Id the unique ID of this touch point
	Id int32 `wayland:"int"`
	// Orientation angle between major axis and positive surface y-axis in degrees
	Orientation float64 `wayland:"fixed"`
}

func (e *WlTouchOrientationEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Id = UnmarshallInt32(offset, d)
	offset, e.Orientation = UnmarshallFixed(offset, d)
	return nil
}

type WlTouchListener interface {
	Down(WlTouchDownEvent)
	Up(WlTouchUpEvent)
	Motion(WlTouchMotionEvent)
	Frame(WlTouchFrameEvent)
	Cancel(WlTouchCancelEvent)
	Shape(WlTouchShapeEvent)
	Orientation(WlTouchOrientationEvent)
}
type UnimplementedWlTouchListener struct{}

func (e *UnimplementedWlTouchListener) Down(_ WlTouchDownEvent) {
	return
}
func (e *UnimplementedWlTouchListener) Up(_ WlTouchUpEvent) {
	return
}
func (e *UnimplementedWlTouchListener) Motion(_ WlTouchMotionEvent) {
	return
}
func (e *UnimplementedWlTouchListener) Frame(_ WlTouchFrameEvent) {
	return
}
func (e *UnimplementedWlTouchListener) Cancel(_ WlTouchCancelEvent) {
	return
}
func (e *UnimplementedWlTouchListener) Shape(_ WlTouchShapeEvent) {
	return
}
func (e *UnimplementedWlTouchListener) Orientation(_ WlTouchOrientationEvent) {
	return
}

type WlTouchDispatcher struct {
	WlTouchListener
}

func NewWlTouchDispatcher() *WlTouchDispatcher {
	return &WlTouchDispatcher{WlTouchListener: &UnimplementedWlTouchListener{}}
}
func (i *WlTouchDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	case 0:
		var e WlTouchDownEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Down(e)
	case 1:
		var e WlTouchUpEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Up(e)
	case 2:
		var e WlTouchMotionEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Motion(e)
	case 3:
		var e WlTouchFrameEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Frame(e)
	case 4:
		var e WlTouchCancelEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Cancel(e)
	case 5:
		var e WlTouchShapeEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Shape(e)
	case 6:
		var e WlTouchOrientationEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Orientation(e)
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlTouchDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlTouchListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlTouchListener = listener
}

// WlOutput compositor output region
/*
	An output describes part of the compositor geometry.  The
	compositor works in the 'compositor coordinate system' and an
	output corresponds to a rectangular area in that space that is
	actually visible.  This typically corresponds to a monitor that
	displays part of the compositor space.  This object is published
	as global during start up, or when a monitor is hotplugged.
*/
type WlOutput struct {
	Proxy
}

func NewWlOutput(p Proxy) WlOutput {
	return WlOutput{Proxy: p}
}
func (w *WlOutput) AttachListener(l WlOutputListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

// WlOutputSubpixel subpixel geometry information
/*
	This enumeration describes how the physical
	pixels on an output are laid out.
*/
type WlOutputSubpixel uint32

const (
	// WlOutputSubpixel_Unknown unknown geometry
	WlOutputSubpixel_Unknown WlOutputSubpixel = 0
	// WlOutputSubpixel_None no geometry
	WlOutputSubpixel_None WlOutputSubpixel = 1
	// WlOutputSubpixel_HorizontalRgb horizontal RGB
	WlOutputSubpixel_HorizontalRgb WlOutputSubpixel = 2
	// WlOutputSubpixel_HorizontalBgr horizontal BGR
	WlOutputSubpixel_HorizontalBgr WlOutputSubpixel = 3
	// WlOutputSubpixel_VerticalRgb vertical RGB
	WlOutputSubpixel_VerticalRgb WlOutputSubpixel = 4
	// WlOutputSubpixel_VerticalBgr vertical BGR
	WlOutputSubpixel_VerticalBgr WlOutputSubpixel = 5
)

// WlOutputTransform transform from framebuffer to output
/*
	This describes the transform that a compositor will apply to a
	surface to compensate for the rotation or mirroring of an
	output device.

	The flipped values correspond to an initial flip around a
	vertical axis followed by rotation.

	The purpose is mainly to allow clients to render accordingly and
	tell the compositor, so that for fullscreen surfaces, the
	compositor will still be able to scan out directly from client
	surfaces.
*/
type WlOutputTransform uint32

const (
	// WlOutputTransform_Normal no transform
	WlOutputTransform_Normal WlOutputTransform = 0
	// WlOutputTransform_90 90 degrees counter-clockwise
	WlOutputTransform_90 WlOutputTransform = 1
	// WlOutputTransform_180 180 degrees counter-clockwise
	WlOutputTransform_180 WlOutputTransform = 2
	// WlOutputTransform_270 270 degrees counter-clockwise
	WlOutputTransform_270 WlOutputTransform = 3
	// WlOutputTransform_Flipped 180 degree flip around a vertical axis
	WlOutputTransform_Flipped WlOutputTransform = 4
	// WlOutputTransform_Flipped90 flip and rotate 90 degrees counter-clockwise
	WlOutputTransform_Flipped90 WlOutputTransform = 5
	// WlOutputTransform_Flipped180 flip and rotate 180 degrees counter-clockwise
	WlOutputTransform_Flipped180 WlOutputTransform = 6
	// WlOutputTransform_Flipped270 flip and rotate 270 degrees counter-clockwise
	WlOutputTransform_Flipped270 WlOutputTransform = 7
)

// WlOutputMode mode information
/*
	These flags describe properties of an output mode.
	They are used in the flags bitfield of the mode event.
*/
type WlOutputMode uint32

const (
	// WlOutputMode_Current indicates this is the current mode
	WlOutputMode_Current WlOutputMode = 0x1
	// WlOutputMode_Preferred indicates this is the preferred mode
	WlOutputMode_Preferred WlOutputMode = 0x2
)

// WlOutputReleaseRequest release the output object
/*
	Using this request a client can tell the server that it is not going to
	use the output object anymore.
*/
type WlOutputReleaseRequest struct{}

// WlOutputGeometryEvent properties of the output
/*
	The geometry event describes geometric properties of the output.
	The event is sent when binding to the output object and whenever
	any of the properties change.

	The physical size can be set to zero if it doesn't make sense for this
	output (e.g. for projectors or virtual outputs).

	The geometry event will be followed by a done event (starting from
	version 2).

	Note: wl_output only advertises partial information about the output
	position and identification. Some compositors, for instance those not
	implementing a desktop-style output layout or those exposing virtual
	outputs, might fake this information. Instead of using x and y, clients
	should use xdg_output.logical_position. Instead of using make and model,
	clients should use name and description.
*/
type WlOutputGeometryEvent struct {
	// X  position within the global compositor space
	X int32 `wayland:"int"`
	// Y  position within the global compositor space
	Y int32 `wayland:"int"`
	// PhysicalWidth width in millimeters of the output
	PhysicalWidth int32 `wayland:"int"`
	// PhysicalHeight height in millimeters of the output
	PhysicalHeight int32 `wayland:"int"`
	// Subpixel  orientation of the output
	Subpixel int32 `wayland:"int"`
	// Make textual description of the manufacturer
	Make string `wayland:"string"`
	// Model textual description of the model
	Model string `wayland:"string"`
	// Transform  that maps framebuffer to output
	Transform int32 `wayland:"int"`
}

func (e *WlOutputGeometryEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.X = UnmarshallInt32(offset, d)
	offset, e.Y = UnmarshallInt32(offset, d)
	offset, e.PhysicalWidth = UnmarshallInt32(offset, d)
	offset, e.PhysicalHeight = UnmarshallInt32(offset, d)
	offset, e.Subpixel = UnmarshallInt32(offset, d)
	offset, e.Make = UnmarshallString(offset, d)
	offset, e.Model = UnmarshallString(offset, d)
	offset, e.Transform = UnmarshallInt32(offset, d)
	return nil
}

// WlOutputModeEvent advertise available modes for the output
/*
	The mode event describes an available mode for the output.

	The event is sent when binding to the output object and there
	will always be one mode, the current mode.  The event is sent
	again if an output changes mode, for the mode that is now
	current.  In other words, the current mode is always the last
	mode that was received with the current flag set.

	Non-current modes are deprecated. A compositor can decide to only
	advertise the current mode and never send other modes. Clients
	should not rely on non-current modes.

	The size of a mode is given in physical hardware units of
	the output device. This is not necessarily the same as
	the output size in the global compositor space. For instance,
	the output may be scaled, as described in wl_output.scale,
	or transformed, as described in wl_output.transform. Clients
	willing to retrieve the output size in the global compositor
	space should use xdg_output.logical_size instead.

	The vertical refresh rate can be set to zero if it doesn't make
	sense for this output (e.g. for virtual outputs).

	The mode event will be followed by a done event (starting from
	version 2).

	Clients should not use the refresh rate to schedule frames. Instead,
	they should use the wl_surface.frame event or the presentation-time
	protocol.

	Note: this information is not always meaningful for all outputs. Some
	compositors, such as those exposing virtual outputs, might fake the
	refresh rate or the size.
*/
type WlOutputModeEvent struct {
	// Flags bitfield of mode flags
	Flags uint32 `wayland:"uint"`
	// Width  of the mode in hardware units
	Width int32 `wayland:"int"`
	// Height  of the mode in hardware units
	Height int32 `wayland:"int"`
	// Refresh vertical refresh rate in mHz
	Refresh int32 `wayland:"int"`
}

func (e *WlOutputModeEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Flags = UnmarshallUint32(offset, d)
	offset, e.Width = UnmarshallInt32(offset, d)
	offset, e.Height = UnmarshallInt32(offset, d)
	offset, e.Refresh = UnmarshallInt32(offset, d)
	return nil
}

// WlOutputDoneEvent sent all information about output
/*
	This event is sent after all other properties have been
	sent after binding to the output object and after any
	other property changes done after that. This allows
	changes to the output properties to be seen as
	atomic, even if they happen via multiple events.
*/
type WlOutputDoneEvent struct{}

func (e *WlOutputDoneEvent) UnmarshallBinary(d []byte) error {
	return nil
}

// WlOutputScaleEvent output scaling properties
/*
	This event contains scaling geometry information
	that is not in the geometry event. It may be sent after
	binding the output object or if the output scale changes
	later. If it is not sent, the client should assume a
	scale of 1.

	A scale larger than 1 means that the compositor will
	automatically scale surface buffers by this amount
	when rendering. This is used for very high resolution
	displays where applications rendering at the native
	resolution would be too small to be legible.

	It is intended that scaling aware clients track the
	current output of a surface, and if it is on a scaled
	output it should use wl_surface.set_buffer_scale with
	the scale of the output. That way the compositor can
	avoid scaling the surface, and the client can supply
	a higher detail image.

	The scale event will be followed by a done event.
*/
type WlOutputScaleEvent struct {
	// Factor scaling factor of output
	Factor int32 `wayland:"int"`
}

func (e *WlOutputScaleEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Factor = UnmarshallInt32(offset, d)
	return nil
}

// WlOutputNameEvent name of this output
/*
	Many compositors will assign user-friendly names to their outputs, show
	them to the user, allow the user to refer to an output, etc. The client
	may wish to know this name as well to offer the user similar behaviors.

	The name is a UTF-8 string with no convention defined for its contents.
	Each name is unique among all wl_output globals. The name is only
	guaranteed to be unique for the compositor instance.

	The same output name is used for all clients for a given wl_output
	global. Thus, the name can be shared across processes to refer to a
	specific wl_output global.

	The name is not guaranteed to be persistent across sessions, thus cannot
	be used to reliably identify an output in e.g. configuration files.

	Examples of names include 'HDMI-A-1', 'WL-1', 'X11-1', etc. However, do
	not assume that the name is a reflection of an underlying DRM connector,
	X11 connection, etc.

	The name event is sent after binding the output object. This event is
	only sent once per output object, and the name does not change over the
	lifetime of the wl_output global.

	Compositors may re-use the same output name if the wl_output global is
	destroyed and re-created later. Compositors should avoid re-using the
	same name if possible.

	The name event will be followed by a done event.
*/
type WlOutputNameEvent struct {
	// Name output name
	Name string `wayland:"string"`
}

func (e *WlOutputNameEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Name = UnmarshallString(offset, d)
	return nil
}

// WlOutputDescriptionEvent human-readable description of this output
/*
	Many compositors can produce human-readable descriptions of their
	outputs. The client may wish to know this description as well, e.g. for
	output selection purposes.

	The description is a UTF-8 string with no convention defined for its
	contents. The description is not guaranteed to be unique among all
	wl_output globals. Examples might include 'Foocorp 11" Display' or
	'Virtual X11 output via :1'.

	The description event is sent after binding the output object and
	whenever the description changes. The description is optional, and may
	not be sent at all.

	The description event will be followed by a done event.
*/
type WlOutputDescriptionEvent struct {
	// Description output description
	Description string `wayland:"string"`
}

func (e *WlOutputDescriptionEvent) UnmarshallBinary(d []byte) error {
	offset := 0
	offset, e.Description = UnmarshallString(offset, d)
	return nil
}

type WlOutputListener interface {
	Geometry(WlOutputGeometryEvent)
	Mode(WlOutputModeEvent)
	Done(WlOutputDoneEvent)
	Scale(WlOutputScaleEvent)
	Name(WlOutputNameEvent)
	Description(WlOutputDescriptionEvent)
}
type UnimplementedWlOutputListener struct{}

func (e *UnimplementedWlOutputListener) Geometry(_ WlOutputGeometryEvent) {
	return
}
func (e *UnimplementedWlOutputListener) Mode(_ WlOutputModeEvent) {
	return
}
func (e *UnimplementedWlOutputListener) Done(_ WlOutputDoneEvent) {
	return
}
func (e *UnimplementedWlOutputListener) Scale(_ WlOutputScaleEvent) {
	return
}
func (e *UnimplementedWlOutputListener) Name(_ WlOutputNameEvent) {
	return
}
func (e *UnimplementedWlOutputListener) Description(_ WlOutputDescriptionEvent) {
	return
}

type WlOutputDispatcher struct {
	WlOutputListener
}

func NewWlOutputDispatcher() *WlOutputDispatcher {
	return &WlOutputDispatcher{WlOutputListener: &UnimplementedWlOutputListener{}}
}
func (i *WlOutputDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	case 0:
		var e WlOutputGeometryEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Geometry(e)
	case 1:
		var e WlOutputModeEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Mode(e)
	case 2:
		var e WlOutputDoneEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Done(e)
	case 3:
		var e WlOutputScaleEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Scale(e)
	case 4:
		var e WlOutputNameEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Name(e)
	case 5:
		var e WlOutputDescriptionEvent
		err := e.UnmarshallBinary(b)
		if err != nil {
			return err
		}
		i.Description(e)
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlOutputDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlOutputListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlOutputListener = listener
}

// WlRegion region interface
/*
	A region object describes an area.

	Region objects are used to describe the opaque and input
	regions of a surface.
*/
type WlRegion struct {
	Proxy
}

func NewWlRegion(p Proxy) WlRegion {
	return WlRegion{Proxy: p}
}
func (w *WlRegion) AttachListener(l WlRegionListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

// WlRegionDestroyRequest destroy region
//
//	Destroy the region.  This will invalidate the object ID.
type WlRegionDestroyRequest struct{}

// WlRegionAddRequest add rectangle to region
//
//	Add the specified rectangle to the region.
type WlRegionAddRequest struct {
	// X region-local x coordinate
	X int32 `wayland:"int"`
	// Y region-local y coordinate
	Y int32 `wayland:"int"`
	// Width rectangle width
	Width int32 `wayland:"int"`
	// Height rectangle height
	Height int32 `wayland:"int"`
}

// WlRegionSubtractRequest subtract rectangle from region
//
//	Subtract the specified rectangle from the region.
type WlRegionSubtractRequest struct {
	// X region-local x coordinate
	X int32 `wayland:"int"`
	// Y region-local y coordinate
	Y int32 `wayland:"int"`
	// Width rectangle width
	Width int32 `wayland:"int"`
	// Height rectangle height
	Height int32 `wayland:"int"`
}
type WlRegionListener interface{}
type UnimplementedWlRegionListener struct{}
type WlRegionDispatcher struct {
	WlRegionListener
}

func NewWlRegionDispatcher() *WlRegionDispatcher {
	return &WlRegionDispatcher{WlRegionListener: &UnimplementedWlRegionListener{}}
}
func (i *WlRegionDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlRegionDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlRegionListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlRegionListener = listener
}

// WlSubcompositor sub-surface compositing
/*
	The global interface exposing sub-surface compositing capabilities.
	A wl_surface, that has sub-surfaces associated, is called the
	parent surface. Sub-surfaces can be arbitrarily nested and create
	a tree of sub-surfaces.

	The root surface in a tree of sub-surfaces is the main
	surface. The main surface cannot be a sub-surface, because
	sub-surfaces must always have a parent.

	A main surface with its sub-surfaces forms a (compound) window.
	For window management purposes, this set of wl_surface objects is
	to be considered as a single window, and it should also behave as
	such.

	The aim of sub-surfaces is to offload some of the compositing work
	within a window from clients to the compositor. A prime example is
	a video player with decorations and video in separate wl_surface
	objects. This should allow the compositor to pass YUV video buffer
	processing to dedicated overlay hardware when possible.
*/
type WlSubcompositor struct {
	Proxy
}

func NewWlSubcompositor(p Proxy) WlSubcompositor {
	return WlSubcompositor{Proxy: p}
}
func (w *WlSubcompositor) AttachListener(l WlSubcompositorListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

type WlSubcompositorError uint32

const (
	// WlSubcompositorError_BadSurface the to-be sub-surface is invalid
	WlSubcompositorError_BadSurface WlSubcompositorError = 0
)

// WlSubcompositorDestroyRequest unbind from the subcompositor interface
/*
	Informs the server that the client will not be using this
	protocol object anymore. This does not affect any other
	objects, wl_subsurface objects included.
*/
type WlSubcompositorDestroyRequest struct{}

// WlSubcompositorGetSubsurfaceRequest give a surface the role sub-surface
/*
	Create a sub-surface interface for the given surface, and
	associate it with the given parent surface. This turns a
	plain wl_surface into a sub-surface.

	The to-be sub-surface must not already have another role, and it
	must not have an existing wl_subsurface object. Otherwise a protocol
	error is raised.

	Adding sub-surfaces to a parent is a double-buffered operation on the
	parent (see wl_surface.commit). The effect of adding a sub-surface
	becomes visible on the next time the state of the parent surface is
	applied.

	This request modifies the behaviour of wl_surface.commit request on
	the sub-surface, see the documentation on wl_subsurface interface.
*/
type WlSubcompositorGetSubsurfaceRequest struct {
	// Id the new sub-surface object ID
	Id uint32 `wayland:"new_id"`
	// Surface the surface to be turned into a sub-surface
	Surface uint32 `wayland:"object"`
	// Parent the parent surface
	Parent uint32 `wayland:"object"`
}
type WlSubcompositorListener interface{}
type UnimplementedWlSubcompositorListener struct{}
type WlSubcompositorDispatcher struct {
	WlSubcompositorListener
}

func NewWlSubcompositorDispatcher() *WlSubcompositorDispatcher {
	return &WlSubcompositorDispatcher{WlSubcompositorListener: &UnimplementedWlSubcompositorListener{}}
}
func (i *WlSubcompositorDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlSubcompositorDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlSubcompositorListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlSubcompositorListener = listener
}

// WlSubsurface sub-surface interface to a wl_surface
/*
	An additional interface to a wl_surface object, which has been
	made a sub-surface. A sub-surface has one parent surface. A
	sub-surface's size and position are not limited to that of the parent.
	Particularly, a sub-surface is not automatically clipped to its
	parent's area.

	A sub-surface becomes mapped, when a non-NULL wl_buffer is applied
	and the parent surface is mapped. The order of which one happens
	first is irrelevant. A sub-surface is hidden if the parent becomes
	hidden, or if a NULL wl_buffer is applied. These rules apply
	recursively through the tree of surfaces.

	The behaviour of a wl_surface.commit request on a sub-surface
	depends on the sub-surface's mode. The possible modes are
	synchronized and desynchronized, see methods
	wl_subsurface.set_sync and wl_subsurface.set_desync. Synchronized
	mode caches the wl_surface state to be applied when the parent's
	state gets applied, and desynchronized mode applies the pending
	wl_surface state directly. A sub-surface is initially in the
	synchronized mode.

	Sub-surfaces also have another kind of state, which is managed by
	wl_subsurface requests, as opposed to wl_surface requests. This
	state includes the sub-surface position relative to the parent
	surface (wl_subsurface.set_position), and the stacking order of
	the parent and its sub-surfaces (wl_subsurface.place_above and
	.place_below). This state is applied when the parent surface's
	wl_surface state is applied, regardless of the sub-surface's mode.
	As the exception, set_sync and set_desync are effective immediately.

	The main surface can be thought to be always in desynchronized mode,
	since it does not have a parent in the sub-surfaces sense.

	Even if a sub-surface is in desynchronized mode, it will behave as
	in synchronized mode, if its parent surface behaves as in
	synchronized mode. This rule is applied recursively throughout the
	tree of surfaces. This means, that one can set a sub-surface into
	synchronized mode, and then assume that all its child and grand-child
	sub-surfaces are synchronized, too, without explicitly setting them.

	If the wl_surface associated with the wl_subsurface is destroyed, the
	wl_subsurface object becomes inert. Note, that destroying either object
	takes effect immediately. If you need to synchronize the removal
	of a sub-surface to the parent surface update, unmap the sub-surface
	first by attaching a NULL wl_buffer, update parent, and then destroy
	the sub-surface.

	If the parent wl_surface object is destroyed, the sub-surface is
	unmapped.
*/
type WlSubsurface struct {
	Proxy
}

func NewWlSubsurface(p Proxy) WlSubsurface {
	return WlSubsurface{Proxy: p}
}
func (w *WlSubsurface) AttachListener(l WlSubsurfaceListener) {
	w.Proxy.Dispatcher().AttachListener(l)
}

type WlSubsurfaceError uint32

const (
	// WlSubsurfaceError_BadSurface wl_surface is not a sibling or the parent
	WlSubsurfaceError_BadSurface WlSubsurfaceError = 0
)

// WlSubsurfaceDestroyRequest remove sub-surface interface
/*
	The sub-surface interface is removed from the wl_surface object
	that was turned into a sub-surface with a
	wl_subcompositor.get_subsurface request. The wl_surface's association
	to the parent is deleted, and the wl_surface loses its role as
	a sub-surface. The wl_surface is unmapped immediately.
*/
type WlSubsurfaceDestroyRequest struct{}

// WlSubsurfaceSetPositionRequest reposition the sub-surface
/*
	This schedules a sub-surface position change.
	The sub-surface will be moved so that its origin (top left
	corner pixel) will be at the location x, y of the parent surface
	coordinate system. The coordinates are not restricted to the parent
	surface area. Negative values are allowed.

	The scheduled coordinates will take effect whenever the state of the
	parent surface is applied. When this happens depends on whether the
	parent surface is in synchronized mode or not. See
	wl_subsurface.set_sync and wl_subsurface.set_desync for details.

	If more than one set_position request is invoked by the client before
	the commit of the parent surface, the position of a new request always
	replaces the scheduled position from any previous request.

	The initial position is 0, 0.
*/
type WlSubsurfaceSetPositionRequest struct {
	// X  coordinate in the parent surface
	X int32 `wayland:"int"`
	// Y  coordinate in the parent surface
	Y int32 `wayland:"int"`
}

// WlSubsurfacePlaceAboveRequest restack the sub-surface
/*
	This sub-surface is taken from the stack, and put back just
	above the reference surface, changing the z-order of the sub-surfaces.
	The reference surface must be one of the sibling surfaces, or the
	parent surface. Using any other surface, including this sub-surface,
	will cause a protocol error.

	The z-order is double-buffered. Requests are handled in order and
	applied immediately to a pending state. The final pending state is
	copied to the active state the next time the state of the parent
	surface is applied. When this happens depends on whether the parent
	surface is in synchronized mode or not. See wl_subsurface.set_sync and
	wl_subsurface.set_desync for details.

	A new sub-surface is initially added as the top-most in the stack
	of its siblings and parent.
*/
type WlSubsurfacePlaceAboveRequest struct {
	// Sibling the reference surface
	Sibling uint32 `wayland:"object"`
}

// WlSubsurfacePlaceBelowRequest restack the sub-surface
/*
	The sub-surface is placed just below the reference surface.
	See wl_subsurface.place_above.
*/
type WlSubsurfacePlaceBelowRequest struct {
	// Sibling the reference surface
	Sibling uint32 `wayland:"object"`
}

// WlSubsurfaceSetSyncRequest set sub-surface to synchronized mode
/*
	Change the commit behaviour of the sub-surface to synchronized
	mode, also described as the parent dependent mode.

	In synchronized mode, wl_surface.commit on a sub-surface will
	accumulate the committed state in a cache, but the state will
	not be applied and hence will not change the compositor output.
	The cached state is applied to the sub-surface immediately after
	the parent surface's state is applied. This ensures atomic
	updates of the parent and all its synchronized sub-surfaces.
	Applying the cached state will invalidate the cache, so further
	parent surface commits do not (re-)apply old state.

	See wl_subsurface for the recursive effect of this mode.
*/
type WlSubsurfaceSetSyncRequest struct{}

// WlSubsurfaceSetDesyncRequest set sub-surface to desynchronized mode
/*
	Change the commit behaviour of the sub-surface to desynchronized
	mode, also described as independent or freely running mode.

	In desynchronized mode, wl_surface.commit on a sub-surface will
	apply the pending state directly, without caching, as happens
	normally with a wl_surface. Calling wl_surface.commit on the
	parent surface has no effect on the sub-surface's wl_surface
	state. This mode allows a sub-surface to be updated on its own.

	If cached state exists when wl_surface.commit is called in
	desynchronized mode, the pending state is added to the cached
	state, and applied as a whole. This invalidates the cache.

	Note: even if a sub-surface is set to desynchronized, a parent
	sub-surface may override it to behave as synchronized. For details,
	see wl_subsurface.

	If a surface's parent surface behaves as desynchronized, then
	the cached state is applied on set_desync.
*/
type WlSubsurfaceSetDesyncRequest struct{}
type WlSubsurfaceListener interface{}
type UnimplementedWlSubsurfaceListener struct{}
type WlSubsurfaceDispatcher struct {
	WlSubsurfaceListener
}

func NewWlSubsurfaceDispatcher() *WlSubsurfaceDispatcher {
	return &WlSubsurfaceDispatcher{WlSubsurfaceListener: &UnimplementedWlSubsurfaceListener{}}
}
func (i *WlSubsurfaceDispatcher) Dispatch(h MessageHeader, b []byte) error {
	switch h.Opcode {
	default:
		return ErrInvalidOp
	}
	return nil
}
func (i *WlSubsurfaceDispatcher) AttachListener(l interface{}) {
	listener, ok := l.(WlSubsurfaceListener)
	if !ok {
		log.Panic().Msg("listener is of wrong type!")
	}
	i.WlSubsurfaceListener = listener
}
