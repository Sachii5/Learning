<?php
require_once "data/Person.php";

    // cara manipulasi properties
$person = new Person("Alber", null);
$person->name = "Alber";
$person->addres = "Kuningan";
$person->country = "Indonesia";

echo $person->name . PHP_EOL;
echo $person->addres . PHP_EOL;
echo $person->country . PHP_EOL;

?>