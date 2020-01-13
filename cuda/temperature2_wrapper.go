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

// CUDA handle for settemperature2 kernel
var settemperature2_code cu.Function

// Stores the arguments for settemperature2 kernel invocation
type settemperature2_args_t struct {
	arg_B            unsafe.Pointer
	arg_noise        unsafe.Pointer
	arg_kB2_VgammaDt float32
	arg_Ms_          unsafe.Pointer
	arg_Ms_mul       float32
	arg_temp_        unsafe.Pointer
	arg_temp_mul     float32
	arg_alpha_       unsafe.Pointer
	arg_alpha_mul    float32
	arg_N            int
	argptr           [10]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for settemperature2 kernel invocation
var settemperature2_args settemperature2_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	settemperature2_args.argptr[0] = unsafe.Pointer(&settemperature2_args.arg_B)
	settemperature2_args.argptr[1] = unsafe.Pointer(&settemperature2_args.arg_noise)
	settemperature2_args.argptr[2] = unsafe.Pointer(&settemperature2_args.arg_kB2_VgammaDt)
	settemperature2_args.argptr[3] = unsafe.Pointer(&settemperature2_args.arg_Ms_)
	settemperature2_args.argptr[4] = unsafe.Pointer(&settemperature2_args.arg_Ms_mul)
	settemperature2_args.argptr[5] = unsafe.Pointer(&settemperature2_args.arg_temp_)
	settemperature2_args.argptr[6] = unsafe.Pointer(&settemperature2_args.arg_temp_mul)
	settemperature2_args.argptr[7] = unsafe.Pointer(&settemperature2_args.arg_alpha_)
	settemperature2_args.argptr[8] = unsafe.Pointer(&settemperature2_args.arg_alpha_mul)
	settemperature2_args.argptr[9] = unsafe.Pointer(&settemperature2_args.arg_N)
}

// Wrapper for settemperature2 CUDA kernel, asynchronous.
func k_settemperature2_async(B unsafe.Pointer, noise unsafe.Pointer, kB2_VgammaDt float32, Ms_ unsafe.Pointer, Ms_mul float32, temp_ unsafe.Pointer, temp_mul float32, alpha_ unsafe.Pointer, alpha_mul float32, N int, cfg *config) {
	if Synchronous { // debug
		Sync()
		timer.Start("settemperature2")
	}

	settemperature2_args.Lock()
	defer settemperature2_args.Unlock()

	if settemperature2_code == 0 {
		settemperature2_code = fatbinLoad(settemperature2_map, "settemperature2")
	}

	settemperature2_args.arg_B = B
	settemperature2_args.arg_noise = noise
	settemperature2_args.arg_kB2_VgammaDt = kB2_VgammaDt
	settemperature2_args.arg_Ms_ = Ms_
	settemperature2_args.arg_Ms_mul = Ms_mul
	settemperature2_args.arg_temp_ = temp_
	settemperature2_args.arg_temp_mul = temp_mul
	settemperature2_args.arg_alpha_ = alpha_
	settemperature2_args.arg_alpha_mul = alpha_mul
	settemperature2_args.arg_N = N

	args := settemperature2_args.argptr[:]
	cu.LaunchKernel(settemperature2_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
		timer.Stop("settemperature2")
	}
}

// maps compute capability on PTX code for settemperature2 kernel.
var settemperature2_map = map[int]string{0: "",
	30: settemperature2_ptx_30,
	35: settemperature2_ptx_35,
	37: settemperature2_ptx_37,
	50: settemperature2_ptx_50,
	52: settemperature2_ptx_52,
	53: settemperature2_ptx_53,
	60: settemperature2_ptx_60,
	61: settemperature2_ptx_61,
	70: settemperature2_ptx_70,
	75: settemperature2_ptx_75}

// settemperature2 PTX code for various compute capabilities.
const (
	settemperature2_ptx_30 = `
.version 6.5
.target sm_30
.address_size 64

	// .globl	settemperature2

.visible .entry settemperature2(
	.param .u64 settemperature2_param_0,
	.param .u64 settemperature2_param_1,
	.param .f32 settemperature2_param_2,
	.param .u64 settemperature2_param_3,
	.param .f32 settemperature2_param_4,
	.param .u64 settemperature2_param_5,
	.param .f32 settemperature2_param_6,
	.param .u64 settemperature2_param_7,
	.param .f32 settemperature2_param_8,
	.param .u32 settemperature2_param_9
)
{
	.reg .pred 	%p<6>;
	.reg .f32 	%f<27>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<20>;


	ld.param.u64 	%rd1, [settemperature2_param_0];
	ld.param.u64 	%rd2, [settemperature2_param_1];
	ld.param.f32 	%f9, [settemperature2_param_2];
	ld.param.u64 	%rd3, [settemperature2_param_3];
	ld.param.f32 	%f23, [settemperature2_param_4];
	ld.param.u64 	%rd4, [settemperature2_param_5];
	ld.param.f32 	%f25, [settemperature2_param_6];
	ld.param.u64 	%rd5, [settemperature2_param_7];
	ld.param.f32 	%f26, [settemperature2_param_8];
	ld.param.u32 	%r2, [settemperature2_param_9];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_10;

	setp.eq.s64	%p2, %rd3, 0;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd6, %rd3;
	mul.wide.s32 	%rd7, %r1, 4;
	add.s64 	%rd8, %rd6, %rd7;
	ld.global.f32 	%f13, [%rd8];
	mul.f32 	%f23, %f13, %f23;

BB0_3:
	setp.eq.f32	%p3, %f23, 0f00000000;
	mov.f32 	%f24, 0f00000000;
	@%p3 bra 	BB0_5;

	rcp.rn.f32 	%f24, %f23;

BB0_5:
	setp.eq.s64	%p4, %rd4, 0;
	@%p4 bra 	BB0_7;

	cvta.to.global.u64 	%rd9, %rd4;
	mul.wide.s32 	%rd10, %r1, 4;
	add.s64 	%rd11, %rd9, %rd10;
	ld.global.f32 	%f15, [%rd11];
	mul.f32 	%f25, %f15, %f25;

BB0_7:
	setp.eq.s64	%p5, %rd5, 0;
	@%p5 bra 	BB0_9;

	cvta.to.global.u64 	%rd12, %rd5;
	mul.wide.s32 	%rd13, %r1, 4;
	add.s64 	%rd14, %rd12, %rd13;
	ld.global.f32 	%f16, [%rd14];
	mul.f32 	%f26, %f16, %f26;

BB0_9:
	cvta.to.global.u64 	%rd15, %rd1;
	cvta.to.global.u64 	%rd16, %rd2;
	mul.wide.s32 	%rd17, %r1, 4;
	add.s64 	%rd18, %rd16, %rd17;
	mul.f32 	%f17, %f26, %f9;
	mul.f32 	%f18, %f25, %f17;
	mul.f32 	%f19, %f24, %f18;
	sqrt.rn.f32 	%f20, %f19;
	ld.global.f32 	%f21, [%rd18];
	mul.f32 	%f22, %f21, %f20;
	add.s64 	%rd19, %rd15, %rd17;
	st.global.f32 	[%rd19], %f22;

BB0_10:
	ret;
}


`
	settemperature2_ptx_35 = `
.version 6.5
.target sm_35
.address_size 64

	// .globl	settemperature2

.visible .entry settemperature2(
	.param .u64 settemperature2_param_0,
	.param .u64 settemperature2_param_1,
	.param .f32 settemperature2_param_2,
	.param .u64 settemperature2_param_3,
	.param .f32 settemperature2_param_4,
	.param .u64 settemperature2_param_5,
	.param .f32 settemperature2_param_6,
	.param .u64 settemperature2_param_7,
	.param .f32 settemperature2_param_8,
	.param .u32 settemperature2_param_9
)
{
	.reg .pred 	%p<6>;
	.reg .f32 	%f<27>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<20>;


	ld.param.u64 	%rd1, [settemperature2_param_0];
	ld.param.u64 	%rd2, [settemperature2_param_1];
	ld.param.f32 	%f9, [settemperature2_param_2];
	ld.param.u64 	%rd3, [settemperature2_param_3];
	ld.param.f32 	%f23, [settemperature2_param_4];
	ld.param.u64 	%rd4, [settemperature2_param_5];
	ld.param.f32 	%f25, [settemperature2_param_6];
	ld.param.u64 	%rd5, [settemperature2_param_7];
	ld.param.f32 	%f26, [settemperature2_param_8];
	ld.param.u32 	%r2, [settemperature2_param_9];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_10;

	setp.eq.s64	%p2, %rd3, 0;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd6, %rd3;
	mul.wide.s32 	%rd7, %r1, 4;
	add.s64 	%rd8, %rd6, %rd7;
	ld.global.nc.f32 	%f13, [%rd8];
	mul.f32 	%f23, %f13, %f23;

BB0_3:
	setp.eq.f32	%p3, %f23, 0f00000000;
	mov.f32 	%f24, 0f00000000;
	@%p3 bra 	BB0_5;

	rcp.rn.f32 	%f24, %f23;

BB0_5:
	setp.eq.s64	%p4, %rd4, 0;
	@%p4 bra 	BB0_7;

	cvta.to.global.u64 	%rd9, %rd4;
	mul.wide.s32 	%rd10, %r1, 4;
	add.s64 	%rd11, %rd9, %rd10;
	ld.global.nc.f32 	%f15, [%rd11];
	mul.f32 	%f25, %f15, %f25;

BB0_7:
	setp.eq.s64	%p5, %rd5, 0;
	@%p5 bra 	BB0_9;

	cvta.to.global.u64 	%rd12, %rd5;
	mul.wide.s32 	%rd13, %r1, 4;
	add.s64 	%rd14, %rd12, %rd13;
	ld.global.nc.f32 	%f16, [%rd14];
	mul.f32 	%f26, %f16, %f26;

BB0_9:
	cvta.to.global.u64 	%rd15, %rd1;
	cvta.to.global.u64 	%rd16, %rd2;
	mul.wide.s32 	%rd17, %r1, 4;
	add.s64 	%rd18, %rd16, %rd17;
	mul.f32 	%f17, %f26, %f9;
	mul.f32 	%f18, %f25, %f17;
	mul.f32 	%f19, %f24, %f18;
	sqrt.rn.f32 	%f20, %f19;
	ld.global.nc.f32 	%f21, [%rd18];
	mul.f32 	%f22, %f21, %f20;
	add.s64 	%rd19, %rd15, %rd17;
	st.global.f32 	[%rd19], %f22;

BB0_10:
	ret;
}


`
	settemperature2_ptx_37 = `
.version 6.5
.target sm_37
.address_size 64

	// .globl	settemperature2

.visible .entry settemperature2(
	.param .u64 settemperature2_param_0,
	.param .u64 settemperature2_param_1,
	.param .f32 settemperature2_param_2,
	.param .u64 settemperature2_param_3,
	.param .f32 settemperature2_param_4,
	.param .u64 settemperature2_param_5,
	.param .f32 settemperature2_param_6,
	.param .u64 settemperature2_param_7,
	.param .f32 settemperature2_param_8,
	.param .u32 settemperature2_param_9
)
{
	.reg .pred 	%p<6>;
	.reg .f32 	%f<27>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<20>;


	ld.param.u64 	%rd1, [settemperature2_param_0];
	ld.param.u64 	%rd2, [settemperature2_param_1];
	ld.param.f32 	%f9, [settemperature2_param_2];
	ld.param.u64 	%rd3, [settemperature2_param_3];
	ld.param.f32 	%f23, [settemperature2_param_4];
	ld.param.u64 	%rd4, [settemperature2_param_5];
	ld.param.f32 	%f25, [settemperature2_param_6];
	ld.param.u64 	%rd5, [settemperature2_param_7];
	ld.param.f32 	%f26, [settemperature2_param_8];
	ld.param.u32 	%r2, [settemperature2_param_9];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_10;

	setp.eq.s64	%p2, %rd3, 0;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd6, %rd3;
	mul.wide.s32 	%rd7, %r1, 4;
	add.s64 	%rd8, %rd6, %rd7;
	ld.global.nc.f32 	%f13, [%rd8];
	mul.f32 	%f23, %f13, %f23;

BB0_3:
	setp.eq.f32	%p3, %f23, 0f00000000;
	mov.f32 	%f24, 0f00000000;
	@%p3 bra 	BB0_5;

	rcp.rn.f32 	%f24, %f23;

BB0_5:
	setp.eq.s64	%p4, %rd4, 0;
	@%p4 bra 	BB0_7;

	cvta.to.global.u64 	%rd9, %rd4;
	mul.wide.s32 	%rd10, %r1, 4;
	add.s64 	%rd11, %rd9, %rd10;
	ld.global.nc.f32 	%f15, [%rd11];
	mul.f32 	%f25, %f15, %f25;

BB0_7:
	setp.eq.s64	%p5, %rd5, 0;
	@%p5 bra 	BB0_9;

	cvta.to.global.u64 	%rd12, %rd5;
	mul.wide.s32 	%rd13, %r1, 4;
	add.s64 	%rd14, %rd12, %rd13;
	ld.global.nc.f32 	%f16, [%rd14];
	mul.f32 	%f26, %f16, %f26;

BB0_9:
	cvta.to.global.u64 	%rd15, %rd1;
	cvta.to.global.u64 	%rd16, %rd2;
	mul.wide.s32 	%rd17, %r1, 4;
	add.s64 	%rd18, %rd16, %rd17;
	mul.f32 	%f17, %f26, %f9;
	mul.f32 	%f18, %f25, %f17;
	mul.f32 	%f19, %f24, %f18;
	sqrt.rn.f32 	%f20, %f19;
	ld.global.nc.f32 	%f21, [%rd18];
	mul.f32 	%f22, %f21, %f20;
	add.s64 	%rd19, %rd15, %rd17;
	st.global.f32 	[%rd19], %f22;

BB0_10:
	ret;
}


`
	settemperature2_ptx_50 = `
.version 6.5
.target sm_50
.address_size 64

	// .globl	settemperature2

.visible .entry settemperature2(
	.param .u64 settemperature2_param_0,
	.param .u64 settemperature2_param_1,
	.param .f32 settemperature2_param_2,
	.param .u64 settemperature2_param_3,
	.param .f32 settemperature2_param_4,
	.param .u64 settemperature2_param_5,
	.param .f32 settemperature2_param_6,
	.param .u64 settemperature2_param_7,
	.param .f32 settemperature2_param_8,
	.param .u32 settemperature2_param_9
)
{
	.reg .pred 	%p<6>;
	.reg .f32 	%f<27>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<20>;


	ld.param.u64 	%rd1, [settemperature2_param_0];
	ld.param.u64 	%rd2, [settemperature2_param_1];
	ld.param.f32 	%f9, [settemperature2_param_2];
	ld.param.u64 	%rd3, [settemperature2_param_3];
	ld.param.f32 	%f23, [settemperature2_param_4];
	ld.param.u64 	%rd4, [settemperature2_param_5];
	ld.param.f32 	%f25, [settemperature2_param_6];
	ld.param.u64 	%rd5, [settemperature2_param_7];
	ld.param.f32 	%f26, [settemperature2_param_8];
	ld.param.u32 	%r2, [settemperature2_param_9];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_10;

	setp.eq.s64	%p2, %rd3, 0;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd6, %rd3;
	mul.wide.s32 	%rd7, %r1, 4;
	add.s64 	%rd8, %rd6, %rd7;
	ld.global.nc.f32 	%f13, [%rd8];
	mul.f32 	%f23, %f13, %f23;

BB0_3:
	setp.eq.f32	%p3, %f23, 0f00000000;
	mov.f32 	%f24, 0f00000000;
	@%p3 bra 	BB0_5;

	rcp.rn.f32 	%f24, %f23;

BB0_5:
	setp.eq.s64	%p4, %rd4, 0;
	@%p4 bra 	BB0_7;

	cvta.to.global.u64 	%rd9, %rd4;
	mul.wide.s32 	%rd10, %r1, 4;
	add.s64 	%rd11, %rd9, %rd10;
	ld.global.nc.f32 	%f15, [%rd11];
	mul.f32 	%f25, %f15, %f25;

BB0_7:
	setp.eq.s64	%p5, %rd5, 0;
	@%p5 bra 	BB0_9;

	cvta.to.global.u64 	%rd12, %rd5;
	mul.wide.s32 	%rd13, %r1, 4;
	add.s64 	%rd14, %rd12, %rd13;
	ld.global.nc.f32 	%f16, [%rd14];
	mul.f32 	%f26, %f16, %f26;

BB0_9:
	cvta.to.global.u64 	%rd15, %rd1;
	cvta.to.global.u64 	%rd16, %rd2;
	mul.wide.s32 	%rd17, %r1, 4;
	add.s64 	%rd18, %rd16, %rd17;
	mul.f32 	%f17, %f26, %f9;
	mul.f32 	%f18, %f25, %f17;
	mul.f32 	%f19, %f24, %f18;
	sqrt.rn.f32 	%f20, %f19;
	ld.global.nc.f32 	%f21, [%rd18];
	mul.f32 	%f22, %f21, %f20;
	add.s64 	%rd19, %rd15, %rd17;
	st.global.f32 	[%rd19], %f22;

BB0_10:
	ret;
}


`
	settemperature2_ptx_52 = `
.version 6.5
.target sm_52
.address_size 64

	// .globl	settemperature2

.visible .entry settemperature2(
	.param .u64 settemperature2_param_0,
	.param .u64 settemperature2_param_1,
	.param .f32 settemperature2_param_2,
	.param .u64 settemperature2_param_3,
	.param .f32 settemperature2_param_4,
	.param .u64 settemperature2_param_5,
	.param .f32 settemperature2_param_6,
	.param .u64 settemperature2_param_7,
	.param .f32 settemperature2_param_8,
	.param .u32 settemperature2_param_9
)
{
	.reg .pred 	%p<6>;
	.reg .f32 	%f<27>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<20>;


	ld.param.u64 	%rd1, [settemperature2_param_0];
	ld.param.u64 	%rd2, [settemperature2_param_1];
	ld.param.f32 	%f9, [settemperature2_param_2];
	ld.param.u64 	%rd3, [settemperature2_param_3];
	ld.param.f32 	%f23, [settemperature2_param_4];
	ld.param.u64 	%rd4, [settemperature2_param_5];
	ld.param.f32 	%f25, [settemperature2_param_6];
	ld.param.u64 	%rd5, [settemperature2_param_7];
	ld.param.f32 	%f26, [settemperature2_param_8];
	ld.param.u32 	%r2, [settemperature2_param_9];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_10;

	setp.eq.s64	%p2, %rd3, 0;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd6, %rd3;
	mul.wide.s32 	%rd7, %r1, 4;
	add.s64 	%rd8, %rd6, %rd7;
	ld.global.nc.f32 	%f13, [%rd8];
	mul.f32 	%f23, %f13, %f23;

BB0_3:
	setp.eq.f32	%p3, %f23, 0f00000000;
	mov.f32 	%f24, 0f00000000;
	@%p3 bra 	BB0_5;

	rcp.rn.f32 	%f24, %f23;

BB0_5:
	setp.eq.s64	%p4, %rd4, 0;
	@%p4 bra 	BB0_7;

	cvta.to.global.u64 	%rd9, %rd4;
	mul.wide.s32 	%rd10, %r1, 4;
	add.s64 	%rd11, %rd9, %rd10;
	ld.global.nc.f32 	%f15, [%rd11];
	mul.f32 	%f25, %f15, %f25;

BB0_7:
	setp.eq.s64	%p5, %rd5, 0;
	@%p5 bra 	BB0_9;

	cvta.to.global.u64 	%rd12, %rd5;
	mul.wide.s32 	%rd13, %r1, 4;
	add.s64 	%rd14, %rd12, %rd13;
	ld.global.nc.f32 	%f16, [%rd14];
	mul.f32 	%f26, %f16, %f26;

BB0_9:
	cvta.to.global.u64 	%rd15, %rd1;
	cvta.to.global.u64 	%rd16, %rd2;
	mul.wide.s32 	%rd17, %r1, 4;
	add.s64 	%rd18, %rd16, %rd17;
	mul.f32 	%f17, %f26, %f9;
	mul.f32 	%f18, %f25, %f17;
	mul.f32 	%f19, %f24, %f18;
	sqrt.rn.f32 	%f20, %f19;
	ld.global.nc.f32 	%f21, [%rd18];
	mul.f32 	%f22, %f21, %f20;
	add.s64 	%rd19, %rd15, %rd17;
	st.global.f32 	[%rd19], %f22;

BB0_10:
	ret;
}


`
	settemperature2_ptx_53 = `
.version 6.5
.target sm_53
.address_size 64

	// .globl	settemperature2

.visible .entry settemperature2(
	.param .u64 settemperature2_param_0,
	.param .u64 settemperature2_param_1,
	.param .f32 settemperature2_param_2,
	.param .u64 settemperature2_param_3,
	.param .f32 settemperature2_param_4,
	.param .u64 settemperature2_param_5,
	.param .f32 settemperature2_param_6,
	.param .u64 settemperature2_param_7,
	.param .f32 settemperature2_param_8,
	.param .u32 settemperature2_param_9
)
{
	.reg .pred 	%p<6>;
	.reg .f32 	%f<27>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<20>;


	ld.param.u64 	%rd1, [settemperature2_param_0];
	ld.param.u64 	%rd2, [settemperature2_param_1];
	ld.param.f32 	%f9, [settemperature2_param_2];
	ld.param.u64 	%rd3, [settemperature2_param_3];
	ld.param.f32 	%f23, [settemperature2_param_4];
	ld.param.u64 	%rd4, [settemperature2_param_5];
	ld.param.f32 	%f25, [settemperature2_param_6];
	ld.param.u64 	%rd5, [settemperature2_param_7];
	ld.param.f32 	%f26, [settemperature2_param_8];
	ld.param.u32 	%r2, [settemperature2_param_9];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_10;

	setp.eq.s64	%p2, %rd3, 0;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd6, %rd3;
	mul.wide.s32 	%rd7, %r1, 4;
	add.s64 	%rd8, %rd6, %rd7;
	ld.global.nc.f32 	%f13, [%rd8];
	mul.f32 	%f23, %f13, %f23;

BB0_3:
	setp.eq.f32	%p3, %f23, 0f00000000;
	mov.f32 	%f24, 0f00000000;
	@%p3 bra 	BB0_5;

	rcp.rn.f32 	%f24, %f23;

BB0_5:
	setp.eq.s64	%p4, %rd4, 0;
	@%p4 bra 	BB0_7;

	cvta.to.global.u64 	%rd9, %rd4;
	mul.wide.s32 	%rd10, %r1, 4;
	add.s64 	%rd11, %rd9, %rd10;
	ld.global.nc.f32 	%f15, [%rd11];
	mul.f32 	%f25, %f15, %f25;

BB0_7:
	setp.eq.s64	%p5, %rd5, 0;
	@%p5 bra 	BB0_9;

	cvta.to.global.u64 	%rd12, %rd5;
	mul.wide.s32 	%rd13, %r1, 4;
	add.s64 	%rd14, %rd12, %rd13;
	ld.global.nc.f32 	%f16, [%rd14];
	mul.f32 	%f26, %f16, %f26;

BB0_9:
	cvta.to.global.u64 	%rd15, %rd1;
	cvta.to.global.u64 	%rd16, %rd2;
	mul.wide.s32 	%rd17, %r1, 4;
	add.s64 	%rd18, %rd16, %rd17;
	mul.f32 	%f17, %f26, %f9;
	mul.f32 	%f18, %f25, %f17;
	mul.f32 	%f19, %f24, %f18;
	sqrt.rn.f32 	%f20, %f19;
	ld.global.nc.f32 	%f21, [%rd18];
	mul.f32 	%f22, %f21, %f20;
	add.s64 	%rd19, %rd15, %rd17;
	st.global.f32 	[%rd19], %f22;

BB0_10:
	ret;
}


`
	settemperature2_ptx_60 = `
.version 6.5
.target sm_60
.address_size 64

	// .globl	settemperature2

.visible .entry settemperature2(
	.param .u64 settemperature2_param_0,
	.param .u64 settemperature2_param_1,
	.param .f32 settemperature2_param_2,
	.param .u64 settemperature2_param_3,
	.param .f32 settemperature2_param_4,
	.param .u64 settemperature2_param_5,
	.param .f32 settemperature2_param_6,
	.param .u64 settemperature2_param_7,
	.param .f32 settemperature2_param_8,
	.param .u32 settemperature2_param_9
)
{
	.reg .pred 	%p<6>;
	.reg .f32 	%f<27>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<20>;


	ld.param.u64 	%rd1, [settemperature2_param_0];
	ld.param.u64 	%rd2, [settemperature2_param_1];
	ld.param.f32 	%f9, [settemperature2_param_2];
	ld.param.u64 	%rd3, [settemperature2_param_3];
	ld.param.f32 	%f23, [settemperature2_param_4];
	ld.param.u64 	%rd4, [settemperature2_param_5];
	ld.param.f32 	%f25, [settemperature2_param_6];
	ld.param.u64 	%rd5, [settemperature2_param_7];
	ld.param.f32 	%f26, [settemperature2_param_8];
	ld.param.u32 	%r2, [settemperature2_param_9];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_10;

	setp.eq.s64	%p2, %rd3, 0;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd6, %rd3;
	mul.wide.s32 	%rd7, %r1, 4;
	add.s64 	%rd8, %rd6, %rd7;
	ld.global.nc.f32 	%f13, [%rd8];
	mul.f32 	%f23, %f13, %f23;

BB0_3:
	setp.eq.f32	%p3, %f23, 0f00000000;
	mov.f32 	%f24, 0f00000000;
	@%p3 bra 	BB0_5;

	rcp.rn.f32 	%f24, %f23;

BB0_5:
	setp.eq.s64	%p4, %rd4, 0;
	@%p4 bra 	BB0_7;

	cvta.to.global.u64 	%rd9, %rd4;
	mul.wide.s32 	%rd10, %r1, 4;
	add.s64 	%rd11, %rd9, %rd10;
	ld.global.nc.f32 	%f15, [%rd11];
	mul.f32 	%f25, %f15, %f25;

BB0_7:
	setp.eq.s64	%p5, %rd5, 0;
	@%p5 bra 	BB0_9;

	cvta.to.global.u64 	%rd12, %rd5;
	mul.wide.s32 	%rd13, %r1, 4;
	add.s64 	%rd14, %rd12, %rd13;
	ld.global.nc.f32 	%f16, [%rd14];
	mul.f32 	%f26, %f16, %f26;

BB0_9:
	cvta.to.global.u64 	%rd15, %rd1;
	cvta.to.global.u64 	%rd16, %rd2;
	mul.wide.s32 	%rd17, %r1, 4;
	add.s64 	%rd18, %rd16, %rd17;
	mul.f32 	%f17, %f26, %f9;
	mul.f32 	%f18, %f25, %f17;
	mul.f32 	%f19, %f24, %f18;
	sqrt.rn.f32 	%f20, %f19;
	ld.global.nc.f32 	%f21, [%rd18];
	mul.f32 	%f22, %f21, %f20;
	add.s64 	%rd19, %rd15, %rd17;
	st.global.f32 	[%rd19], %f22;

BB0_10:
	ret;
}


`
	settemperature2_ptx_61 = `
.version 6.5
.target sm_61
.address_size 64

	// .globl	settemperature2

.visible .entry settemperature2(
	.param .u64 settemperature2_param_0,
	.param .u64 settemperature2_param_1,
	.param .f32 settemperature2_param_2,
	.param .u64 settemperature2_param_3,
	.param .f32 settemperature2_param_4,
	.param .u64 settemperature2_param_5,
	.param .f32 settemperature2_param_6,
	.param .u64 settemperature2_param_7,
	.param .f32 settemperature2_param_8,
	.param .u32 settemperature2_param_9
)
{
	.reg .pred 	%p<6>;
	.reg .f32 	%f<27>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<20>;


	ld.param.u64 	%rd1, [settemperature2_param_0];
	ld.param.u64 	%rd2, [settemperature2_param_1];
	ld.param.f32 	%f9, [settemperature2_param_2];
	ld.param.u64 	%rd3, [settemperature2_param_3];
	ld.param.f32 	%f23, [settemperature2_param_4];
	ld.param.u64 	%rd4, [settemperature2_param_5];
	ld.param.f32 	%f25, [settemperature2_param_6];
	ld.param.u64 	%rd5, [settemperature2_param_7];
	ld.param.f32 	%f26, [settemperature2_param_8];
	ld.param.u32 	%r2, [settemperature2_param_9];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_10;

	setp.eq.s64	%p2, %rd3, 0;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd6, %rd3;
	mul.wide.s32 	%rd7, %r1, 4;
	add.s64 	%rd8, %rd6, %rd7;
	ld.global.nc.f32 	%f13, [%rd8];
	mul.f32 	%f23, %f13, %f23;

BB0_3:
	setp.eq.f32	%p3, %f23, 0f00000000;
	mov.f32 	%f24, 0f00000000;
	@%p3 bra 	BB0_5;

	rcp.rn.f32 	%f24, %f23;

BB0_5:
	setp.eq.s64	%p4, %rd4, 0;
	@%p4 bra 	BB0_7;

	cvta.to.global.u64 	%rd9, %rd4;
	mul.wide.s32 	%rd10, %r1, 4;
	add.s64 	%rd11, %rd9, %rd10;
	ld.global.nc.f32 	%f15, [%rd11];
	mul.f32 	%f25, %f15, %f25;

BB0_7:
	setp.eq.s64	%p5, %rd5, 0;
	@%p5 bra 	BB0_9;

	cvta.to.global.u64 	%rd12, %rd5;
	mul.wide.s32 	%rd13, %r1, 4;
	add.s64 	%rd14, %rd12, %rd13;
	ld.global.nc.f32 	%f16, [%rd14];
	mul.f32 	%f26, %f16, %f26;

BB0_9:
	cvta.to.global.u64 	%rd15, %rd1;
	cvta.to.global.u64 	%rd16, %rd2;
	mul.wide.s32 	%rd17, %r1, 4;
	add.s64 	%rd18, %rd16, %rd17;
	mul.f32 	%f17, %f26, %f9;
	mul.f32 	%f18, %f25, %f17;
	mul.f32 	%f19, %f24, %f18;
	sqrt.rn.f32 	%f20, %f19;
	ld.global.nc.f32 	%f21, [%rd18];
	mul.f32 	%f22, %f21, %f20;
	add.s64 	%rd19, %rd15, %rd17;
	st.global.f32 	[%rd19], %f22;

BB0_10:
	ret;
}


`
	settemperature2_ptx_70 = `
.version 6.5
.target sm_70
.address_size 64

	// .globl	settemperature2

.visible .entry settemperature2(
	.param .u64 settemperature2_param_0,
	.param .u64 settemperature2_param_1,
	.param .f32 settemperature2_param_2,
	.param .u64 settemperature2_param_3,
	.param .f32 settemperature2_param_4,
	.param .u64 settemperature2_param_5,
	.param .f32 settemperature2_param_6,
	.param .u64 settemperature2_param_7,
	.param .f32 settemperature2_param_8,
	.param .u32 settemperature2_param_9
)
{
	.reg .pred 	%p<6>;
	.reg .f32 	%f<27>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<20>;


	ld.param.u64 	%rd1, [settemperature2_param_0];
	ld.param.u64 	%rd2, [settemperature2_param_1];
	ld.param.f32 	%f9, [settemperature2_param_2];
	ld.param.u64 	%rd3, [settemperature2_param_3];
	ld.param.f32 	%f23, [settemperature2_param_4];
	ld.param.u64 	%rd4, [settemperature2_param_5];
	ld.param.f32 	%f25, [settemperature2_param_6];
	ld.param.u64 	%rd5, [settemperature2_param_7];
	ld.param.f32 	%f26, [settemperature2_param_8];
	ld.param.u32 	%r2, [settemperature2_param_9];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_10;

	setp.eq.s64	%p2, %rd3, 0;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd6, %rd3;
	mul.wide.s32 	%rd7, %r1, 4;
	add.s64 	%rd8, %rd6, %rd7;
	ld.global.nc.f32 	%f13, [%rd8];
	mul.f32 	%f23, %f13, %f23;

BB0_3:
	setp.eq.f32	%p3, %f23, 0f00000000;
	mov.f32 	%f24, 0f00000000;
	@%p3 bra 	BB0_5;

	rcp.rn.f32 	%f24, %f23;

BB0_5:
	setp.eq.s64	%p4, %rd4, 0;
	@%p4 bra 	BB0_7;

	cvta.to.global.u64 	%rd9, %rd4;
	mul.wide.s32 	%rd10, %r1, 4;
	add.s64 	%rd11, %rd9, %rd10;
	ld.global.nc.f32 	%f15, [%rd11];
	mul.f32 	%f25, %f15, %f25;

BB0_7:
	setp.eq.s64	%p5, %rd5, 0;
	@%p5 bra 	BB0_9;

	cvta.to.global.u64 	%rd12, %rd5;
	mul.wide.s32 	%rd13, %r1, 4;
	add.s64 	%rd14, %rd12, %rd13;
	ld.global.nc.f32 	%f16, [%rd14];
	mul.f32 	%f26, %f16, %f26;

BB0_9:
	cvta.to.global.u64 	%rd15, %rd1;
	cvta.to.global.u64 	%rd16, %rd2;
	mul.wide.s32 	%rd17, %r1, 4;
	add.s64 	%rd18, %rd16, %rd17;
	mul.f32 	%f17, %f26, %f9;
	mul.f32 	%f18, %f25, %f17;
	mul.f32 	%f19, %f24, %f18;
	sqrt.rn.f32 	%f20, %f19;
	ld.global.nc.f32 	%f21, [%rd18];
	mul.f32 	%f22, %f21, %f20;
	add.s64 	%rd19, %rd15, %rd17;
	st.global.f32 	[%rd19], %f22;

BB0_10:
	ret;
}


`
	settemperature2_ptx_75 = `
.version 6.5
.target sm_75
.address_size 64

	// .globl	settemperature2

.visible .entry settemperature2(
	.param .u64 settemperature2_param_0,
	.param .u64 settemperature2_param_1,
	.param .f32 settemperature2_param_2,
	.param .u64 settemperature2_param_3,
	.param .f32 settemperature2_param_4,
	.param .u64 settemperature2_param_5,
	.param .f32 settemperature2_param_6,
	.param .u64 settemperature2_param_7,
	.param .f32 settemperature2_param_8,
	.param .u32 settemperature2_param_9
)
{
	.reg .pred 	%p<6>;
	.reg .f32 	%f<27>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<20>;


	ld.param.u64 	%rd1, [settemperature2_param_0];
	ld.param.u64 	%rd2, [settemperature2_param_1];
	ld.param.f32 	%f9, [settemperature2_param_2];
	ld.param.u64 	%rd3, [settemperature2_param_3];
	ld.param.f32 	%f23, [settemperature2_param_4];
	ld.param.u64 	%rd4, [settemperature2_param_5];
	ld.param.f32 	%f25, [settemperature2_param_6];
	ld.param.u64 	%rd5, [settemperature2_param_7];
	ld.param.f32 	%f26, [settemperature2_param_8];
	ld.param.u32 	%r2, [settemperature2_param_9];
	mov.u32 	%r3, %nctaid.x;
	mov.u32 	%r4, %ctaid.y;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r3, %r4, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_10;

	setp.eq.s64	%p2, %rd3, 0;
	@%p2 bra 	BB0_3;

	cvta.to.global.u64 	%rd6, %rd3;
	mul.wide.s32 	%rd7, %r1, 4;
	add.s64 	%rd8, %rd6, %rd7;
	ld.global.nc.f32 	%f13, [%rd8];
	mul.f32 	%f23, %f13, %f23;

BB0_3:
	setp.eq.f32	%p3, %f23, 0f00000000;
	mov.f32 	%f24, 0f00000000;
	@%p3 bra 	BB0_5;

	rcp.rn.f32 	%f24, %f23;

BB0_5:
	setp.eq.s64	%p4, %rd4, 0;
	@%p4 bra 	BB0_7;

	cvta.to.global.u64 	%rd9, %rd4;
	mul.wide.s32 	%rd10, %r1, 4;
	add.s64 	%rd11, %rd9, %rd10;
	ld.global.nc.f32 	%f15, [%rd11];
	mul.f32 	%f25, %f15, %f25;

BB0_7:
	setp.eq.s64	%p5, %rd5, 0;
	@%p5 bra 	BB0_9;

	cvta.to.global.u64 	%rd12, %rd5;
	mul.wide.s32 	%rd13, %r1, 4;
	add.s64 	%rd14, %rd12, %rd13;
	ld.global.nc.f32 	%f16, [%rd14];
	mul.f32 	%f26, %f16, %f26;

BB0_9:
	cvta.to.global.u64 	%rd15, %rd1;
	cvta.to.global.u64 	%rd16, %rd2;
	mul.wide.s32 	%rd17, %r1, 4;
	add.s64 	%rd18, %rd16, %rd17;
	mul.f32 	%f17, %f26, %f9;
	mul.f32 	%f18, %f25, %f17;
	mul.f32 	%f19, %f24, %f18;
	sqrt.rn.f32 	%f20, %f19;
	ld.global.nc.f32 	%f21, [%rd18];
	mul.f32 	%f22, %f21, %f20;
	add.s64 	%rd19, %rd15, %rd17;
	st.global.f32 	[%rd19], %f22;

BB0_10:
	ret;
}


`
)
