<?php  
// cara membuat class
namespace Data\Human;

class Person{
    // const di class, cara akses di constant.php
    const NAMA = "Alber Galih Anthoni";
    
    // cara buat properties    
    var string $name; // jika ingin menambahkan tipe data
    var $addres; // jika tanpa tipe data
    var $country = "Indonesia"; // kalo mau dibikin default value
    var ?string $lastName = null ; // nullable properties, gunakan ? sebelum type declaration

    public function __construct(string $name, ?string $addres)
    {
        $this->name = $name;
        $this->addres = $addres;
    }

    
    // contoh function
    function sayHello(?string $name){
    if(is_null($name)){
        // variabel $this buat merepresentasikan object saat ini tapi tidak bisa mengakses class
        // penjelasan lanjutan di file function.php
        echo "Hi, my name is $this->name" . PHP_EOL;
    }else{
        echo "Hi $name, my name is $this->name" . PHP_EOL;
    }
    }

    function info(){
        // untuk akses const di class itu sendiri pake self untuk mempermudah
        echo self::NAMA . PHP_EOL;
    }
}

?>