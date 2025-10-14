<?php 
require_once "data/Person.php";

use Data\Human\Person;
// constant variabel yg tidak bisa diubah
// disarankan pake uppercase
// constant buat php >= 7.4 pake const
const APP_VERSION = "1.0.0";

echo APP_VERSION . PHP_EOL;
// cara akses const di class
// akses const di class tidak perlu object karna const melekat di class bukan object
echo Person::NAMA . PHP_EOL;

?>