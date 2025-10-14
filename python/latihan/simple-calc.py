print("\t    Project #1")
print("\tKalkulator sederhana")
p = True
while p == True:
    print("-------------------------------------")
    print("\t   Pilih operator")
    print("1. Penjumlahan\n2. Pengurangan\n3. Pembagian\n4. Perkalian\n5. Pangkat")
    operator = int(input("Pilih operator : "))
    a = int(input("Masukan angka : "))
    b = int(input("Masukan angka : "))
    if operator == 1:  
        hasil = a + b
        print(a, " + ", b," = ", hasil)
    elif operator == 2:
        hasil = a - b
        print(a, " - ", b," = ", hasil)
    elif operator == 3:
        hasil = a / b
        print(a, " / ", b," = ", hasil)
    elif operator == 4:
        hasil = a * b
        print(a, " x ", b," = ", hasil)
    elif operator == 5:
        hasil = a ** b
        print(a, "^", b," = ", hasil)
    else:
        print("Opsi tidak tersedia")
    pilih = input("Ingin menghitung lagi?(y/t) = ")
    if pilih == "y":
        p = True
    elif pilih == "t":
        p = False
        print("Terima kasih!")
    else:
        print("Pilihan tidak tersedia")
        p = False
        print("Terima kasih!")
