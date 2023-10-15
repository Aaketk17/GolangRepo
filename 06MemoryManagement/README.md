* Memory allocation and de-allocation happens automatically in Golang
* Usually using new() and make() methods to allocate memory
* new() -> allocate memory but not INIT, will get memory address, zerored storage
* make() -> alllocate memory and INIT, will get memory address, non-sero storage
* Garbage collection happens automatically
