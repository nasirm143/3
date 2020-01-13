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

// CUDA handle for reducemaxdiff kernel
var reducemaxdiff_code cu.Function

// Stores the arguments for reducemaxdiff kernel invocation
type reducemaxdiff_args_t struct {
	arg_src1    unsafe.Pointer
	arg_src2    unsafe.Pointer
	arg_dst     unsafe.Pointer
	arg_initVal float32
	arg_n       int
	argptr      [5]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for reducemaxdiff kernel invocation
var reducemaxdiff_args reducemaxdiff_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	reducemaxdiff_args.argptr[0] = unsafe.Pointer(&reducemaxdiff_args.arg_src1)
	reducemaxdiff_args.argptr[1] = unsafe.Pointer(&reducemaxdiff_args.arg_src2)
	reducemaxdiff_args.argptr[2] = unsafe.Pointer(&reducemaxdiff_args.arg_dst)
	reducemaxdiff_args.argptr[3] = unsafe.Pointer(&reducemaxdiff_args.arg_initVal)
	reducemaxdiff_args.argptr[4] = unsafe.Pointer(&reducemaxdiff_args.arg_n)
}

// Wrapper for reducemaxdiff CUDA kernel, asynchronous.
func k_reducemaxdiff_async(src1 unsafe.Pointer, src2 unsafe.Pointer, dst unsafe.Pointer, initVal float32, n int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("reducemaxdiff")
	}

	reducemaxdiff_args.Lock()
	defer reducemaxdiff_args.Unlock()

	if reducemaxdiff_code == 0 {
		reducemaxdiff_code = fatbinLoad(reducemaxdiff_map, "reducemaxdiff")
	}

	reducemaxdiff_args.arg_src1 = src1
	reducemaxdiff_args.arg_src2 = src2
	reducemaxdiff_args.arg_dst = dst
	reducemaxdiff_args.arg_initVal = initVal
	reducemaxdiff_args.arg_n = n

	args := reducemaxdiff_args.argptr[:]
	cu.LaunchKernel(reducemaxdiff_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("reducemaxdiff")
	}
}

// maps compute capability on PTX code for reducemaxdiff kernel.
var reducemaxdiff_map = map[int]string{0: "",
	30: reducemaxdiff_ptx_30,
	35: reducemaxdiff_ptx_35,
	37: reducemaxdiff_ptx_37,
	50: reducemaxdiff_ptx_50,
	52: reducemaxdiff_ptx_52,
	53: reducemaxdiff_ptx_53,
	60: reducemaxdiff_ptx_60,
	61: reducemaxdiff_ptx_61,
	70: reducemaxdiff_ptx_70,
	75: reducemaxdiff_ptx_75}

// reducemaxdiff PTX code for various compute capabilities.
const (
	reducemaxdiff_ptx_30 = `
.version 6.5
.target sm_30
.address_size 64

	// .globl	reducemaxdiff

.visible .entry reducemaxdiff(
	.param .u64 reducemaxdiff_param_0,
	.param .u64 reducemaxdiff_param_1,
	.param .u64 reducemaxdiff_param_2,
	.param .f32 reducemaxdiff_param_3,
	.param .u32 reducemaxdiff_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<34>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<10>;
	// demoted variable
	.shared .align 4 .b8 _ZZ13reducemaxdiffE5sdata[2048];

	ld.param.u64 	%rd4, [reducemaxdiff_param_0];
	ld.param.u64 	%rd5, [reducemaxdiff_param_1];
	ld.param.u64 	%rd3, [reducemaxdiff_param_2];
	ld.param.f32 	%f33, [reducemaxdiff_param_3];
	ld.param.u32 	%r10, [reducemaxdiff_param_4];
	cvta.to.global.u64 	%rd1, %rd5;
	cvta.to.global.u64 	%rd2, %rd4;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd6, %r21, 4;
	add.s64 	%rd7, %rd2, %rd6;
	add.s64 	%rd8, %rd1, %rd6;
	ld.global.f32 	%f5, [%rd8];
	ld.global.f32 	%f6, [%rd7];
	sub.f32 	%f7, %f6, %f5;
	abs.f32 	%f8, %f7;
	max.f32 	%f33, %f33, %f8;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ13reducemaxdiffE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f33;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f9, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f10, [%r18];
	max.f32 	%f11, %f9, %f10;
	st.shared.f32 	[%r7], %f11;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f12, [%r7];
	ld.volatile.shared.f32 	%f13, [%r7+128];
	max.f32 	%f14, %f12, %f13;
	st.volatile.shared.f32 	[%r7], %f14;
	ld.volatile.shared.f32 	%f15, [%r7+64];
	ld.volatile.shared.f32 	%f16, [%r7];
	max.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%r7], %f17;
	ld.volatile.shared.f32 	%f18, [%r7+32];
	ld.volatile.shared.f32 	%f19, [%r7];
	max.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%r7], %f20;
	ld.volatile.shared.f32 	%f21, [%r7+16];
	ld.volatile.shared.f32 	%f22, [%r7];
	max.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%r7], %f23;
	ld.volatile.shared.f32 	%f24, [%r7+8];
	ld.volatile.shared.f32 	%f25, [%r7];
	max.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%r7], %f26;
	ld.volatile.shared.f32 	%f27, [%r7+4];
	ld.volatile.shared.f32 	%f28, [%r7];
	max.f32 	%f29, %f28, %f27;
	st.volatile.shared.f32 	[%r7], %f29;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f30, [_ZZ13reducemaxdiffE5sdata];
	abs.f32 	%f31, %f30;
	mov.b32 	 %r19, %f31;
	cvta.to.global.u64 	%rd9, %rd3;
	atom.global.max.s32 	%r20, [%rd9], %r19;

BB0_10:
	ret;
}


`
	reducemaxdiff_ptx_35 = `
.version 6.5
.target sm_35
.address_size 64

	// .globl	reducemaxdiff

.visible .entry reducemaxdiff(
	.param .u64 reducemaxdiff_param_0,
	.param .u64 reducemaxdiff_param_1,
	.param .u64 reducemaxdiff_param_2,
	.param .f32 reducemaxdiff_param_3,
	.param .u32 reducemaxdiff_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<34>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<10>;
	// demoted variable
	.shared .align 4 .b8 _ZZ13reducemaxdiffE5sdata[2048];

	ld.param.u64 	%rd4, [reducemaxdiff_param_0];
	ld.param.u64 	%rd5, [reducemaxdiff_param_1];
	ld.param.u64 	%rd3, [reducemaxdiff_param_2];
	ld.param.f32 	%f33, [reducemaxdiff_param_3];
	ld.param.u32 	%r10, [reducemaxdiff_param_4];
	cvta.to.global.u64 	%rd1, %rd5;
	cvta.to.global.u64 	%rd2, %rd4;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd6, %r21, 4;
	add.s64 	%rd7, %rd2, %rd6;
	add.s64 	%rd8, %rd1, %rd6;
	ld.global.nc.f32 	%f5, [%rd8];
	ld.global.nc.f32 	%f6, [%rd7];
	sub.f32 	%f7, %f6, %f5;
	abs.f32 	%f8, %f7;
	max.f32 	%f33, %f33, %f8;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ13reducemaxdiffE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f33;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f9, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f10, [%r18];
	max.f32 	%f11, %f9, %f10;
	st.shared.f32 	[%r7], %f11;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f12, [%r7];
	ld.volatile.shared.f32 	%f13, [%r7+128];
	max.f32 	%f14, %f12, %f13;
	st.volatile.shared.f32 	[%r7], %f14;
	ld.volatile.shared.f32 	%f15, [%r7+64];
	ld.volatile.shared.f32 	%f16, [%r7];
	max.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%r7], %f17;
	ld.volatile.shared.f32 	%f18, [%r7+32];
	ld.volatile.shared.f32 	%f19, [%r7];
	max.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%r7], %f20;
	ld.volatile.shared.f32 	%f21, [%r7+16];
	ld.volatile.shared.f32 	%f22, [%r7];
	max.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%r7], %f23;
	ld.volatile.shared.f32 	%f24, [%r7+8];
	ld.volatile.shared.f32 	%f25, [%r7];
	max.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%r7], %f26;
	ld.volatile.shared.f32 	%f27, [%r7+4];
	ld.volatile.shared.f32 	%f28, [%r7];
	max.f32 	%f29, %f28, %f27;
	st.volatile.shared.f32 	[%r7], %f29;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f30, [_ZZ13reducemaxdiffE5sdata];
	abs.f32 	%f31, %f30;
	mov.b32 	 %r19, %f31;
	cvta.to.global.u64 	%rd9, %rd3;
	atom.global.max.s32 	%r20, [%rd9], %r19;

BB0_10:
	ret;
}


`
	reducemaxdiff_ptx_37 = `
.version 6.5
.target sm_37
.address_size 64

	// .globl	reducemaxdiff

.visible .entry reducemaxdiff(
	.param .u64 reducemaxdiff_param_0,
	.param .u64 reducemaxdiff_param_1,
	.param .u64 reducemaxdiff_param_2,
	.param .f32 reducemaxdiff_param_3,
	.param .u32 reducemaxdiff_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<34>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<10>;
	// demoted variable
	.shared .align 4 .b8 _ZZ13reducemaxdiffE5sdata[2048];

	ld.param.u64 	%rd4, [reducemaxdiff_param_0];
	ld.param.u64 	%rd5, [reducemaxdiff_param_1];
	ld.param.u64 	%rd3, [reducemaxdiff_param_2];
	ld.param.f32 	%f33, [reducemaxdiff_param_3];
	ld.param.u32 	%r10, [reducemaxdiff_param_4];
	cvta.to.global.u64 	%rd1, %rd5;
	cvta.to.global.u64 	%rd2, %rd4;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd6, %r21, 4;
	add.s64 	%rd7, %rd2, %rd6;
	add.s64 	%rd8, %rd1, %rd6;
	ld.global.nc.f32 	%f5, [%rd8];
	ld.global.nc.f32 	%f6, [%rd7];
	sub.f32 	%f7, %f6, %f5;
	abs.f32 	%f8, %f7;
	max.f32 	%f33, %f33, %f8;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ13reducemaxdiffE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f33;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f9, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f10, [%r18];
	max.f32 	%f11, %f9, %f10;
	st.shared.f32 	[%r7], %f11;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f12, [%r7];
	ld.volatile.shared.f32 	%f13, [%r7+128];
	max.f32 	%f14, %f12, %f13;
	st.volatile.shared.f32 	[%r7], %f14;
	ld.volatile.shared.f32 	%f15, [%r7+64];
	ld.volatile.shared.f32 	%f16, [%r7];
	max.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%r7], %f17;
	ld.volatile.shared.f32 	%f18, [%r7+32];
	ld.volatile.shared.f32 	%f19, [%r7];
	max.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%r7], %f20;
	ld.volatile.shared.f32 	%f21, [%r7+16];
	ld.volatile.shared.f32 	%f22, [%r7];
	max.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%r7], %f23;
	ld.volatile.shared.f32 	%f24, [%r7+8];
	ld.volatile.shared.f32 	%f25, [%r7];
	max.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%r7], %f26;
	ld.volatile.shared.f32 	%f27, [%r7+4];
	ld.volatile.shared.f32 	%f28, [%r7];
	max.f32 	%f29, %f28, %f27;
	st.volatile.shared.f32 	[%r7], %f29;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f30, [_ZZ13reducemaxdiffE5sdata];
	abs.f32 	%f31, %f30;
	mov.b32 	 %r19, %f31;
	cvta.to.global.u64 	%rd9, %rd3;
	atom.global.max.s32 	%r20, [%rd9], %r19;

BB0_10:
	ret;
}


`
	reducemaxdiff_ptx_50 = `
.version 6.5
.target sm_50
.address_size 64

	// .globl	reducemaxdiff

.visible .entry reducemaxdiff(
	.param .u64 reducemaxdiff_param_0,
	.param .u64 reducemaxdiff_param_1,
	.param .u64 reducemaxdiff_param_2,
	.param .f32 reducemaxdiff_param_3,
	.param .u32 reducemaxdiff_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<34>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<10>;
	// demoted variable
	.shared .align 4 .b8 _ZZ13reducemaxdiffE5sdata[2048];

	ld.param.u64 	%rd4, [reducemaxdiff_param_0];
	ld.param.u64 	%rd5, [reducemaxdiff_param_1];
	ld.param.u64 	%rd3, [reducemaxdiff_param_2];
	ld.param.f32 	%f33, [reducemaxdiff_param_3];
	ld.param.u32 	%r10, [reducemaxdiff_param_4];
	cvta.to.global.u64 	%rd1, %rd5;
	cvta.to.global.u64 	%rd2, %rd4;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd6, %r21, 4;
	add.s64 	%rd7, %rd2, %rd6;
	add.s64 	%rd8, %rd1, %rd6;
	ld.global.nc.f32 	%f5, [%rd8];
	ld.global.nc.f32 	%f6, [%rd7];
	sub.f32 	%f7, %f6, %f5;
	abs.f32 	%f8, %f7;
	max.f32 	%f33, %f33, %f8;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ13reducemaxdiffE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f33;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f9, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f10, [%r18];
	max.f32 	%f11, %f9, %f10;
	st.shared.f32 	[%r7], %f11;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f12, [%r7];
	ld.volatile.shared.f32 	%f13, [%r7+128];
	max.f32 	%f14, %f12, %f13;
	st.volatile.shared.f32 	[%r7], %f14;
	ld.volatile.shared.f32 	%f15, [%r7+64];
	ld.volatile.shared.f32 	%f16, [%r7];
	max.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%r7], %f17;
	ld.volatile.shared.f32 	%f18, [%r7+32];
	ld.volatile.shared.f32 	%f19, [%r7];
	max.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%r7], %f20;
	ld.volatile.shared.f32 	%f21, [%r7+16];
	ld.volatile.shared.f32 	%f22, [%r7];
	max.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%r7], %f23;
	ld.volatile.shared.f32 	%f24, [%r7+8];
	ld.volatile.shared.f32 	%f25, [%r7];
	max.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%r7], %f26;
	ld.volatile.shared.f32 	%f27, [%r7+4];
	ld.volatile.shared.f32 	%f28, [%r7];
	max.f32 	%f29, %f28, %f27;
	st.volatile.shared.f32 	[%r7], %f29;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f30, [_ZZ13reducemaxdiffE5sdata];
	abs.f32 	%f31, %f30;
	mov.b32 	 %r19, %f31;
	cvta.to.global.u64 	%rd9, %rd3;
	atom.global.max.s32 	%r20, [%rd9], %r19;

BB0_10:
	ret;
}


`
	reducemaxdiff_ptx_52 = `
.version 6.5
.target sm_52
.address_size 64

	// .globl	reducemaxdiff

.visible .entry reducemaxdiff(
	.param .u64 reducemaxdiff_param_0,
	.param .u64 reducemaxdiff_param_1,
	.param .u64 reducemaxdiff_param_2,
	.param .f32 reducemaxdiff_param_3,
	.param .u32 reducemaxdiff_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<34>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<10>;
	// demoted variable
	.shared .align 4 .b8 _ZZ13reducemaxdiffE5sdata[2048];

	ld.param.u64 	%rd4, [reducemaxdiff_param_0];
	ld.param.u64 	%rd5, [reducemaxdiff_param_1];
	ld.param.u64 	%rd3, [reducemaxdiff_param_2];
	ld.param.f32 	%f33, [reducemaxdiff_param_3];
	ld.param.u32 	%r10, [reducemaxdiff_param_4];
	cvta.to.global.u64 	%rd1, %rd5;
	cvta.to.global.u64 	%rd2, %rd4;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd6, %r21, 4;
	add.s64 	%rd7, %rd2, %rd6;
	add.s64 	%rd8, %rd1, %rd6;
	ld.global.nc.f32 	%f5, [%rd8];
	ld.global.nc.f32 	%f6, [%rd7];
	sub.f32 	%f7, %f6, %f5;
	abs.f32 	%f8, %f7;
	max.f32 	%f33, %f33, %f8;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ13reducemaxdiffE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f33;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f9, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f10, [%r18];
	max.f32 	%f11, %f9, %f10;
	st.shared.f32 	[%r7], %f11;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f12, [%r7];
	ld.volatile.shared.f32 	%f13, [%r7+128];
	max.f32 	%f14, %f12, %f13;
	st.volatile.shared.f32 	[%r7], %f14;
	ld.volatile.shared.f32 	%f15, [%r7+64];
	ld.volatile.shared.f32 	%f16, [%r7];
	max.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%r7], %f17;
	ld.volatile.shared.f32 	%f18, [%r7+32];
	ld.volatile.shared.f32 	%f19, [%r7];
	max.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%r7], %f20;
	ld.volatile.shared.f32 	%f21, [%r7+16];
	ld.volatile.shared.f32 	%f22, [%r7];
	max.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%r7], %f23;
	ld.volatile.shared.f32 	%f24, [%r7+8];
	ld.volatile.shared.f32 	%f25, [%r7];
	max.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%r7], %f26;
	ld.volatile.shared.f32 	%f27, [%r7+4];
	ld.volatile.shared.f32 	%f28, [%r7];
	max.f32 	%f29, %f28, %f27;
	st.volatile.shared.f32 	[%r7], %f29;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f30, [_ZZ13reducemaxdiffE5sdata];
	abs.f32 	%f31, %f30;
	mov.b32 	 %r19, %f31;
	cvta.to.global.u64 	%rd9, %rd3;
	atom.global.max.s32 	%r20, [%rd9], %r19;

BB0_10:
	ret;
}


`
	reducemaxdiff_ptx_53 = `
.version 6.5
.target sm_53
.address_size 64

	// .globl	reducemaxdiff

.visible .entry reducemaxdiff(
	.param .u64 reducemaxdiff_param_0,
	.param .u64 reducemaxdiff_param_1,
	.param .u64 reducemaxdiff_param_2,
	.param .f32 reducemaxdiff_param_3,
	.param .u32 reducemaxdiff_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<34>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<10>;
	// demoted variable
	.shared .align 4 .b8 _ZZ13reducemaxdiffE5sdata[2048];

	ld.param.u64 	%rd4, [reducemaxdiff_param_0];
	ld.param.u64 	%rd5, [reducemaxdiff_param_1];
	ld.param.u64 	%rd3, [reducemaxdiff_param_2];
	ld.param.f32 	%f33, [reducemaxdiff_param_3];
	ld.param.u32 	%r10, [reducemaxdiff_param_4];
	cvta.to.global.u64 	%rd1, %rd5;
	cvta.to.global.u64 	%rd2, %rd4;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd6, %r21, 4;
	add.s64 	%rd7, %rd2, %rd6;
	add.s64 	%rd8, %rd1, %rd6;
	ld.global.nc.f32 	%f5, [%rd8];
	ld.global.nc.f32 	%f6, [%rd7];
	sub.f32 	%f7, %f6, %f5;
	abs.f32 	%f8, %f7;
	max.f32 	%f33, %f33, %f8;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ13reducemaxdiffE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f33;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f9, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f10, [%r18];
	max.f32 	%f11, %f9, %f10;
	st.shared.f32 	[%r7], %f11;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f12, [%r7];
	ld.volatile.shared.f32 	%f13, [%r7+128];
	max.f32 	%f14, %f12, %f13;
	st.volatile.shared.f32 	[%r7], %f14;
	ld.volatile.shared.f32 	%f15, [%r7+64];
	ld.volatile.shared.f32 	%f16, [%r7];
	max.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%r7], %f17;
	ld.volatile.shared.f32 	%f18, [%r7+32];
	ld.volatile.shared.f32 	%f19, [%r7];
	max.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%r7], %f20;
	ld.volatile.shared.f32 	%f21, [%r7+16];
	ld.volatile.shared.f32 	%f22, [%r7];
	max.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%r7], %f23;
	ld.volatile.shared.f32 	%f24, [%r7+8];
	ld.volatile.shared.f32 	%f25, [%r7];
	max.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%r7], %f26;
	ld.volatile.shared.f32 	%f27, [%r7+4];
	ld.volatile.shared.f32 	%f28, [%r7];
	max.f32 	%f29, %f28, %f27;
	st.volatile.shared.f32 	[%r7], %f29;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f30, [_ZZ13reducemaxdiffE5sdata];
	abs.f32 	%f31, %f30;
	mov.b32 	 %r19, %f31;
	cvta.to.global.u64 	%rd9, %rd3;
	atom.global.max.s32 	%r20, [%rd9], %r19;

BB0_10:
	ret;
}


`
	reducemaxdiff_ptx_60 = `
.version 6.5
.target sm_60
.address_size 64

	// .globl	reducemaxdiff

.visible .entry reducemaxdiff(
	.param .u64 reducemaxdiff_param_0,
	.param .u64 reducemaxdiff_param_1,
	.param .u64 reducemaxdiff_param_2,
	.param .f32 reducemaxdiff_param_3,
	.param .u32 reducemaxdiff_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<34>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<10>;
	// demoted variable
	.shared .align 4 .b8 _ZZ13reducemaxdiffE5sdata[2048];

	ld.param.u64 	%rd4, [reducemaxdiff_param_0];
	ld.param.u64 	%rd5, [reducemaxdiff_param_1];
	ld.param.u64 	%rd3, [reducemaxdiff_param_2];
	ld.param.f32 	%f33, [reducemaxdiff_param_3];
	ld.param.u32 	%r10, [reducemaxdiff_param_4];
	cvta.to.global.u64 	%rd1, %rd5;
	cvta.to.global.u64 	%rd2, %rd4;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd6, %r21, 4;
	add.s64 	%rd7, %rd2, %rd6;
	add.s64 	%rd8, %rd1, %rd6;
	ld.global.nc.f32 	%f5, [%rd8];
	ld.global.nc.f32 	%f6, [%rd7];
	sub.f32 	%f7, %f6, %f5;
	abs.f32 	%f8, %f7;
	max.f32 	%f33, %f33, %f8;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ13reducemaxdiffE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f33;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f9, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f10, [%r18];
	max.f32 	%f11, %f9, %f10;
	st.shared.f32 	[%r7], %f11;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f12, [%r7];
	ld.volatile.shared.f32 	%f13, [%r7+128];
	max.f32 	%f14, %f12, %f13;
	st.volatile.shared.f32 	[%r7], %f14;
	ld.volatile.shared.f32 	%f15, [%r7+64];
	ld.volatile.shared.f32 	%f16, [%r7];
	max.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%r7], %f17;
	ld.volatile.shared.f32 	%f18, [%r7+32];
	ld.volatile.shared.f32 	%f19, [%r7];
	max.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%r7], %f20;
	ld.volatile.shared.f32 	%f21, [%r7+16];
	ld.volatile.shared.f32 	%f22, [%r7];
	max.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%r7], %f23;
	ld.volatile.shared.f32 	%f24, [%r7+8];
	ld.volatile.shared.f32 	%f25, [%r7];
	max.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%r7], %f26;
	ld.volatile.shared.f32 	%f27, [%r7+4];
	ld.volatile.shared.f32 	%f28, [%r7];
	max.f32 	%f29, %f28, %f27;
	st.volatile.shared.f32 	[%r7], %f29;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f30, [_ZZ13reducemaxdiffE5sdata];
	abs.f32 	%f31, %f30;
	mov.b32 	 %r19, %f31;
	cvta.to.global.u64 	%rd9, %rd3;
	atom.global.max.s32 	%r20, [%rd9], %r19;

BB0_10:
	ret;
}


`
	reducemaxdiff_ptx_61 = `
.version 6.5
.target sm_61
.address_size 64

	// .globl	reducemaxdiff

.visible .entry reducemaxdiff(
	.param .u64 reducemaxdiff_param_0,
	.param .u64 reducemaxdiff_param_1,
	.param .u64 reducemaxdiff_param_2,
	.param .f32 reducemaxdiff_param_3,
	.param .u32 reducemaxdiff_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<34>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<10>;
	// demoted variable
	.shared .align 4 .b8 _ZZ13reducemaxdiffE5sdata[2048];

	ld.param.u64 	%rd4, [reducemaxdiff_param_0];
	ld.param.u64 	%rd5, [reducemaxdiff_param_1];
	ld.param.u64 	%rd3, [reducemaxdiff_param_2];
	ld.param.f32 	%f33, [reducemaxdiff_param_3];
	ld.param.u32 	%r10, [reducemaxdiff_param_4];
	cvta.to.global.u64 	%rd1, %rd5;
	cvta.to.global.u64 	%rd2, %rd4;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd6, %r21, 4;
	add.s64 	%rd7, %rd2, %rd6;
	add.s64 	%rd8, %rd1, %rd6;
	ld.global.nc.f32 	%f5, [%rd8];
	ld.global.nc.f32 	%f6, [%rd7];
	sub.f32 	%f7, %f6, %f5;
	abs.f32 	%f8, %f7;
	max.f32 	%f33, %f33, %f8;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ13reducemaxdiffE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f33;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f9, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f10, [%r18];
	max.f32 	%f11, %f9, %f10;
	st.shared.f32 	[%r7], %f11;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f12, [%r7];
	ld.volatile.shared.f32 	%f13, [%r7+128];
	max.f32 	%f14, %f12, %f13;
	st.volatile.shared.f32 	[%r7], %f14;
	ld.volatile.shared.f32 	%f15, [%r7+64];
	ld.volatile.shared.f32 	%f16, [%r7];
	max.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%r7], %f17;
	ld.volatile.shared.f32 	%f18, [%r7+32];
	ld.volatile.shared.f32 	%f19, [%r7];
	max.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%r7], %f20;
	ld.volatile.shared.f32 	%f21, [%r7+16];
	ld.volatile.shared.f32 	%f22, [%r7];
	max.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%r7], %f23;
	ld.volatile.shared.f32 	%f24, [%r7+8];
	ld.volatile.shared.f32 	%f25, [%r7];
	max.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%r7], %f26;
	ld.volatile.shared.f32 	%f27, [%r7+4];
	ld.volatile.shared.f32 	%f28, [%r7];
	max.f32 	%f29, %f28, %f27;
	st.volatile.shared.f32 	[%r7], %f29;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f30, [_ZZ13reducemaxdiffE5sdata];
	abs.f32 	%f31, %f30;
	mov.b32 	 %r19, %f31;
	cvta.to.global.u64 	%rd9, %rd3;
	atom.global.max.s32 	%r20, [%rd9], %r19;

BB0_10:
	ret;
}


`
	reducemaxdiff_ptx_70 = `
.version 6.5
.target sm_70
.address_size 64

	// .globl	reducemaxdiff

.visible .entry reducemaxdiff(
	.param .u64 reducemaxdiff_param_0,
	.param .u64 reducemaxdiff_param_1,
	.param .u64 reducemaxdiff_param_2,
	.param .f32 reducemaxdiff_param_3,
	.param .u32 reducemaxdiff_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<34>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<10>;
	// demoted variable
	.shared .align 4 .b8 _ZZ13reducemaxdiffE5sdata[2048];

	ld.param.u64 	%rd4, [reducemaxdiff_param_0];
	ld.param.u64 	%rd5, [reducemaxdiff_param_1];
	ld.param.u64 	%rd3, [reducemaxdiff_param_2];
	ld.param.f32 	%f33, [reducemaxdiff_param_3];
	ld.param.u32 	%r10, [reducemaxdiff_param_4];
	cvta.to.global.u64 	%rd1, %rd5;
	cvta.to.global.u64 	%rd2, %rd4;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd6, %r21, 4;
	add.s64 	%rd7, %rd2, %rd6;
	add.s64 	%rd8, %rd1, %rd6;
	ld.global.nc.f32 	%f5, [%rd8];
	ld.global.nc.f32 	%f6, [%rd7];
	sub.f32 	%f7, %f6, %f5;
	abs.f32 	%f8, %f7;
	max.f32 	%f33, %f33, %f8;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ13reducemaxdiffE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f33;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f9, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f10, [%r18];
	max.f32 	%f11, %f9, %f10;
	st.shared.f32 	[%r7], %f11;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f12, [%r7];
	ld.volatile.shared.f32 	%f13, [%r7+128];
	max.f32 	%f14, %f12, %f13;
	st.volatile.shared.f32 	[%r7], %f14;
	ld.volatile.shared.f32 	%f15, [%r7+64];
	ld.volatile.shared.f32 	%f16, [%r7];
	max.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%r7], %f17;
	ld.volatile.shared.f32 	%f18, [%r7+32];
	ld.volatile.shared.f32 	%f19, [%r7];
	max.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%r7], %f20;
	ld.volatile.shared.f32 	%f21, [%r7+16];
	ld.volatile.shared.f32 	%f22, [%r7];
	max.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%r7], %f23;
	ld.volatile.shared.f32 	%f24, [%r7+8];
	ld.volatile.shared.f32 	%f25, [%r7];
	max.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%r7], %f26;
	ld.volatile.shared.f32 	%f27, [%r7+4];
	ld.volatile.shared.f32 	%f28, [%r7];
	max.f32 	%f29, %f28, %f27;
	st.volatile.shared.f32 	[%r7], %f29;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f30, [_ZZ13reducemaxdiffE5sdata];
	abs.f32 	%f31, %f30;
	mov.b32 	 %r19, %f31;
	cvta.to.global.u64 	%rd9, %rd3;
	atom.global.max.s32 	%r20, [%rd9], %r19;

BB0_10:
	ret;
}


`
	reducemaxdiff_ptx_75 = `
.version 6.5
.target sm_75
.address_size 64

	// .globl	reducemaxdiff

.visible .entry reducemaxdiff(
	.param .u64 reducemaxdiff_param_0,
	.param .u64 reducemaxdiff_param_1,
	.param .u64 reducemaxdiff_param_2,
	.param .f32 reducemaxdiff_param_3,
	.param .u32 reducemaxdiff_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .f32 	%f<34>;
	.reg .b32 	%r<23>;
	.reg .b64 	%rd<10>;
	// demoted variable
	.shared .align 4 .b8 _ZZ13reducemaxdiffE5sdata[2048];

	ld.param.u64 	%rd4, [reducemaxdiff_param_0];
	ld.param.u64 	%rd5, [reducemaxdiff_param_1];
	ld.param.u64 	%rd3, [reducemaxdiff_param_2];
	ld.param.f32 	%f33, [reducemaxdiff_param_3];
	ld.param.u32 	%r10, [reducemaxdiff_param_4];
	cvta.to.global.u64 	%rd1, %rd5;
	cvta.to.global.u64 	%rd2, %rd4;
	mov.u32 	%r22, %ntid.x;
	mov.u32 	%r11, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r21, %r22, %r11, %r2;
	mov.u32 	%r12, %nctaid.x;
	mul.lo.s32 	%r4, %r12, %r22;
	setp.ge.s32	%p1, %r21, %r10;
	@%p1 bra 	BB0_2;

BB0_1:
	mul.wide.s32 	%rd6, %r21, 4;
	add.s64 	%rd7, %rd2, %rd6;
	add.s64 	%rd8, %rd1, %rd6;
	ld.global.nc.f32 	%f5, [%rd8];
	ld.global.nc.f32 	%f6, [%rd7];
	sub.f32 	%f7, %f6, %f5;
	abs.f32 	%f8, %f7;
	max.f32 	%f33, %f33, %f8;
	add.s32 	%r21, %r21, %r4;
	setp.lt.s32	%p2, %r21, %r10;
	@%p2 bra 	BB0_1;

BB0_2:
	shl.b32 	%r13, %r2, 2;
	mov.u32 	%r14, _ZZ13reducemaxdiffE5sdata;
	add.s32 	%r7, %r14, %r13;
	st.shared.f32 	[%r7], %f33;
	bar.sync 	0;
	setp.lt.u32	%p3, %r22, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	shr.u32 	%r9, %r22, 1;
	setp.ge.u32	%p4, %r2, %r9;
	@%p4 bra 	BB0_5;

	ld.shared.f32 	%f9, [%r7];
	add.s32 	%r15, %r9, %r2;
	shl.b32 	%r16, %r15, 2;
	add.s32 	%r18, %r14, %r16;
	ld.shared.f32 	%f10, [%r18];
	max.f32 	%f11, %f9, %f10;
	st.shared.f32 	[%r7], %f11;

BB0_5:
	bar.sync 	0;
	setp.gt.u32	%p5, %r22, 131;
	mov.u32 	%r22, %r9;
	@%p5 bra 	BB0_3;

BB0_6:
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	ld.volatile.shared.f32 	%f12, [%r7];
	ld.volatile.shared.f32 	%f13, [%r7+128];
	max.f32 	%f14, %f12, %f13;
	st.volatile.shared.f32 	[%r7], %f14;
	ld.volatile.shared.f32 	%f15, [%r7+64];
	ld.volatile.shared.f32 	%f16, [%r7];
	max.f32 	%f17, %f16, %f15;
	st.volatile.shared.f32 	[%r7], %f17;
	ld.volatile.shared.f32 	%f18, [%r7+32];
	ld.volatile.shared.f32 	%f19, [%r7];
	max.f32 	%f20, %f19, %f18;
	st.volatile.shared.f32 	[%r7], %f20;
	ld.volatile.shared.f32 	%f21, [%r7+16];
	ld.volatile.shared.f32 	%f22, [%r7];
	max.f32 	%f23, %f22, %f21;
	st.volatile.shared.f32 	[%r7], %f23;
	ld.volatile.shared.f32 	%f24, [%r7+8];
	ld.volatile.shared.f32 	%f25, [%r7];
	max.f32 	%f26, %f25, %f24;
	st.volatile.shared.f32 	[%r7], %f26;
	ld.volatile.shared.f32 	%f27, [%r7+4];
	ld.volatile.shared.f32 	%f28, [%r7];
	max.f32 	%f29, %f28, %f27;
	st.volatile.shared.f32 	[%r7], %f29;

BB0_8:
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	ld.shared.f32 	%f30, [_ZZ13reducemaxdiffE5sdata];
	abs.f32 	%f31, %f30;
	mov.b32 	 %r19, %f31;
	cvta.to.global.u64 	%rd9, %rd3;
	atom.global.max.s32 	%r20, [%rd9], %r19;

BB0_10:
	ret;
}


`
)
