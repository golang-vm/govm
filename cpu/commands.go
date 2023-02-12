// 
// This file is the command set of the VM
// 

package cpu

type Cmd = uint8

const (
	// Do nothing
	SKIP Cmd = iota
	// return an `*ExitErr`
	EXIT
	// register opers
	// copy value of a register to another one
	MOV
	// set the register value
	SETB
	// load variable value to register
	LOADB
	// store register value to variable
	STOREB
	SETW
	LOADW
	STOREW
	SETI
	LOADI
	STOREI
	SETQ
	LOADQ
	STOREQ
	// TODO: Memory alloc opers
	NEW
	MAKE_SLICE
	MAKE_MAP
	MAKE_CHAN
	LOADP
	STOREP
	// String opers
	SETS
	LOADS
	STORES
	// Stack opers
	PUSHU
	PUSHB
	PUSHW
	PUSHI
	PUSHQ
	PUSHP
	PUSHS
	POP
	POPU
	POPP
	POPS

	B2Q
	W2Q
	I2Q
	I2F
	I2D
	UI2F
	UI2D
	Q2F
	Q2D
	UQ2F
	UQ2D
	F2I
	F2UI
	F2Q
	F2UQ
	F2D
	D2I
	D2UI
	D2Q
	D2UQ
	D2F

	// integer opers
	ADD
	SUB
	MUL
	SMUL
	QUO
	SQUO
	REM
	SREM
	// bit opers
	NOT
	AND
	OR
	XOR
	SHL
	SHR
	SSHR
	SSHRW
	SSHRI
	SSHRQ
	// increse and decrese
	INC
	DEC
	// float opers
	FNEG
	FADD
	FSUB
	FMUL
	FQUO
	DNEG
	DADD
	DSUB
	DMUL
	DQUO
	// Boolean opers and compare
	LNOT
	LAND
	LOR
	EQ
	NE
	GT
	LT
	CMP
	ICMP
	FCMP
	DCMP
	// String opers
	CAT
	CUT
	B2S
	S2B
	// program pointer control
	GOTO
	JMP
	JMP_IF
	JMP_IFN
	// function things
	CALL
	CALLN
	CALLG
	CALLD
	ARG
	RET
	// panic and recover
	PAN
	REC
	// opers of `chan`
	SEND
	RECV
	RECVB
)
