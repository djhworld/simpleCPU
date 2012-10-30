Principles
=====================
To write an assembler for our CPU we are going to keep things as simple as possible

* All programs will begin at 0x00
* Data should be stored in range C8 - FF 
* Assembler will do nothing more than do the 1 - 1 mapping 
* Instructions will be in the format ```<INSTR> <ABC>```

Example

	LOADV 104
	ADD 501
	STORE 5D3
Assembler is still VERY buggy
