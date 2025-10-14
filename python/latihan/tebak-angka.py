import random

print("\t\tProject #2")
print("\t\tTebak angka")
print("-------------------------------------------\n")
print("Pilih angka dari 1 - 100")
angka = random.randint(1, 100)
tebakan = False
while tebakan == False:
    tebak = int(input("Angka pilihan = "))
    if tebak != angka:
        print("Tebakan salah, coba lagi!!")
        if angka < tebak:
            print("Angka terlalu besar\n")
        else:
            print("Angka terlalu kecil\n")
    else:
        print("Tebakan anda benar!!")
        tebakan = True