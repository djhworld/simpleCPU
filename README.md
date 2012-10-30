This is a just a small *realisation* of the computer described in *Computer Science: An overview (Sixth Edition)* by by J.G. Brookshear.

It is very buggy and fragile but it will run the example programs/exercises described in the book, in the form of ROMS assembled by the small assembler included in this package

Notes
--------------------
I used this as an educational exercise to give me a refresher on how *basic* computers work. This is by no means a correct implementation, just my initial interpretation of what is described in the book.

Usage
--------------------
	./machine roms/add.rom

* The ```-p``` flag will pause the machine after each CPU cycle
* The ```-d``` flag will dump the state of the machine after each CPU cycle


Specification
--------------------
* 256 bytes of memory
* 16 general purpose registers (each one byte in size)
* 1 instruction register (two bytes in size)
* 1 program counter (one byte in size)
* Memory addressing is simple. 
* CPU has no clock to regulate it 
* No floating point implementation
* No support for overflow/carry flags 
* Basic instruction set 

<table style="font-size:0.8em; table-layout:fixed;">
	<tr>
		<th>Op-code</th>
		<th>Operand</th>
		<th>Description</th>
	</tr>
	<tr>
		<td>1</td>
		<td>R X Y</td>
		<td>LOAD the register R with the data found in memory cell with address XY</td>
	</tr>
	<tr>
		<td>2</td>
		<td>R X Y</td>
		<td>LOAD the register R with the value XY</td>
	</tr>
	<tr>
		<td>3</td>
		<td>R X Y</td>
		<td>STORE the value in register R in the memory cell with address XY</td>
	</tr>
	<tr>
		<td>4</td>
		<td>0 R S</td>
		<td>MOVE the value in register R to the register S</td>
	</tr>
	<tr>
		<td>5</td>
		<td>R S T</td>
		<td>ADD the values in registers S and T and store the result in register R</td>
	</tr>
	<tr>
		<td>7</td>
		<td>R S T</td>
		<td>OR the values in registers S and T in store the result in register R</td>
	</tr>
	<tr>
		<td>8</td>
		<td>R S T</td>
		<td>AND the values in registers S and T in store the result in register R</td>
	</tr>
	<tr>
		<td>9</td>
		<td>R S T</td>
		<td>XOR the values in registers S and T in store the result in register R</td>
	</tr>
	<tr>
		<td>B</td>
		<td>R X Y</td>
		<td>JUMP to the instruction located in the memory cell XY if the value in register R is equal to the value in register 0</td>
	</tr>
	<tr>
		<td>C</td>
		<td>0 0 0</td>
		<td>HALT execution</td>
	</tr>
	<caption><strong>Table:</strong> The small instruction set for our machine language</caption>
</table>

License
--------------------
Do what you want
