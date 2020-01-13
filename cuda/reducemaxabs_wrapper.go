package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/mumax/3/cuda/cu"
	"github.com/mumax/3/timer"
	"sync"
	"unsafe"
)

// CUDA handle for reducemaxabs kernel
var reducemaxabs_code cu.Function

// Stores the arguments for reducemaxabs kernel invocation
type reducemaxabs_args_t struct {
	arg_src     unsafe.Pointer
	arg_dst     unsafe.Pointer
	arg_initVal float32
	arg_n       int
	argptr      [4]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for reducemaxabs kernel invocation
var reducemaxabs_args reducemaxabs_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	reducemaxabs_args.argptr[0] = unsafe.Pointer(&reducemaxabs_args.arg_src)
	reducemaxabs_args.argptr[1] = unsafe.Pointer(&reducemaxabs_args.arg_dst)
	reducemaxabs_args.argptr[2] = unsafe.Pointer(&reducemaxabs_args.arg_initVal)
	reducemaxabs_args.argptr[3] = unsafe.Pointer(&reducemaxabs_args.arg_n)
}

// Wrapper for reducemaxabs CUDA kernel, asynchronous.
func k_reducemaxabs_async(src unsafe.Pointer, dst unsafe.Pointer, initVal float32, n int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("reducemaxabs")
	}

	reducemaxabs_args.Lock()
	defer reducemaxabs_args.Unlock()

	if reducemaxabs_code == 0 {
		reducemaxabs_code = fatbinLoad(reducemaxabs_map, "reducemaxabs")
	}

	reducemaxabs_args.arg_src = src
	reducemaxabs_args.arg_dst = dst
	reducemaxabs_args.arg_initVal = initVal
	reducemaxabs_args.arg_n = n

	args := reducemaxabs_args.argptr[:]
	cu.LaunchKernel(reducemaxabs_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("reducemaxabs")
	}
}

// maps compute capability on PTX code for reducemaxabs kernel.
var reducemaxabs_map = map[int]string{0: "",
	30: reducemaxabs_ptx_30,
	35: reducemaxabs_ptx_35,
	37: reducemaxabs_ptx_37,
	50: reducemaxabs_ptx_50,
	52: reducemaxabs_ptx_52,
	53: reducemaxabs_ptx_53,
	60: reducemaxabs_ptx_60,
	61: reducemaxabs_ptx_61,
	70: reducemaxabs_ptx_70,
	75: reducemaxabs_ptx_75}

// reducemaxabs PTX code for various compute capabilities.
const (
	reducemaxabs_ptx_30 = `
.version 6.5
.target sm_30
.address_size 64

	// .globl	reducemaxabs

.visible .entry reducemaxabs(
	.param .u64 reducemaxabs_param_0,
	.param .u64 reducemaxabs_param_1,
	.param .f32 reducemaxabs_param_2,
	.param .u32 reducemaxabs_param_3
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<32>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<7>;
	// demoted variable
	.shared .align 4 .b8 _ZZ12reducemaxabsE5sdata[2048];

	ld.param.u64 	%rd3, [reducemaxabs_param_0];
	ld.param.u64 	%rd2, [reducemaxabs_param_1];
	ld.param.f32 	%f31, [reducemaxabs_param_2];
	ld.param.u32 	%r10, [reducemaxabs_param_3];
	cvta.to.global.u64 	%rd1, %rd3;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd4, %r21, 4;
	add.s64 	%rd5, %rd1, %rd4;
	ld.global.f32 	%f5, [%rd5];
	abs.f32 	%f6, %f5;
	max.f32 	%f31, %f31, %f6;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ12reducemaxabsE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f31;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f7, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f8, [%r18];
	max.f32 	%f9, %f7, %f8;
	st.shared.f32 	[%r7], %f9;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f10, [%r7];
	ld.volatile.shared.f32 	%f11, [%r7+128];
	max.f32 	%f12, %f10, %f11;
	st.volatile.shared.f32 	[%r7], %f12;
	ld.volatile.shared.f32 	%f13, [%r7+64];
	ld.volatile.shared.f32 	%f14, [%r7];
	max.f32 	%f15, %f14, %f13;
	st.volatile.shared.f32 	[%r7], %f15;
	ld.volatile.shared.f32 	%f16, [%r7+32];
	ld.volatile.shared.f32 	%f17, [%r7];
	max.f32 	%f18, %f17, %f16;
	st.volatile.shared.f32 	[%r7], %f18;
	ld.volatile.shared.f32 	%f19, [%r7+16];
	ld.volatile.shared.f32 	%f20, [%r7];
	max.f32 	%f21, %f20, %f19;
	st.volatile.shared.f32 	[%r7], %f21;
	ld.volatile.shared.f32 	%f22, [%r7+8];
	ld.volatile.shared.f32 	%f23, [%r7];
	max.f32 	%f24, %f23, %f22;
	st.volatile.shared.f32 	[%r7], %f24;
	ld.volatile.shared.f32 	%f25, [%r7+4];
	ld.volatile.shared.f32 	%f26, [%r7];
	max.f32 	%f27, %f26, %f25;
	st.volatile.shared.f32 	[%r7], %f27;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f28, [_ZZ12reducemaxabsE5sdata];
	abs.f32 	%f29, %f28;
	mov.b32 	 %r19, %f29;
	cvta.to.global.u64 	%rd6, %rd2;
	atom.global.max.s32 	%r20, [%rd6], %r19;

BB0_10:
	ret;
}


`
	reducemaxabs_ptx_35 = `
.version 6.5
.target sm_35
.address_size 64

	// .globl	reducemaxabs

.visible .entry reducemaxabs(
	.param .u64 reducemaxabs_param_0,
	.param .u64 reducemaxabs_param_1,
	.param .f32 reducemaxabs_param_2,
	.param .u32 reducemaxabs_param_3
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<32>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<7>;
	// demoted variable
	.shared .align 4 .b8 _ZZ12reducemaxabsE5sdata[2048];

	ld.param.u64 	%rd3, [reducemaxabs_param_0];
	ld.param.u64 	%rd2, [reducemaxabs_param_1];
	ld.param.f32 	%f31, [reducemaxabs_param_2];
	ld.param.u32 	%r10, [reducemaxabs_param_3];
	cvta.to.global.u64 	%rd1, %rd3;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd4, %r21, 4;
	add.s64 	%rd5, %rd1, %rd4;
	ld.global.nc.f32 	%f5, [%rd5];
	abs.f32 	%f6, %f5;
	max.f32 	%f31, %f31, %f6;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ12reducemaxabsE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f31;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f7, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f8, [%r18];
	max.f32 	%f9, %f7, %f8;
	st.shared.f32 	[%r7], %f9;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f10, [%r7];
	ld.volatile.shared.f32 	%f11, [%r7+128];
	max.f32 	%f12, %f10, %f11;
	st.volatile.shared.f32 	[%r7], %f12;
	ld.volatile.shared.f32 	%f13, [%r7+64];
	ld.volatile.shared.f32 	%f14, [%r7];
	max.f32 	%f15, %f14, %f13;
	st.volatile.shared.f32 	[%r7], %f15;
	ld.volatile.shared.f32 	%f16, [%r7+32];
	ld.volatile.shared.f32 	%f17, [%r7];
	max.f32 	%f18, %f17, %f16;
	st.volatile.shared.f32 	[%r7], %f18;
	ld.volatile.shared.f32 	%f19, [%r7+16];
	ld.volatile.shared.f32 	%f20, [%r7];
	max.f32 	%f21, %f20, %f19;
	st.volatile.shared.f32 	[%r7], %f21;
	ld.volatile.shared.f32 	%f22, [%r7+8];
	ld.volatile.shared.f32 	%f23, [%r7];
	max.f32 	%f24, %f23, %f22;
	st.volatile.shared.f32 	[%r7], %f24;
	ld.volatile.shared.f32 	%f25, [%r7+4];
	ld.volatile.shared.f32 	%f26, [%r7];
	max.f32 	%f27, %f26, %f25;
	st.volatile.shared.f32 	[%r7], %f27;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f28, [_ZZ12reducemaxabsE5sdata];
	abs.f32 	%f29, %f28;
	mov.b32 	 %r19, %f29;
	cvta.to.global.u64 	%rd6, %rd2;
	atom.global.max.s32 	%r20, [%rd6], %r19;

BB0_10:
	ret;
}


`
	reducemaxabs_ptx_37 = `
.version 6.5
.target sm_37
.address_size 64

	// .globl	reducemaxabs

.visible .entry reducemaxabs(
	.param .u64 reducemaxabs_param_0,
	.param .u64 reducemaxabs_param_1,
	.param .f32 reducemaxabs_param_2,
	.param .u32 reducemaxabs_param_3
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<32>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<7>;
	// demoted variable
	.shared .align 4 .b8 _ZZ12reducemaxabsE5sdata[2048];

	ld.param.u64 	%rd3, [reducemaxabs_param_0];
	ld.param.u64 	%rd2, [reducemaxabs_param_1];
	ld.param.f32 	%f31, [reducemaxabs_param_2];
	ld.param.u32 	%r10, [reducemaxabs_param_3];
	cvta.to.global.u64 	%rd1, %rd3;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd4, %r21, 4;
	add.s64 	%rd5, %rd1, %rd4;
	ld.global.nc.f32 	%f5, [%rd5];
	abs.f32 	%f6, %f5;
	max.f32 	%f31, %f31, %f6;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ12reducemaxabsE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f31;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f7, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f8, [%r18];
	max.f32 	%f9, %f7, %f8;
	st.shared.f32 	[%r7], %f9;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f10, [%r7];
	ld.volatile.shared.f32 	%f11, [%r7+128];
	max.f32 	%f12, %f10, %f11;
	st.volatile.shared.f32 	[%r7], %f12;
	ld.volatile.shared.f32 	%f13, [%r7+64];
	ld.volatile.shared.f32 	%f14, [%r7];
	max.f32 	%f15, %f14, %f13;
	st.volatile.shared.f32 	[%r7], %f15;
	ld.volatile.shared.f32 	%f16, [%r7+32];
	ld.volatile.shared.f32 	%f17, [%r7];
	max.f32 	%f18, %f17, %f16;
	st.volatile.shared.f32 	[%r7], %f18;
	ld.volatile.shared.f32 	%f19, [%r7+16];
	ld.volatile.shared.f32 	%f20, [%r7];
	max.f32 	%f21, %f20, %f19;
	st.volatile.shared.f32 	[%r7], %f21;
	ld.volatile.shared.f32 	%f22, [%r7+8];
	ld.volatile.shared.f32 	%f23, [%r7];
	max.f32 	%f24, %f23, %f22;
	st.volatile.shared.f32 	[%r7], %f24;
	ld.volatile.shared.f32 	%f25, [%r7+4];
	ld.volatile.shared.f32 	%f26, [%r7];
	max.f32 	%f27, %f26, %f25;
	st.volatile.shared.f32 	[%r7], %f27;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f28, [_ZZ12reducemaxabsE5sdata];
	abs.f32 	%f29, %f28;
	mov.b32 	 %r19, %f29;
	cvta.to.global.u64 	%rd6, %rd2;
	atom.global.max.s32 	%r20, [%rd6], %r19;

BB0_10:
	ret;
}


`
	reducemaxabs_ptx_50 = `
.version 6.5
.target sm_50
.address_size 64

	// .globl	reducemaxabs

.visible .entry reducemaxabs(
	.param .u64 reducemaxabs_param_0,
	.param .u64 reducemaxabs_param_1,
	.param .f32 reducemaxabs_param_2,
	.param .u32 reducemaxabs_param_3
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<32>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<7>;
	// demoted variable
	.shared .align 4 .b8 _ZZ12reducemaxabsE5sdata[2048];

	ld.param.u64 	%rd3, [reducemaxabs_param_0];
	ld.param.u64 	%rd2, [reducemaxabs_param_1];
	ld.param.f32 	%f31, [reducemaxabs_param_2];
	ld.param.u32 	%r10, [reducemaxabs_param_3];
	cvta.to.global.u64 	%rd1, %rd3;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd4, %r21, 4;
	add.s64 	%rd5, %rd1, %rd4;
	ld.global.nc.f32 	%f5, [%rd5];
	abs.f32 	%f6, %f5;
	max.f32 	%f31, %f31, %f6;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ12reducemaxabsE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f31;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f7, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f8, [%r18];
	max.f32 	%f9, %f7, %f8;
	st.shared.f32 	[%r7], %f9;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f10, [%r7];
	ld.volatile.shared.f32 	%f11, [%r7+128];
	max.f32 	%f12, %f10, %f11;
	st.volatile.shared.f32 	[%r7], %f12;
	ld.volatile.shared.f32 	%f13, [%r7+64];
	ld.volatile.shared.f32 	%f14, [%r7];
	max.f32 	%f15, %f14, %f13;
	st.volatile.shared.f32 	[%r7], %f15;
	ld.volatile.shared.f32 	%f16, [%r7+32];
	ld.volatile.shared.f32 	%f17, [%r7];
	max.f32 	%f18, %f17, %f16;
	st.volatile.shared.f32 	[%r7], %f18;
	ld.volatile.shared.f32 	%f19, [%r7+16];
	ld.volatile.shared.f32 	%f20, [%r7];
	max.f32 	%f21, %f20, %f19;
	st.volatile.shared.f32 	[%r7], %f21;
	ld.volatile.shared.f32 	%f22, [%r7+8];
	ld.volatile.shared.f32 	%f23, [%r7];
	max.f32 	%f24, %f23, %f22;
	st.volatile.shared.f32 	[%r7], %f24;
	ld.volatile.shared.f32 	%f25, [%r7+4];
	ld.volatile.shared.f32 	%f26, [%r7];
	max.f32 	%f27, %f26, %f25;
	st.volatile.shared.f32 	[%r7], %f27;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f28, [_ZZ12reducemaxabsE5sdata];
	abs.f32 	%f29, %f28;
	mov.b32 	 %r19, %f29;
	cvta.to.global.u64 	%rd6, %rd2;
	atom.global.max.s32 	%r20, [%rd6], %r19;

BB0_10:
	ret;
}


`
	reducemaxabs_ptx_52 = `
.version 6.5
.target sm_52
.address_size 64

	// .globl	reducemaxabs

.visible .entry reducemaxabs(
	.param .u64 reducemaxabs_param_0,
	.param .u64 reducemaxabs_param_1,
	.param .f32 reducemaxabs_param_2,
	.param .u32 reducemaxabs_param_3
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<32>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<7>;
	// demoted variable
	.shared .align 4 .b8 _ZZ12reducemaxabsE5sdata[2048];

	ld.param.u64 	%rd3, [reducemaxabs_param_0];
	ld.param.u64 	%rd2, [reducemaxabs_param_1];
	ld.param.f32 	%f31, [reducemaxabs_param_2];
	ld.param.u32 	%r10, [reducemaxabs_param_3];
	cvta.to.global.u64 	%rd1, %rd3;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd4, %r21, 4;
	add.s64 	%rd5, %rd1, %rd4;
	ld.global.nc.f32 	%f5, [%rd5];
	abs.f32 	%f6, %f5;
	max.f32 	%f31, %f31, %f6;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ12reducemaxabsE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f31;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f7, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f8, [%r18];
	max.f32 	%f9, %f7, %f8;
	st.shared.f32 	[%r7], %f9;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f10, [%r7];
	ld.volatile.shared.f32 	%f11, [%r7+128];
	max.f32 	%f12, %f10, %f11;
	st.volatile.shared.f32 	[%r7], %f12;
	ld.volatile.shared.f32 	%f13, [%r7+64];
	ld.volatile.shared.f32 	%f14, [%r7];
	max.f32 	%f15, %f14, %f13;
	st.volatile.shared.f32 	[%r7], %f15;
	ld.volatile.shared.f32 	%f16, [%r7+32];
	ld.volatile.shared.f32 	%f17, [%r7];
	max.f32 	%f18, %f17, %f16;
	st.volatile.shared.f32 	[%r7], %f18;
	ld.volatile.shared.f32 	%f19, [%r7+16];
	ld.volatile.shared.f32 	%f20, [%r7];
	max.f32 	%f21, %f20, %f19;
	st.volatile.shared.f32 	[%r7], %f21;
	ld.volatile.shared.f32 	%f22, [%r7+8];
	ld.volatile.shared.f32 	%f23, [%r7];
	max.f32 	%f24, %f23, %f22;
	st.volatile.shared.f32 	[%r7], %f24;
	ld.volatile.shared.f32 	%f25, [%r7+4];
	ld.volatile.shared.f32 	%f26, [%r7];
	max.f32 	%f27, %f26, %f25;
	st.volatile.shared.f32 	[%r7], %f27;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f28, [_ZZ12reducemaxabsE5sdata];
	abs.f32 	%f29, %f28;
	mov.b32 	 %r19, %f29;
	cvta.to.global.u64 	%rd6, %rd2;
	atom.global.max.s32 	%r20, [%rd6], %r19;

BB0_10:
	ret;
}


`
	reducemaxabs_ptx_53 = `
.version 6.5
.target sm_53
.address_size 64

	// .globl	reducemaxabs

.visible .entry reducemaxabs(
	.param .u64 reducemaxabs_param_0,
	.param .u64 reducemaxabs_param_1,
	.param .f32 reducemaxabs_param_2,
	.param .u32 reducemaxabs_param_3
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<32>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<7>;
	// demoted variable
	.shared .align 4 .b8 _ZZ12reducemaxabsE5sdata[2048];

	ld.param.u64 	%rd3, [reducemaxabs_param_0];
	ld.param.u64 	%rd2, [reducemaxabs_param_1];
	ld.param.f32 	%f31, [reducemaxabs_param_2];
	ld.param.u32 	%r10, [reducemaxabs_param_3];
	cvta.to.global.u64 	%rd1, %rd3;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd4, %r21, 4;
	add.s64 	%rd5, %rd1, %rd4;
	ld.global.nc.f32 	%f5, [%rd5];
	abs.f32 	%f6, %f5;
	max.f32 	%f31, %f31, %f6;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ12reducemaxabsE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f31;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f7, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f8, [%r18];
	max.f32 	%f9, %f7, %f8;
	st.shared.f32 	[%r7], %f9;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f10, [%r7];
	ld.volatile.shared.f32 	%f11, [%r7+128];
	max.f32 	%f12, %f10, %f11;
	st.volatile.shared.f32 	[%r7], %f12;
	ld.volatile.shared.f32 	%f13, [%r7+64];
	ld.volatile.shared.f32 	%f14, [%r7];
	max.f32 	%f15, %f14, %f13;
	st.volatile.shared.f32 	[%r7], %f15;
	ld.volatile.shared.f32 	%f16, [%r7+32];
	ld.volatile.shared.f32 	%f17, [%r7];
	max.f32 	%f18, %f17, %f16;
	st.volatile.shared.f32 	[%r7], %f18;
	ld.volatile.shared.f32 	%f19, [%r7+16];
	ld.volatile.shared.f32 	%f20, [%r7];
	max.f32 	%f21, %f20, %f19;
	st.volatile.shared.f32 	[%r7], %f21;
	ld.volatile.shared.f32 	%f22, [%r7+8];
	ld.volatile.shared.f32 	%f23, [%r7];
	max.f32 	%f24, %f23, %f22;
	st.volatile.shared.f32 	[%r7], %f24;
	ld.volatile.shared.f32 	%f25, [%r7+4];
	ld.volatile.shared.f32 	%f26, [%r7];
	max.f32 	%f27, %f26, %f25;
	st.volatile.shared.f32 	[%r7], %f27;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f28, [_ZZ12reducemaxabsE5sdata];
	abs.f32 	%f29, %f28;
	mov.b32 	 %r19, %f29;
	cvta.to.global.u64 	%rd6, %rd2;
	atom.global.max.s32 	%r20, [%rd6], %r19;

BB0_10:
	ret;
}


`
	reducemaxabs_ptx_60 = `
.version 6.5
.target sm_60
.address_size 64

	// .globl	reducemaxabs

.visible .entry reducemaxabs(
	.param .u64 reducemaxabs_param_0,
	.param .u64 reducemaxabs_param_1,
	.param .f32 reducemaxabs_param_2,
	.param .u32 reducemaxabs_param_3
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<32>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<7>;
	// demoted variable
	.shared .align 4 .b8 _ZZ12reducemaxabsE5sdata[2048];

	ld.param.u64 	%rd3, [reducemaxabs_param_0];
	ld.param.u64 	%rd2, [reducemaxabs_param_1];
	ld.param.f32 	%f31, [reducemaxabs_param_2];
	ld.param.u32 	%r10, [reducemaxabs_param_3];
	cvta.to.global.u64 	%rd1, %rd3;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd4, %r21, 4;
	add.s64 	%rd5, %rd1, %rd4;
	ld.global.nc.f32 	%f5, [%rd5];
	abs.f32 	%f6, %f5;
	max.f32 	%f31, %f31, %f6;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ12reducemaxabsE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f31;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f7, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f8, [%r18];
	max.f32 	%f9, %f7, %f8;
	st.shared.f32 	[%r7], %f9;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f10, [%r7];
	ld.volatile.shared.f32 	%f11, [%r7+128];
	max.f32 	%f12, %f10, %f11;
	st.volatile.shared.f32 	[%r7], %f12;
	ld.volatile.shared.f32 	%f13, [%r7+64];
	ld.volatile.shared.f32 	%f14, [%r7];
	max.f32 	%f15, %f14, %f13;
	st.volatile.shared.f32 	[%r7], %f15;
	ld.volatile.shared.f32 	%f16, [%r7+32];
	ld.volatile.shared.f32 	%f17, [%r7];
	max.f32 	%f18, %f17, %f16;
	st.volatile.shared.f32 	[%r7], %f18;
	ld.volatile.shared.f32 	%f19, [%r7+16];
	ld.volatile.shared.f32 	%f20, [%r7];
	max.f32 	%f21, %f20, %f19;
	st.volatile.shared.f32 	[%r7], %f21;
	ld.volatile.shared.f32 	%f22, [%r7+8];
	ld.volatile.shared.f32 	%f23, [%r7];
	max.f32 	%f24, %f23, %f22;
	st.volatile.shared.f32 	[%r7], %f24;
	ld.volatile.shared.f32 	%f25, [%r7+4];
	ld.volatile.shared.f32 	%f26, [%r7];
	max.f32 	%f27, %f26, %f25;
	st.volatile.shared.f32 	[%r7], %f27;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f28, [_ZZ12reducemaxabsE5sdata];
	abs.f32 	%f29, %f28;
	mov.b32 	 %r19, %f29;
	cvta.to.global.u64 	%rd6, %rd2;
	atom.global.max.s32 	%r20, [%rd6], %r19;

BB0_10:
	ret;
}


`
	reducemaxabs_ptx_61 = `
.version 6.5
.target sm_61
.address_size 64

	// .globl	reducemaxabs

.visible .entry reducemaxabs(
	.param .u64 reducemaxabs_param_0,
	.param .u64 reducemaxabs_param_1,
	.param .f32 reducemaxabs_param_2,
	.param .u32 reducemaxabs_param_3
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<32>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<7>;
	// demoted variable
	.shared .align 4 .b8 _ZZ12reducemaxabsE5sdata[2048];

	ld.param.u64 	%rd3, [reducemaxabs_param_0];
	ld.param.u64 	%rd2, [reducemaxabs_param_1];
	ld.param.f32 	%f31, [reducemaxabs_param_2];
	ld.param.u32 	%r10, [reducemaxabs_param_3];
	cvta.to.global.u64 	%rd1, %rd3;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd4, %r21, 4;
	add.s64 	%rd5, %rd1, %rd4;
	ld.global.nc.f32 	%f5, [%rd5];
	abs.f32 	%f6, %f5;
	max.f32 	%f31, %f31, %f6;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ12reducemaxabsE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f31;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f7, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f8, [%r18];
	max.f32 	%f9, %f7, %f8;
	st.shared.f32 	[%r7], %f9;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f10, [%r7];
	ld.volatile.shared.f32 	%f11, [%r7+128];
	max.f32 	%f12, %f10, %f11;
	st.volatile.shared.f32 	[%r7], %f12;
	ld.volatile.shared.f32 	%f13, [%r7+64];
	ld.volatile.shared.f32 	%f14, [%r7];
	max.f32 	%f15, %f14, %f13;
	st.volatile.shared.f32 	[%r7], %f15;
	ld.volatile.shared.f32 	%f16, [%r7+32];
	ld.volatile.shared.f32 	%f17, [%r7];
	max.f32 	%f18, %f17, %f16;
	st.volatile.shared.f32 	[%r7], %f18;
	ld.volatile.shared.f32 	%f19, [%r7+16];
	ld.volatile.shared.f32 	%f20, [%r7];
	max.f32 	%f21, %f20, %f19;
	st.volatile.shared.f32 	[%r7], %f21;
	ld.volatile.shared.f32 	%f22, [%r7+8];
	ld.volatile.shared.f32 	%f23, [%r7];
	max.f32 	%f24, %f23, %f22;
	st.volatile.shared.f32 	[%r7], %f24;
	ld.volatile.shared.f32 	%f25, [%r7+4];
	ld.volatile.shared.f32 	%f26, [%r7];
	max.f32 	%f27, %f26, %f25;
	st.volatile.shared.f32 	[%r7], %f27;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f28, [_ZZ12reducemaxabsE5sdata];
	abs.f32 	%f29, %f28;
	mov.b32 	 %r19, %f29;
	cvta.to.global.u64 	%rd6, %rd2;
	atom.global.max.s32 	%r20, [%rd6], %r19;

BB0_10:
	ret;
}


`
	reducemaxabs_ptx_70 = `
.version 6.5
.target sm_70
.address_size 64

	// .globl	reducemaxabs

.visible .entry reducemaxabs(
	.param .u64 reducemaxabs_param_0,
	.param .u64 reducemaxabs_param_1,
	.param .f32 reducemaxabs_param_2,
	.param .u32 reducemaxabs_param_3
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<32>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<7>;
	// demoted variable
	.shared .align 4 .b8 _ZZ12reducemaxabsE5sdata[2048];

	ld.param.u64 	%rd3, [reducemaxabs_param_0];
	ld.param.u64 	%rd2, [reducemaxabs_param_1];
	ld.param.f32 	%f31, [reducemaxabs_param_2];
	ld.param.u32 	%r10, [reducemaxabs_param_3];
	cvta.to.global.u64 	%rd1, %rd3;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd4, %r21, 4;
	add.s64 	%rd5, %rd1, %rd4;
	ld.global.nc.f32 	%f5, [%rd5];
	abs.f32 	%f6, %f5;
	max.f32 	%f31, %f31, %f6;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ12reducemaxabsE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f31;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f7, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f8, [%r18];
	max.f32 	%f9, %f7, %f8;
	st.shared.f32 	[%r7], %f9;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f10, [%r7];
	ld.volatile.shared.f32 	%f11, [%r7+128];
	max.f32 	%f12, %f10, %f11;
	st.volatile.shared.f32 	[%r7], %f12;
	ld.volatile.shared.f32 	%f13, [%r7+64];
	ld.volatile.shared.f32 	%f14, [%r7];
	max.f32 	%f15, %f14, %f13;
	st.volatile.shared.f32 	[%r7], %f15;
	ld.volatile.shared.f32 	%f16, [%r7+32];
	ld.volatile.shared.f32 	%f17, [%r7];
	max.f32 	%f18, %f17, %f16;
	st.volatile.shared.f32 	[%r7], %f18;
	ld.volatile.shared.f32 	%f19, [%r7+16];
	ld.volatile.shared.f32 	%f20, [%r7];
	max.f32 	%f21, %f20, %f19;
	st.volatile.shared.f32 	[%r7], %f21;
	ld.volatile.shared.f32 	%f22, [%r7+8];
	ld.volatile.shared.f32 	%f23, [%r7];
	max.f32 	%f24, %f23, %f22;
	st.volatile.shared.f32 	[%r7], %f24;
	ld.volatile.shared.f32 	%f25, [%r7+4];
	ld.volatile.shared.f32 	%f26, [%r7];
	max.f32 	%f27, %f26, %f25;
	st.volatile.shared.f32 	[%r7], %f27;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f28, [_ZZ12reducemaxabsE5sdata];
	abs.f32 	%f29, %f28;
	mov.b32 	 %r19, %f29;
	cvta.to.global.u64 	%rd6, %rd2;
	atom.global.max.s32 	%r20, [%rd6], %r19;

BB0_10:
	ret;
}


`
	reducemaxabs_ptx_75 = `
.version 6.5
.target sm_75
.address_size 64

	// .globl	reducemaxabs

.visible .entry reducemaxabs(
	.param .u64 reducemaxabs_param_0,
	.param .u64 reducemaxabs_param_1,
	.param .f32 reducemaxabs_param_2,
	.param .u32 reducemaxabs_param_3
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<32>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<7>;
	// demoted variable
	.shared .align 4 .b8 _ZZ12reducemaxabsE5sdata[2048];

	ld.param.u64 	%rd3, [reducemaxabs_param_0];
	ld.param.u64 	%rd2, [reducemaxabs_param_1];
	ld.param.f32 	%f31, [reducemaxabs_param_2];
	ld.param.u32 	%r10, [reducemaxabs_param_3];
	cvta.to.global.u64 	%rd1, %rd3;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd4, %r21, 4;
	add.s64 	%rd5, %rd1, %rd4;
	ld.global.nc.f32 	%f5, [%rd5];
	abs.f32 	%f6, %f5;
	max.f32 	%f31, %f31, %f6;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ12reducemaxabsE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f31;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f7, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f8, [%r18];
	max.f32 	%f9, %f7, %f8;
	st.shared.f32 	[%r7], %f9;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f10, [%r7];
	ld.volatile.shared.f32 	%f11, [%r7+128];
	max.f32 	%f12, %f10, %f11;
	st.volatile.shared.f32 	[%r7], %f12;
	ld.volatile.shared.f32 	%f13, [%r7+64];
	ld.volatile.shared.f32 	%f14, [%r7];
	max.f32 	%f15, %f14, %f13;
	st.volatile.shared.f32 	[%r7], %f15;
	ld.volatile.shared.f32 	%f16, [%r7+32];
	ld.volatile.shared.f32 	%f17, [%r7];
	max.f32 	%f18, %f17, %f16;
	st.volatile.shared.f32 	[%r7], %f18;
	ld.volatile.shared.f32 	%f19, [%r7+16];
	ld.volatile.shared.f32 	%f20, [%r7];
	max.f32 	%f21, %f20, %f19;
	st.volatile.shared.f32 	[%r7], %f21;
	ld.volatile.shared.f32 	%f22, [%r7+8];
	ld.volatile.shared.f32 	%f23, [%r7];
	max.f32 	%f24, %f23, %f22;
	st.volatile.shared.f32 	[%r7], %f24;
	ld.volatile.shared.f32 	%f25, [%r7+4];
	ld.volatile.shared.f32 	%f26, [%r7];
	max.f32 	%f27, %f26, %f25;
	st.volatile.shared.f32 	[%r7], %f27;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f28, [_ZZ12reducemaxabsE5sdata];
	abs.f32 	%f29, %f28;
	mov.b32 	 %r19, %f29;
	cvta.to.global.u64 	%rd6, %rd2;
	atom.global.max.s32 	%r20, [%rd6], %r19;

BB0_10:
	ret;
}


`
)
