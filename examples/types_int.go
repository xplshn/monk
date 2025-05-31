//
// Test integer types: initial values, max values, and operations
//

// Initial values of int types
let i8  = int8(0);
let i16 = int16(0);
let i32 = int32(0);
let i64 = int64(0);
let i   = int(0);

puts("Initial values:\n");
puts("int8  = ", i8, "\n");
puts("int16 = ", i16, "\n");
puts("int32 = ", i32, "\n");
puts("int64 = ", i64, "\n");
puts("int   = ", i, "\n");

// Max values of int types
let max_i8  = int8(127);
let max_i16 = int16(32767);
let max_i32 = int32(2147483647);
let max_i64 = int64(9223372036854775807);
let max_i   = int(2147483647);

puts("\nMax values:\n");
puts("int8  max = ", max_i8, "\n");
puts("int16 max = ", max_i16, "\n");
puts("int32 max = ", max_i32, "\n");
puts("int64 max = ", max_i64, "\n");
puts("int   max = ", max_i, "\n");

// Basic arithmetic tests
puts("\nBasic arithmetic:\n");

puts("5 + 3 = ", int8(i) + int8(3), "\n");
puts("10 - 7 = ", int16(10) - int16(7), "\n");
puts("4 * 6 = ", int32(4) * int32(6), "\n");
puts("20 / 5 = ", int64(20) / int64(5), "\n");

// Test type conversions
puts("\nType conversions:\n");

puts("int8(i).type():    ", int8(i).type(), "\n");
puts("int16(i8).type():  ", int16(i8).type(), "\n");
puts("int32(i16).type(): ", int32(i16).type(), "\n");
puts("int64(i32).type(): ", int64(i32).type(), "\n");
puts("int(i64).type():   ", int(i64).type(), "\n");

// Type checks
puts("\nType checks:\n");

puts("i:   ", i.type(), "\n");
puts("i8:  ", i8.type(), "\n");
puts("i16: ", i16.type(), "\n");
puts("i32: ", i32.type(), "\n");
puts("i64: ", i64.type(), "\n");

// Check integer bounds with operations
puts("\nCheck overflow behavior (wrap-around or error expected):\n");
puts("int8 max + 1 = ", max_i8 + int8(1), "\n");
puts("int16 max + 1 = ", max_i16 + int16(1), "\n");
puts("int32 max + 1 = ", max_i32 + int32(1), "\n");
puts("int64 max + 1 = ", max_i64 + int64(1), "\n");

puts("\nDone testing integer types.\n");
