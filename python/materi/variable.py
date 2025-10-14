# di python hanya ada variable integer, float, string, dan boolean
data_integer = 10
data_float = 10.5
data_string = "halo"
data_bool = True
print(data_integer, data_float, data_string, data_bool)
# lalu ada tipe data  khusus yaitu bilangan complex untuk bilangan imajiner
data_complex = complex(2, 3)
print(data_complex)
# lalu bisa import tipe data dari bahasa c
from ctypes import c_double
data_c_double = 10.234343
print(data_c_double)